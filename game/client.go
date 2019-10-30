package game

import (
	"bytes"
	"context"
	"encoding/binary"
	"log"
	"sync"

	"github.com/gogo/protobuf/proto"
	"github.com/hellodudu/Ultimate/utils"
	"github.com/hellodudu/yokai_server/game/define"
	logger "github.com/sirupsen/logrus"
)

type ClientPeersInfo struct {
	ID   uint32 `gorm:"type:int(10);primary_key;column:id;default:0;not null"`
	Name string `gorm:"type:varchar(32);column:name;default:'';not null"`
	c    *TcpCon
	cm   *ClientMgr
}

type Client struct {
	peerInfo *ClientPeersInfo

	ctx       context.Context
	cancel    context.CancelFunc
	waitGroup utils.WaitGroupWrapper
	chw       chan uint32

	mapPlayer map[int64]*pbGame.CrossPlayerInfo
	mapGuild  map[int64]*pbGame.CrossGuildInfo
}

func NewClient(peerInfo *ClientPeersInfo) *Client {
	client := &Client{
		peerInfo:  peerInfo,
		mapPlayer: make(map[int64]*pbGame.CrossPlayerInfo),
		mapGuild:  make(map[int64]*pbGame.CrossGuildInfo),
	}

	client.ctx, client.cancel = context.WithCancel(peerInfo.cm.ctx)

	return client
}

func (Client) TableName() string {
	return "Client"
}

func (c *Client) GetID() uint32 {
	return c.peerInfo.ID
}

func (c *Client) GetName() string {
	return c.peerInfo.Name
}

func (c *Client) GetCon() *TCPConn {
	return c.peerInfo.c
}

func (c *Client) Main() error {
	c.loadFromDB()
	c.saveToDB()

	exitCh := make(chan error)
	var once sync.Once
	exitFunc := func(err error) {
		once.Do(func() {
			if err != nil {
				log.Fatal("Client Main() error:", err)
			}
			exitCh <- err
		})
	}

	c.waitGroup.Wrap(func() {
		exitFunc(c.Run())
	})

	return <-exitCh
}

func (c *Client) loadFromDB() {
	c.peerInfo.cm.g.db.orm.First(c.peerInfo)
}

func (c *Client) saveToDB() {
	c.peerInfo.cm.g.db.orm.Save(c.peerInfo)
}

func (c *Client) Exit() {
	c.cancel()
	c.peerInfo.c.Close()
}

func (c *Client) Run() error {
	for {
		select {
		// context canceled
		case <-c.ctx.Done():
			logger.WithFields(logger.Fields{
				"id": c.GetID(),
			}).Info("Client context done!")
			return nil
		}
	}
}

func (c *Client) SendProtoMessage(p proto.Message) {
	// reply message length = 4bytes size + 8bytes size BaseNetMsg + 2bytes message_name size + message_name + proto_data
	out, err := proto.Marshal(p)
	if err != nil {
		logger.Warn(err)
		return
	}

	typeName := proto.MessageName(p)
	baseMsg := &define.BaseNetMsg{}
	msgSize := binary.Size(baseMsg) + 2 + len(typeName) + len(out)
	baseMsg.ID = utils.Crc32("MUW_DirectProtoMsg")
	baseMsg.Size = uint32(msgSize)

	var resp []byte = make([]byte, 4+msgSize)
	binary.LittleEndian.PutUint32(resp[:4], uint32(msgSize))
	binary.LittleEndian.PutUint32(resp[4:8], baseMsg.ID)
	binary.LittleEndian.PutUint32(resp[8:12], baseMsg.Size)
	binary.LittleEndian.PutUint16(resp[12:12+2], uint16(len(typeName)))
	copy(resp[14:14+len(typeName)], []byte(typeName))
	copy(resp[14+len(typeName):], out)

	if _, err := c.peerInfo.c.Write(resp); err != nil {
		logger.Warn("send proto msg error:", err)
		return
	}
}

func (c *Client) SendTransferMessage(data []byte) {
	resp := make([]byte, 4+len(data))
	binary.LittleEndian.PutUint32(resp[:4], uint32(len(data)))
	copy(resp[4:], data)

	if _, err := c.peerInfo.c.Write(resp); err != nil {
		logger.Warn("send transfer msg error:", err)
		return
	}

	// for testing disconnected from Client server
	transferMsg := &define.TransferNetMsg{}
	byTransferMsg := make([]byte, binary.Size(transferMsg))

	copy(byTransferMsg, data[:binary.Size(transferMsg)])
	buf := &bytes.Buffer{}
	if _, err := buf.Write(byTransferMsg); err != nil {
		return
	}

	// get top 4 bytes messageid
	if err := binary.Read(buf, binary.LittleEndian, transferMsg); err != nil {
		return
	}
}