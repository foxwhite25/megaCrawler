package production

import (
	"strings"

	"github.com/gocolly/colly/v2"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"
)

func init() {
	engine := crawlers.Register("1016", "英国每日邮报", "https://www.dailymail.co.uk/home/index.html")

	engine.SetStartingURLs([]string{"https://www.dailymail.co.uk/google-news-sitemap.xml"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         true,
		Title:        true,
		TextLanguage: "en",
	}

	extractorConfig.Apply(engine)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, ".xml") {
			ctx.Visit(element.Text, crawlers.Index)
			return
		}
		ctx.Visit(element.Text, crawlers.News)
	})
}
