// Package tagmanager provides access to the Tag Manager API.
//
// See https://developers.google.com/tag-manager/api/v1/
//
// Usage example:
//
//   import "google.golang.org/api/tagmanager/v1"
//   ...
//   tagmanagerService, err := tagmanager.New(oauthHttpClient)
package tagmanager

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

const apiId = "tagmanager:v1"
const apiName = "tagmanager"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/tagmanager/v1/"

// OAuth2 scopes used by this API.
const (
	// Delete your Google Tag Manager containers
	TagmanagerDeleteContainersScope = "https://www.googleapis.com/auth/tagmanager.delete.containers"

	// Manage your Google Tag Manager containers
	TagmanagerEditContainersScope = "https://www.googleapis.com/auth/tagmanager.edit.containers"

	// Manage your Google Tag Manager container versions
	TagmanagerEditContainerversionsScope = "https://www.googleapis.com/auth/tagmanager.edit.containerversions"

	// Manage your Google Tag Manager accounts
	TagmanagerManageAccountsScope = "https://www.googleapis.com/auth/tagmanager.manage.accounts"

	// Manage user permissions of your Google Tag Manager data
	TagmanagerManageUsersScope = "https://www.googleapis.com/auth/tagmanager.manage.users"

	// Publish your Google Tag Manager containers
	TagmanagerPublishScope = "https://www.googleapis.com/auth/tagmanager.publish"

	// View your Google Tag Manager containers
	TagmanagerReadonlyScope = "https://www.googleapis.com/auth/tagmanager.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Accounts = NewAccountsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Accounts *AccountsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewAccountsService(s *Service) *AccountsService {
	rs := &AccountsService{s: s}
	rs.Containers = NewAccountsContainersService(s)
	rs.Permissions = NewAccountsPermissionsService(s)
	return rs
}

type AccountsService struct {
	s *Service

	Containers *AccountsContainersService

	Permissions *AccountsPermissionsService
}

func NewAccountsContainersService(s *Service) *AccountsContainersService {
	rs := &AccountsContainersService{s: s}
	rs.Folders = NewAccountsContainersFoldersService(s)
	rs.Macros = NewAccountsContainersMacrosService(s)
	rs.MoveFolders = NewAccountsContainersMoveFoldersService(s)
	rs.Rules = NewAccountsContainersRulesService(s)
	rs.Tags = NewAccountsContainersTagsService(s)
	rs.Triggers = NewAccountsContainersTriggersService(s)
	rs.Variables = NewAccountsContainersVariablesService(s)
	rs.Versions = NewAccountsContainersVersionsService(s)
	return rs
}

type AccountsContainersService struct {
	s *Service

	Folders *AccountsContainersFoldersService

	Macros *AccountsContainersMacrosService

	MoveFolders *AccountsContainersMoveFoldersService

	Rules *AccountsContainersRulesService

	Tags *AccountsContainersTagsService

	Triggers *AccountsContainersTriggersService

	Variables *AccountsContainersVariablesService

	Versions *AccountsContainersVersionsService
}

func NewAccountsContainersFoldersService(s *Service) *AccountsContainersFoldersService {
	rs := &AccountsContainersFoldersService{s: s}
	rs.Entities = NewAccountsContainersFoldersEntitiesService(s)
	return rs
}

type AccountsContainersFoldersService struct {
	s *Service

	Entities *AccountsContainersFoldersEntitiesService
}

func NewAccountsContainersFoldersEntitiesService(s *Service) *AccountsContainersFoldersEntitiesService {
	rs := &AccountsContainersFoldersEntitiesService{s: s}
	return rs
}

type AccountsContainersFoldersEntitiesService struct {
	s *Service
}

func NewAccountsContainersMacrosService(s *Service) *AccountsContainersMacrosService {
	rs := &AccountsContainersMacrosService{s: s}
	return rs
}

type AccountsContainersMacrosService struct {
	s *Service
}

func NewAccountsContainersMoveFoldersService(s *Service) *AccountsContainersMoveFoldersService {
	rs := &AccountsContainersMoveFoldersService{s: s}
	return rs
}

type AccountsContainersMoveFoldersService struct {
	s *Service
}

func NewAccountsContainersRulesService(s *Service) *AccountsContainersRulesService {
	rs := &AccountsContainersRulesService{s: s}
	return rs
}

type AccountsContainersRulesService struct {
	s *Service
}

func NewAccountsContainersTagsService(s *Service) *AccountsContainersTagsService {
	rs := &AccountsContainersTagsService{s: s}
	return rs
}

type AccountsContainersTagsService struct {
	s *Service
}

func NewAccountsContainersTriggersService(s *Service) *AccountsContainersTriggersService {
	rs := &AccountsContainersTriggersService{s: s}
	return rs
}

type AccountsContainersTriggersService struct {
	s *Service
}

func NewAccountsContainersVariablesService(s *Service) *AccountsContainersVariablesService {
	rs := &AccountsContainersVariablesService{s: s}
	return rs
}

type AccountsContainersVariablesService struct {
	s *Service
}

func NewAccountsContainersVersionsService(s *Service) *AccountsContainersVersionsService {
	rs := &AccountsContainersVersionsService{s: s}
	return rs
}

type AccountsContainersVersionsService struct {
	s *Service
}

func NewAccountsPermissionsService(s *Service) *AccountsPermissionsService {
	rs := &AccountsPermissionsService{s: s}
	return rs
}

type AccountsPermissionsService struct {
	s *Service
}

// Account: Represents a Google Tag Manager Account.
type Account struct {
	// AccountId: The Account ID uniquely identifies the GTM Account.
	AccountId string `json:"accountId,omitempty"`

	// Fingerprint: The fingerprint of the GTM Account as computed at
	// storage time. This value is recomputed whenever the account is
	// modified.
	Fingerprint string `json:"fingerprint,omitempty"`

	// Name: Account display name.
	Name string `json:"name,omitempty"`

	// ShareData: Whether the account shares data anonymously with Google
	// and others.
	ShareData bool `json:"shareData,omitempty"`
}

// AccountAccess: Defines the Google Tag Manager Account access
// permissions.
type AccountAccess struct {
	// Permission: List of Account permissions. Valid account permissions
	// are read and manage.
	//
	// Possible values:
	//   "delete"
	//   "edit"
	//   "manage"
	//   "publish"
	//   "read"
	Permission []string `json:"permission,omitempty"`
}

// Condition: Represents a predicate.
type Condition struct {
	// Parameter: A list of named parameters (key/value), depending on the
	// condition's type. Notes:
	// - For binary operators, include parameters named arg0 and arg1 for
	// specifying the left and right operands, respectively.
	// - At this time, the left operand (arg0) must be a reference to a
	// macro.
	// - For case-insensitive Regex matching, include a boolean parameter
	// named ignore_case that is set to true. If not specified or set to any
	// other value, the matching will be case sensitive.
	// - To negate an operator, include a boolean parameter named negate
	// boolean parameter that is set to true.
	Parameter []*Parameter `json:"parameter,omitempty"`

	// Type: The type of operator for this condition.
	//
	// Possible values:
	//   "contains"
	//   "cssSelector"
	//   "endsWith"
	//   "equals"
	//   "greater"
	//   "greaterOrEquals"
	//   "less"
	//   "lessOrEquals"
	//   "matchRegex"
	//   "startsWith"
	//   "urlMatches"
	Type string `json:"type,omitempty"`
}

// Container: Represents a Google Tag Manager Container.
type Container struct {
	// AccountId: GTM Account ID.
	AccountId string `json:"accountId,omitempty"`

	// ContainerId: The Container ID uniquely identifies the GTM Container.
	ContainerId string `json:"containerId,omitempty"`

	// DomainName: Optional list of domain names associated with the
	// Container.
	DomainName []string `json:"domainName,omitempty"`

	// EnabledBuiltInVariable: List of enabled built-in variables. Valid
	// values include: pageUrl, pageHostname, pagePath, referrer, event,
	// clickElement, clickClasses, clickId, clickTarget, clickUrl,
	// clickText, formElement, formClasses, formId, formTarget, formUrl,
	// formText, errorMessage, errorUrl, errorLine, newHistoryFragment,
	// oldHistoryFragment, newHistoryState, oldHistoryState, historySource,
	// containerVersion, debugMode, randomNumber, containerId.
	//
	// Possible values:
	//   "advertiserId"
	//   "advertisingTrackingEnabled"
	//   "appId"
	//   "appName"
	//   "appVersionCode"
	//   "appVersionName"
	//   "clickClasses"
	//   "clickElement"
	//   "clickId"
	//   "clickTarget"
	//   "clickText"
	//   "clickUrl"
	//   "containerId"
	//   "containerVersion"
	//   "debugMode"
	//   "deviceName"
	//   "errorLine"
	//   "errorMessage"
	//   "errorUrl"
	//   "event"
	//   "formClasses"
	//   "formElement"
	//   "formId"
	//   "formTarget"
	//   "formText"
	//   "formUrl"
	//   "historySource"
	//   "htmlId"
	//   "language"
	//   "newHistoryFragment"
	//   "newHistoryState"
	//   "oldHistoryFragment"
	//   "oldHistoryState"
	//   "osVersion"
	//   "pageHostname"
	//   "pagePath"
	//   "pageUrl"
	//   "platform"
	//   "randomNumber"
	//   "referrer"
	//   "resolution"
	//   "sdkVersion"
	EnabledBuiltInVariable []string `json:"enabledBuiltInVariable,omitempty"`

	// Fingerprint: The fingerprint of the GTM Container as computed at
	// storage time. This value is recomputed whenever the account is
	// modified.
	Fingerprint string `json:"fingerprint,omitempty"`

	// Name: Container display name.
	Name string `json:"name,omitempty"`

	// Notes: Container Notes.
	Notes string `json:"notes,omitempty"`

	// PublicId: Container Public ID.
	PublicId string `json:"publicId,omitempty"`

	// TimeZoneCountryId: Container Country ID.
	TimeZoneCountryId string `json:"timeZoneCountryId,omitempty"`

	// TimeZoneId: Container Time Zone ID.
	TimeZoneId string `json:"timeZoneId,omitempty"`

	// UsageContext: List of Usage Contexts for the Container. Valid values
	// include: web, android, ios.
	//
	// Possible values:
	//   "android"
	//   "ios"
	//   "web"
	UsageContext []string `json:"usageContext,omitempty"`
}

// ContainerAccess: Defines the Google Tag Manager Container access
// permissions.
type ContainerAccess struct {
	// ContainerId: GTM Container ID.
	ContainerId string `json:"containerId,omitempty"`

	// Permission: List of Container permissions. Valid container
	// permissions are: read, edit, delete, publish.
	//
	// Possible values:
	//   "delete"
	//   "edit"
	//   "manage"
	//   "publish"
	//   "read"
	Permission []string `json:"permission,omitempty"`
}

// ContainerVersion: Represents a Google Tag Manager Container Version.
type ContainerVersion struct {
	// AccountId: GTM Account ID.
	AccountId string `json:"accountId,omitempty"`

	// Container: The container that this version was taken from.
	Container *Container `json:"container,omitempty"`

	// ContainerId: GTM Container ID.
	ContainerId string `json:"containerId,omitempty"`

	// ContainerVersionId: The Container Version ID uniquely identifies the
	// GTM Container Version.
	ContainerVersionId string `json:"containerVersionId,omitempty"`

	// Deleted: A value of true indicates this container version has been
	// deleted.
	Deleted bool `json:"deleted,omitempty"`

	// Fingerprint: The fingerprint of the GTM Container Version as computed
	// at storage time. This value is recomputed whenever the container
	// version is modified.
	Fingerprint string `json:"fingerprint,omitempty"`

	// Folder: The folders in the container that this version was taken
	// from.
	Folder []*Folder `json:"folder,omitempty"`

	// Macro: The macros in the container that this version was taken from.
	Macro []*Macro `json:"macro,omitempty"`

	// Name: Container version display name.
	Name string `json:"name,omitempty"`

	// Notes: User notes on how to apply this container version in the
	// container.
	Notes string `json:"notes,omitempty"`

	// Rule: The rules in the container that this version was taken from.
	Rule []*Rule `json:"rule,omitempty"`

	// Tag: The tags in the container that this version was taken from.
	Tag []*Tag `json:"tag,omitempty"`

	// Trigger: The triggers in the container that this version was taken
	// from.
	Trigger []*Trigger `json:"trigger,omitempty"`

	// Variable: The variables in the container that this version was taken
	// from.
	Variable []*Variable `json:"variable,omitempty"`
}

// ContainerVersionHeader: Represents a Google Tag Manager Container
// Version Header.
type ContainerVersionHeader struct {
	// AccountId: GTM Account ID.
	AccountId string `json:"accountId,omitempty"`

	// ContainerId: GTM Container ID.
	ContainerId string `json:"containerId,omitempty"`

	// ContainerVersionId: The Container Version ID uniquely identifies the
	// GTM Container Version.
	ContainerVersionId string `json:"containerVersionId,omitempty"`

	// Deleted: A value of true indicates this container version has been
	// deleted.
	Deleted bool `json:"deleted,omitempty"`

	// Name: Container version display name.
	Name string `json:"name,omitempty"`

	// NumMacros: Number of macros in the container version.
	NumMacros string `json:"numMacros,omitempty"`

	// NumRules: Number of rules in the container version.
	NumRules string `json:"numRules,omitempty"`

	// NumTags: Number of tags in the container version.
	NumTags string `json:"numTags,omitempty"`

	// NumTriggers: Number of triggers in the container version.
	NumTriggers string `json:"numTriggers,omitempty"`

	// NumVariables: Number of variables in the container version.
	NumVariables string `json:"numVariables,omitempty"`
}

// CreateContainerVersionRequestVersionOptions: Options for new
// container versions.
type CreateContainerVersionRequestVersionOptions struct {
	// Name: The name of the container version to be created.
	Name string `json:"name,omitempty"`

	// Notes: The notes of the container version to be created.
	Notes string `json:"notes,omitempty"`

	// QuickPreview: The creation of this version may be for quick preview
	// and shouldn't be saved.
	QuickPreview bool `json:"quickPreview,omitempty"`
}

// CreateContainerVersionResponse: Create container versions response.
type CreateContainerVersionResponse struct {
	// CompilerError: Compiler errors or not.
	CompilerError bool `json:"compilerError,omitempty"`

	// ContainerVersion: The container version created.
	ContainerVersion *ContainerVersion `json:"containerVersion,omitempty"`
}

// Folder: Represents a Google Tag Manager Folder.
type Folder struct {
	// AccountId: GTM Account ID.
	AccountId string `json:"accountId,omitempty"`

	// ContainerId: GTM Container ID.
	ContainerId string `json:"containerId,omitempty"`

	// Fingerprint: The fingerprint of the GTM Folder as computed at storage
	// time. This value is recomputed whenever the folder is modified.
	Fingerprint string `json:"fingerprint,omitempty"`

	// FolderId: The Folder ID uniquely identifies the GTM Folder.
	FolderId string `json:"folderId,omitempty"`

	// Name: Folder display name.
	Name string `json:"name,omitempty"`
}

// FolderEntities: Represents a Google Tag Manager Folder's contents.
type FolderEntities struct {
	// Tag: The list of tags inside the folder.
	Tag []*Tag `json:"tag,omitempty"`

	// Trigger: The list of triggers inside the folder.
	Trigger []*Trigger `json:"trigger,omitempty"`

	// Variable: The list of variables inside the folder.
	Variable []*Variable `json:"variable,omitempty"`
}

// ListAccountUsersResponse: List AccountUsers Response.
type ListAccountUsersResponse struct {
	// UserAccess: All GTM AccountUsers of a GTM Account.
	UserAccess []*UserAccess `json:"userAccess,omitempty"`
}

// ListAccountsResponse: List Accounts Response.
type ListAccountsResponse struct {
	// Accounts: List of GTM Accounts that a user has access to.
	Accounts []*Account `json:"accounts,omitempty"`
}

// ListContainerVersionsResponse: List container versions response.
type ListContainerVersionsResponse struct {
	// ContainerVersion: All versions of a GTM Container.
	ContainerVersion []*ContainerVersion `json:"containerVersion,omitempty"`

	// ContainerVersionHeader: All container version headers of a GTM
	// Container.
	ContainerVersionHeader []*ContainerVersionHeader `json:"containerVersionHeader,omitempty"`
}

// ListContainersResponse: List Containers Response.
type ListContainersResponse struct {
	// Containers: All Containers of a GTM Account.
	Containers []*Container `json:"containers,omitempty"`
}

// ListFoldersResponse: List Folders Response.
type ListFoldersResponse struct {
	// Folders: All GTM Folders of a GTM Container.
	Folders []*Folder `json:"folders,omitempty"`
}

// ListMacrosResponse: List Macros Response.
type ListMacrosResponse struct {
	// Macros: All GTM Macros of a GTM Container.
	Macros []*Macro `json:"macros,omitempty"`
}

// ListRulesResponse: List Rules Response.
type ListRulesResponse struct {
	// Rules: All GTM Rules of a GTM Container.
	Rules []*Rule `json:"rules,omitempty"`
}

// ListTagsResponse: List Tags Response.
type ListTagsResponse struct {
	// Tags: All GTM Tags of a GTM Container.
	Tags []*Tag `json:"tags,omitempty"`
}

// ListTriggersResponse: List triggers response.
type ListTriggersResponse struct {
	// Triggers: All GTM Triggers of a GTM Container.
	Triggers []*Trigger `json:"triggers,omitempty"`
}

// ListVariablesResponse: List Variables Response.
type ListVariablesResponse struct {
	// Variables: All GTM Variables of a GTM Container.
	Variables []*Variable `json:"variables,omitempty"`
}

// Macro: Represents a Google Tag Manager Macro.
type Macro struct {
	// AccountId: GTM Account ID.
	AccountId string `json:"accountId,omitempty"`

	// ContainerId: GTM Container ID.
	ContainerId string `json:"containerId,omitempty"`

	// DisablingRuleId: For mobile containers only: A list of rule IDs for
	// disabling conditional macros; the macro is enabled if one of the
	// enabling rules is true while all the disabling rules are false.
	// Treated as an unordered set.
	DisablingRuleId []string `json:"disablingRuleId,omitempty"`

	// EnablingRuleId: For mobile containers only: A list of rule IDs for
	// enabling conditional macros; the macro is enabled if one of the
	// enabling rules is true while all the disabling rules are false.
	// Treated as an unordered set.
	EnablingRuleId []string `json:"enablingRuleId,omitempty"`

	// Fingerprint: The fingerprint of the GTM Macro as computed at storage
	// time. This value is recomputed whenever the macro is modified.
	Fingerprint string `json:"fingerprint,omitempty"`

	// MacroId: The Macro ID uniquely identifies the GTM Macro.
	MacroId string `json:"macroId,omitempty"`

	// Name: Macro display name.
	Name string `json:"name,omitempty"`

	// Notes: User notes on how to apply this macro in the container.
	Notes string `json:"notes,omitempty"`

	// Parameter: The macro's parameters.
	Parameter []*Parameter `json:"parameter,omitempty"`

	// ParentFolderId: Parent folder id.
	ParentFolderId string `json:"parentFolderId,omitempty"`

	// ScheduleEndMs: The end timestamp in milliseconds to schedule a macro.
	ScheduleEndMs int64 `json:"scheduleEndMs,omitempty,string"`

	// ScheduleStartMs: The start timestamp in milliseconds to schedule a
	// macro.
	ScheduleStartMs int64 `json:"scheduleStartMs,omitempty,string"`

	// Type: GTM Macro Type.
	Type string `json:"type,omitempty"`
}

// Parameter: Represents a Google Tag Manager Parameter.
type Parameter struct {
	// Key: The named key that uniquely identifies a parameter. Required for
	// top-level parameters, as well as map values. Ignored for list values.
	Key string `json:"key,omitempty"`

	// List: This list parameter's parameters (keys will be ignored).
	List []*Parameter `json:"list,omitempty"`

	// Map: This map parameter's parameters (must have keys; keys must be
	// unique).
	Map []*Parameter `json:"map,omitempty"`

	// Type: The parameter type. Valid values are:
	// - boolean: The value represents a boolean, represented as 'true' or
	// 'false'
	// - integer: The value represents a 64-bit signed integer value, in
	// base 10
	// - list: A list of parameters should be specified
	// - map: A map of parameters should be specified
	// - template: The value represents any text; this can include macro
	// references (even macro references that might return non-string types)
	//
	// Possible values:
	//   "boolean"
	//   "integer"
	//   "list"
	//   "map"
	//   "template"
	Type string `json:"type,omitempty"`

	// Value: A parameter's value (may contain macro references such as
	// "{{myMacro}}") as appropriate to the specified type.
	Value string `json:"value,omitempty"`
}

// PublishContainerVersionResponse: Publish container version response.
type PublishContainerVersionResponse struct {
	// CompilerError: Compiler errors or not.
	CompilerError bool `json:"compilerError,omitempty"`

	// ContainerVersion: The container version created.
	ContainerVersion *ContainerVersion `json:"containerVersion,omitempty"`
}

// Rule: Represents a Google Tag Manager Rule.
type Rule struct {
	// AccountId: GTM Account ID.
	AccountId string `json:"accountId,omitempty"`

	// Condition: The list of conditions that make up this rule (implicit
	// AND between them).
	Condition []*Condition `json:"condition,omitempty"`

	// ContainerId: GTM Container ID.
	ContainerId string `json:"containerId,omitempty"`

	// Fingerprint: The fingerprint of the GTM Rule as computed at storage
	// time. This value is recomputed whenever the rule is modified.
	Fingerprint string `json:"fingerprint,omitempty"`

	// Name: Rule display name.
	Name string `json:"name,omitempty"`

	// Notes: User notes on how to apply this rule in the container.
	Notes string `json:"notes,omitempty"`

	// RuleId: The Rule ID uniquely identifies the GTM Rule.
	RuleId string `json:"ruleId,omitempty"`
}

type SetupTag struct {
	// StopOnSetupFailure: If true, fire the main tag if and only if the
	// setup tag fires successfully. If false, fire the main tag regardless
	// of setup tag firing status.
	StopOnSetupFailure bool `json:"stopOnSetupFailure,omitempty"`

	// TagName: The name of the setup tag.
	TagName string `json:"tagName,omitempty"`
}

// Tag: Represents a Google Tag Manager Tag.
type Tag struct {
	// AccountId: GTM Account ID.
	AccountId string `json:"accountId,omitempty"`

	// BlockingRuleId: Blocking rule IDs. If any of the listed rules
	// evaluate to true, the tag will not fire.
	BlockingRuleId []string `json:"blockingRuleId,omitempty"`

	// BlockingTriggerId: Blocking trigger IDs. If any of the listed
	// triggers evaluate to true, the tag will not fire.
	BlockingTriggerId []string `json:"blockingTriggerId,omitempty"`

	// ContainerId: GTM Container ID.
	ContainerId string `json:"containerId,omitempty"`

	// Fingerprint: The fingerprint of the GTM Tag as computed at storage
	// time. This value is recomputed whenever the tag is modified.
	Fingerprint string `json:"fingerprint,omitempty"`

	// FiringRuleId: Firing rule IDs. A tag will fire when any of the listed
	// rules are true and all of its blockingRuleIds (if any specified) are
	// false.
	FiringRuleId []string `json:"firingRuleId,omitempty"`

	// FiringTriggerId: Firing trigger IDs. A tag will fire when any of the
	// listed triggers are true and all of its blockingTriggerIds (if any
	// specified) are false.
	FiringTriggerId []string `json:"firingTriggerId,omitempty"`

	// LiveOnly: If set to true, this tag will only fire in the live
	// environment (e.g. not in preview or debug mode).
	LiveOnly bool `json:"liveOnly,omitempty"`

	// Name: Tag display name.
	Name string `json:"name,omitempty"`

	// Notes: User notes on how to apply this tag in the container.
	Notes string `json:"notes,omitempty"`

	// Parameter: The tag's parameters.
	Parameter []*Parameter `json:"parameter,omitempty"`

	// ParentFolderId: Parent folder id.
	ParentFolderId string `json:"parentFolderId,omitempty"`

	// Priority: User defined numeric priority of the tag. Tags are fired
	// asynchronously in order of priority. Tags with higher numeric value
	// fire first. A tag's priority can be a positive or negative value. The
	// default value is 0.
	Priority *Parameter `json:"priority,omitempty"`

	// ScheduleEndMs: The end timestamp in milliseconds to schedule a tag.
	ScheduleEndMs int64 `json:"scheduleEndMs,omitempty,string"`

	// ScheduleStartMs: The start timestamp in milliseconds to schedule a
	// tag.
	ScheduleStartMs int64 `json:"scheduleStartMs,omitempty,string"`

	// SetupTag: The list of setup tags. Currently we only allow one.
	SetupTag []*SetupTag `json:"setupTag,omitempty"`

	// TagId: The Tag ID uniquely identifies the GTM Tag.
	TagId string `json:"tagId,omitempty"`

	// TeardownTag: The list of teardown tags. Currently we only allow one.
	TeardownTag []*TeardownTag `json:"teardownTag,omitempty"`

	// Type: GTM Tag Type.
	Type string `json:"type,omitempty"`
}

type TeardownTag struct {
	// StopTeardownOnFailure: If true, fire the teardown tag if and only if
	// the main tag fires successfully. If false, fire the teardown tag
	// regardless of main tag firing status.
	StopTeardownOnFailure bool `json:"stopTeardownOnFailure,omitempty"`

	// TagName: The name of the teardown tag.
	TagName string `json:"tagName,omitempty"`
}

// Trigger: Represents a Google Tag Manager Trigger
type Trigger struct {
	// AccountId: GTM Account ID.
	AccountId string `json:"accountId,omitempty"`

	// AutoEventFilter: Used in the case of auto event tracking.
	AutoEventFilter []*Condition `json:"autoEventFilter,omitempty"`

	// CheckValidation: Whether or not we should only fire tags if the form
	// submit or link click event is not cancelled by some other event
	// handler (e.g. because of validation). Only valid for Form Submission
	// and Link Click triggers.
	CheckValidation *Parameter `json:"checkValidation,omitempty"`

	// ContainerId: GTM Container ID.
	ContainerId string `json:"containerId,omitempty"`

	// CustomEventFilter: Used in the case of custom event, which is fired
	// iff all Conditions are true.
	CustomEventFilter []*Condition `json:"customEventFilter,omitempty"`

	// EnableAllVideos: Reloads the videos in the page that don't already
	// have the YT API enabled. If false, only capture events from videos
	// that already have the API enabled. Only valid for YouTube triggers.
	EnableAllVideos *Parameter `json:"enableAllVideos,omitempty"`

	// EventName: Name of the GTM event that is fired. Only valid for Timer
	// triggers.
	EventName *Parameter `json:"eventName,omitempty"`

	// Filter: The trigger will only fire iff all Conditions are true.
	Filter []*Condition `json:"filter,omitempty"`

	// Fingerprint: The fingerprint of the GTM Trigger as computed at
	// storage time. This value is recomputed whenever the trigger is
	// modified.
	Fingerprint string `json:"fingerprint,omitempty"`

	// Interval: Time between triggering recurring Timer Events (in
	// milliseconds). Only valid for Timer triggers.
	Interval *Parameter `json:"interval,omitempty"`

	// Limit: Limit of the number of GTM events this Timer Trigger will
	// fire. If no limit is set, we will continue to fire GTM events until
	// the user leaves the page. Only valid for Timer triggers.
	Limit *Parameter `json:"limit,omitempty"`

	// Name: Trigger display name.
	Name string `json:"name,omitempty"`

	// ParentFolderId: Parent folder id.
	ParentFolderId string `json:"parentFolderId,omitempty"`

	// TriggerId: The Trigger ID uniquely identifies the GTM Trigger.
	TriggerId string `json:"triggerId,omitempty"`

	// Type: Defines the data layer event that causes this trigger.
	//
	// Possible values:
	//   "ajaxSubmission"
	//   "always"
	//   "click"
	//   "customEvent"
	//   "domReady"
	//   "formSubmission"
	//   "historyChange"
	//   "jsError"
	//   "linkClick"
	//   "pageview"
	//   "timer"
	//   "windowLoaded"
	//   "youTube"
	Type string `json:"type,omitempty"`

	// UniqueTriggerId: Globally unique id of the trigger that
	// auto-generates this (a Form Submit, Link Click or Timer listener) if
	// any. Used to make incompatible auto-events work together with trigger
	// filtering based on trigger ids. This value is populated during output
	// generation since the tags implied by triggers don't exist until then.
	// Only valid for Form Submit, Link Click and Timer triggers.
	UniqueTriggerId *Parameter `json:"uniqueTriggerId,omitempty"`

	// VideoPercentageList: List of integer percentage values. The trigger
	// will fire as each percentage is reached in any instrumented videos.
	// Only valid for YouTube triggers.
	VideoPercentageList *Parameter `json:"videoPercentageList,omitempty"`

	// WaitForTags: Whether or not we should delay the form submissions or
	// link opening until all of the tags have fired (by preventing the
	// default action and later simulating the default action). Only valid
	// for Form Submission and Link Click triggers.
	WaitForTags *Parameter `json:"waitForTags,omitempty"`

	// WaitForTagsTimeout: How long to wait (in milliseconds) for tags to
	// fire when 'waits_for_tags' above evaluates to true. Only valid for
	// Form Submission and Link Click triggers.
	WaitForTagsTimeout *Parameter `json:"waitForTagsTimeout,omitempty"`
}

// UserAccess: Represents a user's permissions to an account and its
// container.
type UserAccess struct {
	// AccountAccess: GTM Account access permissions.
	AccountAccess *AccountAccess `json:"accountAccess,omitempty"`

	// AccountId: GTM Account ID.
	AccountId string `json:"accountId,omitempty"`

	// ContainerAccess: GTM Container access permissions.
	ContainerAccess []*ContainerAccess `json:"containerAccess,omitempty"`

	// EmailAddress: User's email address.
	EmailAddress string `json:"emailAddress,omitempty"`

	// PermissionId: Account Permission ID.
	PermissionId string `json:"permissionId,omitempty"`
}

// Variable: Represents a Google Tag Manager Variable.
type Variable struct {
	// AccountId: GTM Account ID.
	AccountId string `json:"accountId,omitempty"`

	// ContainerId: GTM Container ID.
	ContainerId string `json:"containerId,omitempty"`

	// DisablingTriggerId: For mobile containers only: A list of trigger IDs
	// for disabling conditional variables; the variable is enabled if one
	// of the enabling trigger is true while all the disabling trigger are
	// false. Treated as an unordered set.
	DisablingTriggerId []string `json:"disablingTriggerId,omitempty"`

	// EnablingTriggerId: For mobile containers only: A list of trigger IDs
	// for enabling conditional variables; the variable is enabled if one of
	// the enabling triggers is true while all the disabling triggers are
	// false. Treated as an unordered set.
	EnablingTriggerId []string `json:"enablingTriggerId,omitempty"`

	// Fingerprint: The fingerprint of the GTM Variable as computed at
	// storage time. This value is recomputed whenever the variable is
	// modified.
	Fingerprint string `json:"fingerprint,omitempty"`

	// Name: Variable display name.
	Name string `json:"name,omitempty"`

	// Notes: User notes on how to apply this variable in the container.
	Notes string `json:"notes,omitempty"`

	// Parameter: The variable's parameters.
	Parameter []*Parameter `json:"parameter,omitempty"`

	// ParentFolderId: Parent folder id.
	ParentFolderId string `json:"parentFolderId,omitempty"`

	// ScheduleEndMs: The end timestamp in milliseconds to schedule a
	// variable.
	ScheduleEndMs int64 `json:"scheduleEndMs,omitempty,string"`

	// ScheduleStartMs: The start timestamp in milliseconds to schedule a
	// variable.
	ScheduleStartMs int64 `json:"scheduleStartMs,omitempty,string"`

	// Type: GTM Variable Type.
	Type string `json:"type,omitempty"`

	// VariableId: The Variable ID uniquely identifies the GTM Variable.
	VariableId string `json:"variableId,omitempty"`
}

// method id "tagmanager.accounts.get":

type AccountsGetCall struct {
	s         *Service
	accountId string
	opt_      map[string]interface{}
}

// Get: Gets a GTM Account.
func (r *AccountsService) Get(accountId string) *AccountsGetCall {
	c := &AccountsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsGetCall) Fields(s ...googleapi.Field) *AccountsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsGetCall) IfNoneMatch(entityTag string) *AccountsGetCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.get" call.
// Exactly one of the return values is non-nil.
func (c *AccountsGetCall) Do() (*Account, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId": c.accountId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.get" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsGetCall) DoHeader() (ret *Account, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Gets a GTM Account.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.get",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}",
	//   "response": {
	//     "$ref": "Account"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.manage.accounts",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.list":

type AccountsListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Lists all GTM Accounts that a user has access to.
func (r *AccountsService) List() *AccountsListCall {
	c := &AccountsListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsListCall) Fields(s ...googleapi.Field) *AccountsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsListCall) IfNoneMatch(entityTag string) *AccountsListCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.list" call.
// Exactly one of the return values is non-nil.
func (c *AccountsListCall) Do() (*ListAccountsResponse, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.list" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsListCall) DoHeader() (ret *ListAccountsResponse, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Lists all GTM Accounts that a user has access to.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.list",
	//   "path": "accounts",
	//   "response": {
	//     "$ref": "ListAccountsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.manage.accounts",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.update":

type AccountsUpdateCall struct {
	s         *Service
	accountId string
	account   *Account
	opt_      map[string]interface{}
}

// Update: Updates a GTM Account.
func (r *AccountsService) Update(accountId string, account *Account) *AccountsUpdateCall {
	c := &AccountsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.account = account
	return c
}

// Fingerprint sets the optional parameter "fingerprint": When provided,
// this fingerprint must match the fingerprint of the account in
// storage.
func (c *AccountsUpdateCall) Fingerprint(fingerprint string) *AccountsUpdateCall {
	c.opt_["fingerprint"] = fingerprint
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsUpdateCall) Fields(s ...googleapi.Field) *AccountsUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsUpdateCall) IfNoneMatch(entityTag string) *AccountsUpdateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.update" call.
// Exactly one of the return values is non-nil.
func (c *AccountsUpdateCall) Do() (*Account, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.account)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fingerprint"]; ok {
		params.Set("fingerprint", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId": c.accountId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.update" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsUpdateCall) DoHeader() (ret *Account, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Updates a GTM Account.",
	//   "httpMethod": "PUT",
	//   "id": "tagmanager.accounts.update",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "fingerprint": {
	//       "description": "When provided, this fingerprint must match the fingerprint of the account in storage.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}",
	//   "request": {
	//     "$ref": "Account"
	//   },
	//   "response": {
	//     "$ref": "Account"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.manage.accounts"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.create":

type AccountsContainersCreateCall struct {
	s         *Service
	accountId string
	container *Container
	opt_      map[string]interface{}
}

// Create: Creates a Container.
func (r *AccountsContainersService) Create(accountId string, container *Container) *AccountsContainersCreateCall {
	c := &AccountsContainersCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.container = container
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersCreateCall) Fields(s ...googleapi.Field) *AccountsContainersCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersCreateCall) IfNoneMatch(entityTag string) *AccountsContainersCreateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.create" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersCreateCall) Do() (*Container, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersCreateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.container)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId": c.accountId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.create" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersCreateCall) DoHeader() (ret *Container, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Creates a Container.",
	//   "httpMethod": "POST",
	//   "id": "tagmanager.accounts.containers.create",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers",
	//   "request": {
	//     "$ref": "Container"
	//   },
	//   "response": {
	//     "$ref": "Container"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.delete":

type AccountsContainersDeleteCall struct {
	s           *Service
	accountId   string
	containerId string
	opt_        map[string]interface{}
}

// Delete: Deletes a Container.
func (r *AccountsContainersService) Delete(accountId string, containerId string) *AccountsContainersDeleteCall {
	c := &AccountsContainersDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersDeleteCall) Fields(s ...googleapi.Field) *AccountsContainersDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersDeleteCall) IfNoneMatch(entityTag string) *AccountsContainersDeleteCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.delete" call.
func (c *AccountsContainersDeleteCall) Do() error {
	_, err := c.DoHeader()
	return err
}

func (c *AccountsContainersDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.delete" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersDeleteCall) DoHeader() (resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return res.Header, err
	}
	return res.Header, nil
	// {
	//   "description": "Deletes a Container.",
	//   "httpMethod": "DELETE",
	//   "id": "tagmanager.accounts.containers.delete",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.delete.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.get":

type AccountsContainersGetCall struct {
	s           *Service
	accountId   string
	containerId string
	opt_        map[string]interface{}
}

// Get: Gets a Container.
func (r *AccountsContainersService) Get(accountId string, containerId string) *AccountsContainersGetCall {
	c := &AccountsContainersGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersGetCall) Fields(s ...googleapi.Field) *AccountsContainersGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersGetCall) IfNoneMatch(entityTag string) *AccountsContainersGetCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.get" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersGetCall) Do() (*Container, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.get" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersGetCall) DoHeader() (ret *Container, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Gets a Container.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}",
	//   "response": {
	//     "$ref": "Container"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.list":

type AccountsContainersListCall struct {
	s         *Service
	accountId string
	opt_      map[string]interface{}
}

// List: Lists all Containers that belongs to a GTM Account.
func (r *AccountsContainersService) List(accountId string) *AccountsContainersListCall {
	c := &AccountsContainersListCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersListCall) Fields(s ...googleapi.Field) *AccountsContainersListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersListCall) IfNoneMatch(entityTag string) *AccountsContainersListCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.list" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersListCall) Do() (*ListContainersResponse, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId": c.accountId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.list" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersListCall) DoHeader() (ret *ListContainersResponse, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Lists all Containers that belongs to a GTM Account.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.list",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers",
	//   "response": {
	//     "$ref": "ListContainersResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.update":

type AccountsContainersUpdateCall struct {
	s           *Service
	accountId   string
	containerId string
	container   *Container
	opt_        map[string]interface{}
}

// Update: Updates a Container.
func (r *AccountsContainersService) Update(accountId string, containerId string, container *Container) *AccountsContainersUpdateCall {
	c := &AccountsContainersUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.container = container
	return c
}

// Fingerprint sets the optional parameter "fingerprint": When provided,
// this fingerprint must match the fingerprint of the container in
// storage.
func (c *AccountsContainersUpdateCall) Fingerprint(fingerprint string) *AccountsContainersUpdateCall {
	c.opt_["fingerprint"] = fingerprint
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersUpdateCall) Fields(s ...googleapi.Field) *AccountsContainersUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersUpdateCall) IfNoneMatch(entityTag string) *AccountsContainersUpdateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.update" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersUpdateCall) Do() (*Container, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.container)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fingerprint"]; ok {
		params.Set("fingerprint", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.update" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersUpdateCall) DoHeader() (ret *Container, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Updates a Container.",
	//   "httpMethod": "PUT",
	//   "id": "tagmanager.accounts.containers.update",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "fingerprint": {
	//       "description": "When provided, this fingerprint must match the fingerprint of the container in storage.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}",
	//   "request": {
	//     "$ref": "Container"
	//   },
	//   "response": {
	//     "$ref": "Container"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.folders.create":

type AccountsContainersFoldersCreateCall struct {
	s           *Service
	accountId   string
	containerId string
	folder      *Folder
	opt_        map[string]interface{}
}

// Create: Creates a GTM Folder.
func (r *AccountsContainersFoldersService) Create(accountId string, containerId string, folder *Folder) *AccountsContainersFoldersCreateCall {
	c := &AccountsContainersFoldersCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.folder = folder
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersFoldersCreateCall) Fields(s ...googleapi.Field) *AccountsContainersFoldersCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersFoldersCreateCall) IfNoneMatch(entityTag string) *AccountsContainersFoldersCreateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.folders.create" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersFoldersCreateCall) Do() (*Folder, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersFoldersCreateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.folder)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/folders")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.folders.create" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersFoldersCreateCall) DoHeader() (ret *Folder, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Creates a GTM Folder.",
	//   "httpMethod": "POST",
	//   "id": "tagmanager.accounts.containers.folders.create",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/folders",
	//   "request": {
	//     "$ref": "Folder"
	//   },
	//   "response": {
	//     "$ref": "Folder"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.folders.delete":

type AccountsContainersFoldersDeleteCall struct {
	s           *Service
	accountId   string
	containerId string
	folderId    string
	opt_        map[string]interface{}
}

// Delete: Deletes a GTM Folder.
func (r *AccountsContainersFoldersService) Delete(accountId string, containerId string, folderId string) *AccountsContainersFoldersDeleteCall {
	c := &AccountsContainersFoldersDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.folderId = folderId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersFoldersDeleteCall) Fields(s ...googleapi.Field) *AccountsContainersFoldersDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersFoldersDeleteCall) IfNoneMatch(entityTag string) *AccountsContainersFoldersDeleteCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.folders.delete" call.
func (c *AccountsContainersFoldersDeleteCall) Do() error {
	_, err := c.DoHeader()
	return err
}

func (c *AccountsContainersFoldersDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/folders/{folderId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"folderId":    c.folderId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.folders.delete" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersFoldersDeleteCall) DoHeader() (resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return res.Header, err
	}
	return res.Header, nil
	// {
	//   "description": "Deletes a GTM Folder.",
	//   "httpMethod": "DELETE",
	//   "id": "tagmanager.accounts.containers.folders.delete",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "folderId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "folderId": {
	//       "description": "The GTM Folder ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/folders/{folderId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.folders.get":

type AccountsContainersFoldersGetCall struct {
	s           *Service
	accountId   string
	containerId string
	folderId    string
	opt_        map[string]interface{}
}

// Get: Gets a GTM Folder.
func (r *AccountsContainersFoldersService) Get(accountId string, containerId string, folderId string) *AccountsContainersFoldersGetCall {
	c := &AccountsContainersFoldersGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.folderId = folderId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersFoldersGetCall) Fields(s ...googleapi.Field) *AccountsContainersFoldersGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersFoldersGetCall) IfNoneMatch(entityTag string) *AccountsContainersFoldersGetCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.folders.get" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersFoldersGetCall) Do() (*Folder, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersFoldersGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/folders/{folderId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"folderId":    c.folderId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.folders.get" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersFoldersGetCall) DoHeader() (ret *Folder, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Gets a GTM Folder.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.folders.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "folderId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "folderId": {
	//       "description": "The GTM Folder ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/folders/{folderId}",
	//   "response": {
	//     "$ref": "Folder"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.folders.list":

type AccountsContainersFoldersListCall struct {
	s           *Service
	accountId   string
	containerId string
	opt_        map[string]interface{}
}

// List: Lists all GTM Folders of a Container.
func (r *AccountsContainersFoldersService) List(accountId string, containerId string) *AccountsContainersFoldersListCall {
	c := &AccountsContainersFoldersListCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersFoldersListCall) Fields(s ...googleapi.Field) *AccountsContainersFoldersListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersFoldersListCall) IfNoneMatch(entityTag string) *AccountsContainersFoldersListCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.folders.list" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersFoldersListCall) Do() (*ListFoldersResponse, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersFoldersListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/folders")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.folders.list" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersFoldersListCall) DoHeader() (ret *ListFoldersResponse, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Lists all GTM Folders of a Container.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.folders.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/folders",
	//   "response": {
	//     "$ref": "ListFoldersResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.folders.update":

type AccountsContainersFoldersUpdateCall struct {
	s           *Service
	accountId   string
	containerId string
	folderId    string
	folder      *Folder
	opt_        map[string]interface{}
}

// Update: Updates a GTM Folder.
func (r *AccountsContainersFoldersService) Update(accountId string, containerId string, folderId string, folder *Folder) *AccountsContainersFoldersUpdateCall {
	c := &AccountsContainersFoldersUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.folderId = folderId
	c.folder = folder
	return c
}

// Fingerprint sets the optional parameter "fingerprint": When provided,
// this fingerprint must match the fingerprint of the folder in storage.
func (c *AccountsContainersFoldersUpdateCall) Fingerprint(fingerprint string) *AccountsContainersFoldersUpdateCall {
	c.opt_["fingerprint"] = fingerprint
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersFoldersUpdateCall) Fields(s ...googleapi.Field) *AccountsContainersFoldersUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersFoldersUpdateCall) IfNoneMatch(entityTag string) *AccountsContainersFoldersUpdateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.folders.update" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersFoldersUpdateCall) Do() (*Folder, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersFoldersUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.folder)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fingerprint"]; ok {
		params.Set("fingerprint", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/folders/{folderId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"folderId":    c.folderId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.folders.update" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersFoldersUpdateCall) DoHeader() (ret *Folder, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Updates a GTM Folder.",
	//   "httpMethod": "PUT",
	//   "id": "tagmanager.accounts.containers.folders.update",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "folderId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "fingerprint": {
	//       "description": "When provided, this fingerprint must match the fingerprint of the folder in storage.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "folderId": {
	//       "description": "The GTM Folder ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/folders/{folderId}",
	//   "request": {
	//     "$ref": "Folder"
	//   },
	//   "response": {
	//     "$ref": "Folder"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.folders.entities.list":

type AccountsContainersFoldersEntitiesListCall struct {
	s           *Service
	accountId   string
	containerId string
	folderId    string
	opt_        map[string]interface{}
}

// List: List all entities in a GTM Folder.
func (r *AccountsContainersFoldersEntitiesService) List(accountId string, containerId string, folderId string) *AccountsContainersFoldersEntitiesListCall {
	c := &AccountsContainersFoldersEntitiesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.folderId = folderId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersFoldersEntitiesListCall) Fields(s ...googleapi.Field) *AccountsContainersFoldersEntitiesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersFoldersEntitiesListCall) IfNoneMatch(entityTag string) *AccountsContainersFoldersEntitiesListCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.folders.entities.list" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersFoldersEntitiesListCall) Do() (*FolderEntities, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersFoldersEntitiesListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/folders/{folderId}/entities")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"folderId":    c.folderId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.folders.entities.list" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersFoldersEntitiesListCall) DoHeader() (ret *FolderEntities, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "List all entities in a GTM Folder.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.folders.entities.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "folderId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "folderId": {
	//       "description": "The GTM Folder ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/folders/{folderId}/entities",
	//   "response": {
	//     "$ref": "FolderEntities"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.macros.create":

type AccountsContainersMacrosCreateCall struct {
	s           *Service
	accountId   string
	containerId string
	macro       *Macro
	opt_        map[string]interface{}
}

// Create: Creates a GTM Macro.
func (r *AccountsContainersMacrosService) Create(accountId string, containerId string, macro *Macro) *AccountsContainersMacrosCreateCall {
	c := &AccountsContainersMacrosCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.macro = macro
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersMacrosCreateCall) Fields(s ...googleapi.Field) *AccountsContainersMacrosCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersMacrosCreateCall) IfNoneMatch(entityTag string) *AccountsContainersMacrosCreateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.macros.create" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersMacrosCreateCall) Do() (*Macro, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersMacrosCreateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.macro)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/macros")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.macros.create" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersMacrosCreateCall) DoHeader() (ret *Macro, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Creates a GTM Macro.",
	//   "httpMethod": "POST",
	//   "id": "tagmanager.accounts.containers.macros.create",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/macros",
	//   "request": {
	//     "$ref": "Macro"
	//   },
	//   "response": {
	//     "$ref": "Macro"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.macros.delete":

type AccountsContainersMacrosDeleteCall struct {
	s           *Service
	accountId   string
	containerId string
	macroId     string
	opt_        map[string]interface{}
}

// Delete: Deletes a GTM Macro.
func (r *AccountsContainersMacrosService) Delete(accountId string, containerId string, macroId string) *AccountsContainersMacrosDeleteCall {
	c := &AccountsContainersMacrosDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.macroId = macroId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersMacrosDeleteCall) Fields(s ...googleapi.Field) *AccountsContainersMacrosDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersMacrosDeleteCall) IfNoneMatch(entityTag string) *AccountsContainersMacrosDeleteCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.macros.delete" call.
func (c *AccountsContainersMacrosDeleteCall) Do() error {
	_, err := c.DoHeader()
	return err
}

func (c *AccountsContainersMacrosDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/macros/{macroId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"macroId":     c.macroId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.macros.delete" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersMacrosDeleteCall) DoHeader() (resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return res.Header, err
	}
	return res.Header, nil
	// {
	//   "description": "Deletes a GTM Macro.",
	//   "httpMethod": "DELETE",
	//   "id": "tagmanager.accounts.containers.macros.delete",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "macroId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "macroId": {
	//       "description": "The GTM Macro ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/macros/{macroId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.macros.get":

type AccountsContainersMacrosGetCall struct {
	s           *Service
	accountId   string
	containerId string
	macroId     string
	opt_        map[string]interface{}
}

// Get: Gets a GTM Macro.
func (r *AccountsContainersMacrosService) Get(accountId string, containerId string, macroId string) *AccountsContainersMacrosGetCall {
	c := &AccountsContainersMacrosGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.macroId = macroId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersMacrosGetCall) Fields(s ...googleapi.Field) *AccountsContainersMacrosGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersMacrosGetCall) IfNoneMatch(entityTag string) *AccountsContainersMacrosGetCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.macros.get" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersMacrosGetCall) Do() (*Macro, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersMacrosGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/macros/{macroId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"macroId":     c.macroId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.macros.get" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersMacrosGetCall) DoHeader() (ret *Macro, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Gets a GTM Macro.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.macros.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "macroId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "macroId": {
	//       "description": "The GTM Macro ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/macros/{macroId}",
	//   "response": {
	//     "$ref": "Macro"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.macros.list":

type AccountsContainersMacrosListCall struct {
	s           *Service
	accountId   string
	containerId string
	opt_        map[string]interface{}
}

// List: Lists all GTM Macros of a Container.
func (r *AccountsContainersMacrosService) List(accountId string, containerId string) *AccountsContainersMacrosListCall {
	c := &AccountsContainersMacrosListCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersMacrosListCall) Fields(s ...googleapi.Field) *AccountsContainersMacrosListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersMacrosListCall) IfNoneMatch(entityTag string) *AccountsContainersMacrosListCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.macros.list" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersMacrosListCall) Do() (*ListMacrosResponse, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersMacrosListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/macros")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.macros.list" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersMacrosListCall) DoHeader() (ret *ListMacrosResponse, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Lists all GTM Macros of a Container.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.macros.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/macros",
	//   "response": {
	//     "$ref": "ListMacrosResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.macros.update":

type AccountsContainersMacrosUpdateCall struct {
	s           *Service
	accountId   string
	containerId string
	macroId     string
	macro       *Macro
	opt_        map[string]interface{}
}

// Update: Updates a GTM Macro.
func (r *AccountsContainersMacrosService) Update(accountId string, containerId string, macroId string, macro *Macro) *AccountsContainersMacrosUpdateCall {
	c := &AccountsContainersMacrosUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.macroId = macroId
	c.macro = macro
	return c
}

// Fingerprint sets the optional parameter "fingerprint": When provided,
// this fingerprint must match the fingerprint of the macro in storage.
func (c *AccountsContainersMacrosUpdateCall) Fingerprint(fingerprint string) *AccountsContainersMacrosUpdateCall {
	c.opt_["fingerprint"] = fingerprint
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersMacrosUpdateCall) Fields(s ...googleapi.Field) *AccountsContainersMacrosUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersMacrosUpdateCall) IfNoneMatch(entityTag string) *AccountsContainersMacrosUpdateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.macros.update" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersMacrosUpdateCall) Do() (*Macro, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersMacrosUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.macro)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fingerprint"]; ok {
		params.Set("fingerprint", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/macros/{macroId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"macroId":     c.macroId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.macros.update" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersMacrosUpdateCall) DoHeader() (ret *Macro, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Updates a GTM Macro.",
	//   "httpMethod": "PUT",
	//   "id": "tagmanager.accounts.containers.macros.update",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "macroId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "fingerprint": {
	//       "description": "When provided, this fingerprint must match the fingerprint of the macro in storage.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "macroId": {
	//       "description": "The GTM Macro ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/macros/{macroId}",
	//   "request": {
	//     "$ref": "Macro"
	//   },
	//   "response": {
	//     "$ref": "Macro"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.move_folders.update":

type AccountsContainersMoveFoldersUpdateCall struct {
	s           *Service
	accountId   string
	containerId string
	folderId    string
	opt_        map[string]interface{}
}

// Update: Moves entities to a GTM Folder.
func (r *AccountsContainersMoveFoldersService) Update(accountId string, containerId string, folderId string) *AccountsContainersMoveFoldersUpdateCall {
	c := &AccountsContainersMoveFoldersUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.folderId = folderId
	return c
}

// TagId sets the optional parameter "tagId": The tags to be moved to
// the folder.
func (c *AccountsContainersMoveFoldersUpdateCall) TagId(tagId string) *AccountsContainersMoveFoldersUpdateCall {
	c.opt_["tagId"] = tagId
	return c
}

// TriggerId sets the optional parameter "triggerId": The triggers to be
// moved to the folder.
func (c *AccountsContainersMoveFoldersUpdateCall) TriggerId(triggerId string) *AccountsContainersMoveFoldersUpdateCall {
	c.opt_["triggerId"] = triggerId
	return c
}

// VariableId sets the optional parameter "variableId": The variables to
// be moved to the folder.
func (c *AccountsContainersMoveFoldersUpdateCall) VariableId(variableId string) *AccountsContainersMoveFoldersUpdateCall {
	c.opt_["variableId"] = variableId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersMoveFoldersUpdateCall) Fields(s ...googleapi.Field) *AccountsContainersMoveFoldersUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersMoveFoldersUpdateCall) IfNoneMatch(entityTag string) *AccountsContainersMoveFoldersUpdateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.move_folders.update" call.
func (c *AccountsContainersMoveFoldersUpdateCall) Do() error {
	_, err := c.DoHeader()
	return err
}

func (c *AccountsContainersMoveFoldersUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["tagId"]; ok {
		params.Set("tagId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["triggerId"]; ok {
		params.Set("triggerId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["variableId"]; ok {
		params.Set("variableId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/move_folders/{folderId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"folderId":    c.folderId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.move_folders.update" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersMoveFoldersUpdateCall) DoHeader() (resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return res.Header, err
	}
	return res.Header, nil
	// {
	//   "description": "Moves entities to a GTM Folder.",
	//   "httpMethod": "PUT",
	//   "id": "tagmanager.accounts.containers.move_folders.update",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "folderId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "folderId": {
	//       "description": "The GTM Folder ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "tagId": {
	//       "description": "The tags to be moved to the folder.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "triggerId": {
	//       "description": "The triggers to be moved to the folder.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "variableId": {
	//       "description": "The variables to be moved to the folder.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/move_folders/{folderId}"
	// }

}

// method id "tagmanager.accounts.containers.rules.create":

type AccountsContainersRulesCreateCall struct {
	s           *Service
	accountId   string
	containerId string
	rule        *Rule
	opt_        map[string]interface{}
}

// Create: Creates a GTM Rule.
func (r *AccountsContainersRulesService) Create(accountId string, containerId string, rule *Rule) *AccountsContainersRulesCreateCall {
	c := &AccountsContainersRulesCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.rule = rule
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersRulesCreateCall) Fields(s ...googleapi.Field) *AccountsContainersRulesCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersRulesCreateCall) IfNoneMatch(entityTag string) *AccountsContainersRulesCreateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.rules.create" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersRulesCreateCall) Do() (*Rule, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersRulesCreateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.rule)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/rules")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.rules.create" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersRulesCreateCall) DoHeader() (ret *Rule, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Creates a GTM Rule.",
	//   "httpMethod": "POST",
	//   "id": "tagmanager.accounts.containers.rules.create",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/rules",
	//   "request": {
	//     "$ref": "Rule"
	//   },
	//   "response": {
	//     "$ref": "Rule"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.rules.delete":

type AccountsContainersRulesDeleteCall struct {
	s           *Service
	accountId   string
	containerId string
	ruleId      string
	opt_        map[string]interface{}
}

// Delete: Deletes a GTM Rule.
func (r *AccountsContainersRulesService) Delete(accountId string, containerId string, ruleId string) *AccountsContainersRulesDeleteCall {
	c := &AccountsContainersRulesDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.ruleId = ruleId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersRulesDeleteCall) Fields(s ...googleapi.Field) *AccountsContainersRulesDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersRulesDeleteCall) IfNoneMatch(entityTag string) *AccountsContainersRulesDeleteCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.rules.delete" call.
func (c *AccountsContainersRulesDeleteCall) Do() error {
	_, err := c.DoHeader()
	return err
}

func (c *AccountsContainersRulesDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/rules/{ruleId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"ruleId":      c.ruleId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.rules.delete" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersRulesDeleteCall) DoHeader() (resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return res.Header, err
	}
	return res.Header, nil
	// {
	//   "description": "Deletes a GTM Rule.",
	//   "httpMethod": "DELETE",
	//   "id": "tagmanager.accounts.containers.rules.delete",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "ruleId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "ruleId": {
	//       "description": "The GTM Rule ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/rules/{ruleId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.rules.get":

type AccountsContainersRulesGetCall struct {
	s           *Service
	accountId   string
	containerId string
	ruleId      string
	opt_        map[string]interface{}
}

// Get: Gets a GTM Rule.
func (r *AccountsContainersRulesService) Get(accountId string, containerId string, ruleId string) *AccountsContainersRulesGetCall {
	c := &AccountsContainersRulesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.ruleId = ruleId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersRulesGetCall) Fields(s ...googleapi.Field) *AccountsContainersRulesGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersRulesGetCall) IfNoneMatch(entityTag string) *AccountsContainersRulesGetCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.rules.get" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersRulesGetCall) Do() (*Rule, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersRulesGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/rules/{ruleId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"ruleId":      c.ruleId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.rules.get" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersRulesGetCall) DoHeader() (ret *Rule, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Gets a GTM Rule.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.rules.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "ruleId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "ruleId": {
	//       "description": "The GTM Rule ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/rules/{ruleId}",
	//   "response": {
	//     "$ref": "Rule"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.rules.list":

type AccountsContainersRulesListCall struct {
	s           *Service
	accountId   string
	containerId string
	opt_        map[string]interface{}
}

// List: Lists all GTM Rules of a Container.
func (r *AccountsContainersRulesService) List(accountId string, containerId string) *AccountsContainersRulesListCall {
	c := &AccountsContainersRulesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersRulesListCall) Fields(s ...googleapi.Field) *AccountsContainersRulesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersRulesListCall) IfNoneMatch(entityTag string) *AccountsContainersRulesListCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.rules.list" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersRulesListCall) Do() (*ListRulesResponse, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersRulesListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/rules")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.rules.list" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersRulesListCall) DoHeader() (ret *ListRulesResponse, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Lists all GTM Rules of a Container.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.rules.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/rules",
	//   "response": {
	//     "$ref": "ListRulesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.rules.update":

type AccountsContainersRulesUpdateCall struct {
	s           *Service
	accountId   string
	containerId string
	ruleId      string
	rule        *Rule
	opt_        map[string]interface{}
}

// Update: Updates a GTM Rule.
func (r *AccountsContainersRulesService) Update(accountId string, containerId string, ruleId string, rule *Rule) *AccountsContainersRulesUpdateCall {
	c := &AccountsContainersRulesUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.ruleId = ruleId
	c.rule = rule
	return c
}

// Fingerprint sets the optional parameter "fingerprint": When provided,
// this fingerprint must match the fingerprint of the rule in storage.
func (c *AccountsContainersRulesUpdateCall) Fingerprint(fingerprint string) *AccountsContainersRulesUpdateCall {
	c.opt_["fingerprint"] = fingerprint
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersRulesUpdateCall) Fields(s ...googleapi.Field) *AccountsContainersRulesUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersRulesUpdateCall) IfNoneMatch(entityTag string) *AccountsContainersRulesUpdateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.rules.update" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersRulesUpdateCall) Do() (*Rule, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersRulesUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.rule)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fingerprint"]; ok {
		params.Set("fingerprint", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/rules/{ruleId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"ruleId":      c.ruleId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.rules.update" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersRulesUpdateCall) DoHeader() (ret *Rule, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Updates a GTM Rule.",
	//   "httpMethod": "PUT",
	//   "id": "tagmanager.accounts.containers.rules.update",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "ruleId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "fingerprint": {
	//       "description": "When provided, this fingerprint must match the fingerprint of the rule in storage.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ruleId": {
	//       "description": "The GTM Rule ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/rules/{ruleId}",
	//   "request": {
	//     "$ref": "Rule"
	//   },
	//   "response": {
	//     "$ref": "Rule"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.tags.create":

type AccountsContainersTagsCreateCall struct {
	s           *Service
	accountId   string
	containerId string
	tag         *Tag
	opt_        map[string]interface{}
}

// Create: Creates a GTM Tag.
func (r *AccountsContainersTagsService) Create(accountId string, containerId string, tag *Tag) *AccountsContainersTagsCreateCall {
	c := &AccountsContainersTagsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.tag = tag
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersTagsCreateCall) Fields(s ...googleapi.Field) *AccountsContainersTagsCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersTagsCreateCall) IfNoneMatch(entityTag string) *AccountsContainersTagsCreateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.tags.create" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersTagsCreateCall) Do() (*Tag, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersTagsCreateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.tag)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/tags")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.tags.create" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersTagsCreateCall) DoHeader() (ret *Tag, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Creates a GTM Tag.",
	//   "httpMethod": "POST",
	//   "id": "tagmanager.accounts.containers.tags.create",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/tags",
	//   "request": {
	//     "$ref": "Tag"
	//   },
	//   "response": {
	//     "$ref": "Tag"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.tags.delete":

type AccountsContainersTagsDeleteCall struct {
	s           *Service
	accountId   string
	containerId string
	tagId       string
	opt_        map[string]interface{}
}

// Delete: Deletes a GTM Tag.
func (r *AccountsContainersTagsService) Delete(accountId string, containerId string, tagId string) *AccountsContainersTagsDeleteCall {
	c := &AccountsContainersTagsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.tagId = tagId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersTagsDeleteCall) Fields(s ...googleapi.Field) *AccountsContainersTagsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersTagsDeleteCall) IfNoneMatch(entityTag string) *AccountsContainersTagsDeleteCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.tags.delete" call.
func (c *AccountsContainersTagsDeleteCall) Do() error {
	_, err := c.DoHeader()
	return err
}

func (c *AccountsContainersTagsDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/tags/{tagId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"tagId":       c.tagId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.tags.delete" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersTagsDeleteCall) DoHeader() (resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return res.Header, err
	}
	return res.Header, nil
	// {
	//   "description": "Deletes a GTM Tag.",
	//   "httpMethod": "DELETE",
	//   "id": "tagmanager.accounts.containers.tags.delete",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "tagId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "tagId": {
	//       "description": "The GTM Tag ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/tags/{tagId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.tags.get":

type AccountsContainersTagsGetCall struct {
	s           *Service
	accountId   string
	containerId string
	tagId       string
	opt_        map[string]interface{}
}

// Get: Gets a GTM Tag.
func (r *AccountsContainersTagsService) Get(accountId string, containerId string, tagId string) *AccountsContainersTagsGetCall {
	c := &AccountsContainersTagsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.tagId = tagId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersTagsGetCall) Fields(s ...googleapi.Field) *AccountsContainersTagsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersTagsGetCall) IfNoneMatch(entityTag string) *AccountsContainersTagsGetCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.tags.get" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersTagsGetCall) Do() (*Tag, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersTagsGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/tags/{tagId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"tagId":       c.tagId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.tags.get" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersTagsGetCall) DoHeader() (ret *Tag, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Gets a GTM Tag.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.tags.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "tagId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "tagId": {
	//       "description": "The GTM Tag ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/tags/{tagId}",
	//   "response": {
	//     "$ref": "Tag"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.tags.list":

type AccountsContainersTagsListCall struct {
	s           *Service
	accountId   string
	containerId string
	opt_        map[string]interface{}
}

// List: Lists all GTM Tags of a Container.
func (r *AccountsContainersTagsService) List(accountId string, containerId string) *AccountsContainersTagsListCall {
	c := &AccountsContainersTagsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersTagsListCall) Fields(s ...googleapi.Field) *AccountsContainersTagsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersTagsListCall) IfNoneMatch(entityTag string) *AccountsContainersTagsListCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.tags.list" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersTagsListCall) Do() (*ListTagsResponse, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersTagsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/tags")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.tags.list" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersTagsListCall) DoHeader() (ret *ListTagsResponse, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Lists all GTM Tags of a Container.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.tags.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/tags",
	//   "response": {
	//     "$ref": "ListTagsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.tags.update":

type AccountsContainersTagsUpdateCall struct {
	s           *Service
	accountId   string
	containerId string
	tagId       string
	tag         *Tag
	opt_        map[string]interface{}
}

// Update: Updates a GTM Tag.
func (r *AccountsContainersTagsService) Update(accountId string, containerId string, tagId string, tag *Tag) *AccountsContainersTagsUpdateCall {
	c := &AccountsContainersTagsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.tagId = tagId
	c.tag = tag
	return c
}

// Fingerprint sets the optional parameter "fingerprint": When provided,
// this fingerprint must match the fingerprint of the tag in storage.
func (c *AccountsContainersTagsUpdateCall) Fingerprint(fingerprint string) *AccountsContainersTagsUpdateCall {
	c.opt_["fingerprint"] = fingerprint
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersTagsUpdateCall) Fields(s ...googleapi.Field) *AccountsContainersTagsUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersTagsUpdateCall) IfNoneMatch(entityTag string) *AccountsContainersTagsUpdateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.tags.update" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersTagsUpdateCall) Do() (*Tag, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersTagsUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.tag)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fingerprint"]; ok {
		params.Set("fingerprint", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/tags/{tagId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"tagId":       c.tagId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.tags.update" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersTagsUpdateCall) DoHeader() (ret *Tag, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Updates a GTM Tag.",
	//   "httpMethod": "PUT",
	//   "id": "tagmanager.accounts.containers.tags.update",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "tagId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "fingerprint": {
	//       "description": "When provided, this fingerprint must match the fingerprint of the tag in storage.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "tagId": {
	//       "description": "The GTM Tag ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/tags/{tagId}",
	//   "request": {
	//     "$ref": "Tag"
	//   },
	//   "response": {
	//     "$ref": "Tag"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.triggers.create":

type AccountsContainersTriggersCreateCall struct {
	s           *Service
	accountId   string
	containerId string
	trigger     *Trigger
	opt_        map[string]interface{}
}

// Create: Creates a GTM Trigger.
func (r *AccountsContainersTriggersService) Create(accountId string, containerId string, trigger *Trigger) *AccountsContainersTriggersCreateCall {
	c := &AccountsContainersTriggersCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.trigger = trigger
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersTriggersCreateCall) Fields(s ...googleapi.Field) *AccountsContainersTriggersCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersTriggersCreateCall) IfNoneMatch(entityTag string) *AccountsContainersTriggersCreateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.triggers.create" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersTriggersCreateCall) Do() (*Trigger, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersTriggersCreateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.trigger)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/triggers")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.triggers.create" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersTriggersCreateCall) DoHeader() (ret *Trigger, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Creates a GTM Trigger.",
	//   "httpMethod": "POST",
	//   "id": "tagmanager.accounts.containers.triggers.create",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/triggers",
	//   "request": {
	//     "$ref": "Trigger"
	//   },
	//   "response": {
	//     "$ref": "Trigger"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.triggers.delete":

type AccountsContainersTriggersDeleteCall struct {
	s           *Service
	accountId   string
	containerId string
	triggerId   string
	opt_        map[string]interface{}
}

// Delete: Deletes a GTM Trigger.
func (r *AccountsContainersTriggersService) Delete(accountId string, containerId string, triggerId string) *AccountsContainersTriggersDeleteCall {
	c := &AccountsContainersTriggersDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.triggerId = triggerId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersTriggersDeleteCall) Fields(s ...googleapi.Field) *AccountsContainersTriggersDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersTriggersDeleteCall) IfNoneMatch(entityTag string) *AccountsContainersTriggersDeleteCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.triggers.delete" call.
func (c *AccountsContainersTriggersDeleteCall) Do() error {
	_, err := c.DoHeader()
	return err
}

func (c *AccountsContainersTriggersDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/triggers/{triggerId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"triggerId":   c.triggerId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.triggers.delete" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersTriggersDeleteCall) DoHeader() (resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return res.Header, err
	}
	return res.Header, nil
	// {
	//   "description": "Deletes a GTM Trigger.",
	//   "httpMethod": "DELETE",
	//   "id": "tagmanager.accounts.containers.triggers.delete",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "triggerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "triggerId": {
	//       "description": "The GTM Trigger ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/triggers/{triggerId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.triggers.get":

type AccountsContainersTriggersGetCall struct {
	s           *Service
	accountId   string
	containerId string
	triggerId   string
	opt_        map[string]interface{}
}

// Get: Gets a GTM Trigger.
func (r *AccountsContainersTriggersService) Get(accountId string, containerId string, triggerId string) *AccountsContainersTriggersGetCall {
	c := &AccountsContainersTriggersGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.triggerId = triggerId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersTriggersGetCall) Fields(s ...googleapi.Field) *AccountsContainersTriggersGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersTriggersGetCall) IfNoneMatch(entityTag string) *AccountsContainersTriggersGetCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.triggers.get" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersTriggersGetCall) Do() (*Trigger, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersTriggersGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/triggers/{triggerId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"triggerId":   c.triggerId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.triggers.get" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersTriggersGetCall) DoHeader() (ret *Trigger, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Gets a GTM Trigger.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.triggers.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "triggerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "triggerId": {
	//       "description": "The GTM Trigger ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/triggers/{triggerId}",
	//   "response": {
	//     "$ref": "Trigger"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.triggers.list":

type AccountsContainersTriggersListCall struct {
	s           *Service
	accountId   string
	containerId string
	opt_        map[string]interface{}
}

// List: Lists all GTM Triggers of a Container.
func (r *AccountsContainersTriggersService) List(accountId string, containerId string) *AccountsContainersTriggersListCall {
	c := &AccountsContainersTriggersListCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersTriggersListCall) Fields(s ...googleapi.Field) *AccountsContainersTriggersListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersTriggersListCall) IfNoneMatch(entityTag string) *AccountsContainersTriggersListCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.triggers.list" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersTriggersListCall) Do() (*ListTriggersResponse, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersTriggersListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/triggers")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.triggers.list" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersTriggersListCall) DoHeader() (ret *ListTriggersResponse, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Lists all GTM Triggers of a Container.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.triggers.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/triggers",
	//   "response": {
	//     "$ref": "ListTriggersResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.triggers.update":

type AccountsContainersTriggersUpdateCall struct {
	s           *Service
	accountId   string
	containerId string
	triggerId   string
	trigger     *Trigger
	opt_        map[string]interface{}
}

// Update: Updates a GTM Trigger.
func (r *AccountsContainersTriggersService) Update(accountId string, containerId string, triggerId string, trigger *Trigger) *AccountsContainersTriggersUpdateCall {
	c := &AccountsContainersTriggersUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.triggerId = triggerId
	c.trigger = trigger
	return c
}

// Fingerprint sets the optional parameter "fingerprint": When provided,
// this fingerprint must match the fingerprint of the trigger in
// storage.
func (c *AccountsContainersTriggersUpdateCall) Fingerprint(fingerprint string) *AccountsContainersTriggersUpdateCall {
	c.opt_["fingerprint"] = fingerprint
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersTriggersUpdateCall) Fields(s ...googleapi.Field) *AccountsContainersTriggersUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersTriggersUpdateCall) IfNoneMatch(entityTag string) *AccountsContainersTriggersUpdateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.triggers.update" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersTriggersUpdateCall) Do() (*Trigger, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersTriggersUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.trigger)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fingerprint"]; ok {
		params.Set("fingerprint", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/triggers/{triggerId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"triggerId":   c.triggerId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.triggers.update" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersTriggersUpdateCall) DoHeader() (ret *Trigger, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Updates a GTM Trigger.",
	//   "httpMethod": "PUT",
	//   "id": "tagmanager.accounts.containers.triggers.update",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "triggerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "fingerprint": {
	//       "description": "When provided, this fingerprint must match the fingerprint of the trigger in storage.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "triggerId": {
	//       "description": "The GTM Trigger ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/triggers/{triggerId}",
	//   "request": {
	//     "$ref": "Trigger"
	//   },
	//   "response": {
	//     "$ref": "Trigger"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.variables.create":

type AccountsContainersVariablesCreateCall struct {
	s           *Service
	accountId   string
	containerId string
	variable    *Variable
	opt_        map[string]interface{}
}

// Create: Creates a GTM Variable.
func (r *AccountsContainersVariablesService) Create(accountId string, containerId string, variable *Variable) *AccountsContainersVariablesCreateCall {
	c := &AccountsContainersVariablesCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.variable = variable
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersVariablesCreateCall) Fields(s ...googleapi.Field) *AccountsContainersVariablesCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersVariablesCreateCall) IfNoneMatch(entityTag string) *AccountsContainersVariablesCreateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.variables.create" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersVariablesCreateCall) Do() (*Variable, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersVariablesCreateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.variable)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/variables")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.variables.create" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersVariablesCreateCall) DoHeader() (ret *Variable, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Creates a GTM Variable.",
	//   "httpMethod": "POST",
	//   "id": "tagmanager.accounts.containers.variables.create",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/variables",
	//   "request": {
	//     "$ref": "Variable"
	//   },
	//   "response": {
	//     "$ref": "Variable"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.variables.delete":

type AccountsContainersVariablesDeleteCall struct {
	s           *Service
	accountId   string
	containerId string
	variableId  string
	opt_        map[string]interface{}
}

// Delete: Deletes a GTM Variable.
func (r *AccountsContainersVariablesService) Delete(accountId string, containerId string, variableId string) *AccountsContainersVariablesDeleteCall {
	c := &AccountsContainersVariablesDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.variableId = variableId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersVariablesDeleteCall) Fields(s ...googleapi.Field) *AccountsContainersVariablesDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersVariablesDeleteCall) IfNoneMatch(entityTag string) *AccountsContainersVariablesDeleteCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.variables.delete" call.
func (c *AccountsContainersVariablesDeleteCall) Do() error {
	_, err := c.DoHeader()
	return err
}

func (c *AccountsContainersVariablesDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/variables/{variableId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"variableId":  c.variableId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.variables.delete" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersVariablesDeleteCall) DoHeader() (resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return res.Header, err
	}
	return res.Header, nil
	// {
	//   "description": "Deletes a GTM Variable.",
	//   "httpMethod": "DELETE",
	//   "id": "tagmanager.accounts.containers.variables.delete",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "variableId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "variableId": {
	//       "description": "The GTM Variable ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/variables/{variableId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.variables.get":

type AccountsContainersVariablesGetCall struct {
	s           *Service
	accountId   string
	containerId string
	variableId  string
	opt_        map[string]interface{}
}

// Get: Gets a GTM Variable.
func (r *AccountsContainersVariablesService) Get(accountId string, containerId string, variableId string) *AccountsContainersVariablesGetCall {
	c := &AccountsContainersVariablesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.variableId = variableId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersVariablesGetCall) Fields(s ...googleapi.Field) *AccountsContainersVariablesGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersVariablesGetCall) IfNoneMatch(entityTag string) *AccountsContainersVariablesGetCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.variables.get" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersVariablesGetCall) Do() (*Variable, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersVariablesGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/variables/{variableId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"variableId":  c.variableId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.variables.get" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersVariablesGetCall) DoHeader() (ret *Variable, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Gets a GTM Variable.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.variables.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "variableId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "variableId": {
	//       "description": "The GTM Variable ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/variables/{variableId}",
	//   "response": {
	//     "$ref": "Variable"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.variables.list":

type AccountsContainersVariablesListCall struct {
	s           *Service
	accountId   string
	containerId string
	opt_        map[string]interface{}
}

// List: Lists all GTM Variables of a Container.
func (r *AccountsContainersVariablesService) List(accountId string, containerId string) *AccountsContainersVariablesListCall {
	c := &AccountsContainersVariablesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersVariablesListCall) Fields(s ...googleapi.Field) *AccountsContainersVariablesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersVariablesListCall) IfNoneMatch(entityTag string) *AccountsContainersVariablesListCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.variables.list" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersVariablesListCall) Do() (*ListVariablesResponse, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersVariablesListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/variables")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.variables.list" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersVariablesListCall) DoHeader() (ret *ListVariablesResponse, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Lists all GTM Variables of a Container.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.variables.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/variables",
	//   "response": {
	//     "$ref": "ListVariablesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.variables.update":

type AccountsContainersVariablesUpdateCall struct {
	s           *Service
	accountId   string
	containerId string
	variableId  string
	variable    *Variable
	opt_        map[string]interface{}
}

// Update: Updates a GTM Variable.
func (r *AccountsContainersVariablesService) Update(accountId string, containerId string, variableId string, variable *Variable) *AccountsContainersVariablesUpdateCall {
	c := &AccountsContainersVariablesUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.variableId = variableId
	c.variable = variable
	return c
}

// Fingerprint sets the optional parameter "fingerprint": When provided,
// this fingerprint must match the fingerprint of the variable in
// storage.
func (c *AccountsContainersVariablesUpdateCall) Fingerprint(fingerprint string) *AccountsContainersVariablesUpdateCall {
	c.opt_["fingerprint"] = fingerprint
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersVariablesUpdateCall) Fields(s ...googleapi.Field) *AccountsContainersVariablesUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersVariablesUpdateCall) IfNoneMatch(entityTag string) *AccountsContainersVariablesUpdateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.variables.update" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersVariablesUpdateCall) Do() (*Variable, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersVariablesUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.variable)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fingerprint"]; ok {
		params.Set("fingerprint", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/variables/{variableId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"variableId":  c.variableId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.variables.update" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersVariablesUpdateCall) DoHeader() (ret *Variable, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Updates a GTM Variable.",
	//   "httpMethod": "PUT",
	//   "id": "tagmanager.accounts.containers.variables.update",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "variableId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "fingerprint": {
	//       "description": "When provided, this fingerprint must match the fingerprint of the variable in storage.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "variableId": {
	//       "description": "The GTM Variable ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/variables/{variableId}",
	//   "request": {
	//     "$ref": "Variable"
	//   },
	//   "response": {
	//     "$ref": "Variable"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.versions.create":

type AccountsContainersVersionsCreateCall struct {
	s                                           *Service
	accountId                                   string
	containerId                                 string
	createcontainerversionrequestversionoptions *CreateContainerVersionRequestVersionOptions
	opt_                                        map[string]interface{}
}

// Create: Creates a Container Version.
func (r *AccountsContainersVersionsService) Create(accountId string, containerId string, createcontainerversionrequestversionoptions *CreateContainerVersionRequestVersionOptions) *AccountsContainersVersionsCreateCall {
	c := &AccountsContainersVersionsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.createcontainerversionrequestversionoptions = createcontainerversionrequestversionoptions
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersVersionsCreateCall) Fields(s ...googleapi.Field) *AccountsContainersVersionsCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersVersionsCreateCall) IfNoneMatch(entityTag string) *AccountsContainersVersionsCreateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.versions.create" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersVersionsCreateCall) Do() (*CreateContainerVersionResponse, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersVersionsCreateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.createcontainerversionrequestversionoptions)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/versions")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.versions.create" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersVersionsCreateCall) DoHeader() (ret *CreateContainerVersionResponse, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Creates a Container Version.",
	//   "httpMethod": "POST",
	//   "id": "tagmanager.accounts.containers.versions.create",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/versions",
	//   "request": {
	//     "$ref": "CreateContainerVersionRequestVersionOptions"
	//   },
	//   "response": {
	//     "$ref": "CreateContainerVersionResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containerversions"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.versions.delete":

type AccountsContainersVersionsDeleteCall struct {
	s                  *Service
	accountId          string
	containerId        string
	containerVersionId string
	opt_               map[string]interface{}
}

// Delete: Deletes a Container Version.
func (r *AccountsContainersVersionsService) Delete(accountId string, containerId string, containerVersionId string) *AccountsContainersVersionsDeleteCall {
	c := &AccountsContainersVersionsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.containerVersionId = containerVersionId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersVersionsDeleteCall) Fields(s ...googleapi.Field) *AccountsContainersVersionsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersVersionsDeleteCall) IfNoneMatch(entityTag string) *AccountsContainersVersionsDeleteCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.versions.delete" call.
func (c *AccountsContainersVersionsDeleteCall) Do() error {
	_, err := c.DoHeader()
	return err
}

func (c *AccountsContainersVersionsDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/versions/{containerVersionId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":          c.accountId,
		"containerId":        c.containerId,
		"containerVersionId": c.containerVersionId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.versions.delete" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersVersionsDeleteCall) DoHeader() (resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return res.Header, err
	}
	return res.Header, nil
	// {
	//   "description": "Deletes a Container Version.",
	//   "httpMethod": "DELETE",
	//   "id": "tagmanager.accounts.containers.versions.delete",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "containerVersionId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerVersionId": {
	//       "description": "The GTM Container Version ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/versions/{containerVersionId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containerversions"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.versions.get":

type AccountsContainersVersionsGetCall struct {
	s                  *Service
	accountId          string
	containerId        string
	containerVersionId string
	opt_               map[string]interface{}
}

// Get: Gets a Container Version.
func (r *AccountsContainersVersionsService) Get(accountId string, containerId string, containerVersionId string) *AccountsContainersVersionsGetCall {
	c := &AccountsContainersVersionsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.containerVersionId = containerVersionId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersVersionsGetCall) Fields(s ...googleapi.Field) *AccountsContainersVersionsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersVersionsGetCall) IfNoneMatch(entityTag string) *AccountsContainersVersionsGetCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.versions.get" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersVersionsGetCall) Do() (*ContainerVersion, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersVersionsGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/versions/{containerVersionId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":          c.accountId,
		"containerId":        c.containerId,
		"containerVersionId": c.containerVersionId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.versions.get" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersVersionsGetCall) DoHeader() (ret *ContainerVersion, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Gets a Container Version.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.versions.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "containerVersionId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerVersionId": {
	//       "description": "The GTM Container Version ID. Specify published to retrieve the currently published version.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/versions/{containerVersionId}",
	//   "response": {
	//     "$ref": "ContainerVersion"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.edit.containerversions",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.versions.list":

type AccountsContainersVersionsListCall struct {
	s           *Service
	accountId   string
	containerId string
	opt_        map[string]interface{}
}

// List: Lists all Container Versions of a GTM Container.
func (r *AccountsContainersVersionsService) List(accountId string, containerId string) *AccountsContainersVersionsListCall {
	c := &AccountsContainersVersionsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	return c
}

// Headers sets the optional parameter "headers": Retrieve headers only
// when true.
func (c *AccountsContainersVersionsListCall) Headers(headers bool) *AccountsContainersVersionsListCall {
	c.opt_["headers"] = headers
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersVersionsListCall) Fields(s ...googleapi.Field) *AccountsContainersVersionsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersVersionsListCall) IfNoneMatch(entityTag string) *AccountsContainersVersionsListCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.versions.list" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersVersionsListCall) Do() (*ListContainerVersionsResponse, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersVersionsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["headers"]; ok {
		params.Set("headers", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/versions")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.versions.list" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersVersionsListCall) DoHeader() (ret *ListContainerVersionsResponse, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Lists all Container Versions of a GTM Container.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.versions.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "headers": {
	//       "default": "false",
	//       "description": "Retrieve headers only when true.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/versions",
	//   "response": {
	//     "$ref": "ListContainerVersionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.edit.containerversions",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.versions.publish":

type AccountsContainersVersionsPublishCall struct {
	s                  *Service
	accountId          string
	containerId        string
	containerVersionId string
	opt_               map[string]interface{}
}

// Publish: Publishes a Container Version.
func (r *AccountsContainersVersionsService) Publish(accountId string, containerId string, containerVersionId string) *AccountsContainersVersionsPublishCall {
	c := &AccountsContainersVersionsPublishCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.containerVersionId = containerVersionId
	return c
}

// Fingerprint sets the optional parameter "fingerprint": When provided,
// this fingerprint must match the fingerprint of the container version
// in storage.
func (c *AccountsContainersVersionsPublishCall) Fingerprint(fingerprint string) *AccountsContainersVersionsPublishCall {
	c.opt_["fingerprint"] = fingerprint
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersVersionsPublishCall) Fields(s ...googleapi.Field) *AccountsContainersVersionsPublishCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersVersionsPublishCall) IfNoneMatch(entityTag string) *AccountsContainersVersionsPublishCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.versions.publish" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersVersionsPublishCall) Do() (*PublishContainerVersionResponse, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersVersionsPublishCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fingerprint"]; ok {
		params.Set("fingerprint", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/versions/{containerVersionId}/publish")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":          c.accountId,
		"containerId":        c.containerId,
		"containerVersionId": c.containerVersionId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.versions.publish" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersVersionsPublishCall) DoHeader() (ret *PublishContainerVersionResponse, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Publishes a Container Version.",
	//   "httpMethod": "POST",
	//   "id": "tagmanager.accounts.containers.versions.publish",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "containerVersionId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerVersionId": {
	//       "description": "The GTM Container Version ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "fingerprint": {
	//       "description": "When provided, this fingerprint must match the fingerprint of the container version in storage.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/versions/{containerVersionId}/publish",
	//   "response": {
	//     "$ref": "PublishContainerVersionResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.publish"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.versions.restore":

type AccountsContainersVersionsRestoreCall struct {
	s                  *Service
	accountId          string
	containerId        string
	containerVersionId string
	opt_               map[string]interface{}
}

// Restore: Restores a Container Version. This will overwrite the
// container's current configuration (including its macros, rules and
// tags). The operation will not have any effect on the version that is
// being served (i.e. the published version).
func (r *AccountsContainersVersionsService) Restore(accountId string, containerId string, containerVersionId string) *AccountsContainersVersionsRestoreCall {
	c := &AccountsContainersVersionsRestoreCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.containerVersionId = containerVersionId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersVersionsRestoreCall) Fields(s ...googleapi.Field) *AccountsContainersVersionsRestoreCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersVersionsRestoreCall) IfNoneMatch(entityTag string) *AccountsContainersVersionsRestoreCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.versions.restore" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersVersionsRestoreCall) Do() (*ContainerVersion, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersVersionsRestoreCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/versions/{containerVersionId}/restore")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":          c.accountId,
		"containerId":        c.containerId,
		"containerVersionId": c.containerVersionId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.versions.restore" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersVersionsRestoreCall) DoHeader() (ret *ContainerVersion, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Restores a Container Version. This will overwrite the container's current configuration (including its macros, rules and tags). The operation will not have any effect on the version that is being served (i.e. the published version).",
	//   "httpMethod": "POST",
	//   "id": "tagmanager.accounts.containers.versions.restore",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "containerVersionId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerVersionId": {
	//       "description": "The GTM Container Version ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/versions/{containerVersionId}/restore",
	//   "response": {
	//     "$ref": "ContainerVersion"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.versions.undelete":

type AccountsContainersVersionsUndeleteCall struct {
	s                  *Service
	accountId          string
	containerId        string
	containerVersionId string
	opt_               map[string]interface{}
}

// Undelete: Undeletes a Container Version.
func (r *AccountsContainersVersionsService) Undelete(accountId string, containerId string, containerVersionId string) *AccountsContainersVersionsUndeleteCall {
	c := &AccountsContainersVersionsUndeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.containerVersionId = containerVersionId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersVersionsUndeleteCall) Fields(s ...googleapi.Field) *AccountsContainersVersionsUndeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersVersionsUndeleteCall) IfNoneMatch(entityTag string) *AccountsContainersVersionsUndeleteCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.versions.undelete" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersVersionsUndeleteCall) Do() (*ContainerVersion, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersVersionsUndeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/versions/{containerVersionId}/undelete")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":          c.accountId,
		"containerId":        c.containerId,
		"containerVersionId": c.containerVersionId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.versions.undelete" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersVersionsUndeleteCall) DoHeader() (ret *ContainerVersion, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Undeletes a Container Version.",
	//   "httpMethod": "POST",
	//   "id": "tagmanager.accounts.containers.versions.undelete",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "containerVersionId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerVersionId": {
	//       "description": "The GTM Container Version ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/versions/{containerVersionId}/undelete",
	//   "response": {
	//     "$ref": "ContainerVersion"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containerversions"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.versions.update":

type AccountsContainersVersionsUpdateCall struct {
	s                  *Service
	accountId          string
	containerId        string
	containerVersionId string
	containerversion   *ContainerVersion
	opt_               map[string]interface{}
}

// Update: Updates a Container Version.
func (r *AccountsContainersVersionsService) Update(accountId string, containerId string, containerVersionId string, containerversion *ContainerVersion) *AccountsContainersVersionsUpdateCall {
	c := &AccountsContainersVersionsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.containerId = containerId
	c.containerVersionId = containerVersionId
	c.containerversion = containerversion
	return c
}

// Fingerprint sets the optional parameter "fingerprint": When provided,
// this fingerprint must match the fingerprint of the container version
// in storage.
func (c *AccountsContainersVersionsUpdateCall) Fingerprint(fingerprint string) *AccountsContainersVersionsUpdateCall {
	c.opt_["fingerprint"] = fingerprint
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersVersionsUpdateCall) Fields(s ...googleapi.Field) *AccountsContainersVersionsUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsContainersVersionsUpdateCall) IfNoneMatch(entityTag string) *AccountsContainersVersionsUpdateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.containers.versions.update" call.
// Exactly one of the return values is non-nil.
func (c *AccountsContainersVersionsUpdateCall) Do() (*ContainerVersion, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsContainersVersionsUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.containerversion)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fingerprint"]; ok {
		params.Set("fingerprint", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/containers/{containerId}/versions/{containerVersionId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":          c.accountId,
		"containerId":        c.containerId,
		"containerVersionId": c.containerVersionId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.containers.versions.update" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsContainersVersionsUpdateCall) DoHeader() (ret *ContainerVersion, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Updates a Container Version.",
	//   "httpMethod": "PUT",
	//   "id": "tagmanager.accounts.containers.versions.update",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "containerVersionId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerVersionId": {
	//       "description": "The GTM Container Version ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "fingerprint": {
	//       "description": "When provided, this fingerprint must match the fingerprint of the container version in storage.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/versions/{containerVersionId}",
	//   "request": {
	//     "$ref": "ContainerVersion"
	//   },
	//   "response": {
	//     "$ref": "ContainerVersion"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containerversions"
	//   ]
	// }

}

// method id "tagmanager.accounts.permissions.create":

type AccountsPermissionsCreateCall struct {
	s          *Service
	accountId  string
	useraccess *UserAccess
	opt_       map[string]interface{}
}

// Create: Creates a user's Account & Container Permissions.
func (r *AccountsPermissionsService) Create(accountId string, useraccess *UserAccess) *AccountsPermissionsCreateCall {
	c := &AccountsPermissionsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.useraccess = useraccess
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsPermissionsCreateCall) Fields(s ...googleapi.Field) *AccountsPermissionsCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsPermissionsCreateCall) IfNoneMatch(entityTag string) *AccountsPermissionsCreateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.permissions.create" call.
// Exactly one of the return values is non-nil.
func (c *AccountsPermissionsCreateCall) Do() (*UserAccess, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsPermissionsCreateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.useraccess)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/permissions")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId": c.accountId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.permissions.create" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsPermissionsCreateCall) DoHeader() (ret *UserAccess, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Creates a user's Account \u0026 Container Permissions.",
	//   "httpMethod": "POST",
	//   "id": "tagmanager.accounts.permissions.create",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/permissions",
	//   "request": {
	//     "$ref": "UserAccess"
	//   },
	//   "response": {
	//     "$ref": "UserAccess"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.manage.users"
	//   ]
	// }

}

// method id "tagmanager.accounts.permissions.delete":

type AccountsPermissionsDeleteCall struct {
	s            *Service
	accountId    string
	permissionId string
	opt_         map[string]interface{}
}

// Delete: Removes a user from the account, revoking access to it and
// all of its containers.
func (r *AccountsPermissionsService) Delete(accountId string, permissionId string) *AccountsPermissionsDeleteCall {
	c := &AccountsPermissionsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.permissionId = permissionId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsPermissionsDeleteCall) Fields(s ...googleapi.Field) *AccountsPermissionsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsPermissionsDeleteCall) IfNoneMatch(entityTag string) *AccountsPermissionsDeleteCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.permissions.delete" call.
func (c *AccountsPermissionsDeleteCall) Do() error {
	_, err := c.DoHeader()
	return err
}

func (c *AccountsPermissionsDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/permissions/{permissionId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":    c.accountId,
		"permissionId": c.permissionId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.permissions.delete" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsPermissionsDeleteCall) DoHeader() (resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return res.Header, err
	}
	return res.Header, nil
	// {
	//   "description": "Removes a user from the account, revoking access to it and all of its containers.",
	//   "httpMethod": "DELETE",
	//   "id": "tagmanager.accounts.permissions.delete",
	//   "parameterOrder": [
	//     "accountId",
	//     "permissionId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "permissionId": {
	//       "description": "The GTM User ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/permissions/{permissionId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.manage.users"
	//   ]
	// }

}

// method id "tagmanager.accounts.permissions.get":

type AccountsPermissionsGetCall struct {
	s            *Service
	accountId    string
	permissionId string
	opt_         map[string]interface{}
}

// Get: Gets a user's Account & Container Permissions.
func (r *AccountsPermissionsService) Get(accountId string, permissionId string) *AccountsPermissionsGetCall {
	c := &AccountsPermissionsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.permissionId = permissionId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsPermissionsGetCall) Fields(s ...googleapi.Field) *AccountsPermissionsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsPermissionsGetCall) IfNoneMatch(entityTag string) *AccountsPermissionsGetCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.permissions.get" call.
// Exactly one of the return values is non-nil.
func (c *AccountsPermissionsGetCall) Do() (*UserAccess, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsPermissionsGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/permissions/{permissionId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":    c.accountId,
		"permissionId": c.permissionId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.permissions.get" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsPermissionsGetCall) DoHeader() (ret *UserAccess, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Gets a user's Account \u0026 Container Permissions.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.permissions.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "permissionId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "permissionId": {
	//       "description": "The GTM User ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/permissions/{permissionId}",
	//   "response": {
	//     "$ref": "UserAccess"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.manage.users"
	//   ]
	// }

}

// method id "tagmanager.accounts.permissions.list":

type AccountsPermissionsListCall struct {
	s         *Service
	accountId string
	opt_      map[string]interface{}
}

// List: List all users that have access to the account along with
// Account and Container Permissions granted to each of them.
func (r *AccountsPermissionsService) List(accountId string) *AccountsPermissionsListCall {
	c := &AccountsPermissionsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsPermissionsListCall) Fields(s ...googleapi.Field) *AccountsPermissionsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsPermissionsListCall) IfNoneMatch(entityTag string) *AccountsPermissionsListCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.permissions.list" call.
// Exactly one of the return values is non-nil.
func (c *AccountsPermissionsListCall) Do() (*ListAccountUsersResponse, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsPermissionsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/permissions")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId": c.accountId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.permissions.list" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsPermissionsListCall) DoHeader() (ret *ListAccountUsersResponse, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "List all users that have access to the account along with Account and Container Permissions granted to each of them.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.permissions.list",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID. @required tagmanager.accounts.permissions.list",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/permissions",
	//   "response": {
	//     "$ref": "ListAccountUsersResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.manage.users"
	//   ]
	// }

}

// method id "tagmanager.accounts.permissions.update":

type AccountsPermissionsUpdateCall struct {
	s            *Service
	accountId    string
	permissionId string
	useraccess   *UserAccess
	opt_         map[string]interface{}
}

// Update: Updates a user's Account & Container Permissions.
func (r *AccountsPermissionsService) Update(accountId string, permissionId string, useraccess *UserAccess) *AccountsPermissionsUpdateCall {
	c := &AccountsPermissionsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.accountId = accountId
	c.permissionId = permissionId
	c.useraccess = useraccess
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsPermissionsUpdateCall) Fields(s ...googleapi.Field) *AccountsPermissionsUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AccountsPermissionsUpdateCall) IfNoneMatch(entityTag string) *AccountsPermissionsUpdateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "tagmanager.accounts.permissions.update" call.
// Exactly one of the return values is non-nil.
func (c *AccountsPermissionsUpdateCall) Do() (*UserAccess, error) {
	v, _, err := c.DoHeader()
	return v, err
}

func (c *AccountsPermissionsUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.useraccess)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/permissions/{permissionId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":    c.accountId,
		"permissionId": c.permissionId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	return c.s.client.Do(req)
}

// DoHeader executes the "tagmanager.accounts.permissions.update" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AccountsPermissionsUpdateCall) DoHeader() (ret *UserAccess, resHeader http.Header, err error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Updates a user's Account \u0026 Container Permissions.",
	//   "httpMethod": "PUT",
	//   "id": "tagmanager.accounts.permissions.update",
	//   "parameterOrder": [
	//     "accountId",
	//     "permissionId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "permissionId": {
	//       "description": "The GTM User ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/permissions/{permissionId}",
	//   "request": {
	//     "$ref": "UserAccess"
	//   },
	//   "response": {
	//     "$ref": "UserAccess"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.manage.users"
	//   ]
	// }

}
