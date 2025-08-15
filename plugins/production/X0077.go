package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("X0077", "Torque", "https://www.torque.com.sg/")

	engine.SetStartingURLs([]string{"https://www.torque.com.sg/news/"})

	extractorConfig := extractors.Config{
		Author:       false,
		Image:        false,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML("div[class=\"grid-title\"] > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("span[class=\"author\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		fulltextwithauthor := element.Text
		textwithauthors := strings.Split(fulltextwithauthor, " | ")[1]
		textwithauthor := strings.Split(textwithauthors, "Story: ")[1]
		ctx.Authors = append(ctx.Authors, textwithauthor)
	})

	engine.OnHTML("span[class=\"date\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		fulltextwithdate := element.Text
		textwithdate := strings.Split(fulltextwithdate, " | ")[1]
		ctx.PublicationTime = textwithdate
	})

	engine.OnHTML(".single-body-content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML(".pagination-box > .next-page-link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
