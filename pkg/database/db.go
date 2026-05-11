package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect() {
	var err error

	DB, err = sql.Open("sqlite3", "./data/mangahub.db")

	if err != nil {
		log.Fatal(err)
	}

	createTables()

	log.Println("Database connected")
}

func createTables() {

	usersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE,
		email TEXT UNIQUE,
		password TEXT
	);
	`

	mangaTable := `
	CREATE TABLE IF NOT EXISTS manga (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		author TEXT,
		status TEXT,
		description TEXT
	);
	`

	progressTable := `
	CREATE TABLE IF NOT EXISTS user_progress (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		manga_id INTEGER,
		chapter INTEGER
	);
	`

	DB.Exec(usersTable)
	DB.Exec(mangaTable)
	DB.Exec(progressTable)
}
