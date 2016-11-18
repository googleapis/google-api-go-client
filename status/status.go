// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package status converts between gRPC errors and google.rpc.Status messages.
package status

import (
	statuspb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// FromError converts the given error to a google.rpc.Status.
// If the error is a gRPC error, the Code and Message fields of the returned Status will match those of the error.
// If the error is not a gRPC error, the Code field will have the value of codes.Unknown, and the Message
// field will be set to the value of err.Error().
// If the error is nil, the Code fiels will have the value of codes.OK, and the Message field will be empty.
// In all cases, the Details field of the returned Status will be empty.
func FromError(err error) *statuspb.Status {
	code := grpc.Code(err)
	desc := grpc.ErrorDesc(err)
	return &statuspb.Status{Code: int32(code), Message: desc}
}

// ToGRPCError converts a google.rpc.Status to an error.
// If the given Status is nil, or its Code is codes.OK (0), then ToGRPCError returns nil.
// Otherwise, ToGRPCError returns a gRPC error whose code and message match those of the Status.
// If the Status's Details field is non-empty, the string " (details omitted)" is appended to the
// message, but otherwise the details are dropped.
func ToGRPCError(s *statuspb.Status) error {
	if s == nil {
		return nil
	}
	format := "%s"
	if len(s.Details) > 0 {
		format += " (details omitted)"
	}
	return grpc.Errorf(codes.Code(s.Code), format, s.Message)
}
