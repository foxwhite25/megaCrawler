package storage

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1467", "印度中央经济情报局", "https://www.nic.in/")

	engine.SetStartingURLs([]string{"https://www.nic.in/wp-sitemap-posts-news-1.xml"})

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
		engine.Visit(element.Text, crawlers.News)
	})

	engine.OnHTML(".section > div > p.news-date", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML(".news-pragh > p, .news-pragh > div", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML(".col-md-4.img-block > img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = append(ctx.Image, element.Attr("src"))
	})
}
