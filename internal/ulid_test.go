package internal

import (
	"testing"
)

func TestGenerateULID(t *testing.T) {
	ulid := GenerateULID()
	if len(ulid) != 26 {
		t.Errorf("Expected ULID length to be 26, but got %d", len(ulid))
	}
}
