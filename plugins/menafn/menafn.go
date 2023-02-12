package menafn

import (
	"github.com/gocolly/colly/v2"
	_ "github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("menafn", "中东北非金融网络", "https://menafn.com/")

	w.SetStartingUrls([]string{"https://menafn.com/qn_news_sections.aspx?src=s",
		"https://menafn.com/qn_news_sections.aspx?src=c&country=Americas",
		"https://menafn.com/qn_news_sections.aspx?src=c&country=Europe",
		"https://menafn.com/qn_news_sections.aspx?src=c&country=ArabWorld",
		"https://menafn.com/qn_news_sections.aspx?src=c&country=Asia",
		"https://menafn.com/qn_news_sections.aspx?src=c&country=Africa",
		"https://menafn.com/pr/index.aspx",
		"https://menafn.com/qn_calendar.aspx?Src1=header",
		"https://menafn.com/research/rc_research.aspx",
		"https://menafn.com/qn_country.aspx?country=SA&src=header",
		"https://menafn.com/qn_country.aspx?country=AE&src=header",
		"https://menafn.com/qn_country.aspx?country=BH&src=header",
		"https://menafn.com/qn_country.aspx?country=QA&src=header",
		"https://menafn.com/qn_country.aspx?country=KW&src=header",
		"https://menafn.com/qn_country.aspx?country=JO&src=header",
		"https://menafn.com/qn_country.aspx?country=OM&src=header",
		"https://menafn.com/qn_country.aspx?country=EG&src=header",
		"https://menafn.com/qn_country.aspx?country=LB&src=header",
		"https://menafn.com/qn_country.aspx?country=IQ&src=header",
		"https://menafn.com/qn_country.aspx?country=PS&src=header",
		"https://menafn.com/qn_country.aspx?country=SY&src=header",
		"https://menafn.com/qn_country.aspx?country=TN&src=header",
		"https://menafn.com/qn_country.aspx?country=DZ&src=header",
		"https://menafn.com/qn_country.aspx?country=MA&src=header",
		"https://menafn.com/qn_country.aspx?country=YE&src=header",
		"https://menafn.com/events/default.aspx"})

	// 从index访问新闻
	w.OnHTML("ul.popular-news-list>li>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	w.OnHTML("div.img-wrapper>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	w.OnHTML("div.img-wrapper>div>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	// report.title
	w.OnHTML("h1.main-entry-title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	// report .content
	w.OnHTML("#ContentPlaceHolder1_div_story > div:nth-child(5)", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
	//report.publish time
	w.OnHTML("span.entry-date", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})
}
