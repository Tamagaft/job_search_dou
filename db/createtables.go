package db

import (
	"database/sql"
	"fmt"
)

func CreateInitialTables() error {
	err := initDB()
	if err != nil {
		return fmt.Errorf("can't init DB: %s", err)
	}
	DB := Connection()

	err = CreateWorksTable(DB)
	if err != nil {
		return err
	}

	err = CreateCategoriesTable(DB)
	if err != nil {
		return err
	}

	return nil
}

func CreateWorksTable(DB *sql.DB) error {
	statement, err := DB.Prepare(`
		CREATE TABLE IF NOT EXISTS works
		(
			id INTEGER PRIMARY KEY,
			title TEXT,
			company TEXT,
			cities TEXT,
			category TEXT,
			years TEXT,
			link TEXT,
			hash TEXT
		)
	`)
	if err != nil {
		return err
	}
	statement.Exec()
	return nil
}

func CreateCategoriesTable(DB *sql.DB) error {
	statement, err := DB.Prepare(`
		CREATE TABLE IF NOT EXISTS categories
		(
			id INTEGER PRIMARY KEY,
			title TEXT
			UNIQUE(title)
		)
	`)
	if err != nil {
		return err
	}
	statement.Exec()
	return nil
}
