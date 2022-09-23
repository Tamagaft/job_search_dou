package parser

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"searcher.com/test/repository"
)

func ParseSaveCategories(doc *goquery.Document) {
	doc.Find("select").Find("option").Each(ParseSaveCategoriesEach)
}

func ParseSaveCategoriesEach(index int, item *goquery.Selection) {
	var cr repository.CategoryRepository
	cr.Save(strings.TrimSpace(item.Text()))
}
