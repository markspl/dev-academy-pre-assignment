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
	// Citybike database name
	dbName := "database.db"
	csvAddress := "dataset/2021-05.csv"

	// Remove existing database when launching the server
	os.Remove("./db/" + dbName)

	db, err := sql.Open("sqlite3", "./db/"+dbName)

	// If error else than zero value ("uninitialized" value)
	errorHandler(err)

	sqlStatement := `
	CREATE TABLE IF NOT EXISTS Journeys (
		Id					 INTEGER PRIMARY KEY,
		Departure            DATETIME     ,
		Return               DATETIME     ,
		DepartureStationId   INTEGER     ,
		DepartureStationName VARCHAR(100)     ,
		ReturnStationId      INTEGER     ,
		ReturnStationName    VARCHAR(100)     ,
		Distance             INTEGER     ,
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

			// Change time values to SQL format (yyyy-mm-dd hh:mm:ss)
			r[0] = strings.Replace(r[0], "T", " ", 1)
			r[1] = strings.Replace(r[1], "T", " ", 1)

			value := "('" + strconv.Itoa(idValue) + "','" + strings.Join(r, "','") + "')"

			// Include to same array
			stmtEnd = append(stmtEnd, value)

			// Keep Id unique
			idValue += 1
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
