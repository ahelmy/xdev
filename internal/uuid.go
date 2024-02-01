package internal

import (
	"bytes"
	"encoding/hex"
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
