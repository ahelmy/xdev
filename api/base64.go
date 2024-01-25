package api

import (
	"encoding/json"

	"github.com/ahelmy/xdev/internal"
	"github.com/gofiber/fiber/v3"
)

const (
	Base64Path = "/base64"
)

type Base64Request struct {
	String string `json:"string"`
}

func base64API(app *fiber.App) {
	app.Post(APIPrefix+Base64Path, func(c fiber.Ctx) error {
		action := c.Query("action")
		base64Request := Base64Request{}
		err := json.Unmarshal(c.Request().Body(), &base64Request)
		if err != nil {
			return c.JSON(Response{Success: false, Message: err.Error()})
		}

		result := ""
		switch action {
		case "encode":
			result = internal.EncodeToBase64(base64Request.String)
		case "decode":
			result = internal.DecodeFromBase64(base64Request.String)
		}
		return c.JSON(Response{Success: true, Data: map[string]interface{}{"base64": result}})
	})
}
