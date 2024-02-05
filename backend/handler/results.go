package handler

import "net/http"

func (h *Handler) AddResults(w http.ResponseWriter, r *http.Request) {
	respondError(w, http.StatusNotImplemented, "add race result handler")
}
