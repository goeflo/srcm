package backend

import "github.com/go-chi/chi"

func publicRoutes(r chi.Router) {
	r.Get("/", h.Home)
	r.Post("/login", h.Login)
}
