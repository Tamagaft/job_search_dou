package repository

import (
	"database/sql"
	"strings"
)

type Database struct {
	database *sql.DB
}

func InitDB() (*Database, error) {
	database, err := sql.Open("sqlite3", "./jobs.db")
	if err != nil {
		return nil, err
	}
	return &Database{database: database}, nil
}

func (db Database) CreateInitialTables() error {
	statement, err := db.database.Prepare("CREATE TABLE IF NOT EXISTS works(id INTEGER PRIMARY KEY,title TEXT, company TEXT, cities TEXT, link TEXT, hash TEXT)")
	if err != nil {
		return err
	}
	statement.Exec()
	return nil
}

func (db Database) GetAllWork() ([]*Work, error) {
	rows, err := db.database.Query("SELECT * FROM works")
	if err != nil {
		return nil, err
	}

	result := make([]*Work, 0, 10)
	var id int
	var title, company, cities, link, hash string
	for rows.Next() {
		rows.Scan(&id, &title, &company, &cities, &link, &hash)
		result = append(result, &Work{Id: id, Title: title, Company: company, Cities: strings.Split(cities, ","), Link: link, Hash: hash})
	}
	return result, nil
}
