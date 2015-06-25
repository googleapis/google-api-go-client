// Package webfonts provides access to the Google Fonts Developer API.
//
// See https://developers.google.com/fonts/docs/developer_api
//
// Usage example:
//
//   import "google.golang.org/api/webfonts/v1"
//   ...
//   webfontsService, err := webfonts.New(oauthHttpClient)
package webfonts

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Background

const apiId = "webfonts:v1"
const apiName = "webfonts"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/webfonts/v1/"

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Webfonts = NewWebfontsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Webfonts *WebfontsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewWebfontsService(s *Service) *WebfontsService {
	rs := &WebfontsService{s: s}
	return rs
}

type WebfontsService struct {
	s *Service
}

type Webfont struct {
	// Category: The category of the font.
	Category string `json:"category,omitempty"`

	// Family: The name of the font.
	Family string `json:"family,omitempty"`

	// Files: The font files (with all supported scripts) for each one of
	// the available variants, as a key : value map.
	Files map[string]string `json:"files,omitempty"`

	// Kind: This kind represents a webfont object in the webfonts service.
	Kind string `json:"kind,omitempty"`

	// LastModified: The date (format "yyyy-MM-dd") the font was modified
	// for the last time.
	LastModified string `json:"lastModified,omitempty"`

	// Subsets: The scripts supported by the font.
	Subsets []string `json:"subsets,omitempty"`

	// Variants: The available variants for the font.
	Variants []string `json:"variants,omitempty"`

	// Version: The font version.
	Version string `json:"version,omitempty"`
}

type WebfontList struct {
	// Items: The list of fonts currently served by the Google Fonts API.
	Items []*Webfont `json:"items,omitempty"`

	// Kind: This kind represents a list of webfont objects in the webfonts
	// service.
	Kind string `json:"kind,omitempty"`
}

// method id "webfonts.webfonts.list":

type WebfontsListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Retrieves the list of fonts currently served by the Google
// Fonts Developer API
func (r *WebfontsService) List() *WebfontsListCall {
	c := &WebfontsListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// Sort sets the optional parameter "sort": Enables sorting of the list
//
// Possible values:
//   "alpha" - Sort alphabetically
//   "date" - Sort by date added
//   "popularity" - Sort by popularity
//   "style" - Sort by number of styles
//   "trending" - Sort by trending
func (c *WebfontsListCall) Sort(sort string) *WebfontsListCall {
	c.opt_["sort"] = sort
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *WebfontsListCall) Fields(s ...googleapi.Field) *WebfontsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *WebfontsListCall) IfNoneMatch(entityTag string) *WebfontsListCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "webfonts.webfonts.list" call.
// Exactly one of the return values is non-nil.
func (c *WebfontsListCall) Do() (*WebfontList, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "webfonts.webfonts.list" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *WebfontsListCall) DoHeader() (ret *WebfontList, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["sort"]; ok {
		params.Set("sort", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "webfonts")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Retrieves the list of fonts currently served by the Google Fonts Developer API",
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
