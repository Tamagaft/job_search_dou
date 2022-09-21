package repository

import (
	"searcher.com/test/db"
	"searcher.com/test/types"
)

const CATEGORIESCOUNT = 48

type CategoryRepository struct{}

func (cr CategoryRepository) Save(category types.Category) error {
	DB := db.Connection()

	statement, err := DB.Prepare(`
		INSERT OR IGNORE INTO categories
		(
			title
		)
		VALUES
		(?)
	`)
	if err != nil {
		return err
	}

	statement.Exec(category.Title)
	return err

}

func (cr CategoryRepository) GetAllCategories() ([]*types.Category, error) {
	DB := db.Connection()
	rows, err := DB.Query("SELECT * FROM categories")
	if err != nil {
		return nil, err
	}

	result := make([]*types.Category, 0, CATEGORIESCOUNT)
	var id int
	var title string
	for rows.Next() {
		rows.Scan(&id, &title)
		result = append(result, &types.Category{Id: id, Title: title})
	}
	return result, nil
}

func (cr CategoryRepository) FillCategories([]types.Category) error {
	return nil
}

func (cr CategoryRepository) IsInitedCategories() (bool, error) {
	DB := db.Connection()
	rows, err := DB.Query("SELECT COUNT(*) FROM categories")
	if err != nil {
		return false, err
	}
	var count int
	rows.Scan(&count)
	if count != CATEGORIESCOUNT {
		return false, nil
	}
	return true, nil
}
