package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/floriwan/srcm/pkg/db/model"
	"github.com/go-chi/chi"
)

type NewSeasonData struct {
	Name string `json:"name"`
}

type NewRaceData struct {
	Name     string `json:"name"`
	SeasonID int    `json:"season_id"`
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
		log.Printf("error creating season '%v': %v\n", d.Name, res.Error)
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

}

func (h *Handler) GetSeason(w http.ResponseWriter, r *http.Request) {
	respondError(w, http.StatusNotImplemented, "get season handler")

	seasons := []model.Season{}
	result := h.DB.Find(&seasons)
	if result.Error != nil {
		respondError(w, http.StatusInternalServerError, result.Error.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	b, err := json.Marshal(seasons)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Write(b)
}

func (h *Handler) AddRace(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	d := &NewRaceData{}
	if err := json.Unmarshal(body, d); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	seasonID := chi.URLParam(r, "id")

	seasonInt, err := strconv.Atoi(seasonID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	race := model.Race{Name: d.Name, SeasonID: seasonInt}
	res := h.DB.Create(&race)
	if res.Error != nil {
		log.Printf("error creating season '%v': %v\n", d.Name, res.Error)
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("new race %+v", res)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	b, err := json.Marshal(race)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Write(b)

	//respondError(w, http.StatusNotImplemented, "add race handler")
}

func (h *Handler) AddResults(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	log.Printf("%+v\n", string(body))
	respondError(w, http.StatusNotImplemented, "add race result handler")

}
