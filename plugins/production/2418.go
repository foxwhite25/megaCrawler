package production

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2418", "United States", "https://www.defensenews.com/")

	engine.SetStartingURLs([]string{
		"https://www.defensenews.com/arc/outboundfeeds/sitemap/?from=100",
		"https://www.defensenews.com/arc/outboundfeeds/sitemap/?from=200",
		"https://www.defensenews.com/arc/outboundfeeds/sitemap/?from=300",
		"https://www.defensenews.com/arc/outboundfeeds/sitemap/?from=400",
		"https://www.defensenews.com/arc/outboundfeeds/sitemap/?from=500",
		"https://www.defensenews.com/arc/outboundfeeds/sitemap/?from=600",
	})

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

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if !strings.Contains(element.Text, "/video/") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML("div.ArticleHeader__LeadArtWrapper-sc-1dhqito-5", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = []string{element.Attr("src")}
	})

	engine.OnHTML("article.default__ArticleBody-sc-1mncpzl-2 > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
