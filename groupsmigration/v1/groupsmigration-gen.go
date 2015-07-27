// Package groupsmigration provides access to the Groups Migration API.
//
// See https://developers.google.com/google-apps/groups-migration/
//
// Usage example:
//
//   import "google.golang.org/api/groupsmigration/v1"
//   ...
//   groupsmigrationService, err := groupsmigration.New(oauthHttpClient)
package groupsmigration

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

const apiId = "groupsmigration:v1"
const apiName = "groupsmigration"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/groups/v1/groups/"

// OAuth2 scopes used by this API.
const (
	// Manage messages in groups on your domain
	AppsGroupsMigrationScope = "https://www.googleapis.com/auth/apps.groups.migration"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Archive = NewArchiveService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Archive *ArchiveService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewArchiveService(s *Service) *ArchiveService {
	rs := &ArchiveService{s: s}
	return rs
}

type ArchiveService struct {
	s *Service
}

// Groups: JSON response template for groups migration API.
type Groups struct {
	// Kind: The kind of insert resource this is.
	Kind string `json:"kind,omitempty"`

	// ResponseCode: The status of the insert request.
	ResponseCode string `json:"responseCode,omitempty"`
}

// method id "groupsmigration.archive.insert":

type ArchiveInsertCall struct {
	s          *Service
	groupId    string
	opt_       map[string]interface{}
	media_     io.Reader
	resumable_ googleapi.SizeReaderAt
	mediaType_ string
	ctx_       context.Context
	protocol_  string
}

// Insert: Inserts a new mail into the archive of the Google group.
func (r *ArchiveService) Insert(groupId string) *ArchiveInsertCall {
	c := &ArchiveInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.groupId = groupId
	return c
}

// Media specifies the media to upload in a single chunk.
// At most one of Media and ResumableMedia may be set.
func (c *ArchiveInsertCall) Media(r io.Reader) *ArchiveInsertCall {
	c.media_ = r
	c.protocol_ = "multipart"
	return c
}

// ResumableMedia specifies the media to upload in chunks and can be cancelled with ctx.
// At most one of Media and ResumableMedia may be set.
// mediaType identifies the MIME media type of the upload, such as "image/png".
// If mediaType is "", it will be auto-detected.
func (c *ArchiveInsertCall) ResumableMedia(ctx context.Context, r io.ReaderAt, size int64, mediaType string) *ArchiveInsertCall {
	c.ctx_ = ctx
	c.resumable_ = io.NewSectionReader(r, 0, size)
	c.mediaType_ = mediaType
	c.protocol_ = "resumable"
	return c
}

// ProgressUpdater provides a callback function that will be called after every chunk.
// It should be a low-latency function in order to not slow down the upload operation.
// This should only be called when using ResumableMedia (as opposed to Media).
func (c *ArchiveInsertCall) ProgressUpdater(pu googleapi.ProgressUpdater) *ArchiveInsertCall {
	c.opt_["progressUpdater"] = pu
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ArchiveInsertCall) Fields(s ...googleapi.Field) *ArchiveInsertCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ArchiveInsertCall) Do() (*Groups, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{groupId}/archive")
	var progressUpdater_ googleapi.ProgressUpdater
	if v, ok := c.opt_["progressUpdater"]; ok {
		if pu, ok := v.(googleapi.ProgressUpdater); ok {
			progressUpdater_ = pu
		}
	}
	if c.media_ != nil || c.resumable_ != nil {
		urls = strings.Replace(urls, "https://www.googleapis.com/", "https://www.googleapis.com/upload/", 1)
		params.Set("uploadType", c.protocol_)
	}
	urls += "?" + params.Encode()
	body = new(bytes.Buffer)
	ctype := "application/json"
	if c.protocol_ != "resumable" {
		var cancel func()
		cancel, _ = googleapi.ConditionallyIncludeMedia(c.media_, &body, &ctype)
		if cancel != nil {
			defer cancel()
		}
	}
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"groupId": c.groupId,
	})
	if c.protocol_ == "resumable" {
		if c.mediaType_ == "" {
			c.mediaType_ = googleapi.DetectMediaType(c.resumable_)
		}
		req.Header.Set("X-Upload-Content-Type", c.mediaType_)
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	} else {
		req.Header.Set("Content-Type", ctype)
	}
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	if c.protocol_ == "resumable" {
		loc := res.Header.Get("Location")
		rx := &googleapi.ResumableUpload{
			Client:        c.s.client,
			UserAgent:     c.s.userAgent(),
			URI:           loc,
			Media:         c.resumable_,
			MediaType:     c.mediaType_,
			ContentLength: c.resumable_.Size(),
			Callback:      progressUpdater_,
		}
		res, err = rx.Upload(c.ctx_)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
	}
	var ret *Groups
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Inserts a new mail into the archive of the Google group.",
	//   "httpMethod": "POST",
	//   "id": "groupsmigration.archive.insert",
	//   "mediaUpload": {
	//     "accept": [
	//       "message/rfc822"
	//     ],
	//     "maxSize": "16MB",
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/groups/v1/groups/{groupId}/archive"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/groups/v1/groups/{groupId}/archive"
	//       }
	//     }
	//   },
	//   "parameterOrder": [
	//     "groupId"
	//   ],
	//   "parameters": {
	//     "groupId": {
	//       "description": "The group ID",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{groupId}/archive",
	//   "response": {
	//     "$ref": "Groups"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.groups.migration"
	//   ],
	//   "supportsMediaUpload": true
	// }

}
