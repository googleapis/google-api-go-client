// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package disco represents Google API discovery documents.
package disco

import (
	"encoding/json"
	"fmt"
	"sort"
)

// A Document is an API discovery document.
type Document struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Version     string             `json:"version"`
	Title       string             `json:"title"`
	RootURL     string             `json:"rootUrl"`
	ServicePath string             `json:"servicePath"`
	BasePath    string             `json:"basePath"`
	Auth        Auth               `json:"auth"`
	Features    []string           `json:"features"`
	Schemas     map[string]*Schema `json:"schemas"`
	// TODO(jba): resources
}

// init performs additional initialization and checks that
// were not done during unmarshaling.
func (d *Document) init() error {
	schemasByID := map[string]*Schema{}
	for _, s := range d.Schemas {
		schemasByID[s.ID] = s
	}
	for name, s := range d.Schemas {
		if s.Ref != "" {
			return fmt.Errorf("top level schema %q is a reference", name)
		}
		s.Name = name
		if err := s.init(schemasByID); err != nil {
			return err
		}
	}
	return nil
}

// NewDocument unmarshals the bytes into a Document.
// It also validates the document to make sure it is error-free.
func NewDocument(bytes []byte) (*Document, error) {
	var doc Document
	if err := json.Unmarshal(bytes, &doc); err != nil {
		return nil, err
	}
	if err := doc.init(); err != nil {
		return nil, err
	}
	return &doc, nil
}

// Auth represents the auth section of a discovery document.
// Only OAuth2 information is retained.
type Auth struct {
	OAuth2Scopes []Scope
}

// A Scope is an OAuth2 scope.
type Scope struct {
	URL         string
	Description string
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (a *Auth) UnmarshalJSON(data []byte) error {
	// Pull out the oauth2 scopes and turn them into nice structs.
	// Ignore other auth information.
	var m struct {
		OAuth2 struct {
			Scopes map[string]struct {
				Description string
			}
		}
	}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	var keys []string
	for k := range m.OAuth2.Scopes {
		keys = append(keys, k)
	}
	// Sort to provide a deterministic ordering, mainly for testing.
	sort.Strings(keys)
	for _, k := range keys {
		a.OAuth2Scopes = append(a.OAuth2Scopes, Scope{
			URL:         k,
			Description: m.OAuth2.Scopes[k].Description,
		})
	}
	return nil
}

// A Schema holds a JSON Schema as defined by
// https://tools.ietf.org/html/draft-zyp-json-schema-03#section-5.1.
// We only support the subset of JSON Schema needed for Google API generation.
type Schema struct {
	ID                   string // union types not supported
	Type                 string // union types not supported
	Format               string
	Description          string
	Properties           map[string]*Schema
	ItemSchema           *Schema `json:"items"` // array of schemas not supported
	AdditionalProperties *Schema // boolean not supported
	Ref                  string  `json:"$ref"`
	Default              string
	Pattern              string
	Enums                []string `json:"enum"`
	// Google extensions to JSON Schema
	EnumDescriptions []string
	Variant          map[string]interface{}

	RefSchema *Schema `json:"-"` // Schema referred to by $ref
	Name      string  `json:"-"` // Schema name, if top level
	Kind      Kind    `json:"-"`
}

func (s *Schema) init(topLevelSchemas map[string]*Schema) error {
	if s == nil {
		return nil
	}
	if s.Ref != "" {
		rs, ok := topLevelSchemas[s.Ref]
		if !ok {
			return fmt.Errorf("could not resolve schema reference %q", s.Ref)
		}
		s.RefSchema = rs
	}
	var err error
	s.Kind, err = s.initKind()
	if err != nil {
		return err
	}
	if s.Kind == ArrayKind && s.ItemSchema == nil {
		return fmt.Errorf("schema %+v: array does not have items", s)
	}
	if s.Kind != ArrayKind && s.ItemSchema != nil {
		return fmt.Errorf("schema %+v: non-array has items", s)
	}
	if err := s.AdditionalProperties.init(topLevelSchemas); err != nil {
		return err
	}
	if err := s.ItemSchema.init(topLevelSchemas); err != nil {
		return err
	}
	for _, p := range s.Properties {
		if err := p.init(topLevelSchemas); err != nil {
			return err
		}
	}
	return nil
}

func (s *Schema) initKind() (Kind, error) {
	if s.Ref != "" {
		return ReferenceKind, nil
	}
	switch s.Type {
	case "string", "number", "integer", "boolean", "any":
		return SimpleKind, nil
	case "object":
		if s.AdditionalProperties != nil {
			if s.AdditionalProperties.Type == "any" {
				return AnyStructKind, nil
			}
			return MapKind, nil
		}
		return StructKind, nil
	case "array":
		return ArrayKind, nil
	default:
		return 0, fmt.Errorf("unknown type %q for schema %q", s.Type, s.ID)
	}
}

// ElementSchema returns the schema for the element type of s. For maps,
// this is the schema of the map values. For arrays, it is the schema
// of the array item type.
//
// ElementSchema panics if called on a schema that is not of kind map or array.
func (s *Schema) ElementSchema() *Schema {
	switch s.Kind {
	case MapKind:
		return s.AdditionalProperties
	case ArrayKind:
		return s.ItemSchema
	default:
		panic("ElementSchema called on schema of type " + s.Type)
	}
}

// Kind classifies a Schema.
type Kind int

const (
	// SimpleKind is the category for any JSON Schema that maps to a
	// primitive Go type: strings, numbers, booleans, and "any" (since it
	// maps to interface{}).
	SimpleKind Kind = iota

	// StructKind is the category for a JSON Schema that declares a JSON
	// object without any additional (arbitrary) properties.
	StructKind

	// MapKind is the category for a JSON Schema that declares a JSON
	// object with additional (arbitrary) properties that have a non-"any"
	// schema type.
	MapKind

	// AnyStructKind is the category for a JSON Schema that declares a
	// JSON object with additional (arbitrary) properties that can be any
	// type.
	AnyStructKind

	// ArrayKind is the category for a JSON Schema that declares an
	// "array" type.
	ArrayKind

	// ReferenceKind is the category for a JSON Schema that is a reference
	// to another JSON Schema.  During code generation, these references
	// are resolved using the API.schemas map.
	// See https://tools.ietf.org/html/draft-zyp-json-schema-03#section-5.28
	// for more details on the format.
	ReferenceKind
)
