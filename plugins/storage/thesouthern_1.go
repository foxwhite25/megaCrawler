package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("thesouthern_1", "南伊利诺伊州报", "https://thesouthern.com/")

	engine.SetStartingURLs([]string{
		"https://thesouthern.com/tncms/sitemap/editorial.xml",
		"https://thesouthern.com/tncms/sitemap/editorial.xml?year=2024",
		"https://thesouthern.com/tncms/sitemap/editorial.xml?year=2023",
		"https://thesouthern.com/tncms/sitemap/editorial.xml?year=2022",
	})

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

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "/sitemap/editorial.xml?date") {
			engine.Visit(element.Text, crawlers.Index)
		} else if !strings.Contains(element.Text, ".xml") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML("li.visible-print > time.tnt-date.asset-date.text-muted", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Attr("datetime"))
	})

	engine.OnHTML(`meta[name="author"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Attr("content"))
	})

	engine.OnHTML("div.subscriber-preview > p, div.subscriber-only > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
