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

func TestJWTAPI(t *testing.T) {
	app := fiber.New()
	jwtAPI(app)
	t.Run("Test JWT encode endpoint", func(t *testing.T) {
		// Prepare request body
		jwtEncodeRequest := JWTEncodeRequest{
			Headers: map[string]interface{}{"alg": "HS256"},
			Claims:  map[string]interface{}{"sub": "Ahelmy", "exp": 1566964030},
			Secret:  "secret-key",
		}
		requestBody, _ := json.Marshal(jwtEncodeRequest)

		// Make request to the endpoint
		resp, err := app.Test(httptest.NewRequest(http.MethodPost, "/api/jwt/encode", bytes.NewReader(requestBody)))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// Parse response body
		var jwtResponse Response
		err = json.NewDecoder(resp.Body).Decode(&jwtResponse)
		assert.NoError(t, err)

		// Assert response
		assert.True(t, jwtResponse.Success)
		assert.NotNil(t, jwtResponse.Data.(map[string]interface{})["jwt"])
	})

	t.Run("Test JWT encode endpoint - invalid req", func(t *testing.T) {
		// Prepare request body
		request := "x{\"headers\":{},\"claims\":{},\"secret\":\"\"}"

		// Make request to the endpoint
		resp, err := app.Test(httptest.NewRequest(http.MethodPost, "/api/jwt/encode", bytes.NewReader([]byte(request))))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// Parse response body
		var jwtResponse Response
		err = json.NewDecoder(resp.Body).Decode(&jwtResponse)
		assert.NoError(t, err)

		// Assert response
		assert.False(t, jwtResponse.Success)
	})

	t.Run("Test JWT encode endpoint - invalid jwt params", func(t *testing.T) {
		// Prepare request body
		request := "{\"headers\":{\"alg\":\"xyz\"},\"claims\":{},\"secret\":\"\"}"

		// Make request to the endpoint
		resp, err := app.Test(httptest.NewRequest(http.MethodPost, "/api/jwt/encode", bytes.NewReader([]byte(request))))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// Parse response body
		var jwtResponse Response
		err = json.NewDecoder(resp.Body).Decode(&jwtResponse)
		assert.NoError(t, err)

		// Assert response
		assert.False(t, jwtResponse.Success)
	})

	t.Run("Test JWT decode endpoint", func(t *testing.T) {
		// Prepare request body
		jwtDecodeRequest := JWTDecodeRequest{
			JWT: "eyJhbGciOiJIUzI1NiIsInMiOiJ4In0.eyJuYW1lIjoiYWxpIn0.R8Uvav2uxGn3vsnSrfDLzK76xzvgrfmv9crfzuREYpc",
		}
		requestBody, _ := json.Marshal(jwtDecodeRequest)

		// Make request to the endpoint
		resp, err := app.Test(httptest.NewRequest(http.MethodPost, "/api/jwt/decode", bytes.NewReader(requestBody)))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// Parse response body
		var jwtResponse Response
		err = json.NewDecoder(resp.Body).Decode(&jwtResponse)
		assert.NoError(t, err)

		// Assert response
		assert.True(t, jwtResponse.Success)
		assert.NotNil(t, jwtResponse.Data.(map[string]interface{})["headers"])
		assert.NotNil(t, jwtResponse.Data.(map[string]interface{})["claims"])
	})

	t.Run("Test JWT decode endpoint - invalid req", func(t *testing.T) {
		// Prepare request body
		request := "x{\"jwt\":\"\"}"

		// Make request to the endpoint
		resp, err := app.Test(httptest.NewRequest(http.MethodPost, "/api/jwt/decode", bytes.NewReader([]byte(request))))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// Parse response body
		var jwtResponse Response
		err = json.NewDecoder(resp.Body).Decode(&jwtResponse)
		assert.NoError(t, err)

		// Assert response
		assert.False(t, jwtResponse.Success)
	})

	t.Run("Test JWT decode endpoint - invalid jwt", func(t *testing.T) {
		// Prepare request body
		request := "{\"jwt\":\"xxyyysss\"}"

		// Make request to the endpoint
		resp, err := app.Test(httptest.NewRequest(http.MethodPost, "/api/jwt/decode", bytes.NewReader([]byte(request))))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// Parse response body
		var jwtResponse Response
		err = json.NewDecoder(resp.Body).Decode(&jwtResponse)
		assert.NoError(t, err)

		// Assert response
		assert.False(t, jwtResponse.Success)
	})

}
