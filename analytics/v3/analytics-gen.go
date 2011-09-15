// Package analytics provides access to the Google Analytics API.
//
// See http://code.google.com/apis/analytics
//
// Usage example:
//
//   import "google-api-go-client.googlecode.com/hg/analytics/v3"
//   ...
//   analyticsService, err := analytics.New(oauthHttpClient)
package analytics

import (
	"bytes"
	"fmt"
	"http"
	"io"
	"json"
	"os"
	"strings"
	"strconv"
	"url"
	"google-api-go-client.googlecode.com/hg/google-api"
)

var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version

const apiId = "analytics:v3"
const apiName = "analytics"
const apiVersion = "v3"
const basePath = "https://www.googleapis.com/analytics/v3/management/"

// OAuth2 scopes used by this API.
const (
	// View your Google Analytics data
	AnalyticsReadonlyScope = "https://www.googleapis.com/auth/analytics.readonly"
)

func New(client *http.Client) (*Service, os.Error) {
	if client == nil {
		return nil, os.NewError("client is nil")
	}
	s := &Service{client: client}
	s.Management = &ManagementService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	Management *ManagementService
}

type ManagementService struct {
	s *Service
}

type Profiles struct {
	// StartIndex: The starting index of the resources, which is 1 by
	// default or otherwise specified by the start-index query parameter.
	StartIndex int64 `json:"startIndex,omitempty"`

	// NextLink: Link to next page for this profile collection.
	NextLink string `json:"nextLink,omitempty"`

	// TotalResults: The total number of results for the query, regardless
	// of the number of results in the response.
	TotalResults int64 `json:"totalResults,omitempty"`

	// Items: A list of profiles.
	Items []*Profile `json:"items,omitempty"`

	// Username: Email ID of the authenticated user
	Username string `json:"username,omitempty"`

	// Kind: Resource type for a profile.
	Kind string `json:"kind,omitempty"`

	// ItemsPerPage: The maximum number of resources the response can
	// contain, regardless of the actual number of resources returned. Its
	// value ranges from 1 to 10,000 with a value of 1000 by default, or
	// otherwise specified by the max-results query parameter.
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`

	// PreviousLink: Link to previous page for this profile collection.
	PreviousLink string `json:"previousLink,omitempty"`
}

type Profile struct {
	// DefaultPage: Default page for this profile.
	DefaultPage string `json:"defaultPage,omitempty"`

	// SelfLink: Link for this profile.
	SelfLink string `json:"selfLink,omitempty"`

	// Currency: The currency type associated with this profile.
	Currency string `json:"currency,omitempty"`

	// ParentLink: Parent link for this profile. Points to the web property
	// to which this profile belongs.
	ParentLink *ProfileParentLink `json:"parentLink,omitempty"`

	// Name: Name of this profile.
	Name string `json:"name,omitempty"`

	// AccountId: Account ID to which this profile belongs.
	AccountId string `json:"accountId,omitempty"`

	// SiteSearchQueryParameters: The site search query parameters for this
	// profile.
	SiteSearchQueryParameters string `json:"siteSearchQueryParameters,omitempty"`

	// Kind: Resource type for Analytics profile.
	Kind string `json:"kind,omitempty"`

	// Updated: Time this profile was last modified.
	Updated string `json:"updated,omitempty"`

	// WebPropertyId: Web property ID of the form UA-XXXXX-YY to which this
	// profile belongs.
	WebPropertyId string `json:"webPropertyId,omitempty"`

	// SiteSearchCategoryParameters: Site search category parameters for
	// this profile.
	SiteSearchCategoryParameters string `json:"siteSearchCategoryParameters,omitempty"`

	// Id: Profile ID.
	Id string `json:"id,omitempty"`

	// Created: Time this profile was created.
	Created string `json:"created,omitempty"`

	// InternalWebPropertyId: Internal ID for the web property to which this
	// profile belongs.
	InternalWebPropertyId string `json:"internalWebPropertyId,omitempty"`

	// ChildLink: Child link for this profile. Points to the list of goals
	// for this profile.
	ChildLink *ProfileChildLink `json:"childLink,omitempty"`

	// ExcludeQueryParameters: The query parameters that are excluded from
	// this profile.
	ExcludeQueryParameters string `json:"excludeQueryParameters,omitempty"`

	// Timezone: Time zone for which this profile has been configured.
	Timezone string `json:"timezone,omitempty"`
}

type ProfileParentLink struct {
	Href string `json:"href,omitempty"`

	Type string `json:"type,omitempty"`
}

type GoalVisitNumPagesDetails struct {
	// ComparisonValue: Value used for this comparison.
	ComparisonValue int64 `json:"comparisonValue,omitempty,string"`

	// ComparisonType: Type of comparison. Possible values are LESS_THAN,
	// GREATER_THAN, or EQUAL.
	ComparisonType string `json:"comparisonType,omitempty"`
}

type ProfileChildLink struct {
	Href string `json:"href,omitempty"`

	Type string `json:"type,omitempty"`
}

type Webproperty struct {
	// SelfLink: Link for this web property.
	SelfLink string `json:"selfLink,omitempty"`

	// ParentLink: Parent link for this web property. Points to the account
	// to which this web property belongs.
	ParentLink *WebpropertyParentLink `json:"parentLink,omitempty"`

	// Name: Name of this web property.
	Name string `json:"name,omitempty"`

	// AccountId: Account ID to which this web property belongs.
	AccountId string `json:"accountId,omitempty"`

	// Kind: Resource type for Analytics WebProperty.
	Kind string `json:"kind,omitempty"`

	// Updated: Time this web property was last modified.
	Updated string `json:"updated,omitempty"`

	// WebsiteUrl: Website url for this web property.
	WebsiteUrl string `json:"websiteUrl,omitempty"`

	// Id: Web property ID of the form UA-XXXXX-YY.
	Id string `json:"id,omitempty"`

	// Created: Time this web property was created.
	Created string `json:"created,omitempty"`

	// InternalWebPropertyId: Internal ID for this web property.
	InternalWebPropertyId string `json:"internalWebPropertyId,omitempty"`

	// ChildLink: Child link for this web property. Points to the list of
	// profiles for this web property.
	ChildLink *WebpropertyChildLink `json:"childLink,omitempty"`
}

type GoalUrlDestinationDetails struct {
	// CaseSensitive: Determines if the goal URL must exactly match the
	// capitalization of visited URLs.
	CaseSensitive bool `json:"caseSensitive,omitempty"`

	// FirstStepRequired: Determines if the first step in this goal is
	// required.
	FirstStepRequired bool `json:"firstStepRequired,omitempty"`

	// Steps: List of steps configured for this goal funnel.
	Steps []*GoalUrlDestinationDetailsSteps `json:"steps,omitempty"`

	// Url: URL for this goal.
	Url string `json:"url,omitempty"`

	// MatchType: Match type for the goal URL. Possible values are HEAD,
	// EXACT, or REGEX.
	MatchType string `json:"matchType,omitempty"`
}

type Accounts struct {
	// StartIndex: The starting index of the entries, which is 1 by default
	// or otherwise specified by the start-index query parameter.
	StartIndex int64 `json:"startIndex,omitempty"`

	// NextLink: Next link for this account collection.
	NextLink string `json:"nextLink,omitempty"`

	// TotalResults: The total number of results for the query, regardless
	// of the number of results in the response.
	TotalResults int64 `json:"totalResults,omitempty"`

	// Items: A list of accounts.
	Items []*Account `json:"items,omitempty"`

	// Username: Email ID of the authenticated user
	Username string `json:"username,omitempty"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`

	// ItemsPerPage: The maximum number of entries the response can contain,
	// regardless of the actual number of entries returned. Its value ranges
	// from 1 to 10,000 with a value of 1000 by default, or otherwise
	// specified by the max-results query parameter.
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`

	// PreviousLink: Previous link for this account collection.
	PreviousLink string `json:"previousLink,omitempty"`
}

type GoalEventDetails struct {
	EventConditions []*GoalEventDetailsEventConditions `json:"eventConditions,omitempty"`

	// UseEventValue: Determines if the event value should be used as the
	// value for this goal.
	UseEventValue bool `json:"useEventValue,omitempty"`
}

type GoalUrlDestinationDetailsSteps struct {
	// Name: Step name.
	Name string `json:"name,omitempty"`

	// Url: URL for this step.
	Url string `json:"url,omitempty"`

	// Number: Step number.
	Number int64 `json:"number,omitempty"`
}

type Segments struct {
	// StartIndex: The starting index of the resources, which is 1 by
	// default or otherwise specified by the start-index query parameter.
	StartIndex int64 `json:"startIndex,omitempty"`

	// NextLink: Link to next page for this advanced segment collection.
	NextLink string `json:"nextLink,omitempty"`

	// TotalResults: The total number of results for the query, regardless
	// of the number of results in the response.
	TotalResults int64 `json:"totalResults,omitempty"`

	// Items: A list of advanced segments.
	Items []*Segment `json:"items,omitempty"`

	// Username: Email ID of the authenticated user
	Username string `json:"username,omitempty"`

	// Kind: Resource type for advanced segments.
	Kind string `json:"kind,omitempty"`

	// ItemsPerPage: The maximum number of resources the response can
	// contain, regardless of the actual number of resources returned. Its
	// value ranges from 1 to 10,000 with a value of 1000 by default, or
	// otherwise specified by the max-results query parameter.
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`

	// PreviousLink: Link to previous page for this advanced segment
	// collection.
	PreviousLink string `json:"previousLink,omitempty"`
}

type Webproperties struct {
	// StartIndex: The starting index of the resources, which is 1 by
	// default or otherwise specified by the start-index query parameter.
	StartIndex int64 `json:"startIndex,omitempty"`

	// NextLink: Link to next page for this web property collection.
	NextLink string `json:"nextLink,omitempty"`

	// TotalResults: The total number of results for the query, regardless
	// of the number of results in the response.
	TotalResults int64 `json:"totalResults,omitempty"`

	// Items: A list of web properties.
	Items []*Webproperty `json:"items,omitempty"`

	// Username: Email ID of the authenticated user
	Username string `json:"username,omitempty"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`

	// ItemsPerPage: The maximum number of resources the response can
	// contain, regardless of the actual number of resources returned. Its
	// value ranges from 1 to 10,000 with a value of 1000 by default, or
	// otherwise specified by the max-results query parameter.
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`

	// PreviousLink: Link to previous page for this web property collection.
	PreviousLink string `json:"previousLink,omitempty"`
}

type Goal struct {
	// Value: Goal value.
	Value float64 `json:"value,omitempty"`

	// SelfLink: Link for this goal.
	SelfLink string `json:"selfLink,omitempty"`

	// ParentLink: Parent link for a goal. Points to the profile to which
	// this goal belongs.
	ParentLink *GoalParentLink `json:"parentLink,omitempty"`

	// Name: Goal name.
	Name string `json:"name,omitempty"`

	// AccountId: Account ID to which this goal belongs.
	AccountId string `json:"accountId,omitempty"`

	// Kind: Resource type for an Analytics goal.
	Kind string `json:"kind,omitempty"`

	// Updated: Time this goal was last modified.
	Updated string `json:"updated,omitempty"`

	// WebPropertyId: Web property ID to which this goal belongs. The web
	// property ID is of the form UA-XXXXX-YY.
	WebPropertyId string `json:"webPropertyId,omitempty"`

	// UrlDestinationDetails: Details for the goal of the type
	// URL_DESTINATION.
	UrlDestinationDetails *GoalUrlDestinationDetails `json:"urlDestinationDetails,omitempty"`

	// EventDetails: Details for the goal of the type EVENT.
	EventDetails *GoalEventDetails `json:"eventDetails,omitempty"`

	// VisitNumPagesDetails: Details for the goal of the type
	// VISIT_NUM_PAGES.
	VisitNumPagesDetails *GoalVisitNumPagesDetails `json:"visitNumPagesDetails,omitempty"`

	// Active: Determines whether this goal is active.
	Active bool `json:"active,omitempty"`

	// Id: Goal ID.
	Id string `json:"id,omitempty"`

	// Created: Time this goal was created.
	Created string `json:"created,omitempty"`

	// InternalWebPropertyId: Internal ID for the web property to which this
	// goal belongs.
	InternalWebPropertyId string `json:"internalWebPropertyId,omitempty"`

	// VisitTimeOnSiteDetails: Details for the goal of the type
	// VISIT_TIME_ON_SITE.
	VisitTimeOnSiteDetails *GoalVisitTimeOnSiteDetails `json:"visitTimeOnSiteDetails,omitempty"`

	// ProfileId: Profile ID to which this goal belongs.
	ProfileId string `json:"profileId,omitempty"`

	// Type: Goal type. Possible values are URL_DESTINATION,
	// VISIT_TIME_ON_SITE, VISIT_NUM_PAGES, AND EVENT.
	Type string `json:"type,omitempty"`
}

type GoalVisitTimeOnSiteDetails struct {
	// ComparisonValue: Value used for this comparison.
	ComparisonValue int64 `json:"comparisonValue,omitempty,string"`

	// ComparisonType: Type of comparison. Possible values are LESS_THAN or
	// GREATER_THAN.
	ComparisonType string `json:"comparisonType,omitempty"`
}

type GoalEventDetailsEventConditions struct {
	// Expression: Expression used for this match.
	Expression string `json:"expression,omitempty"`

	// Type: Type of this event condition. Possible values are CATEGORY,
	// ACTION, LABEL, or VALUE.
	Type string `json:"type,omitempty"`

	// ComparisonValue: Value used for this comparison.
	ComparisonValue int64 `json:"comparisonValue,omitempty,string"`

	// MatchType: Type of the match to be performed. Possible values are
	// REGEXP, BEGINS_WITH, or EXACT.
	MatchType string `json:"matchType,omitempty"`

	// ComparisonType: Type of comparison. Possible values are LESS_THAN,
	// GREATER_THAN or EQUAL.
	ComparisonType string `json:"comparisonType,omitempty"`
}

type WebpropertyParentLink struct {
	// Href: Link to the account for this web property.
	Href string `json:"href,omitempty"`

	// Type: Type of the parent link. Its value is "analytics#account".
	Type string `json:"type,omitempty"`
}

type GoalParentLink struct {
	Href string `json:"href,omitempty"`

	Type string `json:"type,omitempty"`
}

type AccountChildLink struct {
	// Href: Link to the list of web properties for this account.
	Href string `json:"href,omitempty"`

	// Type: Type of the child link. Its value is "analytics#webproperties".
	Type string `json:"type,omitempty"`
}

type Account struct {
	// SelfLink: Link for this account.
	SelfLink string `json:"selfLink,omitempty"`

	// Name: Account name.
	Name string `json:"name,omitempty"`

	// Kind: Resource type for Analytics account.
	Kind string `json:"kind,omitempty"`

	// Updated: Time the account was last modified.
	Updated string `json:"updated,omitempty"`

	// Id: Account ID.
	Id string `json:"id,omitempty"`

	// Created: Time the account was created.
	Created string `json:"created,omitempty"`

	// ChildLink: Child link for an account entry. Points to the list of web
	// properties for this account.
	ChildLink *AccountChildLink `json:"childLink,omitempty"`
}

type Segment struct {
	// SelfLink: Link for this advanced segment.
	SelfLink string `json:"selfLink,omitempty"`

	// Name: Advanced segment name.
	Name string `json:"name,omitempty"`

	// Kind: Resource type for Analytics advanced segment.
	Kind string `json:"kind,omitempty"`

	// Updated: Time the advanced segment was last modified.
	Updated string `json:"updated,omitempty"`

	// Id: Advanced segment ID.
	Id string `json:"id,omitempty"`

	// Created: Time the advanced segment was created.
	Created string `json:"created,omitempty"`

	// Definition: Advanced segment definition.
	Definition string `json:"definition,omitempty"`
}

type WebpropertyChildLink struct {
	// Href: Link to the list of profiles for this web property.
	Href string `json:"href,omitempty"`

	// Type: Type of the parent link. Its value is "analytics#profiles".
	Type string `json:"type,omitempty"`
}

type Goals struct {
	// StartIndex: The starting index of the resources, which is 1 by
	// default or otherwise specified by the start-index query parameter.
	StartIndex int64 `json:"startIndex,omitempty"`

	// NextLink: Link to next page for this goal collection.
	NextLink string `json:"nextLink,omitempty"`

	// TotalResults: The total number of results for the query, regardless
	// of the number of resources in the result.
	TotalResults int64 `json:"totalResults,omitempty"`

	// Items: A list of goals.
	Items []*Goal `json:"items,omitempty"`

	// Username: Email ID of the authenticated user
	Username string `json:"username,omitempty"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`

	// ItemsPerPage: The maximum number of resources the response can
	// contain, regardless of the actual number of resources returned. Its
	// value ranges from 1 to 10,000 with a value of 1000 by default, or
	// otherwise specified by the max-results query parameter.
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`

	// PreviousLink: Link to previous page for this goal collection.
	PreviousLink string `json:"previousLink,omitempty"`
}

func cleanPathString(s string) string {
	return strings.Map(func(r int) int {
		if r >= 0x30 && r <= 0x7a {
			return r
		}
		return -1
	}, s)
}
