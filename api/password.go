package api

import (
	"github.com/ahelmy/xdev/internal"
	"github.com/gofiber/fiber/v3"
)

const (
	PasswordPath = "/password"
)

func passwordAPI(app *fiber.App) {
	app.Get(APIPrefix+PasswordPath, func(c fiber.Ctx) error {
		length := c.QueryInt("length", 10)
		if length == 0 {
			length = 10
		}
		isEspecial := c.QueryBool("especial", true)
		isCapital := c.QueryBool("capital", true)
		isNumberic := c.QueryBool("number", true)
		password := internal.GeneratePassword(length, isEspecial, isNumberic, isCapital)
		passwordResponse := Response{Success: true, Data: map[string]interface{}{"password": password}}
		return c.JSON(passwordResponse)
	})
}
