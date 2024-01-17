package handler

import (
	"fmt"
	"log"
	"net/http"
)

func Admin(w http.ResponseWriter, r *http.Request) {
	log.Printf("admin handler\n")
	w.Write([]byte(fmt.Sprintf("admin handler")))
}
