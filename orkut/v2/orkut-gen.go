// Package orkut provides access to the Orkut API.
//
// See http://code.google.com/apis/orkut/v2/reference.html
//
// Usage example:
//
//   import "google-api-go-client.googlecode.com/hg/orkut/v2"
//   ...
//   orkutService, err := orkut.New(oauthHttpClient)
package orkut

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

const apiId = "orkut:v2"
const apiName = "orkut"
const apiVersion = "v2"
const basePath = "https://www.googleapis.com/orkut/v2/"

// OAuth2 scopes used by this API.
const (
	// View your Orkut data
	OrkutReadonlyScope = "https://www.googleapis.com/auth/orkut.readonly"

	// Manage your Orkut activity
	OrkutScope = "https://www.googleapis.com/auth/orkut"
)

func New(client *http.Client) (*Service, os.Error) {
	if client == nil {
		return nil, os.NewError("client is nil")
	}
	s := &Service{client: client}
	s.Acl = &AclService{s: s}
	s.ActivityVisibility = &ActivityVisibilityService{s: s}
	s.Scraps = &ScrapsService{s: s}
	s.Badges = &BadgesService{s: s}
	s.Comments = &CommentsService{s: s}
	s.Counters = &CountersService{s: s}
	s.Activities = &ActivitiesService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	Acl *AclService

	ActivityVisibility *ActivityVisibilityService

	Scraps *ScrapsService

	Badges *BadgesService

	Comments *CommentsService

	Counters *CountersService

	Activities *ActivitiesService
}

type AclService struct {
	s *Service
}

type ActivityVisibilityService struct {
	s *Service
}

type ScrapsService struct {
	s *Service
}

type BadgesService struct {
	s *Service
}

type CommentsService struct {
	s *Service
}

type CountersService struct {
	s *Service
}

type ActivitiesService struct {
	s *Service
}

type Activity struct {
	// Links: Links to resources related to this activity.
	Links []*OrkutLinkResource `json:"links,omitempty"`

	// Actor: The person who performed the activity.
	Actor *OrkutAuthorResource `json:"actor,omitempty"`

	// Published: The time at which the activity was initially published.
	Published string `json:"published,omitempty"`

	// Kind: The kind of activity. Always orkut#activity.
	Kind string `json:"kind,omitempty"`

	// Object: The activity's object.
	Object *ActivityObject `json:"object,omitempty"`

	// Updated: The time at which the activity was last updated.
	Updated string `json:"updated,omitempty"`

	// Access: Identifies who has access to see this activity.
	Access *Acl `json:"access,omitempty"`

	// Id: The ID for the activity.
	Id string `json:"id,omitempty"`

	// Title: Title of the activity.
	Title string `json:"title,omitempty"`

	// Verb: This activity's verb, indicating what action was performed.
	// Possible values are:  
	// - add - User added new content to profile or
	// album, e.g. video, photo. 
	// - post - User publish content to the
	// stream, e.g. status, scrap. 
	// - update - User commented on an
	// activity. 
	// - make-friend - User added a new friend. 
	// - birthday -
	// User has a birthday.
	Verb string `json:"verb,omitempty"`
}

type Acl struct {
	// Items: The list of ACL entries.
	Items []*AclItems `json:"items,omitempty"`

	// Kind: Identifies this resource as an access control list. Value:
	// "orkut#acl"
	Kind string `json:"kind,omitempty"`

	// TotalParticipants: The total count of participants of the parent
	// resource.
	TotalParticipants int64 `json:"totalParticipants,omitempty"`

	// Description: Human readable description of the access granted.
	Description string `json:"description,omitempty"`
}

type ActivityObject struct {
	// Items: The list of additional items.
	Items []*OrkutActivityobjectsResource `json:"items,omitempty"`

	// Replies: Comments in reply to this activity.
	Replies *ActivityObjectReplies `json:"replies,omitempty"`

	// ObjectType: The type of the object affected by the activity. Clients
	// can use this information to style the rendered activity object
	// differently depending on the content.
	ObjectType string `json:"objectType,omitempty"`

	// Content: The HTML-formatted content, suitable for display. When
	// updating an activity's content, post the changes to this property,
	// using the value of originalContent as a starting point. If the update
	// is successful, the server adds HTML formatting and responds with this
	// formatted content.
	Content string `json:"content,omitempty"`
}

type AclItems struct {
	// Id: The ID of the entity. For entities of type "person" or "circle",
	// this is the ID of the resource. For other types, this will be unset.
	Id string `json:"id,omitempty"`

	// Type: The type of entity to whom access is granted.
	Type string `json:"type,omitempty"`
}

type ActivityObjectReplies struct {
	// Items: The list of comments.
	Items []*Comment `json:"items,omitempty"`

	// TotalItems: Total number of comments.
	TotalItems uint64 `json:"totalItems,omitempty,string"`

	// Url: URL for the collection of comments in reply to this activity.
	Url string `json:"url,omitempty"`
}

type OrkutActivitypersonResource struct {
	// Image: The person's profile photo. This is adapted from Google+ and
	// was originaly introduced as extra OpenSocial convenience fields.
	Image *OrkutActivitypersonResourceImage `json:"image,omitempty"`

	// Name: An object that encapsulates the individual components of a
	// person's name.
	Name *OrkutActivitypersonResourceName `json:"name,omitempty"`

	// Gender: The person's gender. Values include "male", "female", and
	// "other".
	Gender string `json:"gender,omitempty"`

	// Url: The person's profile url. This is adapted from Google+ and was
	// originaly introduced as extra OpenSocial convenience fields.
	Url string `json:"url,omitempty"`

	// Id: The person's opensocial ID.
	Id string `json:"id,omitempty"`

	// Birthday: The person's date of birth, represented as YYYY-MM-DD.
	Birthday string `json:"birthday,omitempty"`
}

type OrkutActivitypersonResourceImage struct {
	// Url: The URL of the person's profile photo.
	Url string `json:"url,omitempty"`
}

type CommentInReplyTo struct {
	// Href: Link to the post on activity stream being commented.
	Href string `json:"href,omitempty"`

	// Rel: Relationship between the comment and the post on activity stream
	// being commented. Always inReplyTo.
	Rel string `json:"rel,omitempty"`

	// Ref: Unique identifier of the post on activity stream being
	// commented.
	Ref string `json:"ref,omitempty"`

	// Type: Type of the post on activity stream being commented. Always
	// text/html.
	Type string `json:"type,omitempty"`
}

type Counters struct {
	// Items: List of counters retrieved.
	Items []*OrkutCounterResource `json:"items,omitempty"`

	// Kind: Identifies this resource as a collection of counters. Value:
	// "orkut#counters"
	Kind string `json:"kind,omitempty"`
}

type OrkutAuthorResource struct {
	// Image: Image data about the actor.
	Image *OrkutAuthorResourceImage `json:"image,omitempty"`

	// Url: The URL of the author who posted the comment [not yet
	// implemented]
	Url string `json:"url,omitempty"`

	// Id: Unique identifier of the person who posted the comment. This is
	// the person's OpenSocial ID.
	Id string `json:"id,omitempty"`

	// DisplayName: The name of the author, suitable for display.
	DisplayName string `json:"displayName,omitempty"`
}

type BadgeList struct {
	// Items: List of badges retrieved.
	Items []*Badge `json:"items,omitempty"`

	// Kind: Identifies this resource as a collection of badges. Value:
	// "orkut#badgeList"
	Kind string `json:"kind,omitempty"`
}

type OrkutActivityobjectsResource struct {
	// Links: Links to other resources related to this object.
	Links []*OrkutLinkResource `json:"links,omitempty"`

	// Person: The person who is related with this activity, e.g. an Added
	// User.
	Person *OrkutActivitypersonResource `json:"person,omitempty"`

	// ObjectType: The object type.
	ObjectType string `json:"objectType,omitempty"`

	// Content: The HTML-formatted content, suitable for display. When
	// updating an activity's content, post the changes to this property,
	// using the value of originalContent as a starting point. If the update
	// is successful, the server adds HTML formatting and responds with this
	// formatted content.
	Content string `json:"content,omitempty"`

	// Id: The ID for the object.
	Id string `json:"id,omitempty"`

	// DisplayName: The title of the object.
	DisplayName string `json:"displayName,omitempty"`
}

type OrkutAuthorResourceImage struct {
	// Url: A URL that points to a thumbnail photo of the author.
	Url string `json:"url,omitempty"`
}

type Comment struct {
	// Links: List of resources for the comment.
	Links []*OrkutLinkResource `json:"links,omitempty"`

	// Actor: The person who posted the comment.
	Actor *OrkutAuthorResource `json:"actor,omitempty"`

	// Published: The time the comment was initially published, in RFC 3339
	// format.
	Published string `json:"published,omitempty"`

	// Kind: Identifies this resource as a comment. Value: "orkut#comment"
	Kind string `json:"kind,omitempty"`

	// Content: The content of the comment in text/html
	Content string `json:"content,omitempty"`

	// Id: The unique ID for the comment.
	Id string `json:"id,omitempty"`

	// InReplyTo: Link to the original activity where this comment was
	// posted.
	InReplyTo *CommentInReplyTo `json:"inReplyTo,omitempty"`
}

type OrkutActivitypersonResourceName struct {
	// GivenName: The given name (first name) of this person.
	GivenName string `json:"givenName,omitempty"`

	// FamilyName: The family name (last name) of this person.
	FamilyName string `json:"familyName,omitempty"`
}

type OrkutLinkResource struct {
	// Href: URL of the link.
	Href string `json:"href,omitempty"`

	// Title: Title of the link.
	Title string `json:"title,omitempty"`

	// Rel: Relation between the resource and the parent object.
	Rel string `json:"rel,omitempty"`

	// Type: Media type of the link.
	Type string `json:"type,omitempty"`
}

type ActivityList struct {
	// Items: List of activities retrieved.
	Items []*Activity `json:"items,omitempty"`

	// NextPageToken: The value of pageToken query parameter in
	// activities.list request to get the next page, if there are more to
	// retrieve.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Kind: Identifies this resource as a collection of activities. Value:
	// "orkut#activityList"
	Kind string `json:"kind,omitempty"`
}

type OrkutCounterResource struct {
	// Name: The name of the counted collection. Currently supported
	// collections are:  
	// - scraps - The scraps of the user. 
	// - photos - The
	// photos of the user. 
	// - videos - The videos of the user.
	Name string `json:"name,omitempty"`

	// Total: The number of resources on the counted collection.
	Total int64 `json:"total,omitempty"`

	// Link: Link to the collection being counted.
	Link *OrkutLinkResource `json:"link,omitempty"`
}

type Badge struct {
	// SponsorLogo: The URL for the 32x32 badge sponsor logo.
	SponsorLogo string `json:"sponsorLogo,omitempty"`

	// Kind: Identifies this resource as a badge. Value: "orkut#badge"
	Kind string `json:"kind,omitempty"`

	// SponsorName: The name of the badge sponsor, suitable for display.
	SponsorName string `json:"sponsorName,omitempty"`

	// Id: The unique ID for the badge.
	Id int64 `json:"id,omitempty,string"`

	// Caption: The name of the badge, suitable for display.
	Caption string `json:"caption,omitempty"`

	// SponsorUrl: The URL for the badge sponsor.
	SponsorUrl string `json:"sponsorUrl,omitempty"`

	// BadgeSmallLogo: The URL for the 24x24 badge logo.
	BadgeSmallLogo string `json:"badgeSmallLogo,omitempty"`

	// BadgeLargeLogo: The URL for the 64x64 badge logo.
	BadgeLargeLogo string `json:"badgeLargeLogo,omitempty"`

	// Description: The description for the badge, suitable for display.
	Description string `json:"description,omitempty"`
}

type CommentList struct {
	// PreviousPageToken: The value of pageToken query parameter in
	// comments.list request to get the previous page, if there are more to
	// retrieve.
	PreviousPageToken string `json:"previousPageToken,omitempty"`

	// Items: List of comments retrieved.
	Items []*Comment `json:"items,omitempty"`

	// NextPageToken: The value of pageToken query parameter in
	// comments.list request to get the next page, if there are more to
	// retrieve.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Kind: Identifies this resource as a collection of comments. Value:
	// "orkut#commentList"
	Kind string `json:"kind,omitempty"`
}

type Visibility struct {
	// Links: List of resources for the visibility item.
	Links []*OrkutLinkResource `json:"links,omitempty"`

	// Kind: Identifies this resource as a visibility item. Value:
	// "orkut#visibility"
	Kind string `json:"kind,omitempty"`

	// Visibility: The visibility of the resource. Possible values are:  
	// -
	// default: not hidden by the user 
	// - hidden: hidden
	Visibility string `json:"visibility,omitempty"`
}

// method id "orkut.acl.delete":

type AclDeleteCall struct {
	s          *Service
	activityId string
	userId     string
	opt_       map[string]interface{}
}

// Delete: Excludes an element from the ACL of the activity.
func (r *AclService) Delete(activityId string, userId string) *AclDeleteCall {
	c := &AclDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.activityId = activityId
	c.userId = userId
	return c
}

func (c *AclDeleteCall) Do() os.Error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/orkut/v2/", "activities/{activityId}/acl/{userId}")
	urls = strings.Replace(urls, "{activityId}", cleanPathString(c.activityId), 1)
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
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
	//   "description": "Excludes an element from the ACL of the activity.",
	//   "httpMethod": "DELETE",
	//   "id": "orkut.acl.delete",
	//   "parameterOrder": [
	//     "activityId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "activityId": {
	//       "description": "ID of the activity.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user to be removed from the activity.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities/{activityId}/acl/{userId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/orkut"
	//   ]
	// }

}

// method id "orkut.activityVisibility.update":

type ActivityVisibilityUpdateCall struct {
	s          *Service
	activityId string
	visibility *Visibility
	opt_       map[string]interface{}
}

// Update: Updates the visibility of an existing activity.
func (r *ActivityVisibilityService) Update(activityId string, visibility *Visibility) *ActivityVisibilityUpdateCall {
	c := &ActivityVisibilityUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.activityId = activityId
	c.visibility = visibility
	return c
}

func (c *ActivityVisibilityUpdateCall) Do() (*Visibility, os.Error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.visibility)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/orkut/v2/", "activities/{activityId}/visibility")
	urls = strings.Replace(urls, "{activityId}", cleanPathString(c.activityId), 1)
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
	ret := new(Visibility)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the visibility of an existing activity.",
	//   "httpMethod": "PUT",
	//   "id": "orkut.activityVisibility.update",
	//   "parameterOrder": [
	//     "activityId"
	//   ],
	//   "parameters": {
	//     "activityId": {
	//       "description": "ID of the activity.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities/{activityId}/visibility",
	//   "request": {
	//     "$ref": "Visibility"
	//   },
	//   "response": {
	//     "$ref": "Visibility"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/orkut"
	//   ]
	// }

}

// method id "orkut.activityVisibility.get":

type ActivityVisibilityGetCall struct {
	s          *Service
	activityId string
	opt_       map[string]interface{}
}

// Get: Gets the visibility of an existing activity.
func (r *ActivityVisibilityService) Get(activityId string) *ActivityVisibilityGetCall {
	c := &ActivityVisibilityGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.activityId = activityId
	return c
}

func (c *ActivityVisibilityGetCall) Do() (*Visibility, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/orkut/v2/", "activities/{activityId}/visibility")
	urls = strings.Replace(urls, "{activityId}", cleanPathString(c.activityId), 1)
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
	ret := new(Visibility)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the visibility of an existing activity.",
	//   "httpMethod": "GET",
	//   "id": "orkut.activityVisibility.get",
	//   "parameterOrder": [
	//     "activityId"
	//   ],
	//   "parameters": {
	//     "activityId": {
	//       "description": "ID of the activity to get the visibility.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities/{activityId}/visibility",
	//   "response": {
	//     "$ref": "Visibility"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/orkut",
	//     "https://www.googleapis.com/auth/orkut.readonly"
	//   ]
	// }

}

// method id "orkut.activityVisibility.patch":

type ActivityVisibilityPatchCall struct {
	s          *Service
	activityId string
	visibility *Visibility
	opt_       map[string]interface{}
}

// Patch: Updates the visibility of an existing activity. This method
// supports patch semantics.
func (r *ActivityVisibilityService) Patch(activityId string, visibility *Visibility) *ActivityVisibilityPatchCall {
	c := &ActivityVisibilityPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.activityId = activityId
	c.visibility = visibility
	return c
}

func (c *ActivityVisibilityPatchCall) Do() (*Visibility, os.Error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.visibility)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/orkut/v2/", "activities/{activityId}/visibility")
	urls = strings.Replace(urls, "{activityId}", cleanPathString(c.activityId), 1)
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
	ret := new(Visibility)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the visibility of an existing activity. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "orkut.activityVisibility.patch",
	//   "parameterOrder": [
	//     "activityId"
	//   ],
	//   "parameters": {
	//     "activityId": {
	//       "description": "ID of the activity.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities/{activityId}/visibility",
	//   "request": {
	//     "$ref": "Visibility"
	//   },
	//   "response": {
	//     "$ref": "Visibility"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/orkut"
	//   ]
	// }

}

// method id "orkut.scraps.insert":

type ScrapsInsertCall struct {
	s        *Service
	activity *Activity
	opt_     map[string]interface{}
}

// Insert: Creates a new scrap.
func (r *ScrapsService) Insert(activity *Activity) *ScrapsInsertCall {
	c := &ScrapsInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.activity = activity
	return c
}

func (c *ScrapsInsertCall) Do() (*Activity, os.Error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.activity)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/orkut/v2/", "activities/scraps")
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
	ret := new(Activity)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new scrap.",
	//   "httpMethod": "POST",
	//   "id": "orkut.scraps.insert",
	//   "path": "activities/scraps",
	//   "request": {
	//     "$ref": "Activity"
	//   },
	//   "response": {
	//     "$ref": "Activity"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/orkut"
	//   ]
	// }

}

// method id "orkut.badges.list":

type BadgesListCall struct {
	s      *Service
	userId string
	opt_   map[string]interface{}
}

// List: Retrieves the list of visible badges of a user.
func (r *BadgesService) List(userId string) *BadgesListCall {
	c := &BadgesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	return c
}

func (c *BadgesListCall) Do() (*BadgeList, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/orkut/v2/", "people/{userId}/badges")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
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
	ret := new(BadgeList)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the list of visible badges of a user.",
	//   "httpMethod": "GET",
	//   "id": "orkut.badges.list",
	//   "parameterOrder": [
	//     "userId"
	//   ],
	//   "parameters": {
	//     "userId": {
	//       "description": "The id of the user whose badges will be listed. Can be me to refer to caller.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "people/{userId}/badges",
	//   "response": {
	//     "$ref": "BadgeList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/orkut",
	//     "https://www.googleapis.com/auth/orkut.readonly"
	//   ]
	// }

}

// method id "orkut.badges.get":

type BadgesGetCall struct {
	s       *Service
	userId  string
	badgeId int64
	opt_    map[string]interface{}
}

// Get: Retrieves a badge from a user.
func (r *BadgesService) Get(userId string, badgeId int64) *BadgesGetCall {
	c := &BadgesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.badgeId = badgeId
	return c
}

func (c *BadgesGetCall) Do() (*Badge, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/orkut/v2/", "people/{userId}/badges/{badgeId}")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{badgeId}", strconv.Itoa64(c.badgeId), 1)
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
	ret := new(Badge)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a badge from a user.",
	//   "httpMethod": "GET",
	//   "id": "orkut.badges.get",
	//   "parameterOrder": [
	//     "userId",
	//     "badgeId"
	//   ],
	//   "parameters": {
	//     "badgeId": {
	//       "description": "The ID of the badge that will be retrieved.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user whose badges will be listed. Can be me to refer to caller.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "people/{userId}/badges/{badgeId}",
	//   "response": {
	//     "$ref": "Badge"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/orkut",
	//     "https://www.googleapis.com/auth/orkut.readonly"
	//   ]
	// }

}

// method id "orkut.comments.list":

type CommentsListCall struct {
	s          *Service
	activityId string
	opt_       map[string]interface{}
}

// List: Retrieves a list of comments, possibly filtered.
func (r *CommentsService) List(activityId string) *CommentsListCall {
	c := &CommentsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.activityId = activityId
	return c
}

// OrderBy sets the optional parameter "orderBy": Sort search results.
func (c *CommentsListCall) OrderBy(orderBy string) *CommentsListCall {
	c.opt_["orderBy"] = orderBy
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of activities to include in the response.
func (c *CommentsListCall) MaxResults(maxResults int64) *CommentsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// Hl sets the optional parameter "hl": Specifies the interface language
// (host language) of your user interface.
func (c *CommentsListCall) Hl(hl string) *CommentsListCall {
	c.opt_["hl"] = hl
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token that allows pagination.
func (c *CommentsListCall) PageToken(pageToken string) *CommentsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

func (c *CommentsListCall) Do() (*CommentList, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["orderBy"]; ok {
		params.Set("orderBy", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/orkut/v2/", "activities/{activityId}/comments")
	urls = strings.Replace(urls, "{activityId}", cleanPathString(c.activityId), 1)
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
	ret := new(CommentList)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of comments, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "orkut.comments.list",
	//   "parameterOrder": [
	//     "activityId"
	//   ],
	//   "parameters": {
	//     "activityId": {
	//       "description": "The ID of the activity containing the comments.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Specifies the interface language (host language) of your user interface.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of activities to include in the response.",
	//       "format": "uint32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "orderBy": {
	//       "default": "DESCENDING_SORT",
	//       "description": "Sort search results.",
	//       "enum": [
	//         "ascending",
	//         "descending"
	//       ],
	//       "enumDescriptions": [
	//         "Use ascending sort order.",
	//         "Use descending sort order."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token that allows pagination.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities/{activityId}/comments",
	//   "response": {
	//     "$ref": "CommentList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/orkut",
	//     "https://www.googleapis.com/auth/orkut.readonly"
	//   ]
	// }

}

// method id "orkut.comments.insert":

type CommentsInsertCall struct {
	s          *Service
	activityId string
	comment    *Comment
	opt_       map[string]interface{}
}

// Insert: Inserts a new comment to an activity.
func (r *CommentsService) Insert(activityId string, comment *Comment) *CommentsInsertCall {
	c := &CommentsInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.activityId = activityId
	c.comment = comment
	return c
}

func (c *CommentsInsertCall) Do() (*Comment, os.Error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.comment)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/orkut/v2/", "activities/{activityId}/comments")
	urls = strings.Replace(urls, "{activityId}", cleanPathString(c.activityId), 1)
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
	ret := new(Comment)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Inserts a new comment to an activity.",
	//   "httpMethod": "POST",
	//   "id": "orkut.comments.insert",
	//   "parameterOrder": [
	//     "activityId"
	//   ],
	//   "parameters": {
	//     "activityId": {
	//       "description": "The ID of the activity to contain the new comment.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities/{activityId}/comments",
	//   "request": {
	//     "$ref": "Comment"
	//   },
	//   "response": {
	//     "$ref": "Comment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/orkut"
	//   ]
	// }

}

// method id "orkut.comments.get":

type CommentsGetCall struct {
	s         *Service
	commentId string
	opt_      map[string]interface{}
}

// Get: Retrieves an existing comment.
func (r *CommentsService) Get(commentId string) *CommentsGetCall {
	c := &CommentsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.commentId = commentId
	return c
}

// Hl sets the optional parameter "hl": Specifies the interface language
// (host language) of your user interface.
func (c *CommentsGetCall) Hl(hl string) *CommentsGetCall {
	c.opt_["hl"] = hl
	return c
}

func (c *CommentsGetCall) Do() (*Comment, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/orkut/v2/", "comments/{commentId}")
	urls = strings.Replace(urls, "{commentId}", cleanPathString(c.commentId), 1)
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
	ret := new(Comment)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves an existing comment.",
	//   "httpMethod": "GET",
	//   "id": "orkut.comments.get",
	//   "parameterOrder": [
	//     "commentId"
	//   ],
	//   "parameters": {
	//     "commentId": {
	//       "description": "ID of the comment to get.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Specifies the interface language (host language) of your user interface.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "comments/{commentId}",
	//   "response": {
	//     "$ref": "Comment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/orkut",
	//     "https://www.googleapis.com/auth/orkut.readonly"
	//   ]
	// }

}

// method id "orkut.comments.delete":

type CommentsDeleteCall struct {
	s         *Service
	commentId string
	opt_      map[string]interface{}
}

// Delete: Deletes an existing comment.
func (r *CommentsService) Delete(commentId string) *CommentsDeleteCall {
	c := &CommentsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.commentId = commentId
	return c
}

func (c *CommentsDeleteCall) Do() os.Error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/orkut/v2/", "comments/{commentId}")
	urls = strings.Replace(urls, "{commentId}", cleanPathString(c.commentId), 1)
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
	//   "description": "Deletes an existing comment.",
	//   "httpMethod": "DELETE",
	//   "id": "orkut.comments.delete",
	//   "parameterOrder": [
	//     "commentId"
	//   ],
	//   "parameters": {
	//     "commentId": {
	//       "description": "ID of the comment to remove.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "comments/{commentId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/orkut"
	//   ]
	// }

}

// method id "orkut.counters.list":

type CountersListCall struct {
	s      *Service
	userId string
	opt_   map[string]interface{}
}

// List: Retrieves the counters of an user.
func (r *CountersService) List(userId string) *CountersListCall {
	c := &CountersListCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	return c
}

func (c *CountersListCall) Do() (*Counters, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/orkut/v2/", "people/{userId}/counters")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
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
	ret := new(Counters)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the counters of an user.",
	//   "httpMethod": "GET",
	//   "id": "orkut.counters.list",
	//   "parameterOrder": [
	//     "userId"
	//   ],
	//   "parameters": {
	//     "userId": {
	//       "description": "The ID of the user whose counters will be listed. Can be me to refer to caller.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "people/{userId}/counters",
	//   "response": {
	//     "$ref": "Counters"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/orkut",
	//     "https://www.googleapis.com/auth/orkut.readonly"
	//   ]
	// }

}

// method id "orkut.activities.list":

type ActivitiesListCall struct {
	s          *Service
	userId     string
	collection string
	opt_       map[string]interface{}
}

// List: Retrieves a list of activities.
func (r *ActivitiesService) List(userId string, collection string) *ActivitiesListCall {
	c := &ActivitiesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.collection = collection
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of activities to include in the response.
func (c *ActivitiesListCall) MaxResults(maxResults int64) *ActivitiesListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// Hl sets the optional parameter "hl": Specifies the interface language
// (host language) of your user interface.
func (c *ActivitiesListCall) Hl(hl string) *ActivitiesListCall {
	c.opt_["hl"] = hl
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token that allows pagination.
func (c *ActivitiesListCall) PageToken(pageToken string) *ActivitiesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

func (c *ActivitiesListCall) Do() (*ActivityList, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/orkut/v2/", "people/{userId}/activities/{collection}")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{collection}", cleanPathString(c.collection), 1)
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
	ret := new(ActivityList)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of activities.",
	//   "httpMethod": "GET",
	//   "id": "orkut.activities.list",
	//   "parameterOrder": [
	//     "userId",
	//     "collection"
	//   ],
	//   "parameters": {
	//     "collection": {
	//       "description": "The collection of activities to list.",
	//       "enum": [
	//         "all",
	//         "scraps",
	//         "stream"
	//       ],
	//       "enumDescriptions": [
	//         "All activities created by the specified user that the authenticated user is authorized to view.",
	//         "The specified user's scrapbook.",
	//         "The specified user's stream feed, intended for consumption. This includes activities posted by people that the user is following, and activities in which the user has been mentioned."
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Specifies the interface language (host language) of your user interface.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of activities to include in the response.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "100",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token that allows pagination.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user whose activities will be listed. Can be me to refer to the viewer (i.e. the authenticated user).",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "people/{userId}/activities/{collection}",
	//   "response": {
	//     "$ref": "ActivityList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/orkut",
	//     "https://www.googleapis.com/auth/orkut.readonly"
	//   ]
	// }

}

// method id "orkut.activities.delete":

type ActivitiesDeleteCall struct {
	s          *Service
	activityId string
	opt_       map[string]interface{}
}

// Delete: Deletes an existing activity, if the access controls allow
// it.
func (r *ActivitiesService) Delete(activityId string) *ActivitiesDeleteCall {
	c := &ActivitiesDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.activityId = activityId
	return c
}

func (c *ActivitiesDeleteCall) Do() os.Error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/orkut/v2/", "activities/{activityId}")
	urls = strings.Replace(urls, "{activityId}", cleanPathString(c.activityId), 1)
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
	//   "description": "Deletes an existing activity, if the access controls allow it.",
	//   "httpMethod": "DELETE",
	//   "id": "orkut.activities.delete",
	//   "parameterOrder": [
	//     "activityId"
	//   ],
	//   "parameters": {
	//     "activityId": {
	//       "description": "ID of the activity to remove.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities/{activityId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/orkut"
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
