package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYaml2JSON(t *testing.T) {
	tests := []struct {
		name         string
		yamlString   string
		expectedJSON string
		expectError  bool
	}{
		{
			name:         "Valid YAML",
			yamlString:   "key: value\nnested:\n  subkey: subvalue",
			expectedJSON: `{"key":"value","nested":{"subkey":"subvalue"}}`,
			expectError:  false,
		},
		{
			name:         "Invalid YAML",
			yamlString:   "invalid: yaml: string",
			expectedJSON: "{}",
			expectError:  true,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Yaml2Json(tt.yamlString)
			if (err != nil) != tt.expectError {
				t.Errorf("Unexpected error status. Got error: %v, expected error: %v", err, tt.expectError)
			}

			// Use assert.JSONEq to compare JSON strings while ignoring key order
			assert.JSONEq(t, tt.expectedJSON, result)
		})
	}
}
