# Helsinki city bike app

This is pre-assignment for Solita Dev Academy Finland (fall 2022).

> Journey dataset:
> - https://dev.hsl.fi/citybikes/od-trips-2021/2021-05.csv
> - https://dev.hsl.fi/citybikes/od-trips-2021/2021-06.csv
> - https://dev.hsl.fi/citybikes/od-trips-2021/2021-07.csv

- - -

### Data validating and checking
- Validates data using regex
    - ID Regex: 
        >`^\d+$`
        > 
        > (E.g. "093")
    - Number Regex:
        > `^\d+(?:\.\d+)?$`
        > 
        > ("215" and "124.4")
    - Time Regex:
        > `^[2][0-9]{3}-(0[0-9]|1[0-2])-([012][0-9]|3[0-1])T([01][0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9])$`
        > 
        >("2021-05-21T18:34:11")

- Checks data
    - If a journey lasted for less than 10 seconds
    - If a journey covered distance is shorter than 10 meters

- Import and insert lasts ~2min

### Database
- Uses [SQLite](https://www.sqlite.org/index.html) with `go-sqlite3` driver by [mattn](https://github.com/mattn/go-sqlite3)
- Creates database table with primary key integer
    > ```txt
    > Id                   INTEGER PRIMARY KEY,
	> Departure            TEXT NOT NULL,
	> Return               TEXT NOT NULL,
	> DepartureStationId   TEXT NOT NULL,
	> DepartureStationName TEXT NOT NULL,
	> ReturnStationId      TEXT NOT NULL,
	> ReturnStationName    TEXT NOT NULL,
	> Distance             TEXT NOT NULL,
	> Duration             TEXT NOT NULL
    >```

- - -

### To-Do

#### Data import
##### Recommended
- [x] Import data from the CSV files to a database or in-memory storage
- [x] Validate data before importing
- [x] Don't import journeys that lasted for less than ten seconds
- [x] Don't import journeys that covered distances shorter than 10 meters

After filtering short (time and distance) journeys, `2021-05.csv` input file's data dropped from `814676` to `784794` rows.

### Learnings ("a-HA" moments)
- Go (Golang) as a new language
    - Using Go packages
    - Use multiple `.go` files
    - The file path must be relative to the folder, where `main.go` file is.
- Creating, connecting and using databases