package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("SR0032", "司法部", "https://www.moj.gov.vn/Pages/home.aspx")

	engine.SetStartingURLs([]string{
		"https://www.moj.gov.vn/en/Pages/Ministry-of-Justices-Activities.aspx",
		"https://www.moj.gov.vn/en/Pages/Coordination-of-International-Legal-Cooperation-Activit.aspx",
		"https://www.moj.gov.vn/en/Pages/Activities-of-public-administrative-and-justice-reform.aspx"})

	extractorConfig := extractors.Config{
		Author:       true,  //无作者
		Image:        true,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML("p > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.News)
	})

	engine.OnHTML("div.content-news", func(element *colly.HTMLElement, ctx *crawlers.Context) {
    	text := element.Text
    	text = strings.ReplaceAll(text, "\n", "")
    	text = strings.ReplaceAll(text, "\r", "")
    	text = strings.TrimSpace(text)
    	ctx.Content += text
	})

	engine.OnHTML(` div.content-news > div > span`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})


	engine.OnHTML("div > a:nth-child(7)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.Index)
	})
}
