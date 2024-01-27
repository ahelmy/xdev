package internal

import (
	"testing"
)

func TestEncodeURL(t *testing.T) {
	// Test case 1: Encode a URL with special characters
	input1 := "https://example.com/?q=hello world"
	expected1 := "https%3A%2F%2Fexample.com%2F%3Fq%3Dhello+world"
	output1 := EncodeURL(input1)
	if output1 != expected1 {
		t.Errorf("Expected encoded URL to be %s, but got %s", expected1, output1)
	}

	// Test case 2: Encode a URL without special characters
	input2 := "https://example.com/"
	expected2 := "https%3A%2F%2Fexample.com%2F"
	output2 := EncodeURL(input2)
	if output2 != expected2 {
		t.Errorf("Expected encoded URL to be %s, but got %s", expected2, output2)
	}

	// Test case 3: Encode an empty URL
	input3 := ""
	expected3 := ""
	output3 := EncodeURL(input3)
	if output3 != expected3 {
		t.Errorf("Expected encoded URL to be %s, but got %s", expected3, output3)
	}
}

func TestDecodeURL(t *testing.T) {
	// Test case 1: Decode a URL with special characters
	input1 := "https%3A%2F%2Fexample.com%2F%3Fq%3Dhello+world"
	expected1 := "https://example.com/?q=hello world"
	output1, err1 := DecodeURL(input1)
	if err1 != nil {
		t.Errorf("Unexpected error: %v", err1)
	}
	if output1 != expected1 {
		t.Errorf("Expected decoded URL to be %s, but got %s", expected1, output1)
	}

	// Test case 2: Decode a URL without special characters
	input2 := "https%3A%2F%2Fexample.com%2F"
	expected2 := "https://example.com/"
	output2, err2 := DecodeURL(input2)
	if err2 != nil {
		t.Errorf("Unexpected error: %v", err2)
	}
	if output2 != expected2 {
		t.Errorf("Expected decoded URL to be %s, but got %s", expected2, output2)
	}

	// Test case 3: Decode an empty URL
	input3 := ""
	expected3 := ""
	output3, err3 := DecodeURL(input3)
	if err3 != nil {
		t.Errorf("Unexpected error: %v", err3)
	}
	if output3 != expected3 {
		t.Errorf("Expected decoded URL to be %s, but got %s", expected3, output3)
	}
}
