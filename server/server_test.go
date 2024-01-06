package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexPage(t *testing.T) {
	app := newApp()
	indexPage(app)
	// Create a test request to the "/" route
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to send test request: %v", err)
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}

	// TODO: Add more assertions for the response body or other expectations
}

func TestJSONPage(t *testing.T) {
	app := newApp()
	jsonPage(app)
	t.Run("Test JSON Page - Beautify", func(t *testing.T) {
		// Create a test request to the "/json" route with action=beautify and json=...
		req := httptest.NewRequest(http.MethodGet, "/json?action=beautify&json=...", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("Failed to send test request: %v", err)
		}

		// Check the response status code
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		// TODO: Add assertions for the response body or other expectations
	})

	t.Run("Test JSON Page - Minify", func(t *testing.T) {
		// Create a test request to the "/json" route with action=minify and json=...
		req := httptest.NewRequest(http.MethodGet, "/json?action=minify&json=...", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("Failed to send test request: %v", err)
		}

		// Check the response status code
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		// TODO: Add assertions for the response body or other expectations
	})

	t.Run("Test JSON Page - json2Yaml", func(t *testing.T) {
		// Create a test request to the "/json" route with action=json2Yaml and json=...
		req := httptest.NewRequest(http.MethodGet, "/json?action=json2Yaml&json=...", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("Failed to send test request: %v", err)
		}

		// Check the response status code
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		// TODO: Add assertions for the response body or other expectations
	})
}
func TestUUIDPage(t *testing.T) {
	app := newApp()
	uuidPage(app)
	// Create a test request to the "/uuid" route
	req := httptest.NewRequest(http.MethodGet, "/uuid", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to send test request: %v", err)
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}

	// TODO: Add assertions for the response body or other expectations
}

func TestULIDPage(t *testing.T) {
	app := newApp()
	ulidPage(app)
	// Create a test request to the "/ulid" route
	req := httptest.NewRequest(http.MethodGet, "/ulid", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to send test request: %v", err)
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}

	// TODO: Add assertions for the response body or other expectations
}

func TestPasswordPage(t *testing.T) {
	app := newApp()
	passwordPage(app)
	t.Run("Test Password Page", func(t *testing.T) {
		// Create a test request to the "/password" route
		req := httptest.NewRequest(http.MethodGet, "/password", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("Failed to send test request: %v", err)
		}

		// Check the response status code
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		// TODO: Add assertions for the response body or other expectations
	})
}
func TestYAMLPage(t *testing.T) {
	app := newApp()
	yamlPage(app)
	t.Run("Test YAML Page - Beautify", func(t *testing.T) {
		// Create a test request to the "/yaml" route with action=beautify and yaml=...
		req := httptest.NewRequest(http.MethodGet, "/yaml?action=beautify&yaml=...", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("Failed to send test request: %v", err)
		}

		// Check the response status code
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		// TODO: Add assertions for the response body or other expectations
	})

	t.Run("Test YAML Page - Minify", func(t *testing.T) {
		// Create a test request to the "/yaml" route with action=minify and yaml=...
		req := httptest.NewRequest(http.MethodGet, "/yaml?action=minify&yaml=...", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("Failed to send test request: %v", err)
		}

		// Check the response status code
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		// TODO: Add assertions for the response body or other expectations
	})

	t.Run("Test YAML Page - yaml2JSON", func(t *testing.T) {
		// Create a test request to the "/yaml" route with action=yaml2JSON and yaml=...
		req := httptest.NewRequest(http.MethodGet, "/yaml?action=yaml2JSON&yaml=...", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("Failed to send test request: %v", err)
		}

		// Check the response status code
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		// TODO: Add assertions for the response body or other expectations
	})
}

func TestJWTPage(t *testing.T) {
	app := newApp()
	jwtPage(app)
	t.Run("Test JWT Page", func(t *testing.T) {
		// Create a test request to the "/jwt" route with jwt=...
		req := httptest.NewRequest(http.MethodGet, "/jwt?jwt=...", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("Failed to send test request: %v", err)
		}

		// Check the response status code
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		// TODO: Add assertions for the response body or other expectations
	})
}

func TestBase64Page(t *testing.T) {
	app := newApp()
	base64Page(app)
	t.Run("Test Base64 Page - Encode", func(t *testing.T) {
		// Create a test request to the "/base64" route with action=encode and decoded=...
		req := httptest.NewRequest(http.MethodGet, "/base64?action=encode&decoded=...", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("Failed to send test request: %v", err)
		}

		// Check the response status code
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		// TODO: Add assertions for the response body or other expectations
	})

	t.Run("Test Base64 Page - Decode", func(t *testing.T) {
		// Create a test request to the "/base64" route with action=decode and encoded=...
		req := httptest.NewRequest(http.MethodGet, "/base64?action=decode&encoded=...", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("Failed to send test request: %v", err)
		}

		// Check the response status code
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		// TODO: Add assertions for the response body or other expectations
	})
}

func TestDefineResources(t *testing.T) {
	app := newApp()
	defineResources(app)

	t.Run("Test Static /css", func(t *testing.T) {
		// Create a test request to the "/css" route
		req := httptest.NewRequest(http.MethodGet, "/css/bootstrap.min.css", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("Failed to send test request: %v", err)
		}

		// Check the response status code
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		// TODO: Add assertions for the response body or other expectations
	})

	t.Run("Test Static /js", func(t *testing.T) {
		// Create a test request to the "/js" route
		req := httptest.NewRequest(http.MethodGet, "/js/bootstrap.min.js", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("Failed to send test request: %v", err)
		}

		// Check the response status code
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		// TODO: Add assertions for the response body or other expectations
	})

	t.Run("Test Static /img", func(t *testing.T) {
		// Create a test request to the "/img" route
		req := httptest.NewRequest(http.MethodGet, "/img/logo-transparent.png", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("Failed to send test request: %v", err)
		}

		// Check the response status code
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		// TODO: Add assertions for the response body or other expectations
	})
}
