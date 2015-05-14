// Package oauth2 provides access to the Google OAuth2 API.
//
// See https://developers.google.com/accounts/docs/OAuth2
//
// Usage example:
//
//   import "google.golang.org/api/oauth2/v2"
//   ...
//   oauth2Service, err := oauth2.New(oauthHttpClient)
package oauth2

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

const apiId = "oauth2:v2"
const apiName = "oauth2"
const apiVersion = "v2"
const basePath = "https://www.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// Know your basic profile info and list of people in your circles.
	PlusLoginScope = "https://www.googleapis.com/auth/plus.login"

	// Know who you are on Google
	PlusMeScope = "https://www.googleapis.com/auth/plus.me"

	// View your email address
	UserinfoEmailScope = "https://www.googleapis.com/auth/userinfo.email"

	// View your basic profile info
	UserinfoProfileScope = "https://www.googleapis.com/auth/userinfo.profile"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Userinfo = NewUserinfoService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Userinfo *UserinfoService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewUserinfoService(s *Service) *UserinfoService {
	rs := &UserinfoService{s: s}
	rs.V2 = NewUserinfoV2Service(s)
	return rs
}

type UserinfoService struct {
	s *Service

	V2 *UserinfoV2Service
}

func NewUserinfoV2Service(s *Service) *UserinfoV2Service {
	rs := &UserinfoV2Service{s: s}
	rs.Me = NewUserinfoV2MeService(s)
	return rs
}

type UserinfoV2Service struct {
	s *Service

	Me *UserinfoV2MeService
}

func NewUserinfoV2MeService(s *Service) *UserinfoV2MeService {
	rs := &UserinfoV2MeService{s: s}
	return rs
}

type UserinfoV2MeService struct {
	s *Service
}

type Jwk struct {
	Keys []*JwkKeys `json:"keys,omitempty"`
}

type JwkKeys struct {
	Alg string `json:"alg,omitempty"`

	E string `json:"e,omitempty"`

	Kid string `json:"kid,omitempty"`

	Kty string `json:"kty,omitempty"`

	N string `json:"n,omitempty"`

	Use string `json:"use,omitempty"`
}

type Tokeninfo struct {
	// AccessType: The access type granted with this token. It can be
	// offline or online.
	AccessType string `json:"access_type,omitempty"`

	// Audience: Who is the intended audience for this token. In general the
	// same as issued_to.
	Audience string `json:"audience,omitempty"`

	// Email: The email address of the user. Present only if the email scope
	// is present in the request.
	Email string `json:"email,omitempty"`

	// ExpiresIn: The expiry time of the token, as number of seconds left
	// until expiry.
	ExpiresIn int64 `json:"expires_in,omitempty"`

	// IssuedTo: To whom was the token issued to. In general the same as
	// audience.
	IssuedTo string `json:"issued_to,omitempty"`

	// Scope: The space separated list of scopes granted to this token.
	Scope string `json:"scope,omitempty"`

	// TokenHandle: The token handle associated with this token.
	TokenHandle string `json:"token_handle,omitempty"`

	// UserId: The obfuscated user id.
	UserId string `json:"user_id,omitempty"`

	// VerifiedEmail: Boolean flag which is true if the email address is
	// verified. Present only if the email scope is present in the request.
	VerifiedEmail bool `json:"verified_email,omitempty"`
}

type Userinfoplus struct {
	// Email: The user's email address.
	Email string `json:"email,omitempty"`

	// FamilyName: The user's last name.
	FamilyName string `json:"family_name,omitempty"`

	// Gender: The user's gender.
	Gender string `json:"gender,omitempty"`

	// GivenName: The user's first name.
	GivenName string `json:"given_name,omitempty"`

	// Hd: The hosted domain e.g. example.com if the user is Google apps
	// user.
	Hd string `json:"hd,omitempty"`

	// Id: The obfuscated ID of the user.
	Id string `json:"id,omitempty"`

	// Link: URL of the profile page.
	Link string `json:"link,omitempty"`

	// Locale: The user's preferred locale.
	Locale string `json:"locale,omitempty"`

	// Name: The user's full name.
	Name string `json:"name,omitempty"`

	// Picture: URL of the user's picture image.
	Picture string `json:"picture,omitempty"`

	// VerifiedEmail: Boolean flag which is true if the email address is
	// verified. Always verified because we only return the user's primary
	// email address.
	//
	// Default: true
	VerifiedEmail *bool `json:"verified_email,omitempty"`
}

// method id "oauth2.getCertForOpenIdConnect":

type GetCertForOpenIdConnectCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// GetCertForOpenIdConnect:
func (s *Service) GetCertForOpenIdConnect() *GetCertForOpenIdConnectCall {
	c := &GetCertForOpenIdConnectCall{s: s, opt_: make(map[string]interface{})}
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GetCertForOpenIdConnectCall) Fields(s ...googleapi.Field) *GetCertForOpenIdConnectCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *GetCertForOpenIdConnectCall) Do() (*Jwk, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "oauth2/v2/certs")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Jwk
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "httpMethod": "GET",
	//   "id": "oauth2.getCertForOpenIdConnect",
	//   "path": "oauth2/v2/certs",
	//   "response": {
	//     "$ref": "Jwk"
	//   }
	// }

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

// AccessToken sets the optional parameter "access_token":
func (c *TokeninfoCall) AccessToken(accessToken string) *TokeninfoCall {
	c.opt_["access_token"] = accessToken
	return c
}

// IdToken sets the optional parameter "id_token":
func (c *TokeninfoCall) IdToken(idToken string) *TokeninfoCall {
	c.opt_["id_token"] = idToken
	return c
}

// TokenHandle sets the optional parameter "token_handle":
func (c *TokeninfoCall) TokenHandle(tokenHandle string) *TokeninfoCall {
	c.opt_["token_handle"] = tokenHandle
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TokeninfoCall) Fields(s ...googleapi.Field) *TokeninfoCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
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
	if v, ok := c.opt_["token_handle"]; ok {
		params.Set("token_handle", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "oauth2/v2/tokeninfo")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Tokeninfo
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
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
	//     },
	//     "token_handle": {
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "oauth2/v2/tokeninfo",
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

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UserinfoGetCall) Fields(s ...googleapi.Field) *UserinfoGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *UserinfoGetCall) Do() (*Userinfoplus, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "oauth2/v2/userinfo")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Userinfoplus
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "httpMethod": "GET",
	//   "id": "oauth2.userinfo.get",
	//   "path": "oauth2/v2/userinfo",
	//   "response": {
	//     "$ref": "Userinfoplus"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/plus.login",
	//     "https://www.googleapis.com/auth/plus.me",
	//     "https://www.googleapis.com/auth/userinfo.email",
	//     "https://www.googleapis.com/auth/userinfo.profile"
	//   ]
	// }

}

// method id "oauth2.userinfo.v2.me.get":

type UserinfoV2MeGetCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// Get:
func (r *UserinfoV2MeService) Get() *UserinfoV2MeGetCall {
	c := &UserinfoV2MeGetCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UserinfoV2MeGetCall) Fields(s ...googleapi.Field) *UserinfoV2MeGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *UserinfoV2MeGetCall) Do() (*Userinfoplus, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "userinfo/v2/me")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Userinfoplus
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "httpMethod": "GET",
	//   "id": "oauth2.userinfo.v2.me.get",
	//   "path": "userinfo/v2/me",
	//   "response": {
	//     "$ref": "Userinfoplus"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/plus.login",
	//     "https://www.googleapis.com/auth/plus.me",
	//     "https://www.googleapis.com/auth/userinfo.email",
	//     "https://www.googleapis.com/auth/userinfo.profile"
	//   ]
	// }

}
