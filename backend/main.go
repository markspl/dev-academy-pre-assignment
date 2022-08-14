package main

import (
	"backend/database"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

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
	// Load configurations
	LoadApiConfig()

	// Count how long running this takes
	timeTrack := time.Now()

	fmt.Print("\n### Initialize database\n\n")
	database.InitDatabase(ApiConfig.DB_ADDRESS)

	fmt.Printf("Loading...\n\n")

	// Import all journeys from ./journeys/*.csv file
	database.ImportJourneys(ApiConfig.JOURNEYS_FOLDER, ApiConfig.STMT_COUNT_QUERY, ApiConfig.MIN_JOURNEY_DIST, ApiConfig.MIN_JOURNEY_TIME)

	// Import stations from ./stations/Helsingin_ja_Espoon_(...).csv file
	database.ImportStations(ApiConfig.STATIONS_FILE)

	fmt.Printf("\n\nLoaded.")

	// Track spent time
	fmt.Printf("\nTime: %v\n", time.Since(timeTrack))

	// Get information about data in database
	reportTableLengths()

	// Initialize new router
	router := mux.NewRouter().StrictSlash(true)

	// Register routes
	router.HandleFunc("/api/journeys", getJourneys).Methods("GET")
	router.HandleFunc("/api/stations", getStations).Methods("GET")

	// Start API
	fmt.Print("\n### Launching server\n\n")
	fmt.Printf("Starting REST API on port %v\n", ApiConfig.API_PORT)
	_, err := fmt.Println(http.ListenAndServe(fmt.Sprintf(":%v", ApiConfig.API_PORT), router))
	if err != nil {
		panic(err)
	}

	database.Database.Close()
}

func reportTableLengths() {
	// Count imported lines
	var countJourneys int
	var countStations int

	err := database.Database.QueryRow("SELECT COUNT(*) FROM Journeys").Scan(&countJourneys)
	errorHandler(err)
	err = database.Database.QueryRow("SELECT COUNT(*) FROM Stations").Scan(&countStations)
	errorHandler(err)

	fmt.Println("\nTotal of imported journeys:", countJourneys)
	fmt.Println("Total of imported stations:", countStations)
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
	writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
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
	writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "%s", string(s))
}

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}
