package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1447", "印度商业线报", "https://www.thehindubusinessline.com/")

	engine.SetStartingURLs([]string{"https://www.thehindubusinessline.com/sitemap/archive.xml"})

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

	//备注：特殊的URL，按天算，程序预运行需要索引约365*14个sitemap
	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, ".xml") {
			engine.Visit(element.Text, crawlers.Index)
		}
		if strings.Contains(element.Text, ".ece") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML("picture  > img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = []string{element.Attr("src")}
	})

	//这个网站未登录有查看限制，但是在开发者工具栏依旧能看到完整的文章（hhh）
	engine.OnHTML(".col-md-12.bl-news-section-split > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
