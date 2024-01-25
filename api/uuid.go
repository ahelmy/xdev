package api

import (
	"github.com/ahelmy/xdev/internal"
	"github.com/gofiber/fiber/v3"
)

const (
	UUIDPath = "/uuid"
)

func uuidAPI(app *fiber.App) {
	app.Get(APIPrefix+UUIDPath, func(c fiber.Ctx) error {
		uuidResponse := Response{Success: true, Data: map[string]interface{}{"uuid": internal.GenerateGUID()}}
		return c.JSON(uuidResponse)
	})
}
