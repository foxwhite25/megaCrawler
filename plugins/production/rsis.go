package production

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	getNextIndexURL := func(currentURL string) string {
		thisURL, _ := url.Parse(currentURL)
		path := thisURL.Path

		reg := regexp.MustCompile(`page/(\d+)`)
		rawStr := reg.FindStringSubmatch(path)
		num, _ := strconv.Atoi(rawStr[1])
		num++
		newURL := reg.ReplaceAllString(currentURL, "page/"+strconv.Itoa(num))

		return newURL
	}

	w := crawlers.Register("rsis", "拉惹勒南国际研究院", "https://www.rsis.edu.sg/")

	w.SetStartingURLs([]string{
		"https://www.rsis.edu.sg/research/cms/",
		"https://www.rsis.edu.sg/research/nts-centre/",
		"https://www.rsis.edu.sg/research/cens/",
		"https://www.rsis.edu.sg/research/idss/",
		"https://www.rsis.edu.sg/research/icpvtr/",
		"https://www.rsis.edu.sg/research/nssp/",
		"https://www.rsis.edu.sg/research/srp/",
		"https://www.rsis.edu.sg/research/other-programmes/stsp/",
	})

	// 访问 Index 从频道入口
	w.OnHTML(`#main > div.landing_page > div.accordion.rsis-accordion > div:nth-child(3) > div > div > div > div > div > div.landing-publication-card-links > span > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	// 访问下一页 Index
	w.OnHTML(`#main > div > div > div.content-view-container > div.col-2-3.print-expand.rsispub-listing-page-container.new-publication > article > section.publication-listing.clearfix > ul.pagination.list.clearfix > li.par-3.pagination-item.active`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if !strings.Contains(ctx.URL, "page") {
			w.Visit(element.ChildAttr("a.link", "href"), crawlers.Index)
		} else {
			nextURL := getNextIndexURL(ctx.URL)
			w.Visit(nextURL, crawlers.Index)
		}
	})

	// 访问 Report 从 Index
	w.OnHTML(`#main > div > div > div.content-view-container > div.col-2-3.print-expand.rsispub-listing-page-container.new-publication > article > section.publication-listing.clearfix > ul:nth-child(1) > div > div > div.rsis-container-card-roll.rsis-container-card-roll-persistant-color > div > p > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 获取 Title
	w.OnHTML(`#main > div > div > div.content-view-container > div > div.publication-details-section > div > div.publication-details-section-pane.pane-left > div > div.title`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`#main > div > div > div.content-view-container > div.col-2-3.print-expand.single-rsispub-page-container.new-publication > div.publication-details-section > div > div.publication-details-section-pane.pane-left > div > div.publication-details-section-content-block > div.publication-details-date > p`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 Authors
	w.OnHTML(`#main > div > div > div.content-view-container > div.col-2-3.print-expand.single-rsispub-page-container.new-publication > div.publication-details-section > div > div.publication-details-section-pane.pane-left > div > div.publication-details-section-content-block > div.publication-details-author`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, ",") {
			ctx.Authors = append(ctx.Authors, crawlers.SplitDelimiters(element.Text, []string{"&", "and"})...)
		} else {
			ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
		}
	})

	// 获取 CommentCount
	w.OnHTML(`#main-nav > div > nav > ul > li.nav-tab.nav-tab--primary.tab-conversation.active > a > span`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		// 裁出数字的字符串并将其转换为 int 类型
		var str = strings.Replace(element.Text, "comments", "", 1)
		str = strings.TrimSpace(str)
		num, _ := strconv.Atoi(str)
		ctx.CommentCount = num
	})

	// 获取 Description
	w.OnHTML(`#main > div > div > div.content-view-container > div.col-2-3.print-expand.single-rsispub-page-container.new-publication > div.publication-content-section > div > div.publication-content-section-pane.pane-left.force-full-width-panel > div > div.publication-content-content`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = strings.TrimSpace(element.Text)
	})

	// 获取 Tags
	w.OnHTML(`#main > div > div > div.content-view-container > div.col-2-3.print-expand.single-rsispub-page-container.new-publication > div.publication-content-section > div > div.publication-content-section-pane.pane-left.force-full-width-panel > div > section > span > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})

	// 获取 File
	w.OnHTML(`#main > div > div > div.content-view-container > div.col-2-3.print-expand.single-rsispub-page-container.new-publication > div.publication-details-section > div > div.publication-details-section-pane.pane-right > div > div > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
	})
}
