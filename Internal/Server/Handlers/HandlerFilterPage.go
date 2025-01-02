package Handlers

import (
	"GroupieTracker/Pkg/Utils"
	"fmt"
	"html/template"
	"net/http"
)

func HandlerFilterPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	templ, err := template.ParseFiles("./Frontend/Templates/FilterModel.gohtml")
	if err != nil {
		fmt.Println("\033[31m[ERR]\033[0m [ROUTAGE] Error occured when handling index page request: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		return
	}

	allArtists, err := Utils.GetArtists()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	data := allArtists

	templ.Execute(w, data)
}
