// Package language provides access to the Google Cloud Natural Language API.
//
// See https://cloud.google.com/natural-language/
//
// Usage example:
//
//   import "google.golang.org/api/language/v1beta1"
//   ...
//   languageService, err := language.New(oauthHttpClient)
package language // import "google.golang.org/api/language/v1beta1"

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

const apiId = "language:v1beta1"
const apiName = "language"
const apiVersion = "v1beta1"
const basePath = "https://language.googleapis.com/"

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
	s.Documents = NewDocumentsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Documents *DocumentsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewDocumentsService(s *Service) *DocumentsService {
	rs := &DocumentsService{s: s}
	return rs
}

type DocumentsService struct {
	s *Service
}

// AnalyzeEntitiesRequest: The entity analysis request message.
type AnalyzeEntitiesRequest struct {
	// Document: Input document.
	Document *Document `json:"document,omitempty"`

	// EncodingType: The encoding type used by the API to calculate offsets.
	//
	// Possible values:
	//   "NONE" - If `EncodingType` is not specified, encoding-dependent
	// information (such as
	// `begin_offset`) will be set at `-1`.
	//   "UTF8" - Encoding-dependent information (such as `begin_offset`) is
	// calculated based
	// on the UTF-8 encoding of the input. C++ and Go are examples of
	// languages
	// that use this encoding natively.
	//   "UTF16" - Encoding-dependent information (such as `begin_offset`)
	// is calculated based
	// on the UTF-16 encoding of the input. Java and Javascript are examples
	// of
	// languages that use this encoding natively.
	//   "UTF32" - Encoding-dependent information (such as `begin_offset`)
	// is calculated based
	// on the UTF-32 encoding of the input. Python is an example of a
	// language
	// that uses this encoding natively.
	EncodingType string `json:"encodingType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Document") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AnalyzeEntitiesRequest) MarshalJSON() ([]byte, error) {
	type noMethod AnalyzeEntitiesRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// AnalyzeEntitiesResponse: The entity analysis response message.
type AnalyzeEntitiesResponse struct {
	// Entities: The recognized entities in the input document.
	Entities []*Entity `json:"entities,omitempty"`

	// Language: The language of the text, which will be the same as the
	// language specified
	// in the request or, if not specified, the automatically-detected
	// language.
	Language string `json:"language,omitempty"`

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
}

func (s *AnalyzeEntitiesResponse) MarshalJSON() ([]byte, error) {
	type noMethod AnalyzeEntitiesResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// AnalyzeSentimentRequest: The sentiment analysis request message.
type AnalyzeSentimentRequest struct {
	// Document: Input document. Currently, `analyzeSentiment` only supports
	// English text
	// (Document.language="EN").
	Document *Document `json:"document,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Document") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AnalyzeSentimentRequest) MarshalJSON() ([]byte, error) {
	type noMethod AnalyzeSentimentRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// AnalyzeSentimentResponse: The sentiment analysis response message.
type AnalyzeSentimentResponse struct {
	// DocumentSentiment: The overall sentiment of the input document.
	DocumentSentiment *Sentiment `json:"documentSentiment,omitempty"`

	// Language: The language of the text, which will be the same as the
	// language specified
	// in the request or, if not specified, the automatically-detected
	// language.
	Language string `json:"language,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "DocumentSentiment")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AnalyzeSentimentResponse) MarshalJSON() ([]byte, error) {
	type noMethod AnalyzeSentimentResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// AnnotateTextRequest: The request message for the advanced text
// annotation API, which performs all
// the above plus syntactic analysis.
type AnnotateTextRequest struct {
	// Document: Input document.
	Document *Document `json:"document,omitempty"`

	// EncodingType: The encoding type used by the API to calculate offsets.
	//
	// Possible values:
	//   "NONE" - If `EncodingType` is not specified, encoding-dependent
	// information (such as
	// `begin_offset`) will be set at `-1`.
	//   "UTF8" - Encoding-dependent information (such as `begin_offset`) is
	// calculated based
	// on the UTF-8 encoding of the input. C++ and Go are examples of
	// languages
	// that use this encoding natively.
	//   "UTF16" - Encoding-dependent information (such as `begin_offset`)
	// is calculated based
	// on the UTF-16 encoding of the input. Java and Javascript are examples
	// of
	// languages that use this encoding natively.
	//   "UTF32" - Encoding-dependent information (such as `begin_offset`)
	// is calculated based
	// on the UTF-32 encoding of the input. Python is an example of a
	// language
	// that uses this encoding natively.
	EncodingType string `json:"encodingType,omitempty"`

	// Features: The enabled features.
	Features *Features `json:"features,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Document") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AnnotateTextRequest) MarshalJSON() ([]byte, error) {
	type noMethod AnnotateTextRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// AnnotateTextResponse: The text annotations response message.
type AnnotateTextResponse struct {
	// DocumentSentiment: The overall sentiment for the document. Populated
	// if the user
	// enables
	// AnnotateTextRequest.Features.extract_document_sentiment.
	DocumentSentiment *Sentiment `json:"documentSentiment,omitempty"`

	// Entities: Entities, along with their semantic information, in the
	// input document.
	// Populated if the user
	// enables
	// AnnotateTextRequest.Features.extract_entities.
	Entities []*Entity `json:"entities,omitempty"`

	// Language: The language of the text, which will be the same as the
	// language specified
	// in the request or, if not specified, the automatically-detected
	// language.
	Language string `json:"language,omitempty"`

	// Sentences: Sentences in the input document. Populated if the user
	// enables
	// AnnotateTextRequest.Features.extract_syntax.
	Sentences []*Sentence `json:"sentences,omitempty"`

	// Tokens: Tokens, along with their syntactic information, in the input
	// document.
	// Populated if the user
	// enables
	// AnnotateTextRequest.Features.extract_syntax.
	Tokens []*Token `json:"tokens,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "DocumentSentiment")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AnnotateTextResponse) MarshalJSON() ([]byte, error) {
	type noMethod AnnotateTextResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// DependencyEdge: Represents dependency parse tree information for a
// token.
type DependencyEdge struct {
	// HeadTokenIndex: Represents the head of this token in the dependency
	// tree.
	// This is the index of the token which has an arc going to this
	// token.
	// The index is the position of the token in the array of tokens
	// returned
	// by the API method. If this token is a root token, then
	// the
	// `head_token_index` is its own index.
	HeadTokenIndex int64 `json:"headTokenIndex,omitempty"`

	// Label: The parse label for the token.
	//
	// Possible values:
	//   "UNKNOWN" - Unknown
	//   "ABBREV" - Abbreviation modifier
	//   "ACOMP" - Adjectival complement
	//   "ADVCL" - Adverbial clause modifier
	//   "ADVMOD" - Adverbial modifier
	//   "AMOD" - Adjectival modifier of an NP
	//   "APPOS" - Appositional modifier of an NP
	//   "ATTR" - Attribute dependent of a copular verb
	//   "AUX" - Auxiliary (non-main) verb
	//   "AUXPASS" - Passive auxiliary
	//   "CC" - Coordinating conjunction
	//   "CCOMP" - Clausal complement of a verb or adjective
	//   "CONJ" - Conjunct
	//   "CSUBJ" - Clausal subject
	//   "CSUBJPASS" - Clausal passive subject
	//   "DEP" - Dependency (unable to determine)
	//   "DET" - Determiner
	//   "DISCOURSE" - Discourse
	//   "DOBJ" - Direct object
	//   "EXPL" - Expletive
	//   "GOESWITH" - Goes with (part of a word in a text not well edited)
	//   "IOBJ" - Indirect object
	//   "MARK" - Marker (word introducing a subordinate clause)
	//   "MWE" - Multi-word expression
	//   "MWV" - Multi-word verbal expression
	//   "NEG" - Negation modifier
	//   "NN" - Noun compound modifier
	//   "NPADVMOD" - Noun phrase used as an adverbial modifier
	//   "NSUBJ" - Nominal subject
	//   "NSUBJPASS" - Passive nominal subject
	//   "NUM" - Numeric modifier of a noun
	//   "NUMBER" - Element of compound number
	//   "P" - Punctuation mark
	//   "PARATAXIS" - Parataxis relation
	//   "PARTMOD" - Participial modifier
	//   "PCOMP" - The complement of a preposition is a clause
	//   "POBJ" - Object of a preposition
	//   "POSS" - Possession modifier
	//   "POSTNEG" - Postverbal negative particle
	//   "PRECOMP" - Predicate complement
	//   "PRECONJ" - Preconjunt
	//   "PREDET" - Predeterminer
	//   "PREF" - Prefix
	//   "PREP" - Prepositional modifier
	//   "PRONL" - The relationship between a verb and verbal morpheme
	//   "PRT" - Particle
	//   "PS" - Associative or possessive marker
	//   "QUANTMOD" - Quantifier phrase modifier
	//   "RCMOD" - Relative clause modifier
	//   "RCMODREL" - Complementizer in relative clause
	//   "RDROP" - Ellipsis without a preceding predicate
	//   "REF" - Referent
	//   "REMNANT" - Remnant
	//   "REPARANDUM" - Reparandum
	//   "ROOT" - Root
	//   "SNUM" - Suffix specifying a unit of number
	//   "SUFF" - Suffix
	//   "TMOD" - Temporal modifier
	//   "TOPIC" - Topic marker
	//   "VMOD" - Clause headed by an infinite form of the verb that
	// modifies a noun
	//   "VOCATIVE" - Vocative
	//   "XCOMP" - Open clausal complement
	//   "SUFFIX" - Name suffix
	//   "TITLE" - Name title
	//   "ADVPHMOD" - Adverbial phrase modifier
	//   "AUXCAUS" - Causative auxiliary
	//   "AUXVV" - Helper auxiliary
	//   "DTMOD" - Rentaishi (Prenominal modifier)
	//   "FOREIGN" - Foreign words
	//   "KW" - Keyword
	//   "LIST" - List for chains of comparable items
	//   "NOMC" - Nominalized clause
	//   "NOMCSUBJ" - Nominalized clausal subject
	//   "NOMCSUBJPASS" - Nominalized clausal passive
	//   "NUMC" - Compound of numeric modifier
	//   "COP" - Copula
	//   "DISLOCATED" - Dislocated relation (for fronted/topicalized
	// elements)
	Label string `json:"label,omitempty"`

	// ForceSendFields is a list of field names (e.g. "HeadTokenIndex") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *DependencyEdge) MarshalJSON() ([]byte, error) {
	type noMethod DependencyEdge
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Document:
// ################################################################
// #
//
// Represents the input to API methods.
type Document struct {
	// Content: The content of the input in string format.
	Content string `json:"content,omitempty"`

	// GcsContentUri: The Google Cloud Storage URI where the file content is
	// located.
	GcsContentUri string `json:"gcsContentUri,omitempty"`

	// Language: The language of the document (if not specified, the
	// language is
	// automatically detected). Both ISO and BCP-47 language codes
	// are
	// accepted.<br>
	// **Current Language Restrictions:**
	//
	//  * Only English, Spanish, and Japanese textual content
	//    are supported, with the following additional restriction:
	//    * `analyzeSentiment` only supports English text.
	// If the language (either specified by the caller or automatically
	// detected)
	// is not supported by the called API method, an `INVALID_ARGUMENT`
	// error
	// is returned.
	Language string `json:"language,omitempty"`

	// Type: Required. If the type is not set or is
	// `TYPE_UNSPECIFIED`,
	// returns an `INVALID_ARGUMENT` error.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED" - The content type is not specified.
	//   "PLAIN_TEXT" - Plain text
	//   "HTML" - HTML
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Content") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Document) MarshalJSON() ([]byte, error) {
	type noMethod Document
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Entity: Represents a phrase in the text that is a known entity, such
// as
// a person, an organization, or location. The API associates
// information, such
// as salience and mentions, with entities.
type Entity struct {
	// Mentions: The mentions of this entity in the input document. The API
	// currently
	// supports proper noun mentions.
	Mentions []*EntityMention `json:"mentions,omitempty"`

	// Metadata: Metadata associated with the entity.
	//
	// Currently, only Wikipedia URLs are provided, if available.
	// The associated key is "wikipedia_url".
	Metadata map[string]string `json:"metadata,omitempty"`

	// Name: The representative name for the entity.
	Name string `json:"name,omitempty"`

	// Salience: The salience score associated with the entity in the [0,
	// 1.0] range.
	//
	// The salience score for an entity provides information about
	// the
	// importance or centrality of that entity to the entire document
	// text.
	// Scores closer to 0 are less salient, while scores closer to 1.0 are
	// highly
	// salient.
	Salience float64 `json:"salience,omitempty"`

	// Type: The entity type.
	//
	// Possible values:
	//   "UNKNOWN" - Unknown
	//   "PERSON" - Person
	//   "LOCATION" - Location
	//   "ORGANIZATION" - Organization
	//   "EVENT" - Event
	//   "WORK_OF_ART" - Work of art
	//   "CONSUMER_GOOD" - Consumer goods
	//   "OTHER" - Other types
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Mentions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Entity) MarshalJSON() ([]byte, error) {
	type noMethod Entity
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// EntityMention: Represents a mention for an entity in the text.
// Currently, proper noun
// mentions are supported.
type EntityMention struct {
	// Text: The mention text.
	Text *TextSpan `json:"text,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Text") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *EntityMention) MarshalJSON() ([]byte, error) {
	type noMethod EntityMention
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Features: All available features for sentiment, syntax, and semantic
// analysis.
// Setting each one to true will enable that specific analysis for the
// input.
type Features struct {
	// ExtractDocumentSentiment: Extract document-level sentiment.
	ExtractDocumentSentiment bool `json:"extractDocumentSentiment,omitempty"`

	// ExtractEntities: Extract entities.
	ExtractEntities bool `json:"extractEntities,omitempty"`

	// ExtractSyntax: Extract syntax information.
	ExtractSyntax bool `json:"extractSyntax,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "ExtractDocumentSentiment") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Features) MarshalJSON() ([]byte, error) {
	type noMethod Features
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// PartOfSpeech: Represents part of speech information for a token.
type PartOfSpeech struct {
	// Tag: The part of speech tag.
	//
	// Possible values:
	//   "UNKNOWN" - Unknown
	//   "ADJ" - Adjective
	//   "ADP" - Adposition (preposition and postposition)
	//   "ADV" - Adverb
	//   "CONJ" - Conjunction
	//   "DET" - Determiner
	//   "NOUN" - Noun (common and proper)
	//   "NUM" - Cardinal number
	//   "PRON" - Pronoun
	//   "PRT" - Particle or other function word
	//   "PUNCT" - Punctuation
	//   "VERB" - Verb (all tenses and modes)
	//   "X" - Other: foreign words, typos, abbreviations
	//   "AFFIX" - Affix
	Tag string `json:"tag,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Tag") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PartOfSpeech) MarshalJSON() ([]byte, error) {
	type noMethod PartOfSpeech
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Sentence: Represents a sentence in the input document.
type Sentence struct {
	// Text: The sentence text.
	Text *TextSpan `json:"text,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Text") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Sentence) MarshalJSON() ([]byte, error) {
	type noMethod Sentence
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Sentiment: Represents the feeling associated with the entire text or
// entities in
// the text.
type Sentiment struct {
	// Magnitude: A non-negative number in the [0, +inf) range, which
	// represents
	// the absolute magnitude of sentiment regardless of polarity (positive
	// or
	// negative).
	Magnitude float64 `json:"magnitude,omitempty"`

	// Polarity: Polarity of the sentiment in the [-1.0, 1.0] range. Larger
	// numbers
	// represent more positive sentiments.
	Polarity float64 `json:"polarity,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Magnitude") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Sentiment) MarshalJSON() ([]byte, error) {
	type noMethod Sentiment
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
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
	Details []StatusDetails `json:"details,omitempty"`

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
}

func (s *Status) MarshalJSON() ([]byte, error) {
	type noMethod Status
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type StatusDetails interface{}

// TextSpan: Represents an output piece of text.
type TextSpan struct {
	// BeginOffset: The API calculates the beginning offset of the content
	// in the original
	// document according to the EncodingType specified in the API request.
	BeginOffset int64 `json:"beginOffset,omitempty"`

	// Content: The content of the output text.
	Content string `json:"content,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BeginOffset") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *TextSpan) MarshalJSON() ([]byte, error) {
	type noMethod TextSpan
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Token: Represents the smallest syntactic building block of the text.
type Token struct {
	// DependencyEdge: Dependency tree parse for this token.
	DependencyEdge *DependencyEdge `json:"dependencyEdge,omitempty"`

	// Lemma: [Lemma](https://en.wikipedia.org/wiki/Lemma_(morphology))
	// of the token.
	Lemma string `json:"lemma,omitempty"`

	// PartOfSpeech: Parts of speech tag for this token.
	PartOfSpeech *PartOfSpeech `json:"partOfSpeech,omitempty"`

	// Text: The token text.
	Text *TextSpan `json:"text,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DependencyEdge") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Token) MarshalJSON() ([]byte, error) {
	type noMethod Token
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// method id "language.documents.analyzeEntities":

type DocumentsAnalyzeEntitiesCall struct {
	s                      *Service
	analyzeentitiesrequest *AnalyzeEntitiesRequest
	urlParams_             gensupport.URLParams
	ctx_                   context.Context
}

// AnalyzeEntities: Finds named entities (currently finds proper names)
// in the text,
// entity types, salience, mentions for each entity, and other
// properties.
func (r *DocumentsService) AnalyzeEntities(analyzeentitiesrequest *AnalyzeEntitiesRequest) *DocumentsAnalyzeEntitiesCall {
	c := &DocumentsAnalyzeEntitiesCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.analyzeentitiesrequest = analyzeentitiesrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DocumentsAnalyzeEntitiesCall) Fields(s ...googleapi.Field) *DocumentsAnalyzeEntitiesCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *DocumentsAnalyzeEntitiesCall) Context(ctx context.Context) *DocumentsAnalyzeEntitiesCall {
	c.ctx_ = ctx
	return c
}

func (c *DocumentsAnalyzeEntitiesCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.analyzeentitiesrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/documents:analyzeEntities")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.SetOpaque(req.URL)
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "language.documents.analyzeEntities" call.
// Exactly one of *AnalyzeEntitiesResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *AnalyzeEntitiesResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *DocumentsAnalyzeEntitiesCall) Do(opts ...googleapi.CallOption) (*AnalyzeEntitiesResponse, error) {
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
	ret := &AnalyzeEntitiesResponse{
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
	//   "description": "Finds named entities (currently finds proper names) in the text,\nentity types, salience, mentions for each entity, and other properties.",
	//   "flatPath": "v1beta1/documents:analyzeEntities",
	//   "httpMethod": "POST",
	//   "id": "language.documents.analyzeEntities",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1beta1/documents:analyzeEntities",
	//   "request": {
	//     "$ref": "AnalyzeEntitiesRequest"
	//   },
	//   "response": {
	//     "$ref": "AnalyzeEntitiesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "language.documents.analyzeSentiment":

type DocumentsAnalyzeSentimentCall struct {
	s                       *Service
	analyzesentimentrequest *AnalyzeSentimentRequest
	urlParams_              gensupport.URLParams
	ctx_                    context.Context
}

// AnalyzeSentiment: Analyzes the sentiment of the provided text.
func (r *DocumentsService) AnalyzeSentiment(analyzesentimentrequest *AnalyzeSentimentRequest) *DocumentsAnalyzeSentimentCall {
	c := &DocumentsAnalyzeSentimentCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.analyzesentimentrequest = analyzesentimentrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DocumentsAnalyzeSentimentCall) Fields(s ...googleapi.Field) *DocumentsAnalyzeSentimentCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *DocumentsAnalyzeSentimentCall) Context(ctx context.Context) *DocumentsAnalyzeSentimentCall {
	c.ctx_ = ctx
	return c
}

func (c *DocumentsAnalyzeSentimentCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.analyzesentimentrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/documents:analyzeSentiment")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.SetOpaque(req.URL)
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "language.documents.analyzeSentiment" call.
// Exactly one of *AnalyzeSentimentResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *AnalyzeSentimentResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *DocumentsAnalyzeSentimentCall) Do(opts ...googleapi.CallOption) (*AnalyzeSentimentResponse, error) {
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
	ret := &AnalyzeSentimentResponse{
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
	//   "description": "Analyzes the sentiment of the provided text.",
	//   "flatPath": "v1beta1/documents:analyzeSentiment",
	//   "httpMethod": "POST",
	//   "id": "language.documents.analyzeSentiment",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1beta1/documents:analyzeSentiment",
	//   "request": {
	//     "$ref": "AnalyzeSentimentRequest"
	//   },
	//   "response": {
	//     "$ref": "AnalyzeSentimentResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "language.documents.annotateText":

type DocumentsAnnotateTextCall struct {
	s                   *Service
	annotatetextrequest *AnnotateTextRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
}

// AnnotateText: Advanced API that analyzes the document and provides a
// full set of text
// annotations, including semantic, syntactic, and sentiment
// information. This
// API is intended for users who are familiar with machine learning and
// need
// in-depth text features to build upon.
func (r *DocumentsService) AnnotateText(annotatetextrequest *AnnotateTextRequest) *DocumentsAnnotateTextCall {
	c := &DocumentsAnnotateTextCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.annotatetextrequest = annotatetextrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DocumentsAnnotateTextCall) Fields(s ...googleapi.Field) *DocumentsAnnotateTextCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *DocumentsAnnotateTextCall) Context(ctx context.Context) *DocumentsAnnotateTextCall {
	c.ctx_ = ctx
	return c
}

func (c *DocumentsAnnotateTextCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.annotatetextrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/documents:annotateText")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.SetOpaque(req.URL)
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "language.documents.annotateText" call.
// Exactly one of *AnnotateTextResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *AnnotateTextResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *DocumentsAnnotateTextCall) Do(opts ...googleapi.CallOption) (*AnnotateTextResponse, error) {
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
	ret := &AnnotateTextResponse{
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
	//   "description": "Advanced API that analyzes the document and provides a full set of text\nannotations, including semantic, syntactic, and sentiment information. This\nAPI is intended for users who are familiar with machine learning and need\nin-depth text features to build upon.",
	//   "flatPath": "v1beta1/documents:annotateText",
	//   "httpMethod": "POST",
	//   "id": "language.documents.annotateText",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1beta1/documents:annotateText",
	//   "request": {
	//     "$ref": "AnnotateTextRequest"
	//   },
	//   "response": {
	//     "$ref": "AnnotateTextResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
