package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	app := fiber.New()
	healthCheck(app)
	// Create a test request
	req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)

	// Check the response status code
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Check the response body
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"success": true, "message": "OK", "data": null}`, string(body))
}
