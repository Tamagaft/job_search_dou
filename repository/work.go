package repository

import (
	"database/sql"
	"strings"

	"searcher.com/test/db"
	"searcher.com/test/types"
)

type WorkRepository struct{}

const EMPTYRESPONSE int64 = 0

func (wr WorkRepository) Save(work types.Work) error {
	DB := db.Connection()

	statement, err := DB.Prepare(`
        INSERT INTO works
        (
            title,
            company,
            cities,
            link,
			category,
            hash
        )
        VALUES
        (?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}

	statement.Exec(
		work.Title,
		work.Company,
		strings.Join(work.Cities[:], ", "),
		work.Link,
		work.Category,
		work.Hash)

	return err
}

func (wr WorkRepository) GetAllWork() ([]*types.Work, error) {
	DB := db.Connection()
	rows, err := DB.Query("SELECT * FROM works")
	if err != nil {
		return nil, err
	}

	result := make([]*types.Work, 0, 10)
	var id int
	var title, company, cities, link, category, hash string
	for rows.Next() {
		rows.Scan(&id, &title, &company, &cities, &link, &category, &hash)
		result = append(result, &types.Work{Id: id, Title: title, Company: company, Cities: strings.Split(cities, ","), Link: link, Category: category, Hash: hash})
	}
	return result, nil
}

func (wr WorkRepository) IsSaved(hash string) (bool, error) {
	DB := db.Connection()
	if err := DB.QueryRow("SELECT * FROM works where hash = ? ", "has").Scan(); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
