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

	app.Get(APIPrefix+ULIDPath+"/convert", func(c fiber.Ctx) error {
		ulid := c.Query("ulid", "")
		if ulid == "" {
			return c.JSON(Response{Success: false, Message: "ulid is required"})
		}

		convertTo := c.Query("to", "uuid")
		var convertedString string
		if convertTo == "uuid" {
			uuid, err := internal.ULIDtoUUID(ulid)
			if err != nil {
				return c.JSON(Response{Success: false, Message: err.Error()})
			}
			convertedString = uuid
		} else {
			return c.JSON(Response{Success: false, Message: "invalid conversion type"})
		}
		return c.JSON(Response{Success: true, Data: map[string]interface{}{"result": convertedString}})
	})
}
