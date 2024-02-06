package handler

import "net/http"

func (h *Handler) AddSeason(w http.ResponseWriter, r *http.Request) {
	respondError(w, http.StatusNotImplemented, "add season handler")
}

func (h *Handler) AddRace(w http.ResponseWriter, r *http.Request) {
	respondError(w, http.StatusNotImplemented, "add race handler")
}

func (h *Handler) AddResults(w http.ResponseWriter, r *http.Request) {
	respondError(w, http.StatusNotImplemented, "add race result handler")
}
