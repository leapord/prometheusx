package v1

import "github.com/gogf/gf/v2/frame/g"

// 新增
type RuleAddReq struct {
	g.Meta    `path:"/rule/add" tags:"Rule" method:"put" summary:"add Rule"`
	GroupName string `p:"groupName" v:"required" dc:"Rule group name"`
	Type      string `p:"type" v:"required" dc:"Rule type alert or record"`
	Content   string `p:"content" v:"required|yaml"`
}

type RuleAddRes struct {
	g.Meta `mime:"application/json"`
	Model  interface{} `json:"model"`
}

// 修改
type RuleUpdateReq struct {
	g.Meta    `path:"/rule/update" tags:"Rule" method:"post" summary:"update Rule"`
	Id        int    `p:"id" v:"required|min:1" dc:"Rule item id"`
	GroupName string `p:"groupName" v:"required" dc:"Rule group name"`
	Type      string `p:"type" v:"required" dc:"Rule type alert or record"`
	Content   string `p:"content" v:"required|yaml"`
	Active    bool   `p:"active" v:"required|boolean" dc:"active or not the default is true"`
}

type RuleUpdateRes struct {
	g.Meta `mime:"application/json"`
	Model  interface{} `json:"model"`
}

// 删除
type RuleRemoveReq struct {
	g.Meta `path:"/rule/remove/{id}" tags:"Rule" method:"delete" summary:"delete Rule"`
	Id     int `p:"id" v:"required" dc:"Rule item id"`
}

type RuleRemoveRes struct {
	g.Meta `mime:"application/json"`
	Model  interface{} `json:"model"`
}

// 单个详情
type RuleDetailReq struct {
	g.Meta `path:"/rule/detail/{id}" tags:"Rule" method:"get" summary:"get Rule detail"`
	Id     int `p:"id" v:"required" dc:"Rule item id"`
}

type RuleDetailRes struct {
	g.Meta `mime:"application/json"`
	Model  interface{} `json:"model"`
}

// 分页查询
type RulePageReq struct {
	g.Meta    `path:"/rule/page" tags:"Rule" method:"post" summary:"get rules with page"`
	GroupName string `p:"groupName" dc:"Rule group name"`
	Type      string `p:"type" dc:"Rule type alert or record"`
	PageNo    int    `p:"page" v:"min:1" d:"1"`
	PageSize  int    `p:"pageSize" v:"min:1" d:"10"`
}

type RulePageRes struct {
	g.Meta   `mime:"application/json"`
	Models   interface{} `json:"rows"`
	Total    int         `json:"total" dc:"total record"`
	PageNo   int         `json:"page" dc:"current page"`
	PageSize int         `json:"pageSize" dc:"current page size"`
}

// 生成规则文件
type RuleFileGeneratedReq struct {
	g.Meta `path:"/rule/generatedFile" tags:"Rule" method:"get" summary:"generated rule files , if rules directory path not set this will do nothing"`
}

type RuleFileGeneratedRes struct {
	g.Meta `mime:"application/json"`
}
