package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("N-0045", "Pasaxon", "https://pasaxon.org.la/")

	engine.SetStartingURLs([]string{"https://pasaxon.org.la/sitemap.xml"})

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
		if strings.Contains(element.Text, "/sitemaps/news-20") {
			engine.Visit(element.Text, crawlers.Index)
		} else if !strings.Contains(element.Text, ".xml") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML("div.article__body > p:last-of-type", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		re := regexp.MustCompile(`ຂ່າວ:\s*([^\s]+)`)
		match := re.FindStringSubmatch(element.Text)
		if len(match) > 1 {
			ctx.Authors = append(ctx.Authors, match[1])
		}
	})

	engine.OnHTML("div.article__body > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
