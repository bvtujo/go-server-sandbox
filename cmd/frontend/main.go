package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(folder string) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", Handler).Methods("GET")

	staticFileDirectory := http.Dir(folder)
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	return r
}
func main() {
	r := NewRouter("src/assets")
	http.ListenAndServe(":80", r)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
