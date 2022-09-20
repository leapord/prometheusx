package vo

import "github.com/gogf/gf/v2/frame/g"

type AlertRule struct {
	Alert       string      `yaml:"alert" json:"alert"`
	Expr        string      `yaml:"expr" json:"expr" dc:"prometheus metrics expression"`
	For         string      `yaml:"for" json:"for" dc:""`
	Labels      g.MapStrAny `yaml:"labels" json:"lables" dc:"alert rule lables"`
	Annotations g.MapStrAny `yaml:"annotations" json:"annotations" dc:"alert rule annotations"`
}
