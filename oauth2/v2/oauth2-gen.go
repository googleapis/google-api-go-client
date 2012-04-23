// Package oauth2 provides access to the .
//
// Usage example:
//
//   import "code.google.com/p/google-api-go-client/oauth2/v2"
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

const apiId = "oauth2:v2"
const apiName = "oauth2"
const apiVersion = "v2"
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

type Userinfo struct {
	Id string `json:"id,omitempty"`

	Verified_email bool `json:"verified_email,omitempty"`

	Timezone string `json:"timezone,omitempty"`

	Birthday string `json:"birthday,omitempty"`

	Email string `json:"email,omitempty"`

	Locale string `json:"locale,omitempty"`

	Given_name string `json:"given_name,omitempty"`

	Picture string `json:"picture,omitempty"`

	Name string `json:"name,omitempty"`

	Family_name string `json:"family_name,omitempty"`

	Link string `json:"link,omitempty"`

	Gender string `json:"gender,omitempty"`
}

type Tokeninfo struct {
	// Issued_to: To whom was the token issued to. In general the same as
	// audience.
	Issued_to string `json:"issued_to,omitempty"`

	// Scope: The space separated list of scopes granted to this token.
	Scope string `json:"scope,omitempty"`

	// Audience: Who is the intended audience for this token. In general the
	// same as issued_to.
	Audience string `json:"audience,omitempty"`

	// Expires_in: The expiry time of the token, as number of seconds left
	// until expiry.
	Expires_in int64 `json:"expires_in,omitempty"`

	// User_id: The Gaia obfuscated user id.
	User_id string `json:"user_id,omitempty"`

	// Access_type: The access type granted with this toke. It can be
	// offline or online.
	Access_type string `json:"access_type,omitempty"`

	// Verified_email: Boolean flag which is true if the email address is
	// verified. Present only if the email scope is present in the request.
	Verified_email bool `json:"verified_email,omitempty"`

	// Email: The email address of the user. Present only if the email scope
	// is present in the request.
	Email string `json:"email,omitempty"`
}

type Oauth2IssueTokenV2ResponseConsentScopes struct {
	Description string `json:"description,omitempty"`

	Detail string `json:"detail,omitempty"`
}

type Oauth2IssueTokenV2Response struct {
	IdToken string `json:"idToken,omitempty"`

	Token string `json:"token,omitempty"`

	Consent *Oauth2IssueTokenV2ResponseConsent `json:"consent,omitempty"`

	Code string `json:"code,omitempty"`

	IssueAdvice string `json:"issueAdvice,omitempty"`
}

type Oauth2IssueTokenV2ResponseConsent struct {
	OauthClient *Oauth2IssueTokenV2ResponseConsentOauthClient `json:"oauthClient,omitempty"`

	Scopes []*Oauth2IssueTokenV2ResponseConsentScopes `json:"scopes,omitempty"`
}

type Oauth2IssueTokenV2ResponseConsentOauthClient struct {
	IconUri string `json:"iconUri,omitempty"`

	Name string `json:"name,omitempty"`

	DeveloperEmail string `json:"developerEmail,omitempty"`
}

// method id "oauth2.issueTokenGet":

type IssueTokenGetCall struct {
	s             *Service
	client_id     string
	response_type string
	scope         string
	opt_          map[string]interface{}
}

// IssueTokenGet: 
func (s *Service) IssueTokenGet(client_id string, response_type string, scope string) *IssueTokenGetCall {
	c := &IssueTokenGetCall{s: s, opt_: make(map[string]interface{})}
	c.client_id = client_id
	c.response_type = response_type
	c.scope = scope
	return c
}

// App_id sets the optional parameter "app_id": 
func (c *IssueTokenGetCall) App_id(app_id string) *IssueTokenGetCall {
	c.opt_["app_id"] = app_id
	return c
}

// Android_device_id sets the optional parameter "android_device_id": 
func (c *IssueTokenGetCall) Android_device_id(android_device_id uint64) *IssueTokenGetCall {
	c.opt_["android_device_id"] = android_device_id
	return c
}

// Force sets the optional parameter "force": 
func (c *IssueTokenGetCall) Force(force bool) *IssueTokenGetCall {
	c.opt_["force"] = force
	return c
}

// Origin sets the optional parameter "origin": 
func (c *IssueTokenGetCall) Origin(origin string) *IssueTokenGetCall {
	c.opt_["origin"] = origin
	return c
}

// Hl sets the optional parameter "hl": 
func (c *IssueTokenGetCall) Hl(hl string) *IssueTokenGetCall {
	c.opt_["hl"] = hl
	return c
}

// Audience sets the optional parameter "audience": 
func (c *IssueTokenGetCall) Audience(audience string) *IssueTokenGetCall {
	c.opt_["audience"] = audience
	return c
}

// Alg sets the optional parameter "alg": 
func (c *IssueTokenGetCall) Alg(alg string) *IssueTokenGetCall {
	c.opt_["alg"] = alg
	return c
}

func (c *IssueTokenGetCall) Do() (*Oauth2IssueTokenV2Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("scope", fmt.Sprintf("%v", c.scope))
	params.Set("client_id", fmt.Sprintf("%v", c.client_id))
	params.Set("response_type", fmt.Sprintf("%v", c.response_type))
	if v, ok := c.opt_["app_id"]; ok {
		params.Set("app_id", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["android_device_id"]; ok {
		params.Set("android_device_id", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["force"]; ok {
		params.Set("force", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["origin"]; ok {
		params.Set("origin", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["audience"]; ok {
		params.Set("audience", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alg"]; ok {
		params.Set("alg", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/", "oauth2/v2/IssueToken")
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
	ret := new(Oauth2IssueTokenV2Response)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "httpMethod": "GET",
	//   "id": "oauth2.issueTokenGet",
	//   "parameterOrder": [
	//     "client_id",
	//     "response_type",
	//     "scope"
	//   ],
	//   "parameters": {
	//     "alg": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "android_device_id": {
	//       "format": "uint64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "app_id": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "audience": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "client_id": {
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "force": {
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "hl": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "origin": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "response_type": {
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "scope": {
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "oauth2/v2/IssueToken",
	//   "response": {
	//     "$ref": "Oauth2IssueTokenV2Response"
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

// Id_token sets the optional parameter "id_token": 
func (c *TokeninfoCall) Id_token(id_token string) *TokeninfoCall {
	c.opt_["id_token"] = id_token
	return c
}

// Access_token sets the optional parameter "access_token": 
func (c *TokeninfoCall) Access_token(access_token string) *TokeninfoCall {
	c.opt_["access_token"] = access_token
	return c
}

func (c *TokeninfoCall) Do() (*Tokeninfo, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["id_token"]; ok {
		params.Set("id_token", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["access_token"]; ok {
		params.Set("access_token", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/", "oauth2/v2/tokeninfo")
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

func (c *UserinfoGetCall) Do() (*Userinfo, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/", "oauth2/v2/userinfo")
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
	//   "path": "oauth2/v2/userinfo",
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
