package backend

import (
	"time"

	"github.com/floriwan/srcm/backend/handler"
	"github.com/floriwan/srcm/backend/middleware"
	"github.com/floriwan/srcm/pkg/jwtauth"
	"github.com/lestrrat-go/jwx/v2/jwt"

	"github.com/go-chi/chi"
)

func protectedRoutes(r chi.Router) {
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil, jwt.WithAcceptableSkew(30*time.Second))
	r.Use(jwtauth.Verifier(tokenAuth))
	//r.Use(jwtauth.Authenticator(tokenAuth))
	r.Group(adminRoutes)
	r.Group(dataRoutes)
}

func adminRoutes(r chi.Router) {
	r.Use(middleware.AdminAuthenticator)
	r.Get("/admin", handler.Admin)
	r.Get("/user/{id}", handler.GetUser)
	r.Delete("/user/{id}", handler.DeleteUser)
	r.Post("/user", handler.CreateUser)

	//r.Post("/event", handler.CreateEvent)
	//r.Get("/event/{id}", handler.GetEvent)

}

func dataRoutes(r chi.Router) {
	r.Use(middleware.UserAuthenticator)
	r.Get("/user/{id}", handler.GetUser)
}
