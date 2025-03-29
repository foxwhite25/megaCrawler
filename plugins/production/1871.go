package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1871", "Liberal Values", "http://liberalvaluesblog.com/")

	engine.SetStartingURLs([]string{"http://liberalvaluesblog.com/"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true, //大部分博客的图片有文字内容，建议不要关闭
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML(".post.type-post > h2 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.News)
	})

	// 大部分内容发布在FaceBook上,该部分无法采集。
	// engine.OnHTML("._5pbx.userContent > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
	// 	ctx.Content += element.Text
	// })

	engine.OnHTML(".content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML(".alignleft > a:nth-child(1)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.Index)
	})
}
