package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("3409", "BNL Center for Functional Nanomaterials", "https://www.bnl.gov/")

	engine.SetTimeout(60 * time.Second)
	engine.SetStartingURLs([]string{
		"https://www.bnl.gov/newsroom/tags/",
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

	engine.OnHTML("ul.articleSummaries > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("div.grid_6 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML(".articleBody > h3.subhead", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.SubTitle += element.Text
	})

	engine.OnHTML(".articleBody > p.dateLine", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML("div.image-100 > a > img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		imageUrl := element.Request.AbsoluteURL(element.Attr("src"))
		ctx.Image = []string{imageUrl}
	})

	engine.OnHTML("div.articleBody > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
