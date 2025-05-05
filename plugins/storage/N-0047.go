package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"regexp"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("N-0047", "Agence Kampuchea Presse", "https://www.akp.gov.kh")

	engine.SetStartingURLs([]string{"https://www.akp.gov.kh/"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        false,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML("div.post-data > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("ul.pagination > li.page-item > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML("div.youtube-video > p:first-child", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		re := regexp.MustCompile(`([A-Z][a-z]+\s+\d{1,2},\s+\d{4})`)
		match := re.FindStringSubmatch(element.Text)
		if len(match) > 1 {
			ctx.PublicationTime += match[1]
		}
	})

	engine.OnHTML(`meta[name="author"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Attr("content"))
	})

	engine.OnHTML("div.youtube-video > p:not(:first-child)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
