package controller

import (
	"context"

	v1 "github.com/leapord/prometheusx/api/v1"
	model "github.com/leapord/prometheusx/internal/model/do"
	"github.com/leapord/prometheusx/internal/service"
)

type cAlert struct{}

var (
	Alert = cAlert{}
)

// 分页查询告警信息
func (c *cAlert) PageQuery(ctx context.Context, req *v1.AlertPageReq) (res *v1.AlertPageRes, err error) {
	models, total, err := service.Alert().Page(ctx, req.PageNo, req.PageSize, model.Alert{Labels: req.Labels})
	if err == nil {
		res = &v1.AlertPageRes{
			Total:    total,
			PageNo:   req.PageNo,
			PageSize: req.PageSize,
			Models:   models,
		}
	}
	return
}
