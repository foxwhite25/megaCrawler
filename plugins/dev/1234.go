package dev

import (
	"net/http"
	"strconv"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/PuerkitoBio/goquery"
)

func init() {
	engine := crawlers.Register("1234", "保护媒体自由", "https://sanef.org.za/")

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
		baseURL := "https://sanef.org.za/category/news/page/"
		for i := 0; true; i++ {
			if engine.Test != nil && engine.Test.Done {
				return
			}

			pageURL := baseURL + strconv.Itoa(i)
			resp, err := http.Get(pageURL)
			if err != nil {
				continue
			}

			dom, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				continue
			}
			urls := dom.Find(".post-title > a")
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
