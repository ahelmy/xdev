package server

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"strconv"

	"github.com/ahelmy/xdev/app"
	"github.com/ahelmy/xdev/internal"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/filesystem"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

const (
	Prefix       = "ui/"
	MainLayout   = Prefix + "layouts/main"
	BasePath     = "ui"
	JSONPath     = "/json"
	YAMLPath     = "/yaml"
	JWTPath      = "/jwt"
	UUIDPath     = "/uuid"
	ULIDPath     = "/ulid"
	PasswordPath = "/password"
	Base64Path   = "/base64"
	URLPath      = "/url"
	HashPath     = "/hash"
)

//go:embed ui/*
var uiFS embed.FS

//go:embed ui/css/*
var cssFS embed.FS

//go:embed ui/js/*
var jsFS embed.FS

//go:embed ui/img/*
var imgFS embed.FS

func newApp() *fiber.App {
	engine := html.NewFileSystem(http.FS(uiFS), ".html")

	return fiber.New(fiber.Config{
		Views:          engine,
		ReadBufferSize: 4096 * 3,
	})
}

func newMap(fiberMap map[string]any) fiber.Map {
	mp := fiber.Map{}
	mp["AppVersion"] = app.AppVersion
	for k, v := range fiberMap {
		mp[k] = v
	}
	return mp
}

func StartServer(port int32, isVerbose bool) {
	app := newApp()
	if isVerbose {
		app.Use(logger.New())
	}
	defineResources(app)
	indexPage(app)
	jsonPage(app)
	uuidPage(app)
	ulidPage(app)
	passwordPage(app)
	yamlPage(app)
	jwtPage(app)
	base64Page(app)
	urlPage(app)
	hashPage(app)

	log.Fatal(app.Listen(":"+strconv.FormatInt(int64(port), 10), fiber.ListenConfig{EnablePrefork: true}))
}

func indexPage(app *fiber.App) {
	app.Get("/", func(c fiber.Ctx) error {
		return c.Render(Prefix+"index", newMap(map[string]any{
			"Title": "Hello, Xdev!",
		}), MainLayout)
	})
}

func jsonPage(app *fiber.App) {
	app.Get(JSONPath, func(c fiber.Ctx) error {
		action := c.FormValue("action")
		errorStr := ""
		json := c.FormValue("json")
		result := ""
		if len(json) > 0 {
			localResult := ""
			err := error(nil)
			switch action {
			case "beautify":
				localResult, err = internal.IndentJSON(json)
			case "minify":
				localResult, err = internal.MinifyJSON(json)
			case "json2Yaml":
				localResult, err = internal.Json2Yaml(json)
			}
			if err != nil {
				errorStr = err.Error()
			} else {
				result = localResult
			}
		}

		// Render index within layouts/main
		return c.Render(Prefix+"json", newMap(map[string]any{
			"Title":  "JSON",
			"JSON":   c.FormValue("json"),
			"Result": result,
			"Error":  errorStr,
			"action": action,
		}), MainLayout)
	})
}

func uuidPage(app *fiber.App) {
	app.Get(UUIDPath, func(c fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render(Prefix+"uuid", newMap(map[string]any{
			"Title": "UUID Generator",
			"UUID":  internal.GenerateGUID(),
		}), MainLayout)
	})
}

func ulidPage(app *fiber.App) {
	app.Get(ULIDPath, func(c fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render(Prefix+"ulid", newMap(map[string]any{
			"Title": "ULID Generator",
			"ULID":  internal.GenerateULID(),
		}), MainLayout)
	})
}

func passwordPage(app *fiber.App) {
	app.Get(PasswordPath, func(c fiber.Ctx) error {
		// Render index within layouts/main
		length, err := strconv.Atoi(c.FormValue("length", "10"))
		if err != nil {
			length = 10
		}
		isEspecial := c.FormValue("especial") == "on"
		isCapital := c.FormValue("capital") == "on"
		isNumberic := c.FormValue("number") == "on"
		capital := "on"
		if !isCapital {
			capital = "off"
		}
		especial := "on"
		if !isEspecial {
			especial = "off"
		}
		number := "on"
		if !isNumberic {
			number = "off"
		}
		password := internal.GeneratePassword(length, isEspecial, isNumberic, isCapital)
		return c.Render(Prefix+"password", newMap(map[string]any{
			"Title":    "Password Generator",
			"Password": password,
			"length":   length,
			"especial": especial,
			"capital":  capital,
			"number":   number,
		}), MainLayout)
	})
}

func yamlPage(app *fiber.App) {
	app.Get(YAMLPath, func(c fiber.Ctx) error {
		action := c.FormValue("action")
		errorStr := ""
		yaml := c.FormValue("yaml")
		result := ""
		if len(yaml) > 0 {
			if action == "yaml2JSON" {
				_yaml, err := internal.Yaml2Json(yaml)
				if err != nil {
					errorStr = err.Error()
				} else {
					result = _yaml
				}
			}
		}

		// Render index within layouts/main
		return c.Render(Prefix+"yaml", newMap(map[string]any{
			"Title":  "YAML",
			"YAML":   c.FormValue("yaml"),
			"Result": result,
			"Error":  errorStr,
			"action": action,
		}), MainLayout)
	})
}

func jwtPage(app *fiber.App) {
	app.Get(JWTPath, func(c fiber.Ctx) error {
		jwtToken := c.FormValue("jwt")
		jwt := internal.JWT{}
		errorStr := ""
		if len(jwtToken) > 0 {
			_jwt, err := internal.DecodeJWT(c.FormValue("jwt"))
			if err != nil {
				errorStr = err.Error()
			} else {
				jwt = _jwt
			}
		}

		// Render index within layouts/main
		return c.Render(Prefix+"jwt", newMap(map[string]any{
			"Title":     "JWT",
			"JWT":       c.FormValue("jwt"),
			"Header":    jwt.Header,
			"Claims":    jwt.Claims,
			"Signature": jwt.Signature,
			"Error":     errorStr,
		}), MainLayout)
	})
}

func base64Page(app *fiber.App) {
	app.Get(Base64Path, func(c fiber.Ctx) error {
		action := c.FormValue("action")

		decoded := ""
		encoded := ""
		if action == "encode" {
			encoded = internal.EncodeToBase64(c.FormValue("decoded"))
			decoded = c.FormValue("decoded")
		} else if action == "decode" {
			decoded = internal.DecodeFromBase64(c.FormValue("encoded"))
			encoded = c.FormValue("encoded")
		}
		errorStr := ""

		// Render index within layouts/main
		return c.Render(Prefix+"base64", newMap(map[string]any{
			"Title":   "Base64 encode/decode",
			"Decoded": decoded,
			"Encoded": encoded,
			"Error":   errorStr,
		}), MainLayout)
	})
}

func urlPage(app *fiber.App) {
	app.Get(URLPath, func(c fiber.Ctx) error {
		action := c.FormValue("action")

		decoded := ""
		encoded := ""
		errorStr := ""

		if action == "encode" {
			encoded = internal.EncodeURL(c.FormValue("decoded"))
			decoded = c.FormValue("decoded")
		} else if action == "decode" {
			_decoded, err := internal.DecodeURL(c.FormValue("encoded"))
			if err != nil {
				errorStr = err.Error()
			} else {
				decoded = _decoded
			}
			encoded = c.FormValue("encoded")
		}

		// Render index within layouts/main
		return c.Render(Prefix+"url", newMap(map[string]any{
			"Title":   "URL encode/decode",
			"Decoded": decoded,
			"Encoded": encoded,
			"Error":   errorStr,
		}), MainLayout)
	})
}

func hashPage(app *fiber.App) {
	app.Get(HashPath, func(c fiber.Ctx) error {
		algorithm := c.FormValue("action")

		str := c.FormValue("string")
		result := ""
		errorStr := ""
		if len(str) > 0 {
			_result, err := internal.HashString(str, algorithm)
			if err != nil {
				errorStr = err.Error()
			} else {
				result = _result
			}
		}

		// Render index within layouts/main
		return c.Render(Prefix+"hash", newMap(map[string]any{
			"Title":  "Hashing",
			"String": str,
			"Result": result,
			"Error":  errorStr,
		}), MainLayout)
	})
}

func defineResources(app *fiber.App) {
	app.Use("/css", filesystem.New(filesystem.Config{
		Root:       fs.FS(cssFS),
		PathPrefix: "ui/css",
		Browse:     true,
	}))
	app.Use("/js", filesystem.New(filesystem.Config{
		Root:       fs.FS(jsFS),
		PathPrefix: "ui/js",
		Browse:     true,
	}))
	app.Use("/img", filesystem.New(filesystem.Config{
		Root:       fs.FS(imgFS),
		PathPrefix: "ui/img",
		Browse:     true,
	}))
}
