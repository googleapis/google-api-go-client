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

// Package bundler supports bundling (batching) of items. Bundling amortizes an
// action with fixed costs over multiple items. For example, if an API provides
// an RPC that accepts a list of items as input, but clients would prefer
// adding items one at a time, then a Bundler can accept individual items from
// the client and bundle many of them into a single RPC.
//
// This package is experimental and subject to change without notice.
package bundler

import (
	"errors"
	"reflect"
	"sync"
	"time"

	"golang.org/x/net/context"
)

const (
	DefaultDelayThreshold       = time.Second
	DefaultBundleCountThreshold = 10
	DefaultBundleByteThreshold  = 1e6 // 1M
	DefaultBufferedByteLimit    = 1e9 // 1G
)

var (
	// ErrOverflow indicates that Bundler's stored bytes exceeds its BufferedByteLimit.
	ErrOverflow = errors.New("bundler reached buffered byte limit")

	// ErrOversizedItem indicates that an item's size exceeds the maximum bundle size.
	ErrOversizedItem = errors.New("item size exceeds bundle byte limit")
)

// A Bundler collects items added to it into a bundle until the bundle
// exceeds a given size, then calls a user-provided function to handle the bundle.
type Bundler struct {
	// Starting from the time that the first message is added to a bundle, once
	// this delay has passed, handle the bundle. The default is DefaultDelayThreshold.
	DelayThreshold time.Duration

	// Once a bundle has this many items, handle the bundle. Since only one
	// item at a time is added to a bundle, no bundle will exceed this
	// threshold, so it also serves as a limit. The default is
	// DefaultBundleCountThreshold.
	BundleCountThreshold int

	// Once the number of bytes in current bundle reaches this threshold, handle
	// the bundle. The default is DefaultBundleByteThreshold. This triggers handling,
	// but does not cap the total size of a bundle.
	BundleByteThreshold int

	// The maximum size of a bundle, in bytes. Zero means unlimited.
	BundleByteLimit int

	// The maximum number of bytes that the Bundler will keep in memory before
	// returning ErrOverflow. The default is DefaultBufferedByteLimit.
	BufferedByteLimit int

	handler       func(interface{}) // called to handle a bundle
	itemSliceZero reflect.Value     // nil (zero value) for slice of items
	flushTimer    *time.Timer       // implements DelayThreshold

	mu             sync.Mutex
	spaceAvailable chan struct{}   // closed and replaced when space is available
	bufferedSize   int             // total bytes buffered
	curBundle      bundle          // incoming items added to this bundle
	handlingc      <-chan struct{} // set to non-nil while a handler is running; closed when it returns
}

type bundle struct {
	items reflect.Value // slice of item type
	size  int           // size in bytes of all items
}

// NewBundler creates a new Bundler.
//
// itemExample is a value of the type that will be bundled. For example, if you
// want to create bundles of *Entry, you could pass &Entry{} for itemExample.
//
// handler is a function that will be called on each bundle. If itemExample is
// of type T, the argument to handler is of type []T. handler is always called
// sequentially for each bundle, and never in parallel.
func NewBundler(itemExample interface{}, handler func(interface{})) *Bundler {
	b := &Bundler{
		DelayThreshold:       DefaultDelayThreshold,
		BundleCountThreshold: DefaultBundleCountThreshold,
		BundleByteThreshold:  DefaultBundleByteThreshold,
		BufferedByteLimit:    DefaultBufferedByteLimit,

		handler:       handler,
		itemSliceZero: reflect.Zero(reflect.SliceOf(reflect.TypeOf(itemExample))),
	}
	b.curBundle.items = b.itemSliceZero
	return b
}

// Add adds item to the current bundle. It marks the bundle for handling and
// starts a new one if any of the thresholds or limits are exceeded.
//
// If the item's size exceeds the maximum bundle size (Bundler.BundleByteLimit), then
// the item can never be handled. Add returns ErrOversizedItem in this case.
//
// If adding the item would exceed the maximum memory allowed (Bundler.BufferedByteLimit),
// Add returns ErrOverflow.
//
// Add never blocks.
func (b *Bundler) Add(item interface{}, size int) error {
	// If this item exceeds the maximum size of a bundle,
	// we can never send it.
	if b.BundleByteLimit > 0 && size > b.BundleByteLimit {
		return ErrOversizedItem
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	// If adding this item would exceed our allotted memory
	// footprint, we can't accept it.
	if b.bufferedSize+size > b.BufferedByteLimit {
		return ErrOverflow
	}
	b.addLocked(item, size)
	return nil
}

// addLocked adds item to the current bundle. It marks the bundle for handling and
// starts a new one if any of the thresholds or limits are exceeded.
//
// addLocked is called with the lock held.
func (b *Bundler) addLocked(item interface{}, size int) {
	// If adding this item to the current bundle would cause it to exceed the
	// maximum bundle size, close the current bundle and start a new one.
	if b.BundleByteLimit > 0 && b.curBundle.size+size > b.BundleByteLimit {
		b.startFlushLocked()
	}
	// Add the item.
	b.curBundle.items = reflect.Append(b.curBundle.items, reflect.ValueOf(item))
	b.curBundle.size += size
	b.bufferedSize += size

	// Start a timer to flush the item if one isn't already running.
	// startFlushLocked clears the timer and closes the bundle at the same time,
	// so we only allocate a new timer for the first item in each bundle.
	// (We could try to call Reset on the timer instead, but that would add a lot
	// of complexity to the code just to save one small allocation.)
	if b.flushTimer == nil {
		b.flushTimer = time.AfterFunc(b.DelayThreshold, b.Flush)
	}

	// If the current bundle equals the count threshold, close it.
	if b.curBundle.items.Len() == b.BundleCountThreshold {
		b.startFlushLocked()
	}
	// If the current bundle equals or exceeds the byte threshold, close it.
	if b.curBundle.size >= b.BundleByteThreshold {
		b.startFlushLocked()
	}
}

// AddWait adds item to the current bundle. It marks the bundle for handling and
// starts a new one if any of the thresholds or limits are exceeded.
//
// If the item's size exceeds the maximum bundle size (Bundler.BundleByteLimit), then
// the item can never be handled. AddWait returns ErrOversizedItem in this case.
//
// If adding the item would exceed the maximum memory allowed (Bundler.BufferedByteLimit),
// AddWait blocks until space is available or ctx is done.
func (b *Bundler) AddWait(ctx context.Context, item interface{}, size int) error {
	// If this item exceeds the maximum size of a bundle,
	// we can never send it.
	if b.BundleByteLimit > 0 && size > b.BundleByteLimit {
		return ErrOversizedItem
	}
	b.mu.Lock()
	// If adding this item would exceed our allotted memory
	// footprint, block until space is available.
	// TODO(jba): avoid starvation of large items.
	for b.bufferedSize+size > b.BufferedByteLimit {
		if b.spaceAvailable == nil {
			b.spaceAvailable = make(chan struct{})
		}
		avail := b.spaceAvailable
		b.mu.Unlock()
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-avail:
			b.mu.Lock()
		}
	}
	b.addLocked(item, size)
	b.mu.Unlock()
	return nil
}

// Flush invokes the handler for all remaining items in the Bundler and waits
// for it to return.
func (b *Bundler) Flush() {
	b.mu.Lock()
	b.startFlushLocked()
	done := b.handlingc
	b.mu.Unlock()

	if done != nil {
		<-done
	}
}

func (b *Bundler) startFlushLocked() {
	if b.flushTimer != nil {
		b.flushTimer.Stop()
		b.flushTimer = nil
	}

	if b.curBundle.items.Len() == 0 {
		return
	}
	bun := b.curBundle
	b.curBundle = bundle{items: b.itemSliceZero}

	done := make(chan struct{})
	var running <-chan struct{}
	running, b.handlingc = b.handlingc, done

	go func() {
		defer func() {
			b.mu.Lock()
			b.bufferedSize -= bun.size
			avail := b.spaceAvailable
			b.spaceAvailable = nil
			b.mu.Unlock()

			if avail != nil {
				close(avail)
			}
			close(done)
		}()

		if running != nil {
			// Wait for our turn to call the handler.
			<-running
		}

		b.handler(bun.items.Interface())
	}()
}
