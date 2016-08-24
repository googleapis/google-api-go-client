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

package iterator_test

import (
	"reflect"
	"strconv"
	"testing"

	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

type service struct {
	end int
}

// List simulates an API List RPC.
func (s *service) List(pageSize int, pageToken string) ([]int, string, error) {
	// Never give back any more than 3.
	if pageSize <= 0 || pageSize > 3 {
		pageSize = 3
	}
	var start int
	if pageToken != "" {
		var err error
		start, err = strconv.Atoi(pageToken)
		if err != nil {
			return nil, "", err
		}
	}
	var ints []int
	var i int
	for i = start; i < s.end && len(ints) < pageSize; i++ {
		ints = append(ints, i)
	}
	if i == s.end {
		pageToken = ""
	} else {
		pageToken = strconv.Itoa(i)
	}
	return ints, pageToken, nil
}

type Client struct{ s *service }

// ItemIterator is a sample implementation of a standard iterator.
type ItemIterator struct {
	pageInfo *iterator.PageInfo
	nextFunc func(error) (error, bool)
	s        *service
	items    []int
	err      error
}

// Page returns a PageInfo, which supports pagination.
func (it *ItemIterator) PageInfo() *iterator.PageInfo { return it.pageInfo }

// Items is a sample implementation of an iterator-creating method.
func (c *Client) Items(ctx context.Context) *ItemIterator {
	it := &ItemIterator{s: c.s}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(
		it.fetch,
		func() int { return len(it.items) },
		func() interface{} { b := it.items; it.items = nil; return b })
	return it
}

func (it *ItemIterator) fetch(pageSize int, pageToken string) (string, error) {
	items, tok, err := it.s.List(pageSize, pageToken)
	it.items = append(it.items, items...)
	return tok, err
}

func (it *ItemIterator) Next() (int, error) {
	var ok bool
	it.err, ok = it.nextFunc(it.err)
	if !ok {
		return 0, it.err
	}
	item := it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func TestIterator(t *testing.T) {
	client := &Client{&service{end: 10}}
	it := client.Items(ctx)
	for i := 0; i < 10; i++ {
		n, err := it.Next()
		if err != nil {
			t.Fatalf("%d: got %v, want nil", i, err)
		}
		if got, want := n, i; got != want {
			t.Fatalf("got %d, want %d", got, want)
		}
	}
	_, err := it.Next()
	if err != iterator.Done {
		t.Fatalf("end: got %v, want Done", err)
	}
}

func TestPager(t *testing.T) {
	const pageSize = 4
	client := &Client{&service{end: 11}}
	it := client.Items(ctx)
	p := iterator.NewPager(it, pageSize, "")
	for i := 0; i < client.s.end; i += pageSize {
		var items []int
		tok, err := p.NextPage(&items)
		if err != nil {
			t.Fatal(err)
		}
		end := i + pageSize
		if end > client.s.end {
			end = client.s.end
		}
		if got, want := items, seq(i, end); !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
		if (tok == "") != (i == 8) {
			t.Fatalf("got token %q, want the opposite", tok)
		}
	}
}

// Verify that, for an iterator that uses PageInfo.next to implement its Next
// method, using Next and NextPage together result in an error.
func TestNextWithNextPage(t *testing.T) {
	client := &Client{&service{end: 11}}
	var items []int

	// Calling Next before NextPage.
	it := client.Items(ctx)
	it.Next()
	_, err := iterator.NewPager(it, 1, "").NextPage(&items)
	if err == nil {
		t.Error("NextPage after Next: got nil, want error")
	}
	_, err = it.Next()
	if err == nil {
		t.Error("Next after NextPage: got nil, want error")
	}

	// Next between two calls to NextPage.
	it = client.Items(ctx)
	p := iterator.NewPager(it, 1, "")
	p.NextPage(&items)
	_, err = it.Next()
	if err == nil {
		t.Error("Next after NextPage: got nil, want error")
	}
	_, err = p.NextPage(&items)
	if err == nil {
		t.Error("second NextPage after Next: got nil, want error")
	}
}

// Verify that we turn various potential reflection panics into errors.
func TestNextPageReflectionErrors(t *testing.T) {
	client := &Client{&service{end: 1}}
	p := iterator.NewPager(client.Items(ctx), 1, "")

	// Passing the nil interface value.
	_, err := p.NextPage(nil)
	if err == nil {
		t.Error("nil: got nil, want error")
	}

	// Passing a non-slice.
	_, err = p.NextPage(17)
	if err == nil {
		t.Error("non-slice: got nil, want error")
	}

	// Passing a slice of the wrong type.
	// TODO(jba): make this pass.
	// var bools []bool
	// _, err = p.NextPage(&bools)
	// if err == nil {
	// 	t.Error("wrong type: got nil, want error")
	// }

	// Using a slice of the right type, but not passing a pointer to it.
	var ints []int
	_, err = p.NextPage(ints)
	if err == nil {
		t.Error("not a pointer: got nil, want error")
	}
}

func seq(from, to int) []int {
	var r []int
	for i := from; i < to; i++ {
		r = append(r, i)
	}
	return r
}
