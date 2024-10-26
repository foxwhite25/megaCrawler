package storage

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1465", "参议院", "http://www.sangiin.go.jp/")

	engine.SetStartingURLs([]string{
		"https://www.sangiin.go.jp/japanese/ugoki/index.html",
	})

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

	// 总索引页和具体年份的索引页的链接selector是基本相同的，所以通过它们的strings区分
	engine.OnHTML(".exp_list_icn02 > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), "index.html") {
			url, err := element.Request.URL.Parse(element.Attr("href"))
			if err != nil {
				crawlers.Sugar.Error(err.Error())
				return
			}
			engine.Visit(url.String(), crawlers.Index)
		}
	})

	engine.OnHTML(".exp_list_icn02 > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if !strings.Contains(element.Attr("href"), "index.html") {
			url, err := element.Request.URL.Parse(element.Attr("href"))
			if err != nil {
				crawlers.Sugar.Error(err.Error())
				return
			}
			engine.Visit(url.String(), crawlers.News)
		}
	})

	engine.OnHTML("#ContentsBox > p, #ContentsBox > div > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
