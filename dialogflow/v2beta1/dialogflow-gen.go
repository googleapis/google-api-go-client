// Package dialogflow provides access to the Dialogflow API.
//
// This package is DEPRECATED. Use package cloud.google.com/go/dialogflow/apiv2 instead.
//
// See https://cloud.google.com/dialogflow-enterprise/
//
// Usage example:
//
//   import "google.golang.org/api/dialogflow/v2beta1"
//   ...
//   dialogflowService, err := dialogflow.New(oauthHttpClient)
package dialogflow // import "google.golang.org/api/dialogflow/v2beta1"

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

const apiId = "dialogflow:v2beta1"
const apiName = "dialogflow"
const apiVersion = "v2beta1"
const basePath = "https://dialogflow.googleapis.com/"

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
	rs.Agent = NewProjectsAgentService(s)
	rs.ConversationProfiles = NewProjectsConversationProfilesService(s)
	rs.Conversations = NewProjectsConversationsService(s)
	rs.Environments = NewProjectsEnvironmentsService(s)
	rs.HumanAgentAssistants = NewProjectsHumanAgentAssistantsService(s)
	rs.KnowledgeBases = NewProjectsKnowledgeBasesService(s)
	rs.Operations = NewProjectsOperationsService(s)
	rs.PhoneNumberOrders = NewProjectsPhoneNumberOrdersService(s)
	rs.PhoneNumbers = NewProjectsPhoneNumbersService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Agent *ProjectsAgentService

	ConversationProfiles *ProjectsConversationProfilesService

	Conversations *ProjectsConversationsService

	Environments *ProjectsEnvironmentsService

	HumanAgentAssistants *ProjectsHumanAgentAssistantsService

	KnowledgeBases *ProjectsKnowledgeBasesService

	Operations *ProjectsOperationsService

	PhoneNumberOrders *ProjectsPhoneNumberOrdersService

	PhoneNumbers *ProjectsPhoneNumbersService
}

func NewProjectsAgentService(s *Service) *ProjectsAgentService {
	rs := &ProjectsAgentService{s: s}
	rs.EntityTypes = NewProjectsAgentEntityTypesService(s)
	rs.Environments = NewProjectsAgentEnvironmentsService(s)
	rs.Intents = NewProjectsAgentIntentsService(s)
	rs.KnowledgeBases = NewProjectsAgentKnowledgeBasesService(s)
	rs.Sessions = NewProjectsAgentSessionsService(s)
	return rs
}

type ProjectsAgentService struct {
	s *Service

	EntityTypes *ProjectsAgentEntityTypesService

	Environments *ProjectsAgentEnvironmentsService

	Intents *ProjectsAgentIntentsService

	KnowledgeBases *ProjectsAgentKnowledgeBasesService

	Sessions *ProjectsAgentSessionsService
}

func NewProjectsAgentEntityTypesService(s *Service) *ProjectsAgentEntityTypesService {
	rs := &ProjectsAgentEntityTypesService{s: s}
	rs.Entities = NewProjectsAgentEntityTypesEntitiesService(s)
	return rs
}

type ProjectsAgentEntityTypesService struct {
	s *Service

	Entities *ProjectsAgentEntityTypesEntitiesService
}

func NewProjectsAgentEntityTypesEntitiesService(s *Service) *ProjectsAgentEntityTypesEntitiesService {
	rs := &ProjectsAgentEntityTypesEntitiesService{s: s}
	return rs
}

type ProjectsAgentEntityTypesEntitiesService struct {
	s *Service
}

func NewProjectsAgentEnvironmentsService(s *Service) *ProjectsAgentEnvironmentsService {
	rs := &ProjectsAgentEnvironmentsService{s: s}
	rs.Users = NewProjectsAgentEnvironmentsUsersService(s)
	return rs
}

type ProjectsAgentEnvironmentsService struct {
	s *Service

	Users *ProjectsAgentEnvironmentsUsersService
}

func NewProjectsAgentEnvironmentsUsersService(s *Service) *ProjectsAgentEnvironmentsUsersService {
	rs := &ProjectsAgentEnvironmentsUsersService{s: s}
	rs.Sessions = NewProjectsAgentEnvironmentsUsersSessionsService(s)
	return rs
}

type ProjectsAgentEnvironmentsUsersService struct {
	s *Service

	Sessions *ProjectsAgentEnvironmentsUsersSessionsService
}

func NewProjectsAgentEnvironmentsUsersSessionsService(s *Service) *ProjectsAgentEnvironmentsUsersSessionsService {
	rs := &ProjectsAgentEnvironmentsUsersSessionsService{s: s}
	rs.Contexts = NewProjectsAgentEnvironmentsUsersSessionsContextsService(s)
	rs.EntityTypes = NewProjectsAgentEnvironmentsUsersSessionsEntityTypesService(s)
	return rs
}

type ProjectsAgentEnvironmentsUsersSessionsService struct {
	s *Service

	Contexts *ProjectsAgentEnvironmentsUsersSessionsContextsService

	EntityTypes *ProjectsAgentEnvironmentsUsersSessionsEntityTypesService
}

func NewProjectsAgentEnvironmentsUsersSessionsContextsService(s *Service) *ProjectsAgentEnvironmentsUsersSessionsContextsService {
	rs := &ProjectsAgentEnvironmentsUsersSessionsContextsService{s: s}
	return rs
}

type ProjectsAgentEnvironmentsUsersSessionsContextsService struct {
	s *Service
}

func NewProjectsAgentEnvironmentsUsersSessionsEntityTypesService(s *Service) *ProjectsAgentEnvironmentsUsersSessionsEntityTypesService {
	rs := &ProjectsAgentEnvironmentsUsersSessionsEntityTypesService{s: s}
	return rs
}

type ProjectsAgentEnvironmentsUsersSessionsEntityTypesService struct {
	s *Service
}

func NewProjectsAgentIntentsService(s *Service) *ProjectsAgentIntentsService {
	rs := &ProjectsAgentIntentsService{s: s}
	return rs
}

type ProjectsAgentIntentsService struct {
	s *Service
}

func NewProjectsAgentKnowledgeBasesService(s *Service) *ProjectsAgentKnowledgeBasesService {
	rs := &ProjectsAgentKnowledgeBasesService{s: s}
	rs.Documents = NewProjectsAgentKnowledgeBasesDocumentsService(s)
	return rs
}

type ProjectsAgentKnowledgeBasesService struct {
	s *Service

	Documents *ProjectsAgentKnowledgeBasesDocumentsService
}

func NewProjectsAgentKnowledgeBasesDocumentsService(s *Service) *ProjectsAgentKnowledgeBasesDocumentsService {
	rs := &ProjectsAgentKnowledgeBasesDocumentsService{s: s}
	return rs
}

type ProjectsAgentKnowledgeBasesDocumentsService struct {
	s *Service
}

func NewProjectsAgentSessionsService(s *Service) *ProjectsAgentSessionsService {
	rs := &ProjectsAgentSessionsService{s: s}
	rs.Contexts = NewProjectsAgentSessionsContextsService(s)
	rs.EntityTypes = NewProjectsAgentSessionsEntityTypesService(s)
	return rs
}

type ProjectsAgentSessionsService struct {
	s *Service

	Contexts *ProjectsAgentSessionsContextsService

	EntityTypes *ProjectsAgentSessionsEntityTypesService
}

func NewProjectsAgentSessionsContextsService(s *Service) *ProjectsAgentSessionsContextsService {
	rs := &ProjectsAgentSessionsContextsService{s: s}
	return rs
}

type ProjectsAgentSessionsContextsService struct {
	s *Service
}

func NewProjectsAgentSessionsEntityTypesService(s *Service) *ProjectsAgentSessionsEntityTypesService {
	rs := &ProjectsAgentSessionsEntityTypesService{s: s}
	return rs
}

type ProjectsAgentSessionsEntityTypesService struct {
	s *Service
}

func NewProjectsConversationProfilesService(s *Service) *ProjectsConversationProfilesService {
	rs := &ProjectsConversationProfilesService{s: s}
	return rs
}

type ProjectsConversationProfilesService struct {
	s *Service
}

func NewProjectsConversationsService(s *Service) *ProjectsConversationsService {
	rs := &ProjectsConversationsService{s: s}
	rs.Messages = NewProjectsConversationsMessagesService(s)
	rs.Participants = NewProjectsConversationsParticipantsService(s)
	return rs
}

type ProjectsConversationsService struct {
	s *Service

	Messages *ProjectsConversationsMessagesService

	Participants *ProjectsConversationsParticipantsService
}

func NewProjectsConversationsMessagesService(s *Service) *ProjectsConversationsMessagesService {
	rs := &ProjectsConversationsMessagesService{s: s}
	return rs
}

type ProjectsConversationsMessagesService struct {
	s *Service
}

func NewProjectsConversationsParticipantsService(s *Service) *ProjectsConversationsParticipantsService {
	rs := &ProjectsConversationsParticipantsService{s: s}
	rs.Suggestions = NewProjectsConversationsParticipantsSuggestionsService(s)
	return rs
}

type ProjectsConversationsParticipantsService struct {
	s *Service

	Suggestions *ProjectsConversationsParticipantsSuggestionsService
}

func NewProjectsConversationsParticipantsSuggestionsService(s *Service) *ProjectsConversationsParticipantsSuggestionsService {
	rs := &ProjectsConversationsParticipantsSuggestionsService{s: s}
	return rs
}

type ProjectsConversationsParticipantsSuggestionsService struct {
	s *Service
}

func NewProjectsEnvironmentsService(s *Service) *ProjectsEnvironmentsService {
	rs := &ProjectsEnvironmentsService{s: s}
	rs.Users = NewProjectsEnvironmentsUsersService(s)
	return rs
}

type ProjectsEnvironmentsService struct {
	s *Service

	Users *ProjectsEnvironmentsUsersService
}

func NewProjectsEnvironmentsUsersService(s *Service) *ProjectsEnvironmentsUsersService {
	rs := &ProjectsEnvironmentsUsersService{s: s}
	rs.Conversations = NewProjectsEnvironmentsUsersConversationsService(s)
	return rs
}

type ProjectsEnvironmentsUsersService struct {
	s *Service

	Conversations *ProjectsEnvironmentsUsersConversationsService
}

func NewProjectsEnvironmentsUsersConversationsService(s *Service) *ProjectsEnvironmentsUsersConversationsService {
	rs := &ProjectsEnvironmentsUsersConversationsService{s: s}
	rs.Contexts = NewProjectsEnvironmentsUsersConversationsContextsService(s)
	return rs
}

type ProjectsEnvironmentsUsersConversationsService struct {
	s *Service

	Contexts *ProjectsEnvironmentsUsersConversationsContextsService
}

func NewProjectsEnvironmentsUsersConversationsContextsService(s *Service) *ProjectsEnvironmentsUsersConversationsContextsService {
	rs := &ProjectsEnvironmentsUsersConversationsContextsService{s: s}
	return rs
}

type ProjectsEnvironmentsUsersConversationsContextsService struct {
	s *Service
}

func NewProjectsHumanAgentAssistantsService(s *Service) *ProjectsHumanAgentAssistantsService {
	rs := &ProjectsHumanAgentAssistantsService{s: s}
	return rs
}

type ProjectsHumanAgentAssistantsService struct {
	s *Service
}

func NewProjectsKnowledgeBasesService(s *Service) *ProjectsKnowledgeBasesService {
	rs := &ProjectsKnowledgeBasesService{s: s}
	rs.Documents = NewProjectsKnowledgeBasesDocumentsService(s)
	return rs
}

type ProjectsKnowledgeBasesService struct {
	s *Service

	Documents *ProjectsKnowledgeBasesDocumentsService
}

func NewProjectsKnowledgeBasesDocumentsService(s *Service) *ProjectsKnowledgeBasesDocumentsService {
	rs := &ProjectsKnowledgeBasesDocumentsService{s: s}
	return rs
}

type ProjectsKnowledgeBasesDocumentsService struct {
	s *Service
}

func NewProjectsOperationsService(s *Service) *ProjectsOperationsService {
	rs := &ProjectsOperationsService{s: s}
	return rs
}

type ProjectsOperationsService struct {
	s *Service
}

func NewProjectsPhoneNumberOrdersService(s *Service) *ProjectsPhoneNumberOrdersService {
	rs := &ProjectsPhoneNumberOrdersService{s: s}
	return rs
}

type ProjectsPhoneNumberOrdersService struct {
	s *Service
}

func NewProjectsPhoneNumbersService(s *Service) *ProjectsPhoneNumbersService {
	rs := &ProjectsPhoneNumbersService{s: s}
	return rs
}

type ProjectsPhoneNumbersService struct {
	s *Service
}

// GoogleCloudDialogflowV2BatchUpdateEntityTypesResponse: The response
// message for EntityTypes.BatchUpdateEntityTypes.
type GoogleCloudDialogflowV2BatchUpdateEntityTypesResponse struct {
	// EntityTypes: The collection of updated or created entity types.
	EntityTypes []*GoogleCloudDialogflowV2EntityType `json:"entityTypes,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EntityTypes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EntityTypes") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2BatchUpdateEntityTypesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2BatchUpdateEntityTypesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2BatchUpdateIntentsResponse: The response
// message for Intents.BatchUpdateIntents.
type GoogleCloudDialogflowV2BatchUpdateIntentsResponse struct {
	// Intents: The collection of updated or created intents.
	Intents []*GoogleCloudDialogflowV2Intent `json:"intents,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Intents") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Intents") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2BatchUpdateIntentsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2BatchUpdateIntentsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2Context: Represents a context.
type GoogleCloudDialogflowV2Context struct {
	// LifespanCount: Optional. The number of conversational query requests
	// after which the
	// context expires. If set to `0` (the default) the context
	// expires
	// immediately. Contexts expire automatically after 20 minutes even if
	// there
	// are no matching queries.
	LifespanCount int64 `json:"lifespanCount,omitempty"`

	// Name: Required. The unique identifier of the context.
	// Format:
	// `projects/<Project ID>/agent/sessions/<Session ID>/contexts/<Context
	// ID>`.
	Name string `json:"name,omitempty"`

	// Parameters: Optional. The collection of parameters associated with
	// this context.
	// Refer to [this
	// doc](https://dialogflow.com/docs/actions-and-parameters) for
	// syntax.
	Parameters googleapi.RawMessage `json:"parameters,omitempty"`

	// ForceSendFields is a list of field names (e.g. "LifespanCount") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "LifespanCount") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2Context) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2Context
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2EntityType: Represents an entity type.
// Entity types serve as a tool for extracting parameter values from
// natural
// language queries.
type GoogleCloudDialogflowV2EntityType struct {
	// AutoExpansionMode: Optional. Indicates whether the entity type can be
	// automatically
	// expanded.
	//
	// Possible values:
	//   "AUTO_EXPANSION_MODE_UNSPECIFIED" - Auto expansion disabled for the
	// entity.
	//   "AUTO_EXPANSION_MODE_DEFAULT" - Allows an agent to recognize values
	// that have not been explicitly
	// listed in the entity.
	AutoExpansionMode string `json:"autoExpansionMode,omitempty"`

	// DisplayName: Required. The name of the entity type.
	DisplayName string `json:"displayName,omitempty"`

	// Entities: Optional. The collection of entities associated with the
	// entity type.
	Entities []*GoogleCloudDialogflowV2EntityTypeEntity `json:"entities,omitempty"`

	// Kind: Required. Indicates the kind of entity type.
	//
	// Possible values:
	//   "KIND_UNSPECIFIED" - Not specified. This value should be never
	// used.
	//   "KIND_MAP" - Map entity types allow mapping of a group of synonyms
	// to a canonical
	// value.
	//   "KIND_LIST" - List entity types contain a set of entries that do
	// not map to canonical
	// values. However, list entity types can contain references to other
	// entity
	// types (with or without aliases).
	Kind string `json:"kind,omitempty"`

	// Name: Required for all methods except `create` (`create` populates
	// the name
	// automatically.
	// The unique identifier of the entity type. Format:
	// `projects/<Project ID>/agent/entityTypes/<Entity Type ID>`.
	Name string `json:"name,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AutoExpansionMode")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AutoExpansionMode") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2EntityType) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2EntityType
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2EntityTypeEntity: Optional. Represents an
// entity.
type GoogleCloudDialogflowV2EntityTypeEntity struct {
	// Synonyms: Required. A collection of synonyms. For `KIND_LIST` entity
	// types this
	// must contain exactly one synonym equal to `value`.
	Synonyms []string `json:"synonyms,omitempty"`

	// Value: Required.
	// For `KIND_MAP` entity types:
	//   A canonical name to be used in place of synonyms.
	// For `KIND_LIST` entity types:
	//   A string that can contain references to other entity types (with
	// or
	//   without aliases).
	Value string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Synonyms") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Synonyms") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2EntityTypeEntity) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2EntityTypeEntity
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2EventInput: Events allow for matching intents
// by event name instead of the natural
// language input. For instance, input `<event: { name:
// “welcome_event”,
// parameters: { name: “Sam” } }>` can trigger a personalized
// welcome response.
// The parameter `name` may be used by the agent in the
// response:
// `“Hello #welcome_event.name! What can I do for you today?”`.
type GoogleCloudDialogflowV2EventInput struct {
	// LanguageCode: Required. The language of this query. See
	// [Language
	// Support](https://dialogflow.com/docs/languages) for a list of
	// the
	// currently supported language codes. Note that queries in the same
	// session
	// do not necessarily need to specify the same language.
	LanguageCode string `json:"languageCode,omitempty"`

	// Name: Required. The unique identifier of the event.
	Name string `json:"name,omitempty"`

	// Parameters: Optional. The collection of parameters associated with
	// the event.
	Parameters googleapi.RawMessage `json:"parameters,omitempty"`

	// ForceSendFields is a list of field names (e.g. "LanguageCode") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "LanguageCode") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2EventInput) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2EventInput
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2ExportAgentResponse: The response message for
// Agents.ExportAgent.
type GoogleCloudDialogflowV2ExportAgentResponse struct {
	// AgentContent: The exported agent.
	//
	// Example for how to export an agent to a zip file via a command
	// line:
	// <pre>curl \
	//
	// 'https://dialogflow.googleapis.com/v2/projects/&lt;project_name&gt;/ag
	// ent:export'\
	//   -X POST \
	//   -H 'Authorization: Bearer '$(gcloud auth application-default
	//   print-access-token) \
	//   -H 'Accept: application/json' \
	//   -H 'Content-Type: application/json' \
	//   --compressed \
	//   --data-binary '{}' \
	// | grep agentContent | sed -e 's/.*"agentContent": "\([^"]*\)".*/\1/'
	// \
	// | base64 --decode > &lt;agent zip file&gt;</pre>
	AgentContent string `json:"agentContent,omitempty"`

	// AgentUri: The URI to a file containing the exported agent. This field
	// is populated
	// only if `agent_uri` is specified in `ExportAgentRequest`.
	AgentUri string `json:"agentUri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AgentContent") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AgentContent") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2ExportAgentResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2ExportAgentResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2Intent: Represents an intent.
// Intents convert a number of user expressions or patterns into an
// action. An
// action is an extraction of a user command or sentence semantics.
type GoogleCloudDialogflowV2Intent struct {
	// Action: Optional. The name of the action associated with the
	// intent.
	// Note: The action name must not contain whitespaces.
	Action string `json:"action,omitempty"`

	// DefaultResponsePlatforms: Optional. The list of platforms for which
	// the first response will be
	// taken from among the messages assigned to the DEFAULT_PLATFORM.
	//
	// Possible values:
	//   "PLATFORM_UNSPECIFIED" - Not specified.
	//   "FACEBOOK" - Facebook.
	//   "SLACK" - Slack.
	//   "TELEGRAM" - Telegram.
	//   "KIK" - Kik.
	//   "SKYPE" - Skype.
	//   "LINE" - Line.
	//   "VIBER" - Viber.
	//   "ACTIONS_ON_GOOGLE" - Actions on Google.
	// When using Actions on Google, you can choose one of the
	// specific
	// Intent.Message types that mention support for Actions on Google,
	// or you can use the advanced Intent.Message.payload field.
	// The payload field provides access to AoG features not available in
	// the
	// specific message types.
	// If using the Intent.Message.payload field, it should have a
	// structure
	// similar to the JSON message shown here. For more information,
	// see
	// [Actions on Google
	// Webhook
	// Format](https://developers.google.com/actions/dialogflow/webho
	// ok)
	// <pre>{
	//   "expectUserResponse": true,
	//   "isSsml": false,
	//   "noInputPrompts": [],
	//   "richResponse": {
	//     "items": [
	//       {
	//         "simpleResponse": {
	//           "displayText": "hi",
	//           "textToSpeech": "hello"
	//         }
	//       }
	//     ],
	//     "suggestions": [
	//       {
	//         "title": "Say this"
	//       },
	//       {
	//         "title": "or this"
	//       }
	//     ]
	//   },
	//   "systemIntent": {
	//     "data": {
	//       "@type":
	// "type.googleapis.com/google.actions.v2.OptionValueSpec",
	//       "listSelect": {
	//         "items": [
	//           {
	//             "optionInfo": {
	//               "key": "key1",
	//               "synonyms": [
	//                 "key one"
	//               ]
	//             },
	//             "title": "must not be empty, but unique"
	//           },
	//           {
	//             "optionInfo": {
	//               "key": "key2",
	//               "synonyms": [
	//                 "key two"
	//               ]
	//             },
	//             "title": "must not be empty, but unique"
	//           }
	//         ]
	//       }
	//     },
	//     "intent": "actions.intent.OPTION"
	//   }
	// }</pre>
	DefaultResponsePlatforms []string `json:"defaultResponsePlatforms,omitempty"`

	// DisplayName: Required. The name of this intent.
	DisplayName string `json:"displayName,omitempty"`

	// Events: Optional. The collection of event names that trigger the
	// intent.
	// If the collection of input contexts is not empty, all of the contexts
	// must
	// be present in the active user session for an event to trigger this
	// intent.
	Events []string `json:"events,omitempty"`

	// FollowupIntentInfo: Read-only. Information about all followup intents
	// that have this intent as
	// a direct or indirect parent. We populate this field only in the
	// output.
	FollowupIntentInfo []*GoogleCloudDialogflowV2IntentFollowupIntentInfo `json:"followupIntentInfo,omitempty"`

	// InputContextNames: Optional. The list of context names required for
	// this intent to be
	// triggered.
	// Format: `projects/<Project ID>/agent/sessions/-/contexts/<Context
	// ID>`.
	InputContextNames []string `json:"inputContextNames,omitempty"`

	// IsFallback: Optional. Indicates whether this is a fallback intent.
	IsFallback bool `json:"isFallback,omitempty"`

	// Messages: Optional. The collection of rich messages corresponding to
	// the
	// `Response` field in the Dialogflow console.
	Messages []*GoogleCloudDialogflowV2IntentMessage `json:"messages,omitempty"`

	// MlDisabled: Optional. Indicates whether Machine Learning is disabled
	// for the intent.
	// Note: If `ml_diabled` setting is set to true, then this intent is
	// not
	// taken into account during inference in `ML ONLY` match mode.
	// Also,
	// auto-markup in the UI is turned off.
	MlDisabled bool `json:"mlDisabled,omitempty"`

	// Name: Required for all methods except `create` (`create` populates
	// the name
	// automatically.
	// The unique identifier of this intent.
	// Format: `projects/<Project ID>/agent/intents/<Intent ID>`.
	Name string `json:"name,omitempty"`

	// OutputContexts: Optional. The collection of contexts that are
	// activated when the intent
	// is matched. Context messages in this collection should not set
	// the
	// parameters field. Setting the `lifespan_count` to 0 will reset the
	// context
	// when the intent is matched.
	// Format: `projects/<Project ID>/agent/sessions/-/contexts/<Context
	// ID>`.
	OutputContexts []*GoogleCloudDialogflowV2Context `json:"outputContexts,omitempty"`

	// Parameters: Optional. The collection of parameters associated with
	// the intent.
	Parameters []*GoogleCloudDialogflowV2IntentParameter `json:"parameters,omitempty"`

	// ParentFollowupIntentName: Read-only after creation. The unique
	// identifier of the parent intent in the
	// chain of followup intents. You can set this field when creating an
	// intent,
	// for example with CreateIntent or BatchUpdateIntents, in order to
	// make this intent a followup intent.
	//
	// It identifies the parent followup intent.
	// Format: `projects/<Project ID>/agent/intents/<Intent ID>`.
	ParentFollowupIntentName string `json:"parentFollowupIntentName,omitempty"`

	// Priority: Optional. The priority of this intent. Higher numbers
	// represent higher
	// priorities. Zero or negative numbers mean that the intent is
	// disabled.
	Priority int64 `json:"priority,omitempty"`

	// ResetContexts: Optional. Indicates whether to delete all contexts in
	// the current
	// session when this intent is matched.
	ResetContexts bool `json:"resetContexts,omitempty"`

	// RootFollowupIntentName: Read-only. The unique identifier of the root
	// intent in the chain of
	// followup intents. It identifies the correct followup intents chain
	// for
	// this intent. We populate this field only in the output.
	//
	// Format: `projects/<Project ID>/agent/intents/<Intent ID>`.
	RootFollowupIntentName string `json:"rootFollowupIntentName,omitempty"`

	// TrainingPhrases: Optional. The collection of examples/templates that
	// the agent is
	// trained on.
	TrainingPhrases []*GoogleCloudDialogflowV2IntentTrainingPhrase `json:"trainingPhrases,omitempty"`

	// WebhookState: Optional. Indicates whether webhooks are enabled for
	// the intent.
	//
	// Possible values:
	//   "WEBHOOK_STATE_UNSPECIFIED" - Webhook is disabled in the agent and
	// in the intent.
	//   "WEBHOOK_STATE_ENABLED" - Webhook is enabled in the agent and in
	// the intent.
	//   "WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING" - Webhook is enabled in
	// the agent and in the intent. Also, each slot
	// filling prompt is forwarded to the webhook.
	WebhookState string `json:"webhookState,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Action") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Action") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2Intent) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2Intent
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2IntentFollowupIntentInfo: Represents a single
// followup intent in the chain.
type GoogleCloudDialogflowV2IntentFollowupIntentInfo struct {
	// FollowupIntentName: The unique identifier of the followup
	// intent.
	// Format: `projects/<Project ID>/agent/intents/<Intent ID>`.
	FollowupIntentName string `json:"followupIntentName,omitempty"`

	// ParentFollowupIntentName: The unique identifier of the followup
	// intent's parent.
	// Format: `projects/<Project ID>/agent/intents/<Intent ID>`.
	ParentFollowupIntentName string `json:"parentFollowupIntentName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FollowupIntentName")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FollowupIntentName") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2IntentFollowupIntentInfo) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2IntentFollowupIntentInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2IntentMessage: Corresponds to the `Response`
// field in the Dialogflow console.
type GoogleCloudDialogflowV2IntentMessage struct {
	// BasicCard: The basic card response for Actions on Google.
	BasicCard *GoogleCloudDialogflowV2IntentMessageBasicCard `json:"basicCard,omitempty"`

	// Card: The card response.
	Card *GoogleCloudDialogflowV2IntentMessageCard `json:"card,omitempty"`

	// CarouselSelect: The carousel card response for Actions on Google.
	CarouselSelect *GoogleCloudDialogflowV2IntentMessageCarouselSelect `json:"carouselSelect,omitempty"`

	// Image: The image response.
	Image *GoogleCloudDialogflowV2IntentMessageImage `json:"image,omitempty"`

	// LinkOutSuggestion: The link out suggestion chip for Actions on
	// Google.
	LinkOutSuggestion *GoogleCloudDialogflowV2IntentMessageLinkOutSuggestion `json:"linkOutSuggestion,omitempty"`

	// ListSelect: The list card response for Actions on Google.
	ListSelect *GoogleCloudDialogflowV2IntentMessageListSelect `json:"listSelect,omitempty"`

	// Payload: Returns a response containing a custom, platform-specific
	// payload.
	// See the Intent.Message.Platform type for a description of
	// the
	// structure that may be required for your platform.
	Payload googleapi.RawMessage `json:"payload,omitempty"`

	// Platform: Optional. The platform that this message is intended for.
	//
	// Possible values:
	//   "PLATFORM_UNSPECIFIED" - Not specified.
	//   "FACEBOOK" - Facebook.
	//   "SLACK" - Slack.
	//   "TELEGRAM" - Telegram.
	//   "KIK" - Kik.
	//   "SKYPE" - Skype.
	//   "LINE" - Line.
	//   "VIBER" - Viber.
	//   "ACTIONS_ON_GOOGLE" - Actions on Google.
	// When using Actions on Google, you can choose one of the
	// specific
	// Intent.Message types that mention support for Actions on Google,
	// or you can use the advanced Intent.Message.payload field.
	// The payload field provides access to AoG features not available in
	// the
	// specific message types.
	// If using the Intent.Message.payload field, it should have a
	// structure
	// similar to the JSON message shown here. For more information,
	// see
	// [Actions on Google
	// Webhook
	// Format](https://developers.google.com/actions/dialogflow/webho
	// ok)
	// <pre>{
	//   "expectUserResponse": true,
	//   "isSsml": false,
	//   "noInputPrompts": [],
	//   "richResponse": {
	//     "items": [
	//       {
	//         "simpleResponse": {
	//           "displayText": "hi",
	//           "textToSpeech": "hello"
	//         }
	//       }
	//     ],
	//     "suggestions": [
	//       {
	//         "title": "Say this"
	//       },
	//       {
	//         "title": "or this"
	//       }
	//     ]
	//   },
	//   "systemIntent": {
	//     "data": {
	//       "@type":
	// "type.googleapis.com/google.actions.v2.OptionValueSpec",
	//       "listSelect": {
	//         "items": [
	//           {
	//             "optionInfo": {
	//               "key": "key1",
	//               "synonyms": [
	//                 "key one"
	//               ]
	//             },
	//             "title": "must not be empty, but unique"
	//           },
	//           {
	//             "optionInfo": {
	//               "key": "key2",
	//               "synonyms": [
	//                 "key two"
	//               ]
	//             },
	//             "title": "must not be empty, but unique"
	//           }
	//         ]
	//       }
	//     },
	//     "intent": "actions.intent.OPTION"
	//   }
	// }</pre>
	Platform string `json:"platform,omitempty"`

	// QuickReplies: The quick replies response.
	QuickReplies *GoogleCloudDialogflowV2IntentMessageQuickReplies `json:"quickReplies,omitempty"`

	// SimpleResponses: The voice and text-only responses for Actions on
	// Google.
	SimpleResponses *GoogleCloudDialogflowV2IntentMessageSimpleResponses `json:"simpleResponses,omitempty"`

	// Suggestions: The suggestion chips for Actions on Google.
	Suggestions *GoogleCloudDialogflowV2IntentMessageSuggestions `json:"suggestions,omitempty"`

	// Text: The text response.
	Text *GoogleCloudDialogflowV2IntentMessageText `json:"text,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BasicCard") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BasicCard") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2IntentMessage) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2IntentMessage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2IntentMessageBasicCard: The basic card
// message. Useful for displaying information.
type GoogleCloudDialogflowV2IntentMessageBasicCard struct {
	// Buttons: Optional. The collection of card buttons.
	Buttons []*GoogleCloudDialogflowV2IntentMessageBasicCardButton `json:"buttons,omitempty"`

	// FormattedText: Required, unless image is present. The body text of
	// the card.
	FormattedText string `json:"formattedText,omitempty"`

	// Image: Optional. The image for the card.
	Image *GoogleCloudDialogflowV2IntentMessageImage `json:"image,omitempty"`

	// Subtitle: Optional. The subtitle of the card.
	Subtitle string `json:"subtitle,omitempty"`

	// Title: Optional. The title of the card.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Buttons") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Buttons") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2IntentMessageBasicCard) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2IntentMessageBasicCard
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2IntentMessageBasicCardButton: The button
// object that appears at the bottom of a card.
type GoogleCloudDialogflowV2IntentMessageBasicCardButton struct {
	// OpenUriAction: Required. Action to take when a user taps on the
	// button.
	OpenUriAction *GoogleCloudDialogflowV2IntentMessageBasicCardButtonOpenUriAction `json:"openUriAction,omitempty"`

	// Title: Required. The title of the button.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "OpenUriAction") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "OpenUriAction") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2IntentMessageBasicCardButton) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2IntentMessageBasicCardButton
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2IntentMessageBasicCardButtonOpenUriAction:
// Opens the given URI.
type GoogleCloudDialogflowV2IntentMessageBasicCardButtonOpenUriAction struct {
	// Uri: Required. The HTTP or HTTPS scheme URI.
	Uri string `json:"uri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Uri") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Uri") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2IntentMessageBasicCardButtonOpenUriAction) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2IntentMessageBasicCardButtonOpenUriAction
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2IntentMessageCard: The card response message.
type GoogleCloudDialogflowV2IntentMessageCard struct {
	// Buttons: Optional. The collection of card buttons.
	Buttons []*GoogleCloudDialogflowV2IntentMessageCardButton `json:"buttons,omitempty"`

	// ImageUri: Optional. The public URI to an image file for the card.
	ImageUri string `json:"imageUri,omitempty"`

	// Subtitle: Optional. The subtitle of the card.
	Subtitle string `json:"subtitle,omitempty"`

	// Title: Optional. The title of the card.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Buttons") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Buttons") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2IntentMessageCard) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2IntentMessageCard
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2IntentMessageCardButton: Optional. Contains
// information about a button.
type GoogleCloudDialogflowV2IntentMessageCardButton struct {
	// Postback: Optional. The text to send back to the Dialogflow API or a
	// URI to
	// open.
	Postback string `json:"postback,omitempty"`

	// Text: Optional. The text to show on the button.
	Text string `json:"text,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Postback") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Postback") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2IntentMessageCardButton) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2IntentMessageCardButton
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2IntentMessageCarouselSelect: The card for
// presenting a carousel of options to select from.
type GoogleCloudDialogflowV2IntentMessageCarouselSelect struct {
	// Items: Required. Carousel items.
	Items []*GoogleCloudDialogflowV2IntentMessageCarouselSelectItem `json:"items,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2IntentMessageCarouselSelect) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2IntentMessageCarouselSelect
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2IntentMessageCarouselSelectItem: An item in
// the carousel.
type GoogleCloudDialogflowV2IntentMessageCarouselSelectItem struct {
	// Description: Optional. The body text of the card.
	Description string `json:"description,omitempty"`

	// Image: Optional. The image to display.
	Image *GoogleCloudDialogflowV2IntentMessageImage `json:"image,omitempty"`

	// Info: Required. Additional info about the option item.
	Info *GoogleCloudDialogflowV2IntentMessageSelectItemInfo `json:"info,omitempty"`

	// Title: Required. Title of the carousel item.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2IntentMessageCarouselSelectItem) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2IntentMessageCarouselSelectItem
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2IntentMessageImage: The image response
// message.
type GoogleCloudDialogflowV2IntentMessageImage struct {
	// AccessibilityText: Optional. A text description of the image to be
	// used for accessibility,
	// e.g., screen readers.
	AccessibilityText string `json:"accessibilityText,omitempty"`

	// ImageUri: Optional. The public URI to an image file.
	ImageUri string `json:"imageUri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AccessibilityText")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AccessibilityText") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2IntentMessageImage) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2IntentMessageImage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2IntentMessageLinkOutSuggestion: The suggestion
// chip message that allows the user to jump out to the app
// or website associated with this agent.
type GoogleCloudDialogflowV2IntentMessageLinkOutSuggestion struct {
	// DestinationName: Required. The name of the app or site this chip is
	// linking to.
	DestinationName string `json:"destinationName,omitempty"`

	// Uri: Required. The URI of the app or site to open when the user taps
	// the
	// suggestion chip.
	Uri string `json:"uri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DestinationName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DestinationName") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2IntentMessageLinkOutSuggestion) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2IntentMessageLinkOutSuggestion
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2IntentMessageListSelect: The card for
// presenting a list of options to select from.
type GoogleCloudDialogflowV2IntentMessageListSelect struct {
	// Items: Required. List items.
	Items []*GoogleCloudDialogflowV2IntentMessageListSelectItem `json:"items,omitempty"`

	// Title: Optional. The overall title of the list.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2IntentMessageListSelect) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2IntentMessageListSelect
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2IntentMessageListSelectItem: An item in the
// list.
type GoogleCloudDialogflowV2IntentMessageListSelectItem struct {
	// Description: Optional. The main text describing the item.
	Description string `json:"description,omitempty"`

	// Image: Optional. The image to display.
	Image *GoogleCloudDialogflowV2IntentMessageImage `json:"image,omitempty"`

	// Info: Required. Additional information about this option.
	Info *GoogleCloudDialogflowV2IntentMessageSelectItemInfo `json:"info,omitempty"`

	// Title: Required. The title of the list item.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2IntentMessageListSelectItem) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2IntentMessageListSelectItem
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2IntentMessageQuickReplies: The quick replies
// response message.
type GoogleCloudDialogflowV2IntentMessageQuickReplies struct {
	// QuickReplies: Optional. The collection of quick replies.
	QuickReplies []string `json:"quickReplies,omitempty"`

	// Title: Optional. The title of the collection of quick replies.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "QuickReplies") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "QuickReplies") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2IntentMessageQuickReplies) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2IntentMessageQuickReplies
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2IntentMessageSelectItemInfo: Additional info
// about the select item for when it is triggered in a
// dialog.
type GoogleCloudDialogflowV2IntentMessageSelectItemInfo struct {
	// Key: Required. A unique key that will be sent back to the agent if
	// this
	// response is given.
	Key string `json:"key,omitempty"`

	// Synonyms: Optional. A list of synonyms that can also be used to
	// trigger this
	// item in dialog.
	Synonyms []string `json:"synonyms,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Key") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Key") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2IntentMessageSelectItemInfo) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2IntentMessageSelectItemInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2IntentMessageSimpleResponse: The simple
// response message containing speech or text.
type GoogleCloudDialogflowV2IntentMessageSimpleResponse struct {
	// DisplayText: Optional. The text to display.
	DisplayText string `json:"displayText,omitempty"`

	// Ssml: One of text_to_speech or ssml must be provided. Structured
	// spoken
	// response to the user in the SSML format. Mutually exclusive
	// with
	// text_to_speech.
	Ssml string `json:"ssml,omitempty"`

	// TextToSpeech: One of text_to_speech or ssml must be provided. The
	// plain text of the
	// speech output. Mutually exclusive with ssml.
	TextToSpeech string `json:"textToSpeech,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DisplayText") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DisplayText") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2IntentMessageSimpleResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2IntentMessageSimpleResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2IntentMessageSimpleResponses: The collection
// of simple response candidates.
// This message in `QueryResult.fulfillment_messages`
// and
// `WebhookResponse.fulfillment_messages` should contain only
// one
// `SimpleResponse`.
type GoogleCloudDialogflowV2IntentMessageSimpleResponses struct {
	// SimpleResponses: Required. The list of simple responses.
	SimpleResponses []*GoogleCloudDialogflowV2IntentMessageSimpleResponse `json:"simpleResponses,omitempty"`

	// ForceSendFields is a list of field names (e.g. "SimpleResponses") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "SimpleResponses") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2IntentMessageSimpleResponses) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2IntentMessageSimpleResponses
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2IntentMessageSuggestion: The suggestion chip
// message that the user can tap to quickly post a reply
// to the conversation.
type GoogleCloudDialogflowV2IntentMessageSuggestion struct {
	// Title: Required. The text shown the in the suggestion chip.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Title") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Title") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2IntentMessageSuggestion) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2IntentMessageSuggestion
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2IntentMessageSuggestions: The collection of
// suggestions.
type GoogleCloudDialogflowV2IntentMessageSuggestions struct {
	// Suggestions: Required. The list of suggested replies.
	Suggestions []*GoogleCloudDialogflowV2IntentMessageSuggestion `json:"suggestions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Suggestions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Suggestions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2IntentMessageSuggestions) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2IntentMessageSuggestions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2IntentMessageText: The text response message.
type GoogleCloudDialogflowV2IntentMessageText struct {
	// Text: Optional. The collection of the agent's responses.
	Text []string `json:"text,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Text") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Text") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2IntentMessageText) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2IntentMessageText
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2IntentParameter: Represents intent parameters.
type GoogleCloudDialogflowV2IntentParameter struct {
	// DefaultValue: Optional. The default value to use when the `value`
	// yields an empty
	// result.
	// Default values can be extracted from contexts by using the
	// following
	// syntax: `#context_name.parameter_name`.
	DefaultValue string `json:"defaultValue,omitempty"`

	// DisplayName: Required. The name of the parameter.
	DisplayName string `json:"displayName,omitempty"`

	// EntityTypeDisplayName: Optional. The name of the entity type,
	// prefixed with `@`, that
	// describes values of the parameter. If the parameter is
	// required, this must be provided.
	EntityTypeDisplayName string `json:"entityTypeDisplayName,omitempty"`

	// IsList: Optional. Indicates whether the parameter represents a list
	// of values.
	IsList bool `json:"isList,omitempty"`

	// Mandatory: Optional. Indicates whether the parameter is required.
	// That is,
	// whether the intent cannot be completed without collecting the
	// parameter
	// value.
	Mandatory bool `json:"mandatory,omitempty"`

	// Name: The unique identifier of this parameter.
	Name string `json:"name,omitempty"`

	// Prompts: Optional. The collection of prompts that the agent can
	// present to the
	// user in order to collect value for the parameter.
	Prompts []string `json:"prompts,omitempty"`

	// Value: Optional. The definition of the parameter value. It can be:
	// - a constant string,
	// - a parameter value defined as `$parameter_name`,
	// - an original parameter value defined as
	// `$parameter_name.original`,
	// - a parameter value from some context defined as
	//   `#context_name.parameter_name`.
	Value string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DefaultValue") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DefaultValue") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2IntentParameter) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2IntentParameter
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2IntentTrainingPhrase: Represents an example or
// template that the agent is trained on.
type GoogleCloudDialogflowV2IntentTrainingPhrase struct {
	// Name: Output only. The unique identifier of this training phrase.
	Name string `json:"name,omitempty"`

	// Parts: Required. The collection of training phrase parts (can be
	// annotated).
	// Fields: `entity_type`, `alias` and `user_defined` should be
	// populated
	// only for the annotated parts of the training phrase.
	Parts []*GoogleCloudDialogflowV2IntentTrainingPhrasePart `json:"parts,omitempty"`

	// TimesAddedCount: Optional. Indicates how many times this example or
	// template was added to
	// the intent. Each time a developer adds an existing sample by editing
	// an
	// intent or training, this counter is increased.
	TimesAddedCount int64 `json:"timesAddedCount,omitempty"`

	// Type: Required. The type of the training phrase.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED" - Not specified. This value should never be
	// used.
	//   "EXAMPLE" - Examples do not contain @-prefixed entity type names,
	// but example parts
	// can be annotated with entity types.
	//   "TEMPLATE" - Templates are not annotated with entity types, but
	// they can contain
	// @-prefixed entity type names as substrings.
	Type string `json:"type,omitempty"`

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

func (s *GoogleCloudDialogflowV2IntentTrainingPhrase) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2IntentTrainingPhrase
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2IntentTrainingPhrasePart: Represents a part of
// a training phrase.
type GoogleCloudDialogflowV2IntentTrainingPhrasePart struct {
	// Alias: Optional. The parameter name for the value extracted from
	// the
	// annotated part of the example.
	Alias string `json:"alias,omitempty"`

	// EntityType: Optional. The entity type name prefixed with `@`. This
	// field is
	// required for the annotated part of the text and applies only
	// to
	// examples.
	EntityType string `json:"entityType,omitempty"`

	// Text: Required. The text corresponding to the example or template,
	// if there are no annotations. For
	// annotated examples, it is the text for one of the example's parts.
	Text string `json:"text,omitempty"`

	// UserDefined: Optional. Indicates whether the text was manually
	// annotated by the
	// developer.
	UserDefined bool `json:"userDefined,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Alias") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Alias") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2IntentTrainingPhrasePart) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2IntentTrainingPhrasePart
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2OriginalDetectIntentRequest: Represents the
// contents of the original request that was passed to
// the `[Streaming]DetectIntent` call.
type GoogleCloudDialogflowV2OriginalDetectIntentRequest struct {
	// Payload: Optional. This field is set to the value of the
	// `QueryParameters.payload`
	// field passed in the request. Some integrations that query a
	// Dialogflow
	// agent may provide additional information in the payload.
	//
	// In particular for the Telephony Gateway this field has the
	// form:
	// <pre>{
	//  "telephony": {
	//    "caller_id": "+18558363987"
	//  }
	// }</pre>
	// Note: The caller ID field (`caller_id`) will be redacted for
	// Standard
	// Edition agents and populated with the caller ID in
	// [E.164
	// format](https://en.wikipedia.org/wiki/E.164) for Enterprise Edition
	// agents.
	Payload googleapi.RawMessage `json:"payload,omitempty"`

	// Source: The source of this request, e.g., `google`, `facebook`,
	// `slack`. It is set
	// by Dialogflow-owned servers.
	Source string `json:"source,omitempty"`

	// Version: Optional. The version of the protocol used for this
	// request.
	// This field is AoG-specific.
	Version string `json:"version,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Payload") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Payload") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2OriginalDetectIntentRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2OriginalDetectIntentRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2QueryResult: Represents the result of
// conversational query or event processing.
type GoogleCloudDialogflowV2QueryResult struct {
	// Action: The action name from the matched intent.
	Action string `json:"action,omitempty"`

	// AllRequiredParamsPresent: This field is set to:
	// - `false` if the matched intent has required parameters and not all
	// of
	//    the required parameter values have been collected.
	// - `true` if all required parameter values have been collected, or if
	// the
	//    matched intent doesn't contain any required parameters.
	AllRequiredParamsPresent bool `json:"allRequiredParamsPresent,omitempty"`

	// DiagnosticInfo: The free-form diagnostic info. For example, this
	// field
	// could contain webhook call latency.
	DiagnosticInfo googleapi.RawMessage `json:"diagnosticInfo,omitempty"`

	// FulfillmentMessages: The collection of rich messages to present to
	// the user.
	FulfillmentMessages []*GoogleCloudDialogflowV2IntentMessage `json:"fulfillmentMessages,omitempty"`

	// FulfillmentText: The text to be pronounced to the user or shown on
	// the screen.
	FulfillmentText string `json:"fulfillmentText,omitempty"`

	// Intent: The intent that matched the conversational query. Some,
	// not
	// all fields are filled in this message, including but not limited
	// to:
	// `name`, `display_name` and `webhook_state`.
	Intent *GoogleCloudDialogflowV2Intent `json:"intent,omitempty"`

	// IntentDetectionConfidence: The intent detection confidence. Values
	// range from 0.0
	// (completely uncertain) to 1.0 (completely certain).
	// If there are `multiple knowledge_answers` messages, this value is set
	// to
	// the greatest `knowledgeAnswers.match_confidence` value in the list.
	IntentDetectionConfidence float64 `json:"intentDetectionConfidence,omitempty"`

	// LanguageCode: The language that was triggered during intent
	// detection.
	// See [Language
	// Support](https://dialogflow.com/docs/reference/language)
	// for a list of the currently supported language codes.
	LanguageCode string `json:"languageCode,omitempty"`

	// OutputContexts: The collection of output contexts. If
	// applicable,
	// `output_contexts.parameters` contains entries with name
	// `<parameter name>.original` containing the original parameter
	// values
	// before the query.
	OutputContexts []*GoogleCloudDialogflowV2Context `json:"outputContexts,omitempty"`

	// Parameters: The collection of extracted parameters.
	Parameters googleapi.RawMessage `json:"parameters,omitempty"`

	// QueryText: The original conversational query text:
	// - If natural language text was provided as input, `query_text`
	// contains
	//   a copy of the input.
	// - If natural language speech audio was provided as input,
	// `query_text`
	//   contains the speech recognition result. If speech recognizer
	// produced
	//   multiple alternatives, a particular one is picked.
	// - If an event was provided as input, `query_text` is not set.
	QueryText string `json:"queryText,omitempty"`

	// SpeechRecognitionConfidence: The Speech recognition confidence
	// between 0.0 and 1.0. A higher number
	// indicates an estimated greater likelihood that the recognized words
	// are
	// correct. The default of 0.0 is a sentinel value indicating that
	// confidence
	// was not set.
	//
	// This field is not guaranteed to be accurate or set. In particular
	// this
	// field isn't set for StreamingDetectIntent since the streaming
	// endpoint has
	// separate confidence estimates per portion of the audio
	// in
	// StreamingRecognitionResult.
	SpeechRecognitionConfidence float64 `json:"speechRecognitionConfidence,omitempty"`

	// WebhookPayload: If the query was fulfilled by a webhook call, this
	// field is set to the
	// value of the `payload` field returned in the webhook response.
	WebhookPayload googleapi.RawMessage `json:"webhookPayload,omitempty"`

	// WebhookSource: If the query was fulfilled by a webhook call, this
	// field is set to the
	// value of the `source` field returned in the webhook response.
	WebhookSource string `json:"webhookSource,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Action") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Action") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2QueryResult) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2QueryResult
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudDialogflowV2QueryResult) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudDialogflowV2QueryResult
	var s1 struct {
		IntentDetectionConfidence   gensupport.JSONFloat64 `json:"intentDetectionConfidence"`
		SpeechRecognitionConfidence gensupport.JSONFloat64 `json:"speechRecognitionConfidence"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.IntentDetectionConfidence = float64(s1.IntentDetectionConfidence)
	s.SpeechRecognitionConfidence = float64(s1.SpeechRecognitionConfidence)
	return nil
}

// GoogleCloudDialogflowV2WebhookRequest: The request message for a
// webhook call.
type GoogleCloudDialogflowV2WebhookRequest struct {
	// OriginalDetectIntentRequest: Optional. The contents of the original
	// request that was passed to
	// `[Streaming]DetectIntent` call.
	OriginalDetectIntentRequest *GoogleCloudDialogflowV2OriginalDetectIntentRequest `json:"originalDetectIntentRequest,omitempty"`

	// QueryResult: The result of the conversational query or event
	// processing. Contains the
	// same value as `[Streaming]DetectIntentResponse.query_result`.
	QueryResult *GoogleCloudDialogflowV2QueryResult `json:"queryResult,omitempty"`

	// ResponseId: The unique identifier of the response. Contains the same
	// value as
	// `[Streaming]DetectIntentResponse.response_id`.
	ResponseId string `json:"responseId,omitempty"`

	// Session: The unique identifier of detectIntent request session.
	// Can be used to identify end-user inside webhook
	// implementation.
	// Format: `projects/<Project ID>/agent/sessions/<Session ID>`.
	Session string `json:"session,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "OriginalDetectIntentRequest") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g.
	// "OriginalDetectIntentRequest") to include in API requests with the
	// JSON null value. By default, fields with empty values are omitted
	// from API requests. However, any field with an empty value appearing
	// in NullFields will be sent to the server as null. It is an error if a
	// field in this list has a non-empty value. This may be used to include
	// null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2WebhookRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2WebhookRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2WebhookResponse: The response message for a
// webhook call.
type GoogleCloudDialogflowV2WebhookResponse struct {
	// FollowupEventInput: Optional. Makes the platform immediately invoke
	// another `DetectIntent` call
	// internally with the specified event as input.
	FollowupEventInput *GoogleCloudDialogflowV2EventInput `json:"followupEventInput,omitempty"`

	// FulfillmentMessages: Optional. The collection of rich messages to
	// present to the user. This
	// value is passed directly to `QueryResult.fulfillment_messages`.
	FulfillmentMessages []*GoogleCloudDialogflowV2IntentMessage `json:"fulfillmentMessages,omitempty"`

	// FulfillmentText: Optional. The text to be shown on the screen. This
	// value is passed directly
	// to `QueryResult.fulfillment_text`.
	FulfillmentText string `json:"fulfillmentText,omitempty"`

	// OutputContexts: Optional. The collection of output contexts. This
	// value is passed directly
	// to `QueryResult.output_contexts`.
	OutputContexts []*GoogleCloudDialogflowV2Context `json:"outputContexts,omitempty"`

	// Payload: Optional. This value is passed directly to
	// `QueryResult.webhook_payload`.
	// See the related `fulfillment_messages[i].payload field`, which may be
	// used
	// as an alternative to this field.
	//
	// This field can be used for Actions on Google responses.
	// It should have a structure similar to the JSON message shown here.
	// For more
	// information, see
	// [Actions on Google
	// Webhook
	// Format](https://developers.google.com/actions/dialogflow/webho
	// ok)
	// <pre>{
	//   "google": {
	//     "expectUserResponse": true,
	//     "richResponse": {
	//       "items": [
	//         {
	//           "simpleResponse": {
	//             "textToSpeech": "this is a simple response"
	//           }
	//         }
	//       ]
	//     }
	//   }
	// }</pre>
	Payload googleapi.RawMessage `json:"payload,omitempty"`

	// Source: Optional. This value is passed directly to
	// `QueryResult.webhook_source`.
	Source string `json:"source,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FollowupEventInput")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FollowupEventInput") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2WebhookResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2WebhookResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1AddConversationPhoneNumberRequest: The
// request message for Conversations.AddConversationPhoneNumber.
type GoogleCloudDialogflowV2beta1AddConversationPhoneNumberRequest struct {
}

// GoogleCloudDialogflowV2beta1Agent: Represents a conversational agent.
type GoogleCloudDialogflowV2beta1Agent struct {
	// AvatarUri: Optional. The URI of the agent's avatar.
	// Avatars are used throughout the Dialogflow console and in the
	// self-hosted
	// [Web Demo](https://dialogflow.com/docs/integrations/web-demo)
	// integration.
	AvatarUri string `json:"avatarUri,omitempty"`

	// ClassificationThreshold: Optional. To filter out false positive
	// results and still get variety in
	// matched natural language inputs for your agent, you can tune the
	// machine
	// learning classification threshold. If the returned score value is
	// less than
	// the threshold value, then a fallback intent is be triggered or, if
	// there
	// are no fallback intents defined, no intent will be triggered. The
	// score
	// values range from 0.0 (completely uncertain) to 1.0 (completely
	// certain).
	// If set to 0.0, the default of 0.3 is used.
	ClassificationThreshold float64 `json:"classificationThreshold,omitempty"`

	// DefaultLanguageCode: Required. The default language of the agent as a
	// language tag. See
	// [Language Support](https://dialogflow.com/docs/reference/language)
	// for a
	// list of the currently supported language codes.
	// This field cannot be set by the `Update` method.
	DefaultLanguageCode string `json:"defaultLanguageCode,omitempty"`

	// Description: Optional. The description of this agent.
	// The maximum length is 500 characters. If exceeded, the request is
	// rejected.
	Description string `json:"description,omitempty"`

	// DisplayName: Required. The name of this agent.
	DisplayName string `json:"displayName,omitempty"`

	// EnableLogging: Optional. Determines whether this agent should log
	// conversation queries.
	EnableLogging bool `json:"enableLogging,omitempty"`

	// MatchMode: Optional. Determines how intents are detected from user
	// queries.
	//
	// Possible values:
	//   "MATCH_MODE_UNSPECIFIED" - Not specified.
	//   "MATCH_MODE_HYBRID" - Best for agents with a small number of
	// examples in intents and/or wide
	// use of templates syntax and composite entities.
	//   "MATCH_MODE_ML_ONLY" - Can be used for agents with a large number
	// of examples in intents,
	// especially the ones using @sys.any or very large developer entities.
	MatchMode string `json:"matchMode,omitempty"`

	// Parent: Required. The project of this agent.
	// Format: `projects/<Project ID>`.
	Parent string `json:"parent,omitempty"`

	// SupportedLanguageCodes: Optional. The list of all languages supported
	// by this agent (except for the
	// `default_language_code`).
	SupportedLanguageCodes []string `json:"supportedLanguageCodes,omitempty"`

	// TimeZone: Required. The time zone of this agent from the
	// [time zone database](https://www.iana.org/time-zones),
	// e.g.,
	// America/New_York, Europe/Paris.
	TimeZone string `json:"timeZone,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AvatarUri") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AvatarUri") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1Agent) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1Agent
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudDialogflowV2beta1Agent) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudDialogflowV2beta1Agent
	var s1 struct {
		ClassificationThreshold gensupport.JSONFloat64 `json:"classificationThreshold"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.ClassificationThreshold = float64(s1.ClassificationThreshold)
	return nil
}

// GoogleCloudDialogflowV2beta1AnalyzeContentRequest: The request
// message for Conversations.AnalyzeContent.
type GoogleCloudDialogflowV2beta1AnalyzeContentRequest struct {
	// Audio: The natural language speech audio to be processed.
	Audio *GoogleCloudDialogflowV2beta1InputAudio `json:"audio,omitempty"`

	// ReplyAudioConfig: Optional. Instructs the speech synthesizer how to
	// generate the output
	// audio.
	ReplyAudioConfig *GoogleCloudDialogflowV2beta1OutputAudioConfig `json:"replyAudioConfig,omitempty"`

	// Text: The natural language text to be processed.
	Text *GoogleCloudDialogflowV2beta1InputText `json:"text,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Audio") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Audio") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1AnalyzeContentRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1AnalyzeContentRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1AnalyzeContentResponse: The response
// message for Conversations.AnalyzeContent.
type GoogleCloudDialogflowV2beta1AnalyzeContentResponse struct {
	// AutomatedAgentReply: Optional. Only set if a Dialogflow automated
	// agent has responded.
	AutomatedAgentReply *GoogleCloudDialogflowV2beta1AutomatedAgentReply `json:"automatedAgentReply,omitempty"`

	// ReplyAudio: Optional. The audio data bytes encoded as specified in
	// the request.
	// This field is set if:
	//  - `reply_audio_config` was specified in the request, or
	//  - The automated agent responded with audio to play to the user. In
	// such
	//    case, `reply_audio.config` contains settings used to synthesize
	// the
	//    speech.
	ReplyAudio *GoogleCloudDialogflowV2beta1OutputAudio `json:"replyAudio,omitempty"`

	// ReplyText: Output only. The output text content.
	// This field is set if the automated agent responded with text to show
	// to
	// the user.
	ReplyText string `json:"replyText,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AutomatedAgentReply")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AutomatedAgentReply") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1AnalyzeContentResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1AnalyzeContentResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1ArticleSuggestionConfig: Defines article
// suggestions that a human agent assistant can provide.
type GoogleCloudDialogflowV2beta1ArticleSuggestionConfig struct {
	// KnowledgeBaseName: Required. Settings for knowledge base,
	// Format:
	// `projects/<Project ID>/knowledgeBases/<Knowledge Base ID>`.
	KnowledgeBaseName string `json:"knowledgeBaseName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "KnowledgeBaseName")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "KnowledgeBaseName") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1ArticleSuggestionConfig) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1ArticleSuggestionConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1AutomatedAgentConfig: Defines the
// Automated Agent to connect to a conversation.
type GoogleCloudDialogflowV2beta1AutomatedAgentConfig struct {
	// Agent: Required. ID of the Dialogflow agent environment to use.
	//
	// This project needs to either be the same project as the conversation
	// or you
	// need to grant `service-<Conversation
	// Project
	// Number>@gcp-sa-dialogflow.iam.gserviceaccount.com` the `Dialogflow
	// API
	// Service Agent` role in this project.
	//
	// Format: `projects/<Project ID>/agent/environments/<Environment ID or
	// '-'>`
	// If environment is not specified, the default `draft` environment
	// is
	// used. Refer
	// to
	// [DetectIntentRequest](/dialogflow-enterprise/docs/reference/rpc/goo
	// gle.cloud.dialogflow.v2beta1#google.cloud.dialogflow.v2beta1.DetectInt
	// entRequest)
	// for more details.
	Agent string `json:"agent,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Agent") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Agent") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1AutomatedAgentConfig) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1AutomatedAgentConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1AutomatedAgentReply: Represent a response
// from an automated agent.
type GoogleCloudDialogflowV2beta1AutomatedAgentReply struct {
	// DetectIntentResponse: Required. Response of the Dialogflow
	// Sessions.DetectIntent call.
	DetectIntentResponse *GoogleCloudDialogflowV2beta1DetectIntentResponse `json:"detectIntentResponse,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "DetectIntentResponse") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DetectIntentResponse") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1AutomatedAgentReply) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1AutomatedAgentReply
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1BatchCreateEntitiesRequest: The request
// message for EntityTypes.BatchCreateEntities.
type GoogleCloudDialogflowV2beta1BatchCreateEntitiesRequest struct {
	// Entities: Required. The entities to create.
	Entities []*GoogleCloudDialogflowV2beta1EntityTypeEntity `json:"entities,omitempty"`

	// LanguageCode: Optional. The language of entity synonyms defined in
	// `entities`. If not
	// specified, the agent's default language is used.
	// [More than a
	// dozen
	// languages](https://dialogflow.com/docs/reference/language) are
	// supported.
	// Note: languages must be enabled in the agent, before they can be
	// used.
	LanguageCode string `json:"languageCode,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Entities") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Entities") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1BatchCreateEntitiesRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1BatchCreateEntitiesRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1BatchDeleteEntitiesRequest: The request
// message for EntityTypes.BatchDeleteEntities.
type GoogleCloudDialogflowV2beta1BatchDeleteEntitiesRequest struct {
	// EntityValues: Required. The canonical `values` of the entities to
	// delete. Note that
	// these are not fully-qualified names, i.e. they don't start
	// with
	// `projects/<Project ID>`.
	EntityValues []string `json:"entityValues,omitempty"`

	// LanguageCode: Optional. The language of entity synonyms defined in
	// `entities`. If not
	// specified, the agent's default language is used.
	// [More than a
	// dozen
	// languages](https://dialogflow.com/docs/reference/language) are
	// supported.
	// Note: languages must be enabled in the agent, before they can be
	// used.
	LanguageCode string `json:"languageCode,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EntityValues") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EntityValues") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1BatchDeleteEntitiesRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1BatchDeleteEntitiesRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1BatchDeleteEntityTypesRequest: The
// request message for EntityTypes.BatchDeleteEntityTypes.
type GoogleCloudDialogflowV2beta1BatchDeleteEntityTypesRequest struct {
	// EntityTypeNames: Required. The names entity types to delete. All
	// names must point to the
	// same agent as `parent`.
	EntityTypeNames []string `json:"entityTypeNames,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EntityTypeNames") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EntityTypeNames") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1BatchDeleteEntityTypesRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1BatchDeleteEntityTypesRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1BatchDeleteIntentsRequest: The request
// message for Intents.BatchDeleteIntents.
type GoogleCloudDialogflowV2beta1BatchDeleteIntentsRequest struct {
	// Intents: Required. The collection of intents to delete. Only intent
	// `name` must be
	// filled in.
	Intents []*GoogleCloudDialogflowV2beta1Intent `json:"intents,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Intents") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Intents") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1BatchDeleteIntentsRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1BatchDeleteIntentsRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1BatchUpdateEntitiesRequest: The request
// message for EntityTypes.BatchUpdateEntities.
type GoogleCloudDialogflowV2beta1BatchUpdateEntitiesRequest struct {
	// Entities: Required. The entities to update or create.
	Entities []*GoogleCloudDialogflowV2beta1EntityTypeEntity `json:"entities,omitempty"`

	// LanguageCode: Optional. The language of entity synonyms defined in
	// `entities`. If not
	// specified, the agent's default language is used.
	// [More than a
	// dozen
	// languages](https://dialogflow.com/docs/reference/language) are
	// supported.
	// Note: languages must be enabled in the agent, before they can be
	// used.
	LanguageCode string `json:"languageCode,omitempty"`

	// UpdateMask: Optional. The mask to control which fields get updated.
	UpdateMask string `json:"updateMask,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Entities") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Entities") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1BatchUpdateEntitiesRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1BatchUpdateEntitiesRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1BatchUpdateEntityTypesRequest: The
// request message for EntityTypes.BatchUpdateEntityTypes.
type GoogleCloudDialogflowV2beta1BatchUpdateEntityTypesRequest struct {
	// EntityTypeBatchInline: The collection of entity types to update or
	// create.
	EntityTypeBatchInline *GoogleCloudDialogflowV2beta1EntityTypeBatch `json:"entityTypeBatchInline,omitempty"`

	// EntityTypeBatchUri: The URI to a Google Cloud Storage file containing
	// entity types to update
	// or create. The file format can either be a serialized proto
	// (of
	// EntityBatch type) or a JSON object. Note: The URI must start
	// with
	// "gs://".
	EntityTypeBatchUri string `json:"entityTypeBatchUri,omitempty"`

	// LanguageCode: Optional. The language of entity synonyms defined in
	// `entity_types`. If not
	// specified, the agent's default language is used.
	// [More than a
	// dozen
	// languages](https://dialogflow.com/docs/reference/language) are
	// supported.
	// Note: languages must be enabled in the agent, before they can be
	// used.
	LanguageCode string `json:"languageCode,omitempty"`

	// UpdateMask: Optional. The mask to control which fields get updated.
	UpdateMask string `json:"updateMask,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "EntityTypeBatchInline") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EntityTypeBatchInline") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1BatchUpdateEntityTypesRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1BatchUpdateEntityTypesRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1BatchUpdateEntityTypesResponse: The
// response message for EntityTypes.BatchUpdateEntityTypes.
type GoogleCloudDialogflowV2beta1BatchUpdateEntityTypesResponse struct {
	// EntityTypes: The collection of updated or created entity types.
	EntityTypes []*GoogleCloudDialogflowV2beta1EntityType `json:"entityTypes,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EntityTypes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EntityTypes") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1BatchUpdateEntityTypesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1BatchUpdateEntityTypesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1BatchUpdateIntentsRequest: The request
// message for Intents.BatchUpdateIntents.
type GoogleCloudDialogflowV2beta1BatchUpdateIntentsRequest struct {
	// IntentBatchInline: The collection of intents to update or create.
	IntentBatchInline *GoogleCloudDialogflowV2beta1IntentBatch `json:"intentBatchInline,omitempty"`

	// IntentBatchUri: The URI to a Google Cloud Storage file containing
	// intents to update or
	// create. The file format can either be a serialized proto (of
	// IntentBatch
	// type) or JSON object. Note: The URI must start with "gs://".
	IntentBatchUri string `json:"intentBatchUri,omitempty"`

	// IntentView: Optional. The resource view to apply to the returned
	// intent.
	//
	// Possible values:
	//   "INTENT_VIEW_UNSPECIFIED" - Training phrases field is not populated
	// in the response.
	//   "INTENT_VIEW_FULL" - All fields are populated.
	IntentView string `json:"intentView,omitempty"`

	// LanguageCode: Optional. The language of training phrases, parameters
	// and rich messages
	// defined in `intents`. If not specified, the agent's default language
	// is
	// used. [More than a
	// dozen
	// languages](https://dialogflow.com/docs/reference/language) are
	// supported.
	// Note: languages must be enabled in the agent, before they can be
	// used.
	LanguageCode string `json:"languageCode,omitempty"`

	// UpdateMask: Optional. The mask to control which fields get updated.
	UpdateMask string `json:"updateMask,omitempty"`

	// ForceSendFields is a list of field names (e.g. "IntentBatchInline")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "IntentBatchInline") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1BatchUpdateIntentsRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1BatchUpdateIntentsRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1BatchUpdateIntentsResponse: The response
// message for Intents.BatchUpdateIntents.
type GoogleCloudDialogflowV2beta1BatchUpdateIntentsResponse struct {
	// Intents: The collection of updated or created intents.
	Intents []*GoogleCloudDialogflowV2beta1Intent `json:"intents,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Intents") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Intents") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1BatchUpdateIntentsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1BatchUpdateIntentsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1CancelPhoneNumberOrderRequest: The
// request message for PhoneNumberOrders.CancelPhoneNumberOrder.
type GoogleCloudDialogflowV2beta1CancelPhoneNumberOrderRequest struct {
}

// GoogleCloudDialogflowV2beta1CompileSuggestionsRequest: The request
// message for
// HumanAgentAssistants.RequestCompileSuggestions.
type GoogleCloudDialogflowV2beta1CompileSuggestionsRequest struct {
	// Messages: Required. List of messages in a conversation in
	// chronological order.
	Messages []*GoogleCloudDialogflowV2beta1Message `json:"messages,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Messages") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Messages") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1CompileSuggestionsRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1CompileSuggestionsRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1CompileSuggestionsResponse: The response
// message for
// HumanAgentAssistants.RequestCompileSuggestions
type GoogleCloudDialogflowV2beta1CompileSuggestionsResponse struct {
	// Suggestions: Required.
	Suggestions []*GoogleCloudDialogflowV2beta1Suggestion `json:"suggestions,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Suggestions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Suggestions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1CompileSuggestionsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1CompileSuggestionsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1CompleteConversationRequest: The request
// message for Conversations.CompleteConversation.
type GoogleCloudDialogflowV2beta1CompleteConversationRequest struct {
}

// GoogleCloudDialogflowV2beta1Context: Represents a context.
type GoogleCloudDialogflowV2beta1Context struct {
	// LifespanCount: Optional. The number of conversational query requests
	// after which the
	// context expires. If set to `0` (the default) the context
	// expires
	// immediately. Contexts expire automatically after 10 minutes even if
	// there
	// are no matching queries.
	LifespanCount int64 `json:"lifespanCount,omitempty"`

	// Name: Required. The unique identifier of the context.
	// Format:
	// `projects/<Project ID>/agent/sessions/<Session ID>/contexts/<Context
	// ID>`,
	// or `projects/<Project ID>/agent/environments/<Environment
	// ID>/users/<User
	// ID>/sessions/<Session ID>/contexts/<Context ID>`. The `Context ID`
	// is
	// always converted to lowercase. If `Environment ID` is not specified,
	// we
	// assume default 'draft' environment. If `User ID` is not specified,
	// we
	// assume default '-' user.
	Name string `json:"name,omitempty"`

	// Parameters: Optional. The collection of parameters associated with
	// this context.
	// Refer to [this
	// doc](https://dialogflow.com/docs/actions-and-parameters) for
	// syntax.
	Parameters googleapi.RawMessage `json:"parameters,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "LifespanCount") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "LifespanCount") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1Context) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1Context
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1Conversation: Represents a
// conversation.
// A conversation is an interaction between an agent, including live
// agents
// and Dialogflow agents, and a support customer. Conversations
// can
// include phone calls and text-based chat sessions.
type GoogleCloudDialogflowV2beta1Conversation struct {
	// ConversationProfile: Required. The Conversation Profile to be used to
	// configure this
	// Conversation. This field cannot be updated.
	// Format: `projects/<Project ID>/conversationProfiles/<Conversation
	// Profile
	// ID>`.
	ConversationProfile string `json:"conversationProfile,omitempty"`

	// EndTime: Output only. The time the conversation was finished.
	EndTime string `json:"endTime,omitempty"`

	// LifecycleState: Output only. The current state of the Conversation.
	//
	// Possible values:
	//   "LIFECYCLE_STATE_UNSPECIFIED" - Unknown.
	//   "IN_PROGRESS" - Conversation is currently open for media analysis.
	//   "COMPLETED" - Conversation has been completed.
	LifecycleState string `json:"lifecycleState,omitempty"`

	// Name: The unique identifier of this conversation.
	// Required for all methods except `create` (`create` populates the
	// name
	// automatically).
	// Format: `projects/<Project ID>/conversations/<Conversation ID>`.
	Name string `json:"name,omitempty"`

	// PhoneNumber: Output only. Required if the conversation is to be
	// connected over
	// telephony.
	PhoneNumber *GoogleCloudDialogflowV2beta1ConversationPhoneNumber `json:"phoneNumber,omitempty"`

	// StartTime: Output only. The time the conversation was started.
	StartTime string `json:"startTime,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ConversationProfile")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ConversationProfile") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1Conversation) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1Conversation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1ConversationPhoneNumber: Represents a
// phone number for telephony integration. It allows for connecting
// a particular conversation over telephony.
type GoogleCloudDialogflowV2beta1ConversationPhoneNumber struct {
	// PhoneNumber: Output only. The phone number to connect to this
	// conversation.
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "PhoneNumber") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "PhoneNumber") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1ConversationPhoneNumber) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1ConversationPhoneNumber
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1ConversationProfile: Defines the services
// to connect to incoming Dialogflow conversations.
type GoogleCloudDialogflowV2beta1ConversationProfile struct {
	// AutomatedAgentConfig: Optional. Configuration for an automated agent
	// to use with this profile.
	AutomatedAgentConfig *GoogleCloudDialogflowV2beta1AutomatedAgentConfig `json:"automatedAgentConfig,omitempty"`

	// DisplayName: Required. Human readable name for this profile. Max
	// length 1024 bytes.
	DisplayName string `json:"displayName,omitempty"`

	// HumanAgentAssistantConfig: Optional. Configuration for agent
	// assistance to use with this profile.
	HumanAgentAssistantConfig *GoogleCloudDialogflowV2beta1HumanAgentAssistantConfig `json:"humanAgentAssistantConfig,omitempty"`

	// LoggingConfig: Optional. Configuration for logging conversation
	// lifecycle events.
	LoggingConfig *GoogleCloudDialogflowV2beta1LoggingConfig `json:"loggingConfig,omitempty"`

	// Name: Required for all methods except `create` (`create` populates
	// the name
	// automatically).
	// The unique identifier of this conversation profile.
	// Format: `projects/<Project ID>/conversationProfiles/<Conversation
	// Profile
	// ID>`.
	Name string `json:"name,omitempty"`

	// NotificationConfig: Optional. Configuration for publishing
	// conversation lifecycle events.
	NotificationConfig *GoogleCloudDialogflowV2beta1NotificationConfig `json:"notificationConfig,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g.
	// "AutomatedAgentConfig") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AutomatedAgentConfig") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1ConversationProfile) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1ConversationProfile
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1DetectIntentRequest: The request to
// detect user's intent.
type GoogleCloudDialogflowV2beta1DetectIntentRequest struct {
	// InputAudio: Optional. The natural language speech audio to be
	// processed. This field
	// should be populated iff `query_input` is set to an input audio
	// config.
	// A single request can contain up to 1 minute of speech audio data.
	InputAudio string `json:"inputAudio,omitempty"`

	// OutputAudioConfig: Optional. Instructs the speech synthesizer how to
	// generate the output
	// audio. If this field is not set and agent-level speech synthesizer is
	// not
	// configured, no output audio is generated.
	OutputAudioConfig *GoogleCloudDialogflowV2beta1OutputAudioConfig `json:"outputAudioConfig,omitempty"`

	// QueryInput: Required. The input specification. It can be set to:
	//
	// 1.  an audio config
	//     which instructs the speech recognizer how to process the speech
	// audio,
	//
	// 2.  a conversational query in the form of text, or
	//
	// 3.  an event that specifies which intent to trigger.
	QueryInput *GoogleCloudDialogflowV2beta1QueryInput `json:"queryInput,omitempty"`

	// QueryParams: Optional. The parameters of this query.
	QueryParams *GoogleCloudDialogflowV2beta1QueryParameters `json:"queryParams,omitempty"`

	// ForceSendFields is a list of field names (e.g. "InputAudio") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "InputAudio") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1DetectIntentRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1DetectIntentRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1DetectIntentResponse: The message
// returned from the DetectIntent method.
type GoogleCloudDialogflowV2beta1DetectIntentResponse struct {
	// AlternativeQueryResults: If Knowledge Connectors are enabled, there
	// could be more than one result
	// returned for a given query or event, and this field will contain
	// all
	// results except for the top one, which is captured in query_result.
	// The
	// alternative results are ordered by
	// decreasing
	// `QueryResult.intent_detection_confidence`. If Knowledge Connectors
	// are
	// disabled, this field will be empty until multiple responses for
	// regular
	// intents are supported, at which point those additional results will
	// be
	// surfaced here.
	AlternativeQueryResults []*GoogleCloudDialogflowV2beta1QueryResult `json:"alternativeQueryResults,omitempty"`

	// OutputAudio: The audio data bytes encoded as specified in the
	// request.
	OutputAudio string `json:"outputAudio,omitempty"`

	// OutputAudioConfig: Instructs the speech synthesizer how to generate
	// the output audio. This
	// field is populated from the agent-level speech synthesizer
	// configuration,
	// if enabled.
	OutputAudioConfig *GoogleCloudDialogflowV2beta1OutputAudioConfig `json:"outputAudioConfig,omitempty"`

	// QueryResult: The selected results of the conversational query or
	// event processing.
	// See `alternative_query_results` for additional potential results.
	QueryResult *GoogleCloudDialogflowV2beta1QueryResult `json:"queryResult,omitempty"`

	// ResponseId: The unique identifier of the response. It can be used
	// to
	// locate a response in the training example set or for reporting
	// issues.
	ResponseId string `json:"responseId,omitempty"`

	// WebhookStatus: Specifies the status of the webhook request.
	// `webhook_status`
	// is never populated in webhook requests.
	WebhookStatus *GoogleRpcStatus `json:"webhookStatus,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g.
	// "AlternativeQueryResults") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AlternativeQueryResults")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1DetectIntentResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1DetectIntentResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1Document: A document resource.
//
// Note: resource `projects.agent.knowledgeBases.documents` is
// deprecated,
// please use `projects.knowledgeBases.documents` instead.
type GoogleCloudDialogflowV2beta1Document struct {
	// Content: The raw content of the document. This field is only
	// permitted for
	// EXTRACTIVE_QA and FAQ knowledge types.
	// Note: This field is in the process of being deprecated, please
	// use
	// raw_content instead.
	Content string `json:"content,omitempty"`

	// ContentUri: The URI where the file content is located.
	//
	// For documents stored in Google Cloud Storage, these URIs must
	// have
	// the form `gs://<bucket-name>/<object-name>`.
	//
	// NOTE: External URLs must correspond to public webpages, i.e., they
	// must
	// be indexed by Google Search. In particular, URLs for showing
	// documents in
	// Google Cloud Storage (i.e. the URL in your browser) are not
	// supported.
	// Instead use the `gs://` format URI described above.
	ContentUri string `json:"contentUri,omitempty"`

	// DisplayName: Required. The display name of the document. The name
	// must be 1024 bytes or
	// less; otherwise, the creation request fails.
	DisplayName string `json:"displayName,omitempty"`

	// KnowledgeTypes: Required. The knowledge type of document content.
	//
	// Possible values:
	//   "KNOWLEDGE_TYPE_UNSPECIFIED" - The type is unspecified or
	// arbitrary.
	//   "FAQ" - The document content contains question and answer pairs as
	// either HTML or
	// CSV. Typical FAQ HTML formats are parsed accurately, but unusual
	// formats
	// may fail to be parsed.
	//
	// CSV must have questions in the first column and answers in the
	// second,
	// with no header. Because of this explicit format, they are always
	// parsed
	// accurately.
	//   "EXTRACTIVE_QA" - Documents for which unstructured text is
	// extracted and used for
	// question answering.
	KnowledgeTypes []string `json:"knowledgeTypes,omitempty"`

	// MimeType: Required. The MIME type of this document.
	MimeType string `json:"mimeType,omitempty"`

	// Name: The document resource name.
	// The name must be empty when creating a document.
	// Format: `projects/<Project ID>/knowledgeBases/<Knowledge
	// Base
	// ID>/documents/<Document ID>`.
	Name string `json:"name,omitempty"`

	// RawContent: The raw content of the document. This field is only
	// permitted for
	// EXTRACTIVE_QA and FAQ knowledge types.
	RawContent string `json:"rawContent,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Content") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Content") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1Document) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1Document
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1EntityType: Represents an entity
// type.
// Entity types serve as a tool for extracting parameter values from
// natural
// language queries.
type GoogleCloudDialogflowV2beta1EntityType struct {
	// AutoExpansionMode: Optional. Indicates whether the entity type can be
	// automatically
	// expanded.
	//
	// Possible values:
	//   "AUTO_EXPANSION_MODE_UNSPECIFIED" - Auto expansion disabled for the
	// entity.
	//   "AUTO_EXPANSION_MODE_DEFAULT" - Allows an agent to recognize values
	// that have not been explicitly
	// listed in the entity.
	AutoExpansionMode string `json:"autoExpansionMode,omitempty"`

	// DisplayName: Required. The name of the entity type.
	DisplayName string `json:"displayName,omitempty"`

	// Entities: Optional. The collection of entities associated with the
	// entity type.
	Entities []*GoogleCloudDialogflowV2beta1EntityTypeEntity `json:"entities,omitempty"`

	// Kind: Required. Indicates the kind of entity type.
	//
	// Possible values:
	//   "KIND_UNSPECIFIED" - Not specified. This value should be never
	// used.
	//   "KIND_MAP" - Map entity types allow mapping of a group of synonyms
	// to a canonical
	// value.
	//   "KIND_LIST" - List entity types contain a set of entries that do
	// not map to canonical
	// values. However, list entity types can contain references to other
	// entity
	// types (with or without aliases).
	Kind string `json:"kind,omitempty"`

	// Name: Required for all methods except `create` (`create` populates
	// the name
	// automatically.
	// The unique identifier of the entity type. Format:
	// `projects/<Project ID>/agent/entityTypes/<Entity Type ID>`.
	Name string `json:"name,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AutoExpansionMode")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AutoExpansionMode") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1EntityType) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1EntityType
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1EntityTypeBatch: This message is a
// wrapper around a collection of entity types.
type GoogleCloudDialogflowV2beta1EntityTypeBatch struct {
	// EntityTypes: A collection of entity types.
	EntityTypes []*GoogleCloudDialogflowV2beta1EntityType `json:"entityTypes,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EntityTypes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EntityTypes") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1EntityTypeBatch) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1EntityTypeBatch
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1EntityTypeEntity: Optional. Represents an
// entity.
type GoogleCloudDialogflowV2beta1EntityTypeEntity struct {
	// Synonyms: Required. A collection of synonyms. For `KIND_LIST` entity
	// types this
	// must contain exactly one synonym equal to `value`.
	Synonyms []string `json:"synonyms,omitempty"`

	// Value: Required.
	// For `KIND_MAP` entity types:
	//   A canonical name to be used in place of synonyms.
	// For `KIND_LIST` entity types:
	//   A string that can contain references to other entity types (with
	// or
	//   without aliases).
	Value string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Synonyms") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Synonyms") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1EntityTypeEntity) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1EntityTypeEntity
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1EventInput: Events allow for matching
// intents by event name instead of the natural
// language input. For instance, input `<event: { name:
// “welcome_event”,
// parameters: { name: “Sam” } }>` can trigger a personalized
// welcome response.
// The parameter `name` may be used by the agent in the
// response:
// `“Hello #welcome_event.name! What can I do for you today?”`.
type GoogleCloudDialogflowV2beta1EventInput struct {
	// LanguageCode: Required. The language of this query. See
	// [Language
	// Support](https://dialogflow.com/docs/languages) for a list of
	// the
	// currently supported language codes. Note that queries in the same
	// session
	// do not necessarily need to specify the same language.
	LanguageCode string `json:"languageCode,omitempty"`

	// Name: Required. The unique identifier of the event.
	Name string `json:"name,omitempty"`

	// Parameters: Optional. The collection of parameters associated with
	// the event.
	Parameters googleapi.RawMessage `json:"parameters,omitempty"`

	// ForceSendFields is a list of field names (e.g. "LanguageCode") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "LanguageCode") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1EventInput) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1EventInput
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1ExportAgentRequest: The request message
// for Agents.ExportAgent.
type GoogleCloudDialogflowV2beta1ExportAgentRequest struct {
	// AgentUri: Optional. The
	// [Google Cloud Storage](https://cloud.google.com/storage/docs/)
	// URI to export the agent to.
	// The format of this URI must be `gs://<bucket-name>/<object-name>`.
	// If left unspecified, the serialized agent is returned inline.
	AgentUri string `json:"agentUri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AgentUri") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AgentUri") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1ExportAgentRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1ExportAgentRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1ExportAgentResponse: The response message
// for Agents.ExportAgent.
type GoogleCloudDialogflowV2beta1ExportAgentResponse struct {
	// AgentContent: The exported agent.
	//
	// Example for how to export an agent to a zip file via a command
	// line:
	// <pre>curl \
	//
	// 'https://dialogflow.googleapis.com/v2beta1/projects/&lt;project_name&g
	// t;/agent:export'\
	//   -X POST \
	//   -H 'Authorization: Bearer '$(gcloud auth application-default
	//   print-access-token) \
	//   -H 'Accept: application/json' \
	//   -H 'Content-Type: application/json' \
	//   --compressed \
	//   --data-binary '{}' \
	// | grep agentContent | sed -e 's/.*"agentContent": "\([^"]*\)".*/\1/'
	// \
	// | base64 --decode > &lt;agent zip file&gt;</pre>
	AgentContent string `json:"agentContent,omitempty"`

	// AgentUri: The URI to a file containing the exported agent. This field
	// is populated
	// only if `agent_uri` is specified in `ExportAgentRequest`.
	AgentUri string `json:"agentUri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AgentContent") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AgentContent") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1ExportAgentResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1ExportAgentResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1FaqAnswersConfig: Defines FAQ responses
// that a human agent assistant can provide.
type GoogleCloudDialogflowV2beta1FaqAnswersConfig struct {
	// KnowledgeBaseName: Required. Settings for knowledge base,
	// Format:
	// `projects/<Project ID>/knowledgeBases/<Knowledge Base ID>`.
	KnowledgeBaseName string `json:"knowledgeBaseName,omitempty"`

	// MaxResults: Optional. Maximum number of results to return. If unset,
	// defaults to 10.
	MaxResults int64 `json:"maxResults,omitempty"`

	// ForceSendFields is a list of field names (e.g. "KnowledgeBaseName")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "KnowledgeBaseName") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1FaqAnswersConfig) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1FaqAnswersConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1HumanAgentAssistant: Represents a human
// agent assistant that provides suggestions to help
// human agents to resolve customer issues. This defines the types of
// content
// that the human agent assistant can present to a human agent.
type GoogleCloudDialogflowV2beta1HumanAgentAssistant struct {
	// ArticleSuggestionConfig: Optional. Settings for article suggestion.
	ArticleSuggestionConfig *GoogleCloudDialogflowV2beta1ArticleSuggestionConfig `json:"articleSuggestionConfig,omitempty"`

	// FaqAnswersConfig: Optional. Settings for knowledge service.
	FaqAnswersConfig *GoogleCloudDialogflowV2beta1FaqAnswersConfig `json:"faqAnswersConfig,omitempty"`

	// Name: Required for all methods except `create` (`create` populates
	// the name
	// automatically).
	// The unique identifier of human agent assistant.
	// Format: `projects/<Project ID>/humanAgentAssistants/<Human Agent
	// Assistant
	// ID>`.
	Name string `json:"name,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g.
	// "ArticleSuggestionConfig") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ArticleSuggestionConfig")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1HumanAgentAssistant) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1HumanAgentAssistant
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1HumanAgentAssistantConfig: Defines the
// Human Agent Assistant to connect to a conversation.
type GoogleCloudDialogflowV2beta1HumanAgentAssistantConfig struct {
	// Name: Required. ID of the agent assistant to use.
	// Format: `projects/<Project ID>/humanAgentAssistants/<Human Agent
	// Assistant
	// ID>`.
	Name string `json:"name,omitempty"`

	// NotificationConfig: Optional. Pub/Sub topic on which to publish new
	// agent assistant events.
	NotificationConfig *GoogleCloudDialogflowV2beta1NotificationConfig `json:"notificationConfig,omitempty"`

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

func (s *GoogleCloudDialogflowV2beta1HumanAgentAssistantConfig) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1HumanAgentAssistantConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1ImportAgentRequest: The request message
// for Agents.ImportAgent.
type GoogleCloudDialogflowV2beta1ImportAgentRequest struct {
	// AgentContent: The agent to import.
	//
	// Example for how to import an agent via the command line:
	// <pre>curl \
	//
	// 'https://dialogflow.googleapis.com/v2beta1/projects/&lt;project_name&g
	// t;/agent:import\
	//    -X POST \
	//    -H 'Authorization: Bearer '$(gcloud auth application-default
	//    print-access-token) \
	//    -H 'Accept: application/json' \
	//    -H 'Content-Type: application/json' \
	//    --compressed \
	//    --data-binary "{
	//       'agentContent': '$(cat &lt;agent zip file&gt; | base64 -w 0)'
	//    }"</pre>
	AgentContent string `json:"agentContent,omitempty"`

	// AgentUri: The URI to a Google Cloud Storage file containing the agent
	// to import.
	// Note: The URI must start with "gs://".
	AgentUri string `json:"agentUri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AgentContent") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AgentContent") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1ImportAgentRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1ImportAgentRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1InputAudio: Represents the natural
// language speech audio to be processed.
type GoogleCloudDialogflowV2beta1InputAudio struct {
	// Audio: Required. The natural language speech audio to be processed.
	// A single request can contain up to 1 minute of speech audio data.
	// The transcribed text cannot contain more than 256 bytes.
	Audio string `json:"audio,omitempty"`

	// Config: Required. Instructs the speech recognizer how to process the
	// speech audio.
	Config *GoogleCloudDialogflowV2beta1InputAudioConfig `json:"config,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Audio") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Audio") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1InputAudio) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1InputAudio
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1InputAudioConfig: Instructs the speech
// recognizer how to process the audio content.
type GoogleCloudDialogflowV2beta1InputAudioConfig struct {
	// AudioEncoding: Required. Audio encoding of the audio content to
	// process.
	//
	// Possible values:
	//   "AUDIO_ENCODING_UNSPECIFIED" - Not specified.
	//   "AUDIO_ENCODING_LINEAR_16" - Uncompressed 16-bit signed
	// little-endian samples (Linear PCM).
	//   "AUDIO_ENCODING_FLAC" -
	// [`FLAC`](https://xiph.org/flac/documentation.html) (Free Lossless
	// Audio
	// Codec) is the recommended encoding because it is lossless
	// (therefore
	// recognition is not compromised) and requires only about half
	// the
	// bandwidth of `LINEAR16`. `FLAC` stream encoding supports 16-bit
	// and
	// 24-bit samples, however, not all fields in `STREAMINFO` are
	// supported.
	//   "AUDIO_ENCODING_MULAW" - 8-bit samples that compand 14-bit audio
	// samples using G.711 PCMU/mu-law.
	//   "AUDIO_ENCODING_AMR" - Adaptive Multi-Rate Narrowband codec.
	// `sample_rate_hertz` must be 8000.
	//   "AUDIO_ENCODING_AMR_WB" - Adaptive Multi-Rate Wideband codec.
	// `sample_rate_hertz` must be 16000.
	//   "AUDIO_ENCODING_OGG_OPUS" - Opus encoded audio frames in Ogg
	// container
	// ([OggOpus](https://wiki.xiph.org/OggOpus)).
	// `sample_rate_her
	// tz` must be 16000.
	//   "AUDIO_ENCODING_SPEEX_WITH_HEADER_BYTE" - Although the use of lossy
	// encodings is not recommended, if a very low
	// bitrate encoding is required, `OGG_OPUS` is highly preferred
	// over
	// Speex encoding. The [Speex](https://speex.org/) encoding supported
	// by
	// Dialogflow API has a header byte in each block, as in MIME
	// type
	// `audio/x-speex-with-header-byte`.
	// It is a variant of the RTP Speex encoding defined in
	// [RFC 5574](https://tools.ietf.org/html/rfc5574).
	// The stream is a sequence of blocks, one block per RTP packet. Each
	// block
	// starts with a byte containing the length of the block, in bytes,
	// followed
	// by one or more frames of Speex data, padded to an integral number
	// of
	// bytes (octets) as specified in RFC 5574. In other words, each RTP
	// header
	// is replaced with a single byte containing the block length. Only
	// Speex
	// wideband is supported. `sample_rate_hertz` must be 16000.
	AudioEncoding string `json:"audioEncoding,omitempty"`

	// LanguageCode: Required. The language of the supplied audio.
	// Dialogflow does not do
	// translations. See
	// [Language
	// Support](https://dialogflow.com/docs/languages) for a list of
	// the
	// currently supported language codes. Note that queries in the same
	// session
	// do not necessarily need to specify the same language.
	LanguageCode string `json:"languageCode,omitempty"`

	// Model: Optional. Which Speech model to select for the given request.
	// Select the
	// model best suited to your domain to get best results. If a model is
	// not
	// explicitly specified, then we auto-select a model based on the
	// parameters
	// in the InputAudioConfig.
	// If enhanced speech model is enabled for the agent and an
	// enhanced
	// version of the specified model for the language does not exist, then
	// the
	// speech is recognized using the standard version of the specified
	// model.
	// Refer to
	// [Cloud Speech
	// API
	// documentation](https://cloud.google.com/speech-to-text/docs/basics
	// #select-model)
	// for more details.
	Model string `json:"model,omitempty"`

	// PhraseHints: Optional. The collection of phrase hints which are used
	// to boost accuracy
	// of speech recognition.
	// Refer to
	// [Cloud Speech
	// API
	// documentation](https://cloud.google.com/speech-to-text/docs/basics
	// #phrase-hints)
	// for more details.
	PhraseHints []string `json:"phraseHints,omitempty"`

	// SampleRateHertz: Required. Sample rate (in Hertz) of the audio
	// content sent in the query.
	// Refer to
	// [Cloud Speech
	// API
	// documentation](https://cloud.google.com/speech-to-text/docs/basics
	// ) for
	// more details.
	SampleRateHertz int64 `json:"sampleRateHertz,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AudioEncoding") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AudioEncoding") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1InputAudioConfig) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1InputAudioConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1InputText: Represents the natural
// language text to be processed.
type GoogleCloudDialogflowV2beta1InputText struct {
	// LanguageCode: Required. The language of this conversational query.
	// See [Language
	// Support](https://dialogflow.com/docs/languages) for a list of
	// the
	// currently supported language codes.
	LanguageCode string `json:"languageCode,omitempty"`

	// Text: Required. The UTF-8 encoded natural language text to be
	// processed.
	// Text length must not exceed 256 bytes.
	Text string `json:"text,omitempty"`

	// ForceSendFields is a list of field names (e.g. "LanguageCode") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "LanguageCode") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1InputText) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1InputText
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1InputTextConfig: Defines the language
// used in the input text.
type GoogleCloudDialogflowV2beta1InputTextConfig struct {
	// LanguageCode: Required. The language of this conversational query.
	// See [Language
	// Support](https://dialogflow.com/docs/languages) for a list of
	// the
	// currently supported language codes.
	LanguageCode string `json:"languageCode,omitempty"`

	// ForceSendFields is a list of field names (e.g. "LanguageCode") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "LanguageCode") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1InputTextConfig) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1InputTextConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1Intent: Represents an intent.
// Intents convert a number of user expressions or patterns into an
// action. An
// action is an extraction of a user command or sentence semantics.
type GoogleCloudDialogflowV2beta1Intent struct {
	// Action: Optional. The name of the action associated with the
	// intent.
	// Note: The action name must not contain whitespaces.
	Action string `json:"action,omitempty"`

	// DefaultResponsePlatforms: Optional. The list of platforms for which
	// the first response will be
	// taken from among the messages assigned to the DEFAULT_PLATFORM.
	//
	// Possible values:
	//   "PLATFORM_UNSPECIFIED" - Not specified.
	//   "FACEBOOK" - Facebook.
	//   "SLACK" - Slack.
	//   "TELEGRAM" - Telegram.
	//   "KIK" - Kik.
	//   "SKYPE" - Skype.
	//   "LINE" - Line.
	//   "VIBER" - Viber.
	//   "ACTIONS_ON_GOOGLE" - Actions on Google.
	// When using Actions on Google, you can choose one of the
	// specific
	// Intent.Message types that mention support for Actions on Google,
	// or you can use the advanced Intent.Message.payload field.
	// The payload field provides access to AoG features not available in
	// the
	// specific message types.
	// If using the Intent.Message.payload field, it should have a
	// structure
	// similar to the JSON message shown here. For more information,
	// see
	// [Actions on Google
	// Webhook
	// Format](https://developers.google.com/actions/dialogflow/webho
	// ok)
	// <pre>{
	//   "expectUserResponse": true,
	//   "isSsml": false,
	//   "noInputPrompts": [],
	//   "richResponse": {
	//     "items": [
	//       {
	//         "simpleResponse": {
	//           "displayText": "hi",
	//           "textToSpeech": "hello"
	//         }
	//       }
	//     ],
	//     "suggestions": [
	//       {
	//         "title": "Say this"
	//       },
	//       {
	//         "title": "or this"
	//       }
	//     ]
	//   },
	//   "systemIntent": {
	//     "data": {
	//       "@type":
	// "type.googleapis.com/google.actions.v2.OptionValueSpec",
	//       "listSelect": {
	//         "items": [
	//           {
	//             "optionInfo": {
	//               "key": "key1",
	//               "synonyms": [
	//                 "key one"
	//               ]
	//             },
	//             "title": "must not be empty, but unique"
	//           },
	//           {
	//             "optionInfo": {
	//               "key": "key2",
	//               "synonyms": [
	//                 "key two"
	//               ]
	//             },
	//             "title": "must not be empty, but unique"
	//           }
	//         ]
	//       }
	//     },
	//     "intent": "actions.intent.OPTION"
	//   }
	// }</pre>
	//   "TELEPHONY" - Telephony Gateway.
	DefaultResponsePlatforms []string `json:"defaultResponsePlatforms,omitempty"`

	// DisplayName: Required. The name of this intent.
	DisplayName string `json:"displayName,omitempty"`

	// EndInteraction: Optional. Indicates that this intent ends an
	// interaction. Some integrations
	// (e.g., Actions on Google or Dialogflow phone gateway) use this
	// information
	// to close interaction with an end user. Default is false.
	EndInteraction bool `json:"endInteraction,omitempty"`

	// Events: Optional. The collection of event names that trigger the
	// intent.
	// If the collection of input contexts is not empty, all of the contexts
	// must
	// be present in the active user session for an event to trigger this
	// intent.
	Events []string `json:"events,omitempty"`

	// FollowupIntentInfo: Read-only. Information about all followup intents
	// that have this intent as
	// a direct or indirect parent. We populate this field only in the
	// output.
	FollowupIntentInfo []*GoogleCloudDialogflowV2beta1IntentFollowupIntentInfo `json:"followupIntentInfo,omitempty"`

	// InputContextNames: Optional. The list of context names required for
	// this intent to be
	// triggered.
	// Format: `projects/<Project ID>/agent/sessions/-/contexts/<Context
	// ID>`.
	InputContextNames []string `json:"inputContextNames,omitempty"`

	// IsFallback: Optional. Indicates whether this is a fallback intent.
	IsFallback bool `json:"isFallback,omitempty"`

	// Messages: Optional. The collection of rich messages corresponding to
	// the
	// `Response` field in the Dialogflow console.
	Messages []*GoogleCloudDialogflowV2beta1IntentMessage `json:"messages,omitempty"`

	// MlDisabled: Optional. Indicates whether Machine Learning is disabled
	// for the intent.
	// Note: If `ml_disabled` setting is set to true, then this intent is
	// not
	// taken into account during inference in `ML ONLY` match mode.
	// Also,
	// auto-markup in the UI is turned off.
	MlDisabled bool `json:"mlDisabled,omitempty"`

	// MlEnabled: Optional. Indicates whether Machine Learning is enabled
	// for the intent.
	// Note: If `ml_enabled` setting is set to false, then this intent is
	// not
	// taken into account during inference in `ML ONLY` match mode.
	// Also,
	// auto-markup in the UI is turned off.
	// DEPRECATED! Please use `ml_disabled` field instead.
	// NOTE: If both `ml_enabled` and `ml_disabled` are either not set or
	// false,
	// then the default value is determined as follows:
	// - Before April 15th, 2018 the default is:
	//   ml_enabled = false / ml_disabled = true.
	// - After April 15th, 2018 the default is:
	//   ml_enabled = true / ml_disabled = false.
	MlEnabled bool `json:"mlEnabled,omitempty"`

	// Name: Required for all methods except `create` (`create` populates
	// the name
	// automatically.
	// The unique identifier of this intent.
	// Format: `projects/<Project ID>/agent/intents/<Intent ID>`.
	Name string `json:"name,omitempty"`

	// OutputContexts: Optional. The collection of contexts that are
	// activated when the intent
	// is matched. Context messages in this collection should not set
	// the
	// parameters field. Setting the `lifespan_count` to 0 will reset the
	// context
	// when the intent is matched.
	// Format: `projects/<Project ID>/agent/sessions/-/contexts/<Context
	// ID>`.
	OutputContexts []*GoogleCloudDialogflowV2beta1Context `json:"outputContexts,omitempty"`

	// Parameters: Optional. The collection of parameters associated with
	// the intent.
	Parameters []*GoogleCloudDialogflowV2beta1IntentParameter `json:"parameters,omitempty"`

	// ParentFollowupIntentName: Read-only after creation. The unique
	// identifier of the parent intent in the
	// chain of followup intents. You can set this field when creating an
	// intent,
	// for example with CreateIntent or BatchUpdateIntents, in order to
	// make this intent a followup intent.
	//
	// It identifies the parent followup intent.
	// Format: `projects/<Project ID>/agent/intents/<Intent ID>`.
	ParentFollowupIntentName string `json:"parentFollowupIntentName,omitempty"`

	// Priority: Optional. The priority of this intent. Higher numbers
	// represent higher
	// priorities. Zero or negative numbers mean that the intent is
	// disabled.
	Priority int64 `json:"priority,omitempty"`

	// ResetContexts: Optional. Indicates whether to delete all contexts in
	// the current
	// session when this intent is matched.
	ResetContexts bool `json:"resetContexts,omitempty"`

	// RootFollowupIntentName: Read-only. The unique identifier of the root
	// intent in the chain of
	// followup intents. It identifies the correct followup intents chain
	// for
	// this intent. We populate this field only in the output.
	//
	// Format: `projects/<Project ID>/agent/intents/<Intent ID>`.
	RootFollowupIntentName string `json:"rootFollowupIntentName,omitempty"`

	// TrainingPhrases: Optional. The collection of examples/templates that
	// the agent is
	// trained on.
	TrainingPhrases []*GoogleCloudDialogflowV2beta1IntentTrainingPhrase `json:"trainingPhrases,omitempty"`

	// WebhookState: Optional. Indicates whether webhooks are enabled for
	// the intent.
	//
	// Possible values:
	//   "WEBHOOK_STATE_UNSPECIFIED" - Webhook is disabled in the agent and
	// in the intent.
	//   "WEBHOOK_STATE_ENABLED" - Webhook is enabled in the agent and in
	// the intent.
	//   "WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING" - Webhook is enabled in
	// the agent and in the intent. Also, each slot
	// filling prompt is forwarded to the webhook.
	WebhookState string `json:"webhookState,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Action") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Action") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1Intent) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1Intent
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1IntentBatch: This message is a wrapper
// around a collection of intents.
type GoogleCloudDialogflowV2beta1IntentBatch struct {
	// Intents: A collection of intents.
	Intents []*GoogleCloudDialogflowV2beta1Intent `json:"intents,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Intents") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Intents") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1IntentBatch) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1IntentBatch
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1IntentFollowupIntentInfo: Represents a
// single followup intent in the chain.
type GoogleCloudDialogflowV2beta1IntentFollowupIntentInfo struct {
	// FollowupIntentName: The unique identifier of the followup
	// intent.
	// Format: `projects/<Project ID>/agent/intents/<Intent ID>`.
	FollowupIntentName string `json:"followupIntentName,omitempty"`

	// ParentFollowupIntentName: The unique identifier of the followup
	// intent's parent.
	// Format: `projects/<Project ID>/agent/intents/<Intent ID>`.
	ParentFollowupIntentName string `json:"parentFollowupIntentName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FollowupIntentName")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FollowupIntentName") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1IntentFollowupIntentInfo) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1IntentFollowupIntentInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1IntentMessage: Corresponds to the
// `Response` field in the Dialogflow console.
type GoogleCloudDialogflowV2beta1IntentMessage struct {
	// BasicCard: Displays a basic card for Actions on Google.
	BasicCard *GoogleCloudDialogflowV2beta1IntentMessageBasicCard `json:"basicCard,omitempty"`

	// Card: Displays a card.
	Card *GoogleCloudDialogflowV2beta1IntentMessageCard `json:"card,omitempty"`

	// CarouselSelect: Displays a carousel card for Actions on Google.
	CarouselSelect *GoogleCloudDialogflowV2beta1IntentMessageCarouselSelect `json:"carouselSelect,omitempty"`

	// Image: Displays an image.
	Image *GoogleCloudDialogflowV2beta1IntentMessageImage `json:"image,omitempty"`

	// LinkOutSuggestion: Displays a link out suggestion chip for Actions on
	// Google.
	LinkOutSuggestion *GoogleCloudDialogflowV2beta1IntentMessageLinkOutSuggestion `json:"linkOutSuggestion,omitempty"`

	// ListSelect: Displays a list card for Actions on Google.
	ListSelect *GoogleCloudDialogflowV2beta1IntentMessageListSelect `json:"listSelect,omitempty"`

	// Payload: Returns a response containing a custom, platform-specific
	// payload.
	// See the Intent.Message.Platform type for a description of
	// the
	// structure that may be required for your platform.
	Payload googleapi.RawMessage `json:"payload,omitempty"`

	// Platform: Optional. The platform that this message is intended for.
	//
	// Possible values:
	//   "PLATFORM_UNSPECIFIED" - Not specified.
	//   "FACEBOOK" - Facebook.
	//   "SLACK" - Slack.
	//   "TELEGRAM" - Telegram.
	//   "KIK" - Kik.
	//   "SKYPE" - Skype.
	//   "LINE" - Line.
	//   "VIBER" - Viber.
	//   "ACTIONS_ON_GOOGLE" - Actions on Google.
	// When using Actions on Google, you can choose one of the
	// specific
	// Intent.Message types that mention support for Actions on Google,
	// or you can use the advanced Intent.Message.payload field.
	// The payload field provides access to AoG features not available in
	// the
	// specific message types.
	// If using the Intent.Message.payload field, it should have a
	// structure
	// similar to the JSON message shown here. For more information,
	// see
	// [Actions on Google
	// Webhook
	// Format](https://developers.google.com/actions/dialogflow/webho
	// ok)
	// <pre>{
	//   "expectUserResponse": true,
	//   "isSsml": false,
	//   "noInputPrompts": [],
	//   "richResponse": {
	//     "items": [
	//       {
	//         "simpleResponse": {
	//           "displayText": "hi",
	//           "textToSpeech": "hello"
	//         }
	//       }
	//     ],
	//     "suggestions": [
	//       {
	//         "title": "Say this"
	//       },
	//       {
	//         "title": "or this"
	//       }
	//     ]
	//   },
	//   "systemIntent": {
	//     "data": {
	//       "@type":
	// "type.googleapis.com/google.actions.v2.OptionValueSpec",
	//       "listSelect": {
	//         "items": [
	//           {
	//             "optionInfo": {
	//               "key": "key1",
	//               "synonyms": [
	//                 "key one"
	//               ]
	//             },
	//             "title": "must not be empty, but unique"
	//           },
	//           {
	//             "optionInfo": {
	//               "key": "key2",
	//               "synonyms": [
	//                 "key two"
	//               ]
	//             },
	//             "title": "must not be empty, but unique"
	//           }
	//         ]
	//       }
	//     },
	//     "intent": "actions.intent.OPTION"
	//   }
	// }</pre>
	//   "TELEPHONY" - Telephony Gateway.
	Platform string `json:"platform,omitempty"`

	// QuickReplies: Displays quick replies.
	QuickReplies *GoogleCloudDialogflowV2beta1IntentMessageQuickReplies `json:"quickReplies,omitempty"`

	// SimpleResponses: Returns a voice or text-only response for Actions on
	// Google.
	SimpleResponses *GoogleCloudDialogflowV2beta1IntentMessageSimpleResponses `json:"simpleResponses,omitempty"`

	// Suggestions: Displays suggestion chips for Actions on Google.
	Suggestions *GoogleCloudDialogflowV2beta1IntentMessageSuggestions `json:"suggestions,omitempty"`

	// TelephonyPlayAudio: Plays audio from a file in Telephony Gateway.
	TelephonyPlayAudio *GoogleCloudDialogflowV2beta1IntentMessageTelephonyPlayAudio `json:"telephonyPlayAudio,omitempty"`

	// TelephonySynthesizeSpeech: Synthesizes speech in Telephony Gateway.
	TelephonySynthesizeSpeech *GoogleCloudDialogflowV2beta1IntentMessageTelephonySynthesizeSpeech `json:"telephonySynthesizeSpeech,omitempty"`

	// TelephonyTransferCall: Transfers the call in Telephony Gateway.
	TelephonyTransferCall *GoogleCloudDialogflowV2beta1IntentMessageTelephonyTransferCall `json:"telephonyTransferCall,omitempty"`

	// Text: Returns a text response.
	Text *GoogleCloudDialogflowV2beta1IntentMessageText `json:"text,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BasicCard") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BasicCard") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1IntentMessage) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1IntentMessage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1IntentMessageBasicCard: The basic card
// message. Useful for displaying information.
type GoogleCloudDialogflowV2beta1IntentMessageBasicCard struct {
	// Buttons: Optional. The collection of card buttons.
	Buttons []*GoogleCloudDialogflowV2beta1IntentMessageBasicCardButton `json:"buttons,omitempty"`

	// FormattedText: Required, unless image is present. The body text of
	// the card.
	FormattedText string `json:"formattedText,omitempty"`

	// Image: Optional. The image for the card.
	Image *GoogleCloudDialogflowV2beta1IntentMessageImage `json:"image,omitempty"`

	// Subtitle: Optional. The subtitle of the card.
	Subtitle string `json:"subtitle,omitempty"`

	// Title: Optional. The title of the card.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Buttons") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Buttons") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1IntentMessageBasicCard) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1IntentMessageBasicCard
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1IntentMessageBasicCardButton: The button
// object that appears at the bottom of a card.
type GoogleCloudDialogflowV2beta1IntentMessageBasicCardButton struct {
	// OpenUriAction: Required. Action to take when a user taps on the
	// button.
	OpenUriAction *GoogleCloudDialogflowV2beta1IntentMessageBasicCardButtonOpenUriAction `json:"openUriAction,omitempty"`

	// Title: Required. The title of the button.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "OpenUriAction") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "OpenUriAction") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1IntentMessageBasicCardButton) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1IntentMessageBasicCardButton
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1IntentMessageBasicCardButtonOpenUriAction:
//  Opens the given URI.
type GoogleCloudDialogflowV2beta1IntentMessageBasicCardButtonOpenUriAction struct {
	// Uri: Required. The HTTP or HTTPS scheme URI.
	Uri string `json:"uri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Uri") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Uri") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1IntentMessageBasicCardButtonOpenUriAction) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1IntentMessageBasicCardButtonOpenUriAction
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1IntentMessageCard: The card response
// message.
type GoogleCloudDialogflowV2beta1IntentMessageCard struct {
	// Buttons: Optional. The collection of card buttons.
	Buttons []*GoogleCloudDialogflowV2beta1IntentMessageCardButton `json:"buttons,omitempty"`

	// ImageUri: Optional. The public URI to an image file for the card.
	ImageUri string `json:"imageUri,omitempty"`

	// Subtitle: Optional. The subtitle of the card.
	Subtitle string `json:"subtitle,omitempty"`

	// Title: Optional. The title of the card.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Buttons") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Buttons") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1IntentMessageCard) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1IntentMessageCard
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1IntentMessageCardButton: Optional.
// Contains information about a button.
type GoogleCloudDialogflowV2beta1IntentMessageCardButton struct {
	// Postback: Optional. The text to send back to the Dialogflow API or a
	// URI to
	// open.
	Postback string `json:"postback,omitempty"`

	// Text: Optional. The text to show on the button.
	Text string `json:"text,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Postback") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Postback") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1IntentMessageCardButton) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1IntentMessageCardButton
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1IntentMessageCarouselSelect: The card for
// presenting a carousel of options to select from.
type GoogleCloudDialogflowV2beta1IntentMessageCarouselSelect struct {
	// Items: Required. Carousel items.
	Items []*GoogleCloudDialogflowV2beta1IntentMessageCarouselSelectItem `json:"items,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1IntentMessageCarouselSelect) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1IntentMessageCarouselSelect
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1IntentMessageCarouselSelectItem: An item
// in the carousel.
type GoogleCloudDialogflowV2beta1IntentMessageCarouselSelectItem struct {
	// Description: Optional. The body text of the card.
	Description string `json:"description,omitempty"`

	// Image: Optional. The image to display.
	Image *GoogleCloudDialogflowV2beta1IntentMessageImage `json:"image,omitempty"`

	// Info: Required. Additional info about the option item.
	Info *GoogleCloudDialogflowV2beta1IntentMessageSelectItemInfo `json:"info,omitempty"`

	// Title: Required. Title of the carousel item.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1IntentMessageCarouselSelectItem) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1IntentMessageCarouselSelectItem
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1IntentMessageImage: The image response
// message.
type GoogleCloudDialogflowV2beta1IntentMessageImage struct {
	// AccessibilityText: A text description of the image to be used for
	// accessibility,
	// e.g., screen readers. Required if image_uri is set for
	// CarouselSelect.
	AccessibilityText string `json:"accessibilityText,omitempty"`

	// ImageUri: Optional. The public URI to an image file.
	ImageUri string `json:"imageUri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AccessibilityText")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AccessibilityText") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1IntentMessageImage) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1IntentMessageImage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1IntentMessageLinkOutSuggestion: The
// suggestion chip message that allows the user to jump out to the
// app
// or website associated with this agent.
type GoogleCloudDialogflowV2beta1IntentMessageLinkOutSuggestion struct {
	// DestinationName: Required. The name of the app or site this chip is
	// linking to.
	DestinationName string `json:"destinationName,omitempty"`

	// Uri: Required. The URI of the app or site to open when the user taps
	// the
	// suggestion chip.
	Uri string `json:"uri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DestinationName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DestinationName") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1IntentMessageLinkOutSuggestion) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1IntentMessageLinkOutSuggestion
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1IntentMessageListSelect: The card for
// presenting a list of options to select from.
type GoogleCloudDialogflowV2beta1IntentMessageListSelect struct {
	// Items: Required. List items.
	Items []*GoogleCloudDialogflowV2beta1IntentMessageListSelectItem `json:"items,omitempty"`

	// Title: Optional. The overall title of the list.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1IntentMessageListSelect) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1IntentMessageListSelect
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1IntentMessageListSelectItem: An item in
// the list.
type GoogleCloudDialogflowV2beta1IntentMessageListSelectItem struct {
	// Description: Optional. The main text describing the item.
	Description string `json:"description,omitempty"`

	// Image: Optional. The image to display.
	Image *GoogleCloudDialogflowV2beta1IntentMessageImage `json:"image,omitempty"`

	// Info: Required. Additional information about this option.
	Info *GoogleCloudDialogflowV2beta1IntentMessageSelectItemInfo `json:"info,omitempty"`

	// Title: Required. The title of the list item.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1IntentMessageListSelectItem) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1IntentMessageListSelectItem
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1IntentMessageQuickReplies: The quick
// replies response message.
type GoogleCloudDialogflowV2beta1IntentMessageQuickReplies struct {
	// QuickReplies: Optional. The collection of quick replies.
	QuickReplies []string `json:"quickReplies,omitempty"`

	// Title: Optional. The title of the collection of quick replies.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "QuickReplies") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "QuickReplies") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1IntentMessageQuickReplies) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1IntentMessageQuickReplies
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1IntentMessageSelectItemInfo: Additional
// info about the select item for when it is triggered in a
// dialog.
type GoogleCloudDialogflowV2beta1IntentMessageSelectItemInfo struct {
	// Key: Required. A unique key that will be sent back to the agent if
	// this
	// response is given.
	Key string `json:"key,omitempty"`

	// Synonyms: Optional. A list of synonyms that can also be used to
	// trigger this
	// item in dialog.
	Synonyms []string `json:"synonyms,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Key") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Key") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1IntentMessageSelectItemInfo) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1IntentMessageSelectItemInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1IntentMessageSimpleResponse: The simple
// response message containing speech or text.
type GoogleCloudDialogflowV2beta1IntentMessageSimpleResponse struct {
	// DisplayText: Optional. The text to display.
	DisplayText string `json:"displayText,omitempty"`

	// Ssml: One of text_to_speech or ssml must be provided. Structured
	// spoken
	// response to the user in the SSML format. Mutually exclusive
	// with
	// text_to_speech.
	Ssml string `json:"ssml,omitempty"`

	// TextToSpeech: One of text_to_speech or ssml must be provided. The
	// plain text of the
	// speech output. Mutually exclusive with ssml.
	TextToSpeech string `json:"textToSpeech,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DisplayText") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DisplayText") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1IntentMessageSimpleResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1IntentMessageSimpleResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1IntentMessageSimpleResponses: The
// collection of simple response candidates.
// This message in `QueryResult.fulfillment_messages`
// and
// `WebhookResponse.fulfillment_messages` should contain only
// one
// `SimpleResponse`.
type GoogleCloudDialogflowV2beta1IntentMessageSimpleResponses struct {
	// SimpleResponses: Required. The list of simple responses.
	SimpleResponses []*GoogleCloudDialogflowV2beta1IntentMessageSimpleResponse `json:"simpleResponses,omitempty"`

	// ForceSendFields is a list of field names (e.g. "SimpleResponses") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "SimpleResponses") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1IntentMessageSimpleResponses) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1IntentMessageSimpleResponses
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1IntentMessageSuggestion: The suggestion
// chip message that the user can tap to quickly post a reply
// to the conversation.
type GoogleCloudDialogflowV2beta1IntentMessageSuggestion struct {
	// Title: Required. The text shown the in the suggestion chip.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Title") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Title") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1IntentMessageSuggestion) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1IntentMessageSuggestion
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1IntentMessageSuggestions: The collection
// of suggestions.
type GoogleCloudDialogflowV2beta1IntentMessageSuggestions struct {
	// Suggestions: Required. The list of suggested replies.
	Suggestions []*GoogleCloudDialogflowV2beta1IntentMessageSuggestion `json:"suggestions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Suggestions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Suggestions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1IntentMessageSuggestions) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1IntentMessageSuggestions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1IntentMessageTelephonyPlayAudio: Plays
// audio from a file in Telephony Gateway.
type GoogleCloudDialogflowV2beta1IntentMessageTelephonyPlayAudio struct {
	// AudioUri: Required. URI to a Google Cloud Storage object containing
	// the audio to
	// play, e.g., "gs://bucket/object". The object must contain a
	// single
	// channel (mono) of linear PCM audio (2 bytes / sample) at 8kHz.
	//
	// This object must be readable by the
	// `service-<Project
	// Number>@gcp-sa-dialogflow.iam.gserviceaccount.com` service
	// account
	// where <Project Number> is the number of the Telephony Gateway
	// project
	// (usually the same as the Dialogflow agent project). If the Google
	// Cloud
	// Storage bucket is in the Telephony Gateway project, this permission
	// is
	// added by default when enabling the Dialogflow V2 API.
	//
	// For audio from other sources, consider using
	// the
	// `TelephonySynthesizeSpeech` message with SSML.
	AudioUri string `json:"audioUri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AudioUri") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AudioUri") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1IntentMessageTelephonyPlayAudio) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1IntentMessageTelephonyPlayAudio
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1IntentMessageTelephonySynthesizeSpeech:
// Synthesizes speech and plays back the synthesized audio to the caller
// in
// Telephony Gateway.
//
// Telephony Gateway takes the synthesizer settings
// from
// `DetectIntentResponse.output_audio_config` which can either be set
// at request-level or can come from the agent-level synthesizer config.
type GoogleCloudDialogflowV2beta1IntentMessageTelephonySynthesizeSpeech struct {
	// Ssml: The SSML to be synthesized. For more information,
	// see
	// [SSML](https://developers.google.com/actions/reference/ssml).
	Ssml string `json:"ssml,omitempty"`

	// Text: The raw text to be synthesized.
	Text string `json:"text,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Ssml") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Ssml") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1IntentMessageTelephonySynthesizeSpeech) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1IntentMessageTelephonySynthesizeSpeech
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1IntentMessageTelephonyTransferCall:
// Transfers the call in Telephony Gateway.
type GoogleCloudDialogflowV2beta1IntentMessageTelephonyTransferCall struct {
	// PhoneNumber: Required. The phone number to transfer the call to
	// in [E.164 format](https://en.wikipedia.org/wiki/E.164).
	//
	// We currently only allow transferring to US numbers (+1xxxyyyzzzz).
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// ForceSendFields is a list of field names (e.g. "PhoneNumber") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "PhoneNumber") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1IntentMessageTelephonyTransferCall) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1IntentMessageTelephonyTransferCall
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1IntentMessageText: The text response
// message.
type GoogleCloudDialogflowV2beta1IntentMessageText struct {
	// Text: Optional. The collection of the agent's responses.
	Text []string `json:"text,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Text") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Text") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1IntentMessageText) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1IntentMessageText
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1IntentParameter: Represents intent
// parameters.
type GoogleCloudDialogflowV2beta1IntentParameter struct {
	// DefaultValue: Optional. The default value to use when the `value`
	// yields an empty
	// result.
	// Default values can be extracted from contexts by using the
	// following
	// syntax: `#context_name.parameter_name`.
	DefaultValue string `json:"defaultValue,omitempty"`

	// DisplayName: Required. The name of the parameter.
	DisplayName string `json:"displayName,omitempty"`

	// EntityTypeDisplayName: Optional. The name of the entity type,
	// prefixed with `@`, that
	// describes values of the parameter. If the parameter is
	// required, this must be provided.
	EntityTypeDisplayName string `json:"entityTypeDisplayName,omitempty"`

	// IsList: Optional. Indicates whether the parameter represents a list
	// of values.
	IsList bool `json:"isList,omitempty"`

	// Mandatory: Optional. Indicates whether the parameter is required.
	// That is,
	// whether the intent cannot be completed without collecting the
	// parameter
	// value.
	Mandatory bool `json:"mandatory,omitempty"`

	// Name: The unique identifier of this parameter.
	Name string `json:"name,omitempty"`

	// Prompts: Optional. The collection of prompts that the agent can
	// present to the
	// user in order to collect value for the parameter.
	Prompts []string `json:"prompts,omitempty"`

	// Value: Optional. The definition of the parameter value. It can be:
	// - a constant string,
	// - a parameter value defined as `$parameter_name`,
	// - an original parameter value defined as
	// `$parameter_name.original`,
	// - a parameter value from some context defined as
	//   `#context_name.parameter_name`.
	Value string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DefaultValue") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DefaultValue") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1IntentParameter) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1IntentParameter
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1IntentTrainingPhrase: Represents an
// example or template that the agent is trained on.
type GoogleCloudDialogflowV2beta1IntentTrainingPhrase struct {
	// Name: Output only. The unique identifier of this training phrase.
	Name string `json:"name,omitempty"`

	// Parts: Required. The collection of training phrase parts (can be
	// annotated).
	// Fields: `entity_type`, `alias` and `user_defined` should be
	// populated
	// only for the annotated parts of the training phrase.
	Parts []*GoogleCloudDialogflowV2beta1IntentTrainingPhrasePart `json:"parts,omitempty"`

	// TimesAddedCount: Optional. Indicates how many times this example or
	// template was added to
	// the intent. Each time a developer adds an existing sample by editing
	// an
	// intent or training, this counter is increased.
	TimesAddedCount int64 `json:"timesAddedCount,omitempty"`

	// Type: Required. The type of the training phrase.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED" - Not specified. This value should never be
	// used.
	//   "EXAMPLE" - Examples do not contain @-prefixed entity type names,
	// but example parts
	// can be annotated with entity types.
	//   "TEMPLATE" - Templates are not annotated with entity types, but
	// they can contain
	// @-prefixed entity type names as substrings.
	Type string `json:"type,omitempty"`

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

func (s *GoogleCloudDialogflowV2beta1IntentTrainingPhrase) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1IntentTrainingPhrase
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1IntentTrainingPhrasePart: Represents a
// part of a training phrase.
type GoogleCloudDialogflowV2beta1IntentTrainingPhrasePart struct {
	// Alias: Optional. The parameter name for the value extracted from
	// the
	// annotated part of the example.
	Alias string `json:"alias,omitempty"`

	// EntityType: Optional. The entity type name prefixed with `@`. This
	// field is
	// required for the annotated part of the text and applies only
	// to
	// examples.
	EntityType string `json:"entityType,omitempty"`

	// Text: Required. The text corresponding to the example or template,
	// if there are no annotations. For
	// annotated examples, it is the text for one of the example's parts.
	Text string `json:"text,omitempty"`

	// UserDefined: Optional. Indicates whether the text was manually
	// annotated by the
	// developer.
	UserDefined bool `json:"userDefined,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Alias") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Alias") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1IntentTrainingPhrasePart) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1IntentTrainingPhrasePart
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1KnowledgeAnswers: Represents the result
// of querying a Knowledge base.
type GoogleCloudDialogflowV2beta1KnowledgeAnswers struct {
	// Answers: A list of answers from Knowledge Connector.
	Answers []*GoogleCloudDialogflowV2beta1KnowledgeAnswersAnswer `json:"answers,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Answers") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Answers") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1KnowledgeAnswers) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1KnowledgeAnswers
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1KnowledgeAnswersAnswer: An answer from
// Knowledge Connector.
type GoogleCloudDialogflowV2beta1KnowledgeAnswersAnswer struct {
	// Answer: The piece of text from the `source` knowledge base document
	// that answers
	// this conversational query.
	Answer string `json:"answer,omitempty"`

	// FaqQuestion: The corresponding FAQ question if the answer was
	// extracted from a FAQ
	// Document, empty otherwise.
	FaqQuestion string `json:"faqQuestion,omitempty"`

	// MatchConfidence: The system's confidence score that this Knowledge
	// answer is a good match
	// for this converstational query, range from 0.0 (completely
	// uncertain)
	// to 1.0 (completely certain).
	// Note: The confidence score is likely to vary somewhat (possibly even
	// for
	// identical requests), as the underlying model is under
	// constant
	// improvement, we may deprecate it in the future. We recommend
	// using
	// `match_confidence_level` which should be generally more stable.
	MatchConfidence float64 `json:"matchConfidence,omitempty"`

	// MatchConfidenceLevel: The system's confidence level that this
	// knowledge answer is a good match
	// for this conversational query.
	// NOTE: The confidence level for a given `<query, answer>` pair may
	// change
	// without notice, as it depends on models that are constantly
	// being
	// improved. However, it will change less frequently than the
	// confidence
	// score below, and should be preferred for referencing the quality of
	// an
	// answer.
	//
	// Possible values:
	//   "MATCH_CONFIDENCE_LEVEL_UNSPECIFIED" - Not specified.
	//   "LOW" - Indicates that the confidence is low.
	//   "MEDIUM" - Indicates our confidence is medium.
	//   "HIGH" - Indicates our confidence is high.
	MatchConfidenceLevel string `json:"matchConfidenceLevel,omitempty"`

	// Source: Indicates which Knowledge Document this answer was extracted
	// from.
	// Format: `projects/<Project ID>/knowledgeBases/<Knowledge
	// Base
	// ID>/documents/<Document ID>`.
	Source string `json:"source,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Answer") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Answer") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1KnowledgeAnswersAnswer) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1KnowledgeAnswersAnswer
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudDialogflowV2beta1KnowledgeAnswersAnswer) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudDialogflowV2beta1KnowledgeAnswersAnswer
	var s1 struct {
		MatchConfidence gensupport.JSONFloat64 `json:"matchConfidence"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.MatchConfidence = float64(s1.MatchConfidence)
	return nil
}

// GoogleCloudDialogflowV2beta1KnowledgeBase: Represents knowledge base
// resource.
//
// Note: resource `projects.agent.knowledgeBases` is deprecated, please
// use
// `projects.knowledgeBases` instead.
type GoogleCloudDialogflowV2beta1KnowledgeBase struct {
	// DisplayName: Required. The display name of the knowledge base. The
	// name must be 1024
	// bytes or less; otherwise, the creation request fails.
	DisplayName string `json:"displayName,omitempty"`

	// Name: The knowledge base resource name.
	// The name must be empty when creating a knowledge base.
	// Format: `projects/<Project ID>/knowledgeBases/<Knowledge Base ID>`.
	Name string `json:"name,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "DisplayName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DisplayName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1KnowledgeBase) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1KnowledgeBase
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1KnowledgeOperationMetadata: Metadata in
// google::longrunning::Operation for Knowledge operations.
type GoogleCloudDialogflowV2beta1KnowledgeOperationMetadata struct {
	// State: Required. The current state of this operation.
	//
	// Possible values:
	//   "STATE_UNSPECIFIED" - State unspecified.
	//   "PENDING" - The operation has been created.
	//   "RUNNING" - The operation is currently running.
	//   "DONE" - The operation is done, either cancelled or completed.
	State string `json:"state,omitempty"`

	// ForceSendFields is a list of field names (e.g. "State") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "State") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1KnowledgeOperationMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1KnowledgeOperationMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1ListContextsResponse: The response
// message for Contexts.ListContexts.
type GoogleCloudDialogflowV2beta1ListContextsResponse struct {
	// Contexts: The list of contexts. There will be a maximum number of
	// items
	// returned based on the page_size field in the request.
	Contexts []*GoogleCloudDialogflowV2beta1Context `json:"contexts,omitempty"`

	// NextPageToken: Token to retrieve the next page of results, or empty
	// if there are no
	// more results in the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Contexts") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Contexts") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1ListContextsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1ListContextsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1ListConversationProfilesResponse: The
// response message for ConversationProfiles.ListConversationProfiles.
type GoogleCloudDialogflowV2beta1ListConversationProfilesResponse struct {
	// ConversationProfiles: The list of project conversation profiles.
	// There is a maximum number
	// of items returned based on the page_size field in the request.
	ConversationProfiles []*GoogleCloudDialogflowV2beta1ConversationProfile `json:"conversationProfiles,omitempty"`

	// NextPageToken: Token to retrieve the next page of results, or empty
	// if there are no
	// more results in the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g.
	// "ConversationProfiles") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ConversationProfiles") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1ListConversationProfilesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1ListConversationProfilesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1ListConversationsResponse: The response
// message for Conversations.ListConversations.
type GoogleCloudDialogflowV2beta1ListConversationsResponse struct {
	// Conversations: The list of conversations. There will be a maximum
	// number of items
	// returned based on the page_size field in the request.
	Conversations []*GoogleCloudDialogflowV2beta1Conversation `json:"conversations,omitempty"`

	// NextPageToken: Token to retrieve the next page of results, or empty
	// if there are no
	// more results in the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Conversations") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Conversations") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1ListConversationsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1ListConversationsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1ListDocumentsResponse: Response message
// for Documents.ListDocuments.
type GoogleCloudDialogflowV2beta1ListDocumentsResponse struct {
	// Documents: The list of documents.
	Documents []*GoogleCloudDialogflowV2beta1Document `json:"documents,omitempty"`

	// NextPageToken: Token to retrieve the next page of results, or empty
	// if there are no
	// more results in the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Documents") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Documents") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1ListDocumentsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1ListDocumentsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1ListEntityTypesResponse: The response
// message for EntityTypes.ListEntityTypes.
type GoogleCloudDialogflowV2beta1ListEntityTypesResponse struct {
	// EntityTypes: The list of agent entity types. There will be a maximum
	// number of items
	// returned based on the page_size field in the request.
	EntityTypes []*GoogleCloudDialogflowV2beta1EntityType `json:"entityTypes,omitempty"`

	// NextPageToken: Token to retrieve the next page of results, or empty
	// if there are no
	// more results in the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "EntityTypes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EntityTypes") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1ListEntityTypesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1ListEntityTypesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1ListHumanAgentAssistantsResponse: The
// response message for HumanAgentAssistants.ListHumanAgentAssistants.
type GoogleCloudDialogflowV2beta1ListHumanAgentAssistantsResponse struct {
	// HumanAgentAssistants: The list of project agent assistants. There is
	// a maximum number of
	// items returned based on the page_size field in the request.
	HumanAgentAssistants []*GoogleCloudDialogflowV2beta1HumanAgentAssistant `json:"humanAgentAssistants,omitempty"`

	// NextPageToken: Token to retrieve the next page of results or empty if
	// there are no
	// more results in the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g.
	// "HumanAgentAssistants") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "HumanAgentAssistants") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1ListHumanAgentAssistantsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1ListHumanAgentAssistantsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1ListIntentsResponse: The response message
// for Intents.ListIntents.
type GoogleCloudDialogflowV2beta1ListIntentsResponse struct {
	// Intents: The list of agent intents. There will be a maximum number of
	// items
	// returned based on the page_size field in the request.
	Intents []*GoogleCloudDialogflowV2beta1Intent `json:"intents,omitempty"`

	// NextPageToken: Token to retrieve the next page of results, or empty
	// if there are no
	// more results in the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Intents") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Intents") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1ListIntentsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1ListIntentsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1ListKnowledgeBasesResponse: Response
// message for KnowledgeBases.ListKnowledgeBases.
type GoogleCloudDialogflowV2beta1ListKnowledgeBasesResponse struct {
	// KnowledgeBases: The list of knowledge bases.
	KnowledgeBases []*GoogleCloudDialogflowV2beta1KnowledgeBase `json:"knowledgeBases,omitempty"`

	// NextPageToken: Token to retrieve the next page of results, or empty
	// if there are no
	// more results in the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "KnowledgeBases") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "KnowledgeBases") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1ListKnowledgeBasesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1ListKnowledgeBasesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1ListMessagesResponse: The response
// message for Conversations.ListMessages.
type GoogleCloudDialogflowV2beta1ListMessagesResponse struct {
	// Messages: Required. The list of messages. There will be a maximum
	// number of items
	// returned based on the page_size field in the request.
	Messages []*GoogleCloudDialogflowV2beta1Message `json:"messages,omitempty"`

	// NextPageToken: Optional. Token to retrieve the next page of results,
	// or empty if there are
	// no more results in the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Messages") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Messages") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1ListMessagesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1ListMessagesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1ListParticipantsResponse: The response
// message for Conversations.ListParticipants.
type GoogleCloudDialogflowV2beta1ListParticipantsResponse struct {
	// NextPageToken: Token to retrieve the next page of results or empty if
	// there are no
	// more results in the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Participants: The list of participants. There is a maximum number of
	// items
	// returned based on the page_size field in the request.
	Participants []*GoogleCloudDialogflowV2beta1Participant `json:"participants,omitempty"`

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

func (s *GoogleCloudDialogflowV2beta1ListParticipantsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1ListParticipantsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1ListPhoneNumberOrdersResponse: The
// response message for PhoneNumberOrders.ListPhoneNumberOrders.
type GoogleCloudDialogflowV2beta1ListPhoneNumberOrdersResponse struct {
	// NextPageToken: Token to retrieve the next page of results, or empty
	// if there are no
	// more results in the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PhoneNumberOrders: The list of orders. There is a maximum number of
	// items returned based
	// on the page_size field in the request.
	PhoneNumberOrders []*GoogleCloudDialogflowV2beta1PhoneNumberOrder `json:"phoneNumberOrders,omitempty"`

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

func (s *GoogleCloudDialogflowV2beta1ListPhoneNumberOrdersResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1ListPhoneNumberOrdersResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1ListPhoneNumbersResponse: The response
// message for PhoneNumbers.ListPhoneNumbers.
type GoogleCloudDialogflowV2beta1ListPhoneNumbersResponse struct {
	// NextPageToken: Token to retrieve the next page of results, or empty
	// if there are no
	// more results in the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PhoneNumbers: The list of `PhoneNumber` resources. There is a maximum
	// number of items
	// returned based on the page_size field in the request.
	PhoneNumbers []*GoogleCloudDialogflowV2beta1PhoneNumber `json:"phoneNumbers,omitempty"`

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

func (s *GoogleCloudDialogflowV2beta1ListPhoneNumbersResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1ListPhoneNumbersResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1ListSessionEntityTypesResponse: The
// response message for SessionEntityTypes.ListSessionEntityTypes.
type GoogleCloudDialogflowV2beta1ListSessionEntityTypesResponse struct {
	// NextPageToken: Token to retrieve the next page of results, or empty
	// if there are no
	// more results in the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SessionEntityTypes: The list of session entity types. There will be a
	// maximum number of items
	// returned based on the page_size field in the request.
	SessionEntityTypes []*GoogleCloudDialogflowV2beta1SessionEntityType `json:"sessionEntityTypes,omitempty"`

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

func (s *GoogleCloudDialogflowV2beta1ListSessionEntityTypesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1ListSessionEntityTypesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1ListSuggestionsResponse: The response
// message for [Conversations.ListSuggestions]
type GoogleCloudDialogflowV2beta1ListSuggestionsResponse struct {
	// NextPageToken: Optional. Token to retrieve the next page of results
	// or empty if there are
	// no more results in the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Suggestions: Required.
	Suggestions []*GoogleCloudDialogflowV2beta1Suggestion `json:"suggestions,omitempty"`

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

func (s *GoogleCloudDialogflowV2beta1ListSuggestionsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1ListSuggestionsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1LoggingConfig: Defines logging behavior
// for conversation lifecycle events.
type GoogleCloudDialogflowV2beta1LoggingConfig struct {
}

// GoogleCloudDialogflowV2beta1Message: Represents a message posted into
// a conversation.
type GoogleCloudDialogflowV2beta1Message struct {
	// Content: Required. The message content.
	Content string `json:"content,omitempty"`

	// CreateTime: Optional. The time when the message was sent.
	CreateTime string `json:"createTime,omitempty"`

	// LanguageCode: Required. The message language.
	// This should be a
	// [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt)
	// language tag. Example: "en-US".
	LanguageCode string `json:"languageCode,omitempty"`

	// Name: Required. The unique identifier of the message.
	// Format: `projects/<Project
	// ID>/conversations/<Conversation
	// ID>/messages/<Message ID>`.
	Name string `json:"name,omitempty"`

	// Participant: Required. The participant that said this message.
	Participant string `json:"participant,omitempty"`

	// ParticipantRole: Optional. The role of the participant.
	//
	// Possible values:
	//   "ROLE_UNSPECIFIED" - Participant role not set.
	//   "HUMAN_AGENT" - Participant is a human agent.
	//   "AUTOMATED_AGENT" - Participant is an automated agent, such as a
	// Dialogflow agent.
	//   "END_USER" - Participant is an end user that has called or chatted
	// with
	// Dialogflow services.
	ParticipantRole string `json:"participantRole,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Content") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Content") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1Message) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1Message
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1NotificationConfig: Defines notification
// behavior for conversation lifecycle events.
type GoogleCloudDialogflowV2beta1NotificationConfig struct {
	// Topic: Optional. Name of the Cloud Pub/Sub topic to publish
	// conversation
	// events like
	// CONVERSATION_STARTED as
	// serialized ConversationEvent protos.
	//
	// If enable_notifications is
	// `true` and no topic is supplied, a new topic is created and
	// listed
	// here.
	//
	// Notification works for phone calls, if this topic either is in the
	// same
	// project as the conversation or you grant `service-<Conversation
	// Project
	// Number>@gcp-sa-dialogflow.iam.gserviceaccount.com` the `Dialogflow
	// Service
	// Agent` role in the topic project.
	//
	// Format: `projects/<Project ID>/topics/<Topic ID>`.
	Topic string `json:"topic,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Topic") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Topic") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1NotificationConfig) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1NotificationConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1OriginalDetectIntentRequest: Represents
// the contents of the original request that was passed to
// the `[Streaming]DetectIntent` call.
type GoogleCloudDialogflowV2beta1OriginalDetectIntentRequest struct {
	// Payload: Optional. This field is set to the value of the
	// `QueryParameters.payload`
	// field passed in the request. Some integrations that query a
	// Dialogflow
	// agent may provide additional information in the payload.
	//
	// In particular for the Telephony Gateway this field has the
	// form:
	// <pre>{
	//  "telephony": {
	//    "caller_id": "+18558363987"
	//  }
	// }</pre>
	// Note: The caller ID field (`caller_id`) will be redacted for
	// Standard
	// Edition agents and populated with the caller ID in
	// [E.164
	// format](https://en.wikipedia.org/wiki/E.164) for Enterprise Edition
	// agents.
	Payload googleapi.RawMessage `json:"payload,omitempty"`

	// Source: The source of this request, e.g., `google`, `facebook`,
	// `slack`. It is set
	// by Dialogflow-owned servers.
	Source string `json:"source,omitempty"`

	// Version: Optional. The version of the protocol used for this
	// request.
	// This field is AoG-specific.
	Version string `json:"version,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Payload") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Payload") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1OriginalDetectIntentRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1OriginalDetectIntentRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1OutputAudio: Represents the natural
// language speech audio to be played to the end user.
type GoogleCloudDialogflowV2beta1OutputAudio struct {
	// Audio: Required. The natural language speech audio.
	Audio string `json:"audio,omitempty"`

	// Config: Required. Instructs the speech synthesizer how to generate
	// the speech
	// audio.
	Config *GoogleCloudDialogflowV2beta1OutputAudioConfig `json:"config,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Audio") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Audio") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1OutputAudio) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1OutputAudio
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1OutputAudioConfig: Instructs the speech
// synthesizer how to generate the output audio content.
type GoogleCloudDialogflowV2beta1OutputAudioConfig struct {
	// AudioEncoding: Required. Audio encoding of the synthesized audio
	// content.
	//
	// Possible values:
	//   "OUTPUT_AUDIO_ENCODING_UNSPECIFIED" - Not specified.
	//   "OUTPUT_AUDIO_ENCODING_LINEAR_16" - Uncompressed 16-bit signed
	// little-endian samples (Linear PCM).
	// Audio content returned as LINEAR16 also contains a WAV header.
	//   "OUTPUT_AUDIO_ENCODING_MP3" - MP3 audio.
	//   "OUTPUT_AUDIO_ENCODING_OGG_OPUS" - Opus encoded audio wrapped in an
	// ogg container. The result will be a
	// file which can be played natively on Android, and in browsers (at
	// least
	// Chrome and Firefox). The quality of the encoding is considerably
	// higher
	// than MP3 while using approximately the same bitrate.
	AudioEncoding string `json:"audioEncoding,omitempty"`

	// SampleRateHertz: Optional. The synthesis sample rate (in hertz) for
	// this audio. If not
	// provided, then the synthesizer will use the default sample rate based
	// on
	// the audio encoding. If this is different from the voice's natural
	// sample
	// rate, then the synthesizer will honor this request by converting to
	// the
	// desired sample rate (which might result in worse audio quality).
	SampleRateHertz int64 `json:"sampleRateHertz,omitempty"`

	// SynthesizeSpeechConfig: Optional. Configuration of how speech should
	// be synthesized.
	SynthesizeSpeechConfig *GoogleCloudDialogflowV2beta1SynthesizeSpeechConfig `json:"synthesizeSpeechConfig,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AudioEncoding") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AudioEncoding") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1OutputAudioConfig) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1OutputAudioConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1Participant: Represents a single side of
// the conversation.
type GoogleCloudDialogflowV2beta1Participant struct {
	// Name: Required. The unique identifier of this participant.
	// Format: `projects/<Project
	// ID>/conversations/<Conversation
	// ID>/participants/<Participant ID>`.
	Name string `json:"name,omitempty"`

	// Role: Required. The role this participant plays in the conversation.
	//
	// Possible values:
	//   "ROLE_UNSPECIFIED" - Participant role not set.
	//   "HUMAN_AGENT" - Participant is a human agent.
	//   "AUTOMATED_AGENT" - Participant is an automated agent, such as a
	// Dialogflow agent.
	//   "END_USER" - Participant is an end user that has called or chatted
	// with
	// Dialogflow services.
	Role string `json:"role,omitempty"`

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

func (s *GoogleCloudDialogflowV2beta1Participant) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1Participant
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1PhoneNumber: Represents a phone
// number.
// `PhoneNumber` resources enable phone calls to be answered by
// Dialogflow
// services and are added to a project through a `PhoneNumberOrder`.
type GoogleCloudDialogflowV2beta1PhoneNumber struct {
	// ConversationProfile: Optional. The conversation profile calls to this
	// `PhoneNumber` should use.
	// Format: `projects/<Project
	// ID>/conversationProfiles/<ConversationProfile
	// ID>`.
	ConversationProfile string `json:"conversationProfile,omitempty"`

	// LifecycleState: Output only. The state of the `PhoneNumber`. Defaults
	// to `ACTIVE`.
	// `PhoneNumber` objects set to `DELETE_REQUESTED` always decline
	// incoming
	// calls and can be removed completely within 30 days.
	//
	// Possible values:
	//   "LIFECYCLE_STATE_UNSPECIFIED" - This value is never used.
	//   "ACTIVE" - Number is active and can receive phone calls.
	//   "DELETE_REQUESTED" - Number is pending deletion, and cannot receive
	// calls.
	LifecycleState string `json:"lifecycleState,omitempty"`

	// Name: Required. The unique identifier of this phone number.
	// Format: `projects/<Project ID>/phoneNumbers/<PhoneNumber ID>`.
	Name string `json:"name,omitempty"`

	// PhoneNumber: Output only. Phone number in
	// [E.164](https://en.wikipedia.org/wiki/E.164)
	// format. An example of a correctly formatted phone number:
	// +15556767888.
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ConversationProfile")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ConversationProfile") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1PhoneNumber) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1PhoneNumber
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1PhoneNumberOrder: Represents a phone
// number order.
// Orders can assign phone numbers to projects.
type GoogleCloudDialogflowV2beta1PhoneNumberOrder struct {
	// CreateTime: Output only. The time this order was created.
	CreateTime string `json:"createTime,omitempty"`

	// Description: Optional. A description of the order, limit is 1024
	// bytes.
	Description string `json:"description,omitempty"`

	// LifecycleState: Output only. The current status of the order.
	//
	// Possible values:
	//   "LIFECYCLE_STATE_UNSPECIFIED" - Unknown.
	//   "PENDING" - Order is awaiting action.
	//   "IN_PROGRESS" - Order is being worked on, and is partially
	// fulfilled.
	//   "COMPLETED" - Order has been fulfilled.
	//   "CANCELLED" - Order has been cancelled.
	LifecycleState string `json:"lifecycleState,omitempty"`

	// Name: Required. The unique identifier of this order.
	// Format: `projects/<Project ID>/phoneNumberOrders/<Order ID>`.
	Name string `json:"name,omitempty"`

	// PhoneNumberSpec: Order is for new numbers.
	PhoneNumberSpec *GoogleCloudDialogflowV2beta1PhoneNumberSpec `json:"phoneNumberSpec,omitempty"`

	// PhoneNumbers: Output only. A map of ordered numbers filled so far,
	// keyed by their
	// resource name. Key format:
	// `projects/<Project ID>/phoneNumbers/<PhoneNumber ID>`.
	// Value format: E.164 phone number. Output only.
	PhoneNumbers map[string]string `json:"phoneNumbers,omitempty"`

	// UpdateTime: Output only. The time this order was last updated.
	UpdateTime string `json:"updateTime,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CreateTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CreateTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1PhoneNumberOrder) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1PhoneNumberOrder
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1PhoneNumberSpec: Request for new numbers
// fitting a set of parameters.
// The country code for newly requested numbers defaults to 1 (US) until
// the
// service is available in other regions.
type GoogleCloudDialogflowV2beta1PhoneNumberSpec struct {
	// Count: Required. Total numbers requested, between 1 and 10 inclusive.
	Count int64 `json:"count,omitempty"`

	// PreferredAreaCodes: Optional. Area codes to use. An empty list means
	// 'any code'. Each value
	// is treated as equally preferred. Each entry has a limit of 10
	// bytes.
	// "area code" corresponds to "National Destination Code" described
	// in
	// [E.164](https://en.wikipedia.org/wiki/E.164) standard.
	PreferredAreaCodes []string `json:"preferredAreaCodes,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Count") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Count") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1PhoneNumberSpec) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1PhoneNumberSpec
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1QueryInput: Represents the query input.
// It can contain either:
//
// 1.  An audio config which
//     instructs the speech recognizer how to process the speech
// audio.
//
// 2.  A conversational query in the form of text,.
//
// 3.  An event that specifies which intent to trigger.
type GoogleCloudDialogflowV2beta1QueryInput struct {
	// AudioConfig: Instructs the speech recognizer how to process the
	// speech audio.
	AudioConfig *GoogleCloudDialogflowV2beta1InputAudioConfig `json:"audioConfig,omitempty"`

	// Event: The event to be processed.
	Event *GoogleCloudDialogflowV2beta1EventInput `json:"event,omitempty"`

	// Text: The natural language text to be processed.
	Text *GoogleCloudDialogflowV2beta1TextInput `json:"text,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AudioConfig") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AudioConfig") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1QueryInput) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1QueryInput
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1QueryParameters: Represents the
// parameters of the conversational query.
type GoogleCloudDialogflowV2beta1QueryParameters struct {
	// Contexts: Optional. The collection of contexts to be activated before
	// this query is
	// executed.
	Contexts []*GoogleCloudDialogflowV2beta1Context `json:"contexts,omitempty"`

	// GeoLocation: Optional. The geo location of this conversational query.
	GeoLocation *GoogleTypeLatLng `json:"geoLocation,omitempty"`

	// KnowledgeBaseNames: Optional. KnowledgeBases to get alternative
	// results from. If not set, the
	// KnowledgeBases enabled in the agent (through UI) will be
	// used.
	// Format:  `projects/<Project ID>/knowledgeBases/<Knowledge Base ID>`.
	KnowledgeBaseNames []string `json:"knowledgeBaseNames,omitempty"`

	// Payload: Optional. This field can be used to pass custom data into
	// the webhook
	// associated with the agent. Arbitrary JSON objects are supported.
	Payload googleapi.RawMessage `json:"payload,omitempty"`

	// ResetContexts: Optional. Specifies whether to delete all contexts in
	// the current session
	// before the new ones are activated.
	ResetContexts bool `json:"resetContexts,omitempty"`

	// SentimentAnalysisRequestConfig: Optional. Configures the type of
	// sentiment analysis to perform. If not
	// provided, sentiment analysis is not performed.
	// Note: Sentiment Analysis is only currently available for Enterprise
	// Edition
	// agents.
	SentimentAnalysisRequestConfig *GoogleCloudDialogflowV2beta1SentimentAnalysisRequestConfig `json:"sentimentAnalysisRequestConfig,omitempty"`

	// SessionEntityTypes: Optional. Additional session entity types to
	// replace or extend developer
	// entity types with. The entity synonyms apply to all languages and
	// persist
	// for the session of this query.
	SessionEntityTypes []*GoogleCloudDialogflowV2beta1SessionEntityType `json:"sessionEntityTypes,omitempty"`

	// TimeZone: Optional. The time zone of this conversational query from
	// the
	// [time zone database](https://www.iana.org/time-zones),
	// e.g.,
	// America/New_York, Europe/Paris. If not provided, the time zone
	// specified in
	// agent settings is used.
	TimeZone string `json:"timeZone,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Contexts") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Contexts") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1QueryParameters) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1QueryParameters
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1QueryResult: Represents the result of
// conversational query or event processing.
type GoogleCloudDialogflowV2beta1QueryResult struct {
	// Action: The action name from the matched intent.
	Action string `json:"action,omitempty"`

	// AllRequiredParamsPresent: This field is set to:
	// - `false` if the matched intent has required parameters and not all
	// of
	//    the required parameter values have been collected.
	// - `true` if all required parameter values have been collected, or if
	// the
	//    matched intent doesn't contain any required parameters.
	AllRequiredParamsPresent bool `json:"allRequiredParamsPresent,omitempty"`

	// DiagnosticInfo: The free-form diagnostic info. For example, this
	// field
	// could contain webhook call latency.
	DiagnosticInfo googleapi.RawMessage `json:"diagnosticInfo,omitempty"`

	// FulfillmentMessages: The collection of rich messages to present to
	// the user.
	FulfillmentMessages []*GoogleCloudDialogflowV2beta1IntentMessage `json:"fulfillmentMessages,omitempty"`

	// FulfillmentText: The text to be pronounced to the user or shown on
	// the screen.
	FulfillmentText string `json:"fulfillmentText,omitempty"`

	// Intent: The intent that matched the conversational query. Some,
	// not
	// all fields are filled in this message, including but not limited
	// to:
	// `name`, `display_name` and `webhook_state`.
	Intent *GoogleCloudDialogflowV2beta1Intent `json:"intent,omitempty"`

	// IntentDetectionConfidence: The intent detection confidence. Values
	// range from 0.0
	// (completely uncertain) to 1.0 (completely certain).
	// If there are `multiple knowledge_answers` messages, this value is set
	// to
	// the greatest `knowledgeAnswers.match_confidence` value in the list.
	IntentDetectionConfidence float64 `json:"intentDetectionConfidence,omitempty"`

	// KnowledgeAnswers: The result from Knowledge Connector (if any),
	// ordered by decreasing
	// `KnowledgeAnswers.match_confidence`.
	KnowledgeAnswers *GoogleCloudDialogflowV2beta1KnowledgeAnswers `json:"knowledgeAnswers,omitempty"`

	// LanguageCode: The language that was triggered during intent
	// detection.
	// See [Language
	// Support](https://dialogflow.com/docs/reference/language)
	// for a list of the currently supported language codes.
	LanguageCode string `json:"languageCode,omitempty"`

	// OutputContexts: The collection of output contexts. If
	// applicable,
	// `output_contexts.parameters` contains entries with name
	// `<parameter name>.original` containing the original parameter
	// values
	// before the query.
	OutputContexts []*GoogleCloudDialogflowV2beta1Context `json:"outputContexts,omitempty"`

	// Parameters: The collection of extracted parameters.
	Parameters googleapi.RawMessage `json:"parameters,omitempty"`

	// QueryText: The original conversational query text:
	// - If natural language text was provided as input, `query_text`
	// contains
	//   a copy of the input.
	// - If natural language speech audio was provided as input,
	// `query_text`
	//   contains the speech recognition result. If speech recognizer
	// produced
	//   multiple alternatives, a particular one is picked.
	// - If an event was provided as input, `query_text` is not set.
	QueryText string `json:"queryText,omitempty"`

	// SentimentAnalysisResult: The sentiment analysis result, which depends
	// on the
	// `sentiment_analysis_request_config` specified in the request.
	SentimentAnalysisResult *GoogleCloudDialogflowV2beta1SentimentAnalysisResult `json:"sentimentAnalysisResult,omitempty"`

	// SpeechRecognitionConfidence: The Speech recognition confidence
	// between 0.0 and 1.0. A higher number
	// indicates an estimated greater likelihood that the recognized words
	// are
	// correct. The default of 0.0 is a sentinel value indicating that
	// confidence
	// was not set.
	//
	// This field is not guaranteed to be accurate or set. In particular
	// this
	// field isn't set for StreamingDetectIntent since the streaming
	// endpoint has
	// separate confidence estimates per portion of the audio
	// in
	// StreamingRecognitionResult.
	SpeechRecognitionConfidence float64 `json:"speechRecognitionConfidence,omitempty"`

	// WebhookPayload: If the query was fulfilled by a webhook call, this
	// field is set to the
	// value of the `payload` field returned in the webhook response.
	WebhookPayload googleapi.RawMessage `json:"webhookPayload,omitempty"`

	// WebhookSource: If the query was fulfilled by a webhook call, this
	// field is set to the
	// value of the `source` field returned in the webhook response.
	WebhookSource string `json:"webhookSource,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Action") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Action") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1QueryResult) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1QueryResult
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudDialogflowV2beta1QueryResult) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudDialogflowV2beta1QueryResult
	var s1 struct {
		IntentDetectionConfidence   gensupport.JSONFloat64 `json:"intentDetectionConfidence"`
		SpeechRecognitionConfidence gensupport.JSONFloat64 `json:"speechRecognitionConfidence"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.IntentDetectionConfidence = float64(s1.IntentDetectionConfidence)
	s.SpeechRecognitionConfidence = float64(s1.SpeechRecognitionConfidence)
	return nil
}

// GoogleCloudDialogflowV2beta1RestoreAgentRequest: The request message
// for Agents.RestoreAgent.
type GoogleCloudDialogflowV2beta1RestoreAgentRequest struct {
	// AgentContent: The agent to restore.
	//
	// Example for how to restore an agent via the command line:
	// <pre>curl \
	//
	// 'https://dialogflow.googleapis.com/v2beta1/projects/&lt;project_name&g
	// t;/agent:restore\
	//    -X POST \
	//    -H 'Authorization: Bearer '$(gcloud auth application-default
	//    print-access-token) \
	//    -H 'Accept: application/json' \
	//    -H 'Content-Type: application/json' \
	//    --compressed \
	//    --data-binary "{
	//        'agentContent': '$(cat &lt;agent zip file&gt; | base64 -w 0)'
	//    }"</pre>
	AgentContent string `json:"agentContent,omitempty"`

	// AgentUri: The URI to a Google Cloud Storage file containing the agent
	// to restore.
	// Note: The URI must start with "gs://".
	AgentUri string `json:"agentUri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AgentContent") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AgentContent") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1RestoreAgentRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1RestoreAgentRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1SearchAgentsResponse: The response
// message for Agents.SearchAgents.
type GoogleCloudDialogflowV2beta1SearchAgentsResponse struct {
	// Agents: The list of agents. There will be a maximum number of items
	// returned based
	// on the page_size field in the request.
	Agents []*GoogleCloudDialogflowV2beta1Agent `json:"agents,omitempty"`

	// NextPageToken: Token to retrieve the next page of results, or empty
	// if there are no
	// more results in the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Agents") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Agents") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1SearchAgentsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1SearchAgentsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1Sentiment: The sentiment, such as
// positive/negative feeling or association, for a unit
// of analysis, such as the query text.
type GoogleCloudDialogflowV2beta1Sentiment struct {
	// Magnitude: A non-negative number in the [0, +inf) range, which
	// represents the absolute
	// magnitude of sentiment, regardless of score (positive or negative).
	Magnitude float64 `json:"magnitude,omitempty"`

	// Score: Sentiment score between -1.0 (negative sentiment) and 1.0
	// (positive
	// sentiment).
	Score float64 `json:"score,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Magnitude") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Magnitude") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1Sentiment) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1Sentiment
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudDialogflowV2beta1Sentiment) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudDialogflowV2beta1Sentiment
	var s1 struct {
		Magnitude gensupport.JSONFloat64 `json:"magnitude"`
		Score     gensupport.JSONFloat64 `json:"score"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Magnitude = float64(s1.Magnitude)
	s.Score = float64(s1.Score)
	return nil
}

// GoogleCloudDialogflowV2beta1SentimentAnalysisRequestConfig:
// Configures the types of sentiment analysis to perform.
type GoogleCloudDialogflowV2beta1SentimentAnalysisRequestConfig struct {
	// AnalyzeQueryTextSentiment: Optional. Instructs the service to perform
	// sentiment analysis on
	// `query_text`. If not provided, sentiment analysis is not performed
	// on
	// `query_text`.
	AnalyzeQueryTextSentiment bool `json:"analyzeQueryTextSentiment,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "AnalyzeQueryTextSentiment") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g.
	// "AnalyzeQueryTextSentiment") to include in API requests with the JSON
	// null value. By default, fields with empty values are omitted from API
	// requests. However, any field with an empty value appearing in
	// NullFields will be sent to the server as null. It is an error if a
	// field in this list has a non-empty value. This may be used to include
	// null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1SentimentAnalysisRequestConfig) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1SentimentAnalysisRequestConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1SentimentAnalysisResult: The result of
// sentiment analysis as configured
// by
// `sentiment_analysis_request_config`.
type GoogleCloudDialogflowV2beta1SentimentAnalysisResult struct {
	// QueryTextSentiment: The sentiment analysis result for `query_text`.
	QueryTextSentiment *GoogleCloudDialogflowV2beta1Sentiment `json:"queryTextSentiment,omitempty"`

	// ForceSendFields is a list of field names (e.g. "QueryTextSentiment")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "QueryTextSentiment") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1SentimentAnalysisResult) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1SentimentAnalysisResult
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1SessionEntityType: Represents a session
// entity type.
//
// Extends or replaces a developer entity type at the user session level
// (we
// refer to the entity types defined at the agent level as "developer
// entity
// types").
//
// Note: session entity types apply to all queries, regardless of the
// language.
type GoogleCloudDialogflowV2beta1SessionEntityType struct {
	// Entities: Required. The collection of entities associated with this
	// session entity
	// type.
	Entities []*GoogleCloudDialogflowV2beta1EntityTypeEntity `json:"entities,omitempty"`

	// EntityOverrideMode: Required. Indicates whether the additional data
	// should override or
	// supplement the developer entity type definition.
	//
	// Possible values:
	//   "ENTITY_OVERRIDE_MODE_UNSPECIFIED" - Not specified. This value
	// should be never used.
	//   "ENTITY_OVERRIDE_MODE_OVERRIDE" - The collection of session
	// entities overrides the collection of entities
	// in the corresponding developer entity type.
	//   "ENTITY_OVERRIDE_MODE_SUPPLEMENT" - The collection of session
	// entities extends the collection of entities in
	// the corresponding developer entity type.
	// Calls to `ListSessionEntityTypes`,
	// `GetSessionEntityType`,
	// `CreateSessionEntityType` and `UpdateSessionEntityType` return the
	// full
	// collection of entities from the developer entity type in the
	// agent's
	// default language and the session entity type.
	EntityOverrideMode string `json:"entityOverrideMode,omitempty"`

	// Name: Required. The unique identifier of this session entity type.
	// Format:
	// `projects/<Project ID>/agent/sessions/<Session
	// ID>/entityTypes/<Entity Type
	// Display Name>`, or
	// `projects/<Project ID>/agent/environments/<Environment
	// ID>/users/<User
	// ID>/sessions/<Session ID>/entityTypes/<Entity Type Display Name>`.
	// If `Environment ID` is not specified, we assume default
	// 'draft'
	// environment. If `User ID` is not specified, we assume default '-'
	// user.
	//
	// `<Entity Type Display Name>` must be the display name of an existing
	// entity
	// type in the same agent that will be overridden or supplemented.
	Name string `json:"name,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Entities") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Entities") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1SessionEntityType) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1SessionEntityType
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1StreamingAnalyzeContentRequest: The
// top-level message sent by the client to the
// `StreamingAnalyzeContent`
// method.
//
// Multiple request messages must be sent in the following order:
//
// 1.  The first message must contain `participant` and `config` fields.
// To
//     receive an audio response, the first message must also contain
// the
//     `reply_audio_config` field. The first message must not contain
// `input`.
//
// 2.  All subsequent messages must contain only input data.
// Specifically:
//     - If the `config` in the first message was set to `audio_config`,
// then
//       all subsequent messages must contain only `input_audio`. It is
// a good
//       practice to split the input audio into short chunks and deliver
// each
//       chunk in a separate message.
//     - If the `config` in the first message was set to `text_config`,
// then
//       the second message must contain only `input_text`. Moreover,
// the
//       `input_text` field can be only sent once.
//     After all input is delivered, the client must half-close, or
// abort the
//     request stream.
type GoogleCloudDialogflowV2beta1StreamingAnalyzeContentRequest struct {
	// AudioConfig: Instructs the speech recognizer how to process the
	// speech audio.
	AudioConfig *GoogleCloudDialogflowV2beta1InputAudioConfig `json:"audioConfig,omitempty"`

	// InputAudio: The input audio content to be recognized. Must be sent if
	// `audio_config`
	// is set in the first message. The complete audio over all
	// streaming
	// messages must not exceed 1 minute.
	InputAudio string `json:"inputAudio,omitempty"`

	// InputText: The UTF-8 encoded natural language text to be processed.
	// Must be sent if
	// `text_config` is set in the first message. Text length must not
	// exceed
	// 256 bytes. The `input_text` field can be only sent once.
	InputText string `json:"inputText,omitempty"`

	// ReplyAudioConfig: Optional. Instructs the speech synthesizer how to
	// generate the output
	// audio.
	ReplyAudioConfig *GoogleCloudDialogflowV2beta1OutputAudioConfig `json:"replyAudioConfig,omitempty"`

	// TextConfig: The natural language text to be processed.
	TextConfig *GoogleCloudDialogflowV2beta1InputTextConfig `json:"textConfig,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AudioConfig") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AudioConfig") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1StreamingAnalyzeContentRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1StreamingAnalyzeContentRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1StreamingAnalyzeContentResponse: The
// top-level message returned from the `StreamingAnalyzeContent`
// method.
//
// Multiple response messages can be returned in order:
//
// 1.  If the input was set to streaming audio, the first one or more
// messages
//     contain `recognition_result`. Each `recognition_result`
// represents a more
//     complete transcript of what the user said. The last
// `recognition_result`
//     has `is_final` set to `true`.
//
// 2.  The next message contains `reply_text` and optionally
// `reply_audio`
//     returned by an agent. This message may also contain
//     `automated_agent_reply`.
type GoogleCloudDialogflowV2beta1StreamingAnalyzeContentResponse struct {
	// AutomatedAgentReply: Optional. Only set if a Dialogflow automated
	// agent has responded.
	AutomatedAgentReply *GoogleCloudDialogflowV2beta1AutomatedAgentReply `json:"automatedAgentReply,omitempty"`

	// RecognitionResult: The result of speech recognition.
	RecognitionResult *GoogleCloudDialogflowV2beta1StreamingRecognitionResult `json:"recognitionResult,omitempty"`

	// ReplyAudio: Optional. The audio data bytes encoded as specified in
	// the request.
	// This field is set if:
	//  - The `reply_audio_config` field is specified in the request.
	//  - The automated agent, which this output comes from, responded with
	// audio.
	//    In such case, the `reply_audio.config` field contains settings
	// used to
	//    synthesize the speech.
	ReplyAudio *GoogleCloudDialogflowV2beta1OutputAudio `json:"replyAudio,omitempty"`

	// ReplyText: Optional. The output text content.
	// This field is set if an automated agent responded with a text for the
	// user.
	ReplyText string `json:"replyText,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AutomatedAgentReply")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AutomatedAgentReply") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1StreamingAnalyzeContentResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1StreamingAnalyzeContentResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1StreamingRecognitionResult: Contains a
// speech recognition result corresponding to a portion of the
// audio
// that is currently being processed or an indication that this is the
// end
// of the single requested utterance.
//
// Example:
//
// 1.  transcript: "tube"
//
// 2.  transcript: "to be a"
//
// 3.  transcript: "to be"
//
// 4.  transcript: "to be or not to be"
//     is_final: true
//
// 5.  transcript: " that's"
//
// 6.  transcript: " that is"
//
// 7.  recognition_event_type:
// `RECOGNITION_EVENT_END_OF_SINGLE_UTTERANCE`
//
// 8.  transcript: " that is the question"
//     is_final: true
//
// Only two of the responses contain final results (#4 and #8 indicated
// by
// `is_final: true`). Concatenating these generates the full transcript:
// "to be
// or not to be that is the question".
//
// In each response we populate:
//
// *  for `MESSAGE_TYPE_TRANSCRIPT`: `transcript` and possibly
// `is_final`.
//
// *  for `MESSAGE_TYPE_END_OF_SINGLE_UTTERANCE`: only `event_type`.
type GoogleCloudDialogflowV2beta1StreamingRecognitionResult struct {
	// Confidence: The Speech confidence between 0.0 and 1.0 for the current
	// portion of audio.
	// A higher number indicates an estimated greater likelihood that
	// the
	// recognized words are correct. The default of 0.0 is a sentinel
	// value
	// indicating that confidence was not set.
	//
	// This field is typically only provided if `is_final` is true and you
	// should
	// not rely on it being accurate or even set.
	Confidence float64 `json:"confidence,omitempty"`

	// IsFinal: The default of 0.0 is a sentinel value indicating
	// `confidence` was not set.
	// If `false`, the `StreamingRecognitionResult` represents an
	// interim result that may change. If `true`, the recognizer will not
	// return
	// any further hypotheses about this piece of the audio. May only be
	// populated
	// for `event_type` = `RECOGNITION_EVENT_TRANSCRIPT`.
	IsFinal bool `json:"isFinal,omitempty"`

	// MessageType: Type of the result message.
	//
	// Possible values:
	//   "MESSAGE_TYPE_UNSPECIFIED" - Not specified. Should never be used.
	//   "TRANSCRIPT" - Message contains a (possibly partial) transcript.
	//   "END_OF_SINGLE_UTTERANCE" - Event indicates that the server has
	// detected the end of the user's speech
	// utterance and expects no additional speech. Therefore, the server
	// will
	// not process additional audio (although it may subsequently
	// return
	// additional results). The client should stop sending additional
	// audio
	// data, half-close the gRPC connection, and wait for any additional
	// results
	// until the server closes the gRPC connection. This message is only
	// sent if
	// `single_utterance` was set to `true`, and is not used otherwise.
	MessageType string `json:"messageType,omitempty"`

	// Transcript: Transcript text representing the words that the user
	// spoke.
	// Populated if and only if `event_type` =
	// `RECOGNITION_EVENT_TRANSCRIPT`.
	Transcript string `json:"transcript,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Confidence") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Confidence") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1StreamingRecognitionResult) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1StreamingRecognitionResult
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudDialogflowV2beta1StreamingRecognitionResult) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudDialogflowV2beta1StreamingRecognitionResult
	var s1 struct {
		Confidence gensupport.JSONFloat64 `json:"confidence"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Confidence = float64(s1.Confidence)
	return nil
}

// GoogleCloudDialogflowV2beta1Suggestion: Represents a suggestion for a
// human agent.
type GoogleCloudDialogflowV2beta1Suggestion struct {
	// Articles: Output only. Articles ordered by score in descending order.
	Articles []*GoogleCloudDialogflowV2beta1SuggestionArticle `json:"articles,omitempty"`

	// CreateTime: Output only. The time the suggestion was created.
	CreateTime string `json:"createTime,omitempty"`

	// FaqAnswers: Optional. Answers extracted from FAQ documents.
	FaqAnswers []*GoogleCloudDialogflowV2beta1SuggestionFaqAnswer `json:"faqAnswers,omitempty"`

	// Name: Output only. The name of this suggestion.
	// Format:
	// `projects/<Project
	// ID>/conversations/<Conversation
	// ID>/participants/*/suggestions/<Sugges
	// tion ID>`.
	Name string `json:"name,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Articles") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Articles") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1Suggestion) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1Suggestion
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1SuggestionArticle: Represents suggested
// article.
type GoogleCloudDialogflowV2beta1SuggestionArticle struct {
	// Metadata: Output only. A map that contains metadata about the answer
	// and the
	// document from which it originates.
	Metadata map[string]string `json:"metadata,omitempty"`

	// Snippets: Output only. Article snippets.
	Snippets []string `json:"snippets,omitempty"`

	// Title: Output only. The article title.
	Title string `json:"title,omitempty"`

	// Uri: Output only. The article URI.
	Uri string `json:"uri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Metadata") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Metadata") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1SuggestionArticle) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1SuggestionArticle
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1SuggestionFaqAnswer: Represents suggested
// answer from "frequently asked questions".
type GoogleCloudDialogflowV2beta1SuggestionFaqAnswer struct {
	// Answer: Output only. The piece of text from the `source` knowledge
	// base document.
	Answer string `json:"answer,omitempty"`

	// Confidence: The system's confidence score that this Knowledge answer
	// is a good match
	// for this conversational query, range from 0.0 (completely
	// uncertain)
	// to 1.0 (completely certain).
	Confidence float64 `json:"confidence,omitempty"`

	// Metadata: Output only. A map that contains metadata about the answer
	// and the
	// document from which it originates.
	Metadata map[string]string `json:"metadata,omitempty"`

	// Question: Output only. The corresponding FAQ question.
	Question string `json:"question,omitempty"`

	// Source: Output only. Indicates which Knowledge Document this answer
	// was extracted
	// from.
	// Format: `projects/<Project ID>/agent/knowledgeBases/<Knowledge
	// Base
	// ID>/documents/<Document ID>`.
	Source string `json:"source,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Answer") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Answer") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1SuggestionFaqAnswer) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1SuggestionFaqAnswer
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudDialogflowV2beta1SuggestionFaqAnswer) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudDialogflowV2beta1SuggestionFaqAnswer
	var s1 struct {
		Confidence gensupport.JSONFloat64 `json:"confidence"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Confidence = float64(s1.Confidence)
	return nil
}

// GoogleCloudDialogflowV2beta1SynthesizeSpeechConfig: Configuration of
// how speech should be synthesized.
type GoogleCloudDialogflowV2beta1SynthesizeSpeechConfig struct {
	// EffectsProfileId: Optional. An identifier which selects 'audio
	// effects' profiles that are
	// applied on (post synthesized) text to speech. Effects are applied on
	// top of
	// each other in the order they are given.
	EffectsProfileId []string `json:"effectsProfileId,omitempty"`

	// Pitch: Optional. Speaking pitch, in the range [-20.0, 20.0]. 20 means
	// increase 20
	// semitones from the original pitch. -20 means decrease 20 semitones
	// from the
	// original pitch.
	Pitch float64 `json:"pitch,omitempty"`

	// SpeakingRate: Optional. Speaking rate/speed, in the range [0.25,
	// 4.0]. 1.0 is the normal
	// native speed supported by the specific voice. 2.0 is twice as fast,
	// and
	// 0.5 is half as fast. If unset(0.0), defaults to the native 1.0 speed.
	// Any
	// other values < 0.25 or > 4.0 will return an error.
	SpeakingRate float64 `json:"speakingRate,omitempty"`

	// Voice: Optional. The desired voice of the synthesized audio.
	Voice *GoogleCloudDialogflowV2beta1VoiceSelectionParams `json:"voice,omitempty"`

	// VolumeGainDb: Optional. Volume gain (in dB) of the normal native
	// volume supported by the
	// specific voice, in the range [-96.0, 16.0]. If unset, or set to a
	// value of
	// 0.0 (dB), will play at normal native signal amplitude. A value of
	// -6.0 (dB)
	// will play at approximately half the amplitude of the normal native
	// signal
	// amplitude. A value of +6.0 (dB) will play at approximately twice
	// the
	// amplitude of the normal native signal amplitude. We strongly
	// recommend not
	// to exceed +10 (dB) as there's usually no effective increase in
	// loudness for
	// any value greater than that.
	VolumeGainDb float64 `json:"volumeGainDb,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EffectsProfileId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EffectsProfileId") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1SynthesizeSpeechConfig) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1SynthesizeSpeechConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudDialogflowV2beta1SynthesizeSpeechConfig) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudDialogflowV2beta1SynthesizeSpeechConfig
	var s1 struct {
		Pitch        gensupport.JSONFloat64 `json:"pitch"`
		SpeakingRate gensupport.JSONFloat64 `json:"speakingRate"`
		VolumeGainDb gensupport.JSONFloat64 `json:"volumeGainDb"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Pitch = float64(s1.Pitch)
	s.SpeakingRate = float64(s1.SpeakingRate)
	s.VolumeGainDb = float64(s1.VolumeGainDb)
	return nil
}

// GoogleCloudDialogflowV2beta1TextInput: Represents the natural
// language text to be processed.
type GoogleCloudDialogflowV2beta1TextInput struct {
	// LanguageCode: Required. The language of this conversational query.
	// See [Language
	// Support](https://dialogflow.com/docs/languages) for a list of
	// the
	// currently supported language codes. Note that queries in the same
	// session
	// do not necessarily need to specify the same language.
	LanguageCode string `json:"languageCode,omitempty"`

	// Text: Required. The UTF-8 encoded natural language text to be
	// processed.
	// Text length must not exceed 256 bytes.
	Text string `json:"text,omitempty"`

	// ForceSendFields is a list of field names (e.g. "LanguageCode") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "LanguageCode") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1TextInput) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1TextInput
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1TrainAgentRequest: The request message
// for Agents.TrainAgent.
type GoogleCloudDialogflowV2beta1TrainAgentRequest struct {
}

// GoogleCloudDialogflowV2beta1UndeletePhoneNumberRequest: The request
// message for PhoneNumbers.UndeletePhoneNumber.
type GoogleCloudDialogflowV2beta1UndeletePhoneNumberRequest struct {
}

// GoogleCloudDialogflowV2beta1VoiceSelectionParams: Description of
// which voice to use for speech synthesis.
type GoogleCloudDialogflowV2beta1VoiceSelectionParams struct {
	// Name: Optional. The name of the voice. If not set, the service will
	// choose a
	// voice based on the other parameters such as language_code and gender.
	Name string `json:"name,omitempty"`

	// SsmlGender: Optional. The preferred gender of the voice. If not set,
	// the service will
	// choose a voice based on the other parameters such as language_code
	// and
	// name. Note that this is only a preference, not requirement. If
	// a
	// voice of the appropriate gender is not available, the synthesizer
	// should
	// substitute a voice with a different gender rather than failing the
	// request.
	//
	// Possible values:
	//   "SSML_VOICE_GENDER_UNSPECIFIED" - An unspecified gender, which
	// means that the client doesn't care which
	// gender the selected voice will have.
	//   "SSML_VOICE_GENDER_MALE" - A male voice.
	//   "SSML_VOICE_GENDER_FEMALE" - A female voice.
	//   "SSML_VOICE_GENDER_NEUTRAL" - A gender-neutral voice.
	SsmlGender string `json:"ssmlGender,omitempty"`

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

func (s *GoogleCloudDialogflowV2beta1VoiceSelectionParams) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1VoiceSelectionParams
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1WebhookRequest: The request message for a
// webhook call.
type GoogleCloudDialogflowV2beta1WebhookRequest struct {
	// AlternativeQueryResults: Alternative query results from
	// KnowledgeService.
	AlternativeQueryResults []*GoogleCloudDialogflowV2beta1QueryResult `json:"alternativeQueryResults,omitempty"`

	// OriginalDetectIntentRequest: Optional. The contents of the original
	// request that was passed to
	// `[Streaming]DetectIntent` call.
	OriginalDetectIntentRequest *GoogleCloudDialogflowV2beta1OriginalDetectIntentRequest `json:"originalDetectIntentRequest,omitempty"`

	// QueryResult: The result of the conversational query or event
	// processing. Contains the
	// same value as `[Streaming]DetectIntentResponse.query_result`.
	QueryResult *GoogleCloudDialogflowV2beta1QueryResult `json:"queryResult,omitempty"`

	// ResponseId: The unique identifier of the response. Contains the same
	// value as
	// `[Streaming]DetectIntentResponse.response_id`.
	ResponseId string `json:"responseId,omitempty"`

	// Session: The unique identifier of detectIntent request session.
	// Can be used to identify end-user inside webhook
	// implementation.
	// Format: `projects/<Project ID>/agent/sessions/<Session ID>`.
	Session string `json:"session,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "AlternativeQueryResults") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AlternativeQueryResults")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1WebhookRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1WebhookRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudDialogflowV2beta1WebhookResponse: The response message for
// a webhook call.
type GoogleCloudDialogflowV2beta1WebhookResponse struct {
	// EndInteraction: Optional. Indicates that this intent ends an
	// interaction. Some integrations
	// (e.g., Actions on Google or Dialogflow phone gateway) use this
	// information
	// to close interaction with an end user. Default is false.
	EndInteraction bool `json:"endInteraction,omitempty"`

	// FollowupEventInput: Optional. Makes the platform immediately invoke
	// another `DetectIntent` call
	// internally with the specified event as input.
	FollowupEventInput *GoogleCloudDialogflowV2beta1EventInput `json:"followupEventInput,omitempty"`

	// FulfillmentMessages: Optional. The collection of rich messages to
	// present to the user. This
	// value is passed directly to `QueryResult.fulfillment_messages`.
	FulfillmentMessages []*GoogleCloudDialogflowV2beta1IntentMessage `json:"fulfillmentMessages,omitempty"`

	// FulfillmentText: Optional. The text to be shown on the screen. This
	// value is passed directly
	// to `QueryResult.fulfillment_text`.
	FulfillmentText string `json:"fulfillmentText,omitempty"`

	// OutputContexts: Optional. The collection of output contexts. This
	// value is passed directly
	// to `QueryResult.output_contexts`.
	OutputContexts []*GoogleCloudDialogflowV2beta1Context `json:"outputContexts,omitempty"`

	// Payload: Optional. This value is passed directly to
	// `QueryResult.webhook_payload`.
	// See the related `fulfillment_messages[i].payload field`, which may be
	// used
	// as an alternative to this field.
	//
	// This field can be used for Actions on Google responses.
	// It should have a structure similar to the JSON message shown here.
	// For more
	// information, see
	// [Actions on Google
	// Webhook
	// Format](https://developers.google.com/actions/dialogflow/webho
	// ok)
	// <pre>{
	//   "google": {
	//     "expectUserResponse": true,
	//     "richResponse": {
	//       "items": [
	//         {
	//           "simpleResponse": {
	//             "textToSpeech": "this is a simple response"
	//           }
	//         }
	//       ]
	//     }
	//   }
	// }</pre>
	Payload googleapi.RawMessage `json:"payload,omitempty"`

	// Source: Optional. This value is passed directly to
	// `QueryResult.webhook_source`.
	Source string `json:"source,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EndInteraction") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EndInteraction") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudDialogflowV2beta1WebhookResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudDialogflowV2beta1WebhookResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleLongrunningOperation: This resource represents a long-running
// operation that is the result of a
// network API call.
type GoogleLongrunningOperation struct {
	// Done: If the value is `false`, it means the operation is still in
	// progress.
	// If `true`, the operation is completed, and either `error` or
	// `response` is
	// available.
	Done bool `json:"done,omitempty"`

	// Error: The error result of the operation in case of failure or
	// cancellation.
	Error *GoogleRpcStatus `json:"error,omitempty"`

	// Metadata: Service-specific metadata associated with the operation.
	// It typically
	// contains progress information and common metadata such as create
	// time.
	// Some services might not provide such metadata.  Any method that
	// returns a
	// long-running operation should document the metadata type, if any.
	Metadata googleapi.RawMessage `json:"metadata,omitempty"`

	// Name: The server-assigned name, which is only unique within the same
	// service that
	// originally returns it. If you use the default HTTP mapping,
	// the
	// `name` should have the format of `operations/some/unique/name`.
	Name string `json:"name,omitempty"`

	// Response: The normal response of the operation in case of success.
	// If the original
	// method returns no data on success, such as `Delete`, the response
	// is
	// `google.protobuf.Empty`.  If the original method is
	// standard
	// `Get`/`Create`/`Update`, the response should be the resource.  For
	// other
	// methods, the response should have the type `XxxResponse`, where
	// `Xxx`
	// is the original method name.  For example, if the original method
	// name
	// is `TakeSnapshot()`, the inferred response type
	// is
	// `TakeSnapshotResponse`.
	Response googleapi.RawMessage `json:"response,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Done") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Done") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleLongrunningOperation) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleLongrunningOperation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleProtobufEmpty: A generic empty message that you can re-use to
// avoid defining duplicated
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
type GoogleProtobufEmpty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// GoogleRpcStatus: The `Status` type defines a logical error model that
// is suitable for different
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
// in the package `google.rpc` that can be used for common error
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
//     have a `Status` message for error reporting.
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
type GoogleRpcStatus struct {
	// Code: The status code, which should be an enum value of
	// google.rpc.Code.
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details.  There is a
	// common set of
	// message types for APIs to use.
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

func (s *GoogleRpcStatus) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleRpcStatus
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleTypeLatLng: An object representing a latitude/longitude pair.
// This is expressed as a pair
// of doubles representing degrees latitude and degrees longitude.
// Unless
// specified otherwise, this must conform to the
// <a
// href="http://www.unoosa.org/pdf/icg/2012/template/WGS_84.pdf">WGS84
// st
// andard</a>. Values must be within normalized ranges.
type GoogleTypeLatLng struct {
	// Latitude: The latitude in degrees. It must be in the range [-90.0,
	// +90.0].
	Latitude float64 `json:"latitude,omitempty"`

	// Longitude: The longitude in degrees. It must be in the range [-180.0,
	// +180.0].
	Longitude float64 `json:"longitude,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Latitude") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Latitude") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleTypeLatLng) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleTypeLatLng
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleTypeLatLng) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleTypeLatLng
	var s1 struct {
		Latitude  gensupport.JSONFloat64 `json:"latitude"`
		Longitude gensupport.JSONFloat64 `json:"longitude"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Latitude = float64(s1.Latitude)
	s.Longitude = float64(s1.Longitude)
	return nil
}

// method id "dialogflow.projects.getAgent":

type ProjectsGetAgentCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// GetAgent: Retrieves the specified agent.
func (r *ProjectsService) GetAgent(parent string) *ProjectsGetAgentCall {
	c := &ProjectsGetAgentCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsGetAgentCall) Fields(s ...googleapi.Field) *ProjectsGetAgentCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsGetAgentCall) IfNoneMatch(entityTag string) *ProjectsGetAgentCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsGetAgentCall) Context(ctx context.Context) *ProjectsGetAgentCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsGetAgentCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsGetAgentCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/agent")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.getAgent" call.
// Exactly one of *GoogleCloudDialogflowV2beta1Agent or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleCloudDialogflowV2beta1Agent.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsGetAgentCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1Agent, error) {
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
	ret := &GoogleCloudDialogflowV2beta1Agent{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the specified agent.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.getAgent",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The project that the agent to fetch is associated with.\nFormat: `projects/\u003cProject ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/agent",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Agent"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.export":

type ProjectsAgentExportCall struct {
	s                                              *Service
	parent                                         string
	googleclouddialogflowv2beta1exportagentrequest *GoogleCloudDialogflowV2beta1ExportAgentRequest
	urlParams_                                     gensupport.URLParams
	ctx_                                           context.Context
	header_                                        http.Header
}

// Export: Exports the specified agent to a ZIP file.
//
//
// Operation <response: ExportAgentResponse,
//            metadata: google.protobuf.Struct>
func (r *ProjectsAgentService) Export(parent string, googleclouddialogflowv2beta1exportagentrequest *GoogleCloudDialogflowV2beta1ExportAgentRequest) *ProjectsAgentExportCall {
	c := &ProjectsAgentExportCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.googleclouddialogflowv2beta1exportagentrequest = googleclouddialogflowv2beta1exportagentrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentExportCall) Fields(s ...googleapi.Field) *ProjectsAgentExportCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentExportCall) Context(ctx context.Context) *ProjectsAgentExportCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentExportCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentExportCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1exportagentrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/agent:export")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.export" call.
// Exactly one of *GoogleLongrunningOperation or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *GoogleLongrunningOperation.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentExportCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunningOperation, error) {
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
	ret := &GoogleLongrunningOperation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Exports the specified agent to a ZIP file.\n\n\nOperation \u003cresponse: ExportAgentResponse,\n           metadata: google.protobuf.Struct\u003e",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent:export",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.export",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The project that the agent to export is associated with.\nFormat: `projects/\u003cProject ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/agent:export",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1ExportAgentRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleLongrunningOperation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.import":

type ProjectsAgentImportCall struct {
	s                                              *Service
	parent                                         string
	googleclouddialogflowv2beta1importagentrequest *GoogleCloudDialogflowV2beta1ImportAgentRequest
	urlParams_                                     gensupport.URLParams
	ctx_                                           context.Context
	header_                                        http.Header
}

// Import: Imports the specified agent from a ZIP file.
//
// Uploads new intents and entity types without deleting the existing
// ones.
// Intents and entity types with the same name are replaced with the
// new
// versions from ImportAgentRequest.
//
//
// Operation <response: google.protobuf.Empty,
//            metadata: google.protobuf.Struct>
func (r *ProjectsAgentService) Import(parent string, googleclouddialogflowv2beta1importagentrequest *GoogleCloudDialogflowV2beta1ImportAgentRequest) *ProjectsAgentImportCall {
	c := &ProjectsAgentImportCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.googleclouddialogflowv2beta1importagentrequest = googleclouddialogflowv2beta1importagentrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentImportCall) Fields(s ...googleapi.Field) *ProjectsAgentImportCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentImportCall) Context(ctx context.Context) *ProjectsAgentImportCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentImportCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentImportCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1importagentrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/agent:import")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.import" call.
// Exactly one of *GoogleLongrunningOperation or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *GoogleLongrunningOperation.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentImportCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunningOperation, error) {
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
	ret := &GoogleLongrunningOperation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Imports the specified agent from a ZIP file.\n\nUploads new intents and entity types without deleting the existing ones.\nIntents and entity types with the same name are replaced with the new\nversions from ImportAgentRequest.\n\n\nOperation \u003cresponse: google.protobuf.Empty,\n           metadata: google.protobuf.Struct\u003e",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent:import",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.import",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The project that the agent to import is associated with.\nFormat: `projects/\u003cProject ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/agent:import",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1ImportAgentRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleLongrunningOperation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.restore":

type ProjectsAgentRestoreCall struct {
	s                                               *Service
	parent                                          string
	googleclouddialogflowv2beta1restoreagentrequest *GoogleCloudDialogflowV2beta1RestoreAgentRequest
	urlParams_                                      gensupport.URLParams
	ctx_                                            context.Context
	header_                                         http.Header
}

// Restore: Restores the specified agent from a ZIP file.
//
// Replaces the current agent version with a new one. All the intents
// and
// entity types in the older version are deleted.
//
//
// Operation <response: google.protobuf.Empty,
//            metadata: google.protobuf.Struct>
func (r *ProjectsAgentService) Restore(parent string, googleclouddialogflowv2beta1restoreagentrequest *GoogleCloudDialogflowV2beta1RestoreAgentRequest) *ProjectsAgentRestoreCall {
	c := &ProjectsAgentRestoreCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.googleclouddialogflowv2beta1restoreagentrequest = googleclouddialogflowv2beta1restoreagentrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentRestoreCall) Fields(s ...googleapi.Field) *ProjectsAgentRestoreCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentRestoreCall) Context(ctx context.Context) *ProjectsAgentRestoreCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentRestoreCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentRestoreCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1restoreagentrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/agent:restore")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.restore" call.
// Exactly one of *GoogleLongrunningOperation or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *GoogleLongrunningOperation.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentRestoreCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunningOperation, error) {
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
	ret := &GoogleLongrunningOperation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Restores the specified agent from a ZIP file.\n\nReplaces the current agent version with a new one. All the intents and\nentity types in the older version are deleted.\n\n\nOperation \u003cresponse: google.protobuf.Empty,\n           metadata: google.protobuf.Struct\u003e",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent:restore",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.restore",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The project that the agent to restore is associated with.\nFormat: `projects/\u003cProject ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/agent:restore",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1RestoreAgentRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleLongrunningOperation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.search":

type ProjectsAgentSearchCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Search: Returns the list of agents.
//
// Since there is at most one conversational agent per project, this
// method is
// useful primarily for listing all agents across projects the caller
// has
// access to. One can achieve that with a wildcard project collection id
// "-".
// Refer to
// [List
// Sub-Collections](https://cloud.google.com/apis/design/design_pat
// terns#list_sub-collections).
func (r *ProjectsAgentService) Search(parent string) *ProjectsAgentSearchCall {
	c := &ProjectsAgentSearchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of items to return in a single page. By
// default 100 and at most 1000.
func (c *ProjectsAgentSearchCall) PageSize(pageSize int64) *ProjectsAgentSearchCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous list request.
func (c *ProjectsAgentSearchCall) PageToken(pageToken string) *ProjectsAgentSearchCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentSearchCall) Fields(s ...googleapi.Field) *ProjectsAgentSearchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsAgentSearchCall) IfNoneMatch(entityTag string) *ProjectsAgentSearchCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentSearchCall) Context(ctx context.Context) *ProjectsAgentSearchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentSearchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentSearchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/agent:search")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.search" call.
// Exactly one of *GoogleCloudDialogflowV2beta1SearchAgentsResponse or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1SearchAgentsResponse.ServerResponse.Heade
// r or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsAgentSearchCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1SearchAgentsResponse, error) {
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
	ret := &GoogleCloudDialogflowV2beta1SearchAgentsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the list of agents.\n\nSince there is at most one conversational agent per project, this method is\nuseful primarily for listing all agents across projects the caller has\naccess to. One can achieve that with a wildcard project collection id \"-\".\nRefer to [List\nSub-Collections](https://cloud.google.com/apis/design/design_patterns#list_sub-collections).",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent:search",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.agent.search",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Optional. The maximum number of items to return in a single page. By\ndefault 100 and at most 1000.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. The next_page_token value returned from a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The project to list agents from.\nFormat: `projects/\u003cProject ID or '-'\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/agent:search",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1SearchAgentsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsAgentSearchCall) Pages(ctx context.Context, f func(*GoogleCloudDialogflowV2beta1SearchAgentsResponse) error) error {
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

// method id "dialogflow.projects.agent.train":

type ProjectsAgentTrainCall struct {
	s                                             *Service
	parent                                        string
	googleclouddialogflowv2beta1trainagentrequest *GoogleCloudDialogflowV2beta1TrainAgentRequest
	urlParams_                                    gensupport.URLParams
	ctx_                                          context.Context
	header_                                       http.Header
}

// Train: Trains the specified agent.
//
//
// Operation <response: google.protobuf.Empty,
//            metadata: google.protobuf.Struct>
func (r *ProjectsAgentService) Train(parent string, googleclouddialogflowv2beta1trainagentrequest *GoogleCloudDialogflowV2beta1TrainAgentRequest) *ProjectsAgentTrainCall {
	c := &ProjectsAgentTrainCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.googleclouddialogflowv2beta1trainagentrequest = googleclouddialogflowv2beta1trainagentrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentTrainCall) Fields(s ...googleapi.Field) *ProjectsAgentTrainCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentTrainCall) Context(ctx context.Context) *ProjectsAgentTrainCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentTrainCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentTrainCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1trainagentrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/agent:train")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.train" call.
// Exactly one of *GoogleLongrunningOperation or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *GoogleLongrunningOperation.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentTrainCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunningOperation, error) {
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
	ret := &GoogleLongrunningOperation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Trains the specified agent.\n\n\nOperation \u003cresponse: google.protobuf.Empty,\n           metadata: google.protobuf.Struct\u003e",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent:train",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.train",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The project that the agent to train is associated with.\nFormat: `projects/\u003cProject ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/agent:train",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1TrainAgentRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleLongrunningOperation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.entityTypes.batchDelete":

type ProjectsAgentEntityTypesBatchDeleteCall struct {
	s                                                         *Service
	parent                                                    string
	googleclouddialogflowv2beta1batchdeleteentitytypesrequest *GoogleCloudDialogflowV2beta1BatchDeleteEntityTypesRequest
	urlParams_                                                gensupport.URLParams
	ctx_                                                      context.Context
	header_                                                   http.Header
}

// BatchDelete: Deletes entity types in the specified agent.
//
// Operation <response: google.protobuf.Empty,
//            metadata: google.protobuf.Struct>
func (r *ProjectsAgentEntityTypesService) BatchDelete(parent string, googleclouddialogflowv2beta1batchdeleteentitytypesrequest *GoogleCloudDialogflowV2beta1BatchDeleteEntityTypesRequest) *ProjectsAgentEntityTypesBatchDeleteCall {
	c := &ProjectsAgentEntityTypesBatchDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.googleclouddialogflowv2beta1batchdeleteentitytypesrequest = googleclouddialogflowv2beta1batchdeleteentitytypesrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEntityTypesBatchDeleteCall) Fields(s ...googleapi.Field) *ProjectsAgentEntityTypesBatchDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEntityTypesBatchDeleteCall) Context(ctx context.Context) *ProjectsAgentEntityTypesBatchDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEntityTypesBatchDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEntityTypesBatchDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1batchdeleteentitytypesrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/entityTypes:batchDelete")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.entityTypes.batchDelete" call.
// Exactly one of *GoogleLongrunningOperation or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *GoogleLongrunningOperation.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentEntityTypesBatchDeleteCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunningOperation, error) {
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
	ret := &GoogleLongrunningOperation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes entity types in the specified agent.\n\nOperation \u003cresponse: google.protobuf.Empty,\n           metadata: google.protobuf.Struct\u003e",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/entityTypes:batchDelete",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.entityTypes.batchDelete",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The name of the agent to delete all entities types for. Format:\n`projects/\u003cProject ID\u003e/agent`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/entityTypes:batchDelete",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1BatchDeleteEntityTypesRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleLongrunningOperation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.entityTypes.batchUpdate":

type ProjectsAgentEntityTypesBatchUpdateCall struct {
	s                                                         *Service
	parent                                                    string
	googleclouddialogflowv2beta1batchupdateentitytypesrequest *GoogleCloudDialogflowV2beta1BatchUpdateEntityTypesRequest
	urlParams_                                                gensupport.URLParams
	ctx_                                                      context.Context
	header_                                                   http.Header
}

// BatchUpdate: Updates/Creates multiple entity types in the specified
// agent.
//
// Operation <response: BatchUpdateEntityTypesResponse,
//            metadata: google.protobuf.Struct>
func (r *ProjectsAgentEntityTypesService) BatchUpdate(parent string, googleclouddialogflowv2beta1batchupdateentitytypesrequest *GoogleCloudDialogflowV2beta1BatchUpdateEntityTypesRequest) *ProjectsAgentEntityTypesBatchUpdateCall {
	c := &ProjectsAgentEntityTypesBatchUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.googleclouddialogflowv2beta1batchupdateentitytypesrequest = googleclouddialogflowv2beta1batchupdateentitytypesrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEntityTypesBatchUpdateCall) Fields(s ...googleapi.Field) *ProjectsAgentEntityTypesBatchUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEntityTypesBatchUpdateCall) Context(ctx context.Context) *ProjectsAgentEntityTypesBatchUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEntityTypesBatchUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEntityTypesBatchUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1batchupdateentitytypesrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/entityTypes:batchUpdate")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.entityTypes.batchUpdate" call.
// Exactly one of *GoogleLongrunningOperation or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *GoogleLongrunningOperation.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentEntityTypesBatchUpdateCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunningOperation, error) {
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
	ret := &GoogleLongrunningOperation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates/Creates multiple entity types in the specified agent.\n\nOperation \u003cresponse: BatchUpdateEntityTypesResponse,\n           metadata: google.protobuf.Struct\u003e",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/entityTypes:batchUpdate",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.entityTypes.batchUpdate",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The name of the agent to update or create entity types in.\nFormat: `projects/\u003cProject ID\u003e/agent`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/entityTypes:batchUpdate",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1BatchUpdateEntityTypesRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleLongrunningOperation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.entityTypes.create":

type ProjectsAgentEntityTypesCreateCall struct {
	s                                      *Service
	parent                                 string
	googleclouddialogflowv2beta1entitytype *GoogleCloudDialogflowV2beta1EntityType
	urlParams_                             gensupport.URLParams
	ctx_                                   context.Context
	header_                                http.Header
}

// Create: Creates an entity type in the specified agent.
func (r *ProjectsAgentEntityTypesService) Create(parent string, googleclouddialogflowv2beta1entitytype *GoogleCloudDialogflowV2beta1EntityType) *ProjectsAgentEntityTypesCreateCall {
	c := &ProjectsAgentEntityTypesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.googleclouddialogflowv2beta1entitytype = googleclouddialogflowv2beta1entitytype
	return c
}

// LanguageCode sets the optional parameter "languageCode": The language
// of entity synonyms defined in `entity_type`. If not
// specified, the agent's default language is used.
// [More than a
// dozen
// languages](https://dialogflow.com/docs/reference/language) are
// supported.
// Note: languages must be enabled in the agent, before they can be
// used.
func (c *ProjectsAgentEntityTypesCreateCall) LanguageCode(languageCode string) *ProjectsAgentEntityTypesCreateCall {
	c.urlParams_.Set("languageCode", languageCode)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEntityTypesCreateCall) Fields(s ...googleapi.Field) *ProjectsAgentEntityTypesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEntityTypesCreateCall) Context(ctx context.Context) *ProjectsAgentEntityTypesCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEntityTypesCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEntityTypesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1entitytype)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/entityTypes")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.entityTypes.create" call.
// Exactly one of *GoogleCloudDialogflowV2beta1EntityType or error will
// be non-nil. Any non-2xx status code is an error. Response headers are
// in either
// *GoogleCloudDialogflowV2beta1EntityType.ServerResponse.Header or (if
// a response was returned at all) in error.(*googleapi.Error).Header.
// Use googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentEntityTypesCreateCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1EntityType, error) {
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
	ret := &GoogleCloudDialogflowV2beta1EntityType{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates an entity type in the specified agent.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/entityTypes",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.entityTypes.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "languageCode": {
	//       "description": "Optional. The language of entity synonyms defined in `entity_type`. If not\nspecified, the agent's default language is used.\n[More than a dozen\nlanguages](https://dialogflow.com/docs/reference/language) are supported.\nNote: languages must be enabled in the agent, before they can be used.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The agent to create a entity type for.\nFormat: `projects/\u003cProject ID\u003e/agent`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/entityTypes",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1EntityType"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1EntityType"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.entityTypes.delete":

type ProjectsAgentEntityTypesDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes the specified entity type.
func (r *ProjectsAgentEntityTypesService) Delete(name string) *ProjectsAgentEntityTypesDeleteCall {
	c := &ProjectsAgentEntityTypesDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEntityTypesDeleteCall) Fields(s ...googleapi.Field) *ProjectsAgentEntityTypesDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEntityTypesDeleteCall) Context(ctx context.Context) *ProjectsAgentEntityTypesDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEntityTypesDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEntityTypesDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.entityTypes.delete" call.
// Exactly one of *GoogleProtobufEmpty or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleProtobufEmpty.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentEntityTypesDeleteCall) Do(opts ...googleapi.CallOption) (*GoogleProtobufEmpty, error) {
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
	ret := &GoogleProtobufEmpty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the specified entity type.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/entityTypes/{entityTypesId}",
	//   "httpMethod": "DELETE",
	//   "id": "dialogflow.projects.agent.entityTypes.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the entity type to delete.\nFormat: `projects/\u003cProject ID\u003e/agent/entityTypes/\u003cEntityType ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/entityTypes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleProtobufEmpty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.entityTypes.get":

type ProjectsAgentEntityTypesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves the specified entity type.
func (r *ProjectsAgentEntityTypesService) Get(name string) *ProjectsAgentEntityTypesGetCall {
	c := &ProjectsAgentEntityTypesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// LanguageCode sets the optional parameter "languageCode": The language
// to retrieve entity synonyms for. If not specified,
// the agent's default language is used.
// [More than a
// dozen
// languages](https://dialogflow.com/docs/reference/language) are
// supported.
// Note: languages must be enabled in the agent, before they can be
// used.
func (c *ProjectsAgentEntityTypesGetCall) LanguageCode(languageCode string) *ProjectsAgentEntityTypesGetCall {
	c.urlParams_.Set("languageCode", languageCode)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEntityTypesGetCall) Fields(s ...googleapi.Field) *ProjectsAgentEntityTypesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsAgentEntityTypesGetCall) IfNoneMatch(entityTag string) *ProjectsAgentEntityTypesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEntityTypesGetCall) Context(ctx context.Context) *ProjectsAgentEntityTypesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEntityTypesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEntityTypesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.entityTypes.get" call.
// Exactly one of *GoogleCloudDialogflowV2beta1EntityType or error will
// be non-nil. Any non-2xx status code is an error. Response headers are
// in either
// *GoogleCloudDialogflowV2beta1EntityType.ServerResponse.Header or (if
// a response was returned at all) in error.(*googleapi.Error).Header.
// Use googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentEntityTypesGetCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1EntityType, error) {
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
	ret := &GoogleCloudDialogflowV2beta1EntityType{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the specified entity type.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/entityTypes/{entityTypesId}",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.agent.entityTypes.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "languageCode": {
	//       "description": "Optional. The language to retrieve entity synonyms for. If not specified,\nthe agent's default language is used.\n[More than a dozen\nlanguages](https://dialogflow.com/docs/reference/language) are supported.\nNote: languages must be enabled in the agent, before they can be used.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "Required. The name of the entity type.\nFormat: `projects/\u003cProject ID\u003e/agent/entityTypes/\u003cEntityType ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/entityTypes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1EntityType"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.entityTypes.list":

type ProjectsAgentEntityTypesListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Returns the list of all entity types in the specified agent.
func (r *ProjectsAgentEntityTypesService) List(parent string) *ProjectsAgentEntityTypesListCall {
	c := &ProjectsAgentEntityTypesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// LanguageCode sets the optional parameter "languageCode": The language
// to list entity synonyms for. If not specified,
// the agent's default language is used.
// [More than a
// dozen
// languages](https://dialogflow.com/docs/reference/language) are
// supported.
// Note: languages must be enabled in the agent, before they can be
// used.
func (c *ProjectsAgentEntityTypesListCall) LanguageCode(languageCode string) *ProjectsAgentEntityTypesListCall {
	c.urlParams_.Set("languageCode", languageCode)
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of items to return in a single page. By
// default 100 and at most 1000.
func (c *ProjectsAgentEntityTypesListCall) PageSize(pageSize int64) *ProjectsAgentEntityTypesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous list request.
func (c *ProjectsAgentEntityTypesListCall) PageToken(pageToken string) *ProjectsAgentEntityTypesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEntityTypesListCall) Fields(s ...googleapi.Field) *ProjectsAgentEntityTypesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsAgentEntityTypesListCall) IfNoneMatch(entityTag string) *ProjectsAgentEntityTypesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEntityTypesListCall) Context(ctx context.Context) *ProjectsAgentEntityTypesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEntityTypesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEntityTypesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/entityTypes")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.entityTypes.list" call.
// Exactly one of *GoogleCloudDialogflowV2beta1ListEntityTypesResponse
// or error will be non-nil. Any non-2xx status code is an error.
// Response headers are in either
// *GoogleCloudDialogflowV2beta1ListEntityTypesResponse.ServerResponse.He
// ader or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsAgentEntityTypesListCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1ListEntityTypesResponse, error) {
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
	ret := &GoogleCloudDialogflowV2beta1ListEntityTypesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the list of all entity types in the specified agent.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/entityTypes",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.agent.entityTypes.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "languageCode": {
	//       "description": "Optional. The language to list entity synonyms for. If not specified,\nthe agent's default language is used.\n[More than a dozen\nlanguages](https://dialogflow.com/docs/reference/language) are supported.\nNote: languages must be enabled in the agent, before they can be used.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Optional. The maximum number of items to return in a single page. By\ndefault 100 and at most 1000.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. The next_page_token value returned from a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The agent to list all entity types from.\nFormat: `projects/\u003cProject ID\u003e/agent`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/entityTypes",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1ListEntityTypesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsAgentEntityTypesListCall) Pages(ctx context.Context, f func(*GoogleCloudDialogflowV2beta1ListEntityTypesResponse) error) error {
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

// method id "dialogflow.projects.agent.entityTypes.patch":

type ProjectsAgentEntityTypesPatchCall struct {
	s                                      *Service
	nameid                                 string
	googleclouddialogflowv2beta1entitytype *GoogleCloudDialogflowV2beta1EntityType
	urlParams_                             gensupport.URLParams
	ctx_                                   context.Context
	header_                                http.Header
}

// Patch: Updates the specified entity type.
func (r *ProjectsAgentEntityTypesService) Patch(nameid string, googleclouddialogflowv2beta1entitytype *GoogleCloudDialogflowV2beta1EntityType) *ProjectsAgentEntityTypesPatchCall {
	c := &ProjectsAgentEntityTypesPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.nameid = nameid
	c.googleclouddialogflowv2beta1entitytype = googleclouddialogflowv2beta1entitytype
	return c
}

// LanguageCode sets the optional parameter "languageCode": The language
// of entity synonyms defined in `entity_type`. If not
// specified, the agent's default language is used.
// [More than a
// dozen
// languages](https://dialogflow.com/docs/reference/language) are
// supported.
// Note: languages must be enabled in the agent, before they can be
// used.
func (c *ProjectsAgentEntityTypesPatchCall) LanguageCode(languageCode string) *ProjectsAgentEntityTypesPatchCall {
	c.urlParams_.Set("languageCode", languageCode)
	return c
}

// UpdateMask sets the optional parameter "updateMask": The mask to
// control which fields get updated.
func (c *ProjectsAgentEntityTypesPatchCall) UpdateMask(updateMask string) *ProjectsAgentEntityTypesPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEntityTypesPatchCall) Fields(s ...googleapi.Field) *ProjectsAgentEntityTypesPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEntityTypesPatchCall) Context(ctx context.Context) *ProjectsAgentEntityTypesPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEntityTypesPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEntityTypesPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1entitytype)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.nameid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.entityTypes.patch" call.
// Exactly one of *GoogleCloudDialogflowV2beta1EntityType or error will
// be non-nil. Any non-2xx status code is an error. Response headers are
// in either
// *GoogleCloudDialogflowV2beta1EntityType.ServerResponse.Header or (if
// a response was returned at all) in error.(*googleapi.Error).Header.
// Use googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentEntityTypesPatchCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1EntityType, error) {
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
	ret := &GoogleCloudDialogflowV2beta1EntityType{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the specified entity type.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/entityTypes/{entityTypesId}",
	//   "httpMethod": "PATCH",
	//   "id": "dialogflow.projects.agent.entityTypes.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "languageCode": {
	//       "description": "Optional. The language of entity synonyms defined in `entity_type`. If not\nspecified, the agent's default language is used.\n[More than a dozen\nlanguages](https://dialogflow.com/docs/reference/language) are supported.\nNote: languages must be enabled in the agent, before they can be used.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "Required for all methods except `create` (`create` populates the name\nautomatically.\nThe unique identifier of the entity type. Format:\n`projects/\u003cProject ID\u003e/agent/entityTypes/\u003cEntity Type ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/entityTypes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Optional. The mask to control which fields get updated.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1EntityType"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1EntityType"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.entityTypes.entities.batchCreate":

type ProjectsAgentEntityTypesEntitiesBatchCreateCall struct {
	s                                                      *Service
	parent                                                 string
	googleclouddialogflowv2beta1batchcreateentitiesrequest *GoogleCloudDialogflowV2beta1BatchCreateEntitiesRequest
	urlParams_                                             gensupport.URLParams
	ctx_                                                   context.Context
	header_                                                http.Header
}

// BatchCreate: Creates multiple new entities in the specified entity
// type.
//
// Operation <response: google.protobuf.Empty>
func (r *ProjectsAgentEntityTypesEntitiesService) BatchCreate(parent string, googleclouddialogflowv2beta1batchcreateentitiesrequest *GoogleCloudDialogflowV2beta1BatchCreateEntitiesRequest) *ProjectsAgentEntityTypesEntitiesBatchCreateCall {
	c := &ProjectsAgentEntityTypesEntitiesBatchCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.googleclouddialogflowv2beta1batchcreateentitiesrequest = googleclouddialogflowv2beta1batchcreateentitiesrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEntityTypesEntitiesBatchCreateCall) Fields(s ...googleapi.Field) *ProjectsAgentEntityTypesEntitiesBatchCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEntityTypesEntitiesBatchCreateCall) Context(ctx context.Context) *ProjectsAgentEntityTypesEntitiesBatchCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEntityTypesEntitiesBatchCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEntityTypesEntitiesBatchCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1batchcreateentitiesrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/entities:batchCreate")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.entityTypes.entities.batchCreate" call.
// Exactly one of *GoogleLongrunningOperation or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *GoogleLongrunningOperation.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentEntityTypesEntitiesBatchCreateCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunningOperation, error) {
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
	ret := &GoogleLongrunningOperation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates multiple new entities in the specified entity type.\n\nOperation \u003cresponse: google.protobuf.Empty\u003e",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/entityTypes/{entityTypesId}/entities:batchCreate",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.entityTypes.entities.batchCreate",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The name of the entity type to create entities in. Format:\n`projects/\u003cProject ID\u003e/agent/entityTypes/\u003cEntity Type ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/entityTypes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/entities:batchCreate",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1BatchCreateEntitiesRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleLongrunningOperation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.entityTypes.entities.batchDelete":

type ProjectsAgentEntityTypesEntitiesBatchDeleteCall struct {
	s                                                      *Service
	parent                                                 string
	googleclouddialogflowv2beta1batchdeleteentitiesrequest *GoogleCloudDialogflowV2beta1BatchDeleteEntitiesRequest
	urlParams_                                             gensupport.URLParams
	ctx_                                                   context.Context
	header_                                                http.Header
}

// BatchDelete: Deletes entities in the specified entity
// type.
//
// Operation <response: google.protobuf.Empty,
//            metadata: google.protobuf.Struct>
func (r *ProjectsAgentEntityTypesEntitiesService) BatchDelete(parent string, googleclouddialogflowv2beta1batchdeleteentitiesrequest *GoogleCloudDialogflowV2beta1BatchDeleteEntitiesRequest) *ProjectsAgentEntityTypesEntitiesBatchDeleteCall {
	c := &ProjectsAgentEntityTypesEntitiesBatchDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.googleclouddialogflowv2beta1batchdeleteentitiesrequest = googleclouddialogflowv2beta1batchdeleteentitiesrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEntityTypesEntitiesBatchDeleteCall) Fields(s ...googleapi.Field) *ProjectsAgentEntityTypesEntitiesBatchDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEntityTypesEntitiesBatchDeleteCall) Context(ctx context.Context) *ProjectsAgentEntityTypesEntitiesBatchDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEntityTypesEntitiesBatchDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEntityTypesEntitiesBatchDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1batchdeleteentitiesrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/entities:batchDelete")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.entityTypes.entities.batchDelete" call.
// Exactly one of *GoogleLongrunningOperation or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *GoogleLongrunningOperation.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentEntityTypesEntitiesBatchDeleteCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunningOperation, error) {
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
	ret := &GoogleLongrunningOperation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes entities in the specified entity type.\n\nOperation \u003cresponse: google.protobuf.Empty,\n           metadata: google.protobuf.Struct\u003e",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/entityTypes/{entityTypesId}/entities:batchDelete",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.entityTypes.entities.batchDelete",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The name of the entity type to delete entries for. Format:\n`projects/\u003cProject ID\u003e/agent/entityTypes/\u003cEntity Type ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/entityTypes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/entities:batchDelete",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1BatchDeleteEntitiesRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleLongrunningOperation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.entityTypes.entities.batchUpdate":

type ProjectsAgentEntityTypesEntitiesBatchUpdateCall struct {
	s                                                      *Service
	parent                                                 string
	googleclouddialogflowv2beta1batchupdateentitiesrequest *GoogleCloudDialogflowV2beta1BatchUpdateEntitiesRequest
	urlParams_                                             gensupport.URLParams
	ctx_                                                   context.Context
	header_                                                http.Header
}

// BatchUpdate: Updates or creates multiple entities in the specified
// entity type. This
// method does not affect entities in the entity type that aren't
// explicitly
// specified in the request.
//
// Operation <response: google.protobuf.Empty,
//            metadata: google.protobuf.Struct>
func (r *ProjectsAgentEntityTypesEntitiesService) BatchUpdate(parent string, googleclouddialogflowv2beta1batchupdateentitiesrequest *GoogleCloudDialogflowV2beta1BatchUpdateEntitiesRequest) *ProjectsAgentEntityTypesEntitiesBatchUpdateCall {
	c := &ProjectsAgentEntityTypesEntitiesBatchUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.googleclouddialogflowv2beta1batchupdateentitiesrequest = googleclouddialogflowv2beta1batchupdateentitiesrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEntityTypesEntitiesBatchUpdateCall) Fields(s ...googleapi.Field) *ProjectsAgentEntityTypesEntitiesBatchUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEntityTypesEntitiesBatchUpdateCall) Context(ctx context.Context) *ProjectsAgentEntityTypesEntitiesBatchUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEntityTypesEntitiesBatchUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEntityTypesEntitiesBatchUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1batchupdateentitiesrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/entities:batchUpdate")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.entityTypes.entities.batchUpdate" call.
// Exactly one of *GoogleLongrunningOperation or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *GoogleLongrunningOperation.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentEntityTypesEntitiesBatchUpdateCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunningOperation, error) {
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
	ret := &GoogleLongrunningOperation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates or creates multiple entities in the specified entity type. This\nmethod does not affect entities in the entity type that aren't explicitly\nspecified in the request.\n\nOperation \u003cresponse: google.protobuf.Empty,\n           metadata: google.protobuf.Struct\u003e",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/entityTypes/{entityTypesId}/entities:batchUpdate",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.entityTypes.entities.batchUpdate",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The name of the entity type to update or create entities in.\nFormat: `projects/\u003cProject ID\u003e/agent/entityTypes/\u003cEntity Type ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/entityTypes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/entities:batchUpdate",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1BatchUpdateEntitiesRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleLongrunningOperation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.environments.users.sessions.deleteContexts":

type ProjectsAgentEnvironmentsUsersSessionsDeleteContextsCall struct {
	s          *Service
	parent     string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// DeleteContexts: Deletes all active contexts in the specified session.
func (r *ProjectsAgentEnvironmentsUsersSessionsService) DeleteContexts(parent string) *ProjectsAgentEnvironmentsUsersSessionsDeleteContextsCall {
	c := &ProjectsAgentEnvironmentsUsersSessionsDeleteContextsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEnvironmentsUsersSessionsDeleteContextsCall) Fields(s ...googleapi.Field) *ProjectsAgentEnvironmentsUsersSessionsDeleteContextsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEnvironmentsUsersSessionsDeleteContextsCall) Context(ctx context.Context) *ProjectsAgentEnvironmentsUsersSessionsDeleteContextsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEnvironmentsUsersSessionsDeleteContextsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEnvironmentsUsersSessionsDeleteContextsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/contexts")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.environments.users.sessions.deleteContexts" call.
// Exactly one of *GoogleProtobufEmpty or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleProtobufEmpty.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentEnvironmentsUsersSessionsDeleteContextsCall) Do(opts ...googleapi.CallOption) (*GoogleProtobufEmpty, error) {
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
	ret := &GoogleProtobufEmpty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes all active contexts in the specified session.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/environments/{environmentsId}/users/{usersId}/sessions/{sessionsId}/contexts",
	//   "httpMethod": "DELETE",
	//   "id": "dialogflow.projects.agent.environments.users.sessions.deleteContexts",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The name of the session to delete all contexts from. Format:\n`projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e` or `projects/\u003cProject\nID\u003e/agent/environments/\u003cEnvironment ID\u003e/users/\u003cUser ID\u003e/sessions/\u003cSession\nID\u003e`. If `Environment ID` is not specified we assume default 'draft'\nenvironment. If `User ID` is not specified, we assume default '-' user.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/environments/[^/]+/users/[^/]+/sessions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/contexts",
	//   "response": {
	//     "$ref": "GoogleProtobufEmpty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.environments.users.sessions.detectIntent":

type ProjectsAgentEnvironmentsUsersSessionsDetectIntentCall struct {
	s                                               *Service
	sessionid                                       string
	googleclouddialogflowv2beta1detectintentrequest *GoogleCloudDialogflowV2beta1DetectIntentRequest
	urlParams_                                      gensupport.URLParams
	ctx_                                            context.Context
	header_                                         http.Header
}

// DetectIntent: Processes a natural language query and returns
// structured, actionable data
// as a result. This method is not idempotent, because it may cause
// contexts
// and session entity types to be updated, which in turn might
// affect
// results of future queries.
func (r *ProjectsAgentEnvironmentsUsersSessionsService) DetectIntent(sessionid string, googleclouddialogflowv2beta1detectintentrequest *GoogleCloudDialogflowV2beta1DetectIntentRequest) *ProjectsAgentEnvironmentsUsersSessionsDetectIntentCall {
	c := &ProjectsAgentEnvironmentsUsersSessionsDetectIntentCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.sessionid = sessionid
	c.googleclouddialogflowv2beta1detectintentrequest = googleclouddialogflowv2beta1detectintentrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEnvironmentsUsersSessionsDetectIntentCall) Fields(s ...googleapi.Field) *ProjectsAgentEnvironmentsUsersSessionsDetectIntentCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEnvironmentsUsersSessionsDetectIntentCall) Context(ctx context.Context) *ProjectsAgentEnvironmentsUsersSessionsDetectIntentCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEnvironmentsUsersSessionsDetectIntentCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEnvironmentsUsersSessionsDetectIntentCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1detectintentrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+session}:detectIntent")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"session": c.sessionid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.environments.users.sessions.detectIntent" call.
// Exactly one of *GoogleCloudDialogflowV2beta1DetectIntentResponse or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1DetectIntentResponse.ServerResponse.Heade
// r or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsAgentEnvironmentsUsersSessionsDetectIntentCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1DetectIntentResponse, error) {
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
	ret := &GoogleCloudDialogflowV2beta1DetectIntentResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Processes a natural language query and returns structured, actionable data\nas a result. This method is not idempotent, because it may cause contexts\nand session entity types to be updated, which in turn might affect\nresults of future queries.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/environments/{environmentsId}/users/{usersId}/sessions/{sessionsId}:detectIntent",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.environments.users.sessions.detectIntent",
	//   "parameterOrder": [
	//     "session"
	//   ],
	//   "parameters": {
	//     "session": {
	//       "description": "Required. The name of the session this query is sent to. Format:\n`projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e`, or\n`projects/\u003cProject ID\u003e/agent/environments/\u003cEnvironment ID\u003e/users/\u003cUser\nID\u003e/sessions/\u003cSession ID\u003e`. If `Environment ID` is not specified, we assume\ndefault 'draft' environment. If `User ID` is not specified, we are using\n\"-\". It’s up to the API caller to choose an appropriate `Session ID` and\n`User Id`. They can be a random numbers or some type of user and session\nidentifiers (preferably hashed). The length of the `Session ID` and\n`User ID` must not exceed 36 characters.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/environments/[^/]+/users/[^/]+/sessions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+session}:detectIntent",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1DetectIntentRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1DetectIntentResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.environments.users.sessions.contexts.create":

type ProjectsAgentEnvironmentsUsersSessionsContextsCreateCall struct {
	s                                   *Service
	parent                              string
	googleclouddialogflowv2beta1context *GoogleCloudDialogflowV2beta1Context
	urlParams_                          gensupport.URLParams
	ctx_                                context.Context
	header_                             http.Header
}

// Create: Creates a context.
//
// If the specified context already exists, overrides the context.
func (r *ProjectsAgentEnvironmentsUsersSessionsContextsService) Create(parent string, googleclouddialogflowv2beta1context *GoogleCloudDialogflowV2beta1Context) *ProjectsAgentEnvironmentsUsersSessionsContextsCreateCall {
	c := &ProjectsAgentEnvironmentsUsersSessionsContextsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.googleclouddialogflowv2beta1context = googleclouddialogflowv2beta1context
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEnvironmentsUsersSessionsContextsCreateCall) Fields(s ...googleapi.Field) *ProjectsAgentEnvironmentsUsersSessionsContextsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEnvironmentsUsersSessionsContextsCreateCall) Context(ctx context.Context) *ProjectsAgentEnvironmentsUsersSessionsContextsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEnvironmentsUsersSessionsContextsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEnvironmentsUsersSessionsContextsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1context)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/contexts")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.environments.users.sessions.contexts.create" call.
// Exactly one of *GoogleCloudDialogflowV2beta1Context or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleCloudDialogflowV2beta1Context.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsAgentEnvironmentsUsersSessionsContextsCreateCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1Context, error) {
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
	ret := &GoogleCloudDialogflowV2beta1Context{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a context.\n\nIf the specified context already exists, overrides the context.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/environments/{environmentsId}/users/{usersId}/sessions/{sessionsId}/contexts",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.environments.users.sessions.contexts.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The session to create a context for.\nFormat: `projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e` or\n`projects/\u003cProject ID\u003e/agent/environments/\u003cEnvironment ID\u003e/users/\u003cUser\nID\u003e/sessions/\u003cSession ID\u003e`. If `Environment ID` is not specified, we assume\ndefault 'draft' environment. If `User ID` is not specified, we assume\ndefault '-' user.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/environments/[^/]+/users/[^/]+/sessions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/contexts",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Context"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Context"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.environments.users.sessions.contexts.delete":

type ProjectsAgentEnvironmentsUsersSessionsContextsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes the specified context.
func (r *ProjectsAgentEnvironmentsUsersSessionsContextsService) Delete(name string) *ProjectsAgentEnvironmentsUsersSessionsContextsDeleteCall {
	c := &ProjectsAgentEnvironmentsUsersSessionsContextsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEnvironmentsUsersSessionsContextsDeleteCall) Fields(s ...googleapi.Field) *ProjectsAgentEnvironmentsUsersSessionsContextsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEnvironmentsUsersSessionsContextsDeleteCall) Context(ctx context.Context) *ProjectsAgentEnvironmentsUsersSessionsContextsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEnvironmentsUsersSessionsContextsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEnvironmentsUsersSessionsContextsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.environments.users.sessions.contexts.delete" call.
// Exactly one of *GoogleProtobufEmpty or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleProtobufEmpty.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentEnvironmentsUsersSessionsContextsDeleteCall) Do(opts ...googleapi.CallOption) (*GoogleProtobufEmpty, error) {
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
	ret := &GoogleProtobufEmpty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the specified context.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/environments/{environmentsId}/users/{usersId}/sessions/{sessionsId}/contexts/{contextsId}",
	//   "httpMethod": "DELETE",
	//   "id": "dialogflow.projects.agent.environments.users.sessions.contexts.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the context to delete. Format:\n`projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e/contexts/\u003cContext ID\u003e`\nor `projects/\u003cProject ID\u003e/agent/environments/\u003cEnvironment ID\u003e/users/\u003cUser\nID\u003e/sessions/\u003cSession ID\u003e/contexts/\u003cContext ID\u003e`. If `Environment ID` is\nnot specified, we assume default 'draft' environment. If `User ID` is not\nspecified, we assume default '-' user.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/environments/[^/]+/users/[^/]+/sessions/[^/]+/contexts/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleProtobufEmpty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.environments.users.sessions.contexts.get":

type ProjectsAgentEnvironmentsUsersSessionsContextsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves the specified context.
func (r *ProjectsAgentEnvironmentsUsersSessionsContextsService) Get(name string) *ProjectsAgentEnvironmentsUsersSessionsContextsGetCall {
	c := &ProjectsAgentEnvironmentsUsersSessionsContextsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEnvironmentsUsersSessionsContextsGetCall) Fields(s ...googleapi.Field) *ProjectsAgentEnvironmentsUsersSessionsContextsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsAgentEnvironmentsUsersSessionsContextsGetCall) IfNoneMatch(entityTag string) *ProjectsAgentEnvironmentsUsersSessionsContextsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEnvironmentsUsersSessionsContextsGetCall) Context(ctx context.Context) *ProjectsAgentEnvironmentsUsersSessionsContextsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEnvironmentsUsersSessionsContextsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEnvironmentsUsersSessionsContextsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.environments.users.sessions.contexts.get" call.
// Exactly one of *GoogleCloudDialogflowV2beta1Context or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleCloudDialogflowV2beta1Context.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsAgentEnvironmentsUsersSessionsContextsGetCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1Context, error) {
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
	ret := &GoogleCloudDialogflowV2beta1Context{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the specified context.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/environments/{environmentsId}/users/{usersId}/sessions/{sessionsId}/contexts/{contextsId}",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.agent.environments.users.sessions.contexts.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the context. Format:\n`projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e/contexts/\u003cContext ID\u003e`\nor `projects/\u003cProject ID\u003e/agent/environments/\u003cEnvironment ID\u003e/users/\u003cUser\nID\u003e/sessions/\u003cSession ID\u003e/contexts/\u003cContext ID\u003e`. If `Environment ID` is\nnot specified, we assume default 'draft' environment. If `User ID` is not\nspecified, we assume default '-' user.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/environments/[^/]+/users/[^/]+/sessions/[^/]+/contexts/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Context"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.environments.users.sessions.contexts.list":

type ProjectsAgentEnvironmentsUsersSessionsContextsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Returns the list of all contexts in the specified session.
func (r *ProjectsAgentEnvironmentsUsersSessionsContextsService) List(parent string) *ProjectsAgentEnvironmentsUsersSessionsContextsListCall {
	c := &ProjectsAgentEnvironmentsUsersSessionsContextsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of items to return in a single page. By
// default 100 and at most 1000.
func (c *ProjectsAgentEnvironmentsUsersSessionsContextsListCall) PageSize(pageSize int64) *ProjectsAgentEnvironmentsUsersSessionsContextsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous list request.
func (c *ProjectsAgentEnvironmentsUsersSessionsContextsListCall) PageToken(pageToken string) *ProjectsAgentEnvironmentsUsersSessionsContextsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEnvironmentsUsersSessionsContextsListCall) Fields(s ...googleapi.Field) *ProjectsAgentEnvironmentsUsersSessionsContextsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsAgentEnvironmentsUsersSessionsContextsListCall) IfNoneMatch(entityTag string) *ProjectsAgentEnvironmentsUsersSessionsContextsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEnvironmentsUsersSessionsContextsListCall) Context(ctx context.Context) *ProjectsAgentEnvironmentsUsersSessionsContextsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEnvironmentsUsersSessionsContextsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEnvironmentsUsersSessionsContextsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/contexts")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.environments.users.sessions.contexts.list" call.
// Exactly one of *GoogleCloudDialogflowV2beta1ListContextsResponse or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1ListContextsResponse.ServerResponse.Heade
// r or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsAgentEnvironmentsUsersSessionsContextsListCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1ListContextsResponse, error) {
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
	ret := &GoogleCloudDialogflowV2beta1ListContextsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the list of all contexts in the specified session.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/environments/{environmentsId}/users/{usersId}/sessions/{sessionsId}/contexts",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.agent.environments.users.sessions.contexts.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Optional. The maximum number of items to return in a single page. By\ndefault 100 and at most 1000.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. The next_page_token value returned from a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The session to list all contexts from.\nFormat: `projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e` or\n`projects/\u003cProject ID\u003e/agent/environments/\u003cEnvironment ID\u003e/users/\u003cUser\nID\u003e/sessions/\u003cSession ID\u003e`. If `Environment ID` is not specified, we assume\ndefault 'draft' environment. If `User ID` is not specified, we assume\ndefault '-' user.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/environments/[^/]+/users/[^/]+/sessions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/contexts",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1ListContextsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsAgentEnvironmentsUsersSessionsContextsListCall) Pages(ctx context.Context, f func(*GoogleCloudDialogflowV2beta1ListContextsResponse) error) error {
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

// method id "dialogflow.projects.agent.environments.users.sessions.contexts.patch":

type ProjectsAgentEnvironmentsUsersSessionsContextsPatchCall struct {
	s                                   *Service
	nameid                              string
	googleclouddialogflowv2beta1context *GoogleCloudDialogflowV2beta1Context
	urlParams_                          gensupport.URLParams
	ctx_                                context.Context
	header_                             http.Header
}

// Patch: Updates the specified context.
func (r *ProjectsAgentEnvironmentsUsersSessionsContextsService) Patch(nameid string, googleclouddialogflowv2beta1context *GoogleCloudDialogflowV2beta1Context) *ProjectsAgentEnvironmentsUsersSessionsContextsPatchCall {
	c := &ProjectsAgentEnvironmentsUsersSessionsContextsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.nameid = nameid
	c.googleclouddialogflowv2beta1context = googleclouddialogflowv2beta1context
	return c
}

// UpdateMask sets the optional parameter "updateMask": The mask to
// control which fields get updated.
func (c *ProjectsAgentEnvironmentsUsersSessionsContextsPatchCall) UpdateMask(updateMask string) *ProjectsAgentEnvironmentsUsersSessionsContextsPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEnvironmentsUsersSessionsContextsPatchCall) Fields(s ...googleapi.Field) *ProjectsAgentEnvironmentsUsersSessionsContextsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEnvironmentsUsersSessionsContextsPatchCall) Context(ctx context.Context) *ProjectsAgentEnvironmentsUsersSessionsContextsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEnvironmentsUsersSessionsContextsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEnvironmentsUsersSessionsContextsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1context)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.nameid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.environments.users.sessions.contexts.patch" call.
// Exactly one of *GoogleCloudDialogflowV2beta1Context or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleCloudDialogflowV2beta1Context.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsAgentEnvironmentsUsersSessionsContextsPatchCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1Context, error) {
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
	ret := &GoogleCloudDialogflowV2beta1Context{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the specified context.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/environments/{environmentsId}/users/{usersId}/sessions/{sessionsId}/contexts/{contextsId}",
	//   "httpMethod": "PATCH",
	//   "id": "dialogflow.projects.agent.environments.users.sessions.contexts.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The unique identifier of the context. Format:\n`projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e/contexts/\u003cContext ID\u003e`,\nor `projects/\u003cProject ID\u003e/agent/environments/\u003cEnvironment ID\u003e/users/\u003cUser\nID\u003e/sessions/\u003cSession ID\u003e/contexts/\u003cContext ID\u003e`. The `Context ID` is\nalways converted to lowercase. If `Environment ID` is not specified, we\nassume default 'draft' environment. If `User ID` is not specified, we\nassume default '-' user.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/environments/[^/]+/users/[^/]+/sessions/[^/]+/contexts/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Optional. The mask to control which fields get updated.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Context"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Context"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.environments.users.sessions.entityTypes.create":

type ProjectsAgentEnvironmentsUsersSessionsEntityTypesCreateCall struct {
	s                                             *Service
	parent                                        string
	googleclouddialogflowv2beta1sessionentitytype *GoogleCloudDialogflowV2beta1SessionEntityType
	urlParams_                                    gensupport.URLParams
	ctx_                                          context.Context
	header_                                       http.Header
}

// Create: Creates a session entity type.
//
// If the specified session entity type already exists, overrides
// the
// session entity type.
func (r *ProjectsAgentEnvironmentsUsersSessionsEntityTypesService) Create(parent string, googleclouddialogflowv2beta1sessionentitytype *GoogleCloudDialogflowV2beta1SessionEntityType) *ProjectsAgentEnvironmentsUsersSessionsEntityTypesCreateCall {
	c := &ProjectsAgentEnvironmentsUsersSessionsEntityTypesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.googleclouddialogflowv2beta1sessionentitytype = googleclouddialogflowv2beta1sessionentitytype
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesCreateCall) Fields(s ...googleapi.Field) *ProjectsAgentEnvironmentsUsersSessionsEntityTypesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesCreateCall) Context(ctx context.Context) *ProjectsAgentEnvironmentsUsersSessionsEntityTypesCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1sessionentitytype)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/entityTypes")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.environments.users.sessions.entityTypes.create" call.
// Exactly one of *GoogleCloudDialogflowV2beta1SessionEntityType or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1SessionEntityType.ServerResponse.Header
// or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesCreateCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1SessionEntityType, error) {
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
	ret := &GoogleCloudDialogflowV2beta1SessionEntityType{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a session entity type.\n\nIf the specified session entity type already exists, overrides the\nsession entity type.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/environments/{environmentsId}/users/{usersId}/sessions/{sessionsId}/entityTypes",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.environments.users.sessions.entityTypes.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The session to create a session entity type for.\nFormat: `projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e` or\n`projects/\u003cProject ID\u003e/agent/environments/\u003cEnvironment ID\u003e/users/\u003cUser ID\u003e/\nsessions/\u003cSession ID\u003e`. If `Environment ID` is not specified, we assume\ndefault 'draft' environment. If `User ID` is not specified, we assume\ndefault '-' user.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/environments/[^/]+/users/[^/]+/sessions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/entityTypes",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1SessionEntityType"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1SessionEntityType"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.environments.users.sessions.entityTypes.delete":

type ProjectsAgentEnvironmentsUsersSessionsEntityTypesDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes the specified session entity type.
func (r *ProjectsAgentEnvironmentsUsersSessionsEntityTypesService) Delete(name string) *ProjectsAgentEnvironmentsUsersSessionsEntityTypesDeleteCall {
	c := &ProjectsAgentEnvironmentsUsersSessionsEntityTypesDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesDeleteCall) Fields(s ...googleapi.Field) *ProjectsAgentEnvironmentsUsersSessionsEntityTypesDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesDeleteCall) Context(ctx context.Context) *ProjectsAgentEnvironmentsUsersSessionsEntityTypesDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.environments.users.sessions.entityTypes.delete" call.
// Exactly one of *GoogleProtobufEmpty or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleProtobufEmpty.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesDeleteCall) Do(opts ...googleapi.CallOption) (*GoogleProtobufEmpty, error) {
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
	ret := &GoogleProtobufEmpty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the specified session entity type.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/environments/{environmentsId}/users/{usersId}/sessions/{sessionsId}/entityTypes/{entityTypesId}",
	//   "httpMethod": "DELETE",
	//   "id": "dialogflow.projects.agent.environments.users.sessions.entityTypes.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the entity type to delete. Format:\n`projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e/entityTypes/\u003cEntity Type\nDisplay Name\u003e` or `projects/\u003cProject ID\u003e/agent/environments/\u003cEnvironment\nID\u003e/users/\u003cUser ID\u003e/sessions/\u003cSession ID\u003e/entityTypes/\u003cEntity Type Display\nName\u003e`. If `Environment ID` is not specified, we assume default 'draft'\nenvironment. If `User ID` is not specified, we assume default '-' user.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/environments/[^/]+/users/[^/]+/sessions/[^/]+/entityTypes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleProtobufEmpty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.environments.users.sessions.entityTypes.get":

type ProjectsAgentEnvironmentsUsersSessionsEntityTypesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves the specified session entity type.
func (r *ProjectsAgentEnvironmentsUsersSessionsEntityTypesService) Get(name string) *ProjectsAgentEnvironmentsUsersSessionsEntityTypesGetCall {
	c := &ProjectsAgentEnvironmentsUsersSessionsEntityTypesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesGetCall) Fields(s ...googleapi.Field) *ProjectsAgentEnvironmentsUsersSessionsEntityTypesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesGetCall) IfNoneMatch(entityTag string) *ProjectsAgentEnvironmentsUsersSessionsEntityTypesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesGetCall) Context(ctx context.Context) *ProjectsAgentEnvironmentsUsersSessionsEntityTypesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.environments.users.sessions.entityTypes.get" call.
// Exactly one of *GoogleCloudDialogflowV2beta1SessionEntityType or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1SessionEntityType.ServerResponse.Header
// or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesGetCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1SessionEntityType, error) {
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
	ret := &GoogleCloudDialogflowV2beta1SessionEntityType{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the specified session entity type.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/environments/{environmentsId}/users/{usersId}/sessions/{sessionsId}/entityTypes/{entityTypesId}",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.agent.environments.users.sessions.entityTypes.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the session entity type. Format:\n`projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e/entityTypes/\u003cEntity Type\nDisplay Name\u003e` or `projects/\u003cProject ID\u003e/agent/environments/\u003cEnvironment\nID\u003e/users/\u003cUser ID\u003e/sessions/\u003cSession ID\u003e/entityTypes/\u003cEntity Type Display\nName\u003e`. If `Environment ID` is not specified, we assume default 'draft'\nenvironment. If `User ID` is not specified, we assume default '-' user.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/environments/[^/]+/users/[^/]+/sessions/[^/]+/entityTypes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1SessionEntityType"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.environments.users.sessions.entityTypes.list":

type ProjectsAgentEnvironmentsUsersSessionsEntityTypesListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Returns the list of all session entity types in the specified
// session.
func (r *ProjectsAgentEnvironmentsUsersSessionsEntityTypesService) List(parent string) *ProjectsAgentEnvironmentsUsersSessionsEntityTypesListCall {
	c := &ProjectsAgentEnvironmentsUsersSessionsEntityTypesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of items to return in a single page. By
// default 100 and at most 1000.
func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesListCall) PageSize(pageSize int64) *ProjectsAgentEnvironmentsUsersSessionsEntityTypesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous list request.
func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesListCall) PageToken(pageToken string) *ProjectsAgentEnvironmentsUsersSessionsEntityTypesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesListCall) Fields(s ...googleapi.Field) *ProjectsAgentEnvironmentsUsersSessionsEntityTypesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesListCall) IfNoneMatch(entityTag string) *ProjectsAgentEnvironmentsUsersSessionsEntityTypesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesListCall) Context(ctx context.Context) *ProjectsAgentEnvironmentsUsersSessionsEntityTypesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/entityTypes")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.environments.users.sessions.entityTypes.list" call.
// Exactly one of
// *GoogleCloudDialogflowV2beta1ListSessionEntityTypesResponse or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1ListSessionEntityTypesResponse.ServerResp
// onse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesListCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1ListSessionEntityTypesResponse, error) {
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
	ret := &GoogleCloudDialogflowV2beta1ListSessionEntityTypesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the list of all session entity types in the specified session.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/environments/{environmentsId}/users/{usersId}/sessions/{sessionsId}/entityTypes",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.agent.environments.users.sessions.entityTypes.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Optional. The maximum number of items to return in a single page. By\ndefault 100 and at most 1000.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. The next_page_token value returned from a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The session to list all session entity types from.\nFormat: `projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e` or\n`projects/\u003cProject ID\u003e/agent/environments/\u003cEnvironment ID\u003e/users/\u003cUser ID\u003e/\nsessions/\u003cSession ID\u003e`.\nIf `Environment ID` is not specified, we assume default 'draft'\nenvironment. If `User ID` is not specified, we assume default '-' user.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/environments/[^/]+/users/[^/]+/sessions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/entityTypes",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1ListSessionEntityTypesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesListCall) Pages(ctx context.Context, f func(*GoogleCloudDialogflowV2beta1ListSessionEntityTypesResponse) error) error {
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

// method id "dialogflow.projects.agent.environments.users.sessions.entityTypes.patch":

type ProjectsAgentEnvironmentsUsersSessionsEntityTypesPatchCall struct {
	s                                             *Service
	nameid                                        string
	googleclouddialogflowv2beta1sessionentitytype *GoogleCloudDialogflowV2beta1SessionEntityType
	urlParams_                                    gensupport.URLParams
	ctx_                                          context.Context
	header_                                       http.Header
}

// Patch: Updates the specified session entity type.
func (r *ProjectsAgentEnvironmentsUsersSessionsEntityTypesService) Patch(nameid string, googleclouddialogflowv2beta1sessionentitytype *GoogleCloudDialogflowV2beta1SessionEntityType) *ProjectsAgentEnvironmentsUsersSessionsEntityTypesPatchCall {
	c := &ProjectsAgentEnvironmentsUsersSessionsEntityTypesPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.nameid = nameid
	c.googleclouddialogflowv2beta1sessionentitytype = googleclouddialogflowv2beta1sessionentitytype
	return c
}

// UpdateMask sets the optional parameter "updateMask": The mask to
// control which fields get updated.
func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesPatchCall) UpdateMask(updateMask string) *ProjectsAgentEnvironmentsUsersSessionsEntityTypesPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesPatchCall) Fields(s ...googleapi.Field) *ProjectsAgentEnvironmentsUsersSessionsEntityTypesPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesPatchCall) Context(ctx context.Context) *ProjectsAgentEnvironmentsUsersSessionsEntityTypesPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1sessionentitytype)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.nameid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.environments.users.sessions.entityTypes.patch" call.
// Exactly one of *GoogleCloudDialogflowV2beta1SessionEntityType or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1SessionEntityType.ServerResponse.Header
// or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsAgentEnvironmentsUsersSessionsEntityTypesPatchCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1SessionEntityType, error) {
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
	ret := &GoogleCloudDialogflowV2beta1SessionEntityType{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the specified session entity type.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/environments/{environmentsId}/users/{usersId}/sessions/{sessionsId}/entityTypes/{entityTypesId}",
	//   "httpMethod": "PATCH",
	//   "id": "dialogflow.projects.agent.environments.users.sessions.entityTypes.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The unique identifier of this session entity type. Format:\n`projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e/entityTypes/\u003cEntity Type\nDisplay Name\u003e`, or\n`projects/\u003cProject ID\u003e/agent/environments/\u003cEnvironment ID\u003e/users/\u003cUser\nID\u003e/sessions/\u003cSession ID\u003e/entityTypes/\u003cEntity Type Display Name\u003e`.\nIf `Environment ID` is not specified, we assume default 'draft'\nenvironment. If `User ID` is not specified, we assume default '-' user.\n\n`\u003cEntity Type Display Name\u003e` must be the display name of an existing entity\ntype in the same agent that will be overridden or supplemented.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/environments/[^/]+/users/[^/]+/sessions/[^/]+/entityTypes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Optional. The mask to control which fields get updated.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1SessionEntityType"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1SessionEntityType"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.intents.batchDelete":

type ProjectsAgentIntentsBatchDeleteCall struct {
	s                                                     *Service
	parent                                                string
	googleclouddialogflowv2beta1batchdeleteintentsrequest *GoogleCloudDialogflowV2beta1BatchDeleteIntentsRequest
	urlParams_                                            gensupport.URLParams
	ctx_                                                  context.Context
	header_                                               http.Header
}

// BatchDelete: Deletes intents in the specified agent.
//
// Operation <response: google.protobuf.Empty>
func (r *ProjectsAgentIntentsService) BatchDelete(parent string, googleclouddialogflowv2beta1batchdeleteintentsrequest *GoogleCloudDialogflowV2beta1BatchDeleteIntentsRequest) *ProjectsAgentIntentsBatchDeleteCall {
	c := &ProjectsAgentIntentsBatchDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.googleclouddialogflowv2beta1batchdeleteintentsrequest = googleclouddialogflowv2beta1batchdeleteintentsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentIntentsBatchDeleteCall) Fields(s ...googleapi.Field) *ProjectsAgentIntentsBatchDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentIntentsBatchDeleteCall) Context(ctx context.Context) *ProjectsAgentIntentsBatchDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentIntentsBatchDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentIntentsBatchDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1batchdeleteintentsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/intents:batchDelete")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.intents.batchDelete" call.
// Exactly one of *GoogleLongrunningOperation or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *GoogleLongrunningOperation.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentIntentsBatchDeleteCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunningOperation, error) {
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
	ret := &GoogleLongrunningOperation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes intents in the specified agent.\n\nOperation \u003cresponse: google.protobuf.Empty\u003e",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/intents:batchDelete",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.intents.batchDelete",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The name of the agent to delete all entities types for. Format:\n`projects/\u003cProject ID\u003e/agent`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/intents:batchDelete",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1BatchDeleteIntentsRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleLongrunningOperation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.intents.batchUpdate":

type ProjectsAgentIntentsBatchUpdateCall struct {
	s                                                     *Service
	parent                                                string
	googleclouddialogflowv2beta1batchupdateintentsrequest *GoogleCloudDialogflowV2beta1BatchUpdateIntentsRequest
	urlParams_                                            gensupport.URLParams
	ctx_                                                  context.Context
	header_                                               http.Header
}

// BatchUpdate: Updates/Creates multiple intents in the specified
// agent.
//
// Operation <response: BatchUpdateIntentsResponse>
func (r *ProjectsAgentIntentsService) BatchUpdate(parent string, googleclouddialogflowv2beta1batchupdateintentsrequest *GoogleCloudDialogflowV2beta1BatchUpdateIntentsRequest) *ProjectsAgentIntentsBatchUpdateCall {
	c := &ProjectsAgentIntentsBatchUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.googleclouddialogflowv2beta1batchupdateintentsrequest = googleclouddialogflowv2beta1batchupdateintentsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentIntentsBatchUpdateCall) Fields(s ...googleapi.Field) *ProjectsAgentIntentsBatchUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentIntentsBatchUpdateCall) Context(ctx context.Context) *ProjectsAgentIntentsBatchUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentIntentsBatchUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentIntentsBatchUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1batchupdateintentsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/intents:batchUpdate")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.intents.batchUpdate" call.
// Exactly one of *GoogleLongrunningOperation or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *GoogleLongrunningOperation.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentIntentsBatchUpdateCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunningOperation, error) {
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
	ret := &GoogleLongrunningOperation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates/Creates multiple intents in the specified agent.\n\nOperation \u003cresponse: BatchUpdateIntentsResponse\u003e",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/intents:batchUpdate",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.intents.batchUpdate",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The name of the agent to update or create intents in.\nFormat: `projects/\u003cProject ID\u003e/agent`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/intents:batchUpdate",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1BatchUpdateIntentsRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleLongrunningOperation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.intents.create":

type ProjectsAgentIntentsCreateCall struct {
	s                                  *Service
	parent                             string
	googleclouddialogflowv2beta1intent *GoogleCloudDialogflowV2beta1Intent
	urlParams_                         gensupport.URLParams
	ctx_                               context.Context
	header_                            http.Header
}

// Create: Creates an intent in the specified agent.
func (r *ProjectsAgentIntentsService) Create(parent string, googleclouddialogflowv2beta1intent *GoogleCloudDialogflowV2beta1Intent) *ProjectsAgentIntentsCreateCall {
	c := &ProjectsAgentIntentsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.googleclouddialogflowv2beta1intent = googleclouddialogflowv2beta1intent
	return c
}

// IntentView sets the optional parameter "intentView": The resource
// view to apply to the returned intent.
//
// Possible values:
//   "INTENT_VIEW_UNSPECIFIED"
//   "INTENT_VIEW_FULL"
func (c *ProjectsAgentIntentsCreateCall) IntentView(intentView string) *ProjectsAgentIntentsCreateCall {
	c.urlParams_.Set("intentView", intentView)
	return c
}

// LanguageCode sets the optional parameter "languageCode": The language
// of training phrases, parameters and rich messages
// defined in `intent`. If not specified, the agent's default language
// is
// used. [More than a
// dozen
// languages](https://dialogflow.com/docs/reference/language) are
// supported.
// Note: languages must be enabled in the agent, before they can be
// used.
func (c *ProjectsAgentIntentsCreateCall) LanguageCode(languageCode string) *ProjectsAgentIntentsCreateCall {
	c.urlParams_.Set("languageCode", languageCode)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentIntentsCreateCall) Fields(s ...googleapi.Field) *ProjectsAgentIntentsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentIntentsCreateCall) Context(ctx context.Context) *ProjectsAgentIntentsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentIntentsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentIntentsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1intent)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/intents")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.intents.create" call.
// Exactly one of *GoogleCloudDialogflowV2beta1Intent or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleCloudDialogflowV2beta1Intent.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsAgentIntentsCreateCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1Intent, error) {
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
	ret := &GoogleCloudDialogflowV2beta1Intent{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates an intent in the specified agent.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/intents",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.intents.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "intentView": {
	//       "description": "Optional. The resource view to apply to the returned intent.",
	//       "enum": [
	//         "INTENT_VIEW_UNSPECIFIED",
	//         "INTENT_VIEW_FULL"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "languageCode": {
	//       "description": "Optional. The language of training phrases, parameters and rich messages\ndefined in `intent`. If not specified, the agent's default language is\nused. [More than a dozen\nlanguages](https://dialogflow.com/docs/reference/language) are supported.\nNote: languages must be enabled in the agent, before they can be used.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The agent to create a intent for.\nFormat: `projects/\u003cProject ID\u003e/agent`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/intents",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Intent"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Intent"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.intents.delete":

type ProjectsAgentIntentsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes the specified intent and its direct or indirect
// followup intents.
func (r *ProjectsAgentIntentsService) Delete(name string) *ProjectsAgentIntentsDeleteCall {
	c := &ProjectsAgentIntentsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentIntentsDeleteCall) Fields(s ...googleapi.Field) *ProjectsAgentIntentsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentIntentsDeleteCall) Context(ctx context.Context) *ProjectsAgentIntentsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentIntentsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentIntentsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.intents.delete" call.
// Exactly one of *GoogleProtobufEmpty or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleProtobufEmpty.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentIntentsDeleteCall) Do(opts ...googleapi.CallOption) (*GoogleProtobufEmpty, error) {
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
	ret := &GoogleProtobufEmpty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the specified intent and its direct or indirect followup intents.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/intents/{intentsId}",
	//   "httpMethod": "DELETE",
	//   "id": "dialogflow.projects.agent.intents.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the intent to delete. If this intent has direct or\nindirect followup intents, we also delete them.\n\nFormat: `projects/\u003cProject ID\u003e/agent/intents/\u003cIntent ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/intents/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleProtobufEmpty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.intents.get":

type ProjectsAgentIntentsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves the specified intent.
func (r *ProjectsAgentIntentsService) Get(name string) *ProjectsAgentIntentsGetCall {
	c := &ProjectsAgentIntentsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// IntentView sets the optional parameter "intentView": The resource
// view to apply to the returned intent.
//
// Possible values:
//   "INTENT_VIEW_UNSPECIFIED"
//   "INTENT_VIEW_FULL"
func (c *ProjectsAgentIntentsGetCall) IntentView(intentView string) *ProjectsAgentIntentsGetCall {
	c.urlParams_.Set("intentView", intentView)
	return c
}

// LanguageCode sets the optional parameter "languageCode": The language
// to retrieve training phrases, parameters and rich
// messages for. If not specified, the agent's default language is
// used.
// [More than a
// dozen
// languages](https://dialogflow.com/docs/reference/language) are
// supported.
// Note: languages must be enabled in the agent, before they can be
// used.
func (c *ProjectsAgentIntentsGetCall) LanguageCode(languageCode string) *ProjectsAgentIntentsGetCall {
	c.urlParams_.Set("languageCode", languageCode)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentIntentsGetCall) Fields(s ...googleapi.Field) *ProjectsAgentIntentsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsAgentIntentsGetCall) IfNoneMatch(entityTag string) *ProjectsAgentIntentsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentIntentsGetCall) Context(ctx context.Context) *ProjectsAgentIntentsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentIntentsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentIntentsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.intents.get" call.
// Exactly one of *GoogleCloudDialogflowV2beta1Intent or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleCloudDialogflowV2beta1Intent.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsAgentIntentsGetCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1Intent, error) {
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
	ret := &GoogleCloudDialogflowV2beta1Intent{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the specified intent.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/intents/{intentsId}",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.agent.intents.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "intentView": {
	//       "description": "Optional. The resource view to apply to the returned intent.",
	//       "enum": [
	//         "INTENT_VIEW_UNSPECIFIED",
	//         "INTENT_VIEW_FULL"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "languageCode": {
	//       "description": "Optional. The language to retrieve training phrases, parameters and rich\nmessages for. If not specified, the agent's default language is used.\n[More than a dozen\nlanguages](https://dialogflow.com/docs/reference/language) are supported.\nNote: languages must be enabled in the agent, before they can be used.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "Required. The name of the intent.\nFormat: `projects/\u003cProject ID\u003e/agent/intents/\u003cIntent ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/intents/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Intent"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.intents.list":

type ProjectsAgentIntentsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Returns the list of all intents in the specified agent.
func (r *ProjectsAgentIntentsService) List(parent string) *ProjectsAgentIntentsListCall {
	c := &ProjectsAgentIntentsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// IntentView sets the optional parameter "intentView": The resource
// view to apply to the returned intent.
//
// Possible values:
//   "INTENT_VIEW_UNSPECIFIED"
//   "INTENT_VIEW_FULL"
func (c *ProjectsAgentIntentsListCall) IntentView(intentView string) *ProjectsAgentIntentsListCall {
	c.urlParams_.Set("intentView", intentView)
	return c
}

// LanguageCode sets the optional parameter "languageCode": The language
// to list training phrases, parameters and rich
// messages for. If not specified, the agent's default language is
// used.
// [More than a
// dozen
// languages](https://dialogflow.com/docs/reference/language) are
// supported.
// Note: languages must be enabled in the agent before they can be used.
func (c *ProjectsAgentIntentsListCall) LanguageCode(languageCode string) *ProjectsAgentIntentsListCall {
	c.urlParams_.Set("languageCode", languageCode)
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of items to return in a single page. By
// default 100 and at most 1000.
func (c *ProjectsAgentIntentsListCall) PageSize(pageSize int64) *ProjectsAgentIntentsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous list request.
func (c *ProjectsAgentIntentsListCall) PageToken(pageToken string) *ProjectsAgentIntentsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentIntentsListCall) Fields(s ...googleapi.Field) *ProjectsAgentIntentsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsAgentIntentsListCall) IfNoneMatch(entityTag string) *ProjectsAgentIntentsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentIntentsListCall) Context(ctx context.Context) *ProjectsAgentIntentsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentIntentsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentIntentsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/intents")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.intents.list" call.
// Exactly one of *GoogleCloudDialogflowV2beta1ListIntentsResponse or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1ListIntentsResponse.ServerResponse.Header
//  or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsAgentIntentsListCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1ListIntentsResponse, error) {
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
	ret := &GoogleCloudDialogflowV2beta1ListIntentsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the list of all intents in the specified agent.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/intents",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.agent.intents.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "intentView": {
	//       "description": "Optional. The resource view to apply to the returned intent.",
	//       "enum": [
	//         "INTENT_VIEW_UNSPECIFIED",
	//         "INTENT_VIEW_FULL"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "languageCode": {
	//       "description": "Optional. The language to list training phrases, parameters and rich\nmessages for. If not specified, the agent's default language is used.\n[More than a dozen\nlanguages](https://dialogflow.com/docs/reference/language) are supported.\nNote: languages must be enabled in the agent before they can be used.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Optional. The maximum number of items to return in a single page. By\ndefault 100 and at most 1000.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. The next_page_token value returned from a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The agent to list all intents from.\nFormat: `projects/\u003cProject ID\u003e/agent`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/intents",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1ListIntentsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsAgentIntentsListCall) Pages(ctx context.Context, f func(*GoogleCloudDialogflowV2beta1ListIntentsResponse) error) error {
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

// method id "dialogflow.projects.agent.intents.patch":

type ProjectsAgentIntentsPatchCall struct {
	s                                  *Service
	nameid                             string
	googleclouddialogflowv2beta1intent *GoogleCloudDialogflowV2beta1Intent
	urlParams_                         gensupport.URLParams
	ctx_                               context.Context
	header_                            http.Header
}

// Patch: Updates the specified intent.
func (r *ProjectsAgentIntentsService) Patch(nameid string, googleclouddialogflowv2beta1intent *GoogleCloudDialogflowV2beta1Intent) *ProjectsAgentIntentsPatchCall {
	c := &ProjectsAgentIntentsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.nameid = nameid
	c.googleclouddialogflowv2beta1intent = googleclouddialogflowv2beta1intent
	return c
}

// IntentView sets the optional parameter "intentView": The resource
// view to apply to the returned intent.
//
// Possible values:
//   "INTENT_VIEW_UNSPECIFIED"
//   "INTENT_VIEW_FULL"
func (c *ProjectsAgentIntentsPatchCall) IntentView(intentView string) *ProjectsAgentIntentsPatchCall {
	c.urlParams_.Set("intentView", intentView)
	return c
}

// LanguageCode sets the optional parameter "languageCode": The language
// of training phrases, parameters and rich messages
// defined in `intent`. If not specified, the agent's default language
// is
// used. [More than a
// dozen
// languages](https://dialogflow.com/docs/reference/language) are
// supported.
// Note: languages must be enabled in the agent, before they can be
// used.
func (c *ProjectsAgentIntentsPatchCall) LanguageCode(languageCode string) *ProjectsAgentIntentsPatchCall {
	c.urlParams_.Set("languageCode", languageCode)
	return c
}

// UpdateMask sets the optional parameter "updateMask": The mask to
// control which fields get updated.
func (c *ProjectsAgentIntentsPatchCall) UpdateMask(updateMask string) *ProjectsAgentIntentsPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentIntentsPatchCall) Fields(s ...googleapi.Field) *ProjectsAgentIntentsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentIntentsPatchCall) Context(ctx context.Context) *ProjectsAgentIntentsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentIntentsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentIntentsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1intent)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.nameid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.intents.patch" call.
// Exactly one of *GoogleCloudDialogflowV2beta1Intent or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleCloudDialogflowV2beta1Intent.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsAgentIntentsPatchCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1Intent, error) {
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
	ret := &GoogleCloudDialogflowV2beta1Intent{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the specified intent.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/intents/{intentsId}",
	//   "httpMethod": "PATCH",
	//   "id": "dialogflow.projects.agent.intents.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "intentView": {
	//       "description": "Optional. The resource view to apply to the returned intent.",
	//       "enum": [
	//         "INTENT_VIEW_UNSPECIFIED",
	//         "INTENT_VIEW_FULL"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "languageCode": {
	//       "description": "Optional. The language of training phrases, parameters and rich messages\ndefined in `intent`. If not specified, the agent's default language is\nused. [More than a dozen\nlanguages](https://dialogflow.com/docs/reference/language) are supported.\nNote: languages must be enabled in the agent, before they can be used.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "Required for all methods except `create` (`create` populates the name\nautomatically.\nThe unique identifier of this intent.\nFormat: `projects/\u003cProject ID\u003e/agent/intents/\u003cIntent ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/intents/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Optional. The mask to control which fields get updated.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Intent"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Intent"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.knowledgeBases.create":

type ProjectsAgentKnowledgeBasesCreateCall struct {
	s                                         *Service
	parent                                    string
	googleclouddialogflowv2beta1knowledgebase *GoogleCloudDialogflowV2beta1KnowledgeBase
	urlParams_                                gensupport.URLParams
	ctx_                                      context.Context
	header_                                   http.Header
}

// Create: Creates a knowledge base.
func (r *ProjectsAgentKnowledgeBasesService) Create(parent string, googleclouddialogflowv2beta1knowledgebase *GoogleCloudDialogflowV2beta1KnowledgeBase) *ProjectsAgentKnowledgeBasesCreateCall {
	c := &ProjectsAgentKnowledgeBasesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.googleclouddialogflowv2beta1knowledgebase = googleclouddialogflowv2beta1knowledgebase
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentKnowledgeBasesCreateCall) Fields(s ...googleapi.Field) *ProjectsAgentKnowledgeBasesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentKnowledgeBasesCreateCall) Context(ctx context.Context) *ProjectsAgentKnowledgeBasesCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentKnowledgeBasesCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentKnowledgeBasesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1knowledgebase)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/knowledgeBases")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.knowledgeBases.create" call.
// Exactly one of *GoogleCloudDialogflowV2beta1KnowledgeBase or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1KnowledgeBase.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsAgentKnowledgeBasesCreateCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1KnowledgeBase, error) {
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
	ret := &GoogleCloudDialogflowV2beta1KnowledgeBase{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a knowledge base.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/knowledgeBases",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.knowledgeBases.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The project to create a knowledge base for.\nFormat: `projects/\u003cProject ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/knowledgeBases",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1KnowledgeBase"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1KnowledgeBase"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.knowledgeBases.delete":

type ProjectsAgentKnowledgeBasesDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes the specified knowledge base.
func (r *ProjectsAgentKnowledgeBasesService) Delete(name string) *ProjectsAgentKnowledgeBasesDeleteCall {
	c := &ProjectsAgentKnowledgeBasesDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Force sets the optional parameter "force": Force deletes the
// knowledge base. When set to true, any documents
// in the knowledge base are also deleted.
func (c *ProjectsAgentKnowledgeBasesDeleteCall) Force(force bool) *ProjectsAgentKnowledgeBasesDeleteCall {
	c.urlParams_.Set("force", fmt.Sprint(force))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentKnowledgeBasesDeleteCall) Fields(s ...googleapi.Field) *ProjectsAgentKnowledgeBasesDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentKnowledgeBasesDeleteCall) Context(ctx context.Context) *ProjectsAgentKnowledgeBasesDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentKnowledgeBasesDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentKnowledgeBasesDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.knowledgeBases.delete" call.
// Exactly one of *GoogleProtobufEmpty or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleProtobufEmpty.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentKnowledgeBasesDeleteCall) Do(opts ...googleapi.CallOption) (*GoogleProtobufEmpty, error) {
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
	ret := &GoogleProtobufEmpty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the specified knowledge base.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/knowledgeBases/{knowledgeBasesId}",
	//   "httpMethod": "DELETE",
	//   "id": "dialogflow.projects.agent.knowledgeBases.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "force": {
	//       "description": "Optional. Force deletes the knowledge base. When set to true, any documents\nin the knowledge base are also deleted.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "name": {
	//       "description": "Required. The name of the knowledge base to delete.\nFormat: `projects/\u003cProject ID\u003e/knowledgeBases/\u003cKnowledge Base ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/knowledgeBases/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleProtobufEmpty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.knowledgeBases.get":

type ProjectsAgentKnowledgeBasesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves the specified knowledge base.
func (r *ProjectsAgentKnowledgeBasesService) Get(name string) *ProjectsAgentKnowledgeBasesGetCall {
	c := &ProjectsAgentKnowledgeBasesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentKnowledgeBasesGetCall) Fields(s ...googleapi.Field) *ProjectsAgentKnowledgeBasesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsAgentKnowledgeBasesGetCall) IfNoneMatch(entityTag string) *ProjectsAgentKnowledgeBasesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentKnowledgeBasesGetCall) Context(ctx context.Context) *ProjectsAgentKnowledgeBasesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentKnowledgeBasesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentKnowledgeBasesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.knowledgeBases.get" call.
// Exactly one of *GoogleCloudDialogflowV2beta1KnowledgeBase or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1KnowledgeBase.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsAgentKnowledgeBasesGetCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1KnowledgeBase, error) {
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
	ret := &GoogleCloudDialogflowV2beta1KnowledgeBase{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the specified knowledge base.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/knowledgeBases/{knowledgeBasesId}",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.agent.knowledgeBases.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the knowledge base to retrieve.\nFormat `projects/\u003cProject ID\u003e/knowledgeBases/\u003cKnowledge Base ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/knowledgeBases/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1KnowledgeBase"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.knowledgeBases.list":

type ProjectsAgentKnowledgeBasesListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Returns the list of all knowledge bases of the specified agent.
func (r *ProjectsAgentKnowledgeBasesService) List(parent string) *ProjectsAgentKnowledgeBasesListCall {
	c := &ProjectsAgentKnowledgeBasesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of items to return in a single page. By
// default 10 and at most 100.
func (c *ProjectsAgentKnowledgeBasesListCall) PageSize(pageSize int64) *ProjectsAgentKnowledgeBasesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous list request.
func (c *ProjectsAgentKnowledgeBasesListCall) PageToken(pageToken string) *ProjectsAgentKnowledgeBasesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentKnowledgeBasesListCall) Fields(s ...googleapi.Field) *ProjectsAgentKnowledgeBasesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsAgentKnowledgeBasesListCall) IfNoneMatch(entityTag string) *ProjectsAgentKnowledgeBasesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentKnowledgeBasesListCall) Context(ctx context.Context) *ProjectsAgentKnowledgeBasesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentKnowledgeBasesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentKnowledgeBasesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/knowledgeBases")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.knowledgeBases.list" call.
// Exactly one of
// *GoogleCloudDialogflowV2beta1ListKnowledgeBasesResponse or error will
// be non-nil. Any non-2xx status code is an error. Response headers are
// in either
// *GoogleCloudDialogflowV2beta1ListKnowledgeBasesResponse.ServerResponse
// .Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsAgentKnowledgeBasesListCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1ListKnowledgeBasesResponse, error) {
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
	ret := &GoogleCloudDialogflowV2beta1ListKnowledgeBasesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the list of all knowledge bases of the specified agent.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/knowledgeBases",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.agent.knowledgeBases.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Optional. The maximum number of items to return in a single page. By\ndefault 10 and at most 100.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. The next_page_token value returned from a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The project to list of knowledge bases for.\nFormat: `projects/\u003cProject ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/knowledgeBases",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1ListKnowledgeBasesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsAgentKnowledgeBasesListCall) Pages(ctx context.Context, f func(*GoogleCloudDialogflowV2beta1ListKnowledgeBasesResponse) error) error {
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

// method id "dialogflow.projects.agent.knowledgeBases.documents.create":

type ProjectsAgentKnowledgeBasesDocumentsCreateCall struct {
	s                                    *Service
	parent                               string
	googleclouddialogflowv2beta1document *GoogleCloudDialogflowV2beta1Document
	urlParams_                           gensupport.URLParams
	ctx_                                 context.Context
	header_                              http.Header
}

// Create: Creates a new document.
//
// Operation <response: Document,
//            metadata: KnowledgeOperationMetadata>
func (r *ProjectsAgentKnowledgeBasesDocumentsService) Create(parent string, googleclouddialogflowv2beta1document *GoogleCloudDialogflowV2beta1Document) *ProjectsAgentKnowledgeBasesDocumentsCreateCall {
	c := &ProjectsAgentKnowledgeBasesDocumentsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.googleclouddialogflowv2beta1document = googleclouddialogflowv2beta1document
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentKnowledgeBasesDocumentsCreateCall) Fields(s ...googleapi.Field) *ProjectsAgentKnowledgeBasesDocumentsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentKnowledgeBasesDocumentsCreateCall) Context(ctx context.Context) *ProjectsAgentKnowledgeBasesDocumentsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentKnowledgeBasesDocumentsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentKnowledgeBasesDocumentsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1document)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/documents")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.knowledgeBases.documents.create" call.
// Exactly one of *GoogleLongrunningOperation or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *GoogleLongrunningOperation.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentKnowledgeBasesDocumentsCreateCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunningOperation, error) {
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
	ret := &GoogleLongrunningOperation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new document.\n\nOperation \u003cresponse: Document,\n           metadata: KnowledgeOperationMetadata\u003e",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/knowledgeBases/{knowledgeBasesId}/documents",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.knowledgeBases.documents.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The knoweldge base to create a document for.\nFormat: `projects/\u003cProject ID\u003e/knowledgeBases/\u003cKnowledge Base ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/knowledgeBases/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/documents",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Document"
	//   },
	//   "response": {
	//     "$ref": "GoogleLongrunningOperation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.knowledgeBases.documents.delete":

type ProjectsAgentKnowledgeBasesDocumentsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes the specified document.
//
// Operation <response: google.protobuf.Empty,
//            metadata: KnowledgeOperationMetadata>
func (r *ProjectsAgentKnowledgeBasesDocumentsService) Delete(name string) *ProjectsAgentKnowledgeBasesDocumentsDeleteCall {
	c := &ProjectsAgentKnowledgeBasesDocumentsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentKnowledgeBasesDocumentsDeleteCall) Fields(s ...googleapi.Field) *ProjectsAgentKnowledgeBasesDocumentsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentKnowledgeBasesDocumentsDeleteCall) Context(ctx context.Context) *ProjectsAgentKnowledgeBasesDocumentsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentKnowledgeBasesDocumentsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentKnowledgeBasesDocumentsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.knowledgeBases.documents.delete" call.
// Exactly one of *GoogleLongrunningOperation or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *GoogleLongrunningOperation.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentKnowledgeBasesDocumentsDeleteCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunningOperation, error) {
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
	ret := &GoogleLongrunningOperation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the specified document.\n\nOperation \u003cresponse: google.protobuf.Empty,\n           metadata: KnowledgeOperationMetadata\u003e",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/knowledgeBases/{knowledgeBasesId}/documents/{documentsId}",
	//   "httpMethod": "DELETE",
	//   "id": "dialogflow.projects.agent.knowledgeBases.documents.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the document to delete.\nFormat: `projects/\u003cProject ID\u003e/knowledgeBases/\u003cKnowledge Base\nID\u003e/documents/\u003cDocument ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/knowledgeBases/[^/]+/documents/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleLongrunningOperation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.knowledgeBases.documents.get":

type ProjectsAgentKnowledgeBasesDocumentsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves the specified document.
func (r *ProjectsAgentKnowledgeBasesDocumentsService) Get(name string) *ProjectsAgentKnowledgeBasesDocumentsGetCall {
	c := &ProjectsAgentKnowledgeBasesDocumentsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentKnowledgeBasesDocumentsGetCall) Fields(s ...googleapi.Field) *ProjectsAgentKnowledgeBasesDocumentsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsAgentKnowledgeBasesDocumentsGetCall) IfNoneMatch(entityTag string) *ProjectsAgentKnowledgeBasesDocumentsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentKnowledgeBasesDocumentsGetCall) Context(ctx context.Context) *ProjectsAgentKnowledgeBasesDocumentsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentKnowledgeBasesDocumentsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentKnowledgeBasesDocumentsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.knowledgeBases.documents.get" call.
// Exactly one of *GoogleCloudDialogflowV2beta1Document or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleCloudDialogflowV2beta1Document.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsAgentKnowledgeBasesDocumentsGetCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1Document, error) {
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
	ret := &GoogleCloudDialogflowV2beta1Document{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the specified document.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/knowledgeBases/{knowledgeBasesId}/documents/{documentsId}",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.agent.knowledgeBases.documents.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the document to retrieve.\nFormat `projects/\u003cProject ID\u003e/knowledgeBases/\u003cKnowledge Base\nID\u003e/documents/\u003cDocument ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/knowledgeBases/[^/]+/documents/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Document"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.knowledgeBases.documents.list":

type ProjectsAgentKnowledgeBasesDocumentsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Returns the list of all documents of the knowledge base.
func (r *ProjectsAgentKnowledgeBasesDocumentsService) List(parent string) *ProjectsAgentKnowledgeBasesDocumentsListCall {
	c := &ProjectsAgentKnowledgeBasesDocumentsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of items to return in a single page. By
// default 10 and at most 100.
func (c *ProjectsAgentKnowledgeBasesDocumentsListCall) PageSize(pageSize int64) *ProjectsAgentKnowledgeBasesDocumentsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous list request.
func (c *ProjectsAgentKnowledgeBasesDocumentsListCall) PageToken(pageToken string) *ProjectsAgentKnowledgeBasesDocumentsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentKnowledgeBasesDocumentsListCall) Fields(s ...googleapi.Field) *ProjectsAgentKnowledgeBasesDocumentsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsAgentKnowledgeBasesDocumentsListCall) IfNoneMatch(entityTag string) *ProjectsAgentKnowledgeBasesDocumentsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentKnowledgeBasesDocumentsListCall) Context(ctx context.Context) *ProjectsAgentKnowledgeBasesDocumentsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentKnowledgeBasesDocumentsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentKnowledgeBasesDocumentsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/documents")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.knowledgeBases.documents.list" call.
// Exactly one of *GoogleCloudDialogflowV2beta1ListDocumentsResponse or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1ListDocumentsResponse.ServerResponse.Head
// er or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsAgentKnowledgeBasesDocumentsListCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1ListDocumentsResponse, error) {
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
	ret := &GoogleCloudDialogflowV2beta1ListDocumentsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the list of all documents of the knowledge base.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/knowledgeBases/{knowledgeBasesId}/documents",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.agent.knowledgeBases.documents.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Optional. The maximum number of items to return in a single page. By\ndefault 10 and at most 100.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. The next_page_token value returned from a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The knowledge base to list all documents for.\nFormat: `projects/\u003cProject ID\u003e/knowledgeBases/\u003cKnowledge Base ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/knowledgeBases/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/documents",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1ListDocumentsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsAgentKnowledgeBasesDocumentsListCall) Pages(ctx context.Context, f func(*GoogleCloudDialogflowV2beta1ListDocumentsResponse) error) error {
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

// method id "dialogflow.projects.agent.sessions.deleteContexts":

type ProjectsAgentSessionsDeleteContextsCall struct {
	s          *Service
	parent     string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// DeleteContexts: Deletes all active contexts in the specified session.
func (r *ProjectsAgentSessionsService) DeleteContexts(parent string) *ProjectsAgentSessionsDeleteContextsCall {
	c := &ProjectsAgentSessionsDeleteContextsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentSessionsDeleteContextsCall) Fields(s ...googleapi.Field) *ProjectsAgentSessionsDeleteContextsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentSessionsDeleteContextsCall) Context(ctx context.Context) *ProjectsAgentSessionsDeleteContextsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentSessionsDeleteContextsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentSessionsDeleteContextsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/contexts")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.sessions.deleteContexts" call.
// Exactly one of *GoogleProtobufEmpty or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleProtobufEmpty.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentSessionsDeleteContextsCall) Do(opts ...googleapi.CallOption) (*GoogleProtobufEmpty, error) {
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
	ret := &GoogleProtobufEmpty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes all active contexts in the specified session.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/sessions/{sessionsId}/contexts",
	//   "httpMethod": "DELETE",
	//   "id": "dialogflow.projects.agent.sessions.deleteContexts",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The name of the session to delete all contexts from. Format:\n`projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e` or `projects/\u003cProject\nID\u003e/agent/environments/\u003cEnvironment ID\u003e/users/\u003cUser ID\u003e/sessions/\u003cSession\nID\u003e`. If `Environment ID` is not specified we assume default 'draft'\nenvironment. If `User ID` is not specified, we assume default '-' user.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/sessions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/contexts",
	//   "response": {
	//     "$ref": "GoogleProtobufEmpty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.sessions.detectIntent":

type ProjectsAgentSessionsDetectIntentCall struct {
	s                                               *Service
	sessionid                                       string
	googleclouddialogflowv2beta1detectintentrequest *GoogleCloudDialogflowV2beta1DetectIntentRequest
	urlParams_                                      gensupport.URLParams
	ctx_                                            context.Context
	header_                                         http.Header
}

// DetectIntent: Processes a natural language query and returns
// structured, actionable data
// as a result. This method is not idempotent, because it may cause
// contexts
// and session entity types to be updated, which in turn might
// affect
// results of future queries.
func (r *ProjectsAgentSessionsService) DetectIntent(sessionid string, googleclouddialogflowv2beta1detectintentrequest *GoogleCloudDialogflowV2beta1DetectIntentRequest) *ProjectsAgentSessionsDetectIntentCall {
	c := &ProjectsAgentSessionsDetectIntentCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.sessionid = sessionid
	c.googleclouddialogflowv2beta1detectintentrequest = googleclouddialogflowv2beta1detectintentrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentSessionsDetectIntentCall) Fields(s ...googleapi.Field) *ProjectsAgentSessionsDetectIntentCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentSessionsDetectIntentCall) Context(ctx context.Context) *ProjectsAgentSessionsDetectIntentCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentSessionsDetectIntentCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentSessionsDetectIntentCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1detectintentrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+session}:detectIntent")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"session": c.sessionid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.sessions.detectIntent" call.
// Exactly one of *GoogleCloudDialogflowV2beta1DetectIntentResponse or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1DetectIntentResponse.ServerResponse.Heade
// r or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsAgentSessionsDetectIntentCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1DetectIntentResponse, error) {
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
	ret := &GoogleCloudDialogflowV2beta1DetectIntentResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Processes a natural language query and returns structured, actionable data\nas a result. This method is not idempotent, because it may cause contexts\nand session entity types to be updated, which in turn might affect\nresults of future queries.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/sessions/{sessionsId}:detectIntent",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.sessions.detectIntent",
	//   "parameterOrder": [
	//     "session"
	//   ],
	//   "parameters": {
	//     "session": {
	//       "description": "Required. The name of the session this query is sent to. Format:\n`projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e`, or\n`projects/\u003cProject ID\u003e/agent/environments/\u003cEnvironment ID\u003e/users/\u003cUser\nID\u003e/sessions/\u003cSession ID\u003e`. If `Environment ID` is not specified, we assume\ndefault 'draft' environment. If `User ID` is not specified, we are using\n\"-\". It’s up to the API caller to choose an appropriate `Session ID` and\n`User Id`. They can be a random numbers or some type of user and session\nidentifiers (preferably hashed). The length of the `Session ID` and\n`User ID` must not exceed 36 characters.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/sessions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+session}:detectIntent",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1DetectIntentRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1DetectIntentResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.sessions.contexts.create":

type ProjectsAgentSessionsContextsCreateCall struct {
	s                                   *Service
	parent                              string
	googleclouddialogflowv2beta1context *GoogleCloudDialogflowV2beta1Context
	urlParams_                          gensupport.URLParams
	ctx_                                context.Context
	header_                             http.Header
}

// Create: Creates a context.
//
// If the specified context already exists, overrides the context.
func (r *ProjectsAgentSessionsContextsService) Create(parent string, googleclouddialogflowv2beta1context *GoogleCloudDialogflowV2beta1Context) *ProjectsAgentSessionsContextsCreateCall {
	c := &ProjectsAgentSessionsContextsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.googleclouddialogflowv2beta1context = googleclouddialogflowv2beta1context
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentSessionsContextsCreateCall) Fields(s ...googleapi.Field) *ProjectsAgentSessionsContextsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentSessionsContextsCreateCall) Context(ctx context.Context) *ProjectsAgentSessionsContextsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentSessionsContextsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentSessionsContextsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1context)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/contexts")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.sessions.contexts.create" call.
// Exactly one of *GoogleCloudDialogflowV2beta1Context or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleCloudDialogflowV2beta1Context.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsAgentSessionsContextsCreateCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1Context, error) {
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
	ret := &GoogleCloudDialogflowV2beta1Context{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a context.\n\nIf the specified context already exists, overrides the context.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/sessions/{sessionsId}/contexts",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.sessions.contexts.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The session to create a context for.\nFormat: `projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e` or\n`projects/\u003cProject ID\u003e/agent/environments/\u003cEnvironment ID\u003e/users/\u003cUser\nID\u003e/sessions/\u003cSession ID\u003e`. If `Environment ID` is not specified, we assume\ndefault 'draft' environment. If `User ID` is not specified, we assume\ndefault '-' user.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/sessions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/contexts",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Context"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Context"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.sessions.contexts.delete":

type ProjectsAgentSessionsContextsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes the specified context.
func (r *ProjectsAgentSessionsContextsService) Delete(name string) *ProjectsAgentSessionsContextsDeleteCall {
	c := &ProjectsAgentSessionsContextsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentSessionsContextsDeleteCall) Fields(s ...googleapi.Field) *ProjectsAgentSessionsContextsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentSessionsContextsDeleteCall) Context(ctx context.Context) *ProjectsAgentSessionsContextsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentSessionsContextsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentSessionsContextsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.sessions.contexts.delete" call.
// Exactly one of *GoogleProtobufEmpty or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleProtobufEmpty.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentSessionsContextsDeleteCall) Do(opts ...googleapi.CallOption) (*GoogleProtobufEmpty, error) {
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
	ret := &GoogleProtobufEmpty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the specified context.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/sessions/{sessionsId}/contexts/{contextsId}",
	//   "httpMethod": "DELETE",
	//   "id": "dialogflow.projects.agent.sessions.contexts.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the context to delete. Format:\n`projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e/contexts/\u003cContext ID\u003e`\nor `projects/\u003cProject ID\u003e/agent/environments/\u003cEnvironment ID\u003e/users/\u003cUser\nID\u003e/sessions/\u003cSession ID\u003e/contexts/\u003cContext ID\u003e`. If `Environment ID` is\nnot specified, we assume default 'draft' environment. If `User ID` is not\nspecified, we assume default '-' user.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/sessions/[^/]+/contexts/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleProtobufEmpty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.sessions.contexts.get":

type ProjectsAgentSessionsContextsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves the specified context.
func (r *ProjectsAgentSessionsContextsService) Get(name string) *ProjectsAgentSessionsContextsGetCall {
	c := &ProjectsAgentSessionsContextsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentSessionsContextsGetCall) Fields(s ...googleapi.Field) *ProjectsAgentSessionsContextsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsAgentSessionsContextsGetCall) IfNoneMatch(entityTag string) *ProjectsAgentSessionsContextsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentSessionsContextsGetCall) Context(ctx context.Context) *ProjectsAgentSessionsContextsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentSessionsContextsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentSessionsContextsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.sessions.contexts.get" call.
// Exactly one of *GoogleCloudDialogflowV2beta1Context or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleCloudDialogflowV2beta1Context.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsAgentSessionsContextsGetCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1Context, error) {
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
	ret := &GoogleCloudDialogflowV2beta1Context{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the specified context.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/sessions/{sessionsId}/contexts/{contextsId}",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.agent.sessions.contexts.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the context. Format:\n`projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e/contexts/\u003cContext ID\u003e`\nor `projects/\u003cProject ID\u003e/agent/environments/\u003cEnvironment ID\u003e/users/\u003cUser\nID\u003e/sessions/\u003cSession ID\u003e/contexts/\u003cContext ID\u003e`. If `Environment ID` is\nnot specified, we assume default 'draft' environment. If `User ID` is not\nspecified, we assume default '-' user.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/sessions/[^/]+/contexts/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Context"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.sessions.contexts.list":

type ProjectsAgentSessionsContextsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Returns the list of all contexts in the specified session.
func (r *ProjectsAgentSessionsContextsService) List(parent string) *ProjectsAgentSessionsContextsListCall {
	c := &ProjectsAgentSessionsContextsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of items to return in a single page. By
// default 100 and at most 1000.
func (c *ProjectsAgentSessionsContextsListCall) PageSize(pageSize int64) *ProjectsAgentSessionsContextsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous list request.
func (c *ProjectsAgentSessionsContextsListCall) PageToken(pageToken string) *ProjectsAgentSessionsContextsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentSessionsContextsListCall) Fields(s ...googleapi.Field) *ProjectsAgentSessionsContextsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsAgentSessionsContextsListCall) IfNoneMatch(entityTag string) *ProjectsAgentSessionsContextsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentSessionsContextsListCall) Context(ctx context.Context) *ProjectsAgentSessionsContextsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentSessionsContextsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentSessionsContextsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/contexts")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.sessions.contexts.list" call.
// Exactly one of *GoogleCloudDialogflowV2beta1ListContextsResponse or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1ListContextsResponse.ServerResponse.Heade
// r or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsAgentSessionsContextsListCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1ListContextsResponse, error) {
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
	ret := &GoogleCloudDialogflowV2beta1ListContextsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the list of all contexts in the specified session.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/sessions/{sessionsId}/contexts",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.agent.sessions.contexts.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Optional. The maximum number of items to return in a single page. By\ndefault 100 and at most 1000.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. The next_page_token value returned from a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The session to list all contexts from.\nFormat: `projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e` or\n`projects/\u003cProject ID\u003e/agent/environments/\u003cEnvironment ID\u003e/users/\u003cUser\nID\u003e/sessions/\u003cSession ID\u003e`. If `Environment ID` is not specified, we assume\ndefault 'draft' environment. If `User ID` is not specified, we assume\ndefault '-' user.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/sessions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/contexts",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1ListContextsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsAgentSessionsContextsListCall) Pages(ctx context.Context, f func(*GoogleCloudDialogflowV2beta1ListContextsResponse) error) error {
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

// method id "dialogflow.projects.agent.sessions.contexts.patch":

type ProjectsAgentSessionsContextsPatchCall struct {
	s                                   *Service
	nameid                              string
	googleclouddialogflowv2beta1context *GoogleCloudDialogflowV2beta1Context
	urlParams_                          gensupport.URLParams
	ctx_                                context.Context
	header_                             http.Header
}

// Patch: Updates the specified context.
func (r *ProjectsAgentSessionsContextsService) Patch(nameid string, googleclouddialogflowv2beta1context *GoogleCloudDialogflowV2beta1Context) *ProjectsAgentSessionsContextsPatchCall {
	c := &ProjectsAgentSessionsContextsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.nameid = nameid
	c.googleclouddialogflowv2beta1context = googleclouddialogflowv2beta1context
	return c
}

// UpdateMask sets the optional parameter "updateMask": The mask to
// control which fields get updated.
func (c *ProjectsAgentSessionsContextsPatchCall) UpdateMask(updateMask string) *ProjectsAgentSessionsContextsPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentSessionsContextsPatchCall) Fields(s ...googleapi.Field) *ProjectsAgentSessionsContextsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentSessionsContextsPatchCall) Context(ctx context.Context) *ProjectsAgentSessionsContextsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentSessionsContextsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentSessionsContextsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1context)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.nameid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.sessions.contexts.patch" call.
// Exactly one of *GoogleCloudDialogflowV2beta1Context or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleCloudDialogflowV2beta1Context.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsAgentSessionsContextsPatchCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1Context, error) {
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
	ret := &GoogleCloudDialogflowV2beta1Context{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the specified context.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/sessions/{sessionsId}/contexts/{contextsId}",
	//   "httpMethod": "PATCH",
	//   "id": "dialogflow.projects.agent.sessions.contexts.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The unique identifier of the context. Format:\n`projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e/contexts/\u003cContext ID\u003e`,\nor `projects/\u003cProject ID\u003e/agent/environments/\u003cEnvironment ID\u003e/users/\u003cUser\nID\u003e/sessions/\u003cSession ID\u003e/contexts/\u003cContext ID\u003e`. The `Context ID` is\nalways converted to lowercase. If `Environment ID` is not specified, we\nassume default 'draft' environment. If `User ID` is not specified, we\nassume default '-' user.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/sessions/[^/]+/contexts/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Optional. The mask to control which fields get updated.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Context"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Context"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.sessions.entityTypes.create":

type ProjectsAgentSessionsEntityTypesCreateCall struct {
	s                                             *Service
	parent                                        string
	googleclouddialogflowv2beta1sessionentitytype *GoogleCloudDialogflowV2beta1SessionEntityType
	urlParams_                                    gensupport.URLParams
	ctx_                                          context.Context
	header_                                       http.Header
}

// Create: Creates a session entity type.
//
// If the specified session entity type already exists, overrides
// the
// session entity type.
func (r *ProjectsAgentSessionsEntityTypesService) Create(parent string, googleclouddialogflowv2beta1sessionentitytype *GoogleCloudDialogflowV2beta1SessionEntityType) *ProjectsAgentSessionsEntityTypesCreateCall {
	c := &ProjectsAgentSessionsEntityTypesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.googleclouddialogflowv2beta1sessionentitytype = googleclouddialogflowv2beta1sessionentitytype
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentSessionsEntityTypesCreateCall) Fields(s ...googleapi.Field) *ProjectsAgentSessionsEntityTypesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentSessionsEntityTypesCreateCall) Context(ctx context.Context) *ProjectsAgentSessionsEntityTypesCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentSessionsEntityTypesCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentSessionsEntityTypesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1sessionentitytype)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/entityTypes")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.sessions.entityTypes.create" call.
// Exactly one of *GoogleCloudDialogflowV2beta1SessionEntityType or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1SessionEntityType.ServerResponse.Header
// or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsAgentSessionsEntityTypesCreateCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1SessionEntityType, error) {
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
	ret := &GoogleCloudDialogflowV2beta1SessionEntityType{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a session entity type.\n\nIf the specified session entity type already exists, overrides the\nsession entity type.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/sessions/{sessionsId}/entityTypes",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.sessions.entityTypes.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The session to create a session entity type for.\nFormat: `projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e` or\n`projects/\u003cProject ID\u003e/agent/environments/\u003cEnvironment ID\u003e/users/\u003cUser ID\u003e/\nsessions/\u003cSession ID\u003e`. If `Environment ID` is not specified, we assume\ndefault 'draft' environment. If `User ID` is not specified, we assume\ndefault '-' user.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/sessions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/entityTypes",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1SessionEntityType"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1SessionEntityType"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.sessions.entityTypes.delete":

type ProjectsAgentSessionsEntityTypesDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes the specified session entity type.
func (r *ProjectsAgentSessionsEntityTypesService) Delete(name string) *ProjectsAgentSessionsEntityTypesDeleteCall {
	c := &ProjectsAgentSessionsEntityTypesDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentSessionsEntityTypesDeleteCall) Fields(s ...googleapi.Field) *ProjectsAgentSessionsEntityTypesDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentSessionsEntityTypesDeleteCall) Context(ctx context.Context) *ProjectsAgentSessionsEntityTypesDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentSessionsEntityTypesDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentSessionsEntityTypesDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.sessions.entityTypes.delete" call.
// Exactly one of *GoogleProtobufEmpty or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleProtobufEmpty.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentSessionsEntityTypesDeleteCall) Do(opts ...googleapi.CallOption) (*GoogleProtobufEmpty, error) {
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
	ret := &GoogleProtobufEmpty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the specified session entity type.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/sessions/{sessionsId}/entityTypes/{entityTypesId}",
	//   "httpMethod": "DELETE",
	//   "id": "dialogflow.projects.agent.sessions.entityTypes.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the entity type to delete. Format:\n`projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e/entityTypes/\u003cEntity Type\nDisplay Name\u003e` or `projects/\u003cProject ID\u003e/agent/environments/\u003cEnvironment\nID\u003e/users/\u003cUser ID\u003e/sessions/\u003cSession ID\u003e/entityTypes/\u003cEntity Type Display\nName\u003e`. If `Environment ID` is not specified, we assume default 'draft'\nenvironment. If `User ID` is not specified, we assume default '-' user.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/sessions/[^/]+/entityTypes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleProtobufEmpty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.sessions.entityTypes.get":

type ProjectsAgentSessionsEntityTypesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves the specified session entity type.
func (r *ProjectsAgentSessionsEntityTypesService) Get(name string) *ProjectsAgentSessionsEntityTypesGetCall {
	c := &ProjectsAgentSessionsEntityTypesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentSessionsEntityTypesGetCall) Fields(s ...googleapi.Field) *ProjectsAgentSessionsEntityTypesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsAgentSessionsEntityTypesGetCall) IfNoneMatch(entityTag string) *ProjectsAgentSessionsEntityTypesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentSessionsEntityTypesGetCall) Context(ctx context.Context) *ProjectsAgentSessionsEntityTypesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentSessionsEntityTypesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentSessionsEntityTypesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.sessions.entityTypes.get" call.
// Exactly one of *GoogleCloudDialogflowV2beta1SessionEntityType or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1SessionEntityType.ServerResponse.Header
// or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsAgentSessionsEntityTypesGetCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1SessionEntityType, error) {
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
	ret := &GoogleCloudDialogflowV2beta1SessionEntityType{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the specified session entity type.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/sessions/{sessionsId}/entityTypes/{entityTypesId}",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.agent.sessions.entityTypes.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the session entity type. Format:\n`projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e/entityTypes/\u003cEntity Type\nDisplay Name\u003e` or `projects/\u003cProject ID\u003e/agent/environments/\u003cEnvironment\nID\u003e/users/\u003cUser ID\u003e/sessions/\u003cSession ID\u003e/entityTypes/\u003cEntity Type Display\nName\u003e`. If `Environment ID` is not specified, we assume default 'draft'\nenvironment. If `User ID` is not specified, we assume default '-' user.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/sessions/[^/]+/entityTypes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1SessionEntityType"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.sessions.entityTypes.list":

type ProjectsAgentSessionsEntityTypesListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Returns the list of all session entity types in the specified
// session.
func (r *ProjectsAgentSessionsEntityTypesService) List(parent string) *ProjectsAgentSessionsEntityTypesListCall {
	c := &ProjectsAgentSessionsEntityTypesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of items to return in a single page. By
// default 100 and at most 1000.
func (c *ProjectsAgentSessionsEntityTypesListCall) PageSize(pageSize int64) *ProjectsAgentSessionsEntityTypesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous list request.
func (c *ProjectsAgentSessionsEntityTypesListCall) PageToken(pageToken string) *ProjectsAgentSessionsEntityTypesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentSessionsEntityTypesListCall) Fields(s ...googleapi.Field) *ProjectsAgentSessionsEntityTypesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsAgentSessionsEntityTypesListCall) IfNoneMatch(entityTag string) *ProjectsAgentSessionsEntityTypesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentSessionsEntityTypesListCall) Context(ctx context.Context) *ProjectsAgentSessionsEntityTypesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentSessionsEntityTypesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentSessionsEntityTypesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/entityTypes")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.sessions.entityTypes.list" call.
// Exactly one of
// *GoogleCloudDialogflowV2beta1ListSessionEntityTypesResponse or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1ListSessionEntityTypesResponse.ServerResp
// onse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsAgentSessionsEntityTypesListCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1ListSessionEntityTypesResponse, error) {
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
	ret := &GoogleCloudDialogflowV2beta1ListSessionEntityTypesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the list of all session entity types in the specified session.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/sessions/{sessionsId}/entityTypes",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.agent.sessions.entityTypes.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Optional. The maximum number of items to return in a single page. By\ndefault 100 and at most 1000.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. The next_page_token value returned from a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The session to list all session entity types from.\nFormat: `projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e` or\n`projects/\u003cProject ID\u003e/agent/environments/\u003cEnvironment ID\u003e/users/\u003cUser ID\u003e/\nsessions/\u003cSession ID\u003e`.\nIf `Environment ID` is not specified, we assume default 'draft'\nenvironment. If `User ID` is not specified, we assume default '-' user.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/sessions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/entityTypes",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1ListSessionEntityTypesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsAgentSessionsEntityTypesListCall) Pages(ctx context.Context, f func(*GoogleCloudDialogflowV2beta1ListSessionEntityTypesResponse) error) error {
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

// method id "dialogflow.projects.agent.sessions.entityTypes.patch":

type ProjectsAgentSessionsEntityTypesPatchCall struct {
	s                                             *Service
	nameid                                        string
	googleclouddialogflowv2beta1sessionentitytype *GoogleCloudDialogflowV2beta1SessionEntityType
	urlParams_                                    gensupport.URLParams
	ctx_                                          context.Context
	header_                                       http.Header
}

// Patch: Updates the specified session entity type.
func (r *ProjectsAgentSessionsEntityTypesService) Patch(nameid string, googleclouddialogflowv2beta1sessionentitytype *GoogleCloudDialogflowV2beta1SessionEntityType) *ProjectsAgentSessionsEntityTypesPatchCall {
	c := &ProjectsAgentSessionsEntityTypesPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.nameid = nameid
	c.googleclouddialogflowv2beta1sessionentitytype = googleclouddialogflowv2beta1sessionentitytype
	return c
}

// UpdateMask sets the optional parameter "updateMask": The mask to
// control which fields get updated.
func (c *ProjectsAgentSessionsEntityTypesPatchCall) UpdateMask(updateMask string) *ProjectsAgentSessionsEntityTypesPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentSessionsEntityTypesPatchCall) Fields(s ...googleapi.Field) *ProjectsAgentSessionsEntityTypesPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentSessionsEntityTypesPatchCall) Context(ctx context.Context) *ProjectsAgentSessionsEntityTypesPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentSessionsEntityTypesPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentSessionsEntityTypesPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1sessionentitytype)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.nameid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.sessions.entityTypes.patch" call.
// Exactly one of *GoogleCloudDialogflowV2beta1SessionEntityType or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1SessionEntityType.ServerResponse.Header
// or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsAgentSessionsEntityTypesPatchCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1SessionEntityType, error) {
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
	ret := &GoogleCloudDialogflowV2beta1SessionEntityType{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the specified session entity type.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/sessions/{sessionsId}/entityTypes/{entityTypesId}",
	//   "httpMethod": "PATCH",
	//   "id": "dialogflow.projects.agent.sessions.entityTypes.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The unique identifier of this session entity type. Format:\n`projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e/entityTypes/\u003cEntity Type\nDisplay Name\u003e`, or\n`projects/\u003cProject ID\u003e/agent/environments/\u003cEnvironment ID\u003e/users/\u003cUser\nID\u003e/sessions/\u003cSession ID\u003e/entityTypes/\u003cEntity Type Display Name\u003e`.\nIf `Environment ID` is not specified, we assume default 'draft'\nenvironment. If `User ID` is not specified, we assume default '-' user.\n\n`\u003cEntity Type Display Name\u003e` must be the display name of an existing entity\ntype in the same agent that will be overridden or supplemented.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/sessions/[^/]+/entityTypes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Optional. The mask to control which fields get updated.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1SessionEntityType"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1SessionEntityType"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.conversationProfiles.create":

type ProjectsConversationProfilesCreateCall struct {
	s                                               *Service
	parent                                          string
	googleclouddialogflowv2beta1conversationprofile *GoogleCloudDialogflowV2beta1ConversationProfile
	urlParams_                                      gensupport.URLParams
	ctx_                                            context.Context
	header_                                         http.Header
}

// Create: Creates a conversation profile in the specified project.
func (r *ProjectsConversationProfilesService) Create(parent string, googleclouddialogflowv2beta1conversationprofile *GoogleCloudDialogflowV2beta1ConversationProfile) *ProjectsConversationProfilesCreateCall {
	c := &ProjectsConversationProfilesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.googleclouddialogflowv2beta1conversationprofile = googleclouddialogflowv2beta1conversationprofile
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConversationProfilesCreateCall) Fields(s ...googleapi.Field) *ProjectsConversationProfilesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConversationProfilesCreateCall) Context(ctx context.Context) *ProjectsConversationProfilesCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsConversationProfilesCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsConversationProfilesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1conversationprofile)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/conversationProfiles")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.conversationProfiles.create" call.
// Exactly one of *GoogleCloudDialogflowV2beta1ConversationProfile or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1ConversationProfile.ServerResponse.Header
//  or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsConversationProfilesCreateCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1ConversationProfile, error) {
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
	ret := &GoogleCloudDialogflowV2beta1ConversationProfile{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a conversation profile in the specified project.",
	//   "flatPath": "v2beta1/projects/{projectsId}/conversationProfiles",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.conversationProfiles.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The project to create a conversation profile for.\nFormat: `projects/\u003cProject ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/conversationProfiles",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1ConversationProfile"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1ConversationProfile"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.conversationProfiles.delete":

type ProjectsConversationProfilesDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes the specified conversation profile.
func (r *ProjectsConversationProfilesService) Delete(name string) *ProjectsConversationProfilesDeleteCall {
	c := &ProjectsConversationProfilesDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConversationProfilesDeleteCall) Fields(s ...googleapi.Field) *ProjectsConversationProfilesDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConversationProfilesDeleteCall) Context(ctx context.Context) *ProjectsConversationProfilesDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsConversationProfilesDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsConversationProfilesDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.conversationProfiles.delete" call.
// Exactly one of *GoogleProtobufEmpty or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleProtobufEmpty.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsConversationProfilesDeleteCall) Do(opts ...googleapi.CallOption) (*GoogleProtobufEmpty, error) {
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
	ret := &GoogleProtobufEmpty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the specified conversation profile.",
	//   "flatPath": "v2beta1/projects/{projectsId}/conversationProfiles/{conversationProfilesId}",
	//   "httpMethod": "DELETE",
	//   "id": "dialogflow.projects.conversationProfiles.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the conversation profile to delete.\nFormat: `projects/\u003cProject ID\u003e/conversationProfiles/\u003cConversation Profile\nID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/conversationProfiles/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleProtobufEmpty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.conversationProfiles.get":

type ProjectsConversationProfilesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves the specified conversation profile.
func (r *ProjectsConversationProfilesService) Get(name string) *ProjectsConversationProfilesGetCall {
	c := &ProjectsConversationProfilesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConversationProfilesGetCall) Fields(s ...googleapi.Field) *ProjectsConversationProfilesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsConversationProfilesGetCall) IfNoneMatch(entityTag string) *ProjectsConversationProfilesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConversationProfilesGetCall) Context(ctx context.Context) *ProjectsConversationProfilesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsConversationProfilesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsConversationProfilesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.conversationProfiles.get" call.
// Exactly one of *GoogleCloudDialogflowV2beta1ConversationProfile or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1ConversationProfile.ServerResponse.Header
//  or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsConversationProfilesGetCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1ConversationProfile, error) {
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
	ret := &GoogleCloudDialogflowV2beta1ConversationProfile{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the specified conversation profile.",
	//   "flatPath": "v2beta1/projects/{projectsId}/conversationProfiles/{conversationProfilesId}",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.conversationProfiles.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The resource name of the conversation profile.\nFormat: `projects/\u003cProject ID\u003e/conversationProfiles/\u003cConversation Profile\nID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/conversationProfiles/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1ConversationProfile"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.conversationProfiles.list":

type ProjectsConversationProfilesListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Returns the list of all conversation profiles in the specified
// project.
func (r *ProjectsConversationProfilesService) List(parent string) *ProjectsConversationProfilesListCall {
	c := &ProjectsConversationProfilesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of items to return in a single page. By
// default 100 and at most 1000.
func (c *ProjectsConversationProfilesListCall) PageSize(pageSize int64) *ProjectsConversationProfilesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous list request.
func (c *ProjectsConversationProfilesListCall) PageToken(pageToken string) *ProjectsConversationProfilesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConversationProfilesListCall) Fields(s ...googleapi.Field) *ProjectsConversationProfilesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsConversationProfilesListCall) IfNoneMatch(entityTag string) *ProjectsConversationProfilesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConversationProfilesListCall) Context(ctx context.Context) *ProjectsConversationProfilesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsConversationProfilesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsConversationProfilesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/conversationProfiles")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.conversationProfiles.list" call.
// Exactly one of
// *GoogleCloudDialogflowV2beta1ListConversationProfilesResponse or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1ListConversationProfilesResponse.ServerRe
// sponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsConversationProfilesListCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1ListConversationProfilesResponse, error) {
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
	ret := &GoogleCloudDialogflowV2beta1ListConversationProfilesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the list of all conversation profiles in the specified project.",
	//   "flatPath": "v2beta1/projects/{projectsId}/conversationProfiles",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.conversationProfiles.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Optional. The maximum number of items to return in a single page. By\ndefault 100 and at most 1000.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. The next_page_token value returned from a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The project to list all conversation profiles from.\nFormat: `projects/\u003cProject ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/conversationProfiles",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1ListConversationProfilesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsConversationProfilesListCall) Pages(ctx context.Context, f func(*GoogleCloudDialogflowV2beta1ListConversationProfilesResponse) error) error {
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

// method id "dialogflow.projects.conversationProfiles.patch":

type ProjectsConversationProfilesPatchCall struct {
	s                                               *Service
	nameid                                          string
	googleclouddialogflowv2beta1conversationprofile *GoogleCloudDialogflowV2beta1ConversationProfile
	urlParams_                                      gensupport.URLParams
	ctx_                                            context.Context
	header_                                         http.Header
}

// Patch: Updates the specified conversation profile.
func (r *ProjectsConversationProfilesService) Patch(nameid string, googleclouddialogflowv2beta1conversationprofile *GoogleCloudDialogflowV2beta1ConversationProfile) *ProjectsConversationProfilesPatchCall {
	c := &ProjectsConversationProfilesPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.nameid = nameid
	c.googleclouddialogflowv2beta1conversationprofile = googleclouddialogflowv2beta1conversationprofile
	return c
}

// UpdateMask sets the optional parameter "updateMask": Required. The
// mask to control which fields to update.
func (c *ProjectsConversationProfilesPatchCall) UpdateMask(updateMask string) *ProjectsConversationProfilesPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConversationProfilesPatchCall) Fields(s ...googleapi.Field) *ProjectsConversationProfilesPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConversationProfilesPatchCall) Context(ctx context.Context) *ProjectsConversationProfilesPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsConversationProfilesPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsConversationProfilesPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1conversationprofile)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.nameid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.conversationProfiles.patch" call.
// Exactly one of *GoogleCloudDialogflowV2beta1ConversationProfile or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1ConversationProfile.ServerResponse.Header
//  or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsConversationProfilesPatchCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1ConversationProfile, error) {
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
	ret := &GoogleCloudDialogflowV2beta1ConversationProfile{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the specified conversation profile.",
	//   "flatPath": "v2beta1/projects/{projectsId}/conversationProfiles/{conversationProfilesId}",
	//   "httpMethod": "PATCH",
	//   "id": "dialogflow.projects.conversationProfiles.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required for all methods except `create` (`create` populates the name\nautomatically).\nThe unique identifier of this conversation profile.\nFormat: `projects/\u003cProject ID\u003e/conversationProfiles/\u003cConversation Profile\nID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/conversationProfiles/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Required. The mask to control which fields to update.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1ConversationProfile"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1ConversationProfile"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.conversations.addConversationPhoneNumber":

type ProjectsConversationsAddConversationPhoneNumberCall struct {
	s                                                             *Service
	name                                                          string
	googleclouddialogflowv2beta1addconversationphonenumberrequest *GoogleCloudDialogflowV2beta1AddConversationPhoneNumberRequest
	urlParams_                                                    gensupport.URLParams
	ctx_                                                          context.Context
	header_                                                       http.Header
}

// AddConversationPhoneNumber: Sets a phone number for this converstion
// to connect to.
func (r *ProjectsConversationsService) AddConversationPhoneNumber(name string, googleclouddialogflowv2beta1addconversationphonenumberrequest *GoogleCloudDialogflowV2beta1AddConversationPhoneNumberRequest) *ProjectsConversationsAddConversationPhoneNumberCall {
	c := &ProjectsConversationsAddConversationPhoneNumberCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.googleclouddialogflowv2beta1addconversationphonenumberrequest = googleclouddialogflowv2beta1addconversationphonenumberrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConversationsAddConversationPhoneNumberCall) Fields(s ...googleapi.Field) *ProjectsConversationsAddConversationPhoneNumberCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConversationsAddConversationPhoneNumberCall) Context(ctx context.Context) *ProjectsConversationsAddConversationPhoneNumberCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsConversationsAddConversationPhoneNumberCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsConversationsAddConversationPhoneNumberCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1addconversationphonenumberrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}:addConversationPhoneNumber")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.conversations.addConversationPhoneNumber" call.
// Exactly one of *GoogleCloudDialogflowV2beta1ConversationPhoneNumber
// or error will be non-nil. Any non-2xx status code is an error.
// Response headers are in either
// *GoogleCloudDialogflowV2beta1ConversationPhoneNumber.ServerResponse.He
// ader or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsConversationsAddConversationPhoneNumberCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1ConversationPhoneNumber, error) {
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
	ret := &GoogleCloudDialogflowV2beta1ConversationPhoneNumber{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Sets a phone number for this converstion to connect to.",
	//   "flatPath": "v2beta1/projects/{projectsId}/conversations/{conversationsId}:addConversationPhoneNumber",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.conversations.addConversationPhoneNumber",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of this conversation.\nFormat: `projects/\u003cProject ID\u003e/conversations/\u003cConversation ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/conversations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}:addConversationPhoneNumber",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1AddConversationPhoneNumberRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1ConversationPhoneNumber"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.conversations.complete":

type ProjectsConversationsCompleteCall struct {
	s                                                       *Service
	nameid                                                  string
	googleclouddialogflowv2beta1completeconversationrequest *GoogleCloudDialogflowV2beta1CompleteConversationRequest
	urlParams_                                              gensupport.URLParams
	ctx_                                                    context.Context
	header_                                                 http.Header
}

// Complete: Completes the specified conversation. Finished
// conversations are purged
// from the database after 30 days.
func (r *ProjectsConversationsService) Complete(nameid string, googleclouddialogflowv2beta1completeconversationrequest *GoogleCloudDialogflowV2beta1CompleteConversationRequest) *ProjectsConversationsCompleteCall {
	c := &ProjectsConversationsCompleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.nameid = nameid
	c.googleclouddialogflowv2beta1completeconversationrequest = googleclouddialogflowv2beta1completeconversationrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConversationsCompleteCall) Fields(s ...googleapi.Field) *ProjectsConversationsCompleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConversationsCompleteCall) Context(ctx context.Context) *ProjectsConversationsCompleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsConversationsCompleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsConversationsCompleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1completeconversationrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}:complete")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.nameid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.conversations.complete" call.
// Exactly one of *GoogleCloudDialogflowV2beta1Conversation or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1Conversation.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsConversationsCompleteCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1Conversation, error) {
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
	ret := &GoogleCloudDialogflowV2beta1Conversation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Completes the specified conversation. Finished conversations are purged\nfrom the database after 30 days.",
	//   "flatPath": "v2beta1/projects/{projectsId}/conversations/{conversationsId}:complete",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.conversations.complete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. Resource identifier of the conversation to close.\nFormat: `projects/\u003cProject ID\u003e/conversations/\u003cConversation ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/conversations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}:complete",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1CompleteConversationRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Conversation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.conversations.create":

type ProjectsConversationsCreateCall struct {
	s                                        *Service
	parentid                                 string
	googleclouddialogflowv2beta1conversation *GoogleCloudDialogflowV2beta1Conversation
	urlParams_                               gensupport.URLParams
	ctx_                                     context.Context
	header_                                  http.Header
}

// Create: Creates a new conversation. Conversation are auto-completed
// after 24 hours.
func (r *ProjectsConversationsService) Create(parentid string, googleclouddialogflowv2beta1conversation *GoogleCloudDialogflowV2beta1Conversation) *ProjectsConversationsCreateCall {
	c := &ProjectsConversationsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parentid = parentid
	c.googleclouddialogflowv2beta1conversation = googleclouddialogflowv2beta1conversation
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConversationsCreateCall) Fields(s ...googleapi.Field) *ProjectsConversationsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConversationsCreateCall) Context(ctx context.Context) *ProjectsConversationsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsConversationsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsConversationsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1conversation)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/conversations")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parentid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.conversations.create" call.
// Exactly one of *GoogleCloudDialogflowV2beta1Conversation or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1Conversation.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsConversationsCreateCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1Conversation, error) {
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
	ret := &GoogleCloudDialogflowV2beta1Conversation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new conversation. Conversation are auto-completed after 24 hours.",
	//   "flatPath": "v2beta1/projects/{projectsId}/conversations",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.conversations.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. Resource identifier of the project creating the conversation.\nFormat: `projects/\u003cProject ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/conversations",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Conversation"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Conversation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.conversations.get":

type ProjectsConversationsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves the specific conversation.
func (r *ProjectsConversationsService) Get(name string) *ProjectsConversationsGetCall {
	c := &ProjectsConversationsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConversationsGetCall) Fields(s ...googleapi.Field) *ProjectsConversationsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsConversationsGetCall) IfNoneMatch(entityTag string) *ProjectsConversationsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConversationsGetCall) Context(ctx context.Context) *ProjectsConversationsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsConversationsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsConversationsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.conversations.get" call.
// Exactly one of *GoogleCloudDialogflowV2beta1Conversation or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1Conversation.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsConversationsGetCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1Conversation, error) {
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
	ret := &GoogleCloudDialogflowV2beta1Conversation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the specific conversation.",
	//   "flatPath": "v2beta1/projects/{projectsId}/conversations/{conversationsId}",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.conversations.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the conversation. Format:\n`projects/\u003cProject ID\u003e/conversations/\u003cConversation ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/conversations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Conversation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.conversations.list":

type ProjectsConversationsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Returns the list of all conversations in the specified project.
func (r *ProjectsConversationsService) List(parent string) *ProjectsConversationsListCall {
	c := &ProjectsConversationsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// Filter sets the optional parameter "filter": A filter expression that
// filters conversations listed in the response. In
// general, the expression must specify the field name, a comparison
// operator,
// and the value to use for filtering:
// <ul>
//   <li>The value must be a string, a number, or a boolean.</li>
//   <li>The comparison operator must be either `=`,`!=`, `>`, or
// `<`.</li>
//   <li>To filter on multiple expressions, separate the
//       expressions with `AND` or `OR` (omitting both implies
// `AND`).</li>
//   <li>For clarity, expressions can be enclosed in
// parentheses.</li>
// </ul>
// Only `lifecycle_state` can be filtered on in this way. For
// example,
// the following expression only returns `FINISHED`
// conversations:
//
// `lifecycle_state = "FINISHED"
func (c *ProjectsConversationsListCall) Filter(filter string) *ProjectsConversationsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of items to return in a single page. By
// default 100 and at most 1000.
func (c *ProjectsConversationsListCall) PageSize(pageSize int64) *ProjectsConversationsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous list request.
func (c *ProjectsConversationsListCall) PageToken(pageToken string) *ProjectsConversationsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConversationsListCall) Fields(s ...googleapi.Field) *ProjectsConversationsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsConversationsListCall) IfNoneMatch(entityTag string) *ProjectsConversationsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConversationsListCall) Context(ctx context.Context) *ProjectsConversationsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsConversationsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsConversationsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/conversations")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.conversations.list" call.
// Exactly one of *GoogleCloudDialogflowV2beta1ListConversationsResponse
// or error will be non-nil. Any non-2xx status code is an error.
// Response headers are in either
// *GoogleCloudDialogflowV2beta1ListConversationsResponse.ServerResponse.
// Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsConversationsListCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1ListConversationsResponse, error) {
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
	ret := &GoogleCloudDialogflowV2beta1ListConversationsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the list of all conversations in the specified project.",
	//   "flatPath": "v2beta1/projects/{projectsId}/conversations",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.conversations.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "A filter expression that filters conversations listed in the response. In\ngeneral, the expression must specify the field name, a comparison operator,\nand the value to use for filtering:\n\u003cul\u003e\n  \u003cli\u003eThe value must be a string, a number, or a boolean.\u003c/li\u003e\n  \u003cli\u003eThe comparison operator must be either `=`,`!=`, `\u003e`, or `\u003c`.\u003c/li\u003e\n  \u003cli\u003eTo filter on multiple expressions, separate the\n      expressions with `AND` or `OR` (omitting both implies `AND`).\u003c/li\u003e\n  \u003cli\u003eFor clarity, expressions can be enclosed in parentheses.\u003c/li\u003e\n\u003c/ul\u003e\nOnly `lifecycle_state` can be filtered on in this way. For example,\nthe following expression only returns `FINISHED` conversations:\n\n`lifecycle_state = \"FINISHED\"`",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Optional. The maximum number of items to return in a single page. By\ndefault 100 and at most 1000.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. The next_page_token value returned from a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The project from which to list all conversation.\nFormat: `projects/\u003cProject ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/conversations",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1ListConversationsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsConversationsListCall) Pages(ctx context.Context, f func(*GoogleCloudDialogflowV2beta1ListConversationsResponse) error) error {
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

// method id "dialogflow.projects.conversations.messages.list":

type ProjectsConversationsMessagesListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists messages that belong to a given conversation.
func (r *ProjectsConversationsMessagesService) List(parent string) *ProjectsConversationsMessagesListCall {
	c := &ProjectsConversationsMessagesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of items to return in a single page. By
// default 100 and at most 1000.
func (c *ProjectsConversationsMessagesListCall) PageSize(pageSize int64) *ProjectsConversationsMessagesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous list request.
func (c *ProjectsConversationsMessagesListCall) PageToken(pageToken string) *ProjectsConversationsMessagesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConversationsMessagesListCall) Fields(s ...googleapi.Field) *ProjectsConversationsMessagesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsConversationsMessagesListCall) IfNoneMatch(entityTag string) *ProjectsConversationsMessagesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConversationsMessagesListCall) Context(ctx context.Context) *ProjectsConversationsMessagesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsConversationsMessagesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsConversationsMessagesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/messages")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.conversations.messages.list" call.
// Exactly one of *GoogleCloudDialogflowV2beta1ListMessagesResponse or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1ListMessagesResponse.ServerResponse.Heade
// r or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsConversationsMessagesListCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1ListMessagesResponse, error) {
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
	ret := &GoogleCloudDialogflowV2beta1ListMessagesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists messages that belong to a given conversation.",
	//   "flatPath": "v2beta1/projects/{projectsId}/conversations/{conversationsId}/messages",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.conversations.messages.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Optional. The maximum number of items to return in a single page. By\ndefault 100 and at most 1000.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. The next_page_token value returned from a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The name of the conversation to list messages for.\nFormat: `projects/\u003cProject ID\u003e/conversations/\u003cConversation ID\u003e`",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/conversations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/messages",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1ListMessagesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsConversationsMessagesListCall) Pages(ctx context.Context, f func(*GoogleCloudDialogflowV2beta1ListMessagesResponse) error) error {
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

// method id "dialogflow.projects.conversations.participants.analyzeContent":

type ProjectsConversationsParticipantsAnalyzeContentCall struct {
	s                                                 *Service
	participant                                       string
	googleclouddialogflowv2beta1analyzecontentrequest *GoogleCloudDialogflowV2beta1AnalyzeContentRequest
	urlParams_                                        gensupport.URLParams
	ctx_                                              context.Context
	header_                                           http.Header
}

// AnalyzeContent: Adds a text (chat, for example), or audio (phone
// recording, for example)
// message from a participan  into the conversation.
func (r *ProjectsConversationsParticipantsService) AnalyzeContent(participant string, googleclouddialogflowv2beta1analyzecontentrequest *GoogleCloudDialogflowV2beta1AnalyzeContentRequest) *ProjectsConversationsParticipantsAnalyzeContentCall {
	c := &ProjectsConversationsParticipantsAnalyzeContentCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.participant = participant
	c.googleclouddialogflowv2beta1analyzecontentrequest = googleclouddialogflowv2beta1analyzecontentrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConversationsParticipantsAnalyzeContentCall) Fields(s ...googleapi.Field) *ProjectsConversationsParticipantsAnalyzeContentCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConversationsParticipantsAnalyzeContentCall) Context(ctx context.Context) *ProjectsConversationsParticipantsAnalyzeContentCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsConversationsParticipantsAnalyzeContentCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsConversationsParticipantsAnalyzeContentCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1analyzecontentrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+participant}:analyzeContent")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"participant": c.participant,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.conversations.participants.analyzeContent" call.
// Exactly one of *GoogleCloudDialogflowV2beta1AnalyzeContentResponse or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1AnalyzeContentResponse.ServerResponse.Hea
// der or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsConversationsParticipantsAnalyzeContentCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1AnalyzeContentResponse, error) {
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
	ret := &GoogleCloudDialogflowV2beta1AnalyzeContentResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Adds a text (chat, for example), or audio (phone recording, for example)\nmessage from a participan  into the conversation.",
	//   "flatPath": "v2beta1/projects/{projectsId}/conversations/{conversationsId}/participants/{participantsId}:analyzeContent",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.conversations.participants.analyzeContent",
	//   "parameterOrder": [
	//     "participant"
	//   ],
	//   "parameters": {
	//     "participant": {
	//       "description": "Required. The name of the participant this text comes from.\nFormat: `projects/\u003cProject ID\u003e/conversations/\u003cConversation\nID\u003e/participants/\u003cParticipant ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/conversations/[^/]+/participants/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+participant}:analyzeContent",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1AnalyzeContentRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1AnalyzeContentResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.conversations.participants.create":

type ProjectsConversationsParticipantsCreateCall struct {
	s                                       *Service
	parentid                                string
	googleclouddialogflowv2beta1participant *GoogleCloudDialogflowV2beta1Participant
	urlParams_                              gensupport.URLParams
	ctx_                                    context.Context
	header_                                 http.Header
}

// Create: Creates a new participant in a conversation.
func (r *ProjectsConversationsParticipantsService) Create(parentid string, googleclouddialogflowv2beta1participant *GoogleCloudDialogflowV2beta1Participant) *ProjectsConversationsParticipantsCreateCall {
	c := &ProjectsConversationsParticipantsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parentid = parentid
	c.googleclouddialogflowv2beta1participant = googleclouddialogflowv2beta1participant
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConversationsParticipantsCreateCall) Fields(s ...googleapi.Field) *ProjectsConversationsParticipantsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConversationsParticipantsCreateCall) Context(ctx context.Context) *ProjectsConversationsParticipantsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsConversationsParticipantsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsConversationsParticipantsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1participant)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/participants")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parentid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.conversations.participants.create" call.
// Exactly one of *GoogleCloudDialogflowV2beta1Participant or error will
// be non-nil. Any non-2xx status code is an error. Response headers are
// in either
// *GoogleCloudDialogflowV2beta1Participant.ServerResponse.Header or (if
// a response was returned at all) in error.(*googleapi.Error).Header.
// Use googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsConversationsParticipantsCreateCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1Participant, error) {
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
	ret := &GoogleCloudDialogflowV2beta1Participant{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new participant in a conversation.",
	//   "flatPath": "v2beta1/projects/{projectsId}/conversations/{conversationsId}/participants",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.conversations.participants.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. Resource identifier of the conversation adding the participant.\nFormat: `projects/\u003cProject ID\u003e/conversations/\u003cConversation ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/conversations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/participants",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Participant"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Participant"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.conversations.participants.get":

type ProjectsConversationsParticipantsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves a conversation participant.
func (r *ProjectsConversationsParticipantsService) Get(name string) *ProjectsConversationsParticipantsGetCall {
	c := &ProjectsConversationsParticipantsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConversationsParticipantsGetCall) Fields(s ...googleapi.Field) *ProjectsConversationsParticipantsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsConversationsParticipantsGetCall) IfNoneMatch(entityTag string) *ProjectsConversationsParticipantsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConversationsParticipantsGetCall) Context(ctx context.Context) *ProjectsConversationsParticipantsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsConversationsParticipantsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsConversationsParticipantsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.conversations.participants.get" call.
// Exactly one of *GoogleCloudDialogflowV2beta1Participant or error will
// be non-nil. Any non-2xx status code is an error. Response headers are
// in either
// *GoogleCloudDialogflowV2beta1Participant.ServerResponse.Header or (if
// a response was returned at all) in error.(*googleapi.Error).Header.
// Use googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsConversationsParticipantsGetCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1Participant, error) {
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
	ret := &GoogleCloudDialogflowV2beta1Participant{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a conversation participant.",
	//   "flatPath": "v2beta1/projects/{projectsId}/conversations/{conversationsId}/participants/{participantsId}",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.conversations.participants.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the participant. Format:\n`projects/\u003cProject ID\u003e/conversations/\u003cConversation\nID\u003e/participants/\u003cParticipant ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/conversations/[^/]+/participants/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Participant"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.conversations.participants.list":

type ProjectsConversationsParticipantsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Returns the list of all participants in the specified
// conversation.
func (r *ProjectsConversationsParticipantsService) List(parent string) *ProjectsConversationsParticipantsListCall {
	c := &ProjectsConversationsParticipantsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of items to return in a single page. By
// default 100 and at most 1000.
func (c *ProjectsConversationsParticipantsListCall) PageSize(pageSize int64) *ProjectsConversationsParticipantsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous list request.
func (c *ProjectsConversationsParticipantsListCall) PageToken(pageToken string) *ProjectsConversationsParticipantsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConversationsParticipantsListCall) Fields(s ...googleapi.Field) *ProjectsConversationsParticipantsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsConversationsParticipantsListCall) IfNoneMatch(entityTag string) *ProjectsConversationsParticipantsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConversationsParticipantsListCall) Context(ctx context.Context) *ProjectsConversationsParticipantsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsConversationsParticipantsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsConversationsParticipantsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/participants")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.conversations.participants.list" call.
// Exactly one of *GoogleCloudDialogflowV2beta1ListParticipantsResponse
// or error will be non-nil. Any non-2xx status code is an error.
// Response headers are in either
// *GoogleCloudDialogflowV2beta1ListParticipantsResponse.ServerResponse.H
// eader or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsConversationsParticipantsListCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1ListParticipantsResponse, error) {
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
	ret := &GoogleCloudDialogflowV2beta1ListParticipantsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the list of all participants in the specified conversation.",
	//   "flatPath": "v2beta1/projects/{projectsId}/conversations/{conversationsId}/participants",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.conversations.participants.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Optional. The maximum number of items to return in a single page. By\ndefault 100 and at most 1000.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. The next_page_token value returned from a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The conversation to list all contexts from.\nFormat: `projects/\u003cProject ID\u003e/conversations/\u003cConversation ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/conversations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/participants",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1ListParticipantsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsConversationsParticipantsListCall) Pages(ctx context.Context, f func(*GoogleCloudDialogflowV2beta1ListParticipantsResponse) error) error {
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

// method id "dialogflow.projects.conversations.participants.streamingAnalyzeContent":

type ProjectsConversationsParticipantsStreamingAnalyzeContentCall struct {
	s                                                          *Service
	participant                                                string
	googleclouddialogflowv2beta1streaminganalyzecontentrequest *GoogleCloudDialogflowV2beta1StreamingAnalyzeContentRequest
	urlParams_                                                 gensupport.URLParams
	ctx_                                                       context.Context
	header_                                                    http.Header
}

// StreamingAnalyzeContent: Adds a text (chat, for example), or audio
// (phone recording, for example)
// message from a participan  into the conversation.
// Note: This method is only available through the gRPC API (not
// REST).
//
// The top-level message sent to the client by the server
// is
// `StreamingAnalyzeContentResponse`. Multiple response messages can
// be
// returned in order. The first one or more messages contain
// the
// `recognition_result` field. Each result represents a more
// complete
// transcript of what the user said. The next message contains
// the
// `reply_text` field and potentially the `reply_audio` field. The
// message can
// also contain the `automated_agent_reply` field.
func (r *ProjectsConversationsParticipantsService) StreamingAnalyzeContent(participant string, googleclouddialogflowv2beta1streaminganalyzecontentrequest *GoogleCloudDialogflowV2beta1StreamingAnalyzeContentRequest) *ProjectsConversationsParticipantsStreamingAnalyzeContentCall {
	c := &ProjectsConversationsParticipantsStreamingAnalyzeContentCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.participant = participant
	c.googleclouddialogflowv2beta1streaminganalyzecontentrequest = googleclouddialogflowv2beta1streaminganalyzecontentrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConversationsParticipantsStreamingAnalyzeContentCall) Fields(s ...googleapi.Field) *ProjectsConversationsParticipantsStreamingAnalyzeContentCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConversationsParticipantsStreamingAnalyzeContentCall) Context(ctx context.Context) *ProjectsConversationsParticipantsStreamingAnalyzeContentCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsConversationsParticipantsStreamingAnalyzeContentCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsConversationsParticipantsStreamingAnalyzeContentCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1streaminganalyzecontentrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+participant}:streamingAnalyzeContent")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"participant": c.participant,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.conversations.participants.streamingAnalyzeContent" call.
// Exactly one of
// *GoogleCloudDialogflowV2beta1StreamingAnalyzeContentResponse or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1StreamingAnalyzeContentResponse.ServerRes
// ponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsConversationsParticipantsStreamingAnalyzeContentCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1StreamingAnalyzeContentResponse, error) {
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
	ret := &GoogleCloudDialogflowV2beta1StreamingAnalyzeContentResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Adds a text (chat, for example), or audio (phone recording, for example)\nmessage from a participan  into the conversation.\nNote: This method is only available through the gRPC API (not REST).\n\nThe top-level message sent to the client by the server is\n`StreamingAnalyzeContentResponse`. Multiple response messages can be\nreturned in order. The first one or more messages contain the\n`recognition_result` field. Each result represents a more complete\ntranscript of what the user said. The next message contains the\n`reply_text` field and potentially the `reply_audio` field. The message can\nalso contain the `automated_agent_reply` field.",
	//   "flatPath": "v2beta1/projects/{projectsId}/conversations/{conversationsId}/participants/{participantsId}:streamingAnalyzeContent",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.conversations.participants.streamingAnalyzeContent",
	//   "parameterOrder": [
	//     "participant"
	//   ],
	//   "parameters": {
	//     "participant": {
	//       "description": "Required. The name of the participant this text comes from.\nFormat: `projects/\u003cProject ID\u003e/conversations/\u003cConversation\nID\u003e/participants/\u003cParticipant ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/conversations/[^/]+/participants/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+participant}:streamingAnalyzeContent",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1StreamingAnalyzeContentRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1StreamingAnalyzeContentResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.conversations.participants.suggestions.list":

type ProjectsConversationsParticipantsSuggestionsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Retrieves suggestions for live agents.
//
// This method should be used by human agent client software to
// fetch
// suggestions in real-time, while the conversation with an end user is
// in
// progress. The functionality is implemented in terms of the
// [list
// pagination](/apis/design/design_patterns#list_pagination)
// design pattern. The client app should use the `next_page_token`
// field
// to fetch the next batch of suggestions.
func (r *ProjectsConversationsParticipantsSuggestionsService) List(parent string) *ProjectsConversationsParticipantsSuggestionsListCall {
	c := &ProjectsConversationsParticipantsSuggestionsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of items to return in a single page. The
// default value is 100; the maximum value is 1000.
func (c *ProjectsConversationsParticipantsSuggestionsListCall) PageSize(pageSize int64) *ProjectsConversationsParticipantsSuggestionsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous list request.
func (c *ProjectsConversationsParticipantsSuggestionsListCall) PageToken(pageToken string) *ProjectsConversationsParticipantsSuggestionsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConversationsParticipantsSuggestionsListCall) Fields(s ...googleapi.Field) *ProjectsConversationsParticipantsSuggestionsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsConversationsParticipantsSuggestionsListCall) IfNoneMatch(entityTag string) *ProjectsConversationsParticipantsSuggestionsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConversationsParticipantsSuggestionsListCall) Context(ctx context.Context) *ProjectsConversationsParticipantsSuggestionsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsConversationsParticipantsSuggestionsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsConversationsParticipantsSuggestionsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/suggestions")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.conversations.participants.suggestions.list" call.
// Exactly one of *GoogleCloudDialogflowV2beta1ListSuggestionsResponse
// or error will be non-nil. Any non-2xx status code is an error.
// Response headers are in either
// *GoogleCloudDialogflowV2beta1ListSuggestionsResponse.ServerResponse.He
// ader or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsConversationsParticipantsSuggestionsListCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1ListSuggestionsResponse, error) {
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
	ret := &GoogleCloudDialogflowV2beta1ListSuggestionsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves suggestions for live agents.\n\nThis method should be used by human agent client software to fetch\nsuggestions in real-time, while the conversation with an end user is in\nprogress. The functionality is implemented in terms of the\n[list pagination](/apis/design/design_patterns#list_pagination)\ndesign pattern. The client app should use the `next_page_token` field\nto fetch the next batch of suggestions.",
	//   "flatPath": "v2beta1/projects/{projectsId}/conversations/{conversationsId}/participants/{participantsId}/suggestions",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.conversations.participants.suggestions.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Optional. The maximum number of items to return in a single page. The\ndefault value is 100; the maximum value is 1000.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. The next_page_token value returned from a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The name of the conversation participant for whom to fetch\nsuggestions.\nFormat: `projects/\u003cProject ID\u003e/conversations/\u003cConversation\nID\u003e/participants/\u003cParticipant ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/conversations/[^/]+/participants/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/suggestions",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1ListSuggestionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsConversationsParticipantsSuggestionsListCall) Pages(ctx context.Context, f func(*GoogleCloudDialogflowV2beta1ListSuggestionsResponse) error) error {
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

// method id "dialogflow.projects.environments.users.conversations.deleteContexts":

type ProjectsEnvironmentsUsersConversationsDeleteContextsCall struct {
	s          *Service
	parent     string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// DeleteContexts: Deletes all active contexts in the specified session.
func (r *ProjectsEnvironmentsUsersConversationsService) DeleteContexts(parent string) *ProjectsEnvironmentsUsersConversationsDeleteContextsCall {
	c := &ProjectsEnvironmentsUsersConversationsDeleteContextsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsEnvironmentsUsersConversationsDeleteContextsCall) Fields(s ...googleapi.Field) *ProjectsEnvironmentsUsersConversationsDeleteContextsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsEnvironmentsUsersConversationsDeleteContextsCall) Context(ctx context.Context) *ProjectsEnvironmentsUsersConversationsDeleteContextsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsEnvironmentsUsersConversationsDeleteContextsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsEnvironmentsUsersConversationsDeleteContextsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/contexts")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.environments.users.conversations.deleteContexts" call.
// Exactly one of *GoogleProtobufEmpty or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleProtobufEmpty.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsEnvironmentsUsersConversationsDeleteContextsCall) Do(opts ...googleapi.CallOption) (*GoogleProtobufEmpty, error) {
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
	ret := &GoogleProtobufEmpty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes all active contexts in the specified session.",
	//   "flatPath": "v2beta1/projects/{projectsId}/environments/{environmentsId}/users/{usersId}/conversations/{conversationsId}/contexts",
	//   "httpMethod": "DELETE",
	//   "id": "dialogflow.projects.environments.users.conversations.deleteContexts",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The name of the session to delete all contexts from. Format:\n`projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e` or `projects/\u003cProject\nID\u003e/agent/environments/\u003cEnvironment ID\u003e/users/\u003cUser ID\u003e/sessions/\u003cSession\nID\u003e`. If `Environment ID` is not specified we assume default 'draft'\nenvironment. If `User ID` is not specified, we assume default '-' user.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/environments/[^/]+/users/[^/]+/conversations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/contexts",
	//   "response": {
	//     "$ref": "GoogleProtobufEmpty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.environments.users.conversations.contexts.create":

type ProjectsEnvironmentsUsersConversationsContextsCreateCall struct {
	s                                   *Service
	parent                              string
	googleclouddialogflowv2beta1context *GoogleCloudDialogflowV2beta1Context
	urlParams_                          gensupport.URLParams
	ctx_                                context.Context
	header_                             http.Header
}

// Create: Creates a context.
//
// If the specified context already exists, overrides the context.
func (r *ProjectsEnvironmentsUsersConversationsContextsService) Create(parent string, googleclouddialogflowv2beta1context *GoogleCloudDialogflowV2beta1Context) *ProjectsEnvironmentsUsersConversationsContextsCreateCall {
	c := &ProjectsEnvironmentsUsersConversationsContextsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.googleclouddialogflowv2beta1context = googleclouddialogflowv2beta1context
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsEnvironmentsUsersConversationsContextsCreateCall) Fields(s ...googleapi.Field) *ProjectsEnvironmentsUsersConversationsContextsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsEnvironmentsUsersConversationsContextsCreateCall) Context(ctx context.Context) *ProjectsEnvironmentsUsersConversationsContextsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsEnvironmentsUsersConversationsContextsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsEnvironmentsUsersConversationsContextsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1context)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/contexts")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.environments.users.conversations.contexts.create" call.
// Exactly one of *GoogleCloudDialogflowV2beta1Context or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleCloudDialogflowV2beta1Context.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsEnvironmentsUsersConversationsContextsCreateCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1Context, error) {
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
	ret := &GoogleCloudDialogflowV2beta1Context{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a context.\n\nIf the specified context already exists, overrides the context.",
	//   "flatPath": "v2beta1/projects/{projectsId}/environments/{environmentsId}/users/{usersId}/conversations/{conversationsId}/contexts",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.environments.users.conversations.contexts.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The session to create a context for.\nFormat: `projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e` or\n`projects/\u003cProject ID\u003e/agent/environments/\u003cEnvironment ID\u003e/users/\u003cUser\nID\u003e/sessions/\u003cSession ID\u003e`. If `Environment ID` is not specified, we assume\ndefault 'draft' environment. If `User ID` is not specified, we assume\ndefault '-' user.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/environments/[^/]+/users/[^/]+/conversations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/contexts",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Context"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Context"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.environments.users.conversations.contexts.delete":

type ProjectsEnvironmentsUsersConversationsContextsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes the specified context.
func (r *ProjectsEnvironmentsUsersConversationsContextsService) Delete(name string) *ProjectsEnvironmentsUsersConversationsContextsDeleteCall {
	c := &ProjectsEnvironmentsUsersConversationsContextsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsEnvironmentsUsersConversationsContextsDeleteCall) Fields(s ...googleapi.Field) *ProjectsEnvironmentsUsersConversationsContextsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsEnvironmentsUsersConversationsContextsDeleteCall) Context(ctx context.Context) *ProjectsEnvironmentsUsersConversationsContextsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsEnvironmentsUsersConversationsContextsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsEnvironmentsUsersConversationsContextsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.environments.users.conversations.contexts.delete" call.
// Exactly one of *GoogleProtobufEmpty or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleProtobufEmpty.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsEnvironmentsUsersConversationsContextsDeleteCall) Do(opts ...googleapi.CallOption) (*GoogleProtobufEmpty, error) {
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
	ret := &GoogleProtobufEmpty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the specified context.",
	//   "flatPath": "v2beta1/projects/{projectsId}/environments/{environmentsId}/users/{usersId}/conversations/{conversationsId}/contexts/{contextsId}",
	//   "httpMethod": "DELETE",
	//   "id": "dialogflow.projects.environments.users.conversations.contexts.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the context to delete. Format:\n`projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e/contexts/\u003cContext ID\u003e`\nor `projects/\u003cProject ID\u003e/agent/environments/\u003cEnvironment ID\u003e/users/\u003cUser\nID\u003e/sessions/\u003cSession ID\u003e/contexts/\u003cContext ID\u003e`. If `Environment ID` is\nnot specified, we assume default 'draft' environment. If `User ID` is not\nspecified, we assume default '-' user.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/environments/[^/]+/users/[^/]+/conversations/[^/]+/contexts/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleProtobufEmpty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.environments.users.conversations.contexts.get":

type ProjectsEnvironmentsUsersConversationsContextsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves the specified context.
func (r *ProjectsEnvironmentsUsersConversationsContextsService) Get(name string) *ProjectsEnvironmentsUsersConversationsContextsGetCall {
	c := &ProjectsEnvironmentsUsersConversationsContextsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsEnvironmentsUsersConversationsContextsGetCall) Fields(s ...googleapi.Field) *ProjectsEnvironmentsUsersConversationsContextsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsEnvironmentsUsersConversationsContextsGetCall) IfNoneMatch(entityTag string) *ProjectsEnvironmentsUsersConversationsContextsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsEnvironmentsUsersConversationsContextsGetCall) Context(ctx context.Context) *ProjectsEnvironmentsUsersConversationsContextsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsEnvironmentsUsersConversationsContextsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsEnvironmentsUsersConversationsContextsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.environments.users.conversations.contexts.get" call.
// Exactly one of *GoogleCloudDialogflowV2beta1Context or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleCloudDialogflowV2beta1Context.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsEnvironmentsUsersConversationsContextsGetCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1Context, error) {
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
	ret := &GoogleCloudDialogflowV2beta1Context{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the specified context.",
	//   "flatPath": "v2beta1/projects/{projectsId}/environments/{environmentsId}/users/{usersId}/conversations/{conversationsId}/contexts/{contextsId}",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.environments.users.conversations.contexts.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the context. Format:\n`projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e/contexts/\u003cContext ID\u003e`\nor `projects/\u003cProject ID\u003e/agent/environments/\u003cEnvironment ID\u003e/users/\u003cUser\nID\u003e/sessions/\u003cSession ID\u003e/contexts/\u003cContext ID\u003e`. If `Environment ID` is\nnot specified, we assume default 'draft' environment. If `User ID` is not\nspecified, we assume default '-' user.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/environments/[^/]+/users/[^/]+/conversations/[^/]+/contexts/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Context"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.environments.users.conversations.contexts.list":

type ProjectsEnvironmentsUsersConversationsContextsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Returns the list of all contexts in the specified session.
func (r *ProjectsEnvironmentsUsersConversationsContextsService) List(parent string) *ProjectsEnvironmentsUsersConversationsContextsListCall {
	c := &ProjectsEnvironmentsUsersConversationsContextsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of items to return in a single page. By
// default 100 and at most 1000.
func (c *ProjectsEnvironmentsUsersConversationsContextsListCall) PageSize(pageSize int64) *ProjectsEnvironmentsUsersConversationsContextsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous list request.
func (c *ProjectsEnvironmentsUsersConversationsContextsListCall) PageToken(pageToken string) *ProjectsEnvironmentsUsersConversationsContextsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsEnvironmentsUsersConversationsContextsListCall) Fields(s ...googleapi.Field) *ProjectsEnvironmentsUsersConversationsContextsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsEnvironmentsUsersConversationsContextsListCall) IfNoneMatch(entityTag string) *ProjectsEnvironmentsUsersConversationsContextsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsEnvironmentsUsersConversationsContextsListCall) Context(ctx context.Context) *ProjectsEnvironmentsUsersConversationsContextsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsEnvironmentsUsersConversationsContextsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsEnvironmentsUsersConversationsContextsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/contexts")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.environments.users.conversations.contexts.list" call.
// Exactly one of *GoogleCloudDialogflowV2beta1ListContextsResponse or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1ListContextsResponse.ServerResponse.Heade
// r or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsEnvironmentsUsersConversationsContextsListCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1ListContextsResponse, error) {
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
	ret := &GoogleCloudDialogflowV2beta1ListContextsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the list of all contexts in the specified session.",
	//   "flatPath": "v2beta1/projects/{projectsId}/environments/{environmentsId}/users/{usersId}/conversations/{conversationsId}/contexts",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.environments.users.conversations.contexts.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Optional. The maximum number of items to return in a single page. By\ndefault 100 and at most 1000.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. The next_page_token value returned from a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The session to list all contexts from.\nFormat: `projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e` or\n`projects/\u003cProject ID\u003e/agent/environments/\u003cEnvironment ID\u003e/users/\u003cUser\nID\u003e/sessions/\u003cSession ID\u003e`. If `Environment ID` is not specified, we assume\ndefault 'draft' environment. If `User ID` is not specified, we assume\ndefault '-' user.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/environments/[^/]+/users/[^/]+/conversations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/contexts",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1ListContextsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsEnvironmentsUsersConversationsContextsListCall) Pages(ctx context.Context, f func(*GoogleCloudDialogflowV2beta1ListContextsResponse) error) error {
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

// method id "dialogflow.projects.environments.users.conversations.contexts.patch":

type ProjectsEnvironmentsUsersConversationsContextsPatchCall struct {
	s                                   *Service
	nameid                              string
	googleclouddialogflowv2beta1context *GoogleCloudDialogflowV2beta1Context
	urlParams_                          gensupport.URLParams
	ctx_                                context.Context
	header_                             http.Header
}

// Patch: Updates the specified context.
func (r *ProjectsEnvironmentsUsersConversationsContextsService) Patch(nameid string, googleclouddialogflowv2beta1context *GoogleCloudDialogflowV2beta1Context) *ProjectsEnvironmentsUsersConversationsContextsPatchCall {
	c := &ProjectsEnvironmentsUsersConversationsContextsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.nameid = nameid
	c.googleclouddialogflowv2beta1context = googleclouddialogflowv2beta1context
	return c
}

// UpdateMask sets the optional parameter "updateMask": The mask to
// control which fields get updated.
func (c *ProjectsEnvironmentsUsersConversationsContextsPatchCall) UpdateMask(updateMask string) *ProjectsEnvironmentsUsersConversationsContextsPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsEnvironmentsUsersConversationsContextsPatchCall) Fields(s ...googleapi.Field) *ProjectsEnvironmentsUsersConversationsContextsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsEnvironmentsUsersConversationsContextsPatchCall) Context(ctx context.Context) *ProjectsEnvironmentsUsersConversationsContextsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsEnvironmentsUsersConversationsContextsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsEnvironmentsUsersConversationsContextsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1context)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.nameid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.environments.users.conversations.contexts.patch" call.
// Exactly one of *GoogleCloudDialogflowV2beta1Context or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleCloudDialogflowV2beta1Context.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsEnvironmentsUsersConversationsContextsPatchCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1Context, error) {
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
	ret := &GoogleCloudDialogflowV2beta1Context{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the specified context.",
	//   "flatPath": "v2beta1/projects/{projectsId}/environments/{environmentsId}/users/{usersId}/conversations/{conversationsId}/contexts/{contextsId}",
	//   "httpMethod": "PATCH",
	//   "id": "dialogflow.projects.environments.users.conversations.contexts.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The unique identifier of the context. Format:\n`projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e/contexts/\u003cContext ID\u003e`,\nor `projects/\u003cProject ID\u003e/agent/environments/\u003cEnvironment ID\u003e/users/\u003cUser\nID\u003e/sessions/\u003cSession ID\u003e/contexts/\u003cContext ID\u003e`. The `Context ID` is\nalways converted to lowercase. If `Environment ID` is not specified, we\nassume default 'draft' environment. If `User ID` is not specified, we\nassume default '-' user.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/environments/[^/]+/users/[^/]+/conversations/[^/]+/contexts/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Optional. The mask to control which fields get updated.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Context"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Context"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.humanAgentAssistants.compileSuggestions":

type ProjectsHumanAgentAssistantsCompileSuggestionsCall struct {
	s                                                     *Service
	name                                                  string
	googleclouddialogflowv2beta1compilesuggestionsrequest *GoogleCloudDialogflowV2beta1CompileSuggestionsRequest
	urlParams_                                            gensupport.URLParams
	ctx_                                                  context.Context
	header_                                               http.Header
}

// CompileSuggestions: Uses the specified human agent assistant to come
// up with suggestions
// (relevant articles and FAQs) on how to respond to a given
// conversation.
func (r *ProjectsHumanAgentAssistantsService) CompileSuggestions(name string, googleclouddialogflowv2beta1compilesuggestionsrequest *GoogleCloudDialogflowV2beta1CompileSuggestionsRequest) *ProjectsHumanAgentAssistantsCompileSuggestionsCall {
	c := &ProjectsHumanAgentAssistantsCompileSuggestionsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.googleclouddialogflowv2beta1compilesuggestionsrequest = googleclouddialogflowv2beta1compilesuggestionsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsHumanAgentAssistantsCompileSuggestionsCall) Fields(s ...googleapi.Field) *ProjectsHumanAgentAssistantsCompileSuggestionsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsHumanAgentAssistantsCompileSuggestionsCall) Context(ctx context.Context) *ProjectsHumanAgentAssistantsCompileSuggestionsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsHumanAgentAssistantsCompileSuggestionsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsHumanAgentAssistantsCompileSuggestionsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1compilesuggestionsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}:compileSuggestions")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.humanAgentAssistants.compileSuggestions" call.
// Exactly one of
// *GoogleCloudDialogflowV2beta1CompileSuggestionsResponse or error will
// be non-nil. Any non-2xx status code is an error. Response headers are
// in either
// *GoogleCloudDialogflowV2beta1CompileSuggestionsResponse.ServerResponse
// .Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsHumanAgentAssistantsCompileSuggestionsCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1CompileSuggestionsResponse, error) {
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
	ret := &GoogleCloudDialogflowV2beta1CompileSuggestionsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Uses the specified human agent assistant to come up with suggestions\n(relevant articles and FAQs) on how to respond to a given conversation.",
	//   "flatPath": "v2beta1/projects/{projectsId}/humanAgentAssistants/{humanAgentAssistantsId}:compileSuggestions",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.humanAgentAssistants.compileSuggestions",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The resource name of the agent assistant.\nFormat: `projects/\u003cProject ID\u003e/humanAgentAssistants/\u003cHuman Agent Assistant\nID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/humanAgentAssistants/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}:compileSuggestions",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1CompileSuggestionsRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1CompileSuggestionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.humanAgentAssistants.create":

type ProjectsHumanAgentAssistantsCreateCall struct {
	s                                               *Service
	parent                                          string
	googleclouddialogflowv2beta1humanagentassistant *GoogleCloudDialogflowV2beta1HumanAgentAssistant
	urlParams_                                      gensupport.URLParams
	ctx_                                            context.Context
	header_                                         http.Header
}

// Create: Creates a human agent assistant.
func (r *ProjectsHumanAgentAssistantsService) Create(parent string, googleclouddialogflowv2beta1humanagentassistant *GoogleCloudDialogflowV2beta1HumanAgentAssistant) *ProjectsHumanAgentAssistantsCreateCall {
	c := &ProjectsHumanAgentAssistantsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.googleclouddialogflowv2beta1humanagentassistant = googleclouddialogflowv2beta1humanagentassistant
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsHumanAgentAssistantsCreateCall) Fields(s ...googleapi.Field) *ProjectsHumanAgentAssistantsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsHumanAgentAssistantsCreateCall) Context(ctx context.Context) *ProjectsHumanAgentAssistantsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsHumanAgentAssistantsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsHumanAgentAssistantsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1humanagentassistant)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/humanAgentAssistants")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.humanAgentAssistants.create" call.
// Exactly one of *GoogleCloudDialogflowV2beta1HumanAgentAssistant or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1HumanAgentAssistant.ServerResponse.Header
//  or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsHumanAgentAssistantsCreateCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1HumanAgentAssistant, error) {
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
	ret := &GoogleCloudDialogflowV2beta1HumanAgentAssistant{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a human agent assistant.",
	//   "flatPath": "v2beta1/projects/{projectsId}/humanAgentAssistants",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.humanAgentAssistants.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The project to create a agent assistant for.\nFormat: `projects/\u003cProject ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/humanAgentAssistants",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1HumanAgentAssistant"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1HumanAgentAssistant"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.humanAgentAssistants.delete":

type ProjectsHumanAgentAssistantsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes the specified human agent assistant.
func (r *ProjectsHumanAgentAssistantsService) Delete(name string) *ProjectsHumanAgentAssistantsDeleteCall {
	c := &ProjectsHumanAgentAssistantsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsHumanAgentAssistantsDeleteCall) Fields(s ...googleapi.Field) *ProjectsHumanAgentAssistantsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsHumanAgentAssistantsDeleteCall) Context(ctx context.Context) *ProjectsHumanAgentAssistantsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsHumanAgentAssistantsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsHumanAgentAssistantsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.humanAgentAssistants.delete" call.
// Exactly one of *GoogleProtobufEmpty or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleProtobufEmpty.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsHumanAgentAssistantsDeleteCall) Do(opts ...googleapi.CallOption) (*GoogleProtobufEmpty, error) {
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
	ret := &GoogleProtobufEmpty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the specified human agent assistant.",
	//   "flatPath": "v2beta1/projects/{projectsId}/humanAgentAssistants/{humanAgentAssistantsId}",
	//   "httpMethod": "DELETE",
	//   "id": "dialogflow.projects.humanAgentAssistants.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The resource name of the agent assistant.\nFormat: `projects/\u003cProject ID\u003e/humanAgentAssistants/\u003cHuman Agent Assistant\nID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/humanAgentAssistants/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleProtobufEmpty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.humanAgentAssistants.get":

type ProjectsHumanAgentAssistantsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves a human agent assistant.
func (r *ProjectsHumanAgentAssistantsService) Get(name string) *ProjectsHumanAgentAssistantsGetCall {
	c := &ProjectsHumanAgentAssistantsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsHumanAgentAssistantsGetCall) Fields(s ...googleapi.Field) *ProjectsHumanAgentAssistantsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsHumanAgentAssistantsGetCall) IfNoneMatch(entityTag string) *ProjectsHumanAgentAssistantsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsHumanAgentAssistantsGetCall) Context(ctx context.Context) *ProjectsHumanAgentAssistantsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsHumanAgentAssistantsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsHumanAgentAssistantsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.humanAgentAssistants.get" call.
// Exactly one of *GoogleCloudDialogflowV2beta1HumanAgentAssistant or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1HumanAgentAssistant.ServerResponse.Header
//  or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsHumanAgentAssistantsGetCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1HumanAgentAssistant, error) {
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
	ret := &GoogleCloudDialogflowV2beta1HumanAgentAssistant{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a human agent assistant.",
	//   "flatPath": "v2beta1/projects/{projectsId}/humanAgentAssistants/{humanAgentAssistantsId}",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.humanAgentAssistants.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The resource name of the agent assistant.\nFormat: `projects/\u003cProject ID\u003e/humanAgentAssistants/\u003cHuman Agent Assistant\nID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/humanAgentAssistants/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1HumanAgentAssistant"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.humanAgentAssistants.list":

type ProjectsHumanAgentAssistantsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Returns the list of all human agent assistants.
func (r *ProjectsHumanAgentAssistantsService) List(parent string) *ProjectsHumanAgentAssistantsListCall {
	c := &ProjectsHumanAgentAssistantsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of items to return in a single page.
// The default value is 100; the maximum value is 1000.
func (c *ProjectsHumanAgentAssistantsListCall) PageSize(pageSize int64) *ProjectsHumanAgentAssistantsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous list request.
func (c *ProjectsHumanAgentAssistantsListCall) PageToken(pageToken string) *ProjectsHumanAgentAssistantsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsHumanAgentAssistantsListCall) Fields(s ...googleapi.Field) *ProjectsHumanAgentAssistantsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsHumanAgentAssistantsListCall) IfNoneMatch(entityTag string) *ProjectsHumanAgentAssistantsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsHumanAgentAssistantsListCall) Context(ctx context.Context) *ProjectsHumanAgentAssistantsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsHumanAgentAssistantsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsHumanAgentAssistantsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/humanAgentAssistants")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.humanAgentAssistants.list" call.
// Exactly one of
// *GoogleCloudDialogflowV2beta1ListHumanAgentAssistantsResponse or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1ListHumanAgentAssistantsResponse.ServerRe
// sponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsHumanAgentAssistantsListCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1ListHumanAgentAssistantsResponse, error) {
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
	ret := &GoogleCloudDialogflowV2beta1ListHumanAgentAssistantsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the list of all human agent assistants.",
	//   "flatPath": "v2beta1/projects/{projectsId}/humanAgentAssistants",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.humanAgentAssistants.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Optional. The maximum number of items to return in a single page.\nThe default value is 100; the maximum value is 1000.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. The next_page_token value returned from a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The project to list all agent assistants from.\nFormat: `projects/\u003cProject ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/humanAgentAssistants",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1ListHumanAgentAssistantsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsHumanAgentAssistantsListCall) Pages(ctx context.Context, f func(*GoogleCloudDialogflowV2beta1ListHumanAgentAssistantsResponse) error) error {
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

// method id "dialogflow.projects.humanAgentAssistants.patch":

type ProjectsHumanAgentAssistantsPatchCall struct {
	s                                               *Service
	nameid                                          string
	googleclouddialogflowv2beta1humanagentassistant *GoogleCloudDialogflowV2beta1HumanAgentAssistant
	urlParams_                                      gensupport.URLParams
	ctx_                                            context.Context
	header_                                         http.Header
}

// Patch: Updates the specified human agent assistant.
func (r *ProjectsHumanAgentAssistantsService) Patch(nameid string, googleclouddialogflowv2beta1humanagentassistant *GoogleCloudDialogflowV2beta1HumanAgentAssistant) *ProjectsHumanAgentAssistantsPatchCall {
	c := &ProjectsHumanAgentAssistantsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.nameid = nameid
	c.googleclouddialogflowv2beta1humanagentassistant = googleclouddialogflowv2beta1humanagentassistant
	return c
}

// UpdateMask sets the optional parameter "updateMask": The mask to
// specify which fields to update.
func (c *ProjectsHumanAgentAssistantsPatchCall) UpdateMask(updateMask string) *ProjectsHumanAgentAssistantsPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsHumanAgentAssistantsPatchCall) Fields(s ...googleapi.Field) *ProjectsHumanAgentAssistantsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsHumanAgentAssistantsPatchCall) Context(ctx context.Context) *ProjectsHumanAgentAssistantsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsHumanAgentAssistantsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsHumanAgentAssistantsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1humanagentassistant)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.nameid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.humanAgentAssistants.patch" call.
// Exactly one of *GoogleCloudDialogflowV2beta1HumanAgentAssistant or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1HumanAgentAssistant.ServerResponse.Header
//  or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsHumanAgentAssistantsPatchCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1HumanAgentAssistant, error) {
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
	ret := &GoogleCloudDialogflowV2beta1HumanAgentAssistant{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the specified human agent assistant.",
	//   "flatPath": "v2beta1/projects/{projectsId}/humanAgentAssistants/{humanAgentAssistantsId}",
	//   "httpMethod": "PATCH",
	//   "id": "dialogflow.projects.humanAgentAssistants.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required for all methods except `create` (`create` populates the name\nautomatically).\nThe unique identifier of human agent assistant.\nFormat: `projects/\u003cProject ID\u003e/humanAgentAssistants/\u003cHuman Agent Assistant\nID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/humanAgentAssistants/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Optional. The mask to specify which fields to update.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1HumanAgentAssistant"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1HumanAgentAssistant"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.knowledgeBases.create":

type ProjectsKnowledgeBasesCreateCall struct {
	s                                         *Service
	parent                                    string
	googleclouddialogflowv2beta1knowledgebase *GoogleCloudDialogflowV2beta1KnowledgeBase
	urlParams_                                gensupport.URLParams
	ctx_                                      context.Context
	header_                                   http.Header
}

// Create: Creates a knowledge base.
func (r *ProjectsKnowledgeBasesService) Create(parent string, googleclouddialogflowv2beta1knowledgebase *GoogleCloudDialogflowV2beta1KnowledgeBase) *ProjectsKnowledgeBasesCreateCall {
	c := &ProjectsKnowledgeBasesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.googleclouddialogflowv2beta1knowledgebase = googleclouddialogflowv2beta1knowledgebase
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsKnowledgeBasesCreateCall) Fields(s ...googleapi.Field) *ProjectsKnowledgeBasesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsKnowledgeBasesCreateCall) Context(ctx context.Context) *ProjectsKnowledgeBasesCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsKnowledgeBasesCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsKnowledgeBasesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1knowledgebase)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/knowledgeBases")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.knowledgeBases.create" call.
// Exactly one of *GoogleCloudDialogflowV2beta1KnowledgeBase or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1KnowledgeBase.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsKnowledgeBasesCreateCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1KnowledgeBase, error) {
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
	ret := &GoogleCloudDialogflowV2beta1KnowledgeBase{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a knowledge base.",
	//   "flatPath": "v2beta1/projects/{projectsId}/knowledgeBases",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.knowledgeBases.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The project to create a knowledge base for.\nFormat: `projects/\u003cProject ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/knowledgeBases",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1KnowledgeBase"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1KnowledgeBase"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.knowledgeBases.delete":

type ProjectsKnowledgeBasesDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes the specified knowledge base.
func (r *ProjectsKnowledgeBasesService) Delete(name string) *ProjectsKnowledgeBasesDeleteCall {
	c := &ProjectsKnowledgeBasesDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Force sets the optional parameter "force": Force deletes the
// knowledge base. When set to true, any documents
// in the knowledge base are also deleted.
func (c *ProjectsKnowledgeBasesDeleteCall) Force(force bool) *ProjectsKnowledgeBasesDeleteCall {
	c.urlParams_.Set("force", fmt.Sprint(force))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsKnowledgeBasesDeleteCall) Fields(s ...googleapi.Field) *ProjectsKnowledgeBasesDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsKnowledgeBasesDeleteCall) Context(ctx context.Context) *ProjectsKnowledgeBasesDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsKnowledgeBasesDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsKnowledgeBasesDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.knowledgeBases.delete" call.
// Exactly one of *GoogleProtobufEmpty or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleProtobufEmpty.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsKnowledgeBasesDeleteCall) Do(opts ...googleapi.CallOption) (*GoogleProtobufEmpty, error) {
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
	ret := &GoogleProtobufEmpty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the specified knowledge base.",
	//   "flatPath": "v2beta1/projects/{projectsId}/knowledgeBases/{knowledgeBasesId}",
	//   "httpMethod": "DELETE",
	//   "id": "dialogflow.projects.knowledgeBases.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "force": {
	//       "description": "Optional. Force deletes the knowledge base. When set to true, any documents\nin the knowledge base are also deleted.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "name": {
	//       "description": "Required. The name of the knowledge base to delete.\nFormat: `projects/\u003cProject ID\u003e/knowledgeBases/\u003cKnowledge Base ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/knowledgeBases/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleProtobufEmpty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.knowledgeBases.get":

type ProjectsKnowledgeBasesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves the specified knowledge base.
func (r *ProjectsKnowledgeBasesService) Get(name string) *ProjectsKnowledgeBasesGetCall {
	c := &ProjectsKnowledgeBasesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsKnowledgeBasesGetCall) Fields(s ...googleapi.Field) *ProjectsKnowledgeBasesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsKnowledgeBasesGetCall) IfNoneMatch(entityTag string) *ProjectsKnowledgeBasesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsKnowledgeBasesGetCall) Context(ctx context.Context) *ProjectsKnowledgeBasesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsKnowledgeBasesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsKnowledgeBasesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.knowledgeBases.get" call.
// Exactly one of *GoogleCloudDialogflowV2beta1KnowledgeBase or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1KnowledgeBase.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsKnowledgeBasesGetCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1KnowledgeBase, error) {
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
	ret := &GoogleCloudDialogflowV2beta1KnowledgeBase{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the specified knowledge base.",
	//   "flatPath": "v2beta1/projects/{projectsId}/knowledgeBases/{knowledgeBasesId}",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.knowledgeBases.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the knowledge base to retrieve.\nFormat `projects/\u003cProject ID\u003e/knowledgeBases/\u003cKnowledge Base ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/knowledgeBases/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1KnowledgeBase"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.knowledgeBases.list":

type ProjectsKnowledgeBasesListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Returns the list of all knowledge bases of the specified agent.
func (r *ProjectsKnowledgeBasesService) List(parent string) *ProjectsKnowledgeBasesListCall {
	c := &ProjectsKnowledgeBasesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of items to return in a single page. By
// default 10 and at most 100.
func (c *ProjectsKnowledgeBasesListCall) PageSize(pageSize int64) *ProjectsKnowledgeBasesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous list request.
func (c *ProjectsKnowledgeBasesListCall) PageToken(pageToken string) *ProjectsKnowledgeBasesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsKnowledgeBasesListCall) Fields(s ...googleapi.Field) *ProjectsKnowledgeBasesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsKnowledgeBasesListCall) IfNoneMatch(entityTag string) *ProjectsKnowledgeBasesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsKnowledgeBasesListCall) Context(ctx context.Context) *ProjectsKnowledgeBasesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsKnowledgeBasesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsKnowledgeBasesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/knowledgeBases")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.knowledgeBases.list" call.
// Exactly one of
// *GoogleCloudDialogflowV2beta1ListKnowledgeBasesResponse or error will
// be non-nil. Any non-2xx status code is an error. Response headers are
// in either
// *GoogleCloudDialogflowV2beta1ListKnowledgeBasesResponse.ServerResponse
// .Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsKnowledgeBasesListCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1ListKnowledgeBasesResponse, error) {
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
	ret := &GoogleCloudDialogflowV2beta1ListKnowledgeBasesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the list of all knowledge bases of the specified agent.",
	//   "flatPath": "v2beta1/projects/{projectsId}/knowledgeBases",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.knowledgeBases.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Optional. The maximum number of items to return in a single page. By\ndefault 10 and at most 100.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. The next_page_token value returned from a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The project to list of knowledge bases for.\nFormat: `projects/\u003cProject ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/knowledgeBases",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1ListKnowledgeBasesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsKnowledgeBasesListCall) Pages(ctx context.Context, f func(*GoogleCloudDialogflowV2beta1ListKnowledgeBasesResponse) error) error {
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

// method id "dialogflow.projects.knowledgeBases.documents.create":

type ProjectsKnowledgeBasesDocumentsCreateCall struct {
	s                                    *Service
	parent                               string
	googleclouddialogflowv2beta1document *GoogleCloudDialogflowV2beta1Document
	urlParams_                           gensupport.URLParams
	ctx_                                 context.Context
	header_                              http.Header
}

// Create: Creates a new document.
//
// Operation <response: Document,
//            metadata: KnowledgeOperationMetadata>
func (r *ProjectsKnowledgeBasesDocumentsService) Create(parent string, googleclouddialogflowv2beta1document *GoogleCloudDialogflowV2beta1Document) *ProjectsKnowledgeBasesDocumentsCreateCall {
	c := &ProjectsKnowledgeBasesDocumentsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.googleclouddialogflowv2beta1document = googleclouddialogflowv2beta1document
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsKnowledgeBasesDocumentsCreateCall) Fields(s ...googleapi.Field) *ProjectsKnowledgeBasesDocumentsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsKnowledgeBasesDocumentsCreateCall) Context(ctx context.Context) *ProjectsKnowledgeBasesDocumentsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsKnowledgeBasesDocumentsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsKnowledgeBasesDocumentsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1document)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/documents")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.knowledgeBases.documents.create" call.
// Exactly one of *GoogleLongrunningOperation or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *GoogleLongrunningOperation.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsKnowledgeBasesDocumentsCreateCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunningOperation, error) {
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
	ret := &GoogleLongrunningOperation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new document.\n\nOperation \u003cresponse: Document,\n           metadata: KnowledgeOperationMetadata\u003e",
	//   "flatPath": "v2beta1/projects/{projectsId}/knowledgeBases/{knowledgeBasesId}/documents",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.knowledgeBases.documents.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The knoweldge base to create a document for.\nFormat: `projects/\u003cProject ID\u003e/knowledgeBases/\u003cKnowledge Base ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/knowledgeBases/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/documents",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Document"
	//   },
	//   "response": {
	//     "$ref": "GoogleLongrunningOperation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.knowledgeBases.documents.delete":

type ProjectsKnowledgeBasesDocumentsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes the specified document.
//
// Operation <response: google.protobuf.Empty,
//            metadata: KnowledgeOperationMetadata>
func (r *ProjectsKnowledgeBasesDocumentsService) Delete(name string) *ProjectsKnowledgeBasesDocumentsDeleteCall {
	c := &ProjectsKnowledgeBasesDocumentsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsKnowledgeBasesDocumentsDeleteCall) Fields(s ...googleapi.Field) *ProjectsKnowledgeBasesDocumentsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsKnowledgeBasesDocumentsDeleteCall) Context(ctx context.Context) *ProjectsKnowledgeBasesDocumentsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsKnowledgeBasesDocumentsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsKnowledgeBasesDocumentsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.knowledgeBases.documents.delete" call.
// Exactly one of *GoogleLongrunningOperation or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *GoogleLongrunningOperation.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsKnowledgeBasesDocumentsDeleteCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunningOperation, error) {
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
	ret := &GoogleLongrunningOperation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the specified document.\n\nOperation \u003cresponse: google.protobuf.Empty,\n           metadata: KnowledgeOperationMetadata\u003e",
	//   "flatPath": "v2beta1/projects/{projectsId}/knowledgeBases/{knowledgeBasesId}/documents/{documentsId}",
	//   "httpMethod": "DELETE",
	//   "id": "dialogflow.projects.knowledgeBases.documents.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the document to delete.\nFormat: `projects/\u003cProject ID\u003e/knowledgeBases/\u003cKnowledge Base\nID\u003e/documents/\u003cDocument ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/knowledgeBases/[^/]+/documents/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleLongrunningOperation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.knowledgeBases.documents.get":

type ProjectsKnowledgeBasesDocumentsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves the specified document.
func (r *ProjectsKnowledgeBasesDocumentsService) Get(name string) *ProjectsKnowledgeBasesDocumentsGetCall {
	c := &ProjectsKnowledgeBasesDocumentsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsKnowledgeBasesDocumentsGetCall) Fields(s ...googleapi.Field) *ProjectsKnowledgeBasesDocumentsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsKnowledgeBasesDocumentsGetCall) IfNoneMatch(entityTag string) *ProjectsKnowledgeBasesDocumentsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsKnowledgeBasesDocumentsGetCall) Context(ctx context.Context) *ProjectsKnowledgeBasesDocumentsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsKnowledgeBasesDocumentsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsKnowledgeBasesDocumentsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.knowledgeBases.documents.get" call.
// Exactly one of *GoogleCloudDialogflowV2beta1Document or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleCloudDialogflowV2beta1Document.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsKnowledgeBasesDocumentsGetCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1Document, error) {
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
	ret := &GoogleCloudDialogflowV2beta1Document{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the specified document.",
	//   "flatPath": "v2beta1/projects/{projectsId}/knowledgeBases/{knowledgeBasesId}/documents/{documentsId}",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.knowledgeBases.documents.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the document to retrieve.\nFormat `projects/\u003cProject ID\u003e/knowledgeBases/\u003cKnowledge Base\nID\u003e/documents/\u003cDocument ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/knowledgeBases/[^/]+/documents/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1Document"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.knowledgeBases.documents.list":

type ProjectsKnowledgeBasesDocumentsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Returns the list of all documents of the knowledge base.
func (r *ProjectsKnowledgeBasesDocumentsService) List(parent string) *ProjectsKnowledgeBasesDocumentsListCall {
	c := &ProjectsKnowledgeBasesDocumentsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of items to return in a single page. By
// default 10 and at most 100.
func (c *ProjectsKnowledgeBasesDocumentsListCall) PageSize(pageSize int64) *ProjectsKnowledgeBasesDocumentsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous list request.
func (c *ProjectsKnowledgeBasesDocumentsListCall) PageToken(pageToken string) *ProjectsKnowledgeBasesDocumentsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsKnowledgeBasesDocumentsListCall) Fields(s ...googleapi.Field) *ProjectsKnowledgeBasesDocumentsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsKnowledgeBasesDocumentsListCall) IfNoneMatch(entityTag string) *ProjectsKnowledgeBasesDocumentsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsKnowledgeBasesDocumentsListCall) Context(ctx context.Context) *ProjectsKnowledgeBasesDocumentsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsKnowledgeBasesDocumentsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsKnowledgeBasesDocumentsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/documents")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.knowledgeBases.documents.list" call.
// Exactly one of *GoogleCloudDialogflowV2beta1ListDocumentsResponse or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1ListDocumentsResponse.ServerResponse.Head
// er or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsKnowledgeBasesDocumentsListCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1ListDocumentsResponse, error) {
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
	ret := &GoogleCloudDialogflowV2beta1ListDocumentsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the list of all documents of the knowledge base.",
	//   "flatPath": "v2beta1/projects/{projectsId}/knowledgeBases/{knowledgeBasesId}/documents",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.knowledgeBases.documents.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Optional. The maximum number of items to return in a single page. By\ndefault 10 and at most 100.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. The next_page_token value returned from a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The knowledge base to list all documents for.\nFormat: `projects/\u003cProject ID\u003e/knowledgeBases/\u003cKnowledge Base ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/knowledgeBases/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/documents",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1ListDocumentsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsKnowledgeBasesDocumentsListCall) Pages(ctx context.Context, f func(*GoogleCloudDialogflowV2beta1ListDocumentsResponse) error) error {
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

// method id "dialogflow.projects.operations.get":

type ProjectsOperationsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets the latest state of a long-running operation.  Clients can
// use this
// method to poll the operation result at intervals as recommended by
// the API
// service.
func (r *ProjectsOperationsService) Get(name string) *ProjectsOperationsGetCall {
	c := &ProjectsOperationsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOperationsGetCall) Fields(s ...googleapi.Field) *ProjectsOperationsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsOperationsGetCall) IfNoneMatch(entityTag string) *ProjectsOperationsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsOperationsGetCall) Context(ctx context.Context) *ProjectsOperationsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsOperationsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsOperationsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.operations.get" call.
// Exactly one of *GoogleLongrunningOperation or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *GoogleLongrunningOperation.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsOperationsGetCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunningOperation, error) {
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
	ret := &GoogleLongrunningOperation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the latest state of a long-running operation.  Clients can use this\nmethod to poll the operation result at intervals as recommended by the API\nservice.",
	//   "flatPath": "v2beta1/projects/{projectsId}/operations/{operationsId}",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.operations.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/operations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleLongrunningOperation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.phoneNumberOrders.cancel":

type ProjectsPhoneNumberOrdersCancelCall struct {
	s                                                         *Service
	nameid                                                    string
	googleclouddialogflowv2beta1cancelphonenumberorderrequest *GoogleCloudDialogflowV2beta1CancelPhoneNumberOrderRequest
	urlParams_                                                gensupport.URLParams
	ctx_                                                      context.Context
	header_                                                   http.Header
}

// Cancel: Cancels an `PhoneNumberOrder`.
// Returns an error if the order is in state
// IN_PROGRESS or
// COMPLETED.
func (r *ProjectsPhoneNumberOrdersService) Cancel(nameid string, googleclouddialogflowv2beta1cancelphonenumberorderrequest *GoogleCloudDialogflowV2beta1CancelPhoneNumberOrderRequest) *ProjectsPhoneNumberOrdersCancelCall {
	c := &ProjectsPhoneNumberOrdersCancelCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.nameid = nameid
	c.googleclouddialogflowv2beta1cancelphonenumberorderrequest = googleclouddialogflowv2beta1cancelphonenumberorderrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsPhoneNumberOrdersCancelCall) Fields(s ...googleapi.Field) *ProjectsPhoneNumberOrdersCancelCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsPhoneNumberOrdersCancelCall) Context(ctx context.Context) *ProjectsPhoneNumberOrdersCancelCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsPhoneNumberOrdersCancelCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsPhoneNumberOrdersCancelCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1cancelphonenumberorderrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}:cancel")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.nameid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.phoneNumberOrders.cancel" call.
// Exactly one of *GoogleProtobufEmpty or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleProtobufEmpty.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsPhoneNumberOrdersCancelCall) Do(opts ...googleapi.CallOption) (*GoogleProtobufEmpty, error) {
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
	ret := &GoogleProtobufEmpty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Cancels an `PhoneNumberOrder`.\nReturns an error if the order is in state\nIN_PROGRESS or\nCOMPLETED.",
	//   "flatPath": "v2beta1/projects/{projectsId}/phoneNumberOrders/{phoneNumberOrdersId}:cancel",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.phoneNumberOrders.cancel",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The unique identifier of the order to delete.\nFormat: `projects/\u003cProject ID\u003e/phoneNumberOrders/\u003cOrder ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/phoneNumberOrders/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}:cancel",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1CancelPhoneNumberOrderRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleProtobufEmpty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.phoneNumberOrders.create":

type ProjectsPhoneNumberOrdersCreateCall struct {
	s                                            *Service
	parentid                                     string
	googleclouddialogflowv2beta1phonenumberorder *GoogleCloudDialogflowV2beta1PhoneNumberOrder
	urlParams_                                   gensupport.URLParams
	ctx_                                         context.Context
	header_                                      http.Header
}

// Create: Creates an order to request phone numbers be added to a
// project.
// The initial `LifecycleState` of a newly created order is
// PENDING.
func (r *ProjectsPhoneNumberOrdersService) Create(parentid string, googleclouddialogflowv2beta1phonenumberorder *GoogleCloudDialogflowV2beta1PhoneNumberOrder) *ProjectsPhoneNumberOrdersCreateCall {
	c := &ProjectsPhoneNumberOrdersCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parentid = parentid
	c.googleclouddialogflowv2beta1phonenumberorder = googleclouddialogflowv2beta1phonenumberorder
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsPhoneNumberOrdersCreateCall) Fields(s ...googleapi.Field) *ProjectsPhoneNumberOrdersCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsPhoneNumberOrdersCreateCall) Context(ctx context.Context) *ProjectsPhoneNumberOrdersCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsPhoneNumberOrdersCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsPhoneNumberOrdersCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1phonenumberorder)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/phoneNumberOrders")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parentid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.phoneNumberOrders.create" call.
// Exactly one of *GoogleCloudDialogflowV2beta1PhoneNumberOrder or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1PhoneNumberOrder.ServerResponse.Header
// or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsPhoneNumberOrdersCreateCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1PhoneNumberOrder, error) {
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
	ret := &GoogleCloudDialogflowV2beta1PhoneNumberOrder{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates an order to request phone numbers be added to a project.\nThe initial `LifecycleState` of a newly created order is\nPENDING.",
	//   "flatPath": "v2beta1/projects/{projectsId}/phoneNumberOrders",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.phoneNumberOrders.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. Resource identifier of the project requesting the orders.\nFormat: `projects/\u003cProject ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/phoneNumberOrders",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1PhoneNumberOrder"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1PhoneNumberOrder"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.phoneNumberOrders.get":

type ProjectsPhoneNumberOrdersGetCall struct {
	s            *Service
	nameid       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Returns a specific `PhoneNumberOrder`.
func (r *ProjectsPhoneNumberOrdersService) Get(nameid string) *ProjectsPhoneNumberOrdersGetCall {
	c := &ProjectsPhoneNumberOrdersGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.nameid = nameid
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsPhoneNumberOrdersGetCall) Fields(s ...googleapi.Field) *ProjectsPhoneNumberOrdersGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsPhoneNumberOrdersGetCall) IfNoneMatch(entityTag string) *ProjectsPhoneNumberOrdersGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsPhoneNumberOrdersGetCall) Context(ctx context.Context) *ProjectsPhoneNumberOrdersGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsPhoneNumberOrdersGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsPhoneNumberOrdersGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.nameid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.phoneNumberOrders.get" call.
// Exactly one of *GoogleCloudDialogflowV2beta1PhoneNumberOrder or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1PhoneNumberOrder.ServerResponse.Header
// or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsPhoneNumberOrdersGetCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1PhoneNumberOrder, error) {
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
	ret := &GoogleCloudDialogflowV2beta1PhoneNumberOrder{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns a specific `PhoneNumberOrder`.",
	//   "flatPath": "v2beta1/projects/{projectsId}/phoneNumberOrders/{phoneNumberOrdersId}",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.phoneNumberOrders.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The unique identifier of the order to retrieve.\nFormat: `projects/\u003cProject ID\u003e/phoneNumberOrders/\u003cOrder ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/phoneNumberOrders/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1PhoneNumberOrder"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.phoneNumberOrders.list":

type ProjectsPhoneNumberOrdersListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists of all `PhoneNumberOrder` resources in the specified
// project.
func (r *ProjectsPhoneNumberOrdersService) List(parent string) *ProjectsPhoneNumberOrdersListCall {
	c := &ProjectsPhoneNumberOrdersListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of items to return in a single page.
// The default value is 100. The maximum value is 1000.
func (c *ProjectsPhoneNumberOrdersListCall) PageSize(pageSize int64) *ProjectsPhoneNumberOrdersListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous list request.
func (c *ProjectsPhoneNumberOrdersListCall) PageToken(pageToken string) *ProjectsPhoneNumberOrdersListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsPhoneNumberOrdersListCall) Fields(s ...googleapi.Field) *ProjectsPhoneNumberOrdersListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsPhoneNumberOrdersListCall) IfNoneMatch(entityTag string) *ProjectsPhoneNumberOrdersListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsPhoneNumberOrdersListCall) Context(ctx context.Context) *ProjectsPhoneNumberOrdersListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsPhoneNumberOrdersListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsPhoneNumberOrdersListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/phoneNumberOrders")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.phoneNumberOrders.list" call.
// Exactly one of
// *GoogleCloudDialogflowV2beta1ListPhoneNumberOrdersResponse or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1ListPhoneNumberOrdersResponse.ServerRespo
// nse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsPhoneNumberOrdersListCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1ListPhoneNumberOrdersResponse, error) {
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
	ret := &GoogleCloudDialogflowV2beta1ListPhoneNumberOrdersResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists of all `PhoneNumberOrder` resources in the specified project.",
	//   "flatPath": "v2beta1/projects/{projectsId}/phoneNumberOrders",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.phoneNumberOrders.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Optional. The maximum number of items to return in a single page.\nThe default value is 100. The maximum value is 1000.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. The next_page_token value returned from a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The project to list all orders from.\nFormat: `projects/\u003cProject ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/phoneNumberOrders",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1ListPhoneNumberOrdersResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsPhoneNumberOrdersListCall) Pages(ctx context.Context, f func(*GoogleCloudDialogflowV2beta1ListPhoneNumberOrdersResponse) error) error {
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

// method id "dialogflow.projects.phoneNumberOrders.patch":

type ProjectsPhoneNumberOrdersPatchCall struct {
	s                                            *Service
	nameid                                       string
	googleclouddialogflowv2beta1phonenumberorder *GoogleCloudDialogflowV2beta1PhoneNumberOrder
	urlParams_                                   gensupport.URLParams
	ctx_                                         context.Context
	header_                                      http.Header
}

// Patch: Updates the specified `PhoneNumberOrder` resource.
// Returns an error if the order is in state
// IN_PROGRESS or
// COMPLETED.
func (r *ProjectsPhoneNumberOrdersService) Patch(nameid string, googleclouddialogflowv2beta1phonenumberorder *GoogleCloudDialogflowV2beta1PhoneNumberOrder) *ProjectsPhoneNumberOrdersPatchCall {
	c := &ProjectsPhoneNumberOrdersPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.nameid = nameid
	c.googleclouddialogflowv2beta1phonenumberorder = googleclouddialogflowv2beta1phonenumberorder
	return c
}

// UpdateMask sets the optional parameter "updateMask": The mask to
// control which fields get updated.
func (c *ProjectsPhoneNumberOrdersPatchCall) UpdateMask(updateMask string) *ProjectsPhoneNumberOrdersPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsPhoneNumberOrdersPatchCall) Fields(s ...googleapi.Field) *ProjectsPhoneNumberOrdersPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsPhoneNumberOrdersPatchCall) Context(ctx context.Context) *ProjectsPhoneNumberOrdersPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsPhoneNumberOrdersPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsPhoneNumberOrdersPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1phonenumberorder)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.nameid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.phoneNumberOrders.patch" call.
// Exactly one of *GoogleCloudDialogflowV2beta1PhoneNumberOrder or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudDialogflowV2beta1PhoneNumberOrder.ServerResponse.Header
// or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsPhoneNumberOrdersPatchCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1PhoneNumberOrder, error) {
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
	ret := &GoogleCloudDialogflowV2beta1PhoneNumberOrder{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the specified `PhoneNumberOrder` resource.\nReturns an error if the order is in state\nIN_PROGRESS or\nCOMPLETED.",
	//   "flatPath": "v2beta1/projects/{projectsId}/phoneNumberOrders/{phoneNumberOrdersId}",
	//   "httpMethod": "PATCH",
	//   "id": "dialogflow.projects.phoneNumberOrders.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The unique identifier of this order.\nFormat: `projects/\u003cProject ID\u003e/phoneNumberOrders/\u003cOrder ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/phoneNumberOrders/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Optional. The mask to control which fields get updated.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1PhoneNumberOrder"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1PhoneNumberOrder"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.phoneNumbers.delete":

type ProjectsPhoneNumbersDeleteCall struct {
	s          *Service
	nameid     string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Requests deletion of a `PhoneNumber`. The `PhoneNumber` is
// moved into the
// DELETE_REQUESTED state
// immediately, and is deleted approximately 30 days later. This method
// may
// only be called on a `PhoneNumber` in the
// ACTIVE state.
func (r *ProjectsPhoneNumbersService) Delete(nameid string) *ProjectsPhoneNumbersDeleteCall {
	c := &ProjectsPhoneNumbersDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.nameid = nameid
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsPhoneNumbersDeleteCall) Fields(s ...googleapi.Field) *ProjectsPhoneNumbersDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsPhoneNumbersDeleteCall) Context(ctx context.Context) *ProjectsPhoneNumbersDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsPhoneNumbersDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsPhoneNumbersDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.nameid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.phoneNumbers.delete" call.
// Exactly one of *GoogleCloudDialogflowV2beta1PhoneNumber or error will
// be non-nil. Any non-2xx status code is an error. Response headers are
// in either
// *GoogleCloudDialogflowV2beta1PhoneNumber.ServerResponse.Header or (if
// a response was returned at all) in error.(*googleapi.Error).Header.
// Use googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsPhoneNumbersDeleteCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1PhoneNumber, error) {
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
	ret := &GoogleCloudDialogflowV2beta1PhoneNumber{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Requests deletion of a `PhoneNumber`. The `PhoneNumber` is moved into the\nDELETE_REQUESTED state\nimmediately, and is deleted approximately 30 days later. This method may\nonly be called on a `PhoneNumber` in the\nACTIVE state.",
	//   "flatPath": "v2beta1/projects/{projectsId}/phoneNumbers/{phoneNumbersId}",
	//   "httpMethod": "DELETE",
	//   "id": "dialogflow.projects.phoneNumbers.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The unique identifier of the `PhoneNumber` to delete.\nFormat: `projects/\u003cProject ID\u003e/phoneNumbers/\u003cPhoneNumber ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/phoneNumbers/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1PhoneNumber"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.phoneNumbers.list":

type ProjectsPhoneNumbersListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Returns the list of all phone numbers in the specified project.
func (r *ProjectsPhoneNumbersService) List(parent string) *ProjectsPhoneNumbersListCall {
	c := &ProjectsPhoneNumbersListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of items to return in a single page.
// The default value is 100. The maximum value is 1000.
func (c *ProjectsPhoneNumbersListCall) PageSize(pageSize int64) *ProjectsPhoneNumbersListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous list request.
func (c *ProjectsPhoneNumbersListCall) PageToken(pageToken string) *ProjectsPhoneNumbersListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// ShowDeleted sets the optional parameter "showDeleted": Controls
// whether `PhoneNumber` resources in the
// DELETE_REQUESTED
// state should be returned. Defaults to false.
func (c *ProjectsPhoneNumbersListCall) ShowDeleted(showDeleted bool) *ProjectsPhoneNumbersListCall {
	c.urlParams_.Set("showDeleted", fmt.Sprint(showDeleted))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsPhoneNumbersListCall) Fields(s ...googleapi.Field) *ProjectsPhoneNumbersListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsPhoneNumbersListCall) IfNoneMatch(entityTag string) *ProjectsPhoneNumbersListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsPhoneNumbersListCall) Context(ctx context.Context) *ProjectsPhoneNumbersListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsPhoneNumbersListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsPhoneNumbersListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/phoneNumbers")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.phoneNumbers.list" call.
// Exactly one of *GoogleCloudDialogflowV2beta1ListPhoneNumbersResponse
// or error will be non-nil. Any non-2xx status code is an error.
// Response headers are in either
// *GoogleCloudDialogflowV2beta1ListPhoneNumbersResponse.ServerResponse.H
// eader or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsPhoneNumbersListCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1ListPhoneNumbersResponse, error) {
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
	ret := &GoogleCloudDialogflowV2beta1ListPhoneNumbersResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the list of all phone numbers in the specified project.",
	//   "flatPath": "v2beta1/projects/{projectsId}/phoneNumbers",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.phoneNumbers.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Optional. The maximum number of items to return in a single page.\nThe default value is 100. The maximum value is 1000.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. The next_page_token value returned from a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The project to list all `PhoneNumber` resources from.\nFormat: `projects/\u003cProject ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "showDeleted": {
	//       "description": "Optional. Controls whether `PhoneNumber` resources in the\nDELETE_REQUESTED\nstate should be returned. Defaults to false.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/phoneNumbers",
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1ListPhoneNumbersResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsPhoneNumbersListCall) Pages(ctx context.Context, f func(*GoogleCloudDialogflowV2beta1ListPhoneNumbersResponse) error) error {
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

// method id "dialogflow.projects.phoneNumbers.patch":

type ProjectsPhoneNumbersPatchCall struct {
	s                                       *Service
	nameid                                  string
	googleclouddialogflowv2beta1phonenumber *GoogleCloudDialogflowV2beta1PhoneNumber
	urlParams_                              gensupport.URLParams
	ctx_                                    context.Context
	header_                                 http.Header
}

// Patch: Updates the specified `PhoneNumber`.
func (r *ProjectsPhoneNumbersService) Patch(nameid string, googleclouddialogflowv2beta1phonenumber *GoogleCloudDialogflowV2beta1PhoneNumber) *ProjectsPhoneNumbersPatchCall {
	c := &ProjectsPhoneNumbersPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.nameid = nameid
	c.googleclouddialogflowv2beta1phonenumber = googleclouddialogflowv2beta1phonenumber
	return c
}

// UpdateMask sets the optional parameter "updateMask": The mask to
// control which fields get updated.
func (c *ProjectsPhoneNumbersPatchCall) UpdateMask(updateMask string) *ProjectsPhoneNumbersPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsPhoneNumbersPatchCall) Fields(s ...googleapi.Field) *ProjectsPhoneNumbersPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsPhoneNumbersPatchCall) Context(ctx context.Context) *ProjectsPhoneNumbersPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsPhoneNumbersPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsPhoneNumbersPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1phonenumber)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.nameid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.phoneNumbers.patch" call.
// Exactly one of *GoogleCloudDialogflowV2beta1PhoneNumber or error will
// be non-nil. Any non-2xx status code is an error. Response headers are
// in either
// *GoogleCloudDialogflowV2beta1PhoneNumber.ServerResponse.Header or (if
// a response was returned at all) in error.(*googleapi.Error).Header.
// Use googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsPhoneNumbersPatchCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1PhoneNumber, error) {
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
	ret := &GoogleCloudDialogflowV2beta1PhoneNumber{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the specified `PhoneNumber`.",
	//   "flatPath": "v2beta1/projects/{projectsId}/phoneNumbers/{phoneNumbersId}",
	//   "httpMethod": "PATCH",
	//   "id": "dialogflow.projects.phoneNumbers.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The unique identifier of this phone number.\nFormat: `projects/\u003cProject ID\u003e/phoneNumbers/\u003cPhoneNumber ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/phoneNumbers/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Optional. The mask to control which fields get updated.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1PhoneNumber"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1PhoneNumber"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.phoneNumbers.undelete":

type ProjectsPhoneNumbersUndeleteCall struct {
	s                                                      *Service
	nameid                                                 string
	googleclouddialogflowv2beta1undeletephonenumberrequest *GoogleCloudDialogflowV2beta1UndeletePhoneNumberRequest
	urlParams_                                             gensupport.URLParams
	ctx_                                                   context.Context
	header_                                                http.Header
}

// Undelete: Cancels the deletion request for a `PhoneNumber`. This
// method may only be
// called on a `PhoneNumber` in the
// DELETE_REQUESTED state.
func (r *ProjectsPhoneNumbersService) Undelete(nameid string, googleclouddialogflowv2beta1undeletephonenumberrequest *GoogleCloudDialogflowV2beta1UndeletePhoneNumberRequest) *ProjectsPhoneNumbersUndeleteCall {
	c := &ProjectsPhoneNumbersUndeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.nameid = nameid
	c.googleclouddialogflowv2beta1undeletephonenumberrequest = googleclouddialogflowv2beta1undeletephonenumberrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsPhoneNumbersUndeleteCall) Fields(s ...googleapi.Field) *ProjectsPhoneNumbersUndeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsPhoneNumbersUndeleteCall) Context(ctx context.Context) *ProjectsPhoneNumbersUndeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsPhoneNumbersUndeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsPhoneNumbersUndeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleclouddialogflowv2beta1undeletephonenumberrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}:undelete")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.nameid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.phoneNumbers.undelete" call.
// Exactly one of *GoogleCloudDialogflowV2beta1PhoneNumber or error will
// be non-nil. Any non-2xx status code is an error. Response headers are
// in either
// *GoogleCloudDialogflowV2beta1PhoneNumber.ServerResponse.Header or (if
// a response was returned at all) in error.(*googleapi.Error).Header.
// Use googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsPhoneNumbersUndeleteCall) Do(opts ...googleapi.CallOption) (*GoogleCloudDialogflowV2beta1PhoneNumber, error) {
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
	ret := &GoogleCloudDialogflowV2beta1PhoneNumber{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Cancels the deletion request for a `PhoneNumber`. This method may only be\ncalled on a `PhoneNumber` in the\nDELETE_REQUESTED state.",
	//   "flatPath": "v2beta1/projects/{projectsId}/phoneNumbers/{phoneNumbersId}:undelete",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.phoneNumbers.undelete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The unique identifier of the `PhoneNumber` to delete.\nFormat: `projects/\u003cProject ID\u003e/phoneNumbers/\u003cPhoneNumber ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/phoneNumbers/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}:undelete",
	//   "request": {
	//     "$ref": "GoogleCloudDialogflowV2beta1UndeletePhoneNumberRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudDialogflowV2beta1PhoneNumber"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
