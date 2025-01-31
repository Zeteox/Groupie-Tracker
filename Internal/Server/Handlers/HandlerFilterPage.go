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
	templ, err := template.New("FilterModel.gohtml").Funcs(template.FuncMap{
		"Iterate": func(count int) []int {
			var Items []int
			for i := 0; i < count; i++ {
				Items = append(Items, i)
			}
			return Items
		},
		"GetCareerOrAlbumYears": func(index int) string {
			careerAndAlbumDates := []string{"1970-1979", "1980-1989", "1990-1999", "2000-2009", "2010-2019", "2020-2025"}
			return careerAndAlbumDates[index]
		},
		"GetAllGroupsName": func() string {
			var Names string
			artists, _ := Utils.GetArtists()
			for x := 0; x < len(artists); x++ {
				Names += artists[x].Name + ","
			}
			return Names[:len(Names)-1]
		},
	}).ParseFiles("./Frontend/Templates/FilterModel.gohtml")
	if err != nil {
		fmt.Println("\033[31m[ERR]\033[0m [ROUTAGE] Error occured when handling index page request: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		return
	}

	data := DataStruct.ArtistsAndForm{}

	allArtists, err := Utils.GetArtists()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	data.Artists = allArtists
	Utils.SetupForm(&data.FormData)

	if r.Method == http.MethodPost {

		numberOfMember, err := strconv.Atoi(r.FormValue("MembersNumber"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		data.FormData.MemberNumber = r.FormValue("MembersNumber")
		allCareerAndAlbumDates := []string{"1970-1979", "1980-1989", "1990-1999", "2000-2009", "2010-2019", "2020-2025"}

		for x := 0; x < len(allCareerAndAlbumDates); x++ {
			data.FormData.CareerYears[x] = r.FormValue("career:" + allCareerAndAlbumDates[x])
			if data.FormData.CareerYears[x] == "1" {
				var tmpData []DataStruct.Artist
				minYear, _ := strconv.Atoi(allCareerAndAlbumDates[x][:4])
				maxYear, _ := strconv.Atoi(allCareerAndAlbumDates[x][5:])
				for y := 0; y < len(data.Artists); y++ {
					if data.Artists[y].CreationDate >= minYear && data.Artists[y].CreationDate <= maxYear {
						tmpData = append(tmpData, data.Artists[y])
					}
				}
				data.Artists = tmpData
			}

			data.FormData.AlbumYears[x] = r.FormValue("album:" + allCareerAndAlbumDates[x])
			if data.FormData.AlbumYears[x] == "1" {
				var tmpData []DataStruct.Artist
				minYear, _ := strconv.Atoi(allCareerAndAlbumDates[x][:4])
				maxYear, _ := strconv.Atoi(allCareerAndAlbumDates[x][5:])
				for y := 0; y < len(data.Artists); y++ {
					albumYear, _ := strconv.Atoi(data.Artists[y].FirstAlbum[6:])
					if albumYear >= minYear && albumYear <= maxYear {
						tmpData = append(tmpData, data.Artists[y])
					}
				}
				data.Artists = tmpData
			}
		}

		if numberOfMember >= 1 {
			var tmpData []DataStruct.Artist
			for x := 0; x < len(data.Artists); x++ {
				if len(data.Artists[x].Members) == numberOfMember {
					tmpData = append(tmpData, data.Artists[x])
				}
			}
			data.Artists = tmpData
		}

		fmt.Println("nb m:" + data.FormData.MemberNumber)
		data.FormData.SearchbarContent = r.FormValue("Searchbar")
		fmt.Println("search:" + data.FormData.SearchbarContent)

	}

	err = templ.Execute(w, data)
	if err != nil {
		return
	}
}
