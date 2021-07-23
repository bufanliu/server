package rank

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"sync"
	"time"

	"e.coding.net/mmstudio/blade/server/define"
	"e.coding.net/mmstudio/blade/server/store"
	"e.coding.net/mmstudio/blade/server/utils"
	"e.coding.net/mmstudio/blade/server/utils/cache"
	"github.com/hellodudu/task"
	log "github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

var (
	rankCleanupInterval = 1 * time.Minute // cache cleanup interval
	rankCacheExpire     = 1 * time.Hour   // cache缓存1小时
)

type RankManager struct {
	r              *Rank
	cacheRankDatas *cache.Cache
	rankPool       sync.Pool
	wg             utils.WaitGroupWrapper
}

func NewRankManager(ctx *cli.Context, r *Rank) *RankManager {
	manager := &RankManager{
		r:              r,
		cacheRankDatas: cache.New(rankCacheExpire, rankCleanupInterval),
	}

	// 排行榜池
	manager.rankPool.New = NewRankData

	// 排行缓存删除时处理
	manager.cacheRankDatas.OnEvicted(func(k, v interface{}) {
		v.(*RankData).Stop()
		manager.rankPool.Put(v)
	})

	// 初始化db
	store.GetStore().AddStoreInfo(define.StoreType_LocalRank, "local_rank", "_id")   // _id = "nodeid-rank_type_id"
	store.GetStore().AddStoreInfo(define.StoreType_GlobalRank, "global_rank", "_id") // _id = "global-rank_type_id"
	if err := store.GetStore().MigrateDbTable("local_rank", "rank_list._id"); err != nil {
		log.Fatal().Err(err).Msg("migrate collection local_rank failed")
	}

	if err := store.GetStore().MigrateDbTable("global_rank", "rank_list._id"); err != nil {
		log.Fatal().Err(err).Msg("migrate collection global_rank failed")
	}

	log.Info().Msg("RankManager init ok ...")
	return manager
}

func (m *RankManager) Run(ctx *cli.Context) error {
	<-ctx.Done()
	log.Info().Msg("RankManager context done...")
	return nil
}

func (m *RankManager) Exit(ctx *cli.Context) {
	m.wg.Wait()
	log.Info().Msg("RankManager exit...")
}

func (m *RankManager) KickAllRankData() {
	m.cacheRankDatas.DeleteAll()
}

// 踢掉排行缓存
func (m *RankManager) KickRankData(rankId int32, rankNodeId int32) error {
	if rankId == -1 {
		return nil
	}

	// 踢掉本服rankdata
	if rankNodeId == int32(m.r.ID) {
		rd, ok := m.cacheRankDatas.Get(rankId)
		if !ok {
			return nil
		}

		rd.(*RankData).Stop()
		store.GetStore().Flush()
		return nil

	} else {
		// rank节点不存在的话不用发送rpc
		nodeId := fmt.Sprintf("rank-%d", rankNodeId)
		srvs, err := m.r.mi.srv.Options().Registry.GetService("rank")
		if err != nil {
			return nil
		}

		hit := false
		for _, srv := range srvs {
			for _, node := range srv.Nodes {
				if node.Id == nodeId {
					hit = true
					break
				}
			}
		}

		if !hit {
			return nil
		}

		// 发送rpc踢掉其他服rankdata
		rs, err := m.r.rpcHandler.CallKickRankData(rankId, rankNodeId)
		if !utils.ErrCheck(err, "kick rank data failed", rankId, rankNodeId, rs) {
			return err
		}

		// rpc调用成功
		if rs.GetRankId() == rankId {
			return nil
		}

		return errors.New("kick rank data invalid error")
	}
}

// 获取rank数据
func (m *RankManager) getRankData(rankId int32) (*RankData, error) {
	if rankId == -1 {
		return nil, ErrInvalidRank
	}

	cache, ok := m.cacheRankDatas.Get(rankId)

	if ok {
		rd := cache.(*RankData)
		if rd.IsTaskRunning() {
			return rd, nil
		}

	} else {

		// 缓存没有，从db加载
		cache = m.rankPool.Get()
		rd := cache.(*RankData)
		rd.Init(m.r.ID, m.r.rpcHandler)
		err := rd.Load(rankId)
		if !utils.ErrCheck(err, "RankData Load failed when RankManager.getRankData", rankId) {
			m.rankPool.Put(cache)
			return nil, err
		}

		// 踢掉上一个节点的缓存
		if rd.LastSaveNodeId != -1 && rd.LastSaveNodeId != int32(m.r.ID) {
			err := m.KickRankData(rd.Id, rd.LastSaveNodeId)
			if !utils.ErrCheck(err, "kick RankData failed", rd.Id, rd.LastSaveNodeId, m.r.ID) {
				return nil, err
			}
		}

		m.cacheRankDatas.Set(rankId, cache, rankCacheExpire)
	}

	rd := cache.(*RankData)
	rd.InitTask()
	m.wg.Wrap(func() {
		defer func() {
			if err := recover(); err != nil {
				stack := string(debug.Stack())
				log.Error().Msgf("catch exception:%v, panic recovered with stack:%s", err, stack)
			}

			// 立即删除缓存
			m.cacheRankDatas.Delete(cache.(*RankData).Id)
		}()

		ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
		err := cache.(*RankData).TaskRun(ctx)
		utils.ErrPrint(err, "RankData run failed", cache.(*RankData).Id)
	})

	return rd, nil
}

func (m *RankManager) AddTask(ctx context.Context, rankId int32, fn task.TaskHandler) error {
	rd, err := m.getRankData(rankId)
	if err != nil {
		return err
	}

	return rd.AddTask(ctx, fn, rd)
}

// 查询排行
func (m *RankManager) QueryRankByKey(ctx context.Context, rankId int32, key int64) (*define.RankRaw, error) {
	var raw *define.RankRaw
	err := m.AddTask(
		ctx,
		rankId,
		func(c context.Context, p ...interface{}) error {
			var e error
			rankData := p[0].(*RankData)
			raw, e = rankData.GetRankByKey(c, key)
			return e
		},
	)

	if !utils.ErrCheck(err, "AddTask failed when RankManager.QueryRankByKey", rankId, key) {
		return nil, err
	}

	return raw, err
}

func (m *RankManager) QueryRankByScore(ctx context.Context, rankId int32, start, end int64) ([]*define.RankRaw, error) {
	var raws []*define.RankRaw
	err := m.AddTask(
		ctx,
		rankId,
		func(c context.Context, p ...interface{}) error {
			var e error
			rankData := p[0].(*RankData)
			raws, e = rankData.GetRankByIndex(c, start, end)
			return e
		},
	)

	if !utils.ErrCheck(err, "AddTask failed when RankManager.QueryRankByScore", rankId, start, end) {
		return nil, err
	}

	return raws, err
}

// 设置排行积分
func (m *RankManager) SetRankScore(ctx context.Context, rankId int32, rankRaw *define.RankRaw) error {
	return m.AddTask(ctx, rankId, func(c context.Context, p ...interface{}) error {
		rankData := p[0].(*RankData)
		return rankData.SetScore(ctx, rankRaw)
	})
}
