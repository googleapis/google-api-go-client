// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package disco

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestDocument(t *testing.T) {
	f, err := os.Open("testdata/storage-api.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	got, err := ReadDocument(f)
	if err != nil {
		t.Fatal(err)
	}
	got.jsonBytes = nil // don't clutter the Errorf output
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
		_, err := ReadDocument(strings.NewReader(in))
		if err == nil {
			t.Errorf("%s: got nil, want error", in)
		}
	}
}
