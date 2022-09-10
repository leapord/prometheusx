package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/leapord/prometheusx/internal/cmd"
	_ "github.com/leapord/prometheusx/internal/packed"
)

func main() {
	cmd.Main.Run(gctx.New())
}
