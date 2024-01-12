package handler

import (
	"log"
	"net/http"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	log.Printf("home router\n")
	w.Write([]byte("home"))
}
