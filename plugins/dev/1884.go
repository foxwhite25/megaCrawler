package dev

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1884", "MLB Trade Rumors", "https://www.mlbtraderumors.com/")
	// 网站来源："https://mlbreports.com/"主页的“Baseball News”链接
	engine.SetStartingURLs([]string{"https://www.mlbtraderumors.com/wp-sitemap.xml"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		switch {
		case strings.Contains(element.Text, "posts"):
			engine.Visit(element.Text, crawlers.Index)
		case strings.Contains(element.Text, "20"):
			engine.Visit(element.Text, crawlers.News)
		}
	})
	// xml链接过多，容易超时，但能采集
	engine.OnHTML(".entry-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
