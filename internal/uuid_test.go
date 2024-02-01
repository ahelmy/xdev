package internal

import (
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	guid := GenerateUUID()
	if len(guid) != 36 {
		t.Errorf("Expected GUID length to be 36, but got %d", len(guid))
	}
}

func TestUUIDtoULID(t *testing.T) {
	t.Run("Valid UUID", func(t *testing.T) {
		uuid := "860d3040-1e23-416f-a1fd-f4ccc773833d"
		expectedULID := "461MR407H385QT3ZFMSK3Q70SX"
		ulid, err := UUIDtoULID(uuid)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if ulid != expectedULID {
			t.Errorf("Expected ULID to be %s, but got %s", expectedULID, ulid)
		}
	})
	t.Run("Invalid UUID", func(t *testing.T) {
		uuid := "860d3040-1e23-416f-a1fd-f4ccc773833"
		_, err := UUIDtoULID(uuid)
		if err == nil {
			t.Errorf("Expected error, but got nil")
		}
	})
}
