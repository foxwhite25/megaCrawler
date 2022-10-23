package heritage

import (
	"github.com/gocolly/colly/v2"
	_ "github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	_ "megaCrawler/megaCrawler"
	"regexp"
	_ "regexp"
	"strings"
	_ "strings"
)

func init() {
	w := megaCrawler.Register("heritage", "美国传统基金会", "https://www.heritage.org/")
	w.SetStartingUrls([]string{"https://www.heritage.org/about-heritage/staff/experts", "https://www.heritage.org/"})
	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告 1
	w.OnHTML("a.button", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = megaCrawler.Report
		}
	})

	// 从翻页器获取链接并访问 1
	w.OnHTML(".button-more ", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Index)
	})

	// 尝试访问作者并添加到ctx 1
	w.OnHTML(".person-list-small__image-wrapper", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if element.ChildAttr("a", "href") != "" {
			w.Visit(element.ChildAttr("a", "href"), megaCrawler.Expert)
		}

	})

	w.OnHTML(".js-hover-target >div", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	// 访问新闻 1
	w.OnHTML("article >a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.News)
	})
	// 访问新闻 1
	w.OnHTML("a[hreflang=\"en\"]", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.News)
	})

	//访问新闻 1
	w.OnHTML(".result-card", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if strings.Contains(element.Attr("class"), "result-card__video") {
			return
		} else if strings.Contains(element.Attr("class"), "_has-video") {
			return
		}
		emailRegex, _ := regexp.Compile("href=\\\"([a-zA-Z/-]+)\"")
		emailMatch := emailRegex.FindStringSubmatch(element.Text)
		w.Visit(emailMatch[1], megaCrawler.News)
	})
	/*

		#block-mainpagecontent > article > div.results-wrapper > div > div.views-element-container > div > div > div > div:nth-child(4) > section > div > div.result-card__info-wrapper > a
						<section class="result-card result-card__commentary js-hover-container _no-image">
						  <div class="result-card__info-container result-card__slide-container">
						    <p class="result-card__eyebrow">Commentary</p>
						        <div class="result-card__info-wrapper">

						      <a href="/africa/commentary/why-tunisias-future-matters-the-west" class="result-card__title js-hover-target" title="article link">Why Tunisia’s Future Matters to the West</a>
						                                <p class="result-card__date"><span>Sep 15, 2022</span>   3 min read
						</p>
						    </div>
						  </div>
						  <i class="heritage-icon-arrow_long"></i>
						</section>
					<section class="result-card result-card__commentary js-hover-container _no-image">
					  <div class="result-card__info-container result-card__slide-container">
					    <p class="result-card__eyebrow">Commentary</p>
					        <div class="result-card__info-wrapper">

					      <a href="/africa/commentary/why-tunisias-future-matters-the-west" class="result-card__title js-hover-target" title="article link">Why Tunisia’s Future Matters to the West</a>
					                                <p class="result-card__date"><span>Sep 15, 2022</span>   3 min read
					</p>
					    </div>
					  </div>
					  <i class="heritage-icon-arrow_long"></i>
					</section>
				<section class="result-card result-card__report js-hover-container">
				  <div class="result-card__info-wrapper result-card__slide-container">
				    <p class="result-card__eyebrow">report</p>
				        <a href="/europe/report/understanding-russias-threat-employ-nuclear-weapons-its-war-against-ukraine" title="heritage report" class="result-card__title js-hover-target">Understanding Russia’s Threat to Employ Nuclear Weapons Against Ukraine</a>
				          <a href="/staff/james-carafano" class="result-card__link">        <span>James Jay Carafano</span>
				      </a>
				    <p class="result-card__date"><span>Oct 6, 2022</span>   4 min read
				</p>
				  </div>
				  <i class="heritage-icon-arrow_long"></i>
				</section>
				<section class="result-card result-card__video js-hover-container">
				  <div class="result-card__info-wrapper result-card__slide-container">
				    <p class="result-card__eyebrow">video</p>

				    <a data-video-overlay="[{&quot;videoTitle&quot;:&quot;Russia and China Aren't Intimidated By Biden's Weak Foreign Policy | James Carafano on Fox Business&quot;,&quot;videoDate&quot;:&quot;Jun 14, 2021&quot;,&quot;videoLength&quot;:&quot;5 min video&quot;,&quot;videoScreenshot&quot;:&quot;https:\/\/www.heritage.org\/sites\/default\/files\/styles\/video_overlay_thumb\/public\/2021-06\/Screen%20Shot%202021-06-18%20at%209.37.43%20AM.png?itok=-YVKa20J&quot;,&quot;youtubeID&quot;:&quot;qwIE9Gf3p6k&quot;,&quot;videoBody&quot;:&quot;Heritage&amp;#039;s James Carafano joined Fox Business, Monday, June 14, to talk about why our enemies like Russia and China don&amp;#039;t take Biden&amp;#039;s foreign policy seriously.&quot;},{&quot;videoTitle&quot;:&quot;How Should NATO Handle Russia, China Challenges? | Luke Coffey on CNBC Asia&quot;,&quot;videoDate&quot;:&quot;Jun 14, 2021&quot;,&quot;videoLength&quot;:&quot;8 min video&quot;,&quot;videoScreenshot&quot;:&quot;https:\/\/www.heritage.org\/sites\/default\/files\/styles\/video_overlay_thumb\/public\/2021-06\/Screen%20Shot%202021-06-15%20at%2010.03.34%20AM.png?itok=W2fHRjs2&quot;,&quot;youtubeID&quot;:&quot;boemiuwHtDw&quot;,&quot;videoBody&quot;:&quot;Heritage&amp;#039;s Luke Coffey joined CNBC Asia, Monday, June 14, to talk about Pres. Biden&amp;#039;s NATO summit, how NATO needs to address the challenges of Russia and China.&quot;},{&quot;videoTitle&quot;:&quot;Does Joe Biden Have What It Takes to Stand Up to China? | Robert Wilkie on Fox News&quot;,&quot;videoDate&quot;:&quot;Jun 13, 2021&quot;,&quot;videoLength&quot;:&quot;4 min video&quot;,&quot;videoScreenshot&quot;:&quot;https:\/\/www.heritage.org\/sites\/default\/files\/styles\/video_overlay_thumb\/public\/2021-06\/Screen%20Shot%202021-06-14%20at%201.16.28%20PM.png?itok=BJOh_fye&quot;,&quot;youtubeID&quot;:&quot;j0nhviP7yhY&quot;,&quot;videoBody&quot;:&quot;Heritage Visiting Fellow Robert Wilkie joined Fox News, Sunday, June 13, to talk about whether the Biden administration will stand up to the Chinese Communist Party.&quot;}]" href="" class="result-card__title video-overlay-cover__trigger js-video-trigger js-hover-target" title="open video overlay">Russia and China Aren't Intimidated By Biden's Weak Foreign Policy | James Ca...</a>
				                      <p class="result-card__date"><span>Jun 14, 2021</span> </p>
				  </div>
				  <a data-video-overlay="[{&quot;videoTitle&quot;:&quot;Russia and China Aren't Intimidated By Biden's Weak Foreign Policy | James Carafano on Fox Business&quot;,&quot;videoDate&quot;:&quot;Jun 14, 2021&quot;,&quot;videoLength&quot;:&quot;5 min video&quot;,&quot;videoScreenshot&quot;:&quot;https:\/\/www.heritage.org\/sites\/default\/files\/styles\/video_overlay_thumb\/public\/2021-06\/Screen%20Shot%202021-06-18%20at%209.37.43%20AM.png?itok=-YVKa20J&quot;,&quot;youtubeID&quot;:&quot;qwIE9Gf3p6k&quot;,&quot;videoBody&quot;:&quot;Heritage&amp;#039;s James Carafano joined Fox Business, Monday, June 14, to talk about why our enemies like Russia and China don&amp;#039;t take Biden&amp;#039;s foreign policy seriously.&quot;},{&quot;videoTitle&quot;:&quot;How Should NATO Handle Russia, China Challenges? | Luke Coffey on CNBC Asia&quot;,&quot;videoDate&quot;:&quot;Jun 14, 2021&quot;,&quot;videoLength&quot;:&quot;8 min video&quot;,&quot;videoScreenshot&quot;:&quot;https:\/\/www.heritage.org\/sites\/default\/files\/styles\/video_overlay_thumb\/public\/2021-06\/Screen%20Shot%202021-06-15%20at%2010.03.34%20AM.png?itok=W2fHRjs2&quot;,&quot;youtubeID&quot;:&quot;boemiuwHtDw&quot;,&quot;videoBody&quot;:&quot;Heritage&amp;#039;s Luke Coffey joined CNBC Asia, Monday, June 14, to talk about Pres. Biden&amp;#039;s NATO summit, how NATO needs to address the challenges of Russia and China.&quot;},{&quot;videoTitle&quot;:&quot;Does Joe Biden Have What It Takes to Stand Up to China? | Robert Wilkie on Fox News&quot;,&quot;videoDate&quot;:&quot;Jun 13, 2021&quot;,&quot;videoLength&quot;:&quot;4 min video&quot;,&quot;videoScreenshot&quot;:&quot;https:\/\/www.heritage.org\/sites\/default\/files\/styles\/video_overlay_thumb\/public\/2021-06\/Screen%20Shot%202021-06-14%20at%201.16.28%20PM.png?itok=BJOh_fye&quot;,&quot;youtubeID&quot;:&quot;j0nhviP7yhY&quot;,&quot;videoBody&quot;:&quot;Heritage Visiting Fellow Robert Wilkie joined Fox News, Sunday, June 13, to talk about whether the Biden administration will stand up to the Chinese Communist Party.&quot;}]" href="#" class="result-card__image-wrapper js-video-trigger" style="background-image: url('https://www.heritage.org/sites/default/files/styles/content_listing_295x205/public/2021-06/Screen%20Shot%202021-06-18%20at%209.37.43%20AM.png?h=50e31a8f&amp;itok=ogth0w34')" title="video overlay">
				    <i class="heritage-icon-play_module"></i>
				  </a>
				</section>
			<section class="result-card result-card__event result-card__past-event js-hover-container _has-video">
				<div class="result-card__slide-container">
			  	<p class="result-card__eyebrow">Event</p>
					<div class="result-card__date-wrapper">
					  <p class="result-card__event-date">Oct 4</p><p class="result-card__time">2:15pm - 3:00pm</p>
					</div>
					<div class="result-card__info-wrapper">

					  <a href="/africa/event/us-uganda-partnership-turbulent-region-conversation-ugandas-minister-foreign-affairs" class="result-card__title js-hover-target">U.S.-Uganda Partnership in a Turbulent Region</a>
					  <p class="result-card__place"><span></span>  </p>
					</div>
				</div>
				    <a href="/africa/event/us-uganda-partnership-turbulent-region-conversation-ugandas-minister-foreign-affairs" class="result-card__image-wrapper js-hover-target" style="background-image: url(https://www.heritage.org/sites/default/files/styles/event_listing_image_430x205/public/images/2022-09/GettyImages-181147783.jpg?h=4e553640&amp;itok=u2O0L_oZ)">
			      <i class="heritage-icon-play_module"></i>
			    </a>

			</section>
	*/

	// 添加正文到ctx 1
	w.OnHTML(".person-list-small__panelist", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Content = element.Text
	})

	// 人物头衔到ctx 1
	w.OnHTML(".person-list-small__title", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Title = megaCrawler.StandardizeSpaces(element.Text)
	})

	// 人物描述到ctx 1
	w.OnHTML(".expert-bio-body__copy  >p", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Description = element.Text
	})

	// 专家领域到ctx 1
	w.OnHTML("font[style=\"vertical-align: inherit;\"]", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Area = element.Text
	})

	//专家图片到ctx
	w.OnHTML("a.expert-bio-card__download-headshot", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Image = []string{element.Attr("href")}
	})

	// 访问新闻
	w.OnHTML("article[role=\"article\"] >div>a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.News)
	})

	//new . author_name
	w.OnHTML(".author-card__author-info-wrapper>a>span", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	//new . author_information
	w.OnHTML(".author-card__card-info >p>font", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	//new . content
	w.OnHTML("#block-mainpagecontent > article > div > div > div", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Content = element.Text
	})
	//new . publish_time
	w.OnHTML("div.article-general-info", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.PublicationTime = element.Text
	})
	//new . author_url
	w.OnHTML(" div.commentary__intro-wrapper>a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Attr("href"))
	})
	// new. keyword
	w.OnHTML(" div.key-takeaways__takeaway >p", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Keywords = append(ctx.Keywords, element.Text)
	})

	// new .image
	w.OnHTML(" figure.image-with-caption__image-wrapper>img ", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Image = append(ctx.Keywords, element.Attr("srcset"))
	})

}
