package Utils

import (
	"GroupieTracker/Pkg/DataStruct"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetArtists() ([]DataStruct.Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var Artists []DataStruct.Artist
	if err := json.NewDecoder(resp.Body).Decode(&Artists); err != nil {
		return nil, err
	}

	for x := 0; x < len(Artists); x++ {
		Artists[x].AllLocations, _ = GetLocationsByID(Artists[x].Id)
		Artists[x].AllDates, _ = GetDatesByID(Artists[x].Id)
		Artists[x].AllRelations, _ = GetRelationsByID(Artists[x].Id)
	}

	return Artists, nil
}

func GetArtistsById(id int) (DataStruct.Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + strconv.Itoa(id))
	if err != nil {
		return DataStruct.Artist{}, err
	}
	defer resp.Body.Close()

	var Artist DataStruct.Artist
	if err := json.NewDecoder(resp.Body).Decode(&Artist); err != nil {
		return DataStruct.Artist{}, err
	}

	Artist.AllLocations, _ = GetLocationsByID(Artist.Id)
	Artist.AllDates, _ = GetDatesByID(Artist.Id)
	Artist.AllRelations, _ = GetRelationsByID(Artist.Id)

	return Artist, nil
}

func GetLocationsByID(id int) (DataStruct.Locations, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + strconv.Itoa(id))
	if err != nil {
		return DataStruct.Locations{}, err
	}
	defer resp.Body.Close()

	var Locations DataStruct.Locations
	if err := json.NewDecoder(resp.Body).Decode(&Locations); err != nil {
		return DataStruct.Locations{}, err
	}
	return Locations, nil
}

func GetRelationsByID(id int) (DataStruct.Relations, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + strconv.Itoa(id))
	if err != nil {
		return DataStruct.Relations{}, err
	}
	defer resp.Body.Close()

	var Relations DataStruct.Relations
	if err := json.NewDecoder(resp.Body).Decode(&Relations); err != nil {
		return DataStruct.Relations{}, err
	}
	return Relations, nil
}

func GetDatesByID(id int) (DataStruct.Dates, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates/" + strconv.Itoa(id))
	if err != nil {
		return DataStruct.Dates{}, err
	}
	defer resp.Body.Close()

	var Dates DataStruct.Dates
	if err := json.NewDecoder(resp.Body).Decode(&Dates); err != nil {
		return DataStruct.Dates{}, err
	}
	return Dates, nil
}
