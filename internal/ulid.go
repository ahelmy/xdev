package internal

import (
	"encoding/hex"

	"github.com/oklog/ulid/v2"
)

func GenerateULID() string {
	return ulid.Make().String()
}

func ULIDtoUUID(ulidStr string) (string, error) {
	_, err := ulid.Parse(ulidStr)
	if err != nil {
		return "", err
	}
	decoded, _ := crockfordDecode(ulidStr)
	uuid := hex.EncodeToString(decoded)
	uuid = uuid[0:8] + "-" + uuid[8:12] + "-" + uuid[12:16] + "-" + uuid[16:20] + "-" + uuid[20:]
	return uuid, nil
}
