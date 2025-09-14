package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("X0086", "FICOBank", "https://www.ficobank.com/")

	engine.SetStartingURLs([]string{"https://www.ficobank.com/news/NewsArchive.htm"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML(".newsLink > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		FICOBankBasicURL := "https://www.ficobank.com/news/"
		FICOBankJSONURL := element.Attr("href")
		FICOBankPartURL := string(FICOBankJSONURL[22 : len(FICOBankJSONURL)-3])
		FICOBankBasicURL += FICOBankPartURL
		engine.Visit(FICOBankBasicURL, crawlers.News)
	})

	engine.OnHTML("p[align=\"center\"] > strong", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	engine.OnHTML("p[align=\"center\"] ~ *", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
