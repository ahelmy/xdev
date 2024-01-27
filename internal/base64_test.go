package internal

import (
	"testing"
)

func TestDecodeFromBase64(t *testing.T) {
	// Test case 1: Valid base64 input
	input1 := "SGVsbG8gV29ybGQh"
	expected1 := "Hello World!"
	result1 := DecodeFromBase64(input1)
	if result1 != expected1 {
		t.Errorf("Expected decoded string to be %q, but got %q", expected1, result1)
	}

	// Test case 2: Invalid base64 input
	input2 := "InvalidBase64"
	expected2 := ""
	result2 := DecodeFromBase64(input2)
	if result2 != expected2 {
		t.Errorf("Expected decoded string to be %q, but got %q", expected2, result2)
	}
}

func TestEncodeToBase64(t *testing.T) {
	// Test case 1: Valid input
	input1 := "Hello World!"
	expected1 := "SGVsbG8gV29ybGQh"
	result1 := EncodeToBase64(input1)
	if result1 != expected1 {
		t.Errorf("Expected encoded string to be %q, but got %q", expected1, result1)
	}

	// Test case 2: Empty input
	input2 := ""
	expected2 := ""
	result2 := EncodeToBase64(input2)
	if result2 != expected2 {
		t.Errorf("Expected encoded string to be %q, but got %q", expected2, result2)
	}
}
