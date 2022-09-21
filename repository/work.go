package repository

import (
	"strings"

	"searcher.com/test/db"
	"searcher.com/test/types"
)

type WorkRepository struct{}

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
