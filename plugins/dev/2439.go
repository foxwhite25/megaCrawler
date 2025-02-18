package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2439", "Association, Organization and Government News", "https://www.thomasnet.com/")

	engine.SetStartingURLs([]string{
		"https://www.thomasnet.com/insights/browse/business-industry/",
		"https://www.thomasnet.com/insights/browse/career/",
		"https://www.thomasnet.com/insights/browse/engineering-design/",
		"https://www.thomasnet.com/insights/browse/industry-trends/",
		"https://www.thomasnet.com/insights/browse/manufacturing-innovation/",
		"https://www.thomasnet.com/insights/browse/sales-marketing/",
		"https://www.thomasnet.com/insights/browse/supply-chain/",
	})

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

	engine.OnHTML("div.css-1oteowz.ejfvg3z0 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.News)
	})

	engine.OnHTML(`div.css-1l4w6pd.evgnxt60 > div > a[aria-label="Next Page"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.Index)
	})

	engine.OnHTML("div.css-9mso2o > img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		imageUrl := element.Request.AbsoluteURL(element.Attr("src"))
		ctx.Image = []string{imageUrl}
	})

	engine.OnHTML("section.css-1a2qmc3 > p, section.css-1a2qmc3 > h2, section.css-1a2qmc3 > ul",
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			ctx.Content += element.Text
		})
}
