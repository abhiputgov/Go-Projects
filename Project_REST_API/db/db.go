package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func IniDB() {
	db, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
}

func createTables() error {
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			datetime DATETIME NOT NULL,
			user_id INTEGER NOT NULL,
			FOREIGN KEY(user_id) REFERENCES users(id)
		);
	`
	_, err := db.Exec(createEventsTable)
	if err != nil {
		panic(err)
	}

	return nil
}
