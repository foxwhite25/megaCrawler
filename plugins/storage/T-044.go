package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("T-044", "Ayala Corporation ", "https://ayalaland.com/")

	engine.SetStartingURLs([]string{"https://ayalaland.com/blog"})

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

	engine.OnHTML(".flex.flex-col.gap-5 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".flex.flex-col.divide-y  a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".text-black.py-12.scroll-mt-8> div  > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".relative.z-10 > div > div > span", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML(".mb-6.w-full.overflow-hidden.rounded-md > img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = []string{element.Attr("src")}
	})

	engine.OnHTML(".mx-auto.flex.max-w-\\[46rem\\].flex-col.gap-2 > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find(".mx-auto.flex.max-w-\\[46rem\\].flex-col.gap-2 > p > a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

	engine.OnHTML(".min-h-screen h2,.min-h-screen p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find(".min-h-screen p > a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

	engine.OnHTML(".ml-0.flex.w-full.flex-row.flex-wrap.items-center.justify-between.gap-4 > li + li + li > a.items-center.justify-center.h-12", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

}
