// Package proximitybeacon provides access to the Google Proximity Beacon API.
//
// See https://developers.google.com/beacons/proximity/
//
// Usage example:
//
//   import "google.golang.org/api/proximitybeacon/v1beta1"
//   ...
//   proximitybeaconService, err := proximitybeacon.New(oauthHttpClient)
package proximitybeacon

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

const apiId = "proximitybeacon:v1beta1"
const apiName = "proximitybeacon"
const apiVersion = "v1beta1"
const basePath = "https://proximitybeacon.googleapis.com/"

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Beaconinfo = NewBeaconinfoService(s)
	s.Beacons = NewBeaconsService(s)
	s.Namespaces = NewNamespacesService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Beaconinfo *BeaconinfoService

	Beacons *BeaconsService

	Namespaces *NamespacesService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewBeaconinfoService(s *Service) *BeaconinfoService {
	rs := &BeaconinfoService{s: s}
	return rs
}

type BeaconinfoService struct {
	s *Service
}

func NewBeaconsService(s *Service) *BeaconsService {
	rs := &BeaconsService{s: s}
	rs.Attachments = NewBeaconsAttachmentsService(s)
	rs.Diagnostics = NewBeaconsDiagnosticsService(s)
	return rs
}

type BeaconsService struct {
	s *Service

	Attachments *BeaconsAttachmentsService

	Diagnostics *BeaconsDiagnosticsService
}

func NewBeaconsAttachmentsService(s *Service) *BeaconsAttachmentsService {
	rs := &BeaconsAttachmentsService{s: s}
	return rs
}

type BeaconsAttachmentsService struct {
	s *Service
}

func NewBeaconsDiagnosticsService(s *Service) *BeaconsDiagnosticsService {
	rs := &BeaconsDiagnosticsService{s: s}
	return rs
}

type BeaconsDiagnosticsService struct {
	s *Service
}

func NewNamespacesService(s *Service) *NamespacesService {
	rs := &NamespacesService{s: s}
	return rs
}

type NamespacesService struct {
	s *Service
}

// AdvertisedId: Defines a unique identifier of a beacon as broadcast by
// the device.
type AdvertisedId struct {
	// Id: The actual beacon identifier, as broadcast by the beacon
	// hardware. Must be
	// [base64](http://tools.ietf.org/html/rfc4648#section-4) encoded in
	// HTTP requests, and will be so encoded (with padding) in responses.
	// The base64 encoding should be of the binary byte-stream and not any
	// textual (such as hex) representation thereof. Required.
	Id string `json:"id,omitempty"`

	// Type: Specifies the identifier type. Required.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED"
	//   "EDDYSTONE"
	//   "IBEACON"
	//   "ALTBEACON"
	Type string `json:"type,omitempty"`
}

// AttachmentInfo: A subset of attachment information served via the
// `beaconinfo.getforobserved` method, used when your users encounter
// your beacons.
type AttachmentInfo struct {
	// Data: An opaque data container for client-provided data.
	Data string `json:"data,omitempty"`

	// NamespacedType: Specifies what kind of attachment this is. Tells a
	// client how to interpret the `data` field. Format is namespace/type,
	// for example scrupulous-wombat-12345/welcome-message
	NamespacedType string `json:"namespacedType,omitempty"`
}

// Beacon: Details of a beacon device.
type Beacon struct {
	// AdvertisedId: The identifier of a beacon as advertised by it. This
	// field must be populated when registering. It may be empty when
	// updating a beacon record because it is ignored in updates.
	AdvertisedId *AdvertisedId `json:"advertisedId,omitempty"`

	// BeaconName: Resource name of this beacon. A beacon name has the
	// format "beacons/N!beaconId" where the beaconId is the base16 ID
	// broadcast by the beacon and N is a code for the beacon's type.
	// Possible values are `3` for Eddystone, `1` for iBeacon, or `5` for
	// AltBeacon. This field must be left empty when registering. After
	// reading a beacon, clients can use the name for future operations.
	BeaconName string `json:"beaconName,omitempty"`

	// Description: Free text used to identify and describe the beacon.
	// Maximum length 140 characters. Optional.
	Description string `json:"description,omitempty"`

	// ExpectedStability: Expected location stability. This is set when the
	// beacon is registered or updated, not automatically detected in any
	// way. Optional.
	//
	// Possible values:
	//   "STABILITY_UNSPECIFIED"
	//   "STABLE"
	//   "PORTABLE"
	//   "MOBILE"
	//   "ROVING"
	ExpectedStability string `json:"expectedStability,omitempty"`

	// IndoorLevel: The indoor level information for this beacon, if known.
	// As returned by the Google Maps API. Optional.
	IndoorLevel *IndoorLevel `json:"indoorLevel,omitempty"`

	// LatLng: The location of the beacon, expressed as a latitude and
	// longitude pair. This location is given when the beacon is registered
	// or updated. It does not necessarily indicate the actual current
	// location of the beacon. Optional.
	LatLng *LatLng `json:"latLng,omitempty"`

	// PlaceId: The [Google Places API](/places/place-id) Place ID of the
	// place where the beacon is deployed. This is given when the beacon is
	// registered or updated, not automatically detected in any way.
	// Optional.
	PlaceId string `json:"placeId,omitempty"`

	// Properties: Properties of the beacon device, for example battery type
	// or firmware version. Optional.
	Properties map[string]string `json:"properties,omitempty"`

	// Status: Current status of the beacon. Required.
	//
	// Possible values:
	//   "STATUS_UNSPECIFIED"
	//   "ACTIVE"
	//   "DECOMMISSIONED"
	//   "INACTIVE"
	Status string `json:"status,omitempty"`
}

// BeaconAttachment: Project-specific data associated with a beacon.
type BeaconAttachment struct {
	// AttachmentName: Resource name of this attachment. Attachment names
	// have the format: beacons/beacon_id/attachments/attachment_id. Leave
	// this empty on creation.
	AttachmentName string `json:"attachmentName,omitempty"`

	// Data: An opaque data container for client-provided data. Must be
	// [base64](http://tools.ietf.org/html/rfc4648#section-4) encoded in
	// HTTP requests, and will be so encoded (with padding) in responses.
	// Required.
	Data string `json:"data,omitempty"`

	// NamespacedType: Specifies what kind of attachment this is. Tells a
	// client how to interpret the `data` field. Format is namespace/type.
	// Namespace provides type separation between clients. Type describes
	// the type of `data`, for use by the client when parsing the `data`
	// field. Required.
	NamespacedType string `json:"namespacedType,omitempty"`
}

// BeaconInfo: A subset of beacon information served via the
// `beaconinfo.getforobserved` method, which you call when users of your
// app encounter your beacons.
type BeaconInfo struct {
	// AdvertisedId: The ID advertised by the beacon.
	AdvertisedId *AdvertisedId `json:"advertisedId,omitempty"`

	// Attachments: Attachments matching the type(s) requested. May be empty
	// if no attachment types were requested, or if none matched.
	Attachments []*AttachmentInfo `json:"attachments,omitempty"`

	// BeaconName: The name under which the beacon is registered.
	BeaconName string `json:"beaconName,omitempty"`

	// Description: Free text used to identify or describe the beacon in a
	// registered establishment. For example: "entrance", "room 101", etc.
	// May be empty.
	Description string `json:"description,omitempty"`
}

// Date: Represents a whole calendar date, e.g. date of birth. The time
// of day and time zone are either specified elsewhere or are not
// significant. The date is relative to the Proleptic Gregorian
// Calendar. The day may be 0 to represent a year and month where the
// day is not significant, e.g. credit card expiration date. The year
// may be 0 to represent a month and day independent of year, e.g.
// anniversary date. Related types are [google.type.TimeOfDay][] and
// `google.protobuf.Timestamp`.
type Date struct {
	// Day: Day of month. Must be from 1 to 31 and valid for the year and
	// month, or 0 if specifying a year/month where the day is not
	// sigificant.
	Day int64 `json:"day,omitempty"`

	// Month: Month of year of date. Must be from 1 to 12.
	Month int64 `json:"month,omitempty"`

	// Year: Year of date. Must be from 1 to 9,999, or 0 if specifying a
	// date without a year.
	Year int64 `json:"year,omitempty"`
}

// DeleteAttachmentsResponse: Response for a request to delete
// attachments.
type DeleteAttachmentsResponse struct {
	// NumDeleted: The number of attachments that were deleted.
	NumDeleted int64 `json:"numDeleted,omitempty"`
}

// Diagnostics: Diagnostics for a single beacon.
type Diagnostics struct {
	// Alerts: An unordered list of Alerts that the beacon has.
	//
	// Possible values:
	//   "ALERT_UNSPECIFIED" - Invalid value. Should never appear.
	//   "WRONG_LOCATION" - The beacon has been reported in a location
	// different than its registered location. This may indicate that the
	// beacon has been moved. This signal is not 100% accurate, but
	// indicates that further investigation is worth while.
	//   "LOW_BATTERY" - The battery level for the beacon is low enough
	// that, given the beacon's current use, its battery will run out with
	// in the next 60 days. This indicates that the battery should be
	// replaced soon.
	Alerts []string `json:"alerts,omitempty"`

	// BeaconName: Resource name of the beacon.
	BeaconName string `json:"beaconName,omitempty"`

	// EstimatedLowBatteryDate: The date when the battery is expected to be
	// low. If the value is missing then there is no estimate for when the
	// battery will be low. This value is only an estimate, not an exact
	// date.
	EstimatedLowBatteryDate *Date `json:"estimatedLowBatteryDate,omitempty"`
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated empty messages in your APIs. A typical example is to use
// it as the request or the response type of an API method. For
// instance: service Foo { rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty); } The JSON representation for `Empty` is
// empty JSON object `{}`.
type Empty struct {
}

// GetInfoForObservedBeaconsRequest: Request for beacon and attachment
// information about beacons that a mobile client has encountered "in
// the wild".
type GetInfoForObservedBeaconsRequest struct {
	// NamespacedTypes: Specifies what kind of attachments to include in the
	// response. When given, the response will include only attachments of
	// the given types. When empty, no attachments will be returned. Must be
	// in the format namespace/type. Accepts `*` to specify all types in all
	// namespaces. Optional.
	NamespacedTypes []string `json:"namespacedTypes,omitempty"`

	// Observations: The beacons that the client has encountered. At least
	// one must be given.
	Observations []*Observation `json:"observations,omitempty"`
}

// GetInfoForObservedBeaconsResponse: Information about the requested
// beacons, optionally including attachment data.
type GetInfoForObservedBeaconsResponse struct {
	// Beacons: Public information about beacons. May be empty if the
	// request matched no beacons.
	Beacons []*BeaconInfo `json:"beacons,omitempty"`
}

// IndoorLevel: Indoor level, a human-readable string as returned by
// Google Maps APIs, useful to indicate which floor of a building a
// beacon is located on.
type IndoorLevel struct {
	// Name: The name of this level.
	Name string `json:"name,omitempty"`
}

// LatLng: An object representing a latitude/longitude pair. This is
// expressed as a pair of doubles representing degrees latitude and
// degrees longitude. Unless specified otherwise, this must conform to
// the WGS84 standard. Values must be within normalized ranges. Example
// of normalization code in Python: def NormalizeLongitude(longitude):
// """Wrapsdecimal degrees longitude to [-180.0, 180.0].""" q, r =
// divmod(longitude, 360.0) if r > 180.0 or (r == 180.0 and q <= -1.0):
// return r - 360.0 return r def NormalizeLatLng(latitude, longitude):
// """Wraps decimal degrees latitude and longitude to [-180.0, 180.0]
// and [-90.0, 90.0], respectively.""" r = latitude % 360.0 if r =
// 270.0: return r - 360, NormalizeLongitude(longitude) else: return 180
// - r, NormalizeLongitude(longitude + 180.0) assert 180.0 ==
// NormalizeLongitude(180.0) assert -180.0 == NormalizeLongitude(-180.0)
// assert -179.0 == NormalizeLongitude(181.0) assert (0.0, 0.0) ==
// NormalizeLatLng(360.0, 0.0) assert (0.0, 0.0) ==
// NormalizeLatLng(-360.0, 0.0) assert (85.0, 180.0) ==
// NormalizeLatLng(95.0, 0.0) assert (-85.0, -170.0) ==
// NormalizeLatLng(-95.0, 10.0) assert (90.0, 10.0) ==
// NormalizeLatLng(90.0, 10.0) assert (-90.0, -10.0) ==
// NormalizeLatLng(-90.0, -10.0) assert (0.0, -170.0) ==
// NormalizeLatLng(-180.0, 10.0) assert (0.0, -170.0) ==
// NormalizeLatLng(180.0, 10.0) assert (-90.0, 10.0) ==
// NormalizeLatLng(270.0, 10.0) assert (90.0, 10.0) ==
// NormalizeLatLng(-270.0, 10.0)
type LatLng struct {
	// Latitude: The latitude in degrees. It must be in the range [-90.0,
	// +90.0].
	Latitude float64 `json:"latitude,omitempty"`

	// Longitude: The longitude in degrees. It must be in the range [-180.0,
	// +180.0].
	Longitude float64 `json:"longitude,omitempty"`
}

// ListBeaconAttachmentsResponse: Response to ListBeaconAttachments that
// contains the requested attachments.
type ListBeaconAttachmentsResponse struct {
	// Attachments: The attachments that corresponded to the request params.
	Attachments []*BeaconAttachment `json:"attachments,omitempty"`
}

// ListBeaconsResponse: Response that contains list beacon results and
// pagination help.
type ListBeaconsResponse struct {
	// Beacons: The beacons that matched the search criteria.
	Beacons []*Beacon `json:"beacons,omitempty"`

	// NextPageToken: An opaque pagination token that the client may provide
	// in their next request to retrieve the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalCount: Estimate of the total number of beacons matched by the
	// query. Higher values may be less accurate.
	TotalCount int64 `json:"totalCount,omitempty,string"`
}

// ListDiagnosticsResponse: Response that contains the requested
// diagnostics.
type ListDiagnosticsResponse struct {
	// Diagnostics: The diagnostics matching the given request.
	Diagnostics []*Diagnostics `json:"diagnostics,omitempty"`

	// NextPageToken: Token that can be used for pagination. Returned only
	// if the request matches more beacons than can be returned in this
	// response.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

// ListNamespacesResponse: Response to ListNamespacesRequest that
// contains all the project's namespaces.
type ListNamespacesResponse struct {
	// Namespaces: The attachments that corresponded to the request params.
	Namespaces []*Namespace `json:"namespaces,omitempty"`
}

// Namespace: An attachment namespace defines read and write access for
// all the attachments created under it. Each namespace is globally
// unique, and owned by one project which is the only project that can
// create attachments under it.
type Namespace struct {
	// NamespaceName: Resource name of this namespace. Namespaces names have
	// the format: namespaces/namespace.
	NamespaceName string `json:"namespaceName,omitempty"`

	// ServingVisibility: Specifies what clients may receive attachments
	// under this namespace via `beaconinfo.getforobserved`.
	//
	// Possible values:
	//   "VISIBILITY_UNSPECIFIED"
	//   "UNLISTED"
	//   "PUBLIC"
	ServingVisibility string `json:"servingVisibility,omitempty"`
}

// Observation: Represents one beacon observed once.
type Observation struct {
	// AdvertisedId: The ID advertised by the beacon the client has
	// encountered. Required.
	AdvertisedId *AdvertisedId `json:"advertisedId,omitempty"`

	// Telemetry: The array of telemetry bytes received from the beacon. The
	// server is responsible for parsing it. This field may frequently be
	// empty, as with a beacon that transmits telemetry only occasionally.
	Telemetry string `json:"telemetry,omitempty"`

	// TimestampMs: Time when the beacon was observed. Being sourced from a
	// mobile device, this time may be suspect.
	TimestampMs string `json:"timestampMs,omitempty"`
}

// method id "proximitybeacon.beaconinfo.getforobserved":

type BeaconinfoGetforobservedCall struct {
	s                                *Service
	getinfoforobservedbeaconsrequest *GetInfoForObservedBeaconsRequest
	opt_                             map[string]interface{}
}

// Getforobserved: Given one or more beacon observations, returns any
// beacon information and attachments accessible to your application.
func (r *BeaconinfoService) Getforobserved(getinfoforobservedbeaconsrequest *GetInfoForObservedBeaconsRequest) *BeaconinfoGetforobservedCall {
	c := &BeaconinfoGetforobservedCall{s: r.s, opt_: make(map[string]interface{})}
	c.getinfoforobservedbeaconsrequest = getinfoforobservedbeaconsrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BeaconinfoGetforobservedCall) Fields(s ...googleapi.Field) *BeaconinfoGetforobservedCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *BeaconinfoGetforobservedCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.getinfoforobservedbeaconsrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/beaconinfo:getforobserved")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *BeaconinfoGetforobservedCall) Do() (*GetInfoForObservedBeaconsResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *GetInfoForObservedBeaconsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Given one or more beacon observations, returns any beacon information and attachments accessible to your application.",
	//   "httpMethod": "POST",
	//   "id": "proximitybeacon.beaconinfo.getforobserved",
	//   "path": "v1beta1/beaconinfo:getforobserved",
	//   "request": {
	//     "$ref": "GetInfoForObservedBeaconsRequest"
	//   },
	//   "response": {
	//     "$ref": "GetInfoForObservedBeaconsResponse"
	//   }
	// }

}

// method id "proximitybeacon.beacons.activate":

type BeaconsActivateCall struct {
	s          *Service
	beaconName string
	opt_       map[string]interface{}
}

// Activate: (Re)activates a beacon. A beacon that is active will return
// information and attachment data when queried via
// `beaconinfo.getforobserved`. Calling this method on an already active
// beacon will do nothing (but will return a successful response code).
func (r *BeaconsService) Activate(beaconName string) *BeaconsActivateCall {
	c := &BeaconsActivateCall{s: r.s, opt_: make(map[string]interface{})}
	c.beaconName = beaconName
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BeaconsActivateCall) Fields(s ...googleapi.Field) *BeaconsActivateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *BeaconsActivateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+beaconName}:activate")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"beaconName": c.beaconName,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *BeaconsActivateCall) Do() (*Empty, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "(Re)activates a beacon. A beacon that is active will return information and attachment data when queried via `beaconinfo.getforobserved`. Calling this method on an already active beacon will do nothing (but will return a successful response code).",
	//   "httpMethod": "POST",
	//   "id": "proximitybeacon.beacons.activate",
	//   "parameterOrder": [
	//     "beaconName"
	//   ],
	//   "parameters": {
	//     "beaconName": {
	//       "description": "The beacon to activate. Required.",
	//       "location": "path",
	//       "pattern": "^beacons/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+beaconName}:activate",
	//   "response": {
	//     "$ref": "Empty"
	//   }
	// }

}

// method id "proximitybeacon.beacons.deactivate":

type BeaconsDeactivateCall struct {
	s          *Service
	beaconName string
	opt_       map[string]interface{}
}

// Deactivate: Deactivates a beacon. Once deactivated, the API will not
// return information nor attachment data for the beacon when queried
// via `beaconinfo.getforobserved`. Calling this method on an already
// inactive beacon will do nothing (but will return a successful
// response code).
func (r *BeaconsService) Deactivate(beaconName string) *BeaconsDeactivateCall {
	c := &BeaconsDeactivateCall{s: r.s, opt_: make(map[string]interface{})}
	c.beaconName = beaconName
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BeaconsDeactivateCall) Fields(s ...googleapi.Field) *BeaconsDeactivateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *BeaconsDeactivateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+beaconName}:deactivate")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"beaconName": c.beaconName,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *BeaconsDeactivateCall) Do() (*Empty, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deactivates a beacon. Once deactivated, the API will not return information nor attachment data for the beacon when queried via `beaconinfo.getforobserved`. Calling this method on an already inactive beacon will do nothing (but will return a successful response code).",
	//   "httpMethod": "POST",
	//   "id": "proximitybeacon.beacons.deactivate",
	//   "parameterOrder": [
	//     "beaconName"
	//   ],
	//   "parameters": {
	//     "beaconName": {
	//       "description": "The beacon name of this beacon.",
	//       "location": "path",
	//       "pattern": "^beacons/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+beaconName}:deactivate",
	//   "response": {
	//     "$ref": "Empty"
	//   }
	// }

}

// method id "proximitybeacon.beacons.decommission":

type BeaconsDecommissionCall struct {
	s          *Service
	beaconName string
	opt_       map[string]interface{}
}

// Decommission: Decommissions the specified beacon in the service. This
// beacon will no longer be returned from `beaconinfo.getforobserved`.
// This operation is permanent -- you will not be able to re-register a
// beacon with this ID again.
func (r *BeaconsService) Decommission(beaconName string) *BeaconsDecommissionCall {
	c := &BeaconsDecommissionCall{s: r.s, opt_: make(map[string]interface{})}
	c.beaconName = beaconName
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BeaconsDecommissionCall) Fields(s ...googleapi.Field) *BeaconsDecommissionCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *BeaconsDecommissionCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+beaconName}:decommission")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"beaconName": c.beaconName,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *BeaconsDecommissionCall) Do() (*Empty, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Decommissions the specified beacon in the service. This beacon will no longer be returned from `beaconinfo.getforobserved`. This operation is permanent -- you will not be able to re-register a beacon with this ID again.",
	//   "httpMethod": "POST",
	//   "id": "proximitybeacon.beacons.decommission",
	//   "parameterOrder": [
	//     "beaconName"
	//   ],
	//   "parameters": {
	//     "beaconName": {
	//       "description": "Beacon that should be decommissioned. Required.",
	//       "location": "path",
	//       "pattern": "^beacons/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+beaconName}:decommission",
	//   "response": {
	//     "$ref": "Empty"
	//   }
	// }

}

// method id "proximitybeacon.beacons.get":

type BeaconsGetCall struct {
	s          *Service
	beaconName string
	opt_       map[string]interface{}
}

// Get: Returns detailed information about the specified beacon.
func (r *BeaconsService) Get(beaconName string) *BeaconsGetCall {
	c := &BeaconsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.beaconName = beaconName
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BeaconsGetCall) Fields(s ...googleapi.Field) *BeaconsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *BeaconsGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+beaconName}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"beaconName": c.beaconName,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *BeaconsGetCall) Do() (*Beacon, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Beacon
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns detailed information about the specified beacon.",
	//   "httpMethod": "GET",
	//   "id": "proximitybeacon.beacons.get",
	//   "parameterOrder": [
	//     "beaconName"
	//   ],
	//   "parameters": {
	//     "beaconName": {
	//       "description": "Beacon that is requested.",
	//       "location": "path",
	//       "pattern": "^beacons/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+beaconName}",
	//   "response": {
	//     "$ref": "Beacon"
	//   }
	// }

}

// method id "proximitybeacon.beacons.list":

type BeaconsListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Searches the beacon registry for beacons that match the given
// search criteria. Only those beacons that the client has permission to
// list will be returned.
func (r *BeaconsService) List() *BeaconsListCall {
	c := &BeaconsListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of records to return for this request, up to a server-defined upper
// limit.
func (c *BeaconsListCall) PageSize(pageSize int64) *BeaconsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": A pagination token
// obtained from a previous request to list beacons.
func (c *BeaconsListCall) PageToken(pageToken string) *BeaconsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Q sets the optional parameter "q": Filter query string that supports
// the following field filters: * `description:"" For example:
// `description:"Room 3" Returns beacons whose description matches
// tokens in the string "Room 3" (not necessarily that exact string).
// The string must be double-quoted. * `status:` For example:
// `status:active` Returns beacons whose status matches the given value.
// Values must be one of the Beacon.Status enum values (case
// insensitive). Accepts multiple filters which will be combined with OR
// logic. * `stability:` For example: `stability:mobile` Returns beacons
// whose expected stability matches the given value. Values must be one
// of the Beacon.Stability enum values (case insensitive). Accepts
// multiple filters which will be combined with OR logic. *
// `place_id:"" For example: `place_id:"ChIJVSZzVR8FdkgRXGmmm6SslKw="
// Returns beacons explicitly registered at the given place, expressed
// as a Place ID obtained from [Google Places API](/places/place-id).
// Does not match places inside the given place. Does not consider the
// beacon's actual location (which may be different from its registered
// place). Accepts multiple filters that will be combined with OR logic.
// The place ID must be double-quoted. * `registration_time[|=]` For
// example: `registration_time>=1433116800` Returns beacons whose
// registration time matches the given filter. Supports the operators: ,
// =. Timestamp must be expressed as an integer number of seconds since
// midnight January 1, 1970 UTC. Accepts at most two filters that will
// be combined with AND logic, to support "between" semantics. If more
// than two are supplied, the latter ones are ignored. * `lat: lng:
// radius:` For example: `lat:51.1232343 lng:-1.093852 radius:1000`
// Returns beacons whose registered location is within the given circle.
// When any of these fields are given, all are required. Latitude and
// longitude must be decimal degrees between -90.0 and 90.0 and between
// -180.0 and 180.0 respectively. Radius must be an integer number of
// meters less than 1,000,000 (1000 km). * `property:"=" For example:
// `property:"battery-type=CR2032" Returns beacons which have a
// property of the given name and value. Supports multiple filters which
// will be combined with OR logic. The entire name=value string must be
// double-quoted as one string. * `attachment_type:"" For example:
// `attachment_type:"my-namespace/my-type" Returns beacons having at
// least one attachment of the given namespaced type. Supports "any
// within this namespace" via the partial wildcard syntax:
// "my-namespace/*". Supports multiple filters which will be combined
// with OR logic. The string must be double-quoted. Multiple filters on
// the same field are combined with OR logic (except registration_time
// which is combined with AND logic). Multiple filters on different
// fields are combined with AND logic. Filters should be separated by
// spaces. As with any HTTP query string parameter, the whole filter
// expression must be URL-encoded. Example REST request: `GET
// /v1beta1/beacons?q=status:active%20lat:51.123%20lng:-1.095%20radius:10
// 00`
func (c *BeaconsListCall) Q(q string) *BeaconsListCall {
	c.opt_["q"] = q
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BeaconsListCall) Fields(s ...googleapi.Field) *BeaconsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *BeaconsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["q"]; ok {
		params.Set("q", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/beacons")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *BeaconsListCall) Do() (*ListBeaconsResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListBeaconsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Searches the beacon registry for beacons that match the given search criteria. Only those beacons that the client has permission to list will be returned.",
	//   "httpMethod": "GET",
	//   "id": "proximitybeacon.beacons.list",
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The maximum number of records to return for this request, up to a server-defined upper limit.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A pagination token obtained from a previous request to list beacons.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "q": {
	//       "description": "Filter query string that supports the following field filters: * `description:\"\"` For example: `description:\"Room 3\"` Returns beacons whose description matches tokens in the string \"Room 3\" (not necessarily that exact string). The string must be double-quoted. * `status:` For example: `status:active` Returns beacons whose status matches the given value. Values must be one of the Beacon.Status enum values (case insensitive). Accepts multiple filters which will be combined with OR logic. * `stability:` For example: `stability:mobile` Returns beacons whose expected stability matches the given value. Values must be one of the Beacon.Stability enum values (case insensitive). Accepts multiple filters which will be combined with OR logic. * `place_id:\"\"` For example: `place_id:\"ChIJVSZzVR8FdkgRXGmmm6SslKw=\"` Returns beacons explicitly registered at the given place, expressed as a Place ID obtained from [Google Places API](/places/place-id). Does not match places inside the given place. Does not consider the beacon's actual location (which may be different from its registered place). Accepts multiple filters that will be combined with OR logic. The place ID must be double-quoted. * `registration_time[|=]` For example: `registration_time\u003e=1433116800` Returns beacons whose registration time matches the given filter. Supports the operators: , =. Timestamp must be expressed as an integer number of seconds since midnight January 1, 1970 UTC. Accepts at most two filters that will be combined with AND logic, to support \"between\" semantics. If more than two are supplied, the latter ones are ignored. * `lat: lng: radius:` For example: `lat:51.1232343 lng:-1.093852 radius:1000` Returns beacons whose registered location is within the given circle. When any of these fields are given, all are required. Latitude and longitude must be decimal degrees between -90.0 and 90.0 and between -180.0 and 180.0 respectively. Radius must be an integer number of meters less than 1,000,000 (1000 km). * `property:\"=\"` For example: `property:\"battery-type=CR2032\"` Returns beacons which have a property of the given name and value. Supports multiple filters which will be combined with OR logic. The entire name=value string must be double-quoted as one string. * `attachment_type:\"\"` For example: `attachment_type:\"my-namespace/my-type\"` Returns beacons having at least one attachment of the given namespaced type. Supports \"any within this namespace\" via the partial wildcard syntax: \"my-namespace/*\". Supports multiple filters which will be combined with OR logic. The string must be double-quoted. Multiple filters on the same field are combined with OR logic (except registration_time which is combined with AND logic). Multiple filters on different fields are combined with AND logic. Filters should be separated by spaces. As with any HTTP query string parameter, the whole filter expression must be URL-encoded. Example REST request: `GET /v1beta1/beacons?q=status:active%20lat:51.123%20lng:-1.095%20radius:1000`",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/beacons",
	//   "response": {
	//     "$ref": "ListBeaconsResponse"
	//   }
	// }

}

// method id "proximitybeacon.beacons.register":

type BeaconsRegisterCall struct {
	s      *Service
	beacon *Beacon
	opt_   map[string]interface{}
}

// Register: Registers a previously unregistered beacon given its
// `advertisedId`. These IDs are unique within the system. An ID can be
// registered only once.
func (r *BeaconsService) Register(beacon *Beacon) *BeaconsRegisterCall {
	c := &BeaconsRegisterCall{s: r.s, opt_: make(map[string]interface{})}
	c.beacon = beacon
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BeaconsRegisterCall) Fields(s ...googleapi.Field) *BeaconsRegisterCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *BeaconsRegisterCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.beacon)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/beacons:register")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *BeaconsRegisterCall) Do() (*Beacon, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Beacon
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Registers a previously unregistered beacon given its `advertisedId`. These IDs are unique within the system. An ID can be registered only once.",
	//   "httpMethod": "POST",
	//   "id": "proximitybeacon.beacons.register",
	//   "path": "v1beta1/beacons:register",
	//   "request": {
	//     "$ref": "Beacon"
	//   },
	//   "response": {
	//     "$ref": "Beacon"
	//   }
	// }

}

// method id "proximitybeacon.beacons.update":

type BeaconsUpdateCall struct {
	s          *Service
	beaconName string
	beacon     *Beacon
	opt_       map[string]interface{}
}

// Update: Updates the information about the specified beacon. **Any
// field that you do not populate in the submitted beacon will be
// permanently erased**, so you should follow the "read, modify, write"
// pattern to avoid inadvertently destroying data. Changes to the beacon
// status via this method will be silently ignored. To update beacon
// status, use the separate methods on this API for (de)activation and
// decommissioning.
func (r *BeaconsService) Update(beaconName string, beacon *Beacon) *BeaconsUpdateCall {
	c := &BeaconsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.beaconName = beaconName
	c.beacon = beacon
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BeaconsUpdateCall) Fields(s ...googleapi.Field) *BeaconsUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *BeaconsUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.beacon)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+beaconName}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"beaconName": c.beaconName,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *BeaconsUpdateCall) Do() (*Beacon, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Beacon
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the information about the specified beacon. **Any field that you do not populate in the submitted beacon will be permanently erased**, so you should follow the \"read, modify, write\" pattern to avoid inadvertently destroying data. Changes to the beacon status via this method will be silently ignored. To update beacon status, use the separate methods on this API for (de)activation and decommissioning.",
	//   "httpMethod": "PUT",
	//   "id": "proximitybeacon.beacons.update",
	//   "parameterOrder": [
	//     "beaconName"
	//   ],
	//   "parameters": {
	//     "beaconName": {
	//       "description": "Resource name of this beacon. A beacon name has the format \"beacons/N!beaconId\" where the beaconId is the base16 ID broadcast by the beacon and N is a code for the beacon's type. Possible values are `3` for Eddystone, `1` for iBeacon, or `5` for AltBeacon. This field must be left empty when registering. After reading a beacon, clients can use the name for future operations.",
	//       "location": "path",
	//       "pattern": "^beacons/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+beaconName}",
	//   "request": {
	//     "$ref": "Beacon"
	//   },
	//   "response": {
	//     "$ref": "Beacon"
	//   }
	// }

}

// method id "proximitybeacon.beacons.attachments.batchDelete":

type BeaconsAttachmentsBatchDeleteCall struct {
	s          *Service
	beaconName string
	opt_       map[string]interface{}
}

// BatchDelete: Deletes multiple attachments on a given beacon. This
// operation is permanent and cannot be undone. You can optionally
// specify `namespacedType` to choose which attachments should be
// deleted. If you do not specify `namespacedType`, all your attachments
// on the given beacon will be deleted. You also may explicitly specify
// `*/*` to delete all.
func (r *BeaconsAttachmentsService) BatchDelete(beaconName string) *BeaconsAttachmentsBatchDeleteCall {
	c := &BeaconsAttachmentsBatchDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.beaconName = beaconName
	return c
}

// NamespacedType sets the optional parameter "namespacedType":
// Specifies the namespace and type of attachments to delete in
// `namespace/type` format. Accepts `*/*` to specify "all types in all
// namespaces".
func (c *BeaconsAttachmentsBatchDeleteCall) NamespacedType(namespacedType string) *BeaconsAttachmentsBatchDeleteCall {
	c.opt_["namespacedType"] = namespacedType
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BeaconsAttachmentsBatchDeleteCall) Fields(s ...googleapi.Field) *BeaconsAttachmentsBatchDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *BeaconsAttachmentsBatchDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["namespacedType"]; ok {
		params.Set("namespacedType", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+beaconName}/attachments:batchDelete")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"beaconName": c.beaconName,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *BeaconsAttachmentsBatchDeleteCall) Do() (*DeleteAttachmentsResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *DeleteAttachmentsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes multiple attachments on a given beacon. This operation is permanent and cannot be undone. You can optionally specify `namespacedType` to choose which attachments should be deleted. If you do not specify `namespacedType`, all your attachments on the given beacon will be deleted. You also may explicitly specify `*/*` to delete all.",
	//   "httpMethod": "POST",
	//   "id": "proximitybeacon.beacons.attachments.batchDelete",
	//   "parameterOrder": [
	//     "beaconName"
	//   ],
	//   "parameters": {
	//     "beaconName": {
	//       "description": "The beacon whose attachments are to be deleted. Required.",
	//       "location": "path",
	//       "pattern": "^beacons/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "namespacedType": {
	//       "description": "Specifies the namespace and type of attachments to delete in `namespace/type` format. Accepts `*/*` to specify \"all types in all namespaces\". Optional.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+beaconName}/attachments:batchDelete",
	//   "response": {
	//     "$ref": "DeleteAttachmentsResponse"
	//   }
	// }

}

// method id "proximitybeacon.beacons.attachments.create":

type BeaconsAttachmentsCreateCall struct {
	s                *Service
	beaconName       string
	beaconattachment *BeaconAttachment
	opt_             map[string]interface{}
}

// Create: Associates the given data with the specified beacon.
// Attachment data must contain two parts:
// - A namespaced type.
// - The actual attachment data itself.  The namespaced type consists of
// two parts, the namespace and the type. The namespace must be one of
// the values returned by the `namespaces` endpoint, while the type can
// be a string of any characters except for the forward slash (`/`) up
// to 100 characters in length. Attachment data can be up to 1024 bytes
// long.
func (r *BeaconsAttachmentsService) Create(beaconName string, beaconattachment *BeaconAttachment) *BeaconsAttachmentsCreateCall {
	c := &BeaconsAttachmentsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.beaconName = beaconName
	c.beaconattachment = beaconattachment
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BeaconsAttachmentsCreateCall) Fields(s ...googleapi.Field) *BeaconsAttachmentsCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *BeaconsAttachmentsCreateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.beaconattachment)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+beaconName}/attachments")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"beaconName": c.beaconName,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *BeaconsAttachmentsCreateCall) Do() (*BeaconAttachment, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *BeaconAttachment
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Associates the given data with the specified beacon. Attachment data must contain two parts:  \n- A namespaced type. \n- The actual attachment data itself.  The namespaced type consists of two parts, the namespace and the type. The namespace must be one of the values returned by the `namespaces` endpoint, while the type can be a string of any characters except for the forward slash (`/`) up to 100 characters in length. Attachment data can be up to 1024 bytes long.",
	//   "httpMethod": "POST",
	//   "id": "proximitybeacon.beacons.attachments.create",
	//   "parameterOrder": [
	//     "beaconName"
	//   ],
	//   "parameters": {
	//     "beaconName": {
	//       "description": "The beacon on which the attachment should be created. Required.",
	//       "location": "path",
	//       "pattern": "^beacons/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+beaconName}/attachments",
	//   "request": {
	//     "$ref": "BeaconAttachment"
	//   },
	//   "response": {
	//     "$ref": "BeaconAttachment"
	//   }
	// }

}

// method id "proximitybeacon.beacons.attachments.delete":

type BeaconsAttachmentsDeleteCall struct {
	s              *Service
	attachmentName string
	opt_           map[string]interface{}
}

// Delete: Deletes the specified attachment for the given beacon. Each
// attachment has a unique attachment name (`attachmentName`) which is
// returned when you fetch the attachment data via this API. You specify
// this with the delete request to control which attachment is removed.
// This operation cannot be undone.
func (r *BeaconsAttachmentsService) Delete(attachmentName string) *BeaconsAttachmentsDeleteCall {
	c := &BeaconsAttachmentsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.attachmentName = attachmentName
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BeaconsAttachmentsDeleteCall) Fields(s ...googleapi.Field) *BeaconsAttachmentsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *BeaconsAttachmentsDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+attachmentName}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"attachmentName": c.attachmentName,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *BeaconsAttachmentsDeleteCall) Do() (*Empty, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the specified attachment for the given beacon. Each attachment has a unique attachment name (`attachmentName`) which is returned when you fetch the attachment data via this API. You specify this with the delete request to control which attachment is removed. This operation cannot be undone.",
	//   "httpMethod": "DELETE",
	//   "id": "proximitybeacon.beacons.attachments.delete",
	//   "parameterOrder": [
	//     "attachmentName"
	//   ],
	//   "parameters": {
	//     "attachmentName": {
	//       "description": "The attachment name (`attachmentName`) of the attachment to remove. For example: `beacons/3!893737abc9/attachments/c5e937-af0-494-959-ec49d12738` Required.",
	//       "location": "path",
	//       "pattern": "^beacons/[^/]*/attachments/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+attachmentName}",
	//   "response": {
	//     "$ref": "Empty"
	//   }
	// }

}

// method id "proximitybeacon.beacons.attachments.list":

type BeaconsAttachmentsListCall struct {
	s          *Service
	beaconName string
	opt_       map[string]interface{}
}

// List: Returns the attachments for the specified beacon that match the
// specified namespaced-type pattern. To control which namespaced types
// are returned, you add the `namespacedType` query parameter to the
// request. You must either use `*/*`, to return all attachments, or the
// namespace must be one of the ones returned from the `namespaces`
// endpoint.
func (r *BeaconsAttachmentsService) List(beaconName string) *BeaconsAttachmentsListCall {
	c := &BeaconsAttachmentsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.beaconName = beaconName
	return c
}

// NamespacedType sets the optional parameter "namespacedType":
// Specifies the namespace and type of attachment to include in response
// in namespace/type format. Accepts `*/*` to specify "all types in all
// namespaces".
func (c *BeaconsAttachmentsListCall) NamespacedType(namespacedType string) *BeaconsAttachmentsListCall {
	c.opt_["namespacedType"] = namespacedType
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BeaconsAttachmentsListCall) Fields(s ...googleapi.Field) *BeaconsAttachmentsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *BeaconsAttachmentsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["namespacedType"]; ok {
		params.Set("namespacedType", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+beaconName}/attachments")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"beaconName": c.beaconName,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *BeaconsAttachmentsListCall) Do() (*ListBeaconAttachmentsResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListBeaconAttachmentsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the attachments for the specified beacon that match the specified namespaced-type pattern. To control which namespaced types are returned, you add the `namespacedType` query parameter to the request. You must either use `*/*`, to return all attachments, or the namespace must be one of the ones returned from the `namespaces` endpoint.",
	//   "httpMethod": "GET",
	//   "id": "proximitybeacon.beacons.attachments.list",
	//   "parameterOrder": [
	//     "beaconName"
	//   ],
	//   "parameters": {
	//     "beaconName": {
	//       "description": "The beacon whose attachments are to be fetched. Required.",
	//       "location": "path",
	//       "pattern": "^beacons/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "namespacedType": {
	//       "description": "Specifies the namespace and type of attachment to include in response in namespace/type format. Accepts `*/*` to specify \"all types in all namespaces\".",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+beaconName}/attachments",
	//   "response": {
	//     "$ref": "ListBeaconAttachmentsResponse"
	//   }
	// }

}

// method id "proximitybeacon.beacons.diagnostics.list":

type BeaconsDiagnosticsListCall struct {
	s          *Service
	beaconName string
	opt_       map[string]interface{}
}

// List: List the diagnostics for a single beacon. You can also list
// diagnostics for all the beacons owned by your Google Developers
// Console project by using the beacon name `beacons/-`.
func (r *BeaconsDiagnosticsService) List(beaconName string) *BeaconsDiagnosticsListCall {
	c := &BeaconsDiagnosticsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.beaconName = beaconName
	return c
}

// AlertFilter sets the optional parameter "alertFilter": Requests only
// beacons that have the given alert. For example, to find beacons that
// have low batteries use `alert_filter=LOW_BATTERY`.
//
// Possible values:
//   "ALERT_UNSPECIFIED"
//   "WRONG_LOCATION"
//   "LOW_BATTERY"
func (c *BeaconsDiagnosticsListCall) AlertFilter(alertFilter string) *BeaconsDiagnosticsListCall {
	c.opt_["alertFilter"] = alertFilter
	return c
}

// PageSize sets the optional parameter "pageSize": Specifies the
// maximum number of results to return. Defaults to 10. Maximum 1000.
func (c *BeaconsDiagnosticsListCall) PageSize(pageSize int64) *BeaconsDiagnosticsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": Requests results
// that occur after the `page_token`, obtained from the response to a
// previous request.
func (c *BeaconsDiagnosticsListCall) PageToken(pageToken string) *BeaconsDiagnosticsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BeaconsDiagnosticsListCall) Fields(s ...googleapi.Field) *BeaconsDiagnosticsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *BeaconsDiagnosticsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["alertFilter"]; ok {
		params.Set("alertFilter", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+beaconName}/diagnostics")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"beaconName": c.beaconName,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *BeaconsDiagnosticsListCall) Do() (*ListDiagnosticsResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListDiagnosticsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List the diagnostics for a single beacon. You can also list diagnostics for all the beacons owned by your Google Developers Console project by using the beacon name `beacons/-`.",
	//   "httpMethod": "GET",
	//   "id": "proximitybeacon.beacons.diagnostics.list",
	//   "parameterOrder": [
	//     "beaconName"
	//   ],
	//   "parameters": {
	//     "alertFilter": {
	//       "description": "Requests only beacons that have the given alert. For example, to find beacons that have low batteries use `alert_filter=LOW_BATTERY`.",
	//       "enum": [
	//         "ALERT_UNSPECIFIED",
	//         "WRONG_LOCATION",
	//         "LOW_BATTERY"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "beaconName": {
	//       "description": "Beacon that the diagnostics are for.",
	//       "location": "path",
	//       "pattern": "^beacons/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Specifies the maximum number of results to return. Defaults to 10. Maximum 1000. Optional.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Requests results that occur after the `page_token`, obtained from the response to a previous request. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+beaconName}/diagnostics",
	//   "response": {
	//     "$ref": "ListDiagnosticsResponse"
	//   }
	// }

}

// method id "proximitybeacon.namespaces.list":

type NamespacesListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Lists all attachment namespaces owned by your Google Developers
// Console project. Attachment data associated with a beacon must
// include a namespaced type, and the namespace must be owned by your
// project.
func (r *NamespacesService) List() *NamespacesListCall {
	c := &NamespacesListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *NamespacesListCall) Fields(s ...googleapi.Field) *NamespacesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *NamespacesListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/namespaces")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *NamespacesListCall) Do() (*ListNamespacesResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListNamespacesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all attachment namespaces owned by your Google Developers Console project. Attachment data associated with a beacon must include a namespaced type, and the namespace must be owned by your project.",
	//   "httpMethod": "GET",
	//   "id": "proximitybeacon.namespaces.list",
	//   "path": "v1beta1/namespaces",
	//   "response": {
	//     "$ref": "ListNamespacesResponse"
	//   }
	// }

}
