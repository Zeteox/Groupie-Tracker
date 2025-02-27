package Handlers

import (
	"GroupieTracker/Pkg/Utils"
	"fmt"
	"html/template"
	"net/http"
)

func HandlerIndexPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	templ, err := template.ParseFiles("./Frontend/Templates/index.gohtml")
	if err != nil {
		fmt.Println("\033[31m[ERR]\033[0m [ROUTAGE] Error occured when handling index page request: ", err)
		w.WriteHeader(404)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		return
	}

	data, err := Utils.GetArtists()
	if err != nil {
		w.WriteHeader(503)
	}

	err = templ.Execute(w, data)
	if err != nil {
		w.WriteHeader(503)
		return
	}
}
