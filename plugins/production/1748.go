package production

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

type structure struct {
	SiteInfo struct {
		DallowedActions interface{} `json:"dallowedActions"`
		Base            struct {
			Properties struct {
				ThemeName             string        `json:"themeName"`
				SiteName              string        `json:"siteName"`
				SiteRootPrefix        string        `json:"siteRootPrefix"`
				IsLive                interface{}   `json:"isLive"`
				Description           string        `json:"description"`
				Keywords              string        `json:"keywords"`
				Header                string        `json:"header"`
				Footer                string        `json:"footer"`
				HideFromSearchEngines bool          `json:"hideFromSearchEngines"`
				ErrorPage             int           `json:"errorPage"`
				SiteURL               string        `json:"siteURL"`
				NoIndex               bool          `json:"noIndex"`
				NoFollow              bool          `json:"noFollow"`
				NoArchive             bool          `json:"noArchive"`
				NoSnippet             bool          `json:"noSnippet"`
				CollectionID          string        `json:"collectionId"`
				TargetID              string        `json:"targetId"`
				TargetAccessTokens    []interface{} `json:"targetAccessTokens"`
				IsCobrowseEnabled     bool          `json:"isCobrowseEnabled"`
				CobrowseID            string        `json:"cobrowseId"`
				SiteConnections       struct {
					OPAConnection  interface{} `json:"OPAConnection"`
					VBCSConnection string      `json:"VBCSConnection"`
				} `json:"siteConnections"`
				ConversationID      interface{} `json:"conversationId"`
				MapProvider         string      `json:"mapProvider"`
				MapAPIKey           string      `json:"mapAPIKey"`
				IsEnterprise        bool        `json:"isEnterprise"`
				RepositoryID        string      `json:"repositoryId"`
				ChannelID           string      `json:"channelId"`
				ChannelAccessTokens []struct {
					Name           string `json:"name"`
					Value          string `json:"value"`
					ExpirationDate string `json:"expirationDate"`
				} `json:"channelAccessTokens"`
				ArCollectionID             string   `json:"arCollectionId"`
				DefaultLanguage            string   `json:"defaultLanguage"`
				LocalizationPolicy         string   `json:"localizationPolicy"`
				AvailableLanguages         []string `json:"availableLanguages"`
				IsJSModuleEnabled          bool     `json:"isJSModuleEnabled"`
				IsAssetAnalyticsEnabled    bool     `json:"isAssetAnalyticsEnabled"`
				UseStandardAnalyticsScript bool     `json:"useStandardAnalyticsScript"`
				IsWebAnalyticsEnabled      bool     `json:"isWebAnalyticsEnabled"`
				WebAnalyticsScript         string   `json:"webAnalyticsScript"`
				CustomProperties           struct {
					OverwriteCounter string `json:"OverwriteCounter"`
				} `json:"customProperties"`
				LocaleAliases struct {
				} `json:"localeAliases"`
				PageProperties struct {
					ShowCallForExperts struct {
						Value string `json:"value"`
					} `json:"showCallForExperts"`
					GlobalNavUseLine struct {
						Value string `json:"value"`
					} `json:"globalNavUseLine"`
					ShowProjectMaterials struct {
						Value string `json:"value"`
					} `json:"showProjectMaterials"`
					GlobalNavStyle struct {
						Value string `json:"value"`
					} `json:"globalNavStyle"`
					GlobalNavColorStyle struct {
						Value string `json:"value"`
					} `json:"globalNavColorStyle"`
				} `json:"pageProperties"`
			} `json:"properties"`
		} `json:"base"`
		Variants                  interface{} `json:"variants"`
		DAllowedActions           interface{} `json:"dAllowedActions"`
		DefaultWebAnalyticsScript interface{} `json:"defaultWebAnalyticsScript"`
		HasExternalUserAccess     bool        `json:"hasExternalUserAccess"`
	} `json:"siteInfo"`
	Base struct {
		Pages []struct {
			ID               int         `json:"id"`
			Name             string      `json:"name"`
			ParentID         interface{} `json:"parentId"`
			PageURL          string      `json:"pageUrl"`
			HideInNavigation bool        `json:"hideInNavigation"`
			LinkURL          string      `json:"linkUrl"`
			LinkTarget       string      `json:"linkTarget"`
			Children         []int       `json:"children"`
			OverrideURL      bool        `json:"overrideUrl"`
			IsDetailPage     bool        `json:"isDetailPage"`
			IsSearchPage     bool        `json:"isSearchPage"`
			PageTemplateID   interface{} `json:"pageTemplateId"`
			Properties       struct {
				ShowPublicEvents string `json:"ShowPublicEvents"`
			} `json:"properties,omitempty"`
		} `json:"pages"`
	} `json:"base"`
	Variants interface{} `json:"variants"`
}

type page struct {
	Base struct {
		Properties struct {
			Title                 string      `json:"title"`
			PageLayout            string      `json:"pageLayout"`
			MobileLayout          string      `json:"mobileLayout"`
			PageDescription       string      `json:"pageDescription"`
			Keywords              string      `json:"keywords"`
			HideFromSearchEngines bool        `json:"hideFromSearchEngines"`
			Header                string      `json:"header"`
			Footer                string      `json:"footer"`
			NoIndex               bool        `json:"noIndex"`
			NoFollow              bool        `json:"noFollow"`
			NoArchive             bool        `json:"noArchive"`
			NoSnippet             bool        `json:"noSnippet"`
			IsCobrowseEnabled     bool        `json:"isCobrowseEnabled"`
			OverrideWebAnalytics  bool        `json:"overrideWebAnalytics"`
			WebAnalyticsScript    interface{} `json:"webAnalyticsScript"`
		} `json:"properties"`
		Slots struct {
			SlotContent struct {
				Components []string `json:"components"`
				Grid       string   `json:"grid"`
			} `json:"slot-content"`
		} `json:"slots"`
		ComponentInstances map[string]struct {
			Type string `json:"type"`
			ID   string `json:"id"`
			Data struct {
				UserText string `json:"userText"`
			} `json:"data"`
		} `json:"componentInstances"`
	} `json:"base"`
}

func init() {
	engine := crawlers.Register("1748", "国家研究委员会/美国国家科学院、工程院和医学院", "https://www.nationalacademies.org/")

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

	engine.OnEngineStart(func() {
		res, err := http.Get("https://www.nationalacademies.org/_cache_1b21/structure.json")
		if err != nil {
			crawlers.Sugar.Errorf("Unable to visit 1748 structure api: %s", err)
			return
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			crawlers.Sugar.Errorf("1748 structure api status code error: %d %s", res.StatusCode, res.Status)
			return
		}

		var data structure
		b, err := io.ReadAll(res.Body)
		if err != nil {
			crawlers.Sugar.Errorf("Unable to read json response from 1748 structure api: %s", err)
			return
		}
		err = json.Unmarshal(b, &data)
		if err != nil {
			crawlers.Sugar.Errorf("Unable parse json response from 1748 structure api: %s", err)
			return
		}
		for _, page := range data.Base.Pages {
			engine.Visit(fmt.Sprintf("https://www.nationalacademies.org/pages/%d.json", page.ID), crawlers.News)
		}
	})

	engine.OnResponse(func(response *colly.Response, ctx *crawlers.Context) {
		var p page
		err := json.Unmarshal(response.Body, &p)
		if err != nil {
			crawlers.Sugar.Errorf("Unable parse json response from %s: %s", response.Request.URL.String(), err)
			return
		}
		if p.Base.Properties.PageLayout != "detail_news.html" {
			ctx.PageType = crawlers.Index
			return
		}
		ctx.Keywords = strings.Split(p.Base.Properties.Keywords, ",")
		ctx.Title = p.Base.Properties.Title
		for _, v := range p.Base.ComponentInstances {
			if v.Type != "scs-paragraph" {
				continue
			}
			doc, err := goquery.NewDocumentFromReader(strings.NewReader(v.Data.UserText))
			if err != nil {
				crawlers.Sugar.Error(err)
				return
			}
			ctx.Content = doc.Text()
			break
		}
	})
}
