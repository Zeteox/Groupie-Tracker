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

	//Creation of route for statics files
	fs := http.FileServer(http.Dir("./Frontend/Static"))
	http.Handle("/Static/", http.StripPrefix("/Static/", fs))

	//Creating a server and listening on the given port
	fmt.Println("\u001B[36m[INFO]\u001B[0m Listening on port " + strconv.Itoa(port) + " (http://localhost:" + strconv.Itoa(port) + ")")
	http.ListenAndServe("localhost:"+strconv.Itoa(port), nil)
}
