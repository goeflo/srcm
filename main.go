package main

import (
	"log"

	"github.com/floriwan/srcm/handler"
	"github.com/floriwan/srcm/pkg/config"
	"github.com/floriwan/srcm/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {
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
	router.Run(":8081")
}

func initRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	// https://github.com/gin-gonic/gin/blob/v1.9.0/docs/doc.md#dont-trust-all-proxies
	router.SetTrustedProxies([]string{"localhost"})

	// handle homepage templates
	router.LoadHTMLGlob("templates/**/*.tmpl")
	router.GET("/", handler.Homepage)
	router.GET("/login", handler.Login)
	router.Static("/css", "templates/css")
	return router
}
