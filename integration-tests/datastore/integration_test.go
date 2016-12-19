// Copyright 2017 Google Inc. All Rights Reserved.
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

// +build integration

package storage

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	datastore "google.golang.org/api/datastore/v1"
)

type object struct {
	name, contents string
}

var (
	projectID string
	ds        *datastore.Service
)

const (
	envProject    = "GCLOUD_TESTS_GOLANG_PROJECT_ID"
	envPrivateKey = "GCLOUD_TESTS_GOLANG_KEY"
)

func TestMain(m *testing.M) {
	ds = createService()
	if ds == nil {
		os.Exit(1)
	}
	if err := cleanup(); err != nil {
		log.Fatalf("Pre-test cleanup failed: %v", err)
	}
	exit := m.Run()
	if err := cleanup(); err != nil {
		log.Fatalf("Post-test cleanup failed: %v", err)
	}
	os.Exit(exit)
}

func createService() *datastore.Service {
	if projectID = os.Getenv(envProject); projectID == "" {
		log.Print("no project ID specified")
		return nil
	}
	ctx := context.Background()
	ts, err := tokenSource(ctx, datastore.DatastoreScope)
	if err != nil {
		log.Printf("tokenSource: %v", err)
		return nil
	}
	client := oauth2.NewClient(ctx, ts)
	s, err := datastore.New(client)
	if err != nil {
		log.Printf("unable to create service: %v", err)
		return nil
	}
	return s
}

// TODO(jba): Move this to a common location, point storage/integration_test.go to it as well.
func tokenSource(ctx context.Context, scopes ...string) (oauth2.TokenSource, error) {
	keyFile := os.Getenv(envPrivateKey)
	if keyFile == "" {
		return nil, errors.New(envPrivateKey + " not set")
	}
	jsonKey, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read %q: %v", keyFile, err)
	}
	conf, err := google.JWTConfigFromJSON(jsonKey, scopes...)
	if err != nil {
		return nil, fmt.Errorf("google.JWTConfigFromJSON: %v", err)
	}
	return conf.TokenSource(ctx), nil
}

func TestSpecialFloats(t *testing.T) {
}
