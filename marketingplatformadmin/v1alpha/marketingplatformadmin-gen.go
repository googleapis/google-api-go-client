// Copyright 2025 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package marketingplatformadmin provides access to the Google Marketing Platform Admin API.
//
// For product documentation, see: https://developers.google.com/analytics/devguides/config/gmp/v1
//
// # Library status
//
// These client libraries are officially supported by Google. However, this
// library is considered complete and is in maintenance mode. This means
// that we will address critical bugs and security issues but will not add
// any new features.
//
// When possible, we recommend using our newer
// [Cloud Client Libraries for Go](https://pkg.go.dev/cloud.google.com/go)
// that are still actively being worked and iterated on.
//
// # Creating a client
//
// Usage example:
//
//	import "google.golang.org/api/marketingplatformadmin/v1alpha"
//	...
//	ctx := context.Background()
//	marketingplatformadminService, err := marketingplatformadmin.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for
// authentication. For information on how to create and obtain Application
// Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// # Other authentication options
//
// By default, all available scopes (see "Constants") are used to authenticate.
// To restrict scopes, use [google.golang.org/api/option.WithScopes]:
//
//	marketingplatformadminService, err := marketingplatformadmin.NewService(ctx, option.WithScopes(marketingplatformadmin.MarketingplatformadminAnalyticsUpdateScope))
//
// To use an API key for authentication (note: some APIs do not support API
// keys), use [google.golang.org/api/option.WithAPIKey]:
//
//	marketingplatformadminService, err := marketingplatformadmin.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth
// flow, use [google.golang.org/api/option.WithTokenSource]:
//
//	config := &oauth2.Config{...}
//	// ...
//	token, err := config.Exchange(ctx, ...)
//	marketingplatformadminService, err := marketingplatformadmin.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See [google.golang.org/api/option.ClientOption] for details on options.
package marketingplatformadmin // import "google.golang.org/api/marketingplatformadmin/v1alpha"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/googleapis/gax-go/v2/internallog"
	googleapi "google.golang.org/api/googleapi"
	internal "google.golang.org/api/internal"
	gensupport "google.golang.org/api/internal/gensupport"
	option "google.golang.org/api/option"
	internaloption "google.golang.org/api/option/internaloption"
	htransport "google.golang.org/api/transport/http"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = internaloption.WithDefaultEndpoint
var _ = internal.Version
var _ = internallog.New

const apiId = "marketingplatformadmin:v1alpha"
const apiName = "marketingplatformadmin"
const apiVersion = "v1alpha"
const basePath = "https://marketingplatformadmin.googleapis.com/"
const basePathTemplate = "https://marketingplatformadmin.UNIVERSE_DOMAIN/"
const mtlsBasePath = "https://marketingplatformadmin.mtls.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View your Google Analytics product account data in GMP home
	MarketingplatformadminAnalyticsReadScope = "https://www.googleapis.com/auth/marketingplatformadmin.analytics.read"

	// Manage your Google Analytics product account data in GMP home
	MarketingplatformadminAnalyticsUpdateScope = "https://www.googleapis.com/auth/marketingplatformadmin.analytics.update"
)

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	scopesOption := internaloption.WithDefaultScopes(
		"https://www.googleapis.com/auth/marketingplatformadmin.analytics.read",
		"https://www.googleapis.com/auth/marketingplatformadmin.analytics.update",
	)
	// NOTE: prepend, so we don't override user-specified scopes.
	opts = append([]option.ClientOption{scopesOption}, opts...)
	opts = append(opts, internaloption.WithDefaultEndpoint(basePath))
	opts = append(opts, internaloption.WithDefaultEndpointTemplate(basePathTemplate))
	opts = append(opts, internaloption.WithDefaultMTLSEndpoint(mtlsBasePath))
	opts = append(opts, internaloption.EnableNewAuthLibrary())
	client, endpoint, err := htransport.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	s := &Service{client: client, BasePath: basePath, logger: internaloption.GetLogger(opts)}
	s.Organizations = NewOrganizationsService(s)
	if err != nil {
		return nil, err
	}
	if endpoint != "" {
		s.BasePath = endpoint
	}
	return s, nil
}

// New creates a new Service. It uses the provided http.Client for requests.
//
// Deprecated: please use NewService instead.
// To provide a custom HTTP client, use option.WithHTTPClient.
// If you are using google.golang.org/api/googleapis/transport.APIKey, use option.WithAPIKey with NewService instead.
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	return NewService(context.TODO(), option.WithHTTPClient(client))
}

type Service struct {
	client    *http.Client
	logger    *slog.Logger
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Organizations *OrganizationsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewOrganizationsService(s *Service) *OrganizationsService {
	rs := &OrganizationsService{s: s}
	rs.AnalyticsAccountLinks = NewOrganizationsAnalyticsAccountLinksService(s)
	return rs
}

type OrganizationsService struct {
	s *Service

	AnalyticsAccountLinks *OrganizationsAnalyticsAccountLinksService
}

func NewOrganizationsAnalyticsAccountLinksService(s *Service) *OrganizationsAnalyticsAccountLinksService {
	rs := &OrganizationsAnalyticsAccountLinksService{s: s}
	return rs
}

type OrganizationsAnalyticsAccountLinksService struct {
	s *Service
}

// AnalyticsAccountLink: A resource message representing the link between a
// Google Analytics account and a Google Marketing Platform organization.
type AnalyticsAccountLink struct {
	// AnalyticsAccount: Required. Immutable. The resource name of the
	// AnalyticsAdmin API account. The account ID will be used as the ID of this
	// AnalyticsAccountLink resource, which will become the final component of the
	// resource name. Format: analyticsadmin.googleapis.com/accounts/{account_id}
	AnalyticsAccount string `json:"analyticsAccount,omitempty"`
	// DisplayName: Output only. The human-readable name for the Analytics account.
	DisplayName string `json:"displayName,omitempty"`
	// LinkVerificationState: Output only. The verification state of the link
	// between the Analytics account and the parent organization.
	//
	// Possible values:
	//   "LINK_VERIFICATION_STATE_UNSPECIFIED" - The link state is unknown.
	//   "LINK_VERIFICATION_STATE_VERIFIED" - The link is established.
	//   "LINK_VERIFICATION_STATE_NOT_VERIFIED" - The link is requested, but hasn't
	// been approved by the product account admin.
	LinkVerificationState string `json:"linkVerificationState,omitempty"`
	// Name: Identifier. Resource name of this AnalyticsAccountLink. Note the
	// resource ID is the same as the ID of the Analtyics account. Format:
	// organizations/{org_id}/analyticsAccountLinks/{analytics_account_link_id}
	// Example: "organizations/xyz/analyticsAccountLinks/1234"
	Name string `json:"name,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the server.
	googleapi.ServerResponse `json:"-"`
	// ForceSendFields is a list of field names (e.g. "AnalyticsAccount") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "AnalyticsAccount") to include in
	// API requests with the JSON null value. By default, fields with empty values
	// are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s AnalyticsAccountLink) MarshalJSON() ([]byte, error) {
	type NoMethod AnalyticsAccountLink
	return gensupport.MarshalJSON(NoMethod(s), s.ForceSendFields, s.NullFields)
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated empty messages in your APIs. A typical example is to use it as
// the request or the response type of an API method. For instance: service Foo
// { rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty); }
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the server.
	googleapi.ServerResponse `json:"-"`
}

// ListAnalyticsAccountLinksResponse: Response message for
// ListAnalyticsAccountLinks RPC.
type ListAnalyticsAccountLinksResponse struct {
	// AnalyticsAccountLinks: Analytics account links in this organization.
	AnalyticsAccountLinks []*AnalyticsAccountLink `json:"analyticsAccountLinks,omitempty"`
	// NextPageToken: A token, which can be sent as `page_token` to retrieve the
	// next page. If this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the server.
	googleapi.ServerResponse `json:"-"`
	// ForceSendFields is a list of field names (e.g. "AnalyticsAccountLinks") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "AnalyticsAccountLinks") to
	// include in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s ListAnalyticsAccountLinksResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListAnalyticsAccountLinksResponse
	return gensupport.MarshalJSON(NoMethod(s), s.ForceSendFields, s.NullFields)
}

// Organization: A resource message representing a Google Marketing Platform
// organization.
type Organization struct {
	// DisplayName: The human-readable name for the organization.
	DisplayName string `json:"displayName,omitempty"`
	// Name: Identifier. The resource name of the GMP organization. Format:
	// organizations/{org_id}
	Name string `json:"name,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the server.
	googleapi.ServerResponse `json:"-"`
	// ForceSendFields is a list of field names (e.g. "DisplayName") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "DisplayName") to include in API
	// requests with the JSON null value. By default, fields with empty values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s Organization) MarshalJSON() ([]byte, error) {
	type NoMethod Organization
	return gensupport.MarshalJSON(NoMethod(s), s.ForceSendFields, s.NullFields)
}

// SetPropertyServiceLevelRequest: Request message for SetPropertyServiceLevel
// RPC.
type SetPropertyServiceLevelRequest struct {
	// AnalyticsProperty: Required. The Analytics property to change the
	// ServiceLevel setting. This field is the name of the Google Analytics Admin
	// API property resource. Format:
	// analyticsadmin.googleapis.com/properties/{property_id}
	AnalyticsProperty string `json:"analyticsProperty,omitempty"`
	// ServiceLevel: Required. The service level to set for this property.
	//
	// Possible values:
	//   "ANALYTICS_SERVICE_LEVEL_UNSPECIFIED" - Service level unspecified.
	//   "ANALYTICS_SERVICE_LEVEL_STANDARD" - The standard version of Google
	// Analytics.
	//   "ANALYTICS_SERVICE_LEVEL_360" - The premium version of Google Analytics.
	ServiceLevel string `json:"serviceLevel,omitempty"`
	// ForceSendFields is a list of field names (e.g. "AnalyticsProperty") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "AnalyticsProperty") to include in
	// API requests with the JSON null value. By default, fields with empty values
	// are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s SetPropertyServiceLevelRequest) MarshalJSON() ([]byte, error) {
	type NoMethod SetPropertyServiceLevelRequest
	return gensupport.MarshalJSON(NoMethod(s), s.ForceSendFields, s.NullFields)
}

// SetPropertyServiceLevelResponse: Response message for
// SetPropertyServiceLevel RPC.
type SetPropertyServiceLevelResponse struct {
	// ServerResponse contains the HTTP response code and headers from the server.
	googleapi.ServerResponse `json:"-"`
}

type OrganizationsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Lookup for a single organization.
//
//   - name: The name of the Organization to retrieve. Format:
//     organizations/{org_id}.
func (r *OrganizationsService) Get(name string) *OrganizationsGetCall {
	c := &OrganizationsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more
// details.
func (c *OrganizationsGetCall) Fields(s ...googleapi.Field) *OrganizationsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets an optional parameter which makes the operation fail if the
// object's ETag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *OrganizationsGetCall) IfNoneMatch(entityTag string) *OrganizationsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method.
func (c *OrganizationsGetCall) Context(ctx context.Context) *OrganizationsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns a http.Header that can be modified by the caller to add
// headers to the request.
func (c *OrganizationsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OrganizationsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := gensupport.SetHeaders(c.s.userAgent(), "", c.header_)
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, nil)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	c.s.logger.DebugContext(c.ctx_, "api request", "serviceName", apiName, "rpcName", "marketingplatformadmin.organizations.get", "request", internallog.HTTPRequest(req, nil))
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "marketingplatformadmin.organizations.get" call.
// Any non-2xx status code is an error. Response headers are in either
// *Organization.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was returned.
func (c *OrganizationsGetCall) Do(opts ...googleapi.CallOption) (*Organization, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, gensupport.WrapError(&googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		})
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, gensupport.WrapError(err)
	}
	ret := &Organization{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	b, err := gensupport.DecodeResponseBytes(target, res)
	if err != nil {
		return nil, err
	}
	c.s.logger.DebugContext(c.ctx_, "api response", "serviceName", apiName, "rpcName", "marketingplatformadmin.organizations.get", "response", internallog.HTTPResponse(res, b))
	return ret, nil
}

type OrganizationsAnalyticsAccountLinksCreateCall struct {
	s                    *Service
	parent               string
	analyticsaccountlink *AnalyticsAccountLink
	urlParams_           gensupport.URLParams
	ctx_                 context.Context
	header_              http.Header
}

// Create: Creates the link between the Analytics account and the Google
// Marketing Platform organization. User needs to be an org user, and admin on
// the Analytics account to create the link. If the account is already linked
// to an organization, user needs to unlink the account from the current
// organization, then try link again.
//
//   - parent: The parent resource where this Analytics account link will be
//     created. Format: organizations/{org_id}.
func (r *OrganizationsAnalyticsAccountLinksService) Create(parent string, analyticsaccountlink *AnalyticsAccountLink) *OrganizationsAnalyticsAccountLinksCreateCall {
	c := &OrganizationsAnalyticsAccountLinksCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.analyticsaccountlink = analyticsaccountlink
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more
// details.
func (c *OrganizationsAnalyticsAccountLinksCreateCall) Fields(s ...googleapi.Field) *OrganizationsAnalyticsAccountLinksCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method.
func (c *OrganizationsAnalyticsAccountLinksCreateCall) Context(ctx context.Context) *OrganizationsAnalyticsAccountLinksCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns a http.Header that can be modified by the caller to add
// headers to the request.
func (c *OrganizationsAnalyticsAccountLinksCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OrganizationsAnalyticsAccountLinksCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := gensupport.SetHeaders(c.s.userAgent(), "application/json", c.header_)
	body, err := googleapi.WithoutDataWrapper.JSONBuffer(c.analyticsaccountlink)
	if err != nil {
		return nil, err
	}
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha/{+parent}/analyticsAccountLinks")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	c.s.logger.DebugContext(c.ctx_, "api request", "serviceName", apiName, "rpcName", "marketingplatformadmin.organizations.analyticsAccountLinks.create", "request", internallog.HTTPRequest(req, body.Bytes()))
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "marketingplatformadmin.organizations.analyticsAccountLinks.create" call.
// Any non-2xx status code is an error. Response headers are in either
// *AnalyticsAccountLink.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified was
// returned.
func (c *OrganizationsAnalyticsAccountLinksCreateCall) Do(opts ...googleapi.CallOption) (*AnalyticsAccountLink, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, gensupport.WrapError(&googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		})
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, gensupport.WrapError(err)
	}
	ret := &AnalyticsAccountLink{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	b, err := gensupport.DecodeResponseBytes(target, res)
	if err != nil {
		return nil, err
	}
	c.s.logger.DebugContext(c.ctx_, "api response", "serviceName", apiName, "rpcName", "marketingplatformadmin.organizations.analyticsAccountLinks.create", "response", internallog.HTTPResponse(res, b))
	return ret, nil
}

type OrganizationsAnalyticsAccountLinksDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes the AnalyticsAccountLink, which detaches the Analytics
// account from the Google Marketing Platform organization. User needs to be an
// org user, and admin on the Analytics account in order to delete the link.
//
//   - name: The name of the Analytics account link to delete. Format:
//     organizations/{org_id}/analyticsAccountLinks/{analytics_account_link_id}.
func (r *OrganizationsAnalyticsAccountLinksService) Delete(name string) *OrganizationsAnalyticsAccountLinksDeleteCall {
	c := &OrganizationsAnalyticsAccountLinksDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more
// details.
func (c *OrganizationsAnalyticsAccountLinksDeleteCall) Fields(s ...googleapi.Field) *OrganizationsAnalyticsAccountLinksDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method.
func (c *OrganizationsAnalyticsAccountLinksDeleteCall) Context(ctx context.Context) *OrganizationsAnalyticsAccountLinksDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns a http.Header that can be modified by the caller to add
// headers to the request.
func (c *OrganizationsAnalyticsAccountLinksDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OrganizationsAnalyticsAccountLinksDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := gensupport.SetHeaders(c.s.userAgent(), "", c.header_)
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("DELETE", urls, nil)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	c.s.logger.DebugContext(c.ctx_, "api request", "serviceName", apiName, "rpcName", "marketingplatformadmin.organizations.analyticsAccountLinks.delete", "request", internallog.HTTPRequest(req, nil))
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "marketingplatformadmin.organizations.analyticsAccountLinks.delete" call.
// Any non-2xx status code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was returned.
func (c *OrganizationsAnalyticsAccountLinksDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, gensupport.WrapError(&googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		})
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, gensupport.WrapError(err)
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	b, err := gensupport.DecodeResponseBytes(target, res)
	if err != nil {
		return nil, err
	}
	c.s.logger.DebugContext(c.ctx_, "api response", "serviceName", apiName, "rpcName", "marketingplatformadmin.organizations.analyticsAccountLinks.delete", "response", internallog.HTTPResponse(res, b))
	return ret, nil
}

type OrganizationsAnalyticsAccountLinksListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists the Google Analytics accounts link to the specified Google
// Marketing Platform organization.
//
//   - parent: The parent organization, which owns this collection of Analytics
//     account links. Format: organizations/{org_id}.
func (r *OrganizationsAnalyticsAccountLinksService) List(parent string) *OrganizationsAnalyticsAccountLinksListCall {
	c := &OrganizationsAnalyticsAccountLinksListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number of
// Analytics account links to return in one call. The service may return fewer
// than this value. If unspecified, at most 50 Analytics account links will be
// returned. The maximum value is 1000; values above 1000 will be coerced to
// 1000.
func (c *OrganizationsAnalyticsAccountLinksListCall) PageSize(pageSize int64) *OrganizationsAnalyticsAccountLinksListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A page token, received
// from a previous ListAnalyticsAccountLinks call. Provide this to retrieve the
// subsequent page. When paginating, all other parameters provided to
// `ListAnalyticsAccountLinks` must match the call that provided the page
// token.
func (c *OrganizationsAnalyticsAccountLinksListCall) PageToken(pageToken string) *OrganizationsAnalyticsAccountLinksListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more
// details.
func (c *OrganizationsAnalyticsAccountLinksListCall) Fields(s ...googleapi.Field) *OrganizationsAnalyticsAccountLinksListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets an optional parameter which makes the operation fail if the
// object's ETag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *OrganizationsAnalyticsAccountLinksListCall) IfNoneMatch(entityTag string) *OrganizationsAnalyticsAccountLinksListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method.
func (c *OrganizationsAnalyticsAccountLinksListCall) Context(ctx context.Context) *OrganizationsAnalyticsAccountLinksListCall {
	c.ctx_ = ctx
	return c
}

// Header returns a http.Header that can be modified by the caller to add
// headers to the request.
func (c *OrganizationsAnalyticsAccountLinksListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OrganizationsAnalyticsAccountLinksListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := gensupport.SetHeaders(c.s.userAgent(), "", c.header_)
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha/{+parent}/analyticsAccountLinks")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, nil)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	c.s.logger.DebugContext(c.ctx_, "api request", "serviceName", apiName, "rpcName", "marketingplatformadmin.organizations.analyticsAccountLinks.list", "request", internallog.HTTPRequest(req, nil))
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "marketingplatformadmin.organizations.analyticsAccountLinks.list" call.
// Any non-2xx status code is an error. Response headers are in either
// *ListAnalyticsAccountLinksResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *OrganizationsAnalyticsAccountLinksListCall) Do(opts ...googleapi.CallOption) (*ListAnalyticsAccountLinksResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, gensupport.WrapError(&googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		})
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, gensupport.WrapError(err)
	}
	ret := &ListAnalyticsAccountLinksResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	b, err := gensupport.DecodeResponseBytes(target, res)
	if err != nil {
		return nil, err
	}
	c.s.logger.DebugContext(c.ctx_, "api response", "serviceName", apiName, "rpcName", "marketingplatformadmin.organizations.analyticsAccountLinks.list", "response", internallog.HTTPResponse(res, b))
	return ret, nil
}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *OrganizationsAnalyticsAccountLinksListCall) Pages(ctx context.Context, f func(*ListAnalyticsAccountLinksResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken"))
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

type OrganizationsAnalyticsAccountLinksSetPropertyServiceLevelCall struct {
	s                              *Service
	analyticsAccountLink           string
	setpropertyservicelevelrequest *SetPropertyServiceLevelRequest
	urlParams_                     gensupport.URLParams
	ctx_                           context.Context
	header_                        http.Header
}

// SetPropertyServiceLevel: Updates the service level for an Analytics
// property.
//
//   - analyticsAccountLink: The parent AnalyticsAccountLink scope where this
//     property is in. Format:
//     organizations/{org_id}/analyticsAccountLinks/{analytics_account_link_id}.
func (r *OrganizationsAnalyticsAccountLinksService) SetPropertyServiceLevel(analyticsAccountLink string, setpropertyservicelevelrequest *SetPropertyServiceLevelRequest) *OrganizationsAnalyticsAccountLinksSetPropertyServiceLevelCall {
	c := &OrganizationsAnalyticsAccountLinksSetPropertyServiceLevelCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.analyticsAccountLink = analyticsAccountLink
	c.setpropertyservicelevelrequest = setpropertyservicelevelrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more
// details.
func (c *OrganizationsAnalyticsAccountLinksSetPropertyServiceLevelCall) Fields(s ...googleapi.Field) *OrganizationsAnalyticsAccountLinksSetPropertyServiceLevelCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method.
func (c *OrganizationsAnalyticsAccountLinksSetPropertyServiceLevelCall) Context(ctx context.Context) *OrganizationsAnalyticsAccountLinksSetPropertyServiceLevelCall {
	c.ctx_ = ctx
	return c
}

// Header returns a http.Header that can be modified by the caller to add
// headers to the request.
func (c *OrganizationsAnalyticsAccountLinksSetPropertyServiceLevelCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OrganizationsAnalyticsAccountLinksSetPropertyServiceLevelCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := gensupport.SetHeaders(c.s.userAgent(), "application/json", c.header_)
	body, err := googleapi.WithoutDataWrapper.JSONBuffer(c.setpropertyservicelevelrequest)
	if err != nil {
		return nil, err
	}
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha/{+analyticsAccountLink}:setPropertyServiceLevel")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"analyticsAccountLink": c.analyticsAccountLink,
	})
	c.s.logger.DebugContext(c.ctx_, "api request", "serviceName", apiName, "rpcName", "marketingplatformadmin.organizations.analyticsAccountLinks.setPropertyServiceLevel", "request", internallog.HTTPRequest(req, body.Bytes()))
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "marketingplatformadmin.organizations.analyticsAccountLinks.setPropertyServiceLevel" call.
// Any non-2xx status code is an error. Response headers are in either
// *SetPropertyServiceLevelResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *OrganizationsAnalyticsAccountLinksSetPropertyServiceLevelCall) Do(opts ...googleapi.CallOption) (*SetPropertyServiceLevelResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, gensupport.WrapError(&googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		})
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, gensupport.WrapError(err)
	}
	ret := &SetPropertyServiceLevelResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	b, err := gensupport.DecodeResponseBytes(target, res)
	if err != nil {
		return nil, err
	}
	c.s.logger.DebugContext(c.ctx_, "api response", "serviceName", apiName, "rpcName", "marketingplatformadmin.organizations.analyticsAccountLinks.setPropertyServiceLevel", "response", internallog.HTTPResponse(res, b))
	return ret, nil
}
