package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/render"
)

func (h *Handler) Assets(w http.ResponseWriter, r *http.Request) {
	log.Printf("-> handle assets %s %s %s", r.RemoteAddr, r.Method, r.URL)
	w.Write([]byte("assets"))
}

func (h *Handler) Css(w http.ResponseWriter, r *http.Request) {
	log.Printf("-> handle css %s %s %s", r.RemoteAddr, r.Method, r.URL)

	// add template base directory to request filepath and load file
	assetFilename := filepath.Join("./templates", r.URL.String())
	b, err := os.ReadFile(assetFilename)
	if err != nil {
		render.Render(w, r, ErrStatusInternalServerError(fmt.Errorf("unable to read file: '%v'", assetFilename)))
	}

	// set the content type to text/css because browser will not load
	// css file if content type is default plain/text
	w.Header().Add("Content-Type", "text/css")
	w.Write(b)
}
