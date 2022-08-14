package database

import (
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Citybike database name
const DBADDRESS string = "database/db/database.db"

// Journeys folder (import)
const CSVADDRESS string = "database/dataset/journeys"

// List of stations
const STATIONADDRESS string = "database/dataset/stations/Helsingin_ja_Espoon_kaupunkipyB6rA4asemat_avoin.csv"

const MINJOURNEYDIST float64 = 10.0 // Don't import journeys if m < 10m
const MINJOURNEYTIME int = 10       // Don't import journeys if t < 10s
const STMTCOUNT int = 100000        // How many values are imported to database (max)

var Database *sql.DB

func InitDatabase() {
	// Count how long running this takes
	timeTrack := time.Now()

	// Remove existing database when launching the server
	err := os.Remove(DBADDRESS)
	if err != nil {
		fmt.Println("Can't delete the db file, creating a new.")
	}

	// If error else than zero value ("uninitialized" value)
	db, err := sql.Open("sqlite3", DBADDRESS)
	if err != nil {
		fmt.Println("File not found")
	}

	Database = db

	// Create database tables
	createDatabaseTables()

	fmt.Printf("Loading...\n\n")

	// Import all journeys from ./journeys/*.csv file
	importJourneys()

	// Import stations from ./stations/Helsingin_ja_Espoon_(...).csv file
	importStations()

	fmt.Printf("\n\nLoaded.")

	// Count imported lines
	var countJourneys int
	var countStations int

	err = Database.QueryRow("SELECT COUNT(*) FROM Journeys").Scan(&countJourneys)
	errorHandler(err, "")
	err = Database.QueryRow("SELECT COUNT(*) FROM Stations").Scan(&countStations)
	errorHandler(err, "")

	fmt.Println("\nTotal of imported journeys:", countJourneys)
	fmt.Println("Total of imported stations:", countStations)

	//Database.Close()

	fmt.Printf("\nTime: %v\n", time.Since(timeTrack))
}

func createDatabaseTables() {
	// Journeys table
	sqlStatementJourneys := `
		CREATE TABLE IF NOT EXISTS Journeys (
			Id					 INTEGER PRIMARY KEY,
			Departure            TEXT NOT NULL,
			Return               TEXT NOT NULL,
			DepartureStationId   TEXT NOT NULL,
			DepartureStationName TEXT NOT NULL,
			ReturnStationId      TEXT NOT NULL,
			ReturnStationName    TEXT NOT NULL,
			Distance             TEXT NOT NULL,
			Duration             TEXT NOT NULL
		);
		DELETE FROM Journeys;
	`

	// Stations table
	sqlStatementStations := `
		CREATE TABLE IF NOT EXISTS Stations (
			Fid			INTEGER PRIMARY KEY,
			Id			TEXT NOT NULL,
			Nimi		TEXT NOT NULL,
			Namn		TEXT NOT NULL,
			Name		TEXT NOT NULL,
			Osoite		TEXT NOT NULL,
			Adress		TEXT NOT NULL,
			Kaupunki	TEXT,
			Stad		TEXT,
			Operaattor	TEXT,
			Kapasiteet	TEXT NOT NULL,
			x			REAL NOT NULL,
			y			REAL NOT NULL
		);
		DELETE FROM Stations;
	`

	// Create database table
	_, errJ := Database.Exec(sqlStatementJourneys)
	errorHandler(errJ, "")

	_, errS := Database.Exec(sqlStatementStations)
	errorHandler(errS, "")
}

func importJourneys() {
	// Open all CSV example files
	folder, err := os.Open(CSVADDRESS)
	errorHandler(err, "")

	defer folder.Close()
	files, err := folder.Readdirnames(0)
	errorHandler(err, "")

	idValue := 0

	// Go thru all files
	for _, name := range files {
		// Filter only CSV files
		if !strings.HasSuffix(name, ".csv") {
			fmt.Printf("File %v is not CSV file, skip.\n", name)
			break
		}

		filePath := fmt.Sprintf("%v/%v", CSVADDRESS, name)
		file, err := os.Open(filePath)
		errorHandler(err, "")

		defer file.Close()

		fmt.Printf("File %v loaded. Importing...\n", name)

		read := csv.NewReader(file)

		// Read only the first row
		headers, err := read.Read()
		errorHandler(err, "")

		// Check there is 8 headers
		if len(headers) == 8 {
			for {
				i := 0
				stmtEnd := []string{}
				var errInner error

				// Create a bulk INSERT with STMTCOUNT VALUES
				for i < STMTCOUNT {
					r, err := read.Read()

					errInner = err

					// No more input available
					if errors.Is(err, io.EOF) {
						break
					}

					// Validate data before importing
					// If the row includes incorrect values (detected by regex), skip row
					if validateJourneyDataBeforeImport(r, idValue) {
						// Check if journey lasted for less than 10s and distance over 10m
						dist, err := strconv.ParseFloat(r[6], 64)
						errorHandler(err, "")
						longerTime := (dist > MINJOURNEYDIST)
						timeA, err := strconv.Atoi(r[7])
						errorHandler(err, "")
						longerDist := (timeA > MINJOURNEYTIME)

						if longerTime && longerDist {
							value := "('" + strconv.Itoa(idValue) + "','" + strings.Join(r, "','") + "')"

							// Include to same array
							stmtEnd = append(stmtEnd, value)

							// Keep Id unique
							idValue += 1
							i += 1
						}
					}
				}

				stmtBegin := "INSERT INTO Journeys(Id, Departure, Return, DepartureStationId, DepartureStationName, ReturnStationId, ReturnStationName, Distance, Duration) VALUES"

				completed := stmtBegin + " " + strings.Join(stmtEnd, ",") + ";"

				// Insert data
				_, err := Database.Exec(completed)
				errorHandler(err, completed)

				// If no input happened, let's quit outer for too.
				if errors.Is(errInner, io.EOF) {
					break
				}
			}
		} else {
			fmt.Printf("File %v has incorrect number of headers, skip.\n", name)
		}

		// DEV (run only one file)
		// break // DEV
	}
}

func importStations() {
	// Open all CSV example files
	file, err := os.Open(STATIONADDRESS)
	errorHandler(err, "")

	defer file.Close()

	fmt.Println("Stations file loaded. Importing...")

	// Skip header row
	read := csv.NewReader(file)
	headers, err := read.Read()
	errorHandler(err, "")

	idValue := 0

	// Check there is 13 headers
	if len(headers) == 13 {
		for {
			i := 0
			stmtEnd := []string{}
			var errInner error

			for i < 100 {
				r, err := read.Read()

				errInner = err

				// No more input available
				if errors.Is(err, io.EOF) {
					break
				}

				// Validate data before importing
				// If the row includes incorrect values (detected by regex), skip row
				// FYI, Some strings in dataset uses (") and (') characters
				// Single quotes can be escaped by doubling them up.
				if validateStationDataBeforeImport(r, r[0]) {
					// Detect if value has single quotes
					for i, item := range r {
						if strings.Contains(item, "'") {
							r[i] = strings.Replace(item, "'", "''", -1)
						}
					}

					//value := "('" + strconv.Itoa(idValue) + "','" + strings.Join(r[1:], "','") + "')"
					value := "('" + strings.Join(r, "','") + "')"

					// Include to same array
					stmtEnd = append(stmtEnd, value)
					idValue += 1
				}
				i += 1
			}

			stmtBegin := "INSERT INTO Stations(Fid, Id, Nimi, Namn, Name, Osoite, Adress, Kaupunki, Stad, Operaattor, Kapasiteet, x, y) VALUES"

			completed := stmtBegin + " " + strings.Join(stmtEnd, ",") + ";"

			// Insert data
			_, err := Database.Exec(completed)
			errorHandler(err, completed)

			if errors.Is(errInner, io.EOF) {
				break
			}
		}

	} else {
		fmt.Println("Loaded stations file has incorrect number of headers, skip.")
	}
}

func validateStationDataBeforeImport(data []string, id string) (validated bool) {
	validated = true

	// Stop if value type is not correct
	errHandler := func(value string) {
		// Disabled for faster loading
		fmt.Printf("Validation error: [%v] Value: %v - Not importing row.\n", id, value)

		validated = false
	}

	// Validate values using regex
	// ID only number [0-9+]
	regexId, _ := regexp.Compile(`^\d+$`)

	// Number can have a dot between numbers
	regexNumber, _ := regexp.Compile(`^\d+(?:\.\d+)?$`)

	// Going thru every value in array
	if !regexId.MatchString(data[0]) || data[0] == "" {
		errHandler(data[0])
	}
	if !regexId.MatchString(data[1]) || data[1] == "" {
		errHandler(data[1])
	}
	if data[2] == "" {
		errHandler(data[2])
	}
	if data[3] == "" {
		errHandler(data[3])
	}
	if data[4] == "" {
		errHandler(data[4])
	}
	if data[5] == "" {
		errHandler(data[5])
	}
	if data[6] == "" {
		errHandler(data[6])
	}

	// data[7-9] can be empty (Kaupunki, Stad, Operaattor)

	if !regexNumber.MatchString(data[11]) || data[10] == "" {
		errHandler(data[10])
	}
	if !regexNumber.MatchString(data[11]) || data[11] == "" {
		errHandler(data[11])
	}
	if !regexNumber.MatchString(data[12]) || data[12] == "" {
		errHandler(data[12])
	}

	// Return true, if everything pass
	// Returns false, if one of the if conditions fails
	return
}

func validateJourneyDataBeforeImport(data []string, id int) (validated bool) {
	validated = true

	// Stop if value type is not correct
	errHandler := func(value string) {
		// Disabled for faster loading
		//fmt.Printf("Validation error: [%v] Value: %v - Not importing row.\n", id, value)

		validated = false
	}

	// Validate values using regex
	// ID only number [0-9+]
	regexId, _ := regexp.Compile(`^\d+$`)

	// Number can have a dot between numbers
	regexNumber, _ := regexp.Compile(`^\d+(?:\.\d+)?$`)

	// Time and time must be realistic (year 2000-2999) and string includes "T"
	regexTime, _ := regexp.Compile("^[2][0-9]{3}-(0[0-9]|1[0-2])-([012][0-9]|3[0-1])T([01][0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9])$")

	// Going thru every value in array
	if !regexTime.MatchString(data[0]) || data[0] == "" {
		errHandler(data[0])
	}
	if !regexTime.MatchString(data[1]) || data[1] == "" {
		errHandler(data[1])
	}
	if !regexId.MatchString(data[2]) || data[2] == "" {
		errHandler(data[2])
	}
	if data[3] == "" {
		errHandler(data[3])
	}
	if !regexId.MatchString(data[4]) || data[4] == "" {
		errHandler(data[4])
	}
	if data[5] == "" {
		errHandler(data[5])
	}
	if !regexNumber.MatchString(data[6]) || data[6] == "" {
		errHandler(data[6])
	}
	if !regexNumber.MatchString(data[7]) || data[7] == "" {
		errHandler(data[7])
	}

	return
}

func errorHandler(err error, completed string) {
	if err != nil {
		//fmt.Println(completed)
		panic(err)
	}
}
