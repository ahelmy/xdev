package internal

import (
	b64 "encoding/base64"
)

// encodeToBase64 encodes a string to base64
func EncodeToBase64(input string) string {
	return b64.StdEncoding.EncodeToString([]byte(input))
}

func DecodeFromBase64(input string) string {
	decoded, err := b64.StdEncoding.DecodeString(input)
	if err != nil {
		return ""
	}
	return string(decoded)
}