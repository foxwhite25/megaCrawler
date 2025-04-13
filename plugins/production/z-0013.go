package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("z-0013", "WHAT'S NEWS", "https://whatsnewindonesia.com/")

	engine.SetStartingURLs([]string{"https://whatsnewindonesia.com/search/node?keys=NEWS&f%5B0%5D=type%3Adeal&f%5B1%5D=type%3Adirectory&f%5B2%5D=type%3Aevent&f%5B3%5D=type%3Afeature&f%5B4%5D=type%3Agallery&f%5B5%5D=type%3Aportfolio&f%5B6%5D=type%3Aultimate_guide&page=0%2C0%2C0%2C0%2C0%2C0%2C0%2C0%2C0%2C0%2C0"})

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
	engine.OnHTML("h3.title > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML("li.pager__item--next > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.Index)
	})
	engine.OnHTML("p.text-align-justify", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
