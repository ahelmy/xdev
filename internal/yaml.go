package internal

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

func Yaml2Json(yamlString string) (string, error) {
	t := make(map[string]any)
	err := yaml.Unmarshal([]byte(yamlString), &t)
	if err != nil {
		return "{}", err
	}

	output, err := json.Marshal(t)
	if err != nil {
		return "{}", err
	}

	return string(output), nil
}
