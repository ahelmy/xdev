package api

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCompareTextAPI(t *testing.T) {
	app := fiber.New()
	compareTextAPI(app)

	t.Run("Valid Request", func(t *testing.T) {
		compareRequest := CompareRequest{
			Text1:     "text",
			Text2:     "text bro",
			CheckLine: false,
		}
		reqBody, _ := json.Marshal(compareRequest)

		req := httptest.NewRequest(http.MethodPost, "/api/text", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		res, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response Response
		err = json.NewDecoder(res.Body).Decode(&response)
		assert.NoError(t, err)
		assert.True(t, response.Success)
	})

	t.Run("JSON Unmarshal Error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/text", bytes.NewReader([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response Response
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.False(t, response.Success)
	})
}
