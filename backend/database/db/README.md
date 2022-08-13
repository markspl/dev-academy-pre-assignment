`backend/database/db` folder for SQLite database

- - -

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