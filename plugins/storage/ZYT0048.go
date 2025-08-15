package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("ZYT0048", "stockmarketsreview", "http://www.stockmarketsreview.com")

	engine.SetStartingURLs([]string{
		"http://www.stockmarketsreview.com/news/category/51-asia/",
		"http://www.stockmarketsreview.com/news/category/48-china/",
		"http://www.stockmarketsreview.com/news/category/50-europe/",
		"http://www.stockmarketsreview.com/news/category/47-india/",
		"http://www.stockmarketsreview.com/news/category/46-malaysia/",
		"http://www.stockmarketsreview.com/news/category/56-metals-industrial/",
		"http://www.stockmarketsreview.com/news/category/55-metals-precious/",
		"http://www.stockmarketsreview.com/news/category/54-oil-gas/",
		"http://www.stockmarketsreview.com/news/category/49-russia/",
		"http://www.stockmarketsreview.com/news/category/45-singapore/",
		"http://www.stockmarketsreview.com/news/category/43-uk/",
		"http://www.stockmarketsreview.com/news/category/41-usa/",
	})

	extractorConfig := extractors.Config{
		Author:       false,
		Image:        false,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         true,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML(".article_preview > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".username > strong", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	engine.OnHTML(".article_username_container", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		fulltextwithdate := element.Text
		textwithdate := strings.Split(fulltextwithdate, "Published on ")[1]
		ctx.PublicationTime = textwithdate
	})

	engine.OnHTML(".postcontainer", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML("a[rel=\"next\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
