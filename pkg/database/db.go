package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
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

func createTables() {
	userTable := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		username TEXT UNIQUE,
		password_hash TEXT
	);`

	mangaTable := `
	CREATE TABLE IF NOT EXISTS manga (
		id TEXT PRIMARY KEY,
		title TEXT,
		author TEXT
	);`

	progressTable := `
	CREATE TABLE IF NOT EXISTS reading_progress (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		manga_id TEXT,
		chapter INTEGER
	);
	`
	readingListTable := `
	CREATE TABLE IF NOT EXISTS reading_list (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		manga_id TEXT
	);
	`
	

	_, err := DB.Exec(userTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(mangaTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(progressTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(readingListTable)
	if err != nil {
		log.Fatal(err)
	}

	DB.Exec("INSERT OR IGNORE INTO manga(id, title, author) VALUES ('blue-box','Blue Box','Kouji Miura')")
	DB.Exec("INSERT OR IGNORE INTO manga(id, title, author) VALUES ('oshi-no-koi','Oshi no Ko','Aka Akasaka')")
}