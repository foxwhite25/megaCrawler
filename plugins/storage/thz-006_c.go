package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("thz-006", "CMS Wire", "https://www.cmswire.com")

	engine.SetStartingURLs([]string{"https://www.cmswire.com/sitemap.xml"})

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
		engine.Visit(element.Text, crawlers.News)
	})

	engine.OnHTML("div.article-title__right-labels > div:nth-child(1), div.styles_article-body__article-meta__3_do0 > div:nth-child(1)",
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			ctx.PublicationTime = strings.TrimSpace(element.Text)
		})

	engine.OnHTML("div.article-title__right-labels > div:nth-child(2)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Attr("content"))
	})

	engine.OnHTML("div.rich-html > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.element += element.Text
	})

}
