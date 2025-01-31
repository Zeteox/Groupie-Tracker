package main

import (
	internal "GroupieTracker/Internal/Server"
	"GroupieTracker/Pkg/Utils"
	"fmt"
)

func main() {
	artist, _ := Utils.GetArtistsById(52)
	fmt.Println(artist.AllLocations.LocationsMap)
	internal.CreateAndListenServer(8080)
}
