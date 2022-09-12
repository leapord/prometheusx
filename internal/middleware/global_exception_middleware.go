package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
)

func GlobalExceptionMiddleware(r *ghttp.Request) {
	r.Middleware.Next()
	err := r.GetError()
	// 记录到自定义错误日志文件
	if err == nil {
		return
	}

	// 重复主键处理
	if gstr.ContainsI(err.Error(), "Duplicate entry") {
		src := err.Error()
		errorString := gstr.SubStr(src, 0, gstr.PosI(src, ","))
		r.Response.ClearBuffer()
		r.Response.WriteJson(g.Map{
			"code":    405,
			"message": errorString,
		})
		return
	}
	// 找不到数据
	if gstr.ContainsI(err.Error(), "sql: no rows in result set") {
		r.Response.ClearBuffer()
		r.Response.WriteJson(g.Map{
			"code":    404,
			"message": "数据不存在，请检查查询条件",
		})
		return
	}

	// 全局统一错误处理
	if err != nil {
		//返回固定的友好信息
		r.Response.ClearBuffer()
		r.Response.WriteJson(g.Map{
			"code":    500,
			"message": "server busy, please retry after sometimes",
		})
	}
}
