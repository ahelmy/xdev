package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestJsonAPI(t *testing.T) {
	app := fiber.New()
	jsonAPI(app)
	t.Run("Test JSON API - Beautify", func(t *testing.T) {
		jsonInput := `{"name":"John","age":30,"city":"New York"}`
		expectedOutput := `{
    "name": "John",
    "age": 30,
    "city": "New York"
}`

		req := httptest.NewRequest(http.MethodPost, "/api/json?action=beautify", bytes.NewBufferString(jsonInput))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		if err != nil {
			t.Fatal(err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		var response map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			t.Fatal(err)
		}

		if response["success"] != true {
			t.Errorf("Expected Success to be true, but got %v", response["Success"])
		}

		if response["data"].(map[string]interface{})["json"] != expectedOutput {
			t.Errorf("Expected JSON output to be %s, but got %v", expectedOutput, response["Data"].(map[string]interface{})["json"])
		}
	})

	t.Run("Test JSON API - Minify", func(t *testing.T) {
		jsonInput := `{
    "name": "John",
    "age": 30,
    "city": "New York"
}`
		expectedOutput := `{"name":"John","age":30,"city":"New York"}`

		req := httptest.NewRequest(http.MethodPost, "/api/json?action=minify", bytes.NewBufferString(jsonInput))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		if err != nil {
			t.Fatal(err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		var response map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			t.Fatal(err)
		}

		if response["success"] != true {
			t.Errorf("Expected Success to be true, but got %v", response["Success"])
		}

		if response["data"].(map[string]interface{})["json"] != expectedOutput {
			t.Errorf("Expected JSON output to be %s, but got %v", expectedOutput, response["Data"].(map[string]interface{})["json"])
		}
	})

	t.Run("Test JSON API - JSON to YAML", func(t *testing.T) {
		jsonInput := `{"name":"John","age":30,"city":"New York"}`
		expectedOutput := `name: John
age: 30
city: New York
`

		req := httptest.NewRequest(http.MethodPost, "/api/json?action=json2Yaml", bytes.NewBufferString(jsonInput))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		if err != nil {
			t.Fatal(err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		var response map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			t.Fatal(err)
		}

		if response["success"] != true {
			t.Errorf("Expected Success to be true, but got %v", response["Success"])
		}

		if response["data"].(map[string]interface{})["json"] != expectedOutput {
			t.Errorf("Expected JSON output to be %s, but got %v", expectedOutput, response["Data"].(map[string]interface{})["json"])
		}
	})

	t.Run("Test JSON API - Invalid Action", func(t *testing.T) {
		jsonInput := `{"name":"John","age":30,"city":"New York"}`

		req := httptest.NewRequest(http.MethodPost, "/api/json?action=invalid", bytes.NewBufferString(jsonInput))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		if err != nil {
			t.Fatal(err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		var response map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			t.Fatal(err)
		}

		if response["success"] != false {
			t.Errorf("Expected Success to be false, but got %v", response["Success"])
		}

	})

	t.Run("Test JSON API - Invalid JSON", func(t *testing.T) {
		jsonInput := `{name":"John","age":30,"city":"New York"}`

		req := httptest.NewRequest(http.MethodPost, "/api/json?action=minify", bytes.NewBufferString(jsonInput))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		if err != nil {
			t.Fatal(err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		var response map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			t.Fatal(err)
		}

		if response["success"] != false {
			t.Errorf("Expected Success to be false, but got %v", response["Success"])
		}

	})

}
