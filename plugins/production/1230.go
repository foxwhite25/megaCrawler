package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1230", "ALUMINUM", "https://www.aluminum.org/")

	engine.SetStartingURLs([]string{"https://www.aluminum.org/latest-news-0"})

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

	engine.OnHTML("h4 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".pager__link--next", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href")) // 补全为完整URL

		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return // 出现错误后打印错误并返回
		}
		engine.Visit(url.String(), crawlers.Index)
	})
}
