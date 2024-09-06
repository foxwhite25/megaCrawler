package storage

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1403", "菲龙网", "https://www.flw.ph/")

	engine.SetStartingURLs([]string{"https://www.flw.ph/forum-40-1.html"})

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

	engine.OnHTML("tr > th > a.s.xst", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("div.pg > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML("td.t_f", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		content := ""

		// 跳过td中的div（div中的文本为广告）
		element.DOM.Contents().Each(func(i int, s *goquery.Selection) {
			if s.Is("div") {
				return
			}
			content += s.Text()
		})
		ctx.Content += content
	})
}
