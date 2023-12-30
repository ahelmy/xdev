package internal

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/itchyny/json2yaml"
)

func IndentJSON(input string) (string, error) {
	var indentedBuffer bytes.Buffer

	// Use json.Indent to format the JSON string with indentation
	err := json.Indent(&indentedBuffer, []byte(input), "", "    ")
	if err != nil {
		return "", err
	}

	return indentedBuffer.String(), nil
}

func MinifyJSON(input string) (string, error) {
	var minifiedBuffer bytes.Buffer

	// Use json.Compact to remove all whitespace from the JSON string
	err := json.Compact(&minifiedBuffer, []byte(input))
	if err != nil {
		return "", err
	}

	return minifiedBuffer.String(), nil
}

func Json2Yaml(json string) (string, error) {
	input := strings.NewReader(json)
	var output strings.Builder
	if err := json2yaml.Convert(&output, input); err != nil {
		return "", err
	}
	return output.String(), nil
}
