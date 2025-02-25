package Handlers

import (
	"GroupieTracker/Pkg/DataStruct"
	"GroupieTracker/Pkg/Utils"
	"fmt"
	"html/template"
	"net/http"
	"slices"
	"strconv"
	"strings"
)

func HandlerDiscoverPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	templ, err := template.New("discoverpage.gohtml").Funcs(template.FuncMap{
		"CityWithCountry": func(cityWithCountry string) []string {
			return Utils.CityWithCountry(cityWithCountry)
		},
		"Iterate": func(count int) []int {
			var Items []int
			for i := 0; i < count; i++ {
				Items = append(Items, i)
			}
			return Items
		},
		"GetCareerOrAlbumYears": func(index int) string {
			careerAndAlbumDates := []string{"1950-1959", "1960-1969", "1970-1979", "1980-1989", "1990-1999", "2000-2009", "2010-2019", "2020-2025"}
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
	}).ParseFiles("./Frontend/Templates/discoverpage.gohtml")
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
	for i := 0; i < len(allArtists); i++ {
		allArtists[i].AllLocations, err = Utils.GetLocationsByID(i + 1)
		if err != nil {
			return
		}
	}

	data.FilteredArtists, data.AllArtists = allArtists, allArtists
	Utils.SetupForm(&data.FormData)

	if r.Method == http.MethodPost {

		allCareerAndAlbumDates := []string{"1950-1959", "1960-1969", "1970-1979", "1980-1989", "1990-1999", "2000-2009", "2010-2019", "2020-2025"}

		switch r.FormValue("Searchbar") {

		case "":

			for x := 0; x < len(allCareerAndAlbumDates); x++ {
				data.FormData.CareerYears[x] = r.FormValue("career:" + allCareerAndAlbumDates[x])
				data.FormData.AlbumYears[x] = r.FormValue("album:" + allCareerAndAlbumDates[x])
			}

			if slices.Contains(data.FormData.CareerYears, "1") {
				data.FilteredArtists = []DataStruct.Artist{}
				for x := 0; x < len(allCareerAndAlbumDates); x++ {
					if data.FormData.CareerYears[x] == "1" {
						FilterWithCareerYear(&data, allCareerAndAlbumDates[x])
					}
				}
			}

			if slices.Contains(data.FormData.AlbumYears, "1") {
				tmpFilteredArtists := []DataStruct.Artist{}
				for x := 0; x < len(allCareerAndAlbumDates); x++ {
					if data.FormData.AlbumYears[x] == "1" {
						FilterWithAlbumYear(&data, &tmpFilteredArtists, allCareerAndAlbumDates[x])
					}
				}
				data.FilteredArtists = tmpFilteredArtists
			}

			//Filter nb membres
			FilterWithMemberNumber(r, w, &data)

			FilterWithCountry(r, w, &data)

		default:
			Utils.SetupForm(&data.FormData)
			FilterWithSearchBar(r, &data)
		}
	}

	err = templ.Execute(w, data)
	if err != nil {
		return
	}
}

func FilterWithSearchBar(r *http.Request, data *DataStruct.ArtistsAndForm) {
	data.FormData.SearchbarContent = r.FormValue("Searchbar")
	var tmpData []DataStruct.Artist
	for x := 0; x < len(data.AllArtists); x++ {
		if strings.Contains(strings.ToLower(data.AllArtists[x].Name), strings.ToLower(r.FormValue("Searchbar"))) {
			tmpData = append(tmpData, data.AllArtists[x])
		}
	}
	data.FilteredArtists = tmpData
}

func FilterWithMemberNumber(r *http.Request, w http.ResponseWriter, data *DataStruct.ArtistsAndForm) {
	numberOfMember, err := strconv.Atoi(r.FormValue("MembersNumber"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	data.FormData.MemberNumber = r.FormValue("MembersNumber")
	if numberOfMember >= 1 {
		var tmpData []DataStruct.Artist
		for x := 0; x < len(data.FilteredArtists); x++ {
			if len(data.FilteredArtists[x].Members) == numberOfMember {
				tmpData = append(tmpData, data.FilteredArtists[x])
			}
		}
		data.FilteredArtists = tmpData
	}
}

func FilterWithCareerYear(data *DataStruct.ArtistsAndForm, year string) {
	var tmpData []DataStruct.Artist
	minYear, _ := strconv.Atoi(year[:4])
	maxYear, _ := strconv.Atoi(year[5:])
	for y := 0; y < len(data.AllArtists); y++ {
		if data.AllArtists[y].CreationDate >= minYear && data.AllArtists[y].CreationDate <= maxYear {
			tmpData = append(tmpData, data.AllArtists[y])
		}
	}
	for _, elem := range tmpData {
		data.FilteredArtists = append(data.FilteredArtists, elem)
	}
}

func FilterWithAlbumYear(data *DataStruct.ArtistsAndForm, tmpFilteredArtists *[]DataStruct.Artist, year string) {
	var tmpData []DataStruct.Artist
	minYear, _ := strconv.Atoi(year[:4])
	maxYear, _ := strconv.Atoi(year[5:])
	for y := 0; y < len(data.FilteredArtists); y++ {
		albumYear, _ := strconv.Atoi(data.FilteredArtists[y].FirstAlbum[6:])
		if albumYear >= minYear && albumYear <= maxYear {
			tmpData = append(tmpData, data.FilteredArtists[y])
		}
	}
	for _, elem := range tmpData {
		*tmpFilteredArtists = append(*tmpFilteredArtists, elem)
	}
}

func FilterWithCountry(r *http.Request, w http.ResponseWriter, data *DataStruct.ArtistsAndForm) {
	data.FormData.Country = r.FormValue("country")
	var tmpData []DataStruct.Artist
	if r.FormValue("country") != "" {
		for x := 0; x < len(data.FilteredArtists); x++ {
			if Utils.Contain(data.FilteredArtists[x].AllLocations.GroupLocations, r.FormValue("country")) {
				tmpData = append(tmpData, data.FilteredArtists[x])
			}
		}
		data.FilteredArtists = tmpData
	}
}
