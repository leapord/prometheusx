package controller

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	v1 "github.com/leapord/prometheusx/api/v1"
	model "github.com/leapord/prometheusx/internal/model/do"
	"github.com/leapord/prometheusx/internal/service"
)

type cRules struct{}

var (
	Rules = cRules{}
	types = g.SliceStr{
		"alert",
		"record",
	}
)

func contains(elems []string, v string) bool {
	for _, s := range elems {
		if gstr.Equal(s, v) {
			return true
		}
	}
	return false
}

// 添加
func (c *cRules) AddRule(ctx context.Context, req *v1.RuleAddReq) (res *v1.RuleAddRes, err error) {
	if !contains(types, req.Type) {
		err = gerror.NewCode(gcode.CodeValidationFailed, "type must be 'alert' or 'record'")
		return
	}
	rules := model.Rules{
		GroupName: req.GroupName,
		Type:      req.Type,
		Content:   req.Content,
	}
	model, err := service.Rules().Add(ctx, rules)
	res = &v1.RuleAddRes{
		Model: model,
	}
	return
}

// 更新
func (c *cRules) UpdateRule(ctx context.Context, req *v1.RuleUpdateReq) (res *v1.RuleUpdateRes, err error) {
	if !contains(types, req.Type) {
		err = gerror.NewCode(gcode.CodeValidationFailed, "type must be 'alert' or 'record'")
		return
	}
	rules := model.Rules{
		Id:        req.Id,
		GroupName: req.GroupName,
		Type:      req.Type,
		Content:   req.Content,
		Active:    req.Active,
	}
	model, err := service.Rules().Update(ctx, rules)
	res = &v1.RuleUpdateRes{
		Model: model,
	}
	return
}

// 删除
func (c *cRules) RemoveRule(ctx context.Context, req *v1.RuleRemoveReq) (res *v1.RuleRemoveRes, err error) {
	rules, err := service.Rules().Remove(ctx, req.Id)
	res = &v1.RuleRemoveRes{
		Model: rules,
	}
	return
}

// 查询单个详情
func (c *cRules) DetailRule(ctx context.Context, req *v1.RuleDetailReq) (res *v1.RuleDetailRes, err error) {
	model, err := service.Rules().Detail(ctx, req.Id)
	res = &v1.RuleDetailRes{
		Model: model,
	}
	return
}

// 分页查询
func (c *cRules) PageRule(ctx context.Context, req *v1.RulePageReq) (res *v1.RulePageRes, err error) {
	if !g.IsEmpty(req.Type) && !contains(types, req.Type) {
		err = gerror.NewCode(gcode.CodeValidationFailed, "type must be 'alert' or 'record'")
		return
	}
	rules := model.Rules{
		GroupName: req.GroupName,
		Type:      req.Type,
	}

	models, total, err := service.Rules().Page(ctx, req.PageNo, req.PageSize, rules)
	res = &v1.RulePageRes{
		Models:   models,
		Total:    total,
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
	}
	return
}

// 生成规则文件
func (c *cRules) GeneratedFile(ctx context.Context, req *v1.RuleFileGeneratedReq) (res *v1.RuleFileGeneratedRes, err error) {
	err = service.Rules().GeneratedFile(ctx)
	return
}

// 改变规则激活状态
func (c *cRules) Active(ctx context.Context, req *v1.RuleActiveReq) (res *v1.RuleActiveRes, err error) {
	err = service.Rules().Active(ctx, req.Id, req.Active)
	return
}
