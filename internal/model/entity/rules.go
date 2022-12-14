// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Rules is the golang structure for table rules.
type Rules struct {
	Id         int64       `json:"id"         ` // 主键
	GroupName  string      `json:"groupName"  ` // 规则组名称
	Type       string      `json:"type"       ` // 规则类型 alert record
	Content    string      `json:"content"    ` // 规则内容
	Active     int         `json:"active"     ` // 是否启用
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
}
