package handler

import (
	"html/template"
	"log"
	"net/http"
)

var ext = ".html"

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {

	log.Printf("home handler\n")

	file, err := template.ParseFiles("templates/index.html")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	tmpl := template.Must(file, err)
	tmpl.Execute(w, nil)

	//h.Tmpl.ExecuteTemplate(w, "public/index.html")

	// log.Printf("-> handle home %s %s %s", r.RemoteAddr, r.Method, r.URL)

	// // only request for html pages are allowed, skip everything else
	// //if !strings.HasSuffix(r.URL.String(), ".html") {
	// //	return
	// //}

	// // request url must start with '/home/' and end with a html filename
	// if !strings.HasPrefix(r.URL.String(), "/home/") || !strings.HasSuffix(r.URL.String(), ".html") {
	// 	render.Render(w, r, ErrInvalidRequest(fmt.Errorf("not a html file '%v'", r.URL.String())))
	// 	return
	// }

	// templateName := "home/index"
	// templateName = strings.TrimSuffix(r.URL.String(), ext)

	// log.Printf("template name: %v\n", templateName)
	// if strings.HasPrefix(templateName, "/") {
	// 	templateName = templateName[1:]
	// }

	// if err := h.Tmpl.ExecuteTemplate(w, templateName,
	// 	map[string]interface{}{"title": "homepage",
	// 		"subtitle": "some nice subtitle"}); err != nil {
	// 	log.Printf("template not found: %v\n", err)
	// 	render.Render(w, r, ErrStatusInternalServerError(fmt.Errorf("template not found '%v'", err)))
	// 	return
	// }

}
