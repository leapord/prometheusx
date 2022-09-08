package v1

import "github.com/gogf/gf/v2/frame/g"

// 添加分组
type GroupAddReq struct {
	g.Meta `path:"/group/add" tags:"group" method:"post" summary:"add prometheus group"`
	Name   string `p:"name" v:"required" dc:"group name"`
}

type GroupAddRes struct {
	g.Meta `mime:"text/html"`
	Model  interface{} `dc:"insert result , final entity"`
}

// 更新分组
type GroupUpdateReq struct {
	g.Meta `path:"/group/update" tags:"group" method:"put" summary:"add prometheus group"`
	Id     string `p:"id" v:"required" dc:"group id,key"`
	Name   string `p:"name" v:"required" dc:"group name"`
}

type GroupUpdateRes struct {
	g.Meta `mime:"text/html"`
	Model  interface{} `dc:"insert result , final entity"`
}

// 删除分组
type GroupDeleteReq struct {
	g.Meta `path:"/group/remove/{id}" tags:"group" method:"delete" summary:"remove the group"`
	Id     string `p:"id" v:"required" dc:"group id,key"`
}

type GroupDeleteRes struct {
	g.Meta `mime:"text/html"`
	Model  interface{} `dc:"insert result , final entity"`
}

// 查询单体详情
type GroupDetailReq struct {
	g.Meta `path:"/group/detail/{id}" tags:"group" method:"get" summary:"get the detail info"`
	Id     string `p:"id" v:"required" dc:"group id,key"`
}

type GroupDetailRes struct {
	g.Meta `mime:"text/html"`
	Model  interface{} `json:"model" dc:"group detail"`
}

//分页查询
type GroupPageReq struct {
	g.Meta   `path:"/group/page" tags:"group" method:"post" summary:"fetch group list by page"`
	PageNo   int    `p:"pageNo" d:"1"  v:"min:1#分页号码错误"`       // 分页号码
	PageSize int    `p:"pageSize" d:"10" v:"max:50#分页数量最大50条"` // 分页数量，最大50
	Name     string `p:"name"`
}

type GroupPageRes struct {
	g.Meta   `mime:"text/html"`
	Model    interface{} ` json:"model" dc:"list of query result"`
	Total    int         `json:"total" dc:"total count"`
	PageNo   int         `json:"pageNo" dc:"current page no"`
	PageSize int         `json:"pageSize" dc:"current page size"`
}

// 查询全部分组
type GroupListReq struct {
	g.Meta `path:"/group/list" tags:"group" method:"get" summary:"fetch group all list"`
}

type GroupListRes struct {
	g.Meta `mime:"text/html"`
	Model  interface{} `json:"models" dc:"entity"`
}
