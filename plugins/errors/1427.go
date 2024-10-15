package errors

import (
	"strings"
	"time"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1427", "BERNAMA", "https://www.bernama.com/en//")

	// 这个网站（没有sitemap）的新闻存放相对有些混乱，但通过下面这个搜索界面能找到大约4000多篇新闻
	engine.SetStartingURLs([]string{"https://www.bernama.com/en/search.php?cat1=all&terms=%20&page=1"})

	engine.SetTimeout(15 * time.Second) // 增加等待服务器回应的时长

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML(".row > div > h6 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	// 处理分页
	engine.OnHTML(".text-center > ul > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		href := element.Attr("href")
		if strings.Contains(href, "?page=") {
			absoluteURL := element.Request.AbsoluteURL(href)
			engine.Visit(absoluteURL, crawlers.Index)
		}
	})

	engine.OnHTML("div.col-12.mt-3.text-dark.text-justify > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
