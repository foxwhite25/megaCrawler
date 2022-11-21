package ewc

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("ewc", "东西方中心", "https://www.eastwestcenter.org/")
	w.SetStartingUrls([]string{"https://www.eastwestcenter.org/publications",
		"https://www.eastwestcenter.org/news",
		"https://www.eastwestcenter.org/staff/experts"})

	//index
	w.OnHTML("body > div.dialog-off-canvas-main-canvas > div > main > div:nth-child(2) > article > div > div > div > div > div > div.layout__region.layout__region--right.views > div > div > div > div > nav > ul > li > a",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.Index)
		})

	//访问新闻
	w.OnHTML("a.teaser__link", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	//报告
	w.OnHTML("a.link--button-primary.fullpage__button", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.PageType = Crawler.Report
			ctx.File = append(ctx.File, element.Attr("href"))
		}
	})

	//专家
	w.OnHTML("a.link--button-primary.fullpage__button", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), ".org") {
			ctx.PageType = Crawler.Expert
		}
	})

	//获取标题,姓名
	w.OnHTML("body > div.dialog-off-canvas-main-canvas > div > main > div:nth-child(2) > article > div > div:nth-child(1) > div > div:nth-child(2) > div",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			if ctx.PageType == Crawler.Expert {
				ctx.Name = element.Text
			} else if (ctx.PageType == Crawler.Report) || (ctx.PageType == Crawler.News) {
				ctx.Title = element.Text
			}
		})

	//获取新闻分类,专家领域
	w.OnHTML("body > div.dialog-off-canvas-main-canvas > div > main > div:nth-child(2) > article > div > div:nth-child(1) > div > div:nth-child(4) > div> a",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			if ctx.PageType == Crawler.News {
				ctx.CategoryText = element.Text
			} else if ctx.PageType == Crawler.Expert {
				ctx.Area += element.Text + " "
			}
		})

	//获取正文,描述
	w.OnHTML("body > div.dialog-off-canvas-main-canvas > div > main > div:nth-child(2) > article > div > div:nth-child(3) > div.layout__region.layout__region--left > div > div > p",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			if ctx.PageType == Crawler.Expert {
				ctx.Description = element.Text
			} else if ctx.PageType == Crawler.News || ctx.PageType == Crawler.Report {
				ctx.Content = element.Text
			}
		})

	//专家头衔
	w.OnHTML("body > div.dialog-off-canvas-main-canvas > div > main > div:nth-child(2) > article > div > div:nth-child(1) > div > div.fullpage__delimiter-wrapper > div",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Title = element.Text
		})
}
