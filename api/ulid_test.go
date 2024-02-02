package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestULIDAPI(t *testing.T) {
	app := fiber.New()
	ulidAPI(app)
	t.Run("Test ULID API", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/ulid", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var ulidResponse Response
		err = json.NewDecoder(resp.Body).Decode(&ulidResponse)
		assert.NoError(t, err)
		assert.True(t, ulidResponse.Success)
		assert.NotNil(t, ulidResponse.Data.(map[string]interface{})["ulid"])
	})
	t.Run("Convert ULID to UUID", func(t *testing.T) {
		ulid := "461MR407H385QT3ZFMSK3Q70SX"
		expectedUUID := "860d3040-1e23-416f-a1fd-f4ccc773833d"

		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/ulid/convert?ulid=%s&to=uuid", ulid), nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		body, _ := io.ReadAll(resp.Body)
		var result Response
		json.Unmarshal(body, &result)

		if !result.Success {
			t.Errorf("Expected success to be true, but got false")
		}

		if _, ok := result.Data.(map[string]interface{})["result"]; !ok {
			t.Errorf("Expected 'result' field in response data, but not found")
		}

		ulid, ok := result.Data.(map[string]interface{})["result"].(string)
		if !ok {
			t.Errorf("Expected 'result' field to be a string, but got %T", result.Data.(map[string]interface{})["result"])
		}

		if ulid != expectedUUID {
			t.Errorf("Expected ULID to be %s, but got %s", expectedUUID, ulid)
		}
	})

	t.Run("Invalid ULID", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/ulid/convert?ulid=ddddddc773833", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		body, _ := io.ReadAll(resp.Body)
		var result Response
		json.Unmarshal(body, &result)

		if result.Success {
			t.Errorf("Expected success to be false, but got true")
		}
	})

	t.Run("Missing ULID", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/ulid/convert?to=uuid", nil)
		resp, _ := app.Test(req)
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

	})

	t.Run("Invalid conversion type", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/ulid/convert?ulid=461MR407H385QT3ZFMSK3Q70SX&to=invalid", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		body, _ := io.ReadAll(resp.Body)
		var result Response
		json.Unmarshal(body, &result)

		if result.Success {
			t.Errorf("Expected success to be false, but got true")
		}

		if result.Message != "invalid conversion type" {
			t.Errorf("Expected message to be 'invalid conversion type', but got '%s'", result.Message)
		}
	})
}
