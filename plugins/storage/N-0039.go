package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("N-0039", "Thai PBS", "https://www.thaipbs.or.th/") //这是个泰语网站

	engine.SetStartingURLs([]string{"https://www.thaipbs.or.th/sitemap/sitemap_news_monthly.xml"})

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

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "/sitemap/sitemap-news/") {
			engine.Visit(element.Text, crawlers.Index)
		} else if !strings.Contains(element.Text, ".xml") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML("div.hyQLlk > div:nth-child(2) > span", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})

	engine.OnHTML("#item-description > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += strings.Join(strings.Fields(element.Text), " ")
	})
}
