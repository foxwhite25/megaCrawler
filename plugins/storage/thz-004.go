package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("thz-004", "国家人口及家庭发展局", "https://www.lppkn.gov.my/lppkngateway/frontend/web/index.php?r=portal%2Findex")

	engine.SetStartingURLs([]string{"https://www.lppkn.gov.my/lppkngateway/frontend/web/index.php?r=portal%2Flist_news&menu=46&id=SkUrdDJmNWpuZDVqbWZuMDRJSzdOQT09&y=2025"})

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

	engine.OnHTML(".col-md-12.berita-content a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".content > p > span > span > span", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find(".content > p > span > span > span > img").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

	engine.OnHTML(".next > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
