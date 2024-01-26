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

func TestBase64API(t *testing.T) {
	app := fiber.New()
	base64API(app)
	t.Run("Test encode action", func(t *testing.T) {
		requestBody := `{"action": "encode", "string": "Hello, World!"}`

		req := httptest.NewRequest(http.MethodPost, "/api/base64?action=encode", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response Response
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, true, response.Success)
	})

	t.Run("Test encode action - invalid req", func(t *testing.T) {
		requestBody := `x{"action": "encode", "string": "Hello, World!"}`

		req := httptest.NewRequest(http.MethodPost, "/api/base64?action=encode", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response Response
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, false, response.Success)
	})

	t.Run("Test decode action", func(t *testing.T) {
		requestBody := `{"action": "decode", "string": "SGVsbG8sIFdvcmxkIQ=="}`

		req := httptest.NewRequest(http.MethodPost, "/api/base64?action=decode", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response Response
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, true, response.Success)
	})
}
