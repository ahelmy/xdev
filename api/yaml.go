package api

import (
	"encoding/json"

	"github.com/ahelmy/xdev/internal"
	"github.com/gofiber/fiber/v3"
)

const (
	YAMLPath = "/yaml"
)

type YamlRequest struct {
	Yaml string `json:"yaml"`
}

func yamlAPI(app *fiber.App) {
	app.Post(APIPrefix+YAMLPath, func(c fiber.Ctx) error {
		action := c.Query("action", "to_json")
		yamlRequest := YamlRequest{}
		err := json.Unmarshal(c.Request().Body(), &yamlRequest)
		if err != nil {
			return c.JSON(Response{Success: false, Message: err.Error()})
		}
		result := ""
		if action == "to_json" {
			_yaml, err := internal.Yaml2Json(yamlRequest.Yaml)
			if err != nil {
				return c.JSON(Response{Success: false, Message: err.Error()})
			}
			result = _yaml
		}
		if action == "to_properties" {
			_yaml, err := internal.Yaml2Properties(yamlRequest.Yaml)
			if err != nil {
				return c.JSON(Response{Success: false, Message: err.Error()})
			}
			result = _yaml
		}
		return c.JSON(Response{Success: true, Data: map[string]interface{}{"yaml": result}})
	})
}
