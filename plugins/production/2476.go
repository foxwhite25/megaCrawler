package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2476", "Alzheimer's Weekly", "https://alzheimersweekly.com/")

	engine.SetStartingURLs([]string{
		"https://alzheimersweekly.com/post-sitemap.xml",
		"https://alzheimersweekly.com/post-sitemap2.xml",
	})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        false,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Text, crawlers.News)
	})

	engine.OnHTML("div.video-container > iframe,div.wp-block-embed__wrapper > iframe", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		fileURL := element.Attr("src")
		if strings.Contains(fileURL, "youtube") {
			url, err := element.Request.URL.Parse(element.Attr("src"))
			if err != nil {
				crawlers.Sugar.Error(err.Error())
				return
			}
			ctx.Video = append(ctx.Video, url.String())
			ctx.PageType = crawlers.Report
		}
	})

	engine.OnHTML("div.jet-listing.jet-listing-dynamic-image > img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = []string{element.Attr("src")}
	})

	engine.OnHTML(`div.elementor-widget.elementor-widget-theme-post-content > div > p:not(:has(script)), 
		div.elementor-widget.elementor-widget-theme-post-content > div > ul,
		div.elementor-widget.elementor-widget-theme-post-content > div > h3`,
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			ctx.Content += element.Text
		})
}
