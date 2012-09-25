// Package plus provides access to the Google+ API.
//
// See https://developers.google.com/+/api/
//
// Usage example:
//
//   import "code.google.com/p/google-api-go-client/plus/v1"
//   ...
//   plusService, err := plus.New(oauthHttpClient)
package plus

import (
	"bytes"
	"code.google.com/p/google-api-go-client/googleapi"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New

const apiId = "plus:v1"
const apiName = "plus"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/plus/v1/"

// OAuth2 scopes used by this API.
const (
	// Know who you are on Google
	PlusMeScope = "https://www.googleapis.com/auth/plus.me"

	// View your email address
	UserinfoEmailScope = "https://www.googleapis.com/auth/userinfo.email"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Activities = &ActivitiesService{s: s}
	s.Comments = &CommentsService{s: s}
	s.People = &PeopleService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	Activities *ActivitiesService

	Comments *CommentsService

	People *PeopleService
}

type ActivitiesService struct {
	s *Service
}

type CommentsService struct {
	s *Service
}

type PeopleService struct {
	s *Service
}

type Acl struct {
	// Description: Description of the access granted, suitable for display.
	Description string `json:"description,omitempty"`

	// Items: The list of access entries.
	Items []*PlusAclentryResource `json:"items,omitempty"`

	// Kind: Identifies this resource as a collection of access controls.
	// Value: "plus#acl".
	Kind string `json:"kind,omitempty"`
}

type Activity struct {
	// Access: Identifies who has access to see this activity.
	Access *Acl `json:"access,omitempty"`

	// Actor: The person who performed this activity.
	Actor *ActivityActor `json:"actor,omitempty"`

	// Address: Street address where this activity occurred.
	Address string `json:"address,omitempty"`

	// Annotation: Additional content added by the person who shared this
	// activity, applicable only when resharing an activity.
	Annotation string `json:"annotation,omitempty"`

	// CrosspostSource: If this activity is a crosspost from another system,
	// this property specifies the ID of the original activity.
	CrosspostSource string `json:"crosspostSource,omitempty"`

	// Etag: ETag of this response for caching purposes.
	Etag string `json:"etag,omitempty"`

	// Geocode: Latitude and longitude where this activity occurred. Format
	// is latitude followed by longitude, space separated.
	Geocode string `json:"geocode,omitempty"`

	// Id: The ID of this activity.
	Id string `json:"id,omitempty"`

	// Kind: Identifies this resource as an activity. Value:
	// "plus#activity".
	Kind string `json:"kind,omitempty"`

	// Object: The object of this activity.
	Object *ActivityObject `json:"object,omitempty"`

	// PlaceId: ID of the place where this activity occurred.
	PlaceId string `json:"placeId,omitempty"`

	// PlaceName: Name of the place where this activity occurred.
	PlaceName string `json:"placeName,omitempty"`

	// Provider: The service provider that initially published this
	// activity.
	Provider *ActivityProvider `json:"provider,omitempty"`

	// Published: The time at which this activity was initially published.
	// Formatted as an RFC 3339 timestamp.
	Published string `json:"published,omitempty"`

	// Radius: Radius, in meters, of the region where this activity
	// occurred, centered at the latitude and longitude identified in
	// geocode.
	Radius string `json:"radius,omitempty"`

	// Title: Title of this activity.
	Title string `json:"title,omitempty"`

	// Updated: The time at which this activity was last updated. Formatted
	// as an RFC 3339 timestamp.
	Updated string `json:"updated,omitempty"`

	// Url: The link to this activity.
	Url string `json:"url,omitempty"`

	// Verb: This activity's verb, indicating what action was performed.
	// Possible values are:  
	// - "post" - Publish content to the stream. 
	// -
	// "checkin" - Check in to a location. 
	// - "share" - Reshare an activity.
	Verb string `json:"verb,omitempty"`
}

type ActivityActor struct {
	// DisplayName: The name of the actor, suitable for display.
	DisplayName string `json:"displayName,omitempty"`

	// Id: The ID of the actor's person resource.
	Id string `json:"id,omitempty"`

	// Image: The image representation of the actor.
	Image *ActivityActorImage `json:"image,omitempty"`

	// Name: An object representation of the individual components of name.
	Name *ActivityActorName `json:"name,omitempty"`

	// Url: The link to the actor's Google profile.
	Url string `json:"url,omitempty"`
}

type ActivityActorImage struct {
	// Url: The URL of the actor's profile photo. To re-size the image and
	// crop it to a square, append the query string ?sz=x, where x is the
	// dimension in pixels of each side.
	Url string `json:"url,omitempty"`
}

type ActivityActorName struct {
	// FamilyName: The family name (last name) of the actor.
	FamilyName string `json:"familyName,omitempty"`

	// GivenName: The given name (first name) of the actor.
	GivenName string `json:"givenName,omitempty"`
}

type ActivityObject struct {
	// Actor: If this activity's object is itself another activity (for
	// example, when a person reshares an activity), this property specifies
	// the original activity's actor.
	Actor *ActivityObjectActor `json:"actor,omitempty"`

	// Attachments: The media objects attached to this activity.
	Attachments []*ActivityObjectAttachments `json:"attachments,omitempty"`

	// Content: The HTML-formatted content, suitable for display. When
	// creating or updating an activity, this value must be supplied as
	// plain text in the request. If successful, the response will contain
	// the HTML-formatted content. When updating an activity, use
	// originalContent as the starting value, then assign the updated text
	// to this property.
	Content string `json:"content,omitempty"`

	// Id: The ID of the object. When resharing an activity, this is the ID
	// of the activity being reshared.
	Id string `json:"id,omitempty"`

	// ObjectType: The type of the object. Possible values are:  
	// - "note" -
	// Textual content. 
	// - "activity" - A Google+ activity.
	ObjectType string `json:"objectType,omitempty"`

	// OriginalContent: The content (text) as provided by the author, stored
	// without any HTML formatting. When updating an activity's content, use
	// the value of originalContent as the starting point from which to make
	// edits.
	OriginalContent string `json:"originalContent,omitempty"`

	// Plusoners: People who +1'd this activity.
	Plusoners *ActivityObjectPlusoners `json:"plusoners,omitempty"`

	// Replies: Comments in reply to this activity.
	Replies *ActivityObjectReplies `json:"replies,omitempty"`

	// Resharers: People who reshared this activity.
	Resharers *ActivityObjectResharers `json:"resharers,omitempty"`

	// Url: The URL that points to the linked resource.
	Url string `json:"url,omitempty"`
}

type ActivityObjectActor struct {
	// DisplayName: The original actor's name, suitable for display.
	DisplayName string `json:"displayName,omitempty"`

	// Id: ID of the original actor.
	Id string `json:"id,omitempty"`

	// Image: The image representation of the original actor.
	Image *ActivityObjectActorImage `json:"image,omitempty"`

	// Url: A link to the original actor's Google profile.
	Url string `json:"url,omitempty"`
}

type ActivityObjectActorImage struct {
	// Url: A URL that points to a thumbnail photo of the original actor.
	Url string `json:"url,omitempty"`
}

type ActivityObjectAttachments struct {
	// Content: If the attachment is an article, this property contains a
	// snippet of text from the article.
	Content string `json:"content,omitempty"`

	// DisplayName: The title of the attachment (such as a photo caption or
	// an article title).
	DisplayName string `json:"displayName,omitempty"`

	// Embed: If the attachment is a video, the embeddable link.
	Embed *ActivityObjectAttachmentsEmbed `json:"embed,omitempty"`

	// FullImage: The full image url for photo attachments.
	FullImage *ActivityObjectAttachmentsFullImage `json:"fullImage,omitempty"`

	// Id: The ID of the media object's resource.
	Id string `json:"id,omitempty"`

	// Image: The preview image for photos or videos.
	Image *ActivityObjectAttachmentsImage `json:"image,omitempty"`

	// ObjectType: The type of media object. Possible values are:  
	// -
	// "photo" - A photo. 
	// - "video" - A video. 
	// - "article" - An article,
	// specified by a link.
	ObjectType string `json:"objectType,omitempty"`

	// Url: The link to the attachment, should be of type text/html.
	Url string `json:"url,omitempty"`
}

type ActivityObjectAttachmentsEmbed struct {
	// Type: Media type of the link.
	Type string `json:"type,omitempty"`

	// Url: URL of the link.
	Url string `json:"url,omitempty"`
}

type ActivityObjectAttachmentsFullImage struct {
	// Height: The height, in pixels, of the linked resource.
	Height int64 `json:"height,omitempty"`

	// Type: Media type of the link.
	Type string `json:"type,omitempty"`

	// Url: URL of the link.
	Url string `json:"url,omitempty"`

	// Width: The width, in pixels, of the linked resource.
	Width int64 `json:"width,omitempty"`
}

type ActivityObjectAttachmentsImage struct {
	// Height: The height, in pixels, of the linked resource.
	Height int64 `json:"height,omitempty"`

	// Type: Media type of the link.
	Type string `json:"type,omitempty"`

	// Url: URL of the link.
	Url string `json:"url,omitempty"`

	// Width: The width, in pixels, of the linked resource.
	Width int64 `json:"width,omitempty"`
}

type ActivityObjectPlusoners struct {
	// SelfLink: The URL for the collection of people who +1'd this
	// activity.
	SelfLink string `json:"selfLink,omitempty"`

	// TotalItems: Total number of people who +1'd this activity.
	TotalItems int64 `json:"totalItems,omitempty"`
}

type ActivityObjectReplies struct {
	// SelfLink: The URL for the collection of comments in reply to this
	// activity.
	SelfLink string `json:"selfLink,omitempty"`

	// TotalItems: Total number of comments on this activity.
	TotalItems int64 `json:"totalItems,omitempty"`
}

type ActivityObjectResharers struct {
	// SelfLink: The URL for the collection of resharers.
	SelfLink string `json:"selfLink,omitempty"`

	// TotalItems: Total number of people who reshared this activity.
	TotalItems int64 `json:"totalItems,omitempty"`
}

type ActivityProvider struct {
	// Title: Name of the service provider.
	Title string `json:"title,omitempty"`
}

type ActivityFeed struct {
	// Etag: ETag of this response for caching purposes.
	Etag string `json:"etag,omitempty"`

	// Id: The ID of this collection of activities.
	Id string `json:"id,omitempty"`

	// Items: The activities in this page of results.
	Items []*Activity `json:"items,omitempty"`

	// Kind: Identifies this resource as a collection of activities. Value:
	// "plus#activityFeed".
	Kind string `json:"kind,omitempty"`

	// NextLink: Link to the next page of activities.
	NextLink string `json:"nextLink,omitempty"`

	// NextPageToken: The continuation token, used to page through large
	// result sets. Provide this value in a subsequent request to return the
	// next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: Link to this activity resource.
	SelfLink string `json:"selfLink,omitempty"`

	// Title: The title of this collection of activities.
	Title string `json:"title,omitempty"`

	// Updated: The time at which this collection of activities was last
	// updated. Formatted as an RFC 3339 timestamp.
	Updated string `json:"updated,omitempty"`
}

type Comment struct {
	// Actor: The person who posted this comment.
	Actor *CommentActor `json:"actor,omitempty"`

	// Etag: ETag of this response for caching purposes.
	Etag string `json:"etag,omitempty"`

	// Id: The ID of this comment.
	Id string `json:"id,omitempty"`

	// InReplyTo: The activity this comment replied to.
	InReplyTo []*CommentInReplyTo `json:"inReplyTo,omitempty"`

	// Kind: Identifies this resource as a comment. Value: "plus#comment".
	Kind string `json:"kind,omitempty"`

	// Object: The object of this comment.
	Object *CommentObject `json:"object,omitempty"`

	// Published: The time at which this comment was initially published.
	// Formatted as an RFC 3339 timestamp.
	Published string `json:"published,omitempty"`

	// SelfLink: Link to this comment resource.
	SelfLink string `json:"selfLink,omitempty"`

	// Updated: The time at which this comment was last updated. Formatted
	// as an RFC 3339 timestamp.
	Updated string `json:"updated,omitempty"`

	// Verb: This comment's verb, indicating what action was performed.
	// Possible values are:  
	// - "post" - Publish content to the stream.
	Verb string `json:"verb,omitempty"`
}

type CommentActor struct {
	// DisplayName: The name of this actor, suitable for display.
	DisplayName string `json:"displayName,omitempty"`

	// Id: The ID of the actor.
	Id string `json:"id,omitempty"`

	// Image: The image representation of this actor.
	Image *CommentActorImage `json:"image,omitempty"`

	// Url: A link to the person resource for this actor.
	Url string `json:"url,omitempty"`
}

type CommentActorImage struct {
	// Url: The URL of the actor's profile photo. To re-size the image and
	// crop it to a square, append the query string ?sz=x, where x is the
	// dimension in pixels of each side.
	Url string `json:"url,omitempty"`
}

type CommentInReplyTo struct {
	// Id: The id of the activity.
	Id string `json:"id,omitempty"`

	// Url: The url of the activity.
	Url string `json:"url,omitempty"`
}

type CommentObject struct {
	// Content: The content of this comment.
	Content string `json:"content,omitempty"`

	// ObjectType: The object type of this comment. Possible values are:  
	// -
	// "comment" - A comment in reply to an activity.
	ObjectType string `json:"objectType,omitempty"`
}

type CommentFeed struct {
	// Etag: ETag of this response for caching purposes.
	Etag string `json:"etag,omitempty"`

	// Id: The ID of this collection of comments.
	Id string `json:"id,omitempty"`

	// Items: The comments in this page of results.
	Items []*Comment `json:"items,omitempty"`

	// Kind: Identifies this resource as a collection of comments. Value:
	// "plus#commentFeed".
	Kind string `json:"kind,omitempty"`

	// NextLink: Link to the next page of activities.
	NextLink string `json:"nextLink,omitempty"`

	// NextPageToken: The continuation token, used to page through large
	// result sets. Provide this value in a subsequent request to return the
	// next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Title: The title of this collection of comments.
	Title string `json:"title,omitempty"`

	// Updated: The time at which this collection of comments was last
	// updated. Formatted as an RFC 3339 timestamp.
	Updated string `json:"updated,omitempty"`
}

type PeopleFeed struct {
	// Etag: ETag of this response for caching purposes.
	Etag string `json:"etag,omitempty"`

	// Items: The people in this page of results. Each item will include the
	// id, displayName, image, and url for the person. To retrieve
	// additional profile data, see the people.get method.
	Items []*Person `json:"items,omitempty"`

	// Kind: Identifies this resource as a collection of people. Value:
	// "plus#peopleFeed".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The continuation token, used to page through large
	// result sets. Provide this value in a subsequent request to return the
	// next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: Link to this resource.
	SelfLink string `json:"selfLink,omitempty"`

	// Title: The title of this collection of people.
	Title string `json:"title,omitempty"`
}

type Person struct {
	// AboutMe: A short biography for this person.
	AboutMe string `json:"aboutMe,omitempty"`

	// Birthday: The person's date of birth, represented as YYYY-MM-DD.
	Birthday string `json:"birthday,omitempty"`

	// CurrentLocation: The current location for this person.
	CurrentLocation string `json:"currentLocation,omitempty"`

	// DisplayName: The name of this person, suitable for display.
	DisplayName string `json:"displayName,omitempty"`

	// Emails: A list of email addresses for this person.
	Emails []*PersonEmails `json:"emails,omitempty"`

	// Etag: ETag of this response for caching purposes.
	Etag string `json:"etag,omitempty"`

	// Gender: The person's gender. Possible values are:  
	// - "male" - Male
	// gender. 
	// - "female" - Female gender. 
	// - "other" - Other.
	Gender string `json:"gender,omitempty"`

	// HasApp: If "true", indicates that the person has installed the app
	// that is making the request and has chosen to expose this install
	// state to the caller. A value of "false" indicates that the install
	// state cannot be determined (it is either not installed or the person
	// has chosen to keep this information private).
	HasApp bool `json:"hasApp,omitempty"`

	// Id: The ID of this person.
	Id string `json:"id,omitempty"`

	// Image: The representation of the person's profile photo.
	Image *PersonImage `json:"image,omitempty"`

	// Kind: Identifies this resource as a person. Value: "plus#person".
	Kind string `json:"kind,omitempty"`

	// LanguagesSpoken: The languages spoken by this person.
	LanguagesSpoken []string `json:"languagesSpoken,omitempty"`

	// Name: An object representation of the individual components of a
	// person's name.
	Name *PersonName `json:"name,omitempty"`

	// Nickname: The nickname of this person.
	Nickname string `json:"nickname,omitempty"`

	// ObjectType: Type of person within Google+. Possible values are:  
	// -
	// "person" - represents an actual person. 
	// - "page" - represents a
	// page.
	ObjectType string `json:"objectType,omitempty"`

	// Organizations: A list of current or past organizations with which
	// this person is associated.
	Organizations []*PersonOrganizations `json:"organizations,omitempty"`

	// PlacesLived: A list of places where this person has lived.
	PlacesLived []*PersonPlacesLived `json:"placesLived,omitempty"`

	// RelationshipStatus: The person's relationship status. Possible values
	// are:  
	// - "single" - Person is single. 
	// - "in_a_relationship" - Person
	// is in a relationship. 
	// - "engaged" - Person is engaged. 
	// - "married"
	// - Person is married. 
	// - "its_complicated" - The relationship is
	// complicated. 
	// - "open_relationship" - Person is in an open
	// relationship. 
	// - "widowed" - Person is widowed. 
	// -
	// "in_domestic_partnership" - Person is in a domestic partnership. 
	// -
	// "in_civil_union" - Person is in a civil union.
	RelationshipStatus string `json:"relationshipStatus,omitempty"`

	// Tagline: The brief description (tagline) of this person.
	Tagline string `json:"tagline,omitempty"`

	// Url: The URL of this person's profile.
	Url string `json:"url,omitempty"`

	// Urls: A list of URLs for this person.
	Urls []*PersonUrls `json:"urls,omitempty"`
}

type PersonEmails struct {
	// Primary: If "true", indicates this email address is the person's
	// primary one.
	Primary bool `json:"primary,omitempty"`

	// Type: The type of address. Possible values are:  
	// - "home" - Home
	// email address. 
	// - "work" - Work email address. 
	// - "other" - Other.
	Type string `json:"type,omitempty"`

	// Value: The email address.
	Value string `json:"value,omitempty"`
}

type PersonImage struct {
	// Url: The URL of the person's profile photo. To re-size the image and
	// crop it to a square, append the query string ?sz=x, where x is the
	// dimension in pixels of each side.
	Url string `json:"url,omitempty"`
}

type PersonName struct {
	// FamilyName: The family name (last name) of this person.
	FamilyName string `json:"familyName,omitempty"`

	// Formatted: The full name of this person, including middle names,
	// suffixes, etc.
	Formatted string `json:"formatted,omitempty"`

	// GivenName: The given name (first name) of this person.
	GivenName string `json:"givenName,omitempty"`

	// HonorificPrefix: The honorific prefixes (such as "Dr." or "Mrs.") for
	// this person.
	HonorificPrefix string `json:"honorificPrefix,omitempty"`

	// HonorificSuffix: The honorific suffixes (such as "Jr.") for this
	// person.
	HonorificSuffix string `json:"honorificSuffix,omitempty"`

	// MiddleName: The middle name of this person.
	MiddleName string `json:"middleName,omitempty"`
}

type PersonOrganizations struct {
	// Department: The department within the organization. Deprecated.
	Department string `json:"department,omitempty"`

	// Description: A short description of the person's role in this
	// organization. Deprecated.
	Description string `json:"description,omitempty"`

	// EndDate: The date the person left this organization.
	EndDate string `json:"endDate,omitempty"`

	// Location: The location of this organization. Deprecated.
	Location string `json:"location,omitempty"`

	// Name: The name of the organization.
	Name string `json:"name,omitempty"`

	// Primary: If "true", indicates this organization is the person's
	// primary one (typically interpreted as current one).
	Primary bool `json:"primary,omitempty"`

	// StartDate: The date the person joined this organization.
	StartDate string `json:"startDate,omitempty"`

	// Title: The person's job title or role within the organization.
	Title string `json:"title,omitempty"`

	// Type: The type of organization. Possible values are:  
	// - "work" -
	// Work. 
	// - "school" - School.
	Type string `json:"type,omitempty"`
}

type PersonPlacesLived struct {
	// Primary: If "true", this place of residence is this person's primary
	// residence.
	Primary bool `json:"primary,omitempty"`

	// Value: A place where this person has lived. For example: "Seattle,
	// WA", "Near Toronto".
	Value string `json:"value,omitempty"`
}

type PersonUrls struct {
	// Primary: If "true", this URL is the person's primary URL.
	Primary bool `json:"primary,omitempty"`

	// Type: The type of URL. Possible values are:  
	// - "home" - URL for
	// home. 
	// - "work" - URL for work. 
	// - "blog" - URL for blog. 
	// -
	// "profile" - URL for profile. 
	// - "other" - Other.
	Type string `json:"type,omitempty"`

	// Value: The URL value.
	Value string `json:"value,omitempty"`
}

type PlusAclentryResource struct {
	// Id: The ID of the entry. For entries of type "person" or "circle",
	// this is the ID of the resource. For other types, this property is not
	// set.
	Id string `json:"id,omitempty"`

	// Type: The type of entry describing to whom access is granted.
	// Possible values are:  
	// - "person" - Access to an individual. 
	// -
	// "circle" - Access to members of a circle. 
	// - "myCircles" - Access to
	// members of all the person's circles. 
	// - "extendedCircles" - Access to
	// members of everyone in a person's circles, plus all of the people in
	// their circles. 
	// - "public" - Access to anyone on the web.
	Type string `json:"type,omitempty"`
}

// method id "plus.activities.get":

type ActivitiesGetCall struct {
	s          *Service
	activityId string
	opt_       map[string]interface{}
}

// Get: Get an activity.
func (r *ActivitiesService) Get(activityId string) *ActivitiesGetCall {
	c := &ActivitiesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.activityId = activityId
	return c
}

func (c *ActivitiesGetCall) Do() (*Activity, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/plus/v1/", "activities/{activityId}")
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
	ret := new(Activity)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get an activity.",
	//   "httpMethod": "GET",
	//   "id": "plus.activities.get",
	//   "parameterOrder": [
	//     "activityId"
	//   ],
	//   "parameters": {
	//     "activityId": {
	//       "description": "The ID of the activity to get.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities/{activityId}",
	//   "response": {
	//     "$ref": "Activity"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/plus.me"
	//   ]
	// }

}

// method id "plus.activities.list":

type ActivitiesListCall struct {
	s          *Service
	userId     string
	collection string
	opt_       map[string]interface{}
}

// List: List all of the activities in the specified collection for a
// particular user.
func (r *ActivitiesService) List(userId string, collection string) *ActivitiesListCall {
	c := &ActivitiesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.collection = collection
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of activities to include in the response, used for paging. For
// any response, the actual number returned may be less than the
// specified maxResults.
func (c *ActivitiesListCall) MaxResults(maxResults int64) *ActivitiesListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, used to page through large result sets. To get the next page
// of results, set this parameter to the value of "nextPageToken" from
// the previous response.
func (c *ActivitiesListCall) PageToken(pageToken string) *ActivitiesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

func (c *ActivitiesListCall) Do() (*ActivityFeed, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/plus/v1/", "people/{userId}/activities/{collection}")
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
	ret := new(ActivityFeed)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all of the activities in the specified collection for a particular user.",
	//   "httpMethod": "GET",
	//   "id": "plus.activities.list",
	//   "parameterOrder": [
	//     "userId",
	//     "collection"
	//   ],
	//   "parameters": {
	//     "collection": {
	//       "description": "The collection of activities to list.",
	//       "enum": [
	//         "public"
	//       ],
	//       "enumDescriptions": [
	//         "All public activities created by the specified user."
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "20",
	//       "description": "The maximum number of activities to include in the response, used for paging. For any response, the actual number returned may be less than the specified maxResults.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "100",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, used to page through large result sets. To get the next page of results, set this parameter to the value of \"nextPageToken\" from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user to get activities for. The special value \"me\" can be used to indicate the authenticated user.",
	//       "location": "path",
	//       "pattern": "me|[0-9]+",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "people/{userId}/activities/{collection}",
	//   "response": {
	//     "$ref": "ActivityFeed"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/plus.me"
	//   ]
	// }

}

// method id "plus.activities.search":

type ActivitiesSearchCall struct {
	s     *Service
	query string
	opt_  map[string]interface{}
}

// Search: Search public activities.
func (r *ActivitiesService) Search(query string) *ActivitiesSearchCall {
	c := &ActivitiesSearchCall{s: r.s, opt_: make(map[string]interface{})}
	c.query = query
	return c
}

// Language sets the optional parameter "language": Specify the
// preferred language to search with. See search language codes for
// available values.
func (c *ActivitiesSearchCall) Language(language string) *ActivitiesSearchCall {
	c.opt_["language"] = language
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of activities to include in the response, used for paging. For
// any response, the actual number returned may be less than the
// specified maxResults.
func (c *ActivitiesSearchCall) MaxResults(maxResults int64) *ActivitiesSearchCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// OrderBy sets the optional parameter "orderBy": Specifies how to order
// search results.
func (c *ActivitiesSearchCall) OrderBy(orderBy string) *ActivitiesSearchCall {
	c.opt_["orderBy"] = orderBy
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, used to page through large result sets. To get the next page
// of results, set this parameter to the value of "nextPageToken" from
// the previous response. This token may be of any length.
func (c *ActivitiesSearchCall) PageToken(pageToken string) *ActivitiesSearchCall {
	c.opt_["pageToken"] = pageToken
	return c
}

func (c *ActivitiesSearchCall) Do() (*ActivityFeed, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("query", fmt.Sprintf("%v", c.query))
	if v, ok := c.opt_["language"]; ok {
		params.Set("language", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["orderBy"]; ok {
		params.Set("orderBy", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/plus/v1/", "activities")
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
	ret := new(ActivityFeed)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Search public activities.",
	//   "httpMethod": "GET",
	//   "id": "plus.activities.search",
	//   "parameterOrder": [
	//     "query"
	//   ],
	//   "parameters": {
	//     "language": {
	//       "default": "",
	//       "description": "Specify the preferred language to search with. See search language codes for available values.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "10",
	//       "description": "The maximum number of activities to include in the response, used for paging. For any response, the actual number returned may be less than the specified maxResults.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "20",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "orderBy": {
	//       "default": "recent",
	//       "description": "Specifies how to order search results.",
	//       "enum": [
	//         "best",
	//         "recent"
	//       ],
	//       "enumDescriptions": [
	//         "Sort activities by relevance to the user, most relevant first.",
	//         "Sort activities by published date, most recent first."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, used to page through large result sets. To get the next page of results, set this parameter to the value of \"nextPageToken\" from the previous response. This token may be of any length.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "query": {
	//       "description": "Full-text search query string.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities",
	//   "response": {
	//     "$ref": "ActivityFeed"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/plus.me"
	//   ]
	// }

}

// method id "plus.comments.get":

type CommentsGetCall struct {
	s         *Service
	commentId string
	opt_      map[string]interface{}
}

// Get: Get a comment.
func (r *CommentsService) Get(commentId string) *CommentsGetCall {
	c := &CommentsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.commentId = commentId
	return c
}

func (c *CommentsGetCall) Do() (*Comment, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/plus/v1/", "comments/{commentId}")
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
	//   "description": "Get a comment.",
	//   "httpMethod": "GET",
	//   "id": "plus.comments.get",
	//   "parameterOrder": [
	//     "commentId"
	//   ],
	//   "parameters": {
	//     "commentId": {
	//       "description": "The ID of the comment to get.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "comments/{commentId}",
	//   "response": {
	//     "$ref": "Comment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/plus.me"
	//   ]
	// }

}

// method id "plus.comments.list":

type CommentsListCall struct {
	s          *Service
	activityId string
	opt_       map[string]interface{}
}

// List: List all of the comments for an activity.
func (r *CommentsService) List(activityId string) *CommentsListCall {
	c := &CommentsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.activityId = activityId
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of comments to include in the response, used for paging. For
// any response, the actual number returned may be less than the
// specified maxResults.
func (c *CommentsListCall) MaxResults(maxResults int64) *CommentsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, used to page through large result sets. To get the next page
// of results, set this parameter to the value of "nextPageToken" from
// the previous response.
func (c *CommentsListCall) PageToken(pageToken string) *CommentsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// SortOrder sets the optional parameter "sortOrder": The order in which
// to sort the list of comments.
func (c *CommentsListCall) SortOrder(sortOrder string) *CommentsListCall {
	c.opt_["sortOrder"] = sortOrder
	return c
}

func (c *CommentsListCall) Do() (*CommentFeed, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["sortOrder"]; ok {
		params.Set("sortOrder", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/plus/v1/", "activities/{activityId}/comments")
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
	ret := new(CommentFeed)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all of the comments for an activity.",
	//   "httpMethod": "GET",
	//   "id": "plus.comments.list",
	//   "parameterOrder": [
	//     "activityId"
	//   ],
	//   "parameters": {
	//     "activityId": {
	//       "description": "The ID of the activity to get comments for.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "20",
	//       "description": "The maximum number of comments to include in the response, used for paging. For any response, the actual number returned may be less than the specified maxResults.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "100",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, used to page through large result sets. To get the next page of results, set this parameter to the value of \"nextPageToken\" from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "default": "ascending",
	//       "description": "The order in which to sort the list of comments.",
	//       "enum": [
	//         "ascending",
	//         "descending"
	//       ],
	//       "enumDescriptions": [
	//         "Sort oldest comments first.",
	//         "Sort newest comments first."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities/{activityId}/comments",
	//   "response": {
	//     "$ref": "CommentFeed"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/plus.me"
	//   ]
	// }

}

// method id "plus.people.get":

type PeopleGetCall struct {
	s      *Service
	userId string
	opt_   map[string]interface{}
}

// Get: Get a person's profile.
func (r *PeopleService) Get(userId string) *PeopleGetCall {
	c := &PeopleGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	return c
}

func (c *PeopleGetCall) Do() (*Person, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/plus/v1/", "people/{userId}")
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
	ret := new(Person)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get a person's profile.",
	//   "httpMethod": "GET",
	//   "id": "plus.people.get",
	//   "parameterOrder": [
	//     "userId"
	//   ],
	//   "parameters": {
	//     "userId": {
	//       "description": "The ID of the person to get the profile for. The special value \"me\" can be used to indicate the authenticated user.",
	//       "location": "path",
	//       "pattern": "me|[0-9]+",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "people/{userId}",
	//   "response": {
	//     "$ref": "Person"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/plus.me",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "plus.people.listByActivity":

type PeopleListByActivityCall struct {
	s          *Service
	activityId string
	collection string
	opt_       map[string]interface{}
}

// ListByActivity: List all of the people in the specified collection
// for a particular activity.
func (r *PeopleService) ListByActivity(activityId string, collection string) *PeopleListByActivityCall {
	c := &PeopleListByActivityCall{s: r.s, opt_: make(map[string]interface{})}
	c.activityId = activityId
	c.collection = collection
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of people to include in the response, used for paging. For any
// response, the actual number returned may be less than the specified
// maxResults.
func (c *PeopleListByActivityCall) MaxResults(maxResults int64) *PeopleListByActivityCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, used to page through large result sets. To get the next page
// of results, set this parameter to the value of "nextPageToken" from
// the previous response.
func (c *PeopleListByActivityCall) PageToken(pageToken string) *PeopleListByActivityCall {
	c.opt_["pageToken"] = pageToken
	return c
}

func (c *PeopleListByActivityCall) Do() (*PeopleFeed, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/plus/v1/", "activities/{activityId}/people/{collection}")
	urls = strings.Replace(urls, "{activityId}", cleanPathString(c.activityId), 1)
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
	ret := new(PeopleFeed)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all of the people in the specified collection for a particular activity.",
	//   "httpMethod": "GET",
	//   "id": "plus.people.listByActivity",
	//   "parameterOrder": [
	//     "activityId",
	//     "collection"
	//   ],
	//   "parameters": {
	//     "activityId": {
	//       "description": "The ID of the activity to get the list of people for.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "collection": {
	//       "description": "The collection of people to list.",
	//       "enum": [
	//         "plusoners",
	//         "resharers"
	//       ],
	//       "enumDescriptions": [
	//         "List all people who have +1'd this activity.",
	//         "List all people who have reshared this activity."
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "20",
	//       "description": "The maximum number of people to include in the response, used for paging. For any response, the actual number returned may be less than the specified maxResults.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "100",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, used to page through large result sets. To get the next page of results, set this parameter to the value of \"nextPageToken\" from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities/{activityId}/people/{collection}",
	//   "response": {
	//     "$ref": "PeopleFeed"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/plus.me"
	//   ]
	// }

}

// method id "plus.people.search":

type PeopleSearchCall struct {
	s     *Service
	query string
	opt_  map[string]interface{}
}

// Search: Search all public profiles.
func (r *PeopleService) Search(query string) *PeopleSearchCall {
	c := &PeopleSearchCall{s: r.s, opt_: make(map[string]interface{})}
	c.query = query
	return c
}

// Language sets the optional parameter "language": Specify the
// preferred language to search with. See search language codes for
// available values.
func (c *PeopleSearchCall) Language(language string) *PeopleSearchCall {
	c.opt_["language"] = language
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of people to include in the response, used for paging. For any
// response, the actual number returned may be less than the specified
// maxResults.
func (c *PeopleSearchCall) MaxResults(maxResults int64) *PeopleSearchCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, used to page through large result sets. To get the next page
// of results, set this parameter to the value of "nextPageToken" from
// the previous response. This token may be of any length.
func (c *PeopleSearchCall) PageToken(pageToken string) *PeopleSearchCall {
	c.opt_["pageToken"] = pageToken
	return c
}

func (c *PeopleSearchCall) Do() (*PeopleFeed, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("query", fmt.Sprintf("%v", c.query))
	if v, ok := c.opt_["language"]; ok {
		params.Set("language", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/plus/v1/", "people")
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
	ret := new(PeopleFeed)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Search all public profiles.",
	//   "httpMethod": "GET",
	//   "id": "plus.people.search",
	//   "parameterOrder": [
	//     "query"
	//   ],
	//   "parameters": {
	//     "language": {
	//       "default": "",
	//       "description": "Specify the preferred language to search with. See search language codes for available values.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "10",
	//       "description": "The maximum number of people to include in the response, used for paging. For any response, the actual number returned may be less than the specified maxResults.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "20",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, used to page through large result sets. To get the next page of results, set this parameter to the value of \"nextPageToken\" from the previous response. This token may be of any length.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "query": {
	//       "description": "Specify a query string for full text search of public text in all profiles.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "people",
	//   "response": {
	//     "$ref": "PeopleFeed"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/plus.me"
	//   ]
	// }

}

func cleanPathString(s string) string {
	return strings.Map(func(r rune) rune {
		if r >= 0x2d && r <= 0x7a {
			return r
		}
		return -1
	}, s)
}
