// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package disco

import (
	"io/ioutil"
	"reflect"
	"testing"
)

var stringSchema = &Schema{
	Type: "string",
	Kind: SimpleKind,
}

func TestDocument(t *testing.T) {
	bytes, err := ioutil.ReadFile("testdata/test-api.json")
	if err != nil {
		t.Fatal(err)
	}
	got, err := NewDocument(bytes)
	if err != nil {
		t.Fatal(err)
	}
	want := &Document{
		ID:          "storage:v1",
		Name:        "storage",
		Version:     "v1",
		Title:       "Cloud Storage JSON API",
		RootURL:     "https://www.googleapis.com/",
		ServicePath: "storage/v1/",
		BasePath:    "/storage/v1/",
		Auth: Auth{
			OAuth2Scopes: []Scope{
				{"https://www.googleapis.com/auth/cloud-platform",
					"View and manage your data across Google Cloud Platform services"},
				{"https://www.googleapis.com/auth/cloud-platform.read-only",
					"View your data across Google Cloud Platform services"},
				{"https://www.googleapis.com/auth/devstorage.full_control",
					"Manage your data and permissions in Google Cloud Storage"},
				{"https://www.googleapis.com/auth/devstorage.read_only",
					"View your data in Google Cloud Storage"},
				{"https://www.googleapis.com/auth/devstorage.read_write",
					"Manage your data in Google Cloud Storage"},
			},
		},
		Features: []string{"dataWrapper"},
		Schemas: map[string]*Schema{
			"Bucket": &Schema{
				Name:        "Bucket",
				ID:          "Bucket",
				Type:        "object",
				Description: "A bucket.",
				Kind:        StructKind,
				Properties: map[string]*Schema{
					"id": stringSchema,
					"kind": &Schema{
						Type:    "string",
						Kind:    SimpleKind,
						Default: "storage#bucket",
					},
					"cors": &Schema{
						Type: "array",
						Kind: ArrayKind,
						ItemSchema: &Schema{
							Type: "object",
							Kind: StructKind,
							Properties: map[string]*Schema{
								"maxAgeSeconds": &Schema{
									Type:   "integer",
									Format: "int32",
									Kind:   SimpleKind,
								},
								"method": &Schema{
									Type:       "array",
									Kind:       ArrayKind,
									ItemSchema: stringSchema,
								},
							},
						},
					},
				},
			},
			"Buckets": &Schema{
				ID:   "Buckets",
				Name: "Buckets",
				Type: "object",
				Kind: StructKind,
				Properties: map[string]*Schema{
					"items": &Schema{
						Type: "array",
						Kind: ArrayKind,
						ItemSchema: &Schema{
							Kind:      ReferenceKind,
							Ref:       "Bucket",
							RefSchema: nil,
						},
					},
				},
			},
		},
	}
	// Resolve schema reference.
	want.Schemas["Buckets"].Properties["items"].ItemSchema.RefSchema = want.Schemas["Bucket"]
	for k, gs := range got.Schemas {
		ws := want.Schemas[k]
		if !reflect.DeepEqual(gs, ws) {
			t.Fatalf("schema %s: got\n%+v\nwant\n%+v", k, gs, ws)
		}
	}
	if len(got.Schemas) != len(want.Schemas) {
		t.Errorf("want %d schemas, got %d", len(got.Schemas), len(want.Schemas))
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}

func TestDocumentErrors(t *testing.T) {
	for _, in := range []string{
		`{"name": "X"`, // malformed JSON
		`{"id": 3}`,    // ID is an int instead of a string
		`{"auth": "oauth2": { "scopes": "string" }}`, // wrong auth structure
	} {
		_, err := NewDocument([]byte(in))
		if err == nil {
			t.Errorf("%s: got nil, want error", in)
		}
	}
}

func TestSchemaErrors(t *testing.T) {
	for _, s := range []*Schema{
		{Type: "array"},                         // missing item schema
		{Type: "string", ItemSchema: &Schema{}}, // items w/o array
		{Type: "moose"},                         // bad kind
		{Ref: "Thing"},                          // unresolved reference
	} {
		if err := s.init(nil); err == nil {
			t.Errorf("%+v: got nil, want error", s)
		}
	}
}
