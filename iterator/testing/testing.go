// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package testing provides support functions for testing iterators conforming
// to the standard pattern.
// See package google.golang.org/api/iterator and
// https://github.com/GoogleCloudPlatform/gcloud-golang/wiki/Iterator-Guidelines.
package testing

import (
	"fmt"
	"reflect"

	"google.golang.org/api/iterator"
)

// TestIterator tests the Next method of a standard iterator. It assumes that
// the underlying sequence to be iterated over already exists.
//
// The want argument should be a slice that contains the elements of this
// sequence. It may be an empty slice, but it must not be the nil interface
// value. The elements must be comparable with reflect.DeepEqual.
//
// The create function should create and return a new iterator.
// It will typically look like
//    func() interface{} { return client.Items(ctx) }
//
// The next function takes the return value of create and should return the
// result of calling Next on the iterator. It can usually be defined as
//     func(it interface{}) (interface{}, error) { return it.(*ItemIterator).Next() }
//
// TestIterator checks that the iterator returns all the elements of want
// in order, followed by (zero, done). It also confirms that subsequent calls
// to next also return (zero, done).
//
// If the iterator implements the method
//     PageInfo() *iterator.PageInfo
// then exact pagination with iterator.Pager is also tested. Pagination testing
// will be more informative if the want slice contains at least three elements.
//
// On success, TestIterator returns ("", true). On failure, it returns a
// suitable error message and false.
func TestIterator(want interface{}, create func() interface{}, next func(interface{}) (interface{}, error)) (string, bool) {
	vWant := reflect.ValueOf(want)
	if vWant.Kind() != reflect.Slice {
		return "'want' must be a slice", false
	}
	it := create()
	msg, ok := testNext(vWant, it, next)
	if !ok {
		return msg, ok
	}
	if _, ok := it.(iterator.Pageable); !ok {
		return "", true
	}
	return testPaging(vWant, create, next)
}

// Check that the iterator returns vWant, the desired sequence.
func testNext(vWant reflect.Value, it interface{}, next func(interface{}) (interface{}, error)) (string, bool) {
	for i := 0; i < vWant.Len(); i++ {
		got, err := next(it)
		if err != nil {
			return fmt.Sprintf("#%d: got %v, expected an item", i, err), false
		}
		w := vWant.Index(i).Interface()
		if !reflect.DeepEqual(got, w) {
			return fmt.Sprintf("#%d: got %+v, want %+v", i, got, w), false
		}
	}
	// We now should see (<zero value of item type>, done), no matter how many
	// additional calls we make.
	zero := reflect.Zero(vWant.Type().Elem()).Interface()
	for i := 0; i < 3; i++ {
		got, err := next(it)
		if err != iterator.Done {
			return fmt.Sprintf("at end: got error %v, want iterator.Done", err), false
		}
		// Since err == iterator.Done, got should be zero.
		if got != zero {
			return fmt.Sprintf("got %+v with done, want zero %T", got, zero), false
		}
	}
	return "", true
}

// Test the iterator's behavior when used with iterator.Pager.
func testPaging(vWant reflect.Value, create func() interface{}, next func(interface{}) (interface{}, error)) (string, bool) {
	for _, pageSize := range []int{1, 2, vWant.Len(), vWant.Len() + 10} {
		// Create the pages we expect to see.
		var wantPages []interface{}
		for i, j := 0, pageSize; i < vWant.Len(); i, j = j, j+pageSize {
			if j > vWant.Len() {
				j = vWant.Len()
			}
			wantPages = append(wantPages, vWant.Slice(i, j).Interface())
		}
		for _, usePageToken := range []bool{false, true} {
			it := create().(iterator.Pageable)
			var pager *iterator.Pager
			if !usePageToken {
				pager = iterator.NewPager(it, pageSize, "")
			}
			tok := ""
			var err error
			for i, wantPage := range wantPages {
				vpagep := reflect.New(vWant.Type())
				pagep := vpagep.Interface()
				if usePageToken {
					tok, err = iterator.NewPager(it, pageSize, tok).NextPage(pagep)
				} else {
					tok, err = pager.NextPage(pagep)
				}
				if err != nil {
					return fmt.Sprintf("use page token %t, pageSize %d, page #%d: got error %v",
						usePageToken, pageSize, i, err), false
				}
				gotPage := vpagep.Elem().Interface()
				if !reflect.DeepEqual(gotPage, wantPage) {
					return fmt.Sprintf("use page token %t, pageSize %d, page #%d:\ngot  %v\nwant %+v",
						usePageToken, pageSize, i, gotPage, wantPage), false
				}
				if tok == "" && i != len(wantPages)-1 {
					return fmt.Sprintf("use page token %t, pageSize %d, page #%d: got empty page token",
						usePageToken, pageSize, i, gotPage, wantPage), false
				}
			}
			if tok != "" {
				return fmt.Sprintf("use page token %t, pageSize %d: non-empty page token at end",
					usePageToken, pageSize), false
			}
		}
	}
	return "", true
}
