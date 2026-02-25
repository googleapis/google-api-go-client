// Copyright 2026 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internaloption

import (
	"google.golang.org/api/internal"
	"google.golang.org/api/option"
)

// UnsafeResolver is used to introspect client options, which
// are opaque by their nature as functional options.
type UnsafeResolver struct {
	ds *internal.DialSettings
}

// NewUnsafeResolver provides and experimental mechanism for introspecting
// passed-in ClientOptions.  It is intended for use by internal Google client
// code, and provides no stability contract as it's dependent on internal
// implementation details.
//
// The method is experimental and subject to change without notice.
func NewUnsafeResolver(opts ...option.ClientOption) (*UnsafeResolver, error) {
	ds := new(internal.DialSettings)
	for _, o := range opts {
		o.Apply(ds)
	}
	return &UnsafeResolver{
		ds: ds,
	}, nil
}

// ResolvedGRPCConnPoolSize provides the passed in value correspnding to the
// WithGRPCConnectionPool option in google.golang.org/option.
func (ur *UnsafeResolver) ResolvedGRPCConnPoolSize() int {
	return ur.ds.GRPCConnPoolSize
}

// ResolvedGRPCEndpoint returns the resolved endpoint address used for
// establishing gRPC connections.
func (ur *UnsafeResolver) ResolvedGRPCEndpoint() (string, error) {
	_, addr, err := internal.GetGRPCTransportConfigAndEndpoint(ur.ds)
	return addr, err
}

// ResolvedGRPCConnIsCustom exposes whether the provided options included
// directives for providing a customized transport, corresponding to the
// WithGRPCConn option.
func (ur *UnsafeResolver) ResolvedGRPCConnIsCustom() bool {
	return ur.ds.GRPCConn != nil
}
