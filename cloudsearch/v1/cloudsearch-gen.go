// Package cloudsearch provides access to the Google Cloud Search API.
//
// Usage example:
//
//   import "google.golang.org/api/cloudsearch/v1"
//   ...
//   cloudsearchService, err := cloudsearch.New(oauthHttpClient)
package cloudsearch

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

const apiId = "cloudsearch:v1"
const apiName = "cloudsearch"
const apiVersion = "v1"
const basePath = "https://cloudsearch.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// FOR TESTING ONLY
	CloudsearchScope = "https://www.googleapis.com/auth/cloudsearch"

	// View your email address
	UserinfoEmailScope = "https://www.googleapis.com/auth/userinfo.email"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Projects *ProjectsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Indexes = NewProjectsIndexesService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Indexes *ProjectsIndexesService
}

func NewProjectsIndexesService(s *Service) *ProjectsIndexesService {
	rs := &ProjectsIndexesService{s: s}
	rs.Documents = NewProjectsIndexesDocumentsService(s)
	return rs
}

type ProjectsIndexesService struct {
	s *Service

	Documents *ProjectsIndexesDocumentsService
}

func NewProjectsIndexesDocumentsService(s *Service) *ProjectsIndexesDocumentsService {
	rs := &ProjectsIndexesDocumentsService{s: s}
	return rs
}

type ProjectsIndexesDocumentsService struct {
	s *Service
}

type Document struct {
	// DocId: The unique identifier of the document. It must contain only
	// visible, printable ASCII characters (ASCII codes 33 through 126
	// inclusive) and be no longer than 500 characters. It cannot begin with
	// an exclamation point ('!'), and it can't begin and end with double
	// underscores ("__"). If missing, it is automatically assigned for the
	// document.
	DocId string `json:"docId,omitempty"`

	// Fields: The list of fields in the document. It cannot be the empty
	// list. Each field has a name and a list of values. The field name is
	// unique to a document and is case sensitive. The name can only contain
	// ASCII characters. It must start with a letter and can contain
	// letters, digits, or underscore. It cannot be longer than 500
	// characters and cannot be the empty string. A field can have multiple
	// values with same or different types, however, it cannot have multiple
	// Timestamp or number values.
	Fields map[string]FieldValueList `json:"fields,omitempty"`

	// Rank: A positive integer which determines the default ordering of
	// documents returned from a search. The rank can be set explicitly when
	// the document is created. It is a bad idea to assign the same rank to
	// many documents, and the same rank should never be assigned to more
	// than 10,000 documents. By default (when it is not specified or set to
	// 0), it is set at the time the document is created to the number of
	// seconds since January 1, 2011. The rank can be used in
	// field_expressions, order_by or return_fields in a search request,
	// where it is referenced as `_rank`.
	Rank int64 `json:"rank,omitempty"`
}

type Empty struct {
}

type FieldNames struct {
	// AtomFields: The names of fields in which ATOM values are stored.
	AtomFields []string `json:"atomFields,omitempty"`

	// DateFields: The names of fields in which DATE values are stored.
	DateFields []string `json:"dateFields,omitempty"`

	// GeoFields: The names of fields in which GEO values are stored.
	GeoFields []string `json:"geoFields,omitempty"`

	// HtmlFields: The names of fields in which HTML values are stored.
	HtmlFields []string `json:"htmlFields,omitempty"`

	// NumberFields: The names of fields in which NUMBER values are stored.
	NumberFields []string `json:"numberFields,omitempty"`

	// TextFields: The names of fields in which TEXT values are stored.
	TextFields []string `json:"textFields,omitempty"`
}

type FieldValue struct {
	// GeoValue: The value of a GEO-valued field, represented in string with
	// any of the listed [ways of writing
	// coordinates](http://en.wikipedia.org/wiki/Geographic_coordinate_conver
	// sion#Ways_of_writing_coordinates)
	GeoValue string `json:"geoValue,omitempty"`

	// Lang: The language of a string value. If given, the language must be
	// a valid `ISO 639-1` code.
	Lang string `json:"lang,omitempty"`

	// NumberValue: The value of a number-valued field.
	NumberValue float64 `json:"numberValue,omitempty"`

	// StringFormat: The format of a string value. By default, the string
	// format is `DEFAULT`, where a format will be automatically detected.
	StringFormat string `json:"stringFormat,omitempty"`

	// StringValue: The value of a string-valued field.
	StringValue string `json:"stringValue,omitempty"`

	// TimestampValue: The value of a timestamp-valued field.
	TimestampValue string `json:"timestampValue,omitempty"`
}

type FieldValueList struct {
	// Values: The list of typed values.
	Values []*FieldValue `json:"values,omitempty"`
}

type IndexInfo struct {
	// IndexId: The index identifier. It cannot be the empty string. It must
	// contain only visible, printable ASCII characters (ASCII codes 33
	// through 126 inclusive) and be no longer than 100 characters. It
	// cannot begin with an exclamation point ('!'), and it can't begin and
	// end with double underscores ("__").
	IndexId string `json:"indexId,omitempty"`

	// IndexedField: Names of indexed fields.
	IndexedField *FieldNames `json:"indexedField,omitempty"`

	// ProjectId: The project associated with the index. It cannot be the
	// empty string.
	ProjectId string `json:"projectId,omitempty"`
}

type ListDocumentsResponse struct {
	// Documents: The list of documents.
	Documents []*Document `json:"documents,omitempty"`

	// NextPageToken: If there are more results, retrieve them by invoking
	// list documents call with the same arguments and this `nextPageToken`.
	// If there are no more results, this field is not set.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListIndexesResponse struct {
	// Indexes: The information about available indexes.
	Indexes []*IndexInfo `json:"indexes,omitempty"`

	// NextPageToken: If there are more results, retrieve them by invoking
	// list indexes call with the same arguments and this `nextPageToken`.
	// If there are no more results, this field is not set.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type SearchResponse struct {
	// MatchedCount: The number of documents that match the query. It is
	// greater than or equal to the number of documents actually returned.
	// This is an approximation and not an exact count unless it is less
	// than or equal to `matchedCountAccuracy` in search parameter.
	MatchedCount int64 `json:"matchedCount,omitempty,string"`

	// Results: The list of documents that match the search query.
	Results []*SearchResult `json:"results,omitempty"`
}

type SearchResult struct {
	// DocId: The unique identifier of the document.
	DocId string `json:"docId,omitempty"`

	// Fields: The list of fields in the result. Each field is either from
	// the stored document, the built-in fields (`_rank`, the document rank,
	// and `_score` if scoring is enabled), or computed from any extra
	// `fieldExpressions` defined in the request. For example, if a request
	// contains a `fieldExpressions` named `"TotalPrice"` and expressed as
	// `"Price + Tax"`, the result will have a field whose name is
	// `"TotalPrice"` and whose value is set to the computed sum of the
	// value of field `"Price"` and the value of field `"Tax"`. If a request
	// contains a `fieldExpressions` named `"snippet"` and expressed as
	// `"snippet(\"good times\", content)"`, the result will have a field
	// whose name is `"snippet"` and whose value contains a snippet of text
	// from field `"content"` matching the query "good times".
	Fields map[string]FieldValueList `json:"fields,omitempty"`

	// NextPageToken: If there are more results, retrieve them by invoking
	// search call with the same arguments and this `nextPageToken`. If
	// there are no more results, this field is not set.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

// method id "cloudsearch.projects.indexes.list":

type ProjectsIndexesListCall struct {
	s               *Service
	projectId       string
	indexNamePrefix string
	pageSize        int64
	pageToken       string
	view            string
	opt_            map[string]interface{}
}

// List: Lists search indexes belonging to the specified project.
func (r *ProjectsIndexesService) List(projectId string, indexNamePrefix string, pageSize int64, pageToken string, view string) *ProjectsIndexesListCall {
	c := &ProjectsIndexesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.indexNamePrefix = indexNamePrefix
	c.pageSize = pageSize
	c.pageToken = pageToken
	c.view = view
	return c
}

// IndexNamePrefix sets the optional parameter "indexNamePrefix": The
// prefix of the index name. It is used to list all indexes with names
// that have this prefix.
func (c *ProjectsIndexesListCall) IndexNamePrefix(indexNamePrefix string) *ProjectsIndexesListCall {
	c.opt_["indexNamePrefix"] = indexNamePrefix
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of indexes to return per page. If not specified, 100 indexes are
// returned per page.
func (c *ProjectsIndexesListCall) PageSize(pageSize int64) *ProjectsIndexesListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": A `nextPageToken`
// returned from previous list indexes call as the starting point for
// this call. If not specified, list indexes from the beginning.
func (c *ProjectsIndexesListCall) PageToken(pageToken string) *ProjectsIndexesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// View sets the optional parameter "view": Specifies which parts of the
// IndexInfo resource is returned in the response. If not specified,
// `ID_ONLY` is used.
func (c *ProjectsIndexesListCall) View(view string) *ProjectsIndexesListCall {
	c.opt_["view"] = view
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsIndexesListCall) Fields(s ...googleapi.Field) *ProjectsIndexesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsIndexesListCall) Do() (*ListIndexesResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["indexNamePrefix"]; ok {
		params.Set("indexNamePrefix", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["view"]; ok {
		params.Set("view", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/indexes")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListIndexesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists search indexes belonging to the specified project.",
	//   "httpMethod": "GET",
	//   "id": "cloudsearch.projects.indexes.list",
	//   "parameterOrder": [
	//     "projectId",
	//     "indexNamePrefix",
	//     "pageSize",
	//     "pageToken",
	//     "view"
	//   ],
	//   "parameters": {
	//     "indexNamePrefix": {
	//       "description": "The prefix of the index name. It is used to list all indexes with names that have this prefix.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of indexes to return per page. If not specified, 100 indexes are returned per page.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A `nextPageToken` returned from previous list indexes call as the starting point for this call. If not specified, list indexes from the beginning.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The project from which to retrieve indexes. It cannot be the empty string.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "view": {
	//       "description": "Specifies which parts of the IndexInfo resource is returned in the response. If not specified, `ID_ONLY` is used.",
	//       "enum": [
	//         "INDEX_VIEW_UNSPECIFIED",
	//         "ID_ONLY",
	//         "FULL"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/indexes",
	//   "response": {
	//     "$ref": "ListIndexesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloudsearch",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "cloudsearch.projects.indexes.search":

type ProjectsIndexesSearchCall struct {
	s                    *Service
	projectId            string
	indexId              string
	query                string
	fieldExpressions     []string
	pageSize             int64
	pageToken            string
	offset               int64
	matchedCountAccuracy int64
	orderBy              string
	scorer               string
	scorerSize           int64
	returnFields         []string
	opt_                 map[string]interface{}
}

// Search: Lists the documents in the named index that match the query.
func (r *ProjectsIndexesService) Search(projectId string, indexId string, query string, fieldExpressions []string, pageSize int64, pageToken string, offset int64, matchedCountAccuracy int64, orderBy string, scorer string, scorerSize int64, returnFields []string) *ProjectsIndexesSearchCall {
	c := &ProjectsIndexesSearchCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.indexId = indexId
	c.query = query
	c.fieldExpressions = fieldExpressions
	c.pageSize = pageSize
	c.pageToken = pageToken
	c.offset = offset
	c.matchedCountAccuracy = matchedCountAccuracy
	c.orderBy = orderBy
	c.scorer = scorer
	c.scorerSize = scorerSize
	c.returnFields = returnFields
	return c
}

// FieldExpressions sets the optional parameter "fieldExpressions":
// Customized expressions used in `orderBy` or `returnFields`. The
// expression can contain fields in `Document`, the built-in fields (
// `_rank`, the document rank, and `_score` if scoring is enabled) and
// fields defined in `fieldExpressions`. Each field expression is
// represented in a json object with the following fields: * `name`: the
// name of the field expression in string. * `expression`: the
// expression to be computed. It can be a combination of supported
// functions encoded in string. Expressions involving number fields can
// use the arithmetical operators (`+`, `-`, `*`, `/`) and the built-in
// numeric functions (`max`, `min`, `pow`, `count`, `log`, `abs`).
// Expressions involving geopoint fields can use the `geopoint` and
// `distance` functions. Expressions for text and html fields can use
// the `snippet` function. For example: ``` fieldExpressions={name:
// "TotalPrice", expression: "(Price+Tax)"} ``` ```
// fieldExpressions={name: "snippet", expression: "snippet('good times',
// content)"} ``` The field expression names can be used in `orderBy`
// and `returnFields` after they are defined in `fieldExpressions`.
func (c *ProjectsIndexesSearchCall) FieldExpressions(fieldExpressions string) *ProjectsIndexesSearchCall {
	c.opt_["fieldExpressions"] = fieldExpressions
	return c
}

// MatchedCountAccuracy sets the optional parameter
// "matchedCountAccuracy": Minimum accuracy requirement for
// `matchedCount` in search response. If specified, `matchedCount` will
// be accurate up to at least that number. For example, when set to 100,
// any `matchedCount <= 100` is accurate. This option may add
// considerable latency/expense. By default (when it is not specified or
// set to 0), the accuracy is the same as `pageSize`.
func (c *ProjectsIndexesSearchCall) MatchedCountAccuracy(matchedCountAccuracy int64) *ProjectsIndexesSearchCall {
	c.opt_["matchedCountAccuracy"] = matchedCountAccuracy
	return c
}

// Offset sets the optional parameter "offset": Offset is used to move
// to an arbitrary result, independent of the previous results. Offsets
// are inefficient when compared to `pageToken`. `pageToken` and
// `offset` cannot be both set. The default value of `offset` is 0.
func (c *ProjectsIndexesSearchCall) Offset(offset int64) *ProjectsIndexesSearchCall {
	c.opt_["offset"] = offset
	return c
}

// OrderBy sets the optional parameter "orderBy": Comma-separated list
// of fields for sorting on the search result, including fields from
// `Document`, the built-in fields (`_rank` and `_score`), and fields
// defined in `fieldExpressions`. For example: `orderBy="foo,bar"`. The
// default sorting order is ascending. To specify descending order for a
// field, a suffix `" desc"` should be appended to the field name. For
// example: `orderBy="foo desc,bar"`. The default value for text sort is
// the empty string, and the default value for numeric sort is 0. If not
// specified, the search results are automatically sorted by descending
// `_rank`. Sorting by ascending `_rank` is not allowed.
func (c *ProjectsIndexesSearchCall) OrderBy(orderBy string) *ProjectsIndexesSearchCall {
	c.opt_["orderBy"] = orderBy
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of search results to return per page. Searches perform best when the
// `pageSize` is kept as small as possible. If not specified, 10 results
// are returned per page.
func (c *ProjectsIndexesSearchCall) PageSize(pageSize int64) *ProjectsIndexesSearchCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": A `nextPageToken`
// returned from previous Search call as the starting point for this
// call. Pagination tokens provide better performance and consistency
// than offsets, and they cannot be used in combination with offsets.
func (c *ProjectsIndexesSearchCall) PageToken(pageToken string) *ProjectsIndexesSearchCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Query sets the optional parameter "query": The query string in search
// query syntax. If the query is missing or empty, all documents are
// returned.
func (c *ProjectsIndexesSearchCall) Query(query string) *ProjectsIndexesSearchCall {
	c.opt_["query"] = query
	return c
}

// ReturnFields sets the optional parameter "returnFields": List of
// fields to return in `SearchResult` objects. It can be fields from
// `Document`, the built-in fields `_rank` and `_score`, and fields
// defined in `fieldExpressions`. Use `"*"` to return all fields from
// `Document`.
func (c *ProjectsIndexesSearchCall) ReturnFields(returnFields string) *ProjectsIndexesSearchCall {
	c.opt_["returnFields"] = returnFields
	return c
}

// Scorer sets the optional parameter "scorer": The scoring function to
// invoke on a search result for this query. If `scorer` is not set,
// scoring is disabled and `_score` is 0 for all documents in the search
// result. To enable document relevancy score based on term frequency,
// set `"scorer=generic"`.
func (c *ProjectsIndexesSearchCall) Scorer(scorer string) *ProjectsIndexesSearchCall {
	c.opt_["scorer"] = scorer
	return c
}

// ScorerSize sets the optional parameter "scorerSize": Maximum number
// of top retrieved results to score. It is valid only when `scorer` is
// set. If not specified, 100 retrieved results are scored.
func (c *ProjectsIndexesSearchCall) ScorerSize(scorerSize int64) *ProjectsIndexesSearchCall {
	c.opt_["scorerSize"] = scorerSize
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsIndexesSearchCall) Fields(s ...googleapi.Field) *ProjectsIndexesSearchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsIndexesSearchCall) Do() (*SearchResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fieldExpressions"]; ok {
		params.Set("fieldExpressions", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["matchedCountAccuracy"]; ok {
		params.Set("matchedCountAccuracy", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["offset"]; ok {
		params.Set("offset", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["orderBy"]; ok {
		params.Set("orderBy", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["query"]; ok {
		params.Set("query", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["returnFields"]; ok {
		params.Set("returnFields", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["scorer"]; ok {
		params.Set("scorer", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["scorerSize"]; ok {
		params.Set("scorerSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/indexes/{indexId}/search")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"indexId":   c.indexId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *SearchResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the documents in the named index that match the query.",
	//   "httpMethod": "GET",
	//   "id": "cloudsearch.projects.indexes.search",
	//   "parameterOrder": [
	//     "projectId",
	//     "indexId",
	//     "query",
	//     "fieldExpressions",
	//     "pageSize",
	//     "pageToken",
	//     "offset",
	//     "matchedCountAccuracy",
	//     "orderBy",
	//     "scorer",
	//     "scorerSize",
	//     "returnFields"
	//   ],
	//   "parameters": {
	//     "fieldExpressions": {
	//       "description": "Customized expressions used in `orderBy` or `returnFields`. The expression can contain fields in `Document`, the built-in fields ( `_rank`, the document rank, and `_score` if scoring is enabled) and fields defined in `fieldExpressions`. Each field expression is represented in a json object with the following fields: * `name`: the name of the field expression in string. * `expression`: the expression to be computed. It can be a combination of supported functions encoded in string. Expressions involving number fields can use the arithmetical operators (`+`, `-`, `*`, `/`) and the built-in numeric functions (`max`, `min`, `pow`, `count`, `log`, `abs`). Expressions involving geopoint fields can use the `geopoint` and `distance` functions. Expressions for text and html fields can use the `snippet` function. For example: ``` fieldExpressions={name: \"TotalPrice\", expression: \"(Price+Tax)\"} ``` ``` fieldExpressions={name: \"snippet\", expression: \"snippet('good times', content)\"} ``` The field expression names can be used in `orderBy` and `returnFields` after they are defined in `fieldExpressions`.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "indexId": {
	//       "description": "The index to search. It cannot be the empty string.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "matchedCountAccuracy": {
	//       "description": "Minimum accuracy requirement for `matchedCount` in search response. If specified, `matchedCount` will be accurate up to at least that number. For example, when set to 100, any `matchedCount \u003c= 100` is accurate. This option may add considerable latency/expense. By default (when it is not specified or set to 0), the accuracy is the same as `pageSize`.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "offset": {
	//       "description": "Offset is used to move to an arbitrary result, independent of the previous results. Offsets are inefficient when compared to `pageToken`. `pageToken` and `offset` cannot be both set. The default value of `offset` is 0.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "orderBy": {
	//       "description": "Comma-separated list of fields for sorting on the search result, including fields from `Document`, the built-in fields (`_rank` and `_score`), and fields defined in `fieldExpressions`. For example: `orderBy=\"foo,bar\"`. The default sorting order is ascending. To specify descending order for a field, a suffix `\" desc\"` should be appended to the field name. For example: `orderBy=\"foo desc,bar\"`. The default value for text sort is the empty string, and the default value for numeric sort is 0. If not specified, the search results are automatically sorted by descending `_rank`. Sorting by ascending `_rank` is not allowed.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of search results to return per page. Searches perform best when the `pageSize` is kept as small as possible. If not specified, 10 results are returned per page.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A `nextPageToken` returned from previous Search call as the starting point for this call. Pagination tokens provide better performance and consistency than offsets, and they cannot be used in combination with offsets.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The project associated with the index for searching document. It cannot be the empty string.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "query": {
	//       "description": "The query string in search query syntax. If the query is missing or empty, all documents are returned.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "returnFields": {
	//       "description": "List of fields to return in `SearchResult` objects. It can be fields from `Document`, the built-in fields `_rank` and `_score`, and fields defined in `fieldExpressions`. Use `\"*\"` to return all fields from `Document`.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "scorer": {
	//       "description": "The scoring function to invoke on a search result for this query. If `scorer` is not set, scoring is disabled and `_score` is 0 for all documents in the search result. To enable document relevancy score based on term frequency, set `\"scorer=generic\"`.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "scorerSize": {
	//       "description": "Maximum number of top retrieved results to score. It is valid only when `scorer` is set. If not specified, 100 retrieved results are scored.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/indexes/{indexId}/search",
	//   "response": {
	//     "$ref": "SearchResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloudsearch",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "cloudsearch.projects.indexes.documents.create":

type ProjectsIndexesDocumentsCreateCall struct {
	s         *Service
	projectId string
	indexId   string
	document  *Document
	opt_      map[string]interface{}
}

// Create: Inserts a document for indexing or updates an indexed
// document. The returned document contains only the ID of the new
// document. When `docId` is absent from the document, it is provided by
// the server.
func (r *ProjectsIndexesDocumentsService) Create(projectId string, indexId string, document *Document) *ProjectsIndexesDocumentsCreateCall {
	c := &ProjectsIndexesDocumentsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.indexId = indexId
	c.document = document
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsIndexesDocumentsCreateCall) Fields(s ...googleapi.Field) *ProjectsIndexesDocumentsCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsIndexesDocumentsCreateCall) Do() (*Document, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.document)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/indexes/{indexId}/documents")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"indexId":   c.indexId,
	})
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
	var ret *Document
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Inserts a document for indexing or updates an indexed document. The returned document contains only the ID of the new document. When `docId` is absent from the document, it is provided by the server.",
	//   "httpMethod": "POST",
	//   "id": "cloudsearch.projects.indexes.documents.create",
	//   "parameterOrder": [
	//     "projectId",
	//     "indexId"
	//   ],
	//   "parameters": {
	//     "indexId": {
	//       "description": "The index to add document to. It cannot be the empty string.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The project associated with the index for adding document. It cannot be the empty string.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/indexes/{indexId}/documents",
	//   "request": {
	//     "$ref": "Document"
	//   },
	//   "response": {
	//     "$ref": "Document"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloudsearch",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "cloudsearch.projects.indexes.documents.delete":

type ProjectsIndexesDocumentsDeleteCall struct {
	s         *Service
	projectId string
	indexId   string
	docId     string
	opt_      map[string]interface{}
}

// Delete: Deletes a document from an index.
func (r *ProjectsIndexesDocumentsService) Delete(projectId string, indexId string, docId string) *ProjectsIndexesDocumentsDeleteCall {
	c := &ProjectsIndexesDocumentsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.indexId = indexId
	c.docId = docId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsIndexesDocumentsDeleteCall) Fields(s ...googleapi.Field) *ProjectsIndexesDocumentsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsIndexesDocumentsDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/indexes/{indexId}/documents/{docId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"indexId":   c.indexId,
		"docId":     c.docId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a document from an index.",
	//   "httpMethod": "DELETE",
	//   "id": "cloudsearch.projects.indexes.documents.delete",
	//   "parameterOrder": [
	//     "projectId",
	//     "indexId",
	//     "docId"
	//   ],
	//   "parameters": {
	//     "docId": {
	//       "description": "The document to be deleted. It cannot be the empty string.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "indexId": {
	//       "description": "The index from which to delete the document. It cannot be the empty string.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The project associated with the index for deleting document. It cannot be the empty string.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/indexes/{indexId}/documents/{docId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloudsearch",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "cloudsearch.projects.indexes.documents.get":

type ProjectsIndexesDocumentsGetCall struct {
	s         *Service
	projectId string
	indexId   string
	docId     string
	opt_      map[string]interface{}
}

// Get: Retrieves a document from an index.
func (r *ProjectsIndexesDocumentsService) Get(projectId string, indexId string, docId string) *ProjectsIndexesDocumentsGetCall {
	c := &ProjectsIndexesDocumentsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.indexId = indexId
	c.docId = docId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsIndexesDocumentsGetCall) Fields(s ...googleapi.Field) *ProjectsIndexesDocumentsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsIndexesDocumentsGetCall) Do() (*Document, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/indexes/{indexId}/documents/{docId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"indexId":   c.indexId,
		"docId":     c.docId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Document
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a document from an index.",
	//   "httpMethod": "GET",
	//   "id": "cloudsearch.projects.indexes.documents.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "indexId",
	//     "docId"
	//   ],
	//   "parameters": {
	//     "docId": {
	//       "description": "The identifier of the document to retrieve. It cannot be the empty string.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "indexId": {
	//       "description": "The index from which to retrieve the document. It cannot be the empty string.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The project associated with the index for retrieving the document. It cannot be the empty string.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/indexes/{indexId}/documents/{docId}",
	//   "response": {
	//     "$ref": "Document"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloudsearch",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "cloudsearch.projects.indexes.documents.list":

type ProjectsIndexesDocumentsListCall struct {
	s         *Service
	projectId string
	indexId   string
	pageSize  int64
	pageToken string
	view      string
	opt_      map[string]interface{}
}

// List: Lists documents in the specified search index. Intended for
// batch processing.
func (r *ProjectsIndexesDocumentsService) List(projectId string, indexId string, pageSize int64, pageToken string, view string) *ProjectsIndexesDocumentsListCall {
	c := &ProjectsIndexesDocumentsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.indexId = indexId
	c.pageSize = pageSize
	c.pageToken = pageToken
	c.view = view
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of documents to return per page. If not specified, 100 documents are
// returned per page.
func (c *ProjectsIndexesDocumentsListCall) PageSize(pageSize int64) *ProjectsIndexesDocumentsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": A `nextPageToken`
// returned from previous list documents call as the starting point for
// this call. If not specified, list documents from the beginning.
func (c *ProjectsIndexesDocumentsListCall) PageToken(pageToken string) *ProjectsIndexesDocumentsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// View sets the optional parameter "view": Specifies which part of the
// document resource is returned in the response. If not specified,
// `ID_ONLY` is used.
func (c *ProjectsIndexesDocumentsListCall) View(view string) *ProjectsIndexesDocumentsListCall {
	c.opt_["view"] = view
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsIndexesDocumentsListCall) Fields(s ...googleapi.Field) *ProjectsIndexesDocumentsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsIndexesDocumentsListCall) Do() (*ListDocumentsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["view"]; ok {
		params.Set("view", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/indexes/{indexId}/documents")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"indexId":   c.indexId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListDocumentsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists documents in the specified search index. Intended for batch processing.",
	//   "httpMethod": "GET",
	//   "id": "cloudsearch.projects.indexes.documents.list",
	//   "parameterOrder": [
	//     "projectId",
	//     "indexId",
	//     "pageSize",
	//     "pageToken",
	//     "view"
	//   ],
	//   "parameters": {
	//     "indexId": {
	//       "description": "The index from which to list the documents. It cannot be the empty string.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of documents to return per page. If not specified, 100 documents are returned per page.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A `nextPageToken` returned from previous list documents call as the starting point for this call. If not specified, list documents from the beginning.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The project associated with the index for listing documents. It cannot be the empty string.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "view": {
	//       "description": "Specifies which part of the document resource is returned in the response. If not specified, `ID_ONLY` is used.",
	//       "enum": [
	//         "DOCUMENT_VIEW_UNSPECIFIED",
	//         "ID_ONLY",
	//         "FULL"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/indexes/{indexId}/documents",
	//   "response": {
	//     "$ref": "ListDocumentsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloudsearch",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}
