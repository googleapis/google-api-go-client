// Copyright 2021 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package downscope

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/api/option"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/google/downscope"
)

const (
	envObject1            = "OBJECT_NAME_1"
	envBucket1            = "BUCKET_NAME_1"
	envObject2            = "OBJECT_NAME_2"
	envBucket2            = "BUCKET_NAME_2"
	envServiceAccountFile = "GCLOUD_TESTS_GOLANG_KEY"
	rootTokenScope        = "https://www.googleapis.com/auth/cloud-platform"
	projectID             = "dulcet-port-762"
)

var (
	object1        string
	object2        string
	bucket1 = "dulcet-port-762"
	bucket2        string
	rootCredential *google.Credentials
)

// TestMain contains all of the setup code that needs to be run once before any of the tests are run
func TestMain(m *testing.M) {
	flag.Parse()
	if testing.Short() {
		// This line runs all of our individual tests
		os.Exit(m.Run())
	}
	ctx := context.Background()
	object1 = os.Getenv(envObject1)
	if object1 == "" {
		log.Fatalf("failed to find envObject1")
	}
	object2 = os.Getenv(envObject2)
	if object2 == "" {
		log.Fatalf("failed to find envObject2")
	}
	bucket1 = os.Getenv(envBucket1)
	if bucket1 == "" {
		log.Fatalf("failed to find envBucket1")
	}
	bucket2 = os.Getenv(envBucket2)
	if bucket2 == "" {
		log.Fatalf("failed to find envBucket2")
	}
	credentialFileName := os.Getenv(envServiceAccountFile)
	credentialsFileData, err := os.ReadFile(credentialFileName)
	if err != nil {
		log.Fatalf("failed to open credentials file: %v", err)
	}
	rootCredential, err = google.CredentialsFromJSON(ctx, credentialsFileData, rootTokenScope)
	if err != nil {
		log.Fatalf("failed to construct root credential: %v", err)
	}

	// This line runs all of our individual tests
	os.Exit(m.Run())

}

// downscopeTest holds the parameters necessary for running a test of the token downscoping capabilities implemented in `oauth2/google/downscope`
type downscopeTest struct {
	name                 string
	availableResource    string
	availablePermissions []string
	conditions           []downscope.AvailabilityCondition
	objectName           string
	rootSource           oauth2.TokenSource
	expectError          bool
}


func TestDownscopedToken(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	var downscopeTests = []downscopeTest{
		{
			name:                 "successfulDownscopedRead",
			availableResource:    "//storage.googleapis.com/projects/_/buckets/" + bucket1,
			availablePermissions: []string{"inRole:roles/storage.objectViewer"},
			conditions:			  []downscope.AvailabilityCondition{
				{
					Expression:  "resource.name.startsWith('projects/_/buckets/" + bucket1 + "/objects/" + object1 + "')",
				},
			},
			rootSource:           rootCredential.TokenSource,
			objectName:           object1,
			expectError:          false,
		},
		{
			name:                 "readOWithoutPermission",
			availableResource:    "//storage.googleapis.com/projects/_/buckets/" + bucket1,
			availablePermissions: []string{"inRole:roles/storage.objectViewer"},
			conditions:			  []downscope.AvailabilityCondition{
				{
					Expression:  "resource.name.startsWith('projects/_/buckets/" + bucket1 + "/objects/" + object1 + "')",
				},
			},
			rootSource:           rootCredential.TokenSource,
			objectName:           object2,
			expectError:          true,
		},
	}

	for _, tt := range downscopeTests {
		t.Run(tt.name, func(t *testing.T) {
			err := helper(tt)
			// If a test isn't supposed to fail, it shouldn't fail.
			if !tt.expectError && err != nil {
				t.Errorf("test case %v should have succeeded, but instead returned %v", tt.name, err)
			} else if tt.expectError && err == nil { // If a test is supposed to fail, it should return a non-nil error.
				t.Errorf(" test case %v should have returned an error, but instaed returned nil", tt.name)
			}
		})
	}
}

// I'm not sure what I should name this according to convention.
func helper(tt downscopeTest) error {
	ctx := context.Background()

	// Initializes an accessBoundary
	AccessBoundaryRules := make([]downscope.AccessBoundaryRule, 0)
	AccessBoundaryRules = append(AccessBoundaryRules, downscope.AccessBoundaryRule{AvailableResource: tt.availableResource, AvailablePermissions: tt.availablePermissions})

	downscopedTokenSource, err := downscope.NewTokenSource(context.Background(), downscope.DownscopingConfig{RootSource: tt.rootSource, Rules: AccessBoundaryRules})
	if err != nil {
		return fmt.Errorf("failed to create the initial token source: %v", err)
	}
	downscopedTokenSource = oauth2.ReuseTokenSource(nil, downscopedTokenSource)

	storageClient, err := storage.NewClient(ctx, option.WithTokenSource(downscopedTokenSource))
	if err != nil {
		return fmt.Errorf("error creating storage client: %v", err)
	}

	bkt := storageClient.Bucket(bucket1)
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()
	obj := bkt.Object(tt.objectName)
	rc, err := obj.NewReader(ctx)
	if err != nil {
		return fmt.Errorf("failed to construct the reader: %v\n", err)
	}
	defer rc.Close()

	_, err = ioutil.ReadAll(rc)
	if err != nil {
		return fmt.Errorf("ioutil.ReadAll: %v", err)
	}

	return nil
}
