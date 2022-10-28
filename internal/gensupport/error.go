// Copyright 2020 Google LLC. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gensupport

import (
	"github.com/googleapis/gax-go/v2/apierror"
	"google.golang.org/api/googleapi"
)

// WrapError creates an apierror.APIError from err (one that
// does not wrap err), wraps it in err, and returns err. If
// err is not a googleapi.Error (or a gRPC Status), it returns
// err without modification.
func WrapError(err *googleapi.Error) *googleapi.Error {
	if apiError, ok := apierror.ParseError(err, false); ok {
		err.Err = apiError
	}
	return err
}
