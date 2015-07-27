// Package clouduseraccounts provides access to the Cloud User Accounts API.
//
// See https://cloud.google.com/compute/docs/access/user-accounts/api/latest/
//
// Usage example:
//
//   import "google.golang.org/api/clouduseraccounts/v0.alpha"
//   ...
//   clouduseraccountsService, err := clouduseraccounts.New(oauthHttpClient)
package clouduseraccounts

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

const apiId = "clouduseraccounts:alpha"
const apiName = "clouduseraccounts"
const apiVersion = "alpha"
const basePath = "https://www.googleapis.com/clouduseraccounts/alpha/projects/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// Manage your Google Cloud User Accounts
	CloudUseraccountsScope = "https://www.googleapis.com/auth/cloud.useraccounts"

	// View your Google Cloud User Accounts
	CloudUseraccountsReadonlyScope = "https://www.googleapis.com/auth/cloud.useraccounts.readonly"

	// Manage your Google Compute Accounts
	ComputeaccountsScope = "https://www.googleapis.com/auth/computeaccounts"

	// View your Google Compute Accounts
	ComputeaccountsReadonlyScope = "https://www.googleapis.com/auth/computeaccounts.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.GlobalAccountsOperations = NewGlobalAccountsOperationsService(s)
	s.Groups = NewGroupsService(s)
	s.Linux = NewLinuxService(s)
	s.Users = NewUsersService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	GlobalAccountsOperations *GlobalAccountsOperationsService

	Groups *GroupsService

	Linux *LinuxService

	Users *UsersService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewGlobalAccountsOperationsService(s *Service) *GlobalAccountsOperationsService {
	rs := &GlobalAccountsOperationsService{s: s}
	return rs
}

type GlobalAccountsOperationsService struct {
	s *Service
}

func NewGroupsService(s *Service) *GroupsService {
	rs := &GroupsService{s: s}
	return rs
}

type GroupsService struct {
	s *Service
}

func NewLinuxService(s *Service) *LinuxService {
	rs := &LinuxService{s: s}
	return rs
}

type LinuxService struct {
	s *Service
}

func NewUsersService(s *Service) *UsersService {
	rs := &UsersService{s: s}
	return rs
}

type UsersService struct {
	s *Service
}

// AuthorizedKeysView: A list of authorized public keys for a user
// account.
type AuthorizedKeysView struct {
	// Keys: [Output Only] The list of authorized public keys in SSH format.
	Keys []string `json:"keys,omitempty"`
}

// Group: A Group resource.
type Group struct {
	// CreationTimestamp: [Output Only] Creation timestamp in RFC3339 text
	// format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Description: An optional textual description of the resource;
	// provided by the client when the resource is created.
	Description string `json:"description,omitempty"`

	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// Kind: [Output Only] Type of the resource. Always
	// clouduseraccounts#group for groups.
	Kind string `json:"kind,omitempty"`

	// Members: [Output Only] A list of URLs to User resources who belong to
	// the group. Users may only be members of groups in the same project.
	Members []string `json:"members,omitempty"`

	// Name: Name of the resource; provided by the client when the resource
	// is created.
	Name string `json:"name,omitempty"`

	// SelfLink: [Output Only] Server defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type GroupList struct {
	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id string `json:"id,omitempty"`

	// Items: [Output Only] A list of Group resources.
	Items []*Group `json:"items,omitempty"`

	// Kind: [Output Only] Type of resource. Always
	// clouduseraccounts#groupList for lists of groups.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: [Output Only] A token used to continue a truncated
	// list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: [Output Only] Server defined URL for this resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type GroupsAddMemberRequest struct {
	// Users: Fully-qualified URLs of the User resources to add.
	Users []string `json:"users,omitempty"`
}

type GroupsRemoveMemberRequest struct {
	// Users: Fully-qualified URLs of the User resources to remove.
	Users []string `json:"users,omitempty"`
}

// LinuxAccountViews: A list of all Linux accounts for this project.
// This API is only used by Compute Engine virtual machines to get
// information about user accounts for a project or instance. Linux
// resources are read-only views into users and groups managed by the
// Compute Engine Accounts API.
type LinuxAccountViews struct {
	// GroupViews: [Output Only] A list of all groups within a project.
	GroupViews []*LinuxGroupView `json:"groupViews,omitempty"`

	// Kind: [Output Only] Type of the resource. Always
	// clouduseraccounts#linuxAccountViews for Linux resources.
	Kind string `json:"kind,omitempty"`

	// UserViews: [Output Only] A list of all users within a project.
	UserViews []*LinuxUserView `json:"userViews,omitempty"`
}

type LinuxGetAuthorizedKeysViewResponse struct {
	// Resource: [Output Only] A list of authorized public keys for a user.
	Resource *AuthorizedKeysView `json:"resource,omitempty"`
}

type LinuxGetLinuxAccountViewsResponse struct {
	// Resource: [Output Only] A list of authorized user accounts and
	// groups.
	Resource *LinuxAccountViews `json:"resource,omitempty"`
}

// LinuxGroupView: A detailed view of a Linux group.
type LinuxGroupView struct {
	// Gid: [Output Only] The Group ID.
	Gid int64 `json:"gid,omitempty"`

	// GroupName: [Output Only] Group name.
	GroupName string `json:"groupName,omitempty"`

	// Members: [Output Only] List of user accounts that belong to the
	// group.
	Members []string `json:"members,omitempty"`
}

// LinuxUserView: A detailed view of a Linux user account.
type LinuxUserView struct {
	// Gecos: [Output Only] The GECOS (user information) entry for this
	// account.
	Gecos string `json:"gecos,omitempty"`

	// Gid: [Output Only] User's default group ID.
	Gid int64 `json:"gid,omitempty"`

	// HomeDirectory: [Output Only] The path to the home directory for this
	// account.
	HomeDirectory string `json:"homeDirectory,omitempty"`

	// Shell: [Output Only] The path to the login shell for this account.
	Shell string `json:"shell,omitempty"`

	// Uid: [Output Only] User ID.
	Uid int64 `json:"uid,omitempty"`

	// Username: [Output Only] The username of the account.
	Username string `json:"username,omitempty"`
}

// Operation: An Operation resource, used to manage asynchronous API
// requests.
type Operation struct {
	// ClientOperationId: [Output Only] An optional identifier specified by
	// the client when the mutation was initiated. Must be unique for all
	// Operation resources in the project.
	ClientOperationId string `json:"clientOperationId,omitempty"`

	// CreationTimestamp: [Output Only] Creation timestamp in RFC3339 text
	// format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// EndTime: [Output Only] The time that this operation was completed.
	// This is in RFC3339 text format.
	EndTime string `json:"endTime,omitempty"`

	// Error: [Output Only] If errors are generated during processing of the
	// operation, this field will be populated.
	Error *OperationError `json:"error,omitempty"`

	// HttpErrorMessage: [Output Only] If the operation fails, this field
	// contains the HTTP error message that was returned, such as NOT FOUND.
	HttpErrorMessage string `json:"httpErrorMessage,omitempty"`

	// HttpErrorStatusCode: [Output Only] If the operation fails, this field
	// contains the HTTP error message that was returned, such as 404.
	HttpErrorStatusCode int64 `json:"httpErrorStatusCode,omitempty"`

	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// InsertTime: [Output Only] The time that this operation was requested.
	// This is in RFC3339 text format.
	InsertTime string `json:"insertTime,omitempty"`

	// Kind: [Output Only] Type of the resource. Always compute#Operation
	// for Operation resources.
	Kind string `json:"kind,omitempty"`

	// Name: [Output Only] Name of the resource.
	Name string `json:"name,omitempty"`

	// OperationType: [Output Only] Type of the operation, such as insert,
	// update, and delete.
	OperationType string `json:"operationType,omitempty"`

	// Progress: [Output Only] An optional progress indicator that ranges
	// from 0 to 100. There is no requirement that this be linear or support
	// any granularity of operations. This should not be used to guess at
	// when the operation will be complete. This number should monotonically
	// increase as the operation progresses.
	Progress int64 `json:"progress,omitempty"`

	// Region: [Output Only] URL of the region where the operation resides.
	// Only applicable for regional resources.
	Region string `json:"region,omitempty"`

	// SelfLink: [Output Only] Server defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// StartTime: [Output Only] The time that this operation was started by
	// the server. This is in RFC3339 text format.
	StartTime string `json:"startTime,omitempty"`

	// Status: [Output Only] Status of the operation. Can be one of the
	// following: PENDING, RUNNING, or DONE.
	//
	// Possible values:
	//   "DONE"
	//   "PENDING"
	//   "RUNNING"
	Status string `json:"status,omitempty"`

	// StatusMessage: [Output Only] An optional textual description of the
	// current status of the operation.
	StatusMessage string `json:"statusMessage,omitempty"`

	// TargetId: [Output Only] Unique target ID which identifies a
	// particular incarnation of the target.
	TargetId uint64 `json:"targetId,omitempty,string"`

	// TargetLink: [Output Only] URL of the resource the operation is
	// mutating.
	TargetLink string `json:"targetLink,omitempty"`

	// User: [Output Only] User who requested the operation, for example:
	// user@example.com.
	User string `json:"user,omitempty"`

	// Warnings: [Output Only] If warning messages are generated during
	// processing of the operation, this field will be populated.
	Warnings []*OperationWarnings `json:"warnings,omitempty"`

	// Zone: [Output Only] URL of the zone where the operation resides.
	Zone string `json:"zone,omitempty"`
}

// OperationError: [Output Only] If errors are generated during
// processing of the operation, this field will be populated.
type OperationError struct {
	// Errors: [Output Only] The array of errors encountered while
	// processing this operation.
	Errors []*OperationErrorErrors `json:"errors,omitempty"`
}

type OperationErrorErrors struct {
	// Code: [Output Only] The error type identifier for this error.
	Code string `json:"code,omitempty"`

	// Location: [Output Only] Indicates the field in the request which
	// caused the error. This property is optional.
	Location string `json:"location,omitempty"`

	// Message: [Output Only] An optional, human-readable error message.
	Message string `json:"message,omitempty"`
}

type OperationWarnings struct {
	// Code: [Output Only] The warning type identifier for this warning.
	//
	// Possible values:
	//   "DEPRECATED_RESOURCE_USED"
	//   "DISK_SIZE_LARGER_THAN_IMAGE_SIZE"
	//   "INJECTED_KERNELS_DEPRECATED"
	//   "NEXT_HOP_ADDRESS_NOT_ASSIGNED"
	//   "NEXT_HOP_CANNOT_IP_FORWARD"
	//   "NEXT_HOP_INSTANCE_NOT_FOUND"
	//   "NEXT_HOP_INSTANCE_NOT_ON_NETWORK"
	//   "NEXT_HOP_NOT_RUNNING"
	//   "NOT_CRITICAL_ERROR"
	//   "NO_RESULTS_ON_PAGE"
	//   "REQUIRED_TOS_AGREEMENT"
	//   "RESOURCE_NOT_DELETED"
	//   "SINGLE_INSTANCE_PROPERTY_TEMPLATE"
	//   "UNREACHABLE"
	Code string `json:"code,omitempty"`

	// Data: [Output Only] Metadata for this warning in key: value format.
	Data []*OperationWarningsData `json:"data,omitempty"`

	// Message: [Output Only] Optional human-readable details for this
	// warning.
	Message string `json:"message,omitempty"`
}

type OperationWarningsData struct {
	// Key: [Output Only] A key for the warning data.
	Key string `json:"key,omitempty"`

	// Value: [Output Only] A warning data value corresponding to the key.
	Value string `json:"value,omitempty"`
}

// OperationList: Contains a list of Operation resources.
type OperationList struct {
	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id string `json:"id,omitempty"`

	// Items: [Output Only] The Operation resources.
	Items []*Operation `json:"items,omitempty"`

	// Kind: [Output Only] Type of resource. Always compute#operations for
	// Operations resource.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: [Output Only] A token used to continue a truncate.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: [Output Only] Server defined URL for this resource.
	SelfLink string `json:"selfLink,omitempty"`
}

// PublicKey: A public key for authenticating to guests.
type PublicKey struct {
	// CreationTimestamp: [Output Only] Creation timestamp in RFC3339 text
	// format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Description: An optional textual description of the resource;
	// provided by the client when the resource is created.
	Description string `json:"description,omitempty"`

	// ExpirationTimestamp: Optional expiration timestamp. If provided, the
	// timestamp must be in RFC3339 text format. If not provided, the public
	// key never expires.
	ExpirationTimestamp string `json:"expirationTimestamp,omitempty"`

	// Fingerprint: [Output Only] The fingerprint of the key is defined by
	// RFC4716 to be the MD5 digest of the public key.
	Fingerprint string `json:"fingerprint,omitempty"`

	// Key: Public key text in SSH format, defined by RFC4253 section 6.6.
	Key string `json:"key,omitempty"`
}

// User: A User resource.
type User struct {
	// CreationTimestamp: [Output Only] Creation timestamp in RFC3339 text
	// format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Description: An optional textual description of the resource;
	// provided by the client when the resource is created.
	Description string `json:"description,omitempty"`

	// Groups: [Output Only] A list of URLs to Group resources who contain
	// the user. Users are only members of groups in the same project.
	Groups []string `json:"groups,omitempty"`

	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// Kind: [Output Only] Type of the resource. Always
	// clouduseraccounts#user for users.
	Kind string `json:"kind,omitempty"`

	// Name: Name of the resource; provided by the client when the resource
	// is created.
	Name string `json:"name,omitempty"`

	// Owner: Email address of account's owner. This account will be
	// validated to make sure it exists. The email can belong to any domain,
	// but it must be tied to a Google account.
	Owner string `json:"owner,omitempty"`

	// PublicKeys: [Output Only] Public keys that this user may use to
	// login.
	PublicKeys []*PublicKey `json:"publicKeys,omitempty"`

	// SelfLink: [Output Only] Server defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type UserList struct {
	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id string `json:"id,omitempty"`

	// Items: [Output Only] A list of User resources.
	Items []*User `json:"items,omitempty"`

	// Kind: [Output Only] Type of resource. Always
	// clouduseraccounts#userList for lists of users.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: [Output Only] A token used to continue a truncated
	// list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: [Output Only] Server defined URL for this resource.
	SelfLink string `json:"selfLink,omitempty"`
}

// method id "clouduseraccounts.globalAccountsOperations.delete":

type GlobalAccountsOperationsDeleteCall struct {
	s         *Service
	project   string
	operation string
	opt_      map[string]interface{}
}

// Delete: Deletes the specified operation resource.
func (r *GlobalAccountsOperationsService) Delete(project string, operation string) *GlobalAccountsOperationsDeleteCall {
	c := &GlobalAccountsOperationsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.operation = operation
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GlobalAccountsOperationsDeleteCall) Fields(s ...googleapi.Field) *GlobalAccountsOperationsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *GlobalAccountsOperationsDeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/operations/{operation}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":   c.project,
		"operation": c.operation,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Deletes the specified operation resource.",
	//   "httpMethod": "DELETE",
	//   "id": "clouduseraccounts.globalAccountsOperations.delete",
	//   "parameterOrder": [
	//     "project",
	//     "operation"
	//   ],
	//   "parameters": {
	//     "operation": {
	//       "description": "Name of the Operations resource to delete.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/operations/{operation}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud.useraccounts",
	//     "https://www.googleapis.com/auth/computeaccounts"
	//   ]
	// }

}

// method id "clouduseraccounts.globalAccountsOperations.get":

type GlobalAccountsOperationsGetCall struct {
	s         *Service
	project   string
	operation string
	opt_      map[string]interface{}
}

// Get: Retrieves the specified operation resource.
func (r *GlobalAccountsOperationsService) Get(project string, operation string) *GlobalAccountsOperationsGetCall {
	c := &GlobalAccountsOperationsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.operation = operation
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GlobalAccountsOperationsGetCall) Fields(s ...googleapi.Field) *GlobalAccountsOperationsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *GlobalAccountsOperationsGetCall) Do() (*Operation, error) {
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
	//   "description": "Retrieves the specified operation resource.",
	//   "httpMethod": "GET",
	//   "id": "clouduseraccounts.globalAccountsOperations.get",
	//   "parameterOrder": [
	//     "project",
	//     "operation"
	//   ],
	//   "parameters": {
	//     "operation": {
	//       "description": "Name of the Operations resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
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
	//     "https://www.googleapis.com/auth/cloud.useraccounts",
	//     "https://www.googleapis.com/auth/cloud.useraccounts.readonly",
	//     "https://www.googleapis.com/auth/computeaccounts",
	//     "https://www.googleapis.com/auth/computeaccounts.readonly"
	//   ]
	// }

}

// method id "clouduseraccounts.globalAccountsOperations.list":

type GlobalAccountsOperationsListCall struct {
	s       *Service
	project string
	opt_    map[string]interface{}
}

// List: Retrieves the list of operation resources contained within the
// specified project.
func (r *GlobalAccountsOperationsService) List(project string) *GlobalAccountsOperationsListCall {
	c := &GlobalAccountsOperationsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	return c
}

// Filter sets the optional parameter "filter": Sets a filter expression
// for filtering listed resources, in the form filter={expression}. Your
// {expression} must contain the following:
// FIELD_NAME COMPARISON_STRING LITERAL_STRING
//
// - FIELD_NAME: The name of the field you want to compare. The field
// name must be valid for the type of resource being filtered. Only
// atomic field types are supported (string, number, boolean). Array and
// object fields are not currently supported.
// - COMPARISON_STRING: The comparison string, either eq (equals) or ne
// (not equals).
// - LITERAL_STRING: The literal string value to filter to. The literal
// value must be valid for the type of field (string, number, boolean).
// For string fields, the literal value is interpreted as a regular
// expression using RE2 syntax. The literal value must match the entire
// field.  For example, you can filter by the name of a
// resource:
// filter=name ne example-instance
// The above filter returns only results whose name field does not equal
// example-instance. You can also enclose your literal string in single,
// double, or no quotes.
func (c *GlobalAccountsOperationsListCall) Filter(filter string) *GlobalAccountsOperationsListCall {
	c.opt_["filter"] = filter
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned.
func (c *GlobalAccountsOperationsListCall) MaxResults(maxResults int64) *GlobalAccountsOperationsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// OrderBy sets the optional parameter "orderBy": Sorts list results by
// a certain order. By default, results are returned in alphanumerical
// order based on the resource name.
//
// You can also sort results in descending order based on the creation
// timestamp using orderBy="creationTimestamp desc". This sorts results
// based on the creationTimestamp field in reverse chronological order
// (newest result first). Use this to sort resources like operations so
// that the newest operation is returned first.
//
// Currently, only sorting by name or creationTimestamp desc is
// supported.
func (c *GlobalAccountsOperationsListCall) OrderBy(orderBy string) *GlobalAccountsOperationsListCall {
	c.opt_["orderBy"] = orderBy
	return c
}

// PageToken sets the optional parameter "pageToken": Specifies a page
// token to use. Use this parameter if you want to list the next page of
// results. Set pageToken to the nextPageToken returned by a previous
// list request.
func (c *GlobalAccountsOperationsListCall) PageToken(pageToken string) *GlobalAccountsOperationsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GlobalAccountsOperationsListCall) Fields(s ...googleapi.Field) *GlobalAccountsOperationsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *GlobalAccountsOperationsListCall) Do() (*OperationList, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["filter"]; ok {
		params.Set("filter", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["orderBy"]; ok {
		params.Set("orderBy", fmt.Sprintf("%v", v))
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
	var ret *OperationList
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the list of operation resources contained within the specified project.",
	//   "httpMethod": "GET",
	//   "id": "clouduseraccounts.globalAccountsOperations.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Sets a filter expression for filtering listed resources, in the form filter={expression}. Your {expression} must contain the following:\nFIELD_NAME COMPARISON_STRING LITERAL_STRING\n \n- FIELD_NAME: The name of the field you want to compare. The field name must be valid for the type of resource being filtered. Only atomic field types are supported (string, number, boolean). Array and object fields are not currently supported. \n- COMPARISON_STRING: The comparison string, either eq (equals) or ne (not equals). \n- LITERAL_STRING: The literal string value to filter to. The literal value must be valid for the type of field (string, number, boolean). For string fields, the literal value is interpreted as a regular expression using RE2 syntax. The literal value must match the entire field.  For example, you can filter by the name of a resource:\nfilter=name ne example-instance\nThe above filter returns only results whose name field does not equal example-instance. You can also enclose your literal string in single, double, or no quotes.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Maximum count of results to be returned.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "orderBy": {
	//       "description": "Sorts list results by a certain order. By default, results are returned in alphanumerical order based on the resource name.\n\nYou can also sort results in descending order based on the creation timestamp using orderBy=\"creationTimestamp desc\". This sorts results based on the creationTimestamp field in reverse chronological order (newest result first). Use this to sort resources like operations so that the newest operation is returned first.\n\nCurrently, only sorting by name or creationTimestamp desc is supported.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Specifies a page token to use. Use this parameter if you want to list the next page of results. Set pageToken to the nextPageToken returned by a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/operations",
	//   "response": {
	//     "$ref": "OperationList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud.useraccounts",
	//     "https://www.googleapis.com/auth/cloud.useraccounts.readonly",
	//     "https://www.googleapis.com/auth/computeaccounts",
	//     "https://www.googleapis.com/auth/computeaccounts.readonly"
	//   ]
	// }

}

// method id "clouduseraccounts.groups.addMember":

type GroupsAddMemberCall struct {
	s                      *Service
	project                string
	groupName              string
	groupsaddmemberrequest *GroupsAddMemberRequest
	opt_                   map[string]interface{}
}

// AddMember: Adds users to the specified group.
func (r *GroupsService) AddMember(project string, groupName string, groupsaddmemberrequest *GroupsAddMemberRequest) *GroupsAddMemberCall {
	c := &GroupsAddMemberCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.groupName = groupName
	c.groupsaddmemberrequest = groupsaddmemberrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupsAddMemberCall) Fields(s ...googleapi.Field) *GroupsAddMemberCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *GroupsAddMemberCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.groupsaddmemberrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/groups/{groupName}/addMember")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":   c.project,
		"groupName": c.groupName,
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
	//   "description": "Adds users to the specified group.",
	//   "httpMethod": "POST",
	//   "id": "clouduseraccounts.groups.addMember",
	//   "parameterOrder": [
	//     "project",
	//     "groupName"
	//   ],
	//   "parameters": {
	//     "groupName": {
	//       "description": "Name of the group for this request.",
	//       "location": "path",
	//       "pattern": "[a-z][-a-z0-9_]{0,31}",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/groups/{groupName}/addMember",
	//   "request": {
	//     "$ref": "GroupsAddMemberRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud.useraccounts",
	//     "https://www.googleapis.com/auth/computeaccounts"
	//   ]
	// }

}

// method id "clouduseraccounts.groups.delete":

type GroupsDeleteCall struct {
	s         *Service
	project   string
	groupName string
	opt_      map[string]interface{}
}

// Delete: Deletes the specified Group resource.
func (r *GroupsService) Delete(project string, groupName string) *GroupsDeleteCall {
	c := &GroupsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.groupName = groupName
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupsDeleteCall) Fields(s ...googleapi.Field) *GroupsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *GroupsDeleteCall) Do() (*Operation, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/groups/{groupName}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":   c.project,
		"groupName": c.groupName,
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
	//   "description": "Deletes the specified Group resource.",
	//   "httpMethod": "DELETE",
	//   "id": "clouduseraccounts.groups.delete",
	//   "parameterOrder": [
	//     "project",
	//     "groupName"
	//   ],
	//   "parameters": {
	//     "groupName": {
	//       "description": "Name of the Group resource to delete.",
	//       "location": "path",
	//       "pattern": "[a-z][-a-z0-9_]{0,31}",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/groups/{groupName}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud.useraccounts",
	//     "https://www.googleapis.com/auth/computeaccounts"
	//   ]
	// }

}

// method id "clouduseraccounts.groups.get":

type GroupsGetCall struct {
	s         *Service
	project   string
	groupName string
	opt_      map[string]interface{}
}

// Get: Returns the specified Group resource.
func (r *GroupsService) Get(project string, groupName string) *GroupsGetCall {
	c := &GroupsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.groupName = groupName
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupsGetCall) Fields(s ...googleapi.Field) *GroupsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *GroupsGetCall) Do() (*Group, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/groups/{groupName}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":   c.project,
		"groupName": c.groupName,
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
	var ret *Group
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the specified Group resource.",
	//   "httpMethod": "GET",
	//   "id": "clouduseraccounts.groups.get",
	//   "parameterOrder": [
	//     "project",
	//     "groupName"
	//   ],
	//   "parameters": {
	//     "groupName": {
	//       "description": "Name of the Group resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z][-a-z0-9_]{0,31}",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/groups/{groupName}",
	//   "response": {
	//     "$ref": "Group"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud.useraccounts",
	//     "https://www.googleapis.com/auth/cloud.useraccounts.readonly",
	//     "https://www.googleapis.com/auth/computeaccounts",
	//     "https://www.googleapis.com/auth/computeaccounts.readonly"
	//   ]
	// }

}

// method id "clouduseraccounts.groups.insert":

type GroupsInsertCall struct {
	s       *Service
	project string
	group   *Group
	opt_    map[string]interface{}
}

// Insert: Creates a Group resource in the specified project using the
// data included in the request.
func (r *GroupsService) Insert(project string, group *Group) *GroupsInsertCall {
	c := &GroupsInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.group = group
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupsInsertCall) Fields(s ...googleapi.Field) *GroupsInsertCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *GroupsInsertCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.group)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/groups")
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
	//   "description": "Creates a Group resource in the specified project using the data included in the request.",
	//   "httpMethod": "POST",
	//   "id": "clouduseraccounts.groups.insert",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/groups",
	//   "request": {
	//     "$ref": "Group"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud.useraccounts",
	//     "https://www.googleapis.com/auth/computeaccounts"
	//   ]
	// }

}

// method id "clouduseraccounts.groups.list":

type GroupsListCall struct {
	s       *Service
	project string
	opt_    map[string]interface{}
}

// List: Retrieves the list of groups contained within the specified
// project.
func (r *GroupsService) List(project string) *GroupsListCall {
	c := &GroupsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	return c
}

// Filter sets the optional parameter "filter": Sets a filter expression
// for filtering listed resources, in the form filter={expression}. Your
// {expression} must contain the following:
// FIELD_NAME COMPARISON_STRING LITERAL_STRING
//
// - FIELD_NAME: The name of the field you want to compare. The field
// name must be valid for the type of resource being filtered. Only
// atomic field types are supported (string, number, boolean). Array and
// object fields are not currently supported.
// - COMPARISON_STRING: The comparison string, either eq (equals) or ne
// (not equals).
// - LITERAL_STRING: The literal string value to filter to. The literal
// value must be valid for the type of field (string, number, boolean).
// For string fields, the literal value is interpreted as a regular
// expression using RE2 syntax. The literal value must match the entire
// field.  For example, you can filter by the name of a
// resource:
// filter=name ne example-instance
// The above filter returns only results whose name field does not equal
// example-instance. You can also enclose your literal string in single,
// double, or no quotes.
func (c *GroupsListCall) Filter(filter string) *GroupsListCall {
	c.opt_["filter"] = filter
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned.
func (c *GroupsListCall) MaxResults(maxResults int64) *GroupsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// OrderBy sets the optional parameter "orderBy": Sorts list results by
// a certain order. By default, results are returned in alphanumerical
// order based on the resource name.
//
// You can also sort results in descending order based on the creation
// timestamp using orderBy="creationTimestamp desc". This sorts results
// based on the creationTimestamp field in reverse chronological order
// (newest result first). Use this to sort resources like operations so
// that the newest operation is returned first.
//
// Currently, only sorting by name or creationTimestamp desc is
// supported.
func (c *GroupsListCall) OrderBy(orderBy string) *GroupsListCall {
	c.opt_["orderBy"] = orderBy
	return c
}

// PageToken sets the optional parameter "pageToken": Specifies a page
// token to use. Use this parameter if you want to list the next page of
// results. Set pageToken to the nextPageToken returned by a previous
// list request.
func (c *GroupsListCall) PageToken(pageToken string) *GroupsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupsListCall) Fields(s ...googleapi.Field) *GroupsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *GroupsListCall) Do() (*GroupList, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["filter"]; ok {
		params.Set("filter", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["orderBy"]; ok {
		params.Set("orderBy", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/groups")
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
	var ret *GroupList
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the list of groups contained within the specified project.",
	//   "httpMethod": "GET",
	//   "id": "clouduseraccounts.groups.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Sets a filter expression for filtering listed resources, in the form filter={expression}. Your {expression} must contain the following:\nFIELD_NAME COMPARISON_STRING LITERAL_STRING\n \n- FIELD_NAME: The name of the field you want to compare. The field name must be valid for the type of resource being filtered. Only atomic field types are supported (string, number, boolean). Array and object fields are not currently supported. \n- COMPARISON_STRING: The comparison string, either eq (equals) or ne (not equals). \n- LITERAL_STRING: The literal string value to filter to. The literal value must be valid for the type of field (string, number, boolean). For string fields, the literal value is interpreted as a regular expression using RE2 syntax. The literal value must match the entire field.  For example, you can filter by the name of a resource:\nfilter=name ne example-instance\nThe above filter returns only results whose name field does not equal example-instance. You can also enclose your literal string in single, double, or no quotes.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Maximum count of results to be returned.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "orderBy": {
	//       "description": "Sorts list results by a certain order. By default, results are returned in alphanumerical order based on the resource name.\n\nYou can also sort results in descending order based on the creation timestamp using orderBy=\"creationTimestamp desc\". This sorts results based on the creationTimestamp field in reverse chronological order (newest result first). Use this to sort resources like operations so that the newest operation is returned first.\n\nCurrently, only sorting by name or creationTimestamp desc is supported.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Specifies a page token to use. Use this parameter if you want to list the next page of results. Set pageToken to the nextPageToken returned by a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/groups",
	//   "response": {
	//     "$ref": "GroupList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud.useraccounts",
	//     "https://www.googleapis.com/auth/cloud.useraccounts.readonly",
	//     "https://www.googleapis.com/auth/computeaccounts",
	//     "https://www.googleapis.com/auth/computeaccounts.readonly"
	//   ]
	// }

}

// method id "clouduseraccounts.groups.removeMember":

type GroupsRemoveMemberCall struct {
	s                         *Service
	project                   string
	groupName                 string
	groupsremovememberrequest *GroupsRemoveMemberRequest
	opt_                      map[string]interface{}
}

// RemoveMember: Removes users from the specified group.
func (r *GroupsService) RemoveMember(project string, groupName string, groupsremovememberrequest *GroupsRemoveMemberRequest) *GroupsRemoveMemberCall {
	c := &GroupsRemoveMemberCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.groupName = groupName
	c.groupsremovememberrequest = groupsremovememberrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupsRemoveMemberCall) Fields(s ...googleapi.Field) *GroupsRemoveMemberCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *GroupsRemoveMemberCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.groupsremovememberrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/groups/{groupName}/removeMember")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":   c.project,
		"groupName": c.groupName,
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
	//   "description": "Removes users from the specified group.",
	//   "httpMethod": "POST",
	//   "id": "clouduseraccounts.groups.removeMember",
	//   "parameterOrder": [
	//     "project",
	//     "groupName"
	//   ],
	//   "parameters": {
	//     "groupName": {
	//       "description": "Name of the group for this request.",
	//       "location": "path",
	//       "pattern": "[a-z][-a-z0-9_]{0,31}",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/groups/{groupName}/removeMember",
	//   "request": {
	//     "$ref": "GroupsRemoveMemberRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud.useraccounts",
	//     "https://www.googleapis.com/auth/computeaccounts"
	//   ]
	// }

}

// method id "clouduseraccounts.linux.getAuthorizedKeysView":

type LinuxGetAuthorizedKeysViewCall struct {
	s        *Service
	project  string
	zone     string
	user     string
	instance string
	opt_     map[string]interface{}
}

// GetAuthorizedKeysView: Returns a list of authorized public keys for a
// specific user account.
func (r *LinuxService) GetAuthorizedKeysView(project string, zone string, user string, instance string) *LinuxGetAuthorizedKeysViewCall {
	c := &LinuxGetAuthorizedKeysViewCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.zone = zone
	c.user = user
	c.instance = instance
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LinuxGetAuthorizedKeysViewCall) Fields(s ...googleapi.Field) *LinuxGetAuthorizedKeysViewCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *LinuxGetAuthorizedKeysViewCall) Do() (*LinuxGetAuthorizedKeysViewResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("instance", fmt.Sprintf("%v", c.instance))
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/zones/{zone}/authorizedKeysView/{user}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
		"zone":    c.zone,
		"user":    c.user,
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
	var ret *LinuxGetAuthorizedKeysViewResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns a list of authorized public keys for a specific user account.",
	//   "httpMethod": "POST",
	//   "id": "clouduseraccounts.linux.getAuthorizedKeysView",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "user",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "The fully-qualified URL of the virtual machine requesting the view.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "user": {
	//       "description": "The user account for which you want to get a list of authorized public keys.",
	//       "location": "path",
	//       "pattern": "[a-z][-a-z0-9_]{0,31}",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "Name of the zone for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/authorizedKeysView/{user}",
	//   "response": {
	//     "$ref": "LinuxGetAuthorizedKeysViewResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud.useraccounts",
	//     "https://www.googleapis.com/auth/cloud.useraccounts.readonly",
	//     "https://www.googleapis.com/auth/computeaccounts",
	//     "https://www.googleapis.com/auth/computeaccounts.readonly"
	//   ]
	// }

}

// method id "clouduseraccounts.linux.getLinuxAccountViews":

type LinuxGetLinuxAccountViewsCall struct {
	s        *Service
	project  string
	zone     string
	instance string
	opt_     map[string]interface{}
}

// GetLinuxAccountViews: Retrieves a list of user accounts for an
// instance within a specific project.
func (r *LinuxService) GetLinuxAccountViews(project string, zone string, instance string) *LinuxGetLinuxAccountViewsCall {
	c := &LinuxGetLinuxAccountViewsCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.zone = zone
	c.instance = instance
	return c
}

// Filter sets the optional parameter "filter": Sets a filter expression
// for filtering listed resources, in the form filter={expression}. Your
// {expression} must contain the following:
// FIELD_NAME COMPARISON_STRING LITERAL_STRING
//
// - FIELD_NAME: The name of the field you want to compare. The field
// name must be valid for the type of resource being filtered. Only
// atomic field types are supported (string, number, boolean). Array and
// object fields are not currently supported.
// - COMPARISON_STRING: The comparison string, either eq (equals) or ne
// (not equals).
// - LITERAL_STRING: The literal string value to filter to. The literal
// value must be valid for the type of field (string, number, boolean).
// For string fields, the literal value is interpreted as a regular
// expression using RE2 syntax. The literal value must match the entire
// field.  For example, you can filter by the name of a
// resource:
// filter=name ne example-instance
// The above filter returns only results whose name field does not equal
// example-instance. You can also enclose your literal string in single,
// double, or no quotes.
func (c *LinuxGetLinuxAccountViewsCall) Filter(filter string) *LinuxGetLinuxAccountViewsCall {
	c.opt_["filter"] = filter
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned.
func (c *LinuxGetLinuxAccountViewsCall) MaxResults(maxResults int64) *LinuxGetLinuxAccountViewsCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// OrderBy sets the optional parameter "orderBy": Sorts list results by
// a certain order. By default, results are returned in alphanumerical
// order based on the resource name.
//
// You can also sort results in descending order based on the creation
// timestamp using orderBy="creationTimestamp desc". This sorts results
// based on the creationTimestamp field in reverse chronological order
// (newest result first). Use this to sort resources like operations so
// that the newest operation is returned first.
//
// Currently, only sorting by name or creationTimestamp desc is
// supported.
func (c *LinuxGetLinuxAccountViewsCall) OrderBy(orderBy string) *LinuxGetLinuxAccountViewsCall {
	c.opt_["orderBy"] = orderBy
	return c
}

// PageToken sets the optional parameter "pageToken": Specifies a page
// token to use. Use this parameter if you want to list the next page of
// results. Set pageToken to the nextPageToken returned by a previous
// list request.
func (c *LinuxGetLinuxAccountViewsCall) PageToken(pageToken string) *LinuxGetLinuxAccountViewsCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LinuxGetLinuxAccountViewsCall) Fields(s ...googleapi.Field) *LinuxGetLinuxAccountViewsCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *LinuxGetLinuxAccountViewsCall) Do() (*LinuxGetLinuxAccountViewsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("instance", fmt.Sprintf("%v", c.instance))
	if v, ok := c.opt_["filter"]; ok {
		params.Set("filter", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["orderBy"]; ok {
		params.Set("orderBy", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/zones/{zone}/linuxAccountViews")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
		"zone":    c.zone,
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
	var ret *LinuxGetLinuxAccountViewsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of user accounts for an instance within a specific project.",
	//   "httpMethod": "POST",
	//   "id": "clouduseraccounts.linux.getLinuxAccountViews",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Sets a filter expression for filtering listed resources, in the form filter={expression}. Your {expression} must contain the following:\nFIELD_NAME COMPARISON_STRING LITERAL_STRING\n \n- FIELD_NAME: The name of the field you want to compare. The field name must be valid for the type of resource being filtered. Only atomic field types are supported (string, number, boolean). Array and object fields are not currently supported. \n- COMPARISON_STRING: The comparison string, either eq (equals) or ne (not equals). \n- LITERAL_STRING: The literal string value to filter to. The literal value must be valid for the type of field (string, number, boolean). For string fields, the literal value is interpreted as a regular expression using RE2 syntax. The literal value must match the entire field.  For example, you can filter by the name of a resource:\nfilter=name ne example-instance\nThe above filter returns only results whose name field does not equal example-instance. You can also enclose your literal string in single, double, or no quotes.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "instance": {
	//       "description": "The fully-qualified URL of the virtual machine requesting the views.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Maximum count of results to be returned.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "orderBy": {
	//       "description": "Sorts list results by a certain order. By default, results are returned in alphanumerical order based on the resource name.\n\nYou can also sort results in descending order based on the creation timestamp using orderBy=\"creationTimestamp desc\". This sorts results based on the creationTimestamp field in reverse chronological order (newest result first). Use this to sort resources like operations so that the newest operation is returned first.\n\nCurrently, only sorting by name or creationTimestamp desc is supported.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Specifies a page token to use. Use this parameter if you want to list the next page of results. Set pageToken to the nextPageToken returned by a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "Name of the zone for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/linuxAccountViews",
	//   "response": {
	//     "$ref": "LinuxGetLinuxAccountViewsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud.useraccounts",
	//     "https://www.googleapis.com/auth/cloud.useraccounts.readonly",
	//     "https://www.googleapis.com/auth/computeaccounts",
	//     "https://www.googleapis.com/auth/computeaccounts.readonly"
	//   ]
	// }

}

// method id "clouduseraccounts.users.addPublicKey":

type UsersAddPublicKeyCall struct {
	s         *Service
	project   string
	user      string
	publickey *PublicKey
	opt_      map[string]interface{}
}

// AddPublicKey: Adds a public key to the specified User resource with
// the data included in the request.
func (r *UsersService) AddPublicKey(project string, user string, publickey *PublicKey) *UsersAddPublicKeyCall {
	c := &UsersAddPublicKeyCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.user = user
	c.publickey = publickey
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersAddPublicKeyCall) Fields(s ...googleapi.Field) *UsersAddPublicKeyCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *UsersAddPublicKeyCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.publickey)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/users/{user}/addPublicKey")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
		"user":    c.user,
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
	//   "description": "Adds a public key to the specified User resource with the data included in the request.",
	//   "httpMethod": "POST",
	//   "id": "clouduseraccounts.users.addPublicKey",
	//   "parameterOrder": [
	//     "project",
	//     "user"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "user": {
	//       "description": "Name of the user for this request.",
	//       "location": "path",
	//       "pattern": "[a-z][-a-z0-9_]{0,31}",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/users/{user}/addPublicKey",
	//   "request": {
	//     "$ref": "PublicKey"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud.useraccounts",
	//     "https://www.googleapis.com/auth/computeaccounts"
	//   ]
	// }

}

// method id "clouduseraccounts.users.delete":

type UsersDeleteCall struct {
	s       *Service
	project string
	user    string
	opt_    map[string]interface{}
}

// Delete: Deletes the specified User resource.
func (r *UsersService) Delete(project string, user string) *UsersDeleteCall {
	c := &UsersDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.user = user
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersDeleteCall) Fields(s ...googleapi.Field) *UsersDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *UsersDeleteCall) Do() (*Operation, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/users/{user}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
		"user":    c.user,
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
	//   "description": "Deletes the specified User resource.",
	//   "httpMethod": "DELETE",
	//   "id": "clouduseraccounts.users.delete",
	//   "parameterOrder": [
	//     "project",
	//     "user"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "user": {
	//       "description": "Name of the user resource to delete.",
	//       "location": "path",
	//       "pattern": "[a-z][-a-z0-9_]{0,31}",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/users/{user}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud.useraccounts",
	//     "https://www.googleapis.com/auth/computeaccounts"
	//   ]
	// }

}

// method id "clouduseraccounts.users.get":

type UsersGetCall struct {
	s       *Service
	project string
	user    string
	opt_    map[string]interface{}
}

// Get: Returns the specified User resource.
func (r *UsersService) Get(project string, user string) *UsersGetCall {
	c := &UsersGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.user = user
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersGetCall) Fields(s ...googleapi.Field) *UsersGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *UsersGetCall) Do() (*User, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/users/{user}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
		"user":    c.user,
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
	var ret *User
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the specified User resource.",
	//   "httpMethod": "GET",
	//   "id": "clouduseraccounts.users.get",
	//   "parameterOrder": [
	//     "project",
	//     "user"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "user": {
	//       "description": "Name of the user resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z][-a-z0-9_]{0,31}",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/users/{user}",
	//   "response": {
	//     "$ref": "User"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud.useraccounts",
	//     "https://www.googleapis.com/auth/cloud.useraccounts.readonly",
	//     "https://www.googleapis.com/auth/computeaccounts",
	//     "https://www.googleapis.com/auth/computeaccounts.readonly"
	//   ]
	// }

}

// method id "clouduseraccounts.users.insert":

type UsersInsertCall struct {
	s       *Service
	project string
	user    *User
	opt_    map[string]interface{}
}

// Insert: Creates a User resource in the specified project using the
// data included in the request.
func (r *UsersService) Insert(project string, user *User) *UsersInsertCall {
	c := &UsersInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.user = user
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersInsertCall) Fields(s ...googleapi.Field) *UsersInsertCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *UsersInsertCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.user)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/users")
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
	//   "description": "Creates a User resource in the specified project using the data included in the request.",
	//   "httpMethod": "POST",
	//   "id": "clouduseraccounts.users.insert",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/users",
	//   "request": {
	//     "$ref": "User"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud.useraccounts",
	//     "https://www.googleapis.com/auth/computeaccounts"
	//   ]
	// }

}

// method id "clouduseraccounts.users.list":

type UsersListCall struct {
	s       *Service
	project string
	opt_    map[string]interface{}
}

// List: Retrieves a list of users contained within the specified
// project.
func (r *UsersService) List(project string) *UsersListCall {
	c := &UsersListCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	return c
}

// Filter sets the optional parameter "filter": Sets a filter expression
// for filtering listed resources, in the form filter={expression}. Your
// {expression} must contain the following:
// FIELD_NAME COMPARISON_STRING LITERAL_STRING
//
// - FIELD_NAME: The name of the field you want to compare. The field
// name must be valid for the type of resource being filtered. Only
// atomic field types are supported (string, number, boolean). Array and
// object fields are not currently supported.
// - COMPARISON_STRING: The comparison string, either eq (equals) or ne
// (not equals).
// - LITERAL_STRING: The literal string value to filter to. The literal
// value must be valid for the type of field (string, number, boolean).
// For string fields, the literal value is interpreted as a regular
// expression using RE2 syntax. The literal value must match the entire
// field.  For example, you can filter by the name of a
// resource:
// filter=name ne example-instance
// The above filter returns only results whose name field does not equal
// example-instance. You can also enclose your literal string in single,
// double, or no quotes.
func (c *UsersListCall) Filter(filter string) *UsersListCall {
	c.opt_["filter"] = filter
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned.
func (c *UsersListCall) MaxResults(maxResults int64) *UsersListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// OrderBy sets the optional parameter "orderBy": Sorts list results by
// a certain order. By default, results are returned in alphanumerical
// order based on the resource name.
//
// You can also sort results in descending order based on the creation
// timestamp using orderBy="creationTimestamp desc". This sorts results
// based on the creationTimestamp field in reverse chronological order
// (newest result first). Use this to sort resources like operations so
// that the newest operation is returned first.
//
// Currently, only sorting by name or creationTimestamp desc is
// supported.
func (c *UsersListCall) OrderBy(orderBy string) *UsersListCall {
	c.opt_["orderBy"] = orderBy
	return c
}

// PageToken sets the optional parameter "pageToken": Specifies a page
// token to use. Use this parameter if you want to list the next page of
// results. Set pageToken to the nextPageToken returned by a previous
// list request.
func (c *UsersListCall) PageToken(pageToken string) *UsersListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersListCall) Fields(s ...googleapi.Field) *UsersListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *UsersListCall) Do() (*UserList, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["filter"]; ok {
		params.Set("filter", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["orderBy"]; ok {
		params.Set("orderBy", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/users")
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
	var ret *UserList
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of users contained within the specified project.",
	//   "httpMethod": "GET",
	//   "id": "clouduseraccounts.users.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Sets a filter expression for filtering listed resources, in the form filter={expression}. Your {expression} must contain the following:\nFIELD_NAME COMPARISON_STRING LITERAL_STRING\n \n- FIELD_NAME: The name of the field you want to compare. The field name must be valid for the type of resource being filtered. Only atomic field types are supported (string, number, boolean). Array and object fields are not currently supported. \n- COMPARISON_STRING: The comparison string, either eq (equals) or ne (not equals). \n- LITERAL_STRING: The literal string value to filter to. The literal value must be valid for the type of field (string, number, boolean). For string fields, the literal value is interpreted as a regular expression using RE2 syntax. The literal value must match the entire field.  For example, you can filter by the name of a resource:\nfilter=name ne example-instance\nThe above filter returns only results whose name field does not equal example-instance. You can also enclose your literal string in single, double, or no quotes.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Maximum count of results to be returned.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "orderBy": {
	//       "description": "Sorts list results by a certain order. By default, results are returned in alphanumerical order based on the resource name.\n\nYou can also sort results in descending order based on the creation timestamp using orderBy=\"creationTimestamp desc\". This sorts results based on the creationTimestamp field in reverse chronological order (newest result first). Use this to sort resources like operations so that the newest operation is returned first.\n\nCurrently, only sorting by name or creationTimestamp desc is supported.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Specifies a page token to use. Use this parameter if you want to list the next page of results. Set pageToken to the nextPageToken returned by a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/users",
	//   "response": {
	//     "$ref": "UserList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud.useraccounts",
	//     "https://www.googleapis.com/auth/cloud.useraccounts.readonly",
	//     "https://www.googleapis.com/auth/computeaccounts",
	//     "https://www.googleapis.com/auth/computeaccounts.readonly"
	//   ]
	// }

}

// method id "clouduseraccounts.users.removePublicKey":

type UsersRemovePublicKeyCall struct {
	s           *Service
	project     string
	user        string
	fingerprint string
	opt_        map[string]interface{}
}

// RemovePublicKey: Removes the specified public key from the user.
func (r *UsersService) RemovePublicKey(project string, user string, fingerprint string) *UsersRemovePublicKeyCall {
	c := &UsersRemovePublicKeyCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.user = user
	c.fingerprint = fingerprint
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersRemovePublicKeyCall) Fields(s ...googleapi.Field) *UsersRemovePublicKeyCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *UsersRemovePublicKeyCall) Do() (*Operation, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("fingerprint", fmt.Sprintf("%v", c.fingerprint))
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/users/{user}/removePublicKey")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
		"user":    c.user,
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
	//   "description": "Removes the specified public key from the user.",
	//   "httpMethod": "POST",
	//   "id": "clouduseraccounts.users.removePublicKey",
	//   "parameterOrder": [
	//     "project",
	//     "user",
	//     "fingerprint"
	//   ],
	//   "parameters": {
	//     "fingerprint": {
	//       "description": "The fingerprint of the public key to delete. Public keys are identified by their fingerprint, which is defined by RFC4716 to be the MD5 digest of the public key.",
	//       "location": "query",
	//       "pattern": "[a-f0-9]{32}",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "user": {
	//       "description": "Name of the user for this request.",
	//       "location": "path",
	//       "pattern": "[a-z][-a-z0-9_]{0,31}",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/users/{user}/removePublicKey",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud.useraccounts",
	//     "https://www.googleapis.com/auth/computeaccounts"
	//   ]
	// }

}
