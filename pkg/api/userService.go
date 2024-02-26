package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/floriwan/srcm/pkg/jwtauth"
	"github.com/floriwan/srcm/pkg/store"
	"github.com/gorilla/mux"
)

type loginData struct {
	Email  string `json:"email"`
	Passwd string `json:"passwd"`
}

type UserService struct {
	store   store.Store
	jwtAuth *jwtauth.JWTAuth
}

func NewUserService(s store.Store, auth *jwtauth.JWTAuth) *UserService {
	return &UserService{
		store:   s,
		jwtAuth: auth,
	}
}

func (s *UserService) RegisterRoutes(r *mux.Router) {
	//r.Use(jwtauth.Verifier(s.jwtAuth))
	r.HandleFunc("/user/register", s.handleUserRegister).Methods("POST")
	r.HandleFunc("/user/login", s.handleUserLogin).Methods("POST")
}

func (s *UserService) handleUserRegister(w http.ResponseWriter, r *http.Request) {
	log.Printf("handleUserRegister\n")
	WriteJSON(w, http.StatusNotImplemented, ErrorResponse{Error: ""})
}

func (s *UserService) handleUserLogin(w http.ResponseWriter, r *http.Request) {
	log.Printf("handleUserLogin\n")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	d := &loginData{}
	err = json.Unmarshal(body, &d)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid request payload"})
		return
	}

	log.Printf("login request data %+v\n", d)

	user, err := s.store.GetUser("")
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "unknown user " + d.Email})
		return
	}

	if err := user.CheckPassword(d.Passwd); err != nil {
		WriteJSON(w, http.StatusUnauthorized, ErrorResponse{Error: "wrong password for email " + d.Email})
		return
	}

	_, tokenString, _ := s.jwtAuth.Encode(map[string]interface{}{"email": d.Email})
	WriteJSON(w, http.StatusCreated, tokenString)
}
