package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("X0085", "Producers Bank", "https://www.producersbank.com.ph/")

	engine.SetStartingURLs([]string{"https://www.producersbank.com.ph/?s=bank&searchsubmit=Search"})

	extractorConfig := extractors.Config{
		Author:       false,
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

	engine.OnHTML("h1[class=\"entry-title\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	engine.OnHTML(".author > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Location = "Philippine"
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	engine.OnHTML(".published", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime += element.Attr("title")
	})

	engine.OnHTML(".entry-content > :not(div:nth-last-child(1))", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML(".entry-content a > img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = append(ctx.Image, element.Attr("src"))
	})

	engine.OnHTML("#nav-below > .nav-previous > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
