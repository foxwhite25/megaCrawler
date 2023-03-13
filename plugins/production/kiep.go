package production

import (
	"strconv"
	"strings"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("kiep", "对外经济政策研究所", "https://www.kiep.go.kr/")

	w.SetStartingURLs([]string{
		"https://www.kiep.go.kr/gallery.es?mid=a20301000000&bid=0007",
		"https://www.kiep.go.kr/gallery.es?mid=a20303000000&bid=0001&cg_code=C01,C02,C03,C04,C13,C19",
		"https://www.kiep.go.kr/gallery.es?mid=a20304000000&bid=0001&cg_code=C05,C06,C07,C09,C10,C11,C12",
		"https://www.kiep.go.kr/board.es?mid=a20401000000&bid=0051",
		"https://www.kiep.go.kr/gallery.es?mid=a20308000000&bid=0008",
	})

	// 访问 Report 从 Index
	w.OnHTML(`.board_report > li > a.title`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 访问 Report 从 Index (Type 2)
	w.OnHTML(`.board_book .desc > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 访问 Report 从 Index (Type 3)
	w.OnHTML(`.txt_left[aria-label="Title"] > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 访问下一页 Index
	w.OnHTML(`.board_pager > a[class="arr next"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 获取 Title
	w.OnHTML(`.board_view > h2.title`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Title (Type 2)
	w.OnHTML(`.board_book .desc > strong.title`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Description
	w.OnHTML(`.contents > div`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = strings.TrimSpace(element.Text)
	})

	// 获取 Description (Type 2)
	w.OnHTML(`.cont > div > div.txt`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = strings.TrimSpace(element.Text)
	})

	// 获取 Description (Type 3)
	w.OnHTML(`#contents_body > article.board_view > div.contents`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = strings.TrimSpace(element.Text)
	})

	// 获取 Description (Type 4)
	w.OnHTML(`#contents_body > div > div.contents`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`.info > li.date > span`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime (Type 2)
	w.OnHTML(`.board_book .desc > .info > span:nth-child(4) > strong`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 Authors
	w.OnHTML(`.info > li.write > span`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		// 若有 "," 或者 "and" 则为多人作者
		if strings.Contains(element.Text, ",") {
			ctx.Authors = append(ctx.Authors, crawlers.SplitDelimiters(element.Text, []string{"&", "and"})...)
		} else {
			ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
		}
	})

	// 获取 Authors (Type 2)
	w.OnHTML(`.board_book .desc > .info > span:nth-child(1) > strong`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		// 若有 "," 或者 "and" 则为多人作者
		if strings.Contains(element.Text, ",") {
			ctx.Authors = append(ctx.Authors, crawlers.SplitDelimiters(element.Text, []string{"&", "and"})...)
		} else {
			ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
		}
	})

	// 获取 Language
	w.OnHTML(`.board_book .desc > .info > span:nth-child(3) > strong`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Language = strings.TrimSpace(element.Text)
	})

	// 获取 Content
	w.OnHTML(`div.elementor-widget-container`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 获取 Tags
	w.OnHTML(`.board_book .desc > .category`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		// 若有 "," 或者 "and" 则为多个 Tag
		if strings.Contains(element.Text, ",") {
			ctx.Tags = append(ctx.Tags, crawlers.SplitDelimiters(element.Text, []string{"&", "and"})...)
		} else {
			ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
		}
	})

	// 获取 File
	w.OnHTML(`.file > .list .link > .btn_line`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		fileURL := "https://www.kiep.go.kr" + element.Attr("href")
		ctx.File = append(ctx.File, fileURL)
	})

	// 获取 File (Type 2)
	w.OnHTML(`p[class="btns txt_left"] > a.btn1`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		fileURL := "https://www.kiep.go.kr" + element.Attr("href")
		ctx.File = append(ctx.File, fileURL)
	})

	// 获取 Location
	w.OnHTML(`#contents_body > article.board_view > ul > li:nth-child(4) > span`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Location = strings.TrimSpace(element.Text)
	})

	// 获取 ViewCount
	w.OnHTML(`#contents_body > article > ul > li.hit > span`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		num, _ := strconv.Atoi(strings.TrimSpace(element.Text))
		ctx.ViewCount = num
	})
}
