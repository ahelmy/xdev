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

	output, _ := json.Marshal(t)
	json, err := IndentJSON(string(output))
	return json, err
}
