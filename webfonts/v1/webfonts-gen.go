// Package webfonts provides access to the .
//
// See http://code.google.com/apis/webfonts/docs/developer_api.html
//
// Usage example:
//
//   import "google-api-go-client.googlecode.com/hg/webfonts/v1"
//   ...
//   webfontsService, err := webfonts.New(oauthHttpClient)
package webfonts

import (
	"bytes"
	"fmt"
	"http"
	"io"
	"json"
	"os"
	"strings"
	"strconv"
	"url"
	"google-api-go-client.googlecode.com/hg/google-api"
)

var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version

const apiId = "webfonts:v1"
const apiName = "webfonts"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/webfonts/v1/"

func New(client *http.Client) (*Service, os.Error) {
	if client == nil {
		return nil, os.NewError("client is nil")
	}
	s := &Service{client: client}
	s.Webfonts = &WebfontsService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	Webfonts *WebfontsService
}

type WebfontsService struct {
	s *Service
}

type Webfont struct {
	// Subsets: The scripts supported by the font.
	Subsets interface{} `json:"subsets,omitempty"`

	Kind string `json:"kind,omitempty"`

	// Family: The name of the font.
	Family interface{} `json:"family,omitempty"`

	// Variants: The available variants for the font.
	Variants interface{} `json:"variants,omitempty"`
}

type WebfontList struct {
	// Items: The list of fonts currently served by the Google Fonts API.
	Items []*Webfont `json:"items,omitempty"`

	// Kind: The object kind.
	Kind string `json:"kind,omitempty"`
}

// method id "webfonts.webfonts.list":

type WebfontsListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Retrieves the list of fonts currently served by the Google Web
// Fonts Developer API
func (r *WebfontsService) List() *WebfontsListCall {
	c := &WebfontsListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// Sort sets the optional parameter "sort": Enables sorting of the list
func (c *WebfontsListCall) Sort(sort string) *WebfontsListCall {
	c.opt_["sort"] = sort
	return c
}

func (c *WebfontsListCall) Do() (*WebfontList, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["sort"]; ok {
		params.Set("sort", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/webfonts/v1/", "webfonts")
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
	ret := new(WebfontList)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the list of fonts currently served by the Google Web Fonts Developer API",
	//   "httpMethod": "GET",
	//   "id": "webfonts.webfonts.list",
	//   "parameters": {
	//     "sort": {
	//       "description": "Enables sorting of the list",
	//       "enum": [
	//         "alpha",
	//         "date",
	//         "popularity",
	//         "style",
	//         "trending"
	//       ],
	//       "enumDescriptions": [
	//         "Sort alphabetically",
	//         "Sort by date added",
	//         "Sort by popularity",
	//         "Sort by number of styles",
	//         "Sort by trending"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "webfonts",
	//   "response": {
	//     "$ref": "WebfontList"
	//   }
	// }

}

func cleanPathString(s string) string {
	return strings.Map(func(r int) int {
		if r >= 0x30 && r <= 0x7a {
			return r
		}
		return -1
	}, s)
}
