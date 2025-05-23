// Copyright 2025 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package pollen provides access to the Pollen API.
//
// For product documentation, see: https://developers.google.com/maps/documentation/pollen
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
//	import "google.golang.org/api/pollen/v1"
//	...
//	ctx := context.Background()
//	pollenService, err := pollen.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for
// authentication. For information on how to create and obtain Application
// Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// # Other authentication options
//
// To use an API key for authentication (note: some APIs do not support API
// keys), use [google.golang.org/api/option.WithAPIKey]:
//
//	pollenService, err := pollen.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth
// flow, use [google.golang.org/api/option.WithTokenSource]:
//
//	config := &oauth2.Config{...}
//	// ...
//	token, err := config.Exchange(ctx, ...)
//	pollenService, err := pollen.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See [google.golang.org/api/option.ClientOption] for details on options.
package pollen // import "google.golang.org/api/pollen/v1"

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

const apiId = "pollen:v1"
const apiName = "pollen"
const apiVersion = "v1"
const basePath = "https://pollen.googleapis.com/"
const basePathTemplate = "https://pollen.UNIVERSE_DOMAIN/"
const mtlsBasePath = "https://pollen.mtls.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// See, edit, configure, and delete your Google Cloud data and see the email
	// address for your Google Account.
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
)

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	scopesOption := internaloption.WithDefaultScopes(
		"https://www.googleapis.com/auth/cloud-platform",
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
	s.Forecast = NewForecastService(s)
	s.MapTypes = NewMapTypesService(s)
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

	Forecast *ForecastService

	MapTypes *MapTypesService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewForecastService(s *Service) *ForecastService {
	rs := &ForecastService{s: s}
	return rs
}

type ForecastService struct {
	s *Service
}

func NewMapTypesService(s *Service) *MapTypesService {
	rs := &MapTypesService{s: s}
	rs.HeatmapTiles = NewMapTypesHeatmapTilesService(s)
	return rs
}

type MapTypesService struct {
	s *Service

	HeatmapTiles *MapTypesHeatmapTilesService
}

func NewMapTypesHeatmapTilesService(s *Service) *MapTypesHeatmapTilesService {
	rs := &MapTypesHeatmapTilesService{s: s}
	return rs
}

type MapTypesHeatmapTilesService struct {
	s *Service
}

// Color: Represents a color in the RGBA color space. This representation is
// designed for simplicity of conversion to and from color representations in
// various languages over compactness. For example, the fields of this
// representation can be trivially provided to the constructor of
// `java.awt.Color` in Java; it can also be trivially provided to UIColor's
// `+colorWithRed:green:blue:alpha` method in iOS; and, with just a little
// work, it can be easily formatted into a CSS `rgba()` string in JavaScript.
// This reference page doesn't have information about the absolute color space
// that should be used to interpret the RGB value—for example, sRGB, Adobe
// RGB, DCI-P3, and BT.2020. By default, applications should assume the sRGB
// color space. When color equality needs to be decided, implementations,
// unless documented otherwise, treat two colors as equal if all their red,
// green, blue, and alpha values each differ by at most `1e-5`. Example (Java):
// import com.google.type.Color; // ... public static java.awt.Color
// fromProto(Color protocolor) { float alpha = protocolor.hasAlpha() ?
// protocolor.getAlpha().getValue() : 1.0; return new java.awt.Color(
// protocolor.getRed(), protocolor.getGreen(), protocolor.getBlue(), alpha); }
// public static Color toProto(java.awt.Color color) { float red = (float)
// color.getRed(); float green = (float) color.getGreen(); float blue = (float)
// color.getBlue(); float denominator = 255.0; Color.Builder resultBuilder =
// Color .newBuilder() .setRed(red / denominator) .setGreen(green /
// denominator) .setBlue(blue / denominator); int alpha = color.getAlpha(); if
// (alpha != 255) { result.setAlpha( FloatValue .newBuilder()
// .setValue(((float) alpha) / denominator) .build()); } return
// resultBuilder.build(); } // ... Example (iOS / Obj-C): // ... static
// UIColor* fromProto(Color* protocolor) { float red = [protocolor red]; float
// green = [protocolor green]; float blue = [protocolor blue]; FloatValue*
// alpha_wrapper = [protocolor alpha]; float alpha = 1.0; if (alpha_wrapper !=
// nil) { alpha = [alpha_wrapper value]; } return [UIColor colorWithRed:red
// green:green blue:blue alpha:alpha]; } static Color* toProto(UIColor* color)
// { CGFloat red, green, blue, alpha; if (![color getRed:&red green:&green
// blue:&blue alpha:&alpha]) { return nil; } Color* result = [[Color alloc]
// init]; [result setRed:red]; [result setGreen:green]; [result setBlue:blue];
// if (alpha <= 0.9999) { [result setAlpha:floatWrapperWithValue(alpha)]; }
// [result autorelease]; return result; } // ... Example (JavaScript): // ...
// var protoToCssColor = function(rgb_color) { var redFrac = rgb_color.red ||
// 0.0; var greenFrac = rgb_color.green || 0.0; var blueFrac = rgb_color.blue
// || 0.0; var red = Math.floor(redFrac * 255); var green =
// Math.floor(greenFrac * 255); var blue = Math.floor(blueFrac * 255); if
// (!('alpha' in rgb_color)) { return rgbToCssColor(red, green, blue); } var
// alphaFrac = rgb_color.alpha.value || 0.0; var rgbParams = [red, green,
// blue].join(','); return ['rgba(', rgbParams, ',', alphaFrac, ')'].join(”);
// }; var rgbToCssColor = function(red, green, blue) { var rgbNumber = new
// Number((red << 16) | (green << 8) | blue); var hexString =
// rgbNumber.toString(16); var missingZeros = 6 - hexString.length; var
// resultBuilder = ['#']; for (var i = 0; i < missingZeros; i++) {
// resultBuilder.push('0'); } resultBuilder.push(hexString); return
// resultBuilder.join(”); }; // ...
type Color struct {
	// Alpha: The fraction of this color that should be applied to the pixel. That
	// is, the final pixel color is defined by the equation: `pixel color = alpha *
	// (this color) + (1.0 - alpha) * (background color)` This means that a value
	// of 1.0 corresponds to a solid color, whereas a value of 0.0 corresponds to a
	// completely transparent color. This uses a wrapper message rather than a
	// simple float scalar so that it is possible to distinguish between a default
	// value and the value being unset. If omitted, this color object is rendered
	// as a solid color (as if the alpha value had been explicitly given a value of
	// 1.0).
	Alpha float64 `json:"alpha,omitempty"`
	// Blue: The amount of blue in the color as a value in the interval [0, 1].
	Blue float64 `json:"blue,omitempty"`
	// Green: The amount of green in the color as a value in the interval [0, 1].
	Green float64 `json:"green,omitempty"`
	// Red: The amount of red in the color as a value in the interval [0, 1].
	Red float64 `json:"red,omitempty"`
	// ForceSendFields is a list of field names (e.g. "Alpha") to unconditionally
	// include in API requests. By default, fields with empty or default values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "Alpha") to include in API
	// requests with the JSON null value. By default, fields with empty values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s Color) MarshalJSON() ([]byte, error) {
	type NoMethod Color
	return gensupport.MarshalJSON(NoMethod(s), s.ForceSendFields, s.NullFields)
}

func (s *Color) UnmarshalJSON(data []byte) error {
	type NoMethod Color
	var s1 struct {
		Alpha gensupport.JSONFloat64 `json:"alpha"`
		Blue  gensupport.JSONFloat64 `json:"blue"`
		Green gensupport.JSONFloat64 `json:"green"`
		Red   gensupport.JSONFloat64 `json:"red"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Alpha = float64(s1.Alpha)
	s.Blue = float64(s1.Blue)
	s.Green = float64(s1.Green)
	s.Red = float64(s1.Red)
	return nil
}

// Date: Represents a whole or partial calendar date, such as a birthday. The
// time of day and time zone are either specified elsewhere or are
// insignificant. The date is relative to the Gregorian Calendar. This can
// represent one of the following: * A full date, with non-zero year, month,
// and day values. * A month and day, with a zero year (for example, an
// anniversary). * A year on its own, with a zero month and a zero day. * A
// year and month, with a zero day (for example, a credit card expiration
// date). Related types: * google.type.TimeOfDay * google.type.DateTime *
// google.protobuf.Timestamp
type Date struct {
	// Day: Day of a month. Must be from 1 to 31 and valid for the year and month,
	// or 0 to specify a year by itself or a year and month where the day isn't
	// significant.
	Day int64 `json:"day,omitempty"`
	// Month: Month of a year. Must be from 1 to 12, or 0 to specify a year without
	// a month and day.
	Month int64 `json:"month,omitempty"`
	// Year: Year of the date. Must be from 1 to 9999, or 0 to specify a date
	// without a year.
	Year int64 `json:"year,omitempty"`
	// ForceSendFields is a list of field names (e.g. "Day") to unconditionally
	// include in API requests. By default, fields with empty or default values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "Day") to include in API requests
	// with the JSON null value. By default, fields with empty values are omitted
	// from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s Date) MarshalJSON() ([]byte, error) {
	type NoMethod Date
	return gensupport.MarshalJSON(NoMethod(s), s.ForceSendFields, s.NullFields)
}

// DayInfo: This object contains the daily forecast information for each day
// requested.
type DayInfo struct {
	// Date: The date in UTC at which the pollen forecast data is represented.
	Date *Date `json:"date,omitempty"`
	// PlantInfo: This list will include up to 15 pollen species affecting the
	// location specified in the request.
	PlantInfo []*PlantInfo `json:"plantInfo,omitempty"`
	// PollenTypeInfo: This list will include up to three pollen types (GRASS,
	// WEED, TREE) affecting the location specified in the request.
	PollenTypeInfo []*PollenTypeInfo `json:"pollenTypeInfo,omitempty"`
	// ForceSendFields is a list of field names (e.g. "Date") to unconditionally
	// include in API requests. By default, fields with empty or default values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "Date") to include in API requests
	// with the JSON null value. By default, fields with empty values are omitted
	// from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s DayInfo) MarshalJSON() ([]byte, error) {
	type NoMethod DayInfo
	return gensupport.MarshalJSON(NoMethod(s), s.ForceSendFields, s.NullFields)
}

// HttpBody: Message that represents an arbitrary HTTP body. It should only be
// used for payload formats that can't be represented as JSON, such as raw
// binary or an HTML page. This message can be used both in streaming and
// non-streaming API methods in the request as well as the response. It can be
// used as a top-level request field, which is convenient if one wants to
// extract parameters from either the URL or HTTP template into the request
// fields and also want access to the raw HTTP body. Example: message
// GetResourceRequest { // A unique request id. string request_id = 1; // The
// raw HTTP body is bound to this field. google.api.HttpBody http_body = 2; }
// service ResourceService { rpc GetResource(GetResourceRequest) returns
// (google.api.HttpBody); rpc UpdateResource(google.api.HttpBody) returns
// (google.protobuf.Empty); } Example with streaming methods: service
// CaldavService { rpc GetCalendar(stream google.api.HttpBody) returns (stream
// google.api.HttpBody); rpc UpdateCalendar(stream google.api.HttpBody) returns
// (stream google.api.HttpBody); } Use of this type only changes how the
// request and response bodies are handled, all other features will continue to
// work unchanged.
type HttpBody struct {
	// ContentType: The HTTP Content-Type header value specifying the content type
	// of the body.
	ContentType string `json:"contentType,omitempty"`
	// Data: The HTTP request/response body as raw binary.
	Data string `json:"data,omitempty"`
	// Extensions: Application specific response metadata. Must be set in the first
	// response for streaming APIs.
	Extensions []googleapi.RawMessage `json:"extensions,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the server.
	googleapi.ServerResponse `json:"-"`
	// ForceSendFields is a list of field names (e.g. "ContentType") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "ContentType") to include in API
	// requests with the JSON null value. By default, fields with empty values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s HttpBody) MarshalJSON() ([]byte, error) {
	type NoMethod HttpBody
	return gensupport.MarshalJSON(NoMethod(s), s.ForceSendFields, s.NullFields)
}

// IndexInfo: This object contains data representing specific pollen index
// value, category and description.
type IndexInfo struct {
	// Category: Text classification of index numerical score interpretation. The
	// index consists of six categories: * 0: "None" * 1: "Very low" * 2: "Low" *
	// 3: "Moderate" * 4: "High" * 5: "Very high
	Category string `json:"category,omitempty"`
	// Code: The index's code. This field represents the index for programming
	// purposes by using snake cases instead of spaces. Example: "UPI".
	//
	// Possible values:
	//   "INDEX_UNSPECIFIED" - Unspecified index.
	//   "UPI" - Universal Pollen Index.
	Code string `json:"code,omitempty"`
	// Color: The color used to represent the Pollen Index numeric score.
	Color *Color `json:"color,omitempty"`
	// DisplayName: A human readable representation of the index name. Example:
	// "Universal Pollen Index".
	DisplayName string `json:"displayName,omitempty"`
	// IndexDescription: Textual explanation of current index level.
	IndexDescription string `json:"indexDescription,omitempty"`
	// Value: The index's numeric score. Numeric range is between 0 and 5.
	Value int64 `json:"value,omitempty"`
	// ForceSendFields is a list of field names (e.g. "Category") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "Category") to include in API
	// requests with the JSON null value. By default, fields with empty values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s IndexInfo) MarshalJSON() ([]byte, error) {
	type NoMethod IndexInfo
	return gensupport.MarshalJSON(NoMethod(s), s.ForceSendFields, s.NullFields)
}

type LookupForecastResponse struct {
	// DailyInfo: Required. This object contains the daily forecast information for
	// each day requested.
	DailyInfo []*DayInfo `json:"dailyInfo,omitempty"`
	// NextPageToken: Optional. The token to retrieve the next page.
	NextPageToken string `json:"nextPageToken,omitempty"`
	// RegionCode: The ISO_3166-1 alpha-2 code of the country/region corresponding
	// to the location provided in the request. This field might be omitted from
	// the response if the location provided in the request resides in a disputed
	// territory.
	RegionCode string `json:"regionCode,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the server.
	googleapi.ServerResponse `json:"-"`
	// ForceSendFields is a list of field names (e.g. "DailyInfo") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "DailyInfo") to include in API
	// requests with the JSON null value. By default, fields with empty values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s LookupForecastResponse) MarshalJSON() ([]byte, error) {
	type NoMethod LookupForecastResponse
	return gensupport.MarshalJSON(NoMethod(s), s.ForceSendFields, s.NullFields)
}

// PlantDescription: Contains general information about plants, including
// details on their seasonality, special shapes and colors, information about
// allergic cross-reactions, and plant photos.
type PlantDescription struct {
	// CrossReaction: Textual description of pollen cross reaction plants. Example:
	// Alder, Hazel, Hornbeam, Beech, Willow, and Oak pollen.
	CrossReaction string `json:"crossReaction,omitempty"`
	// Family: A human readable representation of the plant family name. Example:
	// "Betulaceae (the Birch family)".
	Family string `json:"family,omitempty"`
	// Picture: Link to the picture of the plant.
	Picture string `json:"picture,omitempty"`
	// PictureCloseup: Link to a closeup picture of the plant.
	PictureCloseup string `json:"pictureCloseup,omitempty"`
	// Season: Textual list of explanations of seasons where the pollen is active.
	// Example: "Late winter, spring".
	Season string `json:"season,omitempty"`
	// SpecialColors: Textual description of the plants' colors of leaves, bark,
	// flowers or seeds that helps identify the plant.
	SpecialColors string `json:"specialColors,omitempty"`
	// SpecialShapes: Textual description of the plants' shapes of leaves, bark,
	// flowers or seeds that helps identify the plant.
	SpecialShapes string `json:"specialShapes,omitempty"`
	// Type: The plant's pollen type. For example: "GRASS". A list of all available
	// codes could be found here.
	//
	// Possible values:
	//   "POLLEN_TYPE_UNSPECIFIED" - Unspecified plant type.
	//   "GRASS" - Grass pollen type.
	//   "TREE" - Tree pollen type.
	//   "WEED" - Weed pollen type.
	Type string `json:"type,omitempty"`
	// ForceSendFields is a list of field names (e.g. "CrossReaction") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "CrossReaction") to include in API
	// requests with the JSON null value. By default, fields with empty values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s PlantDescription) MarshalJSON() ([]byte, error) {
	type NoMethod PlantDescription
	return gensupport.MarshalJSON(NoMethod(s), s.ForceSendFields, s.NullFields)
}

// PlantInfo: This object contains the daily information on specific plant.
type PlantInfo struct {
	// Code: The plant code name. For example: "COTTONWOOD". A list of all
	// available codes could be found here.
	//
	// Possible values:
	//   "PLANT_UNSPECIFIED" - Unspecified plant code.
	//   "ALDER" - Alder is classified as a tree pollen type.
	//   "ASH" - Ash is classified as a tree pollen type.
	//   "BIRCH" - Birch is classified as a tree pollen type.
	//   "COTTONWOOD" - Cottonwood is classified as a tree pollen type.
	//   "ELM" - Elm is classified as a tree pollen type.
	//   "MAPLE" - Maple is classified as a tree pollen type.
	//   "OLIVE" - Olive is classified as a tree pollen type.
	//   "JUNIPER" - Juniper is classified as a tree pollen type.
	//   "OAK" - Oak is classified as a tree pollen type.
	//   "PINE" - Pine is classified as a tree pollen type.
	//   "CYPRESS_PINE" - Cypress pine is classified as a tree pollen type.
	//   "HAZEL" - Hazel is classified as a tree pollen type.
	//   "GRAMINALES" - Graminales is classified as a grass pollen type.
	//   "RAGWEED" - Ragweed is classified as a weed pollen type.
	//   "MUGWORT" - Mugwort is classified as a weed pollen type.
	//   "JAPANESE_CEDAR" - Japanese cedar is classified as a tree pollen type.
	//   "JAPANESE_CYPRESS" - Japanese cypress is classified as a tree pollen type.
	Code string `json:"code,omitempty"`
	// DisplayName: A human readable representation of the plant name. Example:
	// “Cottonwood".
	DisplayName string `json:"displayName,omitempty"`
	// InSeason: Indication of either the plant is in season or not.
	InSeason bool `json:"inSeason,omitempty"`
	// IndexInfo: This object contains data representing specific pollen index
	// value, category and description.
	IndexInfo *IndexInfo `json:"indexInfo,omitempty"`
	// PlantDescription: Contains general information about plants, including
	// details on their seasonality, special shapes and colors, information about
	// allergic cross-reactions, and plant photos.
	PlantDescription *PlantDescription `json:"plantDescription,omitempty"`
	// ForceSendFields is a list of field names (e.g. "Code") to unconditionally
	// include in API requests. By default, fields with empty or default values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "Code") to include in API requests
	// with the JSON null value. By default, fields with empty values are omitted
	// from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s PlantInfo) MarshalJSON() ([]byte, error) {
	type NoMethod PlantInfo
	return gensupport.MarshalJSON(NoMethod(s), s.ForceSendFields, s.NullFields)
}

// PollenTypeInfo: This object contains the pollen type index and health
// recommendation information on specific pollen type.
type PollenTypeInfo struct {
	// Code: The pollen type's code name. For example: "GRASS"
	//
	// Possible values:
	//   "POLLEN_TYPE_UNSPECIFIED" - Unspecified plant type.
	//   "GRASS" - Grass pollen type.
	//   "TREE" - Tree pollen type.
	//   "WEED" - Weed pollen type.
	Code string `json:"code,omitempty"`
	// DisplayName: A human readable representation of the pollen type name.
	// Example: "Grass"
	DisplayName string `json:"displayName,omitempty"`
	// HealthRecommendations: Textual list of explanations, related to health
	// insights based on the current pollen levels.
	HealthRecommendations []string `json:"healthRecommendations,omitempty"`
	// InSeason: Indication whether the plant is in season or not.
	InSeason bool `json:"inSeason,omitempty"`
	// IndexInfo: Contains the Universal Pollen Index (UPI) data for the pollen
	// type.
	IndexInfo *IndexInfo `json:"indexInfo,omitempty"`
	// ForceSendFields is a list of field names (e.g. "Code") to unconditionally
	// include in API requests. By default, fields with empty or default values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "Code") to include in API requests
	// with the JSON null value. By default, fields with empty values are omitted
	// from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s PollenTypeInfo) MarshalJSON() ([]byte, error) {
	type NoMethod PollenTypeInfo
	return gensupport.MarshalJSON(NoMethod(s), s.ForceSendFields, s.NullFields)
}

type ForecastLookupCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Lookup: Returns up to 5 days of daily pollen information in more than 65
// countries, up to 1km resolution.
func (r *ForecastService) Lookup() *ForecastLookupCall {
	c := &ForecastLookupCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// Days sets the optional parameter "days": Required. A number that indicates
// how many forecast days to request (minimum value 1, maximum value is 5).
func (c *ForecastLookupCall) Days(days int64) *ForecastLookupCall {
	c.urlParams_.Set("days", fmt.Sprint(days))
	return c
}

// LanguageCode sets the optional parameter "languageCode": Allows the client
// to choose the language for the response. If data cannot be provided for that
// language, the API uses the closest match. Allowed values rely on the IETF
// BCP-47 standard. The default value is "en".
func (c *ForecastLookupCall) LanguageCode(languageCode string) *ForecastLookupCall {
	c.urlParams_.Set("languageCode", languageCode)
	return c
}

// LocationLatitude sets the optional parameter "location.latitude": The
// latitude in degrees. It must be in the range [-90.0, +90.0].
func (c *ForecastLookupCall) LocationLatitude(locationLatitude float64) *ForecastLookupCall {
	c.urlParams_.Set("location.latitude", fmt.Sprint(locationLatitude))
	return c
}

// LocationLongitude sets the optional parameter "location.longitude": The
// longitude in degrees. It must be in the range [-180.0, +180.0].
func (c *ForecastLookupCall) LocationLongitude(locationLongitude float64) *ForecastLookupCall {
	c.urlParams_.Set("location.longitude", fmt.Sprint(locationLongitude))
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number of daily
// info records to return per page. The default and max value is 5, indicating
// 5 days of data.
func (c *ForecastLookupCall) PageSize(pageSize int64) *ForecastLookupCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A page token received
// from a previous daily call. It is used to retrieve the subsequent page. Note
// that when providing a value for the page token, all other request parameters
// provided must match the previous call that provided the page token.
func (c *ForecastLookupCall) PageToken(pageToken string) *ForecastLookupCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// PlantsDescription sets the optional parameter "plantsDescription": Contains
// general information about plants, including details on their seasonality,
// special shapes and colors, information about allergic cross-reactions, and
// plant photos. The default value is "true".
func (c *ForecastLookupCall) PlantsDescription(plantsDescription bool) *ForecastLookupCall {
	c.urlParams_.Set("plantsDescription", fmt.Sprint(plantsDescription))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more
// details.
func (c *ForecastLookupCall) Fields(s ...googleapi.Field) *ForecastLookupCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets an optional parameter which makes the operation fail if the
// object's ETag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *ForecastLookupCall) IfNoneMatch(entityTag string) *ForecastLookupCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method.
func (c *ForecastLookupCall) Context(ctx context.Context) *ForecastLookupCall {
	c.ctx_ = ctx
	return c
}

// Header returns a http.Header that can be modified by the caller to add
// headers to the request.
func (c *ForecastLookupCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ForecastLookupCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := gensupport.SetHeaders(c.s.userAgent(), "", c.header_)
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/forecast:lookup")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, nil)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	c.s.logger.DebugContext(c.ctx_, "api request", "serviceName", apiName, "rpcName", "pollen.forecast.lookup", "request", internallog.HTTPRequest(req, nil))
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "pollen.forecast.lookup" call.
// Any non-2xx status code is an error. Response headers are in either
// *LookupForecastResponse.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified was
// returned.
func (c *ForecastLookupCall) Do(opts ...googleapi.CallOption) (*LookupForecastResponse, error) {
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
	ret := &LookupForecastResponse{
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
	c.s.logger.DebugContext(c.ctx_, "api response", "serviceName", apiName, "rpcName", "pollen.forecast.lookup", "response", internallog.HTTPResponse(res, b))
	return ret, nil
}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ForecastLookupCall) Pages(ctx context.Context, f func(*LookupForecastResponse) error) error {
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

type MapTypesHeatmapTilesLookupHeatmapTileCall struct {
	s            *Service
	mapType      string
	zoom         int64
	x            int64
	y            int64
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// LookupHeatmapTile: Returns a byte array containing the data of the tile PNG
// image.
//
//   - mapType: The type of the pollen heatmap. Defines the combination of pollen
//     type and index that the map will graphically represent.
//   - x: Defines the east-west point in the requested tile.
//   - y: Defines the north-south point in the requested tile.
//   - zoom: The map's zoom level. Defines how large or small the contents of a
//     map appear in a map view. * Zoom level 0 is the entire world in a single
//     tile. * Zoom level 1 is the entire world in 4 tiles. * Zoom level 2 is the
//     entire world in 16 tiles. * Zoom level 16 is the entire world in 65,536
//     tiles. Allowed values: 0-16.
func (r *MapTypesHeatmapTilesService) LookupHeatmapTile(mapType string, zoom int64, x int64, y int64) *MapTypesHeatmapTilesLookupHeatmapTileCall {
	c := &MapTypesHeatmapTilesLookupHeatmapTileCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.mapType = mapType
	c.zoom = zoom
	c.x = x
	c.y = y
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more
// details.
func (c *MapTypesHeatmapTilesLookupHeatmapTileCall) Fields(s ...googleapi.Field) *MapTypesHeatmapTilesLookupHeatmapTileCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets an optional parameter which makes the operation fail if the
// object's ETag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *MapTypesHeatmapTilesLookupHeatmapTileCall) IfNoneMatch(entityTag string) *MapTypesHeatmapTilesLookupHeatmapTileCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method.
func (c *MapTypesHeatmapTilesLookupHeatmapTileCall) Context(ctx context.Context) *MapTypesHeatmapTilesLookupHeatmapTileCall {
	c.ctx_ = ctx
	return c
}

// Header returns a http.Header that can be modified by the caller to add
// headers to the request.
func (c *MapTypesHeatmapTilesLookupHeatmapTileCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *MapTypesHeatmapTilesLookupHeatmapTileCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := gensupport.SetHeaders(c.s.userAgent(), "", c.header_)
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/mapTypes/{mapType}/heatmapTiles/{zoom}/{x}/{y}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, nil)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"mapType": c.mapType,
		"zoom":    strconv.FormatInt(c.zoom, 10),
		"x":       strconv.FormatInt(c.x, 10),
		"y":       strconv.FormatInt(c.y, 10),
	})
	c.s.logger.DebugContext(c.ctx_, "api request", "serviceName", apiName, "rpcName", "pollen.mapTypes.heatmapTiles.lookupHeatmapTile", "request", internallog.HTTPRequest(req, nil))
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "pollen.mapTypes.heatmapTiles.lookupHeatmapTile" call.
// Any non-2xx status code is an error. Response headers are in either
// *HttpBody.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was returned.
func (c *MapTypesHeatmapTilesLookupHeatmapTileCall) Do(opts ...googleapi.CallOption) (*HttpBody, error) {
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
	ret := &HttpBody{
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
	c.s.logger.DebugContext(c.ctx_, "api response", "serviceName", apiName, "rpcName", "pollen.mapTypes.heatmapTiles.lookupHeatmapTile", "response", internallog.HTTPResponse(res, b))
	return ret, nil
}
