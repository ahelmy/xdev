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

func TestTimeAPI(t *testing.T) {
	app := fiber.New()
	timeAPI(app)
	// Test /epoch endpoint
	t.Run("Test /epoch endpoint", func(t *testing.T) {
		requestBody := `{"epoch": "1621234567"}`

		req := httptest.NewRequest(http.MethodPost, "/api/time/epoch", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, true, response["success"])
	})

	t.Run("Test /epoch endpoint - invalid req", func(t *testing.T) {
		requestBody := `x{"epoch2": "1621234567"}`

		req := httptest.NewRequest(http.MethodPost, "/api/time/epoch", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, false, response["success"])
	})

	t.Run("Test /epoch endpoint - invalid epoch", func(t *testing.T) {
		requestBody := `{"epoch": "xyz"}`

		req := httptest.NewRequest(http.MethodPost, "/api/time/epoch", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, false, response["success"])
	})

	// Test /datetime endpoint
	t.Run("Test /datetime endpoint", func(t *testing.T) {
		requestBody := `{"datetime": "17-05-2021 12:42:47"}`

		req := httptest.NewRequest(http.MethodPost, "/api/time/datetime", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, true, response["success"])
	})

	t.Run("Test /datetime endpoint - invalid req", func(t *testing.T) {
		requestBody := `x{"datetime": "17-05-2021 12:42:47"}`

		req := httptest.NewRequest(http.MethodPost, "/api/time/datetime", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, false, response["success"])
	})

	t.Run("Test /datetime endpoint - invalid datetime", func(t *testing.T) {
		requestBody := `{"datetime": "17-05-2021T12:42:47"}`

		req := httptest.NewRequest(http.MethodPost, "/api/time/datetime", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, false, response["success"])
	})

	// Test / endpoint
	t.Run("Test / endpoint", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/time", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, true, response["success"])
	})
}
