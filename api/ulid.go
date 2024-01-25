package api

import (
	"github.com/ahelmy/xdev/internal"
	"github.com/gofiber/fiber/v3"
)

const (
	ULIDPath = "/ulid"
)

func ulidAPI(app *fiber.App) {
	app.Get(APIPrefix+ULIDPath, func(c fiber.Ctx) error {
		ulidResponse := Response{Success: true, Data: map[string]interface{}{"ulid": internal.GenerateULID()}}
		return c.JSON(ulidResponse)
	})
}
