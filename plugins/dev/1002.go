package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1002", "峨山研究", "https://www.asaninst.org/contents/category/about-us/media-relation/press-release-media-relation/")

	engine.SetStartingURLs([]string{"https://www.asaninst.org/contents/%ec%84%b8%ea%b3%84%ec%9d%bc%eb%b3%b4-%ed%95%9c%c2%b7%eb%af%b8-%ea%b3%b5%eb%8f%99-%ed%95%b5%ec%9e%91%ea%b3%84%ec%99%80-%eb%8f%99%eb%a7%b9%ec%95%88%eb%b3%b4-%ec%9e%ac%ed%8e%b8/"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         true,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)
	engine.OnHTML("post_tit", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML("page-numbers current", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
	engine.OnHTML("tbody", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})
	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		switch {
		case strings.Contains(ctx.URL, "expert"):
			engine.Visit(element.Text, crawlers.Expert)
		case strings.Contains(ctx.URL, "post"):
			engine.Visit(element.Text, crawlers.News)
		}
	})
	engine.OnHTML("btn_go-list", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
	engine.OnHTML("post-3896", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})
	engine.OnResponse(func(response *colly.Response, ctx *crawlers.Context) {
		crawlers.Sugar.Debugln(response.StatusCode)
		crawlers.Sugar.Debugln(string(response.Body))

	})

}
