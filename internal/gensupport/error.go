// Copyright 2022 Google LLC. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gensupport

import (
	"errors"

	"github.com/googleapis/gax-go/v2/apierror"
	"google.golang.org/api/googleapi"
)

// WrapError creates an [apierror.APIError] from err (one that does not wrap
// err), wraps it in err, and returns err. If err is not a [googleapi.Error]
// (or a gRPC [Status]), it returns err without modification.
//
// [Status]: https://pkg.go.dev/google.golang.org/grpc/status#Status
func WrapError(err error) error {
	if apiError, ok := apierror.ParseError(err, false); ok {
		var herr *googleapi.Error
		if errors.As(err, &herr) {
			herr.Err = apiError
		}
	}
	return err
}
