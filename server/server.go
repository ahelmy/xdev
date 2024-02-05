package server

import (
	"embed"
	"errors"
	"io/fs"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/ahelmy/xdev/api"
	"github.com/ahelmy/xdev/app"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/filesystem"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

const (
	Prefix         = "ui/"
	APIPrefix      = "/api"
	MainLayout     = Prefix + "layouts/main"
	BasePath       = "ui"
	JSONPath       = "/json"
	YAMLPath       = "/yaml"
	JWTPath        = "/jwt"
	UUIDPath       = "/uuid"
	ULIDPath       = "/ulid"
	PasswordPath   = "/password"
	Base64Path     = "/base64"
	URLPath        = "/url"
	HashPath       = "/hash"
	TimePath       = "/time"
	PropertiesPath = "/properties"
	TextsPath      = "/text"
)

//go:embed ui/*
var uiFS embed.FS

//go:embed ui/css/*
var cssFS embed.FS

//go:embed ui/js/*
var jsFS embed.FS

//go:embed ui/img/*
var imgFS embed.FS

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type JWTEncodeRequest struct {
	Headers map[string]interface{} `json:"headers"`
	Claims  map[string]interface{} `json:"claims"`
	Secret  string                 `json:"secret"`
}

type JWTDecodeRequest struct {
	JWT string `json:"jwt"`
}

type YamlRequest struct {
	Yaml string `json:"yaml"`
}

type Base64Request struct {
	String string `json:"string"`
}

type URLRequest struct {
	URL string `json:"url"`
}

type HashRequest struct {
	String string `json:"string"`
}

type EpochTimeRequest struct {
	Epoch string `json:"epoch"`
}

type DateTimeRequest struct {
	DateTime string `json:"datetime"`
}

var errorHandler = func(ctx fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	errorStr := ""
	// Send custom error page
	if err != nil {
		errorStr = err.Error()
	}
	if strings.HasPrefix(ctx.Route().Path, APIPrefix) {
		return ctx.JSON(Response{Success: false, Message: strconv.Itoa(code) + ": " + errorStr})
	}
	// Return from handler
	return ctx.Render(Prefix+"error", newMap(map[string]any{
		"ErrorCode":    code,
		"ErrorMessage": errorStr,
	}), MainLayout)
}

func newApp() *fiber.App {
	engine := html.NewFileSystem(http.FS(uiFS), ".html")
	app := fiber.New(fiber.Config{
		Views:          engine,
		ReadBufferSize: 4096 * 3000,
		// Override default error handler
		ErrorHandler: errorHandler,
	})
	api.AddAPILayer(app)
	return app
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
	uuidPage(app)
	ulidPage(app)
	passwordPage(app)
	jwtPage(app)
	yamlPage(app)
	jsonPage(app)
	base64Page(app)
	urlPage(app)
	hashPage(app)
	timePage(app)
	propertiesPage(app)
	compareTextPage(app)
	log.Fatal(app.Listen(":"+strconv.FormatInt(int64(port), 10), fiber.ListenConfig{EnablePrefork: true}))
}

func indexPage(app *fiber.App) {
	app.Get("/", func(c fiber.Ctx) error {
		return c.Render(Prefix+"index", newMap(map[string]any{
			"Title": "Hello, Xdev!",
		}), MainLayout)
	})
}

func uuidPage(app *fiber.App) {
	app.Get(UUIDPath, func(c fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render(Prefix+"uuid", newMap(map[string]any{
			"Title": "UUID Generator",
		}), MainLayout)
	})
}

func ulidPage(app *fiber.App) {
	app.Get(ULIDPath, func(c fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render(Prefix+"ulid", newMap(map[string]any{
			"Title": "ULID Generator",
		}), MainLayout)
	})
}
func passwordPage(app *fiber.App) {
	app.Get(PasswordPath, func(c fiber.Ctx) error {
		return c.Render(Prefix+"password", newMap(map[string]any{
			"Title": "Password Generator",
		}), MainLayout)
	})
}

func yamlPage(app *fiber.App) {
	app.Get(YAMLPath, func(c fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render(Prefix+"yaml", newMap(map[string]any{
			"Title": "YAML",
		}), MainLayout)
	})
}

func jsonPage(app *fiber.App) {
	app.Get(JSONPath, func(c fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render(Prefix+"json", newMap(map[string]any{
			"Title": "JSON",
		}), MainLayout)
	})
}

func jwtPage(app *fiber.App) {
	app.Get(JWTPath, func(c fiber.Ctx) error {
		return c.Render(Prefix+"jwt", newMap(map[string]any{
			"Title": "JWT",
		}), MainLayout)
	})
}

func base64Page(app *fiber.App) {
	app.Get(Base64Path, func(c fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render(Prefix+"base64", newMap(map[string]any{
			"Title": "Base64 encode/decode",
		}), MainLayout)
	})
}

func urlPage(app *fiber.App) {
	app.Get(URLPath, func(c fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render(Prefix+"url", newMap(map[string]any{
			"Title": "URL encode/decode",
		}), MainLayout)
	})
}

func hashPage(app *fiber.App) {
	app.Get(HashPath, func(c fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render(Prefix+"hash", newMap(map[string]any{
			"Title": "Hashing",
		}), MainLayout)
	})
}

func timePage(app *fiber.App) {
	app.Get(TimePath, func(c fiber.Ctx) error {
		return c.Render(Prefix+"time", newMap(map[string]any{
			"Title": "Time Converter",
		}), MainLayout)
	})
}

func propertiesPage(app *fiber.App) {
	app.Get(PropertiesPath, func(c fiber.Ctx) error {
		return c.Render(Prefix+"properties", newMap(map[string]any{
			"Title": "Properties",
		}), MainLayout)
	})

}
func compareTextPage(app *fiber.App) {
	app.Get(TextsPath, func(c fiber.Ctx) error {
		return c.Render(Prefix+"text", newMap(map[string]any{
			"Title": "Text BS",
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
