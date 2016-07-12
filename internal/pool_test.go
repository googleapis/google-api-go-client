/*
Copyright 2016 Google Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package internal

import (
	"sync"
	"testing"
	"time"
)

func TestConnectionPool(t *testing.T) {
	addr := "127.0.0.1:123"
	ds := DialSettings{Endpoint: addr}
	pr := NewPoolResolver(4, &ds)
	watcher, err := pr.Resolve(addr)
	if err != nil {
		t.Fatalf("Resolve(): %v", err)
	}

	updates, err := watcher.Next()
	if err != nil {
		t.Fatalf("Next(): %v", err)
	}
	if len(updates) != 4 {
		t.Fatalf("Update count: %v", err)
	}
	metaSet := make(map[interface{}]bool)
	for _, u := range updates {
		if u.Addr != addr {
			t.Errorf("Addr from update, wanted %v, got %v", addr, u.Addr)
		}
		// Metadata must be unique
		_, found := metaSet[u.Metadata]
		metaSet[u.Metadata] = true
		if found {
			t.Errorf("wanted %v to be unique, got %v", u.Metadata, metaSet)
		}
	}
	// Test that Next now blocks until Close and returns nil.
	var wg sync.WaitGroup
	closed := false
	wg.Add(1)
	go func() {
		defer wg.Done()
		updates, err := watcher.Next()
		if !closed {
			t.Errorf("Next(): second invocation didn't block")
		}
		if updates != nil {
			t.Errorf("Next(): expected nil, got %v", updates)
		}
		if err != nil {
			t.Errorf("Next(): expected no error, got %v", err)
		}
	}()

	time.Sleep(100 * time.Millisecond)
	watcher.Close()
	closed = true

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Error("Close() has not returned after 100ms")
	}
}
