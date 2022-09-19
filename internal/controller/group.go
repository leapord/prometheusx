package controller

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
	v1 "github.com/leapord/prometheusx/api/v1"
	model "github.com/leapord/prometheusx/internal/model/do"
	"github.com/leapord/prometheusx/internal/service"
)

type cGroup struct{}

var (
	Group = cGroup{}
)

func (c *cGroup) AddGroup(ctx context.Context, req *v1.GroupAddReq) (res v1.GroupAddRes, err error) {
	group := &model.Group{
		Name:           req.Name,
		Identification: req.Identification,
		CreateTime:     gtime.Now(),
	}

	if g.IsEmpty(group.Identification) {
		group.Identification = guid.S()
	}

	err = service.Group().AddGroup(ctx, group)
	if err == nil {
		res.Model = group
	}
	return
}

func (c *cGroup) UpdateGroup(ctx context.Context, req *v1.GroupUpdateReq) (res *v1.GroupUpdateRes, err error) {
	group := &model.Group{
		Id:             req.Id,
		Name:           req.Name,
		Identification: req.Identification,
	}

	err = service.Group().UpdateGroup(ctx, group)
	if err == nil {
		res = &v1.GroupUpdateRes{
			Model: group,
		}
	} else {
		g.Log().Error(ctx, "update err ", err)
	}

	return
}

func (c *cGroup) DeleteGroup(ctx context.Context, req *v1.GroupDeleteReq) (res *v1.GroupDeleteRes, err error) {
	group, err := service.Group().DeleteById(ctx, gconv.Int(req.Id))
	if err == nil {
		res = &v1.GroupDeleteRes{
			Model: group,
		}
	} else {
		g.Log().Error(ctx, "delete group error ", req.Id)
	}
	return
}

func (c *cGroup) DetailGroup(ctx context.Context, req *v1.GroupDetailReq) (res *v1.GroupDetailRes, err error) {
	group, err := service.Group().Detail(ctx, gconv.Int(req.Id))
	if err == nil {
		res = &v1.GroupDetailRes{
			Model: group,
		}
	}
	return
}

func (c *cGroup) Page(ctx context.Context, req *v1.GroupPageReq) (res *v1.GroupPageRes, err error) {
	group := &model.Group{}
	if !g.IsEmpty(req.Name) {
		group.Name = req.Name
	}
	if !g.IsEmpty(req.Identification) {
		group.Identification = req.Identification
	}
	list, total, err := service.Group().Page(ctx, *group, (req.PageNo-1)*req.PageSize, req.PageSize)
	if err == nil {
		res = &v1.GroupPageRes{
			Total:    total,
			PageNo:   req.PageNo,
			PageSize: req.PageSize,
			Models:   list,
		}
	}
	return
}

func (c *cGroup) List(ctx context.Context, req *v1.GroupListReq) (res *v1.GroupListRes, err error) {
	groups, err := service.Group().List(ctx)
	if err == nil {
		res = &v1.GroupListRes{
			Model: groups,
		}
	}

	return
}
