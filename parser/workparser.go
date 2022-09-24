package parser

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"searcher.com/test/config"
	"searcher.com/test/repository"
	"searcher.com/test/types"
)

func Manualupdate(html string) error {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return err
	}
	return ParseSaveWorks(doc)
}

func DownloadParseWorks(category, experience string) error {
	doc, _ := GetPageData(fmt.Sprintf("%s?category=%s&exp=%s", config.URL, category, experience))
	// NON PANIC ERROR
	return ParseSaveWorks(doc)
}

func ParseSaveWorks(doc *goquery.Document) error {
	doc.Find(".lt").Find(".vacancy").Each(parseWorkPage)

	return nil
}

func parseWorkPage(index int, item *goquery.Selection) {
	var wr repository.WorkRepository
	work := types.Work{}

	titleA := item.Find(".title").Find("a")

	link, _ := titleA.Attr("href")
	work.Title = strings.TrimSpace(titleA.Text())
	work.Link = strings.TrimSpace(link)

	work.SetHash()

	isSaved, _ := wr.IsSaved(work.Hash)
	//non panic error
	if isSaved {
		return
	}

	workPage, _ := GetPageData(link)
	//non panic error

	work.Company = strings.TrimSpace(workPage.Find(".l-n").Find("a").Text())
	workPage.Find(".breadcrumbs").Find("a").Each(func(i int, sel *goquery.Selection) {
		if i == 2 {
			work.Category = sel.Text()
		}
	})

	work.Cities = strings.Split(strings.TrimSpace(workPage.Find(".sh-info").Find(".place").Text()), ", ")

	wr.Save(work)
}
