package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("SR0034", "SAFETY4SEA", "https://safety4sea.com/")

	engine.SetStartingURLs([]string{
		"https://safety4sea.com/category/safety-parent/",
		"https://safety4sea.com/category/seafit/",
		"https://safety4sea.com/category/green/",
		"https://safety4sea.com/category/smart-parent/",
		"https://safety4sea.com/category/risk/",
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

	engine.OnHTML("h3 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".jeg_meta_author:not(span)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		fulltextwithauthor := element.Text
		textwithauthor := strings.Split(fulltextwithauthor, "by")
		authorpart := strings.TrimSpace(textwithauthor[1])
		authorpart = strings.ReplaceAll(authorpart, "meta_text", "")
		authorpart = strings.TrimSpace(authorpart)
		ctx.Authors = append(ctx.Authors, authorpart)
	})

	engine.OnHTML(".intro-text > p, .content-inner > .bialty-container > :not(div)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

	engine.OnHTML("a.page_nav.next", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
