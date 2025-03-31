package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2487", "Autoline", "https://www.autoline.tv/")

	engine.SetStartingURLs([]string{
		"https://www.autoline.tv/post-sitemap.xml",
		"https://www.autoline.tv/post-sitemap2.xml",
		"https://www.autoline.tv/post-sitemap3.xml",
	})

	extractorConfig := extractors.Config{
		Author:       false,
		Image:        false,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Text, crawlers.News)
	})

	engine.OnHTML("div.rll-youtube-player", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("data-src"), "https://www.youtube.com") {
			url, err := element.Request.URL.Parse(element.Attr("data-src"))
			if err != nil {
				crawlers.Sugar.Error(err.Error())
				return
			}
			ctx.Video = append(ctx.Video, url.String())
			ctx.PageType = crawlers.Report
		}
	})

	engine.OnHTML("span.entry-author-name", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	engine.OnHTML("div.entry-content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("noscript").Remove()

		directText := element.DOM.Text()
		ctx.Content += directText
	})
}
