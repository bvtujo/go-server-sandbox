package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/bvtujo/go-server-sandbox/internal/pkg/points"
	template "github.com/bvtujo/go-server-sandbox/internal/pkg/templates"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", Handler).Methods("GET")
	r.HandleFunc("/", HandleRoot).Methods("GET")
	r.HandleFunc("/healthcheck", HealthCheck).Methods("GET")
	return r
}
func main() {
	r := NewRouter()
	http.ListenAndServe(":80", r)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	tmpl := template.GetIndex()
	data := []points.User{
		points.User{
			Username: "austiely",
			Points:   5,
		},
	}
	tmpl.Execute(w, data)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
