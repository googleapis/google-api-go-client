// Copyright 2020 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package grpc

import (
	"testing"

	"google.golang.org/grpc"
)

func TestPool(t *testing.T) {
	conn1 := &grpc.ClientConn{}
	conn2 := &grpc.ClientConn{}

	pool := &roundRobinConnPool{
		conns: []*grpc.ClientConn{
			conn1, conn2,
		},
	}

	if got := pool.Conn(); got != conn2 {
		t.Errorf("pool.Conn() #1 got %v; want conn2 (%v)", got, conn1)
	}

	if got := pool.Conn(); got != conn1 {
		t.Errorf("pool.Conn() #2 got %v; want conn1 (%v)", got, conn1)
	}

	if got := pool.Conn(); got != conn2 {
		t.Errorf("pool.Conn() #3 got %v; want conn2 (%v)", got, conn1)
	}
}
