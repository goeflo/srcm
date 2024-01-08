package main

import (
	"log"
	"net/http"

	srcm_config "github.com/floriwan/srcm/pkg/config"
	srcm_db "github.com/floriwan/srcm/pkg/db"
	srcm_router "github.com/floriwan/srcm/router"
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

	e.POST("/login", srcm_router.Login)

	g := e.Group("/restricted")

	g.Use(echojwt.WithConfig(
		echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(srcm_router.JwtCustomClaims)
			},
			SigningKey: []byte("secret")},
	))

	g.GET("", restricted)
	g.GET("/users/:id", srcm_router.GetUser)
	g.POST("/users/:id/update", srcm_router.UpdateUser)
	g.POST("/users/new", srcm_router.NewUser)
	g.GET("/users/list", srcm_router.GetUserList)

	e.Logger.Fatal(e.Start(":" + srcm_config.GlobalConfig.HttpPort))

	/*
		log.Printf("starting ...")

		err := config.LoadConfig(".")
		if err != nil {
			log.Fatal("could not load configation file", err)
		}

		// create database
		db.Initialize()
		db.Migrate()

		// Initialize Router
		router := initRouter()
		router.Run(":" + config.GlobalConfig.HttpPort)
	*/
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*srcm_router.JwtCustomClaims)
	name := claims.Email
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
