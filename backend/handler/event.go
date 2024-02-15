package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/floriwan/srcm/pkg/db/model"
)

type NewSeasonData struct {
	Name string `json:"name"`
}

func (h *Handler) AddSeason(w http.ResponseWriter, r *http.Request) {
	log.Printf("add season handler")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	d := &NewSeasonData{}
	if err := json.Unmarshal(body, d); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	season := model.Season{Name: d.Name}
	res := h.DB.Create(&season)
	if res.Error != nil {
		log.Fatalf("error creating season '%v': %v\n", d.Name, res.Error)
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("new season %+v", res)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	b, err := json.Marshal(season)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Write(b)

	//respondError(w, http.StatusNotImplemented, "add season handler")
}

func (h *Handler) GetSeason(w http.ResponseWriter, r *http.Request) {
	respondError(w, http.StatusNotImplemented, "get season handler")

	seasons := []model.Season{}
	result := h.DB.Find(&seasons)
	log.Printf("seasons %+v\n", result)
	log.Printf("seasons %+v\n", seasons)

}

func (h *Handler) AddRace(w http.ResponseWriter, r *http.Request) {
	respondError(w, http.StatusNotImplemented, "add race handler")
}

func (h *Handler) AddResults(w http.ResponseWriter, r *http.Request) {
	respondError(w, http.StatusNotImplemented, "add race result handler")
}
