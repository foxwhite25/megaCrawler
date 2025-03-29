package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2493", "Avweb", "https://www.avweb.com/")

	engine.SetStartingURLs([]string{"https://www.avweb.com/topics/aviation-news/"})

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

	engine.OnHTML("div.MuiBox-root.css-f5x0ou > div.MuiBox-root.css-13pmzdw > a.MuiTypography-h5",
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			engine.Visit(element.Attr("href"), crawlers.News)
		})

	engine.OnHTML(`ul.MuiPagination-ul.css-nhb8h9 > li > button`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit("https://www.avweb.com/topics/aviation-news/page/"+element.Text, crawlers.Index)
	})

	engine.OnHTML("div.MuiBox-root.css-1n8lbbg > div.MuiBox-root.css-1d29nfv > div > img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = []string{element.Attr("src")}
	})

	engine.OnHTML("div.MuiBox-root.css-1rc4wle > div > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
