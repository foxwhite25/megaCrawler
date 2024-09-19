package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1414", "越通社", "https://vnanet.vn/")

	engine.SetStartingURLs([]string{ // 首页下拉列表中所有新闻页
		"https://vnanet.vn/vi/tin-tuc/chinh-tri-11/",
		"https://vnanet.vn/vi/tin-tuc/kinh-te-4/",
		"https://vnanet.vn/vi/tin-tuc/an-ninh-quoc-phong-16/",
		"https://vnanet.vn/vi/tin-tuc/xa-hoi-14/",
		"https://vnanet.vn/vi/tin-tuc/toi-pham-2/",
		"https://vnanet.vn/vi/tin-tuc/nghe-thuat-van-hoa-va-giai-tri-1/",
		"https://vnanet.vn/vi/tin-tuc/giao-duc-5/",
		"https://vnanet.vn/vi/tin-tuc/khoa-hoc-cong-nghe-13/",
		"https://vnanet.vn/vi/tin-tuc/the-thao-15/",
		"https://vnanet.vn/vi/tin-tuc/moi-quan-tam-con-nguoi-8/",
		"https://vnanet.vn/vi/tin-tuc/moi-truong-6/",
		"https://vnanet.vn/vi/tin-tuc/suc-khoe-7/",
		"https://vnanet.vn/vi/tin-tuc/thoi-tiet-17/",
	})

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

	engine.OnHTML(".grp-list-news-2 > ul >li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".divContent_ListNewsBody_pager > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML("div.article__body.zce-content-body.cms-body > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
