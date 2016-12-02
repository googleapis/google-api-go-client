// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package disco represents Google API discovery documents.
package disco

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"sort"
)

// A Document is an API discovery document.
type Document struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Version       string `json:"version"`
	Title         string `json:"title"`
	DiscoveryLink string `json:"discoveryRestUrl"` // absolute
	RootURL       string `json:"rootUrl"`
	ServicePath   string `json:"servicePath"`
	Preferred     bool   `json:"preferred"`
	BasePath      string `json:"basePath"`
	Auth          Auth   `json:"auth"`
	// TODO(jba): schemas
	// TODO(jba): resources

	jsonBytes []byte // JSON from which this was unmarshaled
}

// ReadDocument unmarshals the entire contents of the reader into a Document.
// It also validates the document to make sure it is error-free.
func ReadDocument(r io.Reader) (*Document, error) {
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	doc := &Document{jsonBytes: bytes}
	if err := json.Unmarshal(bytes, doc); err != nil {
		return nil, err
	}
	if err := doc.init(); err != nil {
		return nil, err
	}
	return doc, nil
}

// JSONBytes returns the JSON from which the document was unmarshaled.
func (d *Document) JSONBytes() []byte {
	return d.jsonBytes
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
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	scopes := jobj(m, "oauth2", "scopes")
	var keys []string
	for k := range scopes {
		keys = append(keys, k)
	}
	// Sort to provide a deterministic ordering, mainly for testing.
	sort.Strings(keys)
	for _, k := range keys {
		a.OAuth2Scopes = append(a.OAuth2Scopes, Scope{
			URL:         k,
			Description: scopes[k].(map[string]interface{})["description"].(string),
		})
	}
	return nil
}

// init performs additional initialization and checks that
// were not done during unmarshaling.
func (d *Document) init() error {
	return nil
}

func jobj(m map[string]interface{}, keys ...string) map[string]interface{} {
	for _, key := range keys {
		m2, ok := m[key].(map[string]interface{})
		if !ok {
			return nil
		}
		m = m2
	}
	return m
}
