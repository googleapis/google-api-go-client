// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !go1.5

// TODO(djd): Delete this file once Go 1.6 is out and Go 1.4 support is
// dropped. Setting RawPath is preferable to Opaque since we only want to
// control the encoding of the Path element, not the whole URL.

package googleapi

import "net/url"

// SetOpaque sets u.Opaque from u.Path such that HTTP requests to it
// don't alter any hex-escaped characters in u.Path.
func SetOpaque(u *url.URL) {
	u.Opaque = "//" + u.Host + u.Path
}
