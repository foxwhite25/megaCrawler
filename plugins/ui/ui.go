package ui

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func cutToList(inputStr string) []string {
	nameList := strings.Split(inputStr, "&")
	for index, value := range nameList {
		nameList[index] = strings.TrimSpace(value)
	}

	return nameList
}

func init() {
	w := Crawler.Register("ui", "瑞典国际事务研究所", "https://www.ui.se/english/")
	w.SetStartingUrls([]string{"https://www.ui.se/english/about/staff/", "https://www.ui.se/english/publications/ui-publications/"})

	//访问专家
	w.OnHTML("#main-search-container > div > div > div.block-bd > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})

	//专家姓名
	w.OnHTML("h1.name", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Name = element.Text
	})

	//专家头衔
	w.OnHTML("div.desc", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//专家介绍
	w.OnHTML("layout-size4of5 layout-padding", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = element.Text
	})

	//获取报告信息
	w.OnHTML("block noborder publication-item", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		subCtx := ctx.CreateSubContext()
		subCtx.PageType = Crawler.Report
		subCtx.File = append(subCtx.File, element.ChildAttr(" div> div > div.section__body > div > div > div > div > a", "href"))
		subCtx.Title = element.ChildText("div.section__body > div > div> div > div > a > h2")
		subCtx.Authors = cutToList(element.ChildText("turkos"))
		subCtx.PublicationTime = element.ChildText(" div > div.section__body > div > div> div > div > a > p")
	})

}
