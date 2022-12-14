// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Alert is the golang structure of table alert for DAO operations like Where/Data.
type Alert struct {
	g.Meta       `orm:"table:alert, do:true"`
	Id           interface{} // 主键
	GroupKey     interface{} // key identifying the group of alerts (e.g. to deduplicate)
	Status       interface{} // resolved|firing
	ExternalUrl  interface{} // alertmanager连接
	Labels       interface{} // 告警labels
	Annotations  interface{} // 告警annotations
	StartsAt     *gtime.Time // 告警开始时间
	EndsAt       *gtime.Time // 告警结束时间
	GeneratorUrl interface{} // 触发告警连接
	Fingerprint  interface{} // 告警指纹
	CreateTime   *gtime.Time // 创建时间
}
