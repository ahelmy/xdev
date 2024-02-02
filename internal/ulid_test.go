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

func TestULIDtoUUID(t *testing.T) {
	t.Run("Valid ULID", func(t *testing.T) {
		ulid := "461MR407H385QT3ZFMSK3Q70SX"
		expectedUUID := "860d3040-1e23-416f-a1fd-f4ccc773833d"
		uuid, err := ULIDtoUUID(ulid)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if uuid != expectedUUID {
			t.Errorf("Expected UUID to be %s, but got %s", expectedUUID, uuid)
		}
	})
	t.Run("Invalid ULID", func(t *testing.T) {
		ulid := "01F7Z6T2T0Z2VXY6J8X1Z5Q3"
		_, err := ULIDtoUUID(ulid)
		if err == nil {
			t.Errorf("Expected error, but got nil")
		}
	})
}
