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

	// Test no expiry
	jwtTokenNoExp := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJBaGVsbXkifQ.AJhNuKaUM-TQ9erj2prelMnzhTPLZo6sOrTxYrzeDsU"
	jwt, err = DecodeJWT(jwtTokenNoExp)
	if err != nil || jwt.Expires != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test case 2: Invalid JWT
	invalidJwtToken := "invalid-token"
	_, err = DecodeJWT(invalidJwtToken)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	// Test case 3: Invalid claims
	invalidClaimsToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJiLCJleHOjE1N5NjQwMzB9.2ZQ5"
	_, err = DecodeJWT(invalidClaimsToken)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}
func TestEncodeJWT(t *testing.T) {
	algorithm := "HS256"
	claims := `{"sub": "Ahelmy", "exp": 1566964030}`
	signature := "secret-key"

	expectedToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjY5NjQwMzAsInN1YiI6IkFoZWxteSJ9.dvfEUt0hp0DFG4Wfxo78wXDFivHWT65paDnKD0-wNgI"

	token, err := EncodeJWT(algorithm, claims, signature)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if token != expectedToken {
		t.Errorf("Expected token to be '%s', but got '%s'", expectedToken, token)
	}

	_, err = EncodeJWT("xyz", claims, signature)
	if err == nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
