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
	"testing"
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
	for _, u := range updates {
		if u.Addr != addr {
			t.Errorf("Addr from update, wanted %v, got %v", addr, u.Addr)
		}
	}
	watcher.Close()
}
