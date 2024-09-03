package storage

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1703", "世界日报", "https://worldnews.net.ph/")

	engine.SetStartingURLs([]string{
		"https://worldnews.net.ph/category/hua-she/",
		"https://worldnews.net.ph/category/ben-dao/",
		"https://worldnews.net.ph/category/guo-ji/",
		"https://worldnews.net.ph/category/jing-ji/",
		"https://worldnews.net.ph/category/bao-dao/",
		"https://worldnews.net.ph/category/guang-chang/",
		"https://worldnews.net.ph/category/she-lun/",
	})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML(".post-content > div > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".post-content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML(".pagination > li:nth-child(13) > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
