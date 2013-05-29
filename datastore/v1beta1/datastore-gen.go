// Package datastore provides access to the Google Cloud Datastore API.
//
// See https://developers.google.com/datastore/
//
// Usage example:
//
//   import "code.google.com/p/google-api-go-client/datastore/v1beta1"
//   ...
//   datastoreService, err := datastore.New(oauthHttpClient)
package datastore

import (
	"bytes"
	"code.google.com/p/google-api-go-client/googleapi"
	"encoding/json"
	"errors"
	"fmt"
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

const apiId = "datastore:v1beta1"
const apiName = "datastore"
const apiVersion = "v1beta1"
const basePath = "https://www.googleapis.com/datastore/v1beta1/datasets/"

// OAuth2 scopes used by this API.
const (
	// View your email address
	UserinfoEmailScope = "https://www.googleapis.com/auth/userinfo.email"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Datasets = NewDatasetsService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Datasets *DatasetsService
}

func NewDatasetsService(s *Service) *DatasetsService {
	rs := &DatasetsService{s: s}
	return rs
}

type DatasetsService struct {
	s *Service
}

type AllocateIdsRequest struct {
	// Keys: A list of keys with incomplete key paths to allocate ids for.
	Keys []*Key `json:"keys,omitempty"`
}

type AllocateIdsResponse struct {
	// Keys: The keys specified in the request (in the same order), each
	// with its key path completed with a newly allocated id.
	Keys []*Key `json:"keys,omitempty"`

	// Kind: The kind, fixed to "datastore#allocateIdsResponse".
	Kind string `json:"kind,omitempty"`
}

type BeginTransactionRequest struct {
	// IsolationLevel: The transaction isolation level. One of "snapshot" or
	// "serializable" (optional, defaults to "snapshot").
	IsolationLevel string `json:"isolationLevel,omitempty"`
}

type BeginTransactionResponse struct {
	// Kind: The kind, fixed to "datastore#beginTransactionResponse".
	Kind string `json:"kind,omitempty"`

	// Transaction: The transaction identifier (always present).
	Transaction string `json:"transaction,omitempty"`
}

type BlindWriteRequest struct {
	// Mutation: The mutation to commit.
	Mutation *Mutation `json:"mutation,omitempty"`
}

type BlindWriteResponse struct {
	// Kind: The kind, fixed to "datastore#blindWriteResponse".
	Kind string `json:"kind,omitempty"`

	// MutationResult: The result of committing the mutation.
	MutationResult *MutationResult `json:"mutationResult,omitempty"`
}

type CommitRequest struct {
	// Mutation: The mutation to perform as part of this transaction
	// (optional).
	Mutation *Mutation `json:"mutation,omitempty"`

	// Transaction: The transaction identifier.
	Transaction string `json:"transaction,omitempty"`
}

type CommitResponse struct {
	// Kind: The kind, fixed to "datastore#commitResponse".
	Kind string `json:"kind,omitempty"`

	// MutationResult: The result of committing the mutation (if any).
	MutationResult *MutationResult `json:"mutationResult,omitempty"`
}

type CompositeFilter struct {
	// Filters: The list of filters to combine. Must contain at least one
	// filter.
	Filters []*Filter `json:"filters,omitempty"`

	// Operator: The operator for combining multiple filters. Only "and" is
	// currently supported.
	Operator string `json:"operator,omitempty"`
}

type Entity struct {
	// Key: The entity's key.
	//
	// An entity must have a key, unless otherwise
	// documented (for example an entity in Value.entity_value may have no
	// key). An entity's kind is its key's path's last element's kind, or
	// null if it has no key.
	Key *Key `json:"key,omitempty"`

	// Properties: The entity's properties. Each property's name must be
	// unique for its entity.
	Properties *EntityProperties `json:"properties,omitempty"`
}

type EntityProperties struct {
}

type EntityResult struct {
	// Entity: The resulting entity.
	Entity *Entity `json:"entity,omitempty"`
}

type Filter struct {
	// CompositeFilter: A composite filter.
	CompositeFilter *CompositeFilter `json:"compositeFilter,omitempty"`

	// PropertyFilter: A filter on a property.
	PropertyFilter *PropertyFilter `json:"propertyFilter,omitempty"`
}

type Key struct {
	// PartitionId: The ID of the partition containing the entity.
	PartitionId *PartitionId `json:"partitionId,omitempty"`

	// Path: The path of the entity.
	//
	// Each path element identifies an
	// ancestor entity. The last path element identifies the entity, the
	// next-to-last path element identifies the entity's parent, and so
	// forth. The last path element and ONLY the last path element may be
	// incomplete. A path is complete if ALL of its path elements are
	// complete. Any other path is incomplete. A path must be complete
	// unless otherwise documented. A path must not be empty.
	Path []*KeyPathElement `json:"path,omitempty"`
}

type KeyPathElement struct {
	// Id: The id of the entity. Must be > 0.
	Id int64 `json:"id,omitempty,string"`

	// Kind: The kind of the entity. Kinds matching regex "__.*__" are
	// reserved/read_only. Must not be "".
	Kind string `json:"kind,omitempty"`

	// Name: The name of the entity. Must not be "", nor match regex
	// "__.*__".
	Name string `json:"name,omitempty"`
}

type KindExpression struct {
	// Name: The name of the kind.
	Name string `json:"name,omitempty"`
}

type LookupRequest struct {
	// Keys: Keys of entities to get from the datastore.
	Keys []*Key `json:"keys,omitempty"`

	// ReadOptions: Options for this get request (optional).
	ReadOptions *ReadOptions `json:"readOptions,omitempty"`
}

type LookupResponse struct {
	// Deferred: A list of keys that were not looked up due to resource
	// constraints.
	Deferred []*Key `json:"deferred,omitempty"`

	// Found: Entities found.
	Found []*EntityResult `json:"found,omitempty"`

	// Kind: The kind, fixed to "datastore#lookupResponse".
	Kind string `json:"kind,omitempty"`

	// Missing: Entities not found, with only the key populated.
	Missing []*EntityResult `json:"missing,omitempty"`
}

type Mutation struct {
	// Delete: Keys of entities to delete.
	Delete []*Key `json:"delete,omitempty"`

	// Force: Ignore a user specified read-only period (optional).
	Force bool `json:"force,omitempty"`

	// Insert: Entities to insert. Inserted entities MUST have a complete
	// key path.
	Insert []*Entity `json:"insert,omitempty"`

	// InsertAutoId: Insert insertAutoId entities with a newly allocated id.
	// Each entity's key must have an incomplete key path.
	InsertAutoId []*Entity `json:"insertAutoId,omitempty"`

	// Update: Entities to update. Updated entities MUST have a complete key
	// path.
	Update []*Entity `json:"update,omitempty"`

	// Upsert: Entities to upsert. Upserted entities MUST have a complete
	// key path.
	Upsert []*Entity `json:"upsert,omitempty"`
}

type MutationResult struct {
	// IndexUpdates: Number of index entries changed.
	IndexUpdates int64 `json:"indexUpdates,omitempty"`

	// InsertAutoIdKeys: Keys for insert_auto_id entities. One per entity
	// from the request, in the same order.
	InsertAutoIdKeys []*Key `json:"insertAutoIdKeys,omitempty"`
}

type PartitionId struct {
	// DatasetId: The dataset id, usually the 'app' id of the owning app.
	DatasetId string `json:"datasetId,omitempty"`

	// Namespace: The namespace.
	Namespace string `json:"namespace,omitempty"`
}

type Property struct {
	// Multi: If this property contains a list of values. Input values may
	// explicitly set multi to false, but otherwise false is always
	// represented by omitting multi.
	Multi bool `json:"multi,omitempty"`

	// Values: The value(s) of the property. When multi is false there must
	// be exactly one value. When multi is true there must be one or more
	// values.
	Values []*Value `json:"values,omitempty"`
}

type PropertyExpression struct {
	// AggregationFunction: The aggregation function to apply to the
	// property (optional). Can only be used when grouping by at least one
	// property. Must then be set on all properties in the projection that
	// are not being grouped by. Aggregation functions: "first" selects the
	// first result as determined by the query's order.
	AggregationFunction string `json:"aggregationFunction,omitempty"`

	// Property: The property to project.
	Property *PropertyReference `json:"property,omitempty"`
}

type PropertyFilter struct {
	// Operator: The operator to filter by. One of "lessThan",
	// "lessThanOrEqual", "greaterThan", "greaterThanOrEqual", "equal",
	// "hasAncestor".
	Operator string `json:"operator,omitempty"`

	// Property: The property to filter by.
	Property *PropertyReference `json:"property,omitempty"`

	// Value: The value to compare the property to.
	Value *Value `json:"value,omitempty"`
}

type PropertyOrder struct {
	// Direction: The direction to order by. One of "ascending" or
	// "descending" (optional, defaults to "ascending").
	Direction string `json:"direction,omitempty"`

	// Property: The property to order by.
	Property *PropertyReference `json:"property,omitempty"`
}

type PropertyReference struct {
	// Name: The name of the property.
	Name string `json:"name,omitempty"`
}

type Query struct {
	// EndCursor: An upper bound on the query results (optional).
	EndCursor string `json:"endCursor,omitempty"`

	// Filter: The filter to apply (optional).
	Filter *Filter `json:"filter,omitempty"`

	// GroupBy: The properties to group by (if empty, no grouping is applied
	// to the result set).
	GroupBy []*PropertyReference `json:"groupBy,omitempty"`

	// Kinds: The kinds to query (if empty, returns entities from all
	// kinds).
	Kinds []*KindExpression `json:"kinds,omitempty"`

	// Limit: The maximum number of results to return. Applies after all
	// other constraints (optional).
	Limit int64 `json:"limit,omitempty"`

	// Offset: The number of results to skip. Applies before limit, but
	// after all other constraints (optional, defaults to 0).
	Offset int64 `json:"offset,omitempty"`

	// Order: The order to apply to the query results (if empty, order is
	// unspecified).
	Order []*PropertyOrder `json:"order,omitempty"`

	// Projection: The projection to return. If not set the entire entity is
	// returned.
	Projection []*PropertyExpression `json:"projection,omitempty"`

	// StartCursor: A lower bound on the query results (optional).
	StartCursor string `json:"startCursor,omitempty"`
}

type QueryResultBatch struct {
	// EndCursor: A cursor that points to the position after the last result
	// in the batch. May be absent.
	EndCursor string `json:"endCursor,omitempty"`

	// EntityResultType: The result type for every entity in entityResults.
	// "full" for full entities, "projection" for entities with only
	// projected properties, "keyOnly" for entities with only a key.
	EntityResultType string `json:"entityResultType,omitempty"`

	// EntityResults: The results for this batch.
	EntityResults []*EntityResult `json:"entityResults,omitempty"`

	// MoreResults: The state of the query after the current batch. One of
	// "notFinished", "moreResultsAfterLimit", "noMoreResults".
	MoreResults string `json:"moreResults,omitempty"`

	// SkippedResults: The number of results skipped because of
	// Query.offset.
	SkippedResults int64 `json:"skippedResults,omitempty"`
}

type ReadOptions struct {
	// ReadConsistency: The read consistency to use. One of "default",
	// "strong" or "eventual". Cannot be set when transaction is set. Get
	// and ancestor Queries default to "strong", global Queries default to
	// "eventual" and cannot be set to "strong" (optional, defaults to
	// "default").
	ReadConsistency string `json:"readConsistency,omitempty"`

	// Transaction: The transaction to use (optional).
	Transaction string `json:"transaction,omitempty"`
}

type RollbackRequest struct {
	// Transaction: The transaction identifier.
	Transaction string `json:"transaction,omitempty"`
}

type RollbackResponse struct {
	// Kind: The kind, fixed to "datastore#rollbackResponse".
	Kind string `json:"kind,omitempty"`
}

type RunQueryRequest struct {
	// PartitionId: The partition to query. Unlike other RPCs, only a single
	// partition can be queried at a time (optional, defaults to empty
	// namespace).
	PartitionId *PartitionId `json:"partitionId,omitempty"`

	// Query: The query to run.
	Query *Query `json:"query,omitempty"`

	// ReadOptions: The options for this query (optional).
	ReadOptions *ReadOptions `json:"readOptions,omitempty"`
}

type RunQueryResponse struct {
	// Batch: A batch of query results (always present).
	Batch *QueryResultBatch `json:"batch,omitempty"`

	// Kind: The kind, fixed to "datastore#runQueryResponse".
	Kind string `json:"kind,omitempty"`
}

type Value struct {
	// BlobKeyValue: A blob key value.
	BlobKeyValue string `json:"blobKeyValue,omitempty"`

	// BlobValue: A blob value.
	BlobValue string `json:"blobValue,omitempty"`

	// BooleanValue: A boolean value.
	BooleanValue bool `json:"booleanValue,omitempty"`

	// DateTimeValue: A timestamp value.
	DateTimeValue string `json:"dateTimeValue,omitempty"`

	// DoubleValue: A double value.
	DoubleValue float64 `json:"doubleValue,omitempty"`

	// EntityValue: An entity value. May have no key.
	EntityValue *Entity `json:"entityValue,omitempty"`

	// Indexed: If the value should be indexed.
	//
	// Indexed may be set for a
	// null value. When indexed is true, string_value is limited to 500
	// characters, blob_value is limited to 500 bytes, meaning 15 and 22 are
	// not allowed, and meaning 16 will be ignored on input (and will never
	// be set on output). Input values may explicitly set indexed to true,
	// but otherwise true is always represented by leaving indexed unset. An
	// entity value must not have indexed set to true, unless meaning is 9,
	// 20 or 21.
	Indexed bool `json:"indexed,omitempty"`

	// IntegerValue: An integer value.
	IntegerValue int64 `json:"integerValue,omitempty,string"`

	// KeyValue: A key value.
	KeyValue *Key `json:"keyValue,omitempty"`

	// Meaning: Explicit use of the 'meaning' field is discouraged. The
	// field is provided only for legacy support.
	Meaning int64 `json:"meaning,omitempty"`

	// StringValue: A Utf8 encoded string value.
	StringValue string `json:"stringValue,omitempty"`
}

// method id "datastore.datasets.allocateIds":

type DatasetsAllocateIdsCall struct {
	s                  *Service
	datasetId          string
	allocateidsrequest *AllocateIdsRequest
	opt_               map[string]interface{}
}

// AllocateIds: Allocate ids for incomplete keys (useful for referencing
// an entity before it is inserted).
func (r *DatasetsService) AllocateIds(datasetId string, allocateidsrequest *AllocateIdsRequest) *DatasetsAllocateIdsCall {
	c := &DatasetsAllocateIdsCall{s: r.s, opt_: make(map[string]interface{})}
	c.datasetId = datasetId
	c.allocateidsrequest = allocateidsrequest
	return c
}

func (c *DatasetsAllocateIdsCall) Do() (*AllocateIdsResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.allocateidsrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/datastore/v1beta1/datasets/", "{datasetId}/allocateIds")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{datasetId}", url.QueryEscape(c.datasetId), 1)
	req.URL.Opaque = "//" + req.URL.Host + req.URL.Path
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(AllocateIdsResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Allocate ids for incomplete keys (useful for referencing an entity before it is inserted).",
	//   "httpMethod": "POST",
	//   "id": "datastore.datasets.allocateIds",
	//   "parameterOrder": [
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Identifies the dataset.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{datasetId}/allocateIds",
	//   "request": {
	//     "$ref": "AllocateIdsRequest"
	//   },
	//   "response": {
	//     "$ref": "AllocateIdsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "datastore.datasets.beginTransaction":

type DatasetsBeginTransactionCall struct {
	s                       *Service
	datasetId               string
	begintransactionrequest *BeginTransactionRequest
	opt_                    map[string]interface{}
}

// BeginTransaction: Begin a new transaction.
func (r *DatasetsService) BeginTransaction(datasetId string, begintransactionrequest *BeginTransactionRequest) *DatasetsBeginTransactionCall {
	c := &DatasetsBeginTransactionCall{s: r.s, opt_: make(map[string]interface{})}
	c.datasetId = datasetId
	c.begintransactionrequest = begintransactionrequest
	return c
}

func (c *DatasetsBeginTransactionCall) Do() (*BeginTransactionResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.begintransactionrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/datastore/v1beta1/datasets/", "{datasetId}/beginTransaction")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{datasetId}", url.QueryEscape(c.datasetId), 1)
	req.URL.Opaque = "//" + req.URL.Host + req.URL.Path
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(BeginTransactionResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Begin a new transaction.",
	//   "httpMethod": "POST",
	//   "id": "datastore.datasets.beginTransaction",
	//   "parameterOrder": [
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Identifies the dataset.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{datasetId}/beginTransaction",
	//   "request": {
	//     "$ref": "BeginTransactionRequest"
	//   },
	//   "response": {
	//     "$ref": "BeginTransactionResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "datastore.datasets.blindWrite":

type DatasetsBlindWriteCall struct {
	s                 *Service
	datasetId         string
	blindwriterequest *BlindWriteRequest
	opt_              map[string]interface{}
}

// BlindWrite: Write a mutation outside of a transaction.
func (r *DatasetsService) BlindWrite(datasetId string, blindwriterequest *BlindWriteRequest) *DatasetsBlindWriteCall {
	c := &DatasetsBlindWriteCall{s: r.s, opt_: make(map[string]interface{})}
	c.datasetId = datasetId
	c.blindwriterequest = blindwriterequest
	return c
}

func (c *DatasetsBlindWriteCall) Do() (*BlindWriteResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.blindwriterequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/datastore/v1beta1/datasets/", "{datasetId}/blindWrite")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{datasetId}", url.QueryEscape(c.datasetId), 1)
	req.URL.Opaque = "//" + req.URL.Host + req.URL.Path
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(BlindWriteResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Write a mutation outside of a transaction.",
	//   "httpMethod": "POST",
	//   "id": "datastore.datasets.blindWrite",
	//   "parameterOrder": [
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Identifies the dataset.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{datasetId}/blindWrite",
	//   "request": {
	//     "$ref": "BlindWriteRequest"
	//   },
	//   "response": {
	//     "$ref": "BlindWriteResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "datastore.datasets.commit":

type DatasetsCommitCall struct {
	s             *Service
	datasetId     string
	commitrequest *CommitRequest
	opt_          map[string]interface{}
}

// Commit: Commit a transaction.
func (r *DatasetsService) Commit(datasetId string, commitrequest *CommitRequest) *DatasetsCommitCall {
	c := &DatasetsCommitCall{s: r.s, opt_: make(map[string]interface{})}
	c.datasetId = datasetId
	c.commitrequest = commitrequest
	return c
}

func (c *DatasetsCommitCall) Do() (*CommitResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.commitrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/datastore/v1beta1/datasets/", "{datasetId}/commit")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{datasetId}", url.QueryEscape(c.datasetId), 1)
	req.URL.Opaque = "//" + req.URL.Host + req.URL.Path
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(CommitResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Commit a transaction.",
	//   "httpMethod": "POST",
	//   "id": "datastore.datasets.commit",
	//   "parameterOrder": [
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Identifies the dataset.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{datasetId}/commit",
	//   "request": {
	//     "$ref": "CommitRequest"
	//   },
	//   "response": {
	//     "$ref": "CommitResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "datastore.datasets.lookup":

type DatasetsLookupCall struct {
	s             *Service
	datasetId     string
	lookuprequest *LookupRequest
	opt_          map[string]interface{}
}

// Lookup: Lookup entities.
func (r *DatasetsService) Lookup(datasetId string, lookuprequest *LookupRequest) *DatasetsLookupCall {
	c := &DatasetsLookupCall{s: r.s, opt_: make(map[string]interface{})}
	c.datasetId = datasetId
	c.lookuprequest = lookuprequest
	return c
}

func (c *DatasetsLookupCall) Do() (*LookupResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.lookuprequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/datastore/v1beta1/datasets/", "{datasetId}/lookup")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{datasetId}", url.QueryEscape(c.datasetId), 1)
	req.URL.Opaque = "//" + req.URL.Host + req.URL.Path
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(LookupResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lookup entities.",
	//   "httpMethod": "POST",
	//   "id": "datastore.datasets.lookup",
	//   "parameterOrder": [
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Identifies the dataset.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{datasetId}/lookup",
	//   "request": {
	//     "$ref": "LookupRequest"
	//   },
	//   "response": {
	//     "$ref": "LookupResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "datastore.datasets.rollback":

type DatasetsRollbackCall struct {
	s               *Service
	datasetId       string
	rollbackrequest *RollbackRequest
	opt_            map[string]interface{}
}

// Rollback: Rollback a transaction.
func (r *DatasetsService) Rollback(datasetId string, rollbackrequest *RollbackRequest) *DatasetsRollbackCall {
	c := &DatasetsRollbackCall{s: r.s, opt_: make(map[string]interface{})}
	c.datasetId = datasetId
	c.rollbackrequest = rollbackrequest
	return c
}

func (c *DatasetsRollbackCall) Do() (*RollbackResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.rollbackrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/datastore/v1beta1/datasets/", "{datasetId}/rollback")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{datasetId}", url.QueryEscape(c.datasetId), 1)
	req.URL.Opaque = "//" + req.URL.Host + req.URL.Path
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(RollbackResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Rollback a transaction.",
	//   "httpMethod": "POST",
	//   "id": "datastore.datasets.rollback",
	//   "parameterOrder": [
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Identifies the dataset.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{datasetId}/rollback",
	//   "request": {
	//     "$ref": "RollbackRequest"
	//   },
	//   "response": {
	//     "$ref": "RollbackResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "datastore.datasets.runQuery":

type DatasetsRunQueryCall struct {
	s               *Service
	datasetId       string
	runqueryrequest *RunQueryRequest
	opt_            map[string]interface{}
}

// RunQuery: Query for entities.
func (r *DatasetsService) RunQuery(datasetId string, runqueryrequest *RunQueryRequest) *DatasetsRunQueryCall {
	c := &DatasetsRunQueryCall{s: r.s, opt_: make(map[string]interface{})}
	c.datasetId = datasetId
	c.runqueryrequest = runqueryrequest
	return c
}

func (c *DatasetsRunQueryCall) Do() (*RunQueryResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.runqueryrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/datastore/v1beta1/datasets/", "{datasetId}/runQuery")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{datasetId}", url.QueryEscape(c.datasetId), 1)
	req.URL.Opaque = "//" + req.URL.Host + req.URL.Path
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(RunQueryResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Query for entities.",
	//   "httpMethod": "POST",
	//   "id": "datastore.datasets.runQuery",
	//   "parameterOrder": [
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Identifies the dataset.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{datasetId}/runQuery",
	//   "request": {
	//     "$ref": "RunQueryRequest"
	//   },
	//   "response": {
	//     "$ref": "RunQueryResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}
