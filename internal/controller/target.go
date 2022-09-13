package controller

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "github.com/leapord/prometheusx/api/v1"
	"github.com/leapord/prometheusx/internal/service"
)

type cTarget struct{}

var (
	Target = cTarget{}
)

func (c *cTarget) Node(ctx context.Context, req *v1.NodeTargetReq) (res *v1.NodeActiveRes, err error) {
	resp := g.RequestFromCtx(ctx).Response
	json, err := service.Node().Target(ctx)
	resp.WriteJson(json)
	return
}
