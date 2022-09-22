package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"github.com/gogf/gf/v2/os/gfile"
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
			if gfile.Exists("./static/public/html/") {
				s.SetIndexFolder(true)
				s.SetServerRoot("./static/public/html/")
			} else {
				g.Log().Error(ctx, "./static/public/html/ directory is not exist")
			}
			// 无需权限
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					ghttp.MiddlewareHandlerResponse,
					middleware.GlobalExceptionMiddleware,
				)
				// 无需权限
				group.Group("/api", func(group *ghttp.RouterGroup) {
					group.Bind(
						controller.Hello,
						controller.Authentication,
						controller.Target,
						controller.AlertWebhook,
					)
				})
				// 权限验证
				group.Group("/api", func(group *ghttp.RouterGroup) {
					group.Middleware(
						middleware.TokenMiddleware,
					)
					group.Bind(
						controller.Group,
						controller.Node,
						controller.User,
						controller.Rules,
						controller.Config,
						controller.Alert,
					)
				})
			})

			s.Run()
			return nil
		},
	}
)
