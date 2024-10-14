package errors

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

// 反爬
func init() {
	engine := crawlers.Register("1029", "国立大学东亚研究所", "https://research.nus.edu.sg/eai/wp-sitemap.xml")

	engine.SetStartingURLs([]string{"https://research.nus.edu.sg/eai/wp-sitemap-posts-post-1.xml"})

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

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Text, crawlers.News)
	})

	engine.OnHTML("article", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		content := strings.SplitN(element.Text, "\n", 1)
		ctx.Title = content[0]
		ctx.Content = content[1]
	})
}
