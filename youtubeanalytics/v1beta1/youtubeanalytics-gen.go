// Package youtubeanalytics provides access to the YouTube Analytics API.
//
// See http://developers.google.com/youtube/analytics/
//
// Usage example:
//
//   import "code.google.com/p/google-api-go-client/youtubeanalytics/v1beta1"
//   ...
//   youtubeanalyticsService, err := youtubeanalytics.New(oauthHttpClient)
package youtubeanalytics

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

const apiId = "youtubeAnalytics:v1beta1"
const apiName = "youtubeAnalytics"
const apiVersion = "v1beta1"
const basePath = "https://www.googleapis.com/youtube/analytics/v1beta1/"

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Reports = &ReportsService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	Reports *ReportsService
}

type ReportsService struct {
	s *Service
}

type ResultTable struct {
	// ColumnHeaders: Contains information about the columns returned in the
	// "rows" fields. The order of the elements matches the order of the
	// corresponding columns in "rows" field.
	ColumnHeaders []*ResultTableColumnHeaders `json:"columnHeaders,omitempty"`

	// Kind: Identifier used to mark the structure as a result table.
	Kind string `json:"kind,omitempty"`

	// Rows: Contains all rows of the result table. Each row contains an
	// array with the values for the columns. The order matches the order of
	// the column information provided in the "columnHeaders" field. If no
	// data is available for the given query, the "rows" element will be
	// omitted from the response. The response for a query with the day
	// dimension will not contain rows for the most recent days.
	Rows [][]interface{} `json:"rows,omitempty"`
}

type ResultTableColumnHeaders struct {
	// ColumnType: The type of the column (DIMENSION, METRIC).
	ColumnType string `json:"columnType,omitempty"`

	// DataType: Type of the data in the column (STRING, INTEGER, FLOAT).
	DataType string `json:"dataType,omitempty"`

	// Name: The name of the dimension or metric.
	Name string `json:"name,omitempty"`
}

// method id "youtubeAnalytics.reports.query":

type ReportsQueryCall struct {
	s         *Service
	ids       string
	startDate string
	endDate   string
	metrics   string
	opt_      map[string]interface{}
}

// Query: Retrieve your YouTube Analytics reports.
func (r *ReportsService) Query(ids string, startDate string, endDate string, metrics string) *ReportsQueryCall {
	c := &ReportsQueryCall{s: r.s, opt_: make(map[string]interface{})}
	c.ids = ids
	c.startDate = startDate
	c.endDate = endDate
	c.metrics = metrics
	return c
}

// Dimensions sets the optional parameter "dimensions": A
// comma-separated list of YouTube Analytics dimensions. E.g., 'video',
// or 'ageGroup,gender'.
func (c *ReportsQueryCall) Dimensions(dimensions string) *ReportsQueryCall {
	c.opt_["dimensions"] = dimensions
	return c
}

// Filters sets the optional parameter "filters": A list of dimension
// filters to be applied to YouTube Analytics data. Multiple filters can
// be joined together with the ';' character. The returned result table
// will satisfy both filters. E.g., video==dMH0bHeiRNg;country==IT will
// restrict the returned stats to the given video and the country Italy.
func (c *ReportsQueryCall) Filters(filters string) *ReportsQueryCall {
	c.opt_["filters"] = filters
	return c
}

// MaxResults sets the optional parameter "max-results": The maximum
// number of rows to include in the response.
func (c *ReportsQueryCall) MaxResults(maxResults int64) *ReportsQueryCall {
	c.opt_["max-results"] = maxResults
	return c
}

// Sort sets the optional parameter "sort": A comma-separated list of
// dimensions or metrics that determine the sort order for YouTube
// Analytics data. By default the sort order is ascending, '-' prefix
// causes descending sort order.
func (c *ReportsQueryCall) Sort(sort string) *ReportsQueryCall {
	c.opt_["sort"] = sort
	return c
}

// StartIndex sets the optional parameter "start-index": An index of the
// first entity to retrieve. Use this parameter as a pagination
// mechanism along with the max-results parameter (one-based,
// inclusive).
func (c *ReportsQueryCall) StartIndex(startIndex int64) *ReportsQueryCall {
	c.opt_["start-index"] = startIndex
	return c
}

func (c *ReportsQueryCall) Do() (*ResultTable, error) {
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
	if v, ok := c.opt_["sort"]; ok {
		params.Set("sort", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["start-index"]; ok {
		params.Set("start-index", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/analytics/v1beta1/", "reports")
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
	ret := new(ResultTable)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieve your YouTube Analytics reports.",
	//   "httpMethod": "GET",
	//   "id": "youtubeAnalytics.reports.query",
	//   "parameterOrder": [
	//     "ids",
	//     "start-date",
	//     "end-date",
	//     "metrics"
	//   ],
	//   "parameters": {
	//     "dimensions": {
	//       "description": "A comma-separated list of YouTube Analytics dimensions. E.g., 'video', or 'ageGroup,gender'.",
	//       "location": "query",
	//       "pattern": "[a-zA-Z,]+",
	//       "type": "string"
	//     },
	//     "end-date": {
	//       "description": "End date for fetching YouTube Analytics data. All requests should specify an end date formatted as YYYY-MM-DD.",
	//       "location": "query",
	//       "pattern": "[0-9]{4}-[0-9]{2}-[0-9]{2}",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "filters": {
	//       "description": "A list of dimension filters to be applied to YouTube Analytics data. Multiple filters can be joined together with the ';' character. The returned result table will satisfy both filters. E.g., video==dMH0bHeiRNg;country==IT will restrict the returned stats to the given video and the country Italy.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ids": {
	//       "description": "Unique channel or content owner ID for retrieving YouTube Analytics data. Either channel==C or contentOwner==O where 'C' is the encrypted channel ID and 'O' is the content owner name.",
	//       "location": "query",
	//       "pattern": "(channel==[a-zA-Z0-9_-]+|contentOwner==[a-zA-Z0-9_+-]+)",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "max-results": {
	//       "description": "The maximum number of rows to include in the response.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "metrics": {
	//       "description": "A comma-separated list of YouTube Analytics metrics. E.g., 'views' or 'likes,dislikes'",
	//       "location": "query",
	//       "pattern": "[a-zA-Z,]+",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sort": {
	//       "description": "A comma-separated list of dimensions or metrics that determine the sort order for YouTube Analytics data. By default the sort order is ascending, '-' prefix causes descending sort order.",
	//       "location": "query",
	//       "pattern": "(-)?[a-zA-Z,]+",
	//       "type": "string"
	//     },
	//     "start-date": {
	//       "description": "Start date for fetching YouTube Analytics data. All requests should specify a start date formatted as YYYY-MM-DD.",
	//       "location": "query",
	//       "pattern": "[0-9]{4}-[0-9]{2}-[0-9]{2}",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "start-index": {
	//       "description": "An index of the first entity to retrieve. Use this parameter as a pagination mechanism along with the max-results parameter (one-based, inclusive).",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "reports",
	//   "response": {
	//     "$ref": "ResultTable"
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
