package internal

import (
	"GroupieTracker/Internal/Server/Handlers"
	"fmt"
	"net/http"
	"strconv"
)

func CreateAndListenServer(port int) {
	//Creation of routes for all pages
	http.HandleFunc("/", Handlers.HandlerIndexPage)
	http.HandleFunc("/Filter", Handlers.HandlerFilterPage)

	//Creation of route for statics file
	fs := http.FileServer(http.Dir("./Frontend/Static"))
	http.Handle("/Static/", http.StripPrefix("/Static/", fs))

	//Creating a server and listening on the given port
	fmt.Println("\u001B[36m[INFO]\u001B[0m Listening on port " + strconv.Itoa(port) + " (http://localhost:" + strconv.Itoa(port) + ")")
	http.ListenAndServe("0.0.0.0:"+strconv.Itoa(port), nil)
}
