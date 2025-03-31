package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1242", "WSRW", "https://wsrw.org/en")

	engine.SetStartingURLs([]string{"https://wsrw.org/en/news"})

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
	engine.OnHTML(".text-lg > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML("a[rel=\"next\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href")) // 补全为完整URL

		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return // 出现错误后打印错误并返回
		}
		engine.Visit(url.String(), crawlers.Index)
	})
}
