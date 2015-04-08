// Package pubsub provides access to the Google Cloud Pub/Sub API.
//
// Usage example:
//
//   import "google.golang.org/api/pubsub/v1beta2"
//   ...
//   pubsubService, err := pubsub.New(oauthHttpClient)
package pubsub

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

const apiId = "pubsub:v1beta2"
const apiName = "pubsub"
const apiVersion = "v1beta2"
const basePath = "https://pubsub.googleapis.com/v1beta2/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View and manage Pub/Sub topics and subscriptions
	PubsubScope = "https://www.googleapis.com/auth/pubsub"
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
	rs.Subscriptions = NewProjectsSubscriptionsService(s)
	rs.Topics = NewProjectsTopicsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Subscriptions *ProjectsSubscriptionsService

	Topics *ProjectsTopicsService
}

func NewProjectsSubscriptionsService(s *Service) *ProjectsSubscriptionsService {
	rs := &ProjectsSubscriptionsService{s: s}
	return rs
}

type ProjectsSubscriptionsService struct {
	s *Service
}

func NewProjectsTopicsService(s *Service) *ProjectsTopicsService {
	rs := &ProjectsTopicsService{s: s}
	rs.Subscriptions = NewProjectsTopicsSubscriptionsService(s)
	return rs
}

type ProjectsTopicsService struct {
	s *Service

	Subscriptions *ProjectsTopicsSubscriptionsService
}

func NewProjectsTopicsSubscriptionsService(s *Service) *ProjectsTopicsSubscriptionsService {
	rs := &ProjectsTopicsSubscriptionsService{s: s}
	return rs
}

type ProjectsTopicsSubscriptionsService struct {
	s *Service
}

type AcknowledgeRequest struct {
	AckIds []string `json:"ackIds,omitempty"`
}

type Empty struct {
}

type ListSubscriptionsResponse struct {
	NextPageToken string `json:"nextPageToken,omitempty"`

	Subscriptions []*Subscription `json:"subscriptions,omitempty"`
}

type ListTopicSubscriptionsResponse struct {
	NextPageToken string `json:"nextPageToken,omitempty"`

	Subscriptions []string `json:"subscriptions,omitempty"`
}

type ListTopicsResponse struct {
	NextPageToken string `json:"nextPageToken,omitempty"`

	Topics []*Topic `json:"topics,omitempty"`
}

type ModifyAckDeadlineRequest struct {
	AckDeadlineSeconds int64 `json:"ackDeadlineSeconds,omitempty"`

	AckId string `json:"ackId,omitempty"`
}

type ModifyPushConfigRequest struct {
	PushConfig *PushConfig `json:"pushConfig,omitempty"`
}

type PublishRequest struct {
	Messages []*PubsubMessage `json:"messages,omitempty"`
}

type PublishResponse struct {
	MessageIds []string `json:"messageIds,omitempty"`
}

type PubsubMessage struct {
	Attributes map[string]string `json:"attributes,omitempty"`

	Data string `json:"data,omitempty"`

	MessageId string `json:"messageId,omitempty"`
}

type PullRequest struct {
	MaxMessages int64 `json:"maxMessages,omitempty"`

	ReturnImmediately bool `json:"returnImmediately,omitempty"`
}

type PullResponse struct {
	ReceivedMessages []*ReceivedMessage `json:"receivedMessages,omitempty"`
}

type PushConfig struct {
	Attributes map[string]string `json:"attributes,omitempty"`

	PushEndpoint string `json:"pushEndpoint,omitempty"`
}

type ReceivedMessage struct {
	AckId string `json:"ackId,omitempty"`

	Message *PubsubMessage `json:"message,omitempty"`
}

type Subscription struct {
	AckDeadlineSeconds int64 `json:"ackDeadlineSeconds,omitempty"`

	Name string `json:"name,omitempty"`

	PushConfig *PushConfig `json:"pushConfig,omitempty"`

	Topic string `json:"topic,omitempty"`
}

type Topic struct {
	Name string `json:"name,omitempty"`
}

// method id "pubsub.projects.subscriptions.acknowledge":

type ProjectsSubscriptionsAcknowledgeCall struct {
	s                  *Service
	subscription       string
	acknowledgerequest *AcknowledgeRequest
	opt_               map[string]interface{}
}

// Acknowledge: Acknowledges the messages associated with the ack tokens
// in the AcknowledgeRequest. The Pub/Sub system can remove the relevant
// messages from the subscription. Acknowledging a message whose ack
// deadline has expired may succeed, but such a message may be
// redelivered later. Acknowledging a message more than once will not
// result in an error.
func (r *ProjectsSubscriptionsService) Acknowledge(subscription string, acknowledgerequest *AcknowledgeRequest) *ProjectsSubscriptionsAcknowledgeCall {
	c := &ProjectsSubscriptionsAcknowledgeCall{s: r.s, opt_: make(map[string]interface{})}
	c.subscription = subscription
	c.acknowledgerequest = acknowledgerequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSubscriptionsAcknowledgeCall) Fields(s ...googleapi.Field) *ProjectsSubscriptionsAcknowledgeCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsSubscriptionsAcknowledgeCall) Do() (*Empty, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.acknowledgerequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+subscription}:acknowledge")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"subscription": c.subscription,
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
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Acknowledges the messages associated with the ack tokens in the AcknowledgeRequest. The Pub/Sub system can remove the relevant messages from the subscription. Acknowledging a message whose ack deadline has expired may succeed, but such a message may be redelivered later. Acknowledging a message more than once will not result in an error.",
	//   "httpMethod": "POST",
	//   "id": "pubsub.projects.subscriptions.acknowledge",
	//   "parameterOrder": [
	//     "subscription"
	//   ],
	//   "parameters": {
	//     "subscription": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+subscription}:acknowledge",
	//   "request": {
	//     "$ref": "AcknowledgeRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.subscriptions.create":

type ProjectsSubscriptionsCreateCall struct {
	s            *Service
	name         string
	subscription *Subscription
	opt_         map[string]interface{}
}

// Create: Creates a subscription to a given topic for a given
// subscriber. If the subscription already exists, returns
// ALREADY_EXISTS. If the corresponding topic doesn't exist, returns
// NOT_FOUND. If the name is not provided in the request, the server
// will assign a random name for this subscription on the same project
// as the topic.
func (r *ProjectsSubscriptionsService) Create(name string, subscription *Subscription) *ProjectsSubscriptionsCreateCall {
	c := &ProjectsSubscriptionsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.name = name
	c.subscription = subscription
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSubscriptionsCreateCall) Fields(s ...googleapi.Field) *ProjectsSubscriptionsCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsSubscriptionsCreateCall) Do() (*Subscription, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.subscription)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+name}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
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
	var ret *Subscription
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a subscription to a given topic for a given subscriber. If the subscription already exists, returns ALREADY_EXISTS. If the corresponding topic doesn't exist, returns NOT_FOUND. If the name is not provided in the request, the server will assign a random name for this subscription on the same project as the topic.",
	//   "httpMethod": "PUT",
	//   "id": "pubsub.projects.subscriptions.create",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+name}",
	//   "request": {
	//     "$ref": "Subscription"
	//   },
	//   "response": {
	//     "$ref": "Subscription"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.subscriptions.delete":

type ProjectsSubscriptionsDeleteCall struct {
	s            *Service
	subscription string
	opt_         map[string]interface{}
}

// Delete: Deletes an existing subscription. All pending messages in the
// subscription are immediately dropped. Calls to Pull after deletion
// will return NOT_FOUND. After a subscription is deleted, a new one may
// be created with the same name, but the new one has no association
// with the old subscription, or its topic unless the same topic is
// specified.
func (r *ProjectsSubscriptionsService) Delete(subscription string) *ProjectsSubscriptionsDeleteCall {
	c := &ProjectsSubscriptionsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.subscription = subscription
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSubscriptionsDeleteCall) Fields(s ...googleapi.Field) *ProjectsSubscriptionsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsSubscriptionsDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+subscription}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"subscription": c.subscription,
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
	//   "description": "Deletes an existing subscription. All pending messages in the subscription are immediately dropped. Calls to Pull after deletion will return NOT_FOUND. After a subscription is deleted, a new one may be created with the same name, but the new one has no association with the old subscription, or its topic unless the same topic is specified.",
	//   "httpMethod": "DELETE",
	//   "id": "pubsub.projects.subscriptions.delete",
	//   "parameterOrder": [
	//     "subscription"
	//   ],
	//   "parameters": {
	//     "subscription": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+subscription}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.subscriptions.get":

type ProjectsSubscriptionsGetCall struct {
	s            *Service
	subscription string
	opt_         map[string]interface{}
}

// Get: Gets the configuration details of a subscription.
func (r *ProjectsSubscriptionsService) Get(subscription string) *ProjectsSubscriptionsGetCall {
	c := &ProjectsSubscriptionsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.subscription = subscription
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSubscriptionsGetCall) Fields(s ...googleapi.Field) *ProjectsSubscriptionsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsSubscriptionsGetCall) Do() (*Subscription, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+subscription}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"subscription": c.subscription,
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
	var ret *Subscription
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the configuration details of a subscription.",
	//   "httpMethod": "GET",
	//   "id": "pubsub.projects.subscriptions.get",
	//   "parameterOrder": [
	//     "subscription"
	//   ],
	//   "parameters": {
	//     "subscription": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+subscription}",
	//   "response": {
	//     "$ref": "Subscription"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.subscriptions.list":

type ProjectsSubscriptionsListCall struct {
	s       *Service
	project string
	opt_    map[string]interface{}
}

// List: Lists matching subscriptions.
func (r *ProjectsSubscriptionsService) List(project string) *ProjectsSubscriptionsListCall {
	c := &ProjectsSubscriptionsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	return c
}

// PageSize sets the optional parameter "pageSize":
func (c *ProjectsSubscriptionsListCall) PageSize(pageSize int64) *ProjectsSubscriptionsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken":
func (c *ProjectsSubscriptionsListCall) PageToken(pageToken string) *ProjectsSubscriptionsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSubscriptionsListCall) Fields(s ...googleapi.Field) *ProjectsSubscriptionsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsSubscriptionsListCall) Do() (*ListSubscriptionsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+project}/subscriptions")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
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
	var ret *ListSubscriptionsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists matching subscriptions.",
	//   "httpMethod": "GET",
	//   "id": "pubsub.projects.subscriptions.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+project}/subscriptions",
	//   "response": {
	//     "$ref": "ListSubscriptionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.subscriptions.modifyAckDeadline":

type ProjectsSubscriptionsModifyAckDeadlineCall struct {
	s                        *Service
	subscription             string
	modifyackdeadlinerequest *ModifyAckDeadlineRequest
	opt_                     map[string]interface{}
}

// ModifyAckDeadline: Modifies the ack deadline for a specific message.
// This method is useful to indicate that more time is needed to process
// a message by the subscriber, or to make the message available for
// redelivery if the processing was interrupted.
func (r *ProjectsSubscriptionsService) ModifyAckDeadline(subscription string, modifyackdeadlinerequest *ModifyAckDeadlineRequest) *ProjectsSubscriptionsModifyAckDeadlineCall {
	c := &ProjectsSubscriptionsModifyAckDeadlineCall{s: r.s, opt_: make(map[string]interface{})}
	c.subscription = subscription
	c.modifyackdeadlinerequest = modifyackdeadlinerequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSubscriptionsModifyAckDeadlineCall) Fields(s ...googleapi.Field) *ProjectsSubscriptionsModifyAckDeadlineCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsSubscriptionsModifyAckDeadlineCall) Do() (*Empty, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.modifyackdeadlinerequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+subscription}:modifyAckDeadline")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"subscription": c.subscription,
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
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Modifies the ack deadline for a specific message. This method is useful to indicate that more time is needed to process a message by the subscriber, or to make the message available for redelivery if the processing was interrupted.",
	//   "httpMethod": "POST",
	//   "id": "pubsub.projects.subscriptions.modifyAckDeadline",
	//   "parameterOrder": [
	//     "subscription"
	//   ],
	//   "parameters": {
	//     "subscription": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+subscription}:modifyAckDeadline",
	//   "request": {
	//     "$ref": "ModifyAckDeadlineRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.subscriptions.modifyPushConfig":

type ProjectsSubscriptionsModifyPushConfigCall struct {
	s                       *Service
	subscription            string
	modifypushconfigrequest *ModifyPushConfigRequest
	opt_                    map[string]interface{}
}

// ModifyPushConfig: Modifies the PushConfig for a specified
// subscription. This may be used to change a push subscription to a
// pull one (signified by an empty PushConfig) or vice versa, or change
// the endpoint URL and other attributes of a push subscription.
// Messages will accumulate for delivery continuously through the call
// regardless of changes to the PushConfig.
func (r *ProjectsSubscriptionsService) ModifyPushConfig(subscription string, modifypushconfigrequest *ModifyPushConfigRequest) *ProjectsSubscriptionsModifyPushConfigCall {
	c := &ProjectsSubscriptionsModifyPushConfigCall{s: r.s, opt_: make(map[string]interface{})}
	c.subscription = subscription
	c.modifypushconfigrequest = modifypushconfigrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSubscriptionsModifyPushConfigCall) Fields(s ...googleapi.Field) *ProjectsSubscriptionsModifyPushConfigCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsSubscriptionsModifyPushConfigCall) Do() (*Empty, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.modifypushconfigrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+subscription}:modifyPushConfig")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"subscription": c.subscription,
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
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Modifies the PushConfig for a specified subscription. This may be used to change a push subscription to a pull one (signified by an empty PushConfig) or vice versa, or change the endpoint URL and other attributes of a push subscription. Messages will accumulate for delivery continuously through the call regardless of changes to the PushConfig.",
	//   "httpMethod": "POST",
	//   "id": "pubsub.projects.subscriptions.modifyPushConfig",
	//   "parameterOrder": [
	//     "subscription"
	//   ],
	//   "parameters": {
	//     "subscription": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+subscription}:modifyPushConfig",
	//   "request": {
	//     "$ref": "ModifyPushConfigRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.subscriptions.pull":

type ProjectsSubscriptionsPullCall struct {
	s            *Service
	subscription string
	pullrequest  *PullRequest
	opt_         map[string]interface{}
}

// Pull: Pulls messages from the server. Returns an empty list if there
// are no messages available in the backlog. The server may return
// UNAVAILABLE if there are too many concurrent pull requests pending
// for the given subscription.
func (r *ProjectsSubscriptionsService) Pull(subscription string, pullrequest *PullRequest) *ProjectsSubscriptionsPullCall {
	c := &ProjectsSubscriptionsPullCall{s: r.s, opt_: make(map[string]interface{})}
	c.subscription = subscription
	c.pullrequest = pullrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSubscriptionsPullCall) Fields(s ...googleapi.Field) *ProjectsSubscriptionsPullCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsSubscriptionsPullCall) Do() (*PullResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.pullrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+subscription}:pull")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"subscription": c.subscription,
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
	var ret *PullResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Pulls messages from the server. Returns an empty list if there are no messages available in the backlog. The server may return UNAVAILABLE if there are too many concurrent pull requests pending for the given subscription.",
	//   "httpMethod": "POST",
	//   "id": "pubsub.projects.subscriptions.pull",
	//   "parameterOrder": [
	//     "subscription"
	//   ],
	//   "parameters": {
	//     "subscription": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+subscription}:pull",
	//   "request": {
	//     "$ref": "PullRequest"
	//   },
	//   "response": {
	//     "$ref": "PullResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.topics.create":

type ProjectsTopicsCreateCall struct {
	s     *Service
	name  string
	topic *Topic
	opt_  map[string]interface{}
}

// Create: Creates the given topic with the given name.
func (r *ProjectsTopicsService) Create(name string, topic *Topic) *ProjectsTopicsCreateCall {
	c := &ProjectsTopicsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.name = name
	c.topic = topic
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTopicsCreateCall) Fields(s ...googleapi.Field) *ProjectsTopicsCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTopicsCreateCall) Do() (*Topic, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.topic)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+name}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
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
	var ret *Topic
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates the given topic with the given name.",
	//   "httpMethod": "PUT",
	//   "id": "pubsub.projects.topics.create",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+name}",
	//   "request": {
	//     "$ref": "Topic"
	//   },
	//   "response": {
	//     "$ref": "Topic"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.topics.delete":

type ProjectsTopicsDeleteCall struct {
	s     *Service
	topic string
	opt_  map[string]interface{}
}

// Delete: Deletes the topic with the given name. Returns NOT_FOUND if
// the topic does not exist. After a topic is deleted, a new topic may
// be created with the same name; this is an entirely new topic with
// none of the old configuration or subscriptions. Existing
// subscriptions to this topic are not deleted.
func (r *ProjectsTopicsService) Delete(topic string) *ProjectsTopicsDeleteCall {
	c := &ProjectsTopicsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.topic = topic
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTopicsDeleteCall) Fields(s ...googleapi.Field) *ProjectsTopicsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTopicsDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+topic}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"topic": c.topic,
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
	//   "description": "Deletes the topic with the given name. Returns NOT_FOUND if the topic does not exist. After a topic is deleted, a new topic may be created with the same name; this is an entirely new topic with none of the old configuration or subscriptions. Existing subscriptions to this topic are not deleted.",
	//   "httpMethod": "DELETE",
	//   "id": "pubsub.projects.topics.delete",
	//   "parameterOrder": [
	//     "topic"
	//   ],
	//   "parameters": {
	//     "topic": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+topic}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.topics.get":

type ProjectsTopicsGetCall struct {
	s     *Service
	topic string
	opt_  map[string]interface{}
}

// Get: Gets the configuration of a topic.
func (r *ProjectsTopicsService) Get(topic string) *ProjectsTopicsGetCall {
	c := &ProjectsTopicsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.topic = topic
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTopicsGetCall) Fields(s ...googleapi.Field) *ProjectsTopicsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTopicsGetCall) Do() (*Topic, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+topic}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"topic": c.topic,
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
	var ret *Topic
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the configuration of a topic.",
	//   "httpMethod": "GET",
	//   "id": "pubsub.projects.topics.get",
	//   "parameterOrder": [
	//     "topic"
	//   ],
	//   "parameters": {
	//     "topic": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+topic}",
	//   "response": {
	//     "$ref": "Topic"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.topics.list":

type ProjectsTopicsListCall struct {
	s       *Service
	project string
	opt_    map[string]interface{}
}

// List: Lists matching topics.
func (r *ProjectsTopicsService) List(project string) *ProjectsTopicsListCall {
	c := &ProjectsTopicsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	return c
}

// PageSize sets the optional parameter "pageSize":
func (c *ProjectsTopicsListCall) PageSize(pageSize int64) *ProjectsTopicsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken":
func (c *ProjectsTopicsListCall) PageToken(pageToken string) *ProjectsTopicsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTopicsListCall) Fields(s ...googleapi.Field) *ProjectsTopicsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTopicsListCall) Do() (*ListTopicsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+project}/topics")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
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
	var ret *ListTopicsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists matching topics.",
	//   "httpMethod": "GET",
	//   "id": "pubsub.projects.topics.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+project}/topics",
	//   "response": {
	//     "$ref": "ListTopicsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.topics.publish":

type ProjectsTopicsPublishCall struct {
	s              *Service
	topic          string
	publishrequest *PublishRequest
	opt_           map[string]interface{}
}

// Publish: Adds one or more messages to the topic. Returns NOT_FOUND if
// the topic does not exist.
func (r *ProjectsTopicsService) Publish(topic string, publishrequest *PublishRequest) *ProjectsTopicsPublishCall {
	c := &ProjectsTopicsPublishCall{s: r.s, opt_: make(map[string]interface{})}
	c.topic = topic
	c.publishrequest = publishrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTopicsPublishCall) Fields(s ...googleapi.Field) *ProjectsTopicsPublishCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTopicsPublishCall) Do() (*PublishResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.publishrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+topic}:publish")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"topic": c.topic,
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
	var ret *PublishResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Adds one or more messages to the topic. Returns NOT_FOUND if the topic does not exist.",
	//   "httpMethod": "POST",
	//   "id": "pubsub.projects.topics.publish",
	//   "parameterOrder": [
	//     "topic"
	//   ],
	//   "parameters": {
	//     "topic": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+topic}:publish",
	//   "request": {
	//     "$ref": "PublishRequest"
	//   },
	//   "response": {
	//     "$ref": "PublishResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.topics.subscriptions.list":

type ProjectsTopicsSubscriptionsListCall struct {
	s     *Service
	topic string
	opt_  map[string]interface{}
}

// List: Lists the name of the subscriptions for this topic.
func (r *ProjectsTopicsSubscriptionsService) List(topic string) *ProjectsTopicsSubscriptionsListCall {
	c := &ProjectsTopicsSubscriptionsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.topic = topic
	return c
}

// PageSize sets the optional parameter "pageSize":
func (c *ProjectsTopicsSubscriptionsListCall) PageSize(pageSize int64) *ProjectsTopicsSubscriptionsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken":
func (c *ProjectsTopicsSubscriptionsListCall) PageToken(pageToken string) *ProjectsTopicsSubscriptionsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTopicsSubscriptionsListCall) Fields(s ...googleapi.Field) *ProjectsTopicsSubscriptionsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTopicsSubscriptionsListCall) Do() (*ListTopicSubscriptionsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+topic}/subscriptions")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"topic": c.topic,
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
	var ret *ListTopicSubscriptionsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the name of the subscriptions for this topic.",
	//   "httpMethod": "GET",
	//   "id": "pubsub.projects.topics.subscriptions.list",
	//   "parameterOrder": [
	//     "topic"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "topic": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+topic}/subscriptions",
	//   "response": {
	//     "$ref": "ListTopicSubscriptionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}
