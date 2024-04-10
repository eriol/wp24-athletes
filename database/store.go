package database // import "github.com/eriol/wp24-athletes/database"

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

// Open an SQLite3 database at the specified path.
func Open(path string) (err error) {

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return err
	}

	database = db

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			if err := createDatabase(); err != nil {
				return err
			}
		}

	}

	return nil
}

// Close SQLite3 database used by Store.
func Close() error {
	return database.Close()
}

func createDatabase() error {

	tables := `
    CREATE TABLE athletes (
        slug TEXT NOT NULL PRIMARY KEY,
        name TEXT COLLATE NOCASE,
        gender TEXT,
        age INTEGER,
        sport_id INTEGER,
        famous_for TEXT
    );

    CREATE TABLE sports (
        id INTEGER NOT NULL PRIMARY KEY autoincrement,
        name TEXT COLLATE NOCASE
    );
    `

	_, err := database.Exec(tables)

	if err != nil {
		return err
	}

	return nil
}
