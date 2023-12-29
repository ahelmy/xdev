package internal

import (
	"math/rand"
	"time"
)

const (
	letterChars   = "abcdefghijklmnopqrstuvwxyz"
	upperLetterChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberChars   = "0123456789"
	specialChars  = "!@#$%^&*()_+{}[]:;<>,.?/~`"
)

func GeneratePassword(length int, enableSpecialChars bool, enableNumeric bool, enableCapital bool) string {
	var chars string
	chars += letterChars

	if enableSpecialChars && len(specialChars) > 0 {
		chars += specialChars
	}

	if enableNumeric && len(numberChars) > 0 {
		chars += numberChars
	}

	if enableCapital && len(upperLetterChars) > 0 {
		chars += upperLetterChars
	}

	rand.NewSource(time.Now().UnixNano())
	password := make([]byte, length)

	// Ensure at least one special char, numeric, or capital in the password
	if enableSpecialChars && len(specialChars) > 0 {
		password[rand.Intn(length-1)] = specialChars[rand.Intn(len(specialChars)-1)]
	}

	if enableNumeric && len(numberChars) > 0 {
		password[rand.Intn(length-1)] = numberChars[rand.Intn(len(numberChars)-1)]
	}

	if enableCapital && len(upperLetterChars) > 0 {
		password[rand.Intn(length-1)] = upperLetterChars[rand.Intn(len(upperLetterChars)-1)]
	}

	for i := 0; i < length; i++ {
		if password[i] == 0 {
			password[i] = chars[rand.Intn(len(chars)-1)]
		}
	}

	return string(password)
}

