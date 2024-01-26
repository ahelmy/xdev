package internal

import (
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

func Properties2Yaml(propertiesData string) (string, error) {
	// Parse properties
	dataMap := make(map[string]interface{})
	lines := strings.Split(propertiesData, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return "", fmt.Errorf("invalid properties line: %s", line)
		}
		key, value := parts[0], parts[1]
		err := addToMap(dataMap, key, value)
		if err != nil {
			return "", err
		}
	}

	// Convert map to YAML
	yamlData, err := yaml.Marshal(dataMap)
	if err != nil {
		return "", err
	}

	return string(yamlData), nil
}

func addToMap(dataMap map[string]interface{}, key string, value string) error {
	parts := strings.Split(key, ".")
	lastMap := dataMap
	for i, part := range parts {
		if strings.HasSuffix(part, "]") {
			err := handleArrayKey(lastMap, part, value)
			if err != nil {
				return err
			}
		} else if i == len(parts)-1 {
			handleNormalKey(lastMap, part, value)
		} else {
			lastMap = handleNestedKey(lastMap, part)
		}
	}
	return nil
}

// handle the logic when the properties key is an array example -> child[0]=value
func handleArrayKey(lastMap map[string]interface{}, part string, value string) error {
	// extra the name of the array example name[0] --> name
	arrayKey := part[:strings.Index(part, "[")]
	// extra the index of the array example name[0] --> 0
	indexStr := part[strings.Index(part, "[")+1 : len(part)-1]
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		return err
	}

	if _, ok := lastMap[arrayKey]; !ok {
		lastMap[arrayKey] = make([]interface{}, index+1)
	} else if len(lastMap[arrayKey].([]interface{})) <= index {
		oldArray := lastMap[arrayKey].([]interface{})
		newArray := make([]interface{}, index+1)
		copy(newArray, oldArray)
		lastMap[arrayKey] = newArray
	}

	lastMap[arrayKey].([]interface{})[index] = value
	return nil
}

// handle the logic when the properties key is a normal key example -> root=value
func handleNormalKey(lastMap map[string]interface{}, part string, value string) {
	lastMap[part] = value
}

// handle the logic when the properties key is a nested key example -> root.child=value
func handleNestedKey(lastMap map[string]interface{}, part string) map[string]interface{} {
	if _, ok := lastMap[part]; !ok {
		lastMap[part] = make(map[string]interface{})
	}
	return lastMap[part].(map[string]interface{})
}
