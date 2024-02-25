package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/floriwan/srcm/pkg/jwtauth"
	"github.com/floriwan/srcm/pkg/store"
	"github.com/gorilla/mux"
)

type APIServer struct {
	port    string
	store   store.Storage
	jwtAuth *jwtauth.JWTAuth
}

func NewAPIServer(port string, store store.Storage) *APIServer {
	return &APIServer{
		port:  port,
		store: store,
		//jwtAuth: jwtauth.New("HS256", []byte("secret"), nil, jwt.WithAcceptableSkew(30*time.Second)),
		jwtAuth: jwtauth.New("HS256", []byte("secret"), nil),
	}
}

func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userService := NewUserService(s.store, s.jwtAuth)
	userService.RegisterRoutes(subrouter)

	eventService := NewEventService(s.store, s.jwtAuth)
	eventService.RegisterRoutes(subrouter)

	addr := ":" + s.port
	log.Printf("starting server on %v\n", addr)
	log.Fatal(http.ListenAndServe(addr, subrouter))

}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

type ErrorResponse struct {
	Error string `json:"error"`
}
