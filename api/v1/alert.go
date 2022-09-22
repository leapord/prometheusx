package v1

import "github.com/gogf/gf/v2/frame/g"

// webhook告警推送
type AlertHookReq struct {
	g.Meta `path:"/alert/webhook" tags:"alert" method:"post" summary:"webhook for alertmanager"`
}
type AlertHookRes struct {
	g.Meta `mime:"application/json"`
}

// 分页查询告警信息
type AlertPageReq struct {
	g.Meta   `path:"/alert/page" tags:"alert" method:"post" summary:"fetch alert list by page"`
	PageNo   int    `p:"page" d:"1"  v:"min:1#分页号码错误"`         // 分页号码
	PageSize int    `p:"pageSize" d:"10" v:"max:50#分页数量最大50条"` // 分页数量，最大50
	Labels   string `p:"labels" dc:"alert labels"`
}

type AlertPageRes struct {
	g.Meta   `mime:"text/html"`
	Models   interface{} `json:"rows" dc:"list of query result"`
	Total    int         `json:"total" dc:"total count"`
	PageNo   int         `json:"page" dc:"current page no"`
	PageSize int         `json:"pageSize" dc:"current page size"`
}
