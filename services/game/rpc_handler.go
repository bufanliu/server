package game

import (
	"context"
	"time"

	"github.com/micro/go-micro/v2/client"
	log "github.com/rs/zerolog/log"
	"github.com/east-eden/server/define"
	pbCombat "github.com/east-eden/server/proto/combat"
	pbGame "github.com/east-eden/server/proto/game"
	pbGate "github.com/east-eden/server/proto/gate"
	"github.com/east-eden/server/services/game/player"
	"github.com/east-eden/server/utils"
)

type RpcHandler struct {
	g         *Game
	gateSrv   pbGate.GateService
	gameSrv   pbGame.GameService
	combatSrv pbCombat.CombatService
}

func NewRpcHandler(g *Game) *RpcHandler {
	h := &RpcHandler{
		g: g,
		gateSrv: pbGate.NewGateService(
			"gate",
			g.mi.srv.Client(),
		),

		gameSrv: pbGame.NewGameService(
			"game",
			g.mi.srv.Client(),
		),

		combatSrv: pbCombat.NewCombatService(
			"combat",
			g.mi.srv.Client(),
		),
	}

	err := pbGame.RegisterGameServiceHandler(g.mi.srv.Server(), h)
	if err != nil {
		log.Fatal().Err(err).Msg("RegisterGameServiceHandler failed")
	}

	return h
}

/////////////////////////////////////////////
// rpc call
/////////////////////////////////////////////
func (h *RpcHandler) CallGetGateStatus(ctx context.Context) (*pbGate.GetGateStatusReply, error) {
	req := &pbGate.GateEmptyMessage{}
	return h.gateSrv.GetGateStatus(ctx, req)
}

func (h *RpcHandler) CallGetRemoteLitePlayer(playerID int64) (*pbGame.GetRemoteLitePlayerReply, error) {
	req := &pbGame.GetRemoteLitePlayerRequest{Id: playerID}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return h.gameSrv.GetRemoteLitePlayer(ctx, req, client.WithSelectOption(utils.SectionIDRandSelector(playerID)))
}

func (h *RpcHandler) CallUpdateUserInfo(c *player.Account) (*pbGate.GateEmptyMessage, error) {
	var playerID int64 = -1
	if len(c.PlayerIDs) > 0 {
		playerID = c.PlayerIDs[0]
	}

	info := &pbGate.UserInfo{
		UserId:      c.UserId,
		AccountId:   c.ID,
		GameId:      int32(c.GameId),
		PlayerId:    playerID,
		PlayerName:  c.Name,
		PlayerLevel: c.Level,
	}

	req := &pbGate.UpdateUserInfoRequest{Info: info}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return h.gateSrv.UpdateUserInfo(ctx, req)
}

func (h *RpcHandler) CallStartStageCombat(p *player.Player) (*pbCombat.StartStageCombatReply, error) {
	sceneId, err := utils.NextID(define.SnowFlake_Scene)
	if err != nil {
		return nil, err
	}

	req := &pbCombat.StartStageCombatReq{
		SceneId:        sceneId,
		SceneType:      define.Scene_TypeStage,
		AttackId:       p.GetID(),
		AttackUnitList: p.HeroManager().GenerateCombatUnitInfo(),
		DefenceId:      -1,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return h.combatSrv.StartStageCombat(ctx, req)
}

func (h *RpcHandler) CallSyncPlayerInfo(userId int64, info *player.LitePlayer) (*pbGate.SyncPlayerInfoReply, error) {
	req := &pbGate.SyncPlayerInfoRequest{
		UserId: userId,
		Info: &pbGame.PlayerInfo{
			LiteInfo: &pbGame.LitePlayer{
				Id:        info.ID,
				AccountId: info.AccountID,
				Name:      info.Name,
				Exp:       info.Exp,
				Level:     info.Level,
			},
			HeroNums: 1,
			ItemNums: 1,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	tm := time.Now()
	defer func() {
		d := time.Since(tm)
		if d > time.Second*2 {
			log.Warn().
				Dur("latency", d).
				Msg("rpc CallSyncPlayerInfo latency")
		}
	}()
	return h.gateSrv.SyncPlayerInfo(ctx, req)
}

/////////////////////////////////////////////
// rpc receive
/////////////////////////////////////////////
func (h *RpcHandler) GetRemoteLitePlayer(ctx context.Context, req *pbGame.GetRemoteLitePlayerRequest, rsp *pbGame.GetRemoteLitePlayerReply) error {
	lp, err := h.g.am.GetLitePlayer(req.Id)
	if err != nil {
		return err
	}

	rsp.Info = &pbGame.LitePlayer{
		Id:        lp.GetID(),
		AccountId: lp.GetAccountID(),
		Name:      lp.GetName(),
		Exp:       lp.GetExp(),
		Level:     lp.GetLevel(),
	}

	return nil
}

func (h *RpcHandler) UpdatePlayerExp(ctx context.Context, req *pbGame.UpdatePlayerExpRequest, rsp *pbGame.UpdatePlayerExpReply) error {
	//lp, err := h.g.pm.getLitePlayer(req.Id)
	//if err != nil {
	//return err
	//}

	//rsp.Id = lp.GetID()
	//rsp.Exp = lp.GetExp()
	return nil
}