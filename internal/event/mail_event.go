package event

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gookit/event"
	"github.com/leapord/prometheusx/internal/consts"
	"github.com/leapord/prometheusx/internal/service"
)

func init() {
	sendEmail()
}

func sendEmail() {
	event.On(consts.ALERT_WEB_HOOK_EVENT, event.ListenerFunc(func(e event.Event) error {
		messageId := g.NewVar(e.Get(consts.ALERT_MESSAGE_ID)).Int64()
		ctx := gctx.New()
		service.Mail().SendAlertEmail(ctx, messageId)
		return nil
	}), event.High)
}
