package api

import (
	"github.com/gofiber/fiber/v3"
)

const (
	APIPrefix = "/api"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func AddAPILayer(app *fiber.App) {
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
}
