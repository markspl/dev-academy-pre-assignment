package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Citybike database name
	dbName := "database.db"

	// Remove existing database when launching the server
	os.Remove("./db/" + dbName)

	db, err := sql.Open("sqlite3", "./db/"+dbName)

	// If error else than zero value ("uninitialized" value)
	errorHandler(err)

	sqlStatement := `
	CREATE TABLE IF NOT EXISTS Journeys (
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

	// Insert data
	stmt, err := db.Prepare(`
	INSERT INTO Journeys(Departure, Return, DepartureStationId, DepartureStationName, ReturnStationId, ReturnStationName, Distance, Duration) values(?,?,?,?,?,?,?,?)
	`)
	errorHandler(err)

	// Example data #1
	_, err = stmt.Exec("2021-05-31T23:57:25", "2021-06-01T00:05:46", "094", "Laajalahden aukio", "100", "Teljäntie", "2043", "500")
	errorHandler(err)

	// Example data #2
	_, err = stmt.Exec("2021-05-31T23:56:59", "2021-06-01T00:07:14", "082", "Töölöntulli", "113", "Pasilan asema", "1870", "611")
	errorHandler(err)

	// Query
	rows, err := db.Query("SELECT * FROM Journeys")
	errorHandler(err)

	var departureTime time.Time
	var returnTime time.Time
	var departureStationId int
	var departureStationName string
	var returnStationId int
	var returnStationName string
	var distance int
	var duration int

	for rows.Next() {
		err = rows.Scan(&departureTime, &returnTime, &departureStationId, &departureStationName, &returnStationId, &returnStationName, &distance, &duration)
		errorHandler(err)

		fmt.Println(departureTime)
		fmt.Println(returnTime)
		fmt.Println(departureStationId)
		fmt.Println(departureStationName)
		fmt.Println(returnStationId)
		fmt.Println(returnStationName)
		fmt.Println(distance)
		fmt.Println(duration)
		fmt.Println("===")
	}

	rows.Close()

	db.Close()
}

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}
