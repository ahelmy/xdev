package internal

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"strings"
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
	err = writeProperties(&propertiesBuilder, "", yamlNode.Content[0])
	if err != nil {
		return "", err
	}

	return propertiesBuilder.String(), nil
}

func writeProperties(builder *strings.Builder, prefix string, node *yaml.Node) error {

	// invalid yaml
	if node.Kind != yaml.MappingNode {
		return fmt.Errorf("expected a MappingNode, got %v", node.Kind)
	}

	// This loop iterates over the key-value pairs in the YAML mapping.
	// It increments by 2 because each pair of nodes in the 'Content' slice represents a key-value pair.
	// 'i' is the index of the key node and 'i+1' is the index of the value node.
	for i := 0; i < len(node.Content); i += 2 {
		keyNode, valueNode := node.Content[i], node.Content[i+1]

		var fullKey string
		if prefix == "" {
			fullKey = keyNode.Value
		} else {
			fullKey = fmt.Sprintf("%v.%v", prefix, keyNode.Value)
		}

		switch valueNode.Kind {
		case yaml.MappingNode:
			// Recursively write nested maps
			err := writeProperties(builder, fullKey, valueNode)
			if err != nil {
				return err
			}
		case yaml.SequenceNode:
			// Handle lists
			for i, element := range valueNode.Content {
				_, err := fmt.Fprintf(builder, "%s[%d]=%v\n", fullKey, i, element.Value)
				if err != nil {
					return err
				}
			}

		default:
			// Write key-value pairs
			_, err := fmt.Fprintf(builder, "%s=%v\n", fullKey, valueNode.Value)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
