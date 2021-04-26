package talent

import (
	"context"
	"errors"

	"bitbucket.org/funplus/server/excel/auto"
	pbGlobal "bitbucket.org/funplus/server/proto/global"
	"bitbucket.org/funplus/server/services/game/costloot"
	"bitbucket.org/funplus/server/store"
	"bitbucket.org/funplus/server/utils"
	"github.com/spf13/cast"
	"github.com/valyala/bytebufferpool"
)

var (
	ErrInvalidStar              = errors.New("invalid star")
	ErrInvalidTalentId          = errors.New("invalid talent id")
	ErrInvalidTalentType        = errors.New("invalid talent type")
	ErrTalentPrevStepLevelLimit = errors.New("talent prev step level limit")
	ErrTalentPrevLimit          = errors.New("prev talent limit")
	ErrTalentLevelMax           = errors.New("talent level max")

	TalentMaxLevel int32 = 5 // 单个天赋最大等级
)

func makeTalentKey(talentId int32, fields ...string) string {
	b := bytebufferpool.Get()
	defer bytebufferpool.Put(b)

	_, _ = b.WriteString("talent_list.")
	_, _ = b.WriteString(cast.ToString(talentId))

	for _, f := range fields {
		_, _ = b.WriteString(".")
		_, _ = b.WriteString(f)
	}

	return b.String()
}

type TalentOwner interface {
	GetId() int64
	GetTypeId() int32
	GetStoreType() int
}

type Talent struct {
	Id    int32 `bson:"-" json:"-"` // 天赋id
	Level int32 `bson:"-" json:"-"` // 天赋等级
}

// 天赋管理
type TalentBox struct {
	owner           TalentOwner               `bson:"-" json:"-"`
	costLoot        *costloot.CostLootManager `bson:"-" json:"-"`
	talentStepLevel map[int32]int32           `bson:"-" json:"-"`
	TalentType      int32                     `bson:"talent_type" json:"talent_type"`
	TalentList      map[int32]*Talent         `bson:"talent_list" json:"talent_list"`
}

func NewTalentBox(owner TalentOwner, cl *costloot.CostLootManager, tp int32) *TalentBox {
	m := &TalentBox{
		owner:           owner,
		costLoot:        cl,
		talentStepLevel: make(map[int32]int32),
		TalentType:      tp,
		TalentList:      make(map[int32]*Talent),
	}

	return m
}

// 获取天赋层级总等级
func (tb *TalentBox) GetStepTotalLevel(step int32) int32 {
	return tb.talentStepLevel[step]
}

// 激活、升级天赋
func (tb *TalentBox) ChooseTalent(talentId int32) error {
	talentEntry, ok := auto.GetTalentEntry(talentId)
	if !ok {
		return ErrInvalidTalentId
	}

	if talentEntry.Type != tb.TalentType {
		return ErrInvalidTalentType
	}

	if talentEntry.OwnerId != tb.owner.GetTypeId() {
		return ErrInvalidTalentId
	}

	// 前一层天赋等级数限制
	stepLevelLimit := talentEntry.PrevStepLevelLimit
	if stepLevelLimit > 0 && stepLevelLimit < tb.GetStepTotalLevel(talentEntry.Step-1) {
		return ErrTalentPrevStepLevelLimit
	}

	// 前置天赋限制
	if talentEntry.PrevTalentId1 != -1 {
		prevTalent, ok := tb.TalentList[talentEntry.PrevTalentId1]
		if !ok {
			return ErrTalentPrevLimit
		}

		if prevTalent.Level < talentEntry.PrevTalentLevel1 {
			return ErrTalentPrevLimit
		}
	}

	if talentEntry.PrevTalentId2 != -1 {
		prevTalent, ok := tb.TalentList[talentEntry.PrevTalentId2]
		if !ok {
			return ErrTalentPrevLimit
		}

		if prevTalent.Level < talentEntry.PrevTalentLevel2 {
			return ErrTalentPrevLimit
		}
	}

	if tb.costLoot != nil {
		err := tb.costLoot.CanCost(talentEntry.CostId)
		if !utils.ErrCheck(err, "CanCost failed when TalentBox.ChooseTalent", talentId) {
			return err
		}
	}

	curTalent, ok := tb.TalentList[talentId]
	if !ok {
		curTalent = &Talent{
			Id:    talentId,
			Level: 0,
		}
		tb.TalentList[talentId] = curTalent
	}

	// 天赋等级上限
	if curTalent.Level >= TalentMaxLevel {
		return ErrTalentLevelMax
	}

	// 消耗
	if tb.costLoot != nil {
		err := tb.costLoot.DoCost(talentEntry.CostId)
		if !utils.ErrCheck(err, "DoCost failed when TalentBox.ChooseTalent", talentId) {
			return err
		}
	}

	curTalent.Level++
	tb.talentStepLevel[talentEntry.Step]++

	fields := map[string]interface{}{
		makeTalentKey(talentId): curTalent,
	}
	err := store.GetStore().UpdateFields(context.Background(), tb.owner.GetStoreType(), tb.owner.GetId(), fields)
	utils.ErrPrint(err, "UpdateFields failed when TalentBox.ChooseTalent", tb.owner.GetId(), fields)

	return nil
}

func (tb *TalentBox) GenTalentList() []*pbGlobal.Talent {
	pb := make([]*pbGlobal.Talent, 0, len(tb.TalentList))
	for _, talent := range tb.TalentList {
		pb = append(pb, &pbGlobal.Talent{
			TalentId:    talent.Id,
			TalentLevel: talent.Level,
		})
	}

	return pb
}
