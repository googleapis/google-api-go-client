// Package buzz provides access to the Buzz API.
//
// See http://code.google.com/apis/buzz/v1/using_rest.html
//
// Usage example:
//
//   import "google-api-go-client.googlecode.com/hg/buzz/v1"
//   ...
//   buzzService, err := buzz.New(oauthHttpClient)
package buzz

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

const apiId = "buzz:v1"
const apiName = "buzz"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/buzz/v1/"

// OAuth2 scopes used by this API.
const (
	// View your Buzz activity and address book
	BuzzReadonlyScope = "https://www.googleapis.com/auth/buzz.readonly"

	// Manage your photos and videos
	PicasaScope = "https://www.googleapis.com/auth/picasa"

	// Manage your Buzz activity and address book
	BuzzScope = "https://www.googleapis.com/auth/buzz"
)

func New(client *http.Client) (*Service, os.Error) {
	if client == nil {
		return nil, os.NewError("client is nil")
	}
	s := &Service{client: client}
	s.Related = &RelatedService{s: s}
	s.Groups = &GroupsService{s: s}
	s.Photos = &PhotosService{s: s}
	s.Comments = &CommentsService{s: s}
	s.PhotoAlbums = &PhotoAlbumsService{s: s}
	s.People = &PeopleService{s: s}
	s.Activities = &ActivitiesService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	Related *RelatedService

	Groups *GroupsService

	Photos *PhotosService

	Comments *CommentsService

	PhotoAlbums *PhotoAlbumsService

	People *PeopleService

	Activities *ActivitiesService
}

type RelatedService struct {
	s *Service
}

type GroupsService struct {
	s *Service
}

type PhotosService struct {
	s *Service
}

type CommentsService struct {
	s *Service
}

type PhotoAlbumsService struct {
	s *Service
}

type PeopleService struct {
	s *Service
}

type ActivitiesService struct {
	s *Service
}

type PhotosFeed struct {
	Items []*ChiliPhotosResourceJson `json:"items,omitempty"`

	Kind string `json:"kind,omitempty"`
}

type Activity struct {
	Categories []*ActivityCategories `json:"categories,omitempty"`

	UntranslatedTitle interface{} `json:"untranslatedTitle,omitempty"`

	Radius interface{} `json:"radius,omitempty"`

	PlaceId interface{} `json:"placeId,omitempty"`

	Links *ActivityLinks `json:"links,omitempty"`

	Actor *ActivityActor `json:"actor,omitempty"`

	Published string `json:"published,omitempty"`

	Geocode interface{} `json:"geocode,omitempty"`

	Kind string `json:"kind,omitempty"`

	CrosspostSource interface{} `json:"crosspostSource,omitempty"`

	Object *ActivityObject `json:"object,omitempty"`

	Updated string `json:"updated,omitempty"`

	Placeholder interface{} `json:"placeholder,omitempty"`

	Visibility *ActivityVisibility `json:"visibility,omitempty"`

	Source *ActivitySource `json:"source,omitempty"`

	Id interface{} `json:"id,omitempty"`

	Title interface{} `json:"title,omitempty"`

	PlaceName interface{} `json:"placeName,omitempty"`

	TargetLang interface{} `json:"targetLang,omitempty"`

	Verbs []string `json:"verbs,omitempty"`

	DetectedlLang interface{} `json:"detectedlLang,omitempty"`

	Annotation interface{} `json:"annotation,omitempty"`

	Address interface{} `json:"address,omitempty"`
}

type PeopleFeed struct {
	StartIndex interface{} `json:"startIndex,omitempty"`

	TotalResults interface{} `json:"totalResults,omitempty"`

	Kind string `json:"kind,omitempty"`

	Entry []*Person `json:"entry,omitempty"`

	ItemsPerPage interface{} `json:"itemsPerPage,omitempty"`
}

type ActivityVisibilityEntries struct {
	Id interface{} `json:"id,omitempty"`

	Title interface{} `json:"title,omitempty"`
}

type ActivityFeedLinks struct {

}

type CountFeedCounts struct {

}

type ActivityObject struct {
	Links *ActivityObjectLinks `json:"links,omitempty"`

	Actor *ActivityObjectActor `json:"actor,omitempty"`

	Liked []*Person `json:"liked,omitempty"`

	Content interface{} `json:"content,omitempty"`

	Id interface{} `json:"id,omitempty"`

	UntranslatedContent interface{} `json:"untranslatedContent,omitempty"`

	TargetLang interface{} `json:"targetLang,omitempty"`

	DetectedlLang interface{} `json:"detectedlLang,omitempty"`

	Comments []*Comment `json:"comments,omitempty"`

	OriginalContent interface{} `json:"originalContent,omitempty"`

	ShareOriginal *Activity `json:"shareOriginal,omitempty"`

	Type string `json:"type,omitempty"`

	Attachments []*ActivityObjectAttachments `json:"attachments,omitempty"`
}

type ChiliPhotosResourceJsonAlbum struct {
	Page_link *Link `json:"page_link,omitempty"`

	Id string `json:"id,omitempty"`
}

type ChiliPhotosResourceJson struct {
	LastModified string `json:"lastModified,omitempty"`

	Album *ChiliPhotosResourceJsonAlbum `json:"album,omitempty"`

	Links *ChiliPhotosResourceJsonLinks `json:"links,omitempty"`

	Video *Video `json:"video,omitempty"`

	Version int64 `json:"version,omitempty,string"`

	Kind string `json:"kind,omitempty"`

	Id string `json:"id,omitempty"`

	Created string `json:"created,omitempty"`

	Timestamp float64 `json:"timestamp,omitempty"`

	Title string `json:"title,omitempty"`

	FileSize int64 `json:"fileSize,omitempty,string"`

	Description string `json:"description,omitempty"`

	Owner *ChiliPhotosResourceJsonOwner `json:"owner,omitempty"`
}

type PersonEmails struct {
	Value interface{} `json:"value,omitempty"`

	Primary interface{} `json:"primary,omitempty"`

	Type interface{} `json:"type,omitempty"`
}

type GroupLinks struct {
	Self []*GroupLinksSelf `json:"self,omitempty"`
}

type ChiliPhotosResourceJsonLinks struct {
	Alternate []*Link `json:"alternate,omitempty"`
}

type PersonUrls struct {
	Value interface{} `json:"value,omitempty"`

	Primary interface{} `json:"primary,omitempty"`

	Type interface{} `json:"type,omitempty"`
}

type AlbumLiteCollectionPhoto struct {
	PhotoUrl interface{} `json:"photoUrl,omitempty"`
}

type ActivityObjectAttachmentsLinks struct {

}

type ChiliPhotosResourceJsonOwner struct {
	ThumbnailUrl string `json:"thumbnailUrl,omitempty"`

	Name string `json:"name,omitempty"`

	Id string `json:"id,omitempty"`

	ProfileUrl string `json:"profileUrl,omitempty"`
}

type AlbumOwner struct {
	ThumbnailUrl string `json:"thumbnailUrl,omitempty"`

	Name string `json:"name,omitempty"`

	Id string `json:"id,omitempty"`

	ProfileUrl string `json:"profileUrl,omitempty"`
}

type PersonAccounts struct {
	Username interface{} `json:"username,omitempty"`

	Domain interface{} `json:"domain,omitempty"`

	Userid interface{} `json:"userid,omitempty"`
}

type GroupLinksSelf struct {
	Href interface{} `json:"href,omitempty"`

	Type string `json:"type,omitempty"`
}

type ActivityLinksLiked struct {
	Count int64 `json:"count,omitempty"`

	Href interface{} `json:"href,omitempty"`

	Type interface{} `json:"type,omitempty"`
}

type PersonAddresses struct {
	Region interface{} `json:"region,omitempty"`

	Primary interface{} `json:"primary,omitempty"`

	Country interface{} `json:"country,omitempty"`

	PostalCode interface{} `json:"postalCode,omitempty"`

	Locality interface{} `json:"locality,omitempty"`

	Formatted interface{} `json:"formatted,omitempty"`

	Type interface{} `json:"type,omitempty"`

	StreetAddress interface{} `json:"streetAddress,omitempty"`
}

type ActivityObjectActor struct {
	ThumbnailUrl interface{} `json:"thumbnailUrl,omitempty"`

	Name interface{} `json:"name,omitempty"`

	Id interface{} `json:"id,omitempty"`

	ProfileUrl interface{} `json:"profileUrl,omitempty"`
}

type CommentFeedLinks struct {

}

type PersonName struct {
	MiddleName interface{} `json:"middleName,omitempty"`

	GivenName interface{} `json:"givenName,omitempty"`

	HonorificSuffix interface{} `json:"honorificSuffix,omitempty"`

	HonorificPrefix interface{} `json:"honorificPrefix,omitempty"`

	Formatted interface{} `json:"formatted,omitempty"`

	FamilyName interface{} `json:"familyName,omitempty"`
}

type RelatedFeed struct {
	Links *RelatedFeedLinks `json:"links,omitempty"`

	Items []*Related `json:"items,omitempty"`

	Kind string `json:"kind,omitempty"`

	Updated string `json:"updated,omitempty"`

	Id interface{} `json:"id,omitempty"`

	Title interface{} `json:"title,omitempty"`
}

type ActivityCategories struct {
	Schema interface{} `json:"schema,omitempty"`

	Label interface{} `json:"label,omitempty"`

	Term interface{} `json:"term,omitempty"`
}

type Album struct {
	LastModified string `json:"lastModified,omitempty"`

	Links *AlbumLinks `json:"links,omitempty"`

	Version int64 `json:"version,omitempty,string"`

	Kind string `json:"kind,omitempty"`

	Tags []string `json:"tags,omitempty"`

	Id string `json:"id,omitempty"`

	Created string `json:"created,omitempty"`

	Title string `json:"title,omitempty"`

	Description string `json:"description,omitempty"`

	Owner *AlbumOwner `json:"owner,omitempty"`

	FirstPhotoId uint64 `json:"firstPhotoId,omitempty,string"`
}

type Person struct {
	Ims []*PersonIms `json:"ims,omitempty"`

	HappiestWhen interface{} `json:"happiestWhen,omitempty"`

	CurrentLocation interface{} `json:"currentLocation,omitempty"`

	Urls []*PersonUrls `json:"urls,omitempty"`

	Fashion interface{} `json:"fashion,omitempty"`

	Pets []interface{} `json:"pets,omitempty"`

	AboutMe interface{} `json:"aboutMe,omitempty"`

	Children []interface{} `json:"children,omitempty"`

	Relationships []interface{} `json:"relationships,omitempty"`

	ThumbnailUrl interface{} `json:"thumbnailUrl,omitempty"`

	Organizations []*PersonOrganizations `json:"organizations,omitempty"`

	ScaredOf interface{} `json:"scaredOf,omitempty"`

	RelationshipStatus interface{} `json:"relationshipStatus,omitempty"`

	Name *PersonName `json:"name,omitempty"`

	Published string `json:"published,omitempty"`

	TurnOffs []interface{} `json:"turnOffs,omitempty"`

	PreferredUsername interface{} `json:"preferredUsername,omitempty"`

	Religion interface{} `json:"religion,omitempty"`

	LookingFor interface{} `json:"lookingFor,omitempty"`

	BodyType interface{} `json:"bodyType,omitempty"`

	Note interface{} `json:"note,omitempty"`

	Drinker interface{} `json:"drinker,omitempty"`

	Accounts []*PersonAccounts `json:"accounts,omitempty"`

	SexualOrientation interface{} `json:"sexualOrientation,omitempty"`

	Addresses []*PersonAddresses `json:"addresses,omitempty"`

	Kind string `json:"kind,omitempty"`

	ProfileSong interface{} `json:"profileSong,omitempty"`

	Gender interface{} `json:"gender,omitempty"`

	JobInterests []interface{} `json:"jobInterests,omitempty"`

	Anniversary interface{} `json:"anniversary,omitempty"`

	Emails []*PersonEmails `json:"emails,omitempty"`

	Updated string `json:"updated,omitempty"`

	Quotes []interface{} `json:"quotes,omitempty"`

	Tags []interface{} `json:"tags,omitempty"`

	Photos []*PersonPhotos `json:"photos,omitempty"`

	Food []interface{} `json:"food,omitempty"`

	Languages []interface{} `json:"languages,omitempty"`

	Id interface{} `json:"id,omitempty"`

	Nickname interface{} `json:"nickname,omitempty"`

	PhoneNumbers []*PersonPhoneNumbers `json:"phoneNumbers,omitempty"`

	DisplayName interface{} `json:"displayName,omitempty"`

	TurnOns []interface{} `json:"turnOns,omitempty"`

	Status interface{} `json:"status,omitempty"`

	Ethnicity interface{} `json:"ethnicity,omitempty"`

	LanguagesSpoken []interface{} `json:"languagesSpoken,omitempty"`

	Music []interface{} `json:"music,omitempty"`

	Sports []interface{} `json:"sports,omitempty"`

	Movies []interface{} `json:"movies,omitempty"`

	HasApp interface{} `json:"hasApp,omitempty"`

	Cars []interface{} `json:"cars,omitempty"`

	LivingArrangement interface{} `json:"livingArrangement,omitempty"`

	PoliticalViews []interface{} `json:"politicalViews,omitempty"`

	Romance interface{} `json:"romance,omitempty"`

	Humor interface{} `json:"humor,omitempty"`

	Books []interface{} `json:"books,omitempty"`

	UtcOffset interface{} `json:"utcOffset,omitempty"`

	ProfileVideo interface{} `json:"profileVideo,omitempty"`

	ProfileUrl interface{} `json:"profileUrl,omitempty"`

	Birthday interface{} `json:"birthday,omitempty"`

	Connected interface{} `json:"connected,omitempty"`

	Activities []interface{} `json:"activities,omitempty"`

	Smoker interface{} `json:"smoker,omitempty"`

	Heroes []interface{} `json:"heroes,omitempty"`

	TvShows []interface{} `json:"tvShows,omitempty"`

	Interests []interface{} `json:"interests,omitempty"`
}

type ActivityVisibility struct {
	Entries []*ActivityVisibilityEntries `json:"entries,omitempty"`
}

type AlbumsFeed struct {
	Items []*Album `json:"items,omitempty"`

	Kind string `json:"kind,omitempty"`
}

type Video struct {
	Streams []*Link `json:"streams,omitempty"`

	Status string `json:"status,omitempty"`

	Duration int64 `json:"duration,omitempty,string"`

	Size uint64 `json:"size,omitempty,string"`
}

type AlbumLite struct {
	Collection *AlbumLiteCollection `json:"collection,omitempty"`

	Kind string `json:"kind,omitempty"`
}

type GroupFeed struct {
	Links *GroupFeedLinks `json:"links,omitempty"`

	Items []*Group `json:"items,omitempty"`

	Kind string `json:"kind,omitempty"`
}

type PersonPhoneNumbers struct {
	Value interface{} `json:"value,omitempty"`

	Primary interface{} `json:"primary,omitempty"`

	Type interface{} `json:"type,omitempty"`
}

type ActivityActor struct {
	ThumbnailUrl interface{} `json:"thumbnailUrl,omitempty"`

	Name interface{} `json:"name,omitempty"`

	Id interface{} `json:"id,omitempty"`

	ProfileUrl interface{} `json:"profileUrl,omitempty"`
}

type AlbumLiteCollection struct {
	AlbumId interface{} `json:"albumId,omitempty"`

	Photo *AlbumLiteCollectionPhoto `json:"photo,omitempty"`

	Album interface{} `json:"album,omitempty"`
}

type ActivityLinks struct {
	Liked []*ActivityLinksLiked `json:"liked,omitempty"`
}

type CountFeed struct {
	Kind string `json:"kind,omitempty"`

	Counts *CountFeedCounts `json:"counts,omitempty"`
}

type ActivityFeed struct {
	Links *ActivityFeedLinks `json:"links,omitempty"`

	Items []*Activity `json:"items,omitempty"`

	Kind string `json:"kind,omitempty"`

	Updated string `json:"updated,omitempty"`

	Id interface{} `json:"id,omitempty"`

	Title interface{} `json:"title,omitempty"`
}

type CommentActor struct {
	ThumbnailUrl interface{} `json:"thumbnailUrl,omitempty"`

	Name interface{} `json:"name,omitempty"`

	Id interface{} `json:"id,omitempty"`

	ProfileUrl interface{} `json:"profileUrl,omitempty"`
}

type GroupFeedLinks struct {

}

type Comment struct {
	Links *CommentLinks `json:"links,omitempty"`

	Actor *CommentActor `json:"actor,omitempty"`

	Published string `json:"published,omitempty"`

	Kind string `json:"kind,omitempty"`

	Updated string `json:"updated,omitempty"`

	Content interface{} `json:"content,omitempty"`

	Placeholder interface{} `json:"placeholder,omitempty"`

	Id interface{} `json:"id,omitempty"`

	UntranslatedContent interface{} `json:"untranslatedContent,omitempty"`

	TargetLang interface{} `json:"targetLang,omitempty"`

	OriginalContent interface{} `json:"originalContent,omitempty"`

	DetectedLang interface{} `json:"detectedLang,omitempty"`
}

type ActivityObjectLinks struct {

}

type Group struct {
	Links *GroupLinks `json:"links,omitempty"`

	Kind string `json:"kind,omitempty"`

	Id interface{} `json:"id,omitempty"`

	Title interface{} `json:"title,omitempty"`

	MemberCount interface{} `json:"memberCount,omitempty"`
}

type CommentLinks struct {
	InReplyTo []*CommentLinksInReplyTo `json:"inReplyTo,omitempty"`
}

type PersonIms struct {
	Value interface{} `json:"value,omitempty"`

	Primary interface{} `json:"primary,omitempty"`

	Type interface{} `json:"type,omitempty"`
}

type Link struct {
	Count int64 `json:"count,omitempty"`

	Height int64 `json:"height,omitempty"`

	Width int64 `json:"width,omitempty"`

	Href string `json:"href,omitempty"`

	Updated string `json:"updated,omitempty"`

	Title string `json:"title,omitempty"`

	Type string `json:"type,omitempty"`
}

type PersonOrganizations struct {
	Location interface{} `json:"location,omitempty"`

	Name interface{} `json:"name,omitempty"`

	Department interface{} `json:"department,omitempty"`

	Primary interface{} `json:"primary,omitempty"`

	Title interface{} `json:"title,omitempty"`

	StartDate interface{} `json:"startDate,omitempty"`

	EndDate interface{} `json:"endDate,omitempty"`

	Type interface{} `json:"type,omitempty"`

	Description interface{} `json:"description,omitempty"`
}

type ActivitySource struct {
	Title interface{} `json:"title,omitempty"`
}

type Related struct {
	Kind string `json:"kind,omitempty"`

	Href interface{} `json:"href,omitempty"`

	Id interface{} `json:"id,omitempty"`

	Title interface{} `json:"title,omitempty"`

	Summary interface{} `json:"summary,omitempty"`
}

type CommentLinksInReplyTo struct {
	Href interface{} `json:"href,omitempty"`

	Source interface{} `json:"source,omitempty"`

	Ref interface{} `json:"ref,omitempty"`
}

type ActivityObjectAttachments struct {
	Links *ActivityObjectAttachmentsLinks `json:"links,omitempty"`

	Content interface{} `json:"content,omitempty"`

	Id interface{} `json:"id,omitempty"`

	Title interface{} `json:"title,omitempty"`

	Type string `json:"type,omitempty"`
}

type RelatedFeedLinks struct {

}

type PersonPhotos struct {
	Value interface{} `json:"value,omitempty"`

	Primary interface{} `json:"primary,omitempty"`

	Height interface{} `json:"height,omitempty"`

	Width interface{} `json:"width,omitempty"`

	Type interface{} `json:"type,omitempty"`
}

type AlbumLinks struct {
	Enclosure *Link `json:"enclosure,omitempty"`

	Alternate *Link `json:"alternate,omitempty"`
}

type CommentFeed struct {
	Links *CommentFeedLinks `json:"links,omitempty"`

	Items []*Comment `json:"items,omitempty"`

	Kind string `json:"kind,omitempty"`

	Updated string `json:"updated,omitempty"`

	Id interface{} `json:"id,omitempty"`

	Title interface{} `json:"title,omitempty"`
}

// method id "chili.related.list":

type RelatedListCall struct {
	s      *Service
	userId string
	scope  string
	postId string
	opt_   map[string]interface{}
}

// List: Get related links for an activity
func (r *RelatedService) List(userId string, scope string, postId string) *RelatedListCall {
	c := &RelatedListCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.scope = scope
	c.postId = postId
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *RelatedListCall) Hl(hl string) *RelatedListCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *RelatedListCall) Alt(alt string) *RelatedListCall {
	c.opt_["alt"] = alt
	return c
}

func (c *RelatedListCall) Do() (*RelatedFeed, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "activities/{userId}/{scope}/{postId}/@related")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{scope}", cleanPathString(c.scope), 1)
	urls = strings.Replace(urls, "{postId}", cleanPathString(c.postId), 1)
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
	ret := new(RelatedFeed)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get related links for an activity",
	//   "httpMethod": "GET",
	//   "id": "chili.related.list",
	//   "parameterOrder": [
	//     "userId",
	//     "scope",
	//     "postId"
	//   ],
	//   "parameters": {
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "postId": {
	//       "description": "ID of the activity to which to get related links.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "scope": {
	//       "description": "The collection to which the activity belongs.",
	//       "enum": [
	//         "@self"
	//       ],
	//       "enumDescriptions": [
	//         "Activities posted by the user."
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities/{userId}/{scope}/{postId}/@related",
	//   "response": {
	//     "$ref": "RelatedFeed"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz",
	//     "https://www.googleapis.com/auth/buzz.readonly"
	//   ]
	// }

}

// method id "chili.groups.list":

type GroupsListCall struct {
	s      *Service
	userId string
	opt_   map[string]interface{}
}

// List: Get a user's groups
func (r *GroupsService) List(userId string) *GroupsListCall {
	c := &GroupsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	return c
}

// C sets the optional parameter "c": A continuation token that allows
// pagination.
func (c *GroupsListCall) C(c1 string) *GroupsListCall {
	c.opt_["c"] = c1
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *GroupsListCall) Hl(hl string) *GroupsListCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *GroupsListCall) Alt(alt string) *GroupsListCall {
	c.opt_["alt"] = alt
	return c
}

// MaxResults sets the optional parameter "max-results": Maximum number
// of results to include.
func (c *GroupsListCall) MaxResults(maxResults int64) *GroupsListCall {
	c.opt_["max-results"] = maxResults
	return c
}

func (c *GroupsListCall) Do() (*GroupFeed, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["c"]; ok {
		params.Set("c", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["max-results"]; ok {
		params.Set("max-results", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "people/{userId}/@groups")
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
	ret := new(GroupFeed)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get a user's groups",
	//   "httpMethod": "GET",
	//   "id": "chili.groups.list",
	//   "parameterOrder": [
	//     "userId"
	//   ],
	//   "parameters": {
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "c": {
	//       "description": "A continuation token that allows pagination.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "max-results": {
	//       "default": "20",
	//       "description": "Maximum number of results to include.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "userId": {
	//       "description": "ID of the user being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "people/{userId}/@groups",
	//   "response": {
	//     "$ref": "GroupFeed"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz",
	//     "https://www.googleapis.com/auth/buzz.readonly"
	//   ]
	// }

}

// method id "chili.groups.update":

type GroupsUpdateCall struct {
	s       *Service
	userId  string
	groupId string
	group   *Group
	opt_    map[string]interface{}
}

// Update: Update a group
func (r *GroupsService) Update(userId string, groupId string, group *Group) *GroupsUpdateCall {
	c := &GroupsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.groupId = groupId
	c.group = group
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *GroupsUpdateCall) Hl(hl string) *GroupsUpdateCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *GroupsUpdateCall) Alt(alt string) *GroupsUpdateCall {
	c.opt_["alt"] = alt
	return c
}

func (c *GroupsUpdateCall) Do() (*Group, os.Error) {
	var body io.Reader = nil
	body, err := googleapi.WithDataWrapper.JSONReader(c.group)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "people/{userId}/@groups/{groupId}/@self")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{groupId}", cleanPathString(c.groupId), 1)
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
	ret := new(Group)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update a group",
	//   "httpMethod": "PUT",
	//   "id": "chili.groups.update",
	//   "parameterOrder": [
	//     "userId",
	//     "groupId"
	//   ],
	//   "parameters": {
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "groupId": {
	//       "description": "ID of the group to update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "people/{userId}/@groups/{groupId}/@self",
	//   "request": {
	//     "$ref": "Group"
	//   },
	//   "response": {
	//     "$ref": "Group"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz"
	//   ]
	// }

}

// method id "chili.groups.insert":

type GroupsInsertCall struct {
	s      *Service
	userId string
	group  *Group
	opt_   map[string]interface{}
}

// Insert: Create a group
func (r *GroupsService) Insert(userId string, group *Group) *GroupsInsertCall {
	c := &GroupsInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.group = group
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *GroupsInsertCall) Hl(hl string) *GroupsInsertCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *GroupsInsertCall) Alt(alt string) *GroupsInsertCall {
	c.opt_["alt"] = alt
	return c
}

func (c *GroupsInsertCall) Do() (*Group, os.Error) {
	var body io.Reader = nil
	body, err := googleapi.WithDataWrapper.JSONReader(c.group)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "people/{userId}/@groups")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
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
	ret := new(Group)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Create a group",
	//   "httpMethod": "POST",
	//   "id": "chili.groups.insert",
	//   "parameterOrder": [
	//     "userId"
	//   ],
	//   "parameters": {
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "people/{userId}/@groups",
	//   "request": {
	//     "$ref": "Group"
	//   },
	//   "response": {
	//     "$ref": "Group"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz"
	//   ]
	// }

}

// method id "chili.groups.get":

type GroupsGetCall struct {
	s       *Service
	userId  string
	groupId string
	opt_    map[string]interface{}
}

// Get: Get a group
func (r *GroupsService) Get(userId string, groupId string) *GroupsGetCall {
	c := &GroupsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.groupId = groupId
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *GroupsGetCall) Hl(hl string) *GroupsGetCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *GroupsGetCall) Alt(alt string) *GroupsGetCall {
	c.opt_["alt"] = alt
	return c
}

func (c *GroupsGetCall) Do() (*Group, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "people/{userId}/@groups/{groupId}/@self")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{groupId}", cleanPathString(c.groupId), 1)
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
	ret := new(Group)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get a group",
	//   "httpMethod": "GET",
	//   "id": "chili.groups.get",
	//   "parameterOrder": [
	//     "userId",
	//     "groupId"
	//   ],
	//   "parameters": {
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "groupId": {
	//       "description": "ID of the group to get.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "people/{userId}/@groups/{groupId}/@self",
	//   "response": {
	//     "$ref": "Group"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz",
	//     "https://www.googleapis.com/auth/buzz.readonly"
	//   ]
	// }

}

// method id "chili.groups.delete":

type GroupsDeleteCall struct {
	s       *Service
	userId  string
	groupId string
	opt_    map[string]interface{}
}

// Delete: Delete a group
func (r *GroupsService) Delete(userId string, groupId string) *GroupsDeleteCall {
	c := &GroupsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.groupId = groupId
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *GroupsDeleteCall) Hl(hl string) *GroupsDeleteCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *GroupsDeleteCall) Alt(alt string) *GroupsDeleteCall {
	c.opt_["alt"] = alt
	return c
}

func (c *GroupsDeleteCall) Do() os.Error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "people/{userId}/@groups/{groupId}")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{groupId}", cleanPathString(c.groupId), 1)
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
	//   "description": "Delete a group",
	//   "httpMethod": "DELETE",
	//   "id": "chili.groups.delete",
	//   "parameterOrder": [
	//     "userId",
	//     "groupId"
	//   ],
	//   "parameters": {
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "groupId": {
	//       "description": "ID of the group to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "people/{userId}/@groups/{groupId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz"
	//   ]
	// }

}

// method id "chili.groups.patch":

type GroupsPatchCall struct {
	s       *Service
	userId  string
	groupId string
	group   *Group
	opt_    map[string]interface{}
}

// Patch: Update a group. This method supports patch semantics.
func (r *GroupsService) Patch(userId string, groupId string, group *Group) *GroupsPatchCall {
	c := &GroupsPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.groupId = groupId
	c.group = group
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *GroupsPatchCall) Hl(hl string) *GroupsPatchCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *GroupsPatchCall) Alt(alt string) *GroupsPatchCall {
	c.opt_["alt"] = alt
	return c
}

func (c *GroupsPatchCall) Do() (*Group, os.Error) {
	var body io.Reader = nil
	body, err := googleapi.WithDataWrapper.JSONReader(c.group)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "people/{userId}/@groups/{groupId}/@self")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{groupId}", cleanPathString(c.groupId), 1)
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
	ret := new(Group)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update a group. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "chili.groups.patch",
	//   "parameterOrder": [
	//     "userId",
	//     "groupId"
	//   ],
	//   "parameters": {
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "groupId": {
	//       "description": "ID of the group to update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "people/{userId}/@groups/{groupId}/@self",
	//   "request": {
	//     "$ref": "Group"
	//   },
	//   "response": {
	//     "$ref": "Group"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz"
	//   ]
	// }

}

// method id "chili.photos.insert":

type PhotosInsertCall struct {
	s         *Service
	userId    string
	albumId   string
	albumlite *AlbumLite
	opt_      map[string]interface{}
	media_    io.Reader
}

// Insert: Upload a photo to an album
func (r *PhotosService) Insert(userId string, albumId string, albumlite *AlbumLite) *PhotosInsertCall {
	c := &PhotosInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.albumId = albumId
	c.albumlite = albumlite
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *PhotosInsertCall) Hl(hl string) *PhotosInsertCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *PhotosInsertCall) Alt(alt string) *PhotosInsertCall {
	c.opt_["alt"] = alt
	return c
}
func (c *PhotosInsertCall) Media(r io.Reader) *PhotosInsertCall {
	c.media_ = r
	return c
}

func (c *PhotosInsertCall) Do() (*AlbumLite, os.Error) {
	var body io.Reader = nil
	body, err := googleapi.WithDataWrapper.JSONReader(c.albumlite)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "photos/{userId}/{albumId}")
	if c.media_ != nil {
		urls = strings.Replace(urls, "https://www.googleapis.com/", "https://www.googleapis.com/upload/", 1)
	}
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{albumId}", cleanPathString(c.albumId), 1)
	urls += "?" + params.Encode()
	contentLength_, hasMedia_ := googleapi.ConditionallyIncludeMedia(c.media_, &body, &ctype)
	req, _ := http.NewRequest("POST", urls, body)
	if hasMedia_ {
		req.ContentLength = contentLength_
	}
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(AlbumLite)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Upload a photo to an album",
	//   "httpMethod": "POST",
	//   "id": "chili.photos.insert",
	//   "mediaUpload": {
	//     "accept": [
	//       "image/*"
	//     ],
	//     "maxSize": "30MB",
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/buzz/v1/photos/{userId}/{albumId}"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/buzz/v1/photos/{userId}/{albumId}"
	//       }
	//     }
	//   },
	//   "parameterOrder": [
	//     "userId",
	//     "albumId"
	//   ],
	//   "parameters": {
	//     "albumId": {
	//       "description": "ID of the album to which to upload.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "photos/{userId}/{albumId}",
	//   "request": {
	//     "$ref": "AlbumLite"
	//   },
	//   "response": {
	//     "$ref": "AlbumLite"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz"
	//   ]
	// }

}

// method id "chili.photos.get":

type PhotosGetCall struct {
	s       *Service
	userId  string
	albumId string
	photoId string
	opt_    map[string]interface{}
}

// Get: Get photo metadata
func (r *PhotosService) Get(userId string, albumId string, photoId string) *PhotosGetCall {
	c := &PhotosGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.albumId = albumId
	c.photoId = photoId
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *PhotosGetCall) Hl(hl string) *PhotosGetCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *PhotosGetCall) Alt(alt string) *PhotosGetCall {
	c.opt_["alt"] = alt
	return c
}

func (c *PhotosGetCall) Do() (*ChiliPhotosResourceJson, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "photos/{userId}/@self/{albumId}/@photos/{photoId}")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{albumId}", cleanPathString(c.albumId), 1)
	urls = strings.Replace(urls, "{photoId}", cleanPathString(c.photoId), 1)
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
	ret := new(ChiliPhotosResourceJson)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get photo metadata",
	//   "httpMethod": "GET",
	//   "id": "chili.photos.get",
	//   "parameterOrder": [
	//     "userId",
	//     "albumId",
	//     "photoId"
	//   ],
	//   "parameters": {
	//     "albumId": {
	//       "description": "ID of the album containing the photo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "photoId": {
	//       "description": "ID of the photo for which to get metadata.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "photos/{userId}/@self/{albumId}/@photos/{photoId}",
	//   "response": {
	//     "$ref": "ChiliPhotosResourceJson"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/picasa"
	//   ]
	// }

}

// method id "chili.photos.insert2":

type PhotosInsert2Call struct {
	s                       *Service
	userId                  string
	albumId                 string
	chiliphotosresourcejson *ChiliPhotosResourceJson
	opt_                    map[string]interface{}
	media_                  io.Reader
}

// Insert2: Upload a photo to an album
func (r *PhotosService) Insert2(userId string, albumId string, chiliphotosresourcejson *ChiliPhotosResourceJson) *PhotosInsert2Call {
	c := &PhotosInsert2Call{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.albumId = albumId
	c.chiliphotosresourcejson = chiliphotosresourcejson
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *PhotosInsert2Call) Hl(hl string) *PhotosInsert2Call {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *PhotosInsert2Call) Alt(alt string) *PhotosInsert2Call {
	c.opt_["alt"] = alt
	return c
}
func (c *PhotosInsert2Call) Media(r io.Reader) *PhotosInsert2Call {
	c.media_ = r
	return c
}

func (c *PhotosInsert2Call) Do() (*ChiliPhotosResourceJson, os.Error) {
	var body io.Reader = nil
	body, err := googleapi.WithDataWrapper.JSONReader(c.chiliphotosresourcejson)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "photos/{userId}/@self/{albumId}/@photos")
	if c.media_ != nil {
		urls = strings.Replace(urls, "https://www.googleapis.com/", "https://www.googleapis.com/upload/", 1)
	}
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{albumId}", cleanPathString(c.albumId), 1)
	urls += "?" + params.Encode()
	contentLength_, hasMedia_ := googleapi.ConditionallyIncludeMedia(c.media_, &body, &ctype)
	req, _ := http.NewRequest("POST", urls, body)
	if hasMedia_ {
		req.ContentLength = contentLength_
	}
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(ChiliPhotosResourceJson)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Upload a photo to an album",
	//   "httpMethod": "POST",
	//   "id": "chili.photos.insert2",
	//   "mediaUpload": {
	//     "accept": [
	//       "image/*"
	//     ],
	//     "maxSize": "30MB",
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/buzz/v1/photos/{userId}/@self/{albumId}/@photos"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/buzz/v1/photos/{userId}/@self/{albumId}/@photos"
	//       }
	//     }
	//   },
	//   "parameterOrder": [
	//     "userId",
	//     "albumId"
	//   ],
	//   "parameters": {
	//     "albumId": {
	//       "description": "ID of the album to which to upload.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "photos/{userId}/@self/{albumId}/@photos",
	//   "request": {
	//     "$ref": "ChiliPhotosResourceJson"
	//   },
	//   "response": {
	//     "$ref": "ChiliPhotosResourceJson"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/picasa"
	//   ]
	// }

}

// method id "chili.photos.delete":

type PhotosDeleteCall struct {
	s       *Service
	userId  string
	albumId string
	photoId string
	opt_    map[string]interface{}
}

// Delete: Delete a photo
func (r *PhotosService) Delete(userId string, albumId string, photoId string) *PhotosDeleteCall {
	c := &PhotosDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.albumId = albumId
	c.photoId = photoId
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *PhotosDeleteCall) Hl(hl string) *PhotosDeleteCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *PhotosDeleteCall) Alt(alt string) *PhotosDeleteCall {
	c.opt_["alt"] = alt
	return c
}

func (c *PhotosDeleteCall) Do() os.Error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "photos/{userId}/@self/{albumId}/@photos/{photoId}")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{albumId}", cleanPathString(c.albumId), 1)
	urls = strings.Replace(urls, "{photoId}", cleanPathString(c.photoId), 1)
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
	//   "description": "Delete a photo",
	//   "httpMethod": "DELETE",
	//   "id": "chili.photos.delete",
	//   "parameterOrder": [
	//     "userId",
	//     "albumId",
	//     "photoId"
	//   ],
	//   "parameters": {
	//     "albumId": {
	//       "description": "ID of the album to which to photo belongs.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "photoId": {
	//       "description": "ID of the photo to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "photos/{userId}/@self/{albumId}/@photos/{photoId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/picasa"
	//   ]
	// }

}

// method id "chili.photos.listByAlbum":

type PhotosListByAlbumCall struct {
	s       *Service
	userId  string
	albumId string
	opt_    map[string]interface{}
}

// ListByAlbum: List photos in an album
func (r *PhotosService) ListByAlbum(userId string, albumId string) *PhotosListByAlbumCall {
	c := &PhotosListByAlbumCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.albumId = albumId
	return c
}

// C sets the optional parameter "c": A continuation token that allows
// pagination.
func (c *PhotosListByAlbumCall) C(c1 string) *PhotosListByAlbumCall {
	c.opt_["c"] = c1
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *PhotosListByAlbumCall) Hl(hl string) *PhotosListByAlbumCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *PhotosListByAlbumCall) Alt(alt string) *PhotosListByAlbumCall {
	c.opt_["alt"] = alt
	return c
}

// MaxResults sets the optional parameter "max-results": Maximum number
// of results to include.
func (c *PhotosListByAlbumCall) MaxResults(maxResults int64) *PhotosListByAlbumCall {
	c.opt_["max-results"] = maxResults
	return c
}

func (c *PhotosListByAlbumCall) Do() (*PhotosFeed, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["c"]; ok {
		params.Set("c", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["max-results"]; ok {
		params.Set("max-results", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "photos/{userId}/@self/{albumId}/@photos")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{albumId}", cleanPathString(c.albumId), 1)
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
	ret := new(PhotosFeed)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List photos in an album",
	//   "httpMethod": "GET",
	//   "id": "chili.photos.listByAlbum",
	//   "parameterOrder": [
	//     "userId",
	//     "albumId"
	//   ],
	//   "parameters": {
	//     "albumId": {
	//       "description": "ID of the album for which to list photos.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "c": {
	//       "description": "A continuation token that allows pagination.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "max-results": {
	//       "default": "20",
	//       "description": "Maximum number of results to include.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "userId": {
	//       "description": "ID of the user being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "photos/{userId}/@self/{albumId}/@photos",
	//   "response": {
	//     "$ref": "PhotosFeed"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/picasa"
	//   ]
	// }

}

// method id "chili.photos.listByScope":

type PhotosListByScopeCall struct {
	s      *Service
	userId string
	scope  string
	opt_   map[string]interface{}
}

// ListByScope: Get a user's photos
func (r *PhotosService) ListByScope(userId string, scope string) *PhotosListByScopeCall {
	c := &PhotosListByScopeCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.scope = scope
	return c
}

// C sets the optional parameter "c": A continuation token that allows
// pagination.
func (c *PhotosListByScopeCall) C(c1 string) *PhotosListByScopeCall {
	c.opt_["c"] = c1
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *PhotosListByScopeCall) Hl(hl string) *PhotosListByScopeCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *PhotosListByScopeCall) Alt(alt string) *PhotosListByScopeCall {
	c.opt_["alt"] = alt
	return c
}

// MaxResults sets the optional parameter "max-results": Maximum number
// of results to include.
func (c *PhotosListByScopeCall) MaxResults(maxResults int64) *PhotosListByScopeCall {
	c.opt_["max-results"] = maxResults
	return c
}

func (c *PhotosListByScopeCall) Do() (*PhotosFeed, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["c"]; ok {
		params.Set("c", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["max-results"]; ok {
		params.Set("max-results", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "photos/{userId}/@self/{scope}/@photos")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{scope}", cleanPathString(c.scope), 1)
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
	ret := new(PhotosFeed)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get a user's photos",
	//   "httpMethod": "GET",
	//   "id": "chili.photos.listByScope",
	//   "parameterOrder": [
	//     "userId",
	//     "scope"
	//   ],
	//   "parameters": {
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "c": {
	//       "description": "A continuation token that allows pagination.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "max-results": {
	//       "default": "20",
	//       "description": "Maximum number of results to include.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "scope": {
	//       "description": "The collection of photos to list.",
	//       "enum": [
	//         "@recent"
	//       ],
	//       "enumDescriptions": [
	//         "Recent photos uploaded by the user."
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "photos/{userId}/@self/{scope}/@photos",
	//   "response": {
	//     "$ref": "PhotosFeed"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/picasa"
	//   ]
	// }

}

// method id "chili.comments.list":

type CommentsListCall struct {
	s      *Service
	userId string
	scope  string
	postId string
	opt_   map[string]interface{}
}

// List: List comments
func (r *CommentsService) List(userId string, scope string, postId string) *CommentsListCall {
	c := &CommentsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.scope = scope
	c.postId = postId
	return c
}

// C sets the optional parameter "c": A continuation token that allows
// pagination.
func (c *CommentsListCall) C(c1 string) *CommentsListCall {
	c.opt_["c"] = c1
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *CommentsListCall) Hl(hl string) *CommentsListCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *CommentsListCall) Alt(alt string) *CommentsListCall {
	c.opt_["alt"] = alt
	return c
}

// MaxResults sets the optional parameter "max-results": Maximum number
// of results to include.
func (c *CommentsListCall) MaxResults(maxResults int64) *CommentsListCall {
	c.opt_["max-results"] = maxResults
	return c
}

func (c *CommentsListCall) Do() (*CommentFeed, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["c"]; ok {
		params.Set("c", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["max-results"]; ok {
		params.Set("max-results", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "activities/{userId}/{scope}/{postId}/@comments")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{scope}", cleanPathString(c.scope), 1)
	urls = strings.Replace(urls, "{postId}", cleanPathString(c.postId), 1)
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
	//   "description": "List comments",
	//   "httpMethod": "GET",
	//   "id": "chili.comments.list",
	//   "parameterOrder": [
	//     "userId",
	//     "scope",
	//     "postId"
	//   ],
	//   "parameters": {
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "c": {
	//       "description": "A continuation token that allows pagination.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "max-results": {
	//       "default": "20",
	//       "description": "Maximum number of results to include.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "postId": {
	//       "description": "ID of the activity for which to get comments.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "scope": {
	//       "description": "The collection to which the activity belongs.",
	//       "enum": [
	//         "@self"
	//       ],
	//       "enumDescriptions": [
	//         "Activities posted by the user."
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user for whose post to get comments.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities/{userId}/{scope}/{postId}/@comments",
	//   "response": {
	//     "$ref": "CommentFeed"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz",
	//     "https://www.googleapis.com/auth/buzz.readonly"
	//   ]
	// }

}

// method id "chili.comments.update":

type CommentsUpdateCall struct {
	s         *Service
	userId    string
	scope     string
	postId    string
	commentId string
	comment   *Comment
	opt_      map[string]interface{}
}

// Update: Update a comment
func (r *CommentsService) Update(userId string, scope string, postId string, commentId string, comment *Comment) *CommentsUpdateCall {
	c := &CommentsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.scope = scope
	c.postId = postId
	c.commentId = commentId
	c.comment = comment
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *CommentsUpdateCall) Hl(hl string) *CommentsUpdateCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *CommentsUpdateCall) Alt(alt string) *CommentsUpdateCall {
	c.opt_["alt"] = alt
	return c
}

// AbuseType sets the optional parameter "abuseType": 
func (c *CommentsUpdateCall) AbuseType(abuseType string) *CommentsUpdateCall {
	c.opt_["abuseType"] = abuseType
	return c
}

func (c *CommentsUpdateCall) Do() (*Comment, os.Error) {
	var body io.Reader = nil
	body, err := googleapi.WithDataWrapper.JSONReader(c.comment)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["abuseType"]; ok {
		params.Set("abuseType", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "activities/{userId}/{scope}/{postId}/@comments/{commentId}")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{scope}", cleanPathString(c.scope), 1)
	urls = strings.Replace(urls, "{postId}", cleanPathString(c.postId), 1)
	urls = strings.Replace(urls, "{commentId}", cleanPathString(c.commentId), 1)
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
	ret := new(Comment)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update a comment",
	//   "httpMethod": "PUT",
	//   "id": "chili.comments.update",
	//   "parameterOrder": [
	//     "userId",
	//     "scope",
	//     "postId",
	//     "commentId"
	//   ],
	//   "parameters": {
	//     "abuseType": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "commentId": {
	//       "description": "ID of the comment being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "postId": {
	//       "description": "ID of the activity for which to update the comment.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "scope": {
	//       "description": "The collection to which the activity belongs.",
	//       "enum": [
	//         "@abuse",
	//         "@self"
	//       ],
	//       "enumDescriptions": [
	//         "Comments reported by the user.",
	//         "Comments posted by the user."
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities/{userId}/{scope}/{postId}/@comments/{commentId}",
	//   "request": {
	//     "$ref": "Comment"
	//   },
	//   "response": {
	//     "$ref": "Comment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz"
	//   ]
	// }

}

// method id "chili.comments.insert":

type CommentsInsertCall struct {
	s       *Service
	userId  string
	postId  string
	comment *Comment
	opt_    map[string]interface{}
}

// Insert: Create a comment
func (r *CommentsService) Insert(userId string, postId string, comment *Comment) *CommentsInsertCall {
	c := &CommentsInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.postId = postId
	c.comment = comment
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *CommentsInsertCall) Hl(hl string) *CommentsInsertCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *CommentsInsertCall) Alt(alt string) *CommentsInsertCall {
	c.opt_["alt"] = alt
	return c
}

func (c *CommentsInsertCall) Do() (*Comment, os.Error) {
	var body io.Reader = nil
	body, err := googleapi.WithDataWrapper.JSONReader(c.comment)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "activities/{userId}/@self/{postId}/@comments")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{postId}", cleanPathString(c.postId), 1)
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
	//   "description": "Create a comment",
	//   "httpMethod": "POST",
	//   "id": "chili.comments.insert",
	//   "parameterOrder": [
	//     "userId",
	//     "postId"
	//   ],
	//   "parameters": {
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "postId": {
	//       "description": "ID of the activity on which to comment.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user on whose behalf to comment.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities/{userId}/@self/{postId}/@comments",
	//   "request": {
	//     "$ref": "Comment"
	//   },
	//   "response": {
	//     "$ref": "Comment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz"
	//   ]
	// }

}

// method id "chili.comments.get":

type CommentsGetCall struct {
	s         *Service
	userId    string
	postId    string
	commentId string
	opt_      map[string]interface{}
}

// Get: Get a comment
func (r *CommentsService) Get(userId string, postId string, commentId string) *CommentsGetCall {
	c := &CommentsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.postId = postId
	c.commentId = commentId
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *CommentsGetCall) Hl(hl string) *CommentsGetCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *CommentsGetCall) Alt(alt string) *CommentsGetCall {
	c.opt_["alt"] = alt
	return c
}

func (c *CommentsGetCall) Do() (*Comment, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "activities/{userId}/@self/{postId}/@comments/{commentId}")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{postId}", cleanPathString(c.postId), 1)
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
	//   "description": "Get a comment",
	//   "httpMethod": "GET",
	//   "id": "chili.comments.get",
	//   "parameterOrder": [
	//     "userId",
	//     "postId",
	//     "commentId"
	//   ],
	//   "parameters": {
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "commentId": {
	//       "description": "ID of the comment being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "postId": {
	//       "description": "ID of the activity for which to get comments.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities/{userId}/@self/{postId}/@comments/{commentId}",
	//   "response": {
	//     "$ref": "Comment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz",
	//     "https://www.googleapis.com/auth/buzz.readonly"
	//   ]
	// }

}

// method id "chili.comments.delete":

type CommentsDeleteCall struct {
	s         *Service
	userId    string
	postId    string
	commentId string
	opt_      map[string]interface{}
}

// Delete: Delete a comment
func (r *CommentsService) Delete(userId string, postId string, commentId string) *CommentsDeleteCall {
	c := &CommentsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.postId = postId
	c.commentId = commentId
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *CommentsDeleteCall) Hl(hl string) *CommentsDeleteCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *CommentsDeleteCall) Alt(alt string) *CommentsDeleteCall {
	c.opt_["alt"] = alt
	return c
}

func (c *CommentsDeleteCall) Do() os.Error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "activities/{userId}/@self/{postId}/@comments/{commentId}")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{postId}", cleanPathString(c.postId), 1)
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
	//   "description": "Delete a comment",
	//   "httpMethod": "DELETE",
	//   "id": "chili.comments.delete",
	//   "parameterOrder": [
	//     "userId",
	//     "postId",
	//     "commentId"
	//   ],
	//   "parameters": {
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "commentId": {
	//       "description": "ID of the comment being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "postId": {
	//       "description": "ID of the activity for which to delete the comment.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities/{userId}/@self/{postId}/@comments/{commentId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz"
	//   ]
	// }

}

// method id "chili.comments.patch":

type CommentsPatchCall struct {
	s         *Service
	userId    string
	scope     string
	postId    string
	commentId string
	comment   *Comment
	opt_      map[string]interface{}
}

// Patch: Update a comment. This method supports patch semantics.
func (r *CommentsService) Patch(userId string, scope string, postId string, commentId string, comment *Comment) *CommentsPatchCall {
	c := &CommentsPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.scope = scope
	c.postId = postId
	c.commentId = commentId
	c.comment = comment
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *CommentsPatchCall) Hl(hl string) *CommentsPatchCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *CommentsPatchCall) Alt(alt string) *CommentsPatchCall {
	c.opt_["alt"] = alt
	return c
}

// AbuseType sets the optional parameter "abuseType": 
func (c *CommentsPatchCall) AbuseType(abuseType string) *CommentsPatchCall {
	c.opt_["abuseType"] = abuseType
	return c
}

func (c *CommentsPatchCall) Do() (*Comment, os.Error) {
	var body io.Reader = nil
	body, err := googleapi.WithDataWrapper.JSONReader(c.comment)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["abuseType"]; ok {
		params.Set("abuseType", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "activities/{userId}/{scope}/{postId}/@comments/{commentId}")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{scope}", cleanPathString(c.scope), 1)
	urls = strings.Replace(urls, "{postId}", cleanPathString(c.postId), 1)
	urls = strings.Replace(urls, "{commentId}", cleanPathString(c.commentId), 1)
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
	ret := new(Comment)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update a comment. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "chili.comments.patch",
	//   "parameterOrder": [
	//     "userId",
	//     "scope",
	//     "postId",
	//     "commentId"
	//   ],
	//   "parameters": {
	//     "abuseType": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "commentId": {
	//       "description": "ID of the comment being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "postId": {
	//       "description": "ID of the activity for which to update the comment.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "scope": {
	//       "description": "The collection to which the activity belongs.",
	//       "enum": [
	//         "@abuse",
	//         "@self"
	//       ],
	//       "enumDescriptions": [
	//         "Comments reported by the user.",
	//         "Comments posted by the user."
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities/{userId}/{scope}/{postId}/@comments/{commentId}",
	//   "request": {
	//     "$ref": "Comment"
	//   },
	//   "response": {
	//     "$ref": "Comment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz"
	//   ]
	// }

}

// method id "chili.photoAlbums.list":

type PhotoAlbumsListCall struct {
	s      *Service
	userId string
	scope  string
	opt_   map[string]interface{}
}

// List: List a user's photo albums
func (r *PhotoAlbumsService) List(userId string, scope string) *PhotoAlbumsListCall {
	c := &PhotoAlbumsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.scope = scope
	return c
}

// C sets the optional parameter "c": A continuation token that allows
// pagination.
func (c *PhotoAlbumsListCall) C(c1 string) *PhotoAlbumsListCall {
	c.opt_["c"] = c1
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *PhotoAlbumsListCall) Hl(hl string) *PhotoAlbumsListCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *PhotoAlbumsListCall) Alt(alt string) *PhotoAlbumsListCall {
	c.opt_["alt"] = alt
	return c
}

// MaxResults sets the optional parameter "max-results": Maximum number
// of results to include.
func (c *PhotoAlbumsListCall) MaxResults(maxResults int64) *PhotoAlbumsListCall {
	c.opt_["max-results"] = maxResults
	return c
}

func (c *PhotoAlbumsListCall) Do() (*AlbumsFeed, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["c"]; ok {
		params.Set("c", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["max-results"]; ok {
		params.Set("max-results", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "photos/{userId}/{scope}")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{scope}", cleanPathString(c.scope), 1)
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
	ret := new(AlbumsFeed)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List a user's photo albums",
	//   "httpMethod": "GET",
	//   "id": "chili.photoAlbums.list",
	//   "parameterOrder": [
	//     "userId",
	//     "scope"
	//   ],
	//   "parameters": {
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "c": {
	//       "description": "A continuation token that allows pagination.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "max-results": {
	//       "default": "20",
	//       "description": "Maximum number of results to include.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "scope": {
	//       "description": "The collection of albums to list.",
	//       "enum": [
	//         "@self"
	//       ],
	//       "enumDescriptions": [
	//         "Albums posted by the user."
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "photos/{userId}/{scope}",
	//   "response": {
	//     "$ref": "AlbumsFeed"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/picasa"
	//   ]
	// }

}

// method id "chili.photoAlbums.insert":

type PhotoAlbumsInsertCall struct {
	s      *Service
	userId string
	album  *Album
	opt_   map[string]interface{}
}

// Insert: Create a photo album
func (r *PhotoAlbumsService) Insert(userId string, album *Album) *PhotoAlbumsInsertCall {
	c := &PhotoAlbumsInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.album = album
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *PhotoAlbumsInsertCall) Hl(hl string) *PhotoAlbumsInsertCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *PhotoAlbumsInsertCall) Alt(alt string) *PhotoAlbumsInsertCall {
	c.opt_["alt"] = alt
	return c
}

func (c *PhotoAlbumsInsertCall) Do() (*Album, os.Error) {
	var body io.Reader = nil
	body, err := googleapi.WithDataWrapper.JSONReader(c.album)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "photos/{userId}/@self")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
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
	ret := new(Album)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Create a photo album",
	//   "httpMethod": "POST",
	//   "id": "chili.photoAlbums.insert",
	//   "parameterOrder": [
	//     "userId"
	//   ],
	//   "parameters": {
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "photos/{userId}/@self",
	//   "request": {
	//     "$ref": "Album"
	//   },
	//   "response": {
	//     "$ref": "Album"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/picasa"
	//   ]
	// }

}

// method id "chili.photoAlbums.get":

type PhotoAlbumsGetCall struct {
	s       *Service
	userId  string
	albumId string
	opt_    map[string]interface{}
}

// Get: Get a photo album
func (r *PhotoAlbumsService) Get(userId string, albumId string) *PhotoAlbumsGetCall {
	c := &PhotoAlbumsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.albumId = albumId
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *PhotoAlbumsGetCall) Hl(hl string) *PhotoAlbumsGetCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *PhotoAlbumsGetCall) Alt(alt string) *PhotoAlbumsGetCall {
	c.opt_["alt"] = alt
	return c
}

func (c *PhotoAlbumsGetCall) Do() (*Album, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "photos/{userId}/@self/{albumId}")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{albumId}", cleanPathString(c.albumId), 1)
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
	ret := new(Album)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get a photo album",
	//   "httpMethod": "GET",
	//   "id": "chili.photoAlbums.get",
	//   "parameterOrder": [
	//     "userId",
	//     "albumId"
	//   ],
	//   "parameters": {
	//     "albumId": {
	//       "description": "ID of the album to get.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "photos/{userId}/@self/{albumId}",
	//   "response": {
	//     "$ref": "Album"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/picasa"
	//   ]
	// }

}

// method id "chili.photoAlbums.delete":

type PhotoAlbumsDeleteCall struct {
	s       *Service
	userId  string
	albumId string
	opt_    map[string]interface{}
}

// Delete: Delete a photo album
func (r *PhotoAlbumsService) Delete(userId string, albumId string) *PhotoAlbumsDeleteCall {
	c := &PhotoAlbumsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.albumId = albumId
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *PhotoAlbumsDeleteCall) Hl(hl string) *PhotoAlbumsDeleteCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *PhotoAlbumsDeleteCall) Alt(alt string) *PhotoAlbumsDeleteCall {
	c.opt_["alt"] = alt
	return c
}

func (c *PhotoAlbumsDeleteCall) Do() os.Error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "photos/{userId}/@self/{albumId}")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{albumId}", cleanPathString(c.albumId), 1)
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
	//   "description": "Delete a photo album",
	//   "httpMethod": "DELETE",
	//   "id": "chili.photoAlbums.delete",
	//   "parameterOrder": [
	//     "userId",
	//     "albumId"
	//   ],
	//   "parameters": {
	//     "albumId": {
	//       "description": "ID of the album to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "photos/{userId}/@self/{albumId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/picasa"
	//   ]
	// }

}

// method id "chili.people.search":

type PeopleSearchCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// Search: Search for people
func (r *PeopleService) Search() *PeopleSearchCall {
	c := &PeopleSearchCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// C sets the optional parameter "c": A continuation token that allows
// pagination.
func (c *PeopleSearchCall) C(c1 string) *PeopleSearchCall {
	c.opt_["c"] = c1
	return c
}

// Q sets the optional parameter "q": Full-text search query string.
func (c *PeopleSearchCall) Q(q string) *PeopleSearchCall {
	c.opt_["q"] = q
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *PeopleSearchCall) Hl(hl string) *PeopleSearchCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *PeopleSearchCall) Alt(alt string) *PeopleSearchCall {
	c.opt_["alt"] = alt
	return c
}

// MaxResults sets the optional parameter "max-results": Maximum number
// of results to include.
func (c *PeopleSearchCall) MaxResults(maxResults int64) *PeopleSearchCall {
	c.opt_["max-results"] = maxResults
	return c
}

func (c *PeopleSearchCall) Do() (*PeopleFeed, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["c"]; ok {
		params.Set("c", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["q"]; ok {
		params.Set("q", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["max-results"]; ok {
		params.Set("max-results", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "people/search")
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
	//   "description": "Search for people",
	//   "httpMethod": "GET",
	//   "id": "chili.people.search",
	//   "parameters": {
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "c": {
	//       "description": "A continuation token that allows pagination.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "max-results": {
	//       "default": "20",
	//       "description": "Maximum number of results to include.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "q": {
	//       "description": "Full-text search query string.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "people/search",
	//   "response": {
	//     "$ref": "PeopleFeed"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz",
	//     "https://www.googleapis.com/auth/buzz.readonly"
	//   ]
	// }

}

// method id "chili.people.list":

type PeopleListCall struct {
	s       *Service
	userId  string
	groupId string
	opt_    map[string]interface{}
}

// List: Get people in a group
func (r *PeopleService) List(userId string, groupId string) *PeopleListCall {
	c := &PeopleListCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.groupId = groupId
	return c
}

// C sets the optional parameter "c": A continuation token that allows
// pagination.
func (c *PeopleListCall) C(c1 string) *PeopleListCall {
	c.opt_["c"] = c1
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *PeopleListCall) Hl(hl string) *PeopleListCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *PeopleListCall) Alt(alt string) *PeopleListCall {
	c.opt_["alt"] = alt
	return c
}

// MaxResults sets the optional parameter "max-results": Maximum number
// of results to include.
func (c *PeopleListCall) MaxResults(maxResults int64) *PeopleListCall {
	c.opt_["max-results"] = maxResults
	return c
}

func (c *PeopleListCall) Do() (*PeopleFeed, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["c"]; ok {
		params.Set("c", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["max-results"]; ok {
		params.Set("max-results", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "people/{userId}/@groups/{groupId}")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{groupId}", cleanPathString(c.groupId), 1)
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
	//   "description": "Get people in a group",
	//   "httpMethod": "GET",
	//   "id": "chili.people.list",
	//   "parameterOrder": [
	//     "userId",
	//     "groupId"
	//   ],
	//   "parameters": {
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "c": {
	//       "description": "A continuation token that allows pagination.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "groupId": {
	//       "description": "ID of the group for which to list users.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "max-results": {
	//       "default": "20",
	//       "description": "Maximum number of results to include.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "userId": {
	//       "description": "ID of the user being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "people/{userId}/@groups/{groupId}",
	//   "response": {
	//     "$ref": "PeopleFeed"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz",
	//     "https://www.googleapis.com/auth/buzz.readonly"
	//   ]
	// }

}

// method id "chili.people.update":

type PeopleUpdateCall struct {
	s        *Service
	userId   string
	groupId  string
	personId string
	person   *Person
	opt_     map[string]interface{}
}

// Update: Add a person to a group
func (r *PeopleService) Update(userId string, groupId string, personId string, person *Person) *PeopleUpdateCall {
	c := &PeopleUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.groupId = groupId
	c.personId = personId
	c.person = person
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *PeopleUpdateCall) Hl(hl string) *PeopleUpdateCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *PeopleUpdateCall) Alt(alt string) *PeopleUpdateCall {
	c.opt_["alt"] = alt
	return c
}

func (c *PeopleUpdateCall) Do() (*Person, os.Error) {
	var body io.Reader = nil
	body, err := googleapi.WithDataWrapper.JSONReader(c.person)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "people/{userId}/@groups/{groupId}/{personId}")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{groupId}", cleanPathString(c.groupId), 1)
	urls = strings.Replace(urls, "{personId}", cleanPathString(c.personId), 1)
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
	ret := new(Person)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Add a person to a group",
	//   "httpMethod": "PUT",
	//   "id": "chili.people.update",
	//   "parameterOrder": [
	//     "userId",
	//     "groupId",
	//     "personId"
	//   ],
	//   "parameters": {
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "groupId": {
	//       "description": "ID of the group to which to add the person.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "personId": {
	//       "description": "ID of the person to add to the group.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the owner of the group.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "people/{userId}/@groups/{groupId}/{personId}",
	//   "request": {
	//     "$ref": "Person"
	//   },
	//   "response": {
	//     "$ref": "Person"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz"
	//   ]
	// }

}

// method id "chili.people.liked":

type PeopleLikedCall struct {
	s       *Service
	userId  string
	scope   string
	postId  string
	groupId string
	opt_    map[string]interface{}
}

// Liked: Get people who liked an activity
func (r *PeopleService) Liked(userId string, scope string, postId string, groupId string) *PeopleLikedCall {
	c := &PeopleLikedCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.scope = scope
	c.postId = postId
	c.groupId = groupId
	return c
}

// C sets the optional parameter "c": A continuation token that allows
// pagination.
func (c *PeopleLikedCall) C(c1 string) *PeopleLikedCall {
	c.opt_["c"] = c1
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *PeopleLikedCall) Hl(hl string) *PeopleLikedCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *PeopleLikedCall) Alt(alt string) *PeopleLikedCall {
	c.opt_["alt"] = alt
	return c
}

// MaxResults sets the optional parameter "max-results": Maximum number
// of results to include.
func (c *PeopleLikedCall) MaxResults(maxResults int64) *PeopleLikedCall {
	c.opt_["max-results"] = maxResults
	return c
}

func (c *PeopleLikedCall) Do() (*PeopleFeed, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["c"]; ok {
		params.Set("c", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["max-results"]; ok {
		params.Set("max-results", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "activities/{userId}/{scope}/{postId}/{groupId}")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{scope}", cleanPathString(c.scope), 1)
	urls = strings.Replace(urls, "{postId}", cleanPathString(c.postId), 1)
	urls = strings.Replace(urls, "{groupId}", cleanPathString(c.groupId), 1)
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
	//   "description": "Get people who liked an activity",
	//   "httpMethod": "GET",
	//   "id": "chili.people.liked",
	//   "parameterOrder": [
	//     "userId",
	//     "scope",
	//     "postId",
	//     "groupId"
	//   ],
	//   "parameters": {
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "c": {
	//       "description": "A continuation token that allows pagination.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "groupId": {
	//       "enum": [
	//         "@liked"
	//       ],
	//       "enumDescriptions": [
	//         "People who liked this activity."
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "max-results": {
	//       "default": "20",
	//       "description": "Maximum number of results to include.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "postId": {
	//       "description": "ID of the activity that was liked.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "scope": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities/{userId}/{scope}/{postId}/{groupId}",
	//   "response": {
	//     "$ref": "PeopleFeed"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz",
	//     "https://www.googleapis.com/auth/buzz.readonly"
	//   ]
	// }

}

// method id "chili.people.get":

type PeopleGetCall struct {
	s      *Service
	userId string
	opt_   map[string]interface{}
}

// Get: Get a user profile
func (r *PeopleService) Get(userId string) *PeopleGetCall {
	c := &PeopleGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *PeopleGetCall) Hl(hl string) *PeopleGetCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *PeopleGetCall) Alt(alt string) *PeopleGetCall {
	c.opt_["alt"] = alt
	return c
}

func (c *PeopleGetCall) Do() (*Person, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "people/{userId}/@self")
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
	//   "description": "Get a user profile",
	//   "httpMethod": "GET",
	//   "id": "chili.people.get",
	//   "parameterOrder": [
	//     "userId"
	//   ],
	//   "parameters": {
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "people/{userId}/@self",
	//   "response": {
	//     "$ref": "Person"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz",
	//     "https://www.googleapis.com/auth/buzz.readonly"
	//   ]
	// }

}

// method id "chili.people.delete":

type PeopleDeleteCall struct {
	s        *Service
	userId   string
	groupId  string
	personId string
	opt_     map[string]interface{}
}

// Delete: Remove a person from a group
func (r *PeopleService) Delete(userId string, groupId string, personId string) *PeopleDeleteCall {
	c := &PeopleDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.groupId = groupId
	c.personId = personId
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *PeopleDeleteCall) Hl(hl string) *PeopleDeleteCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *PeopleDeleteCall) Alt(alt string) *PeopleDeleteCall {
	c.opt_["alt"] = alt
	return c
}

func (c *PeopleDeleteCall) Do() os.Error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "people/{userId}/@groups/{groupId}/{personId}")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{groupId}", cleanPathString(c.groupId), 1)
	urls = strings.Replace(urls, "{personId}", cleanPathString(c.personId), 1)
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
	//   "description": "Remove a person from a group",
	//   "httpMethod": "DELETE",
	//   "id": "chili.people.delete",
	//   "parameterOrder": [
	//     "userId",
	//     "groupId",
	//     "personId"
	//   ],
	//   "parameters": {
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "groupId": {
	//       "description": "ID of the group from which to remove the person.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "personId": {
	//       "description": "ID of the person to remove from the group.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the owner of the group.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "people/{userId}/@groups/{groupId}/{personId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz"
	//   ]
	// }

}

// method id "chili.people.reshared":

type PeopleResharedCall struct {
	s       *Service
	userId  string
	scope   string
	postId  string
	groupId string
	opt_    map[string]interface{}
}

// Reshared: Get people who reshared an activity
func (r *PeopleService) Reshared(userId string, scope string, postId string, groupId string) *PeopleResharedCall {
	c := &PeopleResharedCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.scope = scope
	c.postId = postId
	c.groupId = groupId
	return c
}

// C sets the optional parameter "c": A continuation token that allows
// pagination.
func (c *PeopleResharedCall) C(c1 string) *PeopleResharedCall {
	c.opt_["c"] = c1
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *PeopleResharedCall) Hl(hl string) *PeopleResharedCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *PeopleResharedCall) Alt(alt string) *PeopleResharedCall {
	c.opt_["alt"] = alt
	return c
}

// MaxResults sets the optional parameter "max-results": Maximum number
// of results to include.
func (c *PeopleResharedCall) MaxResults(maxResults int64) *PeopleResharedCall {
	c.opt_["max-results"] = maxResults
	return c
}

func (c *PeopleResharedCall) Do() (*PeopleFeed, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["c"]; ok {
		params.Set("c", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["max-results"]; ok {
		params.Set("max-results", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "activities/{userId}/{scope}/{postId}/{groupId}")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{scope}", cleanPathString(c.scope), 1)
	urls = strings.Replace(urls, "{postId}", cleanPathString(c.postId), 1)
	urls = strings.Replace(urls, "{groupId}", cleanPathString(c.groupId), 1)
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
	//   "description": "Get people who reshared an activity",
	//   "httpMethod": "GET",
	//   "id": "chili.people.reshared",
	//   "parameterOrder": [
	//     "userId",
	//     "scope",
	//     "postId",
	//     "groupId"
	//   ],
	//   "parameters": {
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "c": {
	//       "description": "A continuation token that allows pagination.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "groupId": {
	//       "enum": [
	//         "@reshared"
	//       ],
	//       "enumDescriptions": [
	//         "People who reshared this activity."
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "max-results": {
	//       "default": "20",
	//       "description": "Maximum number of results to include.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "postId": {
	//       "description": "ID of the activity that was reshared.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "scope": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities/{userId}/{scope}/{postId}/{groupId}",
	//   "response": {
	//     "$ref": "PeopleFeed"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz",
	//     "https://www.googleapis.com/auth/buzz.readonly"
	//   ]
	// }

}

// method id "chili.people.patch":

type PeoplePatchCall struct {
	s        *Service
	userId   string
	groupId  string
	personId string
	person   *Person
	opt_     map[string]interface{}
}

// Patch: Add a person to a group. This method supports patch semantics.
func (r *PeopleService) Patch(userId string, groupId string, personId string, person *Person) *PeoplePatchCall {
	c := &PeoplePatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.groupId = groupId
	c.personId = personId
	c.person = person
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *PeoplePatchCall) Hl(hl string) *PeoplePatchCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *PeoplePatchCall) Alt(alt string) *PeoplePatchCall {
	c.opt_["alt"] = alt
	return c
}

func (c *PeoplePatchCall) Do() (*Person, os.Error) {
	var body io.Reader = nil
	body, err := googleapi.WithDataWrapper.JSONReader(c.person)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "people/{userId}/@groups/{groupId}/{personId}")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{groupId}", cleanPathString(c.groupId), 1)
	urls = strings.Replace(urls, "{personId}", cleanPathString(c.personId), 1)
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
	ret := new(Person)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Add a person to a group. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "chili.people.patch",
	//   "parameterOrder": [
	//     "userId",
	//     "groupId",
	//     "personId"
	//   ],
	//   "parameters": {
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "groupId": {
	//       "description": "ID of the group to which to add the person.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "personId": {
	//       "description": "ID of the person to add to the group.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the owner of the group.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "people/{userId}/@groups/{groupId}/{personId}",
	//   "request": {
	//     "$ref": "Person"
	//   },
	//   "response": {
	//     "$ref": "Person"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz"
	//   ]
	// }

}

// method id "chili.activities.count":

type ActivitiesCountCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// Count: Get a count of link shares
func (r *ActivitiesService) Count() *ActivitiesCountCall {
	c := &ActivitiesCountCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *ActivitiesCountCall) Hl(hl string) *ActivitiesCountCall {
	c.opt_["hl"] = hl
	return c
}

// Url sets the optional parameter "url": URLs for which to get share
// counts.
func (c *ActivitiesCountCall) Url(url string) *ActivitiesCountCall {
	c.opt_["url"] = url
	return c
}

func (c *ActivitiesCountCall) Do() (*CountFeed, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["url"]; ok {
		params.Set("url", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "activities/count")
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
	ret := new(CountFeed)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get a count of link shares",
	//   "httpMethod": "GET",
	//   "id": "chili.activities.count",
	//   "parameters": {
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "url": {
	//       "description": "URLs for which to get share counts.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities/count",
	//   "response": {
	//     "$ref": "CountFeed"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz",
	//     "https://www.googleapis.com/auth/buzz.readonly"
	//   ]
	// }

}

// method id "chili.activities.search":

type ActivitiesSearchCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// Search: Search for activities
func (r *ActivitiesService) Search() *ActivitiesSearchCall {
	c := &ActivitiesSearchCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// C sets the optional parameter "c": A continuation token that allows
// pagination.
func (c *ActivitiesSearchCall) C(c1 string) *ActivitiesSearchCall {
	c.opt_["c"] = c1
	return c
}

// Q sets the optional parameter "q": Full-text search query string.
func (c *ActivitiesSearchCall) Q(q string) *ActivitiesSearchCall {
	c.opt_["q"] = q
	return c
}

// Bbox sets the optional parameter "bbox": Bounding box to use in a
// geographic location query.
func (c *ActivitiesSearchCall) Bbox(bbox string) *ActivitiesSearchCall {
	c.opt_["bbox"] = bbox
	return c
}

// TruncateAtom sets the optional parameter "truncateAtom": Truncate the
// value of the atom:content element.
func (c *ActivitiesSearchCall) TruncateAtom(truncateAtom bool) *ActivitiesSearchCall {
	c.opt_["truncateAtom"] = truncateAtom
	return c
}

// Radius sets the optional parameter "radius": Radius to use in a
// geographic location query.
func (c *ActivitiesSearchCall) Radius(radius string) *ActivitiesSearchCall {
	c.opt_["radius"] = radius
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *ActivitiesSearchCall) Hl(hl string) *ActivitiesSearchCall {
	c.opt_["hl"] = hl
	return c
}

// Lat sets the optional parameter "lat": Latitude to use in a
// geographic location query.
func (c *ActivitiesSearchCall) Lat(lat string) *ActivitiesSearchCall {
	c.opt_["lat"] = lat
	return c
}

// Pid sets the optional parameter "pid": ID of a place to use in a
// geographic location query.
func (c *ActivitiesSearchCall) Pid(pid string) *ActivitiesSearchCall {
	c.opt_["pid"] = pid
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *ActivitiesSearchCall) Alt(alt string) *ActivitiesSearchCall {
	c.opt_["alt"] = alt
	return c
}

// MaxResults sets the optional parameter "max-results": Maximum number
// of results to include.
func (c *ActivitiesSearchCall) MaxResults(maxResults int64) *ActivitiesSearchCall {
	c.opt_["max-results"] = maxResults
	return c
}

// Lon sets the optional parameter "lon": Longitude to use in a
// geographic location query.
func (c *ActivitiesSearchCall) Lon(lon string) *ActivitiesSearchCall {
	c.opt_["lon"] = lon
	return c
}

func (c *ActivitiesSearchCall) Do() (*ActivityFeed, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["c"]; ok {
		params.Set("c", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["q"]; ok {
		params.Set("q", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["bbox"]; ok {
		params.Set("bbox", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["truncateAtom"]; ok {
		params.Set("truncateAtom", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["radius"]; ok {
		params.Set("radius", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["lat"]; ok {
		params.Set("lat", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pid"]; ok {
		params.Set("pid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["max-results"]; ok {
		params.Set("max-results", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["lon"]; ok {
		params.Set("lon", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "activities/search")
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
	//   "description": "Search for activities",
	//   "httpMethod": "GET",
	//   "id": "chili.activities.search",
	//   "parameters": {
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "bbox": {
	//       "description": "Bounding box to use in a geographic location query.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "c": {
	//       "description": "A continuation token that allows pagination.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "lat": {
	//       "description": "Latitude to use in a geographic location query.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "lon": {
	//       "description": "Longitude to use in a geographic location query.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "max-results": {
	//       "default": "20",
	//       "description": "Maximum number of results to include.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pid": {
	//       "description": "ID of a place to use in a geographic location query.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "q": {
	//       "description": "Full-text search query string.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "radius": {
	//       "description": "Radius to use in a geographic location query.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "truncateAtom": {
	//       "description": "Truncate the value of the atom:content element.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "activities/search",
	//   "response": {
	//     "$ref": "ActivityFeed"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz",
	//     "https://www.googleapis.com/auth/buzz.readonly"
	//   ]
	// }

}

// method id "chili.activities.list":

type ActivitiesListCall struct {
	s      *Service
	userId string
	scope  string
	opt_   map[string]interface{}
}

// List: List activities
func (r *ActivitiesService) List(userId string, scope string) *ActivitiesListCall {
	c := &ActivitiesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.scope = scope
	return c
}

// C sets the optional parameter "c": A continuation token that allows
// pagination.
func (c *ActivitiesListCall) C(c1 string) *ActivitiesListCall {
	c.opt_["c"] = c1
	return c
}

// TruncateAtom sets the optional parameter "truncateAtom": Truncate the
// value of the atom:content element.
func (c *ActivitiesListCall) TruncateAtom(truncateAtom bool) *ActivitiesListCall {
	c.opt_["truncateAtom"] = truncateAtom
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *ActivitiesListCall) Hl(hl string) *ActivitiesListCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *ActivitiesListCall) Alt(alt string) *ActivitiesListCall {
	c.opt_["alt"] = alt
	return c
}

// MaxComments sets the optional parameter "max-comments": Maximum
// number of comments to include.
func (c *ActivitiesListCall) MaxComments(maxComments int64) *ActivitiesListCall {
	c.opt_["max-comments"] = maxComments
	return c
}

// MaxLiked sets the optional parameter "max-liked": Maximum number of
// likes to include.
func (c *ActivitiesListCall) MaxLiked(maxLiked int64) *ActivitiesListCall {
	c.opt_["max-liked"] = maxLiked
	return c
}

// MaxResults sets the optional parameter "max-results": Maximum number
// of results to include.
func (c *ActivitiesListCall) MaxResults(maxResults int64) *ActivitiesListCall {
	c.opt_["max-results"] = maxResults
	return c
}

func (c *ActivitiesListCall) Do() (*ActivityFeed, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["c"]; ok {
		params.Set("c", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["truncateAtom"]; ok {
		params.Set("truncateAtom", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["max-comments"]; ok {
		params.Set("max-comments", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["max-liked"]; ok {
		params.Set("max-liked", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["max-results"]; ok {
		params.Set("max-results", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "activities/{userId}/{scope}")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{scope}", cleanPathString(c.scope), 1)
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
	//   "description": "List activities",
	//   "httpMethod": "GET",
	//   "id": "chili.activities.list",
	//   "parameterOrder": [
	//     "userId",
	//     "scope"
	//   ],
	//   "parameters": {
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "c": {
	//       "description": "A continuation token that allows pagination.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "max-comments": {
	//       "default": "0",
	//       "description": "Maximum number of comments to include.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "max-liked": {
	//       "default": "0",
	//       "description": "Maximum number of likes to include.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "max-results": {
	//       "default": "20",
	//       "description": "Maximum number of results to include.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "scope": {
	//       "description": "The collection of activities to list.",
	//       "enum": [
	//         "@comments",
	//         "@consumption",
	//         "@liked",
	//         "@public",
	//         "@self"
	//       ],
	//       "enumDescriptions": [
	//         "Limit to activities commented on by the user.",
	//         "Limit to activities to be consumed by the user.",
	//         "Limit to activities liked by the user.",
	//         "Limit to public activities posted by the user.",
	//         "Limit to activities posted by the user."
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "truncateAtom": {
	//       "description": "Truncate the value of the atom:content element.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "userId": {
	//       "description": "ID of the user being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities/{userId}/{scope}",
	//   "response": {
	//     "$ref": "ActivityFeed"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz",
	//     "https://www.googleapis.com/auth/buzz.readonly"
	//   ]
	// }

}

// method id "chili.activities.update":

type ActivitiesUpdateCall struct {
	s        *Service
	userId   string
	scope    string
	postId   string
	activity *Activity
	opt_     map[string]interface{}
}

// Update: Update an activity
func (r *ActivitiesService) Update(userId string, scope string, postId string, activity *Activity) *ActivitiesUpdateCall {
	c := &ActivitiesUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.scope = scope
	c.postId = postId
	c.activity = activity
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *ActivitiesUpdateCall) Hl(hl string) *ActivitiesUpdateCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *ActivitiesUpdateCall) Alt(alt string) *ActivitiesUpdateCall {
	c.opt_["alt"] = alt
	return c
}

// AbuseType sets the optional parameter "abuseType": 
func (c *ActivitiesUpdateCall) AbuseType(abuseType string) *ActivitiesUpdateCall {
	c.opt_["abuseType"] = abuseType
	return c
}

func (c *ActivitiesUpdateCall) Do() (*Activity, os.Error) {
	var body io.Reader = nil
	body, err := googleapi.WithDataWrapper.JSONReader(c.activity)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["abuseType"]; ok {
		params.Set("abuseType", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "activities/{userId}/{scope}/{postId}")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{scope}", cleanPathString(c.scope), 1)
	urls = strings.Replace(urls, "{postId}", cleanPathString(c.postId), 1)
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
	ret := new(Activity)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update an activity",
	//   "httpMethod": "PUT",
	//   "id": "chili.activities.update",
	//   "parameterOrder": [
	//     "userId",
	//     "scope",
	//     "postId"
	//   ],
	//   "parameters": {
	//     "abuseType": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "postId": {
	//       "description": "ID of the activity to update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "scope": {
	//       "description": "The collection to which the activity belongs.",
	//       "enum": [
	//         "@abuse",
	//         "@liked",
	//         "@muted",
	//         "@self"
	//       ],
	//       "enumDescriptions": [
	//         "Activities reported by the user.",
	//         "Activities liked by the user.",
	//         "Activities muted by the user.",
	//         "Activities posted by the user."
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user whose post to update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities/{userId}/{scope}/{postId}",
	//   "request": {
	//     "$ref": "Activity"
	//   },
	//   "response": {
	//     "$ref": "Activity"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz"
	//   ]
	// }

}

// method id "chili.activities.insert":

type ActivitiesInsertCall struct {
	s        *Service
	userId   string
	activity *Activity
	opt_     map[string]interface{}
	media_   io.Reader
}

// Insert: Create a new activity
func (r *ActivitiesService) Insert(userId string, activity *Activity) *ActivitiesInsertCall {
	c := &ActivitiesInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.activity = activity
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *ActivitiesInsertCall) Hl(hl string) *ActivitiesInsertCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *ActivitiesInsertCall) Alt(alt string) *ActivitiesInsertCall {
	c.opt_["alt"] = alt
	return c
}

// Preview sets the optional parameter "preview": If true, only preview
// the action.
func (c *ActivitiesInsertCall) Preview(preview bool) *ActivitiesInsertCall {
	c.opt_["preview"] = preview
	return c
}
func (c *ActivitiesInsertCall) Media(r io.Reader) *ActivitiesInsertCall {
	c.media_ = r
	return c
}

func (c *ActivitiesInsertCall) Do() (*Activity, os.Error) {
	var body io.Reader = nil
	body, err := googleapi.WithDataWrapper.JSONReader(c.activity)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["preview"]; ok {
		params.Set("preview", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "activities/{userId}/@self")
	if c.media_ != nil {
		urls = strings.Replace(urls, "https://www.googleapis.com/", "https://www.googleapis.com/upload/", 1)
	}
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls += "?" + params.Encode()
	contentLength_, hasMedia_ := googleapi.ConditionallyIncludeMedia(c.media_, &body, &ctype)
	req, _ := http.NewRequest("POST", urls, body)
	if hasMedia_ {
		req.ContentLength = contentLength_
	}
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
	//   "description": "Create a new activity",
	//   "httpMethod": "POST",
	//   "id": "chili.activities.insert",
	//   "mediaUpload": {
	//     "accept": [
	//       "image/*"
	//     ],
	//     "maxSize": "10MB",
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/buzz/v1/activities/{userId}/@self"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/buzz/v1/activities/{userId}/@self"
	//       }
	//     }
	//   },
	//   "parameterOrder": [
	//     "userId"
	//   ],
	//   "parameters": {
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "preview": {
	//       "default": "false",
	//       "description": "If true, only preview the action.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "userId": {
	//       "description": "ID of the user being referenced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities/{userId}/@self",
	//   "request": {
	//     "$ref": "Activity"
	//   },
	//   "response": {
	//     "$ref": "Activity"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz"
	//   ]
	// }

}

// method id "chili.activities.track":

type ActivitiesTrackCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// Track: Get real-time activity tracking information
func (r *ActivitiesService) Track() *ActivitiesTrackCall {
	c := &ActivitiesTrackCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// C sets the optional parameter "c": A continuation token that allows
// pagination.
func (c *ActivitiesTrackCall) C(c1 string) *ActivitiesTrackCall {
	c.opt_["c"] = c1
	return c
}

// Q sets the optional parameter "q": Full-text search query string.
func (c *ActivitiesTrackCall) Q(q string) *ActivitiesTrackCall {
	c.opt_["q"] = q
	return c
}

// Bbox sets the optional parameter "bbox": Bounding box to use in a
// geographic location query.
func (c *ActivitiesTrackCall) Bbox(bbox string) *ActivitiesTrackCall {
	c.opt_["bbox"] = bbox
	return c
}

// Radius sets the optional parameter "radius": Radius to use in a
// geographic location query.
func (c *ActivitiesTrackCall) Radius(radius string) *ActivitiesTrackCall {
	c.opt_["radius"] = radius
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *ActivitiesTrackCall) Hl(hl string) *ActivitiesTrackCall {
	c.opt_["hl"] = hl
	return c
}

// Lat sets the optional parameter "lat": Latitude to use in a
// geographic location query.
func (c *ActivitiesTrackCall) Lat(lat string) *ActivitiesTrackCall {
	c.opt_["lat"] = lat
	return c
}

// Pid sets the optional parameter "pid": ID of a place to use in a
// geographic location query.
func (c *ActivitiesTrackCall) Pid(pid string) *ActivitiesTrackCall {
	c.opt_["pid"] = pid
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *ActivitiesTrackCall) Alt(alt string) *ActivitiesTrackCall {
	c.opt_["alt"] = alt
	return c
}

// MaxResults sets the optional parameter "max-results": Maximum number
// of results to include.
func (c *ActivitiesTrackCall) MaxResults(maxResults int64) *ActivitiesTrackCall {
	c.opt_["max-results"] = maxResults
	return c
}

// Lon sets the optional parameter "lon": Longitude to use in a
// geographic location query.
func (c *ActivitiesTrackCall) Lon(lon string) *ActivitiesTrackCall {
	c.opt_["lon"] = lon
	return c
}

func (c *ActivitiesTrackCall) Do() (*ActivityFeed, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["c"]; ok {
		params.Set("c", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["q"]; ok {
		params.Set("q", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["bbox"]; ok {
		params.Set("bbox", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["radius"]; ok {
		params.Set("radius", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["lat"]; ok {
		params.Set("lat", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pid"]; ok {
		params.Set("pid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["max-results"]; ok {
		params.Set("max-results", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["lon"]; ok {
		params.Set("lon", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "activities/track")
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
	//   "description": "Get real-time activity tracking information",
	//   "httpMethod": "GET",
	//   "id": "chili.activities.track",
	//   "parameters": {
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "bbox": {
	//       "description": "Bounding box to use in a geographic location query.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "c": {
	//       "description": "A continuation token that allows pagination.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "lat": {
	//       "description": "Latitude to use in a geographic location query.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "lon": {
	//       "description": "Longitude to use in a geographic location query.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "max-results": {
	//       "default": "20",
	//       "description": "Maximum number of results to include.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pid": {
	//       "description": "ID of a place to use in a geographic location query.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "q": {
	//       "description": "Full-text search query string.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "radius": {
	//       "description": "Radius to use in a geographic location query.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities/track",
	//   "response": {
	//     "$ref": "ActivityFeed"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz",
	//     "https://www.googleapis.com/auth/buzz.readonly"
	//   ]
	// }

}

// method id "chili.activities.get":

type ActivitiesGetCall struct {
	s      *Service
	userId string
	postId string
	opt_   map[string]interface{}
}

// Get: Get an activity
func (r *ActivitiesService) Get(userId string, postId string) *ActivitiesGetCall {
	c := &ActivitiesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.postId = postId
	return c
}

// TruncateAtom sets the optional parameter "truncateAtom": Truncate the
// value of the atom:content element.
func (c *ActivitiesGetCall) TruncateAtom(truncateAtom bool) *ActivitiesGetCall {
	c.opt_["truncateAtom"] = truncateAtom
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *ActivitiesGetCall) Hl(hl string) *ActivitiesGetCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *ActivitiesGetCall) Alt(alt string) *ActivitiesGetCall {
	c.opt_["alt"] = alt
	return c
}

// MaxComments sets the optional parameter "max-comments": Maximum
// number of comments to include.
func (c *ActivitiesGetCall) MaxComments(maxComments int64) *ActivitiesGetCall {
	c.opt_["max-comments"] = maxComments
	return c
}

// MaxLiked sets the optional parameter "max-liked": Maximum number of
// likes to include.
func (c *ActivitiesGetCall) MaxLiked(maxLiked int64) *ActivitiesGetCall {
	c.opt_["max-liked"] = maxLiked
	return c
}

func (c *ActivitiesGetCall) Do() (*Activity, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["truncateAtom"]; ok {
		params.Set("truncateAtom", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["max-comments"]; ok {
		params.Set("max-comments", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["max-liked"]; ok {
		params.Set("max-liked", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "activities/{userId}/@self/{postId}")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{postId}", cleanPathString(c.postId), 1)
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
	//   "description": "Get an activity",
	//   "httpMethod": "GET",
	//   "id": "chili.activities.get",
	//   "parameterOrder": [
	//     "userId",
	//     "postId"
	//   ],
	//   "parameters": {
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "max-comments": {
	//       "default": "0",
	//       "description": "Maximum number of comments to include.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "max-liked": {
	//       "default": "0",
	//       "description": "Maximum number of likes to include.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "postId": {
	//       "description": "ID of the post to get.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "truncateAtom": {
	//       "description": "Truncate the value of the atom:content element.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "userId": {
	//       "description": "ID of the user whose post to get.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities/{userId}/@self/{postId}",
	//   "response": {
	//     "$ref": "Activity"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz",
	//     "https://www.googleapis.com/auth/buzz.readonly"
	//   ]
	// }

}

// method id "chili.activities.delete":

type ActivitiesDeleteCall struct {
	s      *Service
	userId string
	scope  string
	postId string
	opt_   map[string]interface{}
}

// Delete: Delete an activity
func (r *ActivitiesService) Delete(userId string, scope string, postId string) *ActivitiesDeleteCall {
	c := &ActivitiesDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.scope = scope
	c.postId = postId
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *ActivitiesDeleteCall) Hl(hl string) *ActivitiesDeleteCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *ActivitiesDeleteCall) Alt(alt string) *ActivitiesDeleteCall {
	c.opt_["alt"] = alt
	return c
}

func (c *ActivitiesDeleteCall) Do() os.Error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "activities/{userId}/{scope}/{postId}")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{scope}", cleanPathString(c.scope), 1)
	urls = strings.Replace(urls, "{postId}", cleanPathString(c.postId), 1)
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
	//   "description": "Delete an activity",
	//   "httpMethod": "DELETE",
	//   "id": "chili.activities.delete",
	//   "parameterOrder": [
	//     "userId",
	//     "scope",
	//     "postId"
	//   ],
	//   "parameters": {
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "postId": {
	//       "description": "ID of the activity to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "scope": {
	//       "description": "The collection to which the activity belongs.",
	//       "enum": [
	//         "@liked",
	//         "@muted",
	//         "@self"
	//       ],
	//       "enumDescriptions": [
	//         "Activities liked by the user.",
	//         "Activities muted by the user.",
	//         "Activities posted by the user."
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user whose post to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities/{userId}/{scope}/{postId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz"
	//   ]
	// }

}

// method id "chili.activities.extractPeopleFromSearch":

type ActivitiesExtractPeopleFromSearchCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// ExtractPeopleFromSearch: Search for people by topic
func (r *ActivitiesService) ExtractPeopleFromSearch() *ActivitiesExtractPeopleFromSearchCall {
	c := &ActivitiesExtractPeopleFromSearchCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// C sets the optional parameter "c": A continuation token that allows
// pagination.
func (c *ActivitiesExtractPeopleFromSearchCall) C(c1 string) *ActivitiesExtractPeopleFromSearchCall {
	c.opt_["c"] = c1
	return c
}

// Q sets the optional parameter "q": Full-text search query string.
func (c *ActivitiesExtractPeopleFromSearchCall) Q(q string) *ActivitiesExtractPeopleFromSearchCall {
	c.opt_["q"] = q
	return c
}

// Bbox sets the optional parameter "bbox": Bounding box to use in a
// geographic location query.
func (c *ActivitiesExtractPeopleFromSearchCall) Bbox(bbox string) *ActivitiesExtractPeopleFromSearchCall {
	c.opt_["bbox"] = bbox
	return c
}

// Radius sets the optional parameter "radius": Radius to use in a
// geographic location query.
func (c *ActivitiesExtractPeopleFromSearchCall) Radius(radius string) *ActivitiesExtractPeopleFromSearchCall {
	c.opt_["radius"] = radius
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *ActivitiesExtractPeopleFromSearchCall) Hl(hl string) *ActivitiesExtractPeopleFromSearchCall {
	c.opt_["hl"] = hl
	return c
}

// Lat sets the optional parameter "lat": Latitude to use in a
// geographic location query.
func (c *ActivitiesExtractPeopleFromSearchCall) Lat(lat string) *ActivitiesExtractPeopleFromSearchCall {
	c.opt_["lat"] = lat
	return c
}

// Pid sets the optional parameter "pid": ID of a place to use in a
// geographic location query.
func (c *ActivitiesExtractPeopleFromSearchCall) Pid(pid string) *ActivitiesExtractPeopleFromSearchCall {
	c.opt_["pid"] = pid
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *ActivitiesExtractPeopleFromSearchCall) Alt(alt string) *ActivitiesExtractPeopleFromSearchCall {
	c.opt_["alt"] = alt
	return c
}

// MaxResults sets the optional parameter "max-results": Maximum number
// of results to include.
func (c *ActivitiesExtractPeopleFromSearchCall) MaxResults(maxResults int64) *ActivitiesExtractPeopleFromSearchCall {
	c.opt_["max-results"] = maxResults
	return c
}

// Lon sets the optional parameter "lon": Longitude to use in a
// geographic location query.
func (c *ActivitiesExtractPeopleFromSearchCall) Lon(lon string) *ActivitiesExtractPeopleFromSearchCall {
	c.opt_["lon"] = lon
	return c
}

func (c *ActivitiesExtractPeopleFromSearchCall) Do() (*PeopleFeed, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["c"]; ok {
		params.Set("c", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["q"]; ok {
		params.Set("q", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["bbox"]; ok {
		params.Set("bbox", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["radius"]; ok {
		params.Set("radius", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["lat"]; ok {
		params.Set("lat", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pid"]; ok {
		params.Set("pid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["max-results"]; ok {
		params.Set("max-results", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["lon"]; ok {
		params.Set("lon", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "activities/search/@people")
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
	//   "description": "Search for people by topic",
	//   "httpMethod": "GET",
	//   "id": "chili.activities.extractPeopleFromSearch",
	//   "parameters": {
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "bbox": {
	//       "description": "Bounding box to use in a geographic location query.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "c": {
	//       "description": "A continuation token that allows pagination.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "lat": {
	//       "description": "Latitude to use in a geographic location query.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "lon": {
	//       "description": "Longitude to use in a geographic location query.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "max-results": {
	//       "default": "20",
	//       "description": "Maximum number of results to include.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pid": {
	//       "description": "ID of a place to use in a geographic location query.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "q": {
	//       "description": "Full-text search query string.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "radius": {
	//       "description": "Radius to use in a geographic location query.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities/search/@people",
	//   "response": {
	//     "$ref": "PeopleFeed"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz",
	//     "https://www.googleapis.com/auth/buzz.readonly"
	//   ]
	// }

}

// method id "chili.activities.patch":

type ActivitiesPatchCall struct {
	s        *Service
	userId   string
	scope    string
	postId   string
	activity *Activity
	opt_     map[string]interface{}
}

// Patch: Update an activity. This method supports patch semantics.
func (r *ActivitiesService) Patch(userId string, scope string, postId string, activity *Activity) *ActivitiesPatchCall {
	c := &ActivitiesPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.scope = scope
	c.postId = postId
	c.activity = activity
	return c
}

// Hl sets the optional parameter "hl": Language code to limit language
// results.
func (c *ActivitiesPatchCall) Hl(hl string) *ActivitiesPatchCall {
	c.opt_["hl"] = hl
	return c
}

// Alt sets the optional parameter "alt": Specifies an alternative
// representation type.
func (c *ActivitiesPatchCall) Alt(alt string) *ActivitiesPatchCall {
	c.opt_["alt"] = alt
	return c
}

// AbuseType sets the optional parameter "abuseType": 
func (c *ActivitiesPatchCall) AbuseType(abuseType string) *ActivitiesPatchCall {
	c.opt_["abuseType"] = abuseType
	return c
}

func (c *ActivitiesPatchCall) Do() (*Activity, os.Error) {
	var body io.Reader = nil
	body, err := googleapi.WithDataWrapper.JSONReader(c.activity)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["alt"]; ok {
		params.Set("alt", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["abuseType"]; ok {
		params.Set("abuseType", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/buzz/v1/", "activities/{userId}/{scope}/{postId}")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{scope}", cleanPathString(c.scope), 1)
	urls = strings.Replace(urls, "{postId}", cleanPathString(c.postId), 1)
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
	ret := new(Activity)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update an activity. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "chili.activities.patch",
	//   "parameterOrder": [
	//     "userId",
	//     "scope",
	//     "postId"
	//   ],
	//   "parameters": {
	//     "abuseType": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "alt": {
	//       "default": "atom",
	//       "description": "Specifies an alternative representation type.",
	//       "enum": [
	//         "atom",
	//         "json"
	//       ],
	//       "enumDescriptions": [
	//         "Use Atom XML format",
	//         "Use JSON format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Language code to limit language results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "postId": {
	//       "description": "ID of the activity to update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "scope": {
	//       "description": "The collection to which the activity belongs.",
	//       "enum": [
	//         "@abuse",
	//         "@liked",
	//         "@muted",
	//         "@self"
	//       ],
	//       "enumDescriptions": [
	//         "Activities reported by the user.",
	//         "Activities liked by the user.",
	//         "Activities muted by the user.",
	//         "Activities posted by the user."
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user whose post to update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities/{userId}/{scope}/{postId}",
	//   "request": {
	//     "$ref": "Activity"
	//   },
	//   "response": {
	//     "$ref": "Activity"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/buzz"
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
