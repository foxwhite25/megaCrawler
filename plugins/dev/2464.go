package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2464", "Across Arizona Patch", "https://patch.com/")

	// 这个网站新闻按地区划分，且存在重复，这里选取一部分
	engine.SetStartingURLs([]string{
		"https://patch.com/new-jersey/southorange?page=2",
		"https://patch.com/new-jersey/summit?page=2",
		"https://patch.com/new-jersey/westfield?page=2",
		"https://patch.com/new-jersey/scotchplains?page=2",
		"https://patch.com/new-jersey/ridgewood?page=2",
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

	engine.OnHTML(".MuiBox-root.css-1yuhvjn > div.MuiBox-root.css-0 > a.MuiTypography-root", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.News)
	})

	engine.OnHTML("div.MuiBox-root.css-zdpt2t > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML("div.MuiBox-root.css-79elbk > img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = []string{element.Attr("src")}
	})

	engine.OnHTML("div.styles_HTMLContent__LDG2k > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
