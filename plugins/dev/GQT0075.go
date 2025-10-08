package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("GQT0075", "ZBNI", "https://zbni.ph/category/latest/")

	engine.SetStartingURLs([]string{
		"https://zbni.ph/category/latest/"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML("h2>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML("h1.single-title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})
	engine.OnHTML(".single-meta>div:nth-child(1)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		fulltextwithdate := element.Text
		textwithdate := strings.Split(fulltextwithdate, "AM PHT")
		ctx.PublicationTime = textwithdate[0]
	})
	engine.OnHTML(".single-meta>div>img:nth-child(1)+a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	engine.OnHTML(".entry-content>p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	engine.OnHTML("source+img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = []string{element.Attr("src")}
	})
	engine.OnHTML("h2>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Host = element.Text
	})
	engine.OnHTML(".next", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
