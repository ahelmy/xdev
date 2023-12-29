package internal

import (
	"fmt"
	"strings"
	"testing"
)


func TestGeneratePassword(t *testing.T) {
	// Test case 1: Generate password with length 8, including special characters and numeric characters
	password1 := GeneratePassword(8, true, true, false)
	fmt.Print(password1)
	if len(password1) != 8 || strings.ContainsAny(password1, upperLetterChars) || !strings.ContainsAny(password1, numberChars) || !strings.ContainsAny(password1, specialChars) {
		t.Errorf("Expected password length to be 8, including special characters and numeric characters. but got %d", len(password1))
	}

	// Test case 2: Generate password with length 12, including special characters, numeric characters, and capital letters
	password2 := GeneratePassword(10, true, true, true)
	fmt.Print(password2)
	if len(password2) != 10 || !strings.ContainsAny(password2, upperLetterChars) || !strings.ContainsAny(password2, numberChars) || !strings.ContainsAny(password2, specialChars) {
		t.Errorf("Expected password length to be 12, including special characters, numeric characters, and capital letters. but got %d", len(password2))
	}

	// Test case 3: Generate password with length 6, without special characters and numeric characters
	password3 := GeneratePassword(6, false, false, false)
	fmt.Print(password3)
	if len(password3) != 6 || strings.ContainsAny(password3, upperLetterChars) || strings.ContainsAny(password3, numberChars) || strings.ContainsAny(password3, specialChars)  {
		t.Errorf("Expected password length to be 6, without special characters and numeric characters or capital. but got %d", len(password3))
	}
}
