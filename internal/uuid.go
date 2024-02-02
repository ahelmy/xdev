package internal

import (
	"bytes"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/google/uuid"
)

const B32_CHARACTERS = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"

func GenerateUUID() string {
	return uuid.NewString()
}

func UUIDtoULID(uuidStr string) (string, error) {
	_, err := uuid.Parse(uuidStr)
	if err != nil {
		return "", err
	}
	uuidStr = strings.ReplaceAll(uuidStr, "-", "")
	uuidBytes, _ := hex.DecodeString(uuidStr)

	return crockfordEncode(uuidBytes), nil
}

func crockfordEncode(input []byte) string {
	var output []int
	bitsRead := 0
	buffer := 0

	// Work from the end of the buffer
	reversedInput := reverseBuffer(input)

	for _, byteVal := range reversedInput {
		// Add current byte to start of buffer
		buffer |= int(byteVal) << bitsRead
		bitsRead += 8

		for bitsRead >= 5 {
			output = append([]int{buffer & 0x1f}, output...)
			buffer >>= 5
			bitsRead -= 5
		}
	}

	if bitsRead > 0 {
		output = append([]int{buffer & 0x1f}, output...)
	}

	var resultBuffer bytes.Buffer
	for _, byteVal := range output {
		resultBuffer.WriteString(string(B32_CHARACTERS[byteVal]))
	}

	return resultBuffer.String()
}

func reverseBuffer(input []byte) []byte {
	length := len(input)
	reversed := make([]byte, length)
	for i := 0; i < length; i++ {
		reversed[i] = input[length-i-1]
	}
	return reversed
}

func reverseString(s string) string {
	var result []rune
	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, rune(s[i]))
	}
	return string(result)
}

func crockfordDecode(input string) ([]byte, error) {
	sanitizedInput := toUpperCase(input)

	// Work from the end
	sanitizedInput = reverseString(sanitizedInput)

	var output []byte
	var bitsRead, buffer uint
	for _, char := range sanitizedInput {
		byteVal := byte(strings.Index(B32_CHARACTERS, string(char)))
		if byteVal == 255 {
			return nil, errors.New("Invalid base 32 character found in string: " + string(char))
		}

		buffer |= uint(byteVal) << bitsRead
		bitsRead += 5

		for bitsRead >= 8 {
			output = append([]byte{byte(buffer & 0xff)}, output...)
			buffer >>= 8
			bitsRead -= 8
		}
	}

	if bitsRead >= 5 || buffer > 0 {
		output = append([]byte{byte(buffer & 0xff)}, output...)
	}

	return output, nil
}

func toUpperCase(s string) string {
	var result []rune
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			char -= 'a' - 'A'
		}
		result = append(result, char)
	}
	return string(result)
}
