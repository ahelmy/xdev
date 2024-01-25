package api

import (
	"github.com/ahelmy/xdev/internal"
	"github.com/gofiber/fiber/v3"
)

const (
	JSONPath = "/json"
)

func jsonAPI(app *fiber.App) {
	app.Post(APIPrefix+JSONPath, func(c fiber.Ctx) error {
		action := c.Query("action")
		json := string(c.Request().Body())
		result := ""
		if len(json) > 0 {
			localResult := ""
			err := error(nil)
			switch action {
			case "beautify":
				localResult, err = internal.IndentJSON(json)
			case "minify":
				localResult, err = internal.MinifyJSON(json)
			case "json2Yaml":
				localResult, err = internal.Json2Yaml(json)
			default:
				return c.JSON(Response{Success: false, Message: "Invalid action"})
			}
			if err != nil {
				return c.JSON(Response{Success: false, Message: err.Error()})
			} else {
				result = localResult
			}
		}
		return c.JSON(Response{Success: true, Data: map[string]interface{}{"json": result}})
	})
}
