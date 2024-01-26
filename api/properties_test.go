package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestPropertiesAPI(t *testing.T) {
	app := fiber.New()
	propertiesAPI(app)
	t.Run("Valid Request", func(t *testing.T) {
		propertiesRequest := PropertiesRequest{
			Properties: "parent.child=value\n",
		}
		requestBody, _ := json.Marshal(propertiesRequest)

		req := httptest.NewRequest(http.MethodPost, "/api/properties", bytes.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response Response
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.True(t, response.Success)
		assert.Equal(t, map[string]interface{}{"yaml": "parent:\n    child: value\n"}, response.Data)
	})

	t.Run("Invalid Request", func(t *testing.T) {
		propertiesRequest := PropertiesRequest{
			Properties: "invalid: properties: string",
		}
		requestBody, _ := json.Marshal(propertiesRequest)

		req := httptest.NewRequest(http.MethodPost, "/api/properties", bytes.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response Response
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.False(t, response.Success)
	})

	t.Run("JSON Unmarshal Error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/properties", bytes.NewReader([]byte("invalid json")))
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
