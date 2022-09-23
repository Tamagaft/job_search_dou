package main

import (
	"searcher.com/test/config"
	"searcher.com/test/errorhandler"
	"searcher.com/test/parser"
	"searcher.com/test/repository"

	"searcher.com/test/db"
)

func main() {
	err := db.CreateInitialTables()
	errorhandler.HandlePanicError(err)

	var cr repository.CategoryRepository
	isInited, err := cr.IsInitedCategories()
	errorhandler.HandlePanicError(err)
	if !isInited {
		doc, err := parser.GetPageData(config.URL)
		errorhandler.HandlePanicError(err)
		parser.ParseSaveCategories(doc)
	}

}
