// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"encoding/json"
	"reflect"
	"testing"

	"google.golang.org/api/googleapi"
)

type schema struct {
	// Basic types
	B    bool    `json:"b,omitempty"`
	F    float64 `json:"f,omitempty"`
	I    int64   `json:"i,omitempty"`
	Istr int64   `json:"istr,omitempty,string"`
	Str  string  `json:"str,omitempty"`

	// Pointers to basic types
	PB    *bool    `json:"pb,omitempty"`
	PF    *float64 `json:"pf,omitempty"`
	PI    *int64   `json:"pi,omitempty"`
	PIstr *int64   `json:"pistr,omitempty,string"`
	PStr  *string  `json:"pstr,omitempty"`

	// Other types
	Int64s googleapi.Int64s  `json:"i64s,omitempty"`
	S      []int             `json:"s,omitempty"`
	M      map[string]string `json:"m,omitempty"`
	Any    interface{}       `json:"any,omitempty"`
	Child  *child            `json:"child,omitempty"`
}

type child struct {
	B bool `json:"childbool,omitempty"`
}

type testCase struct {
	s           interface{}
	mustInclude []string
	want        string
}

func TestBasics(t *testing.T) {
	for _, tc := range []testCase{
		{
			s:    schema{},
			want: `{}`,
		},
		{
			s:           schema{},
			mustInclude: []string{"B", "F", "I", "Istr", "Str", "PB", "PF", "PI", "PIStr", "PStr"},
			want:        `{"b":false,"f":0.0,"i":0,"istr":"0","str":""}`,
		},
		{
			s: schema{
				B:     true,
				F:     1.2,
				I:     1,
				Istr:  2,
				Str:   "a",
				PB:    googleapi.Bool(true),
				PF:    googleapi.Float64(1.2),
				PI:    googleapi.Int64(int64(1)),
				PIstr: googleapi.Int64(int64(2)),
				PStr:  googleapi.String("a"),
			},
			want: `{"b":true,"f":1.2,"i":1,"istr":"2","str":"a","pb":true,"pf":1.2,"pi":1,"pistr":"2","pstr":"a"}`,
		},
		{
			s: schema{
				B:     false,
				F:     0.0,
				I:     0,
				Istr:  0,
				Str:   "",
				PB:    googleapi.Bool(false),
				PF:    googleapi.Float64(0.0),
				PI:    googleapi.Int64(int64(0)),
				PIstr: googleapi.Int64(int64(0)),
				PStr:  googleapi.String(""),
			},
			want: `{"pb":false,"pf":0.0,"pi":0,"pistr":"0","pstr":""}`,
		},
		{
			s: schema{
				B:     false,
				F:     0.0,
				I:     0,
				Istr:  0,
				Str:   "",
				PB:    googleapi.Bool(false),
				PF:    googleapi.Float64(0.0),
				PI:    googleapi.Int64(int64(0)),
				PIstr: googleapi.Int64(int64(0)),
				PStr:  googleapi.String(""),
			},
			mustInclude: []string{"B", "F", "I", "Istr", "Str", "PB", "PF", "PI", "PIStr", "PStr"},
			want:        `{"b":false,"f":0.0,"i":0,"istr":"0","str":"","pb":false,"pf":0.0,"pi":0,"pistr":"0","pstr":""}`,
		},
	} {
		checkSchemaToMap(t, tc)
	}
}

func TestSliceFields(t *testing.T) {
	for _, tc := range []testCase{
		{
			s:    schema{},
			want: `{}`,
		},
		{
			s:    schema{S: []int{}, Int64s: googleapi.Int64s{}},
			want: `{}`,
		},
		{
			s:    schema{S: []int{1}, Int64s: googleapi.Int64s{1}},
			want: `{"s":[1],"i64s":["1"]}`,
		},
		{
			s:           schema{},
			mustInclude: []string{"S", "Int64s"},
			want:        `{"s":[],"i64s":[]}`,
		},
		{
			s:           schema{S: []int{}, Int64s: googleapi.Int64s{}},
			mustInclude: []string{"S", "Int64s"},
			want:        `{"s":[],"i64s":[]}`,
		},
		{
			s:           schema{S: []int{1}, Int64s: googleapi.Int64s{1}},
			mustInclude: []string{"S", "Int64s"},
			want:        `{"s":[1],"i64s":["1"]}`,
		},
	} {
		checkSchemaToMap(t, tc)
	}
}

func TestMapField(t *testing.T) {
	for _, tc := range []testCase{
		{
			s:    schema{},
			want: `{}`,
		},
		{
			s:    schema{M: make(map[string]string)},
			want: `{}`,
		},
		{
			s:    schema{M: map[string]string{"a": "b"}},
			want: `{"m":{"a":"b"}}`,
		},
		{
			s:           schema{},
			mustInclude: []string{"M"},
			want:        `{"m":{}}`,
		},
		{
			s:           schema{M: make(map[string]string)},
			mustInclude: []string{"M"},
			want:        `{"m":{}}`,
		},
		{
			s:           schema{M: map[string]string{"a": "b"}},
			mustInclude: []string{"M"},
			want:        `{"m":{"a":"b"}}`,
		},
	} {
		checkSchemaToMap(t, tc)
	}
}

type anyType struct {
	Field int
}

func (a anyType) MarshalJSON() ([]byte, error) {
	return []byte(`"anyType value"`), nil
}

func TestAnyField(t *testing.T) {
	// mustInclude has no effect on nil interfaces and interfaces that contain nil pointers.
	var nilAny *anyType
	for _, tc := range []testCase{
		{
			s:    schema{},
			want: `{}`,
		},
		{
			s:    schema{Any: nilAny},
			want: `{}`,
		},
		{
			s:    schema{Any: &anyType{}},
			want: `{"any":"anyType value"}`,
		},
		{
			s:    schema{Any: anyType{}},
			want: `{"any":"anyType value"}`,
		},
		{
			s:           schema{},
			mustInclude: []string{"Any"},
			want:        `{}`,
		},
		{
			s:           schema{Any: nilAny},
			mustInclude: []string{"Any"},
			want:        `{}`,
		},
		{
			s:           schema{Any: &anyType{}},
			mustInclude: []string{"Any"},
			want:        `{"any":"anyType value"}`,
		},
		{
			s:           schema{Any: anyType{}},
			mustInclude: []string{"Any"},
			want:        `{"any":"anyType value"}`,
		},
	} {
		checkSchemaToMap(t, tc)
	}
}

func TestSubschema(t *testing.T) {
	// Subschemas are always stored as pointers, so mustInclude has no effect on them.
	for _, tc := range []testCase{
		{
			s:    schema{},
			want: `{}`,
		},
		{
			s:           schema{},
			mustInclude: []string{"Child"},
			want:        `{}`,
		},
		{
			s:    schema{Child: &child{}},
			want: `{"child":{}}`,
		},
		{
			s:           schema{Child: &child{}},
			mustInclude: []string{"Child"},
			want:        `{"child":{}}`,
		},
		{
			s:    schema{Child: &child{B: true}},
			want: `{"child":{"childbool":true}}`,
		},

		{
			s:           schema{Child: &child{B: true}},
			mustInclude: []string{"Child"},
			want:        `{"child":{"childbool":true}}`,
		},
	} {
		checkSchemaToMap(t, tc)
	}
}

// checkSchemaToMap verifies that calling SchemaToMap on tc.s yields a result which is equivalent to tc.want.
func checkSchemaToMap(t *testing.T, tc testCase) {
	// We want to verify that the result of SchemaToMap is equivalent to
	// the expected result once it has been encoded to JSON.
	toEncode, err := SchemaToMap(tc.s, tc.mustInclude)
	if err != nil {
		t.Fatalf("SchemaToMap:\n got err: %v", err)
	}
	encoded, err := json.Marshal(toEncode)
	if err != nil {
		t.Fatalf("encoding json:\n got err: %v", err)
	}

	// The expected and obtained JSON can differ in field ordering, so unmarshal before comparing.
	var got interface{}
	var want interface{}
	err = json.Unmarshal(encoded, &got)
	if err != nil {
		t.Fatalf("decoding json:\n got err: %v", err)
	}
	err = json.Unmarshal([]byte(tc.want), &want)
	if err != nil {
		t.Fatalf("decoding json:\n got err: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SchemaToMap:\ngot :%s\nwant:%s", got, want)
	}
}
