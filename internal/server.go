package internal

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

const (
	MainLayout = "layouts/main"

	JSONPath     = "/json"
	YAMLPath     = "/yaml"
	JWTPath      = "/jwt"
	UUIDPath     = "/uuid"
	ULIDPath     = "/ulid"
	PasswordPath = "/password"
	Base64Path   = "/base64"
)

func newApp() *fiber.App {
	engine := html.New("./ui", ".html")

	return fiber.New(fiber.Config{
		Views:          engine,
		ReadBufferSize: 4096 * 3,
		Prefork:        true,
	})
}

func StartServer(port int32) {
	app := newApp()
	defineResources(app)
	indexPage(app)
	jsonPage(app)
	uuidPage(app)
	ulidPage(app)
	passwordPage(app)
	yamlPage(app)
	jwtPage(app)
	base64Page(app)

	log.Fatal(app.Listen(":" + strconv.FormatInt(int64(port), 10)))
}

func indexPage(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, Xdev!",
		}, MainLayout)
	})
}

func jsonPage(app *fiber.App) {
	app.Get(JSONPath, func(c *fiber.Ctx) error {
		action := c.FormValue("action")
		errorStr := ""
		json := c.FormValue("json")
		result := ""
		if len(json) > 0 {
			localResult := ""
			err := error(nil)
			switch action {
			case "beautify":
				localResult, err = IndentJSON(json)
			case "minify":
				localResult, err = MinifyJSON(json)
			case "json2Yaml":
				localResult, err = Json2Yaml(json)
			}
			if err != nil {
				errorStr = err.Error()
			} else {
				result = localResult
			}
		}

		// Render index within layouts/main
		return c.Render("json", fiber.Map{
			"Title":  "JSON",
			"JSON":   c.FormValue("json"),
			"Result": result,
			"Error":  errorStr,
			"action": action,
		}, MainLayout)
	})
}

func uuidPage(app *fiber.App) {
	app.Get(UUIDPath, func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("uuid", fiber.Map{
			"Title": "UUID Generator",
			"UUID":  GenerateGUID(),
		}, MainLayout)
	})
}

func ulidPage(app *fiber.App) {
	app.Get(ULIDPath, func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("ulid", fiber.Map{
			"Title": "ULID Generator",
			"ULID":  GenerateULID(),
		}, MainLayout)
	})
}

func passwordPage(app *fiber.App) {
	app.Get(PasswordPath, func(c *fiber.Ctx) error {
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
		password := GeneratePassword(length, isEspecial, isNumberic, isCapital)
		return c.Render("password", fiber.Map{
			"Title":    "Password Generator",
			"Password": password,
			"length":   length,
			"especial": especial,
			"capital":  capital,
			"number":   number,
		}, MainLayout)
	})
}

func yamlPage(app *fiber.App) {
	app.Get(YAMLPath, func(c *fiber.Ctx) error {
		action := c.FormValue("action")
		errorStr := ""
		yaml := c.FormValue("yaml")
		result := ""
		if len(yaml) > 0 {
			if action == "beautify" {
				errorStr = "Not implemented yet"
			} else if action == "minify" {
				errorStr = "Not implemented yet"
			} else if action == "yaml2JSON" {
				_yaml, err := Yaml2Json(yaml)
				if err != nil {
					errorStr = err.Error()
				} else {
					result = _yaml
				}
			}
		}

		// Render index within layouts/main
		return c.Render("yaml", fiber.Map{
			"Title":  "YAML",
			"YAML":   c.FormValue("yaml"),
			"Result": result,
			"Error":  errorStr,
			"action": action,
		}, MainLayout)
	})
}

func jwtPage(app *fiber.App) {
	app.Get(JWTPath, func(c *fiber.Ctx) error {
		jwtToken := c.FormValue("jwt")
		jwt := JWT{}
		errorStr := ""
		if len(jwtToken) > 0 {
			_jwt, err := DecodeJWT(c.FormValue("jwt"))
			if err != nil {
				errorStr = err.Error()
			} else {
				jwt = _jwt
			}
		}

		// Render index within layouts/main
		return c.Render("jwt", fiber.Map{
			"Title":     "JWT",
			"JWT":       c.FormValue("jwt"),
			"Header":    jwt.Header,
			"Claims":    jwt.Claims,
			"Signature": jwt.Signature,
			"Error":     errorStr,
		}, MainLayout)
	})
}

func base64Page(app *fiber.App) {
	app.Get(Base64Path, func(c *fiber.Ctx) error {
		action := c.FormValue("action")

		decoded := ""
		encoded := ""
		if action == "encode" {
			encoded = EncodeToBase64(c.FormValue("decoded"))
			decoded = c.FormValue("decoded")
		} else if action == "decode" {
			decoded = DecodeFromBase64(c.FormValue("encoded"))
			encoded = c.FormValue("encoded")
		}
		errorStr := ""

		// Render index within layouts/main
		return c.Render("base64", fiber.Map{
			"Title":   "Base64 encode/decode",
			"Decoded": decoded,
			"Encoded": encoded,
			"Error":   errorStr,
		}, MainLayout)
	})
}

func defineResources(app *fiber.App) {
	app.Static("/css", "./ui/css")
	app.Static("/js", "./ui/js")
	app.Static("/img", "./ui/img")
}
