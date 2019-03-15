// Copyright 2019 Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package api is the root of the packages used to access Google Cloud
// Services. See https://godoc.org/google.golang.org/api for a full list of
// sub-packages.
//
// Within api there exist numerous clients which connect to Google APIs,
// and various utility packages.
//
//
// Client Options
//
// All clients in sub-packages are configurable via client options. These
// options are described here: https://godoc.org/google.golang.org/api/option.
//
//
// Authentication and Authorization
//
// All the clients in sub-packages support authentication via Google
// Application Default Credentials (see
// https://cloud.google.com/docs/authentication/production), or by providing a
// JSON key file for a Service Account. See the authentication examples in
// https://godoc.org/google.golang.org/api/transport for more details.
//
//
// Versioning and Stability
//
// Clients in these subpackages follow semver as specified in
// https://cloud.google.com/apis/design/versioning.
//
// Note that versioning and stability is strictly not communicated through Go
// modules. Go modules are used only for dependency management. See the link
// above for stability guarantees.
package api
