package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1701", "海军分析中心", "https://www.cna.org/")

	engine.SetStartingURLs([]string{"https://www.cna.org/x1542"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         true,
		Title:        false,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnResponse(func(response *colly.Response, ctx *crawlers.Context) {
		println(string(response.Body))
	})

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.HasSuffix(element.Text, ".xml") {
			return
		}
		if strings.Contains(element.Text, "expert") {
			engine.Visit(element.Text, crawlers.Expert)
		} else if strings.Contains(element.Text, "report") {
			engine.Visit(element.Text, crawlers.Report)
		} else {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML("h1", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	engine.OnHTML(".article-text", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})
}
