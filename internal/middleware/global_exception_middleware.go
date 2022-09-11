package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
)

func GlobalExceptionMiddleware(r *ghttp.Request) {
	r.Middleware.Next()
	err := r.GetError()
	// 记录到自定义错误日志文件
	if err != nil {
		g.Log().Error(gctx.New(), err)
	} else {
		return
	}

	// 重复主键处理
	if gstr.ContainsI(err.Error(), "Duplicate entry") {
		src := err.Error()
		errorString := gstr.SubStr(src, 0, gstr.PosI(src, ","))
		r.Response.ClearBuffer()
		r.Response.Writeln(errorString)
		return
	}

	// 全局统一错误处理
	if err != nil {
		//返回固定的友好信息
		r.Response.ClearBuffer()
		r.Response.Writeln("服务器居然开小差了，请稍后再试吧！")
	}
}
