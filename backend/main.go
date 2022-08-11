package main

import (
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbName := "database.db"             // Citybike database name
	csvAddress := "dataset/2021-05.csv" // CSV file address (import)
	minJourneyDist := 10.0              // Don't import journeys if m < 10m
	minJourneyTime := 10                // Don't import journeys if t < 10s

	// Remove existing database when launching the server
	os.Remove("./db/" + dbName)

	db, err := sql.Open("sqlite3", "./db/"+dbName)

	// If error else than zero value ("uninitialized" value)
	errorHandler(err)

	sqlStatement := `
			CREATE TABLE IF NOT EXISTS Journeys (
				Id					 INTEGER PRIMARY KEY,
				Departure            TEXT,
				Return               TEXT,
				DepartureStationId   INTEGER,
				DepartureStationName TEXT,
				ReturnStationId      INTEGER,
				ReturnStationName    TEXT,
				Distance             INTEGER,
				Duration             INTEGER
			 );
			 DELETE FROM Journeys;
			`

	// Create database table
	_, err = db.Exec(sqlStatement)
	errorHandler(err)

	// Open CSV example file
	file, err := os.Open(csvAddress)
	errorHandler(err)

	read := csv.NewReader(file)

	// Read only the first row
	_, err = read.Read()
	errorHandler(err)

	// Import data from file
	fmt.Println("Loading...")
	idValue := 0

	for {
		i := 0
		stmtEnd := []string{}
		var errInner error

		// Create a bulk INSERT with 100 VALUES
		for i < 100 {
			r, errInner := read.Read()

			// No more input available
			if errors.Is(errInner, io.EOF) {
				break
			}

			// Check if journey lasted for less than 10s and distance over 10m
			dist, err := strconv.ParseFloat(r[6], 64)
			errorHandler(err)
			longerTime := (dist > minJourneyDist)
			timeA, err := strconv.Atoi(r[7])
			errorHandler(err)
			longerDist := (timeA > minJourneyTime)

			if longerTime && longerDist {
				// Change time values to SQL format (yyyy-mm-dd hh:mm:ss)
				r[0] = strings.Replace(r[0], "T", " ", 1)
				r[1] = strings.Replace(r[1], "T", " ", 1)

				value := "('" + strconv.Itoa(idValue) + "','" + strings.Join(r, "','") + "')"

				// Include to same array
				stmtEnd = append(stmtEnd, value)

				// Keep Id unique
				idValue += 1
			}
		}

		stmtBegin := "INSERT INTO Journeys(Id, Departure, Return, DepartureStationId, DepartureStationName, ReturnStationId, ReturnStationName, Distance, Duration) VALUES"

		completed := stmtBegin + " " + strings.Join(stmtEnd, ",") + ";"

		// Insert data
		_, err := db.Exec(completed)
		errorHandler(err)

		// If no input happened, let's quit outer for too.
		if errors.Is(errInner, io.EOF) {
			break
		}
		break
	}
	fmt.Println("Loaded.")

	// Count exported lines
	var count int

	err = db.QueryRow("SELECT COUNT(*) FROM Journeys").Scan(&count)
	errorHandler(err)

	fmt.Println("\nTotal of exported journeys:", count)

	db.Close()
}

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}
