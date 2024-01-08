package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	srcm_config "github.com/floriwan/srcm/pkg/config"
	srcm_db "github.com/floriwan/srcm/pkg/db"
	srcm_handler "github.com/floriwan/srcm/pkg/handler"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

func main() {

	// load config
	err := srcm_config.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load configation file", err)
	}

	// create database
	srcm_db.Initialize()
	srcm_db.Migrate()

	// initialize the handler
	h := &srcm_handler.Handler{DB: srcm_db.Instance}

	// create echo instance with middleware
	e := echo.New()
	if srcm_config.GlobalConfig.LogLevel == "debug" {
		e.Debug = true
	}
	e.HideBanner = true
	e.Use(middleware.Recover())

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = renderer

	// routes
	e.GET("/", h.Root)
	e.POST("/login", h.Login)
	e.POST("/login", h.Register)

	g := e.Group("/restricted")

	g.Use(echojwt.WithConfig(
		echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(srcm_handler.JwtCustomClaims)
			},
			SigningKey: []byte("secret")},
	))

	g.GET("", restricted)
	g.POST("/users", h.CreateUser)
	g.GET("/users", h.GetAllUsers)
	g.GET("/users/:id", h.GetUser)
	g.PUT("/users/:id", h.UpdateUser)
	e.DELETE("/users/:id", h.DeleteUser)

	e.Logger.Fatal(e.Start(":" + srcm_config.GlobalConfig.HttpPort))

}

// test restricted access
func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*srcm_handler.JwtCustomClaims)
	name := claims.Email
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}
