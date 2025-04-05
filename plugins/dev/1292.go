package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1292", "WBEZ", "https://www.wbez.org/")

	engine.SetStartingURLs([]string{"https://www.wbez.org/latest-news"})

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
	engine.OnHTML(".PagePromo-title > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML(".Pagination-nextPage > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href")) //补全为完整URL

		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return //出现错误后打印错误并返回
		}
		engine.Visit(url.String(), crawlers.Index)
	})
	engine.OnHTML(".Image", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = append(ctx.Image, element.Attr("src"))
	})
	engine.OnHTML(".Page-authors > span > a , .Page-authors > span > span", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	engine.OnHTML("body > div.Page-content > div.Page-oneColumn > main > bsp-article-page-body > div.RichTextArticleBody.RichTextBody > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
