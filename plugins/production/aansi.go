package production

import (
	"strconv"
	"strings"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("aansi", "American-Armenian National Security Institute",
		"https://aansi.org/")

	w.SetStartingURLs([]string{
		"https://aansi.org/tag/future/",
		"https://aansi.org/tag/investment/",
	})

	// 访问 News 从 Index
	w.OnHTML(`div > div.post_text > div > h2 > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	// 获取 Title
	w.OnHTML(`div[class="blog_single blog_holder"] h2.entry_title`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`span[class="date entry_date updated"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 CategoryText
	w.OnHTML(`a[rel="category tag"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})

	// 获取 Authors
	w.OnHTML(`a.post_author_link`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 获取 CommentCount
	w.OnHTML(`a.post_comments`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		// 裁出数字的字符串并将其转换为 int 类型
		var str = strings.Replace(element.Text, "Comments", "", 1)
		str = strings.TrimSpace(str)
		num, _ := strconv.Atoi(str)
		ctx.CommentCount = num
	})

	// 获取 LikeCount
	w.OnHTML(`.qode-like`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		// 裁出数字的字符串并将其转换为 int 类型
		var str = strings.Replace(element.Text, "Likes", "", 1)
		str = strings.TrimSpace(str)
		num, _ := strconv.Atoi(str)
		ctx.LikeCount = num
	})

	// 获取 Content
	w.OnHTML(`div.elementor-widget-container`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 获取 Tags
	w.OnHTML(`a[rel="tag"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})
}
