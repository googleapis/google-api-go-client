// Package youtubeanalytics provides access to the YouTube Analytics API.
//
// See http://developers.google.com/youtube/analytics/
//
// Usage example:
//
//   import "google.golang.org/api/youtubeanalytics/v1beta1"
//   ...
//   youtubeanalyticsService, err := youtubeanalytics.New(oauthHttpClient)
package youtubeanalytics

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

const apiId = "youtubeAnalytics:v1beta1"
const apiName = "youtubeAnalytics"
const apiVersion = "v1beta1"
const basePath = "https://www.googleapis.com/youtube/analytics/v1beta1/"

// OAuth2 scopes used by this API.
const (
	// Manage your YouTube account
	YoutubeScope = "https://www.googleapis.com/auth/youtube"

	// View your YouTube account
	YoutubeReadonlyScope = "https://www.googleapis.com/auth/youtube.readonly"

	// View and manage your assets and associated content on YouTube
	YoutubepartnerScope = "https://www.googleapis.com/auth/youtubepartner"

	// View monetary and non-monetary YouTube Analytics reports for your
	// YouTube content
	YtAnalyticsMonetaryReadonlyScope = "https://www.googleapis.com/auth/yt-analytics-monetary.readonly"

	// View YouTube Analytics reports for your YouTube content
	YtAnalyticsReadonlyScope = "https://www.googleapis.com/auth/yt-analytics.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.BatchReportDefinitions = NewBatchReportDefinitionsService(s)
	s.BatchReports = NewBatchReportsService(s)
	s.GroupItems = NewGroupItemsService(s)
	s.Groups = NewGroupsService(s)
	s.Reports = NewReportsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	BatchReportDefinitions *BatchReportDefinitionsService

	BatchReports *BatchReportsService

	GroupItems *GroupItemsService

	Groups *GroupsService

	Reports *ReportsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewBatchReportDefinitionsService(s *Service) *BatchReportDefinitionsService {
	rs := &BatchReportDefinitionsService{s: s}
	return rs
}

type BatchReportDefinitionsService struct {
	s *Service
}

func NewBatchReportsService(s *Service) *BatchReportsService {
	rs := &BatchReportsService{s: s}
	return rs
}

type BatchReportsService struct {
	s *Service
}

func NewGroupItemsService(s *Service) *GroupItemsService {
	rs := &GroupItemsService{s: s}
	return rs
}

type GroupItemsService struct {
	s *Service
}

func NewGroupsService(s *Service) *GroupsService {
	rs := &GroupsService{s: s}
	return rs
}

type GroupsService struct {
	s *Service
}

func NewReportsService(s *Service) *ReportsService {
	rs := &ReportsService{s: s}
	return rs
}

type ReportsService struct {
	s *Service
}

// BatchReport: Contains single batchReport resource.
type BatchReport struct {
	// Id: The ID that YouTube assigns and uses to uniquely identify the
	// report.
	Id string `json:"id,omitempty"`

	// Kind: This value specifies the type of data of this item. For batch
	// report the kind property value is youtubeAnalytics#batchReport.
	Kind string `json:"kind,omitempty"`

	// Outputs: Report outputs.
	Outputs []*BatchReportOutputs `json:"outputs,omitempty"`

	// ReportId: The ID of the the report definition.
	ReportId string `json:"reportId,omitempty"`

	// TimeSpan: Period included in the report. For reports containing all
	// entities endTime is not set. Both startTime and endTime are
	// inclusive.
	TimeSpan *BatchReportTimeSpan `json:"timeSpan,omitempty"`

	// TimeUpdated: The time when the report was updated.
	TimeUpdated string `json:"timeUpdated,omitempty"`
}

type BatchReportOutputs struct {
	// DownloadUrl: Cloud storage URL to download this report. This URL is
	// valid for 30 minutes.
	DownloadUrl string `json:"downloadUrl,omitempty"`

	// Format: Format of the output.
	Format string `json:"format,omitempty"`

	// Type: Type of the output.
	Type string `json:"type,omitempty"`
}

// BatchReportTimeSpan: Period included in the report. For reports
// containing all entities endTime is not set. Both startTime and
// endTime are inclusive.
type BatchReportTimeSpan struct {
	// EndTime: End of the period included in the report. Inclusive. For
	// reports containing all entities endTime is not set.
	EndTime string `json:"endTime,omitempty"`

	// StartTime: Start of the period included in the report. Inclusive.
	StartTime string `json:"startTime,omitempty"`
}

// BatchReportDefinition: Contains single batchReportDefinition
// resource.
type BatchReportDefinition struct {
	// Id: The ID that YouTube assigns and uses to uniquely identify the
	// report definition.
	Id string `json:"id,omitempty"`

	// Kind: This value specifies the type of data of this item. For batch
	// report definition the kind property value is
	// youtubeAnalytics#batchReportDefinition.
	Kind string `json:"kind,omitempty"`

	// Name: Name of the report definition.
	Name string `json:"name,omitempty"`

	// Status: Status of the report definition.
	Status string `json:"status,omitempty"`

	// Type: Type of the report definition.
	Type string `json:"type,omitempty"`
}

// BatchReportDefinitionList: A paginated list of batchReportDefinition
// resources returned in response to a
// youtubeAnalytics.batchReportDefinitions.list request.
type BatchReportDefinitionList struct {
	// Items: A list of batchReportDefinition resources that match the
	// request criteria.
	Items []*BatchReportDefinition `json:"items,omitempty"`

	// Kind: This value specifies the type of data included in the API
	// response. For the list method, the kind property value is
	// youtubeAnalytics#batchReportDefinitionList.
	Kind string `json:"kind,omitempty"`
}

// BatchReportList: A paginated list of batchReport resources returned
// in response to a youtubeAnalytics.batchReport.list request.
type BatchReportList struct {
	// Items: A list of batchReport resources that match the request
	// criteria.
	Items []*BatchReport `json:"items,omitempty"`

	// Kind: This value specifies the type of data included in the API
	// response. For the list method, the kind property value is
	// youtubeAnalytics#batchReportList.
	Kind string `json:"kind,omitempty"`
}

type Group struct {
	ContentDetails *GroupContentDetails `json:"contentDetails,omitempty"`

	Etag string `json:"etag,omitempty"`

	Id string `json:"id,omitempty"`

	Kind string `json:"kind,omitempty"`

	Snippet *GroupSnippet `json:"snippet,omitempty"`
}

type GroupContentDetails struct {
	ItemCount uint64 `json:"itemCount,omitempty,string"`

	ItemType string `json:"itemType,omitempty"`
}

type GroupSnippet struct {
	PublishedAt string `json:"publishedAt,omitempty"`

	Title string `json:"title,omitempty"`
}

type GroupItem struct {
	Etag string `json:"etag,omitempty"`

	GroupId string `json:"groupId,omitempty"`

	Id string `json:"id,omitempty"`

	Kind string `json:"kind,omitempty"`

	Resource *GroupItemResource `json:"resource,omitempty"`
}

type GroupItemResource struct {
	Id string `json:"id,omitempty"`

	Kind string `json:"kind,omitempty"`
}

// GroupItemListResponse: A paginated list of grouList resources
// returned in response to a youtubeAnalytics.groupApi.list request.
type GroupItemListResponse struct {
	Etag string `json:"etag,omitempty"`

	Items []*GroupItem `json:"items,omitempty"`

	Kind string `json:"kind,omitempty"`
}

// GroupListResponse: A paginated list of grouList resources returned in
// response to a youtubeAnalytics.groupApi.list request.
type GroupListResponse struct {
	Etag string `json:"etag,omitempty"`

	Items []*Group `json:"items,omitempty"`

	Kind string `json:"kind,omitempty"`
}

// ResultTable: Contains a single result table. The table is returned as
// an array of rows that contain the values for the cells of the table.
// Depending on the metric or dimension, the cell can contain a string
// (video ID, country code) or a number (number of views or number of
// likes).
type ResultTable struct {
	// ColumnHeaders: This value specifies information about the data
	// returned in the rows fields. Each item in the columnHeaders list
	// identifies a field returned in the rows value, which contains a list
	// of comma-delimited data. The columnHeaders list will begin with the
	// dimensions specified in the API request, which will be followed by
	// the metrics specified in the API request. The order of both
	// dimensions and metrics will match the ordering in the API request.
	// For example, if the API request contains the parameters
	// dimensions=ageGroup,gender&metrics=viewerPercentage, the API response
	// will return columns in this order: ageGroup,gender,viewerPercentage.
	ColumnHeaders []*ResultTableColumnHeaders `json:"columnHeaders,omitempty"`

	// Kind: This value specifies the type of data included in the API
	// response. For the query method, the kind property value will be
	// youtubeAnalytics#resultTable.
	Kind string `json:"kind,omitempty"`

	// Rows: The list contains all rows of the result table. Each item in
	// the list is an array that contains comma-delimited data corresponding
	// to a single row of data. The order of the comma-delimited data fields
	// will match the order of the columns listed in the columnHeaders
	// field. If no data is available for the given query, the rows element
	// will be omitted from the response. The response for a query with the
	// day dimension will not contain rows for the most recent days.
	Rows [][]interface{} `json:"rows,omitempty"`
}

type ResultTableColumnHeaders struct {
	// ColumnType: The type of the column (DIMENSION or METRIC).
	ColumnType string `json:"columnType,omitempty"`

	// DataType: The type of the data in the column (STRING, INTEGER, FLOAT,
	// etc.).
	DataType string `json:"dataType,omitempty"`

	// Name: The name of the dimension or metric.
	Name string `json:"name,omitempty"`
}

// method id "youtubeAnalytics.batchReportDefinitions.list":

type BatchReportDefinitionsListCall struct {
	s                      *Service
	onBehalfOfContentOwner string
	opt_                   map[string]interface{}
}

// List: Retrieves a list of available batch report definitions.
func (r *BatchReportDefinitionsService) List(onBehalfOfContentOwner string) *BatchReportDefinitionsListCall {
	c := &BatchReportDefinitionsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.onBehalfOfContentOwner = onBehalfOfContentOwner
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BatchReportDefinitionsListCall) Fields(s ...googleapi.Field) *BatchReportDefinitionsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *BatchReportDefinitionsListCall) Do() (*BatchReportDefinitionList, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", c.onBehalfOfContentOwner))
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "batchReportDefinitions")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *BatchReportDefinitionList
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of available batch report definitions.",
	//   "httpMethod": "GET",
	//   "id": "youtubeAnalytics.batchReportDefinitions.list",
	//   "parameterOrder": [
	//     "onBehalfOfContentOwner"
	//   ],
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "batchReportDefinitions",
	//   "response": {
	//     "$ref": "BatchReportDefinitionList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/yt-analytics-monetary.readonly",
	//     "https://www.googleapis.com/auth/yt-analytics.readonly"
	//   ]
	// }

}

// method id "youtubeAnalytics.batchReports.list":

type BatchReportsListCall struct {
	s                       *Service
	batchReportDefinitionId string
	onBehalfOfContentOwner  string
	opt_                    map[string]interface{}
}

// List: Retrieves a list of processed batch reports.
func (r *BatchReportsService) List(batchReportDefinitionId string, onBehalfOfContentOwner string) *BatchReportsListCall {
	c := &BatchReportsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.batchReportDefinitionId = batchReportDefinitionId
	c.onBehalfOfContentOwner = onBehalfOfContentOwner
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BatchReportsListCall) Fields(s ...googleapi.Field) *BatchReportsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *BatchReportsListCall) Do() (*BatchReportList, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("batchReportDefinitionId", fmt.Sprintf("%v", c.batchReportDefinitionId))
	params.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", c.onBehalfOfContentOwner))
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "batchReports")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *BatchReportList
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of processed batch reports.",
	//   "httpMethod": "GET",
	//   "id": "youtubeAnalytics.batchReports.list",
	//   "parameterOrder": [
	//     "batchReportDefinitionId",
	//     "onBehalfOfContentOwner"
	//   ],
	//   "parameters": {
	//     "batchReportDefinitionId": {
	//       "description": "The batchReportDefinitionId parameter specifies the ID of the batch reportort definition for which you are retrieving reports.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "batchReports",
	//   "response": {
	//     "$ref": "BatchReportList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/yt-analytics-monetary.readonly",
	//     "https://www.googleapis.com/auth/yt-analytics.readonly"
	//   ]
	// }

}

// method id "youtubeAnalytics.groupItems.delete":

type GroupItemsDeleteCall struct {
	s    *Service
	id   string
	opt_ map[string]interface{}
}

// Delete: Removes an item from a group.
func (r *GroupItemsService) Delete(id string) *GroupItemsDeleteCall {
	c := &GroupItemsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": Note: This parameter is intended
// exclusively for YouTube content partners.
//
// The onBehalfOfContentOwner parameter indicates that the request's
// authorization credentials identify a YouTube CMS user who is acting
// on behalf of the content owner specified in the parameter value. This
// parameter is intended for YouTube content partners that own and
// manage many different YouTube channels. It allows content owners to
// authenticate once and get access to all their video and channel data,
// without having to provide authentication credentials for each
// individual channel. The CMS account that the user authenticates with
// must be linked to the specified YouTube content owner.
func (c *GroupItemsDeleteCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *GroupItemsDeleteCall {
	c.opt_["onBehalfOfContentOwner"] = onBehalfOfContentOwner
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupItemsDeleteCall) Fields(s ...googleapi.Field) *GroupItemsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *GroupItemsDeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("id", fmt.Sprintf("%v", c.id))
	if v, ok := c.opt_["onBehalfOfContentOwner"]; ok {
		params.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "groupItems")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Removes an item from a group.",
	//   "httpMethod": "DELETE",
	//   "id": "youtubeAnalytics.groupItems.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The id parameter specifies the YouTube group item ID for the group that is being deleted.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "Note: This parameter is intended exclusively for YouTube content partners.\n\nThe onBehalfOfContentOwner parameter indicates that the request's authorization credentials identify a YouTube CMS user who is acting on behalf of the content owner specified in the parameter value. This parameter is intended for YouTube content partners that own and manage many different YouTube channels. It allows content owners to authenticate once and get access to all their video and channel data, without having to provide authentication credentials for each individual channel. The CMS account that the user authenticates with must be linked to the specified YouTube content owner.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "groupItems",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubeAnalytics.groupItems.insert":

type GroupItemsInsertCall struct {
	s         *Service
	groupitem *GroupItem
	opt_      map[string]interface{}
}

// Insert: Creates a group item.
func (r *GroupItemsService) Insert(groupitem *GroupItem) *GroupItemsInsertCall {
	c := &GroupItemsInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.groupitem = groupitem
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": Note: This parameter is intended
// exclusively for YouTube content partners.
//
// The onBehalfOfContentOwner parameter indicates that the request's
// authorization credentials identify a YouTube CMS user who is acting
// on behalf of the content owner specified in the parameter value. This
// parameter is intended for YouTube content partners that own and
// manage many different YouTube channels. It allows content owners to
// authenticate once and get access to all their video and channel data,
// without having to provide authentication credentials for each
// individual channel. The CMS account that the user authenticates with
// must be linked to the specified YouTube content owner.
func (c *GroupItemsInsertCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *GroupItemsInsertCall {
	c.opt_["onBehalfOfContentOwner"] = onBehalfOfContentOwner
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupItemsInsertCall) Fields(s ...googleapi.Field) *GroupItemsInsertCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *GroupItemsInsertCall) Do() (*GroupItem, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.groupitem)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["onBehalfOfContentOwner"]; ok {
		params.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "groupItems")
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
	var ret *GroupItem
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a group item.",
	//   "httpMethod": "POST",
	//   "id": "youtubeAnalytics.groupItems.insert",
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "Note: This parameter is intended exclusively for YouTube content partners.\n\nThe onBehalfOfContentOwner parameter indicates that the request's authorization credentials identify a YouTube CMS user who is acting on behalf of the content owner specified in the parameter value. This parameter is intended for YouTube content partners that own and manage many different YouTube channels. It allows content owners to authenticate once and get access to all their video and channel data, without having to provide authentication credentials for each individual channel. The CMS account that the user authenticates with must be linked to the specified YouTube content owner.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "groupItems",
	//   "request": {
	//     "$ref": "GroupItem"
	//   },
	//   "response": {
	//     "$ref": "GroupItem"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubeAnalytics.groupItems.list":

type GroupItemsListCall struct {
	s       *Service
	groupId string
	opt_    map[string]interface{}
}

// List: Returns a collection of group items that match the API request
// parameters.
func (r *GroupItemsService) List(groupId string) *GroupItemsListCall {
	c := &GroupItemsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.groupId = groupId
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": Note: This parameter is intended
// exclusively for YouTube content partners.
//
// The onBehalfOfContentOwner parameter indicates that the request's
// authorization credentials identify a YouTube CMS user who is acting
// on behalf of the content owner specified in the parameter value. This
// parameter is intended for YouTube content partners that own and
// manage many different YouTube channels. It allows content owners to
// authenticate once and get access to all their video and channel data,
// without having to provide authentication credentials for each
// individual channel. The CMS account that the user authenticates with
// must be linked to the specified YouTube content owner.
func (c *GroupItemsListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *GroupItemsListCall {
	c.opt_["onBehalfOfContentOwner"] = onBehalfOfContentOwner
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupItemsListCall) Fields(s ...googleapi.Field) *GroupItemsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *GroupItemsListCall) Do() (*GroupItemListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("groupId", fmt.Sprintf("%v", c.groupId))
	if v, ok := c.opt_["onBehalfOfContentOwner"]; ok {
		params.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "groupItems")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *GroupItemListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns a collection of group items that match the API request parameters.",
	//   "httpMethod": "GET",
	//   "id": "youtubeAnalytics.groupItems.list",
	//   "parameterOrder": [
	//     "groupId"
	//   ],
	//   "parameters": {
	//     "groupId": {
	//       "description": "The id parameter specifies the unique ID of the group for which you want to retrieve group items.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "Note: This parameter is intended exclusively for YouTube content partners.\n\nThe onBehalfOfContentOwner parameter indicates that the request's authorization credentials identify a YouTube CMS user who is acting on behalf of the content owner specified in the parameter value. This parameter is intended for YouTube content partners that own and manage many different YouTube channels. It allows content owners to authenticate once and get access to all their video and channel data, without having to provide authentication credentials for each individual channel. The CMS account that the user authenticates with must be linked to the specified YouTube content owner.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "groupItems",
	//   "response": {
	//     "$ref": "GroupItemListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtube.readonly",
	//     "https://www.googleapis.com/auth/youtubepartner",
	//     "https://www.googleapis.com/auth/yt-analytics.readonly"
	//   ]
	// }

}

// method id "youtubeAnalytics.groups.delete":

type GroupsDeleteCall struct {
	s    *Service
	id   string
	opt_ map[string]interface{}
}

// Delete: Deletes a group.
func (r *GroupsService) Delete(id string) *GroupsDeleteCall {
	c := &GroupsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": Note: This parameter is intended
// exclusively for YouTube content partners.
//
// The onBehalfOfContentOwner parameter indicates that the request's
// authorization credentials identify a YouTube CMS user who is acting
// on behalf of the content owner specified in the parameter value. This
// parameter is intended for YouTube content partners that own and
// manage many different YouTube channels. It allows content owners to
// authenticate once and get access to all their video and channel data,
// without having to provide authentication credentials for each
// individual channel. The CMS account that the user authenticates with
// must be linked to the specified YouTube content owner.
func (c *GroupsDeleteCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *GroupsDeleteCall {
	c.opt_["onBehalfOfContentOwner"] = onBehalfOfContentOwner
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupsDeleteCall) Fields(s ...googleapi.Field) *GroupsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *GroupsDeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("id", fmt.Sprintf("%v", c.id))
	if v, ok := c.opt_["onBehalfOfContentOwner"]; ok {
		params.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "groups")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Deletes a group.",
	//   "httpMethod": "DELETE",
	//   "id": "youtubeAnalytics.groups.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The id parameter specifies the YouTube group ID for the group that is being deleted.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "Note: This parameter is intended exclusively for YouTube content partners.\n\nThe onBehalfOfContentOwner parameter indicates that the request's authorization credentials identify a YouTube CMS user who is acting on behalf of the content owner specified in the parameter value. This parameter is intended for YouTube content partners that own and manage many different YouTube channels. It allows content owners to authenticate once and get access to all their video and channel data, without having to provide authentication credentials for each individual channel. The CMS account that the user authenticates with must be linked to the specified YouTube content owner.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "groups",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubeAnalytics.groups.insert":

type GroupsInsertCall struct {
	s     *Service
	group *Group
	opt_  map[string]interface{}
}

// Insert: Creates a group.
func (r *GroupsService) Insert(group *Group) *GroupsInsertCall {
	c := &GroupsInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.group = group
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": Note: This parameter is intended
// exclusively for YouTube content partners.
//
// The onBehalfOfContentOwner parameter indicates that the request's
// authorization credentials identify a YouTube CMS user who is acting
// on behalf of the content owner specified in the parameter value. This
// parameter is intended for YouTube content partners that own and
// manage many different YouTube channels. It allows content owners to
// authenticate once and get access to all their video and channel data,
// without having to provide authentication credentials for each
// individual channel. The CMS account that the user authenticates with
// must be linked to the specified YouTube content owner.
func (c *GroupsInsertCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *GroupsInsertCall {
	c.opt_["onBehalfOfContentOwner"] = onBehalfOfContentOwner
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupsInsertCall) Fields(s ...googleapi.Field) *GroupsInsertCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *GroupsInsertCall) Do() (*Group, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.group)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["onBehalfOfContentOwner"]; ok {
		params.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "groups")
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
	var ret *Group
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a group.",
	//   "httpMethod": "POST",
	//   "id": "youtubeAnalytics.groups.insert",
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "Note: This parameter is intended exclusively for YouTube content partners.\n\nThe onBehalfOfContentOwner parameter indicates that the request's authorization credentials identify a YouTube CMS user who is acting on behalf of the content owner specified in the parameter value. This parameter is intended for YouTube content partners that own and manage many different YouTube channels. It allows content owners to authenticate once and get access to all their video and channel data, without having to provide authentication credentials for each individual channel. The CMS account that the user authenticates with must be linked to the specified YouTube content owner.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "groups",
	//   "request": {
	//     "$ref": "Group"
	//   },
	//   "response": {
	//     "$ref": "Group"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubeAnalytics.groups.list":

type GroupsListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Returns a collection of groups that match the API request
// parameters. For example, you can retrieve all groups that the
// authenticated user owns, or you can retrieve one or more groups by
// their unique IDs.
func (r *GroupsService) List() *GroupsListCall {
	c := &GroupsListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// Id sets the optional parameter "id": The id parameter specifies a
// comma-separated list of the YouTube group ID(s) for the resource(s)
// that are being retrieved. In a group resource, the id property
// specifies the group's YouTube group ID.
func (c *GroupsListCall) Id(id string) *GroupsListCall {
	c.opt_["id"] = id
	return c
}

// Mine sets the optional parameter "mine": Set this parameter's value
// to true to instruct the API to only return groups owned by the
// authenticated user.
func (c *GroupsListCall) Mine(mine bool) *GroupsListCall {
	c.opt_["mine"] = mine
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": Note: This parameter is intended
// exclusively for YouTube content partners.
//
// The onBehalfOfContentOwner parameter indicates that the request's
// authorization credentials identify a YouTube CMS user who is acting
// on behalf of the content owner specified in the parameter value. This
// parameter is intended for YouTube content partners that own and
// manage many different YouTube channels. It allows content owners to
// authenticate once and get access to all their video and channel data,
// without having to provide authentication credentials for each
// individual channel. The CMS account that the user authenticates with
// must be linked to the specified YouTube content owner.
func (c *GroupsListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *GroupsListCall {
	c.opt_["onBehalfOfContentOwner"] = onBehalfOfContentOwner
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupsListCall) Fields(s ...googleapi.Field) *GroupsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *GroupsListCall) Do() (*GroupListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["id"]; ok {
		params.Set("id", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["mine"]; ok {
		params.Set("mine", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["onBehalfOfContentOwner"]; ok {
		params.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "groups")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *GroupListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns a collection of groups that match the API request parameters. For example, you can retrieve all groups that the authenticated user owns, or you can retrieve one or more groups by their unique IDs.",
	//   "httpMethod": "GET",
	//   "id": "youtubeAnalytics.groups.list",
	//   "parameters": {
	//     "id": {
	//       "description": "The id parameter specifies a comma-separated list of the YouTube group ID(s) for the resource(s) that are being retrieved. In a group resource, the id property specifies the group's YouTube group ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "mine": {
	//       "description": "Set this parameter's value to true to instruct the API to only return groups owned by the authenticated user.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "Note: This parameter is intended exclusively for YouTube content partners.\n\nThe onBehalfOfContentOwner parameter indicates that the request's authorization credentials identify a YouTube CMS user who is acting on behalf of the content owner specified in the parameter value. This parameter is intended for YouTube content partners that own and manage many different YouTube channels. It allows content owners to authenticate once and get access to all their video and channel data, without having to provide authentication credentials for each individual channel. The CMS account that the user authenticates with must be linked to the specified YouTube content owner.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "groups",
	//   "response": {
	//     "$ref": "GroupListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtube.readonly",
	//     "https://www.googleapis.com/auth/youtubepartner",
	//     "https://www.googleapis.com/auth/yt-analytics.readonly"
	//   ]
	// }

}

// method id "youtubeAnalytics.groups.update":

type GroupsUpdateCall struct {
	s     *Service
	group *Group
	opt_  map[string]interface{}
}

// Update: Modifies a group. For example, you could change a group's
// title.
func (r *GroupsService) Update(group *Group) *GroupsUpdateCall {
	c := &GroupsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.group = group
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": Note: This parameter is intended
// exclusively for YouTube content partners.
//
// The onBehalfOfContentOwner parameter indicates that the request's
// authorization credentials identify a YouTube CMS user who is acting
// on behalf of the content owner specified in the parameter value. This
// parameter is intended for YouTube content partners that own and
// manage many different YouTube channels. It allows content owners to
// authenticate once and get access to all their video and channel data,
// without having to provide authentication credentials for each
// individual channel. The CMS account that the user authenticates with
// must be linked to the specified YouTube content owner.
func (c *GroupsUpdateCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *GroupsUpdateCall {
	c.opt_["onBehalfOfContentOwner"] = onBehalfOfContentOwner
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupsUpdateCall) Fields(s ...googleapi.Field) *GroupsUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *GroupsUpdateCall) Do() (*Group, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.group)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["onBehalfOfContentOwner"]; ok {
		params.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "groups")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
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
	var ret *Group
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Modifies a group. For example, you could change a group's title.",
	//   "httpMethod": "PUT",
	//   "id": "youtubeAnalytics.groups.update",
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "Note: This parameter is intended exclusively for YouTube content partners.\n\nThe onBehalfOfContentOwner parameter indicates that the request's authorization credentials identify a YouTube CMS user who is acting on behalf of the content owner specified in the parameter value. This parameter is intended for YouTube content partners that own and manage many different YouTube channels. It allows content owners to authenticate once and get access to all their video and channel data, without having to provide authentication credentials for each individual channel. The CMS account that the user authenticates with must be linked to the specified YouTube content owner.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "groups",
	//   "request": {
	//     "$ref": "Group"
	//   },
	//   "response": {
	//     "$ref": "Group"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

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

// Currency sets the optional parameter "currency": The currency to
// which financial metrics should be converted. The default is US Dollar
// (USD). If the result contains no financial metrics, this flag will be
// ignored. Responds with an error if the specified currency is not
// recognized.
func (c *ReportsQueryCall) Currency(currency string) *ReportsQueryCall {
	c.opt_["currency"] = currency
	return c
}

// Dimensions sets the optional parameter "dimensions": A
// comma-separated list of YouTube Analytics dimensions, such as views
// or ageGroup,gender. See the Available Reports document for a list of
// the reports that you can retrieve and the dimensions used for those
// reports. Also see the Dimensions document for definitions of those
// dimensions.
func (c *ReportsQueryCall) Dimensions(dimensions string) *ReportsQueryCall {
	c.opt_["dimensions"] = dimensions
	return c
}

// Filters sets the optional parameter "filters": A list of filters that
// should be applied when retrieving YouTube Analytics data. The
// Available Reports document identifies the dimensions that can be used
// to filter each report, and the Dimensions document defines those
// dimensions. If a request uses multiple filters, join them together
// with a semicolon (;), and the returned result table will satisfy both
// filters. For example, a filters parameter value of
// video==dMH0bHeiRNg;country==IT restricts the result set to include
// data for the given video in Italy.
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
// Analytics data. By default the sort order is ascending. The '-'
// prefix causes descending sort order.
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

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReportsQueryCall) Fields(s ...googleapi.Field) *ReportsQueryCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
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
	if v, ok := c.opt_["currency"]; ok {
		params.Set("currency", fmt.Sprintf("%v", v))
	}
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
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "reports")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ResultTable
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
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
	//     "currency": {
	//       "description": "The currency to which financial metrics should be converted. The default is US Dollar (USD). If the result contains no financial metrics, this flag will be ignored. Responds with an error if the specified currency is not recognized.",
	//       "location": "query",
	//       "pattern": "[A-Z]{3}",
	//       "type": "string"
	//     },
	//     "dimensions": {
	//       "description": "A comma-separated list of YouTube Analytics dimensions, such as views or ageGroup,gender. See the Available Reports document for a list of the reports that you can retrieve and the dimensions used for those reports. Also see the Dimensions document for definitions of those dimensions.",
	//       "location": "query",
	//       "pattern": "[0-9a-zA-Z,]+",
	//       "type": "string"
	//     },
	//     "end-date": {
	//       "description": "The end date for fetching YouTube Analytics data. The value should be in YYYY-MM-DD format.",
	//       "location": "query",
	//       "pattern": "[0-9]{4}-[0-9]{2}-[0-9]{2}",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "filters": {
	//       "description": "A list of filters that should be applied when retrieving YouTube Analytics data. The Available Reports document identifies the dimensions that can be used to filter each report, and the Dimensions document defines those dimensions. If a request uses multiple filters, join them together with a semicolon (;), and the returned result table will satisfy both filters. For example, a filters parameter value of video==dMH0bHeiRNg;country==IT restricts the result set to include data for the given video in Italy.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ids": {
	//       "description": "Identifies the YouTube channel or content owner for which you are retrieving YouTube Analytics data.\n- To request data for a YouTube user, set the ids parameter value to channel==CHANNEL_ID, where CHANNEL_ID specifies the unique YouTube channel ID.\n- To request data for a YouTube CMS content owner, set the ids parameter value to contentOwner==OWNER_NAME, where OWNER_NAME is the CMS name of the content owner.",
	//       "location": "query",
	//       "pattern": "[a-zA-Z]+==[a-zA-Z0-9_+-]+",
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
	//       "description": "A comma-separated list of YouTube Analytics metrics, such as views or likes,dislikes. See the Available Reports document for a list of the reports that you can retrieve and the metrics available in each report, and see the Metrics document for definitions of those metrics.",
	//       "location": "query",
	//       "pattern": "[0-9a-zA-Z,]+",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sort": {
	//       "description": "A comma-separated list of dimensions or metrics that determine the sort order for YouTube Analytics data. By default the sort order is ascending. The '-' prefix causes descending sort order.",
	//       "location": "query",
	//       "pattern": "[-0-9a-zA-Z,]+",
	//       "type": "string"
	//     },
	//     "start-date": {
	//       "description": "The start date for fetching YouTube Analytics data. The value should be in YYYY-MM-DD format.",
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
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/yt-analytics-monetary.readonly",
	//     "https://www.googleapis.com/auth/yt-analytics.readonly"
	//   ]
	// }

}
