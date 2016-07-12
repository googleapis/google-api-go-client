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

package internal

import (
	"fmt"
	"google.golang.org/grpc/naming"
)

// PoolResolver provides a fixed list of addresses to load balance between
// and does not provide further updates.
type PoolResolver struct {
	poolSize int
	dialOpt  *DialSettings
	addrs    []*naming.Update
	ch       chan struct{}
}

// NewPoolResolver returns a PoolResolver
func NewPoolResolver(size int, o *DialSettings) *PoolResolver {
	return &PoolResolver{poolSize: size, dialOpt: o}
}

// Return the resolver, which also satisfies Watcher.
func (r *PoolResolver) Resolve(target string) (naming.Watcher, error) {
	if r.dialOpt.Endpoint == "" {
		return nil, fmt.Errorf("No endpoint configured")
	}
	for i := 0; i < r.poolSize; i++ {
		r.addrs = append(r.addrs, &naming.Update{Op: naming.Add, Addr: r.dialOpt.Endpoint, Metadata: i})
	}
	r.ch = make(chan struct{})
	return r, nil
}

// Return the static list immediately and then block.
func (r *PoolResolver) Next() ([]*naming.Update, error) {
	if r.addrs != nil {
		addrs := r.addrs
		r.addrs = nil
		return addrs, nil
	}

	<-r.ch
	return nil, nil
}

func (r *PoolResolver) Close() {
	close(r.ch)
}
