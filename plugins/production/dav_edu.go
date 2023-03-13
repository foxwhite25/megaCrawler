package production

import (
	"strings"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("dav_edu", "外交学院", "https://www.dav.edu.vn/")

	w.SetStartingURLs([]string{
		"https://www.dav.edu.vn/su-kien-hoi-thao-toa-dam/?trang=1",
		"https://www.dav.edu.vn/gioi-thieu-chung-nghien-cu/",
		"https://www.dav.edu.vn/an-pham-nghien-cuu/?trang=1",
	})

	// 访问下一页 Index
	w.OnHTML(`[class="page-item active"] > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		indexURL := crawlers.GetNextIndexURL(ctx.URL, element.Text, "trang")
		w.Visit(indexURL, crawlers.Index)
	})

	// 访问 Report 从 Index
	w.OnHTML(`.row .story__title > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 获取 Title
	w.OnHTML(`.detail__title`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Description
	w.OnHTML(`.detail__summary`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`.detail__meta`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 CategoryText
	w.OnHTML(`li[class="breadcrumb-item active"] > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})

	// 获取 Content
	w.OnHTML(`html .detail__content`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})
}
