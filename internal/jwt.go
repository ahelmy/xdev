package internal

import (
	"encoding/json"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Header    string
	Claims    string
	Signature string
	Expires   *time.Time
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
	}else{
		expires = &expireTime.Time
	}
	claimsJSON, _ := json.MarshalIndent(claims, "", "    ")
	signature := token.Signature

	return JWT{Header: string(headersJSON), Claims: string(claimsJSON), Signature: string(signature), Expires: expires}, nil
}
