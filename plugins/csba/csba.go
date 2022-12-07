package csba

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("csba", "战略与预算评估中心", "https://csbaonline.org/")
	w.SetStartingUrls([]string{"https://csbaonline.org/about/events", "https://csbaonline.org/about/people/staff",
		"https://csbaonline.org/research/publications?keywords=&categories%5B%5D=132&categories%5B%5D=131&categories%5B%5D=133&date_from=&date_to="})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML(".article-resource-title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = Crawler.Report
		}
	})
	//index
	w.OnHTML("a.page-link-next", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//访问专家
	w.OnHTML("a.profile", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})

	//访问文章
	w.OnHTML("h2.article-title > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	//专家姓名
	w.OnHTML("article-title people-entry-name", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Name = element.Text
	})

	//专家头衔
	w.OnHTML("people-entry-position", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//专家介绍,文章正文
	w.OnHTML("article-content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.Expert {
			ctx.Description = element.Text
		} else if ctx.PageType == Crawler.News || ctx.PageType == Crawler.Report {
			ctx.Content += element.Text
		}
	})

	//专家标签
	w.OnHTML(".tag-pill", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Tags = append(ctx.Tags, element.Text)
	})

	//文章标题
	w.OnHTML("h1.article-title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//文章时间
	w.OnHTML(" div > div.article-meta > time", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

}
