// Package logging provides access to the Google Cloud Logging API.
//
// See https://cloud.google.com/logging/docs/
//
// Usage example:
//
//   import "google.golang.org/api/logging/v1beta3"
//   ...
//   loggingService, err := logging.New(oauthHttpClient)
package logging

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

const apiId = "logging:v1beta3"
const apiName = "logging"
const apiVersion = "v1beta3"
const basePath = "https://logging.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// Administrate log data for your projects
	LoggingAdminScope = "https://www.googleapis.com/auth/logging.admin"

	// View log data for your projects
	LoggingReadScope = "https://www.googleapis.com/auth/logging.read"

	// Submit log data for your projects
	LoggingWriteScope = "https://www.googleapis.com/auth/logging.write"
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
	rs.LogServices = NewProjectsLogServicesService(s)
	rs.Logs = NewProjectsLogsService(s)
	rs.Sinks = NewProjectsSinksService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	LogServices *ProjectsLogServicesService

	Logs *ProjectsLogsService

	Sinks *ProjectsSinksService
}

func NewProjectsLogServicesService(s *Service) *ProjectsLogServicesService {
	rs := &ProjectsLogServicesService{s: s}
	rs.Indexes = NewProjectsLogServicesIndexesService(s)
	rs.Sinks = NewProjectsLogServicesSinksService(s)
	return rs
}

type ProjectsLogServicesService struct {
	s *Service

	Indexes *ProjectsLogServicesIndexesService

	Sinks *ProjectsLogServicesSinksService
}

func NewProjectsLogServicesIndexesService(s *Service) *ProjectsLogServicesIndexesService {
	rs := &ProjectsLogServicesIndexesService{s: s}
	return rs
}

type ProjectsLogServicesIndexesService struct {
	s *Service
}

func NewProjectsLogServicesSinksService(s *Service) *ProjectsLogServicesSinksService {
	rs := &ProjectsLogServicesSinksService{s: s}
	return rs
}

type ProjectsLogServicesSinksService struct {
	s *Service
}

func NewProjectsLogsService(s *Service) *ProjectsLogsService {
	rs := &ProjectsLogsService{s: s}
	rs.Entries = NewProjectsLogsEntriesService(s)
	rs.Sinks = NewProjectsLogsSinksService(s)
	return rs
}

type ProjectsLogsService struct {
	s *Service

	Entries *ProjectsLogsEntriesService

	Sinks *ProjectsLogsSinksService
}

func NewProjectsLogsEntriesService(s *Service) *ProjectsLogsEntriesService {
	rs := &ProjectsLogsEntriesService{s: s}
	return rs
}

type ProjectsLogsEntriesService struct {
	s *Service
}

func NewProjectsLogsSinksService(s *Service) *ProjectsLogsSinksService {
	rs := &ProjectsLogsSinksService{s: s}
	return rs
}

type ProjectsLogsSinksService struct {
	s *Service
}

func NewProjectsSinksService(s *Service) *ProjectsSinksService {
	rs := &ProjectsSinksService{s: s}
	return rs
}

type ProjectsSinksService struct {
	s *Service
}

// AuditData: BigQuery request and response messages for audit log.
type AuditData struct {
	// DatasetInsertRequest: Dataset insert request.
	DatasetInsertRequest *DatasetInsertRequest `json:"datasetInsertRequest,omitempty"`

	// DatasetInsertResponse: Dataset insert response.
	DatasetInsertResponse *DatasetInsertResponse `json:"datasetInsertResponse,omitempty"`

	// DatasetListRequest: Dataset list request.
	DatasetListRequest *DatasetListRequest `json:"datasetListRequest,omitempty"`

	// DatasetUpdateRequest: Dataset update request.
	DatasetUpdateRequest *DatasetUpdateRequest `json:"datasetUpdateRequest,omitempty"`

	// DatasetUpdateResponse: Dataset update response.
	DatasetUpdateResponse *DatasetUpdateResponse `json:"datasetUpdateResponse,omitempty"`

	// JobGetQueryResultsRequest: Job get query results request.
	JobGetQueryResultsRequest *JobGetQueryResultsRequest `json:"jobGetQueryResultsRequest,omitempty"`

	// JobGetQueryResultsResponse: Job get query results response.
	JobGetQueryResultsResponse *JobGetQueryResultsResponse `json:"jobGetQueryResultsResponse,omitempty"`

	// JobInsertRequest: Job insert request.
	JobInsertRequest *JobInsertRequest `json:"jobInsertRequest,omitempty"`

	// JobQueryDoneResponse: Job query-done response. Use this information
	// for usage analysis.
	JobQueryDoneResponse *JobQueryDoneResponse `json:"jobQueryDoneResponse,omitempty"`

	// JobQueryRequest: Job query request.
	JobQueryRequest *JobQueryRequest `json:"jobQueryRequest,omitempty"`

	// JobQueryResponse: Job query response.
	JobQueryResponse *JobQueryResponse `json:"jobQueryResponse,omitempty"`

	// TableDataListRequest: Table data-list request.
	TableDataListRequest *TableDataListRequest `json:"tableDataListRequest,omitempty"`

	// TableInsertRequest: Table insert request.
	TableInsertRequest *TableInsertRequest `json:"tableInsertRequest,omitempty"`

	// TableInsertResponse: Table insert response.
	TableInsertResponse *TableInsertResponse `json:"tableInsertResponse,omitempty"`

	// TableUpdateRequest: Table update request.
	TableUpdateRequest *TableUpdateRequest `json:"tableUpdateRequest,omitempty"`

	// TableUpdateResponse: Table update response.
	TableUpdateResponse *TableUpdateResponse `json:"tableUpdateResponse,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "DatasetInsertRequest") to unconditionally include in API requests.
	// By default, fields with empty values are ommitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string
}

func (s AuditData) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// AuditLog: Common audit log format for Google Cloud Platform API
// calls.
type AuditLog struct {
	// AuthenticationInfo: Authentication information about the call.
	AuthenticationInfo *AuthenticationInfo `json:"authenticationInfo,omitempty"`

	// AuthorizationInfo: Authorization information about the call. If there
	// are multiple resources or permissions involved in authorizing the
	// request, there will be one AuthorizationInfo element for each
	// {resource, permission} tuple.
	AuthorizationInfo []*AuthorizationInfo `json:"authorizationInfo,omitempty"`

	// BigqueryData: Service-specific data for BigQuery.
	BigqueryData *AuditData `json:"bigqueryData,omitempty"`

	// MethodName: Name of the service method or operation. Defined by the
	// service. For API call events, should match the name of the API
	// method. For example, `google.datastore.v1.Datastore.RunQuery`
	// `google.logging.v1.LoggingService.DeleteLog`
	MethodName string `json:"methodName,omitempty"`

	// NumResponseItems: If applicable, the number of items returned from a
	// List or Query API method.
	NumResponseItems int64 `json:"numResponseItems,omitempty,string"`

	// RequestMetadata: Metadata about the request.
	RequestMetadata *RequestMetadata `json:"requestMetadata,omitempty"`

	// ResourceName: Resource name of the resource or collection that is the
	// target of this request, as a scheme-less URI, not including the API
	// service name. For example: shelves/shelf_id/books
	// shelves/shelf_id/books/book_id
	ResourceName string `json:"resourceName,omitempty"`

	// ServiceData: Service specific data about the request, response, and
	// other event data. This should include all request parameters or
	// response elements, except for parameters that are large or
	// privacy-sensitive. It should never contain user-generated data (such
	// as file contents).
	ServiceData AuditLogServiceData `json:"serviceData,omitempty"`

	// ServiceName: Name of the API service for the request. e.g.,
	// datastore.googleapis.com
	ServiceName string `json:"serviceName,omitempty"`

	// Status: The status of the overall API call.
	Status *Status `json:"status,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AuthenticationInfo")
	// to unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s AuditLog) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

type AuditLogServiceData interface{}

// AuthenticationInfo: Authentication information for the call.
type AuthenticationInfo struct {
	// PrincipalEmail: Email address of the authenticated user making the
	// request
	PrincipalEmail string `json:"principalEmail,omitempty"`

	// ForceSendFields is a list of field names (e.g. "PrincipalEmail") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s AuthenticationInfo) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// AuthorizationInfo: Authorization information for the call.
type AuthorizationInfo struct {
	// Granted: Whether or not authorization for this resource and
	// permission was granted.
	Granted bool `json:"granted,omitempty"`

	// Permission: The required IAM permission.
	Permission string `json:"permission,omitempty"`

	// Resource: The resource being accessed, as a REST-style string. For
	// example:
	// `bigquery.googlapis.com/projects/PROJECTID/datasets/DATASETID`
	Resource string `json:"resource,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Granted") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s AuthorizationInfo) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// BigQueryAcl: Access control list.
type BigQueryAcl struct {
	// Entries: Access control entry list.
	Entries []*Entry `json:"entries,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Entries") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s BigQueryAcl) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// Dataset: BigQuery dataset information.
type Dataset struct {
	// Acl: Access contol list for this dataset.
	Acl *BigQueryAcl `json:"acl,omitempty"`

	// CreateTime: The creation time for this dataset.
	CreateTime string `json:"createTime,omitempty"`

	// DatasetName: The name of this dataset.
	DatasetName *DatasetName `json:"datasetName,omitempty"`

	// DefaultTableExpireDuration: The number of milliseconds which should
	// be added to the creation time to determine the expiration time for
	// newly created tables. If this value is null then no expiration time
	// will be set for new tables.
	DefaultTableExpireDuration string `json:"defaultTableExpireDuration,omitempty"`

	// Info: User-modifiable metadata for this dataset.
	Info *DatasetInfo `json:"info,omitempty"`

	// UpdateTime: The last modified time for this dataset.
	UpdateTime string `json:"updateTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Acl") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s Dataset) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// DatasetInfo: User-provided metadata for a dataset, primarily for
// display in the UI.
type DatasetInfo struct {
	// Description: The description of a dataset. This can be several
	// sentences or paragraphs describing the dataset contents in detail.
	Description string `json:"description,omitempty"`

	// FriendlyName: The human-readable name of a dataset. This should be a
	// short phrase identifying the dataset (e.g., "Analytics Data 2011").
	FriendlyName string `json:"friendlyName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s DatasetInfo) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// DatasetInsertRequest: Dataset insert request.
type DatasetInsertRequest struct {
	// Resource: Dataset insert payload.
	Resource *Dataset `json:"resource,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Resource") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s DatasetInsertRequest) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// DatasetInsertResponse: Dataset insert response.
type DatasetInsertResponse struct {
	// Resource: Final state of inserted dataset.
	Resource *Dataset `json:"resource,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Resource") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s DatasetInsertResponse) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// DatasetListRequest: Dataset list request.
type DatasetListRequest struct {
	// ListAll: Whether to list all datasets, including hidden ones.
	ListAll bool `json:"listAll,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ListAll") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s DatasetListRequest) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// DatasetName: Fully qualified name for a dataset.
type DatasetName struct {
	// DatasetId: The ID of the dataset (scoped to the project above).
	DatasetId string `json:"datasetId,omitempty"`

	// ProjectId: A string containing the id of this project. The id may be
	// the alphanumeric project ID, or the project number.
	ProjectId string `json:"projectId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DatasetId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s DatasetName) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// DatasetUpdateRequest: Dataset update request.
type DatasetUpdateRequest struct {
	// Resource: Dataset update payload.
	Resource *Dataset `json:"resource,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Resource") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s DatasetUpdateRequest) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// DatasetUpdateResponse: Dataset update response.
type DatasetUpdateResponse struct {
	// Resource: Final state of updated dataset.
	Resource *Dataset `json:"resource,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Resource") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s DatasetUpdateResponse) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated empty messages in your APIs. A typical example is to use
// it as the request or the response type of an API method. For
// instance: service Foo { rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty); } The JSON representation for `Empty` is
// empty JSON object `{}`.
type Empty struct {
}

// Entry: Access control entry.
type Entry struct {
	// Domain: Grants access to all members of a domain.
	Domain string `json:"domain,omitempty"`

	// GroupEmail: Grants access to a group, by e-mail.
	GroupEmail string `json:"groupEmail,omitempty"`

	// Role: Granted role. Valid roles are READER, WRITER, OWNER.
	Role string `json:"role,omitempty"`

	// SpecialGroup: Grants access to special groups. Valid groups are
	// PROJECT_OWNERS, PROJECT_READERS, PROJECT_WRITERS and
	// ALL_AUTHENTICATED_USERS.
	SpecialGroup string `json:"specialGroup,omitempty"`

	// UserEmail: Grants access to a user, by e-mail.
	UserEmail string `json:"userEmail,omitempty"`

	// ViewName: Grants access to a BigQuery View.
	ViewName *TableName `json:"viewName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Domain") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s Entry) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// Extract: Describes an extract job, which exports data to an external
// source via the export pipeline.
type Extract struct {
	// DestinationUris: URI or URIs where extracted data should be written.
	// Currently, only Bigstore URIs are supported (e.g.,
	// "gs://bucket/object"). If more than one URI given, output will be
	// divided into 'partitions' of data, with each partition containing one
	// or more files. If more than one URI is given, each URI must contain
	// exactly one '*' which will be replaced with the file number (within
	// the partition) padded out to 9 digits.
	DestinationUris []string `json:"destinationUris,omitempty"`

	// SourceTable: Source table.
	SourceTable *TableName `json:"sourceTable,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DestinationUris") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s Extract) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// FieldSchema: BigQuery field schema.
type FieldSchema struct {
	// Mode: Column mode
	Mode string `json:"mode,omitempty"`

	// Name: Column name Matches: [A-Za-z_][A-Za-z_0-9]{0,127}
	Name string `json:"name,omitempty"`

	// Schema: Present iff type == RECORD.
	Schema *TableSchema `json:"schema,omitempty"`

	// Type: Column type
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Mode") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s FieldSchema) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// HttpRequest: A common proto for logging HTTP requests.
type HttpRequest struct {
	// Referer: Referer (a.k.a. referrer) URL of request, as defined in
	// http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html.
	Referer string `json:"referer,omitempty"`

	// RemoteIp: IP address of the client who issues the HTTP request. Could
	// be either IPv4 or IPv6.
	RemoteIp string `json:"remoteIp,omitempty"`

	// RequestMethod: Request method, such as `GET`, `HEAD`, `PUT` or
	// `POST`.
	RequestMethod string `json:"requestMethod,omitempty"`

	// RequestSize: Size of the HTTP request message in bytes, including
	// request headers and the request body.
	RequestSize int64 `json:"requestSize,omitempty,string"`

	// RequestUrl: Contains the scheme (http|https), the host name, the path
	// and the query portion of the URL that was requested.
	RequestUrl string `json:"requestUrl,omitempty"`

	// ResponseSize: Size of the HTTP response message in bytes sent back to
	// the client, including response headers and response body.
	ResponseSize int64 `json:"responseSize,omitempty,string"`

	// Status: A response code indicates the status of response, e.g., 200.
	Status int64 `json:"status,omitempty"`

	// UserAgent: User agent sent by the client, e.g., "Mozilla/4.0
	// (compatible; MSIE 6.0; Windows 98; Q312461; .NET CLR 1.0.3705)".
	UserAgent string `json:"userAgent,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Referer") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s HttpRequest) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// Job: Combines all of the information about a job.
type Job struct {
	// JobConfiguration: Job configuration.
	JobConfiguration *JobConfiguration `json:"jobConfiguration,omitempty"`

	// JobName: Job name.
	JobName *JobName `json:"jobName,omitempty"`

	// JobStatistics: Job statistics.
	JobStatistics *JobStatistics `json:"jobStatistics,omitempty"`

	// JobStatus: Job status.
	JobStatus *JobStatus `json:"jobStatus,omitempty"`

	// ForceSendFields is a list of field names (e.g. "JobConfiguration") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s Job) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// JobConfiguration: Job configuration information.
type JobConfiguration struct {
	// DryRun: If set, don't actually run the job. Just check that it would
	// run.
	DryRun bool `json:"dryRun,omitempty"`

	// Extract: Extract job information.
	Extract *Extract `json:"extract,omitempty"`

	// Load: Load job information.
	Load *Load `json:"load,omitempty"`

	// Query: Query job information.
	Query *Query `json:"query,omitempty"`

	// TableCopy: TableCopy job information.
	TableCopy *TableCopy `json:"tableCopy,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DryRun") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s JobConfiguration) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// JobGetQueryResultsRequest: Job get-query-results request.
type JobGetQueryResultsRequest struct {
	// MaxResults: Maximum number of results to return.
	MaxResults int64 `json:"maxResults,omitempty"`

	// StartRow: Row number to start returning results from.
	StartRow uint64 `json:"startRow,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "MaxResults") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s JobGetQueryResultsRequest) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// JobGetQueryResultsResponse: Job get-query-results response.
type JobGetQueryResultsResponse struct {
	// Job: Job that was created to run the query. Includes job state, job
	// statistics, and job errors (if any). To determine whether the job has
	// completed, check that job.status.state == DONE. If
	// job.status.error_result is set, then the job failed. If the job has
	// not yet completed, call GetQueryResults again.
	Job *Job `json:"job,omitempty"`

	// TotalResults: Total number of results in query results.
	TotalResults uint64 `json:"totalResults,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "Job") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s JobGetQueryResultsResponse) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// JobInsertRequest: Job insert request.
type JobInsertRequest struct {
	// Resource: Job insert payload.
	Resource *Job `json:"resource,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Resource") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s JobInsertRequest) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// JobName: Fully-qualified name for a job.
type JobName struct {
	// JobId: The ID of the job (scoped to the project above).
	JobId string `json:"jobId,omitempty"`

	// ProjectId: A string containing the id of this project.
	ProjectId string `json:"projectId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "JobId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s JobName) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// JobQueryDoneResponse: Job get query-done response.
type JobQueryDoneResponse struct {
	// Job: Usage information about completed job.
	Job *Job `json:"job,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Job") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s JobQueryDoneResponse) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// JobQueryRequest: Job query request.
type JobQueryRequest struct {
	// DefaultDataset: Default dataset to use when tables in a query do not
	// have a dataset specified.
	DefaultDataset *DatasetName `json:"defaultDataset,omitempty"`

	// DryRun: If set, don't actually run the query.
	DryRun bool `json:"dryRun,omitempty"`

	// MaxResults: Maximum number of results to return.
	MaxResults int64 `json:"maxResults,omitempty"`

	// ProjectId: Project that the query should be charged to.
	ProjectId string `json:"projectId,omitempty"`

	// Query: The query to execute.
	Query string `json:"query,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DefaultDataset") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s JobQueryRequest) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// JobQueryResponse: Job query response.
type JobQueryResponse struct {
	// Job: Information about queried job.
	Job *Job `json:"job,omitempty"`

	// TotalResults: The total number of rows in the complete query result
	// set.
	TotalResults uint64 `json:"totalResults,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "Job") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s JobQueryResponse) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// JobStatistics: Job statistics that may change after a job starts.
type JobStatistics struct {
	// CreateTime: Time when the job was created (in milliseconds since the
	// POSIX epoch).
	CreateTime string `json:"createTime,omitempty"`

	// EndTime: Time when the job ended.
	EndTime string `json:"endTime,omitempty"`

	// StartTime: Time when the job started.
	StartTime string `json:"startTime,omitempty"`

	// TotalProcessedBytes: Total bytes processed for a job.
	TotalProcessedBytes int64 `json:"totalProcessedBytes,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "CreateTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s JobStatistics) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// JobStatus: Running state of a job (whether it is running, failed,
// etc).
type JobStatus struct {
	// Error: If the job did not complete successfully, this will contain an
	// error.
	Error *Status `json:"error,omitempty"`

	// State: State of a job: PENDING, RUNNING, DONE. Includes no
	// information about whether the job was successful or not.
	State string `json:"state,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Error") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s JobStatus) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// ListLogServiceIndexesResponse: Result returned from
// ListLogServiceIndexesRequest.
type ListLogServiceIndexesResponse struct {
	// NextPageToken: If there are more results, then `nextPageToken` is
	// returned in the response. To get the next batch of indexes, use the
	// value of `nextPageToken` as `pageToken` in the next call of
	// `ListLogServiceIndexes`. If `nextPageToken` is empty, then there are
	// no more results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServiceIndexPrefixes: A list of log service index prefixes.
	ServiceIndexPrefixes []string `json:"serviceIndexPrefixes,omitempty"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s ListLogServiceIndexesResponse) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// ListLogServiceSinksResponse: Result returned from
// `ListLogServiceSinks`.
type ListLogServiceSinksResponse struct {
	// Sinks: The requested log service sinks. If any of the returned
	// `LogSink` objects have an empty `destination` field, then call
	// `logServices.sinks.get` to retrieve the complete `LogSink` object.
	Sinks []*LogSink `json:"sinks,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Sinks") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s ListLogServiceSinksResponse) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// ListLogServicesResponse: Result returned from
// `ListLogServicesRequest`.
type ListLogServicesResponse struct {
	// LogServices: A list of log services.
	LogServices []*LogService `json:"logServices,omitempty"`

	// NextPageToken: If there are more results, then `nextPageToken` is
	// returned in the response. To get the next batch of services, use the
	// value of `nextPageToken` as `pageToken` in the next call of
	// `ListLogServices`. If `nextPageToken` is empty, then there are no
	// more results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ForceSendFields is a list of field names (e.g. "LogServices") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s ListLogServicesResponse) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// ListLogSinksResponse: Result returned from `ListLogSinks`.
type ListLogSinksResponse struct {
	// Sinks: The requested log sinks. If any of the returned `LogSink`
	// objects have an empty `destination` field, then call
	// `logServices.sinks.get` to retrieve the complete `LogSink` object.
	Sinks []*LogSink `json:"sinks,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Sinks") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s ListLogSinksResponse) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// ListLogsResponse: Result returned from ListLogs.
type ListLogsResponse struct {
	// Logs: A list of log resources.
	Logs []*Log `json:"logs,omitempty"`

	// NextPageToken: If there are more results, then `nextPageToken` is
	// returned in the response. To get the next batch of logs, use the
	// value of `nextPageToken` as `pageToken` in the next call of
	// `ListLogs`. If `nextPageToken` is empty, then there are no more
	// results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Logs") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s ListLogsResponse) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// ListSinksResponse: Result returned from `ListSinks`.
type ListSinksResponse struct {
	// Sinks: The requested sinks.
	Sinks []*LogSink `json:"sinks,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Sinks") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s ListSinksResponse) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// Load: Describes a load job, which loads data from an external source
// via the import pipeline.
type Load struct {
	// CreateDisposition: Describes when a job should create a table.
	CreateDisposition string `json:"createDisposition,omitempty"`

	// DestinationTable: table where the imported data should be written.
	DestinationTable *TableName `json:"destinationTable,omitempty"`

	// Schema: Schema for the data to be imported.
	Schema *TableSchema `json:"schema,omitempty"`

	// SourceUris: URIs for the data to be imported. Only Bigstore URIs are
	// supported (e.g., "gs://bucket/object").
	SourceUris []string `json:"sourceUris,omitempty"`

	// WriteDisposition: Describes how writes should affect the table
	// associated with the job.
	WriteDisposition string `json:"writeDisposition,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CreateDisposition")
	// to unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s Load) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// Log: A log object.
type Log struct {
	// DisplayName: Name used when displaying the log to the user (for
	// example, in a UI). Example: "activity_log"
	DisplayName string `json:"displayName,omitempty"`

	// Name: REQUIRED: The log's name. Example:
	// "compute.googleapis.com/activity_log".
	Name string `json:"name,omitempty"`

	// PayloadType: Type URL describing the expected payload type for the
	// log.
	PayloadType string `json:"payloadType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DisplayName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s Log) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// LogEntry: An individual entry in a log.
type LogEntry struct {
	// HttpRequest: Information about the HTTP request associated with this
	// log entry, if applicable.
	HttpRequest *HttpRequest `json:"httpRequest,omitempty"`

	// InsertId: A unique ID for the log entry. If you provide this field,
	// the logging service considers other log entries in the same log with
	// the same ID as duplicates which can be removed.
	InsertId string `json:"insertId,omitempty"`

	// Log: The log to which this entry belongs. When a log entry is
	// ingested, the value of this field is set by the logging system.
	Log string `json:"log,omitempty"`

	// Metadata: Information about the log entry.
	Metadata *LogEntryMetadata `json:"metadata,omitempty"`

	// ProtoPayload: The log entry payload, represented as a protocol buffer
	// that is expressed as a JSON object. You can only pass `protoPayload`
	// values that belong to a set of approved types.
	ProtoPayload LogEntryProtoPayload `json:"protoPayload,omitempty"`

	// StructPayload: The log entry payload, represented as a structure that
	// is expressed as a JSON object.
	StructPayload LogEntryStructPayload `json:"structPayload,omitempty"`

	// TextPayload: The log entry payload, represented as a text string.
	TextPayload string `json:"textPayload,omitempty"`

	// ForceSendFields is a list of field names (e.g. "HttpRequest") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s LogEntry) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

type LogEntryProtoPayload interface{}

type LogEntryStructPayload interface{}

// LogEntryMetadata: Additional data that is associated with a log
// entry, set by the service creating the log entry.
type LogEntryMetadata struct {
	// Labels: A set of (key, value) data that provides additional
	// information about the log entry. If the log entry is from one of the
	// Google Cloud Platform sources listed below, the indicated (key,
	// value) information must be provided: Google App Engine, service_name
	// `appengine.googleapis.com`: "appengine.googleapis.com/module_id",
	// "appengine.googleapis.com/version_id",  and one of:
	// "appengine.googleapis.com/replica_index",
	// "appengine.googleapis.com/clone_id",  or else provide the following
	// Compute Engine labels: Google Compute Engine, service_name
	// `compute.googleapis.com`: "compute.googleapis.com/resource_type",
	// "instance" "compute.googleapis.com/resource_id",
	Labels map[string]string `json:"labels,omitempty"`

	// ProjectId: The project ID of the Google Cloud Platform service that
	// created the log entry.
	ProjectId string `json:"projectId,omitempty"`

	// Region: The region name of the Google Cloud Platform service that
	// created the log entry. For example, "us-central1".
	Region string `json:"region,omitempty"`

	// ServiceName: The API name of the Google Cloud Platform service that
	// created the log entry. For example, "compute.googleapis.com".
	ServiceName string `json:"serviceName,omitempty"`

	// Severity: The severity of the log entry.
	//
	// Possible values:
	//   "DEFAULT"
	//   "DEBUG"
	//   "INFO"
	//   "NOTICE"
	//   "WARNING"
	//   "ERROR"
	//   "CRITICAL"
	//   "ALERT"
	//   "EMERGENCY"
	Severity string `json:"severity,omitempty"`

	// Timestamp: The time the event described by the log entry occurred.
	// Timestamps must be later than January 1, 1970.
	Timestamp string `json:"timestamp,omitempty"`

	// UserId: The fully-qualified email address of the authenticated user
	// that performed or requested the action represented by the log entry.
	// If the log entry does not apply to an action taken by an
	// authenticated user, then the field should be empty.
	UserId string `json:"userId,omitempty"`

	// Zone: The zone of the Google Cloud Platform service that created the
	// log entry. For example, "us-central1-a".
	Zone string `json:"zone,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Labels") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s LogEntryMetadata) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// LogError: A problem in a sink or the sink's configuration.
type LogError struct {
	// Resource: The resource associated with the error. It may be different
	// from the sink destination. For example, the sink may point to a
	// BigQuery dataset, but the error may refer to a table resource inside
	// the dataset.
	Resource string `json:"resource,omitempty"`

	// Status: The description of the last error observed.
	Status *Status `json:"status,omitempty"`

	// TimeNanos: The last time the error was observed, in nanoseconds since
	// the Unix epoch.
	TimeNanos int64 `json:"timeNanos,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "Resource") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s LogError) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// LogLine: Application log line emitted while processing a request.
type LogLine struct {
	// LogMessage: App provided log message.
	LogMessage string `json:"logMessage,omitempty"`

	// Severity: Severity of log.
	//
	// Possible values:
	//   "DEFAULT"
	//   "DEBUG"
	//   "INFO"
	//   "NOTICE"
	//   "WARNING"
	//   "ERROR"
	//   "CRITICAL"
	//   "ALERT"
	//   "EMERGENCY"
	Severity string `json:"severity,omitempty"`

	// SourceLocation: Line of code that generated this log message.
	SourceLocation *SourceLocation `json:"sourceLocation,omitempty"`

	// Time: Time when log entry was made. May be inaccurate.
	Time string `json:"time,omitempty"`

	// ForceSendFields is a list of field names (e.g. "LogMessage") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s LogLine) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// LogService: A log service object.
type LogService struct {
	// IndexKeys: Label keys used when labeling log entries for this
	// service. The order of the keys is significant, with higher priority
	// keys coming earlier in the list.
	IndexKeys []string `json:"indexKeys,omitempty"`

	// Name: The service's name.
	Name string `json:"name,omitempty"`

	// ForceSendFields is a list of field names (e.g. "IndexKeys") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s LogService) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// LogSink: An object that describes where a log may be written.
type LogSink struct {
	// Destination: The resource to send log entries to. The supported sink
	// resource types are: + Google Cloud Storage:
	// `storage.googleapis.com/BUCKET` or `BUCKET.storage.googleapis.com/` +
	// Google BigQuery:
	// `bigquery.googleapis.com/projects/PROJECT/datasets/DATASET` Currently
	// the Cloud Logging API supports at most one sink for each resource
	// type per log or log service resource.
	Destination string `json:"destination,omitempty"`

	// Errors: _Output only._ All active errors found for this sink.
	Errors []*LogError `json:"errors,omitempty"`

	// Filter: One Platform filter expression. If provided, only the
	// messages matching the filter will be published.
	Filter string `json:"filter,omitempty"`

	// Name: The name of this sink. This is a client-assigned identifier for
	// the resource. This is ignored by UpdateLogSink and
	// UpdateLogServicesSink.
	Name string `json:"name,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Destination") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s LogSink) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// Money: Represents an amount of money with its currency type.
type Money struct {
	// CurrencyCode: The 3-letter currency code defined in ISO 4217.
	CurrencyCode string `json:"currencyCode,omitempty"`

	// Nanos: Number of nano (10^-9) units of the amount. The value must be
	// between -999,999,999 and +999,999,999 inclusive. If `units` is
	// positive, `nanos` must be positive or zero. If `units` is zero,
	// `nanos` can be positive, zero, or negative. If `units` is negative,
	// `nanos` must be negative or zero. For example $-1.75 is represented
	// as `units`=-1 and `nanos`=-750,000,000.
	Nanos int64 `json:"nanos,omitempty"`

	// Units: The whole units of the amount. For example if `currencyCode`
	// is "USD", then 1 unit is one US dollar.
	Units int64 `json:"units,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "CurrencyCode") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s Money) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// Query: Describes a query job, which executes a SQL-like query.
type Query struct {
	// CreateDisposition: Describe when a job should create a table.
	CreateDisposition string `json:"createDisposition,omitempty"`

	// DefaultDataset: If a table name is specified without a dataset in a
	// query, this dataset will be added to table name.
	DefaultDataset *DatasetName `json:"defaultDataset,omitempty"`

	// DestinationTable: table where results should be written.
	DestinationTable *TableName `json:"destinationTable,omitempty"`

	// Query: SQL query to run.
	Query string `json:"query,omitempty"`

	// TableDefinitions: Additional tables that this query might reference
	// beyond the tables already defined in BigQuery. This is typically used
	// to provide external data references for this query.
	TableDefinitions []*TableDefinition `json:"tableDefinitions,omitempty"`

	// WriteDisposition: Describes how writes should affect the table
	// associated with the job.
	WriteDisposition string `json:"writeDisposition,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CreateDisposition")
	// to unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s Query) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// RequestLog: Complete log information about a single request to an
// application.
type RequestLog struct {
	// AppEngineRelease: App Engine release version string.
	AppEngineRelease string `json:"appEngineRelease,omitempty"`

	// AppId: Identifies the application that handled this request.
	AppId string `json:"appId,omitempty"`

	// Cost: An indication of the relative cost of serving this request.
	Cost float64 `json:"cost,omitempty"`

	// EndTime: Time at which request was known to end processing.
	EndTime string `json:"endTime,omitempty"`

	// Finished: If true, represents a finished request. Otherwise, the
	// request is active.
	Finished bool `json:"finished,omitempty"`

	// Host: The Internet host and port number of the resource being
	// requested.
	Host string `json:"host,omitempty"`

	// HttpVersion: HTTP version of request.
	HttpVersion string `json:"httpVersion,omitempty"`

	// InstanceId: An opaque identifier for the instance that handled the
	// request.
	InstanceId string `json:"instanceId,omitempty"`

	// InstanceIndex: If the instance that processed this request was
	// individually addressable (i.e. belongs to a manually scaled module),
	// this is the index of the instance.
	InstanceIndex int64 `json:"instanceIndex,omitempty"`

	// Ip: Origin IP address.
	Ip string `json:"ip,omitempty"`

	// Latency: Latency of the request.
	Latency string `json:"latency,omitempty"`

	// Line: List of log lines emitted by the application while serving this
	// request, if requested.
	Line []*LogLine `json:"line,omitempty"`

	// MegaCycles: Number of CPU megacycles used to process request.
	MegaCycles int64 `json:"megaCycles,omitempty,string"`

	// Method: Request method, such as `GET`, `HEAD`, `PUT`, `POST`, or
	// `DELETE`.
	Method string `json:"method,omitempty"`

	// ModuleId: Identifies the module of the application that handled this
	// request.
	ModuleId string `json:"moduleId,omitempty"`

	// Nickname: A string that identifies a logged-in user who made this
	// request, or empty if the user is not logged in. Most likely, this is
	// the part of the user's email before the '@' sign. The field value is
	// the same for different requests from the same user, but different
	// users may have a similar name. This information is also available to
	// the application via Users API. This field will be populated starting
	// with App Engine 1.9.21.
	Nickname string `json:"nickname,omitempty"`

	// PendingTime: Time this request spent in the pending request queue, if
	// it was pending at all.
	PendingTime string `json:"pendingTime,omitempty"`

	// Referrer: Referrer URL of request.
	Referrer string `json:"referrer,omitempty"`

	// RequestId: Globally unique identifier for a request, based on request
	// start time. Request IDs for requests which started later will compare
	// greater as binary strings than those for requests which started
	// earlier.
	RequestId string `json:"requestId,omitempty"`

	// Resource: Contains the path and query portion of the URL that was
	// requested. For example, if the URL was
	// "http://example.com/app?name=val", the resource would be
	// "/app?name=val". Any trailing fragment (separated by a '#' character)
	// will not be included.
	Resource string `json:"resource,omitempty"`

	// ResponseSize: Size in bytes sent back to client by request.
	ResponseSize int64 `json:"responseSize,omitempty,string"`

	// SourceReference: Source code for the application that handled this
	// request. There can be more than one source reference per deployed
	// application if source code is distributed among multiple
	// repositories.
	SourceReference []*SourceReference `json:"sourceReference,omitempty"`

	// StartTime: Time at which request was known to have begun processing.
	StartTime string `json:"startTime,omitempty"`

	// Status: Response status of request.
	Status int64 `json:"status,omitempty"`

	// TaskName: Task name of the request (for an offline request).
	TaskName string `json:"taskName,omitempty"`

	// TaskQueueName: Queue name of the request (for an offline request).
	TaskQueueName string `json:"taskQueueName,omitempty"`

	// TraceId: Cloud Trace identifier of the trace for this request.
	TraceId string `json:"traceId,omitempty"`

	// UrlMapEntry: File or class within URL mapping used for request.
	// Useful for tracking down the source code which was responsible for
	// managing request. Especially for multiply mapped handlers.
	UrlMapEntry string `json:"urlMapEntry,omitempty"`

	// UserAgent: User agent used for making request.
	UserAgent string `json:"userAgent,omitempty"`

	// VersionId: Version of the application that handled this request.
	VersionId string `json:"versionId,omitempty"`

	// WasLoadingRequest: Was this request a loading request for this
	// instance?
	WasLoadingRequest bool `json:"wasLoadingRequest,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AppEngineRelease") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s RequestLog) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// RequestMetadata: Metadata about the request.
type RequestMetadata struct {
	// CallerIp: IP address of the caller
	CallerIp string `json:"callerIp,omitempty"`

	// CallerSuppliedUserAgent: User-Agent of the caller. This is not
	// authenticated, so a malicious caller could provide a misleading
	// value. For example: `google-api-python-client/1.4.0` The request was
	// made by the Google API client for Python. `Cloud SDK Command Line
	// Tool apitools-client/1.0 gcloud/0.9.62` The request was made by the
	// Google Cloud SDK CLI (gcloud). `AppEngine-Google;
	// (+http://code.google.com/appengine; appid: s~my-project` The request
	// was made from the `my-project` App Engine app.
	CallerSuppliedUserAgent string `json:"callerSuppliedUserAgent,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CallerIp") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s RequestMetadata) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// SourceLocation: Specifies a location in a source file.
type SourceLocation struct {
	// File: Source file name. May or may not be a fully qualified name,
	// depending on the runtime environment.
	File string `json:"file,omitempty"`

	// FunctionName: Human-readable name of the function or method being
	// invoked, with optional context such as the class or package name, for
	// use in contexts such as the logs viewer where file:line number is
	// less meaningful. This may vary by language, for example: in Java:
	// qual.if.ied.Class.method in Go: dir/package.func in Python: function
	// ...
	FunctionName string `json:"functionName,omitempty"`

	// Line: Line within the source file.
	Line int64 `json:"line,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "File") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s SourceLocation) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// SourceReference: A reference to a particular snapshot of the source
// tree used to build and deploy an application.
type SourceReference struct {
	// Repository: Optional. A URI string identifying the repository.
	// Example: "https://github.com/GoogleCloudPlatform/kubernetes.git"
	Repository string `json:"repository,omitempty"`

	// RevisionId: The canonical (and persistent) identifier of the deployed
	// revision. Example (git): "0035781c50ec7aa23385dc841529ce8a4b70db1b"
	RevisionId string `json:"revisionId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Repository") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s SourceReference) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// Status: The `Status` type defines a logical error model that is
// suitable for different programming environments, including REST APIs
// and RPC APIs. It is used by [gRPC](https://github.com/grpc). The
// error model is designed to be: - Simple to use and understand for
// most users - Flexible enough to meet unexpected needs # Overview The
// `Status` message contains three pieces of data: error code, error
// message, and error details. The error code should be an enum value of
// [google.rpc.Code][], but it may accept additional error codes if
// needed. The error message should be a developer-facing English
// message that helps developers *understand* and *resolve* the error.
// If a localized user-facing error message is needed, put the localized
// message in the error details or localize it in the client. The
// optional error details may contain arbitrary information about the
// error. There is a predefined set of error detail types in the package
// `google.rpc` which can be used for common error conditions. #
// Language mapping The `Status` message is the logical representation
// of the error model, but it is not necessarily the actual wire format.
// When the `Status` message is exposed in different client libraries
// and different wire protocols, it can be mapped differently. For
// example, it will likely be mapped to some exceptions in Java, but
// more likely mapped to some error codes in C. # Other uses The error
// model and the `Status` message can be used in a variety of
// environments, either with or without APIs, to provide a consistent
// developer experience across different environments. Example uses of
// this error model include: - Partial errors. If a service needs to
// return partial errors to the client, it may embed the `Status` in the
// normal response to indicate the partial errors. - Workflow errors. A
// typical workflow has multiple steps. Each step may have a `Status`
// message for error reporting purpose. - Batch operations. If a client
// uses batch request and batch response, the `Status` message should be
// used directly inside batch response, one for each error sub-response.
// - Asynchronous operations. If an API call embeds asynchronous
// operation results in its response, the status of those operations
// should be represented directly using the `Status` message. - Logging.
// If some API errors are stored in logs, the message `Status` could be
// used directly after any stripping needed for security/privacy
// reasons.
type Status struct {
	// Code: The status code, which should be an enum value of
	// [google.rpc.Code][].
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details. There will
	// be a common set of message types for APIs to use.
	Details []StatusDetails `json:"details,omitempty"`

	// Message: A developer-facing error message, which should be in
	// English. Any user-facing error message should be localized and sent
	// in the [google.rpc.Status.details][google.rpc.Status.details] field,
	// or localized by the client.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s Status) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

type StatusDetails interface{}

// Table: Message containing BigQuery table information.
type Table struct {
	// CreateTime: The creation time for this table.
	CreateTime string `json:"createTime,omitempty"`

	// ExpireTime: The expiration date for this table. After this time, the
	// table will not be externally visible and all storage associated with
	// the table may be garbage collected. If this field is not present, the
	// HelixDataset.default_table_expiration_ms value will be used to
	// calculate the expiration time. Otherwise, the table will live until
	// explicitly deleted.
	ExpireTime string `json:"expireTime,omitempty"`

	// Info: User-modifiable metadata for this table.
	Info *TableInfo `json:"info,omitempty"`

	// Schema: The table schema.
	Schema *TableSchema `json:"schema,omitempty"`

	// TableName: The table and dataset IDs uniquely describing this table.
	TableName *TableName `json:"tableName,omitempty"`

	// TruncateTime: The last truncation time for this table. This will only
	// be updated when operation specified with WRITE_TRUNCATE.
	TruncateTime string `json:"truncateTime,omitempty"`

	// View: The table provides a Database View behavior and functionality
	// based on a query.
	View *TableViewDefinition `json:"view,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CreateTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s Table) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// TableCopy: Describes a copy job, which copies an existing table to
// another table.
type TableCopy struct {
	// CreateDisposition: Describe when a job should create a table.
	CreateDisposition string `json:"createDisposition,omitempty"`

	// DestinationTable: Destination table.
	DestinationTable *TableName `json:"destinationTable,omitempty"`

	// SourceTables: Source tables.
	SourceTables []*TableName `json:"sourceTables,omitempty"`

	// WriteDisposition: Describe whether the copy operation should append
	// or not.
	WriteDisposition string `json:"writeDisposition,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CreateDisposition")
	// to unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s TableCopy) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// TableDataListRequest: Table data-list request.
type TableDataListRequest struct {
	// MaxResults: Maximum number of results to return.
	MaxResults int64 `json:"maxResults,omitempty"`

	// StartRow: Starting row offset.
	StartRow uint64 `json:"startRow,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "MaxResults") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s TableDataListRequest) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// TableDefinition: Per Query external tables. These tables can be
// referenced with 'name' in the query and can be read just like any
// other table.
type TableDefinition struct {
	// Name: Name of the table. This will be used to reference this table in
	// the query.
	Name string `json:"name,omitempty"`

	// SourceUris: URIs for the data to be imported.
	SourceUris []string `json:"sourceUris,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Name") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s TableDefinition) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// TableInfo: User-provided metadata for a table, primarily for display
// in the UI.
type TableInfo struct {
	// Description: The description of a table. This can be several
	// sentences or paragraphs describing the table contents in detail.
	Description string `json:"description,omitempty"`

	// FriendlyName: The human-readable name of a table. This should be a
	// short phrase identifying the table (e.g., "Analytics Data - Jan
	// 2011").
	FriendlyName string `json:"friendlyName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s TableInfo) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// TableInsertRequest: ==== Table =======// Table insert request.
type TableInsertRequest struct {
	// Resource: Table insert payload.
	Resource *Table `json:"resource,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Resource") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s TableInsertRequest) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// TableInsertResponse: Table insert response.
type TableInsertResponse struct {
	// Resource: Final state of inserted table.
	Resource *Table `json:"resource,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Resource") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s TableInsertResponse) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// TableName: Fully-qualified name for a table -- referenced through a
// dataset.
type TableName struct {
	// DatasetId: The ID of the dataset (scoped to the project above).
	DatasetId string `json:"datasetId,omitempty"`

	// ProjectId: A string containing the id of this project. The id be the
	// alphanumeric project ID, or the project number.
	ProjectId string `json:"projectId,omitempty"`

	// TableId: The ID of the table (scoped to the dataset above).
	TableId string `json:"tableId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DatasetId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s TableName) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// TableSchema: BigQuery table schema.
type TableSchema struct {
	// Fields: One field per column in the table
	Fields []*FieldSchema `json:"fields,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Fields") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s TableSchema) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// TableUpdateRequest: Table update request.
type TableUpdateRequest struct {
	// Resource: Table update payload.
	Resource *Table `json:"resource,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Resource") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s TableUpdateRequest) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// TableUpdateResponse: Table update response.
type TableUpdateResponse struct {
	// Resource: Final state of updated table.
	Resource *Table `json:"resource,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Resource") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s TableUpdateResponse) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// TableViewDefinition: Metadata for a table to become like a Database
// View based on a SQL-like query.
type TableViewDefinition struct {
	// Query: Sql query to run.
	Query string `json:"query,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Query") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s TableViewDefinition) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// WriteLogEntriesRequest: The parameters to WriteLogEntries.
type WriteLogEntriesRequest struct {
	// CommonLabels: Metadata labels that apply to all log entries in this
	// request, so that you don't have to repeat them in each log entry's
	// `metadata.labels` field. If any of the log entries contains a (key,
	// value) with the same key that is in `commonLabels`, then the entry's
	// (key, value) overrides the one in `commonLabels`.
	CommonLabels map[string]string `json:"commonLabels,omitempty"`

	// Entries: Log entries to insert.
	Entries []*LogEntry `json:"entries,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CommonLabels") to
	// unconditionally include in API requests. By default, fields with
	// empty values are ommitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string
}

func (s WriteLogEntriesRequest) MarshalJSON() ([]byte, error) {
	mustInclude := make(map[string]struct{})
	for _, f := range s.ForceSendFields {
		mustInclude[f] = struct{}{}
	}
	dataMap, err := googleapi.SchemaToMap(s, mustInclude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dataMap)
}

// WriteLogEntriesResponse: Result returned from WriteLogEntries. empty
type WriteLogEntriesResponse struct {
}

// method id "logging.projects.logServices.list":

type ProjectsLogServicesListCall struct {
	s          *Service
	projectsId string
	opt_       map[string]interface{}
}

// List: Lists log services associated with log entries ingested for a
// project.
func (r *ProjectsLogServicesService) List(projectsId string) *ProjectsLogServicesListCall {
	c := &ProjectsLogServicesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	return c
}

// Log sets the optional parameter "log": The name of the log resource
// whose services are to be listed. log for which to list services. When
// empty, all services are listed.
func (c *ProjectsLogServicesListCall) Log(log string) *ProjectsLogServicesListCall {
	c.opt_["log"] = log
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of `LogService` objects to return in one operation.
func (c *ProjectsLogServicesListCall) PageSize(pageSize int64) *ProjectsLogServicesListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": An opaque token,
// returned as `nextPageToken` by a prior `ListLogServices` operation.
// If `pageToken` is supplied, then the other fields of this request are
// ignored, and instead the previous `ListLogServices` operation is
// continued.
func (c *ProjectsLogServicesListCall) PageToken(pageToken string) *ProjectsLogServicesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesListCall) Fields(s ...googleapi.Field) *ProjectsLogServicesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogServicesListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["log"]; ok {
		params.Set("log", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logServices")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *ProjectsLogServicesListCall) Do() (*ListLogServicesResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListLogServicesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists log services associated with log entries ingested for a project.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logServices.list",
	//   "parameterOrder": [
	//     "projectsId"
	//   ],
	//   "parameters": {
	//     "log": {
	//       "description": "The name of the log resource whose services are to be listed. log for which to list services. When empty, all services are listed.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of `LogService` objects to return in one operation.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "An opaque token, returned as `nextPageToken` by a prior `ListLogServices` operation. If `pageToken` is supplied, then the other fields of this request are ignored, and instead the previous `ListLogServices` operation is continued.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `projectName`. The project resource whose services are to be listed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logServices",
	//   "response": {
	//     "$ref": "ListLogServicesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read"
	//   ]
	// }

}

// method id "logging.projects.logServices.indexes.list":

type ProjectsLogServicesIndexesListCall struct {
	s             *Service
	projectsId    string
	logServicesId string
	opt_          map[string]interface{}
}

// List: Lists log service indexes associated with a log service.
func (r *ProjectsLogServicesIndexesService) List(projectsId string, logServicesId string) *ProjectsLogServicesIndexesListCall {
	c := &ProjectsLogServicesIndexesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.logServicesId = logServicesId
	return c
}

// Depth sets the optional parameter "depth": A limit to the number of
// levels of the index hierarchy that are expanded. If `depth` is 0, it
// defaults to the level specified by the prefix field (the number of
// slash separators). The default empty prefix implies a `depth` of 1.
// It is an error for `depth` to be any non-zero value less than the
// number of components in `indexPrefix`.
func (c *ProjectsLogServicesIndexesListCall) Depth(depth int64) *ProjectsLogServicesIndexesListCall {
	c.opt_["depth"] = depth
	return c
}

// IndexPrefix sets the optional parameter "indexPrefix": Restricts the
// indexes returned to be those with a specified prefix. The prefix has
// the form "/label_value/label_value/...", in order corresponding to
// the [`LogService
// indexKeys`][google.logging.v1.LogService.index_keys]. Non-empty
// prefixes must begin with `/` . Example prefixes: + "/myModule/"
// retrieves App Engine versions associated with `myModule`. The
// trailing slash terminates the value. + "/myModule" retrieves App
// Engine modules with names beginning with `myModule`. + "" retrieves
// all indexes.
func (c *ProjectsLogServicesIndexesListCall) IndexPrefix(indexPrefix string) *ProjectsLogServicesIndexesListCall {
	c.opt_["indexPrefix"] = indexPrefix
	return c
}

// Log sets the optional parameter "log": A log resource like
// `/projects/project_id/logs/log_name`, identifying the log for which
// to list service indexes.
func (c *ProjectsLogServicesIndexesListCall) Log(log string) *ProjectsLogServicesIndexesListCall {
	c.opt_["log"] = log
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of log service index resources to return in one operation.
func (c *ProjectsLogServicesIndexesListCall) PageSize(pageSize int64) *ProjectsLogServicesIndexesListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": An opaque token,
// returned as `nextPageToken` by a prior `ListLogServiceIndexes`
// operation. If `pageToken` is supplied, then the other fields of this
// request are ignored, and instead the previous `ListLogServiceIndexes`
// operation is continued.
func (c *ProjectsLogServicesIndexesListCall) PageToken(pageToken string) *ProjectsLogServicesIndexesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesIndexesListCall) Fields(s ...googleapi.Field) *ProjectsLogServicesIndexesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogServicesIndexesListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["depth"]; ok {
		params.Set("depth", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["indexPrefix"]; ok {
		params.Set("indexPrefix", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["log"]; ok {
		params.Set("log", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logServices/{logServicesId}/indexes")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId":    c.projectsId,
		"logServicesId": c.logServicesId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *ProjectsLogServicesIndexesListCall) Do() (*ListLogServiceIndexesResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListLogServiceIndexesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists log service indexes associated with a log service.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logServices.indexes.list",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logServicesId"
	//   ],
	//   "parameters": {
	//     "depth": {
	//       "description": "A limit to the number of levels of the index hierarchy that are expanded. If `depth` is 0, it defaults to the level specified by the prefix field (the number of slash separators). The default empty prefix implies a `depth` of 1. It is an error for `depth` to be any non-zero value less than the number of components in `indexPrefix`.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "indexPrefix": {
	//       "description": "Restricts the indexes returned to be those with a specified prefix. The prefix has the form `\"/label_value/label_value/...\"`, in order corresponding to the [`LogService indexKeys`][google.logging.v1.LogService.index_keys]. Non-empty prefixes must begin with `/` . Example prefixes: + `\"/myModule/\"` retrieves App Engine versions associated with `myModule`. The trailing slash terminates the value. + `\"/myModule\"` retrieves App Engine modules with names beginning with `myModule`. + `\"\"` retrieves all indexes.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "log": {
	//       "description": "A log resource like `/projects/project_id/logs/log_name`, identifying the log for which to list service indexes.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "logServicesId": {
	//       "description": "Part of `serviceName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of log service index resources to return in one operation.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "An opaque token, returned as `nextPageToken` by a prior `ListLogServiceIndexes` operation. If `pageToken` is supplied, then the other fields of this request are ignored, and instead the previous `ListLogServiceIndexes` operation is continued.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `serviceName`. A log service resource of the form `/projects/*/logServices/*`. The service indexes of the log service are returned. Example: `\"/projects/myProj/logServices/appengine.googleapis.com\"`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logServices/{logServicesId}/indexes",
	//   "response": {
	//     "$ref": "ListLogServiceIndexesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read"
	//   ]
	// }

}

// method id "logging.projects.logServices.sinks.create":

type ProjectsLogServicesSinksCreateCall struct {
	s             *Service
	projectsId    string
	logServicesId string
	logsink       *LogSink
	opt_          map[string]interface{}
}

// Create: Creates the specified log service sink resource.
func (r *ProjectsLogServicesSinksService) Create(projectsId string, logServicesId string, logsink *LogSink) *ProjectsLogServicesSinksCreateCall {
	c := &ProjectsLogServicesSinksCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.logServicesId = logServicesId
	c.logsink = logsink
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesSinksCreateCall) Fields(s ...googleapi.Field) *ProjectsLogServicesSinksCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogServicesSinksCreateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.logsink)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId":    c.projectsId,
		"logServicesId": c.logServicesId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *ProjectsLogServicesSinksCreateCall) Do() (*LogSink, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *LogSink
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates the specified log service sink resource.",
	//   "httpMethod": "POST",
	//   "id": "logging.projects.logServices.sinks.create",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logServicesId"
	//   ],
	//   "parameters": {
	//     "logServicesId": {
	//       "description": "Part of `serviceName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `serviceName`. The name of the service in which to create a sink.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks",
	//   "request": {
	//     "$ref": "LogSink"
	//   },
	//   "response": {
	//     "$ref": "LogSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin"
	//   ]
	// }

}

// method id "logging.projects.logServices.sinks.delete":

type ProjectsLogServicesSinksDeleteCall struct {
	s             *Service
	projectsId    string
	logServicesId string
	sinksId       string
	opt_          map[string]interface{}
}

// Delete: Deletes the specified log service sink.
func (r *ProjectsLogServicesSinksService) Delete(projectsId string, logServicesId string, sinksId string) *ProjectsLogServicesSinksDeleteCall {
	c := &ProjectsLogServicesSinksDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.logServicesId = logServicesId
	c.sinksId = sinksId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesSinksDeleteCall) Fields(s ...googleapi.Field) *ProjectsLogServicesSinksDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogServicesSinksDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks/{sinksId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId":    c.projectsId,
		"logServicesId": c.logServicesId,
		"sinksId":       c.sinksId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *ProjectsLogServicesSinksDeleteCall) Do() (*Empty, error) {
	res, err := c.doRequest("json")
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
	//   "description": "Deletes the specified log service sink.",
	//   "httpMethod": "DELETE",
	//   "id": "logging.projects.logServices.sinks.delete",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logServicesId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "logServicesId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `sinkName`. The name of the sink to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sinksId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks/{sinksId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin"
	//   ]
	// }

}

// method id "logging.projects.logServices.sinks.get":

type ProjectsLogServicesSinksGetCall struct {
	s             *Service
	projectsId    string
	logServicesId string
	sinksId       string
	opt_          map[string]interface{}
}

// Get: Gets the specified log service sink resource.
func (r *ProjectsLogServicesSinksService) Get(projectsId string, logServicesId string, sinksId string) *ProjectsLogServicesSinksGetCall {
	c := &ProjectsLogServicesSinksGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.logServicesId = logServicesId
	c.sinksId = sinksId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesSinksGetCall) Fields(s ...googleapi.Field) *ProjectsLogServicesSinksGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogServicesSinksGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks/{sinksId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId":    c.projectsId,
		"logServicesId": c.logServicesId,
		"sinksId":       c.sinksId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *ProjectsLogServicesSinksGetCall) Do() (*LogSink, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *LogSink
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the specified log service sink resource.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logServices.sinks.get",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logServicesId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "logServicesId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `sinkName`. The name of the sink to return.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sinksId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks/{sinksId}",
	//   "response": {
	//     "$ref": "LogSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read"
	//   ]
	// }

}

// method id "logging.projects.logServices.sinks.list":

type ProjectsLogServicesSinksListCall struct {
	s             *Service
	projectsId    string
	logServicesId string
	opt_          map[string]interface{}
}

// List: Lists log service sinks associated with the specified service.
func (r *ProjectsLogServicesSinksService) List(projectsId string, logServicesId string) *ProjectsLogServicesSinksListCall {
	c := &ProjectsLogServicesSinksListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.logServicesId = logServicesId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesSinksListCall) Fields(s ...googleapi.Field) *ProjectsLogServicesSinksListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogServicesSinksListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId":    c.projectsId,
		"logServicesId": c.logServicesId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *ProjectsLogServicesSinksListCall) Do() (*ListLogServiceSinksResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListLogServiceSinksResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists log service sinks associated with the specified service.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logServices.sinks.list",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logServicesId"
	//   ],
	//   "parameters": {
	//     "logServicesId": {
	//       "description": "Part of `serviceName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `serviceName`. The name of the service for which to list sinks.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks",
	//   "response": {
	//     "$ref": "ListLogServiceSinksResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read"
	//   ]
	// }

}

// method id "logging.projects.logServices.sinks.update":

type ProjectsLogServicesSinksUpdateCall struct {
	s             *Service
	projectsId    string
	logServicesId string
	sinksId       string
	logsink       *LogSink
	opt_          map[string]interface{}
}

// Update: Creates or update the specified log service sink resource.
func (r *ProjectsLogServicesSinksService) Update(projectsId string, logServicesId string, sinksId string, logsink *LogSink) *ProjectsLogServicesSinksUpdateCall {
	c := &ProjectsLogServicesSinksUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.logServicesId = logServicesId
	c.sinksId = sinksId
	c.logsink = logsink
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesSinksUpdateCall) Fields(s ...googleapi.Field) *ProjectsLogServicesSinksUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogServicesSinksUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.logsink)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks/{sinksId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId":    c.projectsId,
		"logServicesId": c.logServicesId,
		"sinksId":       c.sinksId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *ProjectsLogServicesSinksUpdateCall) Do() (*LogSink, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *LogSink
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates or update the specified log service sink resource.",
	//   "httpMethod": "PUT",
	//   "id": "logging.projects.logServices.sinks.update",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logServicesId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "logServicesId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `sinkName`. The name of the sink to update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sinksId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks/{sinksId}",
	//   "request": {
	//     "$ref": "LogSink"
	//   },
	//   "response": {
	//     "$ref": "LogSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin"
	//   ]
	// }

}

// method id "logging.projects.logs.delete":

type ProjectsLogsDeleteCall struct {
	s          *Service
	projectsId string
	logsId     string
	opt_       map[string]interface{}
}

// Delete: Deletes the specified log resource and all log entries
// contained in it.
func (r *ProjectsLogsService) Delete(projectsId string, logsId string) *ProjectsLogsDeleteCall {
	c := &ProjectsLogsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.logsId = logsId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsDeleteCall) Fields(s ...googleapi.Field) *ProjectsLogsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogsDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logs/{logsId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"logsId":     c.logsId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *ProjectsLogsDeleteCall) Do() (*Empty, error) {
	res, err := c.doRequest("json")
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
	//   "description": "Deletes the specified log resource and all log entries contained in it.",
	//   "httpMethod": "DELETE",
	//   "id": "logging.projects.logs.delete",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logsId"
	//   ],
	//   "parameters": {
	//     "logsId": {
	//       "description": "Part of `logName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `logName`. The log resource to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logs/{logsId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin"
	//   ]
	// }

}

// method id "logging.projects.logs.list":

type ProjectsLogsListCall struct {
	s          *Service
	projectsId string
	opt_       map[string]interface{}
}

// List: Lists log resources belonging to the specified project.
func (r *ProjectsLogsService) List(projectsId string) *ProjectsLogsListCall {
	c := &ProjectsLogsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to return.
func (c *ProjectsLogsListCall) PageSize(pageSize int64) *ProjectsLogsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": An opaque token,
// returned as `nextPageToken` by a prior `ListLogs` operation. If
// `pageToken` is supplied, then the other fields of this request are
// ignored, and instead the previous `ListLogs` operation is continued.
func (c *ProjectsLogsListCall) PageToken(pageToken string) *ProjectsLogsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// ServiceIndexPrefix sets the optional parameter "serviceIndexPrefix":
// A log service index prefix for which to list logs. Only logs
// containing entries whose metadata that includes these label values
// (associated with index keys) are returned. The prefix is a slash
// separated list of values, and need not specify all index labels. An
// empty index (or a single slash) matches all log service indexes.
func (c *ProjectsLogsListCall) ServiceIndexPrefix(serviceIndexPrefix string) *ProjectsLogsListCall {
	c.opt_["serviceIndexPrefix"] = serviceIndexPrefix
	return c
}

// ServiceName sets the optional parameter "serviceName": A service name
// for which to list logs. Only logs containing entries whose metadata
// includes this service name are returned. If `serviceName` and
// `serviceIndexPrefix` are both empty, then all log names are returned.
// To list all log names, regardless of service, leave both the
// `serviceName` and `serviceIndexPrefix` empty. To list log names
// containing entries with a particular service name (or explicitly
// empty service name) set `serviceName` to the desired value and
// `serviceIndexPrefix` to "/".
func (c *ProjectsLogsListCall) ServiceName(serviceName string) *ProjectsLogsListCall {
	c.opt_["serviceName"] = serviceName
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsListCall) Fields(s ...googleapi.Field) *ProjectsLogsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["serviceIndexPrefix"]; ok {
		params.Set("serviceIndexPrefix", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["serviceName"]; ok {
		params.Set("serviceName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logs")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *ProjectsLogsListCall) Do() (*ListLogsResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListLogsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists log resources belonging to the specified project.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logs.list",
	//   "parameterOrder": [
	//     "projectsId"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "An opaque token, returned as `nextPageToken` by a prior `ListLogs` operation. If `pageToken` is supplied, then the other fields of this request are ignored, and instead the previous `ListLogs` operation is continued.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `projectName`. The project name for which to list the log resources.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "serviceIndexPrefix": {
	//       "description": "A log service index prefix for which to list logs. Only logs containing entries whose metadata that includes these label values (associated with index keys) are returned. The prefix is a slash separated list of values, and need not specify all index labels. An empty index (or a single slash) matches all log service indexes.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "serviceName": {
	//       "description": "A service name for which to list logs. Only logs containing entries whose metadata includes this service name are returned. If `serviceName` and `serviceIndexPrefix` are both empty, then all log names are returned. To list all log names, regardless of service, leave both the `serviceName` and `serviceIndexPrefix` empty. To list log names containing entries with a particular service name (or explicitly empty service name) set `serviceName` to the desired value and `serviceIndexPrefix` to `\"/\"`.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logs",
	//   "response": {
	//     "$ref": "ListLogsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read"
	//   ]
	// }

}

// method id "logging.projects.logs.entries.write":

type ProjectsLogsEntriesWriteCall struct {
	s                      *Service
	projectsId             string
	logsId                 string
	writelogentriesrequest *WriteLogEntriesRequest
	opt_                   map[string]interface{}
}

// Write: Creates one or more log entries in a log. You must supply a
// list of `LogEntry` objects, named `entries`. Each `LogEntry` object
// must contain a payload object and a `LogEntryMetadata` object that
// describes the entry. You must fill in all the fields of the entry,
// metadata, and payload. You can also supply a map, `commonLabels`,
// that supplies default (key, value) data for the
// `entries[].metadata.labels` maps, saving you the trouble of creating
// identical copies for each entry.
func (r *ProjectsLogsEntriesService) Write(projectsId string, logsId string, writelogentriesrequest *WriteLogEntriesRequest) *ProjectsLogsEntriesWriteCall {
	c := &ProjectsLogsEntriesWriteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.logsId = logsId
	c.writelogentriesrequest = writelogentriesrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsEntriesWriteCall) Fields(s ...googleapi.Field) *ProjectsLogsEntriesWriteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogsEntriesWriteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.writelogentriesrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logs/{logsId}/entries:write")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"logsId":     c.logsId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *ProjectsLogsEntriesWriteCall) Do() (*WriteLogEntriesResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *WriteLogEntriesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates one or more log entries in a log. You must supply a list of `LogEntry` objects, named `entries`. Each `LogEntry` object must contain a payload object and a `LogEntryMetadata` object that describes the entry. You must fill in all the fields of the entry, metadata, and payload. You can also supply a map, `commonLabels`, that supplies default (key, value) data for the `entries[].metadata.labels` maps, saving you the trouble of creating identical copies for each entry.",
	//   "httpMethod": "POST",
	//   "id": "logging.projects.logs.entries.write",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logsId"
	//   ],
	//   "parameters": {
	//     "logsId": {
	//       "description": "Part of `logName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `logName`. The name of the log resource into which to insert the log entries.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logs/{logsId}/entries:write",
	//   "request": {
	//     "$ref": "WriteLogEntriesRequest"
	//   },
	//   "response": {
	//     "$ref": "WriteLogEntriesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.write"
	//   ]
	// }

}

// method id "logging.projects.logs.sinks.create":

type ProjectsLogsSinksCreateCall struct {
	s          *Service
	projectsId string
	logsId     string
	logsink    *LogSink
	opt_       map[string]interface{}
}

// Create: Creates the specified log sink resource.
func (r *ProjectsLogsSinksService) Create(projectsId string, logsId string, logsink *LogSink) *ProjectsLogsSinksCreateCall {
	c := &ProjectsLogsSinksCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.logsId = logsId
	c.logsink = logsink
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsSinksCreateCall) Fields(s ...googleapi.Field) *ProjectsLogsSinksCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogsSinksCreateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.logsink)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logs/{logsId}/sinks")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"logsId":     c.logsId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *ProjectsLogsSinksCreateCall) Do() (*LogSink, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *LogSink
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates the specified log sink resource.",
	//   "httpMethod": "POST",
	//   "id": "logging.projects.logs.sinks.create",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logsId"
	//   ],
	//   "parameters": {
	//     "logsId": {
	//       "description": "Part of `logName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `logName`. The log in which to create a sink resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logs/{logsId}/sinks",
	//   "request": {
	//     "$ref": "LogSink"
	//   },
	//   "response": {
	//     "$ref": "LogSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin"
	//   ]
	// }

}

// method id "logging.projects.logs.sinks.delete":

type ProjectsLogsSinksDeleteCall struct {
	s          *Service
	projectsId string
	logsId     string
	sinksId    string
	opt_       map[string]interface{}
}

// Delete: Deletes the specified log sink resource.
func (r *ProjectsLogsSinksService) Delete(projectsId string, logsId string, sinksId string) *ProjectsLogsSinksDeleteCall {
	c := &ProjectsLogsSinksDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.logsId = logsId
	c.sinksId = sinksId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsSinksDeleteCall) Fields(s ...googleapi.Field) *ProjectsLogsSinksDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogsSinksDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logs/{logsId}/sinks/{sinksId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"logsId":     c.logsId,
		"sinksId":    c.sinksId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *ProjectsLogsSinksDeleteCall) Do() (*Empty, error) {
	res, err := c.doRequest("json")
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
	//   "description": "Deletes the specified log sink resource.",
	//   "httpMethod": "DELETE",
	//   "id": "logging.projects.logs.sinks.delete",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logsId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "logsId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `sinkName`. The name of the sink to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sinksId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logs/{logsId}/sinks/{sinksId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin"
	//   ]
	// }

}

// method id "logging.projects.logs.sinks.get":

type ProjectsLogsSinksGetCall struct {
	s          *Service
	projectsId string
	logsId     string
	sinksId    string
	opt_       map[string]interface{}
}

// Get: Gets the specified log sink resource.
func (r *ProjectsLogsSinksService) Get(projectsId string, logsId string, sinksId string) *ProjectsLogsSinksGetCall {
	c := &ProjectsLogsSinksGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.logsId = logsId
	c.sinksId = sinksId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsSinksGetCall) Fields(s ...googleapi.Field) *ProjectsLogsSinksGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogsSinksGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logs/{logsId}/sinks/{sinksId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"logsId":     c.logsId,
		"sinksId":    c.sinksId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *ProjectsLogsSinksGetCall) Do() (*LogSink, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *LogSink
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the specified log sink resource.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logs.sinks.get",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logsId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "logsId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `sinkName`. The name of the sink resource to return.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sinksId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logs/{logsId}/sinks/{sinksId}",
	//   "response": {
	//     "$ref": "LogSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read"
	//   ]
	// }

}

// method id "logging.projects.logs.sinks.list":

type ProjectsLogsSinksListCall struct {
	s          *Service
	projectsId string
	logsId     string
	opt_       map[string]interface{}
}

// List: Lists log sinks associated with the specified log.
func (r *ProjectsLogsSinksService) List(projectsId string, logsId string) *ProjectsLogsSinksListCall {
	c := &ProjectsLogsSinksListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.logsId = logsId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsSinksListCall) Fields(s ...googleapi.Field) *ProjectsLogsSinksListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogsSinksListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logs/{logsId}/sinks")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"logsId":     c.logsId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *ProjectsLogsSinksListCall) Do() (*ListLogSinksResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListLogSinksResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists log sinks associated with the specified log.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logs.sinks.list",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logsId"
	//   ],
	//   "parameters": {
	//     "logsId": {
	//       "description": "Part of `logName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `logName`. The log for which to list sinks.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logs/{logsId}/sinks",
	//   "response": {
	//     "$ref": "ListLogSinksResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read"
	//   ]
	// }

}

// method id "logging.projects.logs.sinks.update":

type ProjectsLogsSinksUpdateCall struct {
	s          *Service
	projectsId string
	logsId     string
	sinksId    string
	logsink    *LogSink
	opt_       map[string]interface{}
}

// Update: Creates or updates the specified log sink resource.
func (r *ProjectsLogsSinksService) Update(projectsId string, logsId string, sinksId string, logsink *LogSink) *ProjectsLogsSinksUpdateCall {
	c := &ProjectsLogsSinksUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.logsId = logsId
	c.sinksId = sinksId
	c.logsink = logsink
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsSinksUpdateCall) Fields(s ...googleapi.Field) *ProjectsLogsSinksUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogsSinksUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.logsink)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logs/{logsId}/sinks/{sinksId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"logsId":     c.logsId,
		"sinksId":    c.sinksId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *ProjectsLogsSinksUpdateCall) Do() (*LogSink, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *LogSink
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates or updates the specified log sink resource.",
	//   "httpMethod": "PUT",
	//   "id": "logging.projects.logs.sinks.update",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logsId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "logsId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `sinkName`. The name of the sink to update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sinksId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logs/{logsId}/sinks/{sinksId}",
	//   "request": {
	//     "$ref": "LogSink"
	//   },
	//   "response": {
	//     "$ref": "LogSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin"
	//   ]
	// }

}

// method id "logging.projects.sinks.create":

type ProjectsSinksCreateCall struct {
	s          *Service
	projectsId string
	logsink    *LogSink
	opt_       map[string]interface{}
}

// Create: Creates the specified sink resource.
func (r *ProjectsSinksService) Create(projectsId string, logsink *LogSink) *ProjectsSinksCreateCall {
	c := &ProjectsSinksCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.logsink = logsink
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSinksCreateCall) Fields(s ...googleapi.Field) *ProjectsSinksCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsSinksCreateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.logsink)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/sinks")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *ProjectsSinksCreateCall) Do() (*LogSink, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *LogSink
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates the specified sink resource.",
	//   "httpMethod": "POST",
	//   "id": "logging.projects.sinks.create",
	//   "parameterOrder": [
	//     "projectsId"
	//   ],
	//   "parameters": {
	//     "projectsId": {
	//       "description": "Part of `projectName`. The name of the project in which to create a sink.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/sinks",
	//   "request": {
	//     "$ref": "LogSink"
	//   },
	//   "response": {
	//     "$ref": "LogSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin"
	//   ]
	// }

}

// method id "logging.projects.sinks.delete":

type ProjectsSinksDeleteCall struct {
	s          *Service
	projectsId string
	sinksId    string
	opt_       map[string]interface{}
}

// Delete: Deletes the specified sink.
func (r *ProjectsSinksService) Delete(projectsId string, sinksId string) *ProjectsSinksDeleteCall {
	c := &ProjectsSinksDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.sinksId = sinksId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSinksDeleteCall) Fields(s ...googleapi.Field) *ProjectsSinksDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsSinksDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/sinks/{sinksId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"sinksId":    c.sinksId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *ProjectsSinksDeleteCall) Do() (*Empty, error) {
	res, err := c.doRequest("json")
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
	//   "description": "Deletes the specified sink.",
	//   "httpMethod": "DELETE",
	//   "id": "logging.projects.sinks.delete",
	//   "parameterOrder": [
	//     "projectsId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "projectsId": {
	//       "description": "Part of `sinkName`. The name of the sink to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sinksId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/sinks/{sinksId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin"
	//   ]
	// }

}

// method id "logging.projects.sinks.get":

type ProjectsSinksGetCall struct {
	s          *Service
	projectsId string
	sinksId    string
	opt_       map[string]interface{}
}

// Get: Gets the specified sink resource.
func (r *ProjectsSinksService) Get(projectsId string, sinksId string) *ProjectsSinksGetCall {
	c := &ProjectsSinksGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.sinksId = sinksId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSinksGetCall) Fields(s ...googleapi.Field) *ProjectsSinksGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsSinksGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/sinks/{sinksId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"sinksId":    c.sinksId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *ProjectsSinksGetCall) Do() (*LogSink, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *LogSink
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the specified sink resource.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.sinks.get",
	//   "parameterOrder": [
	//     "projectsId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "projectsId": {
	//       "description": "Part of `sinkName`. The name of the sink to return.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sinksId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/sinks/{sinksId}",
	//   "response": {
	//     "$ref": "LogSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read"
	//   ]
	// }

}

// method id "logging.projects.sinks.list":

type ProjectsSinksListCall struct {
	s          *Service
	projectsId string
	opt_       map[string]interface{}
}

// List: Lists sinks associated with the specified project.
func (r *ProjectsSinksService) List(projectsId string) *ProjectsSinksListCall {
	c := &ProjectsSinksListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSinksListCall) Fields(s ...googleapi.Field) *ProjectsSinksListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsSinksListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/sinks")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *ProjectsSinksListCall) Do() (*ListSinksResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListSinksResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists sinks associated with the specified project.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.sinks.list",
	//   "parameterOrder": [
	//     "projectsId"
	//   ],
	//   "parameters": {
	//     "projectsId": {
	//       "description": "Part of `projectName`. The name of the project for which to list sinks.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/sinks",
	//   "response": {
	//     "$ref": "ListSinksResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read"
	//   ]
	// }

}

// method id "logging.projects.sinks.update":

type ProjectsSinksUpdateCall struct {
	s          *Service
	projectsId string
	sinksId    string
	logsink    *LogSink
	opt_       map[string]interface{}
}

// Update: Creates or update the specified sink resource.
func (r *ProjectsSinksService) Update(projectsId string, sinksId string, logsink *LogSink) *ProjectsSinksUpdateCall {
	c := &ProjectsSinksUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.sinksId = sinksId
	c.logsink = logsink
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSinksUpdateCall) Fields(s ...googleapi.Field) *ProjectsSinksUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsSinksUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.logsink)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/sinks/{sinksId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"sinksId":    c.sinksId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *ProjectsSinksUpdateCall) Do() (*LogSink, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *LogSink
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates or update the specified sink resource.",
	//   "httpMethod": "PUT",
	//   "id": "logging.projects.sinks.update",
	//   "parameterOrder": [
	//     "projectsId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "projectsId": {
	//       "description": "Part of `sinkName`. The name of the sink to update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sinksId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/sinks/{sinksId}",
	//   "request": {
	//     "$ref": "LogSink"
	//   },
	//   "response": {
	//     "$ref": "LogSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin"
	//   ]
	// }

}
