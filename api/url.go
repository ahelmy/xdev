package api

import (
	"encoding/json"

	"github.com/ahelmy/xdev/internal"
	"github.com/gofiber/fiber/v3"
)

const (
	URLPath = "/url"
)

type URLRequest struct {
	URL string `json:"url"`
}

func urlAPI(app *fiber.App) {
	app.Post(APIPrefix+URLPath, func(c fiber.Ctx) error {
		action := c.Query("action")
		urlRequest := URLRequest{}
		err := json.Unmarshal(c.Request().Body(), &urlRequest)
		if err != nil {
			return c.JSON(Response{Success: false, Message: err.Error()})
		}
		result := ""
		switch action {
		case "encode":
			result = internal.EncodeURL(urlRequest.URL)
		case "decode":
			result, _ = internal.DecodeURL(urlRequest.URL)
		}
		return c.JSON(Response{Success: true, Data: map[string]interface{}{"url": result}})
	})
}
