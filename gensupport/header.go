// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gensupport

import (
	"fmt"
	"runtime"
	"strings"

	"google.golang.org/api/googleapi"
)

// GoogleClientHeader returns the value to use for the x-goog-api-client
// header, which is used internally by Google.
//
// TODO(jba): Use a repo-wide version, instead of one specific to the generator
// or this package.
func GoogleClientHeader(ignore, clientElement string) string {
	elts := []string{"gl-go/" + strings.Replace(runtime.Version(), " ", "_", -1)}
	if clientElement != "" {
		elts = append(elts, clientElement)
	}
	elts = append(elts, fmt.Sprintf("gdcl/%s", googleapi.Version))
	return strings.Join(elts, " ")
}
