package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var name string
var roll int

const DATABASE_URL = "postgres://postgres:postgres@localhost:5432/?sslmode=disable"

// Trying getting data
func main() {
	db, err := sql.Open("postgres", DATABASE_URL)
	if err != nil {
		println("failed")
		os.Exit(1)
	}
	fmt.Scanf("%d", &roll)
	fmt.Scanf("%s", &name)
	db.Exec("INSERT INTO students (roll, name) values ($1, '$2')", roll, name)
	result, err := db.Query("SELECT roll FROM students")
	var rolltemp int
	if err != nil {
		println(err)
	} else {
		for result.Next() {
			result.Scan(&rolltemp)
			println(rolltemp)
		}

	}

}
