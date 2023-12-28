package internal

import (
	"encoding/json"
	"log"

	jwt "github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Header    string
	Claims    string
	Signature string
}

func DecodeJWT(jwtToken string) (JWT, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(jwtToken, jwt.MapClaims{})
	if err != nil {
		return JWT{}, err
	}
	headersJSON, err := json.MarshalIndent(token.Header, "", "    ")
	if err != nil {
		log.Fatal("Error converting heade to JSON:", err)
		return JWT{}, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("Error decoding claims")
		return JWT{}, err
	}
	claimsJSON, err := json.MarshalIndent(claims, "", "    ")
	if err != nil {
		log.Fatal("Error converting claims to JSON:", err)
		return JWT{}, err
	}
	signature := token.Signature

	return JWT{Header: string(headersJSON), Claims: string(claimsJSON), Signature: string(signature)}, nil
}
