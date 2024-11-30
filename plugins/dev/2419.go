package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2419", "The Vancouver Sun (British Columbia)", "https://vancouversun.com/")

	engine.SetStartingURLs([]string{
		"https://vancouversun.com/category/news/local-news/",
		"https://vancouversun.com/category/news/politics/",
		"https://vancouversun.com/category/health/",
		"https://vancouversun.com/category/news/national/",
		"https://vancouversun.com/category/news/true-crime/",
		"https://vancouversun.com/category/news/crime/",
		"https://vancouversun.com/category/news/world/",
		"https://vancouversun.com/tag/education/",
	})

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

	engine.OnHTML("div.article-card__details > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.News)
	})

	engine.OnHTML("ul.pagination.list-unstyled > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.Index)
	})

	engine.OnHTML("figure.featured-image > picture > img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = []string{element.Attr("src")}
	})

	engine.OnHTML("div.article-header__detail__texts > p.article-subtitle", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.SubTitle += element.Text
	})

	engine.OnHTML("section.article-content__content-group > p:not(:has(em))", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
