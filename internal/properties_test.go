package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProperties2Yaml(t *testing.T) {
	tests := []struct {
		name             string
		propertiesString string
		expectedYaml     string
		expectError      bool
	}{

		{
			name:             "Nested Map",
			propertiesString: "parent.child.grandchild=value\n",
			expectedYaml:     "parent:\n    child:\n        grandchild: value\n",
			expectError:      false,
		},
		{
			name:             "List",
			propertiesString: "fruits[0]=apple\nfruits[1]=banana\nfruits[2]=cherry\n",
			expectedYaml:     "fruits:\n    - apple\n    - banana\n    - cherry\n",
			expectError:      false,
		}, {
			name:             "Invalid Properties",
			propertiesString: "invalid: properties: string",
			expectedYaml:     "",
			expectError:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Properties2Yaml(tt.propertiesString)
			if (err != nil) != tt.expectError {
				t.Errorf("Unexpected error status. Got error: %v, expected error: %v", err, tt.expectError)
			}

			assert.Equal(t, tt.expectedYaml, result)
		})
	}
}
