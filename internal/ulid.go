package internal

import "github.com/oklog/ulid/v2"

func GenerateULID() string {
	return ulid.Make().String()
}