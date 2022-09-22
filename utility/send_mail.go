package utility

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/leapord/prometheusx/internal/model/bo"
	"gopkg.in/gomail.v2"
)

func SendMail(ctx context.Context, subject, body string, mailServer bo.MailServer, to ...string) (err error) {
	m := gomail.NewMessage()

	m.SetHeader("From", mailServer.Account)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(mailServer.Host, mailServer.Port, mailServer.Account, mailServer.Password)

	if err := d.DialAndSend(m); err != nil {
		g.Log().Error(ctx, err)
	}
	return
}
