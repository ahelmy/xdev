package internal

import (
	"encoding/json"
	"fmt"
	"strings"

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

func Yaml2Properties(yamlData string) (string, error) {
	// Parse YAML into a Node
	var yamlNode yaml.Node
	err := yaml.Unmarshal([]byte(yamlData), &yamlNode)
	if err != nil {
		return "", err
	}

	// Write Properties
	var propertiesBuilder strings.Builder
	writeProperties(&propertiesBuilder, "", yamlNode.Content[0])

	return propertiesBuilder.String(), nil
}

func writeProperties(builder *strings.Builder, prefix string, node *yaml.Node) {
	if node.Kind != yaml.MappingNode {
		return
	}

	for i := 0; i < len(node.Content); i += 2 {
		keyNode, valueNode := node.Content[i], node.Content[i+1]

		fullKey := getFullKey(prefix, keyNode.Value)

		switch valueNode.Kind {
		case yaml.MappingNode:
			writeNestedMaps(builder, fullKey, valueNode)
		case yaml.SequenceNode:
			writeSequence(builder, fullKey, valueNode)
		default:
			writeKeyValue(builder, fullKey, valueNode.Value)
		}
	}
}

func getFullKey(prefix, key string) string {
	if prefix == "" {
		return key
	}
	return fmt.Sprintf("%v.%v", prefix, key)
}

func writeNestedMaps(builder *strings.Builder, prefix string, node *yaml.Node) {
	writeProperties(builder, prefix, node)
}

func writeSequence(builder *strings.Builder, prefix string, node *yaml.Node) {
	for i, element := range node.Content {
		builder.WriteString(fmt.Sprintf("%s[%d]=%v\n", prefix, i, element.Value))
	}
}

func writeKeyValue(builder *strings.Builder, key, value string) {
	builder.WriteString(fmt.Sprintf("%s=%v\n", key, value))
}
