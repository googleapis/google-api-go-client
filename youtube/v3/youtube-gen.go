// Package youtube provides access to the YouTube API.
//
// See https://developers.google.com/youtube
//
// Usage example:
//
//   import "code.google.com/p/google-api-go-client/youtube/v3"
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

const apiId = "youtube:v3"
const apiName = "youtube"
const apiVersion = "v3"
const basePath = "https://www.googleapis.com/youtube/v3/"

// OAuth2 scopes used by this API.
const (
	// Manage your YouTube account
	YoutubeScope = "https://www.googleapis.com/auth/youtube"

	// View your YouTube account
	YoutubeReadonlyScope = "https://www.googleapis.com/auth/youtube.readonly"

	// Manage your YouTube videos
	YoutubeUploadScope = "https://www.googleapis.com/auth/youtube.upload"

	// View and manage your assets and associated content on YouTube
	YoutubepartnerScope = "https://www.googleapis.com/auth/youtubepartner"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Activities = &ActivitiesService{s: s}
	s.Channels = &ChannelsService{s: s}
	s.GuideCategories = &GuideCategoriesService{s: s}
	s.PlaylistItems = &PlaylistItemsService{s: s}
	s.Playlists = &PlaylistsService{s: s}
	s.Search = &SearchService{s: s}
	s.Subscriptions = &SubscriptionsService{s: s}
	s.VideoCategories = &VideoCategoriesService{s: s}
	s.Videos = &VideosService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	Activities *ActivitiesService

	Channels *ChannelsService

	GuideCategories *GuideCategoriesService

	PlaylistItems *PlaylistItemsService

	Playlists *PlaylistsService

	Search *SearchService

	Subscriptions *SubscriptionsService

	VideoCategories *VideoCategoriesService

	Videos *VideosService
}

type ActivitiesService struct {
	s *Service
}

type ChannelsService struct {
	s *Service
}

type GuideCategoriesService struct {
	s *Service
}

type PlaylistItemsService struct {
	s *Service
}

type PlaylistsService struct {
	s *Service
}

type SearchService struct {
	s *Service
}

type SubscriptionsService struct {
	s *Service
}

type VideoCategoriesService struct {
	s *Service
}

type VideosService struct {
	s *Service
}

type Activity struct {
	// ContentDetails: Type specific information about the activity.
	ContentDetails *ActivityContentDetails `json:"contentDetails,omitempty"`

	// Etag: The eTag of the activity.
	Etag string `json:"etag,omitempty"`

	// Id: The unique ID of the activity.
	Id string `json:"id,omitempty"`

	// Kind: The type of this API response.
	Kind string `json:"kind,omitempty"`

	// Snippet: Basic details about the activity: title, description,
	// thumbnails.
	Snippet *ActivitySnippet `json:"snippet,omitempty"`
}

type ActivityContentDetails struct {
	// Bulletin: Only present if the type is "bulletin".
	Bulletin *ActivityContentDetailsBulletin `json:"bulletin,omitempty"`

	// Comment: Only present if the type is "comment".
	Comment *ActivityContentDetailsComment `json:"comment,omitempty"`

	// Favorite: Only present if the type is "favorite".
	Favorite *ActivityContentDetailsFavorite `json:"favorite,omitempty"`

	// Like: Only present if the type is "like".
	Like *ActivityContentDetailsLike `json:"like,omitempty"`

	// PlaylistItem: Only present if the type is "playlistItem".
	PlaylistItem *ActivityContentDetailsPlaylistItem `json:"playlistItem,omitempty"`

	// Recommendation: Only set if the type is "recommendation".
	Recommendation *ActivityContentDetailsRecommendation `json:"recommendation,omitempty"`

	// Social: Only present if the type is "social".
	Social *ActivityContentDetailsSocial `json:"social,omitempty"`

	// Subscription: Only present if the type is "subscription".
	Subscription *ActivityContentDetailsSubscription `json:"subscription,omitempty"`

	// Upload: Only present if the type is "upload".
	Upload *ActivityContentDetailsUpload `json:"upload,omitempty"`
}

type ActivityContentDetailsBulletin struct {
	// ResourceId: ID of the resource this bulletin is about.
	ResourceId *ResourceId `json:"resourceId,omitempty"`
}

type ActivityContentDetailsComment struct {
	// ResourceId: ID of the commented resource.
	ResourceId *ResourceId `json:"resourceId,omitempty"`
}

type ActivityContentDetailsFavorite struct {
	// ResourceId: ID of the favorited resource.
	ResourceId *ResourceId `json:"resourceId,omitempty"`
}

type ActivityContentDetailsLike struct {
	// ResourceId: ID of the rated resource.
	ResourceId *ResourceId `json:"resourceId,omitempty"`
}

type ActivityContentDetailsPlaylistItem struct {
	// PlaylistId: ID of the playlist the resource was added to.
	PlaylistId string `json:"playlistId,omitempty"`

	// ResourceId: ID of the resource added to the playlist.
	ResourceId *ResourceId `json:"resourceId,omitempty"`
}

type ActivityContentDetailsRecommendation struct {
	// Reason: Reason for which the video was recommended.
	Reason string `json:"reason,omitempty"`

	// ResourceId: ID of the recommended resource.
	ResourceId *ResourceId `json:"resourceId,omitempty"`

	// SeedResourceId: ID of the video that caused this recommendation.
	SeedResourceId *ResourceId `json:"seedResourceId,omitempty"`
}

type ActivityContentDetailsSocial struct {
	// Author: Author of the post.
	Author string `json:"author,omitempty"`

	// ImageUrl: Image of the post author.
	ImageUrl string `json:"imageUrl,omitempty"`

	// ReferenceUrl: Url of the social post.
	ReferenceUrl string `json:"referenceUrl,omitempty"`

	// ResourceId: ID of the resource this social activity is about.
	ResourceId *ResourceId `json:"resourceId,omitempty"`

	// Type: Type of the social network.
	Type string `json:"type,omitempty"`
}

type ActivityContentDetailsSubscription struct {
	// ResourceId: ID of the resource subscribed to.
	ResourceId *ResourceId `json:"resourceId,omitempty"`
}

type ActivityContentDetailsUpload struct {
	// VideoId: ID of the uploaded video.
	VideoId string `json:"videoId,omitempty"`
}

type ActivityListResponse struct {
	// Etag: The eTag of the response.
	Etag string `json:"etag,omitempty"`

	// Items: List of activities matching the request criteria.
	Items []*Activity `json:"items,omitempty"`

	// Kind: The type of this API response.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token to the next page.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PageInfo: Paging information for the list result.
	PageInfo *PageInfo `json:"pageInfo,omitempty"`

	// PrevPageToken: Token to the previous page.
	PrevPageToken string `json:"prevPageToken,omitempty"`
}

type ActivitySnippet struct {
	// ChannelId: Channel publishing the activity.
	ChannelId string `json:"channelId,omitempty"`

	// Description: Description of the activity.
	Description string `json:"description,omitempty"`

	// GroupId: Id of the group that this activity is part of.
	GroupId string `json:"groupId,omitempty"`

	// PublishedAt: Date and time the activity was published at.
	PublishedAt string `json:"publishedAt,omitempty"`

	// Thumbnails: Activity thumbnails.
	Thumbnails *ActivitySnippetThumbnails `json:"thumbnails,omitempty"`

	// Title: Title of the activity.
	Title string `json:"title,omitempty"`

	// Type: Type of the activity.
	Type string `json:"type,omitempty"`
}

type ActivitySnippetThumbnails struct {
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

	// TopicDetails: Information about channel topics
	TopicDetails *ChannelTopicDetails `json:"topicDetails,omitempty"`
}

type ChannelContentDetails struct {
	// PrivacyStatus: Privacy status of the channel.
	PrivacyStatus string `json:"privacyStatus,omitempty"`

	// Uploads: The ID of the playlist containing the uploads of this
	// channel.
	Uploads string `json:"uploads,omitempty"`
}

type ChannelListResponse struct {
	// Etag: The eTag of the response.
	Etag string `json:"etag,omitempty"`

	// Items: List of channels matching the request criteria.
	Items []*Channel `json:"items,omitempty"`

	// Kind: The type of this API response.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token to the next page.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PageInfo: Paging information for the list result.
	PageInfo *PageInfo `json:"pageInfo,omitempty"`

	// PrevPageToken: Token to the previous page.
	PrevPageToken string `json:"prevPageToken,omitempty"`
}

type ChannelSnippet struct {
	// ChannelId: Id of the channel.
	ChannelId string `json:"channelId,omitempty"`

	// Description: Description of the channel.
	Description string `json:"description,omitempty"`

	// PublishedAt: Date and time the channel was published at.
	PublishedAt string `json:"publishedAt,omitempty"`

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

type ChannelTopicDetails struct {
	// TopicIds: List of topic ids for this channel *
	TopicIds []string `json:"topicIds,omitempty"`
}

type GuideCategory struct {
	// Etag: The eTag of the guide category.
	Etag string `json:"etag,omitempty"`

	// Id: The unique ID of the guide category.
	Id string `json:"id,omitempty"`

	// Kind: The type of this API resource.
	Kind string `json:"kind,omitempty"`

	// Snippet: Snippet of the category.
	Snippet *GuideCategorySnippet `json:"snippet,omitempty"`
}

type GuideCategoryListResponse struct {
	// Etag: The eTag of the response.
	Etag string `json:"etag,omitempty"`

	// Items: List of categories matching the request criteria.
	Items []*GuideCategory `json:"items,omitempty"`

	// Kind: The type of this API response.
	Kind string `json:"kind,omitempty"`
}

type GuideCategorySnippet struct {
	// ChannelId: Channel publishing the guide category.
	ChannelId string `json:"channelId,omitempty"`

	// Title: Title of the guide category.
	Title string `json:"title,omitempty"`
}

type PageInfo struct {
	// ResultsPerPage: The number of results to display for each page.
	ResultsPerPage int64 `json:"resultsPerPage,omitempty"`

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

	// Status: Status of the playlist: only privacy_status for now.
	Status *PlaylistStatus `json:"status,omitempty"`
}

type PlaylistItem struct {
	// ContentDetails: Content details about the playlist item: start and
	// end clipping time.
	ContentDetails *PlaylistItemContentDetails `json:"contentDetails,omitempty"`

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

type PlaylistItemContentDetails struct {
	// EndAt: The time video playback ends.
	EndAt string `json:"endAt,omitempty"`

	// Note: The user-generated note for this item.
	Note string `json:"note,omitempty"`

	// StartAt: The time video playback begins.
	StartAt string `json:"startAt,omitempty"`

	// VideoId: ID of the video.
	VideoId string `json:"videoId,omitempty"`
}

type PlaylistItemListResponse struct {
	// Etag: The eTag of the response.
	Etag string `json:"etag,omitempty"`

	// Items: List of playlist items matching the request criteria.
	Items []*PlaylistItem `json:"items,omitempty"`

	// Kind: The type of this API response.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token to the next page.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PageInfo: Paging information for the list result.
	PageInfo *PageInfo `json:"pageInfo,omitempty"`

	// PrevPageToken: Token to the previous page.
	PrevPageToken string `json:"prevPageToken,omitempty"`
}

type PlaylistItemSnippet struct {
	// ChannelId: Channel publishing the playlist item.
	ChannelId string `json:"channelId,omitempty"`

	// Description: Description of the playlist item.
	Description string `json:"description,omitempty"`

	// PlaylistId: The playlist the item is part of.
	PlaylistId string `json:"playlistId,omitempty"`

	// Position: The position of the item within the playlist.
	Position int64 `json:"position,omitempty"`

	// PublishedAt: Date and time the playlist item was published at.
	PublishedAt string `json:"publishedAt,omitempty"`

	// ResourceId: The ID of the resource referenced by the playlist item.
	ResourceId *ResourceId `json:"resourceId,omitempty"`

	// Thumbnails: Playlist item thumbnails.
	Thumbnails *PlaylistItemSnippetThumbnails `json:"thumbnails,omitempty"`

	// Title: Title of the playlist item.
	Title string `json:"title,omitempty"`
}

type PlaylistItemSnippetThumbnails struct {
}

type PlaylistListResponse struct {
	// Etag: The eTag of the response.
	Etag string `json:"etag,omitempty"`

	// Items: List of playlists matching the request criteria.
	Items []*Playlist `json:"items,omitempty"`

	// Kind: The type of this API response.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token to the next page.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PageInfo: Paging information for the list result.
	PageInfo *PageInfo `json:"pageInfo,omitempty"`

	// PrevPageToken: Token to the previous page.
	PrevPageToken string `json:"prevPageToken,omitempty"`
}

type PlaylistSnippet struct {
	// ChannelId: Channel publishing the playlist.
	ChannelId string `json:"channelId,omitempty"`

	// Description: Description of the playlist.
	Description string `json:"description,omitempty"`

	// PublishedAt: Date and time the playlist was published at.
	PublishedAt string `json:"publishedAt,omitempty"`

	// Thumbnails: Playlist thumbnails.
	Thumbnails *PlaylistSnippetThumbnails `json:"thumbnails,omitempty"`

	// Title: Title of the playlist.
	Title string `json:"title,omitempty"`
}

type PlaylistSnippetThumbnails struct {
}

type PlaylistStatus struct {
	// PrivacyStatus: Privacy of the playlist.
	PrivacyStatus string `json:"privacyStatus,omitempty"`
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

	// Items: List of results matching the request criteria.
	Items []*SearchResult `json:"items,omitempty"`

	// Kind: The type of this API response.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token to the next page.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PageInfo: Paging information for the search result.
	PageInfo *PageInfo `json:"pageInfo,omitempty"`

	// PrevPageToken: Token to the previous page.
	PrevPageToken string `json:"prevPageToken,omitempty"`
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
	// ChannelId: Channel publishing the found resource.
	ChannelId string `json:"channelId,omitempty"`

	// Description: Description of the found resource.
	Description string `json:"description,omitempty"`

	// PublishedAt: Date and time the found resource was published at.
	PublishedAt string `json:"publishedAt,omitempty"`

	// Thumbnails: Thumbnails for the found resource.
	Thumbnails *SearchResultSnippetThumbnails `json:"thumbnails,omitempty"`

	// Title: Title of the found resource.
	Title string `json:"title,omitempty"`
}

type SearchResultSnippetThumbnails struct {
}

type Subscription struct {
	// ContentDetails: Basic statistics about the subscription
	ContentDetails *SubscriptionContentDetails `json:"contentDetails,omitempty"`

	// Etag: The eTag of the subscription.
	Etag string `json:"etag,omitempty"`

	// Id: The unique id of the subscription.
	Id string `json:"id,omitempty"`

	// Kind: The type of this API resource.
	Kind string `json:"kind,omitempty"`

	// Snippet: Basic details about the subscription
	Snippet *SubscriptionSnippet `json:"snippet,omitempty"`
}

type SubscriptionContentDetails struct {
	// NewItemCount: Number of new items in the subscription since its
	// content was last read.
	NewItemCount int64 `json:"newItemCount,omitempty"`

	// TotalItemCount: Approximate total number of items the subscription
	// points to.
	TotalItemCount int64 `json:"totalItemCount,omitempty"`
}

type SubscriptionListResponse struct {
	// Etag: The eTag of the response.
	Etag string `json:"etag,omitempty"`

	// Items: List of subscriptions matching the request criteria.
	Items []*Subscription `json:"items,omitempty"`

	// Kind: The type of this API response.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token to the next page.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PageInfo: Paging information for the list result.
	PageInfo *PageInfo `json:"pageInfo,omitempty"`

	// PrevPageToken: Token to the previous page.
	PrevPageToken string `json:"prevPageToken,omitempty"`
}

type SubscriptionSnippet struct {
	// ChannelId: Channel publishing the subscription.
	ChannelId string `json:"channelId,omitempty"`

	// Description: Description of the subscription.
	Description string `json:"description,omitempty"`

	// PublishedAt: Date and time the subscription was published at.
	PublishedAt string `json:"publishedAt,omitempty"`

	// ResourceId: The resource subscribed to.
	ResourceId *ResourceId `json:"resourceId,omitempty"`

	// Thumbnails: Subscription thumbnails.
	Thumbnails *SubscriptionSnippetThumbnails `json:"thumbnails,omitempty"`

	// Title: Title of the subscription.
	Title string `json:"title,omitempty"`
}

type SubscriptionSnippetThumbnails struct {
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

	// TopicDetails: Topics related to the video
	TopicDetails *VideoTopicDetails `json:"topicDetails,omitempty"`
}

type VideoCategory struct {
	// Etag: The eTag of the video.
	Etag string `json:"etag,omitempty"`

	// Id: The unique id of the video category.
	Id string `json:"id,omitempty"`

	// Kind: The type of this API resource.
	Kind string `json:"kind,omitempty"`

	// Snippet: Basic details about the video category.
	Snippet *VideoCategorySnippet `json:"snippet,omitempty"`
}

type VideoCategoryListResponse struct {
	// Etag: The eTag of the response.
	Etag string `json:"etag,omitempty"`

	// Items: List of video categories matching the request criteria.
	Items []*VideoCategory `json:"items,omitempty"`

	// Kind: The type of this API response.
	Kind string `json:"kind,omitempty"`
}

type VideoCategorySnippet struct {
	// ChannelId: Channel publishing the video category.
	ChannelId string `json:"channelId,omitempty"`

	// Title: Title of the video category.
	Title string `json:"title,omitempty"`
}

type VideoContentDetails struct {
	// Duration: Duration of the video.
	Duration string `json:"duration,omitempty"`

	// RegionRestriction: Region restriction of the video.
	RegionRestriction *VideoContentDetailsRegionRestriction `json:"regionRestriction,omitempty"`
}

type VideoContentDetailsRegionRestriction struct {
	// Allowed: List of allowed region codes.
	Allowed []string `json:"allowed,omitempty"`

	// Blocked: List of blocked region codes.
	Blocked []string `json:"blocked,omitempty"`
}

type VideoListResponse struct {
	// Etag: The eTag of the response.
	Etag string `json:"etag,omitempty"`

	// Items: List of videos matching the request criteria.
	Items []*Video `json:"items,omitempty"`

	// Kind: The type of this API response.
	Kind string `json:"kind,omitempty"`
}

type VideoPlayer struct {
	// EmbedHtml: Iframe embed for the video.
	EmbedHtml string `json:"embedHtml,omitempty"`
}

type VideoSnippet struct {
	// CategoryId: Video category the video belongs to.
	CategoryId string `json:"categoryId,omitempty"`

	// ChannelId: Channel publishing the video.
	ChannelId string `json:"channelId,omitempty"`

	// Description: Description of the video.
	Description string `json:"description,omitempty"`

	// PublishedAt: Date and time the video was published at.
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

type VideoTopicDetails struct {
	// TopicIds: List of topic ids for this video *
	TopicIds []string `json:"topicIds,omitempty"`
}

// method id "youtube.activities.insert":

type ActivitiesInsertCall struct {
	s        *Service
	part     string
	activity *Activity
	opt_     map[string]interface{}
}

// Insert: Post a channel bulletin.
func (r *ActivitiesService) Insert(part string, activity *Activity) *ActivitiesInsertCall {
	c := &ActivitiesInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	c.activity = activity
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The authenticated user acts on behalf of
// this content owner.
func (c *ActivitiesInsertCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *ActivitiesInsertCall {
	c.opt_["onBehalfOfContentOwner"] = onBehalfOfContentOwner
	return c
}

func (c *ActivitiesInsertCall) Do() (*Activity, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.activity)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("part", fmt.Sprintf("%v", c.part))
	if v, ok := c.opt_["onBehalfOfContentOwner"]; ok {
		params.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "activities")
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
	//   "description": "Post a channel bulletin.",
	//   "httpMethod": "POST",
	//   "id": "youtube.activities.insert",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The authenticated user acts on behalf of this content owner.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "One or more parts to return on the current request.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities",
	//   "request": {
	//     "$ref": "Activity"
	//   },
	//   "response": {
	//     "$ref": "Activity"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube"
	//   ]
	// }

}

// method id "youtube.activities.list":

type ActivitiesListCall struct {
	s    *Service
	part string
	opt_ map[string]interface{}
}

// List: Browse the YouTube channel activity collection.
func (r *ActivitiesService) List(part string) *ActivitiesListCall {
	c := &ActivitiesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	return c
}

// ChannelId sets the optional parameter "channelId": YouTube ID of the
// channel.
func (c *ActivitiesListCall) ChannelId(channelId string) *ActivitiesListCall {
	c.opt_["channelId"] = channelId
	return c
}

// Home sets the optional parameter "home": Flag indicating to return
// user's homepage feed.
func (c *ActivitiesListCall) Home(home string) *ActivitiesListCall {
	c.opt_["home"] = home
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return
func (c *ActivitiesListCall) MaxResults(maxResults int64) *ActivitiesListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// Mine sets the optional parameter "mine": Flag indicating to return
// user's activities.
func (c *ActivitiesListCall) Mine(mine string) *ActivitiesListCall {
	c.opt_["mine"] = mine
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The authenticated user acts on behalf of
// this content owner.
func (c *ActivitiesListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *ActivitiesListCall {
	c.opt_["onBehalfOfContentOwner"] = onBehalfOfContentOwner
	return c
}

// PageToken sets the optional parameter "pageToken": Token for the page
// selection.
func (c *ActivitiesListCall) PageToken(pageToken string) *ActivitiesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// PublishedAfter sets the optional parameter "publishedAfter": Only
// return activities published after given date (inclusive).
func (c *ActivitiesListCall) PublishedAfter(publishedAfter string) *ActivitiesListCall {
	c.opt_["publishedAfter"] = publishedAfter
	return c
}

// PublishedBefore sets the optional parameter "publishedBefore": Only
// return activities published before given date (exclusive).
func (c *ActivitiesListCall) PublishedBefore(publishedBefore string) *ActivitiesListCall {
	c.opt_["publishedBefore"] = publishedBefore
	return c
}

func (c *ActivitiesListCall) Do() (*ActivityListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("part", fmt.Sprintf("%v", c.part))
	if v, ok := c.opt_["channelId"]; ok {
		params.Set("channelId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["home"]; ok {
		params.Set("home", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["mine"]; ok {
		params.Set("mine", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["onBehalfOfContentOwner"]; ok {
		params.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["publishedAfter"]; ok {
		params.Set("publishedAfter", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["publishedBefore"]; ok {
		params.Set("publishedBefore", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "activities")
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
	ret := new(ActivityListResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Browse the YouTube channel activity collection.",
	//   "httpMethod": "GET",
	//   "id": "youtube.activities.list",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "channelId": {
	//       "description": "YouTube ID of the channel.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "home": {
	//       "description": "Flag indicating to return user's homepage feed.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "5",
	//       "description": "Maximum number of results to return",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "50",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "mine": {
	//       "description": "Flag indicating to return user's activities.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The authenticated user acts on behalf of this content owner.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Token for the page selection.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "Activity parts to include in the returned response. Valid values are: id, snippet and contentDetails.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "publishedAfter": {
	//       "description": "Only return activities published after given date (inclusive).",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "publishedBefore": {
	//       "description": "Only return activities published before given date (exclusive).",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "activities",
	//   "response": {
	//     "$ref": "ActivityListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtube.readonly"
	//   ]
	// }

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

// CategoryId sets the optional parameter "categoryId": Filter to
// retrieve the channels within the given category ID.
func (c *ChannelsListCall) CategoryId(categoryId string) *ChannelsListCall {
	c.opt_["categoryId"] = categoryId
	return c
}

// Id sets the optional parameter "id": YouTube IDs of the channels to
// be returned.
func (c *ChannelsListCall) Id(id string) *ChannelsListCall {
	c.opt_["id"] = id
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return
func (c *ChannelsListCall) MaxResults(maxResults int64) *ChannelsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// Mine sets the optional parameter "mine": Filter to only channels
// owned by authenticated user.
func (c *ChannelsListCall) Mine(mine string) *ChannelsListCall {
	c.opt_["mine"] = mine
	return c
}

// MySubscribers sets the optional parameter "mySubscribers": Filter to
// channels that subscribed to the channel of the authenticated user.
func (c *ChannelsListCall) MySubscribers(mySubscribers string) *ChannelsListCall {
	c.opt_["mySubscribers"] = mySubscribers
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The authenticated user acts on behalf of
// this content owner.
func (c *ChannelsListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *ChannelsListCall {
	c.opt_["onBehalfOfContentOwner"] = onBehalfOfContentOwner
	return c
}

// PageToken sets the optional parameter "pageToken": Token for the page
// selection.
func (c *ChannelsListCall) PageToken(pageToken string) *ChannelsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

func (c *ChannelsListCall) Do() (*ChannelListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("part", fmt.Sprintf("%v", c.part))
	if v, ok := c.opt_["categoryId"]; ok {
		params.Set("categoryId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["id"]; ok {
		params.Set("id", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["mine"]; ok {
		params.Set("mine", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["mySubscribers"]; ok {
		params.Set("mySubscribers", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["onBehalfOfContentOwner"]; ok {
		params.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "channels")
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
	//     "categoryId": {
	//       "description": "Filter to retrieve the channels within the given category ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "id": {
	//       "description": "YouTube IDs of the channels to be returned.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "5",
	//       "description": "Maximum number of results to return",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "50",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "mine": {
	//       "description": "Filter to only channels owned by authenticated user.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "mySubscribers": {
	//       "description": "Filter to channels that subscribed to the channel of the authenticated user.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The authenticated user acts on behalf of this content owner.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Token for the page selection.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "Channel parts to include in the returned response. Valid values are: id, snippet, contentDetails and topicDetails.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "channels",
	//   "response": {
	//     "$ref": "ChannelListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtube.readonly",
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtube.guideCategories.list":

type GuideCategoriesListCall struct {
	s    *Service
	part string
	opt_ map[string]interface{}
}

// List: Browse the YouTube guide category collection.
func (r *GuideCategoriesService) List(part string) *GuideCategoriesListCall {
	c := &GuideCategoriesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	return c
}

// Hl sets the optional parameter "hl": Language for the returned
// channelCategories.
func (c *GuideCategoriesListCall) Hl(hl string) *GuideCategoriesListCall {
	c.opt_["hl"] = hl
	return c
}

// Id sets the optional parameter "id": Comma-separated YouTube IDs of
// the channelCategories to be returned.
func (c *GuideCategoriesListCall) Id(id string) *GuideCategoriesListCall {
	c.opt_["id"] = id
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The authenticated user acts on behalf of
// this content owner.
func (c *GuideCategoriesListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *GuideCategoriesListCall {
	c.opt_["onBehalfOfContentOwner"] = onBehalfOfContentOwner
	return c
}

// RegionCode sets the optional parameter "regionCode": Return the
// channelCategories in the given region code.
func (c *GuideCategoriesListCall) RegionCode(regionCode string) *GuideCategoriesListCall {
	c.opt_["regionCode"] = regionCode
	return c
}

func (c *GuideCategoriesListCall) Do() (*GuideCategoryListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("part", fmt.Sprintf("%v", c.part))
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["id"]; ok {
		params.Set("id", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["onBehalfOfContentOwner"]; ok {
		params.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["regionCode"]; ok {
		params.Set("regionCode", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "guideCategories")
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
	ret := new(GuideCategoryListResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Browse the YouTube guide category collection.",
	//   "httpMethod": "GET",
	//   "id": "youtube.guideCategories.list",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "hl": {
	//       "default": "en-US",
	//       "description": "Language for the returned channelCategories.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "id": {
	//       "description": "Comma-separated YouTube IDs of the channelCategories to be returned.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The authenticated user acts on behalf of this content owner.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "Guide category parts to include in the returned response. Valid values are: id and snippet.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "regionCode": {
	//       "description": "Return the channelCategories in the given region code.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "guideCategories",
	//   "response": {
	//     "$ref": "GuideCategoryListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtube.readonly",
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtube.playlistItems.delete":

type PlaylistItemsDeleteCall struct {
	s    *Service
	id   string
	opt_ map[string]interface{}
}

// Delete: Deletes playlist items by IDs.
func (r *PlaylistItemsService) Delete(id string) *PlaylistItemsDeleteCall {
	c := &PlaylistItemsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The authenticated user acts on behalf of
// this content owner.
func (c *PlaylistItemsDeleteCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *PlaylistItemsDeleteCall {
	c.opt_["onBehalfOfContentOwner"] = onBehalfOfContentOwner
	return c
}

func (c *PlaylistItemsDeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("id", fmt.Sprintf("%v", c.id))
	if v, ok := c.opt_["onBehalfOfContentOwner"]; ok {
		params.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "playlistItems")
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
	//   "description": "Deletes playlist items by IDs.",
	//   "httpMethod": "DELETE",
	//   "id": "youtube.playlistItems.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "YouTube IDs of the playlist items to be deleted.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The authenticated user acts on behalf of this content owner.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "playlistItems",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtube.playlistItems.insert":

type PlaylistItemsInsertCall struct {
	s            *Service
	part         string
	playlistitem *PlaylistItem
	opt_         map[string]interface{}
}

// Insert: Insert a resource into a playlist.
func (r *PlaylistItemsService) Insert(part string, playlistitem *PlaylistItem) *PlaylistItemsInsertCall {
	c := &PlaylistItemsInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	c.playlistitem = playlistitem
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The authenticated user acts on behalf of
// this content owner.
func (c *PlaylistItemsInsertCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *PlaylistItemsInsertCall {
	c.opt_["onBehalfOfContentOwner"] = onBehalfOfContentOwner
	return c
}

func (c *PlaylistItemsInsertCall) Do() (*PlaylistItem, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.playlistitem)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("part", fmt.Sprintf("%v", c.part))
	if v, ok := c.opt_["onBehalfOfContentOwner"]; ok {
		params.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "playlistItems")
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
	ret := new(PlaylistItem)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Insert a resource into a playlist.",
	//   "httpMethod": "POST",
	//   "id": "youtube.playlistItems.insert",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The authenticated user acts on behalf of this content owner.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "One or more parts to return on the current request.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "playlistItems",
	//   "request": {
	//     "$ref": "PlaylistItem"
	//   },
	//   "response": {
	//     "$ref": "PlaylistItem"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtube.playlistItems.list":

type PlaylistItemsListCall struct {
	s    *Service
	part string
	opt_ map[string]interface{}
}

// List: Browse the YouTube playlist collection.
func (r *PlaylistItemsService) List(part string) *PlaylistItemsListCall {
	c := &PlaylistItemsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	return c
}

// Id sets the optional parameter "id": YouTube IDs of the playlist
// items to be returned.
func (c *PlaylistItemsListCall) Id(id string) *PlaylistItemsListCall {
	c.opt_["id"] = id
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return
func (c *PlaylistItemsListCall) MaxResults(maxResults int64) *PlaylistItemsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The authenticated user acts on behalf of
// this content owner.
func (c *PlaylistItemsListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *PlaylistItemsListCall {
	c.opt_["onBehalfOfContentOwner"] = onBehalfOfContentOwner
	return c
}

// PageToken sets the optional parameter "pageToken": Token for the page
// selection.
func (c *PlaylistItemsListCall) PageToken(pageToken string) *PlaylistItemsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// PlaylistId sets the optional parameter "playlistId": Retrieves
// playlist items from the given playlist id.
func (c *PlaylistItemsListCall) PlaylistId(playlistId string) *PlaylistItemsListCall {
	c.opt_["playlistId"] = playlistId
	return c
}

func (c *PlaylistItemsListCall) Do() (*PlaylistItemListResponse, error) {
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
	if v, ok := c.opt_["onBehalfOfContentOwner"]; ok {
		params.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["playlistId"]; ok {
		params.Set("playlistId", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "playlistItems")
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
	//   "id": "youtube.playlistItems.list",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "YouTube IDs of the playlist items to be returned.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "5",
	//       "description": "Maximum number of results to return",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "50",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The authenticated user acts on behalf of this content owner.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Token for the page selection.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "Playlist item parts to include in the returned response. Valid values are: id, snippet and contentDetails.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "playlistId": {
	//       "description": "Retrieves playlist items from the given playlist id.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "playlistItems",
	//   "response": {
	//     "$ref": "PlaylistItemListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtube.readonly",
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtube.playlistItems.update":

type PlaylistItemsUpdateCall struct {
	s            *Service
	part         string
	playlistitem *PlaylistItem
	opt_         map[string]interface{}
}

// Update: Update a playlist item.
func (r *PlaylistItemsService) Update(part string, playlistitem *PlaylistItem) *PlaylistItemsUpdateCall {
	c := &PlaylistItemsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	c.playlistitem = playlistitem
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The authenticated user acts on behalf of
// this content owner.
func (c *PlaylistItemsUpdateCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *PlaylistItemsUpdateCall {
	c.opt_["onBehalfOfContentOwner"] = onBehalfOfContentOwner
	return c
}

func (c *PlaylistItemsUpdateCall) Do() (*PlaylistItem, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.playlistitem)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("part", fmt.Sprintf("%v", c.part))
	if v, ok := c.opt_["onBehalfOfContentOwner"]; ok {
		params.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "playlistItems")
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
	ret := new(PlaylistItem)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update a playlist item.",
	//   "httpMethod": "PUT",
	//   "id": "youtube.playlistItems.update",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The authenticated user acts on behalf of this content owner.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "One or more parts to return on the current request.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "playlistItems",
	//   "request": {
	//     "$ref": "PlaylistItem"
	//   },
	//   "response": {
	//     "$ref": "PlaylistItem"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtube.playlists.delete":

type PlaylistsDeleteCall struct {
	s    *Service
	id   string
	opt_ map[string]interface{}
}

// Delete: Deletes playlists by IDs.
func (r *PlaylistsService) Delete(id string) *PlaylistsDeleteCall {
	c := &PlaylistsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The authenticated user acts on behalf of
// this content owner.
func (c *PlaylistsDeleteCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *PlaylistsDeleteCall {
	c.opt_["onBehalfOfContentOwner"] = onBehalfOfContentOwner
	return c
}

func (c *PlaylistsDeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("id", fmt.Sprintf("%v", c.id))
	if v, ok := c.opt_["onBehalfOfContentOwner"]; ok {
		params.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "playlists")
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
	//   "description": "Deletes playlists by IDs.",
	//   "httpMethod": "DELETE",
	//   "id": "youtube.playlists.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "YouTube IDs of the playlists to be deleted.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The authenticated user acts on behalf of this content owner.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "playlists",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtube.playlists.insert":

type PlaylistsInsertCall struct {
	s        *Service
	part     string
	playlist *Playlist
	opt_     map[string]interface{}
}

// Insert: Create a playlist.
func (r *PlaylistsService) Insert(part string, playlist *Playlist) *PlaylistsInsertCall {
	c := &PlaylistsInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	c.playlist = playlist
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The authenticated user acts on behalf of
// this content owner.
func (c *PlaylistsInsertCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *PlaylistsInsertCall {
	c.opt_["onBehalfOfContentOwner"] = onBehalfOfContentOwner
	return c
}

func (c *PlaylistsInsertCall) Do() (*Playlist, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.playlist)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("part", fmt.Sprintf("%v", c.part))
	if v, ok := c.opt_["onBehalfOfContentOwner"]; ok {
		params.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "playlists")
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
	ret := new(Playlist)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Create a playlist.",
	//   "httpMethod": "POST",
	//   "id": "youtube.playlists.insert",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The authenticated user acts on behalf of this content owner.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "One or more parts to return on the current request.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "playlists",
	//   "request": {
	//     "$ref": "Playlist"
	//   },
	//   "response": {
	//     "$ref": "Playlist"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtube.playlists.list":

type PlaylistsListCall struct {
	s    *Service
	part string
	opt_ map[string]interface{}
}

// List: Browse the YouTube playlist collection.
func (r *PlaylistsService) List(part string) *PlaylistsListCall {
	c := &PlaylistsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	return c
}

// Id sets the optional parameter "id": Comma-separated YouTube IDs of
// the playlists to be returned.
func (c *PlaylistsListCall) Id(id string) *PlaylistsListCall {
	c.opt_["id"] = id
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return
func (c *PlaylistsListCall) MaxResults(maxResults int64) *PlaylistsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// Mine sets the optional parameter "mine": Flag indicating only return
// the playlists of the authenticated user.
func (c *PlaylistsListCall) Mine(mine string) *PlaylistsListCall {
	c.opt_["mine"] = mine
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The authenticated user acts on behalf of
// this content owner.
func (c *PlaylistsListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *PlaylistsListCall {
	c.opt_["onBehalfOfContentOwner"] = onBehalfOfContentOwner
	return c
}

// PageToken sets the optional parameter "pageToken": Token for the page
// selection.
func (c *PlaylistsListCall) PageToken(pageToken string) *PlaylistsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

func (c *PlaylistsListCall) Do() (*PlaylistListResponse, error) {
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
	if v, ok := c.opt_["mine"]; ok {
		params.Set("mine", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["onBehalfOfContentOwner"]; ok {
		params.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "playlists")
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
	//     "part"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Comma-separated YouTube IDs of the playlists to be returned.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "5",
	//       "description": "Maximum number of results to return",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "50",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "mine": {
	//       "description": "Flag indicating only return the playlists of the authenticated user.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The authenticated user acts on behalf of this content owner.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Token for the page selection.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "Playlist parts to include in the returned response. Valid values are: id, snippet and status.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "playlists",
	//   "response": {
	//     "$ref": "PlaylistListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtube.readonly",
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtube.playlists.update":

type PlaylistsUpdateCall struct {
	s        *Service
	part     string
	playlist *Playlist
	opt_     map[string]interface{}
}

// Update: Update a playlist.
func (r *PlaylistsService) Update(part string, playlist *Playlist) *PlaylistsUpdateCall {
	c := &PlaylistsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	c.playlist = playlist
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The authenticated user acts on behalf of
// this content owner.
func (c *PlaylistsUpdateCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *PlaylistsUpdateCall {
	c.opt_["onBehalfOfContentOwner"] = onBehalfOfContentOwner
	return c
}

func (c *PlaylistsUpdateCall) Do() (*Playlist, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.playlist)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("part", fmt.Sprintf("%v", c.part))
	if v, ok := c.opt_["onBehalfOfContentOwner"]; ok {
		params.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "playlists")
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
	ret := new(Playlist)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update a playlist.",
	//   "httpMethod": "PUT",
	//   "id": "youtube.playlists.update",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The authenticated user acts on behalf of this content owner.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "One or more parts to return on the current request.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "playlists",
	//   "request": {
	//     "$ref": "Playlist"
	//   },
	//   "response": {
	//     "$ref": "Playlist"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtube.search.list":

type SearchListCall struct {
	s    *Service
	part string
	opt_ map[string]interface{}
}

// List: Universal search for youtube.
func (r *SearchService) List(part string) *SearchListCall {
	c := &SearchListCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of search results to return per page.
func (c *SearchListCall) MaxResults(maxResults int64) *SearchListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The authenticated user acts on behalf of
// this content owner.
func (c *SearchListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *SearchListCall {
	c.opt_["onBehalfOfContentOwner"] = onBehalfOfContentOwner
	return c
}

// Order sets the optional parameter "order": Sort order.
func (c *SearchListCall) Order(order string) *SearchListCall {
	c.opt_["order"] = order
	return c
}

// PageToken sets the optional parameter "pageToken": Token for the page
// selection.
func (c *SearchListCall) PageToken(pageToken string) *SearchListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Published sets the optional parameter "published": Only search for
// resources uploaded at a specific pediod
func (c *SearchListCall) Published(published string) *SearchListCall {
	c.opt_["published"] = published
	return c
}

// Q sets the optional parameter "q": Query to search in Youtube.
func (c *SearchListCall) Q(q string) *SearchListCall {
	c.opt_["q"] = q
	return c
}

// RelatedToVideo sets the optional parameter "relatedToVideo": Search
// for resources related to this video. Need to be used with type set to
// 'video'
func (c *SearchListCall) RelatedToVideo(relatedToVideo string) *SearchListCall {
	c.opt_["relatedToVideo"] = relatedToVideo
	return c
}

// TopicId sets the optional parameter "topicId": Only search for
// resources with the specified topic
func (c *SearchListCall) TopicId(topicId string) *SearchListCall {
	c.opt_["topicId"] = topicId
	return c
}

// Type sets the optional parameter "type": Type of resource to search.
func (c *SearchListCall) Type(type_ string) *SearchListCall {
	c.opt_["type"] = type_
	return c
}

// VideoCaption sets the optional parameter "videoCaption": Add a filter
// on the the presence of captions on the videos.
func (c *SearchListCall) VideoCaption(videoCaption string) *SearchListCall {
	c.opt_["videoCaption"] = videoCaption
	return c
}

// VideoDefinition sets the optional parameter "videoDefinition": Add a
// filter for the definition of the videos.
func (c *SearchListCall) VideoDefinition(videoDefinition string) *SearchListCall {
	c.opt_["videoDefinition"] = videoDefinition
	return c
}

// VideoDimension sets the optional parameter "videoDimension": Add a
// filter for the number of dimensions in the videos.
func (c *SearchListCall) VideoDimension(videoDimension string) *SearchListCall {
	c.opt_["videoDimension"] = videoDimension
	return c
}

// VideoDuration sets the optional parameter "videoDuration": Add a
// filter on the duration of the videos.
func (c *SearchListCall) VideoDuration(videoDuration string) *SearchListCall {
	c.opt_["videoDuration"] = videoDuration
	return c
}

// VideoLicense sets the optional parameter "videoLicense": Add a filter
// on the licensing of the videos.
func (c *SearchListCall) VideoLicense(videoLicense string) *SearchListCall {
	c.opt_["videoLicense"] = videoLicense
	return c
}

func (c *SearchListCall) Do() (*SearchListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("part", fmt.Sprintf("%v", c.part))
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["onBehalfOfContentOwner"]; ok {
		params.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["order"]; ok {
		params.Set("order", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["published"]; ok {
		params.Set("published", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["q"]; ok {
		params.Set("q", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["relatedToVideo"]; ok {
		params.Set("relatedToVideo", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["topicId"]; ok {
		params.Set("topicId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["type"]; ok {
		params.Set("type", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["videoCaption"]; ok {
		params.Set("videoCaption", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["videoDefinition"]; ok {
		params.Set("videoDefinition", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["videoDimension"]; ok {
		params.Set("videoDimension", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["videoDuration"]; ok {
		params.Set("videoDuration", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["videoLicense"]; ok {
		params.Set("videoLicense", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "search")
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
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "default": "5",
	//       "description": "Maximum number of search results to return per page.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "50",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The authenticated user acts on behalf of this content owner.",
	//       "location": "query",
	//       "type": "string"
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
	//     "pageToken": {
	//       "description": "Token for the page selection.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "Search result parts to include in the returned response. Valid values are: id and snippet.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "published": {
	//       "description": "Only search for resources uploaded at a specific pediod",
	//       "enum": [
	//         "any",
	//         "thisWeek",
	//         "today"
	//       ],
	//       "enumDescriptions": [
	//         "No filter on the release date",
	//         "Videos uploaded this month",
	//         "Videos uploaded today"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "q": {
	//       "description": "Query to search in Youtube.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "relatedToVideo": {
	//       "description": "Search for resources related to this video. Need to be used with type set to 'video'",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "topicId": {
	//       "description": "Only search for resources with the specified topic",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "type": {
	//       "default": "video,channel,playlist",
	//       "description": "Type of resource to search.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "videoCaption": {
	//       "description": "Add a filter on the the presence of captions on the videos.",
	//       "enum": [
	//         "any",
	//         "closedCaption",
	//         "none"
	//       ],
	//       "enumDescriptions": [
	//         "No filter on the captions.",
	//         "Videos with closed captions.",
	//         "Videos without captions."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "videoDefinition": {
	//       "description": "Add a filter for the definition of the videos.",
	//       "enum": [
	//         "any",
	//         "high",
	//         "standard"
	//       ],
	//       "enumDescriptions": [
	//         "No filter on the definition.",
	//         "Videos in high definition.",
	//         "Videos in standard definition."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "videoDimension": {
	//       "description": "Add a filter for the number of dimensions in the videos.",
	//       "enum": [
	//         "2d",
	//         "3d",
	//         "any"
	//       ],
	//       "enumDescriptions": [
	//         "Videos in two dimensions.",
	//         "Videos in three dimensions.",
	//         "No filter on the dimension."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "videoDuration": {
	//       "description": "Add a filter on the duration of the videos.",
	//       "enum": [
	//         "any",
	//         "long",
	//         "medium",
	//         "short"
	//       ],
	//       "enumDescriptions": [
	//         "No filter on the duration.",
	//         "Videos with a duration longer than 20 minutes.",
	//         "Videos with a duration between 4 and 20 minutes.",
	//         "Videos with a duration under 4 minutes."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "videoLicense": {
	//       "description": "Add a filter on the licensing of the videos.",
	//       "enum": [
	//         "any",
	//         "creativeCommon",
	//         "youtube"
	//       ],
	//       "enumDescriptions": [
	//         "No filter on the license.",
	//         "Videos under the Creative Common license.",
	//         "Videos under the YouTube license."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "search",
	//   "response": {
	//     "$ref": "SearchListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtube.readonly",
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtube.subscriptions.delete":

type SubscriptionsDeleteCall struct {
	s    *Service
	id   string
	opt_ map[string]interface{}
}

// Delete: Deletes subscriptions by IDs.
func (r *SubscriptionsService) Delete(id string) *SubscriptionsDeleteCall {
	c := &SubscriptionsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The authenticated user acts on behalf of
// this content owner.
func (c *SubscriptionsDeleteCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *SubscriptionsDeleteCall {
	c.opt_["onBehalfOfContentOwner"] = onBehalfOfContentOwner
	return c
}

func (c *SubscriptionsDeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("id", fmt.Sprintf("%v", c.id))
	if v, ok := c.opt_["onBehalfOfContentOwner"]; ok {
		params.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "subscriptions")
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
	//   "description": "Deletes subscriptions by IDs.",
	//   "httpMethod": "DELETE",
	//   "id": "youtube.subscriptions.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "YouTube IDs of the subscription to be deleted.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The authenticated user acts on behalf of this content owner.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "subscriptions",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtube.subscriptions.insert":

type SubscriptionsInsertCall struct {
	s            *Service
	part         string
	subscription *Subscription
	opt_         map[string]interface{}
}

// Insert: Insert a subscription.
func (r *SubscriptionsService) Insert(part string, subscription *Subscription) *SubscriptionsInsertCall {
	c := &SubscriptionsInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	c.subscription = subscription
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The authenticated user acts on behalf of
// this content owner.
func (c *SubscriptionsInsertCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *SubscriptionsInsertCall {
	c.opt_["onBehalfOfContentOwner"] = onBehalfOfContentOwner
	return c
}

func (c *SubscriptionsInsertCall) Do() (*Subscription, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.subscription)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("part", fmt.Sprintf("%v", c.part))
	if v, ok := c.opt_["onBehalfOfContentOwner"]; ok {
		params.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "subscriptions")
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
	ret := new(Subscription)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Insert a subscription.",
	//   "httpMethod": "POST",
	//   "id": "youtube.subscriptions.insert",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The authenticated user acts on behalf of this content owner.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "One or more parts to return on the current request.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "subscriptions",
	//   "request": {
	//     "$ref": "Subscription"
	//   },
	//   "response": {
	//     "$ref": "Subscription"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtube.subscriptions.list":

type SubscriptionsListCall struct {
	s    *Service
	part string
	opt_ map[string]interface{}
}

// List: Browses the subscriptions collection.
func (r *SubscriptionsService) List(part string) *SubscriptionsListCall {
	c := &SubscriptionsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	return c
}

// ChannelId sets the optional parameter "channelId": Only return
// subscriptions to given channelId.
func (c *SubscriptionsListCall) ChannelId(channelId string) *SubscriptionsListCall {
	c.opt_["channelId"] = channelId
	return c
}

// ForChannelId sets the optional parameter "forChannelId": Comma
// separated list of channel IDs to return subscriptions for.
func (c *SubscriptionsListCall) ForChannelId(forChannelId string) *SubscriptionsListCall {
	c.opt_["forChannelId"] = forChannelId
	return c
}

// Id sets the optional parameter "id": YouTube IDs of the subscriptions
// to be returned.
func (c *SubscriptionsListCall) Id(id string) *SubscriptionsListCall {
	c.opt_["id"] = id
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of search results to return per page.
func (c *SubscriptionsListCall) MaxResults(maxResults int64) *SubscriptionsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// Mine sets the optional parameter "mine": Flag indicating only return
// the subscriptions of the authenticated user.
func (c *SubscriptionsListCall) Mine(mine string) *SubscriptionsListCall {
	c.opt_["mine"] = mine
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The authenticated user acts on behalf of
// this content owner.
func (c *SubscriptionsListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *SubscriptionsListCall {
	c.opt_["onBehalfOfContentOwner"] = onBehalfOfContentOwner
	return c
}

// Order sets the optional parameter "order": Sort order.
func (c *SubscriptionsListCall) Order(order string) *SubscriptionsListCall {
	c.opt_["order"] = order
	return c
}

// PageToken sets the optional parameter "pageToken": Token for the page
// selection.
func (c *SubscriptionsListCall) PageToken(pageToken string) *SubscriptionsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

func (c *SubscriptionsListCall) Do() (*SubscriptionListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("part", fmt.Sprintf("%v", c.part))
	if v, ok := c.opt_["channelId"]; ok {
		params.Set("channelId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["forChannelId"]; ok {
		params.Set("forChannelId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["id"]; ok {
		params.Set("id", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["mine"]; ok {
		params.Set("mine", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["onBehalfOfContentOwner"]; ok {
		params.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["order"]; ok {
		params.Set("order", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "subscriptions")
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
	ret := new(SubscriptionListResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Browses the subscriptions collection.",
	//   "httpMethod": "GET",
	//   "id": "youtube.subscriptions.list",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "channelId": {
	//       "description": "Only return subscriptions to given channelId.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "forChannelId": {
	//       "description": "Comma separated list of channel IDs to return subscriptions for.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "id": {
	//       "description": "YouTube IDs of the subscriptions to be returned.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "5",
	//       "description": "Maximum number of search results to return per page.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "50",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "mine": {
	//       "description": "Flag indicating only return the subscriptions of the authenticated user.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The authenticated user acts on behalf of this content owner.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "order": {
	//       "default": "SUBSCRIPTION_ORDER_RELEVANCE",
	//       "description": "Sort order.",
	//       "enum": [
	//         "alphabetical",
	//         "relevance",
	//         "unread"
	//       ],
	//       "enumDescriptions": [
	//         "Sort alphabetically",
	//         "Sort by relevance.",
	//         "Sort by order of activity."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Token for the page selection.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "Subscription parts to include in the returned response. Valid values are: id, snippet and contentDetails.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "subscriptions",
	//   "response": {
	//     "$ref": "SubscriptionListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtube.videoCategories.list":

type VideoCategoriesListCall struct {
	s    *Service
	part string
	opt_ map[string]interface{}
}

// List: Browse list of video categories.
func (r *VideoCategoriesService) List(part string) *VideoCategoriesListCall {
	c := &VideoCategoriesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	return c
}

// Hl sets the optional parameter "hl": Language used for the title of
// the categories.
func (c *VideoCategoriesListCall) Hl(hl string) *VideoCategoriesListCall {
	c.opt_["hl"] = hl
	return c
}

// Id sets the optional parameter "id": IDs of the categories to be
// returned.
func (c *VideoCategoriesListCall) Id(id string) *VideoCategoriesListCall {
	c.opt_["id"] = id
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The authenticated user acts on behalf of
// this content owner.
func (c *VideoCategoriesListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *VideoCategoriesListCall {
	c.opt_["onBehalfOfContentOwner"] = onBehalfOfContentOwner
	return c
}

// RegionCode sets the optional parameter "regionCode": Return all the
// categories in this region.
func (c *VideoCategoriesListCall) RegionCode(regionCode string) *VideoCategoriesListCall {
	c.opt_["regionCode"] = regionCode
	return c
}

func (c *VideoCategoriesListCall) Do() (*VideoCategoryListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("part", fmt.Sprintf("%v", c.part))
	if v, ok := c.opt_["hl"]; ok {
		params.Set("hl", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["id"]; ok {
		params.Set("id", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["onBehalfOfContentOwner"]; ok {
		params.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["regionCode"]; ok {
		params.Set("regionCode", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "videoCategories")
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
	ret := new(VideoCategoryListResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Browse list of video categories.",
	//   "httpMethod": "GET",
	//   "id": "youtube.videoCategories.list",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "hl": {
	//       "default": "en_US",
	//       "description": "Language used for the title of the categories.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "id": {
	//       "description": "IDs of the categories to be returned.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The authenticated user acts on behalf of this content owner.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "Video category parts to include in the returned response. Valid values are: id and snippet.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "regionCode": {
	//       "description": "Return all the categories in this region.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "videoCategories",
	//   "response": {
	//     "$ref": "VideoCategoryListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtube.readonly",
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtube.videos.delete":

type VideosDeleteCall struct {
	s    *Service
	id   string
	opt_ map[string]interface{}
}

// Delete: Delete a YouTube video.
func (r *VideosService) Delete(id string) *VideosDeleteCall {
	c := &VideosDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The authenticated user acts on behalf of
// this content owner.
func (c *VideosDeleteCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *VideosDeleteCall {
	c.opt_["onBehalfOfContentOwner"] = onBehalfOfContentOwner
	return c
}

func (c *VideosDeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("id", fmt.Sprintf("%v", c.id))
	if v, ok := c.opt_["onBehalfOfContentOwner"]; ok {
		params.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "videos")
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
	//   "description": "Delete a YouTube video.",
	//   "httpMethod": "DELETE",
	//   "id": "youtube.videos.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "YouTube ID of the video to be deleted.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The authenticated user acts on behalf of this content owner.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "videos",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtube.videos.insert":

type VideosInsertCall struct {
	s      *Service
	part   string
	video  *Video
	opt_   map[string]interface{}
	media_ io.Reader
}

// Insert: Upload a video to YouTube.
func (r *VideosService) Insert(part string, video *Video) *VideosInsertCall {
	c := &VideosInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	c.video = video
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The authenticated user acts on behalf of
// this content owner.
func (c *VideosInsertCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *VideosInsertCall {
	c.opt_["onBehalfOfContentOwner"] = onBehalfOfContentOwner
	return c
}
func (c *VideosInsertCall) Media(r io.Reader) *VideosInsertCall {
	c.media_ = r
	return c
}

func (c *VideosInsertCall) Do() (*Video, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.video)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("part", fmt.Sprintf("%v", c.part))
	if v, ok := c.opt_["onBehalfOfContentOwner"]; ok {
		params.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "videos")
	if c.media_ != nil {
		urls = strings.Replace(urls, "https://www.googleapis.com/", "https://www.googleapis.com/upload/", 1)
		params.Set("uploadType", "multipart")
	}
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
	ret := new(Video)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Upload a video to YouTube.",
	//   "httpMethod": "POST",
	//   "id": "youtube.videos.insert",
	//   "mediaUpload": {
	//     "accept": [
	//       "application/octet-stream",
	//       "video/*"
	//     ],
	//     "maxSize": "64GB",
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/youtube/v3/videos"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/youtube/v3/videos"
	//       }
	//     }
	//   },
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The authenticated user acts on behalf of this content owner.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "One or more parts to return on the current request.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "videos",
	//   "request": {
	//     "$ref": "Video"
	//   },
	//   "response": {
	//     "$ref": "Video"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtube.upload"
	//   ],
	//   "supportsMediaUpload": true
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

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The authenticated user acts on behalf of
// this content owner.
func (c *VideosListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *VideosListCall {
	c.opt_["onBehalfOfContentOwner"] = onBehalfOfContentOwner
	return c
}

func (c *VideosListCall) Do() (*VideoListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("id", fmt.Sprintf("%v", c.id))
	params.Set("part", fmt.Sprintf("%v", c.part))
	if v, ok := c.opt_["onBehalfOfContentOwner"]; ok {
		params.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "videos")
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
	//     "onBehalfOfContentOwner": {
	//       "description": "The authenticated user acts on behalf of this content owner.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "Video parts to include in the returned response. Valid values are: id, snippet, contentDetails, player, statistics, status and topicDetails.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "videos",
	//   "response": {
	//     "$ref": "VideoListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtube.readonly",
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtube.videos.update":

type VideosUpdateCall struct {
	s     *Service
	part  string
	video *Video
	opt_  map[string]interface{}
}

// Update: Update a video.
func (r *VideosService) Update(part string, video *Video) *VideosUpdateCall {
	c := &VideosUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	c.video = video
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The authenticated user acts on behalf of
// this content owner.
func (c *VideosUpdateCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *VideosUpdateCall {
	c.opt_["onBehalfOfContentOwner"] = onBehalfOfContentOwner
	return c
}

func (c *VideosUpdateCall) Do() (*Video, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.video)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("part", fmt.Sprintf("%v", c.part))
	if v, ok := c.opt_["onBehalfOfContentOwner"]; ok {
		params.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "videos")
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
	ret := new(Video)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update a video.",
	//   "httpMethod": "PUT",
	//   "id": "youtube.videos.update",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The authenticated user acts on behalf of this content owner.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "One or more parts to return on the current request.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "videos",
	//   "request": {
	//     "$ref": "Video"
	//   },
	//   "response": {
	//     "$ref": "Video"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtubepartner"
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
