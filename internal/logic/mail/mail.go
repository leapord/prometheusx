package mail

import (
	"context"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/leapord/prometheusx/internal/consts"
	"github.com/leapord/prometheusx/internal/model/bo"
	model "github.com/leapord/prometheusx/internal/model/do"
	"github.com/leapord/prometheusx/internal/model/entity"
	"github.com/leapord/prometheusx/internal/model/vo"
	"github.com/leapord/prometheusx/internal/service"
	"github.com/leapord/prometheusx/utility"
)

type sMail struct{}

func init() {
	service.RegisterMail(New())
}

func New() *sMail {
	return &sMail{}
}

func (s *sMail) SendAlertEmail(ctx context.Context, messageId int64) {
	alert := vo.Alert{}
	g.Model(entity.Alert{}).Where(model.Alert{Id: messageId}).Scan(&alert)

	stampServerAddress, _ := g.Model(entity.Config{}).Where(model.Config{Name: consts.STAMP_SERVER_ADDRESS}).Fields("value").Value()
	stampServerPort, _ := g.Model(entity.Config{}).Where(model.Config{Name: consts.STAMP_SERVER_PORT}).Fields("value").Value()
	stampServerAccount, _ := g.Model(entity.Config{}).Where(model.Config{Name: consts.STAMP_SERVER_ACCOUNT}).Fields("value").Value()
	stampServerPassword, _ := g.Model(entity.Config{}).Where(model.Config{Name: consts.STAMP_SERVER_PASSWORD}).Fields("value").Value()

	if g.IsEmpty(stampServerAddress.String()) || g.IsEmpty(stampServerPort.String()) || g.IsEmpty(stampServerAccount.String()) || g.IsEmpty(stampServerPassword.String()) {
		return
	}

	mailServer := bo.MailServer{
		Host:     stampServerAddress.String(),
		Port:     stampServerPort.Int(),
		Account:  stampServerAccount.String(),
		Password: stampServerPassword.String(),
	}
	if g.IsEmpty(alert.Labels["owner"]) {
		return
	}
	user := entity.User{}
	g.Model(entity.User{}).Where(model.User{LoginName: alert.Labels["owner"]}).Scan(&user)

	if !g.IsEmpty(user.Id) {
		if json, err := gjson.EncodeString(alert); err == nil {
			utility.SendMail(ctx, "Alert", json, mailServer, user.Email)
		} else {
			g.Log().Error(ctx, err)
		}
	}

}
