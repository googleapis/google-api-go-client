// Package deploymentmanager provides access to the Google Cloud Deployment Manager API V2.
//
// See https://developers.google.com/deployment-manager/
//
// Usage example:
//
//   import "google.golang.org/api/deploymentmanager/v2beta1"
//   ...
//   deploymentmanagerService, err := deploymentmanager.New(oauthHttpClient)
package deploymentmanager

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

const apiId = "deploymentmanager:v2beta1"
const apiName = "deploymentmanager"
const apiVersion = "v2beta1"
const basePath = "https://www.googleapis.com/deploymentmanager/v2beta1/projects/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View and manage your Google Cloud Platform management resources and
	// deployment status information
	NdevCloudmanScope = "https://www.googleapis.com/auth/ndev.cloudman"

	// View your Google Cloud Platform management resources and deployment
	// status information
	NdevCloudmanReadonlyScope = "https://www.googleapis.com/auth/ndev.cloudman.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Deployments = NewDeploymentsService(s)
	s.Manifests = NewManifestsService(s)
	s.Operations = NewOperationsService(s)
	s.Resources = NewResourcesService(s)
	s.Types = NewTypesService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Deployments *DeploymentsService

	Manifests *ManifestsService

	Operations *OperationsService

	Resources *ResourcesService

	Types *TypesService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewDeploymentsService(s *Service) *DeploymentsService {
	rs := &DeploymentsService{s: s}
	return rs
}

type DeploymentsService struct {
	s *Service
}

func NewManifestsService(s *Service) *ManifestsService {
	rs := &ManifestsService{s: s}
	return rs
}

type ManifestsService struct {
	s *Service
}

func NewOperationsService(s *Service) *OperationsService {
	rs := &OperationsService{s: s}
	return rs
}

type OperationsService struct {
	s *Service
}

func NewResourcesService(s *Service) *ResourcesService {
	rs := &ResourcesService{s: s}
	return rs
}

type ResourcesService struct {
	s *Service
}

func NewTypesService(s *Service) *TypesService {
	rs := &TypesService{s: s}
	return rs
}

type TypesService struct {
	s *Service
}

// Deployment: Next available tag: 8
type Deployment struct {
	// Description: ! An optional user-provided description of the
	// deployment.
	Description string `json:"description,omitempty"`

	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// Manifest: ! [Output Only] URL of the manifest representing the full
	// configuration ! of this deployment.
	Manifest string `json:"manifest,omitempty"`

	// Name: ! The name of the deployment, which must be unique within the
	// project.
	Name string `json:"name,omitempty"`

	// TargetConfig: ! [Input Only] The YAML configuration to use in
	// processing this deployment. ! ! When you create a deployment, the
	// server creates a new manifest with the ! given YAML configuration and
	// sets the `manifest` property to the URL of ! the manifest resource.
	TargetConfig string `json:"targetConfig,omitempty"`
}

// DeploymentsListResponse: ! A response containing a partial list of
// deployments and a page token used ! to build the next request if the
// request has been truncated. Next available tag: 4
type DeploymentsListResponse struct {
	// Deployments: ! The deployments contained in this response.
	Deployments []*Deployment `json:"deployments,omitempty"`

	// NextPageToken: ! A token used to continue a truncated list request.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

// Manifest: Next available tag: 10
type Manifest struct {
	// Config: v2beta1: YAML with config - described above v2beta2: YAML +
	// templates. ! The YAML configuration for this manifest.
	Config string `json:"config,omitempty"`

	// EvaluatedConfig: ! [Output Only] The fully-expanded configuration
	// file, including any ! templates and references.
	EvaluatedConfig string `json:"evaluatedConfig,omitempty"`

	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// Name: ! [Output Only] The name of the manifest.
	Name string `json:"name,omitempty"`

	// SelfLink: [Output Only] Self link for the manifest.
	SelfLink string `json:"selfLink,omitempty"`
}

// ManifestsListResponse: ! A response containing a partial list of
// manifests and a page token used ! to build the next request if the
// request has been truncated. Next available tag: 4
type ManifestsListResponse struct {
	// Manifests: ! Manifests contained in this list response.
	Manifests []*Manifest `json:"manifests,omitempty"`

	// NextPageToken: ! A token used to continue a truncated list request.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

// Operation: ! An operation resource, used to manage asynchronous API
// requests. Next available tag: 24
type Operation struct {
	// CreationTimestamp: ! [Output Only] Creation timestamp in RFC3339 text
	// format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// EndTime: ! [Output Only] The time that this operation was completed.
	// This is in ! RFC3339 format.
	EndTime string `json:"endTime,omitempty"`

	// Error: ! [Output Only] If errors occurred during processing of this
	// operation, ! this field will be populated.
	Error *OperationError `json:"error,omitempty"`

	// HttpErrorMessage: ! [Output Only] If operation fails, the HTTP error
	// message returned, ! e.g. NOT FOUND.
	HttpErrorMessage string `json:"httpErrorMessage,omitempty"`

	// HttpErrorStatusCode: ! [Output Only] If operation fails, the HTTP
	// error status code returned, ! e.g. 404.
	HttpErrorStatusCode int64 `json:"httpErrorStatusCode,omitempty"`

	// Id: ! [Output Only] Unique identifier for the resource; defined by
	// the server.
	Id uint64 `json:"id,omitempty,string"`

	// InsertTime: ! [Output Only] The time that this operation was
	// requested. ! This is in RFC 3339 format.
	InsertTime string `json:"insertTime,omitempty"`

	// Name: ! [Output Only] Name of the operation.
	Name string `json:"name,omitempty"`

	// OperationType: ! [Output Only] Type of the operation. Examples
	// include "insert", or ! "delete"
	OperationType string `json:"operationType,omitempty"`

	// Progress: ! [Output Only] An optional progress indicator that ranges
	// from 0 to 100. ! There is no requirement that this be linear or
	// support any granularity ! of operations. This should not be used to
	// guess at when the operation will ! be complete. This number should be
	// monotonically increasing as the ! operation progresses.
	Progress int64 `json:"progress,omitempty"`

	// SelfLink: [Output Only] Self link for the manifest.
	SelfLink string `json:"selfLink,omitempty"`

	// StartTime: ! [Output Only] The time that this operation was started
	// by the server. ! This is in RFC 3339 format.
	StartTime string `json:"startTime,omitempty"`

	// Status: ! [Output Only] Status of the operation. Can be one of the
	// following: ! "PENDING", "RUNNING", or "DONE".
	Status string `json:"status,omitempty"`

	// StatusMessage: ! [Output Only] An optional textual description of the
	// current status of ! the operation.
	StatusMessage string `json:"statusMessage,omitempty"`

	// TargetId: ! [Output Only] Unique target id which identifies a
	// particular ! incarnation of the target.
	TargetId uint64 `json:"targetId,omitempty,string"`

	// TargetLink: ! [Output Only] URL of the resource the operation is
	// mutating.
	TargetLink string `json:"targetLink,omitempty"`

	// User: ! [Output Only] User who requested the operation, for example !
	// "user@example.com"
	User string `json:"user,omitempty"`

	// Warnings: ! [Output Only] If warning messages generated during
	// processing of this ! operation, this field will be populated.
	Warnings []*OperationWarnings `json:"warnings,omitempty"`
}

// OperationError: ! [Output Only] If errors occurred during processing
// of this operation, ! this field will be populated.
type OperationError struct {
	// Errors: ! The array of errors encountered while processing this
	// operation.
	Errors []*OperationErrorErrors `json:"errors,omitempty"`
}

type OperationErrorErrors struct {
	// Code: ! The error type identifier for this error.
	Code string `json:"code,omitempty"`

	// Location: ! Indicates the field in the request which caused the
	// error. ! This property is optional.
	Location string `json:"location,omitempty"`

	// Message: ! An optional, human-readable error message.
	Message string `json:"message,omitempty"`
}

type OperationWarnings struct {
	// Code: ! The warning type identifier for this warning.
	Code interface{} `json:"code,omitempty"`

	// Data: ! Metadata for this warning in 'key: value' format.
	Data []*OperationWarningsData `json:"data,omitempty"`

	// Message: ! Optional human-readable details for this warning.
	Message string `json:"message,omitempty"`
}

type OperationWarningsData struct {
	// Key: ! A key for the warning data.
	Key string `json:"key,omitempty"`

	// Value: ! A warning data value corresponding to the key.
	Value string `json:"value,omitempty"`
}

// OperationsListResponse: ! A response containing a partial list of
// operations and a page token used ! to build the next request if the
// request has been truncated. Next available tag: 4
type OperationsListResponse struct {
	// NextPageToken: ! A token used to continue a truncated list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Operations: ! Operations contained in this list response.
	Operations []*Operation `json:"operations,omitempty"`
}

// Resource: Next available tag: 12
type Resource struct {
	// Errors: ! [Output Only] A list of any errors that occurred during
	// deployment.
	Errors []string `json:"errors,omitempty"`

	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// Intent: ! [Output Only] The intended state of the resource.
	Intent string `json:"intent,omitempty"`

	// Manifest: ! [Output Only] URL of the manifest representing the
	// current configuration ! of this resource.
	Manifest string `json:"manifest,omitempty"`

	// Name: ! [Output Only] The name of the resource as it appears in the
	// YAML config.
	Name string `json:"name,omitempty"`

	// State: ! [Output Only] The state of the resource.
	State string `json:"state,omitempty"`

	// Type: ! [Output Only] The type of the resource, for example !
	// ?compute.v1.instance?, or ?replicaPools.v1beta2.instanceGroupManager?
	Type string `json:"type,omitempty"`

	// Url: ! [Output Only] The URL of the actual resource.
	Url string `json:"url,omitempty"`
}

// ResourcesListResponse: ! A response containing a partial list of
// resources and a page token used ! to build the next request if the
// request has been truncated. Next available tag: 4
type ResourcesListResponse struct {
	// NextPageToken: ! A token used to continue a truncated list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Resources: ! Resources contained in this list response.
	Resources []*Resource `json:"resources,omitempty"`
}

// Type: ! A type supported by Deployment Manager. Next available tag: 4
type Type struct {
	// Name: ! Name of the type.
	Name string `json:"name,omitempty"`
}

// TypesListResponse: ! A response that returns all Types supported by
// Deployment Manager Next available tag: 3
type TypesListResponse struct {
	// Types: ! Types supported by Deployment Manager
	Types []*Type `json:"types,omitempty"`
}

// method id "deploymentmanager.deployments.delete":

type DeploymentsDeleteCall struct {
	s          *Service
	project    string
	deployment string
	opt_       map[string]interface{}
}

// Delete: ! Deletes a deployment and all of the resources in the
// deployment.
func (r *DeploymentsService) Delete(project string, deployment string) *DeploymentsDeleteCall {
	c := &DeploymentsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.deployment = deployment
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DeploymentsDeleteCall) Fields(s ...googleapi.Field) *DeploymentsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *DeploymentsDeleteCall) Do() (*Operation, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/deployments/{deployment}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":    c.project,
		"deployment": c.deployment,
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
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "! Deletes a deployment and all of the resources in the deployment.",
	//   "httpMethod": "DELETE",
	//   "id": "deploymentmanager.deployments.delete",
	//   "parameterOrder": [
	//     "project",
	//     "deployment"
	//   ],
	//   "parameters": {
	//     "deployment": {
	//       "description": "! The name of the deployment for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "! The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments/{deployment}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman"
	//   ]
	// }

}

// method id "deploymentmanager.deployments.get":

type DeploymentsGetCall struct {
	s          *Service
	project    string
	deployment string
	opt_       map[string]interface{}
}

// Get: ! Gets information about a specific deployment.
func (r *DeploymentsService) Get(project string, deployment string) *DeploymentsGetCall {
	c := &DeploymentsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.deployment = deployment
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DeploymentsGetCall) Fields(s ...googleapi.Field) *DeploymentsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *DeploymentsGetCall) Do() (*Deployment, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/deployments/{deployment}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":    c.project,
		"deployment": c.deployment,
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
	var ret *Deployment
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "! Gets information about a specific deployment.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.deployments.get",
	//   "parameterOrder": [
	//     "project",
	//     "deployment"
	//   ],
	//   "parameters": {
	//     "deployment": {
	//       "description": "! The name of the deployment for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "! The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments/{deployment}",
	//   "response": {
	//     "$ref": "Deployment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// method id "deploymentmanager.deployments.insert":

type DeploymentsInsertCall struct {
	s          *Service
	project    string
	deployment *Deployment
	opt_       map[string]interface{}
}

// Insert: ! Creates a deployment and all of the resources described by
// the ! deployment manifest.
func (r *DeploymentsService) Insert(project string, deployment *Deployment) *DeploymentsInsertCall {
	c := &DeploymentsInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.deployment = deployment
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DeploymentsInsertCall) Fields(s ...googleapi.Field) *DeploymentsInsertCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *DeploymentsInsertCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.deployment)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/deployments")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
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
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "! Creates a deployment and all of the resources described by the ! deployment manifest.",
	//   "httpMethod": "POST",
	//   "id": "deploymentmanager.deployments.insert",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "! The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments",
	//   "request": {
	//     "$ref": "Deployment"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman"
	//   ]
	// }

}

// method id "deploymentmanager.deployments.list":

type DeploymentsListCall struct {
	s       *Service
	project string
	opt_    map[string]interface{}
}

// List: ! Lists all deployments for a given project.
func (r *DeploymentsService) List(project string) *DeploymentsListCall {
	c := &DeploymentsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	return c
}

// MaxResults sets the optional parameter "maxResults": ! Maximum count
// of results to be returned. ! Acceptable values are 0 to 100,
// inclusive. (Default: 50)
func (c *DeploymentsListCall) MaxResults(maxResults int64) *DeploymentsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": ! Specifies a
// nextPageToken returned by a previous list request. This ! token can
// be used to request the next page of results from a previous ! list
// request.
func (c *DeploymentsListCall) PageToken(pageToken string) *DeploymentsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DeploymentsListCall) Fields(s ...googleapi.Field) *DeploymentsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *DeploymentsListCall) Do() (*DeploymentsListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/deployments")
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
	var ret *DeploymentsListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "! Lists all deployments for a given project.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.deployments.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "default": "50",
	//       "description": "! Maximum count of results to be returned. ! Acceptable values are 0 to 100, inclusive. (Default: 50)",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "100",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "! Specifies a nextPageToken returned by a previous list request. This ! token can be used to request the next page of results from a previous ! list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "! The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments",
	//   "response": {
	//     "$ref": "DeploymentsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// method id "deploymentmanager.manifests.get":

type ManifestsGetCall struct {
	s          *Service
	project    string
	deployment string
	manifest   string
	opt_       map[string]interface{}
}

// Get: ! Gets information about a specific manifest.
func (r *ManifestsService) Get(project string, deployment string, manifest string) *ManifestsGetCall {
	c := &ManifestsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.deployment = deployment
	c.manifest = manifest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ManifestsGetCall) Fields(s ...googleapi.Field) *ManifestsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ManifestsGetCall) Do() (*Manifest, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/deployments/{deployment}/manifests/{manifest}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":    c.project,
		"deployment": c.deployment,
		"manifest":   c.manifest,
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
	var ret *Manifest
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "! Gets information about a specific manifest.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.manifests.get",
	//   "parameterOrder": [
	//     "project",
	//     "deployment",
	//     "manifest"
	//   ],
	//   "parameters": {
	//     "deployment": {
	//       "description": "! The name of the deployment for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "manifest": {
	//       "description": "! The name of the manifest for this request.",
	//       "location": "path",
	//       "pattern": "[-a-z0-9]{1,61}",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "! The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments/{deployment}/manifests/{manifest}",
	//   "response": {
	//     "$ref": "Manifest"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// method id "deploymentmanager.manifests.list":

type ManifestsListCall struct {
	s          *Service
	project    string
	deployment string
	opt_       map[string]interface{}
}

// List: ! Lists all manifests for a given deployment.
func (r *ManifestsService) List(project string, deployment string) *ManifestsListCall {
	c := &ManifestsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.deployment = deployment
	return c
}

// MaxResults sets the optional parameter "maxResults": ! Maximum count
// of results to be returned. ! Acceptable values are 0 to 100,
// inclusive. (Default: 50)
func (c *ManifestsListCall) MaxResults(maxResults int64) *ManifestsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": ! Specifies a
// nextPageToken returned by a previous list request. This ! token can
// be used to request the next page of results from a previous ! list
// request.
func (c *ManifestsListCall) PageToken(pageToken string) *ManifestsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ManifestsListCall) Fields(s ...googleapi.Field) *ManifestsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ManifestsListCall) Do() (*ManifestsListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/deployments/{deployment}/manifests")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":    c.project,
		"deployment": c.deployment,
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
	var ret *ManifestsListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "! Lists all manifests for a given deployment.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.manifests.list",
	//   "parameterOrder": [
	//     "project",
	//     "deployment"
	//   ],
	//   "parameters": {
	//     "deployment": {
	//       "description": "! The name of the deployment for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "50",
	//       "description": "! Maximum count of results to be returned. ! Acceptable values are 0 to 100, inclusive. (Default: 50)",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "100",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "! Specifies a nextPageToken returned by a previous list request. This ! token can be used to request the next page of results from a previous ! list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "! The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments/{deployment}/manifests",
	//   "response": {
	//     "$ref": "ManifestsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// method id "deploymentmanager.operations.get":

type OperationsGetCall struct {
	s         *Service
	project   string
	operation string
	opt_      map[string]interface{}
}

// Get: ! Gets information about a specific Operation.
func (r *OperationsService) Get(project string, operation string) *OperationsGetCall {
	c := &OperationsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.operation = operation
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OperationsGetCall) Fields(s ...googleapi.Field) *OperationsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *OperationsGetCall) Do() (*Operation, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/operations/{operation}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":   c.project,
		"operation": c.operation,
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
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "! Gets information about a specific Operation.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.operations.get",
	//   "parameterOrder": [
	//     "project",
	//     "operation"
	//   ],
	//   "parameters": {
	//     "operation": {
	//       "description": "! The name of the operation for this request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "! The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/operations/{operation}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// method id "deploymentmanager.operations.list":

type OperationsListCall struct {
	s       *Service
	project string
	opt_    map[string]interface{}
}

// List: ! Lists all Operations for a project.
func (r *OperationsService) List(project string) *OperationsListCall {
	c := &OperationsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	return c
}

// MaxResults sets the optional parameter "maxResults": ! Maximum count
// of results to be returned. ! Acceptable values are 0 to 100,
// inclusive. (Default: 50)
func (c *OperationsListCall) MaxResults(maxResults int64) *OperationsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": ! Specifies a
// nextPageToken returned by a previous list request. This ! token can
// be used to request the next page of results from a previous ! list
// request.
func (c *OperationsListCall) PageToken(pageToken string) *OperationsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OperationsListCall) Fields(s ...googleapi.Field) *OperationsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *OperationsListCall) Do() (*OperationsListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/operations")
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
	var ret *OperationsListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "! Lists all Operations for a project.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.operations.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "default": "50",
	//       "description": "! Maximum count of results to be returned. ! Acceptable values are 0 to 100, inclusive. (Default: 50)",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "100",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "! Specifies a nextPageToken returned by a previous list request. This ! token can be used to request the next page of results from a previous ! list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "! The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/operations",
	//   "response": {
	//     "$ref": "OperationsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// method id "deploymentmanager.resources.get":

type ResourcesGetCall struct {
	s          *Service
	project    string
	deployment string
	resource   string
	opt_       map[string]interface{}
}

// Get: ! Gets information about a single resource.
func (r *ResourcesService) Get(project string, deployment string, resource string) *ResourcesGetCall {
	c := &ResourcesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.deployment = deployment
	c.resource = resource
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ResourcesGetCall) Fields(s ...googleapi.Field) *ResourcesGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ResourcesGetCall) Do() (*Resource, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/deployments/{deployment}/resources/{resource}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":    c.project,
		"deployment": c.deployment,
		"resource":   c.resource,
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
	var ret *Resource
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "! Gets information about a single resource.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.resources.get",
	//   "parameterOrder": [
	//     "project",
	//     "deployment",
	//     "resource"
	//   ],
	//   "parameters": {
	//     "deployment": {
	//       "description": "! The name of the deployment for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "! The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "resource": {
	//       "description": "! The name of the resource for this request.",
	//       "location": "path",
	//       "pattern": "[-a-z0-9]{1,61}",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments/{deployment}/resources/{resource}",
	//   "response": {
	//     "$ref": "Resource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// method id "deploymentmanager.resources.list":

type ResourcesListCall struct {
	s          *Service
	project    string
	deployment string
	opt_       map[string]interface{}
}

// List: ! Lists all resources in a given deployment.
func (r *ResourcesService) List(project string, deployment string) *ResourcesListCall {
	c := &ResourcesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.deployment = deployment
	return c
}

// MaxResults sets the optional parameter "maxResults": ! Maximum count
// of results to be returned. ! Acceptable values are 0 to 100,
// inclusive. (Default: 50)
func (c *ResourcesListCall) MaxResults(maxResults int64) *ResourcesListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": ! Specifies a
// nextPageToken returned by a previous list request. This ! token can
// be used to request the next page of results from a previous ! list
// request.
func (c *ResourcesListCall) PageToken(pageToken string) *ResourcesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ResourcesListCall) Fields(s ...googleapi.Field) *ResourcesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ResourcesListCall) Do() (*ResourcesListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/deployments/{deployment}/resources")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":    c.project,
		"deployment": c.deployment,
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
	var ret *ResourcesListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "! Lists all resources in a given deployment.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.resources.list",
	//   "parameterOrder": [
	//     "project",
	//     "deployment"
	//   ],
	//   "parameters": {
	//     "deployment": {
	//       "description": "! The name of the deployment for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "50",
	//       "description": "! Maximum count of results to be returned. ! Acceptable values are 0 to 100, inclusive. (Default: 50)",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "100",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "! Specifies a nextPageToken returned by a previous list request. This ! token can be used to request the next page of results from a previous ! list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "! The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments/{deployment}/resources",
	//   "response": {
	//     "$ref": "ResourcesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// method id "deploymentmanager.types.list":

type TypesListCall struct {
	s       *Service
	project string
	opt_    map[string]interface{}
}

// List: ! Lists all Types for Deployment Manager.
func (r *TypesService) List(project string) *TypesListCall {
	c := &TypesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	return c
}

// MaxResults sets the optional parameter "maxResults": ! Maximum count
// of results to be returned. ! Acceptable values are 0 to 100,
// inclusive. (Default: 50)
func (c *TypesListCall) MaxResults(maxResults int64) *TypesListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": ! Specifies a
// nextPageToken returned by a previous list request. This ! token can
// be used to request the next page of results from a previous ! list
// request.
func (c *TypesListCall) PageToken(pageToken string) *TypesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TypesListCall) Fields(s ...googleapi.Field) *TypesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *TypesListCall) Do() (*TypesListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/types")
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
	var ret *TypesListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "! Lists all Types for Deployment Manager.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.types.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "default": "50",
	//       "description": "! Maximum count of results to be returned. ! Acceptable values are 0 to 100, inclusive. (Default: 50)",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "100",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "! Specifies a nextPageToken returned by a previous list request. This ! token can be used to request the next page of results from a previous ! list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "! The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/types",
	//   "response": {
	//     "$ref": "TypesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}
