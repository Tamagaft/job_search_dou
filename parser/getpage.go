package parser

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
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
