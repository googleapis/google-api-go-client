// Package youtube provides access to the YouTube API.
//
// See https://developers.google.com/youtube
//
// Usage example:
//
//   import "code.google.com/p/google-api-go-client/youtube/v3alpha"
//   ...
//   youtubeService, err := youtube.New(oauthHttpClient)
package youtube

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

const apiId = "youtube:v3alpha"
const apiName = "youtube"
const apiVersion = "v3alpha"
const basePath = "https://www.googleapis.com/youtube/v3alpha/"

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Channels = &ChannelsService{s: s}
	s.Playlistitems = &PlaylistitemsService{s: s}
	s.Playlists = &PlaylistsService{s: s}
	s.Search = &SearchService{s: s}
	s.Videos = &VideosService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	Channels *ChannelsService

	Playlistitems *PlaylistitemsService

	Playlists *PlaylistsService

	Search *SearchService

	Videos *VideosService
}

type ChannelsService struct {
	s *Service
}

type PlaylistitemsService struct {
	s *Service
}

type PlaylistsService struct {
	s *Service
}

type SearchService struct {
	s *Service
}

type VideosService struct {
	s *Service
}

type Channel struct {
	// ContentDetails: Information about the channel content: upload
	// playlist id, privacy status.
	ContentDetails *ChannelContentDetails `json:"contentDetails,omitempty"`

	// Etag: The eTag of the channel.
	Etag string `json:"etag,omitempty"`

	// Id: The unique ID of the channel.
	Id string `json:"id,omitempty"`

	// Kind: The type of this API resource.
	Kind string `json:"kind,omitempty"`

	// Snippet: Basic details about the channel: title, description, and
	// thumbnails.
	Snippet *ChannelSnippet `json:"snippet,omitempty"`

	// Statistics: Statistics about the channel: number of subscribers,
	// views, and comments.
	Statistics *ChannelStatistics `json:"statistics,omitempty"`
}

type ChannelContentDetails struct {
	// PrivacyStatus: Privacy status of the channel.
	PrivacyStatus string `json:"privacyStatus,omitempty"`

	// Uploads: The ID of the playlist containing the uploads of this
	// channel.
	Uploads string `json:"uploads,omitempty"`
}

type ChannelListResponse struct {
	// Channels: Map of channels matching the request criteria, keyed by
	// channel id.
	Channels *ChannelListResponseChannels `json:"channels,omitempty"`

	// Etag: The eTag of the response.
	Etag string `json:"etag,omitempty"`

	// Kind: The type of this API response.
	Kind string `json:"kind,omitempty"`
}

type ChannelListResponseChannels struct {
}

type ChannelSnippet struct {
	// Description: Description of the channel.
	Description string `json:"description,omitempty"`

	// Thumbnails: Channel thumbnails.
	Thumbnails *ChannelSnippetThumbnails `json:"thumbnails,omitempty"`

	// Title: Title of the channel.
	Title string `json:"title,omitempty"`
}

type ChannelSnippetThumbnails struct {
}

type ChannelStatistics struct {
	// CommentCount: Number of comments for this channel.
	CommentCount uint64 `json:"commentCount,omitempty,string"`

	// SubscriberCount: Number of subscribers to this channel.
	SubscriberCount uint64 `json:"subscriberCount,omitempty,string"`

	// VideoCount: Number of videos in the channel.
	VideoCount uint64 `json:"videoCount,omitempty,string"`

	// ViewCount: Number of times the channel has been viewed.
	ViewCount uint64 `json:"viewCount,omitempty,string"`
}

type PageInfo struct {
	// ResultPerPage: The number of results to display for each page.
	ResultPerPage int64 `json:"resultPerPage,omitempty"`

	// StartIndex: The index position of the first result to display.
	StartIndex int64 `json:"startIndex,omitempty"`

	// TotalResults: The total number of results.
	TotalResults int64 `json:"totalResults,omitempty"`
}

type Playlist struct {
	// Etag: The eTag of the playlist.
	Etag string `json:"etag,omitempty"`

	// Id: The unique id of the playlist.
	Id string `json:"id,omitempty"`

	// Kind: The type of this API resource.
	Kind string `json:"kind,omitempty"`

	// Snippet: Basic details about the playlist: title, description,
	// thumbnails.
	Snippet *PlaylistSnippet `json:"snippet,omitempty"`
}

type PlaylistItem struct {
	// Etag: The eTag of the playlist item.
	Etag string `json:"etag,omitempty"`

	// Id: The unique id of the playlist item.
	Id string `json:"id,omitempty"`

	// Kind: The type of this API resource.
	Kind string `json:"kind,omitempty"`

	// Snippet: Basic details about the playlist item: title, description,
	// thumbnails.
	Snippet *PlaylistItemSnippet `json:"snippet,omitempty"`
}

type PlaylistItemListResponse struct {
	// Etag: The eTag of the response.
	Etag string `json:"etag,omitempty"`

	// Kind: The type of this API response.
	Kind string `json:"kind,omitempty"`

	// PlaylistItems: Map of playlist items matching the request criteria,
	// keyed by id.
	PlaylistItems *PlaylistItemListResponsePlaylistItems `json:"playlistItems,omitempty"`
}

type PlaylistItemListResponsePlaylistItems struct {
}

type PlaylistItemSnippet struct {
	// ChannelId: Author of the playlist item.
	ChannelId string `json:"channelId,omitempty"`

	// Description: Description of the playlist item.
	Description string `json:"description,omitempty"`

	// PlaylistId: The playlist the item is part of.
	PlaylistId string `json:"playlistId,omitempty"`

	// Position: The position of the item within the playlist.
	Position int64 `json:"position,omitempty"`

	// PublishedAt: The date and time the playlist item was created.
	PublishedAt string `json:"publishedAt,omitempty"`

	// ResourceId: The ID of the resource referenced by the playlist item.
	ResourceId *ResourceId `json:"resourceId,omitempty"`

	// Title: Title of the playlist item.
	Title string `json:"title,omitempty"`
}

type PlaylistListResponse struct {
	// Etag: The eTag of the response.
	Etag string `json:"etag,omitempty"`

	// Kind: The type of this API response.
	Kind string `json:"kind,omitempty"`

	// Playlists: Map of playlists matching the request criteria, keyed by
	// id.
	Playlists *PlaylistListResponsePlaylists `json:"playlists,omitempty"`
}

type PlaylistListResponsePlaylists struct {
}

type PlaylistSnippet struct {
	// ChannelId: Author of the playlist.
	ChannelId string `json:"channelId,omitempty"`

	// Description: Description of the playlist.
	Description string `json:"description,omitempty"`

	// PublishedAt: The date and time the playlist was created.
	PublishedAt string `json:"publishedAt,omitempty"`

	// Tags: Textual tags associated with the playlist.
	Tags []string `json:"tags,omitempty"`

	// Title: Title of the playlist.
	Title string `json:"title,omitempty"`
}

type ResourceId struct {
	// ChannelId: ID of the referred channel. Present only when type is
	// "CHANNEL".
	ChannelId string `json:"channelId,omitempty"`

	// Kind: The kind of the referred resource.
	Kind string `json:"kind,omitempty"`

	// PlaylistId: ID of the referred playlist. Present only when type is
	// "PLAYLIST".
	PlaylistId string `json:"playlistId,omitempty"`

	// VideoId: ID of the referred video. Present only when type is "VIDEO".
	VideoId string `json:"videoId,omitempty"`
}

type SearchListResponse struct {
	// Etag: The eTag of the response.
	Etag string `json:"etag,omitempty"`

	// Kind: The type of this API response.
	Kind string `json:"kind,omitempty"`

	// PageInfo: Paging information for the search result.
	PageInfo *PageInfo `json:"pageInfo,omitempty"`

	// SearchResults: List of results matching the request criteria.
	SearchResults []*SearchResult `json:"searchResults,omitempty"`
}

type SearchResult struct {
	// Etag: The eTag of the search result.
	Etag string `json:"etag,omitempty"`

	// Id: The id of the resource.
	Id *ResourceId `json:"id,omitempty"`

	// Kind: The type of this API resource.
	Kind string `json:"kind,omitempty"`

	// Snippet: Basic details about the search result: title, description,
	// author.
	Snippet *SearchResultSnippet `json:"snippet,omitempty"`
}

type SearchResultSnippet struct {
	// ChannelId: Author of the found resource.
	ChannelId string `json:"channelId,omitempty"`

	// Description: Description of the search result.
	Description string `json:"description,omitempty"`

	// PublishedAt: The date and time the found resource was created.
	PublishedAt string `json:"publishedAt,omitempty"`

	// Title: Title of the search result.
	Title string `json:"title,omitempty"`
}

type Thumbnail struct {
	// Url: The URL for the thumbnail.
	Url string `json:"url,omitempty"`
}

type Video struct {
	// ContentDetails: Information about the video content, media file.
	ContentDetails *VideoContentDetails `json:"contentDetails,omitempty"`

	// Etag: The eTag of the video.
	Etag string `json:"etag,omitempty"`

	// Id: The unique id of the video.
	Id string `json:"id,omitempty"`

	// Kind: The type of this API resource.
	Kind string `json:"kind,omitempty"`

	// Player: Information used to play the video.
	Player *VideoPlayer `json:"player,omitempty"`

	// Snippet: Basic details about the video: title, description,
	// thumbnails.
	Snippet *VideoSnippet `json:"snippet,omitempty"`

	// Statistics: Statistics about the video: number of views, ratings.
	Statistics *VideoStatistics `json:"statistics,omitempty"`

	// Status: Status of the video upload, privacy status.
	Status *VideoStatus `json:"status,omitempty"`
}

type VideoContentDetails struct {
	// AspectRatio: The aspect ratio of the video.
	AspectRatio string `json:"aspectRatio,omitempty"`

	// Duration: Duration of the video.
	Duration uint64 `json:"duration,omitempty,string"`
}

type VideoListResponse struct {
	// Etag: The eTag of the response.
	Etag string `json:"etag,omitempty"`

	// Kind: The type of this API response.
	Kind string `json:"kind,omitempty"`

	// Videos: Map of videos matching the request criteria, keyed by video
	// id.
	Videos *VideoListResponseVideos `json:"videos,omitempty"`
}

type VideoListResponseVideos struct {
}

type VideoPlayer struct {
	// EmbedHtml: Iframe embed for the video.
	EmbedHtml string `json:"embedHtml,omitempty"`
}

type VideoSnippet struct {
	// CategoryId: Video category the video belongs to.
	CategoryId string `json:"categoryId,omitempty"`

	// ChannelId: Channel the video was uploaded into.
	ChannelId string `json:"channelId,omitempty"`

	// Description: Description of the video.
	Description string `json:"description,omitempty"`

	// PublishedAt: Date time the video was uploaded.
	PublishedAt string `json:"publishedAt,omitempty"`

	// Tags: Textual tags associated with the video.
	Tags []string `json:"tags,omitempty"`

	// Thumbnails: Video thumbnails.
	Thumbnails *VideoSnippetThumbnails `json:"thumbnails,omitempty"`

	// Title: Title of the video.
	Title string `json:"title,omitempty"`
}

type VideoSnippetThumbnails struct {
}

type VideoStatistics struct {
	// CommentCount: Number of comments for this video.
	CommentCount uint64 `json:"commentCount,omitempty,string"`

	// DislikeCount: Number of times the video was disliked.
	DislikeCount uint64 `json:"dislikeCount,omitempty,string"`

	// FavoriteCount: Number of times the video was added to a user's
	// favorites list.
	FavoriteCount uint64 `json:"favoriteCount,omitempty,string"`

	// LikeCount: Number of times the video was liked.
	LikeCount uint64 `json:"likeCount,omitempty,string"`

	// ViewCount: Number of times the video was viewed.
	ViewCount uint64 `json:"viewCount,omitempty,string"`
}

type VideoStatus struct {
	// FailureReason: Present only if the uploadStatus indicates a failed
	// upload.
	FailureReason string `json:"failureReason,omitempty"`

	// PrivacyStatus: Privacy of the video.
	PrivacyStatus string `json:"privacyStatus,omitempty"`

	// RejectionReason: Present only if the uploadStatus indicates a
	// rejected upload.
	RejectionReason string `json:"rejectionReason,omitempty"`

	// UploadStatus: Status of the video upload.
	UploadStatus string `json:"uploadStatus,omitempty"`
}

// method id "youtube.channels.list":

type ChannelsListCall struct {
	s    *Service
	part string
	opt_ map[string]interface{}
}

// List: Browse the YouTube channel collection. Either the 'id' or
// 'mine' parameter must be set.
func (r *ChannelsService) List(part string) *ChannelsListCall {
	c := &ChannelsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	return c
}

// Id sets the optional parameter "id": YouTube IDs of the channels to
// be returned.
func (c *ChannelsListCall) Id(id string) *ChannelsListCall {
	c.opt_["id"] = id
	return c
}

// Mine sets the optional parameter "mine": Flag indicating only return
// the channel ids of the authenticated user.
func (c *ChannelsListCall) Mine(mine string) *ChannelsListCall {
	c.opt_["mine"] = mine
	return c
}

func (c *ChannelsListCall) Do() (*ChannelListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("part", fmt.Sprintf("%v", c.part))
	if v, ok := c.opt_["id"]; ok {
		params.Set("id", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["mine"]; ok {
		params.Set("mine", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3alpha/", "channels")
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
	ret := new(ChannelListResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Browse the YouTube channel collection. Either the 'id' or 'mine' parameter must be set.",
	//   "httpMethod": "GET",
	//   "id": "youtube.channels.list",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "YouTube IDs of the channels to be returned.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "mine": {
	//       "description": "Flag indicating only return the channel ids of the authenticated user.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "Parts of the channel resource to be returned.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "channels",
	//   "response": {
	//     "$ref": "ChannelListResponse"
	//   }
	// }

}

// method id "youtube.playlistitems.list":

type PlaylistitemsListCall struct {
	s    *Service
	part string
	opt_ map[string]interface{}
}

// List: Browse the YouTube playlist collection.
func (r *PlaylistitemsService) List(part string) *PlaylistitemsListCall {
	c := &PlaylistitemsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	return c
}

// Id sets the optional parameter "id": YouTube IDs of the playlists to
// be returned.
func (c *PlaylistitemsListCall) Id(id string) *PlaylistitemsListCall {
	c.opt_["id"] = id
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return
func (c *PlaylistitemsListCall) MaxResults(maxResults int64) *PlaylistitemsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PlaylistId sets the optional parameter "playlistId": Retrieves
// playlist items from the given playlist id.
func (c *PlaylistitemsListCall) PlaylistId(playlistId string) *PlaylistitemsListCall {
	c.opt_["playlistId"] = playlistId
	return c
}

// StartIndex sets the optional parameter "startIndex": Index of the
// first element to return (starts at 0)
func (c *PlaylistitemsListCall) StartIndex(startIndex int64) *PlaylistitemsListCall {
	c.opt_["startIndex"] = startIndex
	return c
}

func (c *PlaylistitemsListCall) Do() (*PlaylistItemListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("part", fmt.Sprintf("%v", c.part))
	if v, ok := c.opt_["id"]; ok {
		params.Set("id", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["playlistId"]; ok {
		params.Set("playlistId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["startIndex"]; ok {
		params.Set("startIndex", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3alpha/", "playlistitems")
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
	ret := new(PlaylistItemListResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Browse the YouTube playlist collection.",
	//   "httpMethod": "GET",
	//   "id": "youtube.playlistitems.list",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "YouTube IDs of the playlists to be returned.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "50",
	//       "description": "Maximum number of results to return",
	//       "format": "uint32",
	//       "location": "query",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "part": {
	//       "description": "Parts of the playlist resource to be returned.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "playlistId": {
	//       "description": "Retrieves playlist items from the given playlist id.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "startIndex": {
	//       "description": "Index of the first element to return (starts at 0)",
	//       "format": "uint32",
	//       "location": "query",
	//       "minimum": "0",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "playlistitems",
	//   "response": {
	//     "$ref": "PlaylistItemListResponse"
	//   }
	// }

}

// method id "youtube.playlists.list":

type PlaylistsListCall struct {
	s    *Service
	id   string
	part string
	opt_ map[string]interface{}
}

// List: Browse the YouTube playlist collection.
func (r *PlaylistsService) List(id string, part string) *PlaylistsListCall {
	c := &PlaylistsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	c.part = part
	return c
}

func (c *PlaylistsListCall) Do() (*PlaylistListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("id", fmt.Sprintf("%v", c.id))
	params.Set("part", fmt.Sprintf("%v", c.part))
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3alpha/", "playlists")
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
	ret := new(PlaylistListResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Browse the YouTube playlist collection.",
	//   "httpMethod": "GET",
	//   "id": "youtube.playlists.list",
	//   "parameterOrder": [
	//     "id",
	//     "part"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "YouTube IDs of the playlists to be returned.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "Parts of the playlist resource to be returned.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "playlists",
	//   "response": {
	//     "$ref": "PlaylistListResponse"
	//   }
	// }

}

// method id "youtube.search.list":

type SearchListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Universal search for youtube.
func (r *SearchService) List() *SearchListCall {
	c := &SearchListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of search results to return per page.
func (c *SearchListCall) MaxResults(maxResults int64) *SearchListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// Order sets the optional parameter "order": Sort order.
func (c *SearchListCall) Order(order string) *SearchListCall {
	c.opt_["order"] = order
	return c
}

// Q sets the optional parameter "q": Query to search in Youtube.
func (c *SearchListCall) Q(q string) *SearchListCall {
	c.opt_["q"] = q
	return c
}

// StartIndex sets the optional parameter "startIndex": Index of the
// first search result to return.
func (c *SearchListCall) StartIndex(startIndex int64) *SearchListCall {
	c.opt_["startIndex"] = startIndex
	return c
}

// Type sets the optional parameter "type": Type of resource to search.
func (c *SearchListCall) Type(type_ string) *SearchListCall {
	c.opt_["type"] = type_
	return c
}

func (c *SearchListCall) Do() (*SearchListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["order"]; ok {
		params.Set("order", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["q"]; ok {
		params.Set("q", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["startIndex"]; ok {
		params.Set("startIndex", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["type"]; ok {
		params.Set("type", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3alpha/", "search")
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
	ret := new(SearchListResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Universal search for youtube.",
	//   "httpMethod": "GET",
	//   "id": "youtube.search.list",
	//   "parameters": {
	//     "maxResults": {
	//       "default": "25",
	//       "description": "Maximum number of search results to return per page.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "50",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "order": {
	//       "default": "SEARCH_SORT_RELEVANCE",
	//       "description": "Sort order.",
	//       "enum": [
	//         "date",
	//         "rating",
	//         "relevance",
	//         "view_count"
	//       ],
	//       "enumDescriptions": [
	//         "Sort according to the date.",
	//         "Sort according to the rating.",
	//         "Sort according to the relevance.",
	//         "Sort according to the view count."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "q": {
	//       "description": "Query to search in Youtube.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "startIndex": {
	//       "default": "0",
	//       "description": "Index of the first search result to return.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "1000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "type": {
	//       "description": "Type of resource to search.",
	//       "enum": [
	//         "channel",
	//         "playlist",
	//         "video"
	//       ],
	//       "enumDescriptions": [
	//         "Search for channels.",
	//         "Search for playlists.",
	//         "Search for videos."
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "search",
	//   "response": {
	//     "$ref": "SearchListResponse"
	//   }
	// }

}

// method id "youtube.videos.list":

type VideosListCall struct {
	s    *Service
	id   string
	part string
	opt_ map[string]interface{}
}

// List: Browse the YouTube video collection.
func (r *VideosService) List(id string, part string) *VideosListCall {
	c := &VideosListCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	c.part = part
	return c
}

func (c *VideosListCall) Do() (*VideoListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("id", fmt.Sprintf("%v", c.id))
	params.Set("part", fmt.Sprintf("%v", c.part))
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3alpha/", "videos")
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
	ret := new(VideoListResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Browse the YouTube video collection.",
	//   "httpMethod": "GET",
	//   "id": "youtube.videos.list",
	//   "parameterOrder": [
	//     "id",
	//     "part"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "YouTube IDs of the videos to be returned.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "Parts of the video resource to be returned.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "videos",
	//   "response": {
	//     "$ref": "VideoListResponse"
	//   }
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
