package rules

import (
	"bytes"
	"context"
	"os/exec"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/leapord/prometheusx/internal/consts"
	model "github.com/leapord/prometheusx/internal/model/do"
	"github.com/leapord/prometheusx/internal/model/entity"
	"github.com/leapord/prometheusx/internal/model/vo"
	"github.com/leapord/prometheusx/internal/service"
	"github.com/leapord/prometheusx/utility"
	"gopkg.in/yaml.v3"
)

type sRules struct{}

func init() {
	service.RegisterRules(New())
}

func New() *sRules {
	return &sRules{}
}

// 添加
func (s *sRules) Add(ctx context.Context, rules model.Rules) (result entity.Rules, err error) {
	id, err := g.Model(entity.Rules{}).InsertAndGetId(rules)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	err = g.Model(entity.Rules{}).Where(model.Rules{Id: id}).Scan(&result)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	return
}

// 更新
func (s *sRules) Update(ctx context.Context, rules model.Rules) (rule entity.Rules, err error) {
	affected, err := g.Model(entity.Rules{}).Where(model.Rules{Id: rules.Id}).UpdateAndGetAffected(rules)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	if affected == 0 {
		err = gerror.NewCode(gcode.CodeNotFound, "ID对应的结果找不到")
		return
	}
	err = g.Model(entity.Rules{}).Where(model.Rules{Id: rules.Id}).Scan(&rule)
	return
}

// 删除
func (s *sRules) Remove(ctx context.Context, id int) (rules entity.Rules, err error) {
	gmodel := g.Model(entity.Rules{})
	err = gmodel.Where(model.Rules{Id: id}).Scan(&rules)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	_, err = gmodel.Where(model.Rules{Id: id}).Delete()
	if err != nil {
		g.Log().Error(ctx, err)
	}
	return
}

// 查询单个详情
func (s *sRules) Detail(ctx context.Context, id int) (rules entity.Rules, err error) {
	err = g.Model(entity.Rules{}).Where(model.Rules{Id: id}).Scan(&rules)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	return
}

// 分页
func (s *sRules) Page(ctx context.Context, pageNo int, pageSize int, rules model.Rules) (models []entity.Rules, total int, err error) {
	gmodel := g.Model(entity.Rules{})
	if !g.IsEmpty(rules.GroupName) {
		gmodel.WhereLike("group_name", "%"+g.NewVar(rules.GroupName).String()+"%")
	}
	if !g.IsEmpty(rules.Type) {
		gmodel.WhereLike("type", "%"+g.NewVar(rules.GroupName).String()+"%")
	}
	if !g.IsEmpty(rules.Active) {
		gmodel.Where(model.Rules{Active: rules.Active})
	}
	total, err = gmodel.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	err = gmodel.Limit((pageNo-1)*pageSize, pageSize).Scan(&models)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	return
}

func (s *sRules) GeneratedFile(ctx context.Context) (err error) {
	prometheusConfigFile, _ := g.Model(entity.Config{}).Where(model.Config{Name: consts.PROMETHEUS_CONFIG_PATH}).Fields("value").Value()
	prometheusToolFile, _ := g.Model(entity.Config{}).Where(model.Config{Name: consts.PROMETHEUS_TOOL_PATH}).Fields("value").Value()
	prometheusRuleDirectory, _ := g.Model(entity.Config{}).Where(model.Config{Name: consts.PROMETHEUS_RULE_PATH}).Fields("value").Value()
	prometheusAdminUrl, _ := g.Model(entity.Config{}).Where(model.Config{Name: consts.PROMETHEUS_ADMIN_URL}).Fields("value").Value()

	err = generatedFile(ctx, "alert", prometheusConfigFile.String(), prometheusToolFile.String(), prometheusRuleDirectory.String())
	if err != nil {
		return
	}
	err = generatedFile(ctx, "record", prometheusConfigFile.String(), prometheusToolFile.String(), prometheusRuleDirectory.String())

	if err != nil {
		return
	}
	_, err = g.Client().Post(ctx, prometheusAdminUrl.String())
	if err != nil {
		err = gerror.NewCode(gcode.CodeInvalidRequest)
	}

	return
}

// 生成规则文件
func generatedFile(ctx context.Context, ruleType, prometheusConfigFile, prometheusToolFile, prometheusRuleDirectory string) (err error) {

	results, err := g.DB().GetValue(ctx, "SELECT DISTINCT `group_name` FROM `rules` where type = ? and `active` = true", ruleType)
	if err != nil || len(results.Slice()) == 0 {
		return
	}
	for _, groupName := range results.Slice() {
		filePath := prometheusRuleDirectory + "/" + g.NewVar(groupName).String() + ".yml"
		err = parseRuleToFile(ctx, g.NewVar(groupName).String(), ruleType, filePath)
		if err != nil {
			return
		}
	}

	cmd := exec.Command(prometheusToolFile, "check", "config", prometheusConfigFile)
	var stdin, stdout, stderr bytes.Buffer
	cmd.Stdin = &stdin
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Run()

	outStr, errStr := stdout.String(), stderr.String()

	g.Log().Info(ctx, outStr)
	if gstr.Contains(errStr, "FAILED") {
		err = gerror.NewCode(gcode.New(1006, errStr, nil))
	}
	return
}

// 生成操作
func parseRuleToFile(ctx context.Context, groupName, ruleType, filePath string) (err error) {
	rules := []entity.Rules{}
	err = g.Model(entity.Rules{}).Where(model.Rules{GroupName: groupName, Type: ruleType, Active: true}).Scan(&rules)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	groups := []vo.Group{}

	for _, rule := range rules {
		alert := vo.AlertRule{}
		yaml.Unmarshal([]byte(rule.Content), &alert)
		group := vo.Group{
			Name:  g.NewVar(groupName).String(),
			Rules: g.Slice{alert},
		}
		groups = append(groups, group)
	}

	utility.GenerateRuleFile(gctx.New(), vo.Groups{Groups: groups}, filePath)
	return
}
