package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2436", "Korea News Gazette", "https://www.koreanewsgazette.com/")

	engine.SetStartingURLs([]string{
		"https://www.koreanewsgazette.com/category/general/",
		"https://www.koreanewsgazette.com/category/finance/",
		"https://www.koreanewsgazette.com/category/arts/",
		"https://www.koreanewsgazette.com/category/health/",
		"https://www.koreanewsgazette.com/category/sports/",
		"https://www.koreanewsgazette.com/category/politics/",
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

	engine.OnHTML("h2.entry-title > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("div.nav-previous > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML(".entry-meta > span.posted-on > a > time", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML("div.entry-content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
