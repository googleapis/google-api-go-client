// Package latitude provides access to the Google Latitude API.
//
// See http://code.google.com/apis/latitude/v1/using_rest.html
//
// Usage example:
//
//   import "code.google.com/p/google-api-go-client/latitude/v1"
//   ...
//   latitudeService, err := latitude.New(oauthHttpClient)
package latitude

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

const apiId = "latitude:v1"
const apiName = "latitude"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/latitude/v1/"

// OAuth2 scopes used by this API.
const (
	// Manage your best-available location
	LatitudeCurrentBestScope = "https://www.googleapis.com/auth/latitude.current.best"

	// Manage your city-level location and location history
	LatitudeAllCityScope = "https://www.googleapis.com/auth/latitude.all.city"

	// Manage your best-available location and location history
	LatitudeAllBestScope = "https://www.googleapis.com/auth/latitude.all.best"

	// Manage your city-level location
	LatitudeCurrentCityScope = "https://www.googleapis.com/auth/latitude.current.city"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.CurrentLocation = &CurrentLocationService{s: s}
	s.Location = &LocationService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	CurrentLocation *CurrentLocationService

	Location *LocationService
}

type CurrentLocationService struct {
	s *Service
}

type LocationService struct {
	s *Service
}

type Location struct {
	// Heading: Direction of travel of the user when this location was
	// recorded. In degrees, clockwise relative to true north. Optional.
	Heading interface{} `json:"heading,omitempty"`

	// Speed: Ground speed of the user at the time this location was
	// recorded, in meters per second. Non-negative. Optional.
	Speed interface{} `json:"speed,omitempty"`

	// TimestampMs: Timestamp of the Location Resource, in milliseconds
	// since the epoch (UTC). This is also the Location Resource's unique
	// id.
	TimestampMs interface{} `json:"timestampMs,omitempty"`

	// AltitudeAccuracy: Accuracy of the altitude value, in meters.
	// Optional.
	AltitudeAccuracy interface{} `json:"altitudeAccuracy,omitempty"`

	// Longitude: Longitude of the location, in decimal degrees.
	Longitude interface{} `json:"longitude,omitempty"`

	// Altitude: Altitude of the location, in meters. Optional.
	Altitude interface{} `json:"altitude,omitempty"`

	// Kind: Kind of this item.
	Kind string `json:"kind,omitempty"`

	// Latitude: Latitude of the location, in decimal degrees.
	Latitude interface{} `json:"latitude,omitempty"`

	// Accuracy: Accuracy of the latitude and longitude coordinates, in
	// non-negative meters. Optional.
	Accuracy interface{} `json:"accuracy,omitempty"`

	// ActivityId: Unique ID of the Buzz message that corresponds to the
	// check-in associated with this location. Available only for check-in
	// locations. Optional.
	ActivityId interface{} `json:"activityId,omitempty"`
}

type LatitudeCurrentlocationResourceJson struct {
	Location
}

type LocationFeed struct {
	Items []*Location `json:"items,omitempty"`

	Kind string `json:"kind,omitempty"`
}

// method id "latitude.currentLocation.get":

type CurrentLocationGetCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// Get: Returns the authenticated user's current location.
func (r *CurrentLocationService) Get() *CurrentLocationGetCall {
	c := &CurrentLocationGetCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// Granularity sets the optional parameter "granularity": Granularity of
// the requested location.
func (c *CurrentLocationGetCall) Granularity(granularity string) *CurrentLocationGetCall {
	c.opt_["granularity"] = granularity
	return c
}

func (c *CurrentLocationGetCall) Do() (*LatitudeCurrentlocationResourceJson, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["granularity"]; ok {
		params.Set("granularity", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/latitude/v1/", "currentLocation")
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
	ret := new(LatitudeCurrentlocationResourceJson)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the authenticated user's current location.",
	//   "httpMethod": "GET",
	//   "id": "latitude.currentLocation.get",
	//   "parameters": {
	//     "granularity": {
	//       "description": "Granularity of the requested location.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "currentLocation",
	//   "response": {
	//     "$ref": "LatitudeCurrentlocationResourceJson"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/latitude.all.best",
	//     "https://www.googleapis.com/auth/latitude.all.city",
	//     "https://www.googleapis.com/auth/latitude.current.best",
	//     "https://www.googleapis.com/auth/latitude.current.city"
	//   ]
	// }

}

// method id "latitude.currentLocation.delete":

type CurrentLocationDeleteCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// Delete: Deletes the authenticated user's current location.
func (r *CurrentLocationService) Delete() *CurrentLocationDeleteCall {
	c := &CurrentLocationDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

func (c *CurrentLocationDeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/latitude/v1/", "currentLocation")
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
	//   "description": "Deletes the authenticated user's current location.",
	//   "httpMethod": "DELETE",
	//   "id": "latitude.currentLocation.delete",
	//   "path": "currentLocation",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/latitude.all.best",
	//     "https://www.googleapis.com/auth/latitude.all.city",
	//     "https://www.googleapis.com/auth/latitude.current.best",
	//     "https://www.googleapis.com/auth/latitude.current.city"
	//   ]
	// }

}

// method id "latitude.currentLocation.insert":

type CurrentLocationInsertCall struct {
	s                                   *Service
	latitudecurrentlocationresourcejson *LatitudeCurrentlocationResourceJson
	opt_                                map[string]interface{}
}

// Insert: Updates or creates the user's current location.
func (r *CurrentLocationService) Insert(latitudecurrentlocationresourcejson *LatitudeCurrentlocationResourceJson) *CurrentLocationInsertCall {
	c := &CurrentLocationInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.latitudecurrentlocationresourcejson = latitudecurrentlocationresourcejson
	return c
}

func (c *CurrentLocationInsertCall) Do() (*LatitudeCurrentlocationResourceJson, error) {
	var body io.Reader = nil
	body, err := googleapi.WithDataWrapper.JSONReader(c.latitudecurrentlocationresourcejson)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/latitude/v1/", "currentLocation")
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
	ret := new(LatitudeCurrentlocationResourceJson)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates or creates the user's current location.",
	//   "httpMethod": "POST",
	//   "id": "latitude.currentLocation.insert",
	//   "path": "currentLocation",
	//   "request": {
	//     "$ref": "LatitudeCurrentlocationResourceJson"
	//   },
	//   "response": {
	//     "$ref": "LatitudeCurrentlocationResourceJson"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/latitude.all.best",
	//     "https://www.googleapis.com/auth/latitude.all.city",
	//     "https://www.googleapis.com/auth/latitude.current.best",
	//     "https://www.googleapis.com/auth/latitude.current.city"
	//   ]
	// }

}

// method id "latitude.location.list":

type LocationListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Lists the user's location history.
func (r *LocationService) List() *LocationListCall {
	c := &LocationListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// MinTime sets the optional parameter "min-time": Minimum timestamp of
// locations to return (ms since epoch).
func (c *LocationListCall) MinTime(minTime string) *LocationListCall {
	c.opt_["min-time"] = minTime
	return c
}

// MaxResults sets the optional parameter "max-results": Maximum number
// of locations to return.
func (c *LocationListCall) MaxResults(maxResults string) *LocationListCall {
	c.opt_["max-results"] = maxResults
	return c
}

// Granularity sets the optional parameter "granularity": Granularity of
// the requested locations.
func (c *LocationListCall) Granularity(granularity string) *LocationListCall {
	c.opt_["granularity"] = granularity
	return c
}

// MaxTime sets the optional parameter "max-time": Maximum timestamp of
// locations to return (ms since epoch).
func (c *LocationListCall) MaxTime(maxTime string) *LocationListCall {
	c.opt_["max-time"] = maxTime
	return c
}

func (c *LocationListCall) Do() (*LocationFeed, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["min-time"]; ok {
		params.Set("min-time", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["max-results"]; ok {
		params.Set("max-results", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["granularity"]; ok {
		params.Set("granularity", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["max-time"]; ok {
		params.Set("max-time", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/latitude/v1/", "location")
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
	ret := new(LocationFeed)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the user's location history.",
	//   "httpMethod": "GET",
	//   "id": "latitude.location.list",
	//   "parameters": {
	//     "granularity": {
	//       "description": "Granularity of the requested locations.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "max-results": {
	//       "description": "Maximum number of locations to return.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "max-time": {
	//       "description": "Maximum timestamp of locations to return (ms since epoch).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "min-time": {
	//       "description": "Minimum timestamp of locations to return (ms since epoch).",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "location",
	//   "response": {
	//     "$ref": "LocationFeed"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/latitude.all.best",
	//     "https://www.googleapis.com/auth/latitude.all.city"
	//   ]
	// }

}

// method id "latitude.location.insert":

type LocationInsertCall struct {
	s        *Service
	location *Location
	opt_     map[string]interface{}
}

// Insert: Inserts or updates a location in the user's location history.
func (r *LocationService) Insert(location *Location) *LocationInsertCall {
	c := &LocationInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.location = location
	return c
}

func (c *LocationInsertCall) Do() (*Location, error) {
	var body io.Reader = nil
	body, err := googleapi.WithDataWrapper.JSONReader(c.location)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/latitude/v1/", "location")
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
	ret := new(Location)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Inserts or updates a location in the user's location history.",
	//   "httpMethod": "POST",
	//   "id": "latitude.location.insert",
	//   "path": "location",
	//   "request": {
	//     "$ref": "Location"
	//   },
	//   "response": {
	//     "$ref": "Location"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/latitude.all.best",
	//     "https://www.googleapis.com/auth/latitude.all.city"
	//   ]
	// }

}

// method id "latitude.location.get":

type LocationGetCall struct {
	s          *Service
	locationId string
	opt_       map[string]interface{}
}

// Get: Reads a location from the user's location history.
func (r *LocationService) Get(locationId string) *LocationGetCall {
	c := &LocationGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.locationId = locationId
	return c
}

// Granularity sets the optional parameter "granularity": Granularity of
// the location to return.
func (c *LocationGetCall) Granularity(granularity string) *LocationGetCall {
	c.opt_["granularity"] = granularity
	return c
}

func (c *LocationGetCall) Do() (*Location, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["granularity"]; ok {
		params.Set("granularity", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/latitude/v1/", "location/{locationId}")
	urls = strings.Replace(urls, "{locationId}", cleanPathString(c.locationId), 1)
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
	ret := new(Location)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Reads a location from the user's location history.",
	//   "httpMethod": "GET",
	//   "id": "latitude.location.get",
	//   "parameterOrder": [
	//     "locationId"
	//   ],
	//   "parameters": {
	//     "granularity": {
	//       "description": "Granularity of the location to return.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "locationId": {
	//       "description": "Timestamp of the location to read (ms since epoch).",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "location/{locationId}",
	//   "response": {
	//     "$ref": "Location"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/latitude.all.best",
	//     "https://www.googleapis.com/auth/latitude.all.city"
	//   ]
	// }

}

// method id "latitude.location.delete":

type LocationDeleteCall struct {
	s          *Service
	locationId string
	opt_       map[string]interface{}
}

// Delete: Deletes a location from the user's location history.
func (r *LocationService) Delete(locationId string) *LocationDeleteCall {
	c := &LocationDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.locationId = locationId
	return c
}

func (c *LocationDeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/latitude/v1/", "location/{locationId}")
	urls = strings.Replace(urls, "{locationId}", cleanPathString(c.locationId), 1)
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
	//   "description": "Deletes a location from the user's location history.",
	//   "httpMethod": "DELETE",
	//   "id": "latitude.location.delete",
	//   "parameterOrder": [
	//     "locationId"
	//   ],
	//   "parameters": {
	//     "locationId": {
	//       "description": "Timestamp of the location to delete (ms since epoch).",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "location/{locationId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/latitude.all.best",
	//     "https://www.googleapis.com/auth/latitude.all.city"
	//   ]
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
