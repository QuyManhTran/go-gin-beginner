package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {
	log.SetPrefix("[Database] ")
	dbFile := "vks.db"

	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		log.Println("Database file does not exist, creating...")
		file, err := os.Create(dbFile)
		if err != nil {
			log.Println("Error creating database file:", err)
		}
		file.Close()
	}

	_db, err := sql.Open("sqlite3", dbFile)
	db = _db
	if err != nil {
		panic(err)
	}

	// Create a "users" table
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}
	log.Println("🚀 Connect database successfuly!")
}

func CloseDB() {
	db.Close()
}

func GetDB() *sql.DB {
	return db
}