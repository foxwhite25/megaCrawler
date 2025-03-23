package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2473", "Alizila", "https://www.alizila.com/")

	engine.SetStartingURLs([]string{
		"https://www.alizila.com/post-sitemap.xml",
		"https://www.alizila.com/post-sitemap2.xml",
		"https://www.alizila.com/post-sitemap3.xml",
	})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        false, // 这个网站的图片链接需要网站完全加载后才加载
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

	engine.OnHTML(`.article-body > p:not(:has(noscript)), .article-body > div > span`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
