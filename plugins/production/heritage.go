package production

import (
	"regexp"
	"strings"
	"time"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("heritage", "美国传统基金会", "https://www.heritage.org/")
	w.SetStartingURLs([]string{"https://www.heritage.org/sitemap.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "?page=") {
			w.Visit(element.Text, crawlers.Index)
		} else {
			if strings.Contains(element.Text, "staff") {
				w.Visit(element.Text, crawlers.Expert)
			}
			w.Visit(element.Text, crawlers.News)
		}
	})

	w.OnHTML(".headline", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType == crawlers.Expert {
			ctx.Name = element.Text
		} else {
			ctx.Title = element.Text
		}
	})

	w.OnHTML(".expert-bio-card__expert-title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	w.OnHTML(".expert-card__expert-name", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	w.OnHTML(".author-card__name", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	w.OnHTML(".author-card__multi-name", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	w.OnHTML(".article__body-copy", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType == crawlers.Expert {
			ctx.Description = strings.TrimSpace(element.Text)
		} else {
			ctx.Content = strings.TrimSpace(element.Text)
		}
	})

	w.OnHTML("meta[property=\"og:description\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType != crawlers.Expert {
			ctx.Description = element.Attr("content")
		}
	})

	w.OnHTML(".expert-bio-card__photo", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		style := element.Attr("style")
		reg := regexp.MustCompile(`https?://(www\.)?[-a-zA-Z0-9@:%._+~#=]{1,256}\.[a-zA-Z\d()]{1,6}\b([-a-zA-Z0-9@:%_+.~#?&/=]*)`)
		ctx.Image = []string{reg.FindString(style)}
	})

	w.OnHTML(".article__eyebrow > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.CategoryText = element.Text
	})

	w.OnHTML(".article-general-info", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		reg := regexp.MustCompile(`(\w+ \d+, \d+)`)
		match := reg.FindString(element.Text)
		times, err := time.Parse("Jan 2, 2006", match)
		if err != nil {
			crawlers.Sugar.Errorw(err.Error(), "Original", element.Text, "Regex", match)
			return
		}
		ctx.PublicationTime = times.Format(time.RFC3339)
	})
}
