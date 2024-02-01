package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestYamlAPI(t *testing.T) {
	app := fiber.New()
	yamlAPI(app)

	t.Run("Test YAML to JSON conversion", func(t *testing.T) {
		// Prepare request body
		requestBody := `{"yaml": "key: value"}`

		// Create a new request
		req := httptest.NewRequest(http.MethodPost, "/api/yaml?action=to_json", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")

		// Perform the request
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("Failed to perform request: %v", err)
		}
		defer resp.Body.Close()

		// Check the response status code
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		// Parse the response body
		var response Response
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Fatalf("Failed to parse response body: %v", err)
		}

		// Check the response data
		expectedResult := "{\n    \"key\": \"value\"\n}"
		result := response.Data.(map[string]interface{})["yaml"].(string)
		if response.Success != true || result != expectedResult {
			t.Errorf("Unexpected response data: %+v", response)
		}
	})

	t.Run("Test YAML to JSON conversion - invalid req", func(t *testing.T) {
		// Prepare request body
		requestBody := `x{"yaml": "key: value"}`

		// Create a new request
		req := httptest.NewRequest(http.MethodPost, "/api/yaml?action=to_json", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")

		// Perform the request
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("Failed to perform request: %v", err)
		}
		defer resp.Body.Close()

		// Check the response status code
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		// Parse the response body
		var response Response
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Fatalf("Failed to parse response body: %v", err)
		}

		// Check the response data
		if response.Success {
			t.Errorf("Unexpected response data: %+v", response)
		}
	})

	t.Run("Test YAML to JSON conversion - invalid yaml", func(t *testing.T) {
		// Prepare request body
		requestBody := `{"yaml": "key:value"}`

		// Create a new request
		req := httptest.NewRequest(http.MethodPost, "/api/yaml?action=to_json", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")

		// Perform the request
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("Failed to perform request: %v", err)
		}
		defer resp.Body.Close()

		// Check the response status code
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		// Parse the response body
		var response Response
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Fatalf("Failed to parse response body: %v", err)
		}

		// Check the response data
		if response.Success {
			t.Errorf("Unexpected response data: %+v", response)
		}
	})

	t.Run("Test YAML to Properties conversion", func(t *testing.T) {
		// Prepare request body
		requestBody := `{"yaml": "key: value"}`

		// Create a new request
		req := httptest.NewRequest(http.MethodPost, "/api/yaml?action=to_properties", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")

		// Perform the request
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("Failed to perform request: %v", err)
		}
		defer resp.Body.Close()

		// Check the response status code
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		// Parse the response body
		var response Response
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Fatalf("Failed to parse response body: %v", err)
		}

		// Check the response data
		expectedResult := "key=value\n"
		result := response.Data.(map[string]interface{})["yaml"].(string)
		if response.Success != true || result != expectedResult {
			t.Errorf("Unexpected response data: %+v", response)
		}
	})

	t.Run("Test YAML to Properties conversion - invalid req", func(t *testing.T) {
		requestBody := `x{"yaml": "c"}`

		// Create a new request
		req := httptest.NewRequest(http.MethodPost, "/api/yaml?action=to_properties", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")

		// Perform the request
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("Failed to perform request: %v", err)
		}
		defer resp.Body.Close()

		// Check the response status code
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		// Parse the response body
		var response Response
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Fatalf("Failed to parse response body: %v", err)
		}

		assert.Equal(t, false, response.Success)
	})

	t.Run("Test YAML to Properties conversion - empty yaml", func(t *testing.T) {
		requestBody := `{"yaml": ""}`

		// Create a new request
		req := httptest.NewRequest(http.MethodPost, "/api/yaml?action=to_properties", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")

		// Perform the request
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("Failed to perform request: %v", err)
		}
		defer resp.Body.Close()

		// Check the response status code
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		// Parse the response body
		var response Response
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Fatalf("Failed to parse response body: %v", err)
		}
		fmt.Println(response)

		assert.Equal(t, true, response.Success)
	})
}
