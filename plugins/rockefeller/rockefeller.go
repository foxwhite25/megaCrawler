package rockefeller

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("rockefeller", "洛克菲勒基金会", "https://www.rockefellerfoundation.org")
	w.SetStartingUrls([]string{"https://www.rockefellerfoundation.org/commitment/food/",
		"https://www.rockefellerfoundation.org/commitment/health/",
		"https://www.rockefellerfoundation.org/commitment/clean-energy/",
		"https://www.rockefellerfoundation.org/commitment/economic-equity/",
		"https://www.rockefellerfoundation.org/commitment/innovation/"})

	//访问专家
	w.OnHTML("div > div.offset-full-1 > ul > li> div > div > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})

	w.OnHTML(" ul > span > span > li > div.pic > div > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})

	//专家姓名,新闻标题
	w.OnHTML("#hero > div.container > div > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.Expert {
			ctx.Name = element.Text
		} else if ctx.PageType == Crawler.News {
			ctx.Title = element.Text
		}
	})

	//专家领域
	w.OnHTML("#hero > div.container > div > span > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Area = element.Text
	})

	//专家头衔
	w.OnHTML("#hero > div.container > div > div > span.job_title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//专家描述,新闻正文
	w.OnHTML(".article-body", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.Expert {
			ctx.Description = element.Text
		} else if ctx.PageType == Crawler.News {
			ctx.Content = element.Text
		}
	})

	//访问index
	w.OnHTML("#update-loadmore > ul > li > button", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//访问新闻
	w.OnHTML("section > div > div.authored_content > div > ul> li > article > span.info > span.title > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	//获取作者
	w.OnHTML("#hero > div.container > div > ul > li > a > strong", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	//新闻时间
	w.OnHTML("#hero > div.container > div > div > span", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	//新闻tag
	w.OnHTML("#tags-content > div > div > ul > li > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Tags = append(ctx.Tags, element.Text)
	})

	// Twitter
	w.OnHTML("li.twitter > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.TwitterId = element.Text
	})

	// LinkedIn
	w.OnHTML("li.linkedin > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.TwitterId = element.Text
	})
}
