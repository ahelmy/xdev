package api

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestHashAPI(t *testing.T) {
	app := fiber.New()
	hashAPI(app)
	// Test /hash endpoint
	t.Run("Test /hash endpoint", func(t *testing.T) {
		requestBody := `{"string": "test"}`
		req := httptest.NewRequest("POST", "/api/hash", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("Failed to perform request: %v", err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			t.Errorf("Expected status code %d, but got %d", 200, resp.StatusCode)
		}
		var response Response
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Fatalf("Failed to parse response body: %v", err)
		}
		expectedResult := fmt.Sprintf("%x", md5.Sum([]byte("test")))
		result := response.Data.(map[string]interface{})["hash"].(string)
		if response.Success != true || result != expectedResult {
			t.Errorf("Unexpected response data: %+v", response)
		}
	})

	t.Run("Test /hash endpoint - invalid req", func(t *testing.T) {
		requestBody := `x{"string": "test"}`
		req := httptest.NewRequest("POST", "/api/hash", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("Failed to perform request: %v", err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			t.Errorf("Expected status code %d, but got %d", 200, resp.StatusCode)
		}
		var response Response
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Fatalf("Failed to parse response body: %v", err)
		}
		if response.Success {
			t.Errorf("Unexpected response data: %+v", response)
		}
	})

	t.Run("Test /hash endpoint - invalid algorithm", func(t *testing.T) {
		requestBody := `{"string": "test"}`
		req := httptest.NewRequest("POST", "/api/hash?algorithm=xyz", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("Failed to perform request: %v", err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			t.Errorf("Expected status code %d, but got %d", 200, resp.StatusCode)
		}
		var response Response
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Fatalf("Failed to parse response body: %v", err)
		}
		if response.Success {
			t.Errorf("Unexpected response data: %+v", response)
		}
	})
}
