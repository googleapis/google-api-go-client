// Package freebase provides access to the Freebase API.
//
// See http://wiki.freebase.com/wiki/API
//
// Usage example:
//
//   import "code.google.com/p/google-api-go-client/freebase/v1-sandbox"
//   ...
//   freebaseService, err := freebase.New(oauthHttpClient)
package freebase

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

const apiId = "freebase:v1-sandbox"
const apiName = "freebase"
const apiVersion = "v1-sandbox"
const basePath = "https://www.googleapis.com/freebase/v1-sandbox/"

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Text = &TextService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	Text *TextService
}

type TextService struct {
	s *Service
}

type ContentserviceGet struct {
	// Result: The text requested.
	Result string `json:"result,omitempty"`
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
	urls := googleapi.ResolveRelative("https://www.googleapis.com/freebase/v1-sandbox/", "mqlread")
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
	//   "path": "mqlread"
	// }

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
	urls := googleapi.ResolveRelative("https://www.googleapis.com/freebase/v1-sandbox/", "image{/id*}")
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
	//         "TODO(bendrees)",
	//         "TODO(bendrees)",
	//         "TODO(bendrees)",
	//         "TODO(bendrees)"
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
	//   "path": "image{/id*}"
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
	urls := googleapi.ResolveRelative("https://www.googleapis.com/freebase/v1-sandbox/", "text{/id*}")
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

func cleanPathString(s string) string {
	return strings.Map(func(r rune) rune {
		if r >= 0x30 && r <= 0x7a {
			return r
		}
		return -1
	}, s)
}
