package Handlers

import (
	"GroupieTracker/Pkg/Utils"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func HandlerInfosPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	templ, err := template.New("infos.gohtml").Funcs(template.FuncMap{
		"CityWithCountry": func(cityWithCountry string) []string {
			split := strings.Split(cityWithCountry, "-")
			citySplit := strings.Split(split[0], "_")
			countrySplit := strings.Split(split[1], "_")

			city := strings.ToUpper(citySplit[0][:1]) + citySplit[0][1:]
			if len(citySplit) > 1 {
				for i := 1; i <= len(citySplit)-1; i++ {
					city += " " + strings.ToUpper(citySplit[i][:1]) + citySplit[i][1:]

				}
			}

			country := strings.ToUpper(countrySplit[0][:1]) + countrySplit[0][1:]
			if len(countrySplit) > 1 {
				for i := 1; i <= len(countrySplit)-1; i++ {
					country += " " + strings.ToUpper(countrySplit[i][:1]) + countrySplit[i][1:]

				}
			}

			return []string{city, country}
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
