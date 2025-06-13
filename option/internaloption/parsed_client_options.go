// Copyright 2025 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Package internaloption contains options used internally by Google client code.
package internaloption

import (
	"reflect"

	"google.golang.org/api/internal"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

// ParseClientOptions validates the given ClientOptions and updates the provided
// receiver with the resolved settings. It returns an error if the provided options
// are invalid.
//
// This function allows other Google Cloud client libraries to read configuration
// values set by users via ClientOptions, which are otherwise unreadable outside of
// google.golang.org/api.
func ParseClientOptions(receiver any, opts []option.ClientOption) error {
	var ds internal.DialSettings
	// Apply all options to the internal DialSettings struct.
	for _, opt := range opts {
		opt.Apply(&ds)
	}

	// Validate the combined settings.
	if err := ds.Validate(); err != nil {
		return err
	}

	// Populate the consumer with values from the internal DialSettings.
	applyOption(receiver, "Endpoint", ds.Endpoint)
	applyOption(receiver, "UserAgent", ds.UserAgent)
	applyOption(receiver, "APIKey", ds.APIKey)
	applyOption(receiver, "CredentialsFile", ds.CredentialsFile)
	applyOption(receiver, "TokenSource", ds.TokenSource)
	applyOption(receiver, "Credentials", ds.Credentials)
	applyOption(receiver, "HTTPClient", ds.HTTPClient)
	applyOption(receiver, "GRPCConn", ds.GRPCConn)
	applyOption(receiver, "GRPCConnPoolSize", ds.GRPCConnPoolSize)
	applyOption(receiver, "NoAuth", ds.NoAuth)
	applyOption(receiver, "TelemetryDisabled", ds.TelemetryDisabled)
	applyOption(receiver, "ClientCertSource", ds.ClientCertSource)
	applyOption(receiver, "UniverseDomain", ds.GetUniverseDomain()) // Uses the getter for correct precedence.
	applyOption(receiver, "Logger", ds.Logger)
	applyOption(receiver, "QuotaProject", ds.QuotaProject)
	applyOption(receiver, "RequestReason", ds.RequestReason)
	applyOption(receiver, "AuthCredentials", ds.AuthCredentials)

	if len(ds.CredentialsJSON) > 0 {
		credsJSONCopy := make([]byte, len(ds.CredentialsJSON))
		copy(credsJSONCopy, ds.CredentialsJSON)
		applyOption(receiver, "CredentialsJSON", credsJSONCopy)
	}

	// ds.GetScopes() returns the effective scopes (user-provided or default).
	resolvedScopes := ds.GetScopes()
	if len(resolvedScopes) > 0 {
		scopesCopy := make([]string, len(resolvedScopes))
		copy(scopesCopy, resolvedScopes)
		applyOption(receiver, "Scopes", scopesCopy)
	}

	if len(ds.Audiences) > 0 { // ds.Audiences is already a copy from withAudiences.ApplyClientOption
		audiencesCopy := make([]string, len(ds.Audiences))
		copy(audiencesCopy, ds.Audiences)
		applyOption(receiver, "Audiences", audiencesCopy)
	}

	if len(ds.GRPCDialOpts) > 0 {
		grpcDialOptsCopy := make([]grpc.DialOption, len(ds.GRPCDialOpts))
		copy(grpcDialOptsCopy, ds.GRPCDialOpts)
		applyOption(receiver, "GRPCDialOpts", grpcDialOptsCopy)
	}

	return nil
}

// ApplyOption accesses the field identified by key on consumer and sets
// it to value using reflection. If the field does not exist or the value
// is not assignable, it's a no-op.
func applyOption(receiver any, key string, value any) {
	v := reflect.ValueOf(receiver).Elem()
	field := v.FieldByName(key)

	if !field.IsValid() || !field.CanSet() {
		return // Field doesn't exist or cannot be set
	}

	val := reflect.ValueOf(value)
	// Handle nil interface gracefully for pointer/interface/slice/map types
	if value == nil {
		if field.Kind() == reflect.Ptr || field.Kind() == reflect.Interface || field.Kind() == reflect.Slice || field.Kind() == reflect.Map || field.Kind() == reflect.Func {
			field.Set(reflect.Zero(field.Type())) // Set to nil/zero value for the type
			return
		}
		// For non-pointer/interface/slice/map types, if value is nil but field is not, this would be an error.
		// However, we silently skip as per requirement.
		return
	}

	if !val.Type().AssignableTo(field.Type()) {
		return // Type mismatch
	}
	field.Set(val)
}
