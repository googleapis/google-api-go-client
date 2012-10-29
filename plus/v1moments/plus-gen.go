// Package plus provides access to the Google+ API.
//
// See https://developers.google.com/+/history/
//
// Usage example:
//
//   import "code.google.com/p/google-api-go-client/plus/v1moments"
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

const apiId = "plus:v1moments"
const apiName = "plus"
const apiVersion = "v1moments"
const basePath = "https://www.googleapis.com/plus/v1moments/people/"

// OAuth2 scopes used by this API.
const (
	// Know who you are on Google
	PlusMeScope = "https://www.googleapis.com/auth/plus.me"

	// Send your activity to your private Google+ history
	PlusMomentsWriteScope = "https://www.googleapis.com/auth/plus.moments.write"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Moments = &MomentsService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	Moments *MomentsService
}

type MomentsService struct {
	s *Service
}

type ItemScope struct {
	// About: The subject matter of the content.
	About *ItemScope `json:"about,omitempty"`

	// AdditionalName: An additional name for a Person, can be used for a
	// middle name.
	AdditionalName []string `json:"additionalName,omitempty"`

	// Address: Postal address.
	Address *ItemScope `json:"address,omitempty"`

	// AddressCountry: Address country.
	AddressCountry string `json:"addressCountry,omitempty"`

	// AddressLocality: Address locality.
	AddressLocality string `json:"addressLocality,omitempty"`

	// AddressRegion: Address region.
	AddressRegion string `json:"addressRegion,omitempty"`

	// Associated_media: The encoding.
	Associated_media []*ItemScope `json:"associated_media,omitempty"`

	// AttendeeCount: Number of attendees.
	AttendeeCount int64 `json:"attendeeCount,omitempty"`

	// Attendees: A person attending the event.
	Attendees []*ItemScope `json:"attendees,omitempty"`

	// Audio: From http://schema.org/MusicRecording, the audio file.
	Audio *ItemScope `json:"audio,omitempty"`

	// Author: The person who created this scope.
	Author []*ItemScope `json:"author,omitempty"`

	// BestRating: Best possible rating value.
	BestRating string `json:"bestRating,omitempty"`

	// BirthDate: Date of birth.
	BirthDate string `json:"birthDate,omitempty"`

	// ByArtist: From http://schema.org/MusicRecording, the artist that
	// performed this recording.
	ByArtist *ItemScope `json:"byArtist,omitempty"`

	// Caption: The caption for this object.
	Caption string `json:"caption,omitempty"`

	// ContentSize: File size in (mega/kilo) bytes.
	ContentSize string `json:"contentSize,omitempty"`

	// ContentUrl: Actual bytes of the media object, for example the image
	// file or video file.
	ContentUrl string `json:"contentUrl,omitempty"`

	// Contributor: The list of contributors for this scope.
	Contributor []*ItemScope `json:"contributor,omitempty"`

	// DateCreated: The date this scope was created.
	DateCreated string `json:"dateCreated,omitempty"`

	// DateModified: The date this scope was last modified.
	DateModified string `json:"dateModified,omitempty"`

	// DatePublished: The initial date this scope was published.
	DatePublished string `json:"datePublished,omitempty"`

	// Description: The string describing the content of this scope.
	Description string `json:"description,omitempty"`

	// Duration: The duration of the item (movie, audio recording, event,
	// etc.) in ISO 8601 date format.
	Duration string `json:"duration,omitempty"`

	// EmbedUrl: A URL pointing to a player for a specific video. In
	// general, this is the information in the src element of an embed tag
	// and should not be the same as the content of the loc tag.
	EmbedUrl string `json:"embedUrl,omitempty"`

	// EndDate: The end date and time of the event (in ISO 8601 date
	// format).
	EndDate string `json:"endDate,omitempty"`

	// FamilyName: Family name. In the U.S., the last name of an Person.
	// This can be used along with givenName instead of the Name property.
	FamilyName string `json:"familyName,omitempty"`

	// Gender: Gender of the person.
	Gender string `json:"gender,omitempty"`

	// Geo: Geo coordinates.
	Geo *ItemScope `json:"geo,omitempty"`

	// GivenName: Given name. In the U.S., the first name of a Person. This
	// can be used along with familyName instead of the Name property.
	GivenName string `json:"givenName,omitempty"`

	// Height: The height of the media object.
	Height string `json:"height,omitempty"`

	// Id: The id for this item scope.
	Id string `json:"id,omitempty"`

	// Image: A url to the image for this scope.
	Image string `json:"image,omitempty"`

	// InAlbum: From http://schema.org/MusicRecording, which album a song is
	// in.
	InAlbum *ItemScope `json:"inAlbum,omitempty"`

	// Kind: Identifies this resource as an itemScope.
	Kind string `json:"kind,omitempty"`

	// Latitude: Latitude.
	Latitude float64 `json:"latitude,omitempty"`

	// Location: The location of the event or organization.
	Location *ItemScope `json:"location,omitempty"`

	// Longitude: Longitude.
	Longitude float64 `json:"longitude,omitempty"`

	// Name: The name of this scope.
	Name string `json:"name,omitempty"`

	// PartOfTVSeries: Property of http://schema.org/TVEpisode indicating
	// which series the episode belongs to.
	PartOfTVSeries *ItemScope `json:"partOfTVSeries,omitempty"`

	// Performers: The main performer or performers of the event-for
	// example, a presenter, musician, or actor.
	Performers []*ItemScope `json:"performers,omitempty"`

	// PlayerType: Player type required-for example, Flash or Silverlight.
	PlayerType string `json:"playerType,omitempty"`

	// PostOfficeBoxNumber: Post office box number.
	PostOfficeBoxNumber string `json:"postOfficeBoxNumber,omitempty"`

	// PostalCode: Postal code.
	PostalCode string `json:"postalCode,omitempty"`

	// RatingValue: Rating value.
	RatingValue string `json:"ratingValue,omitempty"`

	// ReviewRating: Review rating.
	ReviewRating *ItemScope `json:"reviewRating,omitempty"`

	// StartDate: The start date and time of the event (in ISO 8601 date
	// format).
	StartDate string `json:"startDate,omitempty"`

	// StreetAddress: Street address.
	StreetAddress string `json:"streetAddress,omitempty"`

	// Text: Comment text, review text, etc.
	Text string `json:"text,omitempty"`

	// Thumbnail: Thumbnail image for an image or video.
	Thumbnail *ItemScope `json:"thumbnail,omitempty"`

	// ThumbnailUrl: A url to a thumbnail image for this scope.
	ThumbnailUrl string `json:"thumbnailUrl,omitempty"`

	// TickerSymbol: The exchange traded instrument associated with a
	// Corporation object. The tickerSymbol is expressed as an exchange and
	// an instrument name separated by a space character. For the exchange
	// component of the tickerSymbol attribute, we reccommend using the
	// controlled vocaulary of Market Identifier Codes (MIC) specified in
	// ISO15022.
	TickerSymbol string `json:"tickerSymbol,omitempty"`

	// Type: The item type.
	Type string `json:"type,omitempty"`

	// Url: A URL for the item upon which the action was performed.
	Url string `json:"url,omitempty"`

	// Width: The width of the media object.
	Width string `json:"width,omitempty"`

	// WorstRating: Worst possible rating value.
	WorstRating string `json:"worstRating,omitempty"`
}

type Moment struct {
	// Id: The moment ID.
	Id string `json:"id,omitempty"`

	// Kind: Identifies this resource as a moment.
	Kind string `json:"kind,omitempty"`

	// Result: The object generated by performing the action on the item
	Result *ItemScope `json:"result,omitempty"`

	// StartDate: Time stamp of when the action occurred in RFC3339 format.
	StartDate string `json:"startDate,omitempty"`

	// Target: The object on which the action was performed.
	Target *ItemScope `json:"target,omitempty"`

	// Type: The schema.org activity type.
	Type string `json:"type,omitempty"`
}

// method id "plus.moments.insert":

type MomentsInsertCall struct {
	s          *Service
	userId     string
	collection string
	moment     *Moment
	opt_       map[string]interface{}
}

// Insert: Record a user activity (e.g Bill watched a video on Youtube)
func (r *MomentsService) Insert(userId string, collection string, moment *Moment) *MomentsInsertCall {
	c := &MomentsInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.collection = collection
	c.moment = moment
	return c
}

// Debug sets the optional parameter "debug": Return the moment as
// written. Should be used only for debugging.
func (c *MomentsInsertCall) Debug(debug bool) *MomentsInsertCall {
	c.opt_["debug"] = debug
	return c
}

func (c *MomentsInsertCall) Do() (*Moment, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.moment)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["debug"]; ok {
		params.Set("debug", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/plus/v1moments/people/", "{userId}/moments/{collection}")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{collection}", cleanPathString(c.collection), 1)
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
	ret := new(Moment)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Record a user activity (e.g Bill watched a video on Youtube)",
	//   "httpMethod": "POST",
	//   "id": "plus.moments.insert",
	//   "parameterOrder": [
	//     "userId",
	//     "collection"
	//   ],
	//   "parameters": {
	//     "collection": {
	//       "description": "The collection to which to write moments.",
	//       "enum": [
	//         "vault"
	//       ],
	//       "enumDescriptions": [
	//         "The default collection for writing new moments."
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "debug": {
	//       "description": "Return the moment as written. Should be used only for debugging.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "userId": {
	//       "description": "The ID of the user to record activities for. The only valid values are \"me\" and the ID of the authenticated user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{userId}/moments/{collection}",
	//   "request": {
	//     "$ref": "Moment"
	//   },
	//   "response": {
	//     "$ref": "Moment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/plus.me",
	//     "https://www.googleapis.com/auth/plus.moments.write"
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
