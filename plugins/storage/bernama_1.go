package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("bernama_1", "Malaysian National News Agency", "https://www.bernama.com/")

	engine.SetTimeout(30 * time.Second)

	engine.SetStartingURLs([]string{"https://www.bernama.com/en/general/"})

	extractorConfig := extractors.Config{
		Author:       false,
		Image:        false,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML("div.row > div.col-7.col-md-12.col-lg-12.mb-3 > h6 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		if strings.Contains(url.String(), "news.php") {
			engine.Visit(url.String(), crawlers.News)
		}
	})

	engine.OnHTML("div.text-center > ul > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.Index)
	})

	engine.OnHTML(`meta[property="article:author"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Attr("content"))
	})

	engine.OnHTML("div.row > div.col-12.mt-3.text-dark.text-justify > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
