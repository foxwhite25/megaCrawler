package allafrica

import (
	"github.com/gocolly/colly/v2"
	_ "github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("allafrica", "AllAfrica Web Publications", "https://allafrica.com/")

	w.SetStartingUrls([]string{"https://allafrica.com/algeria/?page=1",
		"https://allafrica.com/angola/?page=1",
		"https://allafrica.com/benin/?page=1",
		"https://allafrica.com/botswana/?page=1",
		"https://allafrica.com/burkinafaso/?page=1",
		"https://allafrica.com/burundi/?page=1",
		"https://allafrica.com/cameroon/?page=1",
		"https://allafrica.com/centralafricanrepublic/?page=1",
		"https://allafrica.com/chad/?page=1",
		"https://allafrica.com/comoros/?page=1",
		"https://allafrica.com/congo_kinshasa/?page=1",
		"https://allafrica.com/cotedivoire/?page=1",
		"https://allafrica.com/djibouti/?page=1",
		"https://allafrica.com/egypt/?page=1",
		"https://allafrica.com/equatorialguinea/?page=1",
		"https://allafrica.com/eritrea/?page=1",
		"https://allafrica.com/ethiopia/?page=1",
		"https://allafrica.com/gabon/?page=1",
		"https://allafrica.com/gambia/?page=1",
		"https://allafrica.com/ghana/?page=1",
		"https://allafrica.com/guinea/?page=1",
		"https://allafrica.com/guineabissau/?page=1",
		"https://allafrica.com/kenya/?page=1",
		"https://allafrica.com/lesotho/?page=1",
		"https://allafrica.com/liberia/?page=1",
		"https://allafrica.com/libya/?page=1",
		"https://allafrica.com/madagascar/?page=1",
		"https://allafrica.com/malawi/?page=1",
		"https://allafrica.com/mali/?page=1",
		"https://allafrica.com/mauritius/?page=1",
		"https://allafrica.com/morocco/?page=1",
		"https://allafrica.com/mozambique/?page=1",
		"https://allafrica.com/namibia/?page=1",
		"https://allafrica.com/niger/?page=1",
		"https://allafrica.com/nigeria/?page=1",
		"https://allafrica.com/rwanda/?page=1",
		"https://allafrica.com/senegal/?page=1",
		"https://allafrica.com/seychelles/?page=1",
		"https://allafrica.com/sierraleone/?page=1",
		"https://allafrica.com/somalia/?page=1",
		"https://allafrica.com/southafrica/?page=1",
		"https://allafrica.com/southsudan/?page=1",
		"https://allafrica.com/sudan/?page=1",
		"https://allafrica.com/saotomeandprincipe/?page=1",
		"https://allafrica.com/tanzania/?page=1",
		"https://allafrica.com/togo/?page=1",
		"https://allafrica.com/tunisia/?page=1",
		"https://allafrica.com/uganda/?page=1",
		"https://allafrica.com/westernsahara/?page=1",
		"https://allafrica.com/zambia/?page=1",
		"https://allafrica.com/zimbabwe/?page=1",
		"https://allafrica.com/africa/?page=1",
		"https://allafrica.com/centralafrica/?page=1",
		"https://allafrica.com/eastafrica/?page=1",
		"https://allafrica.com/northafrica/?page=1",
		"https://allafrica.com/southernafrica/?page=1",
		"https://allafrica.com/westafrica/?page=1",
		"https://allafrica.com/aids/?page=1",
		"https://allafrica.com/agribusiness/?page=1",
		"https://allafrica.com/aid/?page=1",
		"https://allafrica.com/armsandarmies/?page=1",
		"https://allafrica.com/asiaaustraliaandafrica/?page=1",
		"https://allafrica.com/banking/?page=1",
		"https://allafrica.com/bookreviews/?page=1",
		"https://allafrica.com/books/?page=1",
		"https://allafrica.com/business/?page=1",
		"https://allafrica.com/children/?page=1",
		"https://allafrica.com/climate/?page=1",
		"https://allafrica.com/commodities/?page=1",
		"https://allafrica.com/company/?page=1",
		"https://allafrica.com/conflict/?page=1",
		"https://allafrica.com/construction/?page=1",
		"https://allafrica.com/coronavirus/?page=1",
		"https://allafrica.com/corruption/?page=1",
		"https://allafrica.com/currencies/?page=1",
		"https://allafrica.com/debt/?page=1",
		"https://allafrica.com/ebola/?page=1",
		"https://allafrica.com/ecotourism/?page=1",
		"https://allafrica.com/education/?page=1",
		"https://allafrica.com/energy/?page=1",
		"https://allafrica.com/environment/?page=1",
		"https://allafrica.com/europeandafrica/?page=1",
		"https://allafrica.com/externalrelations/?page=1",
		"https://allafrica.com/agriculture/?page=1",
		"https://allafrica.com/gameparks/?page=1",
		"https://allafrica.com/governance/?page=1",
		"https://allafrica.com/health/?page=1",
		"https://allafrica.com/humanrights/?page=1",
		"https://allafrica.com/ict/?page=1",
		"https://allafrica.com/infrastructure/?page=1",
		"https://allafrica.com/innovation/?page=1",
		"https://allafrica.com/io/?page=1",
		"https://allafrica.com/investment/?page=1",
		"https://allafrica.com/labour/?page=1",
		"https://allafrica.com/land/?page=1",
		"https://allafrica.com/latinamericaandafrica/?page=1",
		"https://allafrica.com/legalaffairs/?page=1",
		"https://allafrica.com/malaria/?page=1",
		"https://allafrica.com/manufacturing/?page=1",
		"https://allafrica.com/media/?page=1",
		"https://allafrica.com/middleeastandafrica/?page=1",
		"https://allafrica.com/migration/?page=1",
		"https://allafrica.com/mining/?page=1",
		"https://allafrica.com/ncds/?page=1",
		"https://allafrica.com/ngo/?page=1",
		"https://allafrica.com/nutrition/?page=1",
		"https://allafrica.com/oceans/?page=1",
		"https://allafrica.com/olympics/?page=1",
		"https://allafrica.com/peacekeeping/?page=1",
		"https://allafrica.com/petroleum/?page=1",
		"https://allafrica.com/polio/?page=1",
		"https://allafrica.com/pregnancy/?page=1",
		"https://allafrica.com/privatization/?page=1",
		"https://allafrica.com/refugees/?page=1",
		"https://allafrica.com/religion/?page=1",
		"https://allafrica.com/science/?page=1",
		"https://allafrica.com/stockmarkets/?page=1",
		"https://allafrica.com/sustainable/?page=1",
		"https://allafrica.com/terrorism/?page=1",
		"https://allafrica.com/trade/?page=1",
		"https://allafrica.com/transport/?page=1",
		"https://allafrica.com/travel/?page=1",
		"https://allafrica.com/tuberculosis/?page=1",
		"https://allafrica.com/usafrica/?page=1",
		"https://allafrica.com/urbanissues/?page=1",
		"https://allafrica.com/water/?page=1",
		"https://allafrica.com/wildlife/?page=1",
		"https://allafrica.com/women/?page=1"})

	// 从翻页器获取链接并访问
	w.OnHTML("div.pagination>div>ul>li>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 从index访问新闻
	w.OnHTML("ul.stories>li>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	// report.title
	w.OnHTML("h2.headline", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	// report .content
	w.OnHTML("div.story-body", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
	//report.publish time
	w.OnHTML("div.publication-date", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})
	// report.author
	w.OnHTML("cite.byline", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
}
