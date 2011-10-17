// Package adexchangebuyer provides access to the Ad Exchange Buyer API.
//
// See http://code.google.com/apis/adexchangebuyer/docs
//
// Usage example:
//
//   import "google-api-go-client.googlecode.com/hg/adexchangebuyer/v1"
//   ...
//   adexchangebuyerService, err := adexchangebuyer.New(oauthHttpClient)
package adexchangebuyer

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

const apiId = "adexchangebuyer:v1"
const apiName = "adexchangebuyer"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/adexchangebuyer/v1/"

func New(client *http.Client) (*Service, os.Error) {
	if client == nil {
		return nil, os.NewError("client is nil")
	}
	s := &Service{client: client}
	s.Accounts = &AccountsService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	Accounts *AccountsService
}

type AccountsService struct {
	s *Service
}

type AccountBidderLocation struct {
	// Url: The URL to which the Ad Exchange will send bid requests.
	Url string `json:"url,omitempty"`

	// MaximumQps: The maximum queries per second the Ad Exchange will send.
	MaximumQps int64 `json:"maximumQps,omitempty"`
}

type AccountsList struct {
	// Items: A list of accounts.
	Items []*Account `json:"items,omitempty"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`
}

type Account struct {
	// BidderLocation: Your bidder locations that have distinct URLs.
	BidderLocation []*AccountBidderLocation `json:"bidderLocation,omitempty"`

	// CookieMatchingNid: The nid parameter value used in cookie match
	// requests. This value cannot by changed through an update operation.
	// Please contact your technical account manager if you need to change
	// it.
	CookieMatchingNid string `json:"cookieMatchingNid,omitempty"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`

	// Id: Account id.
	Id int64 `json:"id,omitempty,string"`

	// CookieMatchingUrl: The base url used in cookie match requests.
	CookieMatchingUrl string `json:"cookieMatchingUrl,omitempty"`

	// MaximumTotalQps: The sum of all bidderLocation.maximumQps values
	// cannot exceed this. This value cannot by changed through an update
	// operation. Please contact your technical account manager if you need
	// to change it.
	MaximumTotalQps int64 `json:"maximumTotalQps,omitempty"`
}

// method id "adexchangebuyer.accounts.list":

type AccountsListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Retrieves a list of accounts.
func (r *AccountsService) List() *AccountsListCall {
	c := &AccountsListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

func (c *AccountsListCall) Do() (*AccountsList, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/adexchangebuyer/v1/", "accounts")
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
	ret := new(AccountsList)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of accounts.",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer.accounts.list",
	//   "path": "accounts",
	//   "response": {
	//     "$ref": "AccountsList"
	//   }
	// }

}

// method id "adexchangebuyer.accounts.update":

type AccountsUpdateCall struct {
	s       *Service
	id      string
	account *Account
	opt_    map[string]interface{}
}

// Update: Updates an existing account.
func (r *AccountsService) Update(id string, account *Account) *AccountsUpdateCall {
	c := &AccountsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	c.account = account
	return c
}

func (c *AccountsUpdateCall) Do() (*Account, os.Error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.account)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/adexchangebuyer/v1/", "accounts/{id}")
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
	ret := new(Account)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an existing account.",
	//   "httpMethod": "PUT",
	//   "id": "adexchangebuyer.accounts.update",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The account id",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{id}",
	//   "request": {
	//     "$ref": "Account"
	//   },
	//   "response": {
	//     "$ref": "Account"
	//   }
	// }

}

// method id "adexchangebuyer.accounts.get":

type AccountsGetCall struct {
	s    *Service
	id   string
	opt_ map[string]interface{}
}

// Get: Gets one account by id.
func (r *AccountsService) Get(id string) *AccountsGetCall {
	c := &AccountsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

func (c *AccountsGetCall) Do() (*Account, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/adexchangebuyer/v1/", "accounts/{id}")
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
	ret := new(Account)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets one account by id.",
	//   "httpMethod": "GET",
	//   "id": "adexchangebuyer.accounts.get",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The account id",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{id}",
	//   "response": {
	//     "$ref": "Account"
	//   }
	// }

}

// method id "adexchangebuyer.accounts.patch":

type AccountsPatchCall struct {
	s       *Service
	id      string
	account *Account
	opt_    map[string]interface{}
}

// Patch: Updates an existing account. This method supports patch
// semantics.
func (r *AccountsService) Patch(id string, account *Account) *AccountsPatchCall {
	c := &AccountsPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	c.account = account
	return c
}

func (c *AccountsPatchCall) Do() (*Account, os.Error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.account)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/adexchangebuyer/v1/", "accounts/{id}")
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
	ret := new(Account)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an existing account. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "adexchangebuyer.accounts.patch",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The account id",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{id}",
	//   "request": {
	//     "$ref": "Account"
	//   },
	//   "response": {
	//     "$ref": "Account"
	//   }
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
