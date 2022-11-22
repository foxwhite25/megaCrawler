package nbr

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("nbr", "国家亚洲研究局海洋意识项目", "https://map.nbr.org/")
	w.SetStartingUrls([]string{"https://map.nbr.org/interactive/incident-timeline/", "https://map.nbr.org/category/analysis/", "https://map.nbr.org/category/international-expert-panel/"})

	//index
	w.OnHTML("a.page-numbers.next", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//访问报告
	w.OnHTML("article.category-analysis > div > div.elementor-post__text > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	//访问专家
	w.OnHTML("article.category-international-expert-panel > div > div > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})

	//获取标题,姓名
	w.OnHTML("h1.elementor-heading-title.elementor-size-default", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.Report {
			ctx.Title = element.Text
		} else if ctx.PageType == Crawler.Expert {
			ctx.Name = element.Text
		}
	})

	//获取作者
	w.OnHTML(" div > div > section.elementor-section.elementor-top-section.elementor-element.elementor-element-33a6437.elementor-section-full_width.elementor-section-height-min-height.elementor-section-items-stretch.elementor-reverse-tablet.elementor-section-height-default > div > div > div > div > div > div > div > ul > li.elementor-icon-list-item.elementor-repeater-item-2381197.elementor-inline-item > span",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Authors = append(ctx.Authors, element.Text)
		})

	//获取时间
	w.OnHTML(" div > div > section.elementor-section.elementor-top-section.elementor-element.elementor-element-33a6437.elementor-section-full_width.elementor-section-height-min-height.elementor-section-items-stretch.elementor-reverse-tablet.elementor-section-height-default > div > div > div > div > div > div > div > ul > li.elementor-icon-list-item.elementor-repeater-item-13c85cb.elementor-inline-item > span",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.PublicationTime = element.Text
		})

	//获取正文,描述
	w.OnHTML("div > div > section> div > div > div > div > div > div > div > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.Report {
			ctx.Content += element.Text
		} else if ctx.PageType == Crawler.Expert {
			ctx.Description += element.Text
		}
	})

	//新闻
	w.OnHTML("#tl-imckn > div.tl-slider-container-mask > div > div > div", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		subCtx := ctx.CreateSubContext()
		subCtx.PageType = Crawler.News
		subCtx.Title = element.ChildText("div > div.tl-text-headline-container > h2")
		subCtx.Content = element.ChildText(" div > div.tl-text-content > p")
	})
}
