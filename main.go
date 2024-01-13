package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/floriwan/srcm/handler"
	"github.com/floriwan/srcm/pkg/config"
	"github.com/floriwan/srcm/pkg/db"
	"github.com/floriwan/srcm/pkg/jwtauth"
	"github.com/floriwan/srcm/pkg/templates"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

func main() {

	// load config
	err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load configation file", err)
	}

	// create database
	db.Initialize()
	db.Migrate()
	db.PolulateInitialData()

	// initialize the handler
	h := &handler.Handler{
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

	// public route
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})
	r.Post("/login", h.Login)
	r.Mount("/home", templateRouter(h))
	r.Mount("/assets", assetsServer(h))

	// protected route
	r.Mount("/restricted", restrictedRouter())

	addr := ":" + config.GlobalConfig.RestPort
	log.Printf("starting server on %v\n", addr)
	http.ListenAndServe(addr, r)
}

func assetsServer(h *handler.Handler) chi.Router {
	r := chi.NewRouter()
	r.Get("/", h.Assets)
	r.Get("/css/*", h.Css)
	return r
}

func templateRouter(h *handler.Handler) chi.Router {
	r := chi.NewRouter()
	r.Get("/*", h.Home)
	return r
}

func restrictedRouter() chi.Router {
	r := chi.NewRouter()

	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil, jwt.WithAcceptableSkew(30*time.Second))

	r.Use(jwtauth.Verifier(tokenAuth))
	r.Use(jwtauth.Authenticator(tokenAuth))

	/*
		curl example for restricted access
		curl localhost:8081/restricted/user/bla -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluIn0.A9dz8H4vRCdMb39m6nOlnl_HbF5zgof5LrLm2i0xEY0"
	*/

	r.Get("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("restricted: get user id %v", chi.URLParam(r, "id"))))
	})
	return r
}
