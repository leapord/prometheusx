package main

import (
	_ "prometheus_ext/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"prometheus_ext/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
