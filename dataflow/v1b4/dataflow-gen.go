// Package dataflow provides access to the Google Dataflow API.
//
// Usage example:
//
//   import "google.golang.org/api/dataflow/v1b4"
//   ...
//   dataflowService, err := dataflow.New(oauthHttpClient)
package dataflow

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Background

const apiId = "dataflow:v1b4"
const apiName = "dataflow"
const apiVersion = "v1b4"
const basePath = "https://www.googleapis.com/"

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	return s, nil
}

type Service struct {
	client   *http.Client
	BasePath string // API endpoint base URL
}
