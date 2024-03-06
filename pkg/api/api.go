package api

import (
	"encoding/json"
	"io/fs"
	"log"
	"net/http"

	"github.com/floriwan/srcm/pkg/jwtauth"
	"github.com/floriwan/srcm/pkg/store"
	"github.com/gorilla/mux"
)

type APIServer struct {
	apiPort        string
	staticHtmlPort string
	store          store.Store
	jwtAuth        *jwtauth.JWTAuth
	public         fs.FS
}

func NewAPIServer(htmlPort string, apiPort string, store store.Store, public fs.FS) *APIServer {
	return &APIServer{
		apiPort:        apiPort,
		staticHtmlPort: htmlPort,
		store:          store,
		//jwtAuth: jwtauth.New("HS256", []byte("secret"), nil, jwt.WithAcceptableSkew(30*time.Second)),
		jwtAuth: jwtauth.New("HS256", []byte("secret"), nil),
		public:  public,
	}
}

func (s *APIServer) Serve() {

	router := mux.NewRouter()
	fs := http.FileServer(http.FS(s.public))
	http.Handle("/", fs)

	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userService := NewUserService(s.store, s.jwtAuth)
	userService.RegisterRoutes(subrouter)

	eventService := NewEventService(s.store, s.jwtAuth)
	eventService.RegisterRoutes(subrouter)

	// start the API server
	go func() {
		addr := ":" + s.apiPort
		log.Printf("starting api server on %v\n", addr)
		log.Fatal(http.ListenAndServe(addr, subrouter))
	}()

	// start static file server
	addr := ":" + s.staticHtmlPort
	log.Printf("starting http server on %v\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))

}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

type ErrorResponse struct {
	Error string `json:"error"`
}
