package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetHeroes(w http.ResponseWriter, r *http.Request) {

	heroes := `[{"id": "11", "name": "Mr. Nice"},
	{"id": "12", "name": "Narco"},
	{"id": "13", "name": "Bombasto"},
	{"id": "14", "name": "Celeritas"},
	{"id": "15", "name": "Magneta"},
	{"id": "16", "name": "RubberMan"},
	{"id": "17", "name": "Dynama"},
	{"id": "18", "name": "Dr IQ"},
	{"id": "19", "name": "Magma"},
	{"id": "20", "name": "Tornado"}]`

	 //Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", heroes)

}

func BaseHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", BaseHandler)
	r.HandleFunc("/heroes", GetHeroes)

	// Bind to a port and pass our router in
	http.ListenAndServe(":8000", r)
}