package main

import (
	"searcher.com/test/errorhandler"

	"searcher.com/test/db"
)

func main() {
	err := db.CreateInitialTables()
	errorhandler.HandlePanicError(err)

	/*
		var cr repository.CategoryRepository
		isInited, err := cr.IsInitedCategories()
		errorhandler.HandlePanicError(err)
		if !isInited {
			doc, err := parser.GetPageData(config.URL)
			errorhandler.HandlePanicError(err)
			parser.ParseSaveCategories(doc)
		}

		categories, err := cr.GetAllCategories()
		errorhandler.HandlePanicError(err)

		for {
			for _, category := range categories {
				for _, year := range config.URLEXP {
					go parser.DownloadParseWorks(category.Title, year)
				}
			}
			time.Sleep(time.Hour)
		}
	*/
}
