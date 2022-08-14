# Helsinki city bike app

This is pre-assignment for Solita Dev Academy Finland (fall 2022).

It was really fun to work with the project, and I'll continue working with this in another branch!

- Backend
    - GoLang (go-sqlite3, mux)
    - SQLite

- Frontend
    - React (react, react-dom, react-router, react-router-dom)
    - Axios
    - Bootstrap (react-boostrap)

> Journey dataset:
> - https://dev.hsl.fi/citybikes/od-trips-2021/2021-05.csv
> - https://dev.hsl.fi/citybikes/od-trips-2021/2021-06.csv
> - https://dev.hsl.fi/citybikes/od-trips-2021/2021-07.csv

> Helsinki Region Transport’s (HSL) city bicycle stations:
> - Dataset: <https://opendata.arcgis.com/datasets/726277c507ef4914b0aec3cbcfcbfafc_0.csv>
> - License and information: <https://www.avoindata.fi/data/en/dataset/hsl-n-kaupunkipyoraasemat/resource/a23eef3a-cc40-4608-8aa2-c730d17e8902>

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

##### Journeys
```txt
Id                   INTEGER PRIMARY KEY,
Departure            TEXT NOT NULL,
Return               TEXT NOT NULL,
DepartureStationId   TEXT NOT NULL,
DepartureStationName TEXT NOT NULL,
ReturnStationId      TEXT NOT NULL,
ReturnStationName    TEXT NOT NULL,
Distance             TEXT NOT NULL,
Duration             TEXT NOT NULL
```

##### Station
```txt
FID        int64   `json:"fid"`
ID         string  `json:"id"` // "103", "014", "001"
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
```

- - -

### To-Do

#### Data import
##### Recommended
- [x] Import data from the CSV files to a database or in-memory storage
- [x] Validate data before importing
- [x] Don't import journeys that lasted for less than ten seconds
- [x] Don't import journeys that covered distances shorter than 10 meters

After filtering short (time and distance) journeys, `2021-05.csv` input file's data dropped from `814676` to `784794` rows.

#### Journey list view
##### Recommended
- [x] List journeys
    - [x] GET all journeys (limited to 100) `localhost:3000/api/journeys/`
    - [x] Show on frontend (show departure and return stations, distance (km), and duration (min))

#### Station list
##### Recommended
- [x] List all stations
    - [x] GET all stations `localhost:3000/api/journeys/`

#### Single station view
##### Recommended
- [x] Show
    - [x] station name
    - [x] station address
    - [x] total number of journeys starting from the station
    - [x] total number of journeys ending at the station

#### Other

List of things, which I wanted to implify before returning.

Will be implemented on another branch: https://github.com/markspl/dev-academy-pre-assignment/tree/update

- Dockerize backend
- Authentication in SQLite
- OpenStreetMap on `stations/:id`

- - -

### Learnings ("a-HA" moments)
- Go (Golang) as a new language
    - Using Go packages
    - Use multiple `.go` files
    - The file path must be relative to the folder, where `main.go` file is.
- Creating, connecting and using databases