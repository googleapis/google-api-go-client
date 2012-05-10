// Package bigquery provides access to the BigQuery API.
//
// See https://code.google.com/apis/bigquery/docs/v2/
//
// Usage example:
//
//   import "code.google.com/p/google-api-go-client/bigquery/v2"
//   ...
//   bigqueryService, err := bigquery.New(oauthHttpClient)
package bigquery

import (
	"bytes"
	"fmt"
	"net/http"
	"io"
	"encoding/json"
	"errors"
	"strings"
	"strconv"
	"net/url"
	"code.google.com/p/google-api-go-client/googleapi"
)

var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New

const apiId = "bigquery:v2"
const apiName = "bigquery"
const apiVersion = "v2"
const basePath = "https://www.googleapis.com/bigquery/v2/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data in Google BigQuery
	BigqueryScope = "https://www.googleapis.com/auth/bigquery"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Datasets = &DatasetsService{s: s}
	s.Tables = &TablesService{s: s}
	s.Tabledata = &TabledataService{s: s}
	s.Projects = &ProjectsService{s: s}
	s.Jobs = &JobsService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	Datasets *DatasetsService

	Tables *TablesService

	Tabledata *TabledataService

	Projects *ProjectsService

	Jobs *JobsService
}

type DatasetsService struct {
	s *Service
}

type TablesService struct {
	s *Service
}

type TabledataService struct {
	s *Service
}

type ProjectsService struct {
	s *Service
}

type JobsService struct {
	s *Service
}

type Dataset struct {
	// SelfLink: [Output-only] An URL that can be used to access this
	// resource again. You can use this URL in Get or Update requests to
	// this resource.
	SelfLink string `json:"selfLink,omitempty"`

	// FriendlyName: [Optional] A descriptive name for this dataset, which
	// might be shown in any BigQuery user interfaces for browsing the
	// dataset. Use datasetId for making API calls.
	FriendlyName string `json:"friendlyName,omitempty"`

	// CreationTime: [Output-only] The time when this dataset was created,
	// in milliseconds since the epoch.
	CreationTime int64 `json:"creationTime,omitempty,string"`

	// Etag: [Output-only] A hash of this resource.
	Etag string `json:"etag,omitempty"`

	// Kind: [Output-only] The resource type.
	Kind string `json:"kind,omitempty"`

	// Access: [Optional] Describes users' rights on the dataset. You can
	// assign the same role to multiple users, and assign multiple roles to
	// the same user.
	// Default values assigned to a new dataset are as
	// follows: OWNER - Project owners, dataset creator READ - Project
	// readers WRITE - Project writers
	// See ACLs and Rights for a description
	// of these rights. If you specify any of these roles when creating a
	// dataset, the assigned roles will overwrite the defaults listed
	// above.
	// To revoke rights to a dataset, call datasets.update() and omit
	// the names of anyone whose rights you wish to revoke. However, every
	// dataset must have at least one entity granted OWNER role.
	// Each access
	// object can have only one of the following members: userByEmail,
	// groupByEmail, domain, or allAuthenticatedUsers.
	Access []*DatasetAccess `json:"access,omitempty"`

	// Id: [Output-only] The fully-qualified unique name of this dataset in
	// the format projectId:datasetId. The dataset name without the project
	// name is given in the datasetId field. When creating a new dataset,
	// leave this field blank, and instead specify the datasetId field.
	Id string `json:"id,omitempty"`

	// LastModifiedTime: [Output-only] The date when this dataset or any of
	// its tables was last modified, in milliseconds since the epoch.
	LastModifiedTime int64 `json:"lastModifiedTime,omitempty,string"`

	// DatasetReference: [Required] Reference identifying dataset.
	DatasetReference *DatasetReference `json:"datasetReference,omitempty"`

	// Description: [Optional] A user-friendly string description for the
	// dataset. This might be shown in BigQuery UI for browsing the dataset.
	Description string `json:"description,omitempty"`
}

type DatasetAccess struct {
	// UserByEmail: [Pick one] A fully qualified email address of a user to
	// grant access to. For example: fred@example.com.
	UserByEmail string `json:"userByEmail,omitempty"`

	// Role: [Required] Describes the rights granted to the user specified
	// by the other member of the access object. The following string values
	// are supported: READ - User can call any list() or get() method on any
	// collection or resource. WRITE - User can call any method on any
	// collection except for datasets, on which they can call list() and
	// get(). OWNER - User can call any method. The dataset creator is
	// granted this role by default.
	Role string `json:"role,omitempty"`

	// GroupByEmail: [Pick one] A fully-qualified email address of a mailing
	// list to grant access to. This must be either a Google Groups mailing
	// list (ends in @googlegroups.com) or a group managed by an enterprise
	// version of Google Groups.
	GroupByEmail string `json:"groupByEmail,omitempty"`

	// SpecialGroup: [Pick one] A special group to grant access to. The
	// valid values are: projectOwners: Owners of the enclosing project.
	// projectReaders: Readers of the enclosing project. projectWriters:
	// Writers of the enclosing project. allAuthenticatedUsers: All
	// authenticated BigQuery users.
	SpecialGroup string `json:"specialGroup,omitempty"`

	// Domain: [Pick one] A domain to grant access to. Any users signed in
	// with the domain specified will be granted the specified access.
	// Example: "example.com".
	Domain string `json:"domain,omitempty"`
}

type DatasetList struct {
	// Kind: The type of list.
	Kind string `json:"kind,omitempty"`

	// Datasets: An array of one or more summarized dataset resources.
	// Absent when there are no datasets in the specified project.
	Datasets []*DatasetListDatasets `json:"datasets,omitempty"`

	// NextPageToken: A token to request the next page of results. Present
	// only when there is more than one page of results.* See Paging Through
	// Results in the developer's guide.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Etag: A hash of this page of results. See Paging Through Results in
	// the developer's guide.
	Etag string `json:"etag,omitempty"`
}

type DatasetListDatasets struct {
	// DatasetReference: Reference identifying dataset.
	DatasetReference *DatasetReference `json:"datasetReference,omitempty"`

	// FriendlyName: A descriptive name for this dataset, if one exists.
	FriendlyName string `json:"friendlyName,omitempty"`

	// Kind: The resource type.
	Kind string `json:"kind,omitempty"`

	// Id: The fully-qualified unique name of this dataset in the format
	// projectId:datasetId.
	Id string `json:"id,omitempty"`
}

type DatasetReference struct {
	// ProjectId: [Optional] The ID of the container project.
	ProjectId string `json:"projectId,omitempty"`

	// DatasetId: [Required] A unique ID for this dataset, without the
	// project name.
	DatasetId string `json:"datasetId,omitempty"`
}

type ErrorProto struct {
	// Message: A human readable explanation of the error.
	Message string `json:"message,omitempty"`

	// Reason: Specifies the error reason. For example, reason will be
	// "required" or "invalid" if some field was missing or malformed.
	Reason string `json:"reason,omitempty"`

	// Location: Specifies where the error occurred, if present.
	Location string `json:"location,omitempty"`

	// DebugInfo: Debugging information for the service, if present. Should
	// be ignored.
	DebugInfo string `json:"debugInfo,omitempty"`
}

type GetQueryResultsResponse struct {
	// Schema: The schema of the results. Present only when the query
	// completes successfully.
	Schema *TableSchema `json:"schema,omitempty"`

	// Etag: A hash of this response.
	Etag string `json:"etag,omitempty"`

	// Kind: The resource type of the response.
	Kind string `json:"kind,omitempty"`

	// TotalRows: The total number of rows in the complete query result set,
	// which can be more than the number of rows in this single page of
	// results. Present only when the query completes successfully.
	TotalRows uint64 `json:"totalRows,omitempty,string"`

	// Rows: An object with as many results as can be contained within the
	// maximum permitted reply size. To get any additional rows, you can
	// call GetQueryResults and specify the jobReference returned above.
	// Present only when the query completes successfully.
	Rows []*TableRow `json:"rows,omitempty"`

	// JobReference: Reference to the Helix Job that was created to run the
	// query. This field will be present even if the original request timed
	// out, in which case GetQueryResults can be used to read the results
	// once the query has completed. Since this API only returns the first
	// page of results, subsequent pages can be fetched via the same
	// mechanism (GetQueryResults).
	JobReference *JobReference `json:"jobReference,omitempty"`

	// JobComplete: Whether the query has completed or not. If rows or
	// totalRows are present, this will always be true. If this is false,
	// totalRows will not be available.
	JobComplete bool `json:"jobComplete,omitempty"`
}

type Job struct {
	// Id: [Output-only] Opaque ID field of the job
	Id string `json:"id,omitempty"`

	// Status: [Output-only] The status of this job. Examine this value when
	// polling an asynchronous job to see if the job is complete.
	Status *JobStatus `json:"status,omitempty"`

	// Configuration: [Required] Describes the job configuration.
	Configuration *JobConfiguration `json:"configuration,omitempty"`

	// JobReference: [Optional] Reference describing the unique-per-user
	// name of the job.
	JobReference *JobReference `json:"jobReference,omitempty"`

	// Statistics: [Output-only] Information about the job, including
	// starting time and ending time of the job.
	Statistics *JobStatistics `json:"statistics,omitempty"`

	// SelfLink: [Output-only] A URL that can be used to access this
	// resource again.
	SelfLink string `json:"selfLink,omitempty"`

	// Etag: [Output-only] A hash of this resource.
	Etag string `json:"etag,omitempty"`

	// Kind: [Output-only] The type of the resource.
	Kind string `json:"kind,omitempty"`
}

type JobConfiguration struct {
	// Query: [Pick one] Configures a query job.
	Query *JobConfigurationQuery `json:"query,omitempty"`

	// Load: [Pick one] Configures a load job.
	Load *JobConfigurationLoad `json:"load,omitempty"`

	// Extract: [Pick one] Configures an extract job.
	Extract *JobConfigurationExtract `json:"extract,omitempty"`

	// Copy: [Pick one] Copies a table.
	Copy *JobConfigurationTableCopy `json:"copy,omitempty"`

	// Link: [Pick one] Configures a link job.
	Link *JobConfigurationLink `json:"link,omitempty"`

	// Properties: [Optional] Properties providing extra details about how
	// the job should be run. Not used for most jobs.
	Properties *JobConfigurationProperties `json:"properties,omitempty"`
}

type JobConfigurationExtract struct {
	// PrintHeader: [Optional] Whether to print out a heder row in the
	// results. Default is true.
	PrintHeader bool `json:"printHeader,omitempty"`

	// FieldDelimiter: [Optional] Delimiter to use between fields in the
	// exported data. Default is ','
	FieldDelimiter string `json:"fieldDelimiter,omitempty"`

	// DestinationUri: [Required] The fully-qualified Google Cloud Storage
	// URI where the extracted table should be written.
	DestinationUri string `json:"destinationUri,omitempty"`

	// SourceTable: [Required] A reference to the table being exported.
	SourceTable *TableReference `json:"sourceTable,omitempty"`
}

type JobConfigurationLink struct {
	// SourceUri: [Required] URI of source table to link.
	SourceUri []string `json:"sourceUri,omitempty"`

	// DestinationTable: [Required] The destination table of the link job.
	DestinationTable *TableReference `json:"destinationTable,omitempty"`

	// CreateDisposition: [Optional] Whether or not to create a new table,
	// if none exists.
	CreateDisposition string `json:"createDisposition,omitempty"`

	// WriteDisposition: [Optional] Whether to overwrite an existing table
	// (WRITE_TRUNCATE), append to an existing table (WRITE_APPEND), or
	// require that the the table is empty (WRITE_EMPTY). Default is
	// WRITE_APPEND.
	WriteDisposition string `json:"writeDisposition,omitempty"`
}

type JobConfigurationLoad struct {
	// FieldDelimiter: [Optional] Delimiter to use between fields in the
	// import data. Default is ','
	FieldDelimiter string `json:"fieldDelimiter,omitempty"`

	// MaxBadRecords: [Optional] Maximum number of bad records that should
	// be ignored before the entire job is aborted and no updates are
	// performed.
	MaxBadRecords int64 `json:"maxBadRecords,omitempty"`

	// Encoding: [Optional] Character encoding of the input data. May be
	// UTF-8 or ISO-8859-1. Default is UTF-8.
	Encoding string `json:"encoding,omitempty"`

	// Schema: [Optional] Schema of the table being written to.
	Schema *TableSchema `json:"schema,omitempty"`

	// DestinationTable: [Required] Table being written to.
	DestinationTable *TableReference `json:"destinationTable,omitempty"`

	// CreateDisposition: [Optional] Whether to create the table if it
	// doesn't already exist (CREATE_IF_NEEDED) or to require the table
	// already exist (CREATE_NEVER). Default is CREATE_IF_NEEDED.
	CreateDisposition string `json:"createDisposition,omitempty"`

	// SkipLeadingRows: [Optional] Number of rows of initial data to skip in
	// the data being imported.
	SkipLeadingRows int64 `json:"skipLeadingRows,omitempty"`

	// WriteDisposition: [Optional] Whether to overwrite an existing table
	// (WRITE_TRUNCATE), append to an existing table (WRITE_APPEND), or
	// require that the the table is empty (WRITE_EMPTY). Default is
	// WRITE_APPEND.
	WriteDisposition string `json:"writeDisposition,omitempty"`

	// SourceUris: [Required] Source URIs describing Google Cloud Storage
	// locations of data to load.
	SourceUris []string `json:"sourceUris,omitempty"`
}

type JobConfigurationProperties struct {
}

type JobConfigurationQuery struct {
	// DestinationTable: [Optional] Describes the table where the query
	// results should be stored. If not present, a new table will be created
	// to store the results.
	DestinationTable *TableReference `json:"destinationTable,omitempty"`

	// CreateDisposition: [Optional] Whether to create the table if it
	// doesn't already exist (CREATE_IF_NEEDED) or to require the table
	// already exist (CREATE_NEVER). Default is CREATE_IF_NEEDED.
	CreateDisposition string `json:"createDisposition,omitempty"`

	// WriteDisposition: [Optional] Whether to overwrite an existing table
	// (WRITE_TRUNCATE), append to an existing table (WRITE_APPEND), or
	// require that the the table is empty (WRITE_EMPTY). Default is
	// WRITE_EMPTY.
	WriteDisposition string `json:"writeDisposition,omitempty"`

	// Query: [Required] BigQuery SQL query to execute.
	Query string `json:"query,omitempty"`

	// DefaultDataset: [Optional] Specifies the default dataset to assume
	// for unqualified table names in the query.
	DefaultDataset *DatasetReference `json:"defaultDataset,omitempty"`
}

type JobConfigurationTableCopy struct {
	// DestinationTable: [Required] The destination table
	DestinationTable *TableReference `json:"destinationTable,omitempty"`

	// SourceTable: [Required] Source table to copy.
	SourceTable *TableReference `json:"sourceTable,omitempty"`

	// CreateDisposition: [Optional] Whether or not to create a new table,
	// if none exists.
	CreateDisposition string `json:"createDisposition,omitempty"`

	// WriteDisposition: [Optional] Whether or not to append or require the
	// table to be empty.
	WriteDisposition string `json:"writeDisposition,omitempty"`
}

type JobList struct {
	// NextPageToken: A token to request the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalItems: Total number of jobs in this collection.
	TotalItems int64 `json:"totalItems,omitempty"`

	// Etag: A hash of this page of results.
	Etag string `json:"etag,omitempty"`

	// Kind: The resource type of the response.
	Kind string `json:"kind,omitempty"`

	// Jobs: List of jobs that were requested.
	Jobs []*JobListJobs `json:"jobs,omitempty"`
}

type JobListJobs struct {
	// ErrorResult: A result object that will be present only if the job has
	// failed.
	ErrorResult *ErrorProto `json:"errorResult,omitempty"`

	// Kind: The resource type.
	Kind string `json:"kind,omitempty"`

	// Id: Unique opaque ID of the job.
	Id string `json:"id,omitempty"`

	// Status: [Full-projection-only] Describes the state of the job.
	Status *JobStatus `json:"status,omitempty"`

	// Configuration: [Full-projection-only] Specifies the job
	// configuration.
	Configuration *JobConfiguration `json:"configuration,omitempty"`

	// JobReference: Job reference uniquely identifying the job.
	JobReference *JobReference `json:"jobReference,omitempty"`

	// Statistics: [Output-only] Information about the job, including
	// starting time and ending time of the job.
	Statistics *JobStatistics `json:"statistics,omitempty"`

	// State: Running state of the job. When the state is DONE, errorResult
	// can be checked to determine whether the job succeeded or failed.
	State string `json:"state,omitempty"`
}

type JobReference struct {
	// JobId: [Required] ID of the job.
	JobId string `json:"jobId,omitempty"`

	// ProjectId: [Required] Project ID being billed for the job.
	ProjectId string `json:"projectId,omitempty"`
}

type JobStatistics struct {
	// TotalBytesProcessed: [Output-only] Total bytes processed for this
	// job.
	TotalBytesProcessed int64 `json:"totalBytesProcessed,omitempty,string"`

	// EndTime: [Output-only] End time of this job, in milliseconds since
	// the epoch.
	EndTime int64 `json:"endTime,omitempty,string"`

	// StartTime: [Output-only] Start time of this job, in milliseconds
	// since the epoch.
	StartTime int64 `json:"startTime,omitempty,string"`
}

type JobStatus struct {
	// ErrorResult: [Output-only] Final error result of the job. If present,
	// indicates that the job has completed and was unsuccessful.
	ErrorResult *ErrorProto `json:"errorResult,omitempty"`

	// Errors: [Output-only] All errors encountered during the running of
	// the job. Errors here do not necessarily mean that the job has
	// completed or was unsuccessful.
	Errors []*ErrorProto `json:"errors,omitempty"`

	// State: [Output-only] Running state of the job.
	State string `json:"state,omitempty"`
}

type ProjectList struct {
	// Projects: Projects to which you have at least READ access.
	Projects []*ProjectListProjects `json:"projects,omitempty"`

	// NextPageToken: A token to request the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalItems: The total number of projects in the list.
	TotalItems int64 `json:"totalItems,omitempty"`

	// Etag: A hash of the page of results
	Etag string `json:"etag,omitempty"`

	// Kind: The type of list.
	Kind string `json:"kind,omitempty"`
}

type ProjectListProjects struct {
	// FriendlyName: A descriptive name for this project.
	FriendlyName string `json:"friendlyName,omitempty"`

	// Kind: The resource type.
	Kind string `json:"kind,omitempty"`

	// Id: An opaque ID of this project.
	Id string `json:"id,omitempty"`

	// ProjectReference: A unique reference to this project.
	ProjectReference *ProjectReference `json:"projectReference,omitempty"`
}

type ProjectReference struct {
	// ProjectId: [Required] ID of the project. Can be either the numeric ID
	// or the assigned ID of the project.
	ProjectId string `json:"projectId,omitempty"`
}

type QueryRequest struct {
	// DefaultDataset: [Optional] Specifies the default datasetId and
	// projectId to assume for any unqualified table names in the query. If
	// not set, all table names in the query string must be fully-qualified
	// in the format projectId:datasetId.tableid.
	DefaultDataset *DatasetReference `json:"defaultDataset,omitempty"`

	// MaxResults: [Optional] The maximum number of results to return per
	// page of results. If the response list exceeds the maximum response
	// size for a single response, you will have to page through the
	// results. Default is to return the maximum response size.
	MaxResults int64 `json:"maxResults,omitempty"`

	// TimeoutMs: [Optional] How long to wait for the query to complete, in
	// milliseconds, before returning. Default is to return immediately. If
	// the timeout passes before the job completes, the request will fail
	// with a TIMEOUT error.
	TimeoutMs int64 `json:"timeoutMs,omitempty"`

	// Kind: The resource type of the request.
	Kind string `json:"kind,omitempty"`

	// Query: [Required] A query string, following the BigQuery query syntax
	// of the query to execute. Table names should be qualified by dataset
	// name in the format projectId:datasetId.tableId unless you specify the
	// defaultDataset value. If the table is in the same project as the job,
	// you can omit the project ID. Example: SELECT f1 FROM
	// myProjectId:myDatasetId.myTableId.
	Query string `json:"query,omitempty"`
}

type QueryResponse struct {
	// JobComplete: Whether the query has completed or not. If rows or
	// totalRows are present, this will always be true. If this is false,
	// totalRows will not be available.
	JobComplete bool `json:"jobComplete,omitempty"`

	// Schema: The schema of the results. Present only when the query
	// completes successfully.
	Schema *TableSchema `json:"schema,omitempty"`

	// Kind: The resource type.
	Kind string `json:"kind,omitempty"`

	// TotalRows: The total number of rows in the complete query result set,
	// which can be more than the number of rows in this single page of
	// results.
	TotalRows uint64 `json:"totalRows,omitempty,string"`

	// Rows: An object with as many results as can be contained within the
	// maximum permitted reply size. To get any additional rows, you can
	// call GetQueryResults and specify the jobReference returned above.
	Rows []*TableRow `json:"rows,omitempty"`

	// JobReference: Reference to the Job that was created to run the query.
	// This field will be present even if the original request timed out, in
	// which case GetQueryResults can be used to read the results once the
	// query has completed. Since this API only returns the first page of
	// results, subsequent pages can be fetched via the same mechanism
	// (GetQueryResults).
	JobReference *JobReference `json:"jobReference,omitempty"`
}

type Table struct {
	// Description: [Optional] A user-friendly description of this table.
	Description string `json:"description,omitempty"`

	// ExpirationTime: [Optional] The time when this table expires, in
	// milliseconds since the epoch. If not present, the table will persist
	// indefinitely. Expired tables will be deleted and their storage
	// reclaimed.
	ExpirationTime int64 `json:"expirationTime,omitempty,string"`

	// TableReference: [Required] Reference describing the ID of this table.
	TableReference *TableReference `json:"tableReference,omitempty"`

	// SelfLink: [Output-only] A URL that can be used to access this
	// resource again.
	SelfLink string `json:"selfLink,omitempty"`

	// Schema: [Optional] Describes the schema of this table.
	Schema *TableSchema `json:"schema,omitempty"`

	// FriendlyName: [Optional] A descriptive name for this table.
	FriendlyName string `json:"friendlyName,omitempty"`

	// CreationTime: [Output-only] The time when this table was created, in
	// milliseconds since the epoch.
	CreationTime int64 `json:"creationTime,omitempty,string"`

	// Etag: [Output-only] A hash of this resource.
	Etag string `json:"etag,omitempty"`

	// Kind: [Output-only] The type of the resource.
	Kind string `json:"kind,omitempty"`

	// NumRows: [Output-only] The number of rows of data in this table.
	NumRows uint64 `json:"numRows,omitempty,string"`

	// Id: [Output-only] An opaque ID uniquely identifying the table.
	Id string `json:"id,omitempty"`

	// LastModifiedTime: [Output-only] The time when this table was last
	// modified, in milliseconds since the epoch.
	LastModifiedTime int64 `json:"lastModifiedTime,omitempty,string"`

	// NumBytes: [Output-only] The size of the table in bytes.
	NumBytes int64 `json:"numBytes,omitempty,string"`
}

type TableDataList struct {
	// TotalRows: The total number of rows in the complete table.
	TotalRows int64 `json:"totalRows,omitempty,string"`

	// Rows: Rows of results.
	Rows []*TableRow `json:"rows,omitempty"`

	// Etag: A hash of this page of results.
	Etag string `json:"etag,omitempty"`

	// Kind: The resource type of the response.
	Kind string `json:"kind,omitempty"`
}

type TableFieldSchema struct {
	// Fields: [Optional] Describes nested fields when type is RECORD.
	Fields []*TableFieldSchema `json:"fields,omitempty"`

	// Type: [Required] Data type of the field.
	Type string `json:"type,omitempty"`

	// Name: [Required] Name of the field.
	Name string `json:"name,omitempty"`

	// Mode: [Optional] Mode of the field (whether or not it can be null.
	// Default is NULLABLE.
	Mode string `json:"mode,omitempty"`
}

type TableList struct {
	// NextPageToken: A token to request the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalItems: The total number of tables in the dataset.
	TotalItems int64 `json:"totalItems,omitempty"`

	// Etag: A hash of this page of results.
	Etag string `json:"etag,omitempty"`

	// Kind: The type of list.
	Kind string `json:"kind,omitempty"`

	// Tables: Tables in the requested dataset.
	Tables []*TableListTables `json:"tables,omitempty"`
}

type TableListTables struct {
	// TableReference: A reference uniquely identifying the table.
	TableReference *TableReference `json:"tableReference,omitempty"`

	// FriendlyName: The user-friendly name for this table.
	FriendlyName string `json:"friendlyName,omitempty"`

	// Kind: The resource type.
	Kind string `json:"kind,omitempty"`

	// Id: An opaque ID of the table
	Id string `json:"id,omitempty"`
}

type TableReference struct {
	// TableId: [Required] ID of the table.
	TableId string `json:"tableId,omitempty"`

	// ProjectId: [Required] ID of the project billed for storage of the
	// table.
	ProjectId string `json:"projectId,omitempty"`

	// DatasetId: [Required] ID of the dataset containing the table.
	DatasetId string `json:"datasetId,omitempty"`
}

type TableRow struct {
	// F: Represents a single row in the result set, consisting of one or
	// more fields.
	F []*TableRowF `json:"f,omitempty"`
}

type TableRowF struct {
	// V: Contains the field value in this row, as a string.
	V string `json:"v,omitempty"`
}

type TableSchema struct {
	// Fields: Describes the fields in a table.
	Fields []*TableFieldSchema `json:"fields,omitempty"`
}

// method id "bigquery.datasets.delete":

type DatasetsDeleteCall struct {
	s         *Service
	projectId string
	datasetId string
	opt_      map[string]interface{}
}

// Delete: Deletes the dataset specified by datasetId value. Before you
// can delete a dataset, you must delete all its tables, either manually
// or by specifying deleteContents. Immediately after deletion, you can
// create another dataset with the same name.
func (r *DatasetsService) Delete(projectId string, datasetId string) *DatasetsDeleteCall {
	c := &DatasetsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.datasetId = datasetId
	return c
}

// DeleteContents sets the optional parameter "deleteContents": If True,
// delete all the tables in the dataset. If False and the dataset
// contains tables, the request will fail. Default is False
func (c *DatasetsDeleteCall) DeleteContents(deleteContents bool) *DatasetsDeleteCall {
	c.opt_["deleteContents"] = deleteContents
	return c
}

func (c *DatasetsDeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["deleteContents"]; ok {
		params.Set("deleteContents", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/bigquery/v2/", "projects/{projectId}/datasets/{datasetId}")
	urls = strings.Replace(urls, "{projectId}", cleanPathString(c.projectId), 1)
	urls = strings.Replace(urls, "{datasetId}", cleanPathString(c.datasetId), 1)
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
	//   "description": "Deletes the dataset specified by datasetId value. Before you can delete a dataset, you must delete all its tables, either manually or by specifying deleteContents. Immediately after deletion, you can create another dataset with the same name.",
	//   "httpMethod": "DELETE",
	//   "id": "bigquery.datasets.delete",
	//   "parameterOrder": [
	//     "projectId",
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Dataset ID of dataset being deleted",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "deleteContents": {
	//       "description": "If True, delete all the tables in the dataset. If False and the dataset contains tables, the request will fail. Default is False",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the dataset being deleted",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/datasets/{datasetId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery"
	//   ]
	// }

}

// method id "bigquery.datasets.get":

type DatasetsGetCall struct {
	s         *Service
	projectId string
	datasetId string
	opt_      map[string]interface{}
}

// Get: Returns the dataset specified by datasetID.
func (r *DatasetsService) Get(projectId string, datasetId string) *DatasetsGetCall {
	c := &DatasetsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.datasetId = datasetId
	return c
}

func (c *DatasetsGetCall) Do() (*Dataset, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/bigquery/v2/", "projects/{projectId}/datasets/{datasetId}")
	urls = strings.Replace(urls, "{projectId}", cleanPathString(c.projectId), 1)
	urls = strings.Replace(urls, "{datasetId}", cleanPathString(c.datasetId), 1)
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
	ret := new(Dataset)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the dataset specified by datasetID.",
	//   "httpMethod": "GET",
	//   "id": "bigquery.datasets.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Dataset ID of the requested dataset",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the requested dataset",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/datasets/{datasetId}",
	//   "response": {
	//     "$ref": "Dataset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery"
	//   ]
	// }

}

// method id "bigquery.datasets.insert":

type DatasetsInsertCall struct {
	s         *Service
	projectId string
	dataset   *Dataset
	opt_      map[string]interface{}
}

// Insert: Creates a new empty dataset.
func (r *DatasetsService) Insert(projectId string, dataset *Dataset) *DatasetsInsertCall {
	c := &DatasetsInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.dataset = dataset
	return c
}

func (c *DatasetsInsertCall) Do() (*Dataset, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.dataset)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/bigquery/v2/", "projects/{projectId}/datasets")
	urls = strings.Replace(urls, "{projectId}", cleanPathString(c.projectId), 1)
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
	ret := new(Dataset)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new empty dataset.",
	//   "httpMethod": "POST",
	//   "id": "bigquery.datasets.insert",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "Project ID of the new dataset",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/datasets",
	//   "request": {
	//     "$ref": "Dataset"
	//   },
	//   "response": {
	//     "$ref": "Dataset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery"
	//   ]
	// }

}

// method id "bigquery.datasets.list":

type DatasetsListCall struct {
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// List: Lists all the datasets in the specified project to which the
// caller has read access; however, a project owner can list (but not
// necessarily get) all datasets in his project.
func (r *DatasetsService) List(projectId string) *DatasetsListCall {
	c := &DatasetsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of results to return
func (c *DatasetsListCall) MaxResults(maxResults int64) *DatasetsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": Page token,
// returned by a previous call, to request the next page of results
func (c *DatasetsListCall) PageToken(pageToken string) *DatasetsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

func (c *DatasetsListCall) Do() (*DatasetList, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/bigquery/v2/", "projects/{projectId}/datasets")
	urls = strings.Replace(urls, "{projectId}", cleanPathString(c.projectId), 1)
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
	ret := new(DatasetList)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all the datasets in the specified project to which the caller has read access; however, a project owner can list (but not necessarily get) all datasets in his project.",
	//   "httpMethod": "GET",
	//   "id": "bigquery.datasets.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "description": "The maximum number of results to return",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Page token, returned by a previous call, to request the next page of results",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the datasets to be listed",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/datasets",
	//   "response": {
	//     "$ref": "DatasetList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery"
	//   ]
	// }

}

// method id "bigquery.datasets.patch":

type DatasetsPatchCall struct {
	s         *Service
	projectId string
	datasetId string
	dataset   *Dataset
	opt_      map[string]interface{}
}

// Patch: Updates information in an existing dataset, specified by
// datasetId. Properties not included in the submitted resource will not
// be changed. If you include the access property without any values
// assigned, the request will fail as you must specify at least one
// owner for a dataset. This method supports patch semantics.
func (r *DatasetsService) Patch(projectId string, datasetId string, dataset *Dataset) *DatasetsPatchCall {
	c := &DatasetsPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.datasetId = datasetId
	c.dataset = dataset
	return c
}

func (c *DatasetsPatchCall) Do() (*Dataset, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.dataset)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/bigquery/v2/", "projects/{projectId}/datasets/{datasetId}")
	urls = strings.Replace(urls, "{projectId}", cleanPathString(c.projectId), 1)
	urls = strings.Replace(urls, "{datasetId}", cleanPathString(c.datasetId), 1)
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
	ret := new(Dataset)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates information in an existing dataset, specified by datasetId. Properties not included in the submitted resource will not be changed. If you include the access property without any values assigned, the request will fail as you must specify at least one owner for a dataset. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "bigquery.datasets.patch",
	//   "parameterOrder": [
	//     "projectId",
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Dataset ID of the dataset being updated",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the dataset being updated",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/datasets/{datasetId}",
	//   "request": {
	//     "$ref": "Dataset"
	//   },
	//   "response": {
	//     "$ref": "Dataset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery"
	//   ]
	// }

}

// method id "bigquery.datasets.update":

type DatasetsUpdateCall struct {
	s         *Service
	projectId string
	datasetId string
	dataset   *Dataset
	opt_      map[string]interface{}
}

// Update: Updates information in an existing dataset, specified by
// datasetId. Properties not included in the submitted resource will not
// be changed. If you include the access property without any values
// assigned, the request will fail as you must specify at least one
// owner for a dataset.
func (r *DatasetsService) Update(projectId string, datasetId string, dataset *Dataset) *DatasetsUpdateCall {
	c := &DatasetsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.datasetId = datasetId
	c.dataset = dataset
	return c
}

func (c *DatasetsUpdateCall) Do() (*Dataset, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.dataset)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/bigquery/v2/", "projects/{projectId}/datasets/{datasetId}")
	urls = strings.Replace(urls, "{projectId}", cleanPathString(c.projectId), 1)
	urls = strings.Replace(urls, "{datasetId}", cleanPathString(c.datasetId), 1)
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
	ret := new(Dataset)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates information in an existing dataset, specified by datasetId. Properties not included in the submitted resource will not be changed. If you include the access property without any values assigned, the request will fail as you must specify at least one owner for a dataset.",
	//   "httpMethod": "PUT",
	//   "id": "bigquery.datasets.update",
	//   "parameterOrder": [
	//     "projectId",
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Dataset ID of the dataset being updated",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the dataset being updated",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/datasets/{datasetId}",
	//   "request": {
	//     "$ref": "Dataset"
	//   },
	//   "response": {
	//     "$ref": "Dataset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery"
	//   ]
	// }

}

// method id "bigquery.tables.delete":

type TablesDeleteCall struct {
	s         *Service
	projectId string
	datasetId string
	tableId   string
	opt_      map[string]interface{}
}

// Delete: Deletes the table specified by tableId from the dataset. If
// the table contains data, all the data will be deleted.
func (r *TablesService) Delete(projectId string, datasetId string, tableId string) *TablesDeleteCall {
	c := &TablesDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.datasetId = datasetId
	c.tableId = tableId
	return c
}

func (c *TablesDeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/bigquery/v2/", "projects/{projectId}/datasets/{datasetId}/tables/{tableId}")
	urls = strings.Replace(urls, "{projectId}", cleanPathString(c.projectId), 1)
	urls = strings.Replace(urls, "{datasetId}", cleanPathString(c.datasetId), 1)
	urls = strings.Replace(urls, "{tableId}", cleanPathString(c.tableId), 1)
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
	//   "description": "Deletes the table specified by tableId from the dataset. If the table contains data, all the data will be deleted.",
	//   "httpMethod": "DELETE",
	//   "id": "bigquery.tables.delete",
	//   "parameterOrder": [
	//     "projectId",
	//     "datasetId",
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Dataset ID of the table to delete",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the table to delete",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "tableId": {
	//       "description": "Table ID of the table to delete",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/datasets/{datasetId}/tables/{tableId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery"
	//   ]
	// }

}

// method id "bigquery.tables.get":

type TablesGetCall struct {
	s         *Service
	projectId string
	datasetId string
	tableId   string
	opt_      map[string]interface{}
}

// Get: Gets the specified table resource by table ID. This method does
// not return the data in the table, it only returns the table resource,
// which describes the structure of this table.
func (r *TablesService) Get(projectId string, datasetId string, tableId string) *TablesGetCall {
	c := &TablesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.datasetId = datasetId
	c.tableId = tableId
	return c
}

func (c *TablesGetCall) Do() (*Table, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/bigquery/v2/", "projects/{projectId}/datasets/{datasetId}/tables/{tableId}")
	urls = strings.Replace(urls, "{projectId}", cleanPathString(c.projectId), 1)
	urls = strings.Replace(urls, "{datasetId}", cleanPathString(c.datasetId), 1)
	urls = strings.Replace(urls, "{tableId}", cleanPathString(c.tableId), 1)
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
	ret := new(Table)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the specified table resource by table ID. This method does not return the data in the table, it only returns the table resource, which describes the structure of this table.",
	//   "httpMethod": "GET",
	//   "id": "bigquery.tables.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "datasetId",
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Dataset ID of the requested table",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the requested table",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "tableId": {
	//       "description": "Table ID of the requested table",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/datasets/{datasetId}/tables/{tableId}",
	//   "response": {
	//     "$ref": "Table"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery"
	//   ]
	// }

}

// method id "bigquery.tables.insert":

type TablesInsertCall struct {
	s         *Service
	projectId string
	datasetId string
	table     *Table
	opt_      map[string]interface{}
}

// Insert: Creates a new, empty table in the dataset.
func (r *TablesService) Insert(projectId string, datasetId string, table *Table) *TablesInsertCall {
	c := &TablesInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.datasetId = datasetId
	c.table = table
	return c
}

func (c *TablesInsertCall) Do() (*Table, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.table)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/bigquery/v2/", "projects/{projectId}/datasets/{datasetId}/tables")
	urls = strings.Replace(urls, "{projectId}", cleanPathString(c.projectId), 1)
	urls = strings.Replace(urls, "{datasetId}", cleanPathString(c.datasetId), 1)
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
	ret := new(Table)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new, empty table in the dataset.",
	//   "httpMethod": "POST",
	//   "id": "bigquery.tables.insert",
	//   "parameterOrder": [
	//     "projectId",
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Dataset ID of the new table",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the new table",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/datasets/{datasetId}/tables",
	//   "request": {
	//     "$ref": "Table"
	//   },
	//   "response": {
	//     "$ref": "Table"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery"
	//   ]
	// }

}

// method id "bigquery.tables.list":

type TablesListCall struct {
	s         *Service
	projectId string
	datasetId string
	opt_      map[string]interface{}
}

// List: Lists all tables in the specified dataset.
func (r *TablesService) List(projectId string, datasetId string) *TablesListCall {
	c := &TablesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.datasetId = datasetId
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return
func (c *TablesListCall) MaxResults(maxResults int64) *TablesListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": Page token,
// returned by a previous call, to request the next page of results
func (c *TablesListCall) PageToken(pageToken string) *TablesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

func (c *TablesListCall) Do() (*TableList, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/bigquery/v2/", "projects/{projectId}/datasets/{datasetId}/tables")
	urls = strings.Replace(urls, "{projectId}", cleanPathString(c.projectId), 1)
	urls = strings.Replace(urls, "{datasetId}", cleanPathString(c.datasetId), 1)
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
	ret := new(TableList)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all tables in the specified dataset.",
	//   "httpMethod": "GET",
	//   "id": "bigquery.tables.list",
	//   "parameterOrder": [
	//     "projectId",
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Dataset ID of the tables to list",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Page token, returned by a previous call, to request the next page of results",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the tables to list",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/datasets/{datasetId}/tables",
	//   "response": {
	//     "$ref": "TableList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery"
	//   ]
	// }

}

// method id "bigquery.tables.patch":

type TablesPatchCall struct {
	s         *Service
	projectId string
	datasetId string
	tableId   string
	table     *Table
	opt_      map[string]interface{}
}

// Patch: Updates information in an existing table, specified by
// tableId. This method supports patch semantics.
func (r *TablesService) Patch(projectId string, datasetId string, tableId string, table *Table) *TablesPatchCall {
	c := &TablesPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.datasetId = datasetId
	c.tableId = tableId
	c.table = table
	return c
}

func (c *TablesPatchCall) Do() (*Table, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.table)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/bigquery/v2/", "projects/{projectId}/datasets/{datasetId}/tables/{tableId}")
	urls = strings.Replace(urls, "{projectId}", cleanPathString(c.projectId), 1)
	urls = strings.Replace(urls, "{datasetId}", cleanPathString(c.datasetId), 1)
	urls = strings.Replace(urls, "{tableId}", cleanPathString(c.tableId), 1)
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
	ret := new(Table)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates information in an existing table, specified by tableId. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "bigquery.tables.patch",
	//   "parameterOrder": [
	//     "projectId",
	//     "datasetId",
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Dataset ID of the table to update",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the table to update",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "tableId": {
	//       "description": "Table ID of the table to update",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/datasets/{datasetId}/tables/{tableId}",
	//   "request": {
	//     "$ref": "Table"
	//   },
	//   "response": {
	//     "$ref": "Table"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery"
	//   ]
	// }

}

// method id "bigquery.tables.update":

type TablesUpdateCall struct {
	s         *Service
	projectId string
	datasetId string
	tableId   string
	table     *Table
	opt_      map[string]interface{}
}

// Update: Updates information in an existing table, specified by
// tableId.
func (r *TablesService) Update(projectId string, datasetId string, tableId string, table *Table) *TablesUpdateCall {
	c := &TablesUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.datasetId = datasetId
	c.tableId = tableId
	c.table = table
	return c
}

func (c *TablesUpdateCall) Do() (*Table, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.table)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/bigquery/v2/", "projects/{projectId}/datasets/{datasetId}/tables/{tableId}")
	urls = strings.Replace(urls, "{projectId}", cleanPathString(c.projectId), 1)
	urls = strings.Replace(urls, "{datasetId}", cleanPathString(c.datasetId), 1)
	urls = strings.Replace(urls, "{tableId}", cleanPathString(c.tableId), 1)
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
	ret := new(Table)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates information in an existing table, specified by tableId.",
	//   "httpMethod": "PUT",
	//   "id": "bigquery.tables.update",
	//   "parameterOrder": [
	//     "projectId",
	//     "datasetId",
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Dataset ID of the table to update",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the table to update",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "tableId": {
	//       "description": "Table ID of the table to update",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/datasets/{datasetId}/tables/{tableId}",
	//   "request": {
	//     "$ref": "Table"
	//   },
	//   "response": {
	//     "$ref": "Table"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery"
	//   ]
	// }

}

// method id "bigquery.tabledata.list":

type TabledataListCall struct {
	s         *Service
	projectId string
	datasetId string
	tableId   string
	opt_      map[string]interface{}
}

// List: Retrieves table data from a specified set of rows.
func (r *TabledataService) List(projectId string, datasetId string, tableId string) *TabledataListCall {
	c := &TabledataListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.datasetId = datasetId
	c.tableId = tableId
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return
func (c *TabledataListCall) MaxResults(maxResults int64) *TabledataListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// StartIndex sets the optional parameter "startIndex": Zero-based index
// of the starting row to read
func (c *TabledataListCall) StartIndex(startIndex uint64) *TabledataListCall {
	c.opt_["startIndex"] = startIndex
	return c
}

func (c *TabledataListCall) Do() (*TableDataList, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["startIndex"]; ok {
		params.Set("startIndex", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/bigquery/v2/", "projects/{projectId}/datasets/{datasetId}/tables/{tableId}/data")
	urls = strings.Replace(urls, "{projectId}", cleanPathString(c.projectId), 1)
	urls = strings.Replace(urls, "{datasetId}", cleanPathString(c.datasetId), 1)
	urls = strings.Replace(urls, "{tableId}", cleanPathString(c.tableId), 1)
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
	ret := new(TableDataList)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves table data from a specified set of rows.",
	//   "httpMethod": "GET",
	//   "id": "bigquery.tabledata.list",
	//   "parameterOrder": [
	//     "projectId",
	//     "datasetId",
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Dataset ID of the table to read",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the table to read",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "startIndex": {
	//       "description": "Zero-based index of the starting row to read",
	//       "format": "uint64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "tableId": {
	//       "description": "Table ID of the table to read",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/datasets/{datasetId}/tables/{tableId}/data",
	//   "response": {
	//     "$ref": "TableDataList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery"
	//   ]
	// }

}

// method id "bigquery.projects.list":

type ProjectsListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Lists the projects to which you have at least read access.
func (r *ProjectsService) List() *ProjectsListCall {
	c := &ProjectsListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return
func (c *ProjectsListCall) MaxResults(maxResults int64) *ProjectsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": Page token,
// returned by a previous call, to request the next page of results
func (c *ProjectsListCall) PageToken(pageToken string) *ProjectsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

func (c *ProjectsListCall) Do() (*ProjectList, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/bigquery/v2/", "projects")
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
	ret := new(ProjectList)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the projects to which you have at least read access.",
	//   "httpMethod": "GET",
	//   "id": "bigquery.projects.list",
	//   "parameters": {
	//     "maxResults": {
	//       "description": "Maximum number of results to return",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Page token, returned by a previous call, to request the next page of results",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects",
	//   "response": {
	//     "$ref": "ProjectList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery"
	//   ]
	// }

}

// method id "bigquery.jobs.get":

type JobsGetCall struct {
	s         *Service
	projectId string
	jobId     string
	opt_      map[string]interface{}
}

// Get: Retrieves the specified job by ID.
func (r *JobsService) Get(projectId string, jobId string) *JobsGetCall {
	c := &JobsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.jobId = jobId
	return c
}

func (c *JobsGetCall) Do() (*Job, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/bigquery/v2/", "projects/{projectId}/jobs/{jobId}")
	urls = strings.Replace(urls, "{projectId}", cleanPathString(c.projectId), 1)
	urls = strings.Replace(urls, "{jobId}", cleanPathString(c.jobId), 1)
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
	ret := new(Job)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the specified job by ID.",
	//   "httpMethod": "GET",
	//   "id": "bigquery.jobs.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "jobId": {
	//       "description": "Job ID of the requested job",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the requested job",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/jobs/{jobId}",
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery"
	//   ]
	// }

}

// method id "bigquery.jobs.getQueryResults":

type JobsGetQueryResultsCall struct {
	s         *Service
	projectId string
	jobId     string
	opt_      map[string]interface{}
}

// GetQueryResults: Retrieves the results of a query job.
func (r *JobsService) GetQueryResults(projectId string, jobId string) *JobsGetQueryResultsCall {
	c := &JobsGetQueryResultsCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.jobId = jobId
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to read
func (c *JobsGetQueryResultsCall) MaxResults(maxResults int64) *JobsGetQueryResultsCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// StartIndex sets the optional parameter "startIndex": Zero-based index
// of the starting row
func (c *JobsGetQueryResultsCall) StartIndex(startIndex uint64) *JobsGetQueryResultsCall {
	c.opt_["startIndex"] = startIndex
	return c
}

// TimeoutMs sets the optional parameter "timeoutMs": How long to wait
// for the query to complete, in milliseconds, before returning. Default
// is to return immediately. If the timeout passes before the job
// completes, the request will fail with a TIMEOUT error
func (c *JobsGetQueryResultsCall) TimeoutMs(timeoutMs int64) *JobsGetQueryResultsCall {
	c.opt_["timeoutMs"] = timeoutMs
	return c
}

func (c *JobsGetQueryResultsCall) Do() (*GetQueryResultsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["startIndex"]; ok {
		params.Set("startIndex", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["timeoutMs"]; ok {
		params.Set("timeoutMs", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/bigquery/v2/", "projects/{projectId}/queries/{jobId}")
	urls = strings.Replace(urls, "{projectId}", cleanPathString(c.projectId), 1)
	urls = strings.Replace(urls, "{jobId}", cleanPathString(c.jobId), 1)
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
	ret := new(GetQueryResultsResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the results of a query job.",
	//   "httpMethod": "GET",
	//   "id": "bigquery.jobs.getQueryResults",
	//   "parameterOrder": [
	//     "projectId",
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "jobId": {
	//       "description": "Job ID of the query job",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to read",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the query job",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "startIndex": {
	//       "description": "Zero-based index of the starting row",
	//       "format": "uint64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "timeoutMs": {
	//       "description": "How long to wait for the query to complete, in milliseconds, before returning. Default is to return immediately. If the timeout passes before the job completes, the request will fail with a TIMEOUT error",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "projects/{projectId}/queries/{jobId}",
	//   "response": {
	//     "$ref": "GetQueryResultsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery"
	//   ]
	// }

}

// method id "bigquery.jobs.insert":

type JobsInsertCall struct {
	s         *Service
	projectId string
	job       *Job
	opt_      map[string]interface{}
	media_    io.Reader
}

// Insert: Starts a new asynchronous job.
func (r *JobsService) Insert(projectId string, job *Job) *JobsInsertCall {
	c := &JobsInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.job = job
	return c
}
func (c *JobsInsertCall) Media(r io.Reader) *JobsInsertCall {
	c.media_ = r
	return c
}

func (c *JobsInsertCall) Do() (*Job, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.job)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/bigquery/v2/", "projects/{projectId}/jobs")
	if c.media_ != nil {
		urls = strings.Replace(urls, "https://www.googleapis.com/", "https://www.googleapis.com/upload/", 1)
	}
	urls = strings.Replace(urls, "{projectId}", cleanPathString(c.projectId), 1)
	urls += "?" + params.Encode()
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
	ret := new(Job)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Starts a new asynchronous job.",
	//   "httpMethod": "POST",
	//   "id": "bigquery.jobs.insert",
	//   "mediaUpload": {
	//     "accept": [
	//       "application/octet-stream"
	//     ],
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/bigquery/v2/projects/{projectId}/jobs"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/bigquery/v2/projects/{projectId}/jobs"
	//       }
	//     }
	//   },
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "Project ID of the project that will be billed for the job",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/jobs",
	//   "request": {
	//     "$ref": "Job"
	//   },
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery"
	//   ]
	// }

}

// method id "bigquery.jobs.list":

type JobsListCall struct {
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// List: Lists all the Jobs in the specified project that were started
// by the user.
func (r *JobsService) List(projectId string) *JobsListCall {
	c := &JobsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// AllUsers sets the optional parameter "allUsers": Whether to display
// jobs owned by all users in the project. Default false
func (c *JobsListCall) AllUsers(allUsers bool) *JobsListCall {
	c.opt_["allUsers"] = allUsers
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return
func (c *JobsListCall) MaxResults(maxResults int64) *JobsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": Page token,
// returned by a previous call, to request the next page of results
func (c *JobsListCall) PageToken(pageToken string) *JobsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Projection sets the optional parameter "projection": Restrict
// information returned to a set of selected fields
func (c *JobsListCall) Projection(projection string) *JobsListCall {
	c.opt_["projection"] = projection
	return c
}

// StateFilter sets the optional parameter "stateFilter": Filter for job
// state
func (c *JobsListCall) StateFilter(stateFilter string) *JobsListCall {
	c.opt_["stateFilter"] = stateFilter
	return c
}

func (c *JobsListCall) Do() (*JobList, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["allUsers"]; ok {
		params.Set("allUsers", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["projection"]; ok {
		params.Set("projection", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["stateFilter"]; ok {
		params.Set("stateFilter", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/bigquery/v2/", "projects/{projectId}/jobs")
	urls = strings.Replace(urls, "{projectId}", cleanPathString(c.projectId), 1)
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
	ret := new(JobList)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all the Jobs in the specified project that were started by the user.",
	//   "httpMethod": "GET",
	//   "id": "bigquery.jobs.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "allUsers": {
	//       "description": "Whether to display jobs owned by all users in the project. Default false",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Page token, returned by a previous call, to request the next page of results",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the jobs to list",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Restrict information returned to a set of selected fields",
	//       "enum": [
	//         "full",
	//         "minimal"
	//       ],
	//       "enumDescriptions": [
	//         "Includes all job data",
	//         "Does not include the job configuration"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "stateFilter": {
	//       "description": "Filter for job state",
	//       "enum": [
	//         "done",
	//         "pending",
	//         "running"
	//       ],
	//       "enumDescriptions": [
	//         "Finished jobs",
	//         "Pending jobs",
	//         "Running jobs"
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/jobs",
	//   "response": {
	//     "$ref": "JobList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery"
	//   ]
	// }

}

// method id "bigquery.jobs.query":

type JobsQueryCall struct {
	s            *Service
	projectId    string
	queryrequest *QueryRequest
	opt_         map[string]interface{}
}

// Query: Runs a BigQuery SQL query synchronously and returns query
// results if the query completes within a specified timeout.
func (r *JobsService) Query(projectId string, queryrequest *QueryRequest) *JobsQueryCall {
	c := &JobsQueryCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.queryrequest = queryrequest
	return c
}

func (c *JobsQueryCall) Do() (*QueryResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.queryrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/bigquery/v2/", "projects/{projectId}/queries")
	urls = strings.Replace(urls, "{projectId}", cleanPathString(c.projectId), 1)
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
	ret := new(QueryResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Runs a BigQuery SQL query synchronously and returns query results if the query completes within a specified timeout.",
	//   "httpMethod": "POST",
	//   "id": "bigquery.jobs.query",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "Project ID of the project billed for the query",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/queries",
	//   "request": {
	//     "$ref": "QueryRequest"
	//   },
	//   "response": {
	//     "$ref": "QueryResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery"
	//   ]
	// }

}

func cleanPathString(s string) string {
	return strings.Map(func(r rune) rune {
		if r >= 0x30 && r <= 0x7a {
			return r
		}
		return -1
	}, s)
}
