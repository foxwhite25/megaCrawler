package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1019", "Institute of Southeast Asian Studies", "https://www.iseas.edu.sg")

	engine.SetStartingURLs([]string{"https://www.iseas.edu.sg"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        false,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnResponse(func(response *colly.Response, ctx *crawlers.Context) {
		println(string(response.Body))
	})

	engine.OnHTML(".mec-color-hover", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	// 获取 Title
	engine.OnHTML(".mec-single-title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Content
	engine.OnHTML(".mec-single-event-description mec-events-content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})

}
