package utility

import (
	"context"

	"github.com/gogf/gf/v2/os/gfile"
	"gopkg.in/yaml.v3"
)

func GenerateRuleFile(ctx context.Context, ruleGroup interface{}, filePath string) (err error) {

	if err = gfile.Remove(filePath); err != nil {
		return
	}

	yamlBytes, err := yaml.Marshal(ruleGroup)
	if err != nil {
		return
	}

	file, err := gfile.Create(filePath)
	if err != nil {
		return
	}

	_, err = file.Write(yamlBytes)

	return
}
