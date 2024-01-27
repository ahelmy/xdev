package api

import (
	"encoding/json"
	"strconv"

	"github.com/ahelmy/xdev/internal"
	"github.com/gofiber/fiber/v3"
)

const (
	TimePath = "/time"
)

type EpochTimeRequest struct {
	Epoch string `json:"epoch"`
}

type DateTimeRequest struct {
	DateTime string `json:"datetime"`
}

func timeAPI(app *fiber.App) {
	app.Post(APIPrefix+TimePath+"/epoch", func(c fiber.Ctx) error {
		epochTimeRequest := EpochTimeRequest{}
		err := json.Unmarshal(c.Request().Body(), &epochTimeRequest)
		if err != nil {
			return c.JSON(Response{Success: false, Message: err.Error()})
		}
		epoch, err := strconv.ParseInt(epochTimeRequest.Epoch, 10, 64)
		if err != nil {
			return c.JSON(Response{Success: false, Message: err.Error()})
		}
		var timeZone = c.Query("timezone")
		time := internal.ConvertTimeFromEpoch(epoch, internal.ToDateFormat, &timeZone)
		return c.JSON(Response{Success: true, Data: map[string]interface{}{"time": time}})
	})
	app.Post(APIPrefix+TimePath+"/datetime", func(c fiber.Ctx) error {
		dateTimeRequest := DateTimeRequest{}
		err := json.Unmarshal(c.Request().Body(), &dateTimeRequest)
		if err != nil {
			return c.JSON(Response{Success: false, Message: err.Error()})
		}
		var timeZone = c.Query("timezone")
		time, err := internal.ConvertTimeFromFormat(dateTimeRequest.DateTime, internal.ParseFormat("dd-MM-yyyy HH:mm:ss"), internal.ToDateFormat, &timeZone)
		if err != nil {
			return c.JSON(Response{Success: false, Message: err.Error()})
		}
		return c.JSON(Response{Success: true, Data: map[string]interface{}{"time": time}})
	})
	app.Get(APIPrefix+TimePath, func(c fiber.Ctx) error {
		var timeZone = c.Query("timezone")
		time, _ := internal.Now(internal.ToDateFormat, &timeZone)
		return c.JSON(Response{Success: true, Data: map[string]interface{}{"time": time}})
	})
}
