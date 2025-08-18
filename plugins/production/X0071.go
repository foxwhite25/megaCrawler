package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("X0071", "Men's Folio", "https://www.mens-folio.com/")

	engine.SetStartingURLs([]string{
		"https://www.mens-folio.com/category/style/style-editors-pick/",
		"https://www.mens-folio.com/category/style/interview/",
		"https://www.mens-folio.com/category/style/style-news/",
		"https://www.mens-folio.com/category/time/time-editors-pick/",
		"https://www.mens-folio.com/category/time/time-news/",
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

	engine.OnHTML(".article-link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".author", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		fulltextwithauthor := element.Text
		textwithauthor := strings.Split(fulltextwithauthor, "By ")[1]
		authorpart := strings.TrimSpace(textwithauthor)
		ctx.Authors = append(ctx.Authors, authorpart)
	})

	engine.OnHTML(".col-12 > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.SubTitle += element.Text
	})

	engine.OnHTML(".the_content-mensfolio > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("a").Remove()
		element.DOM.Find("img").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

	engine.OnHTML(".next", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
