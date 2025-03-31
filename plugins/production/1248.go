package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1248", "THE LID", "https://lidblog.com/")

	engine.SetStartingURLs([]string{
		"https://lidblog.com/category/us-news/",
		"https://lidblog.com/category/politics/",
		"https://lidblog.com/category/culture/",
		"https://lidblog.com/category/econ/",
		"https://lidblog.com/category/faith-morality/",
		"https://lidblog.com/category/israel/",
		"https://lidblog.com/category/foreign-policy/"})

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
	engine.OnHTML(".entry-title > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML(".alignleft > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
