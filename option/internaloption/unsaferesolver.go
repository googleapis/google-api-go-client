// Copyright 2026 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internaloption

import (
	"errors"

	"google.golang.org/api/internal"
	"google.golang.org/api/option"
)

type unsafeResolver struct {
	ds *internal.DialSettings
}

// NewUnsafeResolver provides and experimental mechanism for introspecting
// passed-in ClientOptions.  It is intended for use by internal Google client
// code, and provides no stability contract as it's dependent on internal
// implementation details.
//
// The method is experimental and subject to change without notice.
func NewUnsafeResolver(opts ...option.ClientOption) (*unsafeResolver, error) {
	ds := new(internal.DialSettings)
	for _, o := range opts {
		o.Apply(ds)
	}
	return &unsafeResolver{
		ds: ds,
	}, nil
}

// ResolvedGRPCConnPoolSize provides the passed in value correspnding to the
// WithGRPCConnectionPool option in google.golang.org/option.
func (ur *unsafeResolver) ResolvedGRPCConnPoolSize() int {
	return ur.ds.GRPCConnPoolSize
}

// ResolvedGRPCEndpoint returns the resolved endpoint address used for
// establishing gRPC connections.
func (ur *unsafeResolver) ResolvedGRPCEndpoint() (string, error) {
	_, addr, err := internal.GetGRPCTransportConfigAndEndpoint(ur.ds)
	return addr, err
}

// ResolvedGRPCConnIsCustom exposes whether the provided options included
// directives for providing a customized transport, corresponding to the
// WithGRPCConn option.
func (ur *unsafeResolver) ResolvedGRPCConnIsCustom() bool {
	return ur.ds.GRPCConn != nil
}

// ProbeDirectPathConfiguration attempts to identify if the passed in options
// are expected to yield result in a client that can communicate over direct
// path.  It will return an error with more details about possible causes that
// prevent direct path from being used, or nil if no issues were identified.
func (ur *unsafeResolver) ProbeDirectPathConfiguration() error {
	return errors.New("unimplemented")
}
