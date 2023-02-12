package theglobalobservatory

import (
	"github.com/gocolly/colly/v2"
	_ "github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("theglobalobservatory", "国际和平研究所", "https://theglobalobservatory.org/")

	w.SetStartingUrls([]string{"https://theglobalobservatory.org/tag/arab-spring/",
		"https://theglobalobservatory.org/tag/central-asia/",
		"https://theglobalobservatory.org/tag/climate-change/",
		"https://theglobalobservatory.org/tag/conflict/",
		"https://theglobalobservatory.org/tag/democracy/",
		"https://theglobalobservatory.org/tag/development/",
		"https://theglobalobservatory.org/tag/elections/",
		"https://theglobalobservatory.org/tag/fragile-states/",
		"https://theglobalobservatory.org/tag/health-and-security/",
		"https://theglobalobservatory.org/tag/humanitarian/",
		"https://theglobalobservatory.org/tag/justice/",
		"https://theglobalobservatory.org/tag/mali/",
		"https://theglobalobservatory.org/tag/mass-protest/",
		"https://theglobalobservatory.org/tag/peace-and-security/",
		"https://theglobalobservatory.org/tag/peace-processes/",
		"https://theglobalobservatory.org/tag/peacebuilding/",
		"https://theglobalobservatory.org/tag/peacekeeping/",
		"https://theglobalobservatory.org/tag/rebel-groups/",
		"https://theglobalobservatory.org/tag/resources/",
		"https://theglobalobservatory.org/tag/rule-of-law/",
		"https://theglobalobservatory.org/tag/somalia/",
		"https://theglobalobservatory.org/tag/southeast-asia/",
		"https://theglobalobservatory.org/tag/statebuilding/",
		"https://theglobalobservatory.org/tag/technology/",
		"https://theglobalobservatory.org/tag/terrorism/",
		"https://theglobalobservatory.org/tag/united-nations/",
		"https://theglobalobservatory.org/tag/women-peace-and-security/",
		"https://theglobalobservatory.org/tag/africa/",
		"https://theglobalobservatory.org/tag/americas/",
		"https://theglobalobservatory.org/tag/asia/",
		"https://theglobalobservatory.org/tag/europe/",
		"https://theglobalobservatory.org/tag/middle-east/",
		"https://theglobalobservatory.org/series/",
		"https://theglobalobservatory.org/category/features/"})
	// 从翻页器获取链接并访问
	w.OnHTML("li.next>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 从index访问新闻
	w.OnHTML(".entry-title>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	// report.title
	w.OnHTML("h1.entry-title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})
	//report.publish time
	w.OnHTML(".main>article>header.entry-header>div>div.entry-meta>span.date>a>time.entry-date", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})
	// report .content
	w.OnHTML("div.entry-content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
	// report.author
	w.OnHTML(".main>article>header > div > div > span.byline.NOTauthor.NOTvcard", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

}
