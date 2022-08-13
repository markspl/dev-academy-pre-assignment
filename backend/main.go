package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const PORT int = 3000

type Hello struct {
	Test string
}

func main() {
	fmt.Print("\n### Initialize database\n\n")
	//database.InitDatabase()

	// Initialize new router
	router := mux.NewRouter().StrictSlash(true)

	// Register routes
	router.HandleFunc("/api/journeys", getJourneys).Methods("GET")

	// Start API
	fmt.Print("\n### Launching server\n\n")
	fmt.Printf("Starting REST API on port %v", PORT)
	_, err := fmt.Println(http.ListenAndServe(fmt.Sprintf(":%v", PORT), router))
	if err != nil {
		panic(err)
	}
}

func getJourneys(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(Hello{"Ok!"})
}
