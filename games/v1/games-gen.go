// Package games provides access to the Google Play Game Services API.
//
// See https://developers.google.com/games/services/
//
// Usage example:
//
//   import "code.google.com/p/google-api-go-client/games/v1"
//   ...
//   gamesService, err := games.New(oauthHttpClient)
package games

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

const apiId = "games:v1"
const apiName = "games"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/games/v1/"

// OAuth2 scopes used by this API.
const (
	// View and manage your game activity
	GamesScope = "https://www.googleapis.com/auth/games"

	// Know your basic profile info and list of people in your circles.
	PlusLoginScope = "https://www.googleapis.com/auth/plus.login"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.AchievementDefinitions = NewAchievementDefinitionsService(s)
	s.Achievements = NewAchievementsService(s)
	s.Applications = NewApplicationsService(s)
	s.Leaderboards = NewLeaderboardsService(s)
	s.Players = NewPlayersService(s)
	s.Revisions = NewRevisionsService(s)
	s.Rooms = NewRoomsService(s)
	s.Scores = NewScoresService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	AchievementDefinitions *AchievementDefinitionsService

	Achievements *AchievementsService

	Applications *ApplicationsService

	Leaderboards *LeaderboardsService

	Players *PlayersService

	Revisions *RevisionsService

	Rooms *RoomsService

	Scores *ScoresService
}

func NewAchievementDefinitionsService(s *Service) *AchievementDefinitionsService {
	rs := &AchievementDefinitionsService{s: s}
	return rs
}

type AchievementDefinitionsService struct {
	s *Service
}

func NewAchievementsService(s *Service) *AchievementsService {
	rs := &AchievementsService{s: s}
	return rs
}

type AchievementsService struct {
	s *Service
}

func NewApplicationsService(s *Service) *ApplicationsService {
	rs := &ApplicationsService{s: s}
	return rs
}

type ApplicationsService struct {
	s *Service
}

func NewLeaderboardsService(s *Service) *LeaderboardsService {
	rs := &LeaderboardsService{s: s}
	return rs
}

type LeaderboardsService struct {
	s *Service
}

func NewPlayersService(s *Service) *PlayersService {
	rs := &PlayersService{s: s}
	return rs
}

type PlayersService struct {
	s *Service
}

func NewRevisionsService(s *Service) *RevisionsService {
	rs := &RevisionsService{s: s}
	return rs
}

type RevisionsService struct {
	s *Service
}

func NewRoomsService(s *Service) *RoomsService {
	rs := &RoomsService{s: s}
	return rs
}

type RoomsService struct {
	s *Service
}

func NewScoresService(s *Service) *ScoresService {
	rs := &ScoresService{s: s}
	return rs
}

type ScoresService struct {
	s *Service
}

type AchievementDefinition struct {
	// AchievementType: The type of the achievement.
	// Possible values are:
	//
	// - "STANDARD" - Achievement is either locked or unlocked.
	// -
	// "INCREMENTAL" - Achievement is incremental.
	AchievementType string `json:"achievementType,omitempty"`

	// Description: The description of the achievement.
	Description string `json:"description,omitempty"`

	// FormattedTotalSteps: The total steps for an incremental achievement
	// as a string.
	FormattedTotalSteps string `json:"formattedTotalSteps,omitempty"`

	// Id: The ID of the achievement.
	Id string `json:"id,omitempty"`

	// InitialState: The initial state of the achievement.
	// Possible values
	// are:
	// - "HIDDEN" - Achievement is hidden.
	// - "REVEALED" -
	// Achievement is revealed.
	// - "UNLOCKED" - Achievement is unlocked.
	InitialState string `json:"initialState,omitempty"`

	// IsRevealedIconUrlDefault: Indicates whether the revealed icon image
	// being returned is a default image, or is provided by the game.
	IsRevealedIconUrlDefault bool `json:"isRevealedIconUrlDefault,omitempty"`

	// IsUnlockedIconUrlDefault: Indicates whether the unlocked icon image
	// being returned is a default image, or is game-provided.
	IsUnlockedIconUrlDefault bool `json:"isUnlockedIconUrlDefault,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#achievementDefinition.
	Kind string `json:"kind,omitempty"`

	// Name: The name of the achievement.
	Name string `json:"name,omitempty"`

	// RevealedIconUrl: The image URL for the revealed achievement icon.
	RevealedIconUrl string `json:"revealedIconUrl,omitempty"`

	// TotalSteps: The total steps for an incremental achievement.
	TotalSteps int64 `json:"totalSteps,omitempty"`

	// UnlockedIconUrl: The image URL for the unlocked achievement icon.
	UnlockedIconUrl string `json:"unlockedIconUrl,omitempty"`
}

type AchievementDefinitionsListResponse struct {
	// Items: The achievement definitions.
	Items []*AchievementDefinition `json:"items,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#achievementDefinitionsListResponse.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token corresponding to the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type AchievementIncrementResponse struct {
	// CurrentSteps: The current steps recorded for this incremental
	// achievement.
	CurrentSteps int64 `json:"currentSteps,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#achievementIncrementResponse.
	Kind string `json:"kind,omitempty"`

	// NewlyUnlocked: Whether the the current steps for the achievement has
	// reached the number of steps required to unlock.
	NewlyUnlocked bool `json:"newlyUnlocked,omitempty"`
}

type AchievementRevealResponse struct {
	// CurrentState: The current state of the achievement for which a reveal
	// was attempted. This might be UNLOCKED if the achievement was already
	// unlocked.
	// Possible values are:
	// - "REVEALED" - Achievement is
	// revealed.
	// - "UNLOCKED" - Achievement is unlocked.
	CurrentState string `json:"currentState,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#achievementRevealResponse.
	Kind string `json:"kind,omitempty"`
}

type AchievementSetStepsAtLeastResponse struct {
	// CurrentSteps: The current steps recorded for this incremental
	// achievement.
	CurrentSteps int64 `json:"currentSteps,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#achievementSetStepsAtLeastResponse.
	Kind string `json:"kind,omitempty"`

	// NewlyUnlocked: Whether the the current steps for the achievement has
	// reached the number of steps required to unlock.
	NewlyUnlocked bool `json:"newlyUnlocked,omitempty"`
}

type AchievementUnlockResponse struct {
	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#achievementUnlockResponse.
	Kind string `json:"kind,omitempty"`

	// NewlyUnlocked: Whether this achievement was newly unlocked (that is,
	// whether the unlock request for the achievement was the first for the
	// player).
	NewlyUnlocked bool `json:"newlyUnlocked,omitempty"`
}

type AchievementUpdateMultipleRequest struct {
	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#achievementUpdateMultipleRequest.
	Kind string `json:"kind,omitempty"`

	// Updates: The individual achievement update requests.
	Updates []*AchievementUpdateRequest `json:"updates,omitempty"`
}

type AchievementUpdateMultipleResponse struct {
	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#achievementUpdateListResponse.
	Kind string `json:"kind,omitempty"`

	// UpdatedAchievements: The updated state of the achievements.
	UpdatedAchievements []*AchievementUpdateResponse `json:"updatedAchievements,omitempty"`
}

type AchievementUpdateRequest struct {
	// AchievementId: The achievement this update is being applied to.
	AchievementId string `json:"achievementId,omitempty"`

	// IncrementPayload: The payload if an update of type INCREMENT was
	// requested for the achievement.
	IncrementPayload *GamesAchievementIncrement `json:"incrementPayload,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#achievementUpdateRequest.
	Kind string `json:"kind,omitempty"`

	// SetStepsAtLeastPayload: The payload if an update of type
	// SET_STEPS_AT_LEAST was requested for the achievement.
	SetStepsAtLeastPayload *GamesAchievementSetStepsAtLeast `json:"setStepsAtLeastPayload,omitempty"`

	// UpdateType: The type of update being applied.
	// Possible values are:
	//
	// - "REVEAL" - Achievement is revealed.
	// - "UNLOCK" - Achievement is
	// unlocked.
	// - "INCREMENT" - Achievement is incremented.
	// -
	// "SET_STEPS_AT_LEAST" - Achievement progress is set to at least the
	// passed value.
	UpdateType string `json:"updateType,omitempty"`
}

type AchievementUpdateResponse struct {
	// AchievementId: The achievement this update is was applied to.
	AchievementId string `json:"achievementId,omitempty"`

	// CurrentState: The current state of the achievement.
	// Possible values
	// are:
	// - "HIDDEN" - Achievement is hidden.
	// - "REVEALED" -
	// Achievement is revealed.
	// - "UNLOCKED" - Achievement is unlocked.
	CurrentState string `json:"currentState,omitempty"`

	// CurrentSteps: The current steps recorded for this achievement if it
	// is incremental.
	CurrentSteps int64 `json:"currentSteps,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#achievementUpdateResponse.
	Kind string `json:"kind,omitempty"`

	// NewlyUnlocked: Whether this achievement was newly unlocked (that is,
	// whether the unlock request for the achievement was the first for the
	// player).
	NewlyUnlocked bool `json:"newlyUnlocked,omitempty"`

	// UpdateOccurred: Whether the requested updates actually affected the
	// achievement.
	UpdateOccurred bool `json:"updateOccurred,omitempty"`
}

type AggregateStats struct {
	// Count: The number of messages sent between a pair of peers.
	Count int64 `json:"count,omitempty,string"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#aggregateStats.
	Kind string `json:"kind,omitempty"`

	// Max: The maximum amount.
	Max int64 `json:"max,omitempty,string"`

	// Min: The minimum amount.
	Min int64 `json:"min,omitempty,string"`

	// Sum: The total number of bytes sent for messages between a pair of
	// peers.
	Sum int64 `json:"sum,omitempty,string"`
}

type AnonymousPlayer struct {
	// AvatarImageUrl: The base URL for the image to display for the
	// anonymous player.
	AvatarImageUrl string `json:"avatarImageUrl,omitempty"`

	// DisplayName: The name to display for the anonymous player.
	DisplayName string `json:"displayName,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#anonymousPlayer.
	Kind string `json:"kind,omitempty"`
}

type Application struct {
	// Achievement_count: The number of achievements visible to the
	// currently authenticated player.
	Achievement_count int64 `json:"achievement_count,omitempty"`

	// Assets: The assets of the application.
	Assets []*ImageAsset `json:"assets,omitempty"`

	// Author: The author of the application.
	Author string `json:"author,omitempty"`

	// Category: The category of the application.
	Category *ApplicationCategory `json:"category,omitempty"`

	// Description: The description of the application.
	Description string `json:"description,omitempty"`

	// Id: The ID of the application.
	Id string `json:"id,omitempty"`

	// Instances: The instances of the application.
	Instances []*Instance `json:"instances,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#application.
	Kind string `json:"kind,omitempty"`

	// LastUpdatedTimestamp: The last updated timestamp of the application.
	LastUpdatedTimestamp int64 `json:"lastUpdatedTimestamp,omitempty,string"`

	// Leaderboard_count: The number of leaderboards visible to the
	// currently authenticated player.
	Leaderboard_count int64 `json:"leaderboard_count,omitempty"`

	// Name: The name of the application.
	Name string `json:"name,omitempty"`
}

type ApplicationCategory struct {
	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#applicationCategory.
	Kind string `json:"kind,omitempty"`

	// Primary: The primary category.
	Primary string `json:"primary,omitempty"`

	// Secondary: The secondary category.
	Secondary string `json:"secondary,omitempty"`
}

type GamesAchievementIncrement struct {
	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#GamesAchievementIncrement.
	Kind string `json:"kind,omitempty"`

	// RequestId: The requestId associated with an increment to an
	// achievement.
	RequestId int64 `json:"requestId,omitempty,string"`

	// Steps: The number of steps to be incremented.
	Steps int64 `json:"steps,omitempty"`
}

type GamesAchievementSetStepsAtLeast struct {
	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#GamesAchievementSetStepsAtLeast.
	Kind string `json:"kind,omitempty"`

	// Steps: The minimum number of steps for the achievement to be set to.
	Steps int64 `json:"steps,omitempty"`
}

type ImageAsset struct {
	// Height: The height of the asset.
	Height int64 `json:"height,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#imageAsset.
	Kind string `json:"kind,omitempty"`

	// Name: The name of the asset.
	Name string `json:"name,omitempty"`

	// Url: The URL of the asset.
	Url string `json:"url,omitempty"`

	// Width: The width of the asset.
	Width int64 `json:"width,omitempty"`
}

type Instance struct {
	// AcquisitionUri: URI which shows where a user can acquire this
	// instance.
	AcquisitionUri string `json:"acquisitionUri,omitempty"`

	// AndroidInstance: Platform dependent details for Android.
	AndroidInstance *InstanceAndroidDetails `json:"androidInstance,omitempty"`

	// IosInstance: Platform dependent details for iOS.
	IosInstance *InstanceIosDetails `json:"iosInstance,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#instance.
	Kind string `json:"kind,omitempty"`

	// Name: Localized display name.
	Name string `json:"name,omitempty"`

	// PlatformType: The platform type.
	// Possible values are:
	// - "ANDROID" -
	// Instance is for Android.
	// - "IOS" - Instance is for iOS
	// - "WEB_APP"
	// - Instance is for Web App.
	PlatformType string `json:"platformType,omitempty"`

	// RealtimePlay: Flag to show if this game instance supports realtime
	// play.
	RealtimePlay bool `json:"realtimePlay,omitempty"`

	// TurnBasedPlay: Flag to show if this game instance supports turn based
	// play.
	TurnBasedPlay bool `json:"turnBasedPlay,omitempty"`

	// WebInstance: Platform dependent details for Web.
	WebInstance *InstanceWebDetails `json:"webInstance,omitempty"`
}

type InstanceAndroidDetails struct {
	// EnablePiracyCheck: Flag indicating whether the anti-piracy check is
	// enabled.
	EnablePiracyCheck bool `json:"enablePiracyCheck,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#instanceAndroidDetails.
	Kind string `json:"kind,omitempty"`

	// PackageName: Android package name which maps to Google Play URL.
	PackageName string `json:"packageName,omitempty"`

	// Preferred: Indicates that this instance is the default for new
	// installations.
	Preferred bool `json:"preferred,omitempty"`
}

type InstanceIosDetails struct {
	// BundleIdentifier: Bundle identifier.
	BundleIdentifier string `json:"bundleIdentifier,omitempty"`

	// ItunesAppId: iTunes App ID.
	ItunesAppId string `json:"itunesAppId,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#instanceIosDetails.
	Kind string `json:"kind,omitempty"`

	// PreferredForIpad: Indicates that this instance is the default for new
	// installations on iPad devices.
	PreferredForIpad bool `json:"preferredForIpad,omitempty"`

	// PreferredForIphone: Indicates that this instance is the default for
	// new installations on iPhone devices.
	PreferredForIphone bool `json:"preferredForIphone,omitempty"`

	// SupportIpad: Flag to indicate if this instance supports iPad.
	SupportIpad bool `json:"supportIpad,omitempty"`

	// SupportIphone: Flag to indicate if this instance supports iPhone.
	SupportIphone bool `json:"supportIphone,omitempty"`
}

type InstanceWebDetails struct {
	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#instanceWebDetails.
	Kind string `json:"kind,omitempty"`

	// LaunchUrl: Launch URL for the game.
	LaunchUrl string `json:"launchUrl,omitempty"`

	// Preferred: Indicates that this instance is the default for new
	// installations.
	Preferred bool `json:"preferred,omitempty"`
}

type Leaderboard struct {
	// IconUrl: The icon for the leaderboard.
	IconUrl string `json:"iconUrl,omitempty"`

	// Id: The leaderboard ID.
	Id string `json:"id,omitempty"`

	// IsIconUrlDefault: Indicates whether the icon image being returned is
	// a default image, or is game-provided.
	IsIconUrlDefault bool `json:"isIconUrlDefault,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#leaderboard.
	Kind string `json:"kind,omitempty"`

	// Name: The name of the leaderboard.
	Name string `json:"name,omitempty"`

	// Order: How scores are ordered.
	// Possible values are:
	// -
	// "LARGER_IS_BETTER" - Larger values are better; scores are sorted in
	// descending order.
	// - "SMALLER_IS_BETTER" - Smaller values are better;
	// scores are sorted in ascending order.
	Order string `json:"order,omitempty"`
}

type LeaderboardEntry struct {
	// FormattedScore: The localized string for the numerical value of this
	// score.
	FormattedScore string `json:"formattedScore,omitempty"`

	// FormattedScoreRank: The localized string for the rank of this score
	// for this leaderboard.
	FormattedScoreRank string `json:"formattedScoreRank,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#leaderboardEntry.
	Kind string `json:"kind,omitempty"`

	// Player: The player who holds this score.
	Player *Player `json:"player,omitempty"`

	// ScoreRank: The rank of this score for this leaderboard.
	ScoreRank int64 `json:"scoreRank,omitempty,string"`

	// ScoreTag: Additional information about the score. Values must contain
	// no more than 64 URI-safe characters as defined by section 2.3 of RFC
	// 3986.
	ScoreTag string `json:"scoreTag,omitempty"`

	// ScoreValue: The numerical value of this score.
	ScoreValue int64 `json:"scoreValue,omitempty,string"`

	// TimeSpan: The time span of this high score.
	// Possible values are:
	// -
	// "ALL_TIME" - The score is an all-time high score.
	// - "WEEKLY" - The
	// score is a weekly high score.
	// - "DAILY" - The score is a daily high
	// score.
	TimeSpan string `json:"timeSpan,omitempty"`

	// WriteTimestampMillis: The timestamp at which this score was recorded,
	// in milliseconds since the epoch in UTC.
	WriteTimestampMillis int64 `json:"writeTimestampMillis,omitempty,string"`
}

type LeaderboardListResponse struct {
	// Items: The leaderboards.
	Items []*Leaderboard `json:"items,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#leaderboardListResponse.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token corresponding to the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type LeaderboardScoreRank struct {
	// FormattedNumScores: The number of scores in the leaderboard as a
	// string.
	FormattedNumScores string `json:"formattedNumScores,omitempty"`

	// FormattedRank: The rank in the leaderboard as a string.
	FormattedRank string `json:"formattedRank,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#leaderboardScoreRank.
	Kind string `json:"kind,omitempty"`

	// NumScores: The number of scores in the leaderboard.
	NumScores int64 `json:"numScores,omitempty,string"`

	// Rank: The rank in the leaderboard.
	Rank int64 `json:"rank,omitempty,string"`
}

type LeaderboardScores struct {
	// Items: The scores in the leaderboard.
	Items []*LeaderboardEntry `json:"items,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#leaderboardScores.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The pagination token for the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// NumScores: The total number of scores in the leaderboard.
	NumScores int64 `json:"numScores,omitempty,string"`

	// PlayerScore: The score of the requesting player on the leaderboard.
	// The player's score may appear both here and in the list of scores
	// above. If you are viewing a public leaderboard and the player is not
	// sharing their gameplay information publicly, the scoreRank and
	// formattedScoreRank values will not be present.
	PlayerScore *LeaderboardEntry `json:"playerScore,omitempty"`

	// PrevPageToken: The pagination token for the previous page of results.
	PrevPageToken string `json:"prevPageToken,omitempty"`
}

type NetworkDiagnostics struct {
	// AndroidNetworkSubtype: The Android network subtype.
	AndroidNetworkSubtype int64 `json:"androidNetworkSubtype,omitempty"`

	// AndroidNetworkType: The Android network type.
	AndroidNetworkType int64 `json:"androidNetworkType,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#networkDiagnostics.
	Kind string `json:"kind,omitempty"`

	// RegistrationLatencyMillis: The amount of time in milliseconds it took
	// for the client to establish a connection with the XMPP server.
	RegistrationLatencyMillis int64 `json:"registrationLatencyMillis,omitempty"`
}

type PeerChannelDiagnostics struct {
	// BytesReceived: Number of bytes received.
	BytesReceived *AggregateStats `json:"bytesReceived,omitempty"`

	// BytesSent: Number of bytes sent.
	BytesSent *AggregateStats `json:"bytesSent,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#peerChannelDiagnostics.
	Kind string `json:"kind,omitempty"`

	// NumMessagesLost: Number of messages lost.
	NumMessagesLost int64 `json:"numMessagesLost,omitempty"`

	// NumMessagesReceived: Number of messages received.
	NumMessagesReceived int64 `json:"numMessagesReceived,omitempty"`

	// NumMessagesSent: Number of messages sent.
	NumMessagesSent int64 `json:"numMessagesSent,omitempty"`

	// NumSendFailures: Number of send failures.
	NumSendFailures int64 `json:"numSendFailures,omitempty"`

	// RoundtripLatencyMillis: Roundtrip latency stats in milliseconds.
	RoundtripLatencyMillis *AggregateStats `json:"roundtripLatencyMillis,omitempty"`
}

type PeerSessionDiagnostics struct {
	// ConnectedTimestampMillis: Connected time in milliseconds.
	ConnectedTimestampMillis int64 `json:"connectedTimestampMillis,omitempty,string"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#peerSessionDiagnostics.
	Kind string `json:"kind,omitempty"`

	// ParticipantId: The participant ID of the peer.
	ParticipantId string `json:"participantId,omitempty"`

	// ReliableChannel: Reliable channel diagnostics.
	ReliableChannel *PeerChannelDiagnostics `json:"reliableChannel,omitempty"`

	// UnreliableChannel: Unreliable channel diagnostics.
	UnreliableChannel *PeerChannelDiagnostics `json:"unreliableChannel,omitempty"`
}

type Player struct {
	// AvatarImageUrl: The base URL for the image that represents the
	// player.
	AvatarImageUrl string `json:"avatarImageUrl,omitempty"`

	// DisplayName: The name to display for the player.
	DisplayName string `json:"displayName,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#player.
	Kind string `json:"kind,omitempty"`

	// PlayerId: The ID of the player.
	PlayerId string `json:"playerId,omitempty"`
}

type PlayerAchievement struct {
	// AchievementState: The state of the achievement.
	// Possible values are:
	//
	// - "HIDDEN" - Achievement is hidden.
	// - "REVEALED" - Achievement is
	// revealed.
	// - "UNLOCKED" - Achievement is unlocked.
	AchievementState string `json:"achievementState,omitempty"`

	// CurrentSteps: The current steps for an incremental achievement.
	CurrentSteps int64 `json:"currentSteps,omitempty"`

	// FormattedCurrentStepsString: The current steps for an incremental
	// achievement as a string.
	FormattedCurrentStepsString string `json:"formattedCurrentStepsString,omitempty"`

	// Id: The ID of the achievement.
	Id string `json:"id,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#playerAchievement.
	Kind string `json:"kind,omitempty"`

	// LastUpdatedTimestamp: The timestamp of the last modification to this
	// achievement's state.
	LastUpdatedTimestamp int64 `json:"lastUpdatedTimestamp,omitempty,string"`
}

type PlayerAchievementListResponse struct {
	// Items: The achievements.
	Items []*PlayerAchievement `json:"items,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#playerAchievementListResponse.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token corresponding to the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type PlayerLeaderboardScore struct {
	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#playerLeaderboardScore.
	Kind string `json:"kind,omitempty"`

	// Leaderboard_id: The ID of the leaderboard this score is in.
	Leaderboard_id string `json:"leaderboard_id,omitempty"`

	// PublicRank: The public rank of the score in this leaderboard. This
	// object will not be present if the user is not sharing their scores
	// publicly.
	PublicRank *LeaderboardScoreRank `json:"publicRank,omitempty"`

	// ScoreString: The formatted value of this score.
	ScoreString string `json:"scoreString,omitempty"`

	// ScoreTag: Additional information about the score. Values must contain
	// no more than 64 URI-safe characters as defined by section 2.3 of RFC
	// 3986.
	ScoreTag string `json:"scoreTag,omitempty"`

	// ScoreValue: The numerical value of this score.
	ScoreValue int64 `json:"scoreValue,omitempty,string"`

	// SocialRank: The social rank of the score in this leaderboard.
	SocialRank *LeaderboardScoreRank `json:"socialRank,omitempty"`

	// TimeSpan: The time span of this score.
	// Possible values are:
	// -
	// "ALL_TIME" - The score is an all-time score.
	// - "WEEKLY" - The score
	// is a weekly score.
	// - "DAILY" - The score is a daily score.
	TimeSpan string `json:"timeSpan,omitempty"`

	// WriteTimestamp: The timestamp at which this score was recorded, in
	// milliseconds since the epoch in UTC.
	WriteTimestamp int64 `json:"writeTimestamp,omitempty,string"`
}

type PlayerLeaderboardScoreListResponse struct {
	// Items: The leaderboard scores.
	Items []*PlayerLeaderboardScore `json:"items,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#playerLeaderboardScoreListResponse.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The pagination token for the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Player: The Player resources for the owner of this score.
	Player *Player `json:"player,omitempty"`
}

type PlayerScore struct {
	// FormattedScore: The formatted score for this player score.
	FormattedScore string `json:"formattedScore,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#playerScore.
	Kind string `json:"kind,omitempty"`

	// Score: The numerical value for this player score.
	Score int64 `json:"score,omitempty,string"`

	// ScoreTag: Additional information about this score. Values will
	// contain no more than 64 URI-safe characters as defined by section 2.3
	// of RFC 3986.
	ScoreTag string `json:"scoreTag,omitempty"`

	// TimeSpan: The time span for this player score.
	// Possible values are:
	//
	// - "ALL_TIME" - The score is an all-time score.
	// - "WEEKLY" - The
	// score is a weekly score.
	// - "DAILY" - The score is a daily score.
	TimeSpan string `json:"timeSpan,omitempty"`
}

type PlayerScoreListResponse struct {
	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#playerScoreListResponse.
	Kind string `json:"kind,omitempty"`

	// SubmittedScores: The score submissions statuses.
	SubmittedScores []*PlayerScoreResponse `json:"submittedScores,omitempty"`
}

type PlayerScoreResponse struct {
	// BeatenScoreTimeSpans: The time spans where the submitted score is
	// better than the existing score for that time span.
	// Possible values
	// are:
	// - "ALL_TIME" - The score is an all-time score.
	// - "WEEKLY" -
	// The score is a weekly score.
	// - "DAILY" - The score is a daily score.
	BeatenScoreTimeSpans []string `json:"beatenScoreTimeSpans,omitempty"`

	// FormattedScore: The formatted value of the submitted score.
	FormattedScore string `json:"formattedScore,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#playerScoreResponse.
	Kind string `json:"kind,omitempty"`

	// LeaderboardId: The leaderboard ID that this score was submitted to.
	LeaderboardId string `json:"leaderboardId,omitempty"`

	// ScoreTag: Additional information about this score. Values will
	// contain no more than 64 URI-safe characters as defined by section 2.3
	// of RFC 3986.
	ScoreTag string `json:"scoreTag,omitempty"`

	// UnbeatenScores: The scores in time spans that have not been beaten.
	// As an example, the submitted score may be better than the player's
	// DAILY score, but not better than the player's scores for the WEEKLY
	// or ALL_TIME time spans.
	UnbeatenScores []*PlayerScore `json:"unbeatenScores,omitempty"`
}

type PlayerScoreSubmissionList struct {
	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#playerScoreSubmissionList.
	Kind string `json:"kind,omitempty"`

	// Scores: The score submissions.
	Scores []*ScoreSubmission `json:"scores,omitempty"`
}

type RevisionCheckResponse struct {
	// ApiVersion: The version of the API this client revision should use
	// when calling API methods.
	ApiVersion string `json:"apiVersion,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#revisionCheckResponse.
	Kind string `json:"kind,omitempty"`

	// RevisionStatus: The result of the revision check.
	// Possible values
	// are:
	// - "OK" - The revision being used is current.
	// - "DEPRECATED" -
	// There is currently a newer version available, but the revision being
	// used still works.
	// - "INVALID" - The revision being used is not
	// supported in any released version.
	RevisionStatus string `json:"revisionStatus,omitempty"`
}

type Room struct {
	// ApplicationId: The ID of the application being played.
	ApplicationId string `json:"applicationId,omitempty"`

	// AutoMatchingCriteria: Criteria for auto-matching players into this
	// room.
	AutoMatchingCriteria *RoomAutoMatchingCriteria `json:"autoMatchingCriteria,omitempty"`

	// AutoMatchingStatus: Auto-matching status for this room. Not set if
	// the room is not currently in the auto-matching queue.
	AutoMatchingStatus *RoomAutoMatchStatus `json:"autoMatchingStatus,omitempty"`

	// CreationDetails: Details about the room creation.
	CreationDetails *RoomModification `json:"creationDetails,omitempty"`

	// Description: This short description is generated by our servers and
	// worded relative to the player requesting the room. It is intended to
	// be displayed when the room is shown in a list (that is, an invitation
	// to a room.)
	Description string `json:"description,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#room.
	Kind string `json:"kind,omitempty"`

	// LastUpdateDetails: Details about the last update to the room.
	LastUpdateDetails *RoomModification `json:"lastUpdateDetails,omitempty"`

	// Participants: The participants involved in the room, along with their
	// statuses. Includes participants who have left or declined
	// invitations.
	Participants []*RoomParticipant `json:"participants,omitempty"`

	// RoomId: Globally unique ID for a room.
	RoomId string `json:"roomId,omitempty"`

	// RoomStatusVersion: The version of the room status: an increasing
	// counter, used by the client to ignore out-of-order updates to room
	// status.
	RoomStatusVersion int64 `json:"roomStatusVersion,omitempty"`

	// Status: The status of the room.
	// Possible values are:
	// -
	// "ROOM_INVITING" - One or more players have been invited and not
	// responded.
	// - "ROOM_AUTO_MATCHING" - One or more slots need to be
	// filled by auto-matching.
	// - "ROOM_CONNECTING" - Players have joined
	// and are connecting to each other (either before or after
	// auto-matching).
	// - "ROOM_ACTIVE" - All players have joined and
	// connected to each other.
	// - "ROOM_DELETED" - The room should no
	// longer be shown on the client. Returned in sync calls when a player
	// joins a room (as a tombstone), or for rooms where all joined
	// participants have left.
	Status string `json:"status,omitempty"`

	// Variant: The variant / mode of the application being played; can be
	// any integer value, or left blank.
	Variant int64 `json:"variant,omitempty"`
}

type RoomAutoMatchStatus struct {
	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#roomAutoMatchStatus.
	Kind string `json:"kind,omitempty"`

	// WaitEstimateSeconds: An estimate for the amount of time (in seconds)
	// that auto-matching is expected to take to complete.
	WaitEstimateSeconds int64 `json:"waitEstimateSeconds,omitempty"`
}

type RoomAutoMatchingCriteria struct {
	// ExclusiveBitmask: A bitmask indicating when auto-matches are valid.
	// When ANDed with other exclusive bitmasks, the result must be zero.
	// Can be used to support exclusive roles within a game.
	ExclusiveBitmask int64 `json:"exclusiveBitmask,omitempty,string"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#roomAutoMatchingCriteria.
	Kind string `json:"kind,omitempty"`

	// MaxAutoMatchingPlayers: The maximum number of players that should be
	// added to the room by auto-matching.
	MaxAutoMatchingPlayers int64 `json:"maxAutoMatchingPlayers,omitempty"`

	// MinAutoMatchingPlayers: The minimum number of players that should be
	// added to the room by auto-matching.
	MinAutoMatchingPlayers int64 `json:"minAutoMatchingPlayers,omitempty"`
}

type RoomClientAddress struct {
	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#roomClientAddress.
	Kind string `json:"kind,omitempty"`

	// XmppAddress: The XMPP address of the client on the Google Games XMPP
	// network.
	XmppAddress string `json:"xmppAddress,omitempty"`
}

type RoomCreateRequest struct {
	// AutoMatchingCriteria: Criteria for auto-matching players into this
	// room.
	AutoMatchingCriteria *RoomAutoMatchingCriteria `json:"autoMatchingCriteria,omitempty"`

	// Capabilities: The capabilities that this client supports for realtime
	// communication.
	Capabilities []string `json:"capabilities,omitempty"`

	// ClientAddress: Client address for the player creating the room.
	ClientAddress *RoomClientAddress `json:"clientAddress,omitempty"`

	// InvitedPlayerIds: The player IDs to invite to the room.
	InvitedPlayerIds []string `json:"invitedPlayerIds,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#roomCreateRequest.
	Kind string `json:"kind,omitempty"`

	// NetworkDiagnostics: Network diagnostics for the client creating the
	// room.
	NetworkDiagnostics *NetworkDiagnostics `json:"networkDiagnostics,omitempty"`

	// Variant: The variant / mode of the application to be played. This can
	// be any integer value, or left blank. You should use a small number of
	// variants to keep the auto-matching pool as large as possible.
	Variant int64 `json:"variant,omitempty"`
}

type RoomJoinRequest struct {
	// Capabilities: The capabilities that this client supports for realtime
	// communication.
	Capabilities []string `json:"capabilities,omitempty"`

	// ClientAddress: Client address for the player joining the room.
	ClientAddress *RoomClientAddress `json:"clientAddress,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#roomJoinRequest.
	Kind string `json:"kind,omitempty"`

	// NetworkDiagnostics: Network diagnostics for the client joining the
	// room.
	NetworkDiagnostics *NetworkDiagnostics `json:"networkDiagnostics,omitempty"`
}

type RoomLeaveDiagnostics struct {
	// AndroidNetworkSubtype: Android network subtype.
	// http://developer.android.com/reference/android/net/NetworkInfo.html#ge
	// tSubtype()
	AndroidNetworkSubtype int64 `json:"androidNetworkSubtype,omitempty"`

	// AndroidNetworkType: Android network type.
	// http://developer.android.com/reference/android/net/NetworkInfo.html#ge
	// tType()
	AndroidNetworkType int64 `json:"androidNetworkType,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#roomLeaveDiagnostics.
	Kind string `json:"kind,omitempty"`

	// PeerSession: Diagnostics about all peer sessions.
	PeerSession []*PeerSessionDiagnostics `json:"peerSession,omitempty"`

	// SocketsUsed: Whether or not sockets were used.
	SocketsUsed bool `json:"socketsUsed,omitempty"`
}

type RoomLeaveRequest struct {
	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#roomLeaveRequest.
	Kind string `json:"kind,omitempty"`

	// LeaveDiagnostics: Diagnostics for a player leaving the room.
	LeaveDiagnostics *RoomLeaveDiagnostics `json:"leaveDiagnostics,omitempty"`

	// Reason: Reason for leaving the match.
	// Possible values are:
	// -
	// "PLAYER_LEFT" - The player chose to leave the room..
	// - "GAME_LEFT" -
	// The game chose to remove the player from the room.
	// -
	// "REALTIME_ABANDONED" - The player switched to another application and
	// abandoned the room.
	// - "REALTIME_PEER_CONNECTION_FAILURE" - The
	// client was unable to establish a connection to other peer(s).
	// -
	// "REALTIME_SERVER_CONNECTION_FAILURE" - The client was unable to
	// communicate with the server.
	// - "REALTIME_SERVER_ERROR" - The client
	// received an error response when it tried to communicate with the
	// server.
	// - "REALTIME_TIMEOUT" - The client timed out while waiting
	// for a room.
	Reason string `json:"reason,omitempty"`
}

type RoomList struct {
	// Items: The rooms.
	Items []*Room `json:"items,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#roomList.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The pagination token for the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type RoomModification struct {
	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#roomModification.
	Kind string `json:"kind,omitempty"`

	// ModifiedTimestampMillis: The timestamp at which they modified the
	// room, in milliseconds since the epoch in UTC.
	ModifiedTimestampMillis int64 `json:"modifiedTimestampMillis,omitempty,string"`

	// ParticipantId: The ID of the participant that modified the room.
	ParticipantId string `json:"participantId,omitempty"`
}

type RoomP2PStatus struct {
	// ConnectionSetupLatencyMillis: The amount of time in milliseconds it
	// took to establish connections with this peer.
	ConnectionSetupLatencyMillis int64 `json:"connectionSetupLatencyMillis,omitempty"`

	// Error: The error code in event of a failure.
	// Possible values are:
	// -
	// "P2P_FAILED" - The client failed to establish a P2P connection with
	// the peer.
	// - "PRESENCE_FAILED" - The client failed to register to
	// receive P2P connections.
	// - "RELAY_SERVER_FAILED" - The client
	// received an error when trying to use the relay server to establish a
	// P2P connection with the peer.
	Error string `json:"error,omitempty"`

	// Error_reason: More detailed diagnostic message returned in event of a
	// failure.
	Error_reason string `json:"error_reason,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#roomP2PStatus.
	Kind string `json:"kind,omitempty"`

	// ParticipantId: The ID of the participant.
	ParticipantId string `json:"participantId,omitempty"`

	// Status: The status of the peer in the room.
	// Possible values are:
	// -
	// "CONNECTION_ESTABLISHED" - The client established a P2P connection
	// with the peer.
	// - "CONNECTION_FAILED" - The client failed to
	// establish directed presence with the peer.
	Status string `json:"status,omitempty"`

	// UnreliableRoundtripLatencyMillis: The amount of time in milliseconds
	// it took to send packets back and forth on the unreliable channel with
	// this peer.
	UnreliableRoundtripLatencyMillis int64 `json:"unreliableRoundtripLatencyMillis,omitempty"`
}

type RoomP2PStatuses struct {
	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#roomP2PStatuses.
	Kind string `json:"kind,omitempty"`

	// Updates: The updates for the peers.
	Updates []*RoomP2PStatus `json:"updates,omitempty"`
}

type RoomParticipant struct {
	// AutoMatchedPlayer: Information about a player that has been
	// anonymously auto-matched against the requesting player. (Either
	// player or autoMatchedPlayer will be set.)
	AutoMatchedPlayer *AnonymousPlayer `json:"autoMatchedPlayer,omitempty"`

	// Capabilities: The capabilities which can be used when communicating
	// with this participant.
	Capabilities []string `json:"capabilities,omitempty"`

	// ClientAddress: Client address for the participant.
	ClientAddress *RoomClientAddress `json:"clientAddress,omitempty"`

	// Connected: True if this participant is in the fully connected set of
	// peers in the room.
	Connected bool `json:"connected,omitempty"`

	// Id: An identifier for the participant in the scope of the room.
	// Cannot be used to identify a player across rooms or in other
	// contexts.
	Id string `json:"id,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#roomParticipant.
	Kind string `json:"kind,omitempty"`

	// LeaveReason: The reason the participant left the room; populated if
	// the participant status is PARTICIPANT_LEFT.
	// Possible values are:
	// -
	// "PLAYER_LEFT" - The player explicitly chose to leave the room.
	// -
	// "GAME_LEFT" - The game chose to remove the player from the room.
	// -
	// "ABANDONED" - The player switched to another application and
	// abandoned the room.
	// - "PEER_CONNECTION_FAILURE" - The client was
	// unable to establish or maintain a connection to other peer(s) in the
	// room.
	// - "SERVER_ERROR" - The client received an error response when
	// it tried to communicate with the server.
	// - "TIMEOUT" - The client
	// timed out while waiting for players to join and connect.
	// -
	// "PRESENCE_FAILURE" - The client's XMPP connection ended abruptly.
	LeaveReason string `json:"leaveReason,omitempty"`

	// Player: Information about the player. Not populated if this player
	// was anonymously auto-matched against the requesting player. (Either
	// player or autoMatchedPlayer will be set.)
	Player *Player `json:"player,omitempty"`

	// Status: The status of the participant with respect to the
	// room.
	// Possible values are:
	// - "PARTICIPANT_INVITED" - The
	// participant has been invited to join the room, but has not yet
	// responded.
	// - "PARTICIPANT_JOINED" - The participant has joined the
	// room (either after creating it or accepting an invitation.)
	// -
	// "PARTICIPANT_DECLINED" - The participant declined an invitation to
	// join the room.
	// - "PARTICIPANT_LEFT" - The participant joined the
	// room and then left it.
	Status string `json:"status,omitempty"`
}

type RoomStatus struct {
	// AutoMatchingStatus: Auto-matching status for this room. Not set if
	// the room is not currently in the automatching queue.
	AutoMatchingStatus *RoomAutoMatchStatus `json:"autoMatchingStatus,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#roomStatus.
	Kind string `json:"kind,omitempty"`

	// Participants: The participants involved in the room, along with their
	// statuses. Includes participants who have left or declined
	// invitations.
	Participants []*RoomParticipant `json:"participants,omitempty"`

	// RoomId: Globally unique ID for a room.
	RoomId string `json:"roomId,omitempty"`

	// Status: The status of the room.
	// Possible values are:
	// -
	// "ROOM_INVITING" - One or more players have been invited and not
	// responded.
	// - "ROOM_AUTO_MATCHING" - One or more slots need to be
	// filled by auto-matching.
	// - "ROOM_CONNECTING" - Players have joined
	// are connecting to each other (either before or after auto-matching).
	//
	// - "ROOM_ACTIVE" - All players have joined and connected to each
	// other.
	// - "ROOM_DELETED" - All joined players have left.
	Status string `json:"status,omitempty"`

	// StatusVersion: The version of the status for the room: an increasing
	// counter, used by the client to ignore out-of-order updates to room
	// status.
	StatusVersion int64 `json:"statusVersion,omitempty"`
}

type ScoreSubmission struct {
	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#scoreSubmission.
	Kind string `json:"kind,omitempty"`

	// LeaderboardId: The leaderboard this score is being submitted to.
	LeaderboardId string `json:"leaderboardId,omitempty"`

	// Score: The new score being submitted.
	Score int64 `json:"score,omitempty,string"`

	// ScoreTag: Additional information about this score. Values will
	// contain no more than 64 URI-safe characters as defined by section 2.3
	// of RFC 3986.
	ScoreTag string `json:"scoreTag,omitempty"`
}

// method id "games.achievementDefinitions.list":

type AchievementDefinitionsListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Lists all the achievement definitions for your application.
func (r *AchievementDefinitionsService) List() *AchievementDefinitionsListCall {
	c := &AchievementDefinitionsListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// Language sets the optional parameter "language": The preferred
// language to use for strings returned by this method.
func (c *AchievementDefinitionsListCall) Language(language string) *AchievementDefinitionsListCall {
	c.opt_["language"] = language
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of achievement resources to return in the response, used for
// paging. For any response, the actual number of achievement resources
// returned may be less than the specified maxResults.
func (c *AchievementDefinitionsListCall) MaxResults(maxResults int64) *AchievementDefinitionsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": The token returned
// by the previous request.
func (c *AchievementDefinitionsListCall) PageToken(pageToken string) *AchievementDefinitionsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

func (c *AchievementDefinitionsListCall) Do() (*AchievementDefinitionsListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["language"]; ok {
		params.Set("language", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/games/v1/", "achievements")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(AchievementDefinitionsListResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all the achievement definitions for your application.",
	//   "httpMethod": "GET",
	//   "id": "games.achievementDefinitions.list",
	//   "parameters": {
	//     "language": {
	//       "description": "The preferred language to use for strings returned by this method.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of achievement resources to return in the response, used for paging. For any response, the actual number of achievement resources returned may be less than the specified maxResults.",
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
	//   "path": "achievements",
	//   "response": {
	//     "$ref": "AchievementDefinitionsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "games.achievements.increment":

type AchievementsIncrementCall struct {
	s                *Service
	achievementId    string
	stepsToIncrement int64
	opt_             map[string]interface{}
}

// Increment: Increments the steps of the achievement with the given ID
// for the currently authenticated player.
func (r *AchievementsService) Increment(achievementId string, stepsToIncrement int64) *AchievementsIncrementCall {
	c := &AchievementsIncrementCall{s: r.s, opt_: make(map[string]interface{})}
	c.achievementId = achievementId
	c.stepsToIncrement = stepsToIncrement
	return c
}

// RequestId sets the optional parameter "requestId": A randomly
// generated numeric ID for each request specified by the caller. This
// number is used at the server to ensure that the request is handled
// correctly across retries.
func (c *AchievementsIncrementCall) RequestId(requestId int64) *AchievementsIncrementCall {
	c.opt_["requestId"] = requestId
	return c
}

func (c *AchievementsIncrementCall) Do() (*AchievementIncrementResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("stepsToIncrement", fmt.Sprintf("%v", c.stepsToIncrement))
	if v, ok := c.opt_["requestId"]; ok {
		params.Set("requestId", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/games/v1/", "achievements/{achievementId}/increment")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{achievementId}", url.QueryEscape(c.achievementId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(AchievementIncrementResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Increments the steps of the achievement with the given ID for the currently authenticated player.",
	//   "httpMethod": "POST",
	//   "id": "games.achievements.increment",
	//   "parameterOrder": [
	//     "achievementId",
	//     "stepsToIncrement"
	//   ],
	//   "parameters": {
	//     "achievementId": {
	//       "description": "The ID of the achievement used by this method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "requestId": {
	//       "description": "A randomly generated numeric ID for each request specified by the caller. This number is used at the server to ensure that the request is handled correctly across retries.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "stepsToIncrement": {
	//       "description": "The number of steps to increment.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "required": true,
	//       "type": "integer"
	//     }
	//   },
	//   "path": "achievements/{achievementId}/increment",
	//   "response": {
	//     "$ref": "AchievementIncrementResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "games.achievements.list":

type AchievementsListCall struct {
	s        *Service
	playerId string
	opt_     map[string]interface{}
}

// List: Lists the progress for all your application's achievements for
// the currently authenticated player.
func (r *AchievementsService) List(playerId string) *AchievementsListCall {
	c := &AchievementsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.playerId = playerId
	return c
}

// Language sets the optional parameter "language": The preferred
// language to use for strings returned by this method.
func (c *AchievementsListCall) Language(language string) *AchievementsListCall {
	c.opt_["language"] = language
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of achievement resources to return in the response, used for
// paging. For any response, the actual number of achievement resources
// returned may be less than the specified maxResults.
func (c *AchievementsListCall) MaxResults(maxResults int64) *AchievementsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": The token returned
// by the previous request.
func (c *AchievementsListCall) PageToken(pageToken string) *AchievementsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// State sets the optional parameter "state": Tells the server to return
// only achievements with the specified state. If this parameter isn't
// specified, all achievements are returned.
func (c *AchievementsListCall) State(state string) *AchievementsListCall {
	c.opt_["state"] = state
	return c
}

func (c *AchievementsListCall) Do() (*PlayerAchievementListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["language"]; ok {
		params.Set("language", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["state"]; ok {
		params.Set("state", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/games/v1/", "players/{playerId}/achievements")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{playerId}", url.QueryEscape(c.playerId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(PlayerAchievementListResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the progress for all your application's achievements for the currently authenticated player.",
	//   "httpMethod": "GET",
	//   "id": "games.achievements.list",
	//   "parameterOrder": [
	//     "playerId"
	//   ],
	//   "parameters": {
	//     "language": {
	//       "description": "The preferred language to use for strings returned by this method.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of achievement resources to return in the response, used for paging. For any response, the actual number of achievement resources returned may be less than the specified maxResults.",
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
	//     },
	//     "playerId": {
	//       "description": "A player ID. A value of me may be used in place of the authenticated player's ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "state": {
	//       "description": "Tells the server to return only achievements with the specified state. If this parameter isn't specified, all achievements are returned.",
	//       "enum": [
	//         "ALL",
	//         "HIDDEN",
	//         "REVEALED",
	//         "UNLOCKED"
	//       ],
	//       "enumDescriptions": [
	//         "List all achievements. This is the default.",
	//         "List only hidden achievements.",
	//         "List only revealed achievements.",
	//         "List only unlocked achievements."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "players/{playerId}/achievements",
	//   "response": {
	//     "$ref": "PlayerAchievementListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "games.achievements.reveal":

type AchievementsRevealCall struct {
	s             *Service
	achievementId string
	opt_          map[string]interface{}
}

// Reveal: Sets the state of the achievement with the given ID to
// REVEALED for the currently authenticated player.
func (r *AchievementsService) Reveal(achievementId string) *AchievementsRevealCall {
	c := &AchievementsRevealCall{s: r.s, opt_: make(map[string]interface{})}
	c.achievementId = achievementId
	return c
}

func (c *AchievementsRevealCall) Do() (*AchievementRevealResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/games/v1/", "achievements/{achievementId}/reveal")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{achievementId}", url.QueryEscape(c.achievementId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(AchievementRevealResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Sets the state of the achievement with the given ID to REVEALED for the currently authenticated player.",
	//   "httpMethod": "POST",
	//   "id": "games.achievements.reveal",
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
	//   "path": "achievements/{achievementId}/reveal",
	//   "response": {
	//     "$ref": "AchievementRevealResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "games.achievements.setStepsAtLeast":

type AchievementsSetStepsAtLeastCall struct {
	s             *Service
	achievementId string
	steps         int64
	opt_          map[string]interface{}
}

// SetStepsAtLeast: Sets the steps for the currently authenticated
// player towards unlocking an achievement. If the steps parameter is
// less than the current number of steps that the player already gained
// for the achievement, the achievement is not modified.
func (r *AchievementsService) SetStepsAtLeast(achievementId string, steps int64) *AchievementsSetStepsAtLeastCall {
	c := &AchievementsSetStepsAtLeastCall{s: r.s, opt_: make(map[string]interface{})}
	c.achievementId = achievementId
	c.steps = steps
	return c
}

func (c *AchievementsSetStepsAtLeastCall) Do() (*AchievementSetStepsAtLeastResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("steps", fmt.Sprintf("%v", c.steps))
	urls := googleapi.ResolveRelative("https://www.googleapis.com/games/v1/", "achievements/{achievementId}/setStepsAtLeast")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{achievementId}", url.QueryEscape(c.achievementId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(AchievementSetStepsAtLeastResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Sets the steps for the currently authenticated player towards unlocking an achievement. If the steps parameter is less than the current number of steps that the player already gained for the achievement, the achievement is not modified.",
	//   "httpMethod": "POST",
	//   "id": "games.achievements.setStepsAtLeast",
	//   "parameterOrder": [
	//     "achievementId",
	//     "steps"
	//   ],
	//   "parameters": {
	//     "achievementId": {
	//       "description": "The ID of the achievement used by this method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "steps": {
	//       "description": "The minimum value to set the steps to.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "required": true,
	//       "type": "integer"
	//     }
	//   },
	//   "path": "achievements/{achievementId}/setStepsAtLeast",
	//   "response": {
	//     "$ref": "AchievementSetStepsAtLeastResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "games.achievements.unlock":

type AchievementsUnlockCall struct {
	s             *Service
	achievementId string
	opt_          map[string]interface{}
}

// Unlock: Unlocks this achievement for the currently authenticated
// player.
func (r *AchievementsService) Unlock(achievementId string) *AchievementsUnlockCall {
	c := &AchievementsUnlockCall{s: r.s, opt_: make(map[string]interface{})}
	c.achievementId = achievementId
	return c
}

func (c *AchievementsUnlockCall) Do() (*AchievementUnlockResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/games/v1/", "achievements/{achievementId}/unlock")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{achievementId}", url.QueryEscape(c.achievementId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(AchievementUnlockResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Unlocks this achievement for the currently authenticated player.",
	//   "httpMethod": "POST",
	//   "id": "games.achievements.unlock",
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
	//   "path": "achievements/{achievementId}/unlock",
	//   "response": {
	//     "$ref": "AchievementUnlockResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "games.achievements.updateMultiple":

type AchievementsUpdateMultipleCall struct {
	s                                *Service
	achievementupdatemultiplerequest *AchievementUpdateMultipleRequest
	opt_                             map[string]interface{}
}

// UpdateMultiple: Updates multiple achievements for the currently
// authenticated player.
func (r *AchievementsService) UpdateMultiple(achievementupdatemultiplerequest *AchievementUpdateMultipleRequest) *AchievementsUpdateMultipleCall {
	c := &AchievementsUpdateMultipleCall{s: r.s, opt_: make(map[string]interface{})}
	c.achievementupdatemultiplerequest = achievementupdatemultiplerequest
	return c
}

func (c *AchievementsUpdateMultipleCall) Do() (*AchievementUpdateMultipleResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.achievementupdatemultiplerequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/games/v1/", "achievements/updateMultiple")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(AchievementUpdateMultipleResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates multiple achievements for the currently authenticated player.",
	//   "httpMethod": "POST",
	//   "id": "games.achievements.updateMultiple",
	//   "path": "achievements/updateMultiple",
	//   "request": {
	//     "$ref": "AchievementUpdateMultipleRequest"
	//   },
	//   "response": {
	//     "$ref": "AchievementUpdateMultipleResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "games.applications.get":

type ApplicationsGetCall struct {
	s             *Service
	applicationId string
	opt_          map[string]interface{}
}

// Get: Retrieves the metadata of the application with the given ID. If
// the requested application is not available for the specified
// platformType, the returned response will not include any instance
// data.
func (r *ApplicationsService) Get(applicationId string) *ApplicationsGetCall {
	c := &ApplicationsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.applicationId = applicationId
	return c
}

// Language sets the optional parameter "language": The preferred
// language to use for strings returned by this method.
func (c *ApplicationsGetCall) Language(language string) *ApplicationsGetCall {
	c.opt_["language"] = language
	return c
}

// PlatformType sets the optional parameter "platformType": Restrict
// application details returned to the specific platform.
func (c *ApplicationsGetCall) PlatformType(platformType string) *ApplicationsGetCall {
	c.opt_["platformType"] = platformType
	return c
}

func (c *ApplicationsGetCall) Do() (*Application, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["language"]; ok {
		params.Set("language", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["platformType"]; ok {
		params.Set("platformType", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/games/v1/", "applications/{applicationId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{applicationId}", url.QueryEscape(c.applicationId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Application)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the metadata of the application with the given ID. If the requested application is not available for the specified platformType, the returned response will not include any instance data.",
	//   "httpMethod": "GET",
	//   "id": "games.applications.get",
	//   "parameterOrder": [
	//     "applicationId"
	//   ],
	//   "parameters": {
	//     "applicationId": {
	//       "description": "The application being requested.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "language": {
	//       "description": "The preferred language to use for strings returned by this method.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "platformType": {
	//       "description": "Restrict application details returned to the specific platform.",
	//       "enum": [
	//         "ANDROID",
	//         "IOS",
	//         "WEB_APP"
	//       ],
	//       "enumDescriptions": [
	//         "Retrieve applications that can be played on Android.",
	//         "Retrieve applications that can be played on iOS.",
	//         "Retrieve applications that can be played on desktop web."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "applications/{applicationId}",
	//   "response": {
	//     "$ref": "Application"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "games.applications.played":

type ApplicationsPlayedCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// Played: Indicate that the the currently authenticated user is playing
// your application.
func (r *ApplicationsService) Played() *ApplicationsPlayedCall {
	c := &ApplicationsPlayedCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

func (c *ApplicationsPlayedCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/games/v1/", "applications/played")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Indicate that the the currently authenticated user is playing your application.",
	//   "httpMethod": "POST",
	//   "id": "games.applications.played",
	//   "path": "applications/played",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "games.leaderboards.get":

type LeaderboardsGetCall struct {
	s             *Service
	leaderboardId string
	opt_          map[string]interface{}
}

// Get: Retrieves the metadata of the leaderboard with the given ID.
func (r *LeaderboardsService) Get(leaderboardId string) *LeaderboardsGetCall {
	c := &LeaderboardsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.leaderboardId = leaderboardId
	return c
}

// Language sets the optional parameter "language": The preferred
// language to use for strings returned by this method.
func (c *LeaderboardsGetCall) Language(language string) *LeaderboardsGetCall {
	c.opt_["language"] = language
	return c
}

func (c *LeaderboardsGetCall) Do() (*Leaderboard, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["language"]; ok {
		params.Set("language", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/games/v1/", "leaderboards/{leaderboardId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{leaderboardId}", url.QueryEscape(c.leaderboardId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Leaderboard)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the metadata of the leaderboard with the given ID.",
	//   "httpMethod": "GET",
	//   "id": "games.leaderboards.get",
	//   "parameterOrder": [
	//     "leaderboardId"
	//   ],
	//   "parameters": {
	//     "language": {
	//       "description": "The preferred language to use for strings returned by this method.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "leaderboardId": {
	//       "description": "The ID of the leaderboard.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "leaderboards/{leaderboardId}",
	//   "response": {
	//     "$ref": "Leaderboard"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "games.leaderboards.list":

type LeaderboardsListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Lists all the leaderboard metadata for your application.
func (r *LeaderboardsService) List() *LeaderboardsListCall {
	c := &LeaderboardsListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// Language sets the optional parameter "language": The preferred
// language to use for strings returned by this method.
func (c *LeaderboardsListCall) Language(language string) *LeaderboardsListCall {
	c.opt_["language"] = language
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of leaderboards to return in the response. For any response,
// the actual number of leaderboards returned may be less than the
// specified maxResults.
func (c *LeaderboardsListCall) MaxResults(maxResults int64) *LeaderboardsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": The token returned
// by the previous request.
func (c *LeaderboardsListCall) PageToken(pageToken string) *LeaderboardsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

func (c *LeaderboardsListCall) Do() (*LeaderboardListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["language"]; ok {
		params.Set("language", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/games/v1/", "leaderboards")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(LeaderboardListResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all the leaderboard metadata for your application.",
	//   "httpMethod": "GET",
	//   "id": "games.leaderboards.list",
	//   "parameters": {
	//     "language": {
	//       "description": "The preferred language to use for strings returned by this method.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of leaderboards to return in the response. For any response, the actual number of leaderboards returned may be less than the specified maxResults.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "100",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The token returned by the previous request.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "leaderboards",
	//   "response": {
	//     "$ref": "LeaderboardListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "games.players.get":

type PlayersGetCall struct {
	s        *Service
	playerId string
	opt_     map[string]interface{}
}

// Get: Retrieves the Player resource with the given ID. To retrieve the
// player for the currently authenticated user, set playerId to me.
func (r *PlayersService) Get(playerId string) *PlayersGetCall {
	c := &PlayersGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.playerId = playerId
	return c
}

func (c *PlayersGetCall) Do() (*Player, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/games/v1/", "players/{playerId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{playerId}", url.QueryEscape(c.playerId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Player)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the Player resource with the given ID. To retrieve the player for the currently authenticated user, set playerId to me.",
	//   "httpMethod": "GET",
	//   "id": "games.players.get",
	//   "parameterOrder": [
	//     "playerId"
	//   ],
	//   "parameters": {
	//     "playerId": {
	//       "description": "A player ID. A value of me may be used in place of the authenticated player's ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "players/{playerId}",
	//   "response": {
	//     "$ref": "Player"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "games.revisions.check":

type RevisionsCheckCall struct {
	s              *Service
	clientRevision string
	opt_           map[string]interface{}
}

// Check: Checks whether the games client is out of date.
func (r *RevisionsService) Check(clientRevision string) *RevisionsCheckCall {
	c := &RevisionsCheckCall{s: r.s, opt_: make(map[string]interface{})}
	c.clientRevision = clientRevision
	return c
}

func (c *RevisionsCheckCall) Do() (*RevisionCheckResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("clientRevision", fmt.Sprintf("%v", c.clientRevision))
	urls := googleapi.ResolveRelative("https://www.googleapis.com/games/v1/", "revisions/check")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(RevisionCheckResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Checks whether the games client is out of date.",
	//   "httpMethod": "GET",
	//   "id": "games.revisions.check",
	//   "parameterOrder": [
	//     "clientRevision"
	//   ],
	//   "parameters": {
	//     "clientRevision": {
	//       "description": "The revision of the client SDK used by your application. Format:\n[PLATFORM_TYPE]:[VERSION_NUMBER]. Possible values of PLATFORM_TYPE are:\n \n- \"ANDROID\" - Client is running the Android SDK. \n- \"IOS\" - Client is running the iOS SDK. \n- \"WEB_APP\" - Client is running as a Web App.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "revisions/check",
	//   "response": {
	//     "$ref": "RevisionCheckResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "games.rooms.create":

type RoomsCreateCall struct {
	s                 *Service
	roomcreaterequest *RoomCreateRequest
	opt_              map[string]interface{}
}

// Create: Create a room. For internal use by the Games SDK only.
// Calling this method directly is unsupported.
func (r *RoomsService) Create(roomcreaterequest *RoomCreateRequest) *RoomsCreateCall {
	c := &RoomsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.roomcreaterequest = roomcreaterequest
	return c
}

// Language sets the optional parameter "language": The preferred
// language to use for strings returned by this method.
func (c *RoomsCreateCall) Language(language string) *RoomsCreateCall {
	c.opt_["language"] = language
	return c
}

func (c *RoomsCreateCall) Do() (*Room, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.roomcreaterequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["language"]; ok {
		params.Set("language", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/games/v1/", "rooms/create")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Room)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Create a room. For internal use by the Games SDK only. Calling this method directly is unsupported.",
	//   "httpMethod": "POST",
	//   "id": "games.rooms.create",
	//   "parameters": {
	//     "language": {
	//       "description": "The preferred language to use for strings returned by this method.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "rooms/create",
	//   "request": {
	//     "$ref": "RoomCreateRequest"
	//   },
	//   "response": {
	//     "$ref": "Room"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "games.rooms.decline":

type RoomsDeclineCall struct {
	s      *Service
	roomId string
	opt_   map[string]interface{}
}

// Decline: Decline an invitation to join a room. For internal use by
// the Games SDK only. Calling this method directly is unsupported.
func (r *RoomsService) Decline(roomId string) *RoomsDeclineCall {
	c := &RoomsDeclineCall{s: r.s, opt_: make(map[string]interface{})}
	c.roomId = roomId
	return c
}

func (c *RoomsDeclineCall) Do() (*Room, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/games/v1/", "rooms/{roomId}/decline")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{roomId}", url.QueryEscape(c.roomId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Room)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Decline an invitation to join a room. For internal use by the Games SDK only. Calling this method directly is unsupported.",
	//   "httpMethod": "POST",
	//   "id": "games.rooms.decline",
	//   "parameterOrder": [
	//     "roomId"
	//   ],
	//   "parameters": {
	//     "roomId": {
	//       "description": "The ID of the room.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "rooms/{roomId}/decline",
	//   "response": {
	//     "$ref": "Room"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "games.rooms.dismiss":

type RoomsDismissCall struct {
	s      *Service
	roomId string
	opt_   map[string]interface{}
}

// Dismiss: Dismiss an invitation to join a room. For internal use by
// the Games SDK only. Calling this method directly is unsupported.
func (r *RoomsService) Dismiss(roomId string) *RoomsDismissCall {
	c := &RoomsDismissCall{s: r.s, opt_: make(map[string]interface{})}
	c.roomId = roomId
	return c
}

func (c *RoomsDismissCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/games/v1/", "rooms/{roomId}/dismiss")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{roomId}", url.QueryEscape(c.roomId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Dismiss an invitation to join a room. For internal use by the Games SDK only. Calling this method directly is unsupported.",
	//   "httpMethod": "POST",
	//   "id": "games.rooms.dismiss",
	//   "parameterOrder": [
	//     "roomId"
	//   ],
	//   "parameters": {
	//     "roomId": {
	//       "description": "The ID of the room.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "rooms/{roomId}/dismiss",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "games.rooms.get":

type RoomsGetCall struct {
	s      *Service
	roomId string
	opt_   map[string]interface{}
}

// Get: Get the data for a room.
func (r *RoomsService) Get(roomId string) *RoomsGetCall {
	c := &RoomsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.roomId = roomId
	return c
}

// Language sets the optional parameter "language": Specify the
// preferred language to use to format room info.
func (c *RoomsGetCall) Language(language string) *RoomsGetCall {
	c.opt_["language"] = language
	return c
}

func (c *RoomsGetCall) Do() (*Room, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["language"]; ok {
		params.Set("language", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/games/v1/", "rooms/{roomId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{roomId}", url.QueryEscape(c.roomId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Room)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get the data for a room.",
	//   "httpMethod": "GET",
	//   "id": "games.rooms.get",
	//   "parameterOrder": [
	//     "roomId"
	//   ],
	//   "parameters": {
	//     "language": {
	//       "description": "Specify the preferred language to use to format room info.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "roomId": {
	//       "description": "The ID of the room.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "rooms/{roomId}",
	//   "response": {
	//     "$ref": "Room"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "games.rooms.join":

type RoomsJoinCall struct {
	s               *Service
	roomId          string
	roomjoinrequest *RoomJoinRequest
	opt_            map[string]interface{}
}

// Join: Join a room. For internal use by the Games SDK only. Calling
// this method directly is unsupported.
func (r *RoomsService) Join(roomId string, roomjoinrequest *RoomJoinRequest) *RoomsJoinCall {
	c := &RoomsJoinCall{s: r.s, opt_: make(map[string]interface{})}
	c.roomId = roomId
	c.roomjoinrequest = roomjoinrequest
	return c
}

func (c *RoomsJoinCall) Do() (*Room, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.roomjoinrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/games/v1/", "rooms/{roomId}/join")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{roomId}", url.QueryEscape(c.roomId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Room)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Join a room. For internal use by the Games SDK only. Calling this method directly is unsupported.",
	//   "httpMethod": "POST",
	//   "id": "games.rooms.join",
	//   "parameterOrder": [
	//     "roomId"
	//   ],
	//   "parameters": {
	//     "roomId": {
	//       "description": "The ID of the room.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "rooms/{roomId}/join",
	//   "request": {
	//     "$ref": "RoomJoinRequest"
	//   },
	//   "response": {
	//     "$ref": "Room"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "games.rooms.leave":

type RoomsLeaveCall struct {
	s                *Service
	roomId           string
	roomleaverequest *RoomLeaveRequest
	opt_             map[string]interface{}
}

// Leave: Leave a room. For internal use by the Games SDK only. Calling
// this method directly is unsupported.
func (r *RoomsService) Leave(roomId string, roomleaverequest *RoomLeaveRequest) *RoomsLeaveCall {
	c := &RoomsLeaveCall{s: r.s, opt_: make(map[string]interface{})}
	c.roomId = roomId
	c.roomleaverequest = roomleaverequest
	return c
}

func (c *RoomsLeaveCall) Do() (*Room, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.roomleaverequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/games/v1/", "rooms/{roomId}/leave")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{roomId}", url.QueryEscape(c.roomId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Room)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Leave a room. For internal use by the Games SDK only. Calling this method directly is unsupported.",
	//   "httpMethod": "POST",
	//   "id": "games.rooms.leave",
	//   "parameterOrder": [
	//     "roomId"
	//   ],
	//   "parameters": {
	//     "roomId": {
	//       "description": "The ID of the room.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "rooms/{roomId}/leave",
	//   "request": {
	//     "$ref": "RoomLeaveRequest"
	//   },
	//   "response": {
	//     "$ref": "Room"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "games.rooms.list":

type RoomsListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Returns invitations to join rooms.
func (r *RoomsService) List() *RoomsListCall {
	c := &RoomsListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// Language sets the optional parameter "language": The preferred
// language to use for strings returned by this method.
func (c *RoomsListCall) Language(language string) *RoomsListCall {
	c.opt_["language"] = language
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of rooms to return in the response, used for paging. For any
// response, the actual number of rooms to return may be less than the
// specified maxResults.
func (c *RoomsListCall) MaxResults(maxResults int64) *RoomsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": The token returned
// by the previous request.
func (c *RoomsListCall) PageToken(pageToken string) *RoomsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

func (c *RoomsListCall) Do() (*RoomList, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["language"]; ok {
		params.Set("language", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/games/v1/", "rooms")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(RoomList)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns invitations to join rooms.",
	//   "httpMethod": "GET",
	//   "id": "games.rooms.list",
	//   "parameters": {
	//     "language": {
	//       "description": "The preferred language to use for strings returned by this method.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of rooms to return in the response, used for paging. For any response, the actual number of rooms to return may be less than the specified maxResults.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The token returned by the previous request.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "rooms",
	//   "response": {
	//     "$ref": "RoomList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "games.rooms.reportStatus":

type RoomsReportStatusCall struct {
	s               *Service
	roomId          string
	roomp2pstatuses *RoomP2PStatuses
	opt_            map[string]interface{}
}

// ReportStatus: Updates sent by a client reporting the status of peers
// in a room. For internal use by the Games SDK only. Calling this
// method directly is unsupported.
func (r *RoomsService) ReportStatus(roomId string, roomp2pstatuses *RoomP2PStatuses) *RoomsReportStatusCall {
	c := &RoomsReportStatusCall{s: r.s, opt_: make(map[string]interface{})}
	c.roomId = roomId
	c.roomp2pstatuses = roomp2pstatuses
	return c
}

func (c *RoomsReportStatusCall) Do() (*RoomStatus, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.roomp2pstatuses)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/games/v1/", "rooms/{roomId}/reportstatus")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{roomId}", url.QueryEscape(c.roomId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(RoomStatus)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates sent by a client reporting the status of peers in a room. For internal use by the Games SDK only. Calling this method directly is unsupported.",
	//   "httpMethod": "POST",
	//   "id": "games.rooms.reportStatus",
	//   "parameterOrder": [
	//     "roomId"
	//   ],
	//   "parameters": {
	//     "roomId": {
	//       "description": "The ID of the room.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "rooms/{roomId}/reportstatus",
	//   "request": {
	//     "$ref": "RoomP2PStatuses"
	//   },
	//   "response": {
	//     "$ref": "RoomStatus"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "games.scores.get":

type ScoresGetCall struct {
	s             *Service
	playerId      string
	leaderboardId string
	timeSpan      string
	opt_          map[string]interface{}
}

// Get: Get high scores, and optionally ranks, in leaderboards for the
// currently authenticated player. For a specific time span,
// leaderboardId can be set to ALL to retrieve data for all leaderboards
// in a given time span.
// NOTE: You cannot ask for 'ALL' leaderboards and
// 'ALL' timeSpans in the same request; only one parameter may be set to
// 'ALL'.
func (r *ScoresService) Get(playerId string, leaderboardId string, timeSpan string) *ScoresGetCall {
	c := &ScoresGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.playerId = playerId
	c.leaderboardId = leaderboardId
	c.timeSpan = timeSpan
	return c
}

// IncludeRankType sets the optional parameter "includeRankType": The
// types of ranks to return. If the parameter is omitted, no ranks will
// be returned.
func (c *ScoresGetCall) IncludeRankType(includeRankType string) *ScoresGetCall {
	c.opt_["includeRankType"] = includeRankType
	return c
}

// Language sets the optional parameter "language": The preferred
// language to use for strings returned by this method.
func (c *ScoresGetCall) Language(language string) *ScoresGetCall {
	c.opt_["language"] = language
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of leaderboard scores to return in the response. For any
// response, the actual number of leaderboard scores returned may be
// less than the specified maxResults.
func (c *ScoresGetCall) MaxResults(maxResults int64) *ScoresGetCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": The token returned
// by the previous request.
func (c *ScoresGetCall) PageToken(pageToken string) *ScoresGetCall {
	c.opt_["pageToken"] = pageToken
	return c
}

func (c *ScoresGetCall) Do() (*PlayerLeaderboardScoreListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["includeRankType"]; ok {
		params.Set("includeRankType", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["language"]; ok {
		params.Set("language", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/games/v1/", "players/{playerId}/leaderboards/{leaderboardId}/scores/{timeSpan}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{playerId}", url.QueryEscape(c.playerId), 1)
	req.URL.Path = strings.Replace(req.URL.Path, "{leaderboardId}", url.QueryEscape(c.leaderboardId), 1)
	req.URL.Path = strings.Replace(req.URL.Path, "{timeSpan}", url.QueryEscape(c.timeSpan), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(PlayerLeaderboardScoreListResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get high scores, and optionally ranks, in leaderboards for the currently authenticated player. For a specific time span, leaderboardId can be set to ALL to retrieve data for all leaderboards in a given time span.\nNOTE: You cannot ask for 'ALL' leaderboards and 'ALL' timeSpans in the same request; only one parameter may be set to 'ALL'.",
	//   "httpMethod": "GET",
	//   "id": "games.scores.get",
	//   "parameterOrder": [
	//     "playerId",
	//     "leaderboardId",
	//     "timeSpan"
	//   ],
	//   "parameters": {
	//     "includeRankType": {
	//       "description": "The types of ranks to return. If the parameter is omitted, no ranks will be returned.",
	//       "enum": [
	//         "ALL",
	//         "PUBLIC",
	//         "SOCIAL"
	//       ],
	//       "enumDescriptions": [
	//         "Retrieve public and social ranks.",
	//         "Retrieve public ranks, if the player is sharing their gameplay activity publicly.",
	//         "Retrieve the social rank."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "language": {
	//       "description": "The preferred language to use for strings returned by this method.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "leaderboardId": {
	//       "description": "The ID of the leaderboard. Can be set to 'ALL' to retrieve data for all leaderboards for this application.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of leaderboard scores to return in the response. For any response, the actual number of leaderboard scores returned may be less than the specified maxResults.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "25",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The token returned by the previous request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "playerId": {
	//       "description": "A player ID. A value of me may be used in place of the authenticated player's ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "timeSpan": {
	//       "description": "The time span for the scores and ranks you're requesting.",
	//       "enum": [
	//         "ALL",
	//         "ALL_TIME",
	//         "DAILY",
	//         "WEEKLY"
	//       ],
	//       "enumDescriptions": [
	//         "Get the high scores for all time spans. If this is used, maxResults values will be ignored.",
	//         "Get the all time high score.",
	//         "List the top scores for the current day.",
	//         "List the top scores for the current week."
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "players/{playerId}/leaderboards/{leaderboardId}/scores/{timeSpan}",
	//   "response": {
	//     "$ref": "PlayerLeaderboardScoreListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "games.scores.list":

type ScoresListCall struct {
	s             *Service
	leaderboardId string
	collection    string
	timeSpan      string
	opt_          map[string]interface{}
}

// List: Lists the scores in a leaderboard, starting from the top.
func (r *ScoresService) List(leaderboardId string, collection string, timeSpan string) *ScoresListCall {
	c := &ScoresListCall{s: r.s, opt_: make(map[string]interface{})}
	c.leaderboardId = leaderboardId
	c.collection = collection
	c.timeSpan = timeSpan
	return c
}

// Language sets the optional parameter "language": The preferred
// language to use for strings returned by this method.
func (c *ScoresListCall) Language(language string) *ScoresListCall {
	c.opt_["language"] = language
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of leaderboard scores to return in the response. For any
// response, the actual number of leaderboard scores returned may be
// less than the specified maxResults.
func (c *ScoresListCall) MaxResults(maxResults int64) *ScoresListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": The token returned
// by the previous request.
func (c *ScoresListCall) PageToken(pageToken string) *ScoresListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

func (c *ScoresListCall) Do() (*LeaderboardScores, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("timeSpan", fmt.Sprintf("%v", c.timeSpan))
	if v, ok := c.opt_["language"]; ok {
		params.Set("language", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/games/v1/", "leaderboards/{leaderboardId}/scores/{collection}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{leaderboardId}", url.QueryEscape(c.leaderboardId), 1)
	req.URL.Path = strings.Replace(req.URL.Path, "{collection}", url.QueryEscape(c.collection), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(LeaderboardScores)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the scores in a leaderboard, starting from the top.",
	//   "httpMethod": "GET",
	//   "id": "games.scores.list",
	//   "parameterOrder": [
	//     "leaderboardId",
	//     "collection",
	//     "timeSpan"
	//   ],
	//   "parameters": {
	//     "collection": {
	//       "description": "The collection of scores you're requesting.",
	//       "enum": [
	//         "PUBLIC",
	//         "SOCIAL"
	//       ],
	//       "enumDescriptions": [
	//         "List all scores in the public leaderboard.",
	//         "List only social scores."
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "language": {
	//       "description": "The preferred language to use for strings returned by this method.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "leaderboardId": {
	//       "description": "The ID of the leaderboard.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of leaderboard scores to return in the response. For any response, the actual number of leaderboard scores returned may be less than the specified maxResults.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "25",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The token returned by the previous request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "timeSpan": {
	//       "description": "The time span for the scores and ranks you're requesting.",
	//       "enum": [
	//         "ALL_TIME",
	//         "DAILY",
	//         "WEEKLY"
	//       ],
	//       "enumDescriptions": [
	//         "List the all-time top scores.",
	//         "List the top scores for the current day.",
	//         "List the top scores for the current week."
	//       ],
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "leaderboards/{leaderboardId}/scores/{collection}",
	//   "response": {
	//     "$ref": "LeaderboardScores"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "games.scores.listWindow":

type ScoresListWindowCall struct {
	s             *Service
	leaderboardId string
	collection    string
	timeSpan      string
	opt_          map[string]interface{}
}

// ListWindow: Lists the scores in a leaderboard around (and including)
// a player's score.
func (r *ScoresService) ListWindow(leaderboardId string, collection string, timeSpan string) *ScoresListWindowCall {
	c := &ScoresListWindowCall{s: r.s, opt_: make(map[string]interface{})}
	c.leaderboardId = leaderboardId
	c.collection = collection
	c.timeSpan = timeSpan
	return c
}

// Language sets the optional parameter "language": The preferred
// language to use for strings returned by this method.
func (c *ScoresListWindowCall) Language(language string) *ScoresListWindowCall {
	c.opt_["language"] = language
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of leaderboard scores to return in the response. For any
// response, the actual number of leaderboard scores returned may be
// less than the specified maxResults.
func (c *ScoresListWindowCall) MaxResults(maxResults int64) *ScoresListWindowCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": The token returned
// by the previous request.
func (c *ScoresListWindowCall) PageToken(pageToken string) *ScoresListWindowCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// ResultsAbove sets the optional parameter "resultsAbove": The
// preferred number of scores to return above the player's score. More
// scores may be returned if the player is at the bottom of the
// leaderboard; fewer may be returned if the player is at the top. Must
// be less than or equal to maxResults.
func (c *ScoresListWindowCall) ResultsAbove(resultsAbove int64) *ScoresListWindowCall {
	c.opt_["resultsAbove"] = resultsAbove
	return c
}

// ReturnTopIfAbsent sets the optional parameter "returnTopIfAbsent":
// True if the top scores should be returned when the player is not in
// the leaderboard. Defaults to true.
func (c *ScoresListWindowCall) ReturnTopIfAbsent(returnTopIfAbsent bool) *ScoresListWindowCall {
	c.opt_["returnTopIfAbsent"] = returnTopIfAbsent
	return c
}

func (c *ScoresListWindowCall) Do() (*LeaderboardScores, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("timeSpan", fmt.Sprintf("%v", c.timeSpan))
	if v, ok := c.opt_["language"]; ok {
		params.Set("language", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["resultsAbove"]; ok {
		params.Set("resultsAbove", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["returnTopIfAbsent"]; ok {
		params.Set("returnTopIfAbsent", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/games/v1/", "leaderboards/{leaderboardId}/window/{collection}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{leaderboardId}", url.QueryEscape(c.leaderboardId), 1)
	req.URL.Path = strings.Replace(req.URL.Path, "{collection}", url.QueryEscape(c.collection), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(LeaderboardScores)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the scores in a leaderboard around (and including) a player's score.",
	//   "httpMethod": "GET",
	//   "id": "games.scores.listWindow",
	//   "parameterOrder": [
	//     "leaderboardId",
	//     "collection",
	//     "timeSpan"
	//   ],
	//   "parameters": {
	//     "collection": {
	//       "description": "The collection of scores you're requesting.",
	//       "enum": [
	//         "PUBLIC",
	//         "SOCIAL"
	//       ],
	//       "enumDescriptions": [
	//         "List all scores in the public leaderboard.",
	//         "List only social scores."
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "language": {
	//       "description": "The preferred language to use for strings returned by this method.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "leaderboardId": {
	//       "description": "The ID of the leaderboard.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of leaderboard scores to return in the response. For any response, the actual number of leaderboard scores returned may be less than the specified maxResults.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "25",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The token returned by the previous request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "resultsAbove": {
	//       "description": "The preferred number of scores to return above the player's score. More scores may be returned if the player is at the bottom of the leaderboard; fewer may be returned if the player is at the top. Must be less than or equal to maxResults.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "returnTopIfAbsent": {
	//       "description": "True if the top scores should be returned when the player is not in the leaderboard. Defaults to true.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "timeSpan": {
	//       "description": "The time span for the scores and ranks you're requesting.",
	//       "enum": [
	//         "ALL_TIME",
	//         "DAILY",
	//         "WEEKLY"
	//       ],
	//       "enumDescriptions": [
	//         "List the all-time top scores.",
	//         "List the top scores for the current day.",
	//         "List the top scores for the current week."
	//       ],
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "leaderboards/{leaderboardId}/window/{collection}",
	//   "response": {
	//     "$ref": "LeaderboardScores"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "games.scores.submit":

type ScoresSubmitCall struct {
	s             *Service
	leaderboardId string
	score         int64
	opt_          map[string]interface{}
}

// Submit: Submits a score to the specified leaderboard.
func (r *ScoresService) Submit(leaderboardId string, score int64) *ScoresSubmitCall {
	c := &ScoresSubmitCall{s: r.s, opt_: make(map[string]interface{})}
	c.leaderboardId = leaderboardId
	c.score = score
	return c
}

// Language sets the optional parameter "language": The preferred
// language to use for strings returned by this method.
func (c *ScoresSubmitCall) Language(language string) *ScoresSubmitCall {
	c.opt_["language"] = language
	return c
}

// ScoreTag sets the optional parameter "scoreTag": Additional
// information about the score you're submitting. Values must contain no
// more than 64 URI-safe characters as defined by section 2.3 of RFC
// 3986.
func (c *ScoresSubmitCall) ScoreTag(scoreTag string) *ScoresSubmitCall {
	c.opt_["scoreTag"] = scoreTag
	return c
}

func (c *ScoresSubmitCall) Do() (*PlayerScoreResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("score", fmt.Sprintf("%v", c.score))
	if v, ok := c.opt_["language"]; ok {
		params.Set("language", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["scoreTag"]; ok {
		params.Set("scoreTag", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/games/v1/", "leaderboards/{leaderboardId}/scores")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{leaderboardId}", url.QueryEscape(c.leaderboardId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(PlayerScoreResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Submits a score to the specified leaderboard.",
	//   "httpMethod": "POST",
	//   "id": "games.scores.submit",
	//   "parameterOrder": [
	//     "leaderboardId",
	//     "score"
	//   ],
	//   "parameters": {
	//     "language": {
	//       "description": "The preferred language to use for strings returned by this method.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "leaderboardId": {
	//       "description": "The ID of the leaderboard.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "score": {
	//       "description": "The score you're submitting. The submitted score is ignored if it is worse than a previously submitted score, where worse depends on the leaderboard sort order. The meaning of the score value depends on the leaderboard format type. For fixed-point, the score represents the raw value. For time, the score represents elapsed time in milliseconds. For currency, the score represents a value in micro units.",
	//       "format": "int64",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "scoreTag": {
	//       "description": "Additional information about the score you're submitting. Values must contain no more than 64 URI-safe characters as defined by section 2.3 of RFC 3986.",
	//       "location": "query",
	//       "pattern": "[a-zA-Z0-9-._~]{0,64}",
	//       "type": "string"
	//     }
	//   },
	//   "path": "leaderboards/{leaderboardId}/scores",
	//   "response": {
	//     "$ref": "PlayerScoreResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "games.scores.submitMultiple":

type ScoresSubmitMultipleCall struct {
	s                         *Service
	playerscoresubmissionlist *PlayerScoreSubmissionList
	opt_                      map[string]interface{}
}

// SubmitMultiple: Submits multiple scores to leaderboards.
func (r *ScoresService) SubmitMultiple(playerscoresubmissionlist *PlayerScoreSubmissionList) *ScoresSubmitMultipleCall {
	c := &ScoresSubmitMultipleCall{s: r.s, opt_: make(map[string]interface{})}
	c.playerscoresubmissionlist = playerscoresubmissionlist
	return c
}

// Language sets the optional parameter "language": The preferred
// language to use for strings returned by this method.
func (c *ScoresSubmitMultipleCall) Language(language string) *ScoresSubmitMultipleCall {
	c.opt_["language"] = language
	return c
}

func (c *ScoresSubmitMultipleCall) Do() (*PlayerScoreListResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.playerscoresubmissionlist)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["language"]; ok {
		params.Set("language", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/games/v1/", "leaderboards/scores")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(PlayerScoreListResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Submits multiple scores to leaderboards.",
	//   "httpMethod": "POST",
	//   "id": "games.scores.submitMultiple",
	//   "parameters": {
	//     "language": {
	//       "description": "The preferred language to use for strings returned by this method.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "leaderboards/scores",
	//   "request": {
	//     "$ref": "PlayerScoreSubmissionList"
	//   },
	//   "response": {
	//     "$ref": "PlayerScoreListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}
