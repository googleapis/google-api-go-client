// Package tracing provides access to the Google Tracing API.
//
// See https://cloud.google.com/trace
//
// Usage example:
//
//   import "google.golang.org/api/tracing/v1"
//   ...
//   tracingService, err := tracing.New(oauthHttpClient)
package tracing // import "google.golang.org/api/tracing/v1"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
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
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "tracing:v1"
const apiName = "tracing"
const apiVersion = "v1"
const basePath = "https://tracing.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// Write Trace data for a project or application
	TraceAppendScope = "https://www.googleapis.com/auth/trace.append"

	// Read Trace data for a project or application
	TraceReadonlyScope = "https://www.googleapis.com/auth/trace.readonly"
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
	client                    *http.Client
	BasePath                  string // API endpoint base URL
	UserAgent                 string // optional additional User-Agent fragment
	GoogleClientHeaderElement string // client header fragment, for Google use only

	Projects *ProjectsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func (s *Service) clientHeader() string {
	return gensupport.GoogleClientHeader("20170210", s.GoogleClientHeaderElement)
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Traces = NewProjectsTracesService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Traces *ProjectsTracesService
}

func NewProjectsTracesService(s *Service) *ProjectsTracesService {
	rs := &ProjectsTracesService{s: s}
	return rs
}

type ProjectsTracesService struct {
	s *Service
}

// Annotation: Text annotation with a set of attributes.
type Annotation struct {
	// Attributes: A set of attributes on the annotation.
	Attributes map[string]AttributeValue `json:"attributes,omitempty"`

	// Description: A user-supplied message describing the event.
	Description string `json:"description,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Attributes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Attributes") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Annotation) MarshalJSON() ([]byte, error) {
	type noMethod Annotation
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AttributeValue: Allowed attribute values.
type AttributeValue struct {
	// BoolValue: A boolean value.
	BoolValue bool `json:"boolValue,omitempty"`

	// IntValue: An integer value.
	IntValue int64 `json:"intValue,omitempty,string"`

	// StringValue: A string value.
	StringValue string `json:"stringValue,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BoolValue") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BoolValue") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AttributeValue) MarshalJSON() ([]byte, error) {
	type noMethod AttributeValue
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BatchUpdateSpansRequest: The request message for the
// `BatchUpdateSpans` method.
type BatchUpdateSpansRequest struct {
	// SpanUpdates: A map from trace name to spans to be stored or updated.
	SpanUpdates map[string]SpanUpdates `json:"spanUpdates,omitempty"`

	// ForceSendFields is a list of field names (e.g. "SpanUpdates") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "SpanUpdates") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BatchUpdateSpansRequest) MarshalJSON() ([]byte, error) {
	type noMethod BatchUpdateSpansRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated
// empty messages in your APIs. A typical example is to use it as the
// request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// Link: Link one span with another which may be in a different Trace.
// Used (for
// example) in batching operations, where a single batch handler
// processes
// multiple requests from different traces.
type Link struct {
	// SpanId: The span identifier of the linked span.
	SpanId uint64 `json:"spanId,omitempty,string"`

	// TraceId: The trace identifier of the linked span.
	TraceId string `json:"traceId,omitempty"`

	// Type: The type of the link.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED"
	//   "CHILD"
	//   "PARENT"
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "SpanId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "SpanId") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Link) MarshalJSON() ([]byte, error) {
	type noMethod Link
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListSpansResponse: The response message for the 'ListSpans' method.
type ListSpansResponse struct {
	// NextPageToken: If defined, indicates that there are more spans that
	// match the request
	// and that this value should be passed to the next request to
	// continue
	// retrieving additional spans.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Spans: The requested spans if they are any in the specified trace.
	Spans []*Span `json:"spans,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListSpansResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListSpansResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListTracesResponse: The response message for the `ListTraces` method.
type ListTracesResponse struct {
	// NextPageToken: If defined, indicates that there are more traces that
	// match the request
	// and that this value should be passed to the next request to
	// continue
	// retrieving additional traces.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Traces: List of trace records returned.
	Traces []*Trace `json:"traces,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListTracesResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListTracesResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type Module struct {
	// BuildId: Build_id is a unique identifier for the module,
	// usually a hash of its contents
	BuildId string `json:"buildId,omitempty"`

	// Module: Binary module.
	// E.g. main binary, kernel modules, and dynamic libraries
	// such as libc.so, sharedlib.so
	Module string `json:"module,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BuildId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BuildId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Module) MarshalJSON() ([]byte, error) {
	type noMethod Module
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// NetworkEvent: An event describing an RPC message sent/received on the
// network.
type NetworkEvent struct {
	// KernelTime: If available, this is the kernel time:
	// For sent messages, this is the time at which the first bit was
	// sent.
	// For received messages, this is the time at which the last bit
	// was
	// received.
	KernelTime string `json:"kernelTime,omitempty"`

	// MessageId: Every message has an identifier, which must be different
	// from all the
	// network messages in this span.
	// This is especially important when the request/response are streamed.
	MessageId uint64 `json:"messageId,omitempty,string"`

	// MessageSize: Number of bytes send/receive.
	MessageSize uint64 `json:"messageSize,omitempty,string"`

	// Possible values:
	//   "TYPE_UNSPECIFIED"
	//   "SENT"
	//   "RECV"
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "KernelTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "KernelTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *NetworkEvent) MarshalJSON() ([]byte, error) {
	type noMethod NetworkEvent
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Span: A span represents a single operation within a trace. Spans can
// be nested
// and form a trace tree. Often, a trace contains a root span that
// describes the
// end-to-end latency and, optionally, one or more subspans for
// its sub-operations. Spans do not need to be contiguous. There may be
// gaps
// between spans in a trace.
type Span struct {
	// Attributes: Properties of a span. Attributes at the span
	// level.
	// E.g.
	// "/instance_id": "my-instance"
	// "/zone": "us-central1-a"
	// "/grpc/peer_address": "ip:port" (dns, etc.)
	// "/grpc/deadline":
	// "Duration"
	// "/http/user_agent"
	// "/http/request_bytes": 300
	// "/http/response_bytes": 1200
	// "/http/url": google.com/apis
	// "/pid"
	// "abc.com/myattribute": "my attribute value"
	//
	// Maximum length for attribute key is 128 characters, for string
	// attribute
	// value is 2K characters.
	Attributes map[string]AttributeValue `json:"attributes,omitempty"`

	// HasRemoteParent: True if this Span has a remote parent (is an RPC
	// server Span).
	HasRemoteParent bool `json:"hasRemoteParent,omitempty"`

	// Id: Identifier for the span. Must be a 64-bit integer other than 0
	// and
	// unique within a trace.
	Id uint64 `json:"id,omitempty,string"`

	// Links: A collection of links.
	Links []*Link `json:"links,omitempty"`

	// LocalEndTime: Local machine clock time from the UNIX epoch,
	// at which span execution ended.
	// On the server side these are the times when the server
	// application
	// handler finishes running.
	LocalEndTime string `json:"localEndTime,omitempty"`

	// LocalStartTime: Local machine clock time from the UNIX epoch,
	// at which span execution started.
	// On the server side these are the times when the server
	// application
	// handler starts running.
	LocalStartTime string `json:"localStartTime,omitempty"`

	// Name: Name of the span. The span name is sanitized and displayed in
	// the
	// Stackdriver Trace tool in the {% dynamic print
	// site_values.console_name %}.
	// The name may be a method name or some other per-call site name.
	// For the same executable and the same call point, a best practice
	// is
	// to use a consistent name, which makes it easier to
	// correlate
	// cross-trace spans.
	Name string `json:"name,omitempty"`

	// ParentId: ID of parent span. 0 or missing if this is a root span.
	ParentId uint64 `json:"parentId,omitempty,string"`

	// StackTrace: Stack trace captured at the start of the span. This is
	// optional.
	StackTrace *StackTrace `json:"stackTrace,omitempty"`

	// Status: The final status of the Span. This is optional.
	Status *Status `json:"status,omitempty"`

	// TimeEvents: A collection of time-stamped events.
	TimeEvents []*TimeEvent `json:"timeEvents,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Attributes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Attributes") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Span) MarshalJSON() ([]byte, error) {
	type noMethod Span
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SpanUpdates: Collection of spans.
type SpanUpdates struct {
	Spans []*Span `json:"spans,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Spans") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Spans") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SpanUpdates) MarshalJSON() ([]byte, error) {
	type noMethod SpanUpdates
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// StackFrame: Presents a single stack frame in a stack trace.
type StackFrame struct {
	// ColumnNumber: Column number is important in JavaScript(anonymous
	// functions),
	// Might not be available in some languages.
	ColumnNumber int64 `json:"columnNumber,omitempty,string"`

	// FileName: File name of the frame.
	FileName string `json:"fileName,omitempty"`

	// FunctionName: Fully qualified names which uniquely identify
	// function/method/etc.
	FunctionName string `json:"functionName,omitempty"`

	// LineNumber: Line number of the frame.
	LineNumber int64 `json:"lineNumber,omitempty,string"`

	// LoadModule: Binary module the code is loaded from.
	LoadModule *Module `json:"loadModule,omitempty"`

	// OriginalFunctionName: Used when function name is ‘mangled’. Not
	// guaranteed to be fully
	// qualified but usually it is.
	OriginalFunctionName string `json:"originalFunctionName,omitempty"`

	// SourceVersion: source_version is deployment specific. It might
	// be
	// better to be stored in deployment metadata.
	SourceVersion string `json:"sourceVersion,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ColumnNumber") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ColumnNumber") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *StackFrame) MarshalJSON() ([]byte, error) {
	type noMethod StackFrame
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type StackTrace struct {
	// StackFrame: Stack frames of this stack trace.
	StackFrame []*StackFrame `json:"stackFrame,omitempty"`

	// StackTraceHashId: User can choose to use their own hash function to
	// hash large attributes to
	// save network bandwidth and storage.
	// Typical usage is to pass both stack_frame and stack_trace_hash_id
	// initially
	// to inform the storage of the mapping. And in subsequent calls, pass
	// in
	// stack_trace_hash_id only. User shall verify the hash value
	// is
	// successfully stored.
	StackTraceHashId uint64 `json:"stackTraceHashId,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "StackFrame") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "StackFrame") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *StackTrace) MarshalJSON() ([]byte, error) {
	type noMethod StackTrace
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Status: The `Status` type defines a logical error model that is
// suitable for different
// programming environments, including REST APIs and RPC APIs. It is
// used by
// [gRPC](https://github.com/grpc). The error model is designed to
// be:
//
// - Simple to use and understand for most users
// - Flexible enough to meet unexpected needs
//
// # Overview
//
// The `Status` message contains three pieces of data: error code, error
// message,
// and error details. The error code should be an enum value
// of
// google.rpc.Code, but it may accept additional error codes if needed.
// The
// error message should be a developer-facing English message that
// helps
// developers *understand* and *resolve* the error. If a localized
// user-facing
// error message is needed, put the localized message in the error
// details or
// localize it in the client. The optional error details may contain
// arbitrary
// information about the error. There is a predefined set of error
// detail types
// in the package `google.rpc` which can be used for common error
// conditions.
//
// # Language mapping
//
// The `Status` message is the logical representation of the error
// model, but it
// is not necessarily the actual wire format. When the `Status` message
// is
// exposed in different client libraries and different wire protocols,
// it can be
// mapped differently. For example, it will likely be mapped to some
// exceptions
// in Java, but more likely mapped to some error codes in C.
//
// # Other uses
//
// The error model and the `Status` message can be used in a variety
// of
// environments, either with or without APIs, to provide a
// consistent developer experience across different
// environments.
//
// Example uses of this error model include:
//
// - Partial errors. If a service needs to return partial errors to the
// client,
//     it may embed the `Status` in the normal response to indicate the
// partial
//     errors.
//
// - Workflow errors. A typical workflow has multiple steps. Each step
// may
//     have a `Status` message for error reporting purpose.
//
// - Batch operations. If a client uses batch request and batch
// response, the
//     `Status` message should be used directly inside batch response,
// one for
//     each error sub-response.
//
// - Asynchronous operations. If an API call embeds asynchronous
// operation
//     results in its response, the status of those operations should
// be
//     represented directly using the `Status` message.
//
// - Logging. If some API errors are stored in logs, the message
// `Status` could
//     be used directly after any stripping needed for security/privacy
// reasons.
type Status struct {
	// Code: The status code, which should be an enum value of
	// google.rpc.Code.
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details.  There will
	// be a
	// common set of message types for APIs to use.
	Details []googleapi.RawMessage `json:"details,omitempty"`

	// Message: A developer-facing error message, which should be in
	// English. Any
	// user-facing error message should be localized and sent in
	// the
	// google.rpc.Status.details field, or localized by the client.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Status) MarshalJSON() ([]byte, error) {
	type noMethod Status
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TimeEvent: A time-stamped annotation in the Span.
type TimeEvent struct {
	// Annotation: Optional field for user supplied <string, AttributeValue>
	// map
	Annotation *Annotation `json:"annotation,omitempty"`

	// LocalTime: The local machine absolute timestamp when this event
	// happened.
	LocalTime string `json:"localTime,omitempty"`

	// NetworkEvent: Optional field that can be used only for network
	// events.
	NetworkEvent *NetworkEvent `json:"networkEvent,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Annotation") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Annotation") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TimeEvent) MarshalJSON() ([]byte, error) {
	type noMethod TimeEvent
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Trace: A trace describes how long it takes for an application to
// perform some
// operations. It consists of a set of spans, each of which contains
// details
// about an operation with time information and operation details.
type Trace struct {
	// Name: ID of the trace which is
	// "projects/<project_id>/traces/<trace_id>".
	// trace_id is globally unique identifier for the trace. Common to all
	// the
	// spans. It is conceptually a 128-bit hex-encoded value.
	Name string `json:"name,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Name") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Name") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Trace) MarshalJSON() ([]byte, error) {
	type noMethod Trace
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "tracing.projects.traces.batchUpdate":

type ProjectsTracesBatchUpdateCall struct {
	s                       *Service
	parent                  string
	batchupdatespansrequest *BatchUpdateSpansRequest
	urlParams_              gensupport.URLParams
	ctx_                    context.Context
	header_                 http.Header
}

// BatchUpdate: Sends new spans to Stackdriver Trace or updates existing
// spans. If the
// name of a trace that you send matches that of an existing trace, any
// fields
// in the existing trace and its spans are overwritten by the provided
// values,
// and any new fields provided are merged with the existing trace data.
// If the
// name does not match, a new trace is created with given set of spans.
func (r *ProjectsTracesService) BatchUpdate(parent string, batchupdatespansrequest *BatchUpdateSpansRequest) *ProjectsTracesBatchUpdateCall {
	c := &ProjectsTracesBatchUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.batchupdatespansrequest = batchupdatespansrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTracesBatchUpdateCall) Fields(s ...googleapi.Field) *ProjectsTracesBatchUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsTracesBatchUpdateCall) Context(ctx context.Context) *ProjectsTracesBatchUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsTracesBatchUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsTracesBatchUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	reqHeaders.Set("x-goog-api-client", c.s.clientHeader())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.batchupdatespansrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/traces:batchUpdate")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "tracing.projects.traces.batchUpdate" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsTracesBatchUpdateCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Sends new spans to Stackdriver Trace or updates existing spans. If the\nname of a trace that you send matches that of an existing trace, any fields\nin the existing trace and its spans are overwritten by the provided values,\nand any new fields provided are merged with the existing trace data. If the\nname does not match, a new trace is created with given set of spans.",
	//   "flatPath": "v1/projects/{projectsId}/traces:batchUpdate",
	//   "httpMethod": "PATCH",
	//   "id": "tracing.projects.traces.batchUpdate",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "ID of the Cloud project where the trace data is stored.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/traces:batchUpdate",
	//   "request": {
	//     "$ref": "BatchUpdateSpansRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/trace.append"
	//   ]
	// }

}

// method id "tracing.projects.traces.get":

type ProjectsTracesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Returns a specific trace.
func (r *ProjectsTracesService) Get(name string) *ProjectsTracesGetCall {
	c := &ProjectsTracesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTracesGetCall) Fields(s ...googleapi.Field) *ProjectsTracesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsTracesGetCall) IfNoneMatch(entityTag string) *ProjectsTracesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsTracesGetCall) Context(ctx context.Context) *ProjectsTracesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsTracesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsTracesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	reqHeaders.Set("x-goog-api-client", c.s.clientHeader())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "tracing.projects.traces.get" call.
// Exactly one of *Trace or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Trace.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsTracesGetCall) Do(opts ...googleapi.CallOption) (*Trace, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Trace{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns a specific trace.",
	//   "flatPath": "v1/projects/{projectsId}/traces/{tracesId}",
	//   "httpMethod": "GET",
	//   "id": "tracing.projects.traces.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "ID of the trace which is \"projects/\u003cproject_id\u003e/traces/\u003ctrace_id\u003e\".",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/traces/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "Trace"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/trace.readonly"
	//   ]
	// }

}

// method id "tracing.projects.traces.list":

type ProjectsTracesListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Returns of a list of traces that match the specified filter
// conditions.
func (r *ProjectsTracesService) List(parent string) *ProjectsTracesListCall {
	c := &ProjectsTracesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// EndTime sets the optional parameter "endTime": End of the time
// interval (inclusive) during which the trace data was
// collected from the application.
func (c *ProjectsTracesListCall) EndTime(endTime string) *ProjectsTracesListCall {
	c.urlParams_.Set("endTime", endTime)
	return c
}

// Filter sets the optional parameter "filter": An optional filter for
// the request.
// Example:
// "version_label_key:a some_label:some_label_key"
// returns traces from version a and has some_label with some_label_key.
func (c *ProjectsTracesListCall) Filter(filter string) *ProjectsTracesListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// OrderBy sets the optional parameter "orderBy": Field used to sort the
// returned traces.
// Can be one of the following:
//
// *   `trace_id`
// *   `name` (`name` field of root span in the trace)
// *   `duration` (difference between `end_time` and `start_time` fields
// of
//      the root span)
// *   `start` (`start_time` field of the root span)
//
// Descending order can be specified by appending `desc` to the sort
// field
// (for example, `name desc`).
//
// Only one sort field is permitted.
func (c *ProjectsTracesListCall) OrderBy(orderBy string) *ProjectsTracesListCall {
	c.urlParams_.Set("orderBy", orderBy)
	return c
}

// PageSize sets the optional parameter "pageSize": Maximum number of
// traces to return. If not specified or <= 0, the
// implementation selects a reasonable value.  The implementation
// may
// return fewer traces than the requested page size.
func (c *ProjectsTracesListCall) PageSize(pageSize int64) *ProjectsTracesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": Token identifying
// the page of results to return. If provided, use the
// value of the `next_page_token` field from a previous request.
func (c *ProjectsTracesListCall) PageToken(pageToken string) *ProjectsTracesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// StartTime sets the optional parameter "startTime": Start of the time
// interval (inclusive) during which the trace data was
// collected from the application.
func (c *ProjectsTracesListCall) StartTime(startTime string) *ProjectsTracesListCall {
	c.urlParams_.Set("startTime", startTime)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTracesListCall) Fields(s ...googleapi.Field) *ProjectsTracesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsTracesListCall) IfNoneMatch(entityTag string) *ProjectsTracesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsTracesListCall) Context(ctx context.Context) *ProjectsTracesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsTracesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsTracesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	reqHeaders.Set("x-goog-api-client", c.s.clientHeader())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/traces")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "tracing.projects.traces.list" call.
// Exactly one of *ListTracesResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListTracesResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsTracesListCall) Do(opts ...googleapi.CallOption) (*ListTracesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListTracesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns of a list of traces that match the specified filter conditions.",
	//   "flatPath": "v1/projects/{projectsId}/traces",
	//   "httpMethod": "GET",
	//   "id": "tracing.projects.traces.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "endTime": {
	//       "description": "End of the time interval (inclusive) during which the trace data was\ncollected from the application.",
	//       "format": "google-datetime",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "filter": {
	//       "description": "An optional filter for the request.\nExample:\n\"version_label_key:a some_label:some_label_key\"\nreturns traces from version a and has some_label with some_label_key.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "orderBy": {
	//       "description": "Field used to sort the returned traces. Optional.\nCan be one of the following:\n\n*   `trace_id`\n*   `name` (`name` field of root span in the trace)\n*   `duration` (difference between `end_time` and `start_time` fields of\n     the root span)\n*   `start` (`start_time` field of the root span)\n\nDescending order can be specified by appending `desc` to the sort field\n(for example, `name desc`).\n\nOnly one sort field is permitted.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Maximum number of traces to return. If not specified or \u003c= 0, the\nimplementation selects a reasonable value.  The implementation may\nreturn fewer traces than the requested page size. Optional.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Token identifying the page of results to return. If provided, use the\nvalue of the `next_page_token` field from a previous request. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "ID of the Cloud project where the trace data is stored.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "startTime": {
	//       "description": "Start of the time interval (inclusive) during which the trace data was\ncollected from the application.",
	//       "format": "google-datetime",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/traces",
	//   "response": {
	//     "$ref": "ListTracesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/trace.readonly"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsTracesListCall) Pages(ctx context.Context, f func(*ListTracesResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "tracing.projects.traces.listSpans":

type ProjectsTracesListSpansCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// ListSpans: Returns a list of spans within a trace.
func (r *ProjectsTracesService) ListSpans(name string) *ProjectsTracesListSpansCall {
	c := &ProjectsTracesListSpansCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// PageToken sets the optional parameter "pageToken": Token identifying
// the page of results to return. If provided, use the
// value of the `page_token` field from a previous request.
func (c *ProjectsTracesListSpansCall) PageToken(pageToken string) *ProjectsTracesListSpansCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTracesListSpansCall) Fields(s ...googleapi.Field) *ProjectsTracesListSpansCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsTracesListSpansCall) IfNoneMatch(entityTag string) *ProjectsTracesListSpansCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsTracesListSpansCall) Context(ctx context.Context) *ProjectsTracesListSpansCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsTracesListSpansCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsTracesListSpansCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	reqHeaders.Set("x-goog-api-client", c.s.clientHeader())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}:listSpans")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "tracing.projects.traces.listSpans" call.
// Exactly one of *ListSpansResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListSpansResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsTracesListSpansCall) Do(opts ...googleapi.CallOption) (*ListSpansResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListSpansResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns a list of spans within a trace.",
	//   "flatPath": "v1/projects/{projectsId}/traces/{tracesId}:listSpans",
	//   "httpMethod": "GET",
	//   "id": "tracing.projects.traces.listSpans",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "ID of the span set where is \"projects/\u003cproject_id\u003e/traces/\u003ctrace_id\u003e\".",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/traces/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Token identifying the page of results to return. If provided, use the\nvalue of the `page_token` field from a previous request. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}:listSpans",
	//   "response": {
	//     "$ref": "ListSpansResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/trace.readonly"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsTracesListSpansCall) Pages(ctx context.Context, f func(*ListSpansResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}
