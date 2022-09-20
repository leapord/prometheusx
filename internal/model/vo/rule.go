package vo

import "github.com/gogf/gf/v2/frame/g"

type Group struct {
	Name  string  `yaml:"name" json:"name" dc:"group name"`
	Rules g.Slice `yaml:"rules" json:"rules" dc:"rules"`
}

type Groups struct {
	Groups []Group `yaml:"groups" json:"groups"`
}
