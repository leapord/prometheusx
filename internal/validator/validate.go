package validator

import (
	"context"

	"github.com/gogf/gf/v2/encoding/gyaml"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gvalid"
)

func RuleYamlContent(ctx context.Context, in gvalid.RuleFuncInput) error {
	ruleContent := in.Data.String()
	_, err := gyaml.Decode([]byte(ruleContent))
	if err != nil {
		err = gerror.WrapCode(gcode.CodeBusinessValidationFailed, err, "yaml content valid error")
	}
	return err
}
