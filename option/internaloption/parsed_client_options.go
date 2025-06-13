// Copyright 2025 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Package internaloption contains options used internally by Google client code.
package internaloption

import (
	"google.golang.org/api/internal"
	"google.golang.org/api/option"
)

type ParsedOptions struct {
	internal.DialSettings
}

// ParseClientOptions validates the given option.ClientOption slice and returns
// ParsedOptions with the resolved settings. It returns an error if the
// provided options are invalid.
//
// This function allows other Google Cloud client libraries to read configuration
// values set by users via ClientOptions, which are otherwise unreadable outside of
// google.golang.org/api.
func ParseClientOptions(opts []option.ClientOption) (*ParsedOptions, error) {
	var ds internal.DialSettings
	// Apply all options to the internal DialSettings struct.
	for _, opt := range opts {
		opt.Apply(&ds)
	}

	// Validate the combined settings.
	if err := ds.Validate(); err != nil {
		return nil, err
	}

	return &ParsedOptions{
		DialSettings: ds,
	}, nil
}
