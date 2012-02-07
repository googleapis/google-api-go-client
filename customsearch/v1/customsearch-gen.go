// Package customsearch provides access to the CustomSearch API.
//
// See http://code.google.com/apis/customsearch/v1/using_rest.html
//
// Usage example:
//
//   import "code.google.com/p/google-api-go-client/customsearch/v1"
//   ...
//   customsearchService, err := customsearch.New(oauthHttpClient)
package customsearch

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

const apiId = "customsearch:v1"
const apiName = "customsearch"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/customsearch/"

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Cse = &CseService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	Cse *CseService
}

type CseService struct {
	s *Service
}

type SearchUrl struct {
	Template string `json:"template,omitempty"`

	Type string `json:"type,omitempty"`
}

type Query struct {
	Sort string `json:"sort,omitempty"`

	Cx string `json:"cx,omitempty"`

	Cr string `json:"cr,omitempty"`

	Count int64 `json:"count,omitempty"`

	StartIndex int64 `json:"startIndex,omitempty"`

	Language string `json:"language,omitempty"`

	TotalResults int64 `json:"totalResults,omitempty,string"`

	OutputEncoding string `json:"outputEncoding,omitempty"`

	Filter string `json:"filter,omitempty"`

	Safe string `json:"safe,omitempty"`

	StartPage int64 `json:"startPage,omitempty"`

	InputEncoding string `json:"inputEncoding,omitempty"`

	Cref string `json:"cref,omitempty"`

	GoogleHost string `json:"googleHost,omitempty"`

	Title string `json:"title,omitempty"`

	Gl string `json:"gl,omitempty"`

	SearchTerms string `json:"searchTerms,omitempty"`
}

type PromotionBodyLines struct {
	Title string `json:"title,omitempty"`

	Link string `json:"link,omitempty"`

	Url string `json:"url,omitempty"`
}

type Context struct {
	Title string `json:"title,omitempty"`

	Facets [][]*ContextFacetsItem `json:"facets,omitempty"`
}

type Promotion struct {
	BodyLines []*PromotionBodyLines `json:"bodyLines,omitempty"`

	DisplayLink string `json:"displayLink,omitempty"`

	Image *PromotionImage `json:"image,omitempty"`

	Link string `json:"link,omitempty"`

	Title string `json:"title,omitempty"`
}

type ResultPagemap struct {
}

type PromotionImage struct {
	Height int64 `json:"height,omitempty"`

	Width int64 `json:"width,omitempty"`

	Source string `json:"source,omitempty"`
}

type Result struct {
	HtmlTitle string `json:"htmlTitle,omitempty"`

	Title string `json:"title,omitempty"`

	HtmlSnippet string `json:"htmlSnippet,omitempty"`

	CacheId string `json:"cacheId,omitempty"`

	DisplayLink string `json:"displayLink,omitempty"`

	Snippet string `json:"snippet,omitempty"`

	Pagemap *ResultPagemap `json:"pagemap,omitempty"`

	Kind string `json:"kind,omitempty"`

	Link string `json:"link,omitempty"`
}

type Search struct {
	Items []*Result `json:"items,omitempty"`

	Context *Context `json:"context,omitempty"`

	Kind string `json:"kind,omitempty"`

	Url *SearchUrl `json:"url,omitempty"`

	Queries *SearchQueries `json:"queries,omitempty"`

	Promotions []*Promotion `json:"promotions,omitempty"`
}

type SearchQueries struct {
}

type ContextFacetsItem struct {
	Label string `json:"label,omitempty"`

	Anchor string `json:"anchor,omitempty"`
}

// method id "search.cse.list":

type CseListCall struct {
	s    *Service
	q    string
	opt_ map[string]interface{}
}

// List: Returns metadata about the search performed, metadata about the
// custom search engine used for the search, and the search results.
func (r *CseService) List(q string) *CseListCall {
	c := &CseListCall{s: r.s, opt_: make(map[string]interface{})}
	c.q = q
	return c
}

// Start sets the optional parameter "start": The index of the first
// result to return
func (c *CseListCall) Start(start string) *CseListCall {
	c.opt_["start"] = start
	return c
}

// Googlehost sets the optional parameter "googlehost": The local Google
// domain to use to perform the search.
func (c *CseListCall) Googlehost(googlehost string) *CseListCall {
	c.opt_["googlehost"] = googlehost
	return c
}

// Lr sets the optional parameter "lr": The language restriction for the
// search results
func (c *CseListCall) Lr(lr string) *CseListCall {
	c.opt_["lr"] = lr
	return c
}

// Num sets the optional parameter "num": Number of search results to
// return
func (c *CseListCall) Num(num string) *CseListCall {
	c.opt_["num"] = num
	return c
}

// Filter sets the optional parameter "filter": Controls turning on or
// off the duplicate content filter.
func (c *CseListCall) Filter(filter string) *CseListCall {
	c.opt_["filter"] = filter
	return c
}

// Safe sets the optional parameter "safe": Search safety level
func (c *CseListCall) Safe(safe string) *CseListCall {
	c.opt_["safe"] = safe
	return c
}

// Cref sets the optional parameter "cref": The URL of a linked custom
// search engine
func (c *CseListCall) Cref(cref string) *CseListCall {
	c.opt_["cref"] = cref
	return c
}

// Gl sets the optional parameter "gl": Geolocation of end user.
func (c *CseListCall) Gl(gl string) *CseListCall {
	c.opt_["gl"] = gl
	return c
}

// Sort sets the optional parameter "sort": The sort expression to apply
// to the results
func (c *CseListCall) Sort(sort string) *CseListCall {
	c.opt_["sort"] = sort
	return c
}

// Cx sets the optional parameter "cx": The custom search engine ID to
// scope this search query
func (c *CseListCall) Cx(cx string) *CseListCall {
	c.opt_["cx"] = cx
	return c
}

// Cr sets the optional parameter "cr": Country restrict(s).
func (c *CseListCall) Cr(cr string) *CseListCall {
	c.opt_["cr"] = cr
	return c
}

func (c *CseListCall) Do() (*Search, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("q", fmt.Sprintf("%v", c.q))
	if v, ok := c.opt_["start"]; ok {
		params.Set("start", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["googlehost"]; ok {
		params.Set("googlehost", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["lr"]; ok {
		params.Set("lr", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["num"]; ok {
		params.Set("num", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["filter"]; ok {
		params.Set("filter", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["safe"]; ok {
		params.Set("safe", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cref"]; ok {
		params.Set("cref", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gl"]; ok {
		params.Set("gl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["sort"]; ok {
		params.Set("sort", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cx"]; ok {
		params.Set("cx", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cr"]; ok {
		params.Set("cr", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/customsearch/", "v1")
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
	ret := new(Search)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns metadata about the search performed, metadata about the custom search engine used for the search, and the search results.",
	//   "httpMethod": "GET",
	//   "id": "search.cse.list",
	//   "parameterOrder": [
	//     "q"
	//   ],
	//   "parameters": {
	//     "cr": {
	//       "description": "Country restrict(s).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cref": {
	//       "description": "The URL of a linked custom search engine",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cx": {
	//       "description": "The custom search engine ID to scope this search query",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "filter": {
	//       "description": "Controls turning on or off the duplicate content filter.",
	//       "enum": [
	//         "0",
	//         "1"
	//       ],
	//       "enumDescriptions": [
	//         "Turns off duplicate content filter.",
	//         "Turns on duplicate content filter."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gl": {
	//       "description": "Geolocation of end user.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "googlehost": {
	//       "description": "The local Google domain to use to perform the search.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "lr": {
	//       "description": "The language restriction for the search results",
	//       "enum": [
	//         "lang_ar",
	//         "lang_bg",
	//         "lang_ca",
	//         "lang_cs",
	//         "lang_da",
	//         "lang_de",
	//         "lang_el",
	//         "lang_en",
	//         "lang_es",
	//         "lang_et",
	//         "lang_fi",
	//         "lang_fr",
	//         "lang_hr",
	//         "lang_hu",
	//         "lang_id",
	//         "lang_is",
	//         "lang_it",
	//         "lang_iw",
	//         "lang_ja",
	//         "lang_ko",
	//         "lang_lt",
	//         "lang_lv",
	//         "lang_nl",
	//         "lang_no",
	//         "lang_pl",
	//         "lang_pt",
	//         "lang_ro",
	//         "lang_ru",
	//         "lang_sk",
	//         "lang_sl",
	//         "lang_sr",
	//         "lang_sv",
	//         "lang_tr",
	//         "lang_zh-CN",
	//         "lang_zh-TW"
	//       ],
	//       "enumDescriptions": [
	//         "Arabic",
	//         "Bulgarian",
	//         "Catalan",
	//         "Czech",
	//         "Danish",
	//         "German",
	//         "Greek",
	//         "English",
	//         "Spanish",
	//         "Estonian",
	//         "Finnish",
	//         "French",
	//         "Croatian",
	//         "Hungarian",
	//         "Indonesian",
	//         "Icelandic",
	//         "Italian",
	//         "Hebrew",
	//         "Japanese",
	//         "Korean",
	//         "Lithuanian",
	//         "Latvian",
	//         "Dutch",
	//         "Norwegian",
	//         "Polish",
	//         "Portuguese",
	//         "Romanian",
	//         "Russian",
	//         "Slovak",
	//         "Slovenian",
	//         "Serbian",
	//         "Swedish",
	//         "Turkish",
	//         "Chinese (Simplified)",
	//         "Chinese (Traditional)"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "num": {
	//       "default": "10",
	//       "description": "Number of search results to return",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "q": {
	//       "description": "Query",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "safe": {
	//       "default": "off",
	//       "description": "Search safety level",
	//       "enum": [
	//         "high",
	//         "medium",
	//         "off"
	//       ],
	//       "enumDescriptions": [
	//         "Enables highest level of safe search filtering.",
	//         "Enables moderate safe search filtering.",
	//         "Disables safe search filtering."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sort": {
	//       "description": "The sort expression to apply to the results",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "start": {
	//       "description": "The index of the first result to return",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1",
	//   "response": {
	//     "$ref": "Search"
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
