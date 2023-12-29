package internal

import (
	"testing"
)

func TestDecodeJWT(t *testing.T) {
	jwtToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJBaGVsbXkiLCJleHAiOjE1NjY5NjQwMzB9.2ZQ5"

	jwt, err := DecodeJWT(jwtToken)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Add assertions to validate the decoded JWT
	if jwt.Header != "{\n    \"alg\": \"HS256\",\n    \"typ\": \"JWT\"\n}" {
		t.Errorf("Expected header to be 'expected-header', but got '%s'", jwt.Header)
	}

	if jwt.Claims != "{\n    \"exp\": 1566964030,\n    \"sub\": \"Ahelmy\"\n}" {
		t.Errorf("Expected claims to be 'expected-claims', but got '%s'", jwt.Claims)
	}
}
