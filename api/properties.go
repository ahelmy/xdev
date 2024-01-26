package api

import (
	"encoding/json"

	"github.com/ahelmy/xdev/internal"
	"github.com/gofiber/fiber/v3"
)

const (
	PropertiesPath = "/properties"
)

type PropertiesRequest struct {
	Properties string `json:"properties"`
}

func propertiesAPI(app *fiber.App) {
	app.Post(APIPrefix+PropertiesPath, func(c fiber.Ctx) error {
		propertiesRequest := PropertiesRequest{}
		err := json.Unmarshal(c.Request().Body(), &propertiesRequest)
		if err != nil {
			return c.JSON(Response{Success: false, Message: err.Error()})
		}

		yamlData, err := internal.Properties2Yaml(propertiesRequest.Properties)
		if err != nil {
			return c.JSON(Response{Success: false, Message: err.Error()})
		}

		yamlResponse := Response{
			Success: true, Data: map[string]interface{}{"yaml": yamlData}}
		return c.JSON(yamlResponse)
	})
}
