package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1497", "工业部", "https://www.kemenperin.go.id/")

	engine.SetStartingURLs([]string{
		"https://www.kemenperin.go.id/siaran-pers",
		"https://www.kemenperin.go.id/berita-industri",
		"https://www.kemenperin.go.id/kegiatan",
	})

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

	engine.OnHTML("ul.listitems > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.News)
	})

	engine.OnHTML(".pagination > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.Index)
	})

	engine.OnHTML(".col-md-12.col-lg-12.col-xs-12.col-sm-12 > img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		imageURL := element.Request.AbsoluteURL(element.Attr("src"))
		ctx.Image = []string{imageURL}
	})

	engine.OnHTML(".col-md-12.col-lg-12.col-xs-12.col-sm-12 > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
