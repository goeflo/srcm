package handler

import (
	"log"
	"net/http"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {

	log.Printf("home router\n")

	if err := h.Tmpl.ExecuteTemplate(w, "home/index",
		map[string]interface{}{"title": "homepage",
			"subtitle": "some nice subtitle"}); err != nil {
		log.Fatal(err)
	}

}
