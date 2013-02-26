// Package youtube provides access to the YouTube Data API.
//
// See https://developers.google.com/youtube/v3
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
	s.LiveBroadcasts = &LiveBroadcastsService{s: s}
	s.LiveStreams = &LiveStreamsService{s: s}
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

	LiveBroadcasts *LiveBroadcastsService

	LiveStreams *LiveStreamsService

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

type LiveBroadcastsService struct {
	s *Service
}

type LiveStreamsService struct {
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

type AccessPolicy struct {
	// Allowed: The value of allowed indicates whether the access to the
	// policy is allowed or denied by default.
	Allowed bool `json:"allowed,omitempty"`

	// Exception: A list of region codes that identify countries where the
	// default policy do not apply.
	Exception []string `json:"exception,omitempty"`
}

type Activity struct {
	// ContentDetails: The contentDetails object contains information about
	// the content associated with the activity. For example, if the
	// snippet.type value is videoRated, then the contentDetails object's
	// content identifies the rated video.
	ContentDetails *ActivityContentDetails `json:"contentDetails,omitempty"`

	// Etag: The ETag of the activity resource.
	Etag string `json:"etag,omitempty"`

	// Id: The ID that YouTube uses to uniquely identify the activity.
	Id string `json:"id,omitempty"`

	// Kind: The type of the API resource. For activity resources, the value
	// will be youtube#activity.
	Kind string `json:"kind,omitempty"`

	// Snippet: The snippet object contains basic details about the
	// activity, including the activity's type and group ID.
	Snippet *ActivitySnippet `json:"snippet,omitempty"`
}

type ActivityContentDetails struct {
	// Bulletin: The bulletin object contains details about a channel
	// bulletin post. This object is only present if the snippet.type is
	// bulletin.
	Bulletin *ActivityContentDetailsBulletin `json:"bulletin,omitempty"`

	// ChannelItem: The channelItem object contains details about a resource
	// which was added to a channel. This property is only present if the
	// snippet.type is channelItem.
	ChannelItem *ActivityContentDetailsChannelItem `json:"channelItem,omitempty"`

	// Comment: The comment object contains information about a resource
	// that received a comment. This property is only present if the
	// snippet.type is comment.
	Comment *ActivityContentDetailsComment `json:"comment,omitempty"`

	// Favorite: The favorite object contains information about a video that
	// was marked as a favorite video. This property is only present if the
	// snippet.type is favorite.
	Favorite *ActivityContentDetailsFavorite `json:"favorite,omitempty"`

	// Like: The like object contains information about a resource that
	// received a positive (like) rating. This property is only present if
	// the snippet.type is like.
	Like *ActivityContentDetailsLike `json:"like,omitempty"`

	// PlaylistItem: The playlistItem object contains information about a
	// new playlist item. This property is only present if the snippet.type
	// is playlistItem.
	PlaylistItem *ActivityContentDetailsPlaylistItem `json:"playlistItem,omitempty"`

	// Recommendation: The recommendation object contains information about
	// a recommended resource. This property is only present if the
	// snippet.type is recommendation.
	Recommendation *ActivityContentDetailsRecommendation `json:"recommendation,omitempty"`

	// Social: The social object contains details about a social network
	// post. This property is only present if the snippet.type is social.
	Social *ActivityContentDetailsSocial `json:"social,omitempty"`

	// Subscription: The subscription object contains information about a
	// channel that a user subscribed to. This property is only present if
	// the snippet.type is subscription.
	Subscription *ActivityContentDetailsSubscription `json:"subscription,omitempty"`

	// Upload: The upload object contains information about the uploaded
	// video. This property is only present if the snippet.type is upload.
	Upload *ActivityContentDetailsUpload `json:"upload,omitempty"`
}

type ActivityContentDetailsBulletin struct {
	// ResourceId: The resourceId object contains information that
	// identifies the resource associated with a bulletin post.
	ResourceId *ResourceId `json:"resourceId,omitempty"`
}

type ActivityContentDetailsChannelItem struct {
	// ResourceId: The resourceId object contains information that
	// identifies the resource that was added to the channel.
	ResourceId *ResourceId `json:"resourceId,omitempty"`
}

type ActivityContentDetailsComment struct {
	// ResourceId: The resourceId object contains information that
	// identifies the resource associated with the comment.
	ResourceId *ResourceId `json:"resourceId,omitempty"`
}

type ActivityContentDetailsFavorite struct {
	// ResourceId: The resourceId object contains information that
	// identifies the resource that was marked as a favorite.
	ResourceId *ResourceId `json:"resourceId,omitempty"`
}

type ActivityContentDetailsLike struct {
	// ResourceId: The resourceId object contains information that
	// identifies the rated resource.
	ResourceId *ResourceId `json:"resourceId,omitempty"`
}

type ActivityContentDetailsPlaylistItem struct {
	// PlaylistId: The value that YouTube uses to uniquely identify the
	// playlist.
	PlaylistId string `json:"playlistId,omitempty"`

	// PlaylistItemId: ID of the item within the playlist.
	PlaylistItemId string `json:"playlistItemId,omitempty"`

	// ResourceId: The resourceId object contains information about the
	// resource that was added to the playlist.
	ResourceId *ResourceId `json:"resourceId,omitempty"`
}

type ActivityContentDetailsRecommendation struct {
	// Reason: The reason that the resource is recommended to the user.
	Reason string `json:"reason,omitempty"`

	// ResourceId: The resourceId object contains information that
	// identifies the recommended resource.
	ResourceId *ResourceId `json:"resourceId,omitempty"`

	// SeedResourceId: The seedResourceId object contains information about
	// the resource that caused the recommendation.
	SeedResourceId *ResourceId `json:"seedResourceId,omitempty"`
}

type ActivityContentDetailsSocial struct {
	// Author: The author of the social network post.
	Author string `json:"author,omitempty"`

	// ImageUrl: An image of the post's author.
	ImageUrl string `json:"imageUrl,omitempty"`

	// ReferenceUrl: The URL of the social network post.
	ReferenceUrl string `json:"referenceUrl,omitempty"`

	// ResourceId: The resourceId object encapsulates information that
	// identifies the resource associated with a social network post.
	ResourceId *ResourceId `json:"resourceId,omitempty"`

	// Type: The name of the social network.
	Type string `json:"type,omitempty"`
}

type ActivityContentDetailsSubscription struct {
	// ResourceId: The resourceId object contains information that
	// identifies the resource that the user subscribed to.
	ResourceId *ResourceId `json:"resourceId,omitempty"`
}

type ActivityContentDetailsUpload struct {
	// VideoId: The ID that YouTube uses to uniquely identify the uploaded
	// video.
	VideoId string `json:"videoId,omitempty"`
}

type ActivityListResponse struct {
	// Etag: The ETag of the response.
	Etag string `json:"etag,omitempty"`

	// Items: A list of activities, or events, that match the request
	// criteria.
	Items []*Activity `json:"items,omitempty"`

	// Kind: The type of the API response. For this operation, the value
	// will be youtube#activityListResponse.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The token that can be used as the value of the
	// pageToken parameter to retrieve the next page in the result set.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PageInfo: The pageInfo object encapsulates paging information for the
	// result set.
	PageInfo *PageInfo `json:"pageInfo,omitempty"`

	// PrevPageToken: The token that can be used as the value of the
	// pageToken parameter to retrieve the previous page in the result set.
	PrevPageToken string `json:"prevPageToken,omitempty"`
}

type ActivitySnippet struct {
	// ChannelId: The ID that YouTube uses to uniquely identify the channel
	// associated with the activity.
	ChannelId string `json:"channelId,omitempty"`

	// Description: The description of the resource primarily associated
	// with the activity.
	Description string `json:"description,omitempty"`

	// GroupId: The group ID associated with the activity. A group ID
	// identifies user events that are associated with the same user and
	// resource. For example, if a user rates a video and marks the same
	// video as a favorite, the entries for those events would have the same
	// group ID in the user's activity feed. In your user interface, you can
	// avoid repetition by grouping events with the same groupId value.
	GroupId string `json:"groupId,omitempty"`

	// PublishedAt: The date and time that the activity occurred. The value
	// is specified in ISO 8601 (YYYY-MM-DDThh:mm:ss.sZ) format.
	PublishedAt string `json:"publishedAt,omitempty"`

	// Thumbnails: A map of thumbnail images associated with the resource
	// that is primarily associated with the activity. For each object in
	// the map, the key is the name of the thumbnail image, and the value is
	// an object that contains other information about the thumbnail.
	Thumbnails *ActivitySnippetThumbnails `json:"thumbnails,omitempty"`

	// Title: The title of the resource primarily associated with the
	// activity.
	Title string `json:"title,omitempty"`

	// Type: The type of activity that the resource describes.
	Type string `json:"type,omitempty"`
}

type ActivitySnippetThumbnails struct {
}

type Channel struct {
	// ContentDetails: The contentDetails object encapsulates information
	// about the channel's content.
	ContentDetails *ChannelContentDetails `json:"contentDetails,omitempty"`

	// Etag: The ETag for the channel resource.
	Etag string `json:"etag,omitempty"`

	// Id: The ID that YouTube uses to uniquely identify the channel.
	Id string `json:"id,omitempty"`

	// Kind: The type of the API resource. For channel resources, the value
	// will be youtube#channel.
	Kind string `json:"kind,omitempty"`

	// Snippet: The snippet object contains basic details about the channel,
	// such as its title, description, and thumbnail images.
	Snippet *ChannelSnippet `json:"snippet,omitempty"`

	// Statistics: The statistics object encapsulates statistics for the
	// channel.
	Statistics *ChannelStatistics `json:"statistics,omitempty"`

	// Status: The status object encapsulates information about the privacy
	// status of the channel.
	Status *ChannelStatus `json:"status,omitempty"`

	// TopicDetails: The topicDetails object encapsulates information about
	// Freebase topics associated with the channel.
	TopicDetails *ChannelTopicDetails `json:"topicDetails,omitempty"`
}

type ChannelContentDetails struct {
	// GooglePlusUserId: The googlePlusUserId object identifies the Google+
	// profile ID associated with this channel.
	GooglePlusUserId string `json:"googlePlusUserId,omitempty"`

	// RelatedPlaylists: The relatedPlaylists object is a map that
	// identifies playlists associated with the channel, such as the
	// channel's uploaded videos or favorite videos. You can retrieve any of
	// these playlists using the playlists.list method.
	RelatedPlaylists *ChannelContentDetailsRelatedPlaylists `json:"relatedPlaylists,omitempty"`
}

type ChannelContentDetailsRelatedPlaylists struct {
	// Favorites: The ID of the playlist that contains the channel's
	// favorite videos. Use the playlistItems.insert and
	// playlistItems.delete to add or remove items from that list.
	Favorites string `json:"favorites,omitempty"`

	// Likes: The ID of the playlist that contains the channel's liked
	// videos. Use the playlistItems.insert and playlistItems.delete to add
	// or remove items from that list.
	Likes string `json:"likes,omitempty"`

	// Uploads: The ID of the playlist that contains the channel's uploaded
	// videos. Use the videos.insert method to upload new videos and the
	// videos.delete method to delete previously uploaded videos.
	Uploads string `json:"uploads,omitempty"`

	// WatchHistory: The ID of the playlist that contains the channel's
	// watch history. Use the playlistItems.insert and playlistItems.delete
	// to add or remove items from that list.
	WatchHistory string `json:"watchHistory,omitempty"`

	// WatchLater: The ID of the channel's watch later playlist. Use the
	// playlistItems.insert and playlistItems.delete to add or remove items
	// from that list.
	WatchLater string `json:"watchLater,omitempty"`
}

type ChannelListResponse struct {
	// Etag: The ETag for the response.
	Etag string `json:"etag,omitempty"`

	// Items: A list of channels that match the request criteria.
	Items []*Channel `json:"items,omitempty"`

	// Kind: The type of the API response. For this operation, the value
	// will be youtube#channelListResponse.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The token that can be used as the value of the
	// pageToken parameter to retrieve the next page in the result set.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PageInfo: The pageInfo object encapsulates paging information for the
	// result set.
	PageInfo *PageInfo `json:"pageInfo,omitempty"`

	// PrevPageToken: The token that can be used as the value of the
	// pageToken parameter to retrieve the previous page in the result set.
	PrevPageToken string `json:"prevPageToken,omitempty"`
}

type ChannelSnippet struct {
	// Description: The description of the channel.
	Description string `json:"description,omitempty"`

	// PublishedAt: The date and time that the channel was created. The
	// value is specified in ISO 8601 (YYYY-MM-DDThh:mm:ss.sZ) format.
	PublishedAt string `json:"publishedAt,omitempty"`

	// Thumbnails: A map of thumbnail images associated with the channel.
	// For each object in the map, the key is the name of the thumbnail
	// image, and the value is an object that contains other information
	// about the thumbnail.
	Thumbnails *ChannelSnippetThumbnails `json:"thumbnails,omitempty"`

	// Title: The channel's title.
	Title string `json:"title,omitempty"`
}

type ChannelSnippetThumbnails struct {
}

type ChannelStatistics struct {
	// CommentCount: The number of comments for the channel.
	CommentCount uint64 `json:"commentCount,omitempty,string"`

	// SubscriberCount: The number of subscribers that the channel has.
	SubscriberCount uint64 `json:"subscriberCount,omitempty,string"`

	// VideoCount: The number of videos uploaded to the channel.
	VideoCount uint64 `json:"videoCount,omitempty,string"`

	// ViewCount: The number of times the channel has been viewed.
	ViewCount uint64 `json:"viewCount,omitempty,string"`
}

type ChannelStatus struct {
	// PrivacyStatus: Privacy status of the channel.
	PrivacyStatus string `json:"privacyStatus,omitempty"`
}

type ChannelTopicDetails struct {
	// TopicIds: A list of Freebase topic IDs associated with the channel.
	// You can retrieve information about each topic using the Freebase
	// Topic API.
	TopicIds []string `json:"topicIds,omitempty"`
}

type GeoPoint struct {
	// Elevation: Altitude above the Earth, in meters.
	Elevation float64 `json:"elevation,omitempty"`

	// Latitude: Latitude in degrees.
	Latitude float64 `json:"latitude,omitempty"`

	// Longitude: Longitude in degrees.
	Longitude float64 `json:"longitude,omitempty"`
}

type GuideCategory struct {
	// Etag: The ETag of the guideCategory resource.
	Etag string `json:"etag,omitempty"`

	// Id: The ID that YouTube uses to uniquely identify the guide category.
	Id string `json:"id,omitempty"`

	// Kind: The type of the API resource. For guideCategory resources, the
	// value will be youtube#guideCategory/code>.
	Kind string `json:"kind,omitempty"`

	// Snippet: The snippet object contains basic details about the
	// category, such as its title.
	Snippet *GuideCategorySnippet `json:"snippet,omitempty"`
}

type GuideCategoryListResponse struct {
	// Etag: The ETag of the response.
	Etag string `json:"etag,omitempty"`

	// Items: A list of categories that can be associated with YouTube
	// channels. In this map, the category ID is the map key, and its value
	// is the corresponding guideCategory resource.
	Items []*GuideCategory `json:"items,omitempty"`

	// Kind: The type of the API response. For this operation, the value
	// will be youtube#guideCategoryListResponse.
	Kind string `json:"kind,omitempty"`
}

type GuideCategorySnippet struct {
	// ChannelId: The ID that YouTube uses to uniquely identify the channel
	// publishing the guide category.
	ChannelId string `json:"channelId,omitempty"`

	// Title: The category's title.
	Title string `json:"title,omitempty"`
}

type LiveBroadcast struct {
	// ContentDetails: The content details of the live broadcast.
	ContentDetails *LiveBroadcastContentDetails `json:"contentDetails,omitempty"`

	// Etag: The eTag of the broadcast.
	Etag string `json:"etag,omitempty"`

	// Id: The unique id of the broadcast.
	Id string `json:"id,omitempty"`

	// Kind: The type of this API resource.
	Kind string `json:"kind,omitempty"`

	// SlateSettings: The slate settings of the live broadcast.
	SlateSettings *LiveBroadcastSlateSettings `json:"slateSettings,omitempty"`

	// Snippet: Basic details about the live broadcast.
	Snippet *LiveBroadcastSnippet `json:"snippet,omitempty"`

	// Status: The status of the live broadcast.
	Status *LiveBroadcastStatus `json:"status,omitempty"`
}

type LiveBroadcastContentDetails struct {
	// BoundStreamId: The id of the stream bound to the broadcast.
	BoundStreamId string `json:"boundStreamId,omitempty"`

	// EnableArchive: Whether the live event will be archived or not.
	EnableArchive bool `json:"enableArchive,omitempty"`

	// EnableContentEncryption: Whether to enable or disable content
	// encryption.
	EnableContentEncryption bool `json:"enableContentEncryption,omitempty"`

	// EnableDvr: Whether the dvr (digital video recording) is enabled or
	// not.
	EnableDvr bool `json:"enableDvr,omitempty"`

	// EnableEmbed: Whether to allow the broadcast to be played in an
	// embedded player.
	EnableEmbed bool `json:"enableEmbed,omitempty"`

	// MonitorStream: Information about the monitor stream which helps the
	// broadcaster to review the event content before shown to the public.
	MonitorStream *LiveBroadcastContentDetailsMonitorStream `json:"monitorStream,omitempty"`

	// StartWithSlateCuepoint: Automatically start with a slate cuepoint.
	StartWithSlateCuepoint bool `json:"startWithSlateCuepoint,omitempty"`
}

type LiveBroadcastContentDetailsMonitorStream struct {
	// BroadcastStreamDelayMs: If enableMonitorStream is true, the public
	// broadcast will be delayed by this value.
	BroadcastStreamDelayMs int64 `json:"broadcastStreamDelayMs,omitempty"`

	// EmbedHtml: The html code of the embedded player for the monitor
	// stream.
	EmbedHtml string `json:"embedHtml,omitempty"`

	// EnableMonitorStream: Whether to enable the monitor stream for the
	// broadcast.
	EnableMonitorStream bool `json:"enableMonitorStream,omitempty"`
}

type LiveBroadcastList struct {
	// Etag: The eTag of the chart.
	Etag string `json:"etag,omitempty"`

	// Items: A list of broadcasts that match the request criteria.
	Items []*LiveBroadcast `json:"items,omitempty"`

	// Kind: The type of this API resource.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The token that can be used as the value of the {@code
	// pageInfo} parameter to retrieve the next page in the result set.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PageInfo: The {@code pageInfo} object encapsulates paging information
	// for the result set.
	PageInfo *PageInfo `json:"pageInfo,omitempty"`

	// PrevPageToken: The token that can be used as the value of the {@code
	// pageInfo} parameter to retrieve the previous page in the result set.
	PrevPageToken string `json:"prevPageToken,omitempty"`
}

type LiveBroadcastSlateSettings struct {
	// EnableSlates: Whether slate is enabled or not.
	EnableSlates bool `json:"enableSlates,omitempty"`

	// Slates: Broadcast slates.
	Slates *LiveBroadcastSlateSettingsSlates `json:"slates,omitempty"`
}

type LiveBroadcastSlateSettingsSlates struct {
}

type LiveBroadcastSnippet struct {
	// ActualEndTime: Date and time the broadcast is actual to end. The
	// value is specified in ISO 8601 (YYYY-MM-DDThh:mm:ss.sZ) format.
	ActualEndTime string `json:"actualEndTime,omitempty"`

	// ActualStartTime: Date and time the broadcast is actual to start. The
	// value is specified in ISO 8601 (YYYY-MM-DDThh:mm:ss.sZ) format.
	ActualStartTime string `json:"actualStartTime,omitempty"`

	// ChannelId: Channel publishing the broadcast.
	ChannelId string `json:"channelId,omitempty"`

	// Description: Description of the broadcast.
	Description string `json:"description,omitempty"`

	// PublishedAt: Date and time the broadcast was published at. The value
	// is specified in ISO 8601 (YYYY-MM-DDThh:mm:ss.sZ) format.
	PublishedAt string `json:"publishedAt,omitempty"`

	// ScheduledEndTime: Date and time the broadcast is scheduled to end.
	// The value is specified in ISO 8601 (YYYY-MM-DDThh:mm:ss.sZ) format.
	ScheduledEndTime string `json:"scheduledEndTime,omitempty"`

	// ScheduledStartTime: Date and time the broadcast is scheduled to
	// start. The value is specified in ISO 8601 (YYYY-MM-DDThh:mm:ss.sZ)
	// format.
	ScheduledStartTime string `json:"scheduledStartTime,omitempty"`

	// Thumbnails: Video thumbnails.
	Thumbnails *LiveBroadcastSnippetThumbnails `json:"thumbnails,omitempty"`

	// Title: Title of the broadcast.
	Title string `json:"title,omitempty"`
}

type LiveBroadcastSnippetThumbnails struct {
}

type LiveBroadcastStatus struct {
	// LifeCycleStatus: Life status of the live broadcast.
	LifeCycleStatus string `json:"lifeCycleStatus,omitempty"`

	// PrivacyStatus: Privacy settings of the live broadcast. Allowed
	// values: private, unlisted, public.
	PrivacyStatus string `json:"privacyStatus,omitempty"`
}

type LiveStream struct {
	// Cdn: Cdn settings of the live stream.
	Cdn *LiveStreamCdn `json:"cdn,omitempty"`

	// Etag: The eTag of the stream.
	Etag string `json:"etag,omitempty"`

	// Id: The unique id of the stream.
	Id string `json:"id,omitempty"`

	// Kind: The type of this API resource.
	Kind string `json:"kind,omitempty"`

	// Snippet: Basic details about the live stream.
	Snippet *LiveStreamSnippet `json:"snippet,omitempty"`

	// Status: Status of the live stream.
	Status *LiveStreamStatus `json:"status,omitempty"`
}

type LiveStreamCdn struct {
	// Format: The format of the inbound data. Allowed values: 240p, 360p,
	// 480p, 720p, 1080p, webm_360p, multicast_qcif, multicast_240p,
	// multicast_360p, multicast_480p, multicast_720p, multicast_1080p.
	Format string `json:"format,omitempty"`

	// IngestionInfo: Encapsulates information for ingesting an RTMP or an
	// HTTP stream.
	IngestionInfo *LiveStreamCdnIngestionInfo `json:"ingestionInfo,omitempty"`

	// IngestionType: The live stream ingestion type. Allowed values: rtmp,
	// http, multicast.
	IngestionType string `json:"ingestionType,omitempty"`

	// MulticastIngestionInfo: Encapsulates information for ingesting a
	// multicast stream.
	MulticastIngestionInfo *LiveStreamCdnMulticastIngestionInfo `json:"multicastIngestionInfo,omitempty"`
}

type LiveStreamCdnIngestionInfo struct {
	// BackupIngestionAddress: The backup address of the inbound data.
	BackupIngestionAddress string `json:"backupIngestionAddress,omitempty"`

	// IngestionAddress: The address of the inbound data.
	IngestionAddress string `json:"ingestionAddress,omitempty"`

	// StreamName: Ingestion stream name.
	StreamName string `json:"streamName,omitempty"`
}

type LiveStreamCdnMulticastIngestionInfo struct {
	// MulticastAddress: The IP address of the multicast data.
	MulticastAddress string `json:"multicastAddress,omitempty"`
}

type LiveStreamList struct {
	// Etag: The eTag of the chart.
	Etag string `json:"etag,omitempty"`

	// Items: A list of live streams that match the request criteria.
	Items []*LiveStream `json:"items,omitempty"`

	// Kind: The type of this API resource.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The token that can be used as the value of the {@code
	// pageInfo} parameter to retrieve the next page in the result set.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PageInfo: The {@code pageInfo} object encapsulates paging information
	// for the result set.
	PageInfo *PageInfo `json:"pageInfo,omitempty"`

	// PrevPageToken: The token that can be used as the value of the {@code
	// pageInfo} parameter to retrieve the previous page in the result set.
	PrevPageToken string `json:"prevPageToken,omitempty"`
}

type LiveStreamSnippet struct {
	// ChannelId: Channel publishing the live stream.
	ChannelId string `json:"channelId,omitempty"`

	// Description: Description of the live stream.
	Description string `json:"description,omitempty"`

	// PublishedAt: Date and time the live stream was published at.
	PublishedAt string `json:"publishedAt,omitempty"`

	// Title: Title of the live stream.
	Title string `json:"title,omitempty"`
}

type LiveStreamStatus struct {
	// StreamStatus: The status of the stream.
	StreamStatus string `json:"streamStatus,omitempty"`
}

type PageInfo struct {
	// ResultsPerPage: The number of results included in the API response.
	ResultsPerPage int64 `json:"resultsPerPage,omitempty"`

	// TotalResults: The total number of results in the result set.
	TotalResults int64 `json:"totalResults,omitempty"`
}

type Playlist struct {
	// ContentDetails: The contentDetails object contains information like
	// video count.
	ContentDetails *PlaylistContentDetails `json:"contentDetails,omitempty"`

	// Etag: The ETag for the playlist resource.
	Etag string `json:"etag,omitempty"`

	// Id: The ID that YouTube uses to uniquely identify the playlist.
	Id string `json:"id,omitempty"`

	// Kind: The type of the API resource. For video resources, the value
	// will be youtube#playlist.
	Kind string `json:"kind,omitempty"`

	// Player: The player object contains information that you would use to
	// play the playlist in an embedded player.
	Player *PlaylistPlayer `json:"player,omitempty"`

	// Snippet: The snippet object contains basic details about the
	// playlist, such as its title and description.
	Snippet *PlaylistSnippet `json:"snippet,omitempty"`

	// Status: The status object contains status information for the
	// playlist.
	Status *PlaylistStatus `json:"status,omitempty"`
}

type PlaylistContentDetails struct {
	// ItemCount: The number of videos in the playlist.
	ItemCount int64 `json:"itemCount,omitempty"`
}

type PlaylistItem struct {
	// ContentDetails: The contentDetails object is included in the resource
	// if the included item is a YouTube video. The object contains
	// additional information about the video.
	ContentDetails *PlaylistItemContentDetails `json:"contentDetails,omitempty"`

	// Etag: The ETag for the playlist item resource.
	Etag string `json:"etag,omitempty"`

	// Id: The ID that YouTube uses to uniquely identify the playlist item.
	Id string `json:"id,omitempty"`

	// Kind: The type of the API resource. For playlist item resources, the
	// value will be youtube#playlistItem.
	Kind string `json:"kind,omitempty"`

	// Snippet: The snippet object contains basic details about the playlist
	// item, such as its title and position in the playlist.
	Snippet *PlaylistItemSnippet `json:"snippet,omitempty"`
}

type PlaylistItemContentDetails struct {
	// EndAt: The time, measured in seconds from the start of the video,
	// when the video should stop playing. (The playlist owner can specify
	// the times when the video should start and stop playing when the video
	// is played in the context of the playlist.) By default, assume that
	// the video.endTime is the end of the video.
	EndAt string `json:"endAt,omitempty"`

	// Note: A user-generated note for this item.
	Note string `json:"note,omitempty"`

	// StartAt: The time, measured in seconds from the start of the video,
	// when the video should start playing. (The playlist owner can specify
	// the times when the video should start and stop playing when the video
	// is played in the context of the playlist.) The default value is 0.
	StartAt string `json:"startAt,omitempty"`

	// VideoId: The ID that YouTube uses to uniquely identify a video. To
	// retrieve the video resource, set the id query parameter to this value
	// in your API request.
	VideoId string `json:"videoId,omitempty"`
}

type PlaylistItemListResponse struct {
	// Etag: The ETag for the response.
	Etag string `json:"etag,omitempty"`

	// Items: A list of playlist items that match the request criteria.
	Items []*PlaylistItem `json:"items,omitempty"`

	// Kind: The type of the API response. For this operation, the value
	// will be youtube#playlistItemListResponse.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: A token that can be used as the value of the pageToken
	// parameter to retrieve the next page in the result set.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PageInfo: The pageInfo object encapsulates paging information for the
	// result set.
	PageInfo *PageInfo `json:"pageInfo,omitempty"`

	// PrevPageToken: A token that can be used as the value of the pageToken
	// parameter to retrieve the previous page in the result set.
	PrevPageToken string `json:"prevPageToken,omitempty"`
}

type PlaylistItemSnippet struct {
	// ChannelId: The ID that YouTube uses to uniquely identify the user
	// that added the item to the playlist.
	ChannelId string `json:"channelId,omitempty"`

	// Description: The item's description.
	Description string `json:"description,omitempty"`

	// PlaylistId: The ID that YouTube uses to uniquely identify the
	// playlist that the playlist item is in.
	PlaylistId string `json:"playlistId,omitempty"`

	// Position: The order in which the item appears in the playlist. The
	// value uses a zero-based index, so the first item has a position of 0,
	// the second item has a position of 1, and so forth.
	Position int64 `json:"position,omitempty"`

	// PublishedAt: The date and time that the item was added to the
	// playlist. The value is specified in ISO 8601 (YYYY-MM-DDThh:mm:ss.sZ)
	// format.
	PublishedAt string `json:"publishedAt,omitempty"`

	// ResourceId: The id object contains information that can be used to
	// uniquely identify the resource that is included in the playlist as
	// the playlist item.
	ResourceId *ResourceId `json:"resourceId,omitempty"`

	// Thumbnails: A map of thumbnail images associated with the playlist
	// item. For each object in the map, the key is the name of the
	// thumbnail image, and the value is an object that contains other
	// information about the thumbnail.
	Thumbnails *PlaylistItemSnippetThumbnails `json:"thumbnails,omitempty"`

	// Title: The item's title.
	Title string `json:"title,omitempty"`
}

type PlaylistItemSnippetThumbnails struct {
}

type PlaylistListResponse struct {
	// Etag: The ETag of the response.
	Etag string `json:"etag,omitempty"`

	// Items: A list of playlists that match the request criteria.
	Items []*Playlist `json:"items,omitempty"`

	// Kind: The type of the API response. For this operation, the value
	// will be youtube#playlistListResponse.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The token that can be used as the value of the
	// pageToken parameter to retrieve the next page in the result set.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PageInfo: The pageInfo object encapsulates paging information for the
	// result set.
	PageInfo *PageInfo `json:"pageInfo,omitempty"`

	// PrevPageToken: The token that can be used as the value of the
	// pageToken parameter to retrieve the previous page in the result set.
	PrevPageToken string `json:"prevPageToken,omitempty"`
}

type PlaylistPlayer struct {
	// EmbedHtml: An <iframe> tag that embeds a player that will play the
	// playlist.
	EmbedHtml string `json:"embedHtml,omitempty"`
}

type PlaylistSnippet struct {
	// ChannelId: The ID that YouTube uses to uniquely identify the channel
	// that published the playlist.
	ChannelId string `json:"channelId,omitempty"`

	// Description: The playlist's description.
	Description string `json:"description,omitempty"`

	// PublishedAt: The date and time that the playlist was created. The
	// value is specified in ISO 8601 (YYYY-MM-DDThh:mm:ss.sZ) format.
	PublishedAt string `json:"publishedAt,omitempty"`

	// Thumbnails: A map of thumbnail images associated with the playlist.
	// For each object in the map, the key is the name of the thumbnail
	// image, and the value is an object that contains other information
	// about the thumbnail.
	Thumbnails *PlaylistSnippetThumbnails `json:"thumbnails,omitempty"`

	// Title: The playlist's title.
	Title string `json:"title,omitempty"`
}

type PlaylistSnippetThumbnails struct {
}

type PlaylistStatus struct {
	// PrivacyStatus: The playlist's privacy status.
	PrivacyStatus string `json:"privacyStatus,omitempty"`
}

type ResourceId struct {
	// ChannelId: The ID that YouTube uses to uniquely identify the referred
	// resource, if that resource is a channel. This property is only
	// present if the resourceId.kind value is youtube#channel.
	ChannelId string `json:"channelId,omitempty"`

	// Kind: The kind, or type, of the referred resource.
	Kind string `json:"kind,omitempty"`

	// PlaylistId: The ID that YouTube uses to uniquely identify the
	// referred resource, if that resource is a playlist. This property is
	// only present if the resourceId.kind value is youtube#playlist.
	PlaylistId string `json:"playlistId,omitempty"`

	// VideoId: The ID that YouTube uses to uniquely identify the referred
	// resource, if that resource is a video. This property is only present
	// if the resourceId.kind value is youtube#video.
	VideoId string `json:"videoId,omitempty"`
}

type SearchListResponse struct {
	// Etag: The ETag for the response.
	Etag string `json:"etag,omitempty"`

	// Items: A list of results that match the search criteria.
	Items []*SearchResult `json:"items,omitempty"`

	// Kind: The type of the API response. For this operation, the value
	// will be youtube#searchListResponse.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The token that can be used as the value of the
	// pageToken parameter to retrieve the next page in the result set.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PageInfo: The pageInfo object encapsulates paging information for the
	// search result set.
	PageInfo *PageInfo `json:"pageInfo,omitempty"`

	// PrevPageToken: The token that can be used as the value of the
	// pageToken parameter to retrieve the previous page in the result set.
	PrevPageToken string `json:"prevPageToken,omitempty"`
}

type SearchResult struct {
	// Etag: The ETag of the search result.
	Etag string `json:"etag,omitempty"`

	// Id: The id object contains information that can be used to uniquely
	// identify the resource that matches the search request.
	Id *ResourceId `json:"id,omitempty"`

	// Kind: The type of the API response. For this resource, the value will
	// be youtube#searchResult.
	Kind string `json:"kind,omitempty"`

	// Snippet: The snippet object contains basic details about a search
	// result, such as its title or description. For example, if the search
	// result is a video, then the title will be the video's title and the
	// description will be the video's description.
	Snippet *SearchResultSnippet `json:"snippet,omitempty"`
}

type SearchResultSnippet struct {
	// ChannelId: The value that YouTube uses to uniquely identify the
	// channel that published the resource that the search result
	// identifies.
	ChannelId string `json:"channelId,omitempty"`

	// Description: A description of the search result.
	Description string `json:"description,omitempty"`

	// PublishedAt: The creation date and time of the resource that the
	// search result identifies. The value is specified in ISO 8601
	// (YYYY-MM-DDThh:mm:ss.sZ) format.
	PublishedAt string `json:"publishedAt,omitempty"`

	// Thumbnails: A map of thumbnail images associated with the search
	// result. For each object in the map, the key is the name of the
	// thumbnail image, and the value is an object that contains other
	// information about the thumbnail.
	Thumbnails *SearchResultSnippetThumbnails `json:"thumbnails,omitempty"`

	// Title: The title to display for the search result.
	Title string `json:"title,omitempty"`
}

type SearchResultSnippetThumbnails struct {
}

type Subscription struct {
	// ContentDetails: The contentDetails object contains basic statistics
	// about the subscription.
	ContentDetails *SubscriptionContentDetails `json:"contentDetails,omitempty"`

	// Etag: The ETag of the subscription resource.
	Etag string `json:"etag,omitempty"`

	// Id: The ID that YouTube uses to uniquely identify the subscription.
	Id string `json:"id,omitempty"`

	// Kind: The type of the API resource. For subscription resources, the
	// value will be youtube#subscription.
	Kind string `json:"kind,omitempty"`

	// Snippet: The snippet object contains basic details about the
	// subscription, including its title and the channel that the user
	// subscribed to.
	Snippet *SubscriptionSnippet `json:"snippet,omitempty"`
}

type SubscriptionContentDetails struct {
	// NewItemCount: The number of new items in the subscription since its
	// content was last read.
	NewItemCount int64 `json:"newItemCount,omitempty"`

	// TotalItemCount: The approximate number of items that the subscription
	// points to.
	TotalItemCount int64 `json:"totalItemCount,omitempty"`
}

type SubscriptionListResponse struct {
	// Etag: The ETag of the response.
	Etag string `json:"etag,omitempty"`

	// Items: A list of subscriptions that match the request criteria.
	Items []*Subscription `json:"items,omitempty"`

	// Kind: The type of the API response. For this operation, the value
	// will be youtube#subscriptionListResponse.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The token that can be used as the value of the
	// pageToken parameter to retrieve the next page in the result set.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PageInfo: The pageInfo object encapsulates paging information for the
	// result set.
	PageInfo *PageInfo `json:"pageInfo,omitempty"`

	// PrevPageToken: The token that can be used as the value of the
	// pageToken parameter to retrieve the previous page in the result set.
	PrevPageToken string `json:"prevPageToken,omitempty"`
}

type SubscriptionSnippet struct {
	// ChannelId: The ID that YouTube uses to uniquely identify the
	// subscriber's channel.
	ChannelId string `json:"channelId,omitempty"`

	// Description: The subscription's details.
	Description string `json:"description,omitempty"`

	// PublishedAt: The date and time that the subscription was created. The
	// value is specified in ISO 8601 (YYYY-MM-DDThh:mm:ss.sZ) format.
	PublishedAt string `json:"publishedAt,omitempty"`

	// ResourceId: The id object contains information about the channel that
	// the user subscribed to.
	ResourceId *ResourceId `json:"resourceId,omitempty"`

	// Thumbnails: A map of thumbnail images associated with the
	// subscription. For each object in the map, the key is the name of the
	// thumbnail image, and the value is an object that contains other
	// information about the thumbnail.
	Thumbnails *SubscriptionSnippetThumbnails `json:"thumbnails,omitempty"`

	// Title: The subscription's title.
	Title string `json:"title,omitempty"`
}

type SubscriptionSnippetThumbnails struct {
}

type Thumbnail struct {
	// Height: (Optional) Height of the thumbnail image.
	Height int64 `json:"height,omitempty"`

	// Url: The thumbnail image's URL.
	Url string `json:"url,omitempty"`

	// Width: (Optional) Width of the thumbnail image.
	Width int64 `json:"width,omitempty"`
}

type Video struct {
	// ContentDetails: The contentDetails object contains information about
	// the video content, including the length of the video and its aspect
	// ratio.
	ContentDetails *VideoContentDetails `json:"contentDetails,omitempty"`

	// Etag: The ETag of the video resource.
	Etag string `json:"etag,omitempty"`

	// FileDetails: The fileDetails object encapsulates information about
	// the video file that was uploaded to YouTube, including the file's
	// resolution, duration, audio and video codecs, stream bitrates, and
	// more. This data can only be retrieved by the video owner.
	FileDetails *VideoFileDetails `json:"fileDetails,omitempty"`

	// Id: The ID that YouTube uses to uniquely identify the video.
	Id string `json:"id,omitempty"`

	// Kind: The type of the API resource. For video resources, the value
	// will be youtube#video.
	Kind string `json:"kind,omitempty"`

	// MonetizationDetails: The monetizationDetails object encapsulates
	// information about the monetization status of the video.
	MonetizationDetails *VideoMonetizationDetails `json:"monetizationDetails,omitempty"`

	// Player: The player object contains information that you would use to
	// play the video in an embedded player.
	Player *VideoPlayer `json:"player,omitempty"`

	// ProcessingDetails: The processingProgress object encapsulates
	// information about YouTube's progress in processing the uploaded video
	// file. The properties in the object identify the current processing
	// status and an estimate of the time remaining until YouTube finishes
	// processing the video. This part also indicates whether different
	// types of data or content, such as file details or thumbnail images,
	// are available for the video.
	//
	// The processingProgress object is
	// designed to be polled so that the video uploaded can track the
	// progress that YouTube has made in processing the uploaded video file.
	// This data can only be retrieved by the video owner.
	ProcessingDetails *VideoProcessingDetails `json:"processingDetails,omitempty"`

	// RecordingDetails: The recordingDetails object encapsulates
	// information about the location, date and address where the video was
	// recorded.
	RecordingDetails *VideoRecordingDetails `json:"recordingDetails,omitempty"`

	// Snippet: The snippet object contains basic details about the video,
	// such as its title, description, and category.
	Snippet *VideoSnippet `json:"snippet,omitempty"`

	// Statistics: The statistics object contains statistics about the
	// video.
	Statistics *VideoStatistics `json:"statistics,omitempty"`

	// Status: The status object contains information about the video's
	// uploading, processing, and privacy statuses.
	Status *VideoStatus `json:"status,omitempty"`

	// Suggestions: The suggestions object encapsulates suggestions that
	// identify opportunities to improve the video quality or the metadata
	// for the uploaded video. This data can only be retrieved by the video
	// owner.
	Suggestions *VideoSuggestions `json:"suggestions,omitempty"`

	// TopicDetails: The topicDetails object encapsulates information about
	// Freebase topics associated with the video.
	TopicDetails *VideoTopicDetails `json:"topicDetails,omitempty"`
}

type VideoCategory struct {
	// Etag: The ETag of the videoCategory resource.
	Etag string `json:"etag,omitempty"`

	// Id: The ID that YouTube uses to uniquely identify the video category.
	Id string `json:"id,omitempty"`

	// Kind: The type of the API resource. For video category resources, the
	// value will be youtube#videoCategory.
	Kind string `json:"kind,omitempty"`

	// Snippet: The snippet object contains basic details about the video
	// category, including its title.
	Snippet *VideoCategorySnippet `json:"snippet,omitempty"`
}

type VideoCategoryListResponse struct {
	// Etag: The ETag of the response.
	Etag string `json:"etag,omitempty"`

	// Items: A list of video categories that can be associated with YouTube
	// videos. In this map, the video category ID is the map key, and its
	// value is the corresponding videoCategory resource.
	Items []*VideoCategory `json:"items,omitempty"`

	// Kind: The type of the API response. For this operation, the value
	// will be youtube#videoCategoryListResponse.
	Kind string `json:"kind,omitempty"`
}

type VideoCategorySnippet struct {
	// ChannelId: The YouTube channel that created the video category.
	ChannelId string `json:"channelId,omitempty"`

	// Title: The video category's title.
	Title string `json:"title,omitempty"`
}

type VideoContentDetails struct {
	// Caption: The value of captions indicates whether the video has
	// captions or not.
	Caption string `json:"caption,omitempty"`

	// Definition: The value of definition indicates whether the video is
	// available in high definition or only in standard definition.
	Definition string `json:"definition,omitempty"`

	// Dimension: The value of dimension indicates whether the video is
	// available in 3D or in 2D.
	Dimension string `json:"dimension,omitempty"`

	// Duration: The length of the video. The tag value is an ISO 8601
	// duration in the format PT#M#S, in which the letters PT indicate that
	// the value specifies a period of time, and the letters M and S refer
	// to length in minutes and seconds, respectively. The # characters
	// preceding the M and S letters are both integers that specify the
	// number of minutes (or seconds) of the video. For example, a value of
	// PT15M51S indicates that the video is 15 minutes and 51 seconds long.
	Duration string `json:"duration,omitempty"`

	// LicensedContent: The value of is_license_content indicates whether
	// the video is licensed content.
	LicensedContent bool `json:"licensedContent,omitempty"`

	// RegionRestriction: The regionRestriction object contains information
	// about the countries where a video is (or is not) viewable. The object
	// will contain either the contentDetails.regionRestriction.allowed
	// property or the contentDetails.regionRestriction.blocked property.
	RegionRestriction *VideoContentDetailsRegionRestriction `json:"regionRestriction,omitempty"`
}

type VideoContentDetailsRegionRestriction struct {
	// Allowed: A list of region codes that identify countries where the
	// video is viewable. If this property is present and a country is not
	// listed in its value, then the video is blocked from appearing in that
	// country. If this property is present and contains an empty list, the
	// video is blocked in all countries.
	Allowed []string `json:"allowed,omitempty"`

	// Blocked: A list of region codes that identify countries where the
	// video is blocked. If this property is present and a country is not
	// listed in its value, then the video is viewable in that country. If
	// this property is present and contains an empty list, the video is
	// viewable in all countries.
	Blocked []string `json:"blocked,omitempty"`
}

type VideoFileDetails struct {
	// AudioStreams: Audio streams.
	AudioStreams []*VideoFileDetailsAudioStream `json:"audioStreams,omitempty"`

	// BitrateBps: Combined audio and video bitrate, in bits per second.
	BitrateBps uint64 `json:"bitrateBps,omitempty,string"`

	// Container: Container format used.
	Container string `json:"container,omitempty"`

	// CreationTime: Date and time when the video file was created, in ISO
	// 8601 format. Currently the only ISO 8601 formats produced are: - Date
	// only: YYYY-MM-DD - Naive time: YYYY-MM-DDTHH:MM:SS - Time with
	// timezone: YYYY-MM-DDTHH:MM:SS+HH:MM
	CreationTime string `json:"creationTime,omitempty"`

	// DurationMs: Video duration in milliseconds.
	DurationMs uint64 `json:"durationMs,omitempty,string"`

	// FileName: File name.
	FileName string `json:"fileName,omitempty"`

	// FileSize: File size.
	FileSize uint64 `json:"fileSize,omitempty,string"`

	// FileType: File type.
	FileType string `json:"fileType,omitempty"`

	// RecordingLocation: Location of the place where the video was
	// recorded.
	RecordingLocation *GeoPoint `json:"recordingLocation,omitempty"`

	// VideoStreams: Video streams.
	VideoStreams []*VideoFileDetailsVideoStream `json:"videoStreams,omitempty"`
}

type VideoFileDetailsAudioStream struct {
	// BitrateBps: Audio stream bitrate, in bits per second.
	BitrateBps uint64 `json:"bitrateBps,omitempty,string"`

	// ChannelCount: Number of audio channels.
	ChannelCount int64 `json:"channelCount,omitempty"`

	// Codec: Audio codec used.
	Codec string `json:"codec,omitempty"`

	// Vendor: Audio vendor identifier, typically a four-letter vendor code.
	Vendor string `json:"vendor,omitempty"`
}

type VideoFileDetailsVideoStream struct {
	// AspectRatio: Display aspect ratio, which might differ from
	// width_pixels / height_pixels.
	AspectRatio float64 `json:"aspectRatio,omitempty"`

	// BitrateBps: Video stream bitrate, in bits per second.
	BitrateBps uint64 `json:"bitrateBps,omitempty,string"`

	// Codec: Video codec used.
	Codec string `json:"codec,omitempty"`

	// FrameRateFps: Video frame rate, in frames per second.
	FrameRateFps float64 `json:"frameRateFps,omitempty"`

	// HeightPixels: Video height in pixels.
	HeightPixels int64 `json:"heightPixels,omitempty"`

	// Rotation: Rotation that is necessary to display the video properly.
	Rotation string `json:"rotation,omitempty"`

	// Vendor: Video vendor identifier, typically a four-letter vendor code.
	Vendor string `json:"vendor,omitempty"`

	// WidthPixels: Video width in pixels.
	WidthPixels int64 `json:"widthPixels,omitempty"`
}

type VideoListResponse struct {
	// Etag: The ETag of the response.
	Etag string `json:"etag,omitempty"`

	// Items: A list of videos that match the request criteria.
	Items []*Video `json:"items,omitempty"`

	// Kind: The type of the API response. For this operation, the value
	// will be youtube#videoListResponse.
	Kind string `json:"kind,omitempty"`
}

type VideoMonetizationDetails struct {
	// Access: The value of access indicates whether the video can be
	// monetized or not.
	Access *AccessPolicy `json:"access,omitempty"`
}

type VideoPlayer struct {
	// EmbedHtml: An <iframe> tag that embeds a player that will play the
	// video.
	EmbedHtml string `json:"embedHtml,omitempty"`
}

type VideoProcessingDetails struct {
	// EditorSuggestionsAvailability: Editor suggestions availability.
	EditorSuggestionsAvailability string `json:"editorSuggestionsAvailability,omitempty"`

	// FileDetailsAvailability: File details availability.
	FileDetailsAvailability string `json:"fileDetailsAvailability,omitempty"`

	// ProcessingFailureReason: Reason why video processing has failed.
	ProcessingFailureReason string `json:"processingFailureReason,omitempty"`

	// ProcessingIssuesAvailability: Processing issues availability.
	ProcessingIssuesAvailability string `json:"processingIssuesAvailability,omitempty"`

	// ProcessingProgress: Video processing progress and completion time
	// estimate.
	ProcessingProgress *VideoProcessingDetailsProcessingProgress `json:"processingProgress,omitempty"`

	// ProcessingStatus: Video processing status.
	ProcessingStatus string `json:"processingStatus,omitempty"`

	// TagSuggestionsAvailability: Tag suggestions availability.
	TagSuggestionsAvailability string `json:"tagSuggestionsAvailability,omitempty"`

	// ThumbnailsAvailability: Thumbnails availability.
	ThumbnailsAvailability string `json:"thumbnailsAvailability,omitempty"`
}

type VideoProcessingDetailsProcessingProgress struct {
	// PartsProcessed: Number of parts already processed. Progress expressed
	// in percent should be computed as: 100 * parts_processed /
	// parts_total.
	PartsProcessed uint64 `json:"partsProcessed,omitempty,string"`

	// PartsTotal: An estimate of total number of parts to process. The
	// number might be updated with more precise estimates as the processing
	// progresses.
	PartsTotal uint64 `json:"partsTotal,omitempty,string"`

	// TimeLeftMs: Estimated time till video processing is complete, in
	// milliseconds.
	TimeLeftMs uint64 `json:"timeLeftMs,omitempty,string"`
}

type VideoRecordingDetails struct {
	// Location: The geolocation information associated with the video.
	Location *GeoPoint `json:"location,omitempty"`

	// LocationDescription: The text description of the location where the
	// video was recorded.
	LocationDescription string `json:"locationDescription,omitempty"`

	// RecordingDate: The date and time when the video was recorded. The
	// value is specified in ISO 8601 (YYYY-MM-DDThh:mm:ss.sZ) format.
	RecordingDate string `json:"recordingDate,omitempty"`
}

type VideoSnippet struct {
	// CategoryId: The YouTube video category associated with the video.
	CategoryId string `json:"categoryId,omitempty"`

	// ChannelId: The ID that YouTube uses to uniquely identify the channel
	// that the video was uploaded to.
	ChannelId string `json:"channelId,omitempty"`

	// Description: The video's description.
	Description string `json:"description,omitempty"`

	// PublishedAt: The date and time that the video was uploaded. The value
	// is specified in ISO 8601 (YYYY-MM-DDThh:mm:ss.sZ) format.
	PublishedAt string `json:"publishedAt,omitempty"`

	// Tags: A list of keyword tags associated with the video. Tags may
	// contain spaces. This field is only visible to the video's uploader.
	Tags []string `json:"tags,omitempty"`

	// Thumbnails: A map of thumbnail images associated with the video. For
	// each object in the map, the key is the name of the thumbnail image,
	// and the value is an object that contains other information about the
	// thumbnail.
	Thumbnails *VideoSnippetThumbnails `json:"thumbnails,omitempty"`

	// Title: The video's title.
	Title string `json:"title,omitempty"`
}

type VideoSnippetThumbnails struct {
}

type VideoStatistics struct {
	// CommentCount: The number of comments for the video.
	CommentCount uint64 `json:"commentCount,omitempty,string"`

	// DislikeCount: The number of users who have indicated that they
	// disliked the video by giving it a negative rating.
	DislikeCount uint64 `json:"dislikeCount,omitempty,string"`

	// FavoriteCount: The number of users who currently have the video
	// marked as a favorite video.
	FavoriteCount uint64 `json:"favoriteCount,omitempty,string"`

	// LikeCount: The number of users who have indicated that they liked the
	// video by giving it a positive rating.
	LikeCount uint64 `json:"likeCount,omitempty,string"`

	// ViewCount: The number of times the video has been viewed.
	ViewCount uint64 `json:"viewCount,omitempty,string"`
}

type VideoStatus struct {
	// Embeddable: This value indicates if the video can be embedded on
	// another website.
	Embeddable bool `json:"embeddable,omitempty"`

	// FailureReason: This value explains why a video failed to upload. This
	// property is only present if the uploadStatus property indicates that
	// the upload failed.
	FailureReason string `json:"failureReason,omitempty"`

	// License: The video's license.
	License string `json:"license,omitempty"`

	// PrivacyStatus: The video's privacy status.
	PrivacyStatus string `json:"privacyStatus,omitempty"`

	// RejectionReason: This value explains why YouTube rejected an uploaded
	// video. This property is only present if the uploadStatus property
	// indicates that the upload was rejected.
	RejectionReason string `json:"rejectionReason,omitempty"`

	// UploadStatus: The status of the uploaded video.
	UploadStatus string `json:"uploadStatus,omitempty"`
}

type VideoSuggestions struct {
	// EditorSuggestions: Editor operations that could improve video
	// quality.
	EditorSuggestions []string `json:"editorSuggestions,omitempty"`

	// ProcessingErrors: Errors encountered during video processing.
	ProcessingErrors []string `json:"processingErrors,omitempty"`

	// ProcessingHints: Hints about how to improve video processing.
	ProcessingHints []string `json:"processingHints,omitempty"`

	// ProcessingWarnings: Warnings produced by the video processing engine.
	ProcessingWarnings []string `json:"processingWarnings,omitempty"`

	// TagSuggestions: Tags that could be added to aid video discovery.
	TagSuggestions []*VideoSuggestionsTagSuggestion `json:"tagSuggestions,omitempty"`
}

type VideoSuggestionsTagSuggestion struct {
	// CategoryRestricts: Set of categories this tag should be restricted
	// to. Tag applies to all categories if there are no restricts.
	CategoryRestricts []string `json:"categoryRestricts,omitempty"`

	// Tag: Tag label.
	Tag string `json:"tag,omitempty"`
}

type VideoTopicDetails struct {
	// TopicIds: A list of Freebase topic IDs associated with the video. You
	// can retrieve information about each topic using the Freebase Topic
	// API.
	TopicIds []string `json:"topicIds,omitempty"`
}

// method id "youtube.activities.insert":

type ActivitiesInsertCall struct {
	s        *Service
	part     string
	activity *Activity
	opt_     map[string]interface{}
}

// Insert: Posts a bulletin for a specific channel. (The user submitting
// the request must be authorized to act on the channel's behalf.)
func (r *ActivitiesService) Insert(part string, activity *Activity) *ActivitiesInsertCall {
	c := &ActivitiesInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	c.activity = activity
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
	//   "description": "Posts a bulletin for a specific channel. (The user submitting the request must be authorized to act on the channel's behalf.)",
	//   "httpMethod": "POST",
	//   "id": "youtube.activities.insert",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "part": {
	//       "description": "The part parameter serves two purposes in this operation. It identifies the properties that the write operation will set as well as the properties that the API response will include.\n\nThe part names that you can include in the parameter value are snippet and contentDetails.",
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

// List: Returns a list of channel activity events that match the
// request criteria. For example, you can retrieve events associated
// with a particular channel, events associated with the user's
// subscriptions and Google+ friends, or the YouTube home page feed,
// which is customized for each user.
func (r *ActivitiesService) List(part string) *ActivitiesListCall {
	c := &ActivitiesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	return c
}

// ChannelId sets the optional parameter "channelId": The channelId
// parameter specifies a unique YouTube channel ID. The API will then
// return a list of that channel's activities.
func (c *ActivitiesListCall) ChannelId(channelId string) *ActivitiesListCall {
	c.opt_["channelId"] = channelId
	return c
}

// Home sets the optional parameter "home": Set this parameter's value
// to true to retrieve the activity feed that displays on the YouTube
// home page for the currently authenticated user.
func (c *ActivitiesListCall) Home(home string) *ActivitiesListCall {
	c.opt_["home"] = home
	return c
}

// MaxResults sets the optional parameter "maxResults": USE_DESCRIPTION
// --- channels:list:maxResults
func (c *ActivitiesListCall) MaxResults(maxResults int64) *ActivitiesListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// Mine sets the optional parameter "mine": Set this parameter's value
// to true to retrieve a feed of the authenticated user's activities.
func (c *ActivitiesListCall) Mine(mine bool) *ActivitiesListCall {
	c.opt_["mine"] = mine
	return c
}

// PageToken sets the optional parameter "pageToken": USE_DESCRIPTION
// --- channels:list:pageToken
func (c *ActivitiesListCall) PageToken(pageToken string) *ActivitiesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// PublishedAfter sets the optional parameter "publishedAfter": The
// publishedAfter parameter specifies the earliest date and time that an
// activity could have occurred for that activity to be included in the
// API response. If the parameter value specifies a day, but not a time,
// then any activities that occurred that day will be included in the
// result set. The value is specified in ISO 8601
// (YYYY-MM-DDThh:mm:ss.sZ) format.
func (c *ActivitiesListCall) PublishedAfter(publishedAfter string) *ActivitiesListCall {
	c.opt_["publishedAfter"] = publishedAfter
	return c
}

// PublishedBefore sets the optional parameter "publishedBefore": The
// publishedBefore parameter specifies the date and time before which an
// activity must have occurred for that activity to be included in the
// API response. If the parameter value specifies a day, but not a time,
// then any activities that occurred that day will be excluded from the
// result set. The value is specified in ISO 8601
// (YYYY-MM-DDThh:mm:ss.sZ) format.
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
	//   "description": "Returns a list of channel activity events that match the request criteria. For example, you can retrieve events associated with a particular channel, events associated with the user's subscriptions and Google+ friends, or the YouTube home page feed, which is customized for each user.",
	//   "httpMethod": "GET",
	//   "id": "youtube.activities.list",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "channelId": {
	//       "description": "The channelId parameter specifies a unique YouTube channel ID. The API will then return a list of that channel's activities.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "home": {
	//       "description": "Set this parameter's value to true to retrieve the activity feed that displays on the YouTube home page for the currently authenticated user.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "5",
	//       "description": "USE_DESCRIPTION --- channels:list:maxResults",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "50",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "mine": {
	//       "description": "Set this parameter's value to true to retrieve a feed of the authenticated user's activities.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "pageToken": {
	//       "description": "USE_DESCRIPTION --- channels:list:pageToken",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "The part parameter specifies a comma-separated list of one or more activity resource properties that the API response will include. The part names that you can include in the parameter value are id, snippet, and contentDetails.\n\nIf the parameter identifies a property that contains child properties, the child properties will be included in the response. For example, in a activity resource, the snippet property contains other properties that identify the type of activity, a display title for the activity, and so forth. If you set part=snippet, the API response will also contain all of those nested properties.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "publishedAfter": {
	//       "description": "The publishedAfter parameter specifies the earliest date and time that an activity could have occurred for that activity to be included in the API response. If the parameter value specifies a day, but not a time, then any activities that occurred that day will be included in the result set. The value is specified in ISO 8601 (YYYY-MM-DDThh:mm:ss.sZ) format.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "publishedBefore": {
	//       "description": "The publishedBefore parameter specifies the date and time before which an activity must have occurred for that activity to be included in the API response. If the parameter value specifies a day, but not a time, then any activities that occurred that day will be excluded from the result set. The value is specified in ISO 8601 (YYYY-MM-DDThh:mm:ss.sZ) format.",
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

// List: Returns a collection of zero or more channel resources that
// match the request criteria.
func (r *ChannelsService) List(part string) *ChannelsListCall {
	c := &ChannelsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	return c
}

// CategoryId sets the optional parameter "categoryId": The categoryId
// parameter specifies a YouTube guide category, thereby requesting
// YouTube channels associated with that category.
func (c *ChannelsListCall) CategoryId(categoryId string) *ChannelsListCall {
	c.opt_["categoryId"] = categoryId
	return c
}

// Id sets the optional parameter "id": The id parameter specifies a
// comma-separated list of the YouTube channel ID(s) for the resource(s)
// that are being retrieved. In a channel resource, the id property
// specifies the channel's YouTube channel ID.
func (c *ChannelsListCall) Id(id string) *ChannelsListCall {
	c.opt_["id"] = id
	return c
}

// MaxResults sets the optional parameter "maxResults": The maxResults
// parameter specifies the maximum number of items that should be
// returned in the result set.
func (c *ChannelsListCall) MaxResults(maxResults int64) *ChannelsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// Mine sets the optional parameter "mine": Set this parameter's value
// to true to instruct the API to only return channels owned by the
// authenticated user.
func (c *ChannelsListCall) Mine(mine bool) *ChannelsListCall {
	c.opt_["mine"] = mine
	return c
}

// MySubscribers sets the optional parameter "mySubscribers": Set this
// parameter's value to true to retrieve a list of channels that
// subscribed to the authenticated user's channel.
func (c *ChannelsListCall) MySubscribers(mySubscribers string) *ChannelsListCall {
	c.opt_["mySubscribers"] = mySubscribers
	return c
}

// PageToken sets the optional parameter "pageToken": The pageToken
// parameter identifies a specific page in the result set that should be
// returned. In an API response, the nextPageToken and prevPageToken
// properties identify other pages that could be retrieved.
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
	//   "description": "Returns a collection of zero or more channel resources that match the request criteria.",
	//   "httpMethod": "GET",
	//   "id": "youtube.channels.list",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "categoryId": {
	//       "description": "The categoryId parameter specifies a YouTube guide category, thereby requesting YouTube channels associated with that category.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "id": {
	//       "description": "The id parameter specifies a comma-separated list of the YouTube channel ID(s) for the resource(s) that are being retrieved. In a channel resource, the id property specifies the channel's YouTube channel ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "5",
	//       "description": "The maxResults parameter specifies the maximum number of items that should be returned in the result set.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "50",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "mine": {
	//       "description": "Set this parameter's value to true to instruct the API to only return channels owned by the authenticated user.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "mySubscribers": {
	//       "description": "Set this parameter's value to true to retrieve a list of channels that subscribed to the authenticated user's channel.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The pageToken parameter identifies a specific page in the result set that should be returned. In an API response, the nextPageToken and prevPageToken properties identify other pages that could be retrieved.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "The part parameter specifies a comma-separated list of one or more channel resource properties that the API response will include. The part names that you can include in the parameter value are id, snippet, contentDetails, statistics, and topicDetails.\n\nIf the parameter identifies a property that contains child properties, the child properties will be included in the response. For example, in a channel resource, the contentDetails property contains other properties, such as the uploads properties. As such, if you set part=contentDetails, the API response will also contain all of those nested properties.",
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

// List: Returns a list of categories that can be associated with
// YouTube channels.
func (r *GuideCategoriesService) List(part string) *GuideCategoriesListCall {
	c := &GuideCategoriesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	return c
}

// Hl sets the optional parameter "hl": The hl parameter specifies the
// language that will be used for text values in the API response.
func (c *GuideCategoriesListCall) Hl(hl string) *GuideCategoriesListCall {
	c.opt_["hl"] = hl
	return c
}

// Id sets the optional parameter "id": The id parameter specifies a
// comma-separated list of the YouTube channel category ID(s) for the
// resource(s) that are being retrieved. In a guideCategory resource,
// the id property specifies the YouTube channel category ID.
func (c *GuideCategoriesListCall) Id(id string) *GuideCategoriesListCall {
	c.opt_["id"] = id
	return c
}

// RegionCode sets the optional parameter "regionCode": The regionCode
// parameter instructs the API to return the list of guide categories
// available in the specified country. The parameter value is an ISO
// 3166-1 alpha-2 country code.
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
	//   "description": "Returns a list of categories that can be associated with YouTube channels.",
	//   "httpMethod": "GET",
	//   "id": "youtube.guideCategories.list",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "hl": {
	//       "default": "en-US",
	//       "description": "The hl parameter specifies the language that will be used for text values in the API response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "id": {
	//       "description": "The id parameter specifies a comma-separated list of the YouTube channel category ID(s) for the resource(s) that are being retrieved. In a guideCategory resource, the id property specifies the YouTube channel category ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "The part parameter specifies a comma-separated list of one or more guideCategory resource properties that the API response will include. The part names that you can include in the parameter value are id and snippet.\n\nIf the parameter identifies a property that contains child properties, the child properties will be included in the response. For example, in a guideCategory resource, the snippet property contains other properties, such as the category's title. If you set part=snippet, the API response will also contain all of those nested properties.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "regionCode": {
	//       "description": "The regionCode parameter instructs the API to return the list of guide categories available in the specified country. The parameter value is an ISO 3166-1 alpha-2 country code.",
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

// method id "youtube.liveBroadcasts.bind":

type LiveBroadcastsBindCall struct {
	s    *Service
	id   string
	part string
	opt_ map[string]interface{}
}

// Bind: Bind a YouTube live broadcast to a stream.
func (r *LiveBroadcastsService) Bind(id string, part string) *LiveBroadcastsBindCall {
	c := &LiveBroadcastsBindCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	c.part = part
	return c
}

// StreamId sets the optional parameter "streamId": ID of the stream to
// bind to the broadcast
func (c *LiveBroadcastsBindCall) StreamId(streamId string) *LiveBroadcastsBindCall {
	c.opt_["streamId"] = streamId
	return c
}

func (c *LiveBroadcastsBindCall) Do() (*LiveBroadcast, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("id", fmt.Sprintf("%v", c.id))
	params.Set("part", fmt.Sprintf("%v", c.part))
	if v, ok := c.opt_["streamId"]; ok {
		params.Set("streamId", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "liveBroadcasts/bind")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(LiveBroadcast)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Bind a YouTube live broadcast to a stream.",
	//   "httpMethod": "POST",
	//   "id": "youtube.liveBroadcasts.bind",
	//   "parameterOrder": [
	//     "id",
	//     "part"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "ID of the broadcast to which the stream will be bound",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "Live broadcast parts to be returned in the response. Valid values are: id, snippet, status, slateSettings, contentDetails.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "streamId": {
	//       "description": "ID of the stream to bind to the broadcast",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "liveBroadcasts/bind",
	//   "response": {
	//     "$ref": "LiveBroadcast"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube"
	//   ]
	// }

}

// method id "youtube.liveBroadcasts.delete":

type LiveBroadcastsDeleteCall struct {
	s    *Service
	id   string
	opt_ map[string]interface{}
}

// Delete: Delete a YouTube live broadcast.
func (r *LiveBroadcastsService) Delete(id string) *LiveBroadcastsDeleteCall {
	c := &LiveBroadcastsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

func (c *LiveBroadcastsDeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("id", fmt.Sprintf("%v", c.id))
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "liveBroadcasts")
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
	//   "description": "Delete a YouTube live broadcast.",
	//   "httpMethod": "DELETE",
	//   "id": "youtube.liveBroadcasts.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The id parameter specifies the YouTube live broadcast ID for the resource that is being deleted.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "liveBroadcasts",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube"
	//   ]
	// }

}

// method id "youtube.liveBroadcasts.insert":

type LiveBroadcastsInsertCall struct {
	s             *Service
	part          string
	livebroadcast *LiveBroadcast
	opt_          map[string]interface{}
}

// Insert: Insert a YouTube live broadcast.
func (r *LiveBroadcastsService) Insert(part string, livebroadcast *LiveBroadcast) *LiveBroadcastsInsertCall {
	c := &LiveBroadcastsInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	c.livebroadcast = livebroadcast
	return c
}

func (c *LiveBroadcastsInsertCall) Do() (*LiveBroadcast, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.livebroadcast)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("part", fmt.Sprintf("%v", c.part))
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "liveBroadcasts")
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
	ret := new(LiveBroadcast)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Insert a YouTube live broadcast.",
	//   "httpMethod": "POST",
	//   "id": "youtube.liveBroadcasts.insert",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "part": {
	//       "description": "Live broadcast parts to be set for the broadcast as well as included in the returned response. Valid values are: snippet, status, slateSettings, contentDetails.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "liveBroadcasts",
	//   "request": {
	//     "$ref": "LiveBroadcast"
	//   },
	//   "response": {
	//     "$ref": "LiveBroadcast"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube"
	//   ]
	// }

}

// method id "youtube.liveBroadcasts.list":

type LiveBroadcastsListCall struct {
	s    *Service
	part string
	opt_ map[string]interface{}
}

// List: Browse the YouTube broadcast collection.
func (r *LiveBroadcastsService) List(part string) *LiveBroadcastsListCall {
	c := &LiveBroadcastsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	return c
}

// BroadcastStatus sets the optional parameter "broadcastStatus": Filter
// to only return broadcasts with the given status by the authenticated
// user.
func (c *LiveBroadcastsListCall) BroadcastStatus(broadcastStatus string) *LiveBroadcastsListCall {
	c.opt_["broadcastStatus"] = broadcastStatus
	return c
}

// Id sets the optional parameter "id": IDs of the live broadcasts to be
// returned.
func (c *LiveBroadcastsListCall) Id(id string) *LiveBroadcastsListCall {
	c.opt_["id"] = id
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return
func (c *LiveBroadcastsListCall) MaxResults(maxResults int64) *LiveBroadcastsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// Mine sets the optional parameter "mine": Filter to only return
// broadcasts owned by authenticated user.
func (c *LiveBroadcastsListCall) Mine(mine bool) *LiveBroadcastsListCall {
	c.opt_["mine"] = mine
	return c
}

// OnBehalfOf sets the optional parameter "onBehalfOf": ID of the
// Google+ Page for the channel that the request is be on behalf of
func (c *LiveBroadcastsListCall) OnBehalfOf(onBehalfOf string) *LiveBroadcastsListCall {
	c.opt_["onBehalfOf"] = onBehalfOf
	return c
}

// PageToken sets the optional parameter "pageToken": Token for the page
// selection.
func (c *LiveBroadcastsListCall) PageToken(pageToken string) *LiveBroadcastsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

func (c *LiveBroadcastsListCall) Do() (*LiveBroadcastList, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("part", fmt.Sprintf("%v", c.part))
	if v, ok := c.opt_["broadcastStatus"]; ok {
		params.Set("broadcastStatus", fmt.Sprintf("%v", v))
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
	if v, ok := c.opt_["onBehalfOf"]; ok {
		params.Set("onBehalfOf", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "liveBroadcasts")
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
	ret := new(LiveBroadcastList)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Browse the YouTube broadcast collection.",
	//   "httpMethod": "GET",
	//   "id": "youtube.liveBroadcasts.list",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "broadcastStatus": {
	//       "description": "Filter to only return broadcasts with the given status by the authenticated user.",
	//       "enum": [
	//         "active",
	//         "all",
	//         "completed",
	//         "upcoming"
	//       ],
	//       "enumDescriptions": [
	//         "Return active broadcasts.",
	//         "Return all the broadcasts.",
	//         "Return previously completed broadcasts.",
	//         "Return upcoming broadcasts."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "id": {
	//       "description": "IDs of the live broadcasts to be returned.",
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
	//       "description": "Filter to only return broadcasts owned by authenticated user.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "onBehalfOf": {
	//       "description": "ID of the Google+ Page for the channel that the request is be on behalf of",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Token for the page selection.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "Live broadcast parts to include in the returned response. Valid values are: id, snippet, status, slateSettings, contentDetails.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "liveBroadcasts",
	//   "response": {
	//     "$ref": "LiveBroadcastList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube.readonly"
	//   ]
	// }

}

// method id "youtube.liveBroadcasts.transition":

type LiveBroadcastsTransitionCall struct {
	s               *Service
	broadcastStatus string
	id              string
	part            string
	opt_            map[string]interface{}
}

// Transition: Change the broadcasting status of a YouTube live
// broadcast and start all the processes associated with it.
func (r *LiveBroadcastsService) Transition(broadcastStatus string, id string, part string) *LiveBroadcastsTransitionCall {
	c := &LiveBroadcastsTransitionCall{s: r.s, opt_: make(map[string]interface{})}
	c.broadcastStatus = broadcastStatus
	c.id = id
	c.part = part
	return c
}

func (c *LiveBroadcastsTransitionCall) Do() (*LiveBroadcast, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("broadcastStatus", fmt.Sprintf("%v", c.broadcastStatus))
	params.Set("id", fmt.Sprintf("%v", c.id))
	params.Set("part", fmt.Sprintf("%v", c.part))
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "liveBroadcasts/transition")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(LiveBroadcast)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Change the broadcasting status of a YouTube live broadcast and start all the processes associated with it.",
	//   "httpMethod": "POST",
	//   "id": "youtube.liveBroadcasts.transition",
	//   "parameterOrder": [
	//     "broadcastStatus",
	//     "id",
	//     "part"
	//   ],
	//   "parameters": {
	//     "broadcastStatus": {
	//       "description": "Desired broadcast status.",
	//       "enum": [
	//         "complete",
	//         "live",
	//         "testing"
	//       ],
	//       "enumDescriptions": [
	//         "Stop broadcasting.",
	//         "Start broadcasting.",
	//         "Start broadcast testing."
	//       ],
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "id": {
	//       "description": "ID of the broadcast to change status",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "Live broadcast parts to be returned in the response. Valid values are: id, snippet, status, slateSettings, contentDetails.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "liveBroadcasts/transition",
	//   "response": {
	//     "$ref": "LiveBroadcast"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube"
	//   ]
	// }

}

// method id "youtube.liveBroadcasts.update":

type LiveBroadcastsUpdateCall struct {
	s             *Service
	part          string
	livebroadcast *LiveBroadcast
	opt_          map[string]interface{}
}

// Update: Update a YouTube live broadcast.
func (r *LiveBroadcastsService) Update(part string, livebroadcast *LiveBroadcast) *LiveBroadcastsUpdateCall {
	c := &LiveBroadcastsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	c.livebroadcast = livebroadcast
	return c
}

func (c *LiveBroadcastsUpdateCall) Do() (*LiveBroadcast, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.livebroadcast)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("part", fmt.Sprintf("%v", c.part))
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "liveBroadcasts")
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
	ret := new(LiveBroadcast)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update a YouTube live broadcast.",
	//   "httpMethod": "PUT",
	//   "id": "youtube.liveBroadcasts.update",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "part": {
	//       "description": "The part parameter serves two purposes in this operation. It identifies the properties that the write operation will set as well as the properties that the API response will include.\n\nThe part names that you can include in the parameter value are id, snippet, status, slateSettings, contentDetails.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "liveBroadcasts",
	//   "request": {
	//     "$ref": "LiveBroadcast"
	//   },
	//   "response": {
	//     "$ref": "LiveBroadcast"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube"
	//   ]
	// }

}

// method id "youtube.liveStreams.delete":

type LiveStreamsDeleteCall struct {
	s    *Service
	id   string
	opt_ map[string]interface{}
}

// Delete: Delete a live stream.
func (r *LiveStreamsService) Delete(id string) *LiveStreamsDeleteCall {
	c := &LiveStreamsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

func (c *LiveStreamsDeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("id", fmt.Sprintf("%v", c.id))
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "liveStreams")
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
	//   "description": "Delete a live stream.",
	//   "httpMethod": "DELETE",
	//   "id": "youtube.liveStreams.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The id parameter specifies the YouTube live stream ID for the resource that is being deleted.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "liveStreams",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube"
	//   ]
	// }

}

// method id "youtube.liveStreams.insert":

type LiveStreamsInsertCall struct {
	s          *Service
	part       string
	livestream *LiveStream
	opt_       map[string]interface{}
}

// Insert: Insert a YouTube live stream.
func (r *LiveStreamsService) Insert(part string, livestream *LiveStream) *LiveStreamsInsertCall {
	c := &LiveStreamsInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	c.livestream = livestream
	return c
}

func (c *LiveStreamsInsertCall) Do() (*LiveStream, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.livestream)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("part", fmt.Sprintf("%v", c.part))
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "liveStreams")
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
	ret := new(LiveStream)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Insert a YouTube live stream.",
	//   "httpMethod": "POST",
	//   "id": "youtube.liveStreams.insert",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "part": {
	//       "description": "Live stream parts to include in the returned response. Valid values are: id, snippet, cdn, status.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "liveStreams",
	//   "request": {
	//     "$ref": "LiveStream"
	//   },
	//   "response": {
	//     "$ref": "LiveStream"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube"
	//   ]
	// }

}

// method id "youtube.liveStreams.list":

type LiveStreamsListCall struct {
	s    *Service
	part string
	opt_ map[string]interface{}
}

// List: Browse the YouTube live stream collection.
func (r *LiveStreamsService) List(part string) *LiveStreamsListCall {
	c := &LiveStreamsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	return c
}

// Id sets the optional parameter "id": IDs of the live streams to be
// returned.
func (c *LiveStreamsListCall) Id(id string) *LiveStreamsListCall {
	c.opt_["id"] = id
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return
func (c *LiveStreamsListCall) MaxResults(maxResults int64) *LiveStreamsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// Mine sets the optional parameter "mine": Filter to only live streams
// owned by authenticated user.
func (c *LiveStreamsListCall) Mine(mine bool) *LiveStreamsListCall {
	c.opt_["mine"] = mine
	return c
}

// OnBehalfOf sets the optional parameter "onBehalfOf": ID of the
// Google+ Page for the channel that the request is to be on behalf of
func (c *LiveStreamsListCall) OnBehalfOf(onBehalfOf string) *LiveStreamsListCall {
	c.opt_["onBehalfOf"] = onBehalfOf
	return c
}

// PageToken sets the optional parameter "pageToken": Token for the page
// selection.
func (c *LiveStreamsListCall) PageToken(pageToken string) *LiveStreamsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

func (c *LiveStreamsListCall) Do() (*LiveStreamList, error) {
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
	if v, ok := c.opt_["onBehalfOf"]; ok {
		params.Set("onBehalfOf", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "liveStreams")
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
	ret := new(LiveStreamList)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Browse the YouTube live stream collection.",
	//   "httpMethod": "GET",
	//   "id": "youtube.liveStreams.list",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "IDs of the live streams to be returned.",
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
	//       "description": "Filter to only live streams owned by authenticated user.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "onBehalfOf": {
	//       "description": "ID of the Google+ Page for the channel that the request is to be on behalf of",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Token for the page selection.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "Live stream parts to include in the returned response. Valid values are: id, snippet, cdn, status.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "liveStreams",
	//   "response": {
	//     "$ref": "LiveStreamList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube.readonly"
	//   ]
	// }

}

// method id "youtube.liveStreams.update":

type LiveStreamsUpdateCall struct {
	s          *Service
	part       string
	livestream *LiveStream
	opt_       map[string]interface{}
}

// Update: Update a YouTube live stream.
func (r *LiveStreamsService) Update(part string, livestream *LiveStream) *LiveStreamsUpdateCall {
	c := &LiveStreamsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	c.livestream = livestream
	return c
}

func (c *LiveStreamsUpdateCall) Do() (*LiveStream, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.livestream)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("part", fmt.Sprintf("%v", c.part))
	urls := googleapi.ResolveRelative("https://www.googleapis.com/youtube/v3/", "liveStreams")
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
	ret := new(LiveStream)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update a YouTube live stream.",
	//   "httpMethod": "PUT",
	//   "id": "youtube.liveStreams.update",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "part": {
	//       "description": "The part parameter serves two purposes in this operation. It identifies the properties that the write operation will set as well as the properties that the API response will include.\n\nThe part names that you can include in the parameter value are id, snippet, cdn, status.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "liveStreams",
	//   "request": {
	//     "$ref": "LiveStream"
	//   },
	//   "response": {
	//     "$ref": "LiveStream"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube"
	//   ]
	// }

}

// method id "youtube.playlistItems.delete":

type PlaylistItemsDeleteCall struct {
	s    *Service
	id   string
	opt_ map[string]interface{}
}

// Delete: Deletes a playlist item.
func (r *PlaylistItemsService) Delete(id string) *PlaylistItemsDeleteCall {
	c := &PlaylistItemsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

func (c *PlaylistItemsDeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("id", fmt.Sprintf("%v", c.id))
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
	//   "description": "Deletes a playlist item.",
	//   "httpMethod": "DELETE",
	//   "id": "youtube.playlistItems.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The id parameter specifies the YouTube playlist item ID for the playlist item that is being deleted. In a playlistItem resource, the id property specifies the playlist item's ID.",
	//       "location": "query",
	//       "required": true,
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

// Insert: Adds a resource to a playlist.
func (r *PlaylistItemsService) Insert(part string, playlistitem *PlaylistItem) *PlaylistItemsInsertCall {
	c := &PlaylistItemsInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	c.playlistitem = playlistitem
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
	//   "description": "Adds a resource to a playlist.",
	//   "httpMethod": "POST",
	//   "id": "youtube.playlistItems.insert",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "part": {
	//       "description": "The part parameter serves two purposes in this operation. It identifies the properties that the write operation will set as well as the properties that the API response will include.\n\nThe part names that you can include in the parameter value are snippet and contentDetails.",
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

// List: Returns a collection of playlist items that match the API
// request parameters. You can retrieve all of the playlist items in a
// specified playlist or retrieve one or more playlist items by their
// unique IDs.
func (r *PlaylistItemsService) List(part string) *PlaylistItemsListCall {
	c := &PlaylistItemsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	return c
}

// Id sets the optional parameter "id": The id parameter specifies a
// comma-separated list of one or more unique playlist item IDs.
func (c *PlaylistItemsListCall) Id(id string) *PlaylistItemsListCall {
	c.opt_["id"] = id
	return c
}

// MaxResults sets the optional parameter "maxResults": USE_DESCRIPTION
// --- channels:list:maxResults
func (c *PlaylistItemsListCall) MaxResults(maxResults int64) *PlaylistItemsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": USE_DESCRIPTION
// --- channels:list:pageToken
func (c *PlaylistItemsListCall) PageToken(pageToken string) *PlaylistItemsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// PlaylistId sets the optional parameter "playlistId": The playlistId
// parameter specifies the unique ID of the playlist for which you want
// to retrieve playlist items. Note that even though this is an optional
// parameter, every request to retrieve playlist items must specify a
// value for either the id parameter or the playlistId parameter.
func (c *PlaylistItemsListCall) PlaylistId(playlistId string) *PlaylistItemsListCall {
	c.opt_["playlistId"] = playlistId
	return c
}

// VideoId sets the optional parameter "videoId": The videoId parameter
// specifies that the request should return only the playlist items that
// contain the specified video.
func (c *PlaylistItemsListCall) VideoId(videoId string) *PlaylistItemsListCall {
	c.opt_["videoId"] = videoId
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
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["playlistId"]; ok {
		params.Set("playlistId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["videoId"]; ok {
		params.Set("videoId", fmt.Sprintf("%v", v))
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
	//   "description": "Returns a collection of playlist items that match the API request parameters. You can retrieve all of the playlist items in a specified playlist or retrieve one or more playlist items by their unique IDs.",
	//   "httpMethod": "GET",
	//   "id": "youtube.playlistItems.list",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The id parameter specifies a comma-separated list of one or more unique playlist item IDs.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "5",
	//       "description": "USE_DESCRIPTION --- channels:list:maxResults",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "50",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "USE_DESCRIPTION --- channels:list:pageToken",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "The part parameter specifies a comma-separated list of one or more playlistItem resource properties that the API response will include. The part names that you can include in the parameter value are id, snippet, and contentDetails.\n\nIf the parameter identifies a property that contains child properties, the child properties will be included in the response. For example, in a playlistItem resource, the snippet property contains numerous fields, including the title, description, position, and resourceId properties. As such, if you set part=snippet, the API response will contain all of those properties.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "playlistId": {
	//       "description": "The playlistId parameter specifies the unique ID of the playlist for which you want to retrieve playlist items. Note that even though this is an optional parameter, every request to retrieve playlist items must specify a value for either the id parameter or the playlistId parameter.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "videoId": {
	//       "description": "The videoId parameter specifies that the request should return only the playlist items that contain the specified video.",
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

// Update: Modifies a playlist item. For example, you could update the
// item's position in the playlist.
func (r *PlaylistItemsService) Update(part string, playlistitem *PlaylistItem) *PlaylistItemsUpdateCall {
	c := &PlaylistItemsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	c.playlistitem = playlistitem
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
	//   "description": "Modifies a playlist item. For example, you could update the item's position in the playlist.",
	//   "httpMethod": "PUT",
	//   "id": "youtube.playlistItems.update",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "part": {
	//       "description": "The part parameter serves two purposes in this operation. It identifies the properties that the write operation will set as well as the properties that the API response will include.\n\nThe part names that you can include in the parameter value are snippet and contentDetails.\n\nNote that this method will override the existing values for all of the mutable properties that are contained in any parts that the parameter value specifies. For example, a playlist item can specify a start time and end time, which identify the times portion of the video that should play when users watch the video in the playlist. If your request is updating a playlist item that sets these values, and the request's part parameter value includes the contentDetails part, the playlist item's start and end times will be updated to whatever value the request body specifies. If the request body does not specify values, the existing start and end times will be removed and replaced with the default settings.",
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

// Delete: Deletes a playlist.
func (r *PlaylistsService) Delete(id string) *PlaylistsDeleteCall {
	c := &PlaylistsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

func (c *PlaylistsDeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("id", fmt.Sprintf("%v", c.id))
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
	//   "description": "Deletes a playlist.",
	//   "httpMethod": "DELETE",
	//   "id": "youtube.playlists.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The id parameter specifies the YouTube playlist ID for the playlist that is being deleted. In a playlist resource, the id property specifies the playlist's ID.",
	//       "location": "query",
	//       "required": true,
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

// Insert: Creates a playlist.
func (r *PlaylistsService) Insert(part string, playlist *Playlist) *PlaylistsInsertCall {
	c := &PlaylistsInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	c.playlist = playlist
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
	//   "description": "Creates a playlist.",
	//   "httpMethod": "POST",
	//   "id": "youtube.playlists.insert",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "part": {
	//       "description": "The part parameter serves two purposes in this operation. It identifies the properties that the write operation will set as well as the properties that the API response will include.\n\nThe part names that you can include in the parameter value are snippet and status.",
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

// List: Returns a collection of playlists that match the API request
// parameters. For example, you can retrieve all playlists that the
// authenticated user owns, or you can retrieve one or more playlists by
// their unique IDs.
func (r *PlaylistsService) List(part string) *PlaylistsListCall {
	c := &PlaylistsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	return c
}

// ChannelId sets the optional parameter "channelId": This value
// indicates that the API should only return the specified channel's
// playlists.
func (c *PlaylistsListCall) ChannelId(channelId string) *PlaylistsListCall {
	c.opt_["channelId"] = channelId
	return c
}

// Id sets the optional parameter "id": The id parameter specifies a
// comma-separated list of the YouTube playlist ID(s) for the
// resource(s) that are being retrieved. In a playlist resource, the id
// property specifies the playlist's YouTube playlist ID.
func (c *PlaylistsListCall) Id(id string) *PlaylistsListCall {
	c.opt_["id"] = id
	return c
}

// MaxResults sets the optional parameter "maxResults": USE_DESCRIPTION
// --- channels:list:maxResults
func (c *PlaylistsListCall) MaxResults(maxResults int64) *PlaylistsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// Mine sets the optional parameter "mine": Set this parameter's value
// to true to instruct the API to only return playlists owned by the
// authenticated user.
func (c *PlaylistsListCall) Mine(mine bool) *PlaylistsListCall {
	c.opt_["mine"] = mine
	return c
}

// PageToken sets the optional parameter "pageToken": USE_DESCRIPTION
// --- channels:list:pageToken
func (c *PlaylistsListCall) PageToken(pageToken string) *PlaylistsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

func (c *PlaylistsListCall) Do() (*PlaylistListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("part", fmt.Sprintf("%v", c.part))
	if v, ok := c.opt_["channelId"]; ok {
		params.Set("channelId", fmt.Sprintf("%v", v))
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
	//   "description": "Returns a collection of playlists that match the API request parameters. For example, you can retrieve all playlists that the authenticated user owns, or you can retrieve one or more playlists by their unique IDs.",
	//   "httpMethod": "GET",
	//   "id": "youtube.playlists.list",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "channelId": {
	//       "description": "This value indicates that the API should only return the specified channel's playlists.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "id": {
	//       "description": "The id parameter specifies a comma-separated list of the YouTube playlist ID(s) for the resource(s) that are being retrieved. In a playlist resource, the id property specifies the playlist's YouTube playlist ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "5",
	//       "description": "USE_DESCRIPTION --- channels:list:maxResults",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "50",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "mine": {
	//       "description": "Set this parameter's value to true to instruct the API to only return playlists owned by the authenticated user.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "pageToken": {
	//       "description": "USE_DESCRIPTION --- channels:list:pageToken",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "The part parameter specifies a comma-separated list of one or more playlist resource properties that the API response will include. The part names that you can include in the parameter value are id, snippet, and status.\n\nIf the parameter identifies a property that contains child properties, the child properties will be included in the response. For example, in a playlist resource, the snippet property contains properties like author, title, description, tags, and timeCreated. As such, if you set part=snippet, the API response will contain all of those properties.",
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

// Update: Modifies a playlist. For example, you could change a
// playlist's title, description, or privacy status.
func (r *PlaylistsService) Update(part string, playlist *Playlist) *PlaylistsUpdateCall {
	c := &PlaylistsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	c.playlist = playlist
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
	//   "description": "Modifies a playlist. For example, you could change a playlist's title, description, or privacy status.",
	//   "httpMethod": "PUT",
	//   "id": "youtube.playlists.update",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "part": {
	//       "description": "The part parameter serves two purposes in this operation. It identifies the properties that the write operation will set as well as the properties that the API response will include.\n\nThe part names that you can include in the parameter value are snippet and status.\n\nNote that this method will override the existing values for all of the mutable properties that are contained in any parts that the parameter value specifies. For example, a playlist's privacy setting is contained in the status part. As such, if your request is updating a private playlist, and the request's part parameter value includes the status part, the playlist's privacy setting will be updated to whatever value the request body specifies. If the request body does not specify a value, the existing privacy setting will be removed and the playlist will revert to the default privacy setting.",
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

// List: Returns a collection of search results that match the query
// parameters specified in the API request. By default, a search result
// set identifies matching video, channel, and playlist resources, but
// you can also configure queries to only retrieve a specific type of
// resource.
func (r *SearchService) List(part string) *SearchListCall {
	c := &SearchListCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	return c
}

// ChannelId sets the optional parameter "channelId": The channelId
// parameter indicates that the API response should only contain
// resources created by the channel
func (c *SearchListCall) ChannelId(channelId string) *SearchListCall {
	c.opt_["channelId"] = channelId
	return c
}

// MaxResults sets the optional parameter "maxResults": USE_DESCRIPTION
// --- channels:list:maxResults
func (c *SearchListCall) MaxResults(maxResults int64) *SearchListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// Order sets the optional parameter "order": The order parameter
// specifies the method that will be used to order resources in the API
// response.
func (c *SearchListCall) Order(order string) *SearchListCall {
	c.opt_["order"] = order
	return c
}

// PageToken sets the optional parameter "pageToken": USE_DESCRIPTION
// --- channels:list:pageToken
func (c *SearchListCall) PageToken(pageToken string) *SearchListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// PublishedAfter sets the optional parameter "publishedAfter": The
// publishedAfter parameter indicates that the API response should only
// contain resources created after the specified time. The value is an
// RFC 3339 formatted date-time value (1970-01-01T00:00:00Z).
func (c *SearchListCall) PublishedAfter(publishedAfter string) *SearchListCall {
	c.opt_["publishedAfter"] = publishedAfter
	return c
}

// PublishedBefore sets the optional parameter "publishedBefore": The
// publishedBefore parameter indicates that the API response should only
// contain resources created before the specified time. The value is an
// RFC 3339 formatted date-time value (1970-01-01T00:00:00Z).
func (c *SearchListCall) PublishedBefore(publishedBefore string) *SearchListCall {
	c.opt_["publishedBefore"] = publishedBefore
	return c
}

// Q sets the optional parameter "q": The q parameter specifies the
// query term to search for.
func (c *SearchListCall) Q(q string) *SearchListCall {
	c.opt_["q"] = q
	return c
}

// RegionCode sets the optional parameter "regionCode": The regionCode
// parameter instructs the API to return search results for the
// specified country. The parameter value is an ISO 3166-1 alpha-2
// country code.
func (c *SearchListCall) RegionCode(regionCode string) *SearchListCall {
	c.opt_["regionCode"] = regionCode
	return c
}

// RelatedToVideoId sets the optional parameter "relatedToVideoId": The
// relatedToVideoId parameter retrieves a list of videos that are
// related to the video that the parameter value identifies. The
// parameter value must be set to a YouTube video ID and, if you are
// using this parameter, the type parameter must be set to video.
func (c *SearchListCall) RelatedToVideoId(relatedToVideoId string) *SearchListCall {
	c.opt_["relatedToVideoId"] = relatedToVideoId
	return c
}

// TopicId sets the optional parameter "topicId": The topicId parameter
// indicates that the API response should only contain resources
// associated with the specified topic. The value identifies a Freebase
// topic ID.
func (c *SearchListCall) TopicId(topicId string) *SearchListCall {
	c.opt_["topicId"] = topicId
	return c
}

// Type sets the optional parameter "type": The type parameter restricts
// a search query to only retrieve a particular type of resource.
func (c *SearchListCall) Type(type_ string) *SearchListCall {
	c.opt_["type"] = type_
	return c
}

// VideoCaption sets the optional parameter "videoCaption": The
// videoCaption parameter indicates whether the API should filter video
// search results based on whether they have captions.
func (c *SearchListCall) VideoCaption(videoCaption string) *SearchListCall {
	c.opt_["videoCaption"] = videoCaption
	return c
}

// VideoCategoryId sets the optional parameter "videoCategoryId": The
// videoCategoryId parameter filters video search results based on their
// category.
func (c *SearchListCall) VideoCategoryId(videoCategoryId string) *SearchListCall {
	c.opt_["videoCategoryId"] = videoCategoryId
	return c
}

// VideoDefinition sets the optional parameter "videoDefinition": The
// videoDefinition parameter lets you restrict a search to only include
// either high definition (HD) or standard definition (SD) videos. HD
// videos are available for playback in at least 720p, though higher
// resolutions, like 1080p, might also be available.
func (c *SearchListCall) VideoDefinition(videoDefinition string) *SearchListCall {
	c.opt_["videoDefinition"] = videoDefinition
	return c
}

// VideoDimension sets the optional parameter "videoDimension": The
// videoDimension parameter lets you restrict a search to only retrieve
// 2D or 3D videos.
func (c *SearchListCall) VideoDimension(videoDimension string) *SearchListCall {
	c.opt_["videoDimension"] = videoDimension
	return c
}

// VideoDuration sets the optional parameter "videoDuration": The
// videoDuration parameter filters video search results based on their
// duration.
func (c *SearchListCall) VideoDuration(videoDuration string) *SearchListCall {
	c.opt_["videoDuration"] = videoDuration
	return c
}

// VideoEmbeddable sets the optional parameter "videoEmbeddable": The
// videoEmbeddable parameter lets you to restrict a search to only
// videos that can be embedded into a webpage.
func (c *SearchListCall) VideoEmbeddable(videoEmbeddable string) *SearchListCall {
	c.opt_["videoEmbeddable"] = videoEmbeddable
	return c
}

// VideoLicense sets the optional parameter "videoLicense": The
// videoLicense parameter filters search results to only include videos
// with a particular license. YouTube lets video uploaders choose to
// attach either the Creative Commons license or the standard YouTube
// license to each of their videos.
func (c *SearchListCall) VideoLicense(videoLicense string) *SearchListCall {
	c.opt_["videoLicense"] = videoLicense
	return c
}

// VideoSyndicated sets the optional parameter "videoSyndicated": The
// videoSyndicated parameter lets you to restrict a search to only
// videos that can be played outside youtube.com.
func (c *SearchListCall) VideoSyndicated(videoSyndicated string) *SearchListCall {
	c.opt_["videoSyndicated"] = videoSyndicated
	return c
}

func (c *SearchListCall) Do() (*SearchListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("part", fmt.Sprintf("%v", c.part))
	if v, ok := c.opt_["channelId"]; ok {
		params.Set("channelId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["order"]; ok {
		params.Set("order", fmt.Sprintf("%v", v))
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
	if v, ok := c.opt_["q"]; ok {
		params.Set("q", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["regionCode"]; ok {
		params.Set("regionCode", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["relatedToVideoId"]; ok {
		params.Set("relatedToVideoId", fmt.Sprintf("%v", v))
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
	if v, ok := c.opt_["videoCategoryId"]; ok {
		params.Set("videoCategoryId", fmt.Sprintf("%v", v))
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
	if v, ok := c.opt_["videoEmbeddable"]; ok {
		params.Set("videoEmbeddable", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["videoLicense"]; ok {
		params.Set("videoLicense", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["videoSyndicated"]; ok {
		params.Set("videoSyndicated", fmt.Sprintf("%v", v))
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
	//   "description": "Returns a collection of search results that match the query parameters specified in the API request. By default, a search result set identifies matching video, channel, and playlist resources, but you can also configure queries to only retrieve a specific type of resource.",
	//   "httpMethod": "GET",
	//   "id": "youtube.search.list",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "channelId": {
	//       "description": "The channelId parameter indicates that the API response should only contain resources created by the channel",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "5",
	//       "description": "USE_DESCRIPTION --- channels:list:maxResults",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "50",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "order": {
	//       "default": "SEARCH_SORT_RELEVANCE",
	//       "description": "The order parameter specifies the method that will be used to order resources in the API response.",
	//       "enum": [
	//         "date",
	//         "rating",
	//         "relevance",
	//         "viewCount"
	//       ],
	//       "enumDescriptions": [
	//         "Resources are sorted in reverse chronological order based on the date they were created.",
	//         "Resources are sorted from highest to lowest rating.",
	//         "Resources are sorted based on their relevance to the search query. This is the default value for this parameter.",
	//         "Resources are sorted from highest to lowest number of views."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "USE_DESCRIPTION --- channels:list:pageToken",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "The part parameter specifies a comma-separated list of one or more search resource properties that the API response will include. The part names that you can include in the parameter value are id and snippet.\n\nIf the parameter identifies a property that contains child properties, the child properties will be included in the response. For example, in a search result, the snippet property contains other properties that identify the result's title, description, and so forth. If you set part=snippet, the API response will also contain all of those nested properties.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "publishedAfter": {
	//       "description": "The publishedAfter parameter indicates that the API response should only contain resources created after the specified time. The value is an RFC 3339 formatted date-time value (1970-01-01T00:00:00Z).",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "publishedBefore": {
	//       "description": "The publishedBefore parameter indicates that the API response should only contain resources created before the specified time. The value is an RFC 3339 formatted date-time value (1970-01-01T00:00:00Z).",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "q": {
	//       "description": "The q parameter specifies the query term to search for.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "regionCode": {
	//       "description": "The regionCode parameter instructs the API to return search results for the specified country. The parameter value is an ISO 3166-1 alpha-2 country code.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "relatedToVideoId": {
	//       "description": "The relatedToVideoId parameter retrieves a list of videos that are related to the video that the parameter value identifies. The parameter value must be set to a YouTube video ID and, if you are using this parameter, the type parameter must be set to video.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "topicId": {
	//       "description": "The topicId parameter indicates that the API response should only contain resources associated with the specified topic. The value identifies a Freebase topic ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "type": {
	//       "default": "video,channel,playlist",
	//       "description": "The type parameter restricts a search query to only retrieve a particular type of resource.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "videoCaption": {
	//       "description": "The videoCaption parameter indicates whether the API should filter video search results based on whether they have captions.",
	//       "enum": [
	//         "any",
	//         "closedCaption",
	//         "none"
	//       ],
	//       "enumDescriptions": [
	//         "Do not filter results based on caption availability.",
	//         "Only include videos that have captions.",
	//         "Only include videos that do not have captions."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "videoCategoryId": {
	//       "description": "The videoCategoryId parameter filters video search results based on their category.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "videoDefinition": {
	//       "description": "The videoDefinition parameter lets you restrict a search to only include either high definition (HD) or standard definition (SD) videos. HD videos are available for playback in at least 720p, though higher resolutions, like 1080p, might also be available.",
	//       "enum": [
	//         "any",
	//         "high",
	//         "standard"
	//       ],
	//       "enumDescriptions": [
	//         "Return all videos, regardless of their resolution.",
	//         "Only retrieve HD videos.",
	//         "Only retrieve videos in standard definition."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "videoDimension": {
	//       "description": "The videoDimension parameter lets you restrict a search to only retrieve 2D or 3D videos.",
	//       "enum": [
	//         "2d",
	//         "3d",
	//         "any"
	//       ],
	//       "enumDescriptions": [
	//         "Restrict search results to exclude 3D videos.",
	//         "Restrict search results to only include 3D videos.",
	//         "Include both 3D and non-3D videos in returned results. This is the default value."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "videoDuration": {
	//       "description": "The videoDuration parameter filters video search results based on their duration.",
	//       "enum": [
	//         "any",
	//         "long",
	//         "medium",
	//         "short"
	//       ],
	//       "enumDescriptions": [
	//         "Do not filter video search results based on their duration. This is the default value.",
	//         "Only include videos longer than 20 minutes.",
	//         "Only include videos that are between four and 20 minutes long (inclusive).",
	//         "Only include videos that are less than four minutes long."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "videoEmbeddable": {
	//       "description": "The videoEmbeddable parameter lets you to restrict a search to only videos that can be embedded into a webpage.",
	//       "enum": [
	//         "any",
	//         "true"
	//       ],
	//       "enumDescriptions": [
	//         "Return all videos, embeddable or not.",
	//         "Only retrieve embeddable videos."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "videoLicense": {
	//       "description": "The videoLicense parameter filters search results to only include videos with a particular license. YouTube lets video uploaders choose to attach either the Creative Commons license or the standard YouTube license to each of their videos.",
	//       "enum": [
	//         "any",
	//         "creativeCommon",
	//         "youtube"
	//       ],
	//       "enumDescriptions": [
	//         "Return all videos, regardless of which license they have, that match the query parameters.",
	//         "Only return videos that have a Creative Commons license. Users can reuse videos with this license in other videos that they create. Learn more.",
	//         "Only return videos that have the standard YouTube license."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "videoSyndicated": {
	//       "description": "The videoSyndicated parameter lets you to restrict a search to only videos that can be played outside youtube.com.",
	//       "enum": [
	//         "any",
	//         "true"
	//       ],
	//       "enumDescriptions": [
	//         "Return all videos, syndicated or not.",
	//         "Only retrieve syndicated videos."
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

// Delete: Deletes a subscription.
func (r *SubscriptionsService) Delete(id string) *SubscriptionsDeleteCall {
	c := &SubscriptionsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

func (c *SubscriptionsDeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("id", fmt.Sprintf("%v", c.id))
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
	//   "description": "Deletes a subscription.",
	//   "httpMethod": "DELETE",
	//   "id": "youtube.subscriptions.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The id parameter specifies the YouTube subscription ID for the resource that is being deleted. In a subscription resource, the id property specifies the YouTube subscription ID.",
	//       "location": "query",
	//       "required": true,
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

// Insert: Adds a subscription for the authenticated user's channel.
func (r *SubscriptionsService) Insert(part string, subscription *Subscription) *SubscriptionsInsertCall {
	c := &SubscriptionsInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	c.subscription = subscription
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
	//   "description": "Adds a subscription for the authenticated user's channel.",
	//   "httpMethod": "POST",
	//   "id": "youtube.subscriptions.insert",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "part": {
	//       "description": "The part parameter serves two purposes in this operation. It identifies the properties that the write operation will set as well as the properties that the API response will include.\n\nThe part names that you can include in the parameter value are snippet and contentDetails.",
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

// List: Returns subscription resources that match the API request
// criteria.
func (r *SubscriptionsService) List(part string) *SubscriptionsListCall {
	c := &SubscriptionsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	return c
}

// ChannelId sets the optional parameter "channelId": The channelId
// parameter specifies a YouTube channel ID. The API will only return
// that channel's subscriptions.
func (c *SubscriptionsListCall) ChannelId(channelId string) *SubscriptionsListCall {
	c.opt_["channelId"] = channelId
	return c
}

// ForChannelId sets the optional parameter "forChannelId": The
// forChannelId parameter specifies a comma-separated list of channel
// IDs. The API response will then only contain subscriptions matching
// those channels.
func (c *SubscriptionsListCall) ForChannelId(forChannelId string) *SubscriptionsListCall {
	c.opt_["forChannelId"] = forChannelId
	return c
}

// Id sets the optional parameter "id": The id parameter specifies a
// comma-separated list of the YouTube subscription ID(s) for the
// resource(s) that are being retrieved. In a subscription resource, the
// id property specifies the YouTube subscription ID.
func (c *SubscriptionsListCall) Id(id string) *SubscriptionsListCall {
	c.opt_["id"] = id
	return c
}

// MaxResults sets the optional parameter "maxResults": USE_DESCRIPTION
// --- channels:list:maxResults
func (c *SubscriptionsListCall) MaxResults(maxResults int64) *SubscriptionsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// Mine sets the optional parameter "mine": Set this parameter's value
// to true to retrieve a feed of the authenticated user's subscriptions.
func (c *SubscriptionsListCall) Mine(mine bool) *SubscriptionsListCall {
	c.opt_["mine"] = mine
	return c
}

// Order sets the optional parameter "order": The order parameter
// specifies the method that will be used to sort resources in the API
// response.
func (c *SubscriptionsListCall) Order(order string) *SubscriptionsListCall {
	c.opt_["order"] = order
	return c
}

// PageToken sets the optional parameter "pageToken": USE_DESCRIPTION
// --- channels:list:pageToken
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
	//   "description": "Returns subscription resources that match the API request criteria.",
	//   "httpMethod": "GET",
	//   "id": "youtube.subscriptions.list",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "channelId": {
	//       "description": "The channelId parameter specifies a YouTube channel ID. The API will only return that channel's subscriptions.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "forChannelId": {
	//       "description": "The forChannelId parameter specifies a comma-separated list of channel IDs. The API response will then only contain subscriptions matching those channels.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "id": {
	//       "description": "The id parameter specifies a comma-separated list of the YouTube subscription ID(s) for the resource(s) that are being retrieved. In a subscription resource, the id property specifies the YouTube subscription ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "5",
	//       "description": "USE_DESCRIPTION --- channels:list:maxResults",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "50",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "mine": {
	//       "description": "Set this parameter's value to true to retrieve a feed of the authenticated user's subscriptions.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "order": {
	//       "default": "SUBSCRIPTION_ORDER_RELEVANCE",
	//       "description": "The order parameter specifies the method that will be used to sort resources in the API response.",
	//       "enum": [
	//         "alphabetical",
	//         "relevance",
	//         "unread"
	//       ],
	//       "enumDescriptions": [
	//         "Sort alphabetically.",
	//         "Sort by relevance.",
	//         "Sort by order of activity."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "USE_DESCRIPTION --- channels:list:pageToken",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "The part parameter specifies a comma-separated list of one or more subscription resource properties that the API response will include. The part names that you can include in the parameter value are id, snippet, and contentDetails.\n\nIf the parameter identifies a property that contains child properties, the child properties will be included in the response. For example, in a subscription resource, the snippet property contains other properties, such as a display title for the subscription. If you set part=snippet, the API response will also contain all of those nested properties.",
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

// List: Returns a list of categories that can be associated with
// YouTube videos.
func (r *VideoCategoriesService) List(part string) *VideoCategoriesListCall {
	c := &VideoCategoriesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	return c
}

// Hl sets the optional parameter "hl": The hl parameter specifies the
// language that should be used for text values in the API response.
func (c *VideoCategoriesListCall) Hl(hl string) *VideoCategoriesListCall {
	c.opt_["hl"] = hl
	return c
}

// Id sets the optional parameter "id": The id parameter specifies a
// comma-separated list of video category IDs for the resources that you
// are retrieving.
func (c *VideoCategoriesListCall) Id(id string) *VideoCategoriesListCall {
	c.opt_["id"] = id
	return c
}

// RegionCode sets the optional parameter "regionCode": The regionCode
// parameter instructs the API to return the list of video categories
// available in the specified country. The parameter value is an ISO
// 3166-1 alpha-2 country code.
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
	//   "description": "Returns a list of categories that can be associated with YouTube videos.",
	//   "httpMethod": "GET",
	//   "id": "youtube.videoCategories.list",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "hl": {
	//       "default": "en_US",
	//       "description": "The hl parameter specifies the language that should be used for text values in the API response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "id": {
	//       "description": "The id parameter specifies a comma-separated list of video category IDs for the resources that you are retrieving.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "The part parameter specifies the videoCategory resource parts that the API response will include. Supported values are id and snippet.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "regionCode": {
	//       "description": "The regionCode parameter instructs the API to return the list of video categories available in the specified country. The parameter value is an ISO 3166-1 alpha-2 country code.",
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

// Delete: Deletes a YouTube video.
func (r *VideosService) Delete(id string) *VideosDeleteCall {
	c := &VideosDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

func (c *VideosDeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("id", fmt.Sprintf("%v", c.id))
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
	//   "description": "Deletes a YouTube video.",
	//   "httpMethod": "DELETE",
	//   "id": "youtube.videos.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The id parameter specifies the YouTube video ID for the resource that is being deleted. In a video resource, the id property specifies the video's ID.",
	//       "location": "query",
	//       "required": true,
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

// Insert: Uploads a video to YouTube and optionally sets the video's
// metadata.
func (r *VideosService) Insert(part string, video *Video) *VideosInsertCall {
	c := &VideosInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	c.video = video
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
	//   "description": "Uploads a video to YouTube and optionally sets the video's metadata.",
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
	//     "part": {
	//       "description": "The part parameter serves two purposes in this operation. It identifies the properties that the write operation will set as well as the properties that the API response will include.\n\nThe part names that you can include in the parameter value are snippet, contentDetails, player, statistics, status, and topicDetails. However, not all of those parts contain properties that can be set when setting or updating a video's metadata. For example, the statistics object encapsulates statistics that YouTube calculates for a video and does not contain values that you can set or modify. If the parameter value specifies a part that does not contain mutable values, that part will still be included in the API response.",
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
	//     "https://www.googleapis.com/auth/youtube.upload",
	//     "https://www.googleapis.com/auth/youtubepartner"
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

// List: Returns a list of videos that match the API request parameters.
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
	//   "description": "Returns a list of videos that match the API request parameters.",
	//   "httpMethod": "GET",
	//   "id": "youtube.videos.list",
	//   "parameterOrder": [
	//     "id",
	//     "part"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The id parameter specifies a comma-separated list of the YouTube video ID(s) for the resource(s) that are being retrieved. In a video resource, the id property specifies the video's ID.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "part": {
	//       "description": "The part parameter specifies a comma-separated list of one or more video resource properties that the API response will include. The part names that you can include in the parameter value are id, snippet, contentDetails, player, statistics, status, and topicDetails.\n\nIf the parameter identifies a property that contains child properties, the child properties will be included in the response. For example, in a video resource, the snippet property contains the channelId, title, description, tags, and categoryId properties. As such, if you set part=snippet, the API response will contain all of those properties.",
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

// Update: Updates a video's metadata.
func (r *VideosService) Update(part string, video *Video) *VideosUpdateCall {
	c := &VideosUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.part = part
	c.video = video
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
	//   "description": "Updates a video's metadata.",
	//   "httpMethod": "PUT",
	//   "id": "youtube.videos.update",
	//   "parameterOrder": [
	//     "part"
	//   ],
	//   "parameters": {
	//     "part": {
	//       "description": "The part parameter serves two purposes in this operation. It identifies the properties that the write operation will set as well as the properties that the API response will include.\n\nThe part names that you can include in the parameter value are snippet, contentDetails, player, statistics, status, and topicDetails.\n\nNote that this method will override the existing values for all of the mutable properties that are contained in any parts that the parameter value specifies. For example, a video's privacy setting is contained in the status part. As such, if your request is updating a private video, and the request's part parameter value includes the status part, the video's privacy setting will be updated to whatever value the request body specifies. If the request body does not specify a value, the existing privacy setting will be removed and the video will revert to the default privacy setting.\n\nIn addition, not all of those parts contain properties that can be set when setting or updating a video's metadata. For example, the statistics object encapsulates statistics that YouTube calculates for a video and does not contain values that you can set or modify. If the parameter value specifies a part that does not contain mutable values, that part will still be included in the API response.",
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
		if r >= 0x2d && r <= 0x7a || r == '~' {
			return r
		}
		return -1
	}, s)
}
