package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func Connect(path string) *sql.DB {

	if _, err := os.Stat(path); err != nil {
		file, _ := os.Create(path)
		file.Close()

	}

	connection, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal("error opening database: ", err)

	}

	CreateTables(connection)

	return connection
}
