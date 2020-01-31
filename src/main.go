package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter(folder string) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	staticFileDirectory := http.Dir(folder)
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	return r
}
func main() {
	r := newRouter("src/assets")
	http.ListenAndServe(":80", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
