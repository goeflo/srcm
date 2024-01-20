package backend

import (
	"github.com/go-chi/chi"
)

func (b *Backend) publicRoutes(r chi.Router) {
	r.Get("/", b.handler.Home)
	r.Post("/login", b.handler.Login)
	r.Post("/user", b.handler.CreateUser)
}
