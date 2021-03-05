// Copyright 2021 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build integration

// To run this test locally, you will need to do the following:
// • Navigate to your Google Cloud Project
// • Get a copy of the Service Account Key File from somebody
// • If you are unable to obtain an existing key file, create one:
//    • > IAM and Admin > Service Accounts
//    • Under the needed Service Account > Actions > Manage Keys
//    • Add Key > Create New Key
//    • Select JSON, and the click Create
// • > Compute > Compute Engine > VM Instances
// • Look for an available VM Instance, or create one
// • On the VM Instance, click the SSH Button
// • Upload your Service Account Key File
// • Upload this script, along with setup.sh
// • Get a copy of the needed environment variables from somebody, and upload those too
// • Set your environment variables (Usually this will be `source env.conf`)
// • If the setup script has not yet been run, then run it
// • `go test -tags integration`

package byoid

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"google.golang.org/api/dns/v1"
	"google.golang.org/api/idtoken"
	"google.golang.org/api/option"
)

const (
	envCredentials  = "GCLOUD_TESTS_GOLANG_KEY"
	envAudienceOIDC = "GCLOUD_TESTS_GOLANG_AUDIENCE_OIDC"
	envProject      = "GCLOUD_TESTS_GOLANG_PROJECT_ID"
)

var (
	oidcAudience string
	oidcToken    string
	awsToken     string
	clientID     string
	projectID    string
)

// TestMain contains all of the setup code that needs to be run once before any of the tests are run
func TestMain(m *testing.M) {
	keyFileName := os.Getenv(envCredentials)
	if keyFileName == "" {
		log.Fatalf("Please set %s to your keyfile", envCredentials)
	}

	projectID = os.Getenv(envProject)
	if projectID == "" {
		log.Fatalf("Please set %s to the ID of the project", envProject)
	}

	oidcAudience = os.Getenv(envAudienceOIDC)
	if oidcAudience == "" {
		log.Fatalf("Please set %s to the OIDC Audience", envAudienceOIDC)
	}

	var err error

	clientID, err = getClientID(keyFileName)
	if err != nil {
		log.Fatalf("Error getting Client ID: %v", err)
	}

	oidcToken, err = generateGoogleToken(keyFileName)
	if err != nil {
		log.Fatalf("Error generating Google token: %v", err)
	}

	// This line runs all of our individual tests
	os.Exit(m.Run())
}

// keyFile is a struct to extract the relevant json fields for our ServiceAccount KeyFile
type keyFile struct {
	ClientEmail string `json:"client_email"`
	ClientID    string `json:"client_id"`
}

func getClientID(keyFileName string) (string, error) {
	kf, err := os.Open(keyFileName)
	if err != nil {
		return "", err
	}
	defer kf.Close()

	decoder := json.NewDecoder(kf)
	var keyFileSettings keyFile
	err = decoder.Decode(&keyFileSettings)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("projects/-/serviceAccounts/%s", keyFileSettings.ClientEmail), nil
}

func generateGoogleToken(keyFileName string) (string, error) {
	ts, err := idtoken.NewTokenSource(context.Background(), oidcAudience, option.WithCredentialsFile(keyFileName))
	if err != nil {
		return "", nil
	}

	token, err := ts.Token()
	if err != nil {
		return "", nil
	}

	return token.AccessToken, nil
}

// testBYOID makes sure that the default credentials works for
// whatever preconditions have been set beforehand
// by using those credentials to run our client libraries.
//
// In each test we will set up whatever preconditions we need,
// and then use this function.
func testBYOID(t *testing.T, c config) {
	t.Helper()

	// Set up config file.
	configFile, err := ioutil.TempFile("", "config.json")
	if err != nil {
		t.Fatalf("Error creating config file: %v", err)
	}
	defer os.Remove(configFile.Name())

	err = json.NewEncoder(configFile).Encode(c)
	if err != nil {
		t.Errorf("Error writing to config file: %v", err)
	}
	configFile.Close()

	// Once the default credentials are obtained,
	// we should be able to access Google Cloud resources.
	dnsService, err := dns.NewService(context.Background(), option.WithCredentialsFile(configFile.Name()))
	if err != nil {
		t.Fatalf("Could not establish DNS Service: %v", err)
	}

	_, err = dnsService.Projects.Get(projectID).Do()
	if err != nil {
		t.Fatalf("DNS Service failed: %v", err)
	}
}

// These structs makes writing our config as json to a file much easier.
type config struct {
	Type                           string           `json:"type"`
	Audience                       string           `json:"audience"`
	SubjectTokenType               string           `json:"subject_token_type"`
	TokenURL                       string           `json:"token_url"`
	ServiceAccountImpersonationURL string           `json:"service_account_impersonation_url"`
	CredentialSource               credentialSource `json:"credential_source"`
}

type credentialSource struct {
	File string `json:"file,omitempty"`

	URL     string            `json:"url,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`

	EnvironmentID               string `json:"environment_id,omitempty"`
	RegionalCredVerificationURL string `json:"regional_cred_verification_url,omitempty"`
	Format                      string `json:"format,omitempty"`
}

// Tests to make sure File based external credentials continues to work.
func TestFileBasedCredentials(t *testing.T) {
	// Set up Token as a file
	tokenFile, err := ioutil.TempFile("", "token.txt")
	if err != nil {
		t.Fatalf("Error creating token file:")
	}
	defer os.Remove(tokenFile.Name())

	tokenFile.WriteString(oidcToken)
	tokenFile.Close()

	// Run our test!
	testBYOID(t, config{
		Type:                           "external_account",
		Audience:                       oidcAudience,
		SubjectTokenType:               "urn:ietf:params:oauth:token-type:jwt",
		TokenURL:                       "https://sts.googleapis.com/v1beta/token",
		ServiceAccountImpersonationURL: fmt.Sprintf("https://iamcredentials.googleapis.com/v1/%s:generateAccessToken", clientID),
		CredentialSource: credentialSource{
			File: tokenFile.Name(),
		},
	})
}
