package main

import (
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

	// routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/login", h.Login)

	g := e.Group("/restricted")

	g.Use(echojwt.WithConfig(
		echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(srcm_handler.JwtCustomClaims)
			},
			SigningKey: []byte("secret")},
	))

	g.GET("", restricted)
	g.GET("/users/:id", h.GetUser)
	g.POST("/users/:id/update", h.UpdateUser)
	g.POST("/users/new", h.NewUser)
	g.GET("/users/list", h.GetUserList)

	e.Logger.Fatal(e.Start(":" + srcm_config.GlobalConfig.HttpPort))

}

// test restricted access
func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*srcm_handler.JwtCustomClaims)
	name := claims.Email
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
