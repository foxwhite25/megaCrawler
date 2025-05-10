package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("z-0059", "AIST", "https://www.aist.go.jp/")

	engine.SetStartingURLs([]string{"https://www.aist.go.jp/aist_e/news/topics/l_news_topics.html?PageNo=1"})

	extractorConfig := extractors.Config{
		Author:       false,
		Image:        false,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)
	engine.OnHTML("p.newsTitle > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML("a.next", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.Index)
	})
	engine.OnHTML("div.contents_text > p,div.newsDescription > p,div.newsDescription > div > p,div.contents > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	engine.OnHTML("p.newsDate", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})
}
