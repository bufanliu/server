package player

import (
	"fmt"
	"reflect"
	"sync"

	logger "github.com/sirupsen/logrus"
	"github.com/yokaiio/yokai_server/define"
	"github.com/yokaiio/yokai_server/entries"
	"github.com/yokaiio/yokai_server/game/db"
	"github.com/yokaiio/yokai_server/game/hero"
	pbCombat "github.com/yokaiio/yokai_server/proto/combat"
	pbGame "github.com/yokaiio/yokai_server/proto/game"
	"github.com/yokaiio/yokai_server/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type HeroManager struct {
	owner   *Player
	mapHero map[int64]hero.Hero

	ds   *db.Datastore
	coll *mongo.Collection
	sync.RWMutex
}

func NewHeroManager(owner *Player, ds *db.Datastore) *HeroManager {
	m := &HeroManager{
		owner:   owner,
		ds:      ds,
		mapHero: make(map[int64]hero.Hero, 0),
	}

	if ds != nil {
		m.coll = ds.Database().Collection(m.TableName())
	}

	return m
}

func (m *HeroManager) TableName() string {
	return "hero"
}

func (m *HeroManager) save(h hero.Hero) {
	filter := bson.D{{"_id", h.GetID()}}
	update := bson.D{{"$set", h}}
	op := options.Update().SetUpsert(true)
	id := h.GetID()

	if m.ds == nil {
		return
	}

	m.ds.Wrap(func() {
		if _, err := m.coll.UpdateOne(nil, filter, update, op); err != nil {
			logger.WithFields(logger.Fields{
				"id":    id,
				"error": err,
			}).Warning("hero save failed")
		}
	})
}

func (m *HeroManager) saveField(h hero.Hero, up *bson.D) {
	filter := bson.D{{"_id", h.GetID()}}
	update := up
	id := h.GetID()

	if m.ds == nil {
		return
	}

	m.ds.Wrap(func() {
		if _, err := m.coll.UpdateOne(nil, filter, *update); err != nil {
			logger.WithFields(logger.Fields{
				"id":    id,
				"error": err,
			}).Warning("hero save level failed")
		}
	})
}

func (m *HeroManager) delete(h hero.Hero, filter *bson.D) {
	id := h.GetID()
	f := filter

	if m.ds == nil {
		return
	}

	m.ds.Wrap(func() {
		if _, err := m.coll.DeleteOne(nil, *f); err != nil {
			logger.WithFields(logger.Fields{
				"id":    id,
				"error": err,
			}).Warning("hero delete level failed")
		}
	})
}

func (m *HeroManager) createEntryHero(entry *define.HeroEntry) hero.Hero {
	if entry == nil {
		logger.Error("newEntryHero with nil HeroEntry")
		return nil
	}

	id, err := utils.NextID(define.SnowFlake_Hero)
	if err != nil {
		logger.Error(err)
		return nil
	}

	h := hero.NewHero(
		hero.Id(id),
		hero.OwnerId(m.owner.GetID()),
		hero.Entry(entry),
		hero.TypeId(entry.ID),
	)

	h.GetAttManager().SetBaseAttId(entry.AttID)
	m.mapHero[h.GetID()] = h
	h.GetAttManager().CalcAtt()

	return h
}

func (m *HeroManager) createDBHero(h hero.Hero) hero.Hero {
	entry := entries.GetHeroEntry(h.Options().TypeId)

	newHero := hero.NewHero(
		hero.Id(h.Options().Id),
		hero.OwnerId(h.Options().OwnerId),
		hero.OwnerType(h.Options().OwnerType),
		hero.Level(h.Options().Level),
		hero.TypeId(h.Options().TypeId),
		hero.Entry(entry),
	)

	newHero.GetAttManager().SetBaseAttId(entry.AttID)
	m.mapHero[newHero.GetID()] = newHero
	newHero.CalcAtt()

	return newHero
}

// interface of cost_loot
func (m *HeroManager) GetCostLootType() int32 {
	return define.CostLoot_Hero
}

func (m *HeroManager) CanCost(typeMisc int32, num int32) error {
	if num <= 0 {
		return fmt.Errorf("hero manager check hero<%d> cost failed, wrong number<%d>", typeMisc, num)
	}

	var fixNum int32 = 0
	for _, v := range m.mapHero {
		if v.Options().TypeId == typeMisc {
			eb := v.GetEquipBar()
			hasEquip := false
			var n int32
			for n = 0; n < define.Hero_MaxEquip; n++ {
				if eb.GetEquipByPos(n) != nil {
					hasEquip = true
					break
				}
			}

			if !hasEquip {
				fixNum++
			}
		}
	}

	if fixNum >= num {
		return nil
	}

	return fmt.Errorf("not enough hero<%d>, num<%d>", typeMisc, num)
}

func (m *HeroManager) DoCost(typeMisc int32, num int32) error {
	if num <= 0 {
		return fmt.Errorf("hero manager cost hero<%d> failed, wrong number<%d>", typeMisc, num)
	}

	var costNum int32 = 0
	for _, v := range m.mapHero {
		if v.Options().TypeId == typeMisc {
			eb := v.GetEquipBar()
			hasEquip := false
			var n int32
			for n = 0; n < define.Hero_MaxEquip; n++ {
				if eb.GetEquipByPos(n) != nil {
					hasEquip = true
					break
				}
			}

			if !hasEquip {
				m.DelHero(v.GetID())
				costNum++
			}
		}
	}

	if costNum < num {
		logger.WithFields(logger.Fields{
			"cost_type_misc":  typeMisc,
			"cost_num":        num,
			"actual_cost_num": costNum,
		}).Warn("hero manager cost num error")
		return nil
	}

	return nil
}

func (m *HeroManager) CanGain(typeMisc int32, num int32) error {
	if num <= 0 {
		return fmt.Errorf("hero manager check hero<%d> gain failed, wrong number<%d>", typeMisc, num)
	}

	// todo max hero num
	return nil
}

func (m *HeroManager) GainLoot(typeMisc int32, num int32) error {
	if num <= 0 {
		return fmt.Errorf("hero manager gain hero<%d> failed, wrong number<%d>", typeMisc, num)
	}

	var n int32 = 0
	for ; n < num; n++ {
		h := m.AddHeroByTypeID(typeMisc)
		if h == nil {
			return fmt.Errorf("hero manager gain hero<%d> failed, cannot add new hero<%d>", typeMisc, num)
		}
	}

	return nil
}

func (m *HeroManager) LoadFromDB() {
	l := hero.LoadAll(m.ds, m.owner.GetID(), m.TableName())
	sliceHero := make([]hero.Hero, 0)

	listHero := reflect.ValueOf(l)
	if listHero.Kind() != reflect.Slice {
		logger.Error("load hero returns non-slice type")
		return
	}

	for n := 0; n < listHero.Len(); n++ {
		p := listHero.Index(n)
		sliceHero = append(sliceHero, p.Interface().(hero.Hero))
	}

	for _, v := range sliceHero {
		m.createDBHero(v)
	}
}

func (m *HeroManager) GetHero(id int64) hero.Hero {
	return m.mapHero[id]
}

func (m *HeroManager) GetHeroNums() int {
	return len(m.mapHero)
}

func (m *HeroManager) GetHeroList() []hero.Hero {
	list := make([]hero.Hero, 0)

	m.RLock()
	for _, v := range m.mapHero {
		list = append(list, v)
	}
	m.RUnlock()

	return list
}

func (m *HeroManager) AddHeroByTypeID(typeID int32) hero.Hero {
	heroEntry := entries.GetHeroEntry(typeID)
	h := m.createEntryHero(heroEntry)
	if h == nil {
		return nil
	}

	m.save(h)
	return h
}

func (m *HeroManager) DelHero(id int64) {
	h, ok := m.mapHero[id]
	if !ok {
		return
	}

	eb := h.GetEquipBar()
	var n int32
	for n = 0; n < define.Hero_MaxEquip; n++ {
		eb.TakeoffEquip(n)
	}
	h.BeforeDelete()

	delete(m.mapHero, id)
	m.delete(h, &bson.D{{"_id", id}})
	hero.ReleasePoolHero(h)
}

func (m *HeroManager) HeroSetLevel(level int32) {
	for _, v := range m.mapHero {
		v.Options().Level = level

		update := &bson.D{{"$set",
			bson.D{
				{"level", v.GetLevel()},
			},
		}}

		m.saveField(v, update)
	}
}

func (m *HeroManager) PutonEquip(heroID int64, equipID int64) error {

	equip := m.owner.ItemManager().GetItem(equipID)
	if equip == nil {
		return fmt.Errorf("cannot find equip<%d> while PutonEquip", equipID)
	}

	if objID := equip.GetEquipObj(); objID != -1 {
		return fmt.Errorf("equip has put on another hero<%d>", objID)
	}

	if equip.EquipEnchantEntry() == nil {
		return fmt.Errorf("cannot find equip_enchant_entry<%d> while PutonEquip", equipID)
	}

	h, ok := m.mapHero[heroID]
	if !ok {
		return fmt.Errorf("invalid heroid")
	}

	equipBar := h.GetEquipBar()
	pos := equip.EquipEnchantEntry().EquipPos

	// takeoff previous equip
	if pe := equipBar.GetEquipByPos(pos); pe != nil {
		if err := m.TakeoffEquip(heroID, pos); err != nil {
			return err
		}
	}

	// puton this equip
	if err := equipBar.PutonEquip(equip); err != nil {
		return err
	}

	m.owner.ItemManager().Save(equip.Options().Id)
	m.owner.ItemManager().SendItemUpdate(equip)
	m.SendHeroUpdate(h)

	// att
	equip.GetAttManager().CalcAtt()
	h.GetAttManager().ModAttManager(equip.GetAttManager())
	h.GetAttManager().CalcAtt()
	m.SendHeroAtt(h)

	return nil
}

func (m *HeroManager) TakeoffEquip(heroID int64, pos int32) error {
	if pos < 0 || pos >= define.Hero_MaxEquip {
		return fmt.Errorf("invalid pos")
	}

	h, ok := m.mapHero[heroID]
	if !ok {
		return fmt.Errorf("invalid heroid")
	}

	equipBar := h.GetEquipBar()
	equip := equipBar.GetEquipByPos(pos)
	if equip == nil {
		return fmt.Errorf("cannot find hero<%d> equip by pos<%d> while TakeoffEquip", heroID, pos)
	}

	if objID := equip.GetEquipObj(); objID == -1 {
		return fmt.Errorf("equip<%d> didn't put on this hero<%d> ", equip.Options().Id, heroID)
	}

	// unequip
	if err := equipBar.TakeoffEquip(pos); err != nil {
		return err
	}

	m.owner.ItemManager().Save(equip.Options().Id)
	m.owner.ItemManager().SendItemUpdate(equip)
	m.SendHeroUpdate(h)

	// att
	h.GetAttManager().CalcAtt()
	m.SendHeroAtt(h)

	return nil
}

func (m *HeroManager) PutonRune(heroID int64, runeID int64) error {

	r := m.owner.RuneManager().GetRune(runeID)
	if r == nil {
		return fmt.Errorf("cannot find rune<%d> while PutonRune", runeID)
	}

	if objID := r.GetEquipObj(); objID != -1 {
		return fmt.Errorf("rune has put on another obj<%d>", objID)
	}

	pos := r.Entry().Pos
	if pos < define.Rune_PositionBegin || pos >= define.Rune_PositionEnd {
		return fmt.Errorf("invalid pos<%d>", pos)
	}

	h, ok := m.mapHero[heroID]
	if !ok {
		return fmt.Errorf("invalid heroid<%d>", heroID)
	}

	runeBox := h.GetRuneBox()

	// takeoff previous rune
	if pr := runeBox.GetRuneByPos(pos); pr != nil {
		if err := m.TakeoffRune(heroID, pos); err != nil {
			return err
		}
	}

	// equip new rune
	if err := runeBox.PutonRune(r); err != nil {
		return err
	}

	m.owner.RuneManager().Save(runeID)
	m.owner.RuneManager().SendRuneUpdate(r)
	m.SendHeroUpdate(h)

	// att
	r.GetAttManager().CalcAtt()
	h.GetAttManager().ModAttManager(r.GetAttManager())
	h.GetAttManager().CalcAtt()
	m.SendHeroAtt(h)

	return nil
}

func (m *HeroManager) TakeoffRune(heroID int64, pos int32) error {
	if pos < 0 || pos >= define.Rune_PositionEnd {
		return fmt.Errorf("invalid pos<%d>", pos)
	}

	h, ok := m.mapHero[heroID]
	if !ok {
		return fmt.Errorf("invalid heroid<%d>", heroID)
	}

	r := h.GetRuneBox().GetRuneByPos(pos)
	if r == nil {
		return fmt.Errorf("cannot find rune from hero<%d>'s runebox pos<%d> while TakeoffRune", heroID, pos)
	}

	// unequip
	if err := h.GetRuneBox().TakeoffRune(pos); err != nil {
		return err
	}

	m.owner.RuneManager().Save(r.GetID())
	m.owner.RuneManager().SendRuneUpdate(r)
	m.SendHeroUpdate(h)

	// att
	h.GetAttManager().CalcAtt()
	m.SendHeroAtt(h)

	return nil
}

func (m *HeroManager) GenerateCombatUnitInfo() []*pbCombat.UnitInfo {
	retList := make([]*pbCombat.UnitInfo, 0)

	list := m.GetHeroList()
	for _, hero := range list {
		unitInfo := &pbCombat.UnitInfo{
			UnitTypeId: hero.Options().TypeId,
		}

		for n := define.Att_Begin; n < define.Att_End; n++ {
			unitInfo.UnitAttList = append(unitInfo.UnitAttList, &pbGame.Att{
				AttType:  int32(n),
				AttValue: hero.GetAttManager().GetAttValue(int32(n)),
			})
		}

		retList = append(retList, unitInfo)
	}

	return retList
}

func (m *HeroManager) SendHeroUpdate(h hero.Hero) {
	// send equips update
	reply := &pbGame.M2C_HeroInfo{
		Info: &pbGame.Hero{
			Id:     h.GetID(),
			TypeId: h.Options().TypeId,
			Exp:    h.Options().Exp,
			Level:  h.Options().Level,
		},
	}

	// equip list
	eb := h.GetEquipBar()
	var n int32
	for n = 0; n < define.Hero_MaxEquip; n++ {
		var equipId int64 = -1
		if i := eb.GetEquipByPos(n); i != nil {
			equipId = i.Options().Id
		}

		reply.Info.EquipList = append(reply.Info.EquipList, equipId)
	}

	// rune list
	var pos int32
	for pos = 0; pos < define.Rune_PositionEnd; pos++ {
		var runeId int64 = -1
		if r := h.GetRuneBox().GetRuneByPos(pos); r != nil {
			runeId = r.GetID()
		}

		reply.Info.RuneList = append(reply.Info.RuneList, runeId)
	}

	m.owner.SendProtoMessage(reply)
}

func (m *HeroManager) SendHeroAtt(h hero.Hero) {
	attManager := h.GetAttManager()
	reply := &pbGame.M2C_HeroAttUpdate{
		HeroId: h.GetID(),
	}

	for k := int32(0); k < define.Att_End; k++ {
		att := &pbGame.Att{
			AttType:  k,
			AttValue: attManager.GetAttValue(k),
		}
		reply.AttList = append(reply.AttList, att)
	}

	m.owner.SendProtoMessage(reply)
}
