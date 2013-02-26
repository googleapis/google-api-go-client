// Package analytics provides access to the Google Analytics API.
//
// See https://developers.google.com/analytics/
//
// Usage example:
//
//   import "code.google.com/p/google-api-go-client/analytics/v3"
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

const apiId = "analytics:v3"
const apiName = "analytics"
const apiVersion = "v3"
const basePath = "https://www.googleapis.com/analytics/v3/"

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

type Account struct {
	// ChildLink: Child link for an account entry. Points to the list of web
	// properties for this account.
	ChildLink *AccountChildLink `json:"childLink,omitempty"`

	// Created: Time the account was created.
	Created string `json:"created,omitempty"`

	// Id: Account ID.
	Id string `json:"id,omitempty"`

	// Kind: Resource type for Analytics account.
	Kind string `json:"kind,omitempty"`

	// Name: Account name.
	Name string `json:"name,omitempty"`

	// SelfLink: Link for this account.
	SelfLink string `json:"selfLink,omitempty"`

	// Updated: Time the account was last modified.
	Updated string `json:"updated,omitempty"`
}

type AccountChildLink struct {
	// Href: Link to the list of web properties for this account.
	Href string `json:"href,omitempty"`

	// Type: Type of the child link. Its value is "analytics#webproperties".
	Type string `json:"type,omitempty"`
}

type Accounts struct {
	// Items: A list of accounts.
	Items []*Account `json:"items,omitempty"`

	// ItemsPerPage: The maximum number of entries the response can contain,
	// regardless of the actual number of entries returned. Its value ranges
	// from 1 to 1000 with a value of 1000 by default, or otherwise
	// specified by the max-results query parameter.
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`

	// Kind: Collection type.
	Kind string `json:"kind,omitempty"`

	// NextLink: Next link for this account collection.
	NextLink string `json:"nextLink,omitempty"`

	// PreviousLink: Previous link for this account collection.
	PreviousLink string `json:"previousLink,omitempty"`

	// StartIndex: The starting index of the entries, which is 1 by default
	// or otherwise specified by the start-index query parameter.
	StartIndex int64 `json:"startIndex,omitempty"`

	// TotalResults: The total number of results for the query, regardless
	// of the number of results in the response.
	TotalResults int64 `json:"totalResults,omitempty"`

	// Username: Email ID of the authenticated user
	Username string `json:"username,omitempty"`
}

type CustomDataSource struct {
	// AccountId: Account ID to which this custom data source belongs.
	AccountId string `json:"accountId,omitempty"`

	// ChildLink: Child link for this custom data source. Points to the list
	// of daily uploads for this custom data source.
	ChildLink *CustomDataSourceChildLink `json:"childLink,omitempty"`

	// Created: Time this custom data source was created.
	Created string `json:"created,omitempty"`

	// Description: Description of custom data source.
	Description string `json:"description,omitempty"`

	// Id: Custom data source ID.
	Id string `json:"id,omitempty"`

	// Kind: Resource type for Analytics custom data source.
	Kind string `json:"kind,omitempty"`

	// Name: Name of this custom data source.
	Name string `json:"name,omitempty"`

	// ParentLink: Parent link for this custom data source. Points to the
	// web property to which this custom data source belongs.
	ParentLink *CustomDataSourceParentLink `json:"parentLink,omitempty"`

	ProfilesLinked []string `json:"profilesLinked,omitempty"`

	// SelfLink: Link for this Analytics custom data source.
	SelfLink string `json:"selfLink,omitempty"`

	// Updated: Time this custom data source was last modified.
	Updated string `json:"updated,omitempty"`

	// WebPropertyId: Web property ID of the form UA-XXXXX-YY to which this
	// custom data source belongs.
	WebPropertyId string `json:"webPropertyId,omitempty"`
}

type CustomDataSourceChildLink struct {
	// Href: Link to the list of daily uploads for this custom data source.
	Href string `json:"href,omitempty"`

	// Type: Value is "analytics#dailyUploads".
	Type string `json:"type,omitempty"`
}

type CustomDataSourceParentLink struct {
	// Href: Link to the web property to which this custom data source
	// belongs.
	Href string `json:"href,omitempty"`

	// Type: Value is "analytics#webproperty".
	Type string `json:"type,omitempty"`
}

type CustomDataSources struct {
	// Items: Collection of custom data sources.
	Items []*CustomDataSource `json:"items,omitempty"`

	// ItemsPerPage: The maximum number of resources the response can
	// contain, regardless of the actual number of resources returned. Its
	// value ranges from 1 to 1000 with a value of 1000 by default, or
	// otherwise specified by the max-results query parameter.
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`

	// Kind: Collection type.
	Kind string `json:"kind,omitempty"`

	// NextLink: Link to next page for this custom data source collection.
	NextLink string `json:"nextLink,omitempty"`

	// PreviousLink: Link to previous page for this custom data source
	// collection.
	PreviousLink string `json:"previousLink,omitempty"`

	// StartIndex: The starting index of the resources, which is 1 by
	// default or otherwise specified by the start-index query parameter.
	StartIndex int64 `json:"startIndex,omitempty"`

	// TotalResults: The total number of results for the query, regardless
	// of the number of results in the response.
	TotalResults int64 `json:"totalResults,omitempty"`

	// Username: Email ID of the authenticated user
	Username string `json:"username,omitempty"`
}

type DailyUpload struct {
	// AccountId: Account ID to which this daily upload belongs.
	AccountId string `json:"accountId,omitempty"`

	// AppendCount: Number of appends for this date.
	AppendCount int64 `json:"appendCount,omitempty"`

	// CreatedTime: Time this daily upload was created.
	CreatedTime string `json:"createdTime,omitempty"`

	// CustomDataSourceId: Custom data source ID to which this daily upload
	// belongs.
	CustomDataSourceId string `json:"customDataSourceId,omitempty"`

	// Date: Date associated with daily upload.
	Date string `json:"date,omitempty"`

	// Kind: Resource type for Analytics daily upload.
	Kind string `json:"kind,omitempty"`

	// ModifiedTime: Time this daily upload was last modified.
	ModifiedTime string `json:"modifiedTime,omitempty"`

	// ParentLink: Parent link for a daily upload. Points to the custom data
	// source to which this daily upload belongs.
	ParentLink *DailyUploadParentLink `json:"parentLink,omitempty"`

	// RecentChanges: Change log for last 10 changes in chronological order.
	RecentChanges []*DailyUploadRecentChanges `json:"recentChanges,omitempty"`

	// SelfLink: Link for this daily upload.
	SelfLink string `json:"selfLink,omitempty"`

	// WebPropertyId: Web property ID of the form UA-XXXXX-YY to which this
	// daily upload belongs.
	WebPropertyId string `json:"webPropertyId,omitempty"`
}

type DailyUploadParentLink struct {
	// Href: Link to the custom data source to which this daily upload
	// belongs.
	Href string `json:"href,omitempty"`

	// Type: Value is "analytics#customDataSource".
	Type string `json:"type,omitempty"`
}

type DailyUploadRecentChanges struct {
	Change string `json:"change,omitempty"`

	Time string `json:"time,omitempty"`
}

type DailyUploadAppend struct {
	// AccountId: Account Id to which this daily upload append belongs.
	AccountId string `json:"accountId,omitempty"`

	// AppendNumber: Append number.
	AppendNumber int64 `json:"appendNumber,omitempty"`

	// CustomDataSourceId: Custom data source Id to which this daily upload
	// append belongs.
	CustomDataSourceId string `json:"customDataSourceId,omitempty"`

	// Date: Date associated with daily upload append.
	Date string `json:"date,omitempty"`

	// Kind: Resource type for Analytics daily upload append.
	Kind string `json:"kind,omitempty"`

	NextAppendLink string `json:"nextAppendLink,omitempty"`

	// WebPropertyId: Web property Id of the form UA-XXXXX-YY to which this
	// daily upload append belongs.
	WebPropertyId string `json:"webPropertyId,omitempty"`
}

type DailyUploads struct {
	// Items: A collection of daily uploads.
	Items []*DailyUpload `json:"items,omitempty"`

	// ItemsPerPage: The maximum number of resources the response can
	// contain, regardless of the actual number of resources returned. Its
	// value ranges from 1 to 1000 with a value of 1000 by default, or
	// otherwise specified by the max-results query parameter.
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`

	// Kind: Collection type. Value is analytics#dailyUploads.
	Kind string `json:"kind,omitempty"`

	// NextLink: Link to next page for this daily upload collection.
	NextLink string `json:"nextLink,omitempty"`

	// PreviousLink: Link to previous page for this daily upload collection.
	PreviousLink string `json:"previousLink,omitempty"`

	// StartIndex: The starting index of the resources, which is 1 by
	// default or otherwise specified by the start-index query parameter.
	StartIndex int64 `json:"startIndex,omitempty"`

	// TotalResults: The total number of results for the query, regardless
	// of the number of results in the response.
	TotalResults int64 `json:"totalResults,omitempty"`

	// Username: Email ID of the authenticated user
	Username string `json:"username,omitempty"`
}

type GaData struct {
	// ColumnHeaders: Column headers that list dimension names followed by
	// the metric names. The order of dimensions and metrics is same as
	// specified in the request.
	ColumnHeaders []*GaDataColumnHeaders `json:"columnHeaders,omitempty"`

	// ContainsSampledData: Determines if Analytics data contains samples.
	ContainsSampledData bool `json:"containsSampledData,omitempty"`

	// Id: Unique ID for this data response.
	Id string `json:"id,omitempty"`

	// ItemsPerPage: The maximum number of rows the response can contain,
	// regardless of the actual number of rows returned. Its value ranges
	// from 1 to 10,000 with a value of 1000 by default, or otherwise
	// specified by the max-results query parameter.
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`

	// NextLink: Link to next page for this Analytics data query.
	NextLink string `json:"nextLink,omitempty"`

	// PreviousLink: Link to previous page for this Analytics data query.
	PreviousLink string `json:"previousLink,omitempty"`

	// ProfileInfo: Information for the profile, for which the Analytics
	// data was requested.
	ProfileInfo *GaDataProfileInfo `json:"profileInfo,omitempty"`

	// Query: Analytics data request query parameters.
	Query *GaDataQuery `json:"query,omitempty"`

	// Rows: Analytics data rows, where each row contains a list of
	// dimension values followed by the metric values. The order of
	// dimensions and metrics is same as specified in the request.
	Rows [][]string `json:"rows,omitempty"`

	// SelfLink: Link to this page.
	SelfLink string `json:"selfLink,omitempty"`

	// TotalResults: The total number of rows for the query, regardless of
	// the number of rows in the response.
	TotalResults int64 `json:"totalResults,omitempty"`

	// TotalsForAllResults: Total values for the requested metrics over all
	// the results, not just the results returned in this response. The
	// order of the metric totals is same as the metric order specified in
	// the request.
	TotalsForAllResults *GaDataTotalsForAllResults `json:"totalsForAllResults,omitempty"`
}

type GaDataColumnHeaders struct {
	// ColumnType: Column Type. Either DIMENSION or METRIC.
	ColumnType string `json:"columnType,omitempty"`

	// DataType: Data type. Dimension column headers have only STRING as the
	// data type. Metric column headers have data types for metric values
	// such as INTEGER, DOUBLE, CURRENCY etc.
	DataType string `json:"dataType,omitempty"`

	// Name: Column name.
	Name string `json:"name,omitempty"`
}

type GaDataProfileInfo struct {
	// AccountId: Account ID to which this profile belongs.
	AccountId string `json:"accountId,omitempty"`

	// InternalWebPropertyId: Internal ID for the web property to which this
	// profile belongs.
	InternalWebPropertyId string `json:"internalWebPropertyId,omitempty"`

	// ProfileId: Profile ID.
	ProfileId string `json:"profileId,omitempty"`

	// ProfileName: Profile name.
	ProfileName string `json:"profileName,omitempty"`

	// TableId: Table ID for profile.
	TableId string `json:"tableId,omitempty"`

	// WebPropertyId: Web Property ID to which this profile belongs.
	WebPropertyId string `json:"webPropertyId,omitempty"`
}

type GaDataQuery struct {
	// Dimensions: List of analytics dimensions.
	Dimensions string `json:"dimensions,omitempty"`

	// EndDate: End date.
	EndDate string `json:"end-date,omitempty"`

	// Filters: Comma-separated list of dimension or metric filters.
	Filters string `json:"filters,omitempty"`

	// Ids: Unique table ID.
	Ids string `json:"ids,omitempty"`

	// MaxResults: Maximum results per page.
	MaxResults int64 `json:"max-results,omitempty"`

	// Metrics: List of analytics metrics.
	Metrics []string `json:"metrics,omitempty"`

	// Segment: Analytics advanced segment.
	Segment string `json:"segment,omitempty"`

	// Sort: List of dimensions or metrics based on which Analytics data is
	// sorted.
	Sort []string `json:"sort,omitempty"`

	// StartDate: Start date.
	StartDate string `json:"start-date,omitempty"`

	// StartIndex: Start index.
	StartIndex int64 `json:"start-index,omitempty"`
}

type GaDataTotalsForAllResults struct {
}

type Goal struct {
	// AccountId: Account ID to which this goal belongs.
	AccountId string `json:"accountId,omitempty"`

	// Active: Determines whether this goal is active.
	Active bool `json:"active,omitempty"`

	// Created: Time this goal was created.
	Created string `json:"created,omitempty"`

	// EventDetails: Details for the goal of the type EVENT.
	EventDetails *GoalEventDetails `json:"eventDetails,omitempty"`

	// Id: Goal ID.
	Id string `json:"id,omitempty"`

	// InternalWebPropertyId: Internal ID for the web property to which this
	// goal belongs.
	InternalWebPropertyId string `json:"internalWebPropertyId,omitempty"`

	// Kind: Resource type for an Analytics goal.
	Kind string `json:"kind,omitempty"`

	// Name: Goal name.
	Name string `json:"name,omitempty"`

	// ParentLink: Parent link for a goal. Points to the profile to which
	// this goal belongs.
	ParentLink *GoalParentLink `json:"parentLink,omitempty"`

	// ProfileId: Profile ID to which this goal belongs.
	ProfileId string `json:"profileId,omitempty"`

	// SelfLink: Link for this goal.
	SelfLink string `json:"selfLink,omitempty"`

	// Type: Goal type. Possible values are URL_DESTINATION,
	// VISIT_TIME_ON_SITE, VISIT_NUM_PAGES, AND EVENT.
	Type string `json:"type,omitempty"`

	// Updated: Time this goal was last modified.
	Updated string `json:"updated,omitempty"`

	// UrlDestinationDetails: Details for the goal of the type
	// URL_DESTINATION.
	UrlDestinationDetails *GoalUrlDestinationDetails `json:"urlDestinationDetails,omitempty"`

	// Value: Goal value.
	Value float64 `json:"value,omitempty"`

	// VisitNumPagesDetails: Details for the goal of the type
	// VISIT_NUM_PAGES.
	VisitNumPagesDetails *GoalVisitNumPagesDetails `json:"visitNumPagesDetails,omitempty"`

	// VisitTimeOnSiteDetails: Details for the goal of the type
	// VISIT_TIME_ON_SITE.
	VisitTimeOnSiteDetails *GoalVisitTimeOnSiteDetails `json:"visitTimeOnSiteDetails,omitempty"`

	// WebPropertyId: Web property ID to which this goal belongs. The web
	// property ID is of the form UA-XXXXX-YY.
	WebPropertyId string `json:"webPropertyId,omitempty"`
}

type GoalEventDetails struct {
	// EventConditions: List of event conditions.
	EventConditions []*GoalEventDetailsEventConditions `json:"eventConditions,omitempty"`

	// UseEventValue: Determines if the event value should be used as the
	// value for this goal.
	UseEventValue bool `json:"useEventValue,omitempty"`
}

type GoalEventDetailsEventConditions struct {
	// ComparisonType: Type of comparison. Possible values are LESS_THAN,
	// GREATER_THAN or EQUAL.
	ComparisonType string `json:"comparisonType,omitempty"`

	// ComparisonValue: Value used for this comparison.
	ComparisonValue int64 `json:"comparisonValue,omitempty,string"`

	// Expression: Expression used for this match.
	Expression string `json:"expression,omitempty"`

	// MatchType: Type of the match to be performed. Possible values are
	// REGEXP, BEGINS_WITH, or EXACT.
	MatchType string `json:"matchType,omitempty"`

	// Type: Type of this event condition. Possible values are CATEGORY,
	// ACTION, LABEL, or VALUE.
	Type string `json:"type,omitempty"`
}

type GoalParentLink struct {
	// Href: Link to the profile to which this goal belongs.
	Href string `json:"href,omitempty"`

	// Type: Value is "analytics#profile".
	Type string `json:"type,omitempty"`
}

type GoalUrlDestinationDetails struct {
	// CaseSensitive: Determines if the goal URL must exactly match the
	// capitalization of visited URLs.
	CaseSensitive bool `json:"caseSensitive,omitempty"`

	// FirstStepRequired: Determines if the first step in this goal is
	// required.
	FirstStepRequired bool `json:"firstStepRequired,omitempty"`

	// MatchType: Match type for the goal URL. Possible values are HEAD,
	// EXACT, or REGEX.
	MatchType string `json:"matchType,omitempty"`

	// Steps: List of steps configured for this goal funnel.
	Steps []*GoalUrlDestinationDetailsSteps `json:"steps,omitempty"`

	// Url: URL for this goal.
	Url string `json:"url,omitempty"`
}

type GoalUrlDestinationDetailsSteps struct {
	// Name: Step name.
	Name string `json:"name,omitempty"`

	// Number: Step number.
	Number int64 `json:"number,omitempty"`

	// Url: URL for this step.
	Url string `json:"url,omitempty"`
}

type GoalVisitNumPagesDetails struct {
	// ComparisonType: Type of comparison. Possible values are LESS_THAN,
	// GREATER_THAN, or EQUAL.
	ComparisonType string `json:"comparisonType,omitempty"`

	// ComparisonValue: Value used for this comparison.
	ComparisonValue int64 `json:"comparisonValue,omitempty,string"`
}

type GoalVisitTimeOnSiteDetails struct {
	// ComparisonType: Type of comparison. Possible values are LESS_THAN or
	// GREATER_THAN.
	ComparisonType string `json:"comparisonType,omitempty"`

	// ComparisonValue: Value used for this comparison.
	ComparisonValue int64 `json:"comparisonValue,omitempty,string"`
}

type Goals struct {
	// Items: A list of goals.
	Items []*Goal `json:"items,omitempty"`

	// ItemsPerPage: The maximum number of resources the response can
	// contain, regardless of the actual number of resources returned. Its
	// value ranges from 1 to 1000 with a value of 1000 by default, or
	// otherwise specified by the max-results query parameter.
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`

	// Kind: Collection type.
	Kind string `json:"kind,omitempty"`

	// NextLink: Link to next page for this goal collection.
	NextLink string `json:"nextLink,omitempty"`

	// PreviousLink: Link to previous page for this goal collection.
	PreviousLink string `json:"previousLink,omitempty"`

	// StartIndex: The starting index of the resources, which is 1 by
	// default or otherwise specified by the start-index query parameter.
	StartIndex int64 `json:"startIndex,omitempty"`

	// TotalResults: The total number of results for the query, regardless
	// of the number of resources in the result.
	TotalResults int64 `json:"totalResults,omitempty"`

	// Username: Email ID of the authenticated user
	Username string `json:"username,omitempty"`
}

type McfData struct {
	// ColumnHeaders: Column headers that list dimension names followed by
	// the metric names. The order of dimensions and metrics is same as
	// specified in the request.
	ColumnHeaders []*McfDataColumnHeaders `json:"columnHeaders,omitempty"`

	// ContainsSampledData: Determines if the Analytics data contains
	// sampled data.
	ContainsSampledData bool `json:"containsSampledData,omitempty"`

	// Id: Unique ID for this data response.
	Id string `json:"id,omitempty"`

	// ItemsPerPage: The maximum number of rows the response can contain,
	// regardless of the actual number of rows returned. Its value ranges
	// from 1 to 10,000 with a value of 1000 by default, or otherwise
	// specified by the max-results query parameter.
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`

	// NextLink: Link to next page for this Analytics data query.
	NextLink string `json:"nextLink,omitempty"`

	// PreviousLink: Link to previous page for this Analytics data query.
	PreviousLink string `json:"previousLink,omitempty"`

	// ProfileInfo: Information for the profile, for which the Analytics
	// data was requested.
	ProfileInfo *McfDataProfileInfo `json:"profileInfo,omitempty"`

	// Query: Analytics data request query parameters.
	Query *McfDataQuery `json:"query,omitempty"`

	// Rows: Analytics data rows, where each row contains a list of
	// dimension values followed by the metric values. The order of
	// dimensions and metrics is same as specified in the request.
	Rows [][]*McfDataRowsItem `json:"rows,omitempty"`

	// SelfLink: Link to this page.
	SelfLink string `json:"selfLink,omitempty"`

	// TotalResults: The total number of rows for the query, regardless of
	// the number of rows in the response.
	TotalResults int64 `json:"totalResults,omitempty"`

	// TotalsForAllResults: Total values for the requested metrics over all
	// the results, not just the results returned in this response. The
	// order of the metric totals is same as the metric order specified in
	// the request.
	TotalsForAllResults *McfDataTotalsForAllResults `json:"totalsForAllResults,omitempty"`
}

type McfDataColumnHeaders struct {
	// ColumnType: Column Type. Either DIMENSION or METRIC.
	ColumnType string `json:"columnType,omitempty"`

	// DataType: Data type. Dimension and metric values data types such as
	// INTEGER, DOUBLE, CURRENCY, MCF_SEQUENCE etc.
	DataType string `json:"dataType,omitempty"`

	// Name: Column name.
	Name string `json:"name,omitempty"`
}

type McfDataProfileInfo struct {
	// AccountId: Account ID to which this profile belongs.
	AccountId string `json:"accountId,omitempty"`

	// InternalWebPropertyId: Internal ID for the web property to which this
	// profile belongs.
	InternalWebPropertyId string `json:"internalWebPropertyId,omitempty"`

	// ProfileId: Profile ID.
	ProfileId string `json:"profileId,omitempty"`

	// ProfileName: Profile name.
	ProfileName string `json:"profileName,omitempty"`

	// TableId: Table ID for profile.
	TableId string `json:"tableId,omitempty"`

	// WebPropertyId: Web Property ID to which this profile belongs.
	WebPropertyId string `json:"webPropertyId,omitempty"`
}

type McfDataQuery struct {
	// Dimensions: List of analytics dimensions.
	Dimensions string `json:"dimensions,omitempty"`

	// EndDate: End date.
	EndDate string `json:"end-date,omitempty"`

	// Filters: Comma-separated list of dimension or metric filters.
	Filters string `json:"filters,omitempty"`

	// Ids: Unique table ID.
	Ids string `json:"ids,omitempty"`

	// MaxResults: Maximum results per page.
	MaxResults int64 `json:"max-results,omitempty"`

	// Metrics: List of analytics metrics.
	Metrics []string `json:"metrics,omitempty"`

	// Segment: Analytics advanced segment.
	Segment string `json:"segment,omitempty"`

	// Sort: List of dimensions or metrics based on which Analytics data is
	// sorted.
	Sort []string `json:"sort,omitempty"`

	// StartDate: Start date.
	StartDate string `json:"start-date,omitempty"`

	// StartIndex: Start index.
	StartIndex int64 `json:"start-index,omitempty"`
}

type McfDataRowsItem struct {
	// ConversionPathValue: A conversion path dimension value, containing a
	// list of interactions with their attributes.
	ConversionPathValue []*McfDataRowsItemConversionPathValue `json:"conversionPathValue,omitempty"`

	// PrimitiveValue: A primitive dimension value. A primitive metric
	// value.
	PrimitiveValue string `json:"primitiveValue,omitempty"`
}

type McfDataRowsItemConversionPathValue struct {
	// InteractionType: Type of an interaction on conversion path. Such as
	// CLICK, IMPRESSION etc.
	InteractionType string `json:"interactionType,omitempty"`

	// NodeValue: Node value of an interaction on conversion path. Such as
	// source, medium etc.
	NodeValue string `json:"nodeValue,omitempty"`
}

type McfDataTotalsForAllResults struct {
}

type Profile struct {
	// AccountId: Account ID to which this profile belongs.
	AccountId string `json:"accountId,omitempty"`

	// ChildLink: Child link for this profile. Points to the list of goals
	// for this profile.
	ChildLink *ProfileChildLink `json:"childLink,omitempty"`

	// Created: Time this profile was created.
	Created string `json:"created,omitempty"`

	// Currency: The currency type associated with this profile.
	Currency string `json:"currency,omitempty"`

	// DefaultPage: Default page for this profile.
	DefaultPage string `json:"defaultPage,omitempty"`

	// ECommerceTracking: E-commerce tracking parameter for this profile.
	ECommerceTracking bool `json:"eCommerceTracking,omitempty"`

	// ExcludeQueryParameters: The query parameters that are excluded from
	// this profile.
	ExcludeQueryParameters string `json:"excludeQueryParameters,omitempty"`

	// Id: Profile ID.
	Id string `json:"id,omitempty"`

	// InternalWebPropertyId: Internal ID for the web property to which this
	// profile belongs.
	InternalWebPropertyId string `json:"internalWebPropertyId,omitempty"`

	// Kind: Resource type for Analytics profile.
	Kind string `json:"kind,omitempty"`

	// Name: Name of this profile.
	Name string `json:"name,omitempty"`

	// ParentLink: Parent link for this profile. Points to the web property
	// to which this profile belongs.
	ParentLink *ProfileParentLink `json:"parentLink,omitempty"`

	// SelfLink: Link for this profile.
	SelfLink string `json:"selfLink,omitempty"`

	// SiteSearchCategoryParameters: Site search category parameters for
	// this profile.
	SiteSearchCategoryParameters string `json:"siteSearchCategoryParameters,omitempty"`

	// SiteSearchQueryParameters: The site search query parameters for this
	// profile.
	SiteSearchQueryParameters string `json:"siteSearchQueryParameters,omitempty"`

	// Timezone: Time zone for which this profile has been configured.
	Timezone string `json:"timezone,omitempty"`

	// Type: Profile type. Supported types: WEB or APP.
	Type string `json:"type,omitempty"`

	// Updated: Time this profile was last modified.
	Updated string `json:"updated,omitempty"`

	// WebPropertyId: Web property ID of the form UA-XXXXX-YY to which this
	// profile belongs.
	WebPropertyId string `json:"webPropertyId,omitempty"`

	// WebsiteUrl: Website URL for this profile.
	WebsiteUrl string `json:"websiteUrl,omitempty"`
}

type ProfileChildLink struct {
	// Href: Link to the list of goals for this profile.
	Href string `json:"href,omitempty"`

	// Type: Value is "analytics#goals".
	Type string `json:"type,omitempty"`
}

type ProfileParentLink struct {
	// Href: Link to the web property to which this profile belongs.
	Href string `json:"href,omitempty"`

	// Type: Value is "analytics#webproperty".
	Type string `json:"type,omitempty"`
}

type Profiles struct {
	// Items: A list of profiles.
	Items []*Profile `json:"items,omitempty"`

	// ItemsPerPage: The maximum number of resources the response can
	// contain, regardless of the actual number of resources returned. Its
	// value ranges from 1 to 1000 with a value of 1000 by default, or
	// otherwise specified by the max-results query parameter.
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`

	// Kind: Collection type.
	Kind string `json:"kind,omitempty"`

	// NextLink: Link to next page for this profile collection.
	NextLink string `json:"nextLink,omitempty"`

	// PreviousLink: Link to previous page for this profile collection.
	PreviousLink string `json:"previousLink,omitempty"`

	// StartIndex: The starting index of the resources, which is 1 by
	// default or otherwise specified by the start-index query parameter.
	StartIndex int64 `json:"startIndex,omitempty"`

	// TotalResults: The total number of results for the query, regardless
	// of the number of results in the response.
	TotalResults int64 `json:"totalResults,omitempty"`

	// Username: Email ID of the authenticated user
	Username string `json:"username,omitempty"`
}

type Segment struct {
	// Created: Time the advanced segment was created.
	Created string `json:"created,omitempty"`

	// Definition: Advanced segment definition.
	Definition string `json:"definition,omitempty"`

	// Id: Advanced segment ID.
	Id string `json:"id,omitempty"`

	// Kind: Resource type for Analytics advanced segment.
	Kind string `json:"kind,omitempty"`

	// Name: Advanced segment name.
	Name string `json:"name,omitempty"`

	// SegmentId: Segment ID. Can be used with the 'segment' parameter in
	// Data Feed.
	SegmentId string `json:"segmentId,omitempty"`

	// SelfLink: Link for this advanced segment.
	SelfLink string `json:"selfLink,omitempty"`

	// Updated: Time the advanced segment was last modified.
	Updated string `json:"updated,omitempty"`
}

type Segments struct {
	// Items: A list of advanced segments.
	Items []*Segment `json:"items,omitempty"`

	// ItemsPerPage: The maximum number of resources the response can
	// contain, regardless of the actual number of resources returned. Its
	// value ranges from 1 to 1000 with a value of 1000 by default, or
	// otherwise specified by the max-results query parameter.
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`

	// Kind: Collection type for advanced segments.
	Kind string `json:"kind,omitempty"`

	// NextLink: Link to next page for this advanced segment collection.
	NextLink string `json:"nextLink,omitempty"`

	// PreviousLink: Link to previous page for this advanced segment
	// collection.
	PreviousLink string `json:"previousLink,omitempty"`

	// StartIndex: The starting index of the resources, which is 1 by
	// default or otherwise specified by the start-index query parameter.
	StartIndex int64 `json:"startIndex,omitempty"`

	// TotalResults: The total number of results for the query, regardless
	// of the number of results in the response.
	TotalResults int64 `json:"totalResults,omitempty"`

	// Username: Email ID of the authenticated user
	Username string `json:"username,omitempty"`
}

type Webproperties struct {
	// Items: A list of web properties.
	Items []*Webproperty `json:"items,omitempty"`

	// ItemsPerPage: The maximum number of resources the response can
	// contain, regardless of the actual number of resources returned. Its
	// value ranges from 1 to 1000 with a value of 1000 by default, or
	// otherwise specified by the max-results query parameter.
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`

	// Kind: Collection type.
	Kind string `json:"kind,omitempty"`

	// NextLink: Link to next page for this web property collection.
	NextLink string `json:"nextLink,omitempty"`

	// PreviousLink: Link to previous page for this web property collection.
	PreviousLink string `json:"previousLink,omitempty"`

	// StartIndex: The starting index of the resources, which is 1 by
	// default or otherwise specified by the start-index query parameter.
	StartIndex int64 `json:"startIndex,omitempty"`

	// TotalResults: The total number of results for the query, regardless
	// of the number of results in the response.
	TotalResults int64 `json:"totalResults,omitempty"`

	// Username: Email ID of the authenticated user
	Username string `json:"username,omitempty"`
}

type Webproperty struct {
	// AccountId: Account ID to which this web property belongs.
	AccountId string `json:"accountId,omitempty"`

	// ChildLink: Child link for this web property. Points to the list of
	// profiles for this web property.
	ChildLink *WebpropertyChildLink `json:"childLink,omitempty"`

	// Created: Time this web property was created.
	Created string `json:"created,omitempty"`

	// Id: Web property ID of the form UA-XXXXX-YY.
	Id string `json:"id,omitempty"`

	// InternalWebPropertyId: Internal ID for this web property.
	InternalWebPropertyId string `json:"internalWebPropertyId,omitempty"`

	// Kind: Resource type for Analytics WebProperty.
	Kind string `json:"kind,omitempty"`

	// Name: Name of this web property.
	Name string `json:"name,omitempty"`

	// ParentLink: Parent link for this web property. Points to the account
	// to which this web property belongs.
	ParentLink *WebpropertyParentLink `json:"parentLink,omitempty"`

	// SelfLink: Link for this web property.
	SelfLink string `json:"selfLink,omitempty"`

	// Updated: Time this web property was last modified.
	Updated string `json:"updated,omitempty"`

	// WebsiteUrl: Website url for this web property.
	WebsiteUrl string `json:"websiteUrl,omitempty"`
}

type WebpropertyChildLink struct {
	// Href: Link to the list of profiles for this web property.
	Href string `json:"href,omitempty"`

	// Type: Type of the parent link. Its value is "analytics#profiles".
	Type string `json:"type,omitempty"`
}

type WebpropertyParentLink struct {
	// Href: Link to the account for this web property.
	Href string `json:"href,omitempty"`

	// Type: Type of the parent link. Its value is "analytics#account".
	Type string `json:"type,omitempty"`
}

func cleanPathString(s string) string {
	return strings.Map(func(r rune) rune {
		if r >= 0x2d && r <= 0x7a || r == '~' {
			return r
		}
		return -1
	}, s)
}
