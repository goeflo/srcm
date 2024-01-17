package backend

import (
	"log"
	"net/http"
	"time"

	"github.com/floriwan/srcm/backend/handler"
	"github.com/floriwan/srcm/pkg/config"
	"github.com/floriwan/srcm/pkg/db"
	"github.com/floriwan/srcm/pkg/templates"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

var h *handler.Handler

// Run will initialize database and start and run http api server
// Using global config.GlobalConfig for configuration data
func Run() {
	// create database
	db.Initialize()
	db.Migrate()
	db.PolulateInitialData()

	// the handler for the routes
	h = &handler.Handler{
		DB:     db.Instance,
		Config: config.GlobalConfig,
		Tmpl:   templates.NewTmpl(),
	}

	// load all html templates
	if err := h.Tmpl.Load("./templates/", ".tmpl"); err != nil {
		log.Fatal(err)
	}

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

	r.Group(publicRoutes)
	r.Group(protectedRoutes)

	addr := ":" + config.GlobalConfig.RestPort
	log.Printf("starting server on %v\n", addr)
	http.ListenAndServe(addr, r)
}
