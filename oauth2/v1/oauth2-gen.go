// Package oauth2 provides access to the .
//
// Usage example:
//
//   import "code.google.com/p/google-api-go-client/oauth2/v1"
//   ...
//   oauth2Service, err := oauth2.New(oauthHttpClient)
package oauth2

import (
	"bytes"
	"fmt"
	"net/http"
	"io"
	"encoding/json"
	"errors"
	"strings"
	"strconv"
	"net/url"
	"code.google.com/p/google-api-go-client/googleapi"
)

var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New

const apiId = "oauth2:v1"
const apiName = "oauth2"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/"

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Userinfo = &UserinfoService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	Userinfo *UserinfoService
}

type UserinfoService struct {
	s *Service
}

type Tokeninfo struct {
	// Access_type: The access type granted with this toke. It can be
	// offline or online.
	Access_type string `json:"access_type,omitempty"`

	// Audience: Who is the intended audience for this token. In general the
	// same as issued_to.
	Audience string `json:"audience,omitempty"`

	// Email: The email address of the user. Present only if the email scope
	// is present in the request.
	Email string `json:"email,omitempty"`

	// Expires_in: The expiry time of the token, as number of seconds left
	// until expiry.
	Expires_in int64 `json:"expires_in,omitempty"`

	// Issued_to: To whom was the token issued to. In general the same as
	// audience.
	Issued_to string `json:"issued_to,omitempty"`

	// Scope: The space separated list of scopes granted to this token.
	Scope string `json:"scope,omitempty"`

	// User_id: The Gaia obfuscated user id.
	User_id string `json:"user_id,omitempty"`

	// Verified_email: Boolean flag which is true if the email address is
	// verified. Present only if the email scope is present in the request.
	Verified_email bool `json:"verified_email,omitempty"`
}

type Userinfo struct {
	Birthday string `json:"birthday,omitempty"`

	Email string `json:"email,omitempty"`

	Family_name string `json:"family_name,omitempty"`

	Gender string `json:"gender,omitempty"`

	Given_name string `json:"given_name,omitempty"`

	Id string `json:"id,omitempty"`

	Link string `json:"link,omitempty"`

	Locale string `json:"locale,omitempty"`

	Name string `json:"name,omitempty"`

	Picture string `json:"picture,omitempty"`

	Timezone string `json:"timezone,omitempty"`

	Verified_email bool `json:"verified_email,omitempty"`
}

// method id "oauth2.tokeninfo":

type TokeninfoCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// Tokeninfo: 
func (s *Service) Tokeninfo() *TokeninfoCall {
	c := &TokeninfoCall{s: s, opt_: make(map[string]interface{})}
	return c
}

// Access_token sets the optional parameter "access_token": 
func (c *TokeninfoCall) Access_token(access_token string) *TokeninfoCall {
	c.opt_["access_token"] = access_token
	return c
}

// Id_token sets the optional parameter "id_token": 
func (c *TokeninfoCall) Id_token(id_token string) *TokeninfoCall {
	c.opt_["id_token"] = id_token
	return c
}

func (c *TokeninfoCall) Do() (*Tokeninfo, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["access_token"]; ok {
		params.Set("access_token", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["id_token"]; ok {
		params.Set("id_token", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/", "oauth2/v1/tokeninfo")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Tokeninfo)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "httpMethod": "POST",
	//   "id": "oauth2.tokeninfo",
	//   "parameters": {
	//     "access_token": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "id_token": {
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "oauth2/v1/tokeninfo",
	//   "response": {
	//     "$ref": "Tokeninfo"
	//   }
	// }

}

// method id "oauth2.userinfo.get":

type UserinfoGetCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// Get: 
func (r *UserinfoService) Get() *UserinfoGetCall {
	c := &UserinfoGetCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

func (c *UserinfoGetCall) Do() (*Userinfo, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/", "oauth2/v1/userinfo")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Userinfo)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "httpMethod": "GET",
	//   "id": "oauth2.userinfo.get",
	//   "path": "oauth2/v1/userinfo",
	//   "response": {
	//     "$ref": "Userinfo"
	//   }
	// }

}

func cleanPathString(s string) string {
	return strings.Map(func(r rune) rune {
		if r >= 0x30 && r <= 0x7a {
			return r
		}
		return -1
	}, s)
}
