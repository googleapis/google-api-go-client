// Package analytics provides access to the Google Analytics API.
//
// See https://developers.google.com/analytics/
//
// Usage example:
//
//   import "code.google.com/p/google-api-go-client/analytics/v2.4"
//   ...
//   analyticsService, err := analytics.New(oauthHttpClient)
package analytics

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

const apiId = "analytics:v2.4"
const apiName = "analytics"
const apiVersion = "v2.4"
const basePath = "https://www.googleapis.com/analytics/v2.4/"

// OAuth2 scopes used by this API.
const (
	// View and manage your Google Analytics data
	AnalyticsScope = "https://www.googleapis.com/auth/analytics"

	// View your Google Analytics data
	AnalyticsReadonlyScope = "https://www.googleapis.com/auth/analytics.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Data = &DataService{s: s}
	s.Management = &ManagementService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	Data *DataService

	Management *ManagementService
}

type DataService struct {
	s *Service
}

type ManagementService struct {
	s *Service
}

// method id "analytics.data.get":

type DataGetCall struct {
	s         *Service
	ids       string
	startDate string
	endDate   string
	metrics   string
	opt_      map[string]interface{}
}

// Get: Returns Analytics report data for a profile.
func (r *DataService) Get(ids string, startDate string, endDate string, metrics string) *DataGetCall {
	c := &DataGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.ids = ids
	c.startDate = startDate
	c.endDate = endDate
	c.metrics = metrics
	return c
}

// Dimensions sets the optional parameter "dimensions": A
// comma-separated list of Analytics dimensions. E.g.,
// 'ga:browser,ga:city'.
func (c *DataGetCall) Dimensions(dimensions string) *DataGetCall {
	c.opt_["dimensions"] = dimensions
	return c
}

// Filters sets the optional parameter "filters": A comma-separated list
// of dimension or metric filters to be applied to the report data.
func (c *DataGetCall) Filters(filters string) *DataGetCall {
	c.opt_["filters"] = filters
	return c
}

// MaxResults sets the optional parameter "max-results": The maximum
// number of entries to include in this feed.
func (c *DataGetCall) MaxResults(maxResults int64) *DataGetCall {
	c.opt_["max-results"] = maxResults
	return c
}

// Segment sets the optional parameter "segment": An Analytics advanced
// segment to be applied to the report data.
func (c *DataGetCall) Segment(segment string) *DataGetCall {
	c.opt_["segment"] = segment
	return c
}

// Sort sets the optional parameter "sort": A comma-separated list of
// dimensions or metrics that determine the sort order for the report
// data.
func (c *DataGetCall) Sort(sort string) *DataGetCall {
	c.opt_["sort"] = sort
	return c
}

// StartIndex sets the optional parameter "start-index": An index of the
// first entity to retrieve. Use this parameter as a pagination
// mechanism along with the max-results parameter.
func (c *DataGetCall) StartIndex(startIndex int64) *DataGetCall {
	c.opt_["start-index"] = startIndex
	return c
}

func (c *DataGetCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("end-date", fmt.Sprintf("%v", c.endDate))
	params.Set("ids", fmt.Sprintf("%v", c.ids))
	params.Set("metrics", fmt.Sprintf("%v", c.metrics))
	params.Set("start-date", fmt.Sprintf("%v", c.startDate))
	if v, ok := c.opt_["dimensions"]; ok {
		params.Set("dimensions", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["filters"]; ok {
		params.Set("filters", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["max-results"]; ok {
		params.Set("max-results", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["segment"]; ok {
		params.Set("segment", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["sort"]; ok {
		params.Set("sort", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["start-index"]; ok {
		params.Set("start-index", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/analytics/v2.4/", "data")
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
	//   "description": "Returns Analytics report data for a profile.",
	//   "httpMethod": "GET",
	//   "id": "analytics.data.get",
	//   "parameterOrder": [
	//     "ids",
	//     "start-date",
	//     "end-date",
	//     "metrics"
	//   ],
	//   "parameters": {
	//     "dimensions": {
	//       "description": "A comma-separated list of Analytics dimensions. E.g., 'ga:browser,ga:city'.",
	//       "location": "query",
	//       "pattern": "(ga:.+)?",
	//       "type": "string"
	//     },
	//     "end-date": {
	//       "description": "End date for fetching report data. All requests should specify an end date formatted as YYYY-MM-DD.",
	//       "location": "query",
	//       "pattern": "[0-9]{4}-[0-9]{2}-[0-9]{2}",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "filters": {
	//       "description": "A comma-separated list of dimension or metric filters to be applied to the report data.",
	//       "location": "query",
	//       "pattern": "ga:.+",
	//       "type": "string"
	//     },
	//     "ids": {
	//       "description": "Unique table ID for retrieving report data. Table ID is of the form ga:XXXX, where XXXX is the Analytics profile ID.",
	//       "location": "query",
	//       "pattern": "ga:[0-9]+",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "max-results": {
	//       "description": "The maximum number of entries to include in this feed.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "metrics": {
	//       "description": "A comma-separated list of Analytics metrics. E.g., 'ga:visits,ga:pageviews'. At least one metric must be specified to retrieve a valid Analytics report.",
	//       "location": "query",
	//       "pattern": "ga:.+",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "segment": {
	//       "description": "An Analytics advanced segment to be applied to the report data.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sort": {
	//       "description": "A comma-separated list of dimensions or metrics that determine the sort order for the report data.",
	//       "location": "query",
	//       "pattern": "(-)?ga:.+",
	//       "type": "string"
	//     },
	//     "start-date": {
	//       "description": "Start date for fetching report data. All requests should specify a start date formatted as YYYY-MM-DD.",
	//       "location": "query",
	//       "pattern": "[0-9]{4}-[0-9]{2}-[0-9]{2}",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "start-index": {
	//       "description": "An index of the first entity to retrieve. Use this parameter as a pagination mechanism along with the max-results parameter.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "data",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/analytics",
	//     "https://www.googleapis.com/auth/analytics.readonly"
	//   ]
	// }

}

func cleanPathString(s string) string {
	return strings.Map(func(r rune) rune {
		if r >= 0x2d && r <= 0x7a {
			return r
		}
		return -1
	}, s)
}
