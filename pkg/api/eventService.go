package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/floriwan/srcm/pkg/jwtauth"
	"github.com/floriwan/srcm/pkg/store"
	"github.com/gorilla/mux"
)

type NewSeasonData struct {
	Name string `json:"name"`
}

type NewRaceData struct {
	RaceName string `json:"raceName"`
	SeasonID uint   `json:"seasonID"`
}

type NewDriverData struct {
	Name     string `json:"name"`
	TeamName string `json:"teamName"`
	UserID   uint   `json:"userID"`
}

type EventService struct {
	store   store.Store
	jwtAuth *jwtauth.JWTAuth
}

func NewEventService(s store.Store, auth *jwtauth.JWTAuth) *EventService {
	return &EventService{store: s, jwtAuth: auth}
}

func (e *EventService) RegisterRoutes(r *mux.Router) {
	r.Use(jwtauth.Verifier(e.jwtAuth))
	r.Use(jwtauth.Authenticator(e.jwtAuth))
	r.HandleFunc("/event/season", e.handleCreateSeason).Methods("POST")
	r.HandleFunc("/event/season/{id}", e.handleGetSeason).Methods("GET")
	r.HandleFunc("/event/season", e.handleGetSeasonList).Methods("GET")

	r.HandleFunc("/event/race", e.handleCreateRace).Methods("POST")
	r.HandleFunc("/event/race/{id}", e.handleGetRace).Methods("GET")

	r.HandleFunc("/event/race/{id}/driver", e.handleAddDriver).Methods("POST")
	r.HandleFunc("/event/race/{id}/driver", e.handleGetDriver).Methods("GET")
}

func (s *EventService) handleGetDriver(w http.ResponseWriter, r *http.Request) {
}

func (s *EventService) handleAddDriver(w http.ResponseWriter, r *http.Request) {
	log.Printf("handleAddDriver\n")
	vars := mux.Vars(r)
	if vars["id"] == "" {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "no race ID found in request"})
		return
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	driverData := []NewDriverData{}
	if err := json.Unmarshal(body, &driverData); err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	log.Printf("get race ID %v\n", id)
	race, err := s.store.GetRaceByID(uint(id))
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "no race found with ID " + vars["id"]})
		return
	}

	for _, driver := range driverData {
		log.Printf("add driver %v (%v) to race '%v'\n", driver.Name, driver.TeamName, race.Name)
		driver, err := s.store.AddDriver(driver.Name, driver.TeamName, uint(id))
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "can not add driver '" + driver.Name + "' to race '" + race.Name + "' " + err.Error()})
			return
		}
	}
}

func (s *EventService) handleGetRace(w http.ResponseWriter, r *http.Request) {
	log.Printf("handleGetRace\n")
	vars := mux.Vars(r)
	if vars["id"] == "" {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "no race ID found in request"})
		return
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, err)
		return
	}

	race, err := s.store.GetRaceByID(uint(id))
	if err != nil {
		WriteJSON(w, http.StatusNotFound, err)
		return
	}
	WriteJSON(w, http.StatusOK, race)
}

func (s *EventService) handleCreateRace(w http.ResponseWriter, r *http.Request) {
	log.Printf("handleCreateRace\n")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	d := &NewRaceData{}
	if err := json.Unmarshal(body, d); err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	race, err := s.store.AddRace(d.RaceName, 3)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "unable to create race '" + d.RaceName + "' " + err.Error()})
		return
	}

	WriteJSON(w, http.StatusCreated, race)

}

func (s *EventService) handleGetSeasonList(w http.ResponseWriter, r *http.Request) {
	log.Printf("handleGetSeasonList\n")
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

	season, err := s.store.AddSeason(d.Name)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "unable to create season '" + d.Name + "' " + err.Error()})
		return
	}

	WriteJSON(w, http.StatusCreated, season)

}

func (s *EventService) handleGetSeason(w http.ResponseWriter, r *http.Request) {
	log.Printf("handleGetSeason\n")
	vars := mux.Vars(r)
	if vars["id"] == "" {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "no season ID found in request"})
		return
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	season, err := s.store.GetSeasonByID(uint(id))
	if err != nil {
		WriteJSON(w, http.StatusNotFound, ErrorResponse{Error: err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, season)
}
