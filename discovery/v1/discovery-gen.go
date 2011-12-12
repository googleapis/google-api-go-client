// Package discovery provides access to the APIs Discovery Service.
//
// See http://code.google.com/apis/discovery
//
// Usage example:
//
//   import "google-api-go-client.googlecode.com/hg/discovery/v1"
//   ...
//   discoveryService, err := discovery.New(oauthHttpClient)
package discovery

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

const apiId = "discovery:v1"
const apiName = "discovery"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/discovery/v1/"

func New(client *http.Client) (*Service, os.Error) {
	if client == nil {
		return nil, os.NewError("client is nil")
	}
	s := &Service{client: client}
	s.Apis = &ApisService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	Apis *ApisService
}

type ApisService struct {
	s *Service
}

type JsonSchema struct {
	// Format: An additional regular expression or key that helps constrain
	// the value. For more details see:
	// http://tools.ietf.org/html/draft-zyp-json-schema-03#section-5.23
	Format string `json:"format,omitempty"`

	// Location: Whether this parameter goes in the query or the path for
	// REST requests.
	Location string `json:"location,omitempty"`

	// Pattern: The regular expression this parameter must conform to.
	Pattern string `json:"pattern,omitempty"`

	// Items: If this is a schema for an array, this property is the schema
	// for each element in the array.
	Items *JsonSchema `json:"items,omitempty"`

	// Ref: A reference to another schema. The value of this property is the
	// "id" of another schema.
	Ref string `json:"$ref,omitempty"`

	// EnumDescriptions: The descriptions for the enums. Each position maps
	// to the corresponding value in the "enum" array.
	EnumDescriptions []string `json:"enumDescriptions,omitempty"`

	// Properties: If this is a schema for an object, list the schema for
	// each property of this object.
	Properties *JsonSchemaProperties `json:"properties,omitempty"`

	// Id: Unique identifier for this schema.
	Id string `json:"id,omitempty"`

	// Default: The default value of this property (if one exists).
	Default string `json:"default,omitempty"`

	// Required: Whether the parameter is required.
	Required bool `json:"required,omitempty"`

	// Repeated: Whether this parameter may appear multiple times.
	Repeated bool `json:"repeated,omitempty"`

	// Enum: Values this parameter may take (if it is an enum).
	Enum []string `json:"enum,omitempty"`

	// Maximum: The maximum value of this parameter.
	Maximum string `json:"maximum,omitempty"`

	// Minimum: The minimum value of this parameter.
	Minimum string `json:"minimum,omitempty"`

	// Type: The value type for this schema. A list of values can be found
	// here: http://tools.ietf.org/html/draft-zyp-json-schema-03#section-5.1
	Type string `json:"type,omitempty"`

	// Description: A description of this object.
	Description string `json:"description,omitempty"`

	// AdditionalProperties: If this is a schema for an object, this
	// property is the schema for any additional properties with dynamic
	// keys on this object.
	AdditionalProperties *JsonSchema `json:"additionalProperties,omitempty"`
}

type RestDescriptionIcons struct {
	// X16: The url of the 16x16 icon.
	X16 string `json:"x16,omitempty"`

	// X32: The url of the 32x32 icon.
	X32 string `json:"x32,omitempty"`
}

type RestMethodMediaUpload struct {
	// Protocols: Supported upload protocols.
	Protocols *RestMethodMediaUploadProtocols `json:"protocols,omitempty"`

	// MaxSize: Maximum size of a media upload, such as "1MB", "2GB" or
	// "3TB".
	MaxSize string `json:"maxSize,omitempty"`

	// Accept: MIME Media Ranges for acceptable media uploads to this
	// method.
	Accept []string `json:"accept,omitempty"`
}

type RestDescription struct {
	// Protocol: The protocol described by this document.
	Protocol string `json:"protocol,omitempty"`

	// Labels: Labels for the status of this API, such as labs or
	// deprecated.
	Labels []string `json:"labels,omitempty"`

	// Features: A list of supported features for this API.
	Features []string `json:"features,omitempty"`

	// Name: The name of this API.
	Name string `json:"name,omitempty"`

	// DocumentationLink: A link to human readable documentation for the
	// API.
	DocumentationLink string `json:"documentationLink,omitempty"`

	// Icons: Links to 16x16 and 32x32 icons representing the API.
	Icons *RestDescriptionIcons `json:"icons,omitempty"`

	// Version: The version of this API.
	Version string `json:"version,omitempty"`

	// Kind: The kind for this response.
	Kind string `json:"kind,omitempty"`

	// Schemas: The schemas for this API.
	Schemas *RestDescriptionSchemas `json:"schemas,omitempty"`

	// BasePath: The base URI path for REST requests.
	BasePath string `json:"basePath,omitempty"`

	// Auth: Authentication information.
	Auth *RestDescriptionAuth `json:"auth,omitempty"`

	// Id: The id of this API.
	Id string `json:"id,omitempty"`

	// Parameters: Common parameters that apply across all apis.
	Parameters *RestDescriptionParameters `json:"parameters,omitempty"`

	// Title: The title of this API.
	Title string `json:"title,omitempty"`

	// Resources: The resources in this API.
	Resources *RestDescriptionResources `json:"resources,omitempty"`

	// Methods: API-level methods for this API.
	Methods *RestDescriptionMethods `json:"methods,omitempty"`

	// Description: The description of this API.
	Description string `json:"description,omitempty"`
}

type RestDescriptionMethods struct {

}

type RestMethodRequest struct {
	// Ref: Schema ID for the request schema.
	Ref string `json:"$ref,omitempty"`
}

type RestResource struct {
	// Resources: Sub-resources on this resource.
	Resources *RestResourceResources `json:"resources,omitempty"`

	// Methods: Methods on this resource.
	Methods *RestResourceMethods `json:"methods,omitempty"`
}

type RestMethod struct {
	// Scopes: OAuth 2.0 scopes applicable to this method.
	Scopes []string `json:"scopes,omitempty"`

	// Response: The schema for the response.
	Response *RestMethodResponse `json:"response,omitempty"`

	// Path: The URI path of this REST method. Should be used in conjunction
	// with the basePath property at the api-level.
	Path string `json:"path,omitempty"`

	// Id: A unique ID for this method. This property can be used to match
	// methods between different versions of Discovery.
	Id string `json:"id,omitempty"`

	// Request: The schema for the request.
	Request *RestMethodRequest `json:"request,omitempty"`

	// Parameters: Details for all parameters in this method.
	Parameters *RestMethodParameters `json:"parameters,omitempty"`

	// ParameterOrder: Ordered list of required parameters, serves as a hint
	// to clients on how to structure their method signatures. The array is
	// ordered such that the "most-significant" parameter appears first.
	ParameterOrder []string `json:"parameterOrder,omitempty"`

	// HttpMethod: HTTP method used by this method.
	HttpMethod string `json:"httpMethod,omitempty"`

	// Description: Description of this method.
	Description string `json:"description,omitempty"`

	// MediaUpload: Media upload parameters.
	MediaUpload *RestMethodMediaUpload `json:"mediaUpload,omitempty"`
}

type RestDescriptionAuthOauth2 struct {
	// Scopes: Available OAuth 2.0 scopes.
	Scopes *RestDescriptionAuthOauth2Scopes `json:"scopes,omitempty"`
}

type RestDescriptionAuth struct {
	// Oauth2: OAuth 2.0 authentication information.
	Oauth2 *RestDescriptionAuthOauth2 `json:"oauth2,omitempty"`
}

type RestDescriptionSchemas struct {

}

type JsonSchemaProperties struct {

}

type RestDescriptionAuthOauth2Scopes struct {

}

type RestDescriptionResources struct {

}

type RestResourceResources struct {

}

type DirectoryListItems struct {
	// Labels: Labels for the status of this API, such as labs or
	// deprecated.
	Labels []string `json:"labels,omitempty"`

	// Name: The name of the API.
	Name string `json:"name,omitempty"`

	// Preferred: True if this version is the preferred version to use.
	Preferred bool `json:"preferred,omitempty"`

	// DocumentationLink: A link to human readable documentation for the
	// API.
	DocumentationLink string `json:"documentationLink,omitempty"`

	// Icons: Links to 16x16 and 32x32 icons representing the API.
	Icons *DirectoryListItemsIcons `json:"icons,omitempty"`

	// Version: The version of the API.
	Version string `json:"version,omitempty"`

	// Kind: The kind for this response.
	Kind string `json:"kind,omitempty"`

	// Id: The id of this API.
	Id string `json:"id,omitempty"`

	// Title: The title of this API.
	Title string `json:"title,omitempty"`

	// DiscoveryLink: A link to the discovery document.
	DiscoveryLink string `json:"discoveryLink,omitempty"`

	// Description: The description of this API.
	Description string `json:"description,omitempty"`
}

type RestMethodMediaUploadProtocols struct {
	// Resumable: Supports the Resumable Media Upload protocol.
	Resumable *RestMethodMediaUploadProtocolsResumable `json:"resumable,omitempty"`

	// Simple: Supports uploading as a single HTTP request.
	Simple *RestMethodMediaUploadProtocolsSimple `json:"simple,omitempty"`
}

type RestMethodParameters struct {

}

type RestDescriptionParameters struct {

}

type DirectoryList struct {
	// Items: The individual directory entries. One entry per api/version
	// pair.
	Items []*DirectoryListItems `json:"items,omitempty"`

	// Kind: The kind for this response.
	Kind string `json:"kind,omitempty"`
}

type RestResourceMethods struct {

}

type DirectoryListItemsIcons struct {
	// X16: The url of the 16x16 icon.
	X16 string `json:"x16,omitempty"`

	// X32: The url of the 32x32 icon.
	X32 string `json:"x32,omitempty"`
}

type RestMethodMediaUploadProtocolsSimple struct {
	// Path: The URI path to be used for upload. Should be used in
	// conjunction with the basePath property at the api-level.
	Path string `json:"path,omitempty"`

	// Multipart: True if this endpoint supports upload multipart media.
	Multipart bool `json:"multipart,omitempty"`
}

type RestMethodResponse struct {
	// Ref: Schema ID for the response schema.
	Ref string `json:"$ref,omitempty"`
}

type RestMethodMediaUploadProtocolsResumable struct {
	// Path: The URI path to be used for upload. Should be used in
	// conjunction with the basePath property at the api-level.
	Path string `json:"path,omitempty"`

	// Multipart: True if this endpoint supports uploading multipart media.
	Multipart bool `json:"multipart,omitempty"`
}

// method id "discovery.apis.list":

type ApisListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Retrieve the list of APIs supported at this endpoint.
func (r *ApisService) List() *ApisListCall {
	c := &ApisListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// Name sets the optional parameter "name": Only include APIs with the
// given name.
func (c *ApisListCall) Name(name string) *ApisListCall {
	c.opt_["name"] = name
	return c
}

// Preferred sets the optional parameter "preferred": Return only the
// preferred version of an API.
func (c *ApisListCall) Preferred(preferred bool) *ApisListCall {
	c.opt_["preferred"] = preferred
	return c
}

// Label sets the optional parameter "label": Only include APIs with a
// matching label, such as 'graduated' or 'labs'.
func (c *ApisListCall) Label(label string) *ApisListCall {
	c.opt_["label"] = label
	return c
}

func (c *ApisListCall) Do() (*DirectoryList, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["name"]; ok {
		params.Set("name", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["preferred"]; ok {
		params.Set("preferred", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["label"]; ok {
		params.Set("label", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/discovery/v1/", "apis")
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
	ret := new(DirectoryList)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieve the list of APIs supported at this endpoint.",
	//   "httpMethod": "GET",
	//   "id": "discovery.apis.list",
	//   "parameters": {
	//     "label": {
	//       "description": "Only include APIs with a matching label, such as 'graduated' or 'labs'.",
	//       "enum": [
	//         "deprecated",
	//         "graduated",
	//         "labs"
	//       ],
	//       "enumDescriptions": [
	//         "APIs that have been deprecated.",
	//         "Supported APIs that have graduated from labs.",
	//         "APIs that are experimental"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "Only include APIs with the given name.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "preferred": {
	//       "default": "false",
	//       "description": "Return only the preferred version of an API.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "apis",
	//   "response": {
	//     "$ref": "DirectoryList"
	//   }
	// }

}

// method id "discovery.apis.getRest":

type ApisGetRestCall struct {
	s       *Service
	api     string
	version string
	opt_    map[string]interface{}
}

// GetRest: Retrieve the description of a particular version of an api.
func (r *ApisService) GetRest(api string, version string) *ApisGetRestCall {
	c := &ApisGetRestCall{s: r.s, opt_: make(map[string]interface{})}
	c.api = api
	c.version = version
	return c
}

func (c *ApisGetRestCall) Do() (*RestDescription, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/discovery/v1/", "apis/{api}/{version}/rest")
	urls = strings.Replace(urls, "{api}", cleanPathString(c.api), 1)
	urls = strings.Replace(urls, "{version}", cleanPathString(c.version), 1)
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
	ret := new(RestDescription)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieve the description of a particular version of an api.",
	//   "httpMethod": "GET",
	//   "id": "discovery.apis.getRest",
	//   "parameterOrder": [
	//     "api",
	//     "version"
	//   ],
	//   "parameters": {
	//     "api": {
	//       "description": "The name of the API.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "version": {
	//       "description": "The version of the API.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "apis/{api}/{version}/rest",
	//   "response": {
	//     "$ref": "RestDescription"
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
