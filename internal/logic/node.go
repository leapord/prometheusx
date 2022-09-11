package logic

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	model "github.com/leapord/prometheusx/internal/model/do"
)

type sNode struct{}

var (
	Node = sNode{}
)

func (s *sNode) AddNode(ctx context.Context, node *model.Node) (err error) {
	node.CreateTime = gtime.Now()
	id, err := g.Model(model.Node{}).InsertAndGetId(node)
	node.Id = id

	if err != nil {
		g.Log().Error(ctx, err)
	}

	return
}

func (s *sNode) UpdateNode(ctx context.Context, node *model.Node) (err error) {
	gmodel := g.Model(model.Node{})

	if _, err = gmodel.Where(model.Node{Id: node.Id}).Count(); err == nil {
		if count, err := gmodel.UpdateAndGetAffected(node); err != nil && count == 1 {
			g.Log().Error(ctx, err)
		}
	} else {
		g.Log().Info(ctx, "no record found")
		err = errors.New("can not find record")
	}

	return
}

func (s *sNode) RemoveNode(ctx context.Context, id int) (node model.Node, err error) {
	if err = g.Model(model.Node{}).Where(model.Node{Id: id}).Scan(&node); err == nil {
		if _, err = g.Model(model.Node{}).Delete(model.Node{Id: id}); err != nil {
			g.Log().Error(ctx, err)
		}
	} else {
		g.Log().Info(ctx, "record not found id : ", id)
		err = errors.New("can not find record")
	}
	return
}

func (s *sNode) DetailNode(ctx context.Context, id int) (node model.Node, err error) {
	if err = g.Model(model.Node{}).Where(model.Node{Id: id}).Scan(&node); err != nil {
		g.Log().Info(ctx, "record not found id : ", id)
		err = errors.New("can not find record")
	}
	return
}

func (s *sNode) Page(ctx context.Context, pageNo int, pageSize int, node model.Node) (total int, models []model.Node, err error) {
	gmodel := g.Model(model.Node{})

	if !g.NewVar(node.Host).IsEmpty() {
		gmodel.WhereLike("host", "%"+g.NewVar(node.Host).String()+"%")
	}
	if !g.NewVar(node.Port).IsEmpty() {
		gmodel.WhereLike("port", "%"+g.NewVar(node.Port).String()+"%")
	}
	if !g.NewVar(node.Owner).IsEmpty() {
		gmodel.WhereLike("owner", "%"+g.NewVar(node.Owner).String()+"%")
	}
	if !g.NewVar(node.Group).IsEmpty() {
		gmodel.WhereLike("group", "%"+g.NewVar(node.Group).String()+"%")
	}
	if !g.NewVar(node.JobName).IsEmpty() {
		gmodel.WhereLike("job_name", "%"+g.NewVar(node.JobName).String()+"%")
	}

	if total, err = gmodel.Count(); err != nil {
		g.Log().Error(ctx, err)
		return
	}

	err = gmodel.Limit((pageNo-1)*pageSize, pageSize).Scan(&models)
	if err != nil {
		g.Log().Error(ctx, err)
	}

	return
}
