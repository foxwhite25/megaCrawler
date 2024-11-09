package dev

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1496", "能矿部", "https://www.esdm.go.id/")

	engine.SetStartingURLs([]string{"https://www.esdm.go.id/en/media-center/arsip-berita"})

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

	engine.OnHTML(".title.p-3 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.News)
	})

	engine.OnHTML("li.next > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.Index)
	})

	engine.OnHTML("div.date.mb-3", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		parts := strings.Split(element.Text, "-")

		if len(parts) == 2 {
			ctx.PublicationTime = parts[0] // 获取date
		}
	})

	engine.OnHTML(".col-md-9.news-read.pt-3 > img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		imageURL := element.Request.AbsoluteURL(element.Attr("src"))
		ctx.Image = []string{imageURL}
	})

	engine.OnHTML(`.col-md-9.news-read.pt-3 > p[style*="text-align:justify;"],.col-md-9.news-read.pt-3 > p[align="justify"],
	.col-md-9.news-read.pt-3 > div[align="justify"]`,
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			ctx.Content += element.Text
		})
}
