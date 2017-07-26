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

package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
)

// Creds returns credential information obtained from DialSettings, or if none, then
// it returns default credential information.
func Creds(ctx context.Context, ds *DialSettings) (*google.DefaultCredentials, error) {
	if ds.ServiceAccountJSONFilename != "" {
		return serviceAcctCreds(ctx, ds.ServiceAccountJSONFilename, ds.Scopes...)
	}
	if ds.TokenSource != nil {
		return &google.DefaultCredentials{TokenSource: ds.TokenSource}, nil
	}
	return google.FindDefaultCredentials(ctx, ds.Scopes...)
}

// serviceAcctTokenSource reads a JWT config from filename and returns
// a TokenSource constructed from the config.
func serviceAcctCreds(ctx context.Context, filename string, scope ...string) (*google.DefaultCredentials, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("cannot read service account file: %v", err)
	}
	cfg, err := google.JWTConfigFromJSON(data, scope...)
	if err != nil {
		return nil, fmt.Errorf("google.JWTConfigFromJSON: %v", err)
	}
	// jwt.Config does not expose the project ID, so re-unmarshal to get it.
	var pid struct {
		ProjectID string `json:"project_id"`
	}
	if err := json.Unmarshal(data, &pid); err != nil {
		return nil, err
	}
	return &google.DefaultCredentials{
		ProjectID:   pid.ProjectID,
		TokenSource: cfg.TokenSource(ctx),
		// TODO(jba): uncomment after https://go-review.googlesource.com/c/51111 is in.
		// JSON: data,
	}, nil
}
