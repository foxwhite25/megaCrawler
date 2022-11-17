package inpr

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("inpr", "台湾国策研究院", "http://inpr.org.tw/")

	w.SetStartingUrls([]string{
		"http://inpr.org.tw/m/412-1728-111.php",
		"http://inpr.org.tw/m/412-1728-112.php",
		"http://inpr.org.tw/m/412-1728-113.php",
		"http://inpr.org.tw/m/412-1728-114.php",
	})

	// 从 Index 访问 Report
	w.OnHTML(".d-img>a",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.Report)
		})

	// 添加 Title 到 ctx
	w.OnHTML(".mpgtitle>h3",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Title = strings.TrimSpace(element.Text)
		})

	// 添加 Content 到 ctx
	w.OnHTML("div.mpgdetail",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Content = strings.TrimSpace(element.Text)
		})

}
