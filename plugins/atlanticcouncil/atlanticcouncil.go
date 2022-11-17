package atlanticcouncil

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"strings"
)

func init() {
	w := megaCrawler.Register("atlanticcouncil", "Adrienne Arsht Latin America Center", "https://www.atlanticcouncil.org/")

	w.SetStartingUrls([]string{"https://www.atlanticcouncil.org/programs/adrienne-arsht-latin-america-center/china-latin-america/",
		"https://www.atlanticcouncil.org/programs/adrienne-arsht-latin-america-center/gender-equality-and-diversity-in-latin-america/",
		"https://www.atlanticcouncil.org/region/brazil/",
		"https://www.atlanticcouncil.org/programs/adrienne-arsht-latin-america-center/caribbean-initiative/",
		"https://www.atlanticcouncil.org/region/colombia/",
		"https://www.atlanticcouncil.org/region/Mexico/",
		"https://www.atlanticcouncil.org/region/northern-triangle/",
		"https://www.atlanticcouncil.org/region/venezuela/",
		"https://www.atlanticcouncil.org/programs/adrienne-arsht-latin-america-center/proactivelac-series/",
		"https://www.atlanticcouncil.org/issue/climate-change-climate-action/",
		"https://www.atlanticcouncil.org/issue/economy-business/",
		"https://www.atlanticcouncil.org/issue/coronavirus/",
		"https://www.atlanticcouncil.org/issue/civil-society/",
		"https://www.atlanticcouncil.org/issue/politics-diplomacy/",
		"https://www.atlanticcouncil.org/issue/elections/",
		"https://www.atlanticcouncil.org/issue/democratic-transitions/",
		"https://www.atlanticcouncil.org/issue/freedom-and-prosperity/",
		"https://www.atlanticcouncil.org/issue/international-norms/",
		"https://www.atlanticcouncil.org/issue/political-reform/",
		"https://www.atlanticcouncil.org/issue/corruption/",
		"https://www.atlanticcouncil.org/issue/g20/",
		"https://www.atlanticcouncil.org/issue/rule-of-law/",
		"https://www.atlanticcouncil.org/issue/united-nations/",
		"https://www.atlanticcouncil.org/issue/media/",
		"https://www.atlanticcouncil.org/issue/security-defense/",
		"https://www.atlanticcouncil.org/issue/conflict/",
		"https://www.atlanticcouncil.org/issue/defense-industry/",
		"https://www.atlanticcouncil.org/issue/defense-technologies/",
		"https://www.atlanticcouncil.org/issue/intelligence/",
		"https://www.atlanticcouncil.org/issue/nato/",
		"https://www.atlanticcouncil.org/issue/nuclear-nonproliferation/",
		"https://www.atlanticcouncil.org/issue/security-partnerships/",
		"https://www.atlanticcouncil.org/issue/crisis-management/",
		"https://www.atlanticcouncil.org/issue/arms-control/",
		"https://www.atlanticcouncil.org/issue/defense-policy/",
		"https://www.atlanticcouncil.org/issue/extremism/",
		"https://www.atlanticcouncil.org/issue/national-security/",
		"https://www.atlanticcouncil.org/issue/peacekeeping-and-peacebuilding/",
		"https://www.atlanticcouncil.org/issue/terrorism/",
		"https://www.atlanticcouncil.org/issue/digital-policy/",
		"https://www.atlanticcouncil.org/issue/digital-currencies/",
		"https://www.atlanticcouncil.org/issue/economic-sanctions/",
		"https://www.atlanticcouncil.org/issue/geopolitics-energy-security/",
		"https://www.atlanticcouncil.org/issue/renewables-advanced-energy/",
		"https://www.atlanticcouncil.org/issue/human-rights/",
		"https://www.atlanticcouncil.org/issue/nationalism/",
		"https://www.atlanticcouncil.org/issue/education/",
		"https://www.atlanticcouncil.org/issue/fiscal-and-structural-reform/",
		"https://www.atlanticcouncil.org/issue/disinformation/",
		"https://www.atlanticcouncil.org/issue/cybersecurity/",
		"https://www.atlanticcouncil.org/issue/space/",
		"https://www.atlanticcouncil.org/region/caribbean/",
		"https://www.atlanticcouncil.org/region/united-states-canada/",
		"https://www.atlanticcouncil.org/region/cuba/",
		"https://www.atlanticcouncil.org/region/latin-america/",
		"https://www.atlanticcouncil.org/region/democratic-repubic-congo/",
		"https://www.atlanticcouncil.org/region/eritrea/",
		"https://www.atlanticcouncil.org/region/general-africa/",
		"https://www.atlanticcouncil.org/region/nigeria/",
		"https://www.atlanticcouncil.org/region/sahel/",
		"https://www.atlanticcouncil.org/region/south-africa/",
		"https://www.atlanticcouncil.org/region/somalia/",
		"https://www.atlanticcouncil.org/region/angola/",
		"https://www.atlanticcouncil.org/region/east-africa/",
		"https://www.atlanticcouncil.org/region/ethiopia/",
		"https://www.atlanticcouncil.org/region/morocco/",
		"https://www.atlanticcouncil.org/region/sudan/",
		"https://www.atlanticcouncil.org/region/balkans/",
		"https://www.atlanticcouncil.org/region/the-caucasus/",
		"https://www.atlanticcouncil.org/region/germany/",
		"https://www.atlanticcouncil.org/region/hungary/",
		"https://www.atlanticcouncil.org/region/moldova/",
		"https://www.atlanticcouncil.org/region/russia/",
		"https://www.atlanticcouncil.org/programs/eurasia-center/research/",
		"https://www.atlanticcouncil.org/region/ukraine/",
		"https://www.atlanticcouncil.org/region/european-union/",
		"https://www.atlanticcouncil.org/region/belarus/",
		"https://www.atlanticcouncil.org/region/france/",
		"https://www.atlanticcouncil.org/region/greece/",
		"https://www.atlanticcouncil.org/region/italy/",
		"https://www.atlanticcouncil.org/region/poland/",
		"https://www.atlanticcouncil.org/region/turkey/",
		"https://www.atlanticcouncil.org/issue/eurozone/",
		"https://www.atlanticcouncil.org/region/united-kingdom/",
		"https://www.atlanticcouncil.org/region/australia/",
		"https://www.atlanticcouncil.org/region/china/",
		"https://www.atlanticcouncil.org/region/japan/",
		"https://www.atlanticcouncil.org/region/pakistan/",
		"https://www.atlanticcouncil.org/region/afghanistan/",
		"https://www.atlanticcouncil.org/region/bangladesh/",
		"https://www.atlanticcouncil.org/region/india/",
		"https://www.atlanticcouncil.org/region/korea/",
		"https://www.atlanticcouncil.org/region/taiwan/",
		"https://www.atlanticcouncil.org/region/iraq/",
		"https://www.atlanticcouncil.org/region/lebanon/",
		"https://www.atlanticcouncil.org/region/saudi-arabia/",
		"https://www.atlanticcouncil.org/region/yemen/",
		"https://www.atlanticcouncil.org/region/iran/",
		"https://www.atlanticcouncil.org/region/israel/",
		"https://www.atlanticcouncil.org/region/syria/",
		"https://www.atlanticcouncil.org/in-depth-research-reports/",
		"https://www.atlanticcouncil.org/fastthinking/",
		"https://www.atlanticcouncil.org/new-atlanticist/experts-react/",
		"https://www.atlanticcouncil.org/category/content-series/inflection-points/",
		"https://www.atlanticcouncil.org/category/blogs/ukrainealert/",
		"https://www.atlanticcouncil.org/category/blogs/menasource/",
		"https://www.atlanticcouncil.org/category/blogs/iransource/",
		"https://www.atlanticcouncil.org/category/blogs/africasource/",
		"https://www.atlanticcouncil.org/category/blogs/energysource/",
		"https://www.atlanticcouncil.org/category/blogs/southasiasource/",
		"https://www.atlanticcouncil.org/category/blogs/geotech-cues/",
		"https://www.atlanticcouncil.org/category/blogs/turkeysource/",
		"https://www.atlanticcouncil.org/programs/scowcroft-center-for-strategy-and-security/research/",
		"https://www.atlanticcouncil.org/programs/scowcroft-center-for-strategy-and-security/asia-security-initiative/events/",
		"https://www.atlanticcouncil.org/programs/scowcroft-center-for-strategy-and-security/asia-security-initiative/commentary/",
		"https://www.atlanticcouncil.org/programs/scowcroft-center-for-strategy-and-security/asia-security-initiative/in-the-news/",
		"https://www.atlanticcouncil.org/programs/scowcroft-center-for-strategy-and-security/asia-security-initiative/research/",
		"https://www.atlanticcouncil.org/programs/adrienne-arsht-latin-america-center/all-events/",
		"https://www.atlanticcouncil.org/programs/adrienne-arsht-latin-america-center/commentary/",
		"https://www.atlanticcouncil.org/programs/adrienne-arsht-latin-america-center/research/",
		"https://www.atlanticcouncil.org/programs/adrienne-arsht-latin-america-center/insights-impact/",
		"https://www.atlanticcouncil.org/programs/africa-center/events/",
		"https://www.atlanticcouncil.org/programs/africa-center/commentary/",
		"https://www.atlanticcouncil.org/programs/africa-center/research/",
		"https://www.atlanticcouncil.org/global-africa-africa-center/",
		"https://www.atlanticcouncil.org/prosperity-africa-center/",
		"https://www.atlanticcouncil.org/rule-of-law/",
		"https://www.atlanticcouncil.org/programs/eurasia-center/ukraine-in-europe-initiative/",
		"https://www.atlanticcouncil.org/programs/eurasia-center/commentary/",
		"https://www.atlanticcouncil.org/programs/eurasia-center/all-events/",
		"https://www.atlanticcouncil.org/programs/eurasia-center/insight-impact/",
		"https://www.atlanticcouncil.org/programs/eurasia-center/research/",
		"https://www.atlanticcouncil.org/category/blogs/econographics/",
		"https://www.atlanticcouncil.org/programs/geoeconomics-center/trackers/",
		"https://www.atlanticcouncil.org/programs/geoeconomics-center/global-sanctions-dashboard/",
		"https://www.atlanticcouncil.org/programs/geoeconomics-center/commentary/",
		"https://www.atlanticcouncil.org/programs/geoeconomics-center/all-events/",
		"https://www.atlanticcouncil.org/programs/geoeconomics-center/research/",
		"https://www.atlanticcouncil.org/programs/global-energy-center/energysource-innovation-stream/",
		"https://www.atlanticcouncil.org/programs/global-energy-center/climate-and-advanced-energy/",
		"https://www.atlanticcouncil.org/programs/global-energy-center/european-energy-security/commentary-analysis/",
		"https://www.atlanticcouncil.org/programs/global-energy-center/european-energy-security/events/",
		"https://www.atlanticcouncil.org/programs/global-energy-center/hydrogen-policy-sprint/",
		"https://www.atlanticcouncil.org/programs/global-energy-center/nuclear-energy-policy-initiative/commentary-and-analysis/",
		"https://www.atlanticcouncil.org/programs/global-energy-center/nuclear-energy-policy-initiative/events/"})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = megaCrawler.Report
		}
	})

	//https://www.atlanticcouncil.org/programs/adrienne-arsht-latin-america-center/china-latin-america/
	//expert
	//div.gta-embed--content.gta-expert-embed--content > h3 > a
	// 从翻页器获取链接并访问
	w.OnHTML("div.p-archives--container.row > div> a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Index)
	})
	w.OnHTML("div.j-posts--pagination>a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Index)
	})
	//访问expert
	w.OnHTML("div.gta-embed--content.gta-expert-embed--content > h3 > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Expert)
	})

	// expert.Name
	w.OnHTML("#masthead > section > div > div > div > div.gta-site-banner--event-column-right > h2 > font > font", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Name = element.Text
	})
	// expert.title
	w.OnHTML("#masthead > section > div > div > div > div.gta-site-banner--event-column-right > ul > li:nth-child(1) > font > font", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Title = element.Text
	})
	w.OnHTML("#masthead > section > div > div > div > div.gta-site-banner--event-column-right > ul > li:nth-child(2) > font > font", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Title = ctx.Title + "," + element.Text
	})
	// expert.description
	w.OnHTML("#expert-content > div > div > div", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Content = element.Text
	})
	// expert.area
	w.OnHTML("#content > section > div > div.ac-single-expert--tags-content--left > div > div > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Area = ctx.Area + "," + element.Text
	})
	// expert.link
	w.OnHTML("#content > div > div > a.e-share-link.e-share-link__email", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Link = append(ctx.Link, element.Attr("href"))
	})
	w.OnHTML("#content > div > div > a.e-share-link.e-share-link__twitter", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Link = append(ctx.Link, element.Attr("href"))
	})

	//expert.location
	w.OnHTML("div > div.ac-single-expert--tags-content--left > div > div:nth-child(2) > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Location = ctx.Location + "," + element.Text
	})

	// interview
	//访问new

	//内容一般处理
	//new .content
	w.OnHTML("#content > section > div > div > div > p", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Content = ctx.Content + element.Text
	})

	// 从翻页器获取链接并访问
	w.OnHTML(" div.j-posts--pagination.columns-12.o-archives--pagination > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Index)
	})

	w.OnHTML("a.gta-embed--link", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.News)
	})

	// new.category
	w.OnHTML("p.gta-site-banner--expert-author.upper > span > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.CategoryText = element.Text
	})
	//new.publish time
	w.OnHTML("div.gta-site-banner--heading.gta-post-site-banner--heading > p.gta-site-banner--heading--date.gta-post-site-banner--heading--date", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.PublicationTime = element.Text
	})

	// new.title
	w.OnHTML("#masthead > section > div > div > h2", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Title = element.Text
	})

	// report.author
	w.OnHTML("span.gta-embed--tax--expert.gta-post-embed--tax--expert", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	//

	// new.description
	w.OnHTML("#content > section > div > div > div > p:nth-child(2)", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Description = element.Text
	})

	//in the news
	// new.category
	w.OnHTML("p.gta-site-banner--expert-author.upper > span > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.CategoryText = element.Text
	})
	//new.publish time
	w.OnHTML("div.gta-site-banner--heading.gta-post-site-banner--heading > p.gta-site-banner--heading--date.gta-post-site-banner--heading--date", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.PublicationTime = element.Text
	})
	// new.title
	w.OnHTML("#masthead > section > div > div > h2", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Title = element.Text
	})
	//new.link
	w.OnHTML(" div.wp-container-1.wp-block-buttons > div > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Link = append(ctx.Link, element.Attr("href"))
	})

	//Econographics
	// new.category
	w.OnHTML(" p.ac-single-post--marquee--expert-author.upper > span > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.CategoryText = element.Text
	})
	//new.publish time
	w.OnHTML(" p.ac-single-post--marquee--heading--date", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.PublicationTime = element.Text
	})
	// new.title
	w.OnHTML("#content > section > div > div > div > header > h2", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Title = element.Text
	})

	// new.author
	w.OnHTML("a.gta-site-banner--tax--expert", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	//report
	// new.category
	w.OnHTML("p.gta-site-banner--expert-author.upper > span > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.CategoryText = element.Text
	})

	//new.publish time
	w.OnHTML("div.gta-site-banner--heading.gta-post-site-banner--heading > p.gta-site-banner--heading--date.gta-post-site-banner--heading--date", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.PublicationTime = element.Text
	})

	// new.title
	w.OnHTML("#masthead > section > div > div > h2", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Title = element.Text
	})

	// report.author
	w.OnHTML(" a.gta-site-banner--tax--expert", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	// new.description
	w.OnHTML("#content > section > div > div > div > p:nth-child(2)", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Description = element.Text
	})

	//new Atlanticist 和Engagement Reframed In-Depth Research & Reports
	// new.category
	w.OnHTML(" p.ac-single-post--marquee--expert-author.upper > span > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.CategoryText = element.Text
	})

	//new.publish time
	w.OnHTML(" p.ac-single-post--marquee--heading--date", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.PublicationTime = element.Text
	})
	// new.title
	w.OnHTML("h2.ac-single-post--marquee--title", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Title = element.Text
	})

	// new.author
	w.OnHTML(" span.gta-embed--tax--expert", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	w.OnHTML(" a.gta-site-banner--tax--expert", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	// new .content
	w.OnHTML("div.ac-single-post--content", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Content = element.Text
	})
	// new.link
	w.OnHTML("a.wp-block-button__link", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Link = append(ctx.Link, element.Attr("href"))
	})

	//past event
	// new.category
	w.OnHTML(" span.gta-site-banner--tax--cats", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.CategoryText = element.Text
	})
	//new.publish time
	w.OnHTML("p.gta-site-banner--heading--date", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.PublicationTime = element.Text
	})
	// new.title
	w.OnHTML("h2.gta-site-banner--title ", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Title = element.Text
	})

	// new.author
	w.OnHTML(" span.gta-embed--tax--expert", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	// new .content
	w.OnHTML("div.ac-single-post--content", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Content = element.Text
	})

	//https://www.atlanticcouncil.org/programs/adrienne-arsht-latin-america-center/gender-equality-and-diversity-in-latin-america/
	//访问new

	// 从翻页器获取链接并访问
	w.OnHTML(" div.j-posts--pagination.columns-12.o-archives--pagination > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Index)
	})
	w.OnHTML("a.gta-embed--link", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.News)
	})

	//new Atlanticist   Diversity, Equity, and Inclusion   Expert react   GTA系列
	// new.category
	w.OnHTML(" p.gta-site-banner--expert-author.upper > span > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.CategoryText = element.Text
	})
	//new.publish time
	w.OnHTML("  p.gta-site-banner--heading--date.gta-post-site-banner--heading--date", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.PublicationTime = element.Text
	})
	// new.title
	w.OnHTML("h2.gta-site-banner--title", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Title = element.Text
	})

	// new.author
	w.OnHTML(" span.gta-embed--tax--expert.gta-post-embed--tax--expert", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	// new .content
	w.OnHTML("div.ac-single-post--content", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Content = element.Text
	})

	//https://www.atlanticcouncil.org/region/brazil/

	//Fast Thinking   AC系列
	//访问new - 与上一样

	// new.category
	w.OnHTML(" p.ac-single-post--marquee--expert-author.upper > span > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.CategoryText = element.Text
	})

	//new.publish time
	w.OnHTML(" p.ac-single-post--marquee--heading--date", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.PublicationTime = element.Text
	})
	// new.title
	w.OnHTML("h2.ac-single-post--marquee--title", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Title = element.Text
	})

	// new.author
	w.OnHTML(" span.gta-embed--tax--expert", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	w.OnHTML(" a.gta-site-banner--tax--expert", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	// new .content
	w.OnHTML("div.ac-single-post--content", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Content = element.Text
	})
	// new.link
	w.OnHTML("a.wp-block-button__link", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Link = append(ctx.Link, element.Attr("href"))
	})

	//Aviso LatAm: COVID-19 GTA系列
	//Issue Brief与Econographics一样
	//Spotlight  AC系列
	//TradeWorld AC系列
	//Diversity, Equity, and Inclusion  GTA系列
	//Event Recap
	//Article GTA系列
	//UkraineAlert  GTA系列

	//Featured commentary & analysis
	// 访问新闻
	w.OnHTML("a.gta-horizontal-featured--link", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.News)
	})

}
