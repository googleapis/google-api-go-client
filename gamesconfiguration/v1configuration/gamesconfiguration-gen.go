// Package gamesconfiguration provides access to the Google Play Game Services Publishing API.
//
// See https://developers.google.com/games/services
//
// Usage example:
//
//   import "google.golang.org/api/gamesconfiguration/v1configuration"
//   ...
//   gamesconfigurationService, err := gamesconfiguration.New(oauthHttpClient)
package gamesconfiguration

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

const apiId = "gamesConfiguration:v1configuration"
const apiName = "gamesConfiguration"
const apiVersion = "v1configuration"
const basePath = "https://www.googleapis.com/games/v1configuration/"

// OAuth2 scopes used by this API.
const (
	// View and manage your Google Play Developer account
	AndroidpublisherScope = "https://www.googleapis.com/auth/androidpublisher"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.AchievementConfigurations = NewAchievementConfigurationsService(s)
	s.ImageConfigurations = NewImageConfigurationsService(s)
	s.LeaderboardConfigurations = NewLeaderboardConfigurationsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	AchievementConfigurations *AchievementConfigurationsService

	ImageConfigurations *ImageConfigurationsService

	LeaderboardConfigurations *LeaderboardConfigurationsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewAchievementConfigurationsService(s *Service) *AchievementConfigurationsService {
	rs := &AchievementConfigurationsService{s: s}
	return rs
}

type AchievementConfigurationsService struct {
	s *Service
}

func NewImageConfigurationsService(s *Service) *ImageConfigurationsService {
	rs := &ImageConfigurationsService{s: s}
	return rs
}

type ImageConfigurationsService struct {
	s *Service
}

func NewLeaderboardConfigurationsService(s *Service) *LeaderboardConfigurationsService {
	rs := &LeaderboardConfigurationsService{s: s}
	return rs
}

type LeaderboardConfigurationsService struct {
	s *Service
}

// AchievementConfiguration: This is a JSON template for an achievement
// configuration resource.
type AchievementConfiguration struct {
	// AchievementType: The type of the achievement.
	// Possible values are:
	// - "STANDARD" - Achievement is either locked or unlocked.
	// - "INCREMENTAL" - Achievement is incremental.
	AchievementType string `json:"achievementType,omitempty"`

	// Draft: The draft data of the achievement.
	Draft *AchievementConfigurationDetail `json:"draft,omitempty"`

	// Id: The ID of the achievement.
	Id string `json:"id,omitempty"`

	// InitialState: The initial state of the achievement.
	// Possible values are:
	// - "HIDDEN" - Achievement is hidden.
	// - "REVEALED" - Achievement is revealed.
	// - "UNLOCKED" - Achievement is unlocked.
	InitialState string `json:"initialState,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string gamesConfiguration#achievementConfiguration.
	Kind string `json:"kind,omitempty"`

	// Published: The read-only published data of the achievement.
	Published *AchievementConfigurationDetail `json:"published,omitempty"`

	// StepsToUnlock: Steps to unlock. Only applicable to incremental
	// achievements.
	StepsToUnlock int64 `json:"stepsToUnlock,omitempty"`

	// Token: The token for this resource.
	Token string `json:"token,omitempty"`
}

// AchievementConfigurationDetail: This is a JSON template for an
// achievement configuration detail.
type AchievementConfigurationDetail struct {
	// Description: Localized strings for the achievement description.
	Description *LocalizedStringBundle `json:"description,omitempty"`

	// IconUrl: The icon url of this achievement. Writes to this field are
	// ignored.
	IconUrl string `json:"iconUrl,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string gamesConfiguration#achievementConfigurationDetail.
	Kind string `json:"kind,omitempty"`

	// Name: Localized strings for the achievement name.
	Name *LocalizedStringBundle `json:"name,omitempty"`

	// PointValue: Point value for the achievement.
	PointValue int64 `json:"pointValue,omitempty"`

	// SortRank: The sort rank of this achievement. Writes to this field are
	// ignored.
	SortRank int64 `json:"sortRank,omitempty"`
}

// AchievementConfigurationListResponse: This is a JSON template for a
// ListConfigurations response.
type AchievementConfigurationListResponse struct {
	// Items: The achievement configurations.
	Items []*AchievementConfiguration `json:"items,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#achievementConfigurationListResponse.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The pagination token for the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

// GamesNumberAffixConfiguration: This is a JSON template for a number
// affix resource.
type GamesNumberAffixConfiguration struct {
	// Few: When the language requires special treatment of "small" numbers
	// (as with 2, 3, and 4 in Czech; or numbers ending 2, 3, or 4 but not
	// 12, 13, or 14 in Polish).
	Few *LocalizedStringBundle `json:"few,omitempty"`

	// Many: When the language requires special treatment of "large" numbers
	// (as with numbers ending 11-99 in Maltese).
	Many *LocalizedStringBundle `json:"many,omitempty"`

	// One: When the language requires special treatment of numbers like one
	// (as with the number 1 in English and most other languages; in
	// Russian, any number ending in 1 but not ending in 11 is in this
	// class).
	One *LocalizedStringBundle `json:"one,omitempty"`

	// Other: When the language does not require special treatment of the
	// given quantity (as with all numbers in Chinese, or 42 in English).
	Other *LocalizedStringBundle `json:"other,omitempty"`

	// Two: When the language requires special treatment of numbers like two
	// (as with 2 in Welsh, or 102 in Slovenian).
	Two *LocalizedStringBundle `json:"two,omitempty"`

	// Zero: When the language requires special treatment of the number 0
	// (as in Arabic).
	Zero *LocalizedStringBundle `json:"zero,omitempty"`
}

// GamesNumberFormatConfiguration: This is a JSON template for a number
// format resource.
type GamesNumberFormatConfiguration struct {
	// CurrencyCode: The curreny code string. Only used for CURRENCY format
	// type.
	CurrencyCode string `json:"currencyCode,omitempty"`

	// NumDecimalPlaces: The number of decimal places for number. Only used
	// for NUMERIC format type.
	NumDecimalPlaces int64 `json:"numDecimalPlaces,omitempty"`

	// NumberFormatType: The formatting for the number.
	// Possible values are:
	// - "NUMERIC" - Numbers are formatted to have no digits or a fixed
	// number of digits after the decimal point according to locale. An
	// optional custom unit can be added.
	// - "TIME_DURATION" - Numbers are formatted to hours, minutes and
	// seconds.
	// - "CURRENCY" - Numbers are formatted to currency according to locale.
	NumberFormatType string `json:"numberFormatType,omitempty"`

	// Suffix: An optional suffix for the NUMERIC format type. These strings
	// follow the same  plural rules as all Android string resources.
	Suffix *GamesNumberAffixConfiguration `json:"suffix,omitempty"`
}

// ImageConfiguration: This is a JSON template for an image
// configuration resource.
type ImageConfiguration struct {
	// ImageType: The image type for the image.
	ImageType string `json:"imageType,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string gamesConfiguration#imageConfiguration.
	Kind string `json:"kind,omitempty"`

	// ResourceId: The resource ID of resource which the image belongs to.
	ResourceId string `json:"resourceId,omitempty"`

	// Url: The url for this image.
	Url string `json:"url,omitempty"`
}

// LeaderboardConfiguration: This is a JSON template for an leaderboard
// configuration resource.
type LeaderboardConfiguration struct {
	// Draft: The draft data of the leaderboard.
	Draft *LeaderboardConfigurationDetail `json:"draft,omitempty"`

	// Id: The ID of the leaderboard.
	Id string `json:"id,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string gamesConfiguration#leaderboardConfiguration.
	Kind string `json:"kind,omitempty"`

	// Published: The read-only published data of the leaderboard.
	Published *LeaderboardConfigurationDetail `json:"published,omitempty"`

	// ScoreMax: Maximum score that can be posted to this leaderboard.
	ScoreMax int64 `json:"scoreMax,omitempty,string"`

	// ScoreMin: Minimum score that can be posted to this leaderboard.
	ScoreMin int64 `json:"scoreMin,omitempty,string"`

	// ScoreOrder: The type of the leaderboard.
	// Possible values are:
	// - "LARGER_IS_BETTER" - Larger scores posted are ranked higher.
	// - "SMALLER_IS_BETTER" - Smaller scores posted are ranked higher.
	ScoreOrder string `json:"scoreOrder,omitempty"`

	// Token: The token for this resource.
	Token string `json:"token,omitempty"`
}

// LeaderboardConfigurationDetail: This is a JSON template for a
// leaderboard configuration detail.
type LeaderboardConfigurationDetail struct {
	// IconUrl: The icon url of this leaderboard. Writes to this field are
	// ignored.
	IconUrl string `json:"iconUrl,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string gamesConfiguration#leaderboardConfigurationDetail.
	Kind string `json:"kind,omitempty"`

	// Name: Localized strings for the leaderboard name.
	Name *LocalizedStringBundle `json:"name,omitempty"`

	// ScoreFormat: The score formatting for the leaderboard.
	ScoreFormat *GamesNumberFormatConfiguration `json:"scoreFormat,omitempty"`

	// SortRank: The sort rank of this leaderboard. Writes to this field are
	// ignored.
	SortRank int64 `json:"sortRank,omitempty"`
}

// LeaderboardConfigurationListResponse: This is a JSON template for a
// ListConfigurations response.
type LeaderboardConfigurationListResponse struct {
	// Items: The leaderboard configurations.
	Items []*LeaderboardConfiguration `json:"items,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#leaderboardConfigurationListResponse.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The pagination token for the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

// LocalizedString: This is a JSON template for a localized string
// resource.
type LocalizedString struct {
	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string gamesConfiguration#localizedString.
	Kind string `json:"kind,omitempty"`

	// Locale: The locale string.
	Locale string `json:"locale,omitempty"`

	// Value: The string value.
	Value string `json:"value,omitempty"`
}

// LocalizedStringBundle: This is a JSON template for a localized string
// bundle resource.
type LocalizedStringBundle struct {
	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string gamesConfiguration#localizedStringBundle.
	Kind string `json:"kind,omitempty"`

	// Translations: The locale strings.
	Translations []*LocalizedString `json:"translations,omitempty"`
}

// method id "gamesConfiguration.achievementConfigurations.delete":

type AchievementConfigurationsDeleteCall struct {
	s             *Service
	achievementId string
	opt_          map[string]interface{}
}

// Delete: Delete the achievement configuration with the given ID.
func (r *AchievementConfigurationsService) Delete(achievementId string) *AchievementConfigurationsDeleteCall {
	c := &AchievementConfigurationsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.achievementId = achievementId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AchievementConfigurationsDeleteCall) Fields(s ...googleapi.Field) *AchievementConfigurationsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *AchievementConfigurationsDeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "achievements/{achievementId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"achievementId": c.achievementId,
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
	//   "description": "Delete the achievement configuration with the given ID.",
	//   "httpMethod": "DELETE",
	//   "id": "gamesConfiguration.achievementConfigurations.delete",
	//   "parameterOrder": [
	//     "achievementId"
	//   ],
	//   "parameters": {
	//     "achievementId": {
	//       "description": "The ID of the achievement used by this method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "achievements/{achievementId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ]
	// }

}

// method id "gamesConfiguration.achievementConfigurations.get":

type AchievementConfigurationsGetCall struct {
	s             *Service
	achievementId string
	opt_          map[string]interface{}
}

// Get: Retrieves the metadata of the achievement configuration with the
// given ID.
func (r *AchievementConfigurationsService) Get(achievementId string) *AchievementConfigurationsGetCall {
	c := &AchievementConfigurationsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.achievementId = achievementId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AchievementConfigurationsGetCall) Fields(s ...googleapi.Field) *AchievementConfigurationsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *AchievementConfigurationsGetCall) Do() (*AchievementConfiguration, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "achievements/{achievementId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"achievementId": c.achievementId,
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
	var ret *AchievementConfiguration
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the metadata of the achievement configuration with the given ID.",
	//   "httpMethod": "GET",
	//   "id": "gamesConfiguration.achievementConfigurations.get",
	//   "parameterOrder": [
	//     "achievementId"
	//   ],
	//   "parameters": {
	//     "achievementId": {
	//       "description": "The ID of the achievement used by this method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "achievements/{achievementId}",
	//   "response": {
	//     "$ref": "AchievementConfiguration"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ]
	// }

}

// method id "gamesConfiguration.achievementConfigurations.insert":

type AchievementConfigurationsInsertCall struct {
	s                        *Service
	applicationId            string
	achievementconfiguration *AchievementConfiguration
	opt_                     map[string]interface{}
}

// Insert: Insert a new achievement configuration in this application.
func (r *AchievementConfigurationsService) Insert(applicationId string, achievementconfiguration *AchievementConfiguration) *AchievementConfigurationsInsertCall {
	c := &AchievementConfigurationsInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.applicationId = applicationId
	c.achievementconfiguration = achievementconfiguration
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AchievementConfigurationsInsertCall) Fields(s ...googleapi.Field) *AchievementConfigurationsInsertCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *AchievementConfigurationsInsertCall) Do() (*AchievementConfiguration, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.achievementconfiguration)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "applications/{applicationId}/achievements")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"applicationId": c.applicationId,
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
	var ret *AchievementConfiguration
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Insert a new achievement configuration in this application.",
	//   "httpMethod": "POST",
	//   "id": "gamesConfiguration.achievementConfigurations.insert",
	//   "parameterOrder": [
	//     "applicationId"
	//   ],
	//   "parameters": {
	//     "applicationId": {
	//       "description": "The application ID from the Google Play developer console.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "applications/{applicationId}/achievements",
	//   "request": {
	//     "$ref": "AchievementConfiguration"
	//   },
	//   "response": {
	//     "$ref": "AchievementConfiguration"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ]
	// }

}

// method id "gamesConfiguration.achievementConfigurations.list":

type AchievementConfigurationsListCall struct {
	s             *Service
	applicationId string
	opt_          map[string]interface{}
}

// List: Returns a list of the achievement configurations in this
// application.
func (r *AchievementConfigurationsService) List(applicationId string) *AchievementConfigurationsListCall {
	c := &AchievementConfigurationsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.applicationId = applicationId
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of resource configurations to return in the response, used for
// paging. For any response, the actual number of resources returned may
// be less than the specified maxResults.
func (c *AchievementConfigurationsListCall) MaxResults(maxResults int64) *AchievementConfigurationsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": The token returned
// by the previous request.
func (c *AchievementConfigurationsListCall) PageToken(pageToken string) *AchievementConfigurationsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AchievementConfigurationsListCall) Fields(s ...googleapi.Field) *AchievementConfigurationsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *AchievementConfigurationsListCall) Do() (*AchievementConfigurationListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "applications/{applicationId}/achievements")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"applicationId": c.applicationId,
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
	var ret *AchievementConfigurationListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns a list of the achievement configurations in this application.",
	//   "httpMethod": "GET",
	//   "id": "gamesConfiguration.achievementConfigurations.list",
	//   "parameterOrder": [
	//     "applicationId"
	//   ],
	//   "parameters": {
	//     "applicationId": {
	//       "description": "The application ID from the Google Play developer console.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of resource configurations to return in the response, used for paging. For any response, the actual number of resources returned may be less than the specified maxResults.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "200",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The token returned by the previous request.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "applications/{applicationId}/achievements",
	//   "response": {
	//     "$ref": "AchievementConfigurationListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ]
	// }

}

// method id "gamesConfiguration.achievementConfigurations.patch":

type AchievementConfigurationsPatchCall struct {
	s                        *Service
	achievementId            string
	achievementconfiguration *AchievementConfiguration
	opt_                     map[string]interface{}
}

// Patch: Update the metadata of the achievement configuration with the
// given ID. This method supports patch semantics.
func (r *AchievementConfigurationsService) Patch(achievementId string, achievementconfiguration *AchievementConfiguration) *AchievementConfigurationsPatchCall {
	c := &AchievementConfigurationsPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.achievementId = achievementId
	c.achievementconfiguration = achievementconfiguration
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AchievementConfigurationsPatchCall) Fields(s ...googleapi.Field) *AchievementConfigurationsPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *AchievementConfigurationsPatchCall) Do() (*AchievementConfiguration, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.achievementconfiguration)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "achievements/{achievementId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"achievementId": c.achievementId,
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
	var ret *AchievementConfiguration
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update the metadata of the achievement configuration with the given ID. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "gamesConfiguration.achievementConfigurations.patch",
	//   "parameterOrder": [
	//     "achievementId"
	//   ],
	//   "parameters": {
	//     "achievementId": {
	//       "description": "The ID of the achievement used by this method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "achievements/{achievementId}",
	//   "request": {
	//     "$ref": "AchievementConfiguration"
	//   },
	//   "response": {
	//     "$ref": "AchievementConfiguration"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ]
	// }

}

// method id "gamesConfiguration.achievementConfigurations.update":

type AchievementConfigurationsUpdateCall struct {
	s                        *Service
	achievementId            string
	achievementconfiguration *AchievementConfiguration
	opt_                     map[string]interface{}
}

// Update: Update the metadata of the achievement configuration with the
// given ID.
func (r *AchievementConfigurationsService) Update(achievementId string, achievementconfiguration *AchievementConfiguration) *AchievementConfigurationsUpdateCall {
	c := &AchievementConfigurationsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.achievementId = achievementId
	c.achievementconfiguration = achievementconfiguration
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AchievementConfigurationsUpdateCall) Fields(s ...googleapi.Field) *AchievementConfigurationsUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *AchievementConfigurationsUpdateCall) Do() (*AchievementConfiguration, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.achievementconfiguration)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "achievements/{achievementId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"achievementId": c.achievementId,
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
	var ret *AchievementConfiguration
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update the metadata of the achievement configuration with the given ID.",
	//   "httpMethod": "PUT",
	//   "id": "gamesConfiguration.achievementConfigurations.update",
	//   "parameterOrder": [
	//     "achievementId"
	//   ],
	//   "parameters": {
	//     "achievementId": {
	//       "description": "The ID of the achievement used by this method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "achievements/{achievementId}",
	//   "request": {
	//     "$ref": "AchievementConfiguration"
	//   },
	//   "response": {
	//     "$ref": "AchievementConfiguration"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ]
	// }

}

// method id "gamesConfiguration.imageConfigurations.upload":

type ImageConfigurationsUploadCall struct {
	s          *Service
	resourceId string
	imageType  string
	opt_       map[string]interface{}
	media_     io.Reader
	resumable_ googleapi.SizeReaderAt
	mediaType_ string
	ctx_       context.Context
	protocol_  string
}

// Upload: Uploads an image for a resource with the given ID and image
// type.
func (r *ImageConfigurationsService) Upload(resourceId string, imageType string) *ImageConfigurationsUploadCall {
	c := &ImageConfigurationsUploadCall{s: r.s, opt_: make(map[string]interface{})}
	c.resourceId = resourceId
	c.imageType = imageType
	return c
}

// Media specifies the media to upload in a single chunk.
// At most one of Media and ResumableMedia may be set.
func (c *ImageConfigurationsUploadCall) Media(r io.Reader) *ImageConfigurationsUploadCall {
	c.media_ = r
	c.protocol_ = "multipart"
	return c
}

// ResumableMedia specifies the media to upload in chunks and can be cancelled with ctx.
// At most one of Media and ResumableMedia may be set.
// mediaType identifies the MIME media type of the upload, such as "image/png".
// If mediaType is "", it will be auto-detected.
func (c *ImageConfigurationsUploadCall) ResumableMedia(ctx context.Context, r io.ReaderAt, size int64, mediaType string) *ImageConfigurationsUploadCall {
	c.ctx_ = ctx
	c.resumable_ = io.NewSectionReader(r, 0, size)
	c.mediaType_ = mediaType
	c.protocol_ = "resumable"
	return c
}

// ProgressUpdater provides a callback function that will be called after every chunk.
// It should be a low-latency function in order to not slow down the upload operation.
// This should only be called when using ResumableMedia (as opposed to Media).
func (c *ImageConfigurationsUploadCall) ProgressUpdater(pu googleapi.ProgressUpdater) *ImageConfigurationsUploadCall {
	c.opt_["progressUpdater"] = pu
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ImageConfigurationsUploadCall) Fields(s ...googleapi.Field) *ImageConfigurationsUploadCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ImageConfigurationsUploadCall) Do() (*ImageConfiguration, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "images/{resourceId}/imageType/{imageType}")
	var progressUpdater_ googleapi.ProgressUpdater
	if v, ok := c.opt_["progressUpdater"]; ok {
		if pu, ok := v.(googleapi.ProgressUpdater); ok {
			progressUpdater_ = pu
		}
	}
	if c.media_ != nil || c.resumable_ != nil {
		urls = strings.Replace(urls, "https://www.googleapis.com/", "https://www.googleapis.com/upload/", 1)
		params.Set("uploadType", c.protocol_)
	}
	urls += "?" + params.Encode()
	body = new(bytes.Buffer)
	ctype := "application/json"
	if c.protocol_ != "resumable" {
		var cancel func()
		cancel, _ = googleapi.ConditionallyIncludeMedia(c.media_, &body, &ctype)
		if cancel != nil {
			defer cancel()
		}
	}
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"resourceId": c.resourceId,
		"imageType":  c.imageType,
	})
	if c.protocol_ == "resumable" {
		if c.mediaType_ == "" {
			c.mediaType_ = googleapi.DetectMediaType(c.resumable_)
		}
		req.Header.Set("X-Upload-Content-Type", c.mediaType_)
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	} else {
		req.Header.Set("Content-Type", ctype)
	}
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	if c.protocol_ == "resumable" {
		loc := res.Header.Get("Location")
		rx := &googleapi.ResumableUpload{
			Client:        c.s.client,
			UserAgent:     c.s.userAgent(),
			URI:           loc,
			Media:         c.resumable_,
			MediaType:     c.mediaType_,
			ContentLength: c.resumable_.Size(),
			Callback:      progressUpdater_,
		}
		res, err = rx.Upload(c.ctx_)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
	}
	var ret *ImageConfiguration
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Uploads an image for a resource with the given ID and image type.",
	//   "httpMethod": "POST",
	//   "id": "gamesConfiguration.imageConfigurations.upload",
	//   "mediaUpload": {
	//     "accept": [
	//       "image/*"
	//     ],
	//     "maxSize": "15MB",
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/games/v1configuration/images/{resourceId}/imageType/{imageType}"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/games/v1configuration/images/{resourceId}/imageType/{imageType}"
	//       }
	//     }
	//   },
	//   "parameterOrder": [
	//     "resourceId",
	//     "imageType"
	//   ],
	//   "parameters": {
	//     "imageType": {
	//       "description": "Selects which image in a resource for this method.",
	//       "enum": [
	//         "ACHIEVEMENT_ICON",
	//         "LEADERBOARD_ICON"
	//       ],
	//       "enumDescriptions": [
	//         "The icon image for an achievement resource.",
	//         "The icon image for a leaderboard resource."
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "resourceId": {
	//       "description": "The ID of the resource used by this method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "images/{resourceId}/imageType/{imageType}",
	//   "response": {
	//     "$ref": "ImageConfiguration"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ],
	//   "supportsMediaUpload": true
	// }

}

// method id "gamesConfiguration.leaderboardConfigurations.delete":

type LeaderboardConfigurationsDeleteCall struct {
	s             *Service
	leaderboardId string
	opt_          map[string]interface{}
}

// Delete: Delete the leaderboard configuration with the given ID.
func (r *LeaderboardConfigurationsService) Delete(leaderboardId string) *LeaderboardConfigurationsDeleteCall {
	c := &LeaderboardConfigurationsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.leaderboardId = leaderboardId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LeaderboardConfigurationsDeleteCall) Fields(s ...googleapi.Field) *LeaderboardConfigurationsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *LeaderboardConfigurationsDeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "leaderboards/{leaderboardId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"leaderboardId": c.leaderboardId,
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
	//   "description": "Delete the leaderboard configuration with the given ID.",
	//   "httpMethod": "DELETE",
	//   "id": "gamesConfiguration.leaderboardConfigurations.delete",
	//   "parameterOrder": [
	//     "leaderboardId"
	//   ],
	//   "parameters": {
	//     "leaderboardId": {
	//       "description": "The ID of the leaderboard.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "leaderboards/{leaderboardId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ]
	// }

}

// method id "gamesConfiguration.leaderboardConfigurations.get":

type LeaderboardConfigurationsGetCall struct {
	s             *Service
	leaderboardId string
	opt_          map[string]interface{}
}

// Get: Retrieves the metadata of the leaderboard configuration with the
// given ID.
func (r *LeaderboardConfigurationsService) Get(leaderboardId string) *LeaderboardConfigurationsGetCall {
	c := &LeaderboardConfigurationsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.leaderboardId = leaderboardId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LeaderboardConfigurationsGetCall) Fields(s ...googleapi.Field) *LeaderboardConfigurationsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *LeaderboardConfigurationsGetCall) Do() (*LeaderboardConfiguration, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "leaderboards/{leaderboardId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"leaderboardId": c.leaderboardId,
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
	var ret *LeaderboardConfiguration
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the metadata of the leaderboard configuration with the given ID.",
	//   "httpMethod": "GET",
	//   "id": "gamesConfiguration.leaderboardConfigurations.get",
	//   "parameterOrder": [
	//     "leaderboardId"
	//   ],
	//   "parameters": {
	//     "leaderboardId": {
	//       "description": "The ID of the leaderboard.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "leaderboards/{leaderboardId}",
	//   "response": {
	//     "$ref": "LeaderboardConfiguration"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ]
	// }

}

// method id "gamesConfiguration.leaderboardConfigurations.insert":

type LeaderboardConfigurationsInsertCall struct {
	s                        *Service
	applicationId            string
	leaderboardconfiguration *LeaderboardConfiguration
	opt_                     map[string]interface{}
}

// Insert: Insert a new leaderboard configuration in this application.
func (r *LeaderboardConfigurationsService) Insert(applicationId string, leaderboardconfiguration *LeaderboardConfiguration) *LeaderboardConfigurationsInsertCall {
	c := &LeaderboardConfigurationsInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.applicationId = applicationId
	c.leaderboardconfiguration = leaderboardconfiguration
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LeaderboardConfigurationsInsertCall) Fields(s ...googleapi.Field) *LeaderboardConfigurationsInsertCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *LeaderboardConfigurationsInsertCall) Do() (*LeaderboardConfiguration, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.leaderboardconfiguration)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "applications/{applicationId}/leaderboards")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"applicationId": c.applicationId,
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
	var ret *LeaderboardConfiguration
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Insert a new leaderboard configuration in this application.",
	//   "httpMethod": "POST",
	//   "id": "gamesConfiguration.leaderboardConfigurations.insert",
	//   "parameterOrder": [
	//     "applicationId"
	//   ],
	//   "parameters": {
	//     "applicationId": {
	//       "description": "The application ID from the Google Play developer console.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "applications/{applicationId}/leaderboards",
	//   "request": {
	//     "$ref": "LeaderboardConfiguration"
	//   },
	//   "response": {
	//     "$ref": "LeaderboardConfiguration"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ]
	// }

}

// method id "gamesConfiguration.leaderboardConfigurations.list":

type LeaderboardConfigurationsListCall struct {
	s             *Service
	applicationId string
	opt_          map[string]interface{}
}

// List: Returns a list of the leaderboard configurations in this
// application.
func (r *LeaderboardConfigurationsService) List(applicationId string) *LeaderboardConfigurationsListCall {
	c := &LeaderboardConfigurationsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.applicationId = applicationId
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of resource configurations to return in the response, used for
// paging. For any response, the actual number of resources returned may
// be less than the specified maxResults.
func (c *LeaderboardConfigurationsListCall) MaxResults(maxResults int64) *LeaderboardConfigurationsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": The token returned
// by the previous request.
func (c *LeaderboardConfigurationsListCall) PageToken(pageToken string) *LeaderboardConfigurationsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LeaderboardConfigurationsListCall) Fields(s ...googleapi.Field) *LeaderboardConfigurationsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *LeaderboardConfigurationsListCall) Do() (*LeaderboardConfigurationListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "applications/{applicationId}/leaderboards")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"applicationId": c.applicationId,
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
	var ret *LeaderboardConfigurationListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns a list of the leaderboard configurations in this application.",
	//   "httpMethod": "GET",
	//   "id": "gamesConfiguration.leaderboardConfigurations.list",
	//   "parameterOrder": [
	//     "applicationId"
	//   ],
	//   "parameters": {
	//     "applicationId": {
	//       "description": "The application ID from the Google Play developer console.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of resource configurations to return in the response, used for paging. For any response, the actual number of resources returned may be less than the specified maxResults.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "200",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The token returned by the previous request.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "applications/{applicationId}/leaderboards",
	//   "response": {
	//     "$ref": "LeaderboardConfigurationListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ]
	// }

}

// method id "gamesConfiguration.leaderboardConfigurations.patch":

type LeaderboardConfigurationsPatchCall struct {
	s                        *Service
	leaderboardId            string
	leaderboardconfiguration *LeaderboardConfiguration
	opt_                     map[string]interface{}
}

// Patch: Update the metadata of the leaderboard configuration with the
// given ID. This method supports patch semantics.
func (r *LeaderboardConfigurationsService) Patch(leaderboardId string, leaderboardconfiguration *LeaderboardConfiguration) *LeaderboardConfigurationsPatchCall {
	c := &LeaderboardConfigurationsPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.leaderboardId = leaderboardId
	c.leaderboardconfiguration = leaderboardconfiguration
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LeaderboardConfigurationsPatchCall) Fields(s ...googleapi.Field) *LeaderboardConfigurationsPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *LeaderboardConfigurationsPatchCall) Do() (*LeaderboardConfiguration, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.leaderboardconfiguration)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "leaderboards/{leaderboardId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"leaderboardId": c.leaderboardId,
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
	var ret *LeaderboardConfiguration
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update the metadata of the leaderboard configuration with the given ID. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "gamesConfiguration.leaderboardConfigurations.patch",
	//   "parameterOrder": [
	//     "leaderboardId"
	//   ],
	//   "parameters": {
	//     "leaderboardId": {
	//       "description": "The ID of the leaderboard.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "leaderboards/{leaderboardId}",
	//   "request": {
	//     "$ref": "LeaderboardConfiguration"
	//   },
	//   "response": {
	//     "$ref": "LeaderboardConfiguration"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ]
	// }

}

// method id "gamesConfiguration.leaderboardConfigurations.update":

type LeaderboardConfigurationsUpdateCall struct {
	s                        *Service
	leaderboardId            string
	leaderboardconfiguration *LeaderboardConfiguration
	opt_                     map[string]interface{}
}

// Update: Update the metadata of the leaderboard configuration with the
// given ID.
func (r *LeaderboardConfigurationsService) Update(leaderboardId string, leaderboardconfiguration *LeaderboardConfiguration) *LeaderboardConfigurationsUpdateCall {
	c := &LeaderboardConfigurationsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.leaderboardId = leaderboardId
	c.leaderboardconfiguration = leaderboardconfiguration
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LeaderboardConfigurationsUpdateCall) Fields(s ...googleapi.Field) *LeaderboardConfigurationsUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *LeaderboardConfigurationsUpdateCall) Do() (*LeaderboardConfiguration, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.leaderboardconfiguration)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "leaderboards/{leaderboardId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"leaderboardId": c.leaderboardId,
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
	var ret *LeaderboardConfiguration
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update the metadata of the leaderboard configuration with the given ID.",
	//   "httpMethod": "PUT",
	//   "id": "gamesConfiguration.leaderboardConfigurations.update",
	//   "parameterOrder": [
	//     "leaderboardId"
	//   ],
	//   "parameters": {
	//     "leaderboardId": {
	//       "description": "The ID of the leaderboard.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "leaderboards/{leaderboardId}",
	//   "request": {
	//     "$ref": "LeaderboardConfiguration"
	//   },
	//   "response": {
	//     "$ref": "LeaderboardConfiguration"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ]
	// }

}
