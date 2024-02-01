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
		uuidResponse := Response{Success: true, Data: map[string]interface{}{"uuid": internal.GenerateUUID()}}
		return c.JSON(uuidResponse)
	})

	app.Get(APIPrefix+UUIDPath+"/convert", func(c fiber.Ctx) error {
		uuid := c.Query("uuid", "")
		if uuid == "" {
			return c.JSON(Response{Success: false, Message: "uuid is required"})
		}

		convertTo := c.Query("to", "ulid")
		var convertedString string
		if convertTo == "ulid" {
			ulid, err := internal.UUIDtoULID(uuid)
			if err != nil {
				return c.JSON(Response{Success: false, Message: err.Error()})
			}
			convertedString = ulid
		} else {
			return c.JSON(Response{Success: false, Message: "invalid conversion type"})
		}
		return c.JSON(Response{Success: true, Data: map[string]interface{}{"result": convertedString}})
	})
}
