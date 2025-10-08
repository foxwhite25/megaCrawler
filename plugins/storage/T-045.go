package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("T-045", "Banco de Oro", "https://www.bankcom.com.ph/")

	engine.SetStartingURLs([]string{"https://www.bankcom.com.ph/news/"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        false,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML(".entry-title > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".entry-content img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		imageUrl := element.Request.AbsoluteURL(element.Attr("src"))
		ctx.Image = []string{imageUrl}
	})

	engine.OnHTML(".entry-content span", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find(".entry-content span > a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

	engine.OnHTML(".entry-meta > time", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML(".archive-pagination.pagination li+ li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

}
