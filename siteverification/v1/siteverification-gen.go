// Package siteverification provides access to the Google Site Verification API.
//
// See http://code.google.com/apis/siteverification/
//
// Usage example:
//
//   import "google-api-go-client.googlecode.com/hg/siteverification/v1"
//   ...
//   siteverificationService, err := siteverification.New(oauthHttpClient)
package siteverification

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

const apiId = "siteVerification:v1"
const apiName = "siteVerification"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/siteVerification/v1/"

// OAuth2 scopes used by this API.
const (
	// Manage the list of sites and domains you control
	SiteverificationScope = "https://www.googleapis.com/auth/siteverification"
)

func New(client *http.Client) (*Service, os.Error) {
	if client == nil {
		return nil, os.NewError("client is nil")
	}
	s := &Service{client: client}
	s.WebResource = &WebResourceService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	WebResource *WebResourceService
}

type WebResourceService struct {
	s *Service
}

type SiteVerificationWebResourceResource struct {
	// Id: The string used to identify this site. This value should be used
	// in the "id" portion of the REST URL for the Get, Update, and Delete
	// operations.
	Id string `json:"id,omitempty"`

	// Site: The address and type of a site that is verified or will be
	// verified.
	Site *SiteVerificationWebResourceResourceSite `json:"site,omitempty"`

	// Owners: The email addresses of all verified owners.
	Owners []string `json:"owners,omitempty"`
}

type SiteVerificationWebResourceGettokenResponse struct {
	// Token: The verification token. The token must be placed appropriately
	// in order for verification to succeed.
	Token string `json:"token,omitempty"`

	// Method: The verification method to use in conjunction with this
	// token. For FILE, the token should be placed in the top-level
	// directory of the site, stored inside a file of the same name. For
	// META, the token should be placed in the HEAD tag of the default page
	// that is loaded for the site. For DNS, the token should be placed in a
	// TXT record of the domain.
	Method string `json:"method,omitempty"`
}

type SiteVerificationWebResourceGettokenRequest struct {
	// VerificationMethod: The verification method that will be used to
	// verify this site. For sites, 'FILE' or 'META' methods may be used.
	// For domains, only 'DNS' may be used.
	VerificationMethod string `json:"verificationMethod,omitempty"`

	// Site: The site for which a verification token will be generated.
	Site *SiteVerificationWebResourceGettokenRequestSite `json:"site,omitempty"`
}

type SiteVerificationWebResourceListResponse struct {
	// Items: The list of sites that are owned by the authenticated user.
	Items []*SiteVerificationWebResourceResource `json:"items,omitempty"`
}

type SiteVerificationWebResourceGettokenRequestSite struct {
	// Identifier: The site identifier. If the type is set to SITE, the
	// identifier is a URL. If the type is set to INET_DOMAIN, the site
	// identifier is a domain name.
	Identifier string `json:"identifier,omitempty"`

	// Type: The type of resource to be verified. Can be SITE or INET_DOMAIN
	// (domain name).
	Type string `json:"type,omitempty"`
}

type SiteVerificationWebResourceResourceSite struct {
	// Identifier: The site identifier. If the type is set to SITE, the
	// identifier is a URL. If the type is set to INET_DOMAIN, the site
	// identifier is a domain name.
	Identifier string `json:"identifier,omitempty"`

	// Type: The site type. Can be SITE or INET_DOMAIN (domain name).
	Type string `json:"type,omitempty"`
}

// method id "siteVerification.webResource.list":

type WebResourceListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Get the list of your verified websites and domains.
func (r *WebResourceService) List() *WebResourceListCall {
	c := &WebResourceListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

func (c *WebResourceListCall) Do() (*SiteVerificationWebResourceListResponse, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/siteVerification/v1/", "webResource")
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
	ret := new(SiteVerificationWebResourceListResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get the list of your verified websites and domains.",
	//   "httpMethod": "GET",
	//   "id": "siteVerification.webResource.list",
	//   "path": "webResource",
	//   "response": {
	//     "$ref": "SiteVerificationWebResourceListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/siteverification"
	//   ]
	// }

}

// method id "siteVerification.webResource.update":

type WebResourceUpdateCall struct {
	s                                   *Service
	id                                  string
	siteverificationwebresourceresource *SiteVerificationWebResourceResource
	opt_                                map[string]interface{}
}

// Update: Modify the list of owners for your website or domain.
func (r *WebResourceService) Update(id string, siteverificationwebresourceresource *SiteVerificationWebResourceResource) *WebResourceUpdateCall {
	c := &WebResourceUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	c.siteverificationwebresourceresource = siteverificationwebresourceresource
	return c
}

func (c *WebResourceUpdateCall) Do() (*SiteVerificationWebResourceResource, os.Error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.siteverificationwebresourceresource)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/siteVerification/v1/", "webResource/{id}")
	urls = strings.Replace(urls, "{id}", cleanPathString(c.id), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(SiteVerificationWebResourceResource)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Modify the list of owners for your website or domain.",
	//   "httpMethod": "PUT",
	//   "id": "siteVerification.webResource.update",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The id of a verified site or domain.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "webResource/{id}",
	//   "request": {
	//     "$ref": "SiteVerificationWebResourceResource"
	//   },
	//   "response": {
	//     "$ref": "SiteVerificationWebResourceResource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/siteverification"
	//   ]
	// }

}

// method id "siteVerification.webResource.insert":

type WebResourceInsertCall struct {
	s                                   *Service
	verificationMethod                  string
	siteverificationwebresourceresource *SiteVerificationWebResourceResource
	opt_                                map[string]interface{}
}

// Insert: Attempt verification of a website or domain.
func (r *WebResourceService) Insert(verificationMethod string, siteverificationwebresourceresource *SiteVerificationWebResourceResource) *WebResourceInsertCall {
	c := &WebResourceInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.verificationMethod = verificationMethod
	c.siteverificationwebresourceresource = siteverificationwebresourceresource
	return c
}

func (c *WebResourceInsertCall) Do() (*SiteVerificationWebResourceResource, os.Error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.siteverificationwebresourceresource)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("verificationMethod", fmt.Sprintf("%v", c.verificationMethod))
	urls := googleapi.ResolveRelative("https://www.googleapis.com/siteVerification/v1/", "webResource")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(SiteVerificationWebResourceResource)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Attempt verification of a website or domain.",
	//   "httpMethod": "POST",
	//   "id": "siteVerification.webResource.insert",
	//   "parameterOrder": [
	//     "verificationMethod"
	//   ],
	//   "parameters": {
	//     "verificationMethod": {
	//       "description": "The method to use for verifying a site or domain.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "webResource",
	//   "request": {
	//     "$ref": "SiteVerificationWebResourceResource"
	//   },
	//   "response": {
	//     "$ref": "SiteVerificationWebResourceResource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/siteverification"
	//   ]
	// }

}

// method id "siteVerification.webResource.get":

type WebResourceGetCall struct {
	s    *Service
	id   string
	opt_ map[string]interface{}
}

// Get: Get the most current data for a website or domain.
func (r *WebResourceService) Get(id string) *WebResourceGetCall {
	c := &WebResourceGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

func (c *WebResourceGetCall) Do() (*SiteVerificationWebResourceResource, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/siteVerification/v1/", "webResource/{id}")
	urls = strings.Replace(urls, "{id}", cleanPathString(c.id), 1)
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
	ret := new(SiteVerificationWebResourceResource)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get the most current data for a website or domain.",
	//   "httpMethod": "GET",
	//   "id": "siteVerification.webResource.get",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The id of a verified site or domain.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "webResource/{id}",
	//   "response": {
	//     "$ref": "SiteVerificationWebResourceResource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/siteverification"
	//   ]
	// }

}

// method id "siteVerification.webResource.delete":

type WebResourceDeleteCall struct {
	s    *Service
	id   string
	opt_ map[string]interface{}
}

// Delete: Relinquish ownership of a website or domain.
func (r *WebResourceService) Delete(id string) *WebResourceDeleteCall {
	c := &WebResourceDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

func (c *WebResourceDeleteCall) Do() os.Error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/siteVerification/v1/", "webResource/{id}")
	urls = strings.Replace(urls, "{id}", cleanPathString(c.id), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Relinquish ownership of a website or domain.",
	//   "httpMethod": "DELETE",
	//   "id": "siteVerification.webResource.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The id of a verified site or domain.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "webResource/{id}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/siteverification"
	//   ]
	// }

}

// method id "siteVerification.webResource.getToken":

type WebResourceGetTokenCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// GetToken: Get a verification token for placing on a website or
// domain.
func (r *WebResourceService) GetToken() *WebResourceGetTokenCall {
	c := &WebResourceGetTokenCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// VerificationMethod sets the optional parameter "verificationMethod":
// The method to use for verifying a site or domain.
func (c *WebResourceGetTokenCall) VerificationMethod(verificationMethod string) *WebResourceGetTokenCall {
	c.opt_["verificationMethod"] = verificationMethod
	return c
}

// Identifier sets the optional parameter "identifier": The URL or
// domain to verify.
func (c *WebResourceGetTokenCall) Identifier(identifier string) *WebResourceGetTokenCall {
	c.opt_["identifier"] = identifier
	return c
}

// Type sets the optional parameter "type": Type of resource to verify.
// Can be 'site' (URL) or 'inet_domain' (domain name).
func (c *WebResourceGetTokenCall) Type(type_ string) *WebResourceGetTokenCall {
	c.opt_["type"] = type_
	return c
}

func (c *WebResourceGetTokenCall) Do() (*SiteVerificationWebResourceGettokenResponse, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["verificationMethod"]; ok {
		params.Set("verificationMethod", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["identifier"]; ok {
		params.Set("identifier", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["type"]; ok {
		params.Set("type", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/siteVerification/v1/", "token")
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
	ret := new(SiteVerificationWebResourceGettokenResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get a verification token for placing on a website or domain.",
	//   "httpMethod": "GET",
	//   "id": "siteVerification.webResource.getToken",
	//   "parameters": {
	//     "identifier": {
	//       "description": "The URL or domain to verify.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "type": {
	//       "description": "Type of resource to verify. Can be 'site' (URL) or 'inet_domain' (domain name).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "verificationMethod": {
	//       "description": "The method to use for verifying a site or domain.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "token",
	//   "response": {
	//     "$ref": "SiteVerificationWebResourceGettokenResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/siteverification"
	//   ]
	// }

}

// method id "siteVerification.webResource.patch":

type WebResourcePatchCall struct {
	s                                   *Service
	id                                  string
	siteverificationwebresourceresource *SiteVerificationWebResourceResource
	opt_                                map[string]interface{}
}

// Patch: Modify the list of owners for your website or domain. This
// method supports patch semantics.
func (r *WebResourceService) Patch(id string, siteverificationwebresourceresource *SiteVerificationWebResourceResource) *WebResourcePatchCall {
	c := &WebResourcePatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	c.siteverificationwebresourceresource = siteverificationwebresourceresource
	return c
}

func (c *WebResourcePatchCall) Do() (*SiteVerificationWebResourceResource, os.Error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.siteverificationwebresourceresource)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/siteVerification/v1/", "webResource/{id}")
	urls = strings.Replace(urls, "{id}", cleanPathString(c.id), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(SiteVerificationWebResourceResource)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Modify the list of owners for your website or domain. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "siteVerification.webResource.patch",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The id of a verified site or domain.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "webResource/{id}",
	//   "request": {
	//     "$ref": "SiteVerificationWebResourceResource"
	//   },
	//   "response": {
	//     "$ref": "SiteVerificationWebResourceResource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/siteverification"
	//   ]
	// }

}

func cleanPathString(s string) string {
	return strings.Map(func(r int) int {
		if r >= 0x30 && r <= 0x7a {
			return r
		}
		return -1
	}, s)
}
