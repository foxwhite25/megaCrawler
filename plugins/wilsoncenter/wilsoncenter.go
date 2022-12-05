package wilsoncenter

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("wilsoncenter", "伍德罗·威尔逊国际学者中心", "https://www.wilsoncenter.org/")
	w.SetStartingUrls([]string{"https://www.wilsoncenter.org/sitemap.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Text, "?page=") {
			w.Visit(element.Text, Crawler.Index)
		} else {
			if strings.Contains(element.Text, "/person/") {
				w.Visit(element.Text, Crawler.Expert)
			} else if strings.Contains(element.Text, "/article/") || strings.Contains(element.Text, "/blog-post/") {
				w.Visit(element.Text, Crawler.News)
			}
		}
	})

	//专家姓名
	w.OnHTML("expert-hero-name h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Name = element.Text
	})

	//专家头衔
	w.OnHTML("expert-hero-position", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//新闻标题
	w.OnHTML("insight-detail-hero-title h1 -serif", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//作者
	w.OnHTML("insight-detail-hero-author-byline-link-text", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	//时间
	w.OnHTML("insight-detail-hero-author-byline-text -date", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	//专家描述,新闻正文
	w.OnHTML("text-block-inner", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.Expert {
			ctx.Description = element.Text
		} else if ctx.PageType == Crawler.News {
			ctx.Content = element.Text
		}
	})
}
