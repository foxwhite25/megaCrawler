package diis

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("diis", "国际问题研究所", "https://www.diis.dk/en")
	w.SetStartingUrls([]string{"https://www.diis.dk/en", "https://www.diis.dk/en/experts"})

	//index
	w.OnHTML("ul > li.menu-item.menu-item--expanded.ot-expanded > ul > li> a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})
	w.OnHTML(" div.views-element-container > div > nav > ul > li> a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})
	w.OnHTML("#quicktabs-tabpage-qt_terms-0 > div:nth-child(2) > div > div > nav > ul > li> a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//访问experts
	w.OnHTML("div.field.node-title > h2 > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})

	//获取姓名
	w.OnHTML("body > div > main > div > div > article > div.inner > div.beta > div.group-content-top > div.field.node-title > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Name = element.Text
	})

	//获取人物头衔
	w.OnHTML("body > div > main > div > div > article > div.inner > div.beta > div.group-content-top > div.field.field-category", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//获取人物领域
	w.OnHTML("body > div > main > div > div > article > div.inner > div.beta > div.group-content-top > div.field.field-research-area", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Area = element.Text
	})

	//获取人物介绍
	w.OnHTML("div.content > div > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = element.Text
	})
	w.OnHTML("div.content > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = element.Text
	})

	//获取人物邮箱
	w.OnHTML("body > div > main > div > div > article > div.inner > div.alpha > div.group-contact-wrapper > div.field.field-email > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Email = element.Text
	})

	//访问新闻
	w.OnHTML("#quicktabs-tabpage-qt_terms-0 > div:nth-child(2) > div > div > div > div > ul > li> div > div > div.beta > div.field.node-title > h2 > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	//获取标题
	w.OnHTML("body > div > main > div > div > article > div.inner > div.alpha > div.group-content-top > div.field.node-title > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//获取文章分类
	w.OnHTML("body > div > main > div > div > article > div.inner > div.alpha > div.group-content-top > div.slugline > div.field.octo-slug > div", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.CategoryText = element.Text
	})

	//获取作者
	w.OnHTML("body > div > main > div > div > article > div.inner > div.alpha > div.group-content-top > div.field.field-byline", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	//获取时间
	w.OnHTML("body > div > main > div > div > article > div.inner > div.alpha > div.group-content-top > div.slugline > div.field.field-date > time", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	//获取正文
	w.OnHTML("body > div > main > div > div > article > div.inner > div.alpha > div.field.field-content > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

	//获取标签
	w.OnHTML(" div.inner > div.alpha > div.field.field-topic > ul > li > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Tags = append(ctx.Tags, element.Text)
	})
}
