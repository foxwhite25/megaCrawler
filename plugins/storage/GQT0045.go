package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("GQT0045", "Qiushi", "http://en.qstheory.cn/")

	engine.SetStartingURLs([]string{
		"http://en.qstheory.cn/exclusive.html",
		"http://en.qstheory.cn/xismoments.html",
		"http://en.qstheory.cn/xisspeeches.html",
		"http://en.qstheory.cn/selectedreadingsvolumeI.html",
		"http://en.qstheory.cn/selectedreadingsvolumeII.html",
		"http://en.qstheory.cn/focus.html",
		"http://en.qstheory.cn/policyanalysis.html",
		"http://en.qstheory.cn/developmentexperience.html",
		"http://en.qstheory.cn/opinion.html",
		"http://en.qstheory.cn/chinaandtheworld.html",
		"http://en.qstheory.cn/thegovernanceofchinaI.html",
		"http://en.qstheory.cn/thegovernanceofchinaII.html",
		"http://en.qstheory.cn/thegovernanceofchinaIII.html",
		"http://en.qstheory.cn/thegovernanceofchinaIV.html",
	})

	extractorConfig := extractors.Config{
		Author:       false,
		Image:        false,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML("h3 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".info2_l > span:nth-child(1)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		fulltextwithauthor := element.Text
		textwithauthor := strings.Split(fulltextwithauthor, "Source: ")[1]
		ctx.Authors = append(ctx.Authors, textwithauthor)
	})

	engine.OnHTML("span[class=\"br0\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		fulltextwithdate := element.Text
		textwithdate := strings.Split(fulltextwithdate, "Updated: ")[1]
		ctx.PublicationTime = textwithdate
	})

	engine.OnHTML(".arcCont2 > p:not(p[style=\"text-align:center\"], p[style=\"text-align: center;\"])", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML("center > :nth-last-child(1)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
