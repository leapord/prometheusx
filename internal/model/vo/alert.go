package vo

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type Alert struct {
	Id           int64       `json:"id"           ` // 主键
	GroupKey     string      `json:"groupKey"     ` // key identifying the group of alerts (e.g. to deduplicate)
	Status       string      `json:"status"       ` // resolved|firing
	ExternalUrl  string      `json:"externalUrl"  ` // alertmanager连接
	Labels       g.MapStrAny `json:"labels"       ` // 告警labels
	Annotations  g.MapStrAny `json:"annotations"  ` // 告警annotations
	StartsAt     *gtime.Time `json:"startsAt"     ` // 告警开始时间
	EndsAt       *gtime.Time `json:"endsAt"       ` // 告警结束时间
	GeneratorUrl string      `json:"generatorUrl" ` // 触发告警连接
	Fingerprint  string      `json:"fingerprint"  ` // 告警指纹
	CreateTime   *gtime.Time `json:"createTime"   ` // 创建时间
}
