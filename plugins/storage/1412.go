package storage

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	//NSCB在2013年已并入菲律宾统计局，现采用统计局的官网（GPT）
	engine := crawlers.Register("1412", "菲律宾统计局", "https://psa.gov.ph/")

	engine.SetStartingURLs([]string{"https://psa.gov.ph/press-releases"})

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

	engine.OnHTML(".action.margin-top-10 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	// 处理分页链接
	engine.OnHTML(".pager__items > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		href := element.Attr("href")
		if strings.Contains(href, "?page=") {
			absoluteURL := element.Request.AbsoluteURL(href)
			engine.Visit(absoluteURL, crawlers.Index)
		}
	})

	engine.OnHTML(".layout__region.layout__region--content > span > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
