// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bytestream

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type ExampleReadHandler struct {
	buf  string
	name string // In this example, the service can handle one name only.
}

func (mr *ExampleReadHandler) GetReader(ctx context.Context, name string) (io.ReaderAt, error) {
	if mr.name == "" {
		mr.name = name
		log.Printf("read from name: %q", name)
	} else if mr.name != name {
		return nil, grpc.Errorf(codes.NotFound, "reader already has name=%q, now a new name=%q confuses me", r.name, name)
	}
	return bytes.NewReader([]byte(mr.buf)), nil
}

// Close can be a no-op.
func (mr *ExampleReadHandler) Close(ctx context.Context, name string) {}

type ExampleWriteHandler struct {
	buf  bytes.Buffer // bytes.Buffer implements io.Writer
	name string       // In this example, the service can handle one name only.
}

// Handle writes to a given name.
func (mw *ExampleWriteHandler) GetWriter(ctx context.Context, name string, initOffset int64) (io.Writer, error) {
	if mw.name == "" {
		mw.name = name
		log.Printf("write to name: %q", name)
	} else if mw.name != name {
		return nil, grpc.Errorf(codes.NotFound, "I already have name=%q, now a new name=%q confuses me", mw.name, name)
	}
	// TODO: initOffset is ignored.
	return mw.buf, nil
}

// Close can be a no-op.
func (mw *ExampleWriteHandler) Close(ctx context.Context, name string) {}

func ExampleNewServer() {
	gsrv := grpc.NewServer()
	bytestreamServer, err := bytestream.NewServer(gsrv, &MyReadHandler{buf: "Hello world", name: "foo"}, &MyWriteHandler{})
	if err != nil {
		log.Printf("bytestream.NewServer: %v", err)
		return
	}

	// Start accepting incoming connections. See gRPC docs.
}
