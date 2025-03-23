package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1205", "日本广播协会", "https://www.nhk.or.jp/nhkworld/")

	engine.SetStartingURLs([]string{
		"https://www3.nhk.or.jp/nhkworld/en/news/list/?p=1",
		"https://www3.nhk.or.jp/nhkworld/en/news/tags/2/",
		"https://www3.nhk.or.jp/nhkworld/en/news/tags/58/",
		"https://www3.nhk.or.jp/nhkworld/en/news/backstories/"})

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

	engine.OnHTML(" .c-articleMedia > a  , .c-article >div > a ", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".c-pagenation > a:last-of-type", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.Index)
	})
}
