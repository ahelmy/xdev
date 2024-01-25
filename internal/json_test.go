package internal

import (
	"testing"
)
const (
	EXPECTED_ERROR = "Expected error '%s', but got:\n%s"
	UNEXPECTED_ERROR = "Unexpected error: %v"
)
func TestIndentJSON(t *testing.T) {
	input := `{"name":"John","age":30,"city":"New York"}`
	expectedOutput := `{
    "name": "John",
    "age": 30,
    "city": "New York"
}`

	output, err := IndentJSON(input)
	if err != nil {
		t.Errorf(UNEXPECTED_ERROR, err)
	}

	if output != expectedOutput {
		t.Errorf(EXPECTED_ERROR, expectedOutput, output)
	}

	// Test case 2: Invalid JSON input
	input2 := `{"name":"John","age":30,"city":"New York"`
	_, err = IndentJSON(input2)
	if err != nil && err.Error() != "unexpected end of JSON input" {
		t.Errorf(EXPECTED_ERROR, "unexpected end of JSON input", err)
	}
}

func TestMinifyJSON(t *testing.T) {
	input := `{
    "name": "John",
    "age": 30,
    "city": "New York"
}`
	expectedOutput := `{"name":"John","age":30,"city":"New York"}`

	output, err := MinifyJSON(input)
	if err != nil {
		t.Errorf(UNEXPECTED_ERROR, err)
	}

	if output != expectedOutput {
		t.Errorf(EXPECTED_ERROR, expectedOutput, output)
	}

	// Test case 2: Invalid JSON input
	input2 := `{"name":"John","age":30,"city":"New York"`
	_, err = MinifyJSON(input2)
	if err != nil && err.Error() != "unexpected end of JSON input" {
		t.Errorf(EXPECTED_ERROR, "unexpected end of JSON input", err)
	}
}

func TestJson2Yaml(t *testing.T) {
	input := `{"name":"John","age":30,"city":"New York"}`
	expectedOutput := `name: John
age: 30
city: New York
`

	output, err := Json2Yaml(input)
	if err != nil {
		t.Errorf(UNEXPECTED_ERROR, err)
	}

	if output != expectedOutput {
		t.Errorf(EXPECTED_ERROR, expectedOutput, output)
	}

	// Test case 2: Invalid JSON input
	input2 := `{"name":"John","age":30,"city":"New York"`
	_, err = Json2Yaml(input2)
	if err != nil && err.Error() != "unexpected EOF" {
		t.Errorf(EXPECTED_ERROR, "unexpected EOF", err)
	}
}