package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("SR0012", "Security Info Watch", "https://www.securityinfowatch.com/")

	engine.SetStartingURLs([]string{
		"https://www.securityinfowatch.com/sitemap/News.xml",
		"https://www.securityinfowatch.com/sitemap/News.2.xml",
		"https://www.securityinfowatch.com/sitemap/News.1.xml",
		"https://www.securityinfowatch.com/sitemap/News.3.xml",
		"https://www.securityinfowatch.com/sitemap/News.4.xml"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Text, crawlers.News)
	})

	engine.OnHTML("div.web-body-blocks  p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML(`div.date-wrapper > div.date`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML("div.byline-group", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})
}
