// Package androidenterprise provides access to the Google Play EMM API.
//
// See https://developers.google.com/play/enterprise
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

// AppRestrictionsSchema: Represents the list of app restrictions
// available to be pre-configured for the product.
type AppRestrictionsSchema struct {
	// Restrictions: The set of restrictions that make up this schema.
	Restrictions []*AppRestrictionsSchemaRestriction `json:"restrictions,omitempty"`
}

// AppRestrictionsSchemaRestriction: A restriction in the App
// Restriction Schema represents a piece of configuration that may be
// pre-applied.
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

// AppRestrictionsSchemaRestrictionRestrictionValue: A typed value for
// the restriction.
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

// AppVersion: This represents a single version of the app.
type AppVersion struct {
	// VersionCode: Unique increasing identifier for the app version.
	VersionCode int64 `json:"versionCode,omitempty"`

	// VersionString: The string used in the Play Store by the app developer
	// to identify the version. The string is not necessarily unique or
	// localized (for example, the string could be "1.4").
	VersionString string `json:"versionString,omitempty"`
}

// ApprovalUrlInfo: Information on an approval URL.
type ApprovalUrlInfo struct {
	// ApprovalUrl: A URL that displays a product's permissions and that can
	// also be used to approve the product with the Products.approve call.
	ApprovalUrl string `json:"approvalUrl,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#approvalUrlInfo".
	Kind string `json:"kind,omitempty"`
}

// Collection: A collection resource defines a named set of apps that is
// visible to a set of users in the Google Play Store app running on
// those users' managed devices. Those users can then install any of
// those apps if they wish (which will trigger creation of install and
// entitlement resources). A user cannot install an app on a managed
// device unless the app is listed in at least one collection that is
// visible to that user.
//
// Note that the API can be used to directly install an app regardless
// of whether it is in any collection - so an enterprise has a choice of
// either directly pushing apps to users, or allowing users to install
// apps if they want. Which is appropriate will depend on the
// enterprise's policies and the purpose of the apps concerned.
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

// CollectionViewersListResponse: The user resources for the collection.
type CollectionViewersListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#collectionViewersListResponse".
	Kind string `json:"kind,omitempty"`

	// User: A user of an enterprise.
	User []*User `json:"user,omitempty"`
}

// CollectionsListResponse: The collection resources for the enterprise.
type CollectionsListResponse struct {
	// Collection: An ordered collection of products which can be made
	// visible on the Google Play Store app to a selected group of users.
	Collection []*Collection `json:"collection,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#collectionsListResponse".
	Kind string `json:"kind,omitempty"`
}

// Device: A device resource represents a mobile device managed by the
// MDM and belonging to a specific enterprise user.
//
// This collection cannot be modified via the API; it is automatically
// populated as devices are set up to be managed.
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

// DeviceState: The state of a user's device, as accessed by the
// getState and setState methods on device resources.
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

// DevicesListResponse: The device resources for the user.
type DevicesListResponse struct {
	// Device: A managed device.
	Device []*Device `json:"device,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#devicesListResponse".
	Kind string `json:"kind,omitempty"`
}

// Enterprise: An enterprise resource represents a binding between an
// organisation and their MDM.
//
// To create an enterprise, an admin of the enterprise must first go
// through a Play for Work sign-up flow. At the end of this the admin
// will be presented with a token (a short opaque alphanumeric string).
// They must then present this to the MDM, who then supplies it to the
// enroll method. Until this is done the MDM will not have any access to
// the enterprise.
//
// After calling enroll the MDM should call setAccount to specify the
// service account that will be allowed to act on behalf of the
// enterprise, which will be required for access to the enterprise's
// data through this API. Only one call of setAccount is allowed for a
// given enterprise; the only way to change the account later is to
// unenroll the enterprise and enroll it again (obtaining a new
// token).
//
// The MDM can unenroll an enterprise in order to sever the binding
// between them. Re-enrolling an enterprise is possible, but requires a
// new token to be retrieved. Enterprises.unenroll requires the MDM's
// credentials (as enroll does), not the enterprise's.
// Enterprises.unenroll can only be used for enterprises that were
// previously enrolled with the enroll call. Any enterprises that were
// enrolled using the (deprecated) Enterprises.insert call must be
// unenrolled with Enterprises.delete and can then be re-enrolled using
// the Enterprises.enroll call.
//
// The ID for an enterprise is an opaque string. It is returned by
// insert and enroll and can also be retrieved if the enterprise's
// primary domain is known using the list method.
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

// EnterpriseAccount: A service account that can be used to authenticate
// as the enterprise to API calls that require such authentication.
type EnterpriseAccount struct {
	// AccountEmail: The email address of the service account.
	AccountEmail string `json:"accountEmail,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#enterpriseAccount".
	Kind string `json:"kind,omitempty"`
}

// EnterprisesListResponse: The matching enterprise resources.
type EnterprisesListResponse struct {
	// Enterprise: An enterprise.
	Enterprise []*Enterprise `json:"enterprise,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#enterprisesListResponse".
	Kind string `json:"kind,omitempty"`
}

// Entitlement: The existence of an entitlement resource means that a
// user has the right to use a particular app on any of their devices.
// This might be because the app is free or because they have been
// allocated a license to the app from a group license purchased by the
// enterprise.
//
// It should always be true that a user has an app installed on one of
// their devices only if they have an entitlement to it. So if an
// entitlement is deleted, the app will be uninstalled from all devices.
// Similarly if the user installs an app (and is permitted to do so), or
// the MDM triggers an install of the app, an entitlement to that app is
// automatically created. If this is impossible - e.g. the enterprise
// has not purchased sufficient licenses - then installation
// fails.
//
// Note that entitlements are always user specific, not device specific;
// a user may have an entitlement even though they have not installed
// the app anywhere. Once they have an entitlement they can install the
// app on multiple devices.
//
// The API can be used to create an entitlement. If the app is a free
// app, a group license for that app is created. If it's a paid app,
// creating the entitlement consumes one license; it remains consumed
// until the entitlement is removed. Optionally an installation of the
// app on all the user's managed devices can be triggered at the time
// the entitlement is created. An entitlement cannot be created for an
// app if the app requires permissions that the enterprise has not yet
// accepted.
//
// Entitlements for paid apps that are due to purchases by the user on a
// non-managed profile will have "userPurchase" as entitlement reason;
// those entitlements cannot be removed via the API.
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

// EntitlementsListResponse: The entitlement resources for the user.
type EntitlementsListResponse struct {
	// Entitlement: An entitlement of a user to a product (e.g. an app). For
	// example, a free app that they have installed, or a paid app that they
	// have been allocated a license to.
	Entitlement []*Entitlement `json:"entitlement,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#entitlementsListResponse".
	Kind string `json:"kind,omitempty"`
}

// GroupLicense: A group license object indicates a product that an
// enterprise admin has approved for use in the enterprise. The product
// may be free or paid. For free products, a group license object is
// created in these cases: if the enterprise admin approves a product in
// Google Play, if the product is added to a collection, or if an
// entitlement for the product is created for a user via the API. For
// paid products, a group license object is only created as part of the
// first bulk purchase of that product in Google Play by the enterprise
// admin.
//
// The API can be used to query group licenses; the available
// information includes the total number of licenses purchased (for paid
// products) and the total number of licenses that have been
// provisioned, that is, the total number of user entitlements in
// existence for the product.
//
// Group license objects are never deleted. If, for example, a free app
// is added to a collection and then removed, the group license will
// remain, allowing the enterprise admin to keep track of any remaining
// entitlements. An enterprise admin may indicate they are no longer
// interested in the group license by marking it as unapproved in Google
// Play.
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

// GroupLicenseUsersListResponse: The user resources for the group
// license.
type GroupLicenseUsersListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#groupLicenseUsersListResponse".
	Kind string `json:"kind,omitempty"`

	// User: A user of an enterprise.
	User []*User `json:"user,omitempty"`
}

// GroupLicensesListResponse: The grouplicense resources for the
// enterprise.
type GroupLicensesListResponse struct {
	// GroupLicense: A group license for a product approved for use in the
	// enterprise.
	GroupLicense []*GroupLicense `json:"groupLicense,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#groupLicensesListResponse".
	Kind string `json:"kind,omitempty"`
}

// Install: The existence of an install resource indicates that an app
// is installed on a particular device (or that an install is
// pending).
//
// The API can be used to create an install resource using the update
// method. This triggers the actual install of the app on the device. If
// the user does not already have an entitlement for the app then an
// attempt is made to create one. If this fails (e.g. because the app is
// not free and there is no available license) then the creation of the
// install fails.
//
// The API can also be used to update an installed app. If the update
// method is used on an existing install then the app will be updated to
// the latest available version.
//
// Note that it is not possible to force the installation of a specific
// version of an app; the version code is read-only.
//
// If a user installs an app themselves (as permitted by the
// enterprise), then again an install resource and possibly an
// entitlement resource are automatically created.
//
// The API can also be used to delete an install resource, which
// triggers the removal of the app from the device. Note that deleting
// an install does not automatically remove the corresponding
// entitlement, even if there are no remaining installs. The install
// resource will also be deleted if the user uninstalls the app
// themselves.
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

// InstallsListResponse: The install resources for the device.
type InstallsListResponse struct {
	// Install: An installation of an app for a user on a specific device.
	// The existence of an install implies that the user must have an
	// entitlement to the app.
	Install []*Install `json:"install,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#installsListResponse".
	Kind string `json:"kind,omitempty"`
}

// Permission: A permission represents some extra capability, to be
// granted to an Android app, which requires explicit consent. An
// enterprise admin must consent to these permissions on behalf of their
// users before an entitlement for the app can be created.
//
// The permissions collection is read-only. The information provided for
// each permission (localized name and description) is intended to be
// used in the MDM user interface when obtaining consent from the
// enterprise.
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

// Product: A product represents an app in the Google Play Store that is
// available to at least some users in the enterprise. (Some apps are
// restricted to a single enterprise, and no information about them is
// made available outside that enterprise.)
//
// The information provided for each product (localized name, icon, link
// to the full Google Play details page) is intended to allow a basic
// representation of the product within an MDM user interface.
type Product struct {
	// AppVersion: App versions currently available for this product. The
	// returned list contains only public versions. Alpha and beta versions
	// are not included.
	AppVersion []*AppVersion `json:"appVersion,omitempty"`

	// AuthorName: The name of the author of the product (e.g. the app
	// developer).
	AuthorName string `json:"authorName,omitempty"`

	// DetailsUrl: A link to the (consumer) Google Play details page for the
	// product.
	DetailsUrl string `json:"detailsUrl,omitempty"`

	// DistributionChannel: How and to whom the package is made available.
	// The value publicGoogleHosted means that the package is available
	// through the Play Store and not restricted to a specific enterprise.
	// The value privateGoogleHosted means that the package is a private app
	// (restricted to an enterprise) but hosted by Google. The value
	// privateSelfHosted means that the package is a private app (restricted
	// to an enterprise) and is privately hosted.
	DistributionChannel string `json:"distributionChannel,omitempty"`

	// IconUrl: A link to an image that can be used as an icon for the
	// product.
	IconUrl string `json:"iconUrl,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#product".
	Kind string `json:"kind,omitempty"`

	// ProductId: A string of the form app:
	// . For example, app:com.google.android.gm represents the Gmail app.
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

// ProductPermission: A product permissions resource represents the set
// of permissions required by a specific app and whether or not they
// have been accepted by an enterprise admin.
//
// The API can be used to read the set of permissions, and also to
// update the set to indicate that permissions have been accepted.
type ProductPermission struct {
	// PermissionId: An opaque string uniquely identifying the permission.
	PermissionId string `json:"permissionId,omitempty"`

	// State: Whether the permission has been accepted or not.
	State string `json:"state,omitempty"`
}

// ProductPermissions: Information about the permissions required by a
// specific app and whether they have been accepted by the enterprise.
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
	// ApprovalUrlInfo: The approval URL that was shown to the user. Only
	// the permissions shown to the user with that URL will be accepted,
	// which may not be the product's entire set of permissions. For
	// example, the URL may only display new permissions from an update
	// after the product was approved, or not include new permissions if the
	// product was updated since the URL was generated.
	ApprovalUrlInfo *ApprovalUrlInfo `json:"approvalUrlInfo,omitempty"`
}

type ProductsGenerateApprovalUrlResponse struct {
	// Url: A URL that can be rendered in an iframe to display the
	// permissions (if any) of a product. This URL can be used to approve
	// the product only once and only within 24 hours of being generated,
	// using the Products.approve call. If the product is currently
	// unapproved and has no permissions, this URL will point to an empty
	// page. If the product is currently approved, a URL will only be
	// generated if that product has added permissions since it was last
	// approved, and the URL will only display those new permissions that
	// have not yet been accepted.
	Url string `json:"url,omitempty"`
}

// User: A user resource represents an individual user within the
// enterprise's domain.
//
// Note that each user is associated with a Google account based on the
// user's corporate email address (which must be in one of the
// enterprise's domains). As part of installing an MDM app to manage a
// device the Google account must be provisioned to the device, and so
// the user resource must be created before that. This can be done using
// the Google Admin SDK Directory API.
//
// The ID for a user is an opaque string. It can be retrieved using the
// list method queried by the user's primary email address.
type User struct {
	// Id: The unique ID for the user.
	Id string `json:"id,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#user".
	Kind string `json:"kind,omitempty"`

	// PrimaryEmail: The user's primary email, e.g. "jsmith@example.com".
	PrimaryEmail string `json:"primaryEmail,omitempty"`
}

// UserToken: A UserToken is used by a user when setting up a managed
// device or profile with their work account on a device. When the user
// enters their email address and token (activation code) the
// appropriate MDM app can be automatically downloaded.
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

// UsersListResponse: The matching user resources.
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

func (c *CollectionsDeleteCall) Do() error {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
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

func (c *CollectionsGetCall) Do() (*Collection, error) {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Collection
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *CollectionsInsertCall) Do() (*Collection, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.collection)
	if err != nil {
		return nil, err
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Collection
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *CollectionsListCall) Do() (*CollectionsListResponse, error) {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *CollectionsListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *CollectionsPatchCall) Do() (*Collection, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.collection)
	if err != nil {
		return nil, err
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Collection
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *CollectionsUpdateCall) Do() (*Collection, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.collection)
	if err != nil {
		return nil, err
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Collection
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *CollectionviewersDeleteCall) Do() error {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
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

func (c *CollectionviewersGetCall) Do() (*User, error) {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *User
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *CollectionviewersListCall) Do() (*CollectionViewersListResponse, error) {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *CollectionViewersListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *CollectionviewersPatchCall) Do() (*User, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.user)
	if err != nil {
		return nil, err
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *User
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *CollectionviewersUpdateCall) Do() (*User, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.user)
	if err != nil {
		return nil, err
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *User
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *DevicesGetCall) Do() (*Device, error) {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Device
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *DevicesGetStateCall) Do() (*DeviceState, error) {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *DeviceState
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *DevicesListCall) Do() (*DevicesListResponse, error) {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *DevicesListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *DevicesSetStateCall) Do() (*DeviceState, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.devicestate)
	if err != nil {
		return nil, err
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *DeviceState
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *EnterprisesDeleteCall) Do() error {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
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

func (c *EnterprisesEnrollCall) Do() (*Enterprise, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.enterprise)
	if err != nil {
		return nil, err
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Enterprise
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *EnterprisesGetCall) Do() (*Enterprise, error) {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Enterprise
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *EnterprisesInsertCall) Do() (*Enterprise, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.enterprise)
	if err != nil {
		return nil, err
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Enterprise
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *EnterprisesListCall) Do() (*EnterprisesListResponse, error) {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *EnterprisesListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *EnterprisesSetAccountCall) Do() (*EnterpriseAccount, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.enterpriseaccount)
	if err != nil {
		return nil, err
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *EnterpriseAccount
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *EnterprisesUnenrollCall) Do() error {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
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

func (c *EntitlementsDeleteCall) Do() error {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
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

func (c *EntitlementsGetCall) Do() (*Entitlement, error) {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Entitlement
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *EntitlementsListCall) Do() (*EntitlementsListResponse, error) {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *EntitlementsListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *EntitlementsPatchCall) Do() (*Entitlement, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.entitlement)
	if err != nil {
		return nil, err
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Entitlement
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *EntitlementsUpdateCall) Do() (*Entitlement, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.entitlement)
	if err != nil {
		return nil, err
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Entitlement
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *GrouplicensesGetCall) Do() (*GroupLicense, error) {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *GroupLicense
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *GrouplicensesListCall) Do() (*GroupLicensesListResponse, error) {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *GroupLicensesListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *GrouplicenseusersListCall) Do() (*GroupLicenseUsersListResponse, error) {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *GroupLicenseUsersListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *InstallsDeleteCall) Do() error {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
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

func (c *InstallsGetCall) Do() (*Install, error) {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Install
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *InstallsListCall) Do() (*InstallsListResponse, error) {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *InstallsListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *InstallsPatchCall) Do() (*Install, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.install)
	if err != nil {
		return nil, err
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Install
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *InstallsUpdateCall) Do() (*Install, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.install)
	if err != nil {
		return nil, err
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Install
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *PermissionsGetCall) Do() (*Permission, error) {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Permission
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *ProductsApproveCall) Do() error {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.productsapproverequest)
	if err != nil {
		return err
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
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

// GenerateApprovalUrl: Generates a URL that can be rendered in an
// iframe to display the permissions (if any) of a product. An
// enterprise admin must view these permissions and accept them on
// behalf of their organization in order to approve that
// product.
//
// Admins should accept the displayed permissions by interacting with a
// separate UI element in the EMM console, which in turn should trigger
// the use of this URL as the approvalUrlInfo.approvalUrl property in a
// Products.approve call to approve the product. This URL can only be
// used to display permissions for up to 1 day.
func (r *ProductsService) GenerateApprovalUrl(enterpriseId string, productId string) *ProductsGenerateApprovalUrlCall {
	c := &ProductsGenerateApprovalUrlCall{s: r.s, opt_: make(map[string]interface{})}
	c.enterpriseId = enterpriseId
	c.productId = productId
	return c
}

// LanguageCode sets the optional parameter "languageCode": The BCP 47
// language code used for permission names and descriptions in the
// returned iframe, for instance "en-US".
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

func (c *ProductsGenerateApprovalUrlCall) Do() (*ProductsGenerateApprovalUrlResponse, error) {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ProductsGenerateApprovalUrlResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Generates a URL that can be rendered in an iframe to display the permissions (if any) of a product. An enterprise admin must view these permissions and accept them on behalf of their organization in order to approve that product.\n\nAdmins should accept the displayed permissions by interacting with a separate UI element in the EMM console, which in turn should trigger the use of this URL as the approvalUrlInfo.approvalUrl property in a Products.approve call to approve the product. This URL can only be used to display permissions for up to 1 day.",
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
	//       "description": "The BCP 47 language code used for permission names and descriptions in the returned iframe, for instance \"en-US\".",
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

func (c *ProductsGetCall) Do() (*Product, error) {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Product
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *ProductsGetAppRestrictionsSchemaCall) Do() (*AppRestrictionsSchema, error) {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *AppRestrictionsSchema
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *ProductsGetPermissionsCall) Do() (*ProductPermissions, error) {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ProductPermissions
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *ProductsUpdatePermissionsCall) Do() (*ProductPermissions, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.productpermissions)
	if err != nil {
		return nil, err
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ProductPermissions
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *UsersGenerateTokenCall) Do() (*UserToken, error) {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *UserToken
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *UsersGetCall) Do() (*User, error) {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *User
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *UsersListCall) Do() (*UsersListResponse, error) {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *UsersListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
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

func (c *UsersRevokeTokenCall) Do() error {
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
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
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
