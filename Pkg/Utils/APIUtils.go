package Utils

import (
	"GroupieTracker/Pkg/DataStruct"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func GetAllGroups() []DataStruct.Group {
	url := "https://groupietrackers.herokuapp.com/api/artists"

	spaceClient := http.Client{ // client to make a request to the api
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}
	req1, err := http.NewRequest(http.MethodGet, url, nil) // creation of the request
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := spaceClient.Do(req1)
	if getErr != nil {
		log.Fatal(getErr)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	data, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var Groups []DataStruct.Group
	jsonErr := json.Unmarshal(data, &Groups) // put all data into the data struct
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return Groups
}
