package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("N-0016", "MarketBeat", "https://www.marketbeat.com")

	engine.SetStartingURLs([]string{"https://www.marketbeat.com/sitemap-news.ashx"})

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

	engine.OnHTML(".article-image.mb-3 > img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = []string{element.Attr("src")}
	})

	engine.OnHTML(".d-block.c-gray-8.font-smaller", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML(".keypoints.lh-loose.mt-0.mb-3 > ul > li > strong", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description += element.Text
	})

	engine.OnHTML(".body-copy.lh-loose.article-page > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if element.Text != "View The Five Stocks Here " { //清理冗余文本
			ctx.Content += element.Text
		}
	})

}
