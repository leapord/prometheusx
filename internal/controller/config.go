package controller

import (
	"context"

	v1 "github.com/leapord/prometheusx/api/v1"
	model "github.com/leapord/prometheusx/internal/model/do"
	"github.com/leapord/prometheusx/internal/service"
)

type cConfig struct{}

var (
	Config = cConfig{}
)

// 添加
func (c *cConfig) AddConfig(ctx context.Context, req *v1.ConfigAddReq) (res *v1.ConfigAddRes, err error) {
	config, err := service.Config().Add(ctx, model.Config{Name: req.Name, Value: req.Value})
	if err == nil {
		res = &v1.ConfigAddRes{Model: config}
	}
	return
}

// 更新
func (c *cConfig) UpdateConfig(ctx context.Context, req *v1.ConfigUpdateReq) (res *v1.ConfigUpdateRes, err error) {
	config, err := service.Config().Update(ctx, model.Config{Id: req.Id, Name: req.Name, Value: req.Value})
	if err == nil {
		res = &v1.ConfigUpdateRes{Model: config}
	}
	return
}

// 删除
func (c *cConfig) RemoveConfig(ctx context.Context, req *v1.ConfigRemoveReq) (res *v1.ConfigRemoveRes, err error) {
	config, err := service.Config().Remove(ctx, req.Id)
	if err == nil {
		res = &v1.ConfigRemoveRes{
			Model: config,
		}
	}
	return
}

// 查询单个详情
func (c *cConfig) DetailConfig(ctx context.Context, req *v1.ConfigDetailReq) (res *v1.ConfigDetailRes, err error) {
	config, err := service.Config().Detail(ctx, req.Id)
	if err == nil {
		res = &v1.ConfigDetailRes{Model: config}
	}
	return
}

// 分页查询
func (c *cConfig) PageConfig(ctx context.Context, req *v1.ConfigPageReq) (res *v1.ConfigPageRes, err error) {
	models, total, err := service.Config().Page(ctx, req.PageNo, req.PageSize, model.Config{Name: req.Name})
	if err == nil {
		res = &v1.ConfigPageRes{
			Models:   models,
			Total:    total,
			PageNo:   req.PageNo,
			PageSize: req.PageSize,
		}
	}
	return
}

func (c *cConfig) NameConfigQuery(ctx context.Context, req *v1.ConfigNameReq) (res *v1.ConfigNameRes, err error) {
	model, err := service.Config().NameQuery(ctx, req.Name)
	if err == nil {
		res = &v1.ConfigNameRes{
			Model: model,
		}
	}
	return
}
