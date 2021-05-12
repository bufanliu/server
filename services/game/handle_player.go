package game

import (
	"context"
	"errors"
	"fmt"

	pbGlobal "bitbucket.org/funplus/server/proto/global"
	"bitbucket.org/funplus/server/services/game/player"
)

func (m *MsgRegister) handleCreatePlayer(ctx context.Context, p ...interface{}) error {
	acct := p[0].(*player.Account)
	msg, ok := p[1].(*pbGlobal.C2S_CreatePlayer)
	if !ok {
		return errors.New("handleCreatePlayer failed: recv message body error")
	}

	reply := &pbGlobal.S2C_CreatePlayer{}

	pl, err := m.am.CreatePlayer(acct, msg.Name)
	if err != nil {
		acct.SendProtoMessage(reply)
		return fmt.Errorf("handleCreatePlayer.AccountExecute failed: %w", err)
	}

	reply.Info = &pbGlobal.PlayerInfo{
		Id:        pl.GetId(),
		AccountId: pl.GetAccountID(),
		Name:      pl.GetName(),
		Exp:       pl.GetExp(),
		Level:     pl.GetLevel(),
	}

	acct.SendProtoMessage(reply)
	return nil
}

func (m *MsgRegister) handleWithdrawStrengthen(ctx context.Context, p ...interface{}) error {
	acct := p[0].(*player.Account)
	msg, ok := p[1].(*pbGlobal.C2S_WithdrawStrengthen)
	if !ok {
		return errors.New("handleWithdrawStrengthen failed: recv message body error")
	}

	pl, err := m.am.GetPlayerByAccount(acct)
	if err != nil {
		return err
	}

	return pl.WithdrawStrengthen(msg.GetValue())
}

func (m *MsgRegister) handleBuyStrengthen(ctx context.Context, p ...interface{}) error {
	acct := p[0].(*player.Account)
	pl, err := m.am.GetPlayerByAccount(acct)
	if err != nil {
		return err
	}

	return pl.BuyStrengthen()
}

func (m *MsgRegister) handleGuidePass(ctx context.Context, p ...interface{}) error {
	acct := p[0].(*player.Account)
	msg, ok := p[1].(*pbGlobal.C2S_GuidePass)
	if !ok {
		return errors.New("handleGuidePass failed: recv message body error")
	}

	pl, err := m.am.GetPlayerByAccount(acct)
	if err != nil {
		return err
	}

	return pl.GuideManager.GuidePass(msg.Index)
}

func (m *MsgRegister) handleSaveBattleArray(ctx context.Context, p ...interface{}) error {
	acct := p[0].(*player.Account)
	msg, ok := p[1].(*pbGlobal.C2S_SaveBattleArray)
	if !ok {
		return errors.New("handleSaveBattleArray failed: recv message body error")
	}

	pl, err := m.am.GetPlayerByAccount(acct)
	if err != nil {
		return err
	}

	return pl.SaveBattleArray(msg.GetBattleHeroId())
}
