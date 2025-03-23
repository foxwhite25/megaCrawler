package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1233", "投资协会", "https://www.theia.org/")

	engine.SetStartingURLs([]string{"https://www.theia.org/news/press-releases"})

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
	engine.OnHTML(".weight--semi-bold > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML(".pager__item--next > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href")) // 补全为完整URL

		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return // 出现错误后打印错误并返回
		}
		engine.Visit(url.String(), crawlers.Index)
	})
}
