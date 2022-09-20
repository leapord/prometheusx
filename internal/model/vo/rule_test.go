package vo

import (
	"fmt"
	"testing"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/leapord/prometheusx/utility"
	"gopkg.in/yaml.v3"
)

func TestRule(t *testing.T) {

	alert := AlertRule{
		Alert: "alert Test",
		Expr:  "up == 0",
		For:   "10m",
		Labels: g.MapStrAny{
			"owner": "leapord",
		},
		Annotations: g.MapStrAny{
			"summary": "${value}",
		},
	}

	group := Group{
		Name:  "test",
		Rules: g.Slice{alert},
	}

	groups := Groups{
		Groups: []Group{group},
	}

	jsonStr, _ := gjson.EncodeString(groups)
	g.Dump(jsonStr)
	yamlBytes, _ := yaml.Marshal(groups)
	fmt.Print(g.NewVar(yamlBytes).String())

}

func TestRuleToFile(t *testing.T) {
	alert := AlertRule{
		Alert: "alert Test",
		Expr:  "up == 0",
		For:   "10m",
		Labels: g.MapStrAny{
			"owner": "leapord",
		},
		Annotations: g.MapStrAny{
			"summary": "${value}",
		},
	}

	group := Group{
		Name:  "test",
		Rules: g.Slice{alert},
	}

	groups := Groups{
		Groups: []Group{group},
	}

	utility.GenerateRuleFile(gctx.New(), groups, "d:/test.yml")
}
