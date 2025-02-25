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
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		return
	}

	urlStr := "http://" + r.Host + r.URL.Path + "?" + r.URL.RawQuery
	myUrl, _ := url.Parse(urlStr)
	param, _ := url.ParseQuery(myUrl.RawQuery)
	IdStr := param.Get("Id")
	Id, _ := strconv.Atoi(IdStr)

	if Id < 1 {
		Id = 1
	} else if Id > 52 {
		Id = 52
	}

	data, err := Utils.GetArtistsById(Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	templ.Execute(w, data)
}
