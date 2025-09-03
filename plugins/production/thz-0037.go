package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("thz-0037", "NTP Pool Project", "https://www.ntppool.org/zone/ph")

	engine.SetStartingURLs([]string{"https://news.ntppool.org/post/"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        false,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML(".blog-posts a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("body > main > content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("body > main > content > p > a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})
}
