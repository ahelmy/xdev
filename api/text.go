package api

import (
	"encoding/json"
	"github.com/ahelmy/xdev/internal"
	"github.com/gofiber/fiber/v3"
)

const (
	TextsPath = "/text"
)

type CompareRequest struct {
	Text1     string `json:"text1"`
	Text2     string `json:"text2"`
	CheckLine bool   `json:"check_line"`
}

func compareTextAPI(app *fiber.App) {
	app.Post(APIPrefix+TextsPath, func(c fiber.Ctx) error {
		req := CompareRequest{}
		err := json.Unmarshal(c.Request().Body(), &req)

		if err != nil {
			return c.JSON(Response{Success: false, Message: err.Error()})
		}
		diffs := internal.CompareText(req.Text1, req.Text2, req.CheckLine)

		res := Response{
			Success: true, Data: map[string]interface{}{"diffs": diffs}}
		return c.JSON(res)
	})
}
