package Handlers

import (
	"GroupieTracker/Pkg/Utils"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
)

func HandlerInfosPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	templ, err := template.New("infos.gohtml").Funcs(template.FuncMap{
		"CityWithCountry": func(cityWithCountry string) []string {
			return Utils.CityWithCountry(cityWithCountry)
		},
		"GetCoordinates": func(City string) []float32 {
			return Utils.GetCoordinates(City)
		},
	}).ParseFiles("./Frontend/Templates/infos.gohtml")
	if err != nil {
		fmt.Println("\033[31m[ERR]\033[0m [ROUTAGE] Error occured when handling index page request: ", err)
		w.WriteHeader(404)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		return
	}

	urlStr := "http://" + r.Host + r.URL.Path + "?" + r.URL.RawQuery
	myUrl, err := url.Parse(urlStr)
	if err != nil {
		w.WriteHeader(500)
	}
	param, err := url.ParseQuery(myUrl.RawQuery)
	if err != nil {
		w.WriteHeader(500)
	}
	IdStr := param.Get("Id")
	Id, err := strconv.Atoi(IdStr)
	if err != nil {
		w.WriteHeader(404)
	}

	if Id < 1 {
		Id = 1
	} else if Id > 52 {
		Id = 52
	}

	data, err := Utils.GetArtistsById(Id)
	if err != nil {
		w.WriteHeader(503)
	}

	err = templ.Execute(w, data)
	if err != nil {
		w.WriteHeader(503)
		return
	}
}
