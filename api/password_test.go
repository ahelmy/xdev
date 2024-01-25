package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestPasswordAPI(t *testing.T) {
	app := fiber.New()
	passwordAPI(app)

	t.Run("Test /datetime endpoint - Test password API with default parameters", func(t *testing.T) {
		req1 := httptest.NewRequest(http.MethodGet, "/api/password", nil)
		resp1, err := app.Test(req1)
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp1.Body.Close()

		if resp1.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp1.StatusCode)
		}

		var passwordResponse1 Response
		err = json.NewDecoder(resp1.Body).Decode(&passwordResponse1)
		if err != nil {
			t.Fatalf("Failed to decode response body: %v", err)
		}

		if !passwordResponse1.Success {
			t.Errorf("Expected success to be true, but got false")
		}

		password1, ok := passwordResponse1.Data.(map[string]interface{})["password"].(string)
		if !ok {
			t.Errorf("Expected password to be a string, but got %T", passwordResponse1.Data.(map[string]interface{})["password"])
		}

		if len(password1) != 10 {
			t.Errorf("Expected password length to be 10, but got %d", len(password1))
		}
	})

	t.Run("Test /datetime endpoint - Test password API with length = 0", func(t *testing.T) {
		req1 := httptest.NewRequest(http.MethodGet, "/api/password?length=0", nil)
		resp1, err := app.Test(req1)
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp1.Body.Close()

		if resp1.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp1.StatusCode)
		}

		var passwordResponse1 Response
		err = json.NewDecoder(resp1.Body).Decode(&passwordResponse1)
		if err != nil {
			t.Fatalf("Failed to decode response body: %v", err)
		}

		if !passwordResponse1.Success {
			t.Errorf("Expected success to be true, but got false")
		}

		password1, ok := passwordResponse1.Data.(map[string]interface{})["password"].(string)
		if !ok {
			t.Errorf("Expected password to be a string, but got %T", passwordResponse1.Data.(map[string]interface{})["password"])
		}

		if len(password1) != 10 {
			t.Errorf("Expected password length to be 10, but got %d", len(password1))
		}
	})

	t.Run("Test /password endpoint - Test password API with custom parameters", func(t *testing.T) {
		req2 := httptest.NewRequest(http.MethodGet, "/api/password?length=12&especial=false&capital=false&number=false", nil)
		resp2, err := app.Test(req2)
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp2.Body.Close()

		if resp2.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp2.StatusCode)
		}

		var passwordResponse2 Response
		err = json.NewDecoder(resp2.Body).Decode(&passwordResponse2)
		if err != nil {
			t.Fatalf("Failed to decode response body: %v", err)
		}

		if !passwordResponse2.Success {
			t.Errorf("Expected success to be true, but got false")
		}

		password2, ok := passwordResponse2.Data.(map[string]interface{})["password"].(string)
		if !ok {
			t.Errorf("Expected password to be a string, but got %T", passwordResponse2.Data.(map[string]interface{})["password"])
		}

		if len(password2) != 12 {
			t.Errorf("Expected password length to be 12, but got %d", len(password2))
		}
	})
}
