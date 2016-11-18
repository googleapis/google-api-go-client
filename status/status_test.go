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

package status

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/protobuf/ptypes/any"
	pb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func TestFromError(t *testing.T) {
	for _, test := range []struct {
		err     error
		code    codes.Code
		message string
	}{
		// An actual gRPC error turns into a Status with the same code and message.
		{grpc.Errorf(codes.NotFound, "not found"), codes.NotFound, "not found"},
		// A nil error becomes a Status with code OK and an empty message.
		{nil, codes.OK, ""},
		// A non-gRPC error has an unknown code, and the message is the result of calling Error().
		{errors.New("something"), codes.Unknown, "something"},
	} {
		want := &pb.Status{Code: int32(test.code), Message: test.message}
		if got := FromError(test.err); !reflect.DeepEqual(got, want) {
			t.Errorf("%#v:\n got %+v\nwant %+v", test.err, got, want)
		}
	}
}

func TestToGRPCError(t *testing.T) {
	for _, test := range []struct {
		status *pb.Status
		want   error
	}{
		// A status without details maps directly to a gRPC error.
		{&pb.Status{Code: int32(codes.NotFound), Message: "not found"}, grpc.Errorf(codes.NotFound, "not found")},
		// A status with details loses the details, but amends the message.
		{&pb.Status{Code: int32(codes.NotFound), Message: "not found", Details: []*any.Any{{}}},
			grpc.Errorf(codes.NotFound, "not found (details omitted)")},
		// A status with codes.OK maps to nil, regardless of message or details.
		{&pb.Status{Code: int32(codes.OK), Message: "whatever", Details: []*any.Any{{}}}, nil},
		// A nil status maps to nil.
		{nil, nil},
	} {
		got := ToGRPCError(test.status)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%+v:\n got %#v\nwant %#v", test.status, got, test.want)
		}
	}
}
