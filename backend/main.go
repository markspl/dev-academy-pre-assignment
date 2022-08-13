package main

import (
	"backend/database"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const PORT int = 3000

type Journey struct {
	Id                   int64  `json:"id"`
	Departure            string `json:"departure"`
	Return               string `json:"return"`
	DepartureStationId   string `json:"departureStationId"`
	DepartureStationName string `json:"departureStationName"`
	ReturnStationId      string `json:"returnStationId"`
	ReturnStationName    string `json:"returnStationName"`
	Distance             string `json:"distance"`
	Duration             string `json:"duration"`
}

type Journeys []Journey

func main() {
	fmt.Print("\n### Initialize database\n\n")
	database.InitDatabase()

	// Initialize new router
	router := mux.NewRouter().StrictSlash(true)

	// Register routes
	router.HandleFunc("/api/journeys", getJourneys).Methods("GET")

	// Start API
	fmt.Print("\n### Launching server\n\n")
	fmt.Printf("Starting REST API on port %v\n", PORT)
	_, err := fmt.Println(http.ListenAndServe(fmt.Sprintf(":%v", PORT), router))
	if err != nil {
		panic(err)
	}

	database.Database.Close()
}

func getJourneys(writer http.ResponseWriter, req *http.Request) {
	// Fetch data from database (limit 100)
	rows, err := database.Database.Query("SELECT * FROM journeys ORDER BY id DESC limit 100")
	errorHandler(err)

	var journeys Journeys

	// Add all received rows into journeys array
	for rows.Next() {
		var journey Journey
		err = rows.Scan(&journey.Id, &journey.Departure, &journey.Return, &journey.DepartureStationId, &journey.DepartureStationName, &journey.ReturnStationId, &journey.ReturnStationName, &journey.Distance, &journey.Duration)
		errorHandler(err)

		journeys = append(journeys, journey)
	}

	// Convert to JSON
	j, err := json.Marshal(journeys)
	errorHandler(err)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "%s", string(j))
}

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}
