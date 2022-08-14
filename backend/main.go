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

type Station struct {
	FID        int64   `json:"fid"`
	ID         int64   `json:"id"`
	Nimi       string  `json:"nimi"`
	Namn       string  `json:"namn"`
	Name       string  `json:"name"`
	Osoite     string  `json:"osoite"`
	Adress     string  `json:"adress"`
	Kaupunki   string  `json:"kaupunki"`
	Stad       string  `json:"stad"`
	Operaattor string  `json:"operaattor"`
	Kapasiteet int64   `json:"kapasiteet"`
	X          float64 `json:"x"`
	Y          float64 `json:"y"`
}

type Journeys []Journey
type Stations []Station

func main() {
	fmt.Print("\n### Initialize database\n\n")
	database.InitDatabase()

	// Initialize new router
	router := mux.NewRouter().StrictSlash(true)

	// Register routes
	router.HandleFunc("/api/journeys", getJourneys).Methods("GET")
	router.HandleFunc("/api/stations", getStations).Methods("GET")

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
		var j Journey
		err = rows.Scan(&j.Id, &j.Departure, &j.Return, &j.DepartureStationId, &j.DepartureStationName, &j.ReturnStationId, &j.ReturnStationName, &j.Distance, &j.Duration)
		errorHandler(err)

		journeys = append(journeys, j)
	}

	// Convert to JSON
	j, err := json.Marshal(journeys)
	errorHandler(err)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "%s", string(j))
}

func getStations(writer http.ResponseWriter, req *http.Request) {
	// Fetch all stations from database
	rows, err := database.Database.Query("SELECT * FROM stations ORDER BY fid")
	errorHandler(err)

	var stations Stations

	for rows.Next() {
		var s Station
		err = rows.Scan(&s.FID, &s.ID, &s.Nimi, &s.Namn, &s.Name, &s.Osoite, &s.Adress, &s.Kaupunki, &s.Stad, &s.Operaattor, &s.Kapasiteet, &s.X, &s.Y)
		errorHandler(err)

		stations = append(stations, s)
	}

	// Convert to JSON
	s, err := json.Marshal(stations)
	errorHandler(err)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "%s", string(s))
}

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}
