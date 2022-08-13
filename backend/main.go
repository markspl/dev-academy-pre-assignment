package main

import (
	"fmt"

	"backend/database"
)

func main() {
	fmt.Print("\n### Initialize database\n\n")
	database.InitDatabase()
}
