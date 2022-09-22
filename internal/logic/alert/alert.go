package alert

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gookit/event"
	"github.com/leapord/prometheusx/internal/consts"
	model "github.com/leapord/prometheusx/internal/model/do"
	"github.com/leapord/prometheusx/internal/model/entity"
	vo "github.com/leapord/prometheusx/internal/model/vo"
	"github.com/leapord/prometheusx/internal/service"
)

type sAlert struct{}

func init() {
	service.RegisterAlert(New())
}

func New() *sAlert {
	return &sAlert{}
}

// 添加告警消息
func (s *sAlert) AddAlert(ctx context.Context, alerts []model.Alert) (err error) {
	gmodel := g.Model(entity.Alert{})
	for _, alert := range alerts {
		id, err := gmodel.InsertAndGetId(alert)
		event.Fire(consts.ALERT_WEB_HOOK_EVENT, event.M{consts.ALERT_MESSAGE_ID: id})
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}
	return
}

// 分页查询
func (s *sAlert) Page(ctx context.Context, pageNo int, pageSize int, alert model.Alert) (models []vo.Alert, total int, err error) {
	gmodel := g.Model(entity.Alert{})
	if !g.IsEmpty(alert.Labels) {
		gmodel.WhereLike("labels", "%"+g.NewVar(alert.Labels).String()+"%")
	}

	total, err = gmodel.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	err = gmodel.Order("create_time desc").Scan(&models)
	if err != nil {
		g.Log().Error(ctx, err)
	}

	return
}
