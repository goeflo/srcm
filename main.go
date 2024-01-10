package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/floriwan/srcm/handler"
	"github.com/floriwan/srcm/pkg/config"
	"github.com/floriwan/srcm/pkg/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {

	// create database
	db.Initialize()
	db.Migrate()

	// load config
	err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load configation file", err)
	}

	// initialize the handler
	h := &handler.Handler{DB: db.Instance, Config: config.GlobalConfig}

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

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	r.Post("/login", h.Login)

	// access after login
	r.Mount("/restricted", restrictedRouter())

	addr := ":" + config.GlobalConfig.RestPort
	fmt.Printf("starting server on %v\n", addr)
	http.ListenAndServe(addr, r)
}

func restrictedRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(loginOnly)
	r.Get("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("restricted: get user id %v", chi.URLParam(r, "id"))))
	})
	return r
}

func loginOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return

		// isAdmin, ok := r.Context().Value("acl.admin").(bool)
		// if !ok || !isAdmin {
		// 	http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		// 	return
		// }
		// next.ServeHTTP(w, r)
	})
}

// import (
// 	"context"
// 	"html/template"
// 	"io"
// 	"net/http"
// 	"os"
// 	"os/signal"
// 	"time"

// 	srcm_config "github.com/floriwan/srcm/pkg/config"
// 	srcm_db "github.com/floriwan/srcm/pkg/db"
// 	srcm_handler "github.com/floriwan/srcm/pkg/handler"
// 	"github.com/golang-jwt/jwt/v5"
// 	echojwt "github.com/labstack/echo-jwt/v4"
// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"
// 	"github.com/labstack/gommon/log"
// )

// // TemplateRenderer is a custom html/template renderer for Echo framework
// type TemplateRenderer struct {
// 	templates *template.Template
// }

// func main() {

// 	// load config
// 	err := srcm_config.LoadConfig(".")
// 	if err != nil {
// 		log.Fatal("could not load configation file", err)
// 	}

// 	// create database
// 	srcm_db.Initialize()
// 	srcm_db.Migrate()

// 	// rest api server
// 	restServer := echo.New()
// 	go func() {
// 		runRestServer(restServer)
// 	}()

// 	// html page server
// 	htmlServer := echo.New()
// 	go func() {
// 		runHtmlServer(htmlServer)
// 	}()

// 	// graceful shutdown
// 	quit := make(chan os.Signal, 1)
// 	signal.Notify(quit, os.Interrupt)
// 	<-quit
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	if err := restServer.Shutdown(ctx); err != nil {
// 		restServer.Logger.Fatal(err)
// 	}
// 	if err := htmlServer.Shutdown(ctx); err != nil {
// 		htmlServer.Logger.Fatal(err)
// 	}
// }

// func runRestServer(e *echo.Echo) {

// 	// initialize the handler
// 	h := &srcm_handler.Handler{DB: srcm_db.Instance}

// 	// create echo instance with middleware for rest api
// 	//e := echo.New()
// 	e.Logger.SetLevel(log.DEBUG)
// 	e.HideBanner = true
// 	e.Use(middleware.Recover())
// 	/*e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
// 		Root:   "public",
// 		Browse: true,
// 	}))*/

// 	// renderer := &TemplateRenderer{
// 	// 	templates: template.Must(template.ParseGlob("public/views/*.html")),
// 	// }
// 	// e.Renderer = renderer

// 	// routes
// 	e.GET("/", h.Root)
// 	e.POST("/login", h.Login)
// 	//e.POST("/login", h.Register)

// 	g := e.Group("/restricted")

// 	g.Use(echojwt.WithConfig(
// 		echojwt.Config{
// 			NewClaimsFunc: func(c echo.Context) jwt.Claims {
// 				return new(srcm_handler.JwtCustomClaims)
// 			},
// 			SigningKey: []byte("secret")},
// 	))

// 	g.GET("", restricted)
// 	g.POST("/users", h.CreateUser)
// 	g.GET("/users", h.GetAllUsers)
// 	g.GET("/users/:id", h.GetUser)
// 	g.PUT("/users/:id", h.UpdateUser)
// 	e.DELETE("/users/:id", h.DeleteUser)

// 	if err := e.Start(":" + srcm_config.GlobalConfig.RestPort); err != nil && err != http.ErrServerClosed {
// 		e.Logger.Fatal("shutting down rest server")
// 	}

// }

// func runHtmlServer(e *echo.Echo) {
// 	//htmlServer := echo.New()
// 	e.Logger.SetLevel(log.DEBUG)
// 	e.HideBanner = true

// 	if err := e.Start(":" + srcm_config.GlobalConfig.HttpPort); err != nil && err != http.ErrServerClosed {
// 		e.Logger.Fatal("shutting down rest server")
// 	}

// }

// // test restricted access
// func restricted(c echo.Context) error {
// 	user := c.Get("user").(*jwt.Token)
// 	claims := user.Claims.(*srcm_handler.JwtCustomClaims)
// 	name := claims.Email
// 	return c.String(http.StatusOK, "Welcome "+name+"!")
// }

// // Render renders a template document
// func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

// 	// Add global methods if data is a map
// 	if viewContext, isMap := data.(map[string]interface{}); isMap {
// 		viewContext["reverse"] = c.Echo().Reverse
// 	}

// 	return t.templates.ExecuteTemplate(w, name, data)
// }
