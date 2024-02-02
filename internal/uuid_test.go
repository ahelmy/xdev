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

func TestToUpperCase(t *testing.T) {
	t.Run("Lowercase string", func(t *testing.T) {
		input := "hello world"
		expected := "HELLO WORLD"
		result := toUpperCase(input)
		if result != expected {
			t.Errorf("Expected result to be %s, but got %s", expected, result)
		}
	})

	t.Run("Uppercase string", func(t *testing.T) {
		input := "HELLO WORLD"
		expected := "HELLO WORLD"
		result := toUpperCase(input)
		if result != expected {
			t.Errorf("Expected result to be %s, but got %s", expected, result)
		}
	})

	t.Run("Mixed case string", func(t *testing.T) {
		input := "HeLlO WoRlD"
		expected := "HELLO WORLD"
		result := toUpperCase(input)
		if result != expected {
			t.Errorf("Expected result to be %s, but got %s", expected, result)
		}
	})
}

func TestCrockfordDecode(t *testing.T) {
	t.Run("Valid input", func(t *testing.T) {
		input := "A1B2C3D4E5F6G7H8I9J0K1L2M3N4O5P6Q7R8S9T"
		expected := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF}
		result, err := crockfordDecode(input)
		if err == nil {
			t.Errorf("Expected result length to be %d, but got %d", len(expected), len(result))
		}
		for i := 0; i < len(result); i++ {
			if result[i] != expected[i] {
				t.Errorf("Expected result[%d] to be %d, but got %d", i, expected[i], result[i])
			}
		}
	})

	t.Run("Invalid input", func(t *testing.T) {
		input := "A1B2C3D4E5F6G7H8I9J0K1L2M3N4O5P6Q7R8S9T!"
		_, err := crockfordDecode(input)
		if err == nil {
			t.Errorf("Expected error, but got nil")
		}
	})

}
