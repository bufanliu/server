package game

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/gogo/protobuf/proto"
	"github.com/hellodudu/Ultimate/utils"
	logger "github.com/sirupsen/logrus"
)

type ClientMgr struct {
	mapClient      sync.Map
	mapConn        sync.Map
	g              *Game
	waitGroup      utils.WaitGroupWrapper
	ctx            context.Context
	cancel         context.CancelFunc
	chKickClientID chan uint32
}

func NewClientMgr(game *Game) *ClientMgr {
	cm := &ClientMgr{
		g:              game,
		chKickClientID: make(chan uint32, game.opts.ClientConnectMax),
	}

	cm.ctx, cm.cancel = context.WithCancel(game.ctx.Background())
	cm.g.db.orm.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").AutoMigrate(ClientPeersInfo{})

	return cm, nil
}

func (cm *ClientMgr) Main() error {
	exitCh := make(chan error)
	var once sync.Once
	exitFunc := func(err error) {
		once.Do(func() {
			if err != nil {
				log.Fatal("ClientMgr Main() error:", err)
			}
			exitCh <- err
		})
	}

	cm.waitGroup.Wrap(func() {
		exitFunc(cm.Run())
	})

	return <-exitCh
}

func (cm *ClientMgr) Exit() {
	cm.cancel()
	close(cm.chKickClientID)
	cm.waitGroup.Wait()
}

func (cm *ClientMgr) AddClient(id uint32, name string, c *TCPConn) (*Client, error) {
	if int32(id) == -1 {
		return nil, errors.New("add world id invalid!")
	}

	if _, ok := cm.mapClient.Load(id); ok {
		cm.KickClient(id, "AddClient")
	}

	var numConn uint32
	cm.mapConn.Range(func(_, _ interface{}) bool {
		numConn++
		return true
	})

	if numConn >= uint32(cm.g.opts.ClientConnectMax) {
		return nil, errors.New("reach game server's max client connect num!")
	}

	// new client
	peerInfo := &ClientPeersInfo{
		ID:   id,
		Name: name,
		c:    c,
	}

	client := NewClient(peerInfo, cm)
	cm.mapClient.Store(client.GetID(), client)
	cm.mapConn.Store(c, client)
	logger.Info(fmt.Sprintf("add client <id:%d, name:%s, con:%v> success!", client.GetID(), client.GetName(), client.GetCon()))

	// client main
	cm.waitGroup.Wrap(func() {
		err := client.Main()
		if err != nil {
			client.Exit()
		}
		cm.mapClient.Delete(client.GetID())
		cm.mapConn.Delete(client.peerInfo.c)
	})

	return client, nil
}

func (cm *ClientMgr) GetClientByID(id uint32) *Client {
	v, ok := cm.mapClient.Load(id)
	if !ok {
		return nil
	}

	return v.(*Client)
}

func (cm *ClientMgr) GetClientByCon(con TcpCon) *Client {
	v, ok := cm.mapConn.Load(con)
	if !ok {
		return nil
	}

	return v.(*Client)
}

func (cm *ClientMgr) DisconnectClient(con TcpCon) {
	v, ok := cm.mapConn.Load(con)
	if !ok {
		return
	}

	client, ok := v.(*Client)
	if !ok {
		return
	}

	logger.WithFields(logger.Fields{
		"id": client.GetID(),
	}).Warn("Client disconnected!")

	client.cancel()
}

func (cm *ClientMgr) KickClient(id uint32, reason string) {
	v, ok := cm.mapClient.Load(id)
	if !ok {
		return
	}

	client, ok := v.(*client)
	if !ok {
		return
	}

	logger.WithFields(logger.Fields{
		"id":     client.GetID(),
		"reason": reason,
	}).Warn("Client was kicked!")

	client.cancel()
}

func (cm *ClientMgr) BroadCast(msg proto.Message) {
	cm.mapClient.Range(func(_, v interface{}) bool {
		if client, ok := v.(*client); ok {
			client.SendProtoMessage(msg)
		}
		return true
	})
}

func (cm *ClientMgr) Run() error {
	for {
		select {
		case <-cm.ctx.Done():
			logger.Print("world session context done!")
			return nil
		case wid := <-cm.chKickClientID:
			cm.KickClient(wid, "time out")
		}
	}

	return nil
}