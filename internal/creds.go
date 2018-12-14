// Copyright 2017 Google LLC
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

package internal

import (
	"context"
	"fmt"
	"io/ioutil"

	"golang.org/x/oauth2/google"
)

// Creds returns credential information obtained from DialSettings, or if none, then
// it returns default credential information.
func Creds(ctx context.Context, ds *DialSettings) (*google.Credentials, error) {
	if ds.Credentials != nil {
		return ds.Credentials, nil
	}
	if ds.CredentialsJSON != nil {
		return credentialsFromJSON(ctx, ds.CredentialsJSON, ds.Endpoint, ds.Scopes, ds.Audiences)
	}
	if ds.CredentialsFile != "" {
		data, err := ioutil.ReadFile(ds.CredentialsFile)
		if err != nil {
			return nil, fmt.Errorf("cannot read credentials file: %v", err)
		}
		return credentialsFromJSON(ctx, data, ds.Endpoint, ds.Scopes, ds.Audiences)
	}
	if ds.TokenSource != nil {
		return &google.Credentials{TokenSource: ds.TokenSource}, nil
	}
	cred, err := google.FindDefaultCredentials(ctx, ds.Scopes...)
	if err != nil {
		return nil, err
	}
	return credentialsFromJSON(
		ctx, cred.JSON, ds.Endpoint, ds.Scopes, ds.Audiences)
}

func credentialsFromJSON(ctx context.Context, data []byte, endpoint string,
	scopes []string, audiences []string) (*google.Credentials, error) {
	cred, err := google.CredentialsFromJSON(ctx, data, scopes...)
	if len(scopes) == 0 {
		// Use JWT token if no scopes provided.
		// Construct the default audience field using the service endpoint.
		// The audience field has the following format: https://<service_name>/
		audience := endpoint
		if len(audiences) > 0 {
			// TODO(shinfan): Update golang oauth to support multiple audiences.
			if len(audiences) > 1 {
				return nil, fmt.Errorf("multiple audiences support is not implemented")
			}
			audience = audiences[0]
		}
		ts, err := google.JWTAccessTokenSourceFromJSON(data, audience)
		if err != nil {
			return nil, err
		}
		cred.TokenSource = ts
	}
	return cred, err
}
