package backend

import "github.com/go-chi/chi"

func publicRoutes(r chi.Router) {
	r.Get("/", h.Login)
	r.Post("/login", h.Login)
}
