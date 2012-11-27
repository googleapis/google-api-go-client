// Package audit provides access to the Enterprise Audit API.
//
// See http://code.google.com/googleapps/domain/audit_admin/v1/getting_started.html
//
// Usage example:
//
//   import "code.google.com/p/google-api-go-client/audit/v1"
//   ...
//   auditService, err := audit.New(oauthHttpClient)
package audit

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

const apiId = "audit:v1"
const apiName = "audit"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/apps/reporting/audit/v1/"

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Activities = &ActivitiesService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	Activities *ActivitiesService
}

type ActivitiesService struct {
	s *Service
}

type Activities struct {
	// Items: Each record in read response.
	Items []*Activity `json:"items,omitempty"`

	// Kind: Kind of list response this is.
	Kind string `json:"kind,omitempty"`

	// Next: Next page URL.
	Next string `json:"next,omitempty"`
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
	// ApplicationId: ID of application which interacted on behalf of the
	// user.
	ApplicationId int64 `json:"applicationId,omitempty,string"`

	// CallerType: User or OAuth 2LO request.
	CallerType string `json:"callerType,omitempty"`

	// Email: Email address of the user.
	Email string `json:"email,omitempty"`

	// Key: For OAuth 2LO API requests, consumer_key of the requestor.
	Key string `json:"key,omitempty"`
}

type ActivityEvents struct {
	// EventType: Type of event.
	EventType string `json:"eventType,omitempty"`

	// Name: Name of event.
	Name string `json:"name,omitempty"`

	// Parameters: Event parameters.
	Parameters []*ActivityEventsParameters `json:"parameters,omitempty"`
}

type ActivityEventsParameters struct {
	// Name: Name of the parameter.
	Name string `json:"name,omitempty"`

	// Value: Value of the parameter.
	Value string `json:"value,omitempty"`
}

type ActivityId struct {
	// ApplicationId: Application ID of the source application.
	ApplicationId int64 `json:"applicationId,omitempty,string"`

	// CustomerId: Obfuscated customer ID of the source customer.
	CustomerId string `json:"customerId,omitempty"`

	// Time: Time of occurrence of the activity.
	Time string `json:"time,omitempty"`

	// UniqQualifier: Unique qualifier if multiple events have the same
	// time.
	UniqQualifier int64 `json:"uniqQualifier,omitempty,string"`
}

// method id "audit.activities.list":

type ActivitiesListCall struct {
	s             *Service
	customerId    string
	applicationId int64
	opt_          map[string]interface{}
}

// List: Retrieves a list of activities for a specific customer and
// application.
func (r *ActivitiesService) List(customerId string, applicationId int64) *ActivitiesListCall {
	c := &ActivitiesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.customerId = customerId
	c.applicationId = applicationId
	return c
}

// ActorApplicationId sets the optional parameter "actorApplicationId":
// Application ID of the application which interacted on behalf of the
// user while performing the event.
func (c *ActivitiesListCall) ActorApplicationId(actorApplicationId int64) *ActivitiesListCall {
	c.opt_["actorApplicationId"] = actorApplicationId
	return c
}

// ActorEmail sets the optional parameter "actorEmail": Email address of
// the user who performed the action.
func (c *ActivitiesListCall) ActorEmail(actorEmail string) *ActivitiesListCall {
	c.opt_["actorEmail"] = actorEmail
	return c
}

// ActorIpAddress sets the optional parameter "actorIpAddress": IP
// Address of host where the event was performed. Supports both IPv4 and
// IPv6 addresses.
func (c *ActivitiesListCall) ActorIpAddress(actorIpAddress string) *ActivitiesListCall {
	c.opt_["actorIpAddress"] = actorIpAddress
	return c
}

// Caller sets the optional parameter "caller": Type of the caller.
func (c *ActivitiesListCall) Caller(caller string) *ActivitiesListCall {
	c.opt_["caller"] = caller
	return c
}

// ContinuationToken sets the optional parameter "continuationToken":
// Next page URL.
func (c *ActivitiesListCall) ContinuationToken(continuationToken string) *ActivitiesListCall {
	c.opt_["continuationToken"] = continuationToken
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

// MaxResults sets the optional parameter "maxResults": Number of
// activity records to be shown in each page.
func (c *ActivitiesListCall) MaxResults(maxResults int64) *ActivitiesListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// Parameters sets the optional parameter "parameters": Event parameters
// in the form [parameter1 name]:[parameter1 value],[parameter2
// name]:[parameter2 value],...
func (c *ActivitiesListCall) Parameters(parameters string) *ActivitiesListCall {
	c.opt_["parameters"] = parameters
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
	if v, ok := c.opt_["actorApplicationId"]; ok {
		params.Set("actorApplicationId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["actorEmail"]; ok {
		params.Set("actorEmail", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["actorIpAddress"]; ok {
		params.Set("actorIpAddress", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["caller"]; ok {
		params.Set("caller", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["continuationToken"]; ok {
		params.Set("continuationToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["endTime"]; ok {
		params.Set("endTime", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["eventName"]; ok {
		params.Set("eventName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["parameters"]; ok {
		params.Set("parameters", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["startTime"]; ok {
		params.Set("startTime", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/apps/reporting/audit/v1/", "{customerId}/{applicationId}")
	urls = strings.Replace(urls, "{customerId}", cleanPathString(c.customerId), 1)
	urls = strings.Replace(urls, "{applicationId}", strconv.FormatInt(c.applicationId, 10), 1)
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
	ret := new(Activities)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of activities for a specific customer and application.",
	//   "httpMethod": "GET",
	//   "id": "audit.activities.list",
	//   "parameterOrder": [
	//     "customerId",
	//     "applicationId"
	//   ],
	//   "parameters": {
	//     "actorApplicationId": {
	//       "description": "Application ID of the application which interacted on behalf of the user while performing the event.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "actorEmail": {
	//       "description": "Email address of the user who performed the action.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "actorIpAddress": {
	//       "description": "IP Address of host where the event was performed. Supports both IPv4 and IPv6 addresses.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "applicationId": {
	//       "description": "Application ID of the application on which the event was performed.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "caller": {
	//       "description": "Type of the caller.",
	//       "enum": [
	//         "application_owner",
	//         "customer"
	//       ],
	//       "enumDescriptions": [
	//         "Caller is an application owner.",
	//         "Caller is a customer."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "continuationToken": {
	//       "description": "Next page URL.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "customerId": {
	//       "description": "Represents the customer who is the owner of target object on which action was performed.",
	//       "location": "path",
	//       "pattern": "C.+",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "endTime": {
	//       "description": "Return events which occured at or before this time.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "eventName": {
	//       "description": "Name of the event being queried.",
	//       "location": "query",
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
	//     "parameters": {
	//       "description": "Event parameters in the form [parameter1 name]:[parameter1 value],[parameter2 name]:[parameter2 value],...",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "startTime": {
	//       "description": "Return events which occured at or after this time.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "{customerId}/{applicationId}",
	//   "response": {
	//     "$ref": "Activities"
	//   }
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
