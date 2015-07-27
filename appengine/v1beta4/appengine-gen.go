// Package appengine provides access to the Google App Engine Admin API.
//
// See https://developers.google.com/appengine/
//
// Usage example:
//
//   import "google.golang.org/api/appengine/v1beta4"
//   ...
//   appengineService, err := appengine.New(oauthHttpClient)
package appengine

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

const apiId = "appengine:v1beta4"
const apiName = "appengine"
const apiVersion = "v1beta4"
const basePath = "https://appengine.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Apps = NewAppsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Apps *AppsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewAppsService(s *Service) *AppsService {
	rs := &AppsService{s: s}
	rs.Modules = NewAppsModulesService(s)
	rs.Operations = NewAppsOperationsService(s)
	return rs
}

type AppsService struct {
	s *Service

	Modules *AppsModulesService

	Operations *AppsOperationsService
}

func NewAppsModulesService(s *Service) *AppsModulesService {
	rs := &AppsModulesService{s: s}
	rs.Versions = NewAppsModulesVersionsService(s)
	return rs
}

type AppsModulesService struct {
	s *Service

	Versions *AppsModulesVersionsService
}

func NewAppsModulesVersionsService(s *Service) *AppsModulesVersionsService {
	rs := &AppsModulesVersionsService{s: s}
	return rs
}

type AppsModulesVersionsService struct {
	s *Service
}

func NewAppsOperationsService(s *Service) *AppsOperationsService {
	rs := &AppsOperationsService{s: s}
	return rs
}

type AppsOperationsService struct {
	s *Service
}

// ApiConfigHandler: API Serving configuration for Cloud Endpoints.
type ApiConfigHandler struct {
	// AuthFailAction: For users not logged in, how to handle access to
	// resources with required login. Defaults to "redirect".
	//
	// Possible values:
	//   "AUTH_FAIL_ACTION_UNSPECIFIED"
	//   "AUTH_FAIL_ACTION_REDIRECT"
	//   "AUTH_FAIL_ACTION_UNAUTHORIZED"
	AuthFailAction string `json:"authFailAction,omitempty"`

	// Login: What level of login is required to access this resource.
	// Default is "optional".
	//
	// Possible values:
	//   "LOGIN_UNSPECIFIED"
	//   "LOGIN_OPTIONAL"
	//   "LOGIN_ADMIN"
	//   "LOGIN_REQUIRED"
	Login string `json:"login,omitempty"`

	// Script: Specifies the path to the script from the application root
	// directory.
	Script string `json:"script,omitempty"`

	// SecurityLevel: Configures whether security (HTTPS) should be enforced
	// for this URL.
	//
	// Possible values:
	//   "SECURE_UNSPECIFIED"
	//   "SECURE_DEFAULT"
	//   "SECURE_NEVER"
	//   "SECURE_OPTIONAL"
	//   "SECURE_ALWAYS"
	SecurityLevel string `json:"securityLevel,omitempty"`

	// Url: URL to serve the endpoint at.
	Url string `json:"url,omitempty"`
}

// ApiEndpointHandler: Use Google Cloud Endpoints to handle requests.
type ApiEndpointHandler struct {
	// ScriptPath: Specifies the path to the script from the application
	// root directory.
	ScriptPath string `json:"scriptPath,omitempty"`
}

// Application: An Application contains the top-level configuration of
// an App Engine application.
type Application struct {
	// CodeBucket: A Google Cloud Storage bucket which can be used for
	// storing files associated with an application. This bucket is
	// associated with the application and can be used by the gcloud
	// deployment commands. @OutputOnly
	CodeBucket string `json:"codeBucket,omitempty"`

	// DispatchRules: HTTP path dispatch rules for requests to the app that
	// do not explicitly target a module or version. The rules are
	// order-dependent.
	DispatchRules []*UrlDispatchRule `json:"dispatchRules,omitempty"`

	// Id: The relative name/path of the application. Example: "myapp".
	// @OutputOnly
	Id string `json:"id,omitempty"`

	// Location: The location from which the application will be run.
	// Choices are "us" for United States and "eu" for European Union.
	// Application instances will run out of data centers in the chosen
	// location and all of the application's End User Content will be stored
	// at rest in the chosen location. The default is "us".
	Location string `json:"location,omitempty"`

	// Name: The full path to the application in the API. Example:
	// "apps/myapp". @OutputOnly
	Name string `json:"name,omitempty"`
}

// AutomaticScaling: Automatic scaling is the scaling policy that App
// Engine has used since its inception. It is based on request rate,
// response latencies, and other application metrics.
type AutomaticScaling struct {
	// CoolDownPeriod: The amount of time that the
	// [Autoscaler](https://cloud.google.com/compute/docs/autoscaler/)
	// should wait between changes to the number of virtual machines.
	// Applies only to the VM runtime.
	CoolDownPeriod string `json:"coolDownPeriod,omitempty"`

	// CpuUtilization: Target scaling by CPU usage.
	CpuUtilization *CpuUtilization `json:"cpuUtilization,omitempty"`

	// MaxConcurrentRequests: The number of concurrent requests an automatic
	// scaling instance can accept before the scheduler spawns a new
	// instance. Default value is chosen based on the runtime.
	MaxConcurrentRequests int64 `json:"maxConcurrentRequests,omitempty"`

	// MaxIdleInstances: The maximum number of idle instances that App
	// Engine should maintain for this version.
	MaxIdleInstances int64 `json:"maxIdleInstances,omitempty"`

	// MaxPendingLatency: The maximum amount of time that App Engine should
	// allow a request to wait in the pending queue before starting a new
	// instance to handle it.
	MaxPendingLatency string `json:"maxPendingLatency,omitempty"`

	// MaxTotalInstances: Max number of instances that App Engine should
	// start to handle requests.
	MaxTotalInstances int64 `json:"maxTotalInstances,omitempty"`

	// MinIdleInstances: The minimum number of idle instances that App
	// Engine should maintain for this version. Only applies to the default
	// version of a module, since other versions are not expected to receive
	// significant traffic.
	MinIdleInstances int64 `json:"minIdleInstances,omitempty"`

	// MinPendingLatency: The minimum amount of time that App Engine should
	// allow a request to wait in the pending queue before starting a new
	// instance to handle it.
	MinPendingLatency string `json:"minPendingLatency,omitempty"`

	// MinTotalInstances: Minimum number of instances that App Engine should
	// maintain.
	MinTotalInstances int64 `json:"minTotalInstances,omitempty"`
}

// BasicScaling: A module with basic scaling will create an instance
// when the application receives a request. The instance will be turned
// down when the app becomes idle. Basic scaling is ideal for work that
// is intermittent or driven by user activity.
type BasicScaling struct {
	// IdleTimeout: The instance will be shut down this amount of time after
	// receiving its last request.
	IdleTimeout string `json:"idleTimeout,omitempty"`

	// MaxInstances: The maximum number of instances for App Engine to
	// create for this version.
	MaxInstances int64 `json:"maxInstances,omitempty"`
}

// ContainerInfo: A Docker (container) image which should be used to
// start the application.
type ContainerInfo struct {
	// Image: Reference to a hosted container image. Must be a URI to a
	// resource in a Docker repository. Must be fully qualified, including
	// tag or digest. e.g. gcr.io/my-project/image:tag or
	// gcr.io/my-project/image@digest
	Image string `json:"image,omitempty"`

	// Sha256: The SHA256 hash of the image in hex.
	Sha256 string `json:"sha256,omitempty"`
}

// CpuUtilization: Target scaling by CPU usage.
type CpuUtilization struct {
	// AggregationWindowLength: The period of time over which CPU
	// utilization is calculated.
	AggregationWindowLength string `json:"aggregationWindowLength,omitempty"`

	// TargetUtilization: Target (0-1) CPU utilization ratio to maintain
	// when scaling.
	TargetUtilization float64 `json:"targetUtilization,omitempty"`
}

// Deployment: Code and application artifacts used to deploy a version
// to App Engine.
type Deployment struct {
	// Container: If supplied, a docker (container) image which should be
	// used to start the application. Only applicable to the 'vm' runtime.
	Container *ContainerInfo `json:"container,omitempty"`

	// Files: A manifest of files stored in Google Cloud Storage which
	// should be included as part of this application. All files must be
	// readable using the credentials supplied with this call.
	Files map[string]FileInfo `json:"files,omitempty"`

	// SourceReferences: The origin of the source code for this deployment.
	// There can be more than one source reference per Version if source
	// code is distributed among multiple repositories.
	SourceReferences []*SourceReference `json:"sourceReferences,omitempty"`
}

// ErrorHandler: A custom static error page to be served when an error
// occurs.
type ErrorHandler struct {
	// ErrorCode: The error condition this handler applies to.
	//
	// Possible values:
	//   "ERROR_CODE_UNSPECIFIED"
	//   "ERROR_CODE_DEFAULT"
	//   "ERROR_CODE_OVER_QUOTA"
	//   "ERROR_CODE_DOS_API_DENIAL"
	//   "ERROR_CODE_TIMEOUT"
	ErrorCode string `json:"errorCode,omitempty"`

	// MimeType: MIME type of file. If unspecified, "text/html" is assumed.
	MimeType string `json:"mimeType,omitempty"`

	// StaticFile: Static file content to be served for this error.
	StaticFile string `json:"staticFile,omitempty"`
}

// Field: Field represents a single field of a message type.
type Field struct {
	// Cardinality: The field cardinality, i.e. optional/required/repeated.
	//
	// Possible values:
	//   "CARDINALITY_UNKNOWN"
	//   "CARDINALITY_OPTIONAL"
	//   "CARDINALITY_REQUIRED"
	//   "CARDINALITY_REPEATED"
	Cardinality string `json:"cardinality,omitempty"`

	// JsonName: The JSON name for this field.
	JsonName string `json:"jsonName,omitempty"`

	// Kind: The field kind.
	//
	// Possible values:
	//   "TYPE_UNKNOWN"
	//   "TYPE_DOUBLE"
	//   "TYPE_FLOAT"
	//   "TYPE_INT64"
	//   "TYPE_UINT64"
	//   "TYPE_INT32"
	//   "TYPE_FIXED64"
	//   "TYPE_FIXED32"
	//   "TYPE_BOOL"
	//   "TYPE_STRING"
	//   "TYPE_GROUP"
	//   "TYPE_MESSAGE"
	//   "TYPE_BYTES"
	//   "TYPE_UINT32"
	//   "TYPE_ENUM"
	//   "TYPE_SFIXED32"
	//   "TYPE_SFIXED64"
	//   "TYPE_SINT32"
	//   "TYPE_SINT64"
	Kind string `json:"kind,omitempty"`

	// Name: The field name.
	Name string `json:"name,omitempty"`

	// Number: The proto field number.
	Number int64 `json:"number,omitempty"`

	// OneofIndex: Index in Type.oneofs. Starts at 1. Zero means no oneof
	// mapping.
	OneofIndex int64 `json:"oneofIndex,omitempty"`

	// Options: The proto options.
	Options []*Option `json:"options,omitempty"`

	// Packed: Whether to use alternative packed wire representation.
	Packed bool `json:"packed,omitempty"`

	// TypeUrl: The type URL (without the scheme) when the type is MESSAGE
	// or ENUM, such as `type.googleapis.com/google.protobuf.Empty`.
	TypeUrl string `json:"typeUrl,omitempty"`
}

// FileInfo: A single source file which is part of the application to be
// deployed.
type FileInfo struct {
	// MimeType: The MIME type of the file; if unspecified, the value from
	// Google Cloud Storage will be used.
	MimeType string `json:"mimeType,omitempty"`

	// Sha1Sum: The SHA1 (160 bits) hash of the file in hex.
	Sha1Sum string `json:"sha1Sum,omitempty"`

	// SourceUrl: The URL source to use to fetch this file. Must be a URL to
	// a resource in Google Cloud Storage.
	SourceUrl string `json:"sourceUrl,omitempty"`
}

// HealthCheck: Configure health checking for the VM instances.
// Unhealthy VM instances will be killed and replaced with new
// instances.
type HealthCheck struct {
	// CheckInterval: The interval between health checks.
	CheckInterval string `json:"checkInterval,omitempty"`

	// DisableHealthCheck: Whether to explicitly disable health checks for
	// this instance.
	DisableHealthCheck bool `json:"disableHealthCheck,omitempty"`

	// HealthyThreshold: The number of consecutive successful health checks
	// before receiving traffic.
	HealthyThreshold int64 `json:"healthyThreshold,omitempty"`

	// Host: The host header to send when performing an HTTP health check
	// (e.g. myapp.appspot.com)
	Host string `json:"host,omitempty"`

	// RestartThreshold: The number of consecutive failed health checks
	// before an instance is restarted.
	RestartThreshold int64 `json:"restartThreshold,omitempty"`

	// Timeout: The amount of time before the health check is considered
	// failed.
	Timeout string `json:"timeout,omitempty"`

	// UnhealthyThreshold: The number of consecutive failed health checks
	// before removing traffic.
	UnhealthyThreshold int64 `json:"unhealthyThreshold,omitempty"`
}

// Library: A Python runtime third-party library required by the
// application.
type Library struct {
	// Name: The name of the library, e.g. "PIL" or "django".
	Name string `json:"name,omitempty"`

	// Version: The version of the library to select, or "latest".
	Version string `json:"version,omitempty"`
}

// ListModulesResponse: Response message for `Modules.ListModules`.
type ListModulesResponse struct {
	// Modules: The modules belonging to the requested application.
	Modules []*Module `json:"modules,omitempty"`

	// NextPageToken: Continuation token for fetching the next page of
	// results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

// ListOperationsResponse: The response message for
// [Operations.ListOperations][google.longrunning.Operations.ListOperatio
// ns].
type ListOperationsResponse struct {
	// NextPageToken: The standard List next-page token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Operations: A list of operations that matches the specified filter in
	// the request.
	Operations []*Operation `json:"operations,omitempty"`
}

// ListVersionsResponse: Response message for `Versions.ListVersions`.
type ListVersionsResponse struct {
	// NextPageToken: Continuation token for fetching the next page of
	// results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Versions: The versions belonging to the requested application module.
	Versions []*Version `json:"versions,omitempty"`
}

// ManualScaling: A module with manual scaling runs continuously,
// allowing you to perform complex initialization and rely on the state
// of its memory over time.
type ManualScaling struct {
	// Instances: The number of instances to assign to the module at the
	// start. This number can later be altered by using the [Modules
	// API](https://cloud.google.com/appengine/docs/python/modules/functions)
	//  `set_num_instances()` function.
	Instances int64 `json:"instances,omitempty"`
}

// Module: A module is a component of an application that provides a
// single service or configuration. A module has a collection of
// versions that define a specific set of code used to implement the
// functionality of that module.
type Module struct {
	// Id: The relative name/path of the module within the application.
	// Example: "default" @OutputOnly
	Id string `json:"id,omitempty"`

	// Name: The full path to the Module resource in the API. Example:
	// "apps/myapp/modules/default" @OutputOnly
	Name string `json:"name,omitempty"`

	// Split: A mapping that defines fractional HTTP traffic diversion to
	// different versions within the module.
	Split *TrafficSplit `json:"split,omitempty"`
}

// Network: Used to specify extra network settings (for VM runtimes
// only).
type Network struct {
	// ForwardedPorts: A list of ports (or port pairs) to forward from the
	// VM into the app container.
	ForwardedPorts []string `json:"forwardedPorts,omitempty"`

	// InstanceTag: A tag to apply to the VM instance during creation.
	InstanceTag string `json:"instanceTag,omitempty"`

	// Name: The Google Compute Engine network where the VMs will be
	// created. If not specified, or empty, the network named 'default' will
	// be used. (The short name should be specified, not the resource path.)
	Name string `json:"name,omitempty"`
}

// Operation: This resource represents a long-running operation that is
// the result of a network API call.
type Operation struct {
	// Done: If the value is `false`, it means the operation is still in
	// progress. If true, the operation is completed and the `result` is
	// available.
	Done bool `json:"done,omitempty"`

	// Error: The error result of the operation in case of failure.
	Error *Status `json:"error,omitempty"`

	// Metadata: Service-specific metadata associated with the operation. It
	// typically contains progress information and common metadata such as
	// create time. Some services might not provide such metadata. Any
	// method that returns a long-running operation should document the
	// metadata type, if any.
	Metadata OperationMetadata `json:"metadata,omitempty"`

	// Name: The server-assigned name, which is only unique within the same
	// service that originally returns it. If you use the default HTTP
	// mapping above, the `name` should have the format of
	// `operations/some/unique/name`.
	Name string `json:"name,omitempty"`

	// Response: The normal response of the operation in case of success. If
	// the original method returns no data on success, such as `Delete`, the
	// response is `google.protobuf.Empty`. If the original method is
	// standard `Get`/`Create`/`Update`, the response should be the
	// resource. For other methods, the response should have the type
	// `XxxResponse`, where `Xxx` is the original method name. For example,
	// if the original method name is `TakeSnapshot()`, the inferred
	// response type is `TakeSnapshotResponse`.
	Response OperationResponse `json:"response,omitempty"`
}

type OperationMetadata interface{}

type OperationResponse interface{}

// OperationMetadata1: Metadata for the given
// [google.longrunning.Operation][google.longrunning.Operation].
type OperationMetadata1 struct {
	// EndTime: Timestamp that this operation was completed. (Not present if
	// the operation is still in progress.) @OutputOnly
	EndTime string `json:"endTime,omitempty"`

	// InsertTime: Timestamp that this operation was received. @OutputOnly
	InsertTime string `json:"insertTime,omitempty"`

	// OperationType: The type of the operation, e.g. 'deployment'.
	// @OutputOnly
	OperationType string `json:"operationType,omitempty"`

	// User: The user who requested this operation. @OutputOnly
	User string `json:"user,omitempty"`
}

// Option: Proto option attached to messages/fields/enums etc.
type Option struct {
	// Name: Proto option name.
	Name string `json:"name,omitempty"`

	// Value: Proto option value.
	Value OptionValue `json:"value,omitempty"`
}

type OptionValue interface{}

// Resources: Used to specify how many machine resources an app version
// needs.
type Resources struct {
	// Cpu: How many CPU cores an app version needs.
	Cpu float64 `json:"cpu,omitempty"`

	// DiskGb: How much disk size, in GB, an app version needs.
	DiskGb float64 `json:"diskGb,omitempty"`

	// MemoryGb: How much memory, in GB, an app version needs.
	MemoryGb float64 `json:"memoryGb,omitempty"`
}

// ScriptHandler: Executes a script to handle the request that matches
// the URL pattern.
type ScriptHandler struct {
	// ScriptPath: Specifies the path to the script from the application
	// root directory.
	ScriptPath string `json:"scriptPath,omitempty"`
}

// SourceContext: `SourceContext` represents information about the
// source of a protobuf element, like the file in which it is defined.
type SourceContext struct {
	// FileName: The path-qualified name of the .proto file that contained
	// the associated protobuf element. For example:
	// "google/protobuf/source.proto".
	FileName string `json:"fileName,omitempty"`
}

// SourceReference: A reference to a particular snapshot of the source
// tree used to build and deploy the application.
type SourceReference struct {
	// Repository: Optional. A URI string identifying the repository.
	// Example: "https://source.developers.google.com/p/app-123/r/default"
	Repository string `json:"repository,omitempty"`

	// RevisionId: The canonical (and persistent) identifier of the deployed
	// revision, i.e. any kind of aliases including tags or branch names are
	// not allowed. Example (git):
	// "2198322f89e0bb2e25021667c2ed489d1fd34e6b"
	RevisionId string `json:"revisionId,omitempty"`
}

// StaticDirectoryHandler: Files served directly to the user for a given
// URL, such as images, CSS stylesheets, or JavaScript source files.
// Static directory handlers make it easy to serve the entire contents
// of a directory as static files.
type StaticDirectoryHandler struct {
	// ApplicationReadable: By default, files declared in static file
	// handlers are uploaded as static data and are only served to end
	// users, they cannot be read by an application. If this field is set to
	// true, the files are also uploaded as code data so your application
	// can read them. Both uploads are charged against your code and static
	// data storage resource quotas.
	ApplicationReadable bool `json:"applicationReadable,omitempty"`

	// Directory: The path to the directory containing the static files,
	// from the application root directory. Everything after the end of the
	// matched url pattern is appended to static_dir to form the full path
	// to the requested file.
	Directory string `json:"directory,omitempty"`

	// Expiration: The length of time a static file served by this handler
	// ought to be cached by web proxies and browsers.
	Expiration string `json:"expiration,omitempty"`

	// HttpHeaders: HTTP headers to use for all responses from these URLs.
	HttpHeaders map[string]string `json:"httpHeaders,omitempty"`

	// MimeType: If specified, all files served by this handler will be
	// served using the specified MIME type. If not specified, the MIME type
	// for a file will be derived from the file's filename extension.
	MimeType string `json:"mimeType,omitempty"`

	// RequireMatchingFile: If true, this UrlMap entry does not match the
	// request unless the file referenced by the handler also exists. If no
	// such file exists, processing will continue with the next UrlMap that
	// matches the requested URL.
	RequireMatchingFile bool `json:"requireMatchingFile,omitempty"`
}

// StaticFilesHandler: Files served directly to the user for a given
// URL, such as images, CSS stylesheets, or JavaScript source files.
// Static file handlers describe which files in the application
// directory are static files, and which URLs serve them.
type StaticFilesHandler struct {
	// ApplicationReadable: By default, files declared in static file
	// handlers are uploaded as static data and are only served to end
	// users, they cannot be read by an application. If this field is set to
	// true, the files are also uploaded as code data so your application
	// can read them. Both uploads are charged against your code and static
	// data storage resource quotas.
	ApplicationReadable bool `json:"applicationReadable,omitempty"`

	// Expiration: The length of time a static file served by this handler
	// ought to be cached by web proxies and browsers.
	Expiration string `json:"expiration,omitempty"`

	// HttpHeaders: HTTP headers to use for all responses from these URLs.
	HttpHeaders map[string]string `json:"httpHeaders,omitempty"`

	// MimeType: If specified, all files served by this handler will be
	// served using the specified MIME type. If not specified, the MIME type
	// for a file will be derived from the file's filename extension.
	MimeType string `json:"mimeType,omitempty"`

	// Path: The path to the static files matched by the URL pattern, from
	// the application root directory. The path can refer to text matched in
	// groupings in the URL pattern.
	Path string `json:"path,omitempty"`

	// RequireMatchingFile: If true, this
	// [UrlMap][google.appengine.v1beta4.UrlMap] entry does not match the
	// request unless the file referenced by the handler also exists. If no
	// such file exists, processing will continue with the next
	// [UrlMap][google.appengine.v1beta4.UrlMap] that matches the requested
	// URL.
	RequireMatchingFile bool `json:"requireMatchingFile,omitempty"`

	// UploadPathRegex: A regular expression that matches the file paths for
	// all files that will be referenced by this handler.
	UploadPathRegex string `json:"uploadPathRegex,omitempty"`
}

// Status: The `Status` defines a logical error model that is suitable
// for different programming environments, including REST APIs and RPC
// APIs. It is used by [gRPC](https://github.com/grpc). The error model
// is designed to be: - Simple to use and understand for most users. -
// Flexible enough to meet unexpected needs. # Overview The `Status`
// message contains 3 pieces of data: error code, error message, and
// error details. The error code should be an enum value of
// [google.rpc.Code][google.rpc.Code], but it may accept additional
// error codes if needed. The error message should be a developer-facing
// English message that helps developers *understand* and *resolve* the
// error. If a localized user-facing error message is needed, it can be
// sent in the error details or localized by the client. The optional
// error details may contain arbitrary information about the error.
// There is a predefined set of error detail types in the package
// `google.rpc` which can be used for common error conditions. #
// Language mapping The `Status` message is the logical representation
// of the error model, but it is not necessarily the actual wire format.
// When the `Status` message is exposed in different client libraries
// and different wire protocols, it can be mapped differently. For
// example, it will likely be mapped to some exceptions in Java, but
// more likely mapped to some error codes in C. # Other uses The error
// model and the `Status` message can be used in a variety of
// environments - either with or without APIs - to provide consistent
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
	// [google.rpc.Code][google.rpc.Code].
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details. There will
	// be a common set of message types for APIs to use.
	Details []StatusDetails `json:"details,omitempty"`

	// Message: A developer-facing error message, which should be in
	// English. Any user-facing error message should be localized and sent
	// in the [google.rpc.Status.details][google.rpc.Status.details] field,
	// or localized by the client.
	Message string `json:"message,omitempty"`
}

type StatusDetails interface{}

// TrafficSplit: Configuration for traffic splitting for versions within
// a single module. Traffic splitting allows traffic directed to the
// module to be assigned to one of several versions in a fractional way,
// enabling experiments and canarying new builds, for example.
type TrafficSplit struct {
	// Allocations: Mapping from module version IDs within the module to
	// fractional (0.000, 1] allocations of traffic for that version. Each
	// version may only be specified once, but some versions in the module
	// may not have any traffic allocation. Modules that have traffic
	// allocated in this field may not be deleted until the module is
	// deleted, or their traffic allocation is removed. Allocations must sum
	// to 1. Supports precision up to two decimal places for IP-based splits
	// and up to three decimal places for cookie-based splits.
	Allocations *TrafficSplitAllocations `json:"allocations,omitempty"`

	// ShardBy: Which mechanism should be used as a selector when choosing a
	// version to send a request to. The traffic selection algorithm will be
	// stable for either type until allocations are changed.
	//
	// Possible values:
	//   "UNSPECIFIED"
	//   "COOKIE"
	//   "IP"
	ShardBy string `json:"shardBy,omitempty"`
}

// TrafficSplitAllocations: Mapping from module version IDs within the
// module to fractional (0.000, 1] allocations of traffic for that
// version. Each version may only be specified once, but some versions
// in the module may not have any traffic allocation. Modules that have
// traffic allocated in this field may not be deleted until the module
// is deleted, or their traffic allocation is removed. Allocations must
// sum to 1. Supports precision up to two decimal places for IP-based
// splits and up to three decimal places for cookie-based splits.
type TrafficSplitAllocations struct {
}

// Type: A light-weight descriptor for a proto message type.
type Type struct {
	// Fields: The list of fields.
	Fields []*Field `json:"fields,omitempty"`

	// Name: The fully qualified message name.
	Name string `json:"name,omitempty"`

	// Oneofs: The list of oneof definitions. The list of oneofs declared in
	// this Type
	Oneofs []string `json:"oneofs,omitempty"`

	// Options: The proto options.
	Options []*Option `json:"options,omitempty"`

	// SourceContext: The source context.
	SourceContext *SourceContext `json:"sourceContext,omitempty"`

	// Syntax: The source syntax.
	//
	// Possible values:
	//   "SYNTAX_PROTO2"
	//   "SYNTAX_PROTO3"
	Syntax string `json:"syntax,omitempty"`
}

// UrlDispatchRule: Rules to match an HTTP request and dispatch that
// request to a module.
type UrlDispatchRule struct {
	// Domain: The domain name to match on. Supports '*' (glob) wildcarding
	// on the left-hand side of a '.'. If empty, all domains will be matched
	// (the same as '*').
	Domain string `json:"domain,omitempty"`

	// Module: The resource id of a Module in this application that should
	// service the matched request. The Module must already exist. Example:
	// "default".
	Module string `json:"module,omitempty"`

	// Path: The pathname within the host. This must start with a '/'. A
	// single '*' (glob) can be included at the end of the path. The sum of
	// the lengths of the domain and path may not exceed 100 characters.
	Path string `json:"path,omitempty"`
}

// UrlMap: A URL pattern and description of how it should be handled.
// App Engine can handle URLs by executing application code, or by
// serving static files uploaded with the code, such as images, CSS or
// JavaScript.
type UrlMap struct {
	// ApiEndpoint: Use API Endpoints to handle requests.
	ApiEndpoint *ApiEndpointHandler `json:"apiEndpoint,omitempty"`

	// AuthFailAction: For users not logged in, how to handle access to
	// resources with required login. Defaults to "redirect".
	//
	// Possible values:
	//   "AUTH_FAIL_ACTION_UNSPECIFIED"
	//   "AUTH_FAIL_ACTION_REDIRECT"
	//   "AUTH_FAIL_ACTION_UNAUTHORIZED"
	AuthFailAction string `json:"authFailAction,omitempty"`

	// Login: What level of login is required to access this resource.
	//
	// Possible values:
	//   "LOGIN_UNSPECIFIED"
	//   "LOGIN_OPTIONAL"
	//   "LOGIN_ADMIN"
	//   "LOGIN_REQUIRED"
	Login string `json:"login,omitempty"`

	// RedirectHttpResponseCode: `30x` code to use when performing redirects
	// for the `secure` field. A `302` is used by default.
	//
	// Possible values:
	//   "REDIRECT_HTTP_RESPONSE_CODE_UNSPECIFIED"
	//   "REDIRECT_HTTP_RESPONSE_CODE_301"
	//   "REDIRECT_HTTP_RESPONSE_CODE_302"
	//   "REDIRECT_HTTP_RESPONSE_CODE_303"
	//   "REDIRECT_HTTP_RESPONSE_CODE_307"
	RedirectHttpResponseCode string `json:"redirectHttpResponseCode,omitempty"`

	// Script: Executes a script to handle the request that matches the URL
	// pattern.
	Script *ScriptHandler `json:"script,omitempty"`

	// SecurityLevel: Configures whether security (HTTPS) should be enforced
	// for this URL.
	//
	// Possible values:
	//   "SECURE_UNSPECIFIED"
	//   "SECURE_DEFAULT"
	//   "SECURE_NEVER"
	//   "SECURE_OPTIONAL"
	//   "SECURE_ALWAYS"
	SecurityLevel string `json:"securityLevel,omitempty"`

	// StaticDirectory: Serves the entire contents of a directory as static
	// files. This attribute is deprecated. You can mimic the behavior of
	// static directories using static files.
	StaticDirectory *StaticDirectoryHandler `json:"staticDirectory,omitempty"`

	// StaticFiles: Returns the contents of a file, such as an image, as the
	// response.
	StaticFiles *StaticFilesHandler `json:"staticFiles,omitempty"`

	// UrlRegex: A URL prefix. This value uses regular expression syntax
	// (and so regexp special characters must be escaped), but it should not
	// contain groupings. All URLs that begin with this prefix are handled
	// by this handler, using the portion of the URL after the prefix as
	// part of the file path. This is always required.
	UrlRegex string `json:"urlRegex,omitempty"`
}

// Version: A Version is a specific set of source code and configuration
// files deployed to a module.
type Version struct {
	// ApiConfig: Serving configuration for Google Cloud Endpoints. Only
	// returned in `GET` requests if `view=FULL` is set. May only be set on
	// create requests; once created, is immutable.
	ApiConfig *ApiConfigHandler `json:"apiConfig,omitempty"`

	// AutomaticScaling: Automatic scaling is the scaling policy that App
	// Engine has used since its inception. It is based on request rate,
	// response latencies, and other application metrics.
	AutomaticScaling *AutomaticScaling `json:"automaticScaling,omitempty"`

	// BasicScaling: A module with basic scaling will create an instance
	// when the application receives a request. The instance will be turned
	// down when the app becomes idle. Basic scaling is ideal for work that
	// is intermittent or driven by user activity.
	BasicScaling *BasicScaling `json:"basicScaling,omitempty"`

	// BetaSettings: Beta settings supplied to the application via metadata.
	BetaSettings map[string]string `json:"betaSettings,omitempty"`

	// DefaultExpiration: The length of time a static file served by a
	// static file handler ought to be cached by web proxies and browsers,
	// if the handler does not specify its own expiration. Only returned in
	// `GET` requests if `view=FULL` is set. May only be set on create
	// requests; once created, is immutable.
	DefaultExpiration string `json:"defaultExpiration,omitempty"`

	// Deployment: Code and application artifacts that make up this version.
	// Only returned in `GET` requests if `view=FULL` is set. May only be
	// set on create requests; once created, is immutable.
	Deployment *Deployment `json:"deployment,omitempty"`

	// EnvVariables: Environment variables made available to the
	// application. Only returned in `GET` requests if `view=FULL` is set.
	// May only be set on create requests; once created, is immutable.
	EnvVariables map[string]string `json:"envVariables,omitempty"`

	// ErrorHandlers: Custom static error pages instead of these generic
	// error pages, (limit 10 KB/page) Only returned in `GET` requests if
	// `view=FULL` is set. May only be set on create requests; once created,
	// is immutable.
	ErrorHandlers []*ErrorHandler `json:"errorHandlers,omitempty"`

	// Handlers: An ordered list of URL Matching patterns that should be
	// applied to incoming requests. The first matching URL consumes the
	// request, and subsequent handlers are not attempted. Only returned in
	// `GET` requests if `view=FULL` is set. May only be set on create
	// requests; once created, is immutable.
	Handlers []*UrlMap `json:"handlers,omitempty"`

	// HealthCheck: Configure health checking for the VM instances.
	// Unhealthy VM instances will be stopped and replaced with new
	// instances. Only returned in `GET` requests if `view=FULL` is set. May
	// only be set on create requests; once created, is immutable.
	HealthCheck *HealthCheck `json:"healthCheck,omitempty"`

	// Id: The relative name/path of the Version within the module. Example:
	// "v1"
	Id string `json:"id,omitempty"`

	// InboundServices: Before an application can receive email or XMPP
	// messages, the application must be configured to enable the service.
	//
	// Possible values:
	//   "INBOUND_SERVICE_UNSPECIFIED" - Not specified.
	//   "INBOUND_SERVICE_MAIL" - Allows an application to receive mail.
	//   "INBOUND_SERVICE_MAIL_BOUNCE" - Allows an application receive email
	// bound notifications.
	//   "INBOUND_SERVICE_XMPP_ERROR" - Allows an application to receive
	// error stanzas.
	//   "INBOUND_SERVICE_XMPP_MESSAGE" - Allows an application to receive
	// instant messages.
	//   "INBOUND_SERVICE_XMPP_SUBSCRIBE" - Allows an application to receive
	// user subscription POSTs.
	//   "INBOUND_SERVICE_XMPP_PRESENCE" - Allows an application to receive
	// a user's chat presence.
	//   "INBOUND_SERVICE_CHANNEL_PRESENCE" - Registers an application for
	// notifications when a client connects or disconnects from a channel.
	//   "INBOUND_SERVICE_WARMUP" - Enables warmup requests.
	InboundServices []string `json:"inboundServices,omitempty"`

	// InstanceClass: The frontend instance class to use to run this app.
	// Valid values are `[F1, F2, F4, F4_1G]`.
	InstanceClass string `json:"instanceClass,omitempty"`

	// Libraries: Configuration for Python runtime third-party libraries
	// required by the application. Only returned in `GET` requests if
	// `view=FULL` is set. May only be set on create requests; once created,
	// is immutable.
	Libraries []*Library `json:"libraries,omitempty"`

	// ManualScaling: A module with manual scaling runs continuously,
	// allowing you to perform complex initialization and rely on the state
	// of its memory over time.
	ManualScaling *ManualScaling `json:"manualScaling,omitempty"`

	// Name: The full path to the Version resource in the API. Example:
	// "apps/myapp/modules/default/versions/v1". @OutputOnly
	Name string `json:"name,omitempty"`

	// Network: Used to specify extra network settings (for VM runtimes
	// only).
	Network *Network `json:"network,omitempty"`

	// NobuildFilesRegex: Go only. Files that match this pattern will not be
	// built into the app. May only be set on create requests.
	NobuildFilesRegex string `json:"nobuildFilesRegex,omitempty"`

	// Resources: Used to specify how many machine resources an app version
	// needs (for VM runtimes only).
	Resources *Resources `json:"resources,omitempty"`

	// Runtime: The desired runtime. Values can include python27, java7, go,
	// etc.
	Runtime string `json:"runtime,omitempty"`

	// Threadsafe: If true, multiple requests can be dispatched to the app
	// at once.
	Threadsafe bool `json:"threadsafe,omitempty"`

	// Vm: Whether to deploy this app in a VM container.
	Vm bool `json:"vm,omitempty"`
}

// method id "appengine.apps.get":

type AppsGetCall struct {
	s      *Service
	appsId string
	opt_   map[string]interface{}
}

// Get: Gets information about an application.
func (r *AppsService) Get(appsId string) *AppsGetCall {
	c := &AppsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.appsId = appsId
	return c
}

// EnsureResourcesExist sets the optional parameter
// "ensureResourcesExist": Certain resources associated with an
// application are created on-demand. Controls whether these resources
// should be created when performing the `GET` operation. If specified
// and any resources cloud not be created, the request will fail with an
// error code.
func (c *AppsGetCall) EnsureResourcesExist(ensureResourcesExist bool) *AppsGetCall {
	c.opt_["ensureResourcesExist"] = ensureResourcesExist
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AppsGetCall) Fields(s ...googleapi.Field) *AppsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *AppsGetCall) Do() (*Application, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["ensureResourcesExist"]; ok {
		params.Set("ensureResourcesExist", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta4/apps/{appsId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"appsId": c.appsId,
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
	var ret *Application
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets information about an application.",
	//   "httpMethod": "GET",
	//   "id": "appengine.apps.get",
	//   "parameterOrder": [
	//     "appsId"
	//   ],
	//   "parameters": {
	//     "appsId": {
	//       "description": "Part of `name`. Name of the application to get. For example: \"apps/myapp\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "ensureResourcesExist": {
	//       "description": "Certain resources associated with an application are created on-demand. Controls whether these resources should be created when performing the `GET` operation. If specified and any resources cloud not be created, the request will fail with an error code.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "v1beta4/apps/{appsId}",
	//   "response": {
	//     "$ref": "Application"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "appengine.apps.modules.delete":

type AppsModulesDeleteCall struct {
	s         *Service
	appsId    string
	modulesId string
	opt_      map[string]interface{}
}

// Delete: Deletes a module and all enclosed versions.
func (r *AppsModulesService) Delete(appsId string, modulesId string) *AppsModulesDeleteCall {
	c := &AppsModulesDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.appsId = appsId
	c.modulesId = modulesId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AppsModulesDeleteCall) Fields(s ...googleapi.Field) *AppsModulesDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *AppsModulesDeleteCall) Do() (*Operation, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta4/apps/{appsId}/modules/{modulesId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"appsId":    c.appsId,
		"modulesId": c.modulesId,
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
	//   "description": "Deletes a module and all enclosed versions.",
	//   "httpMethod": "DELETE",
	//   "id": "appengine.apps.modules.delete",
	//   "parameterOrder": [
	//     "appsId",
	//     "modulesId"
	//   ],
	//   "parameters": {
	//     "appsId": {
	//       "description": "Part of `name`. Name of the resource requested. For example: \"apps/myapp/modules/default\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "modulesId": {
	//       "description": "Part of `name`. See documentation of `appsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta4/apps/{appsId}/modules/{modulesId}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "appengine.apps.modules.get":

type AppsModulesGetCall struct {
	s         *Service
	appsId    string
	modulesId string
	opt_      map[string]interface{}
}

// Get: Gets the current configuration of the module.
func (r *AppsModulesService) Get(appsId string, modulesId string) *AppsModulesGetCall {
	c := &AppsModulesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.appsId = appsId
	c.modulesId = modulesId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AppsModulesGetCall) Fields(s ...googleapi.Field) *AppsModulesGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *AppsModulesGetCall) Do() (*Module, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta4/apps/{appsId}/modules/{modulesId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"appsId":    c.appsId,
		"modulesId": c.modulesId,
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
	var ret *Module
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the current configuration of the module.",
	//   "httpMethod": "GET",
	//   "id": "appengine.apps.modules.get",
	//   "parameterOrder": [
	//     "appsId",
	//     "modulesId"
	//   ],
	//   "parameters": {
	//     "appsId": {
	//       "description": "Part of `name`. Name of the resource requested. For example: \"/apps/myapp/modules/default\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "modulesId": {
	//       "description": "Part of `name`. See documentation of `appsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta4/apps/{appsId}/modules/{modulesId}",
	//   "response": {
	//     "$ref": "Module"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "appengine.apps.modules.list":

type AppsModulesListCall struct {
	s      *Service
	appsId string
	opt_   map[string]interface{}
}

// List: Lists all the modules in the application.
func (r *AppsModulesService) List(appsId string) *AppsModulesListCall {
	c := &AppsModulesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.appsId = appsId
	return c
}

// PageSize sets the optional parameter "pageSize": Maximum results to
// return per page.
func (c *AppsModulesListCall) PageSize(pageSize int64) *AppsModulesListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": Continuation token
// for fetching the next page of results.
func (c *AppsModulesListCall) PageToken(pageToken string) *AppsModulesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AppsModulesListCall) Fields(s ...googleapi.Field) *AppsModulesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *AppsModulesListCall) Do() (*ListModulesResponse, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta4/apps/{appsId}/modules")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"appsId": c.appsId,
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
	var ret *ListModulesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all the modules in the application.",
	//   "httpMethod": "GET",
	//   "id": "appengine.apps.modules.list",
	//   "parameterOrder": [
	//     "appsId"
	//   ],
	//   "parameters": {
	//     "appsId": {
	//       "description": "Part of `name`. Name of the resource requested. For example: \"/apps/myapp\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Maximum results to return per page.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Continuation token for fetching the next page of results.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta4/apps/{appsId}/modules",
	//   "response": {
	//     "$ref": "ListModulesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "appengine.apps.modules.patch":

type AppsModulesPatchCall struct {
	s         *Service
	appsId    string
	modulesId string
	module    *Module
	opt_      map[string]interface{}
}

// Patch: Updates the configuration of the specified module.
func (r *AppsModulesService) Patch(appsId string, modulesId string, module *Module) *AppsModulesPatchCall {
	c := &AppsModulesPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.appsId = appsId
	c.modulesId = modulesId
	c.module = module
	return c
}

// Mask sets the optional parameter "mask": Standard field mask for the
// set of fields to be updated.
func (c *AppsModulesPatchCall) Mask(mask string) *AppsModulesPatchCall {
	c.opt_["mask"] = mask
	return c
}

// MigrateTraffic sets the optional parameter "migrateTraffic": Whether
// to use Traffic Migration to shift traffic gradually. Traffic can only
// be migrated from a single version to another single version.
func (c *AppsModulesPatchCall) MigrateTraffic(migrateTraffic bool) *AppsModulesPatchCall {
	c.opt_["migrateTraffic"] = migrateTraffic
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AppsModulesPatchCall) Fields(s ...googleapi.Field) *AppsModulesPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *AppsModulesPatchCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.module)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["mask"]; ok {
		params.Set("mask", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["migrateTraffic"]; ok {
		params.Set("migrateTraffic", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta4/apps/{appsId}/modules/{modulesId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"appsId":    c.appsId,
		"modulesId": c.modulesId,
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
	//   "description": "Updates the configuration of the specified module.",
	//   "httpMethod": "PATCH",
	//   "id": "appengine.apps.modules.patch",
	//   "parameterOrder": [
	//     "appsId",
	//     "modulesId"
	//   ],
	//   "parameters": {
	//     "appsId": {
	//       "description": "Part of `name`. Name of the resource to update. For example: \"apps/myapp/modules/default\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "mask": {
	//       "description": "Standard field mask for the set of fields to be updated.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "migrateTraffic": {
	//       "description": "Whether to use Traffic Migration to shift traffic gradually. Traffic can only be migrated from a single version to another single version.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "modulesId": {
	//       "description": "Part of `name`. See documentation of `appsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta4/apps/{appsId}/modules/{modulesId}",
	//   "request": {
	//     "$ref": "Module"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "appengine.apps.modules.versions.create":

type AppsModulesVersionsCreateCall struct {
	s         *Service
	appsId    string
	modulesId string
	version   *Version
	opt_      map[string]interface{}
}

// Create: Deploys new code and resource files to a version.
func (r *AppsModulesVersionsService) Create(appsId string, modulesId string, version *Version) *AppsModulesVersionsCreateCall {
	c := &AppsModulesVersionsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.appsId = appsId
	c.modulesId = modulesId
	c.version = version
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AppsModulesVersionsCreateCall) Fields(s ...googleapi.Field) *AppsModulesVersionsCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *AppsModulesVersionsCreateCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.version)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta4/apps/{appsId}/modules/{modulesId}/versions")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"appsId":    c.appsId,
		"modulesId": c.modulesId,
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
	//   "description": "Deploys new code and resource files to a version.",
	//   "httpMethod": "POST",
	//   "id": "appengine.apps.modules.versions.create",
	//   "parameterOrder": [
	//     "appsId",
	//     "modulesId"
	//   ],
	//   "parameters": {
	//     "appsId": {
	//       "description": "Part of `name`. Name of the resource to update. For example: \"apps/myapp/modules/default\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "modulesId": {
	//       "description": "Part of `name`. See documentation of `appsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta4/apps/{appsId}/modules/{modulesId}/versions",
	//   "request": {
	//     "$ref": "Version"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "appengine.apps.modules.versions.delete":

type AppsModulesVersionsDeleteCall struct {
	s          *Service
	appsId     string
	modulesId  string
	versionsId string
	opt_       map[string]interface{}
}

// Delete: Deletes an existing version.
func (r *AppsModulesVersionsService) Delete(appsId string, modulesId string, versionsId string) *AppsModulesVersionsDeleteCall {
	c := &AppsModulesVersionsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.appsId = appsId
	c.modulesId = modulesId
	c.versionsId = versionsId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AppsModulesVersionsDeleteCall) Fields(s ...googleapi.Field) *AppsModulesVersionsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *AppsModulesVersionsDeleteCall) Do() (*Operation, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta4/apps/{appsId}/modules/{modulesId}/versions/{versionsId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"appsId":     c.appsId,
		"modulesId":  c.modulesId,
		"versionsId": c.versionsId,
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
	//   "description": "Deletes an existing version.",
	//   "httpMethod": "DELETE",
	//   "id": "appengine.apps.modules.versions.delete",
	//   "parameterOrder": [
	//     "appsId",
	//     "modulesId",
	//     "versionsId"
	//   ],
	//   "parameters": {
	//     "appsId": {
	//       "description": "Part of `name`. Name of the resource requested. For example: \"apps/myapp/modules/default/versions/v1\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "modulesId": {
	//       "description": "Part of `name`. See documentation of `appsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "versionsId": {
	//       "description": "Part of `name`. See documentation of `appsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta4/apps/{appsId}/modules/{modulesId}/versions/{versionsId}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "appengine.apps.modules.versions.get":

type AppsModulesVersionsGetCall struct {
	s          *Service
	appsId     string
	modulesId  string
	versionsId string
	opt_       map[string]interface{}
}

// Get: Gets application deployment information.
func (r *AppsModulesVersionsService) Get(appsId string, modulesId string, versionsId string) *AppsModulesVersionsGetCall {
	c := &AppsModulesVersionsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.appsId = appsId
	c.modulesId = modulesId
	c.versionsId = versionsId
	return c
}

// View sets the optional parameter "view": Controls the set of fields
// returned in the `Get` response.
//
// Possible values:
//   "BASIC"
//   "FULL"
func (c *AppsModulesVersionsGetCall) View(view string) *AppsModulesVersionsGetCall {
	c.opt_["view"] = view
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AppsModulesVersionsGetCall) Fields(s ...googleapi.Field) *AppsModulesVersionsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *AppsModulesVersionsGetCall) Do() (*Version, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["view"]; ok {
		params.Set("view", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta4/apps/{appsId}/modules/{modulesId}/versions/{versionsId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"appsId":     c.appsId,
		"modulesId":  c.modulesId,
		"versionsId": c.versionsId,
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
	var ret *Version
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets application deployment information.",
	//   "httpMethod": "GET",
	//   "id": "appengine.apps.modules.versions.get",
	//   "parameterOrder": [
	//     "appsId",
	//     "modulesId",
	//     "versionsId"
	//   ],
	//   "parameters": {
	//     "appsId": {
	//       "description": "Part of `name`. Name of the resource requested. For example: \"apps/myapp/modules/default/versions/v1\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "modulesId": {
	//       "description": "Part of `name`. See documentation of `appsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "versionsId": {
	//       "description": "Part of `name`. See documentation of `appsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "view": {
	//       "description": "Controls the set of fields returned in the `Get` response.",
	//       "enum": [
	//         "BASIC",
	//         "FULL"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta4/apps/{appsId}/modules/{modulesId}/versions/{versionsId}",
	//   "response": {
	//     "$ref": "Version"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "appengine.apps.modules.versions.list":

type AppsModulesVersionsListCall struct {
	s         *Service
	appsId    string
	modulesId string
	opt_      map[string]interface{}
}

// List: Lists the versions of a module.
func (r *AppsModulesVersionsService) List(appsId string, modulesId string) *AppsModulesVersionsListCall {
	c := &AppsModulesVersionsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.appsId = appsId
	c.modulesId = modulesId
	return c
}

// PageSize sets the optional parameter "pageSize": Maximum results to
// return per page.
func (c *AppsModulesVersionsListCall) PageSize(pageSize int64) *AppsModulesVersionsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": Continuation token
// for fetching the next page of results.
func (c *AppsModulesVersionsListCall) PageToken(pageToken string) *AppsModulesVersionsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// View sets the optional parameter "view": Controls the set of fields
// returned in the `List` response.
//
// Possible values:
//   "BASIC"
//   "FULL"
func (c *AppsModulesVersionsListCall) View(view string) *AppsModulesVersionsListCall {
	c.opt_["view"] = view
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AppsModulesVersionsListCall) Fields(s ...googleapi.Field) *AppsModulesVersionsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *AppsModulesVersionsListCall) Do() (*ListVersionsResponse, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta4/apps/{appsId}/modules/{modulesId}/versions")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"appsId":    c.appsId,
		"modulesId": c.modulesId,
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
	var ret *ListVersionsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the versions of a module.",
	//   "httpMethod": "GET",
	//   "id": "appengine.apps.modules.versions.list",
	//   "parameterOrder": [
	//     "appsId",
	//     "modulesId"
	//   ],
	//   "parameters": {
	//     "appsId": {
	//       "description": "Part of `name`. Name of the resource requested. For example: \"apps/myapp/modules/default\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "modulesId": {
	//       "description": "Part of `name`. See documentation of `appsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Maximum results to return per page.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Continuation token for fetching the next page of results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "view": {
	//       "description": "Controls the set of fields returned in the `List` response.",
	//       "enum": [
	//         "BASIC",
	//         "FULL"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta4/apps/{appsId}/modules/{modulesId}/versions",
	//   "response": {
	//     "$ref": "ListVersionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "appengine.apps.operations.get":

type AppsOperationsGetCall struct {
	s            *Service
	appsId       string
	operationsId string
	opt_         map[string]interface{}
}

// Get: Gets the latest state of a long-running operation. Clients can
// use this method to poll the operation result at intervals as
// recommended by the API service.
func (r *AppsOperationsService) Get(appsId string, operationsId string) *AppsOperationsGetCall {
	c := &AppsOperationsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.appsId = appsId
	c.operationsId = operationsId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AppsOperationsGetCall) Fields(s ...googleapi.Field) *AppsOperationsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *AppsOperationsGetCall) Do() (*Operation, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta4/apps/{appsId}/operations/{operationsId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"appsId":       c.appsId,
		"operationsId": c.operationsId,
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
	//   "description": "Gets the latest state of a long-running operation. Clients can use this method to poll the operation result at intervals as recommended by the API service.",
	//   "httpMethod": "GET",
	//   "id": "appengine.apps.operations.get",
	//   "parameterOrder": [
	//     "appsId",
	//     "operationsId"
	//   ],
	//   "parameters": {
	//     "appsId": {
	//       "description": "Part of `name`. The name of the operation resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "operationsId": {
	//       "description": "Part of `name`. See documentation of `appsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta4/apps/{appsId}/operations/{operationsId}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "appengine.apps.operations.list":

type AppsOperationsListCall struct {
	s      *Service
	appsId string
	opt_   map[string]interface{}
}

// List: Lists operations that match the specified filter in the
// request. If the server doesn't support this method, it returns
// `UNIMPLEMENTED`. NOTE: the `name` binding below allows API services
// to override the binding to use different resource name schemes, such
// as `users/*/operations`.
func (r *AppsOperationsService) List(appsId string) *AppsOperationsListCall {
	c := &AppsOperationsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.appsId = appsId
	return c
}

// Filter sets the optional parameter "filter": The standard List
// filter.
func (c *AppsOperationsListCall) Filter(filter string) *AppsOperationsListCall {
	c.opt_["filter"] = filter
	return c
}

// PageSize sets the optional parameter "pageSize": The standard List
// page size.
func (c *AppsOperationsListCall) PageSize(pageSize int64) *AppsOperationsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": The standard List
// page token.
func (c *AppsOperationsListCall) PageToken(pageToken string) *AppsOperationsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AppsOperationsListCall) Fields(s ...googleapi.Field) *AppsOperationsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *AppsOperationsListCall) Do() (*ListOperationsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["filter"]; ok {
		params.Set("filter", fmt.Sprintf("%v", v))
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta4/apps/{appsId}/operations")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"appsId": c.appsId,
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
	var ret *ListOperationsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists operations that match the specified filter in the request. If the server doesn't support this method, it returns `UNIMPLEMENTED`. NOTE: the `name` binding below allows API services to override the binding to use different resource name schemes, such as `users/*/operations`.",
	//   "httpMethod": "GET",
	//   "id": "appengine.apps.operations.list",
	//   "parameterOrder": [
	//     "appsId"
	//   ],
	//   "parameters": {
	//     "appsId": {
	//       "description": "Part of `name`. The name of the operation collection.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "filter": {
	//       "description": "The standard List filter.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The standard List page size.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The standard List page token.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta4/apps/{appsId}/operations",
	//   "response": {
	//     "$ref": "ListOperationsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
