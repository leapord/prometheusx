package logic

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	model "github.com/leapord/prometheusx/internal/model/do"
	entity "github.com/leapord/prometheusx/internal/model/entity"
	"github.com/leapord/prometheusx/internal/service"
)

type sNode struct{}

func New() *sNode {
	return &sNode{}
}

func init() {
	service.RegisterNode(New())
}

// 添加
func (s *sNode) AddNode(ctx context.Context, node *model.Node) (err error) {
	node.CreateTime = gtime.Now()
	id, err := g.Model(entity.Node{}).InsertAndGetId(node)
	node.Id = id

	if err != nil {
		g.Log().Error(ctx, err)
	}

	return
}

// 修改
func (s *sNode) UpdateNode(ctx context.Context, node *model.Node) (err error) {
	gmodel := g.Model(entity.Node{})

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

// 删除
func (s *sNode) RemoveNode(ctx context.Context, id int) (node entity.Node, err error) {
	if err = g.Model(entity.Node{}).Where(model.Node{Id: id}).Scan(&node); err == nil {
		if _, err = g.Model(model.Node{}).Delete(model.Node{Id: id}); err != nil {
			g.Log().Error(ctx, err)
		}
	} else {
		g.Log().Info(ctx, "record not found id : ", id)
		err = errors.New("can not find record")
	}
	return
}

// 单个详情
func (s *sNode) DetailNode(ctx context.Context, id int) (node entity.Node, err error) {
	if err = g.Model(entity.Node{}).Where(model.Node{Id: id}).Scan(&node); err != nil {
		g.Log().Info(ctx, "record not found id : ", id)
		err = errors.New("can not find record")
	}
	return
}

// 分页
func (s *sNode) Page(ctx context.Context, pageNo int, pageSize int, node model.Node) (total int, models []entity.Node, err error) {
	gmodel := g.Model(entity.Node{})

	if !g.IsEmpty(node.Host) {
		gmodel.WhereLike("host", "%"+g.NewVar(node.Host).String()+"%")
	}
	if !g.IsEmpty(node.Port) {
		gmodel.WhereLike("port", "%"+g.NewVar(node.Port).String()+"%")
	}
	if !g.IsEmpty(node.Owner) {
		gmodel.WhereLike("owner", "%"+g.NewVar(node.Owner).String()+"%")
	}
	if !g.IsEmpty(node.Group) {
		gmodel.WhereLike("group", "%"+g.NewVar(node.Group).String()+"%")
	}
	if !g.IsEmpty(node.JobName) {
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

// 改变active状态
func (s *sNode) Active(ctx context.Context, id int, active bool) (err error) {
	_, err = g.Model(entity.Node{}).Where(model.Node{Id: id}).UpdateAndGetAffected(model.Node{Active: active})
	return
}

func (s *sNode) Target(ctx context.Context) (result string, err error) {
	nodes := []entity.Node{}
	err = g.Model(entity.Node{}).Where(model.Node{Active: true}).Scan(&nodes)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	list := g.List{}

	for i := 0; i < len(nodes); i++ {
		node := nodes[i]
		labels := g.Map{
			"owner":    node.Owner,
			"job_name": node.JobName,
			"group":    node.Group,
		}

		nodeLabels := node.Labels

		if !g.IsEmpty(nodeLabels) && gjson.Valid(nodeLabels) {
			nodeLabelsJson := g.NewVar(nodeLabels).Map()
			for key, value := range nodeLabelsJson {
				labels[key] = value
			}
		}

		list = append(list, g.Map{
			"targets": g.NewVar(node.Host).String() + ":" + g.NewVar(node.Port).String(),
			"labels":  labels,
		})
	}

	result = gjson.MustEncodeString(list)
	return
}
