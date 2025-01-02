package Handlers

import (
	"GroupieTracker/Pkg/DataStruct"
	"GroupieTracker/Pkg/Utils"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
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

	if r.Method == http.MethodPost {

		numberOfMember, err := strconv.Atoi(r.FormValue("MembersNumber"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		if numberOfMember >= 1 {
			var tmpData []DataStruct.Artist
			for x := 0; x < len(data); x++ {
				if len(data[x].Members) == numberOfMember {
					tmpData = append(tmpData, data[x])
				}
			}
			data = tmpData
		}
	}

	templ.Execute(w, data)
}
