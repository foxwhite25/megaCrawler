package heritage

import (
	"github.com/gocolly/colly/v2"
	_ "github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	_ "megaCrawler/megaCrawler"
	_ "regexp"
	"strings"
	_ "strings"
)

func init() {
	w := megaCrawler.Register("heritage", "美国传统基金会", "https://www.heritage.org/")
	w.SetStartingUrls([]string{"https://www.heritage.org/about-heritage/staff/experts", "https://www.heritage.org/"})
	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告 1
	w.OnHTML("a.button", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = megaCrawler.Report
		}
	})

	// 从翻页器获取链接并访问 1
	w.OnHTML(".button-more", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Index)
	})

	// 尝试访问作者并添加到ctx 1
	w.OnHTML(".person-list-small__image-wrapper", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if element.ChildAttr("a", "href") != "" {
			w.Visit(element.ChildAttr("a", "href"), megaCrawler.Expert)
		}

	})

	w.OnHTML(".js-hover-target > div", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	// 访问新闻 1
	w.OnHTML("article > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.News)
	})
	// 访问新闻 1
	w.OnHTML("a[hreflang=\"en\"]", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Index)
	})

	//访问新闻 1
	w.OnHTML(".result-card:not(.result-card__video):not(._has-video)", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.ChildAttr(".result-card__title", "href"), megaCrawler.News)
	})

	// 添加正文到ctx 1
	w.OnHTML(".person-list-small__panelist", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Content = element.Text
	})

	// 人物头衔到ctx 1
	w.OnHTML(".person-list-small__title", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Title = megaCrawler.StandardizeSpaces(element.Text)
	})

	// 人物描述到ctx 1
	w.OnHTML(".expert-bio-body__copy > p", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Description = element.Text
	})

	// 专家领域到ctx 1
	w.OnHTML("font[style=\"vertical-align: inherit;\"]", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Area = element.Text
	})

	//专家图片到ctx
	w.OnHTML("a.expert-bio-card__download-headshot", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Image = []string{element.Attr("href")}
	})

	// 访问新闻
	w.OnHTML("article[role=\"article\"] > div > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.News)
	})

	//new . author_name
	w.OnHTML(".author-card__author-info-wrapper > a > span", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	//new . author_information
	w.OnHTML(".author-card__card-info > p > font", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	//new . content
	w.OnHTML("#block-mainpagecontent > article > div > div > div", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Content = element.Text
	})
	//new . publish_time
	w.OnHTML("div.article-general-info", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.PublicationTime = element.Text
	})
	//new . author_url
	w.OnHTML(" div.commentary__intro-wrapper > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Attr("href"))
	})
	// new. keyword
	w.OnHTML(" div.key-takeaways__takeaway > p", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Keywords = append(ctx.Keywords, element.Text)
	})

	// new .image
	w.OnHTML(" figure.image-with-caption__image-wrapper > img ", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Image = append(ctx.Keywords, element.Attr("srcset"))
	})

}
