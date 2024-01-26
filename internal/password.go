package internal

import (
	"math/rand"
	"time"
)

const (
	letterChars      = "abcdefghijklmnopqrstuvwxyz"
	upperLetterChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberChars      = "0123456789"
	specialChars     = "!@#$%^&*()_+{}[]:;<>,.?/~`"
	MaxLength        = 2048
)

func getCharacters(enableCapital bool, enableNumeric bool, enableSpecialChars bool) string {
	chars := letterChars
	if enableCapital {
		chars += upperLetterChars
	}
	if enableNumeric {
		chars += numberChars
	}
	if enableSpecialChars {
		chars += specialChars
	}
	return chars
}

func GeneratePassword(length int, enableSpecialChars bool, enableNumeric bool, enableCapital bool) string {
	if length > MaxLength {
		length = MaxLength
	}
	chars := getCharacters(enableCapital, enableNumeric, enableSpecialChars)

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
