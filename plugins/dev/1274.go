package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1274", "PETA.UK", "https://www.peta.org.uk/")

	engine.SetStartingURLs([]string{"https://www.peta.org.uk/blog/"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        false,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         true,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)
	engine.OnHTML(".text-content > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML(".pagination >li:last-of-type > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
	engine.OnHTML("P > img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = append(ctx.Image, element.Attr("src"))
	})
	engine.OnHTML("[role=\"article\"] > a , .ytp-title-link , .video-player iframe", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Video = append(ctx.Video, element.Attr("href")) //视频采集
	})
	engine.OnHTML(".video-player iframe", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Video = append(ctx.Video, element.Attr("src")) //视频采集
	})
}
