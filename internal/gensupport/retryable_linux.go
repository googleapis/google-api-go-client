// Copyright 2020 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gensupport

import "syscall"

// +build linux

func init() {
	// Initialize syscallRetryable to return true on transient socket-level
	// errors. These errors are specific to Linux.
	syscallRetryable = func(err error) bool { return err == syscall.ECONNRESET || err == syscall.ECONNREFUSED }
}
