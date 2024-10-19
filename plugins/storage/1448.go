package storage

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1448", "The Peninsula", "https://thepeninsulaqatar.com/")

	engine.SetStartingURLs([]string{"https://thepeninsulaqatar.com/sitemap.xml"})

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

	// 配置文件中定义Crawl-delay: 10，但它似乎并未实施拦截
	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "/sitemap_articles") {
			engine.Visit(element.Text, crawlers.Index)
		}
		if strings.Contains(element.Text, "/article/") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML(".tag-art > div.photo > img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = []string{"https://thepeninsulaqatar.com/" + element.Attr("src")}
	})

	engine.OnHTML(".con-text > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
