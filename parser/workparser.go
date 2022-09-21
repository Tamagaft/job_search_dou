package parser

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"searcher.com/test/repository"
	"searcher.com/test/types"
)

func Manualupdate(html string) error {
	data := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(data)
	if err != nil {
		return err
	}
	return ParseWork(doc)
}

func ParseWork(doc *goquery.Document) error {
	doc.Find(".lt").Find(".vacancy").Each(parseSaveWork)

	return nil
}

func parseSaveWork(index int, item *goquery.Selection) {
	work := types.Work{}

	titleA := item.Find(".title").Find("a")
	work.Title = strings.TrimSpace(titleA.Text())

	link, _ := titleA.Attr("href")
	work.Link = strings.TrimSpace(link)

	work.Company = strings.TrimSpace(item.Find(".company").Text())

	work.Cities = strings.Split(strings.TrimSpace(item.Find(".cities").Text()), ", ")

	work.SetHash()

	var wr repository.WorkRepository
	wr.Save(work)
}
