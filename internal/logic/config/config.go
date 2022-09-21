package config

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	model "github.com/leapord/prometheusx/internal/model/do"
	"github.com/leapord/prometheusx/internal/model/entity"
	"github.com/leapord/prometheusx/internal/service"
)

type sConfig struct{}

func init() {
	service.RegisterConfig(New())
}

func New() *sConfig {
	return &sConfig{}
}

// 添加
func (s *sConfig) Add(ctx context.Context, config model.Config) (result entity.Config, err error) {
	id, err := g.Model(entity.Config{}).InsertAndGetId(config)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	err = g.Model(entity.Config{}).Where(model.Config{Id: id}).Scan(&result)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	return
}

// 更新
func (s *sConfig) Update(ctx context.Context, config model.Config) (result entity.Config, err error) {
	_, err = g.Model(entity.Config{}).Where(model.Config{Id: config.Id}).UpdateAndGetAffected(config)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	err = g.Model(entity.Config{}).Where(model.Config{Id: config.Id}).Scan(&result)
	return
}

// 删除
func (s *sConfig) Remove(ctx context.Context, id int) (config entity.Config, err error) {
	gmodel := g.Model(entity.Config{})
	err = gmodel.Where(model.Config{Id: id}).Scan(&config)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	_, err = gmodel.Where(model.Config{Id: id}).Delete()
	if err != nil {
		g.Log().Error(ctx, err)
	}
	return
}

// 查询单个详情
func (s *sConfig) Detail(ctx context.Context, id int) (config entity.Config, err error) {
	err = g.Model(entity.Config{}).Where(model.Config{Id: id}).Scan(&config)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	return
}

// 分页
func (s *sConfig) Page(ctx context.Context, pageNo int, pageSize int, config model.Config) (configs []entity.Config, total int, err error) {
	gmodel := g.Model(entity.Config{})
	if !g.IsEmpty(config.Name) {
		gmodel.WhereLike("name", "%"+g.NewVar(config.Name).String()+"%")
	}
	total, err = gmodel.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	err = gmodel.Limit((pageNo-1)*pageSize, pageSize).Scan(&configs)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	return
}

// 根据名称查询配置内容
func (s *sConfig) NameQuery(ctx context.Context, name string) (result entity.Config, err error) {
	err = g.Model(entity.Config{}).Where(model.Config{Name: name}).Scan(&result)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	return
}
