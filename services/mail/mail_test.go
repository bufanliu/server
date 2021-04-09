package mail

import (
	"context"
	"flag"
	"log"
	"testing"
	"time"

	"bitbucket.org/funplus/server/define"
	"bitbucket.org/funplus/server/excel"
	"bitbucket.org/funplus/server/logger"
	"bitbucket.org/funplus/server/services/mail/mailbox"
	"bitbucket.org/funplus/server/store"
	"bitbucket.org/funplus/server/utils"
	"github.com/urfave/cli/v2"
)

var (
	gameId      int16 = 404
	mailBoxNum        = 100
	mailNum           = 100
	mailManager *MailManager
	ctx         *cli.Context
)

func init() {
	// snow flake init
	utils.InitMachineID(gameId)

	// reload to project root path
	if err := utils.RelocatePath("/server", "\\server"); err != nil {
		log.Fatalf("relocate path failed: %s", err.Error())
	}

	// logger init
	logger.InitLogger("mail_test")

	// read excel files
	excel.ReadAllEntries("config/excel/")

	set := flag.NewFlagSet("mail_test", flag.ContinueOnError)
	set.String("db_dsn", "mongodb://localhost:27017", "mongodb address")
	set.String("database", "game", "mongodb default database")
	ctx = cli.NewContext(nil, set, nil)
	store.NewStore(ctx)

	mailManager = NewMailManager(ctx, &Mail{})
}

func TestAddMa(t *testing.T) {

	for n := 0; n < mailBoxNum; n++ {
		mb, err := mailManager.getMailBox(int64(n + 1))
		if !utils.ErrCheck(err, "getMailBox failed when TestMail", n+1) {
			continue
		}

		// 创建邮件
		_ = mb.AddResultHandler(ctx, func(c context.Context, mailBox *mailbox.MailBox) error {
			for m := 0; m < mailNum; m++ {
				newMail := &define.Mail{}
				newMail.Init()
				id, _ := utils.NextID(define.SnowFlake_Mail)
				newMail.Id = id
				newMail.OwnerId = int64(n + 1)
				newMail.Type = define.Mail_Type_System
				newMail.Date = int32(time.Now().Unix())
				newMail.ExpireDate = int32(time.Now().Add(time.Hour * 24).Unix())
				newMail.SenderName = "系统"
				newMail.Title = "测试标题"
				newMail.Content = "测试内容发送等待服务器返回消息发送等待服务器返回消息发送等待服务器返回消息发送等待服务器返回消息"
				newMail.Attachments = make([]*define.LootData, 0)
				newMail.Attachments = append(newMail.Attachments, &define.LootData{
					LootType: define.CostLoot_Item,
					LootMisc: 1,
					LootNum:  2,
				})

				_ = mailBox.BenchAddMail(c, newMail)
			}

			return nil
		})
	}
}

// func TestLoadMail(t *testing.T) {
// 	for n := 0; n < mailBoxNum; n++ {
// 		mb, err := mailManager.getMailBox(int64(n + 1))
// 		if !utils.ErrCheck(err, "getMailBox failed when TestMail", n+1) {
// 			continue
// 		}
// 	}

// }