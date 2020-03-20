//api.go

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"internal/pkg/points"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	errPointsInvalid = "points invalid: %v"
)

func validatePoints(pts string) (float64, error) {
	num, err := strconv.ParseFloat(pts, 64)
	if err != nil {
		return 0, fmt.Errorf(errPointsInvalid, err.Error)
	}
	return num, nil
}

func PutPoints(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	psUser := ps.ByName("user")
	psPoints := ps.ByName("points")

	pts, err := validatePoints(psPoints)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	svc := points.NewPointsService(session.Must(
		session.NewSessionWithOptions(
			session.Options{
				SharedConfigState: session.SharedConfigEnable,
			},
		),
	),
	)
	pointsInput := points.PointsInput{
		Username:  psUser,
		From:      "",
		Points:    pts,
		Timestamp: time.Now(),
	}

	err = svc.Put(&pointsInput)
	w.Write([]byte(fmt.Sprintf("%v", pointsInput)))
	return
}

func GetUser(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	serialized, err := json.Marshal(
		points.PointsUser{
			Uuid:   "fab5",
			Points: 5,
			Transactions: []points.PointsInput{
				points.PointsInput{
					Points:    5,
					Username:  "austiely",
					From:      "system",
					Timestamp: time.Now(),
				},
			},
		},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Write(serialized)
	return
}

func main() {
	router := httprouter.New()
	router.POST("/api/:points/points/to/:user", PutPoints)
	router.GET("/api/user/:user", GetUser)

	logrus.Fatal(http.ListenAndServe(":8080", router))
}
