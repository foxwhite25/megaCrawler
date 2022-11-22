package ecfr

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("ecfr", "欧洲对外关系委员会", "https://ecfr.eu/")
	w.SetStartingUrls([]string{"https://ecfr.eu/search/", "https://ecfr.eu/profiles/"})

	//index
	w.OnHTML("li.cat-item > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	w.OnHTML("a.next.btn", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//访问报告
	w.OnHTML("h2.post-title > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	//获取pdf
	w.OnHTML("span.btn-link-label", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
		}
	})

	//标题
	w.OnHTML("h1.post-title.article-h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//获取作者
	w.OnHTML("div.bio-name > a.card-main-link", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.Report {
			ctx.Authors = append(ctx.Authors, element.Text)
		}
	})

	//分类
	w.OnHTML("li > a.a-subtle", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.CategoryText = element.Text
	})

	//正文
	w.OnHTML("div.text-body", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.Report {
			ctx.Content = element.Text
		} else if ctx.PageType == Crawler.Expert {
			ctx.Description = element.Text
		}
	})

	//访问专家
	w.OnHTML("div.bio-name > a.card-main-link", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})

	//专家姓名
	w.OnHTML("h1.profile-name", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Name = element.Text
	})

	//专家头衔
	w.OnHTML("div.profile-jobtitle", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//获取邮箱
	w.OnHTML("div.sidebar > div > ul > li:nth-child(1) > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Email = element.Text
	})

}
