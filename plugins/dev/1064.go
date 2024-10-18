package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"time"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1064", "examinerlive", "https://www.examinerlive.co.uk/")

	engine.SetStartingURLs([]string{
		"https://www.examinerlive.co.uk/news/local-news/",
		"https://www.examinerlive.co.uk/news/west-yorkshire-news/",
		"https://www.examinerlive.co.uk/news/health/",
		"https://www.examinerlive.co.uk/all-about/education",
		"https://www.examinerlive.co.uk/all-about/crime",
		"https://www.examinerlive.co.uk/all-about/politics",
		"https://www.examinerlive.co.uk/news/uk-world-news/",
	})

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

	extractorConfig.Apply(engine)

	engine.OnHTML(".headline", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".pagi-next > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	time.Sleep(400 * time.Microsecond)
}
