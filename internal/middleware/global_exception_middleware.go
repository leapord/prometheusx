package middleware

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func GlobalExceptionMiddleware(r *ghttp.Request) {
	r.Middleware.Next()
	err := r.GetError()
	// 记录到自定义错误日志文件
	if err == nil {
		return
	}
	code := gerror.Code(err)

	r.Response.ClearBuffer()
	r.Response.WriteJson(g.Map{
		"code":    code.Code(),
		"message": code.Message(),
	})
}
