// Package androidenterprise provides access to the Google Play EMM API.
//
// Usage example:
//
//   import "google.golang.org/api/androidenterprise/v1"
//   ...
//   androidenterpriseService, err := androidenterprise.New(oauthHttpClient)
package androidenterprise

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

const apiId = "androidenterprise:v1"
const apiName = "androidenterprise"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/androidenterprise/v1/"

// OAuth2 scopes used by this API.
const (
	// Manage corporate Android devices
	AndroidenterpriseScope = "https://www.googleapis.com/auth/androidenterprise"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Collections = NewCollectionsService(s)
	s.Collectionviewers = NewCollectionviewersService(s)
	s.Devices = NewDevicesService(s)
	s.Enterprises = NewEnterprisesService(s)
	s.Entitlements = NewEntitlementsService(s)
	s.Grouplicenses = NewGrouplicensesService(s)
	s.Grouplicenseusers = NewGrouplicenseusersService(s)
	s.Installs = NewInstallsService(s)
	s.Permissions = NewPermissionsService(s)
	s.Products = NewProductsService(s)
	s.Users = NewUsersService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Collections *CollectionsService

	Collectionviewers *CollectionviewersService

	Devices *DevicesService

	Enterprises *EnterprisesService

	Entitlements *EntitlementsService

	Grouplicenses *GrouplicensesService

	Grouplicenseusers *GrouplicenseusersService

	Installs *InstallsService

	Permissions *PermissionsService

	Products *ProductsService

	Users *UsersService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewCollectionsService(s *Service) *CollectionsService {
	rs := &CollectionsService{s: s}
	return rs
}

type CollectionsService struct {
	s *Service
}

func NewCollectionviewersService(s *Service) *CollectionviewersService {
	rs := &CollectionviewersService{s: s}
	return rs
}

type CollectionviewersService struct {
	s *Service
}

func NewDevicesService(s *Service) *DevicesService {
	rs := &DevicesService{s: s}
	return rs
}

type DevicesService struct {
	s *Service
}

func NewEnterprisesService(s *Service) *EnterprisesService {
	rs := &EnterprisesService{s: s}
	return rs
}

type EnterprisesService struct {
	s *Service
}

func NewEntitlementsService(s *Service) *EntitlementsService {
	rs := &EntitlementsService{s: s}
	return rs
}

type EntitlementsService struct {
	s *Service
}

func NewGrouplicensesService(s *Service) *GrouplicensesService {
	rs := &GrouplicensesService{s: s}
	return rs
}

type GrouplicensesService struct {
	s *Service
}

func NewGrouplicenseusersService(s *Service) *GrouplicenseusersService {
	rs := &GrouplicenseusersService{s: s}
	return rs
}

type GrouplicenseusersService struct {
	s *Service
}

func NewInstallsService(s *Service) *InstallsService {
	rs := &InstallsService{s: s}
	return rs
}

type InstallsService struct {
	s *Service
}

func NewPermissionsService(s *Service) *PermissionsService {
	rs := &PermissionsService{s: s}
	return rs
}

type PermissionsService struct {
	s *Service
}

func NewProductsService(s *Service) *ProductsService {
	rs := &ProductsService{s: s}
	return rs
}

type ProductsService struct {
	s *Service
}

func NewUsersService(s *Service) *UsersService {
	rs := &UsersService{s: s}
	return rs
}

type UsersService struct {
	s *Service
}

type AppRestrictionsSchema struct {
	// Restrictions: The set of restrictions that make up this schema.
	Restrictions []*AppRestrictionsSchemaRestriction `json:"restrictions,omitempty"`
}

type AppRestrictionsSchemaRestriction struct {
	// DefaultValue: The default value of the restriction.
	DefaultValue *AppRestrictionsSchemaRestrictionRestrictionValue `json:"defaultValue,omitempty"`

	// Description: A longer description of the restriction, giving more
	// detail of what it affects.
	Description string `json:"description,omitempty"`

	// Entry: For choice or multiselect restrictions, the list of possible
	// entries' human-readable names.
	Entry []string `json:"entry,omitempty"`

	// EntryValue: For choice or multiselect restrictions, the list of
	// possible entries' machine-readable values.
	EntryValue []string `json:"entryValue,omitempty"`

	// Key: The unique key that the product uses to identify the
	// restriction, e.g. "com.google.android.gm.fieldname".
	Key string `json:"key,omitempty"`

	// RestrictionType: The type of the restriction.
	RestrictionType string `json:"restrictionType,omitempty"`

	// Title: The name of the restriction.
	Title string `json:"title,omitempty"`
}

type AppRestrictionsSchemaRestrictionRestrictionValue struct {
	// Type: The type of the value being provided.
	Type string `json:"type,omitempty"`

	// ValueBool: The boolean value - this will only be present if type is
	// bool.
	ValueBool bool `json:"valueBool,omitempty"`

	// ValueInteger: The integer value - this will only be present if type
	// is integer.
	ValueInteger int64 `json:"valueInteger,omitempty"`

	// ValueMultiselect: The list of string values - this will only be
	// present if type is multiselect.
	ValueMultiselect []string `json:"valueMultiselect,omitempty"`

	// ValueString: The string value - this will be present for types
	// string, choice and hidden.
	ValueString string `json:"valueString,omitempty"`
}

type AppVersion struct {
	// VersionCode: Unique increasing identifier for the apk version.
	VersionCode int64 `json:"versionCode,omitempty"`

	// VersionString: The string used in the Play Store by the app developer
	// to identify a version of an app. The string is not necessarily unique
	// or localized (e.g. "1.4").
	VersionString string `json:"versionString,omitempty"`
}

type ApprovalUrlInfo struct {
	// ApprovalUrl: A URL that displays a product's permissions and that can
	// also be used to approve the product with the Products.approve call.
	ApprovalUrl string `json:"approvalUrl,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#approvalUrlInfo".
	Kind string `json:"kind,omitempty"`
}

type Collection struct {
	// CollectionId: Arbitrary unique ID, allocated by the API on creation.
	CollectionId string `json:"collectionId,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#collection".
	Kind string `json:"kind,omitempty"`

	// Name: A user-friendly name for the collection (should be unique),
	// e.g. "Accounting apps".
	Name string `json:"name,omitempty"`

	// ProductId: The IDs of the products in the collection, in the order in
	// which they should be displayed.
	ProductId []string `json:"productId,omitempty"`

	// Visibility: Whether this collection is visible to all users, or only
	// to the users that have been granted access through the
	// collection_viewers api. Even if a collection is visible to allUsers,
	// it is possible to add and remove viewers, but this will have no
	// effect until the collection's visibility changes to viewersOnly.
	Visibility string `json:"visibility,omitempty"`
}

type CollectionViewersListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#collectionViewersListResponse".
	Kind string `json:"kind,omitempty"`

	// User: A user of an enterprise.
	User []*User `json:"user,omitempty"`
}

type CollectionsListResponse struct {
	// Collection: An ordered collection of products which can be made
	// visible on the Google Play Store app to a selected group of users.
	Collection []*Collection `json:"collection,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#collectionsListResponse".
	Kind string `json:"kind,omitempty"`
}

type Device struct {
	// AndroidId: The Google Play Services Android ID for the device encoded
	// as a lowercase hex string, e.g. "123456789abcdef0".
	AndroidId string `json:"androidId,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#device".
	Kind string `json:"kind,omitempty"`

	// ManagementType: The mechanism by which this device is managed by the
	// MDM. "managedDevice" means that the MDM's app is a device owner.
	// "managedProfile" means that the MDM's app is the profile owner (and
	// there is a separate personal profile which is not managed).
	// "containerApp" means that the MDM's app is managing the Android for
	// Work container app on the device.
	ManagementType string `json:"managementType,omitempty"`
}

type DeviceState struct {
	// AccountState: The state of the Google account on the device.
	// "enabled" indicates that the Google account on the device can be used
	// to access Google services (including Google Play), while "disabled"
	// means that it cannot. A new device is initially in the "disabled"
	// state.
	AccountState string `json:"accountState,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#deviceState".
	Kind string `json:"kind,omitempty"`
}

type DevicesListResponse struct {
	// Device: A managed device.
	Device []*Device `json:"device,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#devicesListResponse".
	Kind string `json:"kind,omitempty"`
}

type Enterprise struct {
	// Id: The unique ID for the enterprise.
	Id string `json:"id,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#enterprise".
	Kind string `json:"kind,omitempty"`

	// Name: The name of the enterprise, e.g. "Example Inc".
	Name string `json:"name,omitempty"`

	// PrimaryDomain: The enterprise's primary domain, e.g. "example.com".
	PrimaryDomain string `json:"primaryDomain,omitempty"`
}

type EnterpriseAccount struct {
	// AccountEmail: The email address of the service account.
	AccountEmail string `json:"accountEmail,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#enterpriseAccount".
	Kind string `json:"kind,omitempty"`
}

type EnterprisesListResponse struct {
	// Enterprise: An enterprise.
	Enterprise []*Enterprise `json:"enterprise,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#enterprisesListResponse".
	Kind string `json:"kind,omitempty"`
}

type Entitlement struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#entitlement".
	Kind string `json:"kind,omitempty"`

	// ProductId: The ID of the product that the entitlement is for, e.g.
	// "app:com.google.android.gm".
	ProductId string `json:"productId,omitempty"`

	// Reason: The reason for the entitlement, e.g. "free" for free apps.
	// This is temporary, it will be replaced by the acquisition kind field
	// of group licenses.
	Reason string `json:"reason,omitempty"`
}

type EntitlementsListResponse struct {
	// Entitlement: An entitlement of a user to a product (e.g. an app). For
	// example, a free app that they have installed, or a paid app that they
	// have been allocated a license to.
	Entitlement []*Entitlement `json:"entitlement,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#entitlementsListResponse".
	Kind string `json:"kind,omitempty"`
}

type GroupLicense struct {
	// AcquisitionKind: How this group license was acquired. "bulkPurchase"
	// means that this group license object was created because the
	// enterprise purchased licenses for this product; this is "free"
	// otherwise (for free products).
	AcquisitionKind string `json:"acquisitionKind,omitempty"`

	// Approval: Whether the product to which this group license relates is
	// currently approved by the enterprise, as either "approved" or
	// "unapproved". Products are approved when a group license is first
	// created, but this approval may be revoked by an enterprise admin via
	// Google Play. Unapproved products will not be visible to end users in
	// collections and new entitlements to them should not normally be
	// created.
	Approval string `json:"approval,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#groupLicense".
	Kind string `json:"kind,omitempty"`

	// NumProvisioned: The total number of provisioned licenses for this
	// product. Returned by read operations, but ignored in write
	// operations.
	NumProvisioned int64 `json:"numProvisioned,omitempty"`

	// NumPurchased: The number of purchased licenses (possibly in multiple
	// purchases). If this field is omitted then there is no limit on the
	// number of licenses that can be provisioned (e.g. if the acquisition
	// kind is "free").
	NumPurchased int64 `json:"numPurchased,omitempty"`

	// ProductId: The ID of the product that the license is for, e.g.
	// "app:com.google.android.gm".
	ProductId string `json:"productId,omitempty"`
}

type GroupLicenseUsersListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#groupLicenseUsersListResponse".
	Kind string `json:"kind,omitempty"`

	// User: A user of an enterprise.
	User []*User `json:"user,omitempty"`
}

type GroupLicensesListResponse struct {
	// GroupLicense: A group license for a product approved for use in the
	// enterprise.
	GroupLicense []*GroupLicense `json:"groupLicense,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#groupLicensesListResponse".
	Kind string `json:"kind,omitempty"`
}

type Install struct {
	// InstallState: Install state. The state "installPending" means that an
	// install request has recently been made and download to the device is
	// in progress. The state "installed" means that the app has been
	// installed. This field is read-only.
	InstallState string `json:"installState,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#install".
	Kind string `json:"kind,omitempty"`

	// ProductId: The ID of the product that the install is for, e.g.
	// "app:com.google.android.gm".
	ProductId string `json:"productId,omitempty"`

	// VersionCode: The version of the installed product. Guaranteed to be
	// set only if the install state is "installed".
	VersionCode int64 `json:"versionCode,omitempty"`
}

type InstallsListResponse struct {
	// Install: An installation of an app for a user on a specific device.
	// The existence of an install implies that the user must have an
	// entitlement to the app.
	Install []*Install `json:"install,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#installsListResponse".
	Kind string `json:"kind,omitempty"`
}

type Permission struct {
	// Description: A longer description of the permissions giving more
	// details of what it affects.
	Description string `json:"description,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#permission".
	Kind string `json:"kind,omitempty"`

	// Name: The name of the permission.
	Name string `json:"name,omitempty"`

	// PermissionId: An opaque string uniquely identifying the permission.
	PermissionId string `json:"permissionId,omitempty"`
}

type Product struct {
	// AppVersion: List of app versions available for this product. The
	// returned list contains only public versions. E.g. alpha, beta or
	// canary versions will not be included.
	AppVersion []*AppVersion `json:"appVersion,omitempty"`

	// AuthorName: The name of the author of the product (e.g. the app
	// developer).
	AuthorName string `json:"authorName,omitempty"`

	// DetailsUrl: A link to the (consumer) Google Play details page for the
	// product.
	DetailsUrl string `json:"detailsUrl,omitempty"`

	// DistributionChannel: How and to whom the package is made available.
	DistributionChannel string `json:"distributionChannel,omitempty"`

	// IconUrl: A link to an image that can be used as an icon for the
	// product.
	IconUrl string `json:"iconUrl,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#product".
	Kind string `json:"kind,omitempty"`

	// ProductId: A string of the form "app:
	// " - e.g. "app:com.google.android.gm" represents the GMail app.
	ProductId string `json:"productId,omitempty"`

	// RequiresContainerApp: Whether this app can only be installed on
	// devices using the Android for Work container app.
	RequiresContainerApp bool `json:"requiresContainerApp,omitempty"`

	// Title: The name of the product.
	Title string `json:"title,omitempty"`

	// WorkDetailsUrl: A link to the Google Play for Work details page for
	// the product, for use by an Enterprise administrator.
	WorkDetailsUrl string `json:"workDetailsUrl,omitempty"`
}

type ProductPermission struct {
	// PermissionId: An opaque string uniquely identifying the permission.
	PermissionId string `json:"permissionId,omitempty"`

	// State: Whether the permission has been accepted or not.
	State string `json:"state,omitempty"`
}

type ProductPermissions struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#productPermissions".
	Kind string `json:"kind,omitempty"`

	// Permission: The permissions required by the app.
	Permission []*ProductPermission `json:"permission,omitempty"`

	// ProductId: The ID of the app that the permissions relate to, e.g.
	// "app:com.google.android.gm".
	ProductId string `json:"productId,omitempty"`
}

type ProductsApproveRequest struct {
	ApprovalUrlInfo *ApprovalUrlInfo `json:"approvalUrlInfo,omitempty"`
}

type ProductsGenerateApprovalUrlResponse struct {
	// Url: A iframe-able URL that displays a product's permissions (if
	// any). This URL can be used to approve the product only once and for a
	// limited time (1 hour), using the Products.approve call. If the
	// product is not currently approved and has no permissions, this URL
	// will point to an empty page. If the product is currently approved and
	// all of its permissions (if any) are also approved, this field will
	// not be populated.
	Url string `json:"url,omitempty"`
}

type User struct {
	// Id: The unique ID for the user.
	Id string `json:"id,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#user".
	Kind string `json:"kind,omitempty"`

	// PrimaryEmail: The user's primary email, e.g. "jsmith@example.com".
	PrimaryEmail string `json:"primaryEmail,omitempty"`
}

type UserToken struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#userToken".
	Kind string `json:"kind,omitempty"`

	// Token: The token (activation code) to be entered by the user. This
	// consists of a sequence of decimal digits. Note that the leading digit
	// may be 0.
	Token string `json:"token,omitempty"`

	// UserId: The unique ID for the user.
	UserId string `json:"userId,omitempty"`
}

type UsersListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#usersListResponse".
	Kind string `json:"kind,omitempty"`

	// User: A user of an enterprise.
	User []*User `json:"user,omitempty"`
}

// method id "androidenterprise.collections.delete":

type CollectionsDeleteCall struct {
	s            *Service
	enterpriseId string
	collectionId string
	opt_         map[string]interface{}
}

// Delete: Deletes a collection.
func (r *CollectionsService) Delete(enterpriseId string, collectionId string) *CollectionsDeleteCall {
	c := &CollectionsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.collectionId = collectionId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CollectionsDeleteCall) Fields(s ...googleapi.Field) *CollectionsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *CollectionsDeleteCall) IfNoneMatch(entityTag string) *CollectionsDeleteCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.collections.delete" call.
func (c *CollectionsDeleteCall) Do() error {
	_, err := c.DoHeader()
	return err
}

// DoHeader executes the "androidenterprise.collections.delete" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *CollectionsDeleteCall) DoHeader() (resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/collections/{collectionId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
		"collectionId": c.collectionId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return res.Header, err
	}
	return res.Header, nil
	// {
	//   "description": "Deletes a collection.",
	//   "httpMethod": "DELETE",
	//   "id": "androidenterprise.collections.delete",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "collectionId"
	//   ],
	//   "parameters": {
	//     "collectionId": {
	//       "description": "The ID of the collection.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/collections/{collectionId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.collections.get":

type CollectionsGetCall struct {
	s            *Service
	enterpriseId string
	collectionId string
	opt_         map[string]interface{}
}

// Get: Retrieves the details of a collection.
func (r *CollectionsService) Get(enterpriseId string, collectionId string) *CollectionsGetCall {
	c := &CollectionsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.collectionId = collectionId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CollectionsGetCall) Fields(s ...googleapi.Field) *CollectionsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *CollectionsGetCall) IfNoneMatch(entityTag string) *CollectionsGetCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.collections.get" call.
// Exactly one of the return values is non-nil.
func (c *CollectionsGetCall) Do() (*Collection, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.collections.get" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *CollectionsGetCall) DoHeader() (ret *Collection, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/collections/{collectionId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
		"collectionId": c.collectionId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Retrieves the details of a collection.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.collections.get",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "collectionId"
	//   ],
	//   "parameters": {
	//     "collectionId": {
	//       "description": "The ID of the collection.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/collections/{collectionId}",
	//   "response": {
	//     "$ref": "Collection"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.collections.insert":

type CollectionsInsertCall struct {
	s            *Service
	enterpriseId string
	collection   *Collection
	opt_         map[string]interface{}
}

// Insert: Creates a new collection.
func (r *CollectionsService) Insert(enterpriseId string, collection *Collection) *CollectionsInsertCall {
	c := &CollectionsInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.collection = collection
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CollectionsInsertCall) Fields(s ...googleapi.Field) *CollectionsInsertCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *CollectionsInsertCall) IfNoneMatch(entityTag string) *CollectionsInsertCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.collections.insert" call.
// Exactly one of the return values is non-nil.
func (c *CollectionsInsertCall) Do() (*Collection, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.collections.insert" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *CollectionsInsertCall) DoHeader() (ret *Collection, resHeader http.Header, err error) {
	var body io.Reader = nil
	body, err = googleapi.WithoutDataWrapper.JSONReader(c.collection)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/collections")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Creates a new collection.",
	//   "httpMethod": "POST",
	//   "id": "androidenterprise.collections.insert",
	//   "parameterOrder": [
	//     "enterpriseId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/collections",
	//   "request": {
	//     "$ref": "Collection"
	//   },
	//   "response": {
	//     "$ref": "Collection"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.collections.list":

type CollectionsListCall struct {
	s            *Service
	enterpriseId string
	opt_         map[string]interface{}
}

// List: Retrieves the IDs of all the collections for an enterprise.
func (r *CollectionsService) List(enterpriseId string) *CollectionsListCall {
	c := &CollectionsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CollectionsListCall) Fields(s ...googleapi.Field) *CollectionsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *CollectionsListCall) IfNoneMatch(entityTag string) *CollectionsListCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.collections.list" call.
// Exactly one of the return values is non-nil.
func (c *CollectionsListCall) Do() (*CollectionsListResponse, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.collections.list" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *CollectionsListCall) DoHeader() (ret *CollectionsListResponse, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/collections")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Retrieves the IDs of all the collections for an enterprise.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.collections.list",
	//   "parameterOrder": [
	//     "enterpriseId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/collections",
	//   "response": {
	//     "$ref": "CollectionsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.collections.patch":

type CollectionsPatchCall struct {
	s            *Service
	enterpriseId string
	collectionId string
	collection   *Collection
	opt_         map[string]interface{}
}

// Patch: Updates a collection. This method supports patch semantics.
func (r *CollectionsService) Patch(enterpriseId string, collectionId string, collection *Collection) *CollectionsPatchCall {
	c := &CollectionsPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.collectionId = collectionId
	c.collection = collection
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CollectionsPatchCall) Fields(s ...googleapi.Field) *CollectionsPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *CollectionsPatchCall) IfNoneMatch(entityTag string) *CollectionsPatchCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.collections.patch" call.
// Exactly one of the return values is non-nil.
func (c *CollectionsPatchCall) Do() (*Collection, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.collections.patch" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *CollectionsPatchCall) DoHeader() (ret *Collection, resHeader http.Header, err error) {
	var body io.Reader = nil
	body, err = googleapi.WithoutDataWrapper.JSONReader(c.collection)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/collections/{collectionId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
		"collectionId": c.collectionId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Updates a collection. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "androidenterprise.collections.patch",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "collectionId"
	//   ],
	//   "parameters": {
	//     "collectionId": {
	//       "description": "The ID of the collection.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/collections/{collectionId}",
	//   "request": {
	//     "$ref": "Collection"
	//   },
	//   "response": {
	//     "$ref": "Collection"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.collections.update":

type CollectionsUpdateCall struct {
	s            *Service
	enterpriseId string
	collectionId string
	collection   *Collection
	opt_         map[string]interface{}
}

// Update: Updates a collection.
func (r *CollectionsService) Update(enterpriseId string, collectionId string, collection *Collection) *CollectionsUpdateCall {
	c := &CollectionsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.collectionId = collectionId
	c.collection = collection
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CollectionsUpdateCall) Fields(s ...googleapi.Field) *CollectionsUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *CollectionsUpdateCall) IfNoneMatch(entityTag string) *CollectionsUpdateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.collections.update" call.
// Exactly one of the return values is non-nil.
func (c *CollectionsUpdateCall) Do() (*Collection, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.collections.update" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *CollectionsUpdateCall) DoHeader() (ret *Collection, resHeader http.Header, err error) {
	var body io.Reader = nil
	body, err = googleapi.WithoutDataWrapper.JSONReader(c.collection)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/collections/{collectionId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
		"collectionId": c.collectionId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Updates a collection.",
	//   "httpMethod": "PUT",
	//   "id": "androidenterprise.collections.update",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "collectionId"
	//   ],
	//   "parameters": {
	//     "collectionId": {
	//       "description": "The ID of the collection.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/collections/{collectionId}",
	//   "request": {
	//     "$ref": "Collection"
	//   },
	//   "response": {
	//     "$ref": "Collection"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.collectionviewers.delete":

type CollectionviewersDeleteCall struct {
	s            *Service
	enterpriseId string
	collectionId string
	userId       string
	opt_         map[string]interface{}
}

// Delete: Removes the user from the list of those specifically allowed
// to see the collection. If the collection's visibility is set to
// viewersOnly then only such users will see the collection.
func (r *CollectionviewersService) Delete(enterpriseId string, collectionId string, userId string) *CollectionviewersDeleteCall {
	c := &CollectionviewersDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.collectionId = collectionId
	c.userId = userId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CollectionviewersDeleteCall) Fields(s ...googleapi.Field) *CollectionviewersDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *CollectionviewersDeleteCall) IfNoneMatch(entityTag string) *CollectionviewersDeleteCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.collectionviewers.delete" call.
func (c *CollectionviewersDeleteCall) Do() error {
	_, err := c.DoHeader()
	return err
}

// DoHeader executes the "androidenterprise.collectionviewers.delete" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *CollectionviewersDeleteCall) DoHeader() (resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/collections/{collectionId}/users/{userId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
		"collectionId": c.collectionId,
		"userId":       c.userId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return res.Header, err
	}
	return res.Header, nil
	// {
	//   "description": "Removes the user from the list of those specifically allowed to see the collection. If the collection's visibility is set to viewersOnly then only such users will see the collection.",
	//   "httpMethod": "DELETE",
	//   "id": "androidenterprise.collectionviewers.delete",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "collectionId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "collectionId": {
	//       "description": "The ID of the collection.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/collections/{collectionId}/users/{userId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.collectionviewers.get":

type CollectionviewersGetCall struct {
	s            *Service
	enterpriseId string
	collectionId string
	userId       string
	opt_         map[string]interface{}
}

// Get: Retrieves the ID of the user if they have been specifically
// allowed to see the collection. If the collection's visibility is set
// to viewersOnly then only these users will see the collection.
func (r *CollectionviewersService) Get(enterpriseId string, collectionId string, userId string) *CollectionviewersGetCall {
	c := &CollectionviewersGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.collectionId = collectionId
	c.userId = userId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CollectionviewersGetCall) Fields(s ...googleapi.Field) *CollectionviewersGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *CollectionviewersGetCall) IfNoneMatch(entityTag string) *CollectionviewersGetCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.collectionviewers.get" call.
// Exactly one of the return values is non-nil.
func (c *CollectionviewersGetCall) Do() (*User, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.collectionviewers.get" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *CollectionviewersGetCall) DoHeader() (ret *User, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/collections/{collectionId}/users/{userId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
		"collectionId": c.collectionId,
		"userId":       c.userId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Retrieves the ID of the user if they have been specifically allowed to see the collection. If the collection's visibility is set to viewersOnly then only these users will see the collection.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.collectionviewers.get",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "collectionId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "collectionId": {
	//       "description": "The ID of the collection.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/collections/{collectionId}/users/{userId}",
	//   "response": {
	//     "$ref": "User"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.collectionviewers.list":

type CollectionviewersListCall struct {
	s            *Service
	enterpriseId string
	collectionId string
	opt_         map[string]interface{}
}

// List: Retrieves the IDs of the users who have been specifically
// allowed to see the collection. If the collection's visibility is set
// to viewersOnly then only these users will see the collection.
func (r *CollectionviewersService) List(enterpriseId string, collectionId string) *CollectionviewersListCall {
	c := &CollectionviewersListCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.collectionId = collectionId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CollectionviewersListCall) Fields(s ...googleapi.Field) *CollectionviewersListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *CollectionviewersListCall) IfNoneMatch(entityTag string) *CollectionviewersListCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.collectionviewers.list" call.
// Exactly one of the return values is non-nil.
func (c *CollectionviewersListCall) Do() (*CollectionViewersListResponse, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.collectionviewers.list" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *CollectionviewersListCall) DoHeader() (ret *CollectionViewersListResponse, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/collections/{collectionId}/users")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
		"collectionId": c.collectionId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Retrieves the IDs of the users who have been specifically allowed to see the collection. If the collection's visibility is set to viewersOnly then only these users will see the collection.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.collectionviewers.list",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "collectionId"
	//   ],
	//   "parameters": {
	//     "collectionId": {
	//       "description": "The ID of the collection.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/collections/{collectionId}/users",
	//   "response": {
	//     "$ref": "CollectionViewersListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.collectionviewers.patch":

type CollectionviewersPatchCall struct {
	s            *Service
	enterpriseId string
	collectionId string
	userId       string
	user         *User
	opt_         map[string]interface{}
}

// Patch: Adds the user to the list of those specifically allowed to see
// the collection. If the collection's visibility is set to viewersOnly
// then only such users will see the collection. This method supports
// patch semantics.
func (r *CollectionviewersService) Patch(enterpriseId string, collectionId string, userId string, user *User) *CollectionviewersPatchCall {
	c := &CollectionviewersPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.collectionId = collectionId
	c.userId = userId
	c.user = user
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CollectionviewersPatchCall) Fields(s ...googleapi.Field) *CollectionviewersPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *CollectionviewersPatchCall) IfNoneMatch(entityTag string) *CollectionviewersPatchCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.collectionviewers.patch" call.
// Exactly one of the return values is non-nil.
func (c *CollectionviewersPatchCall) Do() (*User, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.collectionviewers.patch" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *CollectionviewersPatchCall) DoHeader() (ret *User, resHeader http.Header, err error) {
	var body io.Reader = nil
	body, err = googleapi.WithoutDataWrapper.JSONReader(c.user)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/collections/{collectionId}/users/{userId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
		"collectionId": c.collectionId,
		"userId":       c.userId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Adds the user to the list of those specifically allowed to see the collection. If the collection's visibility is set to viewersOnly then only such users will see the collection. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "androidenterprise.collectionviewers.patch",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "collectionId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "collectionId": {
	//       "description": "The ID of the collection.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/collections/{collectionId}/users/{userId}",
	//   "request": {
	//     "$ref": "User"
	//   },
	//   "response": {
	//     "$ref": "User"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.collectionviewers.update":

type CollectionviewersUpdateCall struct {
	s            *Service
	enterpriseId string
	collectionId string
	userId       string
	user         *User
	opt_         map[string]interface{}
}

// Update: Adds the user to the list of those specifically allowed to
// see the collection. If the collection's visibility is set to
// viewersOnly then only such users will see the collection.
func (r *CollectionviewersService) Update(enterpriseId string, collectionId string, userId string, user *User) *CollectionviewersUpdateCall {
	c := &CollectionviewersUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.collectionId = collectionId
	c.userId = userId
	c.user = user
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CollectionviewersUpdateCall) Fields(s ...googleapi.Field) *CollectionviewersUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *CollectionviewersUpdateCall) IfNoneMatch(entityTag string) *CollectionviewersUpdateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.collectionviewers.update" call.
// Exactly one of the return values is non-nil.
func (c *CollectionviewersUpdateCall) Do() (*User, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.collectionviewers.update" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *CollectionviewersUpdateCall) DoHeader() (ret *User, resHeader http.Header, err error) {
	var body io.Reader = nil
	body, err = googleapi.WithoutDataWrapper.JSONReader(c.user)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/collections/{collectionId}/users/{userId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
		"collectionId": c.collectionId,
		"userId":       c.userId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Adds the user to the list of those specifically allowed to see the collection. If the collection's visibility is set to viewersOnly then only such users will see the collection.",
	//   "httpMethod": "PUT",
	//   "id": "androidenterprise.collectionviewers.update",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "collectionId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "collectionId": {
	//       "description": "The ID of the collection.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/collections/{collectionId}/users/{userId}",
	//   "request": {
	//     "$ref": "User"
	//   },
	//   "response": {
	//     "$ref": "User"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.devices.get":

type DevicesGetCall struct {
	s            *Service
	enterpriseId string
	userId       string
	deviceId     string
	opt_         map[string]interface{}
}

// Get: Retrieves the details of a device.
func (r *DevicesService) Get(enterpriseId string, userId string, deviceId string) *DevicesGetCall {
	c := &DevicesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.userId = userId
	c.deviceId = deviceId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DevicesGetCall) Fields(s ...googleapi.Field) *DevicesGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *DevicesGetCall) IfNoneMatch(entityTag string) *DevicesGetCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.devices.get" call.
// Exactly one of the return values is non-nil.
func (c *DevicesGetCall) Do() (*Device, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.devices.get" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *DevicesGetCall) DoHeader() (ret *Device, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
		"userId":       c.userId,
		"deviceId":     c.deviceId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Retrieves the details of a device.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.devices.get",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId",
	//     "deviceId"
	//   ],
	//   "parameters": {
	//     "deviceId": {
	//       "description": "The ID of the device.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}",
	//   "response": {
	//     "$ref": "Device"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.devices.getState":

type DevicesGetStateCall struct {
	s            *Service
	enterpriseId string
	userId       string
	deviceId     string
	opt_         map[string]interface{}
}

// GetState: Retrieves whether a device is enabled or disabled for
// access by the user to Google services. The device state takes effect
// only if enforcing EMM policies on Android devices is enabled in the
// Google Admin Console. Otherwise, the device state is ignored and all
// devices are allowed access to Google services.
func (r *DevicesService) GetState(enterpriseId string, userId string, deviceId string) *DevicesGetStateCall {
	c := &DevicesGetStateCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.userId = userId
	c.deviceId = deviceId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DevicesGetStateCall) Fields(s ...googleapi.Field) *DevicesGetStateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *DevicesGetStateCall) IfNoneMatch(entityTag string) *DevicesGetStateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.devices.getState" call.
// Exactly one of the return values is non-nil.
func (c *DevicesGetStateCall) Do() (*DeviceState, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.devices.getState" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *DevicesGetStateCall) DoHeader() (ret *DeviceState, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}/state")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
		"userId":       c.userId,
		"deviceId":     c.deviceId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Retrieves whether a device is enabled or disabled for access by the user to Google services. The device state takes effect only if enforcing EMM policies on Android devices is enabled in the Google Admin Console. Otherwise, the device state is ignored and all devices are allowed access to Google services.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.devices.getState",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId",
	//     "deviceId"
	//   ],
	//   "parameters": {
	//     "deviceId": {
	//       "description": "The ID of the device.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}/state",
	//   "response": {
	//     "$ref": "DeviceState"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.devices.list":

type DevicesListCall struct {
	s            *Service
	enterpriseId string
	userId       string
	opt_         map[string]interface{}
}

// List: Retrieves the IDs of all of a user's devices.
func (r *DevicesService) List(enterpriseId string, userId string) *DevicesListCall {
	c := &DevicesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.userId = userId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DevicesListCall) Fields(s ...googleapi.Field) *DevicesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *DevicesListCall) IfNoneMatch(entityTag string) *DevicesListCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.devices.list" call.
// Exactly one of the return values is non-nil.
func (c *DevicesListCall) Do() (*DevicesListResponse, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.devices.list" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *DevicesListCall) DoHeader() (ret *DevicesListResponse, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/users/{userId}/devices")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
		"userId":       c.userId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Retrieves the IDs of all of a user's devices.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.devices.list",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/devices",
	//   "response": {
	//     "$ref": "DevicesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.devices.setState":

type DevicesSetStateCall struct {
	s            *Service
	enterpriseId string
	userId       string
	deviceId     string
	devicestate  *DeviceState
	opt_         map[string]interface{}
}

// SetState: Sets whether a device is enabled or disabled for access by
// the user to Google services. The device state takes effect only if
// enforcing EMM policies on Android devices is enabled in the Google
// Admin Console. Otherwise, the device state is ignored and all devices
// are allowed access to Google services.
func (r *DevicesService) SetState(enterpriseId string, userId string, deviceId string, devicestate *DeviceState) *DevicesSetStateCall {
	c := &DevicesSetStateCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.userId = userId
	c.deviceId = deviceId
	c.devicestate = devicestate
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DevicesSetStateCall) Fields(s ...googleapi.Field) *DevicesSetStateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *DevicesSetStateCall) IfNoneMatch(entityTag string) *DevicesSetStateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.devices.setState" call.
// Exactly one of the return values is non-nil.
func (c *DevicesSetStateCall) Do() (*DeviceState, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.devices.setState" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *DevicesSetStateCall) DoHeader() (ret *DeviceState, resHeader http.Header, err error) {
	var body io.Reader = nil
	body, err = googleapi.WithoutDataWrapper.JSONReader(c.devicestate)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}/state")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
		"userId":       c.userId,
		"deviceId":     c.deviceId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Sets whether a device is enabled or disabled for access by the user to Google services. The device state takes effect only if enforcing EMM policies on Android devices is enabled in the Google Admin Console. Otherwise, the device state is ignored and all devices are allowed access to Google services.",
	//   "httpMethod": "PUT",
	//   "id": "androidenterprise.devices.setState",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId",
	//     "deviceId"
	//   ],
	//   "parameters": {
	//     "deviceId": {
	//       "description": "The ID of the device.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}/state",
	//   "request": {
	//     "$ref": "DeviceState"
	//   },
	//   "response": {
	//     "$ref": "DeviceState"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.enterprises.delete":

type EnterprisesDeleteCall struct {
	s            *Service
	enterpriseId string
	opt_         map[string]interface{}
}

// Delete: Deletes the binding between the MDM and enterprise. This is
// now deprecated; use this to unenroll customers that were previously
// enrolled with the 'insert' call, then enroll them again with the
// 'enroll' call.
func (r *EnterprisesService) Delete(enterpriseId string) *EnterprisesDeleteCall {
	c := &EnterprisesDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EnterprisesDeleteCall) Fields(s ...googleapi.Field) *EnterprisesDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *EnterprisesDeleteCall) IfNoneMatch(entityTag string) *EnterprisesDeleteCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.enterprises.delete" call.
func (c *EnterprisesDeleteCall) Do() error {
	_, err := c.DoHeader()
	return err
}

// DoHeader executes the "androidenterprise.enterprises.delete" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *EnterprisesDeleteCall) DoHeader() (resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return res.Header, err
	}
	return res.Header, nil
	// {
	//   "description": "Deletes the binding between the MDM and enterprise. This is now deprecated; use this to unenroll customers that were previously enrolled with the 'insert' call, then enroll them again with the 'enroll' call.",
	//   "httpMethod": "DELETE",
	//   "id": "androidenterprise.enterprises.delete",
	//   "parameterOrder": [
	//     "enterpriseId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.enterprises.enroll":

type EnterprisesEnrollCall struct {
	s          *Service
	token      string
	enterprise *Enterprise
	opt_       map[string]interface{}
}

// Enroll: Enrolls an enterprise with the calling MDM.
func (r *EnterprisesService) Enroll(token string, enterprise *Enterprise) *EnterprisesEnrollCall {
	c := &EnterprisesEnrollCall{s: r.s, opt_: make(map[string]interface{})}
	c.token = token
	c.enterprise = enterprise
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EnterprisesEnrollCall) Fields(s ...googleapi.Field) *EnterprisesEnrollCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *EnterprisesEnrollCall) IfNoneMatch(entityTag string) *EnterprisesEnrollCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.enterprises.enroll" call.
// Exactly one of the return values is non-nil.
func (c *EnterprisesEnrollCall) Do() (*Enterprise, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.enterprises.enroll" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *EnterprisesEnrollCall) DoHeader() (ret *Enterprise, resHeader http.Header, err error) {
	var body io.Reader = nil
	body, err = googleapi.WithoutDataWrapper.JSONReader(c.enterprise)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("token", fmt.Sprintf("%v", c.token))
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/enroll")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Enrolls an enterprise with the calling MDM.",
	//   "httpMethod": "POST",
	//   "id": "androidenterprise.enterprises.enroll",
	//   "parameterOrder": [
	//     "token"
	//   ],
	//   "parameters": {
	//     "token": {
	//       "description": "The token provided by the enterprise to register the MDM.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/enroll",
	//   "request": {
	//     "$ref": "Enterprise"
	//   },
	//   "response": {
	//     "$ref": "Enterprise"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.enterprises.get":

type EnterprisesGetCall struct {
	s            *Service
	enterpriseId string
	opt_         map[string]interface{}
}

// Get: Retrieves the name and domain of an enterprise.
func (r *EnterprisesService) Get(enterpriseId string) *EnterprisesGetCall {
	c := &EnterprisesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EnterprisesGetCall) Fields(s ...googleapi.Field) *EnterprisesGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *EnterprisesGetCall) IfNoneMatch(entityTag string) *EnterprisesGetCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.enterprises.get" call.
// Exactly one of the return values is non-nil.
func (c *EnterprisesGetCall) Do() (*Enterprise, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.enterprises.get" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *EnterprisesGetCall) DoHeader() (ret *Enterprise, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Retrieves the name and domain of an enterprise.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.enterprises.get",
	//   "parameterOrder": [
	//     "enterpriseId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}",
	//   "response": {
	//     "$ref": "Enterprise"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.enterprises.insert":

type EnterprisesInsertCall struct {
	s          *Service
	token      string
	enterprise *Enterprise
	opt_       map[string]interface{}
}

// Insert: Establishes the binding between the MDM and an enterprise.
// This is now deprecated; use enroll instead.
func (r *EnterprisesService) Insert(token string, enterprise *Enterprise) *EnterprisesInsertCall {
	c := &EnterprisesInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.token = token
	c.enterprise = enterprise
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EnterprisesInsertCall) Fields(s ...googleapi.Field) *EnterprisesInsertCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *EnterprisesInsertCall) IfNoneMatch(entityTag string) *EnterprisesInsertCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.enterprises.insert" call.
// Exactly one of the return values is non-nil.
func (c *EnterprisesInsertCall) Do() (*Enterprise, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.enterprises.insert" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *EnterprisesInsertCall) DoHeader() (ret *Enterprise, resHeader http.Header, err error) {
	var body io.Reader = nil
	body, err = googleapi.WithoutDataWrapper.JSONReader(c.enterprise)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("token", fmt.Sprintf("%v", c.token))
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Establishes the binding between the MDM and an enterprise. This is now deprecated; use enroll instead.",
	//   "httpMethod": "POST",
	//   "id": "androidenterprise.enterprises.insert",
	//   "parameterOrder": [
	//     "token"
	//   ],
	//   "parameters": {
	//     "token": {
	//       "description": "The token provided by the enterprise to register the MDM.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises",
	//   "request": {
	//     "$ref": "Enterprise"
	//   },
	//   "response": {
	//     "$ref": "Enterprise"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.enterprises.list":

type EnterprisesListCall struct {
	s      *Service
	domain string
	opt_   map[string]interface{}
}

// List: Looks up an enterprise by domain name.
func (r *EnterprisesService) List(domain string) *EnterprisesListCall {
	c := &EnterprisesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.domain = domain
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EnterprisesListCall) Fields(s ...googleapi.Field) *EnterprisesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *EnterprisesListCall) IfNoneMatch(entityTag string) *EnterprisesListCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.enterprises.list" call.
// Exactly one of the return values is non-nil.
func (c *EnterprisesListCall) Do() (*EnterprisesListResponse, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.enterprises.list" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *EnterprisesListCall) DoHeader() (ret *EnterprisesListResponse, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("domain", fmt.Sprintf("%v", c.domain))
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Looks up an enterprise by domain name.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.enterprises.list",
	//   "parameterOrder": [
	//     "domain"
	//   ],
	//   "parameters": {
	//     "domain": {
	//       "description": "The exact primary domain name of the enterprise to look up.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises",
	//   "response": {
	//     "$ref": "EnterprisesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.enterprises.setAccount":

type EnterprisesSetAccountCall struct {
	s                 *Service
	enterpriseId      string
	enterpriseaccount *EnterpriseAccount
	opt_              map[string]interface{}
}

// SetAccount: Set the account that will be used to authenticate to the
// API as the enterprise.
func (r *EnterprisesService) SetAccount(enterpriseId string, enterpriseaccount *EnterpriseAccount) *EnterprisesSetAccountCall {
	c := &EnterprisesSetAccountCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.enterpriseaccount = enterpriseaccount
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EnterprisesSetAccountCall) Fields(s ...googleapi.Field) *EnterprisesSetAccountCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *EnterprisesSetAccountCall) IfNoneMatch(entityTag string) *EnterprisesSetAccountCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.enterprises.setAccount" call.
// Exactly one of the return values is non-nil.
func (c *EnterprisesSetAccountCall) Do() (*EnterpriseAccount, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.enterprises.setAccount" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *EnterprisesSetAccountCall) DoHeader() (ret *EnterpriseAccount, resHeader http.Header, err error) {
	var body io.Reader = nil
	body, err = googleapi.WithoutDataWrapper.JSONReader(c.enterpriseaccount)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/account")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Set the account that will be used to authenticate to the API as the enterprise.",
	//   "httpMethod": "PUT",
	//   "id": "androidenterprise.enterprises.setAccount",
	//   "parameterOrder": [
	//     "enterpriseId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/account",
	//   "request": {
	//     "$ref": "EnterpriseAccount"
	//   },
	//   "response": {
	//     "$ref": "EnterpriseAccount"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.enterprises.unenroll":

type EnterprisesUnenrollCall struct {
	s            *Service
	enterpriseId string
	opt_         map[string]interface{}
}

// Unenroll: Unenrolls an enterprise from the calling MDM.
func (r *EnterprisesService) Unenroll(enterpriseId string) *EnterprisesUnenrollCall {
	c := &EnterprisesUnenrollCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EnterprisesUnenrollCall) Fields(s ...googleapi.Field) *EnterprisesUnenrollCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *EnterprisesUnenrollCall) IfNoneMatch(entityTag string) *EnterprisesUnenrollCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.enterprises.unenroll" call.
func (c *EnterprisesUnenrollCall) Do() error {
	_, err := c.DoHeader()
	return err
}

// DoHeader executes the "androidenterprise.enterprises.unenroll" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *EnterprisesUnenrollCall) DoHeader() (resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/unenroll")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return res.Header, err
	}
	return res.Header, nil
	// {
	//   "description": "Unenrolls an enterprise from the calling MDM.",
	//   "httpMethod": "POST",
	//   "id": "androidenterprise.enterprises.unenroll",
	//   "parameterOrder": [
	//     "enterpriseId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/unenroll",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.entitlements.delete":

type EntitlementsDeleteCall struct {
	s             *Service
	enterpriseId  string
	userId        string
	entitlementId string
	opt_          map[string]interface{}
}

// Delete: Removes an entitlement to an app for a user and uninstalls
// it.
func (r *EntitlementsService) Delete(enterpriseId string, userId string, entitlementId string) *EntitlementsDeleteCall {
	c := &EntitlementsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.userId = userId
	c.entitlementId = entitlementId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EntitlementsDeleteCall) Fields(s ...googleapi.Field) *EntitlementsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *EntitlementsDeleteCall) IfNoneMatch(entityTag string) *EntitlementsDeleteCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.entitlements.delete" call.
func (c *EntitlementsDeleteCall) Do() error {
	_, err := c.DoHeader()
	return err
}

// DoHeader executes the "androidenterprise.entitlements.delete" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *EntitlementsDeleteCall) DoHeader() (resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/users/{userId}/entitlements/{entitlementId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId":  c.enterpriseId,
		"userId":        c.userId,
		"entitlementId": c.entitlementId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return res.Header, err
	}
	return res.Header, nil
	// {
	//   "description": "Removes an entitlement to an app for a user and uninstalls it.",
	//   "httpMethod": "DELETE",
	//   "id": "androidenterprise.entitlements.delete",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId",
	//     "entitlementId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "entitlementId": {
	//       "description": "The ID of the entitlement, e.g. \"app:com.google.android.gm\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/entitlements/{entitlementId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.entitlements.get":

type EntitlementsGetCall struct {
	s             *Service
	enterpriseId  string
	userId        string
	entitlementId string
	opt_          map[string]interface{}
}

// Get: Retrieves details of an entitlement.
func (r *EntitlementsService) Get(enterpriseId string, userId string, entitlementId string) *EntitlementsGetCall {
	c := &EntitlementsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.userId = userId
	c.entitlementId = entitlementId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EntitlementsGetCall) Fields(s ...googleapi.Field) *EntitlementsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *EntitlementsGetCall) IfNoneMatch(entityTag string) *EntitlementsGetCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.entitlements.get" call.
// Exactly one of the return values is non-nil.
func (c *EntitlementsGetCall) Do() (*Entitlement, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.entitlements.get" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *EntitlementsGetCall) DoHeader() (ret *Entitlement, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/users/{userId}/entitlements/{entitlementId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId":  c.enterpriseId,
		"userId":        c.userId,
		"entitlementId": c.entitlementId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Retrieves details of an entitlement.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.entitlements.get",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId",
	//     "entitlementId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "entitlementId": {
	//       "description": "The ID of the entitlement, e.g. \"app:com.google.android.gm\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/entitlements/{entitlementId}",
	//   "response": {
	//     "$ref": "Entitlement"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.entitlements.list":

type EntitlementsListCall struct {
	s            *Service
	enterpriseId string
	userId       string
	opt_         map[string]interface{}
}

// List: List of all entitlements for the specified user. Only the ID is
// set.
func (r *EntitlementsService) List(enterpriseId string, userId string) *EntitlementsListCall {
	c := &EntitlementsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.userId = userId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EntitlementsListCall) Fields(s ...googleapi.Field) *EntitlementsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *EntitlementsListCall) IfNoneMatch(entityTag string) *EntitlementsListCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.entitlements.list" call.
// Exactly one of the return values is non-nil.
func (c *EntitlementsListCall) Do() (*EntitlementsListResponse, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.entitlements.list" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *EntitlementsListCall) DoHeader() (ret *EntitlementsListResponse, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/users/{userId}/entitlements")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
		"userId":       c.userId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "List of all entitlements for the specified user. Only the ID is set.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.entitlements.list",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/entitlements",
	//   "response": {
	//     "$ref": "EntitlementsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.entitlements.patch":

type EntitlementsPatchCall struct {
	s             *Service
	enterpriseId  string
	userId        string
	entitlementId string
	entitlement   *Entitlement
	opt_          map[string]interface{}
}

// Patch: Adds or updates an entitlement to an app for a user. This
// method supports patch semantics.
func (r *EntitlementsService) Patch(enterpriseId string, userId string, entitlementId string, entitlement *Entitlement) *EntitlementsPatchCall {
	c := &EntitlementsPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.userId = userId
	c.entitlementId = entitlementId
	c.entitlement = entitlement
	return c
}

// Install sets the optional parameter "install": Set to true to also
// install the product on all the user's devices where possible. Failure
// to install on one or more devices will not prevent this operation
// from returning successfully, as long as the entitlement was
// successfully assigned to the user.
func (c *EntitlementsPatchCall) Install(install bool) *EntitlementsPatchCall {
	c.opt_["install"] = install
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EntitlementsPatchCall) Fields(s ...googleapi.Field) *EntitlementsPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *EntitlementsPatchCall) IfNoneMatch(entityTag string) *EntitlementsPatchCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.entitlements.patch" call.
// Exactly one of the return values is non-nil.
func (c *EntitlementsPatchCall) Do() (*Entitlement, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.entitlements.patch" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *EntitlementsPatchCall) DoHeader() (ret *Entitlement, resHeader http.Header, err error) {
	var body io.Reader = nil
	body, err = googleapi.WithoutDataWrapper.JSONReader(c.entitlement)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["install"]; ok {
		params.Set("install", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/users/{userId}/entitlements/{entitlementId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId":  c.enterpriseId,
		"userId":        c.userId,
		"entitlementId": c.entitlementId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Adds or updates an entitlement to an app for a user. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "androidenterprise.entitlements.patch",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId",
	//     "entitlementId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "entitlementId": {
	//       "description": "The ID of the entitlement, e.g. \"app:com.google.android.gm\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "install": {
	//       "description": "Set to true to also install the product on all the user's devices where possible. Failure to install on one or more devices will not prevent this operation from returning successfully, as long as the entitlement was successfully assigned to the user.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/entitlements/{entitlementId}",
	//   "request": {
	//     "$ref": "Entitlement"
	//   },
	//   "response": {
	//     "$ref": "Entitlement"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.entitlements.update":

type EntitlementsUpdateCall struct {
	s             *Service
	enterpriseId  string
	userId        string
	entitlementId string
	entitlement   *Entitlement
	opt_          map[string]interface{}
}

// Update: Adds or updates an entitlement to an app for a user.
func (r *EntitlementsService) Update(enterpriseId string, userId string, entitlementId string, entitlement *Entitlement) *EntitlementsUpdateCall {
	c := &EntitlementsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.userId = userId
	c.entitlementId = entitlementId
	c.entitlement = entitlement
	return c
}

// Install sets the optional parameter "install": Set to true to also
// install the product on all the user's devices where possible. Failure
// to install on one or more devices will not prevent this operation
// from returning successfully, as long as the entitlement was
// successfully assigned to the user.
func (c *EntitlementsUpdateCall) Install(install bool) *EntitlementsUpdateCall {
	c.opt_["install"] = install
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EntitlementsUpdateCall) Fields(s ...googleapi.Field) *EntitlementsUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *EntitlementsUpdateCall) IfNoneMatch(entityTag string) *EntitlementsUpdateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.entitlements.update" call.
// Exactly one of the return values is non-nil.
func (c *EntitlementsUpdateCall) Do() (*Entitlement, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.entitlements.update" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *EntitlementsUpdateCall) DoHeader() (ret *Entitlement, resHeader http.Header, err error) {
	var body io.Reader = nil
	body, err = googleapi.WithoutDataWrapper.JSONReader(c.entitlement)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["install"]; ok {
		params.Set("install", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/users/{userId}/entitlements/{entitlementId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId":  c.enterpriseId,
		"userId":        c.userId,
		"entitlementId": c.entitlementId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Adds or updates an entitlement to an app for a user.",
	//   "httpMethod": "PUT",
	//   "id": "androidenterprise.entitlements.update",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId",
	//     "entitlementId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "entitlementId": {
	//       "description": "The ID of the entitlement, e.g. \"app:com.google.android.gm\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "install": {
	//       "description": "Set to true to also install the product on all the user's devices where possible. Failure to install on one or more devices will not prevent this operation from returning successfully, as long as the entitlement was successfully assigned to the user.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/entitlements/{entitlementId}",
	//   "request": {
	//     "$ref": "Entitlement"
	//   },
	//   "response": {
	//     "$ref": "Entitlement"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.grouplicenses.get":

type GrouplicensesGetCall struct {
	s              *Service
	enterpriseId   string
	groupLicenseId string
	opt_           map[string]interface{}
}

// Get: Retrieves details of an enterprise's group license for a
// product.
func (r *GrouplicensesService) Get(enterpriseId string, groupLicenseId string) *GrouplicensesGetCall {
	c := &GrouplicensesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.groupLicenseId = groupLicenseId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GrouplicensesGetCall) Fields(s ...googleapi.Field) *GrouplicensesGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *GrouplicensesGetCall) IfNoneMatch(entityTag string) *GrouplicensesGetCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.grouplicenses.get" call.
// Exactly one of the return values is non-nil.
func (c *GrouplicensesGetCall) Do() (*GroupLicense, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.grouplicenses.get" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *GrouplicensesGetCall) DoHeader() (ret *GroupLicense, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/groupLicenses/{groupLicenseId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId":   c.enterpriseId,
		"groupLicenseId": c.groupLicenseId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Retrieves details of an enterprise's group license for a product.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.grouplicenses.get",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "groupLicenseId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "groupLicenseId": {
	//       "description": "The ID of the product the group license is for, e.g. \"app:com.google.android.gm\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/groupLicenses/{groupLicenseId}",
	//   "response": {
	//     "$ref": "GroupLicense"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.grouplicenses.list":

type GrouplicensesListCall struct {
	s            *Service
	enterpriseId string
	opt_         map[string]interface{}
}

// List: Retrieves IDs of all products for which the enterprise has a
// group license.
func (r *GrouplicensesService) List(enterpriseId string) *GrouplicensesListCall {
	c := &GrouplicensesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GrouplicensesListCall) Fields(s ...googleapi.Field) *GrouplicensesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *GrouplicensesListCall) IfNoneMatch(entityTag string) *GrouplicensesListCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.grouplicenses.list" call.
// Exactly one of the return values is non-nil.
func (c *GrouplicensesListCall) Do() (*GroupLicensesListResponse, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.grouplicenses.list" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *GrouplicensesListCall) DoHeader() (ret *GroupLicensesListResponse, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/groupLicenses")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Retrieves IDs of all products for which the enterprise has a group license.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.grouplicenses.list",
	//   "parameterOrder": [
	//     "enterpriseId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/groupLicenses",
	//   "response": {
	//     "$ref": "GroupLicensesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.grouplicenseusers.list":

type GrouplicenseusersListCall struct {
	s              *Service
	enterpriseId   string
	groupLicenseId string
	opt_           map[string]interface{}
}

// List: Retrieves the IDs of the users who have been granted
// entitlements under the license.
func (r *GrouplicenseusersService) List(enterpriseId string, groupLicenseId string) *GrouplicenseusersListCall {
	c := &GrouplicenseusersListCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.groupLicenseId = groupLicenseId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GrouplicenseusersListCall) Fields(s ...googleapi.Field) *GrouplicenseusersListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *GrouplicenseusersListCall) IfNoneMatch(entityTag string) *GrouplicenseusersListCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.grouplicenseusers.list" call.
// Exactly one of the return values is non-nil.
func (c *GrouplicenseusersListCall) Do() (*GroupLicenseUsersListResponse, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.grouplicenseusers.list" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *GrouplicenseusersListCall) DoHeader() (ret *GroupLicenseUsersListResponse, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/groupLicenses/{groupLicenseId}/users")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId":   c.enterpriseId,
		"groupLicenseId": c.groupLicenseId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Retrieves the IDs of the users who have been granted entitlements under the license.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.grouplicenseusers.list",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "groupLicenseId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "groupLicenseId": {
	//       "description": "The ID of the product the group license is for, e.g. \"app:com.google.android.gm\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/groupLicenses/{groupLicenseId}/users",
	//   "response": {
	//     "$ref": "GroupLicenseUsersListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.installs.delete":

type InstallsDeleteCall struct {
	s            *Service
	enterpriseId string
	userId       string
	deviceId     string
	installId    string
	opt_         map[string]interface{}
}

// Delete: Requests to remove an app from a device. A call to get or
// list will still show the app as installed on the device until it is
// actually removed.
func (r *InstallsService) Delete(enterpriseId string, userId string, deviceId string, installId string) *InstallsDeleteCall {
	c := &InstallsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.userId = userId
	c.deviceId = deviceId
	c.installId = installId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstallsDeleteCall) Fields(s ...googleapi.Field) *InstallsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *InstallsDeleteCall) IfNoneMatch(entityTag string) *InstallsDeleteCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.installs.delete" call.
func (c *InstallsDeleteCall) Do() error {
	_, err := c.DoHeader()
	return err
}

// DoHeader executes the "androidenterprise.installs.delete" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *InstallsDeleteCall) DoHeader() (resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}/installs/{installId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
		"userId":       c.userId,
		"deviceId":     c.deviceId,
		"installId":    c.installId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return res.Header, err
	}
	return res.Header, nil
	// {
	//   "description": "Requests to remove an app from a device. A call to get or list will still show the app as installed on the device until it is actually removed.",
	//   "httpMethod": "DELETE",
	//   "id": "androidenterprise.installs.delete",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId",
	//     "deviceId",
	//     "installId"
	//   ],
	//   "parameters": {
	//     "deviceId": {
	//       "description": "The Android ID of the device.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "installId": {
	//       "description": "The ID of the product represented by the install, e.g. \"app:com.google.android.gm\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}/installs/{installId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.installs.get":

type InstallsGetCall struct {
	s            *Service
	enterpriseId string
	userId       string
	deviceId     string
	installId    string
	opt_         map[string]interface{}
}

// Get: Retrieves details of an installation of an app on a device.
func (r *InstallsService) Get(enterpriseId string, userId string, deviceId string, installId string) *InstallsGetCall {
	c := &InstallsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.userId = userId
	c.deviceId = deviceId
	c.installId = installId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstallsGetCall) Fields(s ...googleapi.Field) *InstallsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *InstallsGetCall) IfNoneMatch(entityTag string) *InstallsGetCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.installs.get" call.
// Exactly one of the return values is non-nil.
func (c *InstallsGetCall) Do() (*Install, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.installs.get" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *InstallsGetCall) DoHeader() (ret *Install, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}/installs/{installId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
		"userId":       c.userId,
		"deviceId":     c.deviceId,
		"installId":    c.installId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Retrieves details of an installation of an app on a device.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.installs.get",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId",
	//     "deviceId",
	//     "installId"
	//   ],
	//   "parameters": {
	//     "deviceId": {
	//       "description": "The Android ID of the device.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "installId": {
	//       "description": "The ID of the product represented by the install, e.g. \"app:com.google.android.gm\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}/installs/{installId}",
	//   "response": {
	//     "$ref": "Install"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.installs.list":

type InstallsListCall struct {
	s            *Service
	enterpriseId string
	userId       string
	deviceId     string
	opt_         map[string]interface{}
}

// List: Retrieves the details of all apps installed on the specified
// device.
func (r *InstallsService) List(enterpriseId string, userId string, deviceId string) *InstallsListCall {
	c := &InstallsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.userId = userId
	c.deviceId = deviceId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstallsListCall) Fields(s ...googleapi.Field) *InstallsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *InstallsListCall) IfNoneMatch(entityTag string) *InstallsListCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.installs.list" call.
// Exactly one of the return values is non-nil.
func (c *InstallsListCall) Do() (*InstallsListResponse, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.installs.list" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *InstallsListCall) DoHeader() (ret *InstallsListResponse, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}/installs")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
		"userId":       c.userId,
		"deviceId":     c.deviceId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Retrieves the details of all apps installed on the specified device.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.installs.list",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId",
	//     "deviceId"
	//   ],
	//   "parameters": {
	//     "deviceId": {
	//       "description": "The Android ID of the device.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}/installs",
	//   "response": {
	//     "$ref": "InstallsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.installs.patch":

type InstallsPatchCall struct {
	s            *Service
	enterpriseId string
	userId       string
	deviceId     string
	installId    string
	install      *Install
	opt_         map[string]interface{}
}

// Patch: Requests to install the latest version of an app to a device.
// If the app is already installed then it is updated to the latest
// version if necessary. This method supports patch semantics.
func (r *InstallsService) Patch(enterpriseId string, userId string, deviceId string, installId string, install *Install) *InstallsPatchCall {
	c := &InstallsPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.userId = userId
	c.deviceId = deviceId
	c.installId = installId
	c.install = install
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstallsPatchCall) Fields(s ...googleapi.Field) *InstallsPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *InstallsPatchCall) IfNoneMatch(entityTag string) *InstallsPatchCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.installs.patch" call.
// Exactly one of the return values is non-nil.
func (c *InstallsPatchCall) Do() (*Install, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.installs.patch" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *InstallsPatchCall) DoHeader() (ret *Install, resHeader http.Header, err error) {
	var body io.Reader = nil
	body, err = googleapi.WithoutDataWrapper.JSONReader(c.install)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}/installs/{installId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
		"userId":       c.userId,
		"deviceId":     c.deviceId,
		"installId":    c.installId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Requests to install the latest version of an app to a device. If the app is already installed then it is updated to the latest version if necessary. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "androidenterprise.installs.patch",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId",
	//     "deviceId",
	//     "installId"
	//   ],
	//   "parameters": {
	//     "deviceId": {
	//       "description": "The Android ID of the device.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "installId": {
	//       "description": "The ID of the product represented by the install, e.g. \"app:com.google.android.gm\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}/installs/{installId}",
	//   "request": {
	//     "$ref": "Install"
	//   },
	//   "response": {
	//     "$ref": "Install"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.installs.update":

type InstallsUpdateCall struct {
	s            *Service
	enterpriseId string
	userId       string
	deviceId     string
	installId    string
	install      *Install
	opt_         map[string]interface{}
}

// Update: Requests to install the latest version of an app to a device.
// If the app is already installed then it is updated to the latest
// version if necessary.
func (r *InstallsService) Update(enterpriseId string, userId string, deviceId string, installId string, install *Install) *InstallsUpdateCall {
	c := &InstallsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.userId = userId
	c.deviceId = deviceId
	c.installId = installId
	c.install = install
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstallsUpdateCall) Fields(s ...googleapi.Field) *InstallsUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *InstallsUpdateCall) IfNoneMatch(entityTag string) *InstallsUpdateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.installs.update" call.
// Exactly one of the return values is non-nil.
func (c *InstallsUpdateCall) Do() (*Install, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.installs.update" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *InstallsUpdateCall) DoHeader() (ret *Install, resHeader http.Header, err error) {
	var body io.Reader = nil
	body, err = googleapi.WithoutDataWrapper.JSONReader(c.install)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}/installs/{installId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
		"userId":       c.userId,
		"deviceId":     c.deviceId,
		"installId":    c.installId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Requests to install the latest version of an app to a device. If the app is already installed then it is updated to the latest version if necessary.",
	//   "httpMethod": "PUT",
	//   "id": "androidenterprise.installs.update",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId",
	//     "deviceId",
	//     "installId"
	//   ],
	//   "parameters": {
	//     "deviceId": {
	//       "description": "The Android ID of the device.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "installId": {
	//       "description": "The ID of the product represented by the install, e.g. \"app:com.google.android.gm\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}/installs/{installId}",
	//   "request": {
	//     "$ref": "Install"
	//   },
	//   "response": {
	//     "$ref": "Install"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.permissions.get":

type PermissionsGetCall struct {
	s            *Service
	permissionId string
	opt_         map[string]interface{}
}

// Get: Retrieves details of an Android app permission for display to an
// enterprise admin.
func (r *PermissionsService) Get(permissionId string) *PermissionsGetCall {
	c := &PermissionsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.permissionId = permissionId
	return c
}

// Language sets the optional parameter "language": The BCP47 tag for
// the user's preferred language (e.g. "en-US", "de")
func (c *PermissionsGetCall) Language(language string) *PermissionsGetCall {
	c.opt_["language"] = language
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PermissionsGetCall) Fields(s ...googleapi.Field) *PermissionsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *PermissionsGetCall) IfNoneMatch(entityTag string) *PermissionsGetCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.permissions.get" call.
// Exactly one of the return values is non-nil.
func (c *PermissionsGetCall) Do() (*Permission, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.permissions.get" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *PermissionsGetCall) DoHeader() (ret *Permission, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["language"]; ok {
		params.Set("language", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "permissions/{permissionId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"permissionId": c.permissionId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Retrieves details of an Android app permission for display to an enterprise admin.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.permissions.get",
	//   "parameterOrder": [
	//     "permissionId"
	//   ],
	//   "parameters": {
	//     "language": {
	//       "description": "The BCP47 tag for the user's preferred language (e.g. \"en-US\", \"de\")",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "permissionId": {
	//       "description": "The ID of the permission.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "permissions/{permissionId}",
	//   "response": {
	//     "$ref": "Permission"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.products.approve":

type ProductsApproveCall struct {
	s                      *Service
	enterpriseId           string
	productId              string
	productsapproverequest *ProductsApproveRequest
	opt_                   map[string]interface{}
}

// Approve: Approves the specified product (and the relevant app
// permissions, if any).
func (r *ProductsService) Approve(enterpriseId string, productId string, productsapproverequest *ProductsApproveRequest) *ProductsApproveCall {
	c := &ProductsApproveCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.productId = productId
	c.productsapproverequest = productsapproverequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProductsApproveCall) Fields(s ...googleapi.Field) *ProductsApproveCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *ProductsApproveCall) IfNoneMatch(entityTag string) *ProductsApproveCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.products.approve" call.
func (c *ProductsApproveCall) Do() error {
	_, err := c.DoHeader()
	return err
}

// DoHeader executes the "androidenterprise.products.approve" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *ProductsApproveCall) DoHeader() (resHeader http.Header, err error) {
	var body io.Reader = nil
	body, err = googleapi.WithoutDataWrapper.JSONReader(c.productsapproverequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/products/{productId}/approve")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
		"productId":    c.productId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return res.Header, err
	}
	return res.Header, nil
	// {
	//   "description": "Approves the specified product (and the relevant app permissions, if any).",
	//   "httpMethod": "POST",
	//   "id": "androidenterprise.products.approve",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "productId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "productId": {
	//       "description": "The ID of the product.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/products/{productId}/approve",
	//   "request": {
	//     "$ref": "ProductsApproveRequest"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.products.generateApprovalUrl":

type ProductsGenerateApprovalUrlCall struct {
	s            *Service
	enterpriseId string
	productId    string
	opt_         map[string]interface{}
}

// GenerateApprovalUrl: Generates a URL that can be used to display an
// iframe to view the product's permissions (if any) and approve the
// product. This URL can be used to approve the product for a limited
// time (currently 1 hour) using the Products.approve call.
func (r *ProductsService) GenerateApprovalUrl(enterpriseId string, productId string) *ProductsGenerateApprovalUrlCall {
	c := &ProductsGenerateApprovalUrlCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.productId = productId
	return c
}

// LanguageCode sets the optional parameter "languageCode": The language
// code that will be used for permission names and descriptions in the
// returned iframe.
func (c *ProductsGenerateApprovalUrlCall) LanguageCode(languageCode string) *ProductsGenerateApprovalUrlCall {
	c.opt_["languageCode"] = languageCode
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProductsGenerateApprovalUrlCall) Fields(s ...googleapi.Field) *ProductsGenerateApprovalUrlCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *ProductsGenerateApprovalUrlCall) IfNoneMatch(entityTag string) *ProductsGenerateApprovalUrlCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.products.generateApprovalUrl" call.
// Exactly one of the return values is non-nil.
func (c *ProductsGenerateApprovalUrlCall) Do() (*ProductsGenerateApprovalUrlResponse, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.products.generateApprovalUrl" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *ProductsGenerateApprovalUrlCall) DoHeader() (ret *ProductsGenerateApprovalUrlResponse, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["languageCode"]; ok {
		params.Set("languageCode", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/products/{productId}/generateApprovalUrl")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
		"productId":    c.productId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Generates a URL that can be used to display an iframe to view the product's permissions (if any) and approve the product. This URL can be used to approve the product for a limited time (currently 1 hour) using the Products.approve call.",
	//   "httpMethod": "POST",
	//   "id": "androidenterprise.products.generateApprovalUrl",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "productId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "languageCode": {
	//       "description": "The language code that will be used for permission names and descriptions in the returned iframe.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "productId": {
	//       "description": "The ID of the product.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/products/{productId}/generateApprovalUrl",
	//   "response": {
	//     "$ref": "ProductsGenerateApprovalUrlResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.products.get":

type ProductsGetCall struct {
	s            *Service
	enterpriseId string
	productId    string
	opt_         map[string]interface{}
}

// Get: Retrieves details of a product for display to an enterprise
// admin.
func (r *ProductsService) Get(enterpriseId string, productId string) *ProductsGetCall {
	c := &ProductsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.productId = productId
	return c
}

// Language sets the optional parameter "language": The BCP47 tag for
// the user's preferred language (e.g. "en-US", "de").
func (c *ProductsGetCall) Language(language string) *ProductsGetCall {
	c.opt_["language"] = language
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProductsGetCall) Fields(s ...googleapi.Field) *ProductsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *ProductsGetCall) IfNoneMatch(entityTag string) *ProductsGetCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.products.get" call.
// Exactly one of the return values is non-nil.
func (c *ProductsGetCall) Do() (*Product, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.products.get" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *ProductsGetCall) DoHeader() (ret *Product, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["language"]; ok {
		params.Set("language", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/products/{productId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
		"productId":    c.productId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Retrieves details of a product for display to an enterprise admin.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.products.get",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "productId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "language": {
	//       "description": "The BCP47 tag for the user's preferred language (e.g. \"en-US\", \"de\").",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "productId": {
	//       "description": "The ID of the product, e.g. \"app:com.google.android.gm\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/products/{productId}",
	//   "response": {
	//     "$ref": "Product"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.products.getAppRestrictionsSchema":

type ProductsGetAppRestrictionsSchemaCall struct {
	s            *Service
	enterpriseId string
	productId    string
	opt_         map[string]interface{}
}

// GetAppRestrictionsSchema: Retrieves the schema defining app
// restrictions configurable for this product. All products have a
// schema, but this may be empty if no app restrictions are defined.
func (r *ProductsService) GetAppRestrictionsSchema(enterpriseId string, productId string) *ProductsGetAppRestrictionsSchemaCall {
	c := &ProductsGetAppRestrictionsSchemaCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.productId = productId
	return c
}

// Language sets the optional parameter "language": The BCP47 tag for
// the user's preferred language (e.g. "en-US", "de").
func (c *ProductsGetAppRestrictionsSchemaCall) Language(language string) *ProductsGetAppRestrictionsSchemaCall {
	c.opt_["language"] = language
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProductsGetAppRestrictionsSchemaCall) Fields(s ...googleapi.Field) *ProductsGetAppRestrictionsSchemaCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *ProductsGetAppRestrictionsSchemaCall) IfNoneMatch(entityTag string) *ProductsGetAppRestrictionsSchemaCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.products.getAppRestrictionsSchema" call.
// Exactly one of the return values is non-nil.
func (c *ProductsGetAppRestrictionsSchemaCall) Do() (*AppRestrictionsSchema, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.products.getAppRestrictionsSchema" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *ProductsGetAppRestrictionsSchemaCall) DoHeader() (ret *AppRestrictionsSchema, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["language"]; ok {
		params.Set("language", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/products/{productId}/appRestrictionsSchema")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
		"productId":    c.productId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Retrieves the schema defining app restrictions configurable for this product. All products have a schema, but this may be empty if no app restrictions are defined.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.products.getAppRestrictionsSchema",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "productId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "language": {
	//       "description": "The BCP47 tag for the user's preferred language (e.g. \"en-US\", \"de\").",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "productId": {
	//       "description": "The ID of the product.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/products/{productId}/appRestrictionsSchema",
	//   "response": {
	//     "$ref": "AppRestrictionsSchema"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.products.getPermissions":

type ProductsGetPermissionsCall struct {
	s            *Service
	enterpriseId string
	productId    string
	opt_         map[string]interface{}
}

// GetPermissions: Retrieves the Android app permissions required by
// this app.
func (r *ProductsService) GetPermissions(enterpriseId string, productId string) *ProductsGetPermissionsCall {
	c := &ProductsGetPermissionsCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.productId = productId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProductsGetPermissionsCall) Fields(s ...googleapi.Field) *ProductsGetPermissionsCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *ProductsGetPermissionsCall) IfNoneMatch(entityTag string) *ProductsGetPermissionsCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.products.getPermissions" call.
// Exactly one of the return values is non-nil.
func (c *ProductsGetPermissionsCall) Do() (*ProductPermissions, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.products.getPermissions" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *ProductsGetPermissionsCall) DoHeader() (ret *ProductPermissions, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/products/{productId}/permissions")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
		"productId":    c.productId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Retrieves the Android app permissions required by this app.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.products.getPermissions",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "productId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "productId": {
	//       "description": "The ID of the product.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/products/{productId}/permissions",
	//   "response": {
	//     "$ref": "ProductPermissions"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.products.updatePermissions":

type ProductsUpdatePermissionsCall struct {
	s                  *Service
	enterpriseId       string
	productId          string
	productpermissions *ProductPermissions
	opt_               map[string]interface{}
}

// UpdatePermissions: Updates the set of Android app permissions for
// this app that have been accepted by the enterprise.
func (r *ProductsService) UpdatePermissions(enterpriseId string, productId string, productpermissions *ProductPermissions) *ProductsUpdatePermissionsCall {
	c := &ProductsUpdatePermissionsCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.productId = productId
	c.productpermissions = productpermissions
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProductsUpdatePermissionsCall) Fields(s ...googleapi.Field) *ProductsUpdatePermissionsCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *ProductsUpdatePermissionsCall) IfNoneMatch(entityTag string) *ProductsUpdatePermissionsCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.products.updatePermissions" call.
// Exactly one of the return values is non-nil.
func (c *ProductsUpdatePermissionsCall) Do() (*ProductPermissions, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.products.updatePermissions" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *ProductsUpdatePermissionsCall) DoHeader() (ret *ProductPermissions, resHeader http.Header, err error) {
	var body io.Reader = nil
	body, err = googleapi.WithoutDataWrapper.JSONReader(c.productpermissions)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/products/{productId}/permissions")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
		"productId":    c.productId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Updates the set of Android app permissions for this app that have been accepted by the enterprise.",
	//   "httpMethod": "PUT",
	//   "id": "androidenterprise.products.updatePermissions",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "productId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "productId": {
	//       "description": "The ID of the product.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/products/{productId}/permissions",
	//   "request": {
	//     "$ref": "ProductPermissions"
	//   },
	//   "response": {
	//     "$ref": "ProductPermissions"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.users.generateToken":

type UsersGenerateTokenCall struct {
	s            *Service
	enterpriseId string
	userId       string
	opt_         map[string]interface{}
}

// GenerateToken: Generates a token (activation code) to allow this user
// to configure their work account in the Android Setup Wizard. Revokes
// any previously generated token.
func (r *UsersService) GenerateToken(enterpriseId string, userId string) *UsersGenerateTokenCall {
	c := &UsersGenerateTokenCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.userId = userId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersGenerateTokenCall) Fields(s ...googleapi.Field) *UsersGenerateTokenCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *UsersGenerateTokenCall) IfNoneMatch(entityTag string) *UsersGenerateTokenCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.users.generateToken" call.
// Exactly one of the return values is non-nil.
func (c *UsersGenerateTokenCall) Do() (*UserToken, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.users.generateToken" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *UsersGenerateTokenCall) DoHeader() (ret *UserToken, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/users/{userId}/token")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
		"userId":       c.userId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Generates a token (activation code) to allow this user to configure their work account in the Android Setup Wizard. Revokes any previously generated token.",
	//   "httpMethod": "POST",
	//   "id": "androidenterprise.users.generateToken",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/token",
	//   "response": {
	//     "$ref": "UserToken"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.users.get":

type UsersGetCall struct {
	s            *Service
	enterpriseId string
	userId       string
	opt_         map[string]interface{}
}

// Get: Retrieves a user's details.
func (r *UsersService) Get(enterpriseId string, userId string) *UsersGetCall {
	c := &UsersGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.userId = userId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersGetCall) Fields(s ...googleapi.Field) *UsersGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *UsersGetCall) IfNoneMatch(entityTag string) *UsersGetCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.users.get" call.
// Exactly one of the return values is non-nil.
func (c *UsersGetCall) Do() (*User, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.users.get" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *UsersGetCall) DoHeader() (ret *User, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/users/{userId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
		"userId":       c.userId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Retrieves a user's details.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.users.get",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}",
	//   "response": {
	//     "$ref": "User"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.users.list":

type UsersListCall struct {
	s            *Service
	enterpriseId string
	email        string
	opt_         map[string]interface{}
}

// List: Looks up a user by email address.
func (r *UsersService) List(enterpriseId string, email string) *UsersListCall {
	c := &UsersListCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.email = email
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersListCall) Fields(s ...googleapi.Field) *UsersListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *UsersListCall) IfNoneMatch(entityTag string) *UsersListCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.users.list" call.
// Exactly one of the return values is non-nil.
func (c *UsersListCall) Do() (*UsersListResponse, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "androidenterprise.users.list" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *UsersListCall) DoHeader() (ret *UsersListResponse, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("email", fmt.Sprintf("%v", c.email))
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/users")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Looks up a user by email address.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.users.list",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "email"
	//   ],
	//   "parameters": {
	//     "email": {
	//       "description": "The exact primary email address of the user to look up.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users",
	//   "response": {
	//     "$ref": "UsersListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.users.revokeToken":

type UsersRevokeTokenCall struct {
	s            *Service
	enterpriseId string
	userId       string
	opt_         map[string]interface{}
}

// RevokeToken: Revokes a previously generated token (activation code)
// for the user.
func (r *UsersService) RevokeToken(enterpriseId string, userId string) *UsersRevokeTokenCall {
	c := &UsersRevokeTokenCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.userId = userId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersRevokeTokenCall) Fields(s ...googleapi.Field) *UsersRevokeTokenCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *UsersRevokeTokenCall) IfNoneMatch(entityTag string) *UsersRevokeTokenCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "androidenterprise.users.revokeToken" call.
func (c *UsersRevokeTokenCall) Do() error {
	_, err := c.DoHeader()
	return err
}

// DoHeader executes the "androidenterprise.users.revokeToken" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *UsersRevokeTokenCall) DoHeader() (resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "enterprises/{enterpriseId}/users/{userId}/token")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"enterpriseId": c.enterpriseId,
		"userId":       c.userId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return res.Header, err
	}
	return res.Header, nil
	// {
	//   "description": "Revokes a previously generated token (activation code) for the user.",
	//   "httpMethod": "DELETE",
	//   "id": "androidenterprise.users.revokeToken",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/token",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}
