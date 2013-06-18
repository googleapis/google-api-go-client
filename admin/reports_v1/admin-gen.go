// Package admin provides access to the Admin Reports API.
//
// See https://developers.google.com/admin-sdk/reports/
//
// Usage example:
//
//   import "code.google.com/p/google-api-go-client/admin/reports_v1"
//   ...
//   adminService, err := admin.New(oauthHttpClient)
package admin

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

const apiId = "admin:reports_v1"
const apiName = "admin"
const apiVersion = "reports_v1"
const basePath = "https://www.googleapis.com/admin/reports/v1/"

// OAuth2 scopes used by this API.
const (
	// View audit reports of Google Apps for your domain
	AdminReportsAuditReadonlyScope = "https://www.googleapis.com/auth/admin.reports.audit.readonly"

	// View usage reports of Google Apps for your domain
	AdminReportsUsageReadonlyScope = "https://www.googleapis.com/auth/admin.reports.usage.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Activities = NewActivitiesService(s)
	s.CustomerUsageReports = NewCustomerUsageReportsService(s)
	s.UserUsageReport = NewUserUsageReportService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Activities *ActivitiesService

	CustomerUsageReports *CustomerUsageReportsService

	UserUsageReport *UserUsageReportService
}

func NewActivitiesService(s *Service) *ActivitiesService {
	rs := &ActivitiesService{s: s}
	return rs
}

type ActivitiesService struct {
	s *Service
}

func NewCustomerUsageReportsService(s *Service) *CustomerUsageReportsService {
	rs := &CustomerUsageReportsService{s: s}
	return rs
}

type CustomerUsageReportsService struct {
	s *Service
}

func NewUserUsageReportService(s *Service) *UserUsageReportService {
	rs := &UserUsageReportService{s: s}
	return rs
}

type UserUsageReportService struct {
	s *Service
}

type Activities struct {
	// Items: Each record in read response.
	Items []*Activity `json:"items,omitempty"`

	// Kind: Kind of list response this is.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token for retrieving the next page
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type Activity struct {
	// Actor: User doing the action.
	Actor *ActivityActor `json:"actor,omitempty"`

	// Events: Activity events.
	Events []*ActivityEvents `json:"events,omitempty"`

	// Id: Unique identifier for each activity record.
	Id *ActivityId `json:"id,omitempty"`

	// IpAddress: IP Address of the user doing the action.
	IpAddress string `json:"ipAddress,omitempty"`

	// Kind: Kind of resource this is.
	Kind string `json:"kind,omitempty"`

	// OwnerDomain: Domain of source customer.
	OwnerDomain string `json:"ownerDomain,omitempty"`
}

type ActivityActor struct {
	// CallerType: User or OAuth 2LO request.
	CallerType string `json:"callerType,omitempty"`

	// Email: Email address of the user.
	Email string `json:"email,omitempty"`

	// Key: For OAuth 2LO API requests, consumer_key of the requestor.
	Key string `json:"key,omitempty"`

	// ProfileId: Obfuscated user id of the user.
	ProfileId string `json:"profileId,omitempty"`
}

type ActivityEvents struct {
	// Name: Name of event.
	Name string `json:"name,omitempty"`

	// Parameters: Parameter value pairs for various applications.
	Parameters []*ActivityEventsParameters `json:"parameters,omitempty"`

	// Type: Type of event.
	Type string `json:"type,omitempty"`
}

type ActivityEventsParameters struct {
	// BoolValue: Boolean value of the parameter.
	BoolValue bool `json:"boolValue,omitempty"`

	// IntValue: Integral value of the parameter.
	IntValue int64 `json:"intValue,omitempty,string"`

	// Name: The name of the parameter.
	Name string `json:"name,omitempty"`

	// Value: String value of the parameter.
	Value string `json:"value,omitempty"`
}

type ActivityId struct {
	// ApplicationName: Application name to which the event belongs.
	ApplicationName string `json:"applicationName,omitempty"`

	// CustomerId: Obfuscated customer ID of the source customer.
	CustomerId string `json:"customerId,omitempty"`

	// Time: Time of occurrence of the activity.
	Time string `json:"time,omitempty"`

	// UniqueQualifier: Unique qualifier if multiple events have the same
	// time.
	UniqueQualifier int64 `json:"uniqueQualifier,omitempty,string"`
}

type UsageReport struct {
	// Date: The date to which the record belongs.
	Date string `json:"date,omitempty"`

	// Entity: Information about the type of the item.
	Entity *UsageReportEntity `json:"entity,omitempty"`

	// Kind: The kind of object.
	Kind string `json:"kind,omitempty"`

	// Parameters: Parameter value pairs for various applications.
	Parameters []*UsageReportParameters `json:"parameters,omitempty"`
}

type UsageReportEntity struct {
	// CustomerId: Obfuscated customer id for the record.
	CustomerId string `json:"customerId,omitempty"`

	// ProfileId: Obfuscated user id for the record.
	ProfileId string `json:"profileId,omitempty"`

	// Type: The type of item, can be a customer or user.
	Type string `json:"type,omitempty"`

	// UserEmail: user's email.
	UserEmail string `json:"userEmail,omitempty"`
}

type UsageReportParameters struct {
	// BoolValue: Boolean value of the parameter.
	BoolValue bool `json:"boolValue,omitempty"`

	// DatetimeValue: RFC 3339 formatted value of the parameter.
	DatetimeValue string `json:"datetimeValue,omitempty"`

	// IntValue: Integral value of the parameter.
	IntValue int64 `json:"intValue,omitempty,string"`

	// Name: The name of the parameter.
	Name string `json:"name,omitempty"`

	// StringValue: String value of the parameter.
	StringValue string `json:"stringValue,omitempty"`
}

type UsageReports struct {
	// Kind: The kind of object.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token for retrieving the next page
	NextPageToken string `json:"nextPageToken,omitempty"`

	// UsageReports: Various application parameter records.
	UsageReports []*UsageReport `json:"usageReports,omitempty"`

	// Warnings: Warnings if any.
	Warnings []*UsageReportsWarnings `json:"warnings,omitempty"`
}

type UsageReportsWarnings struct {
	// Code: Machine readable code / warning type.
	Code string `json:"code,omitempty"`

	// Data: Key-Value pairs to give detailed information on the warning.
	Data []*UsageReportsWarningsData `json:"data,omitempty"`

	// Message: Human readable message for the warning.
	Message string `json:"message,omitempty"`
}

type UsageReportsWarningsData struct {
	// Key: Key associated with a key-value pair to give detailed
	// information on the warning.
	Key string `json:"key,omitempty"`

	// Value: Value associated with a key-value pair to give detailed
	// information on the warning.
	Value string `json:"value,omitempty"`
}

// method id "reports.activities.list":

type ActivitiesListCall struct {
	s               *Service
	userKey         string
	applicationName string
	opt_            map[string]interface{}
}

// List: Retrieves a list of activities for a specific customer and
// application.
func (r *ActivitiesService) List(userKey string, applicationName string) *ActivitiesListCall {
	c := &ActivitiesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.userKey = userKey
	c.applicationName = applicationName
	return c
}

// ActorIpAddress sets the optional parameter "actorIpAddress": IP
// Address of host where the event was performed. Supports both IPv4 and
// IPv6 addresses.
func (c *ActivitiesListCall) ActorIpAddress(actorIpAddress string) *ActivitiesListCall {
	c.opt_["actorIpAddress"] = actorIpAddress
	return c
}

// EndTime sets the optional parameter "endTime": Return events which
// occured at or before this time.
func (c *ActivitiesListCall) EndTime(endTime string) *ActivitiesListCall {
	c.opt_["endTime"] = endTime
	return c
}

// EventName sets the optional parameter "eventName": Name of the event
// being queried.
func (c *ActivitiesListCall) EventName(eventName string) *ActivitiesListCall {
	c.opt_["eventName"] = eventName
	return c
}

// Filters sets the optional parameter "filters": Event parameters in
// the form [parameter1 name][operator][parameter1 value],[parameter2
// name][operator][parameter2 value],...
func (c *ActivitiesListCall) Filters(filters string) *ActivitiesListCall {
	c.opt_["filters"] = filters
	return c
}

// MaxResults sets the optional parameter "maxResults": Number of
// activity records to be shown in each page.
func (c *ActivitiesListCall) MaxResults(maxResults int64) *ActivitiesListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": Token to specify
// next page.
func (c *ActivitiesListCall) PageToken(pageToken string) *ActivitiesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// StartTime sets the optional parameter "startTime": Return events
// which occured at or after this time.
func (c *ActivitiesListCall) StartTime(startTime string) *ActivitiesListCall {
	c.opt_["startTime"] = startTime
	return c
}

func (c *ActivitiesListCall) Do() (*Activities, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["actorIpAddress"]; ok {
		params.Set("actorIpAddress", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["endTime"]; ok {
		params.Set("endTime", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["eventName"]; ok {
		params.Set("eventName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["filters"]; ok {
		params.Set("filters", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["startTime"]; ok {
		params.Set("startTime", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/admin/reports/v1/", "activity/users/{userKey}/applications/{applicationName}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{userKey}", url.QueryEscape(c.userKey), 1)
	req.URL.Path = strings.Replace(req.URL.Path, "{applicationName}", url.QueryEscape(c.applicationName), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Activities)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of activities for a specific customer and application.",
	//   "httpMethod": "GET",
	//   "id": "reports.activities.list",
	//   "parameterOrder": [
	//     "userKey",
	//     "applicationName"
	//   ],
	//   "parameters": {
	//     "actorIpAddress": {
	//       "description": "IP Address of host where the event was performed. Supports both IPv4 and IPv6 addresses.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "applicationName": {
	//       "description": "Application name for which the events are to be retrieved.",
	//       "location": "path",
	//       "pattern": "(admin)|(docs)",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "endTime": {
	//       "description": "Return events which occured at or before this time.",
	//       "location": "query",
	//       "pattern": "(\\d\\d\\d\\d)-(\\d\\d)-(\\d\\d)T(\\d\\d):(\\d\\d):(\\d\\d)(?:\\.(\\d+))?(?:(Z)|([-+])(\\d\\d):(\\d\\d))",
	//       "type": "string"
	//     },
	//     "eventName": {
	//       "description": "Name of the event being queried.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "filters": {
	//       "description": "Event parameters in the form [parameter1 name][operator][parameter1 value],[parameter2 name][operator][parameter2 value],...",
	//       "location": "query",
	//       "pattern": "(.+[\u003c,\u003c=,==,\u003e=,\u003e,\u003c\u003e].+,)*(.+[\u003c,\u003c=,==,\u003e=,\u003e,\u003c\u003e].+)",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Number of activity records to be shown in each page.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "1000",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Token to specify next page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "startTime": {
	//       "description": "Return events which occured at or after this time.",
	//       "location": "query",
	//       "pattern": "(\\d\\d\\d\\d)-(\\d\\d)-(\\d\\d)T(\\d\\d):(\\d\\d):(\\d\\d)(?:\\.(\\d+))?(?:(Z)|([-+])(\\d\\d):(\\d\\d))",
	//       "type": "string"
	//     },
	//     "userKey": {
	//       "description": "Represents the profile id or the user email for which the data should be filtered. When 'all' is specified as the userKey, it returns usageReports for all users.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "activity/users/{userKey}/applications/{applicationName}",
	//   "response": {
	//     "$ref": "Activities"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.reports.audit.readonly"
	//   ]
	// }

}

// method id "reports.customerUsageReports.get":

type CustomerUsageReportsGetCall struct {
	s    *Service
	date string
	opt_ map[string]interface{}
}

// Get: Retrieves a report which is a collection of properties /
// statistics for a specific customer.
func (r *CustomerUsageReportsService) Get(date string) *CustomerUsageReportsGetCall {
	c := &CustomerUsageReportsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.date = date
	return c
}

// PageToken sets the optional parameter "pageToken": Token to specify
// next page.
func (c *CustomerUsageReportsGetCall) PageToken(pageToken string) *CustomerUsageReportsGetCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Parameters sets the optional parameter "parameters": Represents the
// application name, parameter name pairs to fetch in csv as
// app_name1:param_name1, app_name2:param_name2.
func (c *CustomerUsageReportsGetCall) Parameters(parameters string) *CustomerUsageReportsGetCall {
	c.opt_["parameters"] = parameters
	return c
}

func (c *CustomerUsageReportsGetCall) Do() (*UsageReports, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["parameters"]; ok {
		params.Set("parameters", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/admin/reports/v1/", "usage/dates/{date}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{date}", url.QueryEscape(c.date), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(UsageReports)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a report which is a collection of properties / statistics for a specific customer.",
	//   "httpMethod": "GET",
	//   "id": "reports.customerUsageReports.get",
	//   "parameterOrder": [
	//     "date"
	//   ],
	//   "parameters": {
	//     "date": {
	//       "description": "Represents the date in yyyy-mm-dd format for which the data is to be fetched.",
	//       "location": "path",
	//       "pattern": "(\\d){4}-(\\d){2}-(\\d){2}",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Token to specify next page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parameters": {
	//       "description": "Represents the application name, parameter name pairs to fetch in csv as app_name1:param_name1, app_name2:param_name2.",
	//       "location": "query",
	//       "pattern": "(((accounts)|(gmail)|(calendar)|(docs)):.+,)*(((accounts)|(gmail)|(calendar)|(docs)):.+)",
	//       "type": "string"
	//     }
	//   },
	//   "path": "usage/dates/{date}",
	//   "response": {
	//     "$ref": "UsageReports"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.reports.usage.readonly"
	//   ]
	// }

}

// method id "reports.userUsageReport.get":

type UserUsageReportGetCall struct {
	s       *Service
	userKey string
	date    string
	opt_    map[string]interface{}
}

// Get: Retrieves a report which is a collection of properties /
// statistics for a set of users.
func (r *UserUsageReportService) Get(userKey string, date string) *UserUsageReportGetCall {
	c := &UserUsageReportGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.userKey = userKey
	c.date = date
	return c
}

// Filters sets the optional parameter "filters": Represents the set of
// filters including parameter operator value.
func (c *UserUsageReportGetCall) Filters(filters string) *UserUsageReportGetCall {
	c.opt_["filters"] = filters
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return. Maximum allowed is 1000
func (c *UserUsageReportGetCall) MaxResults(maxResults int64) *UserUsageReportGetCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": Token to specify
// next page.
func (c *UserUsageReportGetCall) PageToken(pageToken string) *UserUsageReportGetCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Parameters sets the optional parameter "parameters": Represents the
// application name, parameter name pairs to fetch in csv as
// app_name1:param_name1, app_name2:param_name2.
func (c *UserUsageReportGetCall) Parameters(parameters string) *UserUsageReportGetCall {
	c.opt_["parameters"] = parameters
	return c
}

func (c *UserUsageReportGetCall) Do() (*UsageReports, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["filters"]; ok {
		params.Set("filters", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["parameters"]; ok {
		params.Set("parameters", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/admin/reports/v1/", "usage/users/{userKey}/dates/{date}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{userKey}", url.QueryEscape(c.userKey), 1)
	req.URL.Path = strings.Replace(req.URL.Path, "{date}", url.QueryEscape(c.date), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(UsageReports)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a report which is a collection of properties / statistics for a set of users.",
	//   "httpMethod": "GET",
	//   "id": "reports.userUsageReport.get",
	//   "parameterOrder": [
	//     "userKey",
	//     "date"
	//   ],
	//   "parameters": {
	//     "date": {
	//       "description": "Represents the date in yyyy-mm-dd format for which the data is to be fetched.",
	//       "location": "path",
	//       "pattern": "(\\d){4}-(\\d){2}-(\\d){2}",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "filters": {
	//       "description": "Represents the set of filters including parameter operator value.",
	//       "location": "query",
	//       "pattern": "(((accounts)|(gmail)|(calendar)|(docs)):.+[\u003c,\u003c=,==,\u003e=,\u003e,!=].+,)*(((accounts)|(gmail)|(calendar)|(docs)):.+[\u003c,\u003c=,==,\u003e=,\u003e,!=].+)",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return. Maximum allowed is 1000",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "1000",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Token to specify next page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parameters": {
	//       "description": "Represents the application name, parameter name pairs to fetch in csv as app_name1:param_name1, app_name2:param_name2.",
	//       "location": "query",
	//       "pattern": "(((accounts)|(gmail)|(calendar)|(docs)):.+,)*(((accounts)|(gmail)|(calendar)|(docs)):.+)",
	//       "type": "string"
	//     },
	//     "userKey": {
	//       "description": "Represents the profile id or the user email for which the data should be filtered.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "usage/users/{userKey}/dates/{date}",
	//   "response": {
	//     "$ref": "UsageReports"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.reports.usage.readonly"
	//   ]
	// }

}
