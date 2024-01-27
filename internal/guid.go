package internal

import "github.com/google/uuid"

func GenerateGUID() string {
	return uuid.NewString()
}
