package internal

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
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

func TestYaml2Properties(t *testing.T) {
	tests := []struct {
		name             string
		yamlString       string
		expectedProperty string
		expectError      bool
	}{
		{
			name:             "Valid YAML",
			yamlString:       "key: value\nnested:\n  subkey: subvalue",
			expectedProperty: "key=value\nnested.subkey=subvalue\n",
			expectError:      false,
		},
		{
			name:             "Invalid YAML",
			yamlString:       "invalid: yaml: string",
			expectedProperty: "",
			expectError:      true,
		},
		{
			name:             "Nested Map",
			yamlString:       "parent:\n  child:\n    grandchild: value",
			expectedProperty: "parent.child.grandchild=value\n",
			expectError:      false,
		},
		{
			name:             "List",
			yamlString:       "fruits:\n  - apple\n  - banana\n  - cherry",
			expectedProperty: "fruits[0]=apple\nfruits[1]=banana\nfruits[2]=cherry\n",
			expectError:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Yaml2Properties(tt.yamlString)
			if (err != nil) != tt.expectError {
				t.Errorf("Unexpected error status. Got error: %v, expected error: %v", err, tt.expectError)
			}

			assert.Equal(t, tt.expectedProperty, result)
		})
	}

	t.Run("write properties - invalid kind", func(t *testing.T) {
		builder := &strings.Builder{}
		writeProperties(builder, "xyz", &yaml.Node{Kind: yaml.DocumentNode})
		assert.Equal(t, "", builder.String())
	})
}
