package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
<<<<<<< HEAD
	// removes db if exists
=======
	// Deletes DB on each execution
>>>>>>> testing
	os.Remove("database.db")

	// Opens DBs through sqlite3 driver
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	queryDb := `
  CREATE TABLE contacts (
  contact_id INTEGER PRIMARY KEY AUTOINCREMENT,
  first_name TEXT NOT NULL,
  last_name TEXT NOT NULL,
  email TEXT NOT NULL,
  address TEXT NOT NULL)`

	// Executes Query
	_, err = db.Exec(queryDb)
	if err != nil {
		log.Fatal(err)
	}

	// Prepars row into table
	stmt, err := db.Prepare("INSERT INTO contacts (first_name, last_name, email, address) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	// Inserts info to row mentioned above
	_, err = stmt.Exec("Hernan", "Di Tano", "hditano@gmail.com", "Imbramowska 3")
	if err != nil {
		log.Fatal(err)
	}
}
