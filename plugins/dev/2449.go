package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2449", "411MANIA", "https://411mania.com/")

	engine.SetStartingURLs([]string{"https://411mania.com/movies/category/news/"})

	extractorConfig := extractors.Config{
		Author:       false,
		Image:        false,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML(".article_section_link > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("div.next-prev-wrap > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML("div.info", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		parts := strings.Split(element.Text, "|")
		if len(parts) > 0 {
			ctx.PublicationTime = strings.TrimSpace(parts[0])
		}
	})

	engine.OnHTML("span.author.name > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	engine.OnHTML(".thumbnail > img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = []string{element.Attr("src")}
	})

	engine.OnHTML(".post > div.content > p,.post > div.content > blockquote", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
