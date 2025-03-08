package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func init() {
	engine := crawlers.Register("1273", "People for the Ethical Treatment of Animals (PETA)", "https://www.peta.org/")

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         true,
		Title:        true,
		TextLanguage: "",
	}

	engine.OnLaunch(func() {
		baseURL := "https://www.peta.org/blog/page/"
		for i := 1; true; i++ {
			if engine.Test != nil && engine.Test.Done {
				return
			}

			pageURL := baseURL + strconv.Itoa(i) + "/?pl=78_15_0"
			resp, err := http.Get(pageURL)
			if err != nil {
				continue
			}

			dom, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				continue
			}
			urls := dom.Find(".grid-item__thumb > a")
			if len(urls.Nodes) == 0 {
				break
			}
			urls.Each(func(i int, selection *goquery.Selection) {
				pageURL, ok := selection.Attr("href")
				if !ok {
					return
				}
				engine.Visit(pageURL, crawlers.News)
			})

			err = resp.Body.Close()
			if err != nil {
				continue
			}
		}
	})
	extractorConfig.Apply(engine)
}
