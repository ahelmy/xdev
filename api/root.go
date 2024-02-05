package api

import (
	"github.com/gofiber/fiber/v3"
)

const (
	APIPrefix       = "/api"
	healthCheckPath = "/health"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func AddAPILayer(app *fiber.App) {
	healthCheck(app)
	uuidAPI(app)
	ulidAPI(app)
	passwordAPI(app)
	jwtAPI(app)
	yamlAPI(app)
	jsonAPI(app)
	base64API(app)
	urlAPI(app)
	hashAPI(app)
	timeAPI(app)
	propertiesAPI(app)
	compareTextAPI(app)
}

func healthCheck(app *fiber.App) {
	app.Get(APIPrefix+healthCheckPath, func(c fiber.Ctx) error {
		return c.JSON(Response{Success: true, Message: "OK"})
	})
}
