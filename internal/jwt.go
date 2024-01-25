package internal

import (
	"encoding/json"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Header  string
	Claims  string
	Expires *time.Time
}

func DecodeJWT(jwtToken string) (JWT, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(jwtToken, jwt.MapClaims{})
	if err != nil {
		return JWT{}, err
	}
	headersJSON, _ := json.MarshalIndent(token.Header, "", "    ")

	claims := token.Claims.(jwt.MapClaims)
	var expires *time.Time
	expireTime, err := token.Claims.GetExpirationTime()
	if err != nil || expireTime == nil {
		expires = nil
	} else {
		expires = &expireTime.Time
	}
	claimsJSON, _ := json.MarshalIndent(claims, "", "    ")

	return JWT{Header: string(headersJSON), Claims: string(claimsJSON), Expires: expires}, nil
}

func EncodeJWT(headers map[string]interface{}, claims map[string]interface{}, signature string) (string, error) {
	algorithms := jwt.GetAlgorithms()
	algorithmExists := false
	if _, ok := headers["alg"]; !ok {
		headers["alg"] = "HS256"
	}
	algorithm := headers["alg"].(string)
	for _, a := range algorithms {
		if a == headers["alg"] {
			algorithm = a
			algorithmExists = true
			break
		}
	}
	if !algorithmExists {
		return "", fmt.Errorf("algorithm %s is not supported", algorithm)
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod(algorithm), jwt.MapClaims(claims))
	token.Header = headers
	signedString, err := token.SignedString([]byte(signature))
	return signedString, err
}
