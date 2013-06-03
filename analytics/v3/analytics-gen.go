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
	s.Data = NewDataService(s)
	s.Management = NewManagementService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Data *DataService

	Management *ManagementService
}

func NewDataService(s *Service) *DataService {
	rs := &DataService{s: s}
	rs.Ga = NewDataGaService(s)
	rs.Mcf = NewDataMcfService(s)
	return rs
}

type DataService struct {
	s *Service

	Ga *DataGaService

	Mcf *DataMcfService
}

func NewDataGaService(s *Service) *DataGaService {
	rs := &DataGaService{s: s}
	return rs
}

type DataGaService struct {
	s *Service
}

func NewDataMcfService(s *Service) *DataMcfService {
	rs := &DataMcfService{s: s}
	return rs
}

type DataMcfService struct {
	s *Service
}

func NewManagementService(s *Service) *ManagementService {
	rs := &ManagementService{s: s}
	rs.Accounts = NewManagementAccountsService(s)
	rs.CustomDataSources = NewManagementCustomDataSourcesService(s)
	rs.DailyUploads = NewManagementDailyUploadsService(s)
	rs.Experiments = NewManagementExperimentsService(s)
	rs.Goals = NewManagementGoalsService(s)
	rs.Profiles = NewManagementProfilesService(s)
	rs.Segments = NewManagementSegmentsService(s)
	rs.Webproperties = NewManagementWebpropertiesService(s)
	return rs
}

type ManagementService struct {
	s *Service

	Accounts *ManagementAccountsService

	CustomDataSources *ManagementCustomDataSourcesService

	DailyUploads *ManagementDailyUploadsService

	Experiments *ManagementExperimentsService

	Goals *ManagementGoalsService

	Profiles *ManagementProfilesService

	Segments *ManagementSegmentsService

	Webproperties *ManagementWebpropertiesService
}

func NewManagementAccountsService(s *Service) *ManagementAccountsService {
	rs := &ManagementAccountsService{s: s}
	return rs
}

type ManagementAccountsService struct {
	s *Service
}

func NewManagementCustomDataSourcesService(s *Service) *ManagementCustomDataSourcesService {
	rs := &ManagementCustomDataSourcesService{s: s}
	return rs
}

type ManagementCustomDataSourcesService struct {
	s *Service
}

func NewManagementDailyUploadsService(s *Service) *ManagementDailyUploadsService {
	rs := &ManagementDailyUploadsService{s: s}
	return rs
}

type ManagementDailyUploadsService struct {
	s *Service
}

func NewManagementExperimentsService(s *Service) *ManagementExperimentsService {
	rs := &ManagementExperimentsService{s: s}
	return rs
}

type ManagementExperimentsService struct {
	s *Service
}

func NewManagementGoalsService(s *Service) *ManagementGoalsService {
	rs := &ManagementGoalsService{s: s}
	return rs
}

type ManagementGoalsService struct {
	s *Service
}

func NewManagementProfilesService(s *Service) *ManagementProfilesService {
	rs := &ManagementProfilesService{s: s}
	return rs
}

type ManagementProfilesService struct {
	s *Service
}

func NewManagementSegmentsService(s *Service) *ManagementSegmentsService {
	rs := &ManagementSegmentsService{s: s}
	return rs
}

type ManagementSegmentsService struct {
	s *Service
}

func NewManagementWebpropertiesService(s *Service) *ManagementWebpropertiesService {
	rs := &ManagementWebpropertiesService{s: s}
	return rs
}

type ManagementWebpropertiesService struct {
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

type Experiment struct {
	// AccountId: Account ID to which this experiment belongs. This field is
	// read-only.
	AccountId string `json:"accountId,omitempty"`

	// Created: Time the experiment was created. This field is read-only.
	Created string `json:"created,omitempty"`

	// Description: Notes about this experiment.
	Description string `json:"description,omitempty"`

	// EditableInGaUi: If true, the end user will be able to edit the
	// experiment via the Google Analytics user interface.
	EditableInGaUi interface{} `json:"editableInGaUi,omitempty"`

	// EndTime: The ending time of the experiment (the time the status
	// changed from RUNNING to ENDED). This field is present only if the
	// experiment has ended. This field is read-only.
	EndTime string `json:"endTime,omitempty"`

	// Id: Experiment ID. Required for patch and update. Disallowed for
	// create.
	Id string `json:"id,omitempty"`

	// InternalWebPropertyId: Internal ID for the web property to which this
	// experiment belongs. This field is read-only.
	InternalWebPropertyId string `json:"internalWebPropertyId,omitempty"`

	// Kind: Resource type for an Analytics experiment. This field is
	// read-only.
	Kind string `json:"kind,omitempty"`

	// MinimumExperimentLengthInDays: Specifies the minimum length of the
	// experiment. Can be changed for a running experiment. This field may
	// not be changed for an experiments whose status is ENDED.
	MinimumExperimentLengthInDays int64 `json:"minimumExperimentLengthInDays,omitempty"`

	// Name: Experiment name. This field may not be changed for an
	// experiment whose status is ENDED. This field is required when
	// creating an experiment.
	Name string `json:"name,omitempty"`

	// ObjectiveMetric: The metric that the experiment is optimizing. Valid
	// values: "ga:goal(n)Completions", "ga:bounces", "ga:pageviews",
	// "ga:timeOnSite", "ga:transactions", "ga:transactionRevenue". This
	// field is required if status is "RUNNING" and servingFramework is one
	// of "REDIRECT" or "API".
	ObjectiveMetric string `json:"objectiveMetric,omitempty"`

	// OptimizationType: Whether the objectiveMetric should be minimized or
	// maximized. Possible values: "MAXIMUM", "MINIMUM". Optional--defaults
	// to "MAXIMUM". Cannot be specified without objectiveMetric. Cannot be
	// modified when status is "RUNNING" or "ENDED".
	OptimizationType string `json:"optimizationType,omitempty"`

	// ParentLink: Parent link for an experiment. Points to the profile to
	// which this experiment belongs.
	ParentLink *ExperimentParentLink `json:"parentLink,omitempty"`

	// ProfileId: Profile ID to which this experiment belongs. This field is
	// read-only.
	ProfileId string `json:"profileId,omitempty"`

	// ReasonExperimentEnded: Why the experiment ended. Possible values:
	// "STOPPED_BY_USER", "WINNER_FOUND", "EXPERIMENT_EXPIRED",
	// "ENDED_WITH_NO_WINNER", "GOAL_OBJECTIVE_CHANGED".
	// "ENDED_WITH_NO_WINNER" means that the experiment didn't expire but no
	// winner was projected to be found. If the experiment status is changed
	// via the API to ENDED this field is set to STOPPED_BY_USER. This field
	// is read-only.
	ReasonExperimentEnded string `json:"reasonExperimentEnded,omitempty"`

	// RewriteVariationUrlsAsOriginal: Boolean specifying whether variations
	// URLS are rewritten to match those of the original. This field may not
	// be changed for an experiments whose status is ENDED.
	RewriteVariationUrlsAsOriginal bool `json:"rewriteVariationUrlsAsOriginal,omitempty"`

	// SelfLink: Link for this experiment. This field is read-only.
	SelfLink string `json:"selfLink,omitempty"`

	// ServingFramework: The framework used to serve the experiment
	// variations and evaluate the results. One of:
	// - REDIRECT: Google
	// Analytics redirects traffic to different variation pages, reports the
	// chosen variation and evaluates the results.
	// - API: Google Analytics
	// chooses and reports the variation to serve and evaluates the results;
	// the caller is responsible for serving the selected variation.
	// -
	// EXTERNAL: The variations will be served externally and the chosen
	// variation reported to Google Analytics. The caller is responsible for
	// serving the selected variation and evaluating the results.
	ServingFramework interface{} `json:"servingFramework,omitempty"`

	// Snippet: The snippet of code to include on the control page(s). This
	// field is read-only.
	Snippet string `json:"snippet,omitempty"`

	// StartTime: The starting time of the experiment (the time the status
	// changed from READY_TO_RUN to RUNNING). This field is present only if
	// the experiment has started. This field is read-only.
	StartTime string `json:"startTime,omitempty"`

	// Status: Experiment status. Possible values: "DRAFT", "READY_TO_RUN",
	// "RUNNING", "ENDED". Experiments can be created in the "DRAFT",
	// "READY_TO_RUN" or "RUNNING" state. This field is required when
	// creating an experiment.
	Status string `json:"status,omitempty"`

	// TrafficCoverage: A floating-point number between 0 and 1. Specifies
	// the fraction of the traffic that participates in the experiment. Can
	// be changed for a running experiment. This field may not be changed
	// for an experiments whose status is ENDED.
	TrafficCoverage float64 `json:"trafficCoverage,omitempty"`

	// Updated: Time the experiment was last modified. This field is
	// read-only.
	Updated string `json:"updated,omitempty"`

	// Variations: Array of variations. The first variation in the array is
	// the original. The number of variations may not change once an
	// experiment is in the RUNNING state. At least two variations are
	// required before status can be set to RUNNING.
	Variations []*ExperimentVariations `json:"variations,omitempty"`

	// WebPropertyId: Web property ID to which this experiment belongs. The
	// web property ID is of the form UA-XXXXX-YY. This field is read-only.
	WebPropertyId string `json:"webPropertyId,omitempty"`

	// WinnerConfidenceLevel: A floating-point number between 0 and 1.
	// Specifies the necessary confidence level to choose a winner. This
	// field may not be changed for an experiments whose status is ENDED.
	WinnerConfidenceLevel float64 `json:"winnerConfidenceLevel,omitempty"`

	// WinnerFound: Boolean specifying whether a winner has been found for
	// this experiment. This field is read-only.
	WinnerFound bool `json:"winnerFound,omitempty"`
}

type ExperimentParentLink struct {
	// Href: Link to the profile to which this experiment belongs. This
	// field is read-only.
	Href string `json:"href,omitempty"`

	// Type: Value is "analytics#profile". This field is read-only.
	Type string `json:"type,omitempty"`
}

type ExperimentVariations struct {
	// Name: The name of the variation. This field is required when creating
	// an experiment. This field may not be changed for an experiment whose
	// status is ENDED.
	Name string `json:"name,omitempty"`

	// Status: Status of the variation. Possible values: "ACTIVE",
	// "INACTIVE". INACTIVE variations are not served. This field may not be
	// changed for an experiment whose status is ENDED.
	Status string `json:"status,omitempty"`

	// Url: The URL of the variation. This field may not be changed for an
	// experiment whose status is RUNNING or ENDED.
	Url string `json:"url,omitempty"`

	// Weight: Weight that this variation should receive. Only present if
	// the experiment is running. This field is read-only.
	Weight float64 `json:"weight,omitempty"`

	// Won: True if the experiment has ended and this variation performed
	// (statistically) significantly better than the original. This field is
	// read-only.
	Won bool `json:"won,omitempty"`
}

type Experiments struct {
	// Items: A list of experiments.
	Items []*Experiment `json:"items,omitempty"`

	// ItemsPerPage: The maximum number of resources the response can
	// contain, regardless of the actual number of resources returned. Its
	// value ranges from 1 to 1000 with a value of 1000 by default, or
	// otherwise specified by the max-results query parameter.
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`

	// Kind: Collection type.
	Kind string `json:"kind,omitempty"`

	// NextLink: Link to next page for this experiment collection.
	NextLink string `json:"nextLink,omitempty"`

	// PreviousLink: Link to previous page for this experiment collection.
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

// method id "analytics.data.ga.get":

type DataGaGetCall struct {
	s         *Service
	ids       string
	startDate string
	endDate   string
	metrics   string
	opt_      map[string]interface{}
}

// Get: Returns Analytics data for a profile.
func (r *DataGaService) Get(ids string, startDate string, endDate string, metrics string) *DataGaGetCall {
	c := &DataGaGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.ids = ids
	c.startDate = startDate
	c.endDate = endDate
	c.metrics = metrics
	return c
}

// Dimensions sets the optional parameter "dimensions": A
// comma-separated list of Analytics dimensions. E.g.,
// 'ga:browser,ga:city'.
func (c *DataGaGetCall) Dimensions(dimensions string) *DataGaGetCall {
	c.opt_["dimensions"] = dimensions
	return c
}

// Filters sets the optional parameter "filters": A comma-separated list
// of dimension or metric filters to be applied to Analytics data.
func (c *DataGaGetCall) Filters(filters string) *DataGaGetCall {
	c.opt_["filters"] = filters
	return c
}

// MaxResults sets the optional parameter "max-results": The maximum
// number of entries to include in this feed.
func (c *DataGaGetCall) MaxResults(maxResults int64) *DataGaGetCall {
	c.opt_["max-results"] = maxResults
	return c
}

// Segment sets the optional parameter "segment": An Analytics advanced
// segment to be applied to data.
func (c *DataGaGetCall) Segment(segment string) *DataGaGetCall {
	c.opt_["segment"] = segment
	return c
}

// Sort sets the optional parameter "sort": A comma-separated list of
// dimensions or metrics that determine the sort order for Analytics
// data.
func (c *DataGaGetCall) Sort(sort string) *DataGaGetCall {
	c.opt_["sort"] = sort
	return c
}

// StartIndex sets the optional parameter "start-index": An index of the
// first entity to retrieve. Use this parameter as a pagination
// mechanism along with the max-results parameter.
func (c *DataGaGetCall) StartIndex(startIndex int64) *DataGaGetCall {
	c.opt_["start-index"] = startIndex
	return c
}

func (c *DataGaGetCall) Do() (*GaData, error) {
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
	urls := googleapi.ResolveRelative("https://www.googleapis.com/analytics/v3/", "data/ga")
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
	ret := new(GaData)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns Analytics data for a profile.",
	//   "httpMethod": "GET",
	//   "id": "analytics.data.ga.get",
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
	//       "description": "End date for fetching Analytics data. All requests should specify an end date formatted as YYYY-MM-DD.",
	//       "location": "query",
	//       "pattern": "[0-9]{4}-[0-9]{2}-[0-9]{2}",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "filters": {
	//       "description": "A comma-separated list of dimension or metric filters to be applied to Analytics data.",
	//       "location": "query",
	//       "pattern": "ga:.+",
	//       "type": "string"
	//     },
	//     "ids": {
	//       "description": "Unique table ID for retrieving Analytics data. Table ID is of the form ga:XXXX, where XXXX is the Analytics profile ID.",
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
	//       "description": "A comma-separated list of Analytics metrics. E.g., 'ga:visits,ga:pageviews'. At least one metric must be specified.",
	//       "location": "query",
	//       "pattern": "ga:.+",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "segment": {
	//       "description": "An Analytics advanced segment to be applied to data.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sort": {
	//       "description": "A comma-separated list of dimensions or metrics that determine the sort order for Analytics data.",
	//       "location": "query",
	//       "pattern": "(-)?ga:.+",
	//       "type": "string"
	//     },
	//     "start-date": {
	//       "description": "Start date for fetching Analytics data. All requests should specify a start date formatted as YYYY-MM-DD.",
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
	//   "path": "data/ga",
	//   "response": {
	//     "$ref": "GaData"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/analytics",
	//     "https://www.googleapis.com/auth/analytics.readonly"
	//   ]
	// }

}

// method id "analytics.data.mcf.get":

type DataMcfGetCall struct {
	s         *Service
	ids       string
	startDate string
	endDate   string
	metrics   string
	opt_      map[string]interface{}
}

// Get: Returns Analytics Multi-Channel Funnels data for a profile.
func (r *DataMcfService) Get(ids string, startDate string, endDate string, metrics string) *DataMcfGetCall {
	c := &DataMcfGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.ids = ids
	c.startDate = startDate
	c.endDate = endDate
	c.metrics = metrics
	return c
}

// Dimensions sets the optional parameter "dimensions": A
// comma-separated list of Multi-Channel Funnels dimensions. E.g.,
// 'mcf:source,mcf:medium'.
func (c *DataMcfGetCall) Dimensions(dimensions string) *DataMcfGetCall {
	c.opt_["dimensions"] = dimensions
	return c
}

// Filters sets the optional parameter "filters": A comma-separated list
// of dimension or metric filters to be applied to the Analytics data.
func (c *DataMcfGetCall) Filters(filters string) *DataMcfGetCall {
	c.opt_["filters"] = filters
	return c
}

// MaxResults sets the optional parameter "max-results": The maximum
// number of entries to include in this feed.
func (c *DataMcfGetCall) MaxResults(maxResults int64) *DataMcfGetCall {
	c.opt_["max-results"] = maxResults
	return c
}

// Sort sets the optional parameter "sort": A comma-separated list of
// dimensions or metrics that determine the sort order for the Analytics
// data.
func (c *DataMcfGetCall) Sort(sort string) *DataMcfGetCall {
	c.opt_["sort"] = sort
	return c
}

// StartIndex sets the optional parameter "start-index": An index of the
// first entity to retrieve. Use this parameter as a pagination
// mechanism along with the max-results parameter.
func (c *DataMcfGetCall) StartIndex(startIndex int64) *DataMcfGetCall {
	c.opt_["start-index"] = startIndex
	return c
}

func (c *DataMcfGetCall) Do() (*McfData, error) {
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
	urls := googleapi.ResolveRelative("https://www.googleapis.com/analytics/v3/", "data/mcf")
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
	ret := new(McfData)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns Analytics Multi-Channel Funnels data for a profile.",
	//   "httpMethod": "GET",
	//   "id": "analytics.data.mcf.get",
	//   "parameterOrder": [
	//     "ids",
	//     "start-date",
	//     "end-date",
	//     "metrics"
	//   ],
	//   "parameters": {
	//     "dimensions": {
	//       "description": "A comma-separated list of Multi-Channel Funnels dimensions. E.g., 'mcf:source,mcf:medium'.",
	//       "location": "query",
	//       "pattern": "(mcf:.+)?",
	//       "type": "string"
	//     },
	//     "end-date": {
	//       "description": "End date for fetching Analytics data. All requests should specify an end date formatted as YYYY-MM-DD.",
	//       "location": "query",
	//       "pattern": "[0-9]{4}-[0-9]{2}-[0-9]{2}",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "filters": {
	//       "description": "A comma-separated list of dimension or metric filters to be applied to the Analytics data.",
	//       "location": "query",
	//       "pattern": "mcf:.+",
	//       "type": "string"
	//     },
	//     "ids": {
	//       "description": "Unique table ID for retrieving Analytics data. Table ID is of the form ga:XXXX, where XXXX is the Analytics profile ID.",
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
	//       "description": "A comma-separated list of Multi-Channel Funnels metrics. E.g., 'mcf:totalConversions,mcf:totalConversionValue'. At least one metric must be specified.",
	//       "location": "query",
	//       "pattern": "mcf:.+",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sort": {
	//       "description": "A comma-separated list of dimensions or metrics that determine the sort order for the Analytics data.",
	//       "location": "query",
	//       "pattern": "(-)?mcf:.+",
	//       "type": "string"
	//     },
	//     "start-date": {
	//       "description": "Start date for fetching Analytics data. All requests should specify a start date formatted as YYYY-MM-DD.",
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
	//   "path": "data/mcf",
	//   "response": {
	//     "$ref": "McfData"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/analytics",
	//     "https://www.googleapis.com/auth/analytics.readonly"
	//   ]
	// }

}

// method id "analytics.management.accounts.list":

type ManagementAccountsListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Lists all accounts to which the user has access.
func (r *ManagementAccountsService) List() *ManagementAccountsListCall {
	c := &ManagementAccountsListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// MaxResults sets the optional parameter "max-results": The maximum
// number of accounts to include in this response.
func (c *ManagementAccountsListCall) MaxResults(maxResults int64) *ManagementAccountsListCall {
	c.opt_["max-results"] = maxResults
	return c
}

// StartIndex sets the optional parameter "start-index": An index of the
// first account to retrieve. Use this parameter as a pagination
// mechanism along with the max-results parameter.
func (c *ManagementAccountsListCall) StartIndex(startIndex int64) *ManagementAccountsListCall {
	c.opt_["start-index"] = startIndex
	return c
}

func (c *ManagementAccountsListCall) Do() (*Accounts, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["max-results"]; ok {
		params.Set("max-results", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["start-index"]; ok {
		params.Set("start-index", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/analytics/v3/", "management/accounts")
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
	ret := new(Accounts)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all accounts to which the user has access.",
	//   "httpMethod": "GET",
	//   "id": "analytics.management.accounts.list",
	//   "parameters": {
	//     "max-results": {
	//       "description": "The maximum number of accounts to include in this response.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "start-index": {
	//       "description": "An index of the first account to retrieve. Use this parameter as a pagination mechanism along with the max-results parameter.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "management/accounts",
	//   "response": {
	//     "$ref": "Accounts"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/analytics",
	//     "https://www.googleapis.com/auth/analytics.readonly"
	//   ]
	// }

}

// method id "analytics.management.customDataSources.list":

type ManagementCustomDataSourcesListCall struct {
	s             *Service
	accountId     string
	webPropertyId string
	opt_          map[string]interface{}
}

// List: List custom data sources to which the user has access.
func (r *ManagementCustomDataSourcesService) List(accountId string, webPropertyId string) *ManagementCustomDataSourcesListCall {
	c := &ManagementCustomDataSourcesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.webPropertyId = webPropertyId
	return c
}

// MaxResults sets the optional parameter "max-results": The maximum
// number of custom data sources to include in this response.
func (c *ManagementCustomDataSourcesListCall) MaxResults(maxResults int64) *ManagementCustomDataSourcesListCall {
	c.opt_["max-results"] = maxResults
	return c
}

// StartIndex sets the optional parameter "start-index": A 1-based index
// of the first custom data source to retrieve. Use this parameter as a
// pagination mechanism along with the max-results parameter.
func (c *ManagementCustomDataSourcesListCall) StartIndex(startIndex int64) *ManagementCustomDataSourcesListCall {
	c.opt_["start-index"] = startIndex
	return c
}

func (c *ManagementCustomDataSourcesListCall) Do() (*CustomDataSources, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["max-results"]; ok {
		params.Set("max-results", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["start-index"]; ok {
		params.Set("start-index", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/analytics/v3/", "management/accounts/{accountId}/webproperties/{webPropertyId}/customDataSources")
	urls = strings.Replace(urls, "{accountId}", cleanPathString(c.accountId), 1)
	urls = strings.Replace(urls, "{webPropertyId}", cleanPathString(c.webPropertyId), 1)
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
	ret := new(CustomDataSources)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List custom data sources to which the user has access.",
	//   "httpMethod": "GET",
	//   "id": "analytics.management.customDataSources.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "webPropertyId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account Id for the custom data sources to retrieve.",
	//       "location": "path",
	//       "pattern": "\\d+",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "max-results": {
	//       "description": "The maximum number of custom data sources to include in this response.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "start-index": {
	//       "description": "A 1-based index of the first custom data source to retrieve. Use this parameter as a pagination mechanism along with the max-results parameter.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "webPropertyId": {
	//       "description": "Web property Id for the custom data sources to retrieve.",
	//       "location": "path",
	//       "pattern": "UA-(\\d+)-(\\d+)",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "management/accounts/{accountId}/webproperties/{webPropertyId}/customDataSources",
	//   "response": {
	//     "$ref": "CustomDataSources"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/analytics",
	//     "https://www.googleapis.com/auth/analytics.readonly"
	//   ]
	// }

}

// method id "analytics.management.dailyUploads.delete":

type ManagementDailyUploadsDeleteCall struct {
	s                  *Service
	accountId          string
	webPropertyId      string
	customDataSourceId string
	date               string
	type_              string
	opt_               map[string]interface{}
}

// Delete: Delete uploaded data for the given date.
func (r *ManagementDailyUploadsService) Delete(accountId string, webPropertyId string, customDataSourceId string, date string, type_ string) *ManagementDailyUploadsDeleteCall {
	c := &ManagementDailyUploadsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.webPropertyId = webPropertyId
	c.customDataSourceId = customDataSourceId
	c.date = date
	c.type_ = type_
	return c
}

func (c *ManagementDailyUploadsDeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("type", fmt.Sprintf("%v", c.type_))
	urls := googleapi.ResolveRelative("https://www.googleapis.com/analytics/v3/", "management/accounts/{accountId}/webproperties/{webPropertyId}/customDataSources/{customDataSourceId}/dailyUploads/{date}")
	urls = strings.Replace(urls, "{accountId}", cleanPathString(c.accountId), 1)
	urls = strings.Replace(urls, "{webPropertyId}", cleanPathString(c.webPropertyId), 1)
	urls = strings.Replace(urls, "{customDataSourceId}", cleanPathString(c.customDataSourceId), 1)
	urls = strings.Replace(urls, "{date}", cleanPathString(c.date), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
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
	//   "description": "Delete uploaded data for the given date.",
	//   "httpMethod": "DELETE",
	//   "id": "analytics.management.dailyUploads.delete",
	//   "parameterOrder": [
	//     "accountId",
	//     "webPropertyId",
	//     "customDataSourceId",
	//     "date",
	//     "type"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account Id associated with daily upload delete.",
	//       "location": "path",
	//       "pattern": "[0-9]+",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "customDataSourceId": {
	//       "description": "Custom data source Id associated with daily upload delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "date": {
	//       "description": "Date for which data is to be deleted. Date should be formatted as YYYY-MM-DD.",
	//       "location": "path",
	//       "pattern": "[0-9]{4}-[0-9]{2}-[0-9]{2}",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "type": {
	//       "description": "Type of data for this delete.",
	//       "enum": [
	//         "cost"
	//       ],
	//       "enumDescriptions": [
	//         "Value for specifying cost data upload."
	//       ],
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "webPropertyId": {
	//       "description": "Web property Id associated with daily upload delete.",
	//       "location": "path",
	//       "pattern": "UA-[0-9]+-[0-9]+",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "management/accounts/{accountId}/webproperties/{webPropertyId}/customDataSources/{customDataSourceId}/dailyUploads/{date}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/analytics"
	//   ]
	// }

}

// method id "analytics.management.dailyUploads.list":

type ManagementDailyUploadsListCall struct {
	s                  *Service
	accountId          string
	webPropertyId      string
	customDataSourceId string
	startDate          string
	endDate            string
	opt_               map[string]interface{}
}

// List: List daily uploads to which the user has access.
func (r *ManagementDailyUploadsService) List(accountId string, webPropertyId string, customDataSourceId string, startDate string, endDate string) *ManagementDailyUploadsListCall {
	c := &ManagementDailyUploadsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.webPropertyId = webPropertyId
	c.customDataSourceId = customDataSourceId
	c.startDate = startDate
	c.endDate = endDate
	return c
}

// MaxResults sets the optional parameter "max-results": The maximum
// number of custom data sources to include in this response.
func (c *ManagementDailyUploadsListCall) MaxResults(maxResults int64) *ManagementDailyUploadsListCall {
	c.opt_["max-results"] = maxResults
	return c
}

// StartIndex sets the optional parameter "start-index": A 1-based index
// of the first daily upload to retrieve. Use this parameter as a
// pagination mechanism along with the max-results parameter.
func (c *ManagementDailyUploadsListCall) StartIndex(startIndex int64) *ManagementDailyUploadsListCall {
	c.opt_["start-index"] = startIndex
	return c
}

func (c *ManagementDailyUploadsListCall) Do() (*DailyUploads, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("end-date", fmt.Sprintf("%v", c.endDate))
	params.Set("start-date", fmt.Sprintf("%v", c.startDate))
	if v, ok := c.opt_["max-results"]; ok {
		params.Set("max-results", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["start-index"]; ok {
		params.Set("start-index", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/analytics/v3/", "management/accounts/{accountId}/webproperties/{webPropertyId}/customDataSources/{customDataSourceId}/dailyUploads")
	urls = strings.Replace(urls, "{accountId}", cleanPathString(c.accountId), 1)
	urls = strings.Replace(urls, "{webPropertyId}", cleanPathString(c.webPropertyId), 1)
	urls = strings.Replace(urls, "{customDataSourceId}", cleanPathString(c.customDataSourceId), 1)
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
	ret := new(DailyUploads)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List daily uploads to which the user has access.",
	//   "httpMethod": "GET",
	//   "id": "analytics.management.dailyUploads.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "webPropertyId",
	//     "customDataSourceId",
	//     "start-date",
	//     "end-date"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account Id for the daily uploads to retrieve.",
	//       "location": "path",
	//       "pattern": "\\d+",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "customDataSourceId": {
	//       "description": "Custom data source Id for daily uploads to retrieve.",
	//       "location": "path",
	//       "pattern": ".{22}",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "end-date": {
	//       "description": "End date of the form YYYY-MM-DD.",
	//       "location": "query",
	//       "pattern": "[0-9]{4}-[0-9]{2}-[0-9]{2}",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "max-results": {
	//       "description": "The maximum number of custom data sources to include in this response.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "start-date": {
	//       "description": "Start date of the form YYYY-MM-DD.",
	//       "location": "query",
	//       "pattern": "[0-9]{4}-[0-9]{2}-[0-9]{2}",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "start-index": {
	//       "description": "A 1-based index of the first daily upload to retrieve. Use this parameter as a pagination mechanism along with the max-results parameter.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "webPropertyId": {
	//       "description": "Web property Id for the daily uploads to retrieve.",
	//       "location": "path",
	//       "pattern": "UA-(\\d+)-(\\d+)",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "management/accounts/{accountId}/webproperties/{webPropertyId}/customDataSources/{customDataSourceId}/dailyUploads",
	//   "response": {
	//     "$ref": "DailyUploads"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/analytics",
	//     "https://www.googleapis.com/auth/analytics.readonly"
	//   ]
	// }

}

// method id "analytics.management.dailyUploads.upload":

type ManagementDailyUploadsUploadCall struct {
	s                  *Service
	accountId          string
	webPropertyId      string
	customDataSourceId string
	date               string
	appendNumber       int64
	type_              string
	opt_               map[string]interface{}
	media_             io.Reader
}

// Upload: Update/Overwrite data for a custom data source.
func (r *ManagementDailyUploadsService) Upload(accountId string, webPropertyId string, customDataSourceId string, date string, appendNumber int64, type_ string) *ManagementDailyUploadsUploadCall {
	c := &ManagementDailyUploadsUploadCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.webPropertyId = webPropertyId
	c.customDataSourceId = customDataSourceId
	c.date = date
	c.appendNumber = appendNumber
	c.type_ = type_
	return c
}

// Reset sets the optional parameter "reset": Reset/Overwrite all
// previous appends for this date and start over with this file as the
// first upload.
func (c *ManagementDailyUploadsUploadCall) Reset(reset bool) *ManagementDailyUploadsUploadCall {
	c.opt_["reset"] = reset
	return c
}
func (c *ManagementDailyUploadsUploadCall) Media(r io.Reader) *ManagementDailyUploadsUploadCall {
	c.media_ = r
	return c
}

func (c *ManagementDailyUploadsUploadCall) Do() (*DailyUploadAppend, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("appendNumber", fmt.Sprintf("%v", c.appendNumber))
	params.Set("type", fmt.Sprintf("%v", c.type_))
	if v, ok := c.opt_["reset"]; ok {
		params.Set("reset", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/analytics/v3/", "management/accounts/{accountId}/webproperties/{webPropertyId}/customDataSources/{customDataSourceId}/dailyUploads/{date}/uploads")
	if c.media_ != nil {
		urls = strings.Replace(urls, "https://www.googleapis.com/", "https://www.googleapis.com/upload/", 1)
		params.Set("uploadType", "multipart")
	}
	urls = strings.Replace(urls, "{accountId}", cleanPathString(c.accountId), 1)
	urls = strings.Replace(urls, "{webPropertyId}", cleanPathString(c.webPropertyId), 1)
	urls = strings.Replace(urls, "{customDataSourceId}", cleanPathString(c.customDataSourceId), 1)
	urls = strings.Replace(urls, "{date}", cleanPathString(c.date), 1)
	urls += "?" + params.Encode()
	body = new(bytes.Buffer)
	ctype := "application/json"
	contentLength_, hasMedia_ := googleapi.ConditionallyIncludeMedia(c.media_, &body, &ctype)
	req, _ := http.NewRequest("POST", urls, body)
	if hasMedia_ {
		req.ContentLength = contentLength_
	}
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(DailyUploadAppend)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update/Overwrite data for a custom data source.",
	//   "httpMethod": "POST",
	//   "id": "analytics.management.dailyUploads.upload",
	//   "mediaUpload": {
	//     "accept": [
	//       "application/octet-stream"
	//     ],
	//     "maxSize": "5MB",
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/analytics/v3/management/accounts/{accountId}/webproperties/{webPropertyId}/customDataSources/{customDataSourceId}/dailyUploads/{date}/uploads"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/analytics/v3/management/accounts/{accountId}/webproperties/{webPropertyId}/customDataSources/{customDataSourceId}/dailyUploads/{date}/uploads"
	//       }
	//     }
	//   },
	//   "parameterOrder": [
	//     "accountId",
	//     "webPropertyId",
	//     "customDataSourceId",
	//     "date",
	//     "appendNumber",
	//     "type"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account Id associated with daily upload.",
	//       "location": "path",
	//       "pattern": "\\d+",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "appendNumber": {
	//       "description": "Append number for this upload indexed from 1.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "20",
	//       "minimum": "1",
	//       "required": true,
	//       "type": "integer"
	//     },
	//     "customDataSourceId": {
	//       "description": "Custom data source Id to which the data being uploaded belongs.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "date": {
	//       "description": "Date for which data is uploaded. Date should be formatted as YYYY-MM-DD.",
	//       "location": "path",
	//       "pattern": "[0-9]{4}-[0-9]{2}-[0-9]{2}",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "reset": {
	//       "default": "false",
	//       "description": "Reset/Overwrite all previous appends for this date and start over with this file as the first upload.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "type": {
	//       "description": "Type of data for this upload.",
	//       "enum": [
	//         "cost"
	//       ],
	//       "enumDescriptions": [
	//         "Value for specifying cost data upload."
	//       ],
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "webPropertyId": {
	//       "description": "Web property Id associated with daily upload.",
	//       "location": "path",
	//       "pattern": "UA-\\d+-\\d+",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "management/accounts/{accountId}/webproperties/{webPropertyId}/customDataSources/{customDataSourceId}/dailyUploads/{date}/uploads",
	//   "response": {
	//     "$ref": "DailyUploadAppend"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/analytics"
	//   ],
	//   "supportsMediaUpload": true
	// }

}

// method id "analytics.management.experiments.delete":

type ManagementExperimentsDeleteCall struct {
	s             *Service
	accountId     string
	webPropertyId string
	profileId     string
	experimentId  string
	opt_          map[string]interface{}
}

// Delete: Delete an experiment.
func (r *ManagementExperimentsService) Delete(accountId string, webPropertyId string, profileId string, experimentId string) *ManagementExperimentsDeleteCall {
	c := &ManagementExperimentsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.webPropertyId = webPropertyId
	c.profileId = profileId
	c.experimentId = experimentId
	return c
}

func (c *ManagementExperimentsDeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/analytics/v3/", "management/accounts/{accountId}/webproperties/{webPropertyId}/profiles/{profileId}/experiments/{experimentId}")
	urls = strings.Replace(urls, "{accountId}", cleanPathString(c.accountId), 1)
	urls = strings.Replace(urls, "{webPropertyId}", cleanPathString(c.webPropertyId), 1)
	urls = strings.Replace(urls, "{profileId}", cleanPathString(c.profileId), 1)
	urls = strings.Replace(urls, "{experimentId}", cleanPathString(c.experimentId), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
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
	//   "description": "Delete an experiment.",
	//   "httpMethod": "DELETE",
	//   "id": "analytics.management.experiments.delete",
	//   "parameterOrder": [
	//     "accountId",
	//     "webPropertyId",
	//     "profileId",
	//     "experimentId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account ID to which the experiment belongs",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "experimentId": {
	//       "description": "ID of the experiment to delete",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "Profile ID to which the experiment belongs",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "webPropertyId": {
	//       "description": "Web property ID to which the experiment belongs",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "management/accounts/{accountId}/webproperties/{webPropertyId}/profiles/{profileId}/experiments/{experimentId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/analytics"
	//   ]
	// }

}

// method id "analytics.management.experiments.get":

type ManagementExperimentsGetCall struct {
	s             *Service
	accountId     string
	webPropertyId string
	profileId     string
	experimentId  string
	opt_          map[string]interface{}
}

// Get: Returns an experiment to which the user has access.
func (r *ManagementExperimentsService) Get(accountId string, webPropertyId string, profileId string, experimentId string) *ManagementExperimentsGetCall {
	c := &ManagementExperimentsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.webPropertyId = webPropertyId
	c.profileId = profileId
	c.experimentId = experimentId
	return c
}

func (c *ManagementExperimentsGetCall) Do() (*Experiment, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/analytics/v3/", "management/accounts/{accountId}/webproperties/{webPropertyId}/profiles/{profileId}/experiments/{experimentId}")
	urls = strings.Replace(urls, "{accountId}", cleanPathString(c.accountId), 1)
	urls = strings.Replace(urls, "{webPropertyId}", cleanPathString(c.webPropertyId), 1)
	urls = strings.Replace(urls, "{profileId}", cleanPathString(c.profileId), 1)
	urls = strings.Replace(urls, "{experimentId}", cleanPathString(c.experimentId), 1)
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
	ret := new(Experiment)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns an experiment to which the user has access.",
	//   "httpMethod": "GET",
	//   "id": "analytics.management.experiments.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "webPropertyId",
	//     "profileId",
	//     "experimentId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account ID to retrieve the experiment for.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "experimentId": {
	//       "description": "Experiment ID to retrieve the experiment for.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "Profile ID to retrieve the experiment for.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "webPropertyId": {
	//       "description": "Web property ID to retrieve the experiment for.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "management/accounts/{accountId}/webproperties/{webPropertyId}/profiles/{profileId}/experiments/{experimentId}",
	//   "response": {
	//     "$ref": "Experiment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/analytics",
	//     "https://www.googleapis.com/auth/analytics.readonly"
	//   ]
	// }

}

// method id "analytics.management.experiments.insert":

type ManagementExperimentsInsertCall struct {
	s             *Service
	accountId     string
	webPropertyId string
	profileId     string
	experiment    *Experiment
	opt_          map[string]interface{}
}

// Insert: Create a new experiment.
func (r *ManagementExperimentsService) Insert(accountId string, webPropertyId string, profileId string, experiment *Experiment) *ManagementExperimentsInsertCall {
	c := &ManagementExperimentsInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.webPropertyId = webPropertyId
	c.profileId = profileId
	c.experiment = experiment
	return c
}

func (c *ManagementExperimentsInsertCall) Do() (*Experiment, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.experiment)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/analytics/v3/", "management/accounts/{accountId}/webproperties/{webPropertyId}/profiles/{profileId}/experiments")
	urls = strings.Replace(urls, "{accountId}", cleanPathString(c.accountId), 1)
	urls = strings.Replace(urls, "{webPropertyId}", cleanPathString(c.webPropertyId), 1)
	urls = strings.Replace(urls, "{profileId}", cleanPathString(c.profileId), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Experiment)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Create a new experiment.",
	//   "httpMethod": "POST",
	//   "id": "analytics.management.experiments.insert",
	//   "parameterOrder": [
	//     "accountId",
	//     "webPropertyId",
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account ID to create the experiment for.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "Profile ID to create the experiment for.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "webPropertyId": {
	//       "description": "Web property ID to create the experiment for.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "management/accounts/{accountId}/webproperties/{webPropertyId}/profiles/{profileId}/experiments",
	//   "request": {
	//     "$ref": "Experiment"
	//   },
	//   "response": {
	//     "$ref": "Experiment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/analytics"
	//   ]
	// }

}

// method id "analytics.management.experiments.list":

type ManagementExperimentsListCall struct {
	s             *Service
	accountId     string
	webPropertyId string
	profileId     string
	opt_          map[string]interface{}
}

// List: Lists experiments to which the user has access.
func (r *ManagementExperimentsService) List(accountId string, webPropertyId string, profileId string) *ManagementExperimentsListCall {
	c := &ManagementExperimentsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.webPropertyId = webPropertyId
	c.profileId = profileId
	return c
}

// MaxResults sets the optional parameter "max-results": The maximum
// number of experiments to include in this response.
func (c *ManagementExperimentsListCall) MaxResults(maxResults int64) *ManagementExperimentsListCall {
	c.opt_["max-results"] = maxResults
	return c
}

// StartIndex sets the optional parameter "start-index": An index of the
// first experiment to retrieve. Use this parameter as a pagination
// mechanism along with the max-results parameter.
func (c *ManagementExperimentsListCall) StartIndex(startIndex int64) *ManagementExperimentsListCall {
	c.opt_["start-index"] = startIndex
	return c
}

func (c *ManagementExperimentsListCall) Do() (*Experiments, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["max-results"]; ok {
		params.Set("max-results", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["start-index"]; ok {
		params.Set("start-index", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/analytics/v3/", "management/accounts/{accountId}/webproperties/{webPropertyId}/profiles/{profileId}/experiments")
	urls = strings.Replace(urls, "{accountId}", cleanPathString(c.accountId), 1)
	urls = strings.Replace(urls, "{webPropertyId}", cleanPathString(c.webPropertyId), 1)
	urls = strings.Replace(urls, "{profileId}", cleanPathString(c.profileId), 1)
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
	ret := new(Experiments)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists experiments to which the user has access.",
	//   "httpMethod": "GET",
	//   "id": "analytics.management.experiments.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "webPropertyId",
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account ID to retrieve experiments for.",
	//       "location": "path",
	//       "pattern": "\\d+",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "max-results": {
	//       "description": "The maximum number of experiments to include in this response.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "profileId": {
	//       "description": "Profile ID to retrieve experiments for.",
	//       "location": "path",
	//       "pattern": "\\d+",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "start-index": {
	//       "description": "An index of the first experiment to retrieve. Use this parameter as a pagination mechanism along with the max-results parameter.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "webPropertyId": {
	//       "description": "Web property ID to retrieve experiments for.",
	//       "location": "path",
	//       "pattern": "UA-(\\d+)-(\\d+)",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "management/accounts/{accountId}/webproperties/{webPropertyId}/profiles/{profileId}/experiments",
	//   "response": {
	//     "$ref": "Experiments"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/analytics",
	//     "https://www.googleapis.com/auth/analytics.readonly"
	//   ]
	// }

}

// method id "analytics.management.experiments.patch":

type ManagementExperimentsPatchCall struct {
	s             *Service
	accountId     string
	webPropertyId string
	profileId     string
	experimentId  string
	experiment    *Experiment
	opt_          map[string]interface{}
}

// Patch: Update an existing experiment. This method supports patch
// semantics.
func (r *ManagementExperimentsService) Patch(accountId string, webPropertyId string, profileId string, experimentId string, experiment *Experiment) *ManagementExperimentsPatchCall {
	c := &ManagementExperimentsPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.webPropertyId = webPropertyId
	c.profileId = profileId
	c.experimentId = experimentId
	c.experiment = experiment
	return c
}

func (c *ManagementExperimentsPatchCall) Do() (*Experiment, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.experiment)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/analytics/v3/", "management/accounts/{accountId}/webproperties/{webPropertyId}/profiles/{profileId}/experiments/{experimentId}")
	urls = strings.Replace(urls, "{accountId}", cleanPathString(c.accountId), 1)
	urls = strings.Replace(urls, "{webPropertyId}", cleanPathString(c.webPropertyId), 1)
	urls = strings.Replace(urls, "{profileId}", cleanPathString(c.profileId), 1)
	urls = strings.Replace(urls, "{experimentId}", cleanPathString(c.experimentId), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Experiment)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update an existing experiment. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "analytics.management.experiments.patch",
	//   "parameterOrder": [
	//     "accountId",
	//     "webPropertyId",
	//     "profileId",
	//     "experimentId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account ID of the experiment to update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "experimentId": {
	//       "description": "Experiment ID of the experiment to update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "Profile ID of the experiment to update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "webPropertyId": {
	//       "description": "Web property ID of the experiment to update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "management/accounts/{accountId}/webproperties/{webPropertyId}/profiles/{profileId}/experiments/{experimentId}",
	//   "request": {
	//     "$ref": "Experiment"
	//   },
	//   "response": {
	//     "$ref": "Experiment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/analytics"
	//   ]
	// }

}

// method id "analytics.management.experiments.update":

type ManagementExperimentsUpdateCall struct {
	s             *Service
	accountId     string
	webPropertyId string
	profileId     string
	experimentId  string
	experiment    *Experiment
	opt_          map[string]interface{}
}

// Update: Update an existing experiment.
func (r *ManagementExperimentsService) Update(accountId string, webPropertyId string, profileId string, experimentId string, experiment *Experiment) *ManagementExperimentsUpdateCall {
	c := &ManagementExperimentsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.webPropertyId = webPropertyId
	c.profileId = profileId
	c.experimentId = experimentId
	c.experiment = experiment
	return c
}

func (c *ManagementExperimentsUpdateCall) Do() (*Experiment, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.experiment)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/analytics/v3/", "management/accounts/{accountId}/webproperties/{webPropertyId}/profiles/{profileId}/experiments/{experimentId}")
	urls = strings.Replace(urls, "{accountId}", cleanPathString(c.accountId), 1)
	urls = strings.Replace(urls, "{webPropertyId}", cleanPathString(c.webPropertyId), 1)
	urls = strings.Replace(urls, "{profileId}", cleanPathString(c.profileId), 1)
	urls = strings.Replace(urls, "{experimentId}", cleanPathString(c.experimentId), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Experiment)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update an existing experiment.",
	//   "httpMethod": "PUT",
	//   "id": "analytics.management.experiments.update",
	//   "parameterOrder": [
	//     "accountId",
	//     "webPropertyId",
	//     "profileId",
	//     "experimentId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account ID of the experiment to update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "experimentId": {
	//       "description": "Experiment ID of the experiment to update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "Profile ID of the experiment to update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "webPropertyId": {
	//       "description": "Web property ID of the experiment to update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "management/accounts/{accountId}/webproperties/{webPropertyId}/profiles/{profileId}/experiments/{experimentId}",
	//   "request": {
	//     "$ref": "Experiment"
	//   },
	//   "response": {
	//     "$ref": "Experiment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/analytics"
	//   ]
	// }

}

// method id "analytics.management.goals.list":

type ManagementGoalsListCall struct {
	s             *Service
	accountId     string
	webPropertyId string
	profileId     string
	opt_          map[string]interface{}
}

// List: Lists goals to which the user has access.
func (r *ManagementGoalsService) List(accountId string, webPropertyId string, profileId string) *ManagementGoalsListCall {
	c := &ManagementGoalsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.webPropertyId = webPropertyId
	c.profileId = profileId
	return c
}

// MaxResults sets the optional parameter "max-results": The maximum
// number of goals to include in this response.
func (c *ManagementGoalsListCall) MaxResults(maxResults int64) *ManagementGoalsListCall {
	c.opt_["max-results"] = maxResults
	return c
}

// StartIndex sets the optional parameter "start-index": An index of the
// first goal to retrieve. Use this parameter as a pagination mechanism
// along with the max-results parameter.
func (c *ManagementGoalsListCall) StartIndex(startIndex int64) *ManagementGoalsListCall {
	c.opt_["start-index"] = startIndex
	return c
}

func (c *ManagementGoalsListCall) Do() (*Goals, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["max-results"]; ok {
		params.Set("max-results", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["start-index"]; ok {
		params.Set("start-index", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/analytics/v3/", "management/accounts/{accountId}/webproperties/{webPropertyId}/profiles/{profileId}/goals")
	urls = strings.Replace(urls, "{accountId}", cleanPathString(c.accountId), 1)
	urls = strings.Replace(urls, "{webPropertyId}", cleanPathString(c.webPropertyId), 1)
	urls = strings.Replace(urls, "{profileId}", cleanPathString(c.profileId), 1)
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
	ret := new(Goals)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists goals to which the user has access.",
	//   "httpMethod": "GET",
	//   "id": "analytics.management.goals.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "webPropertyId",
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account ID to retrieve goals for. Can either be a specific account ID or '~all', which refers to all the accounts that user has access to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "max-results": {
	//       "description": "The maximum number of goals to include in this response.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "profileId": {
	//       "description": "Profile ID to retrieve goals for. Can either be a specific profile ID or '~all', which refers to all the profiles that user has access to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "start-index": {
	//       "description": "An index of the first goal to retrieve. Use this parameter as a pagination mechanism along with the max-results parameter.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "webPropertyId": {
	//       "description": "Web property ID to retrieve goals for. Can either be a specific web property ID or '~all', which refers to all the web properties that user has access to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "management/accounts/{accountId}/webproperties/{webPropertyId}/profiles/{profileId}/goals",
	//   "response": {
	//     "$ref": "Goals"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/analytics",
	//     "https://www.googleapis.com/auth/analytics.readonly"
	//   ]
	// }

}

// method id "analytics.management.profiles.list":

type ManagementProfilesListCall struct {
	s             *Service
	accountId     string
	webPropertyId string
	opt_          map[string]interface{}
}

// List: Lists profiles to which the user has access.
func (r *ManagementProfilesService) List(accountId string, webPropertyId string) *ManagementProfilesListCall {
	c := &ManagementProfilesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.webPropertyId = webPropertyId
	return c
}

// MaxResults sets the optional parameter "max-results": The maximum
// number of profiles to include in this response.
func (c *ManagementProfilesListCall) MaxResults(maxResults int64) *ManagementProfilesListCall {
	c.opt_["max-results"] = maxResults
	return c
}

// StartIndex sets the optional parameter "start-index": An index of the
// first entity to retrieve. Use this parameter as a pagination
// mechanism along with the max-results parameter.
func (c *ManagementProfilesListCall) StartIndex(startIndex int64) *ManagementProfilesListCall {
	c.opt_["start-index"] = startIndex
	return c
}

func (c *ManagementProfilesListCall) Do() (*Profiles, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["max-results"]; ok {
		params.Set("max-results", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["start-index"]; ok {
		params.Set("start-index", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/analytics/v3/", "management/accounts/{accountId}/webproperties/{webPropertyId}/profiles")
	urls = strings.Replace(urls, "{accountId}", cleanPathString(c.accountId), 1)
	urls = strings.Replace(urls, "{webPropertyId}", cleanPathString(c.webPropertyId), 1)
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
	ret := new(Profiles)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists profiles to which the user has access.",
	//   "httpMethod": "GET",
	//   "id": "analytics.management.profiles.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "webPropertyId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account ID for the profiles to retrieve. Can either be a specific account ID or '~all', which refers to all the accounts to which the user has access.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "max-results": {
	//       "description": "The maximum number of profiles to include in this response.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "start-index": {
	//       "description": "An index of the first entity to retrieve. Use this parameter as a pagination mechanism along with the max-results parameter.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "webPropertyId": {
	//       "description": "Web property ID for the profiles to retrieve. Can either be a specific web property ID or '~all', which refers to all the web properties to which the user has access.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "management/accounts/{accountId}/webproperties/{webPropertyId}/profiles",
	//   "response": {
	//     "$ref": "Profiles"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/analytics",
	//     "https://www.googleapis.com/auth/analytics.readonly"
	//   ]
	// }

}

// method id "analytics.management.segments.list":

type ManagementSegmentsListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Lists advanced segments to which the user has access.
func (r *ManagementSegmentsService) List() *ManagementSegmentsListCall {
	c := &ManagementSegmentsListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// MaxResults sets the optional parameter "max-results": The maximum
// number of advanced segments to include in this response.
func (c *ManagementSegmentsListCall) MaxResults(maxResults int64) *ManagementSegmentsListCall {
	c.opt_["max-results"] = maxResults
	return c
}

// StartIndex sets the optional parameter "start-index": An index of the
// first advanced segment to retrieve. Use this parameter as a
// pagination mechanism along with the max-results parameter.
func (c *ManagementSegmentsListCall) StartIndex(startIndex int64) *ManagementSegmentsListCall {
	c.opt_["start-index"] = startIndex
	return c
}

func (c *ManagementSegmentsListCall) Do() (*Segments, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["max-results"]; ok {
		params.Set("max-results", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["start-index"]; ok {
		params.Set("start-index", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/analytics/v3/", "management/segments")
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
	ret := new(Segments)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists advanced segments to which the user has access.",
	//   "httpMethod": "GET",
	//   "id": "analytics.management.segments.list",
	//   "parameters": {
	//     "max-results": {
	//       "description": "The maximum number of advanced segments to include in this response.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "start-index": {
	//       "description": "An index of the first advanced segment to retrieve. Use this parameter as a pagination mechanism along with the max-results parameter.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "management/segments",
	//   "response": {
	//     "$ref": "Segments"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/analytics",
	//     "https://www.googleapis.com/auth/analytics.readonly"
	//   ]
	// }

}

// method id "analytics.management.webproperties.list":

type ManagementWebpropertiesListCall struct {
	s         *Service
	accountId string
	opt_      map[string]interface{}
}

// List: Lists web properties to which the user has access.
func (r *ManagementWebpropertiesService) List(accountId string) *ManagementWebpropertiesListCall {
	c := &ManagementWebpropertiesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	return c
}

// MaxResults sets the optional parameter "max-results": The maximum
// number of web properties to include in this response.
func (c *ManagementWebpropertiesListCall) MaxResults(maxResults int64) *ManagementWebpropertiesListCall {
	c.opt_["max-results"] = maxResults
	return c
}

// StartIndex sets the optional parameter "start-index": An index of the
// first entity to retrieve. Use this parameter as a pagination
// mechanism along with the max-results parameter.
func (c *ManagementWebpropertiesListCall) StartIndex(startIndex int64) *ManagementWebpropertiesListCall {
	c.opt_["start-index"] = startIndex
	return c
}

func (c *ManagementWebpropertiesListCall) Do() (*Webproperties, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["max-results"]; ok {
		params.Set("max-results", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["start-index"]; ok {
		params.Set("start-index", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/analytics/v3/", "management/accounts/{accountId}/webproperties")
	urls = strings.Replace(urls, "{accountId}", cleanPathString(c.accountId), 1)
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
	ret := new(Webproperties)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists web properties to which the user has access.",
	//   "httpMethod": "GET",
	//   "id": "analytics.management.webproperties.list",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account ID to retrieve web properties for. Can either be a specific account ID or '~all', which refers to all the accounts that user has access to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "max-results": {
	//       "description": "The maximum number of web properties to include in this response.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "start-index": {
	//       "description": "An index of the first entity to retrieve. Use this parameter as a pagination mechanism along with the max-results parameter.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "management/accounts/{accountId}/webproperties",
	//   "response": {
	//     "$ref": "Webproperties"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/analytics",
	//     "https://www.googleapis.com/auth/analytics.readonly"
	//   ]
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
