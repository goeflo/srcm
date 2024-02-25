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

type NewSeasonData struct {
	Name string `json:"name"`
}

type EventService struct {
	store   store.Storage
	jwtAuth *jwtauth.JWTAuth
}

func NewEventService(s store.Storage, auth *jwtauth.JWTAuth) *EventService {
	return &EventService{store: s, jwtAuth: auth}
}

func (e *EventService) RegisterRoutes(r *mux.Router) {
	r.Use(jwtauth.Verifier(e.jwtAuth))
	r.Use(jwtauth.Authenticator(e.jwtAuth))
	r.HandleFunc("/event/season", e.handleCreateSeason).Methods("POST")
	r.HandleFunc("/event/season/{id}", e.handleGetSeason).Methods("GET")
	r.HandleFunc("/event/season", e.handleGetSeasonList).Methods("GET")

	r.HandleFunc("/event/season/{id}/race", e.handleCreateRace).Methods("POST")
}

func (s *EventService) handleCreateRace(w http.ResponseWriter, r *http.Request) {
}

func (s *EventService) handleGetSeasonList(w http.ResponseWriter, r *http.Request) {
	seasons, err := s.store.GetSeasons()
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, seasons)
}

func (s *EventService) handleCreateSeason(w http.ResponseWriter, r *http.Request) {
	log.Printf("handleCreateSeason\n")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	d := &NewSeasonData{}
	if err := json.Unmarshal(body, d); err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	season, err := s.store.CreateSeason(d.Name)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "unable to create season " + d.Name + " " + err.Error()})
		return
	}

	WriteJSON(w, http.StatusCreated, season)

}

func (s *EventService) handleGetSeason(w http.ResponseWriter, r *http.Request) {
	log.Printf("handleGetSeason\n")
	WriteJSON(w, http.StatusNotImplemented, ErrorResponse{Error: ""})
}
