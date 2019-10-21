// Copyright 2016 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"testing"
	"time"

	"google.golang.org/grpc/naming"
)

func TestConnectionPool(t *testing.T) {
	addr := "127.0.0.1:123"
	ds := DialSettings{Endpoint: addr}
	pr := NewPoolResolver(4, &ds)
	watcher, err := pr.Resolve(addr)
	if err != nil {
		t.Fatalf("Resolve: %v", err)
	}

	updates, err := watcher.Next()
	if err != nil {
		t.Fatalf("Next: %v", err)
	}
	if len(updates) != 4 {
		t.Fatalf("Update count: %v", err)
	}
	metaSeen := make(map[interface{}]bool)
	for _, u := range updates {
		if u.Addr != addr {
			t.Errorf("Addr from update: wanted %v, got %v", addr, u.Addr)
		}
		// Metadata must be unique
		if metaSeen[u.Metadata] {
			t.Errorf("Wanted %v to be unique, got %v", u.Metadata, metaSeen)
		}
		metaSeen[u.Metadata] = true
	}

	// Test that Next blocks until Close and returns nil.
	nextc := make(chan []*naming.Update)
	go func() {
		next, err := watcher.Next()
		if err != nil {
			t.Errorf("Next: expected success, got %v", err)
		}
		nextc <- next
	}()
	time.Sleep(50 * time.Millisecond) // wait for watcher.Next goroutine
	select {
	case <-nextc:
		t.Fatal("next should not have been called yet")
	default:
	}
	watcher.Close()
	select {
	case next := <-nextc:
		if next != nil {
			t.Errorf("Next: expected nil, got %v", next)
		}
	case <-time.After(50 * time.Millisecond):
		t.Error("Next: did not return after 100ms")
	}
}
