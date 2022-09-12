package v1

import "github.com/gogf/gf/v2/frame/g"

/*
 * 添加
 */
type NodeAddReq struct {
	g.Meta  `path:"/node/add" tags:"node" method:"put" summary:"add node"`
	Host    string      `p:"host" v:"required|ip" dc:"node host or ip"`
	Port    string      `p:"port" v:"required|min:1|max:65535" dc:"node exporter port"`
	Owner   string      `p:"owner" v:"required" dc:"this job record owner"`
	Group   string      `p:"group" v:"required" dc:"group name"`
	JobName string      `p:"jobName" v:"required" dc:"job name"`
	Labels  g.MapAnyAny `p:"labels" dc:"extend pair values for job"`
}
type NodeAddRes struct {
	g.Meta `mime:"application/json"`
	Model  interface{} `json:"model" dc:"add result"`
}

/*
 * 更新
 */
type NodeUpdateReq struct {
	g.Meta  `path:"/node/update" tags:"node" method:"post" summary:"update node"`
	Id      string      `p:"id" v:"required" dc:"node record id"`
	Host    string      `p:"host" v:"required|ip" dc:"node host or ip"`
	Port    string      `p:"port" v:"required|min:1|max:65535" dc:"node exporter port"`
	Owner   string      `p:"owner" v:"required" dc:"this job record owner"`
	Group   string      `p:"group" v:"required" dc:"group name"`
	JobName string      `p:"jobName" v:"required" dc:"job name"`
	Labels  g.MapAnyAny `p:"labels" dc:"extend pair values for job"`
}

type NodeUpdateRes struct {
	g.Meta `mime:"application/json"`
	Model  interface{} `dc:"the update result"`
}

type NodeRemoveReq struct {
	g.Meta `path:"/node/remove/{id}" tags:"node" method:"delete" summary:"delete node"`
	Id     string `p:"id" v:"required" dc:"node record id"`
}

type NodeRemoveRes struct {
	g.Meta `mime:"application/json"`
	Model  interface{} `json:"model" dc:"the delete result"`
}

type NodeDetailReq struct {
	g.Meta `path:"/node/detail/{id}" tags:"node" method:"get" summary:"delete node"`
	Id     string `p:"id" v:"required" dc:"node record id"`
}
type NodeDetailRes struct {
	g.Meta `mime:"application/json"`
	Model  interface{} `json:"model" dc:"the delete result"`
}
type NodePageReq struct {
	g.Meta   `path:"/node/page" tags:"node" method:"post" summary:"page node query"`
	PageNo   int    `p:"pageNo" d:"1" v:"min:1#can not lower than 1" dc:"page number"`
	PageSize int    `p:"pageSize" d:"10" v:"max:100" dc:"page size of each page"`
	Host     string `p:"host" dc:"node host or ip"`
	Port     string `p:"port" dc:"node exporter port"`
	Owner    string `p:"owner" dc:"this job record owner"`
	Group    string `p:"group" dc:"group name"`
	JobName  string `p:"jobName" dc:"job name"`
}

type NodePageRes struct {
	g.Meta   `mime:"application/json"`
	PageNo   int `json:"pageNo" dc:"current page number "`
	PageSize int `json:"pageSize" dc:"current page size"`
	Total    int `json:"total" dc:"total record for this query condition"`
	Models   interface{}
}

type NodeActiveReq struct {
	g.Meta `path:"/node/active" tags:"node" method:"post" summary:"node active status change"`
	Id     int  `p:"id" v:"required" dc:"Node record id"`
	Active bool `p:"active" v:"required|boolean" dc:"Node active status true or false"`
}

type NodeActiveRes struct {
	g.Meta `mime:"application/json"`
}
