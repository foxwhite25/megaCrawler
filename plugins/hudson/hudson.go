package hudson

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("hudson", "哈德森研究所", "https://www.hudson.org/")
	w.SetStartingUrls([]string{"https://www.hudson.org/experts",
		"https://www.hudson.org/search?article-type-ajax=258",
		"https://www.hudson.org/search?article-type-ajax=259"})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("p:nth-child(1) > a.button.button--secondary", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = Crawler.Report
		}
	})

	//index
	w.OnHTML(".pager__item.pager__item--next > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//访问文章
	w.OnHTML("c-horizontal-card__title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	//文章类型
	w.OnHTML("field name", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.CategoryText = element.Text
	})

	//时间
	w.OnHTML("#block-mainpagecontent > article > div > div.hud-layout__short-form-article-hero.hero.hero--article-short.hero--half-image-right > div.hero__inner > div > div.layout__region.layout__region--content-top.hero__content-top > div.block.block-layout-builder.block-field-blocknodeshort-form-articlefield-date > div > time",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.PublicationTime = element.Text
		})

	//标题
	w.OnHTML("block block-layout-builder block-field-blocknodeshort-form-articletitle", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//作者
	w.OnHTML(".hero__experts > div > div > div.expert-author--content > div > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	//正文
	w.OnHTML(".field.body > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content += element.Text
	})

	//标签
	w.OnHTML(".block-field-blocknodeshort-form-articlefield-topics > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Tags = append(ctx.Tags, element.Text)
	})

	//访问专家
	w.OnHTML(".expert-card__cta > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})

	//专家姓名
	w.OnHTML("block block-layout-builder block-field-blocknodeexpert-biotitle", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Name = element.Text
	})

	//专家头衔
	w.OnHTML("field field-eb-position", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//专家描述
	w.OnHTML("field field-eb-biography", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = element.Text
	})
}
