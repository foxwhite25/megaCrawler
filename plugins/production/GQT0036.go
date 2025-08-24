package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("GQT0036", "Nigerian Bulletin", "https://nigerianbulletin.com/")

	engine.SetStartingURLs([]string{
		"https://nigerianbulletin.com/news/",
		"https://nigerianbulletin.com/lifestyle/"})

	extractorConfig := extractors.Config{
		Author:       false,
		Image:        false,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        false,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML(".structItem-title > a + a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".p-title-value:not(span)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	engine.OnHTML(".u-srOnly + a > span", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	engine.OnHTML(".listInline--bullet .u-concealed > time", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Attr("datetime")
	})

	engine.OnHTML(".bbWrapper", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("a").Remove()
		element.DOM.Find("img").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

	engine.OnHTML(".block-outer--after .pageNav-main + a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
