package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"github.com/leapord/prometheusx/internal/controller"
	"github.com/leapord/prometheusx/internal/middleware"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					ghttp.MiddlewareHandlerResponse,
					middleware.GlobalExceptionMiddleware,
				)
				group.Bind(
					controller.Hello,
					controller.Authentication,
				)
			})
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(
					ghttp.MiddlewareHandlerResponse,
					middleware.TokenMiddleware,
					middleware.GlobalExceptionMiddleware,
				)
				group.Bind(
					controller.Group,
					controller.Node,
				)
			})
			s.Run()
			return nil
		},
	}
)
