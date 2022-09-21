package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *Database

type Database = sql.DB

func initDB() error {
	if db != nil {
		return nil
	}

	var err error
	db, err = sql.Open("sqlite3", "./jobs.db")
	return err
}

func Connection() *Database {
	return db
}
