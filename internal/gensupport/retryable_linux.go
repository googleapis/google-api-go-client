package gensupport

import "syscall"

// +build linux

func init() {
	// Initialize syscallRetryable to return true on transient socket-level
	// errors. These errors are specific to Linux.
	syscallRetryable = func(err error) bool { return err == syscall.ECONNRESET || err == syscall.ECONNREFUSED }
}
