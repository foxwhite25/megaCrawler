package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("z-0039", "PNGIMR", "https://www.pngimr.org.pg/")

	engine.SetStartingURLs([]string{"https://www.pngimr.org.pg/news-events/news/"})

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
	engine.OnHTML("div.blog-entry-readmore > a,div.recent-posts-details-inner > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML("div.older-posts > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.Index)
	})
	engine.OnXML("//html/body//ul//li[2]/text()", func(element *colly.XMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime += strings.TrimSpace(element.Text)
	})
	engine.OnHTML("div.elementor-widget-container > p,div.entry-content p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
