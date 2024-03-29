package production

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("project2049", "2049计划研究所", "https://project2049.net/")

	w.SetStartingURLs([]string{"https://project2049.net/category/blog/",
		"https://project2049.net/category/publications/policy-briefs/",
		"https://project2049.net/category/publications/occasional-papers/",
		"https://project2049.net/category/publications/in-the-news/",
		"https://project2049.net/events",
		"https://project2049.net/people/",
	})

	// 从翻页器获取链接并访问
	w.OnHTML("div.pages > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 从index访问新闻
	w.OnHTML("a.post-more", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	// new.title
	w.OnHTML("h1.entry-title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	// new.publish time
	w.OnHTML("time.entry-date", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})

	// new.author
	w.OnHTML("span.fn > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	// new .content
	w.OnHTML("div.the_content_wrapper", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})

	w.OnHTML(".team", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		subCtx := ctx.CreateSubContext()
		subCtx.PageType = crawlers.Expert
		subCtx.Name = element.ChildText(".title")
		subCtx.Title = element.ChildText(".desc_wrapper > h4")
		subCtx.Description = element.ChildText(".popup-inner")
	})
}
