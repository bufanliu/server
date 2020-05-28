package game

import (
	"context"

	logger "github.com/sirupsen/logrus"
	pbGame "github.com/yokaiio/yokai_server/proto/game"
	"github.com/yokaiio/yokai_server/transport"
)

func (m *MsgHandler) handleAddItem(ctx context.Context, sock transport.Socket, p *transport.Message) {
	acct := m.g.am.GetAccountBySock(sock)
	if acct == nil {
		logger.WithFields(logger.Fields{
			"account_id":   acct.GetID(),
			"account_name": acct.GetName(),
		}).Warn("add item failed")
		return
	}

	pl := m.g.pm.GetPlayerByAccount(acct)
	if pl == nil {
		return
	}

	msg, ok := p.Body.(*pbGame.C2M_AddItem)
	if !ok {
		logger.Warn("Add Item failed, recv message body error")
		return
	}

	acct.PushWrapHandler(func() {
		pl.ItemManager().AddItemByTypeID(msg.TypeId, 1)
	})

}

func (m *MsgHandler) handleDelItem(ctx context.Context, sock transport.Socket, p *transport.Message) {
	acct := m.g.am.GetAccountBySock(sock)
	if acct == nil {
		logger.WithFields(logger.Fields{
			"account_id":   acct.GetID(),
			"account_name": acct.GetName(),
		}).Warn("delete item failed")
		return
	}

	pl := m.g.pm.GetPlayerByAccount(acct)
	if pl == nil {
		return
	}

	msg, ok := p.Body.(*pbGame.C2M_DelItem)
	if !ok {
		logger.Warn("Delete item failed, recv message body error")
		return
	}

	acct.PushWrapHandler(func() {
		item := pl.ItemManager().GetItem(msg.Id)
		if item == nil {
			logger.Warn("Delete item failed, non-existing item_id:", msg.Id)
			return
		}

		// clear hero's equip id before delete item
		equipObjID := item.GetEquipObj()
		if equipObjID != -1 {
			pl.HeroManager().TakeoffEquip(equipObjID, item.EquipEnchantEntry().EquipPos)
		}

		// delete item
		pl.ItemManager().DeleteItem(msg.Id)
	})
}

func (m *MsgHandler) handleUseItem(ctx context.Context, sock transport.Socket, p *transport.Message) {
	acct := m.g.am.GetAccountBySock(sock)
	if acct == nil {
		logger.WithFields(logger.Fields{
			"account_id":   acct.GetID(),
			"account_name": acct.GetName(),
		}).Warn("use item failed")
		return
	}

	pl := m.g.pm.GetPlayerByAccount(acct)
	if pl == nil {
		return
	}

	msg, ok := p.Body.(*pbGame.C2M_UseItem)
	if !ok {
		logger.Warn("Use Item failed, recv message body error")
		return
	}

	acct.PushWrapHandler(func() {
		pl.ItemManager().UseItem(msg.ItemId)
	})
}

func (m *MsgHandler) handleQueryItems(ctx context.Context, sock transport.Socket, p *transport.Message) {
	acct := m.g.am.GetAccountBySock(sock)
	if acct == nil {
		logger.WithFields(logger.Fields{
			"account_id":   acct.GetID(),
			"account_name": acct.GetName(),
		}).Warn("query items failed")
		return
	}

	pl := m.g.pm.GetPlayerByAccount(acct)
	if pl == nil {
		return
	}

	acct.PushWrapHandler(func() {
		reply := &pbGame.M2C_ItemList{}
		list := pl.ItemManager().GetItemList()
		for _, v := range list {
			i := &pbGame.Item{
				Id:     v.Options().Id,
				TypeId: v.Options().TypeId,
			}
			reply.Items = append(reply.Items, i)
		}
		acct.SendProtoMessage(reply)
	})
}

func (m *MsgHandler) handlePutonEquip(ctx context.Context, sock transport.Socket, p *transport.Message) {
	acct := m.g.am.GetAccountBySock(sock)
	if acct == nil {
		logger.WithFields(logger.Fields{
			"account_id":   acct.GetID(),
			"account_name": acct.GetName(),
		}).Warn("Puton equip failed")
		return
	}

	pl := m.g.pm.GetPlayerByAccount(acct)
	if pl == nil {
		return
	}

	msg, ok := p.Body.(*pbGame.C2M_PutonEquip)
	if !ok {
		logger.Warn("Puton equip failed, recv message body error")
		return
	}

	acct.PushWrapHandler(func() {
		if err := pl.HeroManager().PutonEquip(msg.HeroId, msg.EquipId); err != nil {
			logger.Warn(err)
			return
		}
	})
}

func (m *MsgHandler) handleTakeoffEquip(ctx context.Context, sock transport.Socket, p *transport.Message) {
	acct := m.g.am.GetAccountBySock(sock)
	if acct == nil {
		logger.WithFields(logger.Fields{
			"account_id":   acct.GetID(),
			"account_name": acct.GetName(),
		}).Warn("Takeoff equip failed")
		return
	}

	pl := m.g.pm.GetPlayerByAccount(acct)
	if pl == nil {
		return
	}

	msg, ok := p.Body.(*pbGame.C2M_TakeoffEquip)
	if !ok {
		logger.Warn("Takeoff equip failed, recv message body error")
		return
	}

	acct.PushWrapHandler(func() {
		if err := pl.HeroManager().TakeoffEquip(msg.HeroId, msg.Pos); err != nil {
			logger.Warn(err)
			return
		}
	})
}
