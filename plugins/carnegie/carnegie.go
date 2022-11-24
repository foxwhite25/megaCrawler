package carnegie

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func cutToList(inputStr string) []string {
	nameList := strings.Split(inputStr, ",")
	for index, value := range nameList {
		nameList[index] = strings.TrimSpace(value)
	}

	return nameList
}
func init() {
	w := Crawler.Register("carnegie", "卡内基欧洲中心", "https://carnegieeurope.eu/?lang=en")
	w.SetStartingUrls([]string{"https://carnegieeurope.eu/publications/?lang=en", "https://carnegieeurope.eu/experts/?lang=en"})

	//访问专家
	w.OnHTML("li > div > div > h4 > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})

	//专家姓名
	w.OnHTML(" div.gutter-right.tablet-zero.divider.clearfix > div > div > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Name = element.Text
	})

	//专家头衔
	w.OnHTML("div.meta.component.uppercase.roman-normal-bold", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//专家介绍
	w.OnHTML("#bio-panel > div > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description += element.Text
	})

	//专家邮箱
	w.OnHTML("div.col.col-30.tablet-zero > div > div:nth-child(1) > div > div > ul > li:nth-child(1) > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Email = element.Attr("href")
	})

	//专家领域
	w.OnHTML("#expertCenterRegionTags > ul > li> a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Area += element.Text + " "
	})
	//专家
	w.OnHTML("#expertCenterIssueTags > ul > li> a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Keywords = append(ctx.Keywords, element.Text)
	})

	//index
	w.OnHTML("a.page-link", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//访问报告
	w.OnHTML(" div > div.col.col-70.zone-1 > div > ul > li> h4 > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	//获取标题
	w.OnHTML(" div > div.container-headline.foreground > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//获取作者
	w.OnHTML("post-author", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = cutToList(element.Text)
	})

	//获取时间
	w.OnHTML("div.container-headline.foreground > div > div.post-date.col.col-25 > ul > li:nth-child(1)", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	//pdf
	w.OnHTML("a.analytics-download", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
	})

	//获取正文
	w.OnHTML(" div.zone-1 > div > div.article-body > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content += element.Text
	})

	//标签
	w.OnHTML(" div.zone-1 > div > div.section > div > ul > li> a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Tags = append(ctx.Tags, element.Text)
	})

}
