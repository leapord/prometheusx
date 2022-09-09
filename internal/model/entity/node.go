// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Node is the golang structure for table node.
type Node struct {
	Id         int64       `json:"id"         ` // 主键
	Host       string      `json:"host"       ` // 主机地址 IP或者域名
	Port       string      `json:"port"       ` // exporter对应的端口号
	Owner      string      `json:"owner"      ` // 所有者/责任人
	JobName    string      `json:"jobName"    ` // prometheus job_name
	Group      string      `json:"group"      ` // 组名
	Labels     string      `json:"labels"     ` // 标签 对应 prometheus中的label配置选项
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
}
