package internal

import (
	"testing"
)

func TestDecodeURL(t *testing.T) {
	// Test case 1: Decode a URL-encoded parameter
	param1 := "Hello%20World%21"
	expected1 := "Hello World!"
	result1, err1 := DecodeURL(param1)
	if err1 != nil {
		t.Errorf("Unexpected error: %v", err1)
	}
	if result1 != expected1 {
		t.Errorf("Expected decoded parameter to be %q, but got %q", expected1, result1)
	}

	// Test case 2: Decode an empty parameter
	param2 := ""
	expected2 := ""
	result2, err2 := DecodeURL(param2)
	if err2 != nil {
		t.Errorf("Unexpected error: %v", err2)
	}
	if result2 != expected2 {
		t.Errorf("Expected decoded parameter to be %q, but got %q", expected2, result2)
	}

	// Test case 3: Decode a parameter with special characters
	param3 := "%24%25%26"
	expected3 := "$%&"
	result3, err3 := DecodeURL(param3)
	if err3 != nil {
		t.Errorf("Unexpected error: %v", err3)
	}
	if result3 != expected3 {
		t.Errorf("Expected decoded parameter to be %q, but got %q", expected3, result3)
	}
}

func TestEncodeURL(t *testing.T) {
	// Test case 1: Encode a URL parameter
	param1 := "Hello World!"
	expected1 := "Hello%20World%21"
	result1 := EncodeURL(param1)
	if result1 != expected1 {
		t.Errorf("Expected encoded parameter to be %q, but got %q", expected1, result1)
	}

	// Test case 2: Encode an empty parameter
	param2 := ""
	expected2 := ""
	result2 := EncodeURL(param2)
	if result2 != expected2 {
		t.Errorf("Expected encoded parameter to be %q, but got %q", expected2, result2)
	}

	// Test case 3: Encode a parameter with special characters
	param3 := "$%&"
	expected3 := "%24%25%26"
	result3 := EncodeURL(param3)
	if result3 != expected3 {
		t.Errorf("Expected encoded parameter to be %q, but got %q", expected3, result3)
	}
}