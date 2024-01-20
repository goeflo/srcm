package backend

import (
	"log"
	"net/http"
	"time"

	"github.com/floriwan/srcm/backend/handler"
	"github.com/floriwan/srcm/pkg/config"
	"github.com/floriwan/srcm/pkg/db"
	"github.com/floriwan/srcm/pkg/db/migration"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type Backend struct {
	handler *handler.Handler
}

// Initialize will create database with initial data and load config
func Initialize() *Backend {

	// load config
	err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load configation file", err)
	}

	// create and populate database
	db.Initialize()
	m := migration.NewMigrator(db.Instance)
	m.Migration()
	db.PolulateInitialData()

	// create new handler for routes
	h := handler.Handler{
		DB:     db.Instance,
		Config: config.GlobalConfig,
	}

	return &Backend{
		handler: &h,
	}
}

// Run start backend http server with config
func (b *Backend) Run() {

	// load all html templates
	// if err := h.Tmpl.Load("./templates/", ".tmpl"); err != nil {
	// 	log.Fatal(err)
	// }

	r := chi.NewRouter()

	// middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Group(b.publicRoutes)
	r.Group(b.protectedRoutes)

	addr := ":" + b.handler.Config.RestPort
	log.Printf("starting server on %v\n", addr)
	http.ListenAndServe(addr, r)
}
