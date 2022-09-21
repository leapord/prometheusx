package v1

import "github.com/gogf/gf/v2/frame/g"

// 新增
type ConfigAddReq struct {
	g.Meta `path:"/config/add" tags:"config" method:"put" summary:"add configuration"`
	Name   string `p:"name" v:"required" dc:"config name"`
	Value  string `p:"value" v:"required" dc:"config value"`
}

type ConfigAddRes struct {
	g.Meta `mime:"application/json"`
	Model  interface{} `json:"model"`
}

// 修改
type ConfigUpdateReq struct {
	g.Meta `path:"/config/update" tags:"config" method:"post" summary:"update configuration"`
	Id     int    `p:"id" v:"required" dc:"config item id"`
	Name   string `p:"name" v:"required" dc:"config name"`
	Value  string `p:"value" v:"required" dc:"config value"`
}

type ConfigUpdateRes struct {
	g.Meta `mime:"application/json"`
	Model  interface{} `json:"model"`
}

// 删除
type ConfigRemoveReq struct {
	g.Meta `path:"/config/remove/{id}" tags:"config" method:"delete" summary:"delete configuration"`
	Id     int `p:"id" v:"required" dc:"config item id"`
}

type ConfigRemoveRes struct {
	g.Meta `mime:"application/json"`
	Model  interface{} `json:"model"`
}

// 单个详情
type ConfigDetailReq struct {
	g.Meta `path:"/config/detail/{id}" tags:"config" method:"get" summary:"get configuration detail"`
	Id     int `p:"id" v:"required" dc:"config item id"`
}

type ConfigDetailRes struct {
	g.Meta `mime:"application/json"`
	Model  interface{} `json:"model"`
}

// 分页查询
type ConfigPageReq struct {
	g.Meta   `path:"/config/page" tags:"config" method:"post" summary:"get configuration with page"`
	Name     string `p:"name" dc:"config name"`
	PageNo   int    `p:"page" v:"min:1" d:"1"`
	PageSize int    `p:"pageSize" v:"min:1" d:"10"`
}

type ConfigPageRes struct {
	g.Meta   `mime:"application/json"`
	Models   interface{} `json:"rows"`
	Total    int         `p:"total" dc:"total record"`
	PageNo   int         `p:"page" dc:"current page"`
	PageSize int         `p:"pageSize" dc:"current page size"`
}

// 根据配置名查询
type ConfigNameReq struct {
	g.Meta `path:"/config/name/{name}" tags:"config" method:"get" summary:"get configuration with page"`
	Name   string `p:"name" v:"required" dc:"config name"`
}

type ConfigNameRes struct {
	g.Meta `mime:"application/json"`
	Model  interface{} `json:"model"`
}
