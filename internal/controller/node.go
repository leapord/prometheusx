package controller

import (
	"context"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	v1 "github.com/leapord/prometheusx/api/v1"
	model "github.com/leapord/prometheusx/internal/model/do"
	"github.com/leapord/prometheusx/internal/service"
)

type cNode struct{}

var (
	Node = cNode{}
)

// 新增
func (c *cNode) AddNode(ctx context.Context, req *v1.NodeAddReq) (res v1.NodeAddRes, err error) {

	if !gjson.Valid(req.Labels) {
		err = gerror.NewCode(gcode.CodeValidationFailed, "labels 必须是JSON")
		return
	}

	node := model.Node{
		Host:    req.Host,
		Port:    req.Port,
		Owner:   req.Owner,
		Group:   req.Group,
		JobName: req.JobName,
		Labels:  req.Labels,
	}
	err = service.Node().AddNode(ctx, &node)
	if err == nil {
		res.Model = node
	}
	return
}

// 修改
func (c *cNode) UpdateNode(ctx context.Context, req *v1.NodeUpdateReq) (res *v1.NodeUpdateRes, err error) {

	if !gjson.Valid(req.Labels) {
		err = gerror.NewCode(gcode.CodeValidationFailed, "labels 必须是JSON")
		return
	}

	node := model.Node{
		Id:      req.Id,
		Host:    req.Host,
		Port:    req.Port,
		Owner:   req.Owner,
		Group:   req.Group,
		JobName: req.JobName,
	}
	if !g.IsEmpty(req.Labels) {
		node.Labels = req.Labels
	}
	err = service.Node().UpdateNode(ctx, &node)

	if err == nil {
		res = &v1.NodeUpdateRes{
			Model: node,
		}
	}

	return
}

// 删除
func (c *cNode) RemoveNode(ctx context.Context, req *v1.NodeRemoveReq) (res *v1.NodeRemoveRes, err error) {
	node, err := service.Node().RemoveNode(ctx, g.NewVar(req.Id).Int())
	if err == nil {
		res = &v1.NodeRemoveRes{
			Model: node,
		}
	}
	return
}

// 详情
func (c *cNode) DetilNode(ctx context.Context, req *v1.NodeDetailReq) (res *v1.NodeDetailRes, err error) {
	node, err := service.Node().DetailNode(ctx, g.NewVar(req.Id).Int())
	if err == nil {
		res = &v1.NodeDetailRes{
			Model: node,
		}
	}
	return
}

// 分页查询
func (c *cNode) PageNode(ctx context.Context, req *v1.NodePageReq) (res *v1.NodePageRes, err error) {
	node := model.Node{
		Host:    req.Host,
		Port:    req.Port,
		Group:   req.Group,
		JobName: req.JobName,
		Owner:   req.Owner,
	}
	total, models, err := service.Node().Page(ctx, g.NewVar(req.PageNo).Int(), g.NewVar(req.PageSize).Int(), node)
	if err == nil {
		res = &v1.NodePageRes{
			PageNo:   req.PageNo,
			PageSize: req.PageSize,
			Total:    total,
			Models:   models,
		}
	}
	return
}

func (c *cNode) ChanageActiveStatus(ctx context.Context, req *v1.NodeActiveReq) (res *v1.NodeActiveRes, err error) {
	err = service.Node().Active(ctx, req.Id, req.Active)
	return
}
