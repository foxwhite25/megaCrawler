package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1075", "英国每日邮报", "https://www.dailymail.co.uk/home/index.html")

	engine.SetStartingURLs([]string{
		"https://www.dailymail.co.uk/news/index.html",
		"https://www.dailymail.co.uk/news/royals/index.html",
		"https://www.dailymail.co.uk/ushome/index.html",
		"https://www.dailymail.co.uk/health/index.html",
		"https://www.dailymail.co.uk/sciencetech/index.html",
		"https://www.dailymail.co.uk/tvshowbiz/index.html",
		"https://www.dailymail.co.uk/sport/index.html",
		""})

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

	engine.OnHTML(".articletext > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

}
