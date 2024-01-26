package api

import (
	"encoding/json"

	"github.com/ahelmy/xdev/internal"
	"github.com/gofiber/fiber/v3"
)

const (
	JWTPath = "/jwt"
)

type JWTEncodeRequest struct {
	Headers map[string]interface{} `json:"headers"`
	Claims  map[string]interface{} `json:"claims"`
	Secret  string                 `json:"secret"`
}

type JWTDecodeRequest struct {
	JWT string `json:"jwt"`
}

func jwtAPI(app *fiber.App) {
	app.Post(APIPrefix+JWTPath+"/encode", func(c fiber.Ctx) error {
		jwtEncodeRequest := JWTEncodeRequest{}
		err := json.Unmarshal(c.Request().Body(), &jwtEncodeRequest)
		if err != nil {
			return c.JSON(Response{Success: false, Message: err.Error()})
		}

		jwtToken, err := internal.EncodeJWT(jwtEncodeRequest.Headers, jwtEncodeRequest.Claims, jwtEncodeRequest.Secret)
		if err != nil {
			return c.JSON(Response{Success: false, Message: err.Error()})
		}
		jwtResponse := Response{Success: true, Data: map[string]interface{}{"jwt": jwtToken}}
		return c.JSON(jwtResponse)
	})

	app.Post(APIPrefix+JWTPath+"/decode", func(c fiber.Ctx) error {
		jwtDecodeRequest := JWTDecodeRequest{}
		err := json.Unmarshal(c.Request().Body(), &jwtDecodeRequest)
		if err != nil {
			return c.JSON(Response{Success: false, Message: err.Error()})
		}

		jwtToken, err := internal.DecodeJWT(jwtDecodeRequest.JWT)
		if err != nil {
			return c.JSON(Response{Success: false, Message: err.Error()})
		}
		jwtResponse := Response{Success: true, Data: map[string]interface{}{"headers": jwtToken.Header, "claims": jwtToken.Claims, "expires": jwtToken.Expires}}
		return c.JSON(jwtResponse)
	})
}
