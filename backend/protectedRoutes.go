package backend

// import (
// 	"time"

// 	"github.com/floriwan/srcm/backend/middleware"
// 	"github.com/floriwan/srcm/pkg/jwtauth"
// 	"github.com/lestrrat-go/jwx/v2/jwt"

// 	"github.com/go-chi/chi"
// )

// func (b *Backend) protectedRoutes(r chi.Router) {
// 	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil, jwt.WithAcceptableSkew(30*time.Second))
// 	r.Use(jwtauth.Verifier(tokenAuth))

// 	r.Group(b.adminRoutes)
// 	r.Group(b.dataRoutes)
// }

// func (b *Backend) adminRoutes(r chi.Router) {
// 	r.Use(middleware.AdminAuthenticator)
// 	r.Get("/admin", b.handler.Admin)
// 	r.Get("/user/{id}", b.handler.GetUser)
// 	r.Delete("/user/{id}", b.handler.DeleteUser)
// 	r.Post("/user", b.handler.CreateUser)

// 	r.Post("/event/season", b.handler.AddSeason)
// 	r.Get("/event/season", b.handler.GetSeason)
// 	r.Post("/event/season/{id}/race", b.handler.AddRace)

// 	r.Post("/event/race/{id}/results", b.handler.AddResults)
// }

// func (b *Backend) dataRoutes(r chi.Router) {
// 	r.Use(middleware.UserAuthenticator)
// 	r.Get("/user/{id}", b.handler.GetUser)

// }
