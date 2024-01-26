package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestURLAPI(t *testing.T) {
	app := fiber.New()
	urlAPI(app)
	t.Run("Test /url endpoint", func(t *testing.T) {
		requestBody := `{"url": "https://google.com"}`
		req := httptest.NewRequest(http.MethodPost, "/api/url?action=encode", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response Response
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, true, response.Success)
	})
	t.Run("Test /url endpoint - decode", func(t *testing.T) {
		requestBody := `{"url": "https%3A%2F%2Fgoogle.com"}`
		req := httptest.NewRequest(http.MethodPost, "/api/url?action=decode", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response Response
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, true, response.Success)
	})

	t.Run("Test /url endpoint - decode - error", func(t *testing.T) {
		requestBody := `x{"url": "htt'a%%%%%%%'asps%3FFFA%2AF%2AFgoogle.com"}`
		req := httptest.NewRequest(http.MethodPost, "/api/url?action=decode", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response Response
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, false, response.Success)
	})
}
