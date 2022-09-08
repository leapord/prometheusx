package v1

import "github.com/gogf/gf/v2/frame/g"

/*
 * 添加
 */
type NodeAddReq struct {
	g.Meta `path:"/node/add" tags:"node" method:"put" summary:"add node"`
	Host   string                   `p:"host" v:"required" dc:"node host or ip"`
	Port   string                   `p:"port" v:"required" dc:"node exporter port"`
	Owner  string                   `p:"owner" v:"required" dc:"this job record owner"`
	Group  string                   `p:"group" v:"required" dc:"group name"`
	Labels []map[string]interface{} `p:"labels" dc:"extend pair values for job"`
}
type NodeAddRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
