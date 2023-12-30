package internal

import (
	"testing"
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
		t.Errorf("Unexpected error: %v", err)
	}

	if output != expectedOutput {
		t.Errorf("Expected output:\n%s\n\nBut got:\n%s", expectedOutput, output)
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
		t.Errorf("Unexpected error: %v", err)
	}

	if output != expectedOutput {
		t.Errorf("Expected output:\n%s\n\nBut got:\n%s", expectedOutput, output)
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
		t.Errorf("Unexpected error: %v", err)
	}

	if output != expectedOutput {
		t.Errorf("Expected output:\n%s\n\nBut got:\n%s", expectedOutput, output)
	}
}
