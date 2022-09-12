// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Node is the golang structure of table node for DAO operations like Where/Data.
type Node struct {
	g.Meta     `orm:"table:node, do:true"`
	Id         interface{} // 主键
	Host       interface{} // 主机地址 IP或者域名
	Port       interface{} // exporter对应的端口号
	Owner      interface{} // 所有者/责任人
	JobName    interface{} // prometheus job_name
	Group      interface{} // 组名
	Labels     interface{} // 标签 对应 prometheus中的label配置选项
	CreateTime *gtime.Time // 创建时间
	Active     interface{} // 是否启用
}
