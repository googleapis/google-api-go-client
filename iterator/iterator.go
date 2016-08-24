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

// Package iterator provides support for standard Google API iterators.
// See https://github.com/GoogleCloudPlatform/gcloud-golang/wiki/Iterator-Guidelines.
package iterator

import (
	"errors"
	"fmt"
	"reflect"
)

// Done is returned by an iterator's Next method when the iteration is complete.
var Done = errors.New("no more items in iterator")

// We don't support mixed calls to Next and NextPage because they play
// with the paging state in incompatible ways.
var errMixed = errors.New("iterator: Next and NextPage called on same iterator")

// PageInfo contains information about an iterator's paging state.
type PageInfo struct {
	// Token is the token used to retrieve the next page of items from the
	// API. You may set Token immediately after creating an iterator to
	// begin iteration at a particular point. After the underlying API method
	// is called to retrieve a page of items, Token is set to the next-page
	// token in the response.
	Token string

	// MaxSize is the maximum number of items returned by a call to the API.
	// Set MaxSize as a hint to optimize the buffering behavior of the iterator.
	//
	// Use Pager to retrieve a page of a specific, exact size.
	MaxSize int

	// Function that fetches a page from the underlying service.
	fetch func(pageSize int, pageToken string) (nextPageToken string, err error)

	// Function returning the length of the iterator's buffer.
	bufLen func() int

	// Function that returns the buffer, after setting the buffer variable to nil.
	takeBuf func() interface{}

	// Set to true on first call to PageInfo.next or Pager.NextPage. Used to check
	// for calls to both Next and NextPage with the same iterator.
	nextCalled, nextPageCalled bool
}

// NewPageInfo exposes internals for iterator implementations.
// It is not a stable interface.
var NewPageInfo = newPageInfo

// If an iterator can support paging, its iterator-creating method should call
// this.
//
// The fetch, bufLen and takeBuf arguments provide access to the
// iterator's internal slice of buffered items. They behave as described in
// PageInfo, above.
//
// The return value is the PageInfo.next method bound to the returned PageInfo value.
// (Returning it avoids exporting PageInfo.next.)
func newPageInfo(fetch func(int, string) (string, error), bufLen func() int, takeBuf func() interface{}) (*PageInfo, func(error) (error, bool)) {
	pi := &PageInfo{
		fetch:   fetch,
		bufLen:  bufLen,
		takeBuf: takeBuf,
	}
	return pi, pi.next
}

// Remaining returns the number of items available before the iterator makes another API call.
func (pi *PageInfo) Remaining() int { return pi.bufLen() }

// next provides support for an iterator's Next function.
// It take the iterator's error state on input, and return the new error state on output.
// It also reports whether the buffer has at least one element in it.
// Note that the two return values can diverge -- there can be a non-nil error, and a filled
// buffer -- when Done is returned.
func (pi *PageInfo) next(err error) (error, bool) {
	pi.nextCalled = true
	if pi.nextPageCalled {
		return errMixed, false
	}
	for pi.bufLen() == 0 {
		if err != nil {
			return err, false
		}
		if err := pi.fill(pi.MaxSize); err != nil {
			return err, false
		}
		if pi.Token == "" {
			return Done, true
		}
	}
	return nil, true
}

// Call the API to fill the buffer, using size and pi.Token. Set pi.Token to the
// next-page token returned by the call.
func (pi *PageInfo) fill(size int) error {
	tok, err := pi.fetch(size, pi.Token)
	if err != nil {
		pi.takeBuf() // clear the buffer
		return err
	}
	pi.Token = tok
	return nil
}

// Pageable is implemented by iterators that support paging.
type Pageable interface {
	// PageInfo returns paging information associated with the iterator.
	PageInfo() *PageInfo
}

// Pager supports retrieving iterator items a page at a time.
type Pager struct {
	pageInfo *PageInfo
	pageSize int
	err      error
}

// NewPager returns a pager that uses iter. Calls to its NextPage method will
// obtain exactly pageSize items, unless fewer remain. The pageToken argument
// indicates where to start the iteration. Pass the empty string to start at
// the beginning, or pass a token retrieved from a call to Pager.NextPage.
//
// If you use an iterator with a Pager, you must not call Next on the iterator.
func NewPager(iter Pageable, pageSize int, pageToken string) *Pager {
	p := &Pager{
		pageInfo: iter.PageInfo(),
		pageSize: pageSize,
	}
	p.pageInfo.Token = pageToken
	if pageSize <= 0 {
		p.err = errors.New("iterator: page size must be positive")
	}
	return p
}

// NextPage retrieves a sequence of items from the iterator and appends them
// to slicep, which must be a pointer to a slice of the iterator's item type.
// Exactly p.pageSize items will be appended, unless fewer remain.
//
// The first return value is the page token to use for the next page of items.
// If empty, there are no more pages. Aside from checking for the end of the
// iteration, the returned page token is only needed if the iteration is to be
// resumed a later time, in another context (possibly another process).
//
// The second return value is non-nil if an error occurred. It will never be
// the special iterator sentinel value Done. To recognize the end of the
// iteration, compare nextPageToken to the empty string.
//
// After NextPage returns a non-nil error, all subsequent calls will return the
// same error.
func (p *Pager) NextPage(slicep interface{}) (nextPageToken string, err error) {
	p.pageInfo.nextPageCalled = true
	if p.err != nil {
		return "", p.err
	}
	if p.pageInfo.nextCalled {
		p.err = errMixed
		return "", p.err
	}
	if slicep == nil {
		return "", errors.New("nil passed to Pager.NextPage")
	}
	vslicep := reflect.ValueOf(slicep)
	// The buffer must be empty here, so takeBuf is a no-op. We call it just to get
	// the buffer's type.
	wantSliceType := reflect.PtrTo(reflect.ValueOf(p.pageInfo.takeBuf()).Type())
	if vslicep.Type() != wantSliceType {
		return "", fmt.Errorf("slicep should be of type %s, got %T", wantSliceType, slicep)
	}
	if p.pageInfo.bufLen() > 0 {
		p.err = errors.New("must call NextPage with an empty buffer")
		return "", p.err
	}
	for p.pageInfo.bufLen() < p.pageSize {
		if err := p.pageInfo.fill(p.pageSize - p.pageInfo.bufLen()); err != nil {
			p.err = err
			return "", p.err
		}
		if p.pageInfo.Token == "" {
			break
		}
	}
	e := vslicep.Elem()
	e.Set(reflect.AppendSlice(e, reflect.ValueOf(p.pageInfo.takeBuf())))
	return p.pageInfo.Token, nil
}
