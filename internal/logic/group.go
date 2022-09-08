package logic

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	model "github.com/leapord/prometheus_ext/internal/model/do"
	entity "github.com/leapord/prometheus_ext/internal/model/entity"
)

type sGroup struct{}

var (
	Group = sGroup{}
)

func (s *sGroup) AddGroup(ctx context.Context, group *model.Group) (err error) {
	group.Id, err = g.Model(model.Group{}).InsertAndGetId(group)
	if err != nil {
		g.Log().Error(ctx, err)
	} else {
		g.Log().Debug(ctx, "insert prometheus group :", group)
	}
	return
}

func (s *sGroup) UpdateGroup(ctx context.Context, group *model.Group) (err error) {
	_, err = g.Model(model.Group{}).UpdateAndGetAffected(group)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	return
}

func (s *sGroup) DeleteById(ctx context.Context, id int) (group model.Group, err error) {
	gmodel := g.Model(model.Group{})
	err = gmodel.Where(model.Group{Id: id}).Scan(&group)
	if err != nil {
		return
	}
	_, err = gmodel.Delete(model.Group{Id: id})
	return
}

func (s *sGroup) Detail(ctx context.Context, id int) (group model.Group, err error) {
	err = g.Model(model.Group{}).Where(model.Group{
		Id: id,
	}).Scan(&group)
	if err != nil {
		g.Log().Error(ctx, err)
	}

	return
}

func (s *sGroup) Page(ctx context.Context, group model.Group, pageNo int, pageSize int) (list []model.Group, total int, err error) {
	gmodel := g.Model(model.Group{}).Where(group)
	total, err = gmodel.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	err = gmodel.Limit(pageNo, pageSize).Scan(&list)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	return
}

func (s *sGroup) List(ctx context.Context) (groups []entity.Group, err error) {
	err = g.Model(model.Group{}).Scan(&groups)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	return
}
