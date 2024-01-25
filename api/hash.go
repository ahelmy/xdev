package api

import (
	"encoding/json"

	"github.com/ahelmy/xdev/internal"
	"github.com/gofiber/fiber/v3"
)

const (
	HashPath = "/hash"
)

type HashRequest struct {
	String string `json:"string"`
}

func hashAPI(app *fiber.App) {
	app.Post(APIPrefix+HashPath, func(c fiber.Ctx) error {
		algorithm := c.Query("algorithm", "md5")
		hashRequest := HashRequest{}
		err := json.Unmarshal(c.Request().Body(), &hashRequest)
		if err != nil {
			return c.JSON(Response{Success: false, Message: err.Error()})
		}
		result, err := internal.HashString(hashRequest.String, algorithm)
		if err != nil {
			return c.JSON(Response{Success: false, Message: err.Error()})
		}
		return c.JSON(Response{Success: true, Data: map[string]interface{}{"hash": result}})
	})
}
