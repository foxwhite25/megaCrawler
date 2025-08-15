package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("SR0032", "ORIX", "https://www.orix.co.jp/grp/")

	engine.SetStartingURLs([]string{"https://www.orix.co.jp/orem/sitemap.xml"})

	extractorConfig := extractors.Config{
		Author:       false, //无作者
		Image:        true,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "news20") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML(`.date`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
		ctx.Authors = append(ctx.Authors, "None")
	})

	engine.OnHTML(".news", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
