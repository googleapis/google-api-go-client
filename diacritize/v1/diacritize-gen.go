// Package diacritize provides access to the Diacritize API.
//
// See http://code.google.com/apis/language/diacritize/v1/using_rest.html
//
// Usage example:
//
//   import "google-api-go-client.googlecode.com/hg/diacritize/v1"
//   ...
//   diacritizeService, err := diacritize.New(oauthHttpClient)
package diacritize

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

const apiId = "diacritize:v1"
const apiName = "diacritize"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/language/diacritize/"

func New(client *http.Client) (*Service, os.Error) {
	if client == nil {
		return nil, os.NewError("client is nil")
	}
	s := &Service{client: client}
	s.Diacritize = &DiacritizeService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	Diacritize *DiacritizeService
}

type DiacritizeService struct {
	s *Service
}

type LanguageDiacritizeCorpusResource struct {
	Diacritized_text interface{} `json:"diacritized_text,omitempty"`
}

func cleanPathString(s string) string {
	return strings.Map(func(r int) int {
		if r >= 0x30 && r <= 0x7a {
			return r
		}
		return -1
	}, s)
}
