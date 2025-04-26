package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("N-0025", "Business", "https://www.business.com/")

	engine.SetStartingURLs([]string{"https://www.business.com/bdc_article-sitemap.xml"})

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
		engine.Visit(element.Text, crawlers.News)
	})

	engine.OnHTML("h1.bdc-svxg0b-HeroTitle", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	engine.OnHTML("span.ArticlePage-authorName > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	engine.OnHTML("span.bdc-1b17nv0-TypeText-ModifiedAt", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.ReplaceAll(element.Text, "Updated ", "")
	})

	engine.OnHTML("div.e1ryxgnn1 p,div.e1ryxgnn1 ul,div.e1ryxgnn1 h2",
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			ctx.Content += element.Text
		})
}
