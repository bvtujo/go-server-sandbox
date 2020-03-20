package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// User struct represents a user and their crucial data
type User struct {
	Rank   int    `json:"rank"`
	Points string `json:"points"`
	Name   string `json:"user"`
}

// URLCommand encodes the granting of points from the source user to the dest user
type URLCommand struct {
	SourceUser string  `json:"src"`
	DestUser   string  `json:"dest"`
	Points     float64 `json:"pts"`
}

var users []User

func getUserHandler(w http.ResponseWriter, r *http.Request) {

	userListBytes, err := json.Marshal(users)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(userListBytes)
}

func addPointsHandler(w http.ResponseWriter, r *http.Request) {

}

func sanitizePath(path string) string {
	re := regexp.MustCompile("[^./?=]+")
	strs := re.FindAll([]byte(path), -1)
	return string(strs[0])
}
func parseURL(rawurl string) URLCommand {

	u, err := url.Parse(rawurl)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	var cmd URLCommand
	hostParts := strings.Split(u.Hostname(), ".")
	cmd.SourceUser = hostParts[0]
	cmd.DestUser = sanitizePath(u.Path)
	cmd.Points, err = strconv.ParseFloat(hostParts[2], 64)

	return cmd
}

func checkUser(user string) bool {
	return false
}
