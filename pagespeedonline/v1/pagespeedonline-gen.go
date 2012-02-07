// Package pagespeedonline provides access to the Page Speed Online API.
//
// See https://code.google.com/apis/pagespeedonline/v1/getting_started.html
//
// Usage example:
//
//   import "code.google.com/p/google-api-go-client/pagespeedonline/v1"
//   ...
//   pagespeedonlineService, err := pagespeedonline.New(oauthHttpClient)
package pagespeedonline

import (
	"bytes"
	"fmt"
	"net/http"
	"io"
	"encoding/json"
	"errors"
	"strings"
	"strconv"
	"net/url"
	"code.google.com/p/google-api-go-client/googleapi"
)

var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New

const apiId = "pagespeedonline:v1"
const apiName = "pagespeedonline"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/pagespeedonline/v1/"

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Pagespeedapi = &PagespeedapiService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	Pagespeedapi *PagespeedapiService
}

type PagespeedapiService struct {
	s *Service
}

type ResultFormattedResults struct {
	// Locale: The locale of the formattedResults, e.g. "en_US".
	Locale string `json:"locale,omitempty"`

	// RuleResults: Dictionary of formatted rule results, with one entry for
	// each Page Speed rule instantiated and run by the server.
	RuleResults *ResultFormattedResultsRuleResults `json:"ruleResults,omitempty"`
}

type Result struct {
	// Version: The version of the Page Speed SDK used to generate these
	// results.
	Version *ResultVersion `json:"version,omitempty"`

	// Kind: Kind of result.
	Kind string `json:"kind,omitempty"`

	// PageStats: Summary statistics for the page, such as number of
	// JavaScript bytes, number of HTML bytes, etc.
	PageStats *ResultPageStats `json:"pageStats,omitempty"`

	// Id: Canonicalized and final URL for the document, after following
	// page redirects (if any).
	Id string `json:"id,omitempty"`

	// Title: Title of the page, as displayed in the browser's title bar.
	Title string `json:"title,omitempty"`

	// InvalidRules: List of rules that were specified in the request, but
	// which the server did not know how to instantiate.
	InvalidRules []string `json:"invalidRules,omitempty"`

	// ResponseCode: Response code for the document. 200 indicates a normal
	// page load. 4xx/5xx indicates an error.
	ResponseCode int64 `json:"responseCode,omitempty"`

	// Score: The Page Speed Score (0-100), which indicates how much faster
	// a page could be. A high score indicates little room for improvement,
	// while a lower score indicates more room for improvement.
	Score int64 `json:"score,omitempty"`

	// FormattedResults: Localized Page Speed results. Contains a
	// ruleResults entry for each Page Speed rule instantiated and run by
	// the server.
	FormattedResults *ResultFormattedResults `json:"formattedResults,omitempty"`
}

type ResultVersion struct {
	// Minor: The minor version number of the Page Speed SDK used to
	// generate these results.
	Minor int64 `json:"minor,omitempty"`

	// Major: The major version number of the Page Speed SDK used to
	// generate these results.
	Major int64 `json:"major,omitempty"`
}

type ResultFormattedResultsRuleResults struct {
}

type ResultPageStats struct {
	// NumberResources: Number of HTTP resources loaded by the page.
	NumberResources int64 `json:"numberResources,omitempty"`

	// NumberHosts: Number of unique hosts referenced by the page.
	NumberHosts int64 `json:"numberHosts,omitempty"`

	// ImageResponseBytes: Number of response bytes for image resources on
	// the page.
	ImageResponseBytes int64 `json:"imageResponseBytes,omitempty,string"`

	// TextResponseBytes: Number of uncompressed response bytes for text
	// resources not covered by other statistics (i.e non-HTML, non-script,
	// non-CSS resources) on the page.
	TextResponseBytes int64 `json:"textResponseBytes,omitempty,string"`

	// OtherResponseBytes: Number of response bytes for other resources on
	// the page.
	OtherResponseBytes int64 `json:"otherResponseBytes,omitempty,string"`

	// NumberStaticResources: Number of static (i.e. cacheable) resources on
	// the page.
	NumberStaticResources int64 `json:"numberStaticResources,omitempty"`

	// NumberCssResources: Number of CSS resources referenced by the page.
	NumberCssResources int64 `json:"numberCssResources,omitempty"`

	// HtmlResponseBytes: Number of uncompressed response bytes for the main
	// HTML document and all iframes on the page.
	HtmlResponseBytes int64 `json:"htmlResponseBytes,omitempty,string"`

	// FlashResponseBytes: Number of response bytes for flash resources on
	// the page.
	FlashResponseBytes int64 `json:"flashResponseBytes,omitempty,string"`

	// CssResponseBytes: Number of uncompressed response bytes for CSS
	// resources on the page.
	CssResponseBytes int64 `json:"cssResponseBytes,omitempty,string"`

	// NumberJsResources: Number of JavaScript resources referenced by the
	// page.
	NumberJsResources int64 `json:"numberJsResources,omitempty"`

	// JavascriptResponseBytes: Number of uncompressed response bytes for JS
	// resources on the page.
	JavascriptResponseBytes int64 `json:"javascriptResponseBytes,omitempty,string"`

	// TotalRequestBytes: Total size of all request bytes sent by the page.
	TotalRequestBytes int64 `json:"totalRequestBytes,omitempty,string"`
}

// method id "pagespeedonline.pagespeedapi.runpagespeed":

type PagespeedapiRunpagespeedCall struct {
	s    *Service
	url  string
	opt_ map[string]interface{}
}

// Runpagespeed: Runs Page Speed analysis on the page at the specified
// URL, and returns a Page Speed score, a list of suggestions to make
// that page faster, and other information.
func (r *PagespeedapiService) Runpagespeed(url string) *PagespeedapiRunpagespeedCall {
	c := &PagespeedapiRunpagespeedCall{s: r.s, opt_: make(map[string]interface{})}
	c.url = url
	return c
}

// Strategy sets the optional parameter "strategy": The analysis
// strategy to use
func (c *PagespeedapiRunpagespeedCall) Strategy(strategy string) *PagespeedapiRunpagespeedCall {
	c.opt_["strategy"] = strategy
	return c
}

// Rule sets the optional parameter "rule": A Page Speed rule to run; if
// none are given, all rules are run
func (c *PagespeedapiRunpagespeedCall) Rule(rule string) *PagespeedapiRunpagespeedCall {
	c.opt_["rule"] = rule
	return c
}

// Locale sets the optional parameter "locale": The locale used to
// localize formatted results
func (c *PagespeedapiRunpagespeedCall) Locale(locale string) *PagespeedapiRunpagespeedCall {
	c.opt_["locale"] = locale
	return c
}

func (c *PagespeedapiRunpagespeedCall) Do() (*Result, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("url", fmt.Sprintf("%v", c.url))
	if v, ok := c.opt_["strategy"]; ok {
		params.Set("strategy", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["rule"]; ok {
		params.Set("rule", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["locale"]; ok {
		params.Set("locale", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/pagespeedonline/v1/", "runPagespeed")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Result)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Runs Page Speed analysis on the page at the specified URL, and returns a Page Speed score, a list of suggestions to make that page faster, and other information.",
	//   "httpMethod": "GET",
	//   "id": "pagespeedonline.pagespeedapi.runpagespeed",
	//   "parameterOrder": [
	//     "url"
	//   ],
	//   "parameters": {
	//     "locale": {
	//       "description": "The locale used to localize formatted results",
	//       "location": "query",
	//       "pattern": "[a-zA-Z]+(_[a-zA-Z]+)?",
	//       "type": "string"
	//     },
	//     "rule": {
	//       "description": "A Page Speed rule to run; if none are given, all rules are run",
	//       "location": "query",
	//       "pattern": "[a-zA-Z]+",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "strategy": {
	//       "description": "The analysis strategy to use",
	//       "enum": [
	//         "desktop",
	//         "mobile"
	//       ],
	//       "enumDescriptions": [
	//         "Fetch and analyze the URL for desktop browsers",
	//         "Fetch and analyze the URL for mobile devices"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "url": {
	//       "description": "The URL to fetch and analyze",
	//       "location": "query",
	//       "pattern": "http(s)?://.*",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "runPagespeed",
	//   "response": {
	//     "$ref": "Result"
	//   }
	// }

}

func cleanPathString(s string) string {
	return strings.Map(func(r rune) rune {
		if r >= 0x30 && r <= 0x7a {
			return r
		}
		return -1
	}, s)
}
