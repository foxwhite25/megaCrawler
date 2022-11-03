package cato

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"strings"
)

func init() {
	w := megaCrawler.Register("cato", "卡托研究所", "https://www.cato.org/	")

	w.SetStartingUrls([]string{"https://www.cato.org/constitutional-law",
		"https://www.cato.org/criminal-justice",
		"https://www.cato.org/free-speech-civil-liberties",
		"https://www.cato.org/banking-finance",
		"https://www.cato.org/monetary-policy",
		"https://www.cato.org/regulatory-studies",
		"https://www.cato.org/tax-budget-policy",
		"https://www.cato.org/government-politics",
		"https://www.cato.org/poverty-social-welfare",
		"https://www.cato.org/public-opinion",
		"https://www.cato.org/technology-privacy",
		"https://www.cato.org/defense-foreign-policy",
		"https://www.cato.org/global-freedom",
		"https://www.cato.org/immigration",
		"https://www.cato.org/trade-policy",
		"https://www.cato.org/ukraine",
		"https://www.cato.org/project-jones-act-reform",
		"https://www.cato.org/inflation",
		"https://www.cato.org/covid-19",
		"https://www.cato.org/search/category/commentary",
		"https://www.cato.org/search/category/economic-freedom-world",
		"https://www.cato.org/search/category/multimedia+cato-audio",
		"https://www.cato.org/search/category/multimedia+cato-daily-podcast",
		"https://www.cato.org/search/category/multimedia+power-problems",
		"https://www.cato.org/search/category/outside-articles",
		"https://www.cato.org/search/category/reviews-journals+cato-journal",
		"https://www.cato.org/search/category/reviews-journals+policy-report",
		"https://www.cato.org/search/category/reviews-journals+regulation",
		"https://www.cato.org/search/category/reviews-journals+supreme-court-review",
		"https://www.cato.org/search/category/study+briefing-paper",
		"https://www.cato.org/search/category/study+development-policy-analysis",
		"https://www.cato.org/search/category/study+economic-development-bulletin",
		"https://www.cato.org/search/category/study+foreign-policy-briefing",
		"https://www.cato.org/search/category/study+free-trade-bulletin",
		"https://www.cato.org/search/category/study+policy-analysis",
		"https://www.cato.org/search/category/study+pandemics-policy",
		"https://www.cato.org/search/category/study+tax-budget-bulletin"})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = megaCrawler.Report
		}
	})

	// 从翻页器获取链接并访问
	w.OnHTML("#main > article > div.collection-page__canvases.spacer--standout--children.spacer--standout--top > div:nth-child(1) > div.more-button.d-flex.justify-content-center.mt-5.mt-lg-6 > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Index)
	})
	w.OnHTML("#main > div.search-page.js-search-page > div > div:nth-child(4) > div > div > nav > ul > li > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Index)
	})
	//访问new
	w.OnHTML("#main > div.search-page.js-search-page > div > div:nth-child(4) > div > div > ul > li > article > h5 > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.News)
	})

	//https://www.cato.org/ukraine 需要另一种翻页模式

	// 从翻页器获取链接并访问

	w.OnHTML("#main > div.spacer--standout--top > div > div > div.container > div > div > nav > ul > li > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Index)
	})
	//访问new
	w.OnHTML("#main > div.spacer--standout--top > div > div > div.container > div > div > ul > li > article > h5 > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.News)
	})

	//blog
	// 添加标题到ctx
	w.OnHTML("#main > div.container.spacer--xl-responsive--top.sidebar.sidebar--right.sidebar--active > div.sidebar__content.p-mb-last-child-0 > article > div.blog-page__header.mb-5 > h1 > span",
		func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
			ctx.Title = element.Text
		})
	// report.author
	w.OnHTML("#main > div.container.spacer--xl-responsive--top.sidebar.sidebar--right.sidebar--active > div.sidebar__content.p-mb-last-child-0 > article > div.blog-page__separator.d-flex.align-items-center.mb-5 > div.me-4.authors.fs-xs.fw-medium > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	//report.publish time
	w.OnHTML("#main > div.container.spacer--xl-responsive--top.sidebar.sidebar--right.sidebar--active > div.sidebar__content.p-mb-last-child-0 > article > div.blog-page__header.mb-5 > div", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.PublicationTime = element.Text
	})
	// report .content
	w.OnHTML("#main > div.container.spacer--xl-responsive--top.sidebar.sidebar--right.sidebar--active > div.sidebar__content.p-mb-last-child-0 > article > div.blog-page__content.mb-5.pb-4 > div > div > div:nth-child(1) > div", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Content = element.Text
	})
	// report .link
	w.OnHTML("div.popover-body.p-mb-last-child-0 > div > div > div>a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Link = append(ctx.Link, element.Attr("href"))
	})
	//内含Expert
	w.OnHTML("#main > div.container.spacer--xl-responsive--top.sidebar.sidebar--right.sidebar--active > div.sidebar__content.p-mb-last-child-0 > article > div.blog-page__separator.d-flex.align-items-center.mb-5 > div.me-4.authors.fs-xs.fw-medium > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Expert)
	})
	// Expert .link
	w.OnHTML("#main > div.container-lg > article > div.author-page__name > h1 > span", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	// Expert .name
	w.OnHTML("#main > div.container-lg > article > div.author-page__contact > div > div:nth-child(1) > div.field-large-text__content.fs-xxl--responsive", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Phone = element.Text
	})

	//Expert .information
	w.OnHTML("#main > div.container-lg > article > div.author-page__info.border-bottom.mb-5.mb-xl-6.pb-4 > div > div.meta.meta--default.text-uppercase.p-mb-last-child-0 > p > span > span", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)

	})
	w.OnHTML("#main > div.container-lg > article > div.author-page__name > div > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)

	})
	// Expert  .content
	w.OnHTML("#rmjs-2 > div", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Content = element.Text
	})
	// Expert  .link
	w.OnHTML("#main > div.container-lg > article > div.author-page__info.border-bottom.mb-5.mb-xl-6.pb-4 > div > div.field-large-text > div > div>a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Link = append(ctx.Link, element.Attr("href"))
	})

	//commentary  speeches
	// 添加标题到ctx
	w.OnHTML("#main > article > div:nth-child(2) > header > div.article-title > h1 > span", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Title = element.Text
	})

	// report.description
	w.OnHTML("#main > article > div:nth-child(2) > header > div.blurb.p-mb-last-child-0 > p", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Description = element.Text
	})
	//report.publish time
	w.OnHTML("#main > article > div:nth-child(2) > header > div.meta.meta--default.text-uppercase.p-mb-last-child-0", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.PublicationTime = element.Text
	})
	// report.author
	w.OnHTML("#main > article > div:nth-child(2) > header > div.authors.fs-xs.fw-medium > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	//内含Expert
	w.OnHTML("#main > article > div:nth-child(2) > header > div.authors.fs-xs.fw-medium > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Expert)
	})
	// report  .link
	w.OnHTML("#main > article > div:nth-child(2) > header > div.share-icons__wrapper > div.share-icons.js-share-icons.fs-xl.d-print-none.icon-row.align-items-center>a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Link = append(ctx.Link, element.Attr("href"))
	})

	// report .content
	w.OnHTML("#main > article > div.long-form-canvas.long-form-canvas--default.js-long-form-canvas", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Content = element.Text
	})

	//event
	// report.category -> Events
	w.OnHTML("#main > div.site-section-header.d-print-none.theme.quaternary > div > div > div.collapse__header > a > h1 > div > div.collapse__toggler__text > div.collapse__toggler__text__inner.collapse__toggler__text__expand.js-collapse-text-expand", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.CategoryText = element.Text
	})
	// report.title
	w.OnHTML("#main > article > header > div.event-page__title.event-title > h1 > span > em", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Title = element.Text
	})
	//report.publish time
	w.OnHTML("#main > article > header > div.event-page__title.event-title > h3", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.PublicationTime = element.Text
	})

	// report .content
	w.OnHTML("#main > article > div.event-page__content > div:nth-child(1) > p", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Content = element.Text
	})

	//public filing
	// report.title
	w.OnHTML("#main > article > div:nth-child(2) > header > div.article-title > h1 > span", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Title = element.Text
	})
	// report.description
	w.OnHTML("#main > article > div.long-form__wrapper.container--medium-down > header > div.blurb.p-mb-last-child-0 > p", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Description = element.Text
	})
	//report.publish time
	w.OnHTML("#main > article > div.long-form__wrapper.container--medium-down > header > div.meta.meta--default.text-uppercase.p-mb-last-child-0", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.PublicationTime = element.Text
	})
	// report.author
	w.OnHTML("#main > article > div.long-form__wrapper.container--medium-down > header > div.authors.fs-xs.fw-medium", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	//内含Expert
	w.OnHTML("#main > article > div.long-form__wrapper.container--medium-down > header > div.authors.fs-xs.fw-medium > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Expert)
	})
	// report  .link
	w.OnHTML("#main > article > div.long-form__wrapper.container--medium-down > header > div.share-icons__wrapper > div.share-icons.js-share-icons.fs-xl.d-print-none.icon-row.align-items-center>a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Link = append(ctx.Link, element.Attr("href"))
	})
	// report .content
	w.OnHTML("#main > article > div.long-form-canvas.long-form-canvas--default.js-long-form-canvas > div.long-form-canvas__content.js-long-form-canvas-content > div.long-form > div.position-relative.js-read-depth-tracking-container > div:nth-child(1) > div:nth-child(1) > div > p", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Content = element.Text
	})
	// report  .file
	w.OnHTML("#main > article > div.long-form-canvas.long-form-canvas--default.js-long-form-canvas > div.long-form-canvas__content.js-long-form-canvas-content > div.long-form > div:nth-child(3) > div > div.theme.info.book-promo-block__content > div > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.File = append(ctx.Link, element.Attr("href"))
	})

	//reviews and Journals
	// report.title
	w.OnHTML("#main > article > div.long-form__wrapper.container--medium-down > header > div.article-title > h1 > span", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Title = element.Text
	})
	//report.publish time
	w.OnHTML("#main > article > div.long-form__wrapper.container--medium-down > header > div.meta.meta--default.text-uppercase.p-mb-last-child-0", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.PublicationTime = element.Text
	})
	// report.author
	w.OnHTML("#main > article > div.long-form__wrapper.container--medium-down > header > div.authors.fs-xs.fw-medium > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	//内含Expert
	w.OnHTML("#main > article > div.long-form__wrapper.container--medium-down > header > div.authors.fs-xs.fw-medium > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Expert)
	})
	// report  .link
	w.OnHTML("#main > article > div.long-form__wrapper.container--medium-down > header > div.share-icons__wrapper > div.share-icons.js-share-icons.fs-xl.d-print-none.icon-row.align-items-center", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Link = append(ctx.Link, element.Attr("href"))
	})

	// report .content
	w.OnHTML("#main > article > div.long-form-canvas.long-form-canvas--default.js-long-form-canvas > div.long-form-canvas__content.js-long-form-canvas-content > div.long-form > div.position-relative.js-read-depth-tracking-container > div:nth-child(1)", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Content = element.Text
	})

	//Policy Report   regulation
	// report.category
	w.OnHTML("#main > div.site-section-header.d-print-none.theme.policy-report > div > div > div.collapse__header > a > h1 > div > div.collapse__toggler__text > div.collapse__toggler__text__inner.collapse__toggler__text__expand.js-collapse-text-expand", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})
	// report.title
	w.OnHTML("#main > article > div:nth-child(2) > header > div.article-title > h1 > span", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Title = element.Text
	})

	// report.description
	w.OnHTML("#main > article > div.long-form__wrapper.container--medium-down > header > div.blurb.p-mb-last-child-0 > p", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Description = element.Text
	})
	//report.publish time
	w.OnHTML("#main > article > div.long-form__wrapper.container--medium-down > header > div.meta.meta--default.text-uppercase.p-mb-last-child-0", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.PublicationTime = element.Text
	})
	// report  .link
	w.OnHTML("#main > article > div.long-form__wrapper.container--medium-down > header > div.share-icons__wrapper > div.share-icons.js-share-icons.fs-xl.d-print-none.icon-row.align-items-center>a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Link = append(ctx.Link, element.Attr("href"))
	})
	// report .content
	w.OnHTML("#main > article > div.long-form-canvas.long-form-canvas--default.js-long-form-canvas > div.long-form-canvas__content.js-long-form-canvas-content > div.long-form > div.position-relative.js-read-depth-tracking-container > div:nth-child(1)", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Content = element.Text
	})

	// report  .file
	w.OnHTML("#main > article > div.long-form-canvas.long-form-canvas--default.js-long-form-canvas > div.long-form-canvas__content.js-long-form-canvas-content > div.long-form > div:nth-child(2) > div > div.theme.info.book-promo-block__content > div > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.File = append(ctx.Link, element.Attr("href"))
	})

	//briefing paper  development policyanalysis
	// report.title
	w.OnHTML("#main > article > div.long-form__wrapper.container--medium-down > header > div.article-title > h1 > span", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Title = element.Text
	})
	//report.publish time
	w.OnHTML("#main > article > div.long-form__wrapper.container--medium-down > header > div.meta.meta--default.text-uppercase.p-mb-last-child-0", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.PublicationTime = element.Text
	})
	// report.author
	w.OnHTML("#main > article > div.long-form__wrapper.container--medium-down > header > div.authors.fs-xs.fw-medium > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	//内含Expert
	w.OnHTML("#main > article > div.long-form__wrapper.container--medium-down > header > div.authors.fs-xs.fw-medium > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Expert)
	})
	// report  .link
	w.OnHTML("#main > article > div.long-form__wrapper.container--medium-down > header > div.share-icons__wrapper > div.share-icons.js-share-icons.fs-xl.d-print-none.icon-row.align-items-center>a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Link = append(ctx.Link, element.Attr("href"))
	})
	// report .content
	w.OnHTML("#main > article > div.long-form-canvas.long-form-canvas--default.js-long-form-canvas > div.long-form-canvas__content.js-long-form-canvas-content > div.long-form > div.position-relative.js-read-depth-tracking-container > div:nth-child(1)", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Content = element.Text
	})
	// report  .file
	w.OnHTML("#main > article > div.long-form-canvas.long-form-canvas--default.js-long-form-canvas > div.long-form-canvas__content.js-long-form-canvas-content > div.long-form > div:nth-child(2) > div > div.theme.info.book-promo-block__content > div > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.File = append(ctx.Link, element.Attr("href"))
	})

}
