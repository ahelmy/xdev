package internal

import (
	"testing"
)

func TestGenerateGUID(t *testing.T) {
	guid := GenerateGUID()
	if len(guid) != 36 {
		t.Errorf("Expected GUID length to be 36, but got %d", len(guid))
	}
}
