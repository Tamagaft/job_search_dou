package main

import (
	"time"

	"searcher.com/test/config"
	"searcher.com/test/errorhandler"
	"searcher.com/test/loger"
	"searcher.com/test/parser"
	"searcher.com/test/repository"

	"searcher.com/test/db"
)

func main() {
	err := db.CreateInitialTables()
	errorhandler.HandlePanicError(err)
	loger.Info("Initialized database")

	var cr repository.CategoryRepository
	isInited, err := cr.IsInitedCategories()
	errorhandler.HandlePanicError(err)
	if !isInited {
		doc, err := parser.GetPageData(config.URL)
		errorhandler.HandlePanicError(err)
		parser.ParseSaveCategories(doc)
	}
	loger.Info("Checked categories in db")

	categories, err := cr.GetAllCategories()
	errorhandler.HandlePanicError(err)
	loger.Info("Got categories from db")

	for {
		loger.Info("New works parsing")
		for _, category := range categories {
			for _, year := range config.URLEXP {
				go parser.DownloadParseWorks(category.Title, year)
			}
		}

		time.Sleep(time.Hour)
	}

}
