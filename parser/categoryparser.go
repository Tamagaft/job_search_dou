package parser

import (
	"github.com/PuerkitoBio/goquery"
	"searcher.com/test/repository"
	"searcher.com/test/types"
)

func ParseSaveCategories(doc *goquery.Document) {
	doc.Find("select").Find("option").Each(ParseSaveCategoriesEach)
}

func ParseSaveCategoriesEach(index int, item *goquery.Selection) {
	category := types.Category{}

	var cr repository.CategoryRepository

	cr.Save(category)
}
