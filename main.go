package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gvalid"
	"github.com/leapord/prometheusx/internal/cmd"
	_ "github.com/leapord/prometheusx/internal/logic"
	_ "github.com/leapord/prometheusx/internal/packed"
	"github.com/leapord/prometheusx/internal/validator"
)

func init() {
	rule := "yaml"
	gvalid.RegisterRule(rule, validator.RuleYamlContent)
}

func main() {
	cmd.Main.Run(gctx.New())
}
