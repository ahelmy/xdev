package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompareTexts(t *testing.T) {
	tests := []struct {
		name     string
		txt1     string
		txt2     string
		expected string
	}{
		{
			name:     "Same texts",
			txt1:     "Hello, World",
			txt2:     "Hello, World",
			expected: "<span>Hello, World</span>",
		},
		{
			name:     "Texts with additional words",
			txt1:     "Hello World",
			txt2:     "Hello",
			expected: "<span>Hello</span><del style=\"background:#ffe6e6;\"> World</del>",
		},
		{
			name:     "Texts with less words",
			txt1:     "Hello",
			txt2:     "Hello World",
			expected: "<span>Hello</span><ins style=\"background:#e6ffe6;\"> World</ins>",
		},
		{
			name:     "Texts with different characters",
			txt1:     "Hello",
			txt2:     "Helli",
			expected: "<span>Hell</span><del style=\"background:#ffe6e6;\">o</del><ins style=\"background:#e6ffe6;\">i</ins>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CompareText(tt.txt1, tt.txt2)
			assert.Equal(t, tt.expected, result)
		})
	}
}
