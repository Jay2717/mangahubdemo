package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite", "mangahub.db")
	if err != nil {
		log.Fatal(err)
	}

	createTables()
}

func Connect() {
	var err error

	DB, err = sql.Open("sqlite3", "./mangahub.db")
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("DB connection failed:", err)
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
		password_hash TEXT
	);
	`

	mangaTable := `
	CREATE TABLE IF NOT EXISTS manga (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		author TEXT,
		status TEXT,
		description TEXT,
		total_chapters INTEGER
	);
	`

	progressTable := `
	CREATE TABLE IF NOT EXISTS user_progress (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		manga_id INTEGER,
		chapter INTEGER,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		UNIQUE(user_id, manga_id)
	);
	`

	libraryTable := `
	CREATE TABLE IF NOT EXISTS user_manga_library (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		manga_id INTEGER,
		status TEXT DEFAULT 'reading',
		added_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		UNIQUE(user_id, manga_id)
	);
	`

	if _, err := DB.Exec(usersTable); err != nil {
		log.Fatal(err)
	}

	if _, err := DB.Exec(mangaTable); err != nil {
		log.Fatal(err)
	}

	if _, err := DB.Exec(progressTable); err != nil {
		log.Fatal(err)
	}

	if _, err := DB.Exec(libraryTable); err != nil {
		log.Fatal(err)
	}
}
