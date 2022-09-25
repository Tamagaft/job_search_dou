package parser

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"searcher.com/test/loger"
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

	loger.Info(fmt.Sprintf("Downloaded page: %s", url))

	return doc, nil
}
