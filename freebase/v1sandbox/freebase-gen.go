// Package freebase provides access to the Freebase API.
//
// See https://developers.google.com/freebase/
//
// Usage example:
//
//   import "code.google.com/p/google-api-go-client/freebase/v1sandbox"
//   ...
//   freebaseService, err := freebase.New(oauthHttpClient)
package freebase

import (
	"bytes"
	"code.google.com/p/google-api-go-client/googleapi"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New

const apiId = "freebase:v1sandbox"
const apiName = "freebase"
const apiVersion = "v1sandbox"
const basePath = "https://www.googleapis.com/freebase/v1sandbox/"

// OAuth2 scopes used by this API.
const (
	// Sign in to Freebase with your account
	FreebaseScope = "https://www.googleapis.com/auth/freebase"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Text = NewTextService(s)
	s.Topic = NewTopicService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Text *TextService

	Topic *TopicService
}

func NewTextService(s *Service) *TextService {
	rs := &TextService{s: s}
	return rs
}

type TextService struct {
	s *Service
}

func NewTopicService(s *Service) *TopicService {
	rs := &TopicService{s: s}
	return rs
}

type TopicService struct {
	s *Service
}

type ContentserviceGet struct {
	// Result: The text requested.
	Result string `json:"result,omitempty"`
}

type TopicLookup struct {
	Id string `json:"id,omitempty"`

	Property *TopicLookupProperty `json:"property,omitempty"`
}

type TopicLookupProperty struct {
	FreebaseObject_profileLinkcount *TopicStatslinkcount `json:"/freebase/object_profile/linkcount,omitempty"`
}

type TopicPropertyvalue struct {
	Count float64 `json:"count,omitempty"`

	Status string `json:"status,omitempty"`

	Values []*TopicValue `json:"values,omitempty"`

	Valuetype string `json:"valuetype,omitempty"`
}

type TopicStatslinkcount struct {
	Type string `json:"type,omitempty"`

	Values []*TopicStatslinkcountValues `json:"values,omitempty"`
}

type TopicStatslinkcountValues struct {
	Count int64 `json:"count,omitempty"`

	Id string `json:"id,omitempty"`

	Values []*TopicStatslinkcountValuesValues `json:"values,omitempty"`
}

type TopicStatslinkcountValuesValues struct {
	Count int64 `json:"count,omitempty"`

	Id string `json:"id,omitempty"`

	Values []*TopicStatslinkcountValuesValuesValues `json:"values,omitempty"`
}

type TopicStatslinkcountValuesValuesValues struct {
	Count int64 `json:"count,omitempty"`

	Id string `json:"id,omitempty"`
}

type TopicValue struct {
	Citation *TopicValueCitation `json:"citation,omitempty"`

	Creator string `json:"creator,omitempty"`

	Dataset string `json:"dataset,omitempty"`

	Id string `json:"id,omitempty"`

	Lang string `json:"lang,omitempty"`

	Project string `json:"project,omitempty"`

	Property *TopicValueProperty `json:"property,omitempty"`

	Text string `json:"text,omitempty"`

	Timestamp string `json:"timestamp,omitempty"`

	Value interface{} `json:"value,omitempty"`
}

type TopicValueCitation struct {
	Provider string `json:"provider,omitempty"`

	Statement string `json:"statement,omitempty"`

	Uri string `json:"uri,omitempty"`
}

type TopicValueProperty struct {
}

// method id "freebase.image":

type ImageCall struct {
	s    *Service
	id   []string
	opt_ map[string]interface{}
}

// Image: Returns the scaled/cropped image attached to a freebase node.
func (s *Service) Image(id []string) *ImageCall {
	c := &ImageCall{s: s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

// Fallbackid sets the optional parameter "fallbackid": Use the image
// associated with this secondary id if no image is associated with the
// primary id.
func (c *ImageCall) Fallbackid(fallbackid string) *ImageCall {
	c.opt_["fallbackid"] = fallbackid
	return c
}

// Maxheight sets the optional parameter "maxheight": Maximum height in
// pixels for resulting image.
func (c *ImageCall) Maxheight(maxheight int64) *ImageCall {
	c.opt_["maxheight"] = maxheight
	return c
}

// Maxwidth sets the optional parameter "maxwidth": Maximum width in
// pixels for resulting image.
func (c *ImageCall) Maxwidth(maxwidth int64) *ImageCall {
	c.opt_["maxwidth"] = maxwidth
	return c
}

// Mode sets the optional parameter "mode": Method used to scale or crop
// image.
func (c *ImageCall) Mode(mode string) *ImageCall {
	c.opt_["mode"] = mode
	return c
}

// Pad sets the optional parameter "pad": A boolean specifying whether
// the resulting image should be padded up to the requested dimensions.
func (c *ImageCall) Pad(pad bool) *ImageCall {
	c.opt_["pad"] = pad
	return c
}

func (c *ImageCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fallbackid"]; ok {
		params.Set("fallbackid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxheight"]; ok {
		params.Set("maxheight", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxwidth"]; ok {
		params.Set("maxwidth", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["mode"]; ok {
		params.Set("mode", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pad"]; ok {
		params.Set("pad", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/freebase/v1sandbox/", "image{/id*}")
	urls = strings.Replace(urls, "{id}", cleanPathString(c.id[0]), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Returns the scaled/cropped image attached to a freebase node.",
	//   "httpMethod": "GET",
	//   "id": "freebase.image",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "fallbackid": {
	//       "default": "/freebase/no_image_png",
	//       "description": "Use the image associated with this secondary id if no image is associated with the primary id.",
	//       "location": "query",
	//       "pattern": "/[^.]*$",
	//       "type": "string"
	//     },
	//     "id": {
	//       "description": "Freebase entity or content id, mid, or guid.",
	//       "location": "path",
	//       "repeated": true,
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxheight": {
	//       "description": "Maximum height in pixels for resulting image.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "4096",
	//       "type": "integer"
	//     },
	//     "maxwidth": {
	//       "description": "Maximum width in pixels for resulting image.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "4096",
	//       "type": "integer"
	//     },
	//     "mode": {
	//       "default": "fit",
	//       "description": "Method used to scale or crop image.",
	//       "enum": [
	//         "fill",
	//         "fillcrop",
	//         "fillcropmid",
	//         "fit"
	//       ],
	//       "enumDescriptions": [
	//         "Fill rectangle completely with image, relax constraint on one dimension if necessary.",
	//         "Fill rectangle with image, crop image to maintain rectangle dimensions.",
	//         "Fill rectangle with image, center horizontally, crop left and right.",
	//         "Fit image inside rectangle, leave empty space in one dimension if necessary."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pad": {
	//       "default": "false",
	//       "description": "A boolean specifying whether the resulting image should be padded up to the requested dimensions.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "image{/id*}",
	//   "supportsMediaDownload": true
	// }

}

// method id "freebase.mqlread":

type MqlreadCall struct {
	s     *Service
	query string
	opt_  map[string]interface{}
}

// Mqlread: Performs MQL Queries.
func (s *Service) Mqlread(query string) *MqlreadCall {
	c := &MqlreadCall{s: s, opt_: make(map[string]interface{})}
	c.query = query
	return c
}

// As_of_time sets the optional parameter "as_of_time": Run the query as
// it would've been run at the specified point in time.
func (c *MqlreadCall) As_of_time(as_of_time string) *MqlreadCall {
	c.opt_["as_of_time"] = as_of_time
	return c
}

// Callback sets the optional parameter "callback": JS method name for
// JSONP callbacks.
func (c *MqlreadCall) Callback(callback string) *MqlreadCall {
	c.opt_["callback"] = callback
	return c
}

// Cost sets the optional parameter "cost": Show the costs or not.
func (c *MqlreadCall) Cost(cost bool) *MqlreadCall {
	c.opt_["cost"] = cost
	return c
}

// Cursor sets the optional parameter "cursor": The mql cursor.
func (c *MqlreadCall) Cursor(cursor string) *MqlreadCall {
	c.opt_["cursor"] = cursor
	return c
}

// Dateline sets the optional parameter "dateline": The dateline that
// you get in a mqlwrite response to ensure consistent results.
func (c *MqlreadCall) Dateline(dateline string) *MqlreadCall {
	c.opt_["dateline"] = dateline
	return c
}

// Html_escape sets the optional parameter "html_escape": Whether or not
// to escape entities.
func (c *MqlreadCall) Html_escape(html_escape bool) *MqlreadCall {
	c.opt_["html_escape"] = html_escape
	return c
}

// Indent sets the optional parameter "indent": How many spaces to
// indent the json.
func (c *MqlreadCall) Indent(indent int64) *MqlreadCall {
	c.opt_["indent"] = indent
	return c
}

// Lang sets the optional parameter "lang": The language of the results
// - an id of a /type/lang object.
func (c *MqlreadCall) Lang(lang string) *MqlreadCall {
	c.opt_["lang"] = lang
	return c
}

// Uniqueness_failure sets the optional parameter "uniqueness_failure":
// How MQL responds to uniqueness failures.
func (c *MqlreadCall) Uniqueness_failure(uniqueness_failure string) *MqlreadCall {
	c.opt_["uniqueness_failure"] = uniqueness_failure
	return c
}

func (c *MqlreadCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("query", fmt.Sprintf("%v", c.query))
	if v, ok := c.opt_["as_of_time"]; ok {
		params.Set("as_of_time", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["callback"]; ok {
		params.Set("callback", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cost"]; ok {
		params.Set("cost", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cursor"]; ok {
		params.Set("cursor", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["dateline"]; ok {
		params.Set("dateline", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["html_escape"]; ok {
		params.Set("html_escape", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["indent"]; ok {
		params.Set("indent", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["lang"]; ok {
		params.Set("lang", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["uniqueness_failure"]; ok {
		params.Set("uniqueness_failure", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/freebase/v1sandbox/", "mqlread")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Performs MQL Queries.",
	//   "httpMethod": "GET",
	//   "id": "freebase.mqlread",
	//   "parameterOrder": [
	//     "query"
	//   ],
	//   "parameters": {
	//     "as_of_time": {
	//       "description": "Run the query as it would've been run at the specified point in time.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "callback": {
	//       "description": "JS method name for JSONP callbacks.",
	//       "location": "query",
	//       "pattern": "([A-Za-z0-9_$.]|\\[|\\])+",
	//       "type": "string"
	//     },
	//     "cost": {
	//       "default": "false",
	//       "description": "Show the costs or not.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "cursor": {
	//       "description": "The mql cursor.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "dateline": {
	//       "description": "The dateline that you get in a mqlwrite response to ensure consistent results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "html_escape": {
	//       "default": "true",
	//       "description": "Whether or not to escape entities.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "indent": {
	//       "default": "0",
	//       "description": "How many spaces to indent the json.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "10",
	//       "type": "integer"
	//     },
	//     "lang": {
	//       "default": "/lang/en",
	//       "description": "The language of the results - an id of a /type/lang object.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "query": {
	//       "description": "An envelope containing a single MQL query.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "uniqueness_failure": {
	//       "default": "hard",
	//       "description": "How MQL responds to uniqueness failures.",
	//       "enum": [
	//         "hard",
	//         "soft"
	//       ],
	//       "enumDescriptions": [
	//         "Be strict - throw an error.",
	//         "Just return the first encountered object."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "mqlread",
	//   "supportsMediaDownload": true
	// }

}

// method id "freebase.mqlwrite":

type MqlwriteCall struct {
	s     *Service
	query string
	opt_  map[string]interface{}
}

// Mqlwrite: Performs MQL Write Operations.
func (s *Service) Mqlwrite(query string) *MqlwriteCall {
	c := &MqlwriteCall{s: s, opt_: make(map[string]interface{})}
	c.query = query
	return c
}

// Callback sets the optional parameter "callback": JS method name for
// JSONP callbacks.
func (c *MqlwriteCall) Callback(callback string) *MqlwriteCall {
	c.opt_["callback"] = callback
	return c
}

// Dateline sets the optional parameter "dateline": The dateline that
// you get in a mqlwrite response to ensure consistent results.
func (c *MqlwriteCall) Dateline(dateline string) *MqlwriteCall {
	c.opt_["dateline"] = dateline
	return c
}

// Indent sets the optional parameter "indent": How many spaces to
// indent the json.
func (c *MqlwriteCall) Indent(indent int64) *MqlwriteCall {
	c.opt_["indent"] = indent
	return c
}

// Use_permission_of sets the optional parameter "use_permission_of":
// Use the same permission node of the object with the specified id.
func (c *MqlwriteCall) Use_permission_of(use_permission_of string) *MqlwriteCall {
	c.opt_["use_permission_of"] = use_permission_of
	return c
}

func (c *MqlwriteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("query", fmt.Sprintf("%v", c.query))
	if v, ok := c.opt_["callback"]; ok {
		params.Set("callback", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["dateline"]; ok {
		params.Set("dateline", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["indent"]; ok {
		params.Set("indent", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["use_permission_of"]; ok {
		params.Set("use_permission_of", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/freebase/v1sandbox/", "mqlwrite")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Performs MQL Write Operations.",
	//   "httpMethod": "GET",
	//   "id": "freebase.mqlwrite",
	//   "parameterOrder": [
	//     "query"
	//   ],
	//   "parameters": {
	//     "callback": {
	//       "description": "JS method name for JSONP callbacks.",
	//       "location": "query",
	//       "pattern": "([A-Za-z0-9_$.]|\\[|\\])+",
	//       "type": "string"
	//     },
	//     "dateline": {
	//       "description": "The dateline that you get in a mqlwrite response to ensure consistent results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "indent": {
	//       "default": "0",
	//       "description": "How many spaces to indent the json.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "10",
	//       "type": "integer"
	//     },
	//     "query": {
	//       "description": "An MQL query with write directives.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "use_permission_of": {
	//       "description": "Use the same permission node of the object with the specified id.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "mqlwrite",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/freebase"
	//   ],
	//   "supportsMediaDownload": true
	// }

}

// method id "freebase.text.get":

type TextGetCall struct {
	s    *Service
	id   []string
	opt_ map[string]interface{}
}

// Get: Returns blob attached to node at specified id as HTML
func (r *TextService) Get(id []string) *TextGetCall {
	c := &TextGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

// Format sets the optional parameter "format": Sanitizing
// transformation.
func (c *TextGetCall) Format(format string) *TextGetCall {
	c.opt_["format"] = format
	return c
}

// Maxlength sets the optional parameter "maxlength": The max number of
// characters to return. Valid only for 'plain' format.
func (c *TextGetCall) Maxlength(maxlength int64) *TextGetCall {
	c.opt_["maxlength"] = maxlength
	return c
}

func (c *TextGetCall) Do() (*ContentserviceGet, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["format"]; ok {
		params.Set("format", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxlength"]; ok {
		params.Set("maxlength", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/freebase/v1sandbox/", "text{/id*}")
	urls = strings.Replace(urls, "{id}", cleanPathString(c.id[0]), 1)
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
	ret := new(ContentserviceGet)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns blob attached to node at specified id as HTML",
	//   "httpMethod": "GET",
	//   "id": "freebase.text.get",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "format": {
	//       "default": "plain",
	//       "description": "Sanitizing transformation.",
	//       "enum": [
	//         "html",
	//         "plain",
	//         "raw"
	//       ],
	//       "enumDescriptions": [
	//         "Return valid, sanitized html.",
	//         "Return plain text - strip html tags.",
	//         "Return the entire content as-is."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "id": {
	//       "description": "The id of the item that you want data about",
	//       "location": "path",
	//       "repeated": true,
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxlength": {
	//       "description": "The max number of characters to return. Valid only for 'plain' format.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "text{/id*}",
	//   "response": {
	//     "$ref": "ContentserviceGet"
	//   }
	// }

}

// method id "freebase.topic.lookup":

type TopicLookupCall struct {
	s    *Service
	id   []string
	opt_ map[string]interface{}
}

// Lookup: Get properties and meta-data about a topic.
func (r *TopicService) Lookup(id []string) *TopicLookupCall {
	c := &TopicLookupCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

// Dateline sets the optional parameter "dateline": Determines how
// up-to-date the data returned is. A unix epoch time, a guid or a 'now'
func (c *TopicLookupCall) Dateline(dateline string) *TopicLookupCall {
	c.opt_["dateline"] = dateline
	return c
}

// Filter sets the optional parameter "filter": A frebase domain, type
// or property id, 'suggest', 'commons', or 'all'. Filter the results
// and returns only appropriate properties.
func (c *TopicLookupCall) Filter(filter string) *TopicLookupCall {
	c.opt_["filter"] = filter
	return c
}

// Lang sets the optional parameter "lang": The language you 'd like the
// content in - a freebase /type/lang language key.
func (c *TopicLookupCall) Lang(lang string) *TopicLookupCall {
	c.opt_["lang"] = lang
	return c
}

// Limit sets the optional parameter "limit": The maximum number of
// property values to return for each property.
func (c *TopicLookupCall) Limit(limit int64) *TopicLookupCall {
	c.opt_["limit"] = limit
	return c
}

// Raw sets the optional parameter "raw": Do not apply any constraints,
// or get any names.
func (c *TopicLookupCall) Raw(raw bool) *TopicLookupCall {
	c.opt_["raw"] = raw
	return c
}

func (c *TopicLookupCall) Do() (*TopicLookup, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["dateline"]; ok {
		params.Set("dateline", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["filter"]; ok {
		params.Set("filter", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["lang"]; ok {
		params.Set("lang", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["limit"]; ok {
		params.Set("limit", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["raw"]; ok {
		params.Set("raw", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/freebase/v1sandbox/", "topic{/id*}")
	urls = strings.Replace(urls, "{id}", cleanPathString(c.id[0]), 1)
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
	ret := new(TopicLookup)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get properties and meta-data about a topic.",
	//   "httpMethod": "GET",
	//   "id": "freebase.topic.lookup",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "dateline": {
	//       "description": "Determines how up-to-date the data returned is. A unix epoch time, a guid or a 'now'",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "filter": {
	//       "description": "A frebase domain, type or property id, 'suggest', 'commons', or 'all'. Filter the results and returns only appropriate properties.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "id": {
	//       "description": "The id of the item that you want data about.",
	//       "location": "path",
	//       "repeated": true,
	//       "required": true,
	//       "type": "string"
	//     },
	//     "lang": {
	//       "default": "en",
	//       "description": "The language you 'd like the content in - a freebase /type/lang language key.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "limit": {
	//       "default": "10",
	//       "description": "The maximum number of property values to return for each property.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "raw": {
	//       "default": "false",
	//       "description": "Do not apply any constraints, or get any names.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "topic{/id*}",
	//   "response": {
	//     "$ref": "TopicLookup"
	//   }
	// }

}

func cleanPathString(s string) string {
	return strings.Map(func(r rune) rune {
		if r >= 0x2d && r <= 0x7a || r == '~' {
			return r
		}
		return -1
	}, s)
}
