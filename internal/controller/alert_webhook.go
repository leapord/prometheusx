package controller

import (
	"context"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	v1 "github.com/leapord/prometheusx/api/v1"
	model "github.com/leapord/prometheusx/internal/model/do"
	"github.com/leapord/prometheusx/internal/service"
)

type cAlertWebhook struct{}

var (
	AlertWebhook = cAlertWebhook{}
)

// 对接alert manager webhook
func (c *cAlertWebhook) WebHook(ctx context.Context, req *v1.AlertHookReq) (res *v1.AlertHookRes, err error) {
	request := g.RequestFromCtx(ctx)
	json, err := gjson.DecodeToJson(request.GetBody())
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	alertJsonArray := json.GetJsons("alerts")
	alerts := []model.Alert{}

	for _, item := range alertJsonArray {
		alert := model.Alert{
			GroupKey:     json.Get("groupKey").String(),
			ExternalUrl:  json.Get("externalURL").String(),
			Status:       item.Get("status").String(),
			Labels:       item.GetJson("labels").MapStrAny(),
			Annotations:  item.GetJson("annotations").MapStrAny(),
			StartsAt:     item.Get("startsAt").GTime(),
			EndsAt:       item.Get("endsAt").GTime(),
			GeneratorUrl: item.Get("generatorURL").String(),
			Fingerprint:  item.Get("fingerprint").String(),
		}
		alerts = append(alerts, alert)
	}
	service.Alert().AddAlert(ctx, alerts)

	return
}
