package parser

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"searcher.com/test/repository"
)

func GetPageData(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if err != nil {
			return nil, err
		}
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func Parse(doc *goquery.Document) error {
	doc.Find(".lt").Find(".vacancy").Each(parseSaveWork)

	return nil
}

func parseSaveWork(index int, item *goquery.Selection) {
	work := repository.Work{}

	titleA := item.Find(".title").Find("a")
	work.Title = strings.TrimSpace(titleA.Text())

	link, _ := titleA.Attr("href")
	work.Link = strings.TrimSpace(link)

	work.Company = strings.TrimSpace(item.Find(".company").Text())

	work.Cities = strings.Split(strings.TrimSpace(item.Find(".cities").Text()), ", ")

	work.SetHash()

}
