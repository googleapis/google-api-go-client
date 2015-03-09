// Package cloudlatencytest provides access to the Google Cloud Network Performance Monitoring API.
//
// Usage example:
//
//   import "google.golang.org/api/cloudlatencytest/v2"
//   ...
//   cloudlatencytestService, err := cloudlatencytest.New(oauthHttpClient)
package cloudlatencytest

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

const apiId = "cloudlatencytest:v2"
const apiName = "cloudlatencytest"
const apiVersion = "v2"
const basePath = "https://cloudlatencytest-pa.googleapis.com/v2/statscollection/"

// OAuth2 scopes used by this API.
const (
	// View monitoring data for all of your Google Cloud and API projects
	MonitoringReadonlyScope = "https://www.googleapis.com/auth/monitoring.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Statscollection = NewStatscollectionService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Statscollection *StatscollectionService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewStatscollectionService(s *Service) *StatscollectionService {
	rs := &StatscollectionService{s: s}
	return rs
}

type StatscollectionService struct {
	s *Service
}

type AggregatedStats struct {
	Stats []*Stats `json:"stats,omitempty"`
}

type AggregatedStatsReply struct {
	TestValue string `json:"testValue,omitempty"`
}

type DoubleValue struct {
	Label string `json:"label,omitempty"`

	Value float64 `json:"value,omitempty"`
}

type IntValue struct {
	Label string `json:"label,omitempty"`

	Value int64 `json:"value,omitempty,string"`
}

type Stats struct {
	DoubleValues []*DoubleValue `json:"doubleValues,omitempty"`

	IntValues []*IntValue `json:"intValues,omitempty"`

	StringValues []*StringValue `json:"stringValues,omitempty"`

	Time float64 `json:"time,omitempty"`
}

type StatsReply struct {
	TestValue string `json:"testValue,omitempty"`
}

type StringValue struct {
	Label string `json:"label,omitempty"`

	Value string `json:"value,omitempty"`
}

// method id "cloudlatencytest.statscollection.updateaggregatedstats":

type StatscollectionUpdateaggregatedstatsCall struct {
	s               *Service
	aggregatedstats *AggregatedStats
	opt_            map[string]interface{}
}

// Updateaggregatedstats: RPC to update the new TCP stats.
func (r *StatscollectionService) Updateaggregatedstats(aggregatedstats *AggregatedStats) *StatscollectionUpdateaggregatedstatsCall {
	c := &StatscollectionUpdateaggregatedstatsCall{s: r.s, opt_: make(map[string]interface{})}
	c.aggregatedstats = aggregatedstats
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *StatscollectionUpdateaggregatedstatsCall) Fields(s ...googleapi.Field) *StatscollectionUpdateaggregatedstatsCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *StatscollectionUpdateaggregatedstatsCall) Do() (*AggregatedStatsReply, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.aggregatedstats)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "updateaggregatedstats")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *AggregatedStatsReply
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "RPC to update the new TCP stats.",
	//   "httpMethod": "POST",
	//   "id": "cloudlatencytest.statscollection.updateaggregatedstats",
	//   "path": "updateaggregatedstats",
	//   "request": {
	//     "$ref": "AggregatedStats"
	//   },
	//   "response": {
	//     "$ref": "AggregatedStatsReply"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/monitoring.readonly"
	//   ]
	// }

}

// method id "cloudlatencytest.statscollection.updatestats":

type StatscollectionUpdatestatsCall struct {
	s     *Service
	stats *Stats
	opt_  map[string]interface{}
}

// Updatestats: RPC to update the new TCP stats.
func (r *StatscollectionService) Updatestats(stats *Stats) *StatscollectionUpdatestatsCall {
	c := &StatscollectionUpdatestatsCall{s: r.s, opt_: make(map[string]interface{})}
	c.stats = stats
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *StatscollectionUpdatestatsCall) Fields(s ...googleapi.Field) *StatscollectionUpdatestatsCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *StatscollectionUpdatestatsCall) Do() (*StatsReply, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.stats)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "updatestats")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *StatsReply
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "RPC to update the new TCP stats.",
	//   "httpMethod": "POST",
	//   "id": "cloudlatencytest.statscollection.updatestats",
	//   "path": "updatestats",
	//   "request": {
	//     "$ref": "Stats"
	//   },
	//   "response": {
	//     "$ref": "StatsReply"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/monitoring.readonly"
	//   ]
	// }

}
