package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("GQT0066", "St.Luke's Medical Center Quezon City.Global City", "www.stlukes.com.ph")

	engine.SetStartingURLs([]string{
		"https://www.stlukes.com.ph/news-and-events/news-and-press-release"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML("h2>a,.article-details>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML(".page-title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})
	engine.OnHTML("h1>span", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.SubTitle += element.Text
	})
	engine.OnHTML("h1+div", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		fulltextwithdate := element.Text
		textwithdate := strings.Split(fulltextwithdate, "Posted on")
		ctx.PublicationTime = textwithdate[1]
	})
	engine.OnHTML(".article-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	engine.OnHTML(".banner>img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = []string{element.Attr("src")}
	})
	engine.OnHTML("h2>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Host = element.Text
	})
	engine.OnHTML(".col-md-12>ul>li:nth-child(1)>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.FirstChannel = element.Text
	})
	engine.OnHTML(".col-md-12>ul>li:nth-child(2)>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.SecondChannel = element.Text
	})
	engine.OnHTML(".col-md-12>ul>li:nth-child(3)>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.ThirdChannel = element.Text
	})
	engine.OnHTML("[rel=\"next\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
