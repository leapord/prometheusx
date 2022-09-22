package main

import (
	"os"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/util/gvalid"
	"github.com/leapord/prometheusx/internal/cmd"
	_ "github.com/leapord/prometheusx/internal/event"
	_ "github.com/leapord/prometheusx/internal/logic"
	_ "github.com/leapord/prometheusx/internal/packed"
	"github.com/leapord/prometheusx/internal/validator"
)

func init() {
	rule := "yaml"
	gvalid.RegisterRule(rule, validator.RuleYamlContent)
}

func main() {
	ctx := gctx.New()
	if err := gfile.Remove("./static"); err != nil {
		g.Log().Error(ctx, err)
	}
	if err := os.Mkdir("./static", 0755); err != nil {
		g.Log().Error(ctx, err)
	}
	if !gres.IsEmpty() {
		g.Dump()
		if err := gres.Export("public/html", "static"); err != nil {
			g.Log().Error(ctx, err)
		}
	}
	cmd.Main.Run(ctx)
}
