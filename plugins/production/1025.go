package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1025", "韩国时报", "https://www.koreatimes.co.kr/www2/index.asp")

	engine.SetStartingURLs([]string{
		"https://www.koreatimes.co.kr/www/sublist_113.html",
		"https://www.koreatimes.co.kr/www/sublist_103.html",
		"https://www.koreatimes.co.kr/www/sublist_129.html",
		"https://www.koreatimes.co.kr/www/sublist_602.html",
		"https://www.koreatimes.co.kr/www/sublist_398.html",
		"https://www.koreatimes.co.kr/www/sublist_135.html",
		"https://www.koreatimes.co.kr/www/sublist_600.html",
		"https://www.koreatimes.co.kr/www/sublist_501.html",
	})

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

	engine.OnHTML(".list_article_headline > a ", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("span.pagenation > b+a ", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML(".pagenation_img > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
