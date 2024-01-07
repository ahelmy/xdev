package internal

import (
	"crypto"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashString(str string, algorithm string) (string, error) {
	switch algorithm {
	case "md5":
		return hash(str, crypto.MD5)
	case "sha256":
		return hash(str, crypto.SHA256)
	case "sha512":
		return hash(str, crypto.SHA512)
	case "salt":
		return salt(str)
	}
	return str, fmt.Errorf("invalid algorithm")
}

func hash(str string, algorithm crypto.Hash) (string, error) {
	hasher := algorithm.New()

	hasher.Write([]byte(str))

	hashBytes := hasher.Sum(nil)

	return hex.EncodeToString(hashBytes), nil
}
func salt(str string) (string, error) {
	salt, err := bcrypt.GenerateFromPassword([]byte(str), 10)
	if err != nil {
		return "", err
	}
	return string(salt), nil
}
