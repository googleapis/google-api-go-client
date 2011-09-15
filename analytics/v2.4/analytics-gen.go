// Package analytics provides access to the Google Analytics API.
//
// See http://code.google.com/apis/analytics
//
// Usage example:
//
//   import "google-api-go-client.googlecode.com/hg/analytics/v2.4"
//   ...
//   analyticsService, err := analytics.New(oauthHttpClient)
package analytics

import (
	"bytes"
	"fmt"
	"http"
	"io"
	"json"
	"os"
	"strings"
	"strconv"
	"url"
	"google-api-go-client.googlecode.com/hg/google-api"
)

var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version

const apiId = "analytics:v2.4"
const apiName = "analytics"
const apiVersion = "v2.4"
const basePath = "https://www.googleapis.com/analytics/v2.4/management/"

// OAuth2 scopes used by this API.
const (
	// View your Google Analytics data
	AnalyticsReadonlyScope = "https://www.googleapis.com/auth/analytics.readonly"
)

func New(client *http.Client) (*Service, os.Error) {
	if client == nil {
		return nil, os.NewError("client is nil")
	}
	s := &Service{client: client}
	s.Management = &ManagementService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	Management *ManagementService
}

type ManagementService struct {
	s *Service
}

func cleanPathString(s string) string {
	return strings.Map(func(r int) int {
		if r >= 0x30 && r <= 0x7a {
			return r
		}
		return -1
	}, s)
}
