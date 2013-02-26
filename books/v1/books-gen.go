// Package books provides access to the Books API.
//
// See https://developers.google.com/books/docs/v1/getting_started
//
// Usage example:
//
//   import "code.google.com/p/google-api-go-client/books/v1"
//   ...
//   booksService, err := books.New(oauthHttpClient)
package books

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

const apiId = "books:v1"
const apiName = "books"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/books/v1/"

// OAuth2 scopes used by this API.
const (
	// Manage your books
	BooksScope = "https://www.googleapis.com/auth/books"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Bookshelves = &BookshelvesService{s: s}
	s.Layers = &LayersService{s: s}
	s.Myconfig = &MyconfigService{s: s}
	s.Mylibrary = &MylibraryService{s: s}
	s.Volumes = &VolumesService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	Bookshelves *BookshelvesService

	Layers *LayersService

	Myconfig *MyconfigService

	Mylibrary *MylibraryService

	Volumes *VolumesService
}

type BookshelvesService struct {
	s *Service
}

type LayersService struct {
	s *Service
}

type MyconfigService struct {
	s *Service
}

type MylibraryService struct {
	s *Service
}

type VolumesService struct {
	s *Service
}

type Annotation struct {
	// AfterSelectedText: Anchor text after excerpt.
	AfterSelectedText string `json:"afterSelectedText,omitempty"`

	// BeforeSelectedText: Anchor text before excerpt.
	BeforeSelectedText string `json:"beforeSelectedText,omitempty"`

	// ClientVersionRanges: Selection ranges sent from the client.
	ClientVersionRanges *AnnotationClientVersionRanges `json:"clientVersionRanges,omitempty"`

	// Created: Timestamp for the created time of this annotation.
	Created string `json:"created,omitempty"`

	// CurrentVersionRanges: Selection ranges for the most recent content
	// version.
	CurrentVersionRanges *AnnotationCurrentVersionRanges `json:"currentVersionRanges,omitempty"`

	// Data: User-created data for this annotation.
	Data string `json:"data,omitempty"`

	// Deleted: Indicates that this annotation is deleted.
	Deleted bool `json:"deleted,omitempty"`

	// HighlightStyle: The highlight style for this annotation.
	HighlightStyle string `json:"highlightStyle,omitempty"`

	// Id: Id of this annotation, in the form of a GUID.
	Id string `json:"id,omitempty"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`

	// LayerId: The layer this annotation is for.
	LayerId string `json:"layerId,omitempty"`

	// PageIds: Pages that this annotation spans.
	PageIds []string `json:"pageIds,omitempty"`

	// SelectedText: Excerpt from the volume.
	SelectedText string `json:"selectedText,omitempty"`

	// SelfLink: URL to this resource.
	SelfLink string `json:"selfLink,omitempty"`

	// Updated: Timestamp for the last time this annotation was modified.
	Updated string `json:"updated,omitempty"`

	// VolumeId: The volume that this annotation belongs to.
	VolumeId string `json:"volumeId,omitempty"`
}

type AnnotationClientVersionRanges struct {
	// CfiRange: Range in CFI format for this annotation sent by client.
	CfiRange *BooksAnnotationsRange `json:"cfiRange,omitempty"`

	// ContentVersion: Content version the client sent in.
	ContentVersion string `json:"contentVersion,omitempty"`

	// GbImageRange: Range in GB image format for this annotation sent by
	// client.
	GbImageRange *BooksAnnotationsRange `json:"gbImageRange,omitempty"`

	// GbTextRange: Range in GB text format for this annotation sent by
	// client.
	GbTextRange *BooksAnnotationsRange `json:"gbTextRange,omitempty"`
}

type AnnotationCurrentVersionRanges struct {
	// CfiRange: Range in CFI format for this annotation for version above.
	CfiRange *BooksAnnotationsRange `json:"cfiRange,omitempty"`

	// ContentVersion: Content version applicable to ranges below.
	ContentVersion string `json:"contentVersion,omitempty"`

	// GbImageRange: Range in GB image format for this annotation for
	// version above.
	GbImageRange *BooksAnnotationsRange `json:"gbImageRange,omitempty"`

	// GbTextRange: Range in GB text format for this annotation for version
	// above.
	GbTextRange *BooksAnnotationsRange `json:"gbTextRange,omitempty"`
}

type Annotationdata struct {
	// AnnotationType: The type of annotation this data is for.
	AnnotationType string `json:"annotationType,omitempty"`

	Data interface{} `json:"data,omitempty"`

	// Encoded_data: Base64 encoded data for this annotation data.
	Encoded_data string `json:"encoded_data,omitempty"`

	// Id: Unique id for this annotation data.
	Id string `json:"id,omitempty"`

	// Kind: Resource Type
	Kind string `json:"kind,omitempty"`

	// LayerId: The Layer id for this data. *
	LayerId string `json:"layerId,omitempty"`

	// SelfLink: URL for this resource. *
	SelfLink string `json:"selfLink,omitempty"`

	// Updated: Timestamp for the last time this data was updated. (RFC 3339
	// UTC date-time format).
	Updated string `json:"updated,omitempty"`

	// VolumeId: The volume id for this data. *
	VolumeId string `json:"volumeId,omitempty"`
}

type Annotations struct {
	// Items: A list of annotations.
	Items []*Annotation `json:"items,omitempty"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token to pass in for pagination for the next page.
	// This will not be present if this request does not have more results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalItems: Total number of annotations found. This may be greater
	// than the number of notes returned in this response if results have
	// been paginated.
	TotalItems int64 `json:"totalItems,omitempty"`
}

type Annotationsdata struct {
	// Items: A list of Annotation Data.
	Items []*Annotationdata `json:"items,omitempty"`

	// Kind: Resource type
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token to pass in for pagination for the next page.
	// This will not be present if this request does not have more results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalItems: The total number of volume annotations found.
	TotalItems int64 `json:"totalItems,omitempty"`
}

type BooksAnnotationsRange struct {
	// EndOffset: The offset from the ending position.
	EndOffset string `json:"endOffset,omitempty"`

	// EndPosition: The ending position for the range.
	EndPosition string `json:"endPosition,omitempty"`

	// StartOffset: The offset from the starting position.
	StartOffset string `json:"startOffset,omitempty"`

	// StartPosition: The starting position for the range.
	StartPosition string `json:"startPosition,omitempty"`
}

type BooksLayerDictData struct {
	Common *BooksLayerDictDataCommon `json:"common,omitempty"`

	Dict *BooksLayerDictDataDict `json:"dict,omitempty"`
}

type BooksLayerDictDataCommon struct {
	// Title: The display title and localized canonical name to use when
	// searching for this entity on Google search.
	Title string `json:"title,omitempty"`
}

type BooksLayerDictDataDict struct {
	// Source: The source, url and attribution for this dictionary data.
	Source *BooksLayerDictDataDictSource `json:"source,omitempty"`

	Words []*BooksLayerDictDataDictWords `json:"words,omitempty"`
}

type BooksLayerDictDataDictSource struct {
	Attribution string `json:"attribution,omitempty"`

	Url string `json:"url,omitempty"`
}

type BooksLayerDictDataDictWords struct {
	Derivatives []*BooksLayerDictDataDictWordsDerivatives `json:"derivatives,omitempty"`

	Examples []*BooksLayerDictDataDictWordsExamples `json:"examples,omitempty"`

	Senses []*BooksLayerDictDataDictWordsSenses `json:"senses,omitempty"`

	// Source: The words with different meanings but not related words, e.g.
	// "go" (game) and "go" (verb).
	Source *BooksLayerDictDataDictWordsSource `json:"source,omitempty"`
}

type BooksLayerDictDataDictWordsDerivatives struct {
	Source *BooksLayerDictDataDictWordsDerivativesSource `json:"source,omitempty"`

	Text string `json:"text,omitempty"`
}

type BooksLayerDictDataDictWordsDerivativesSource struct {
	Attribution string `json:"attribution,omitempty"`

	Url string `json:"url,omitempty"`
}

type BooksLayerDictDataDictWordsExamples struct {
	Source *BooksLayerDictDataDictWordsExamplesSource `json:"source,omitempty"`

	Text string `json:"text,omitempty"`
}

type BooksLayerDictDataDictWordsExamplesSource struct {
	Attribution string `json:"attribution,omitempty"`

	Url string `json:"url,omitempty"`
}

type BooksLayerDictDataDictWordsSenses struct {
	Conjugations []*BooksLayerDictDataDictWordsSensesConjugations `json:"conjugations,omitempty"`

	Definitions []*BooksLayerDictDataDictWordsSensesDefinitions `json:"definitions,omitempty"`

	PartOfSpeech string `json:"partOfSpeech,omitempty"`

	Pronunciation string `json:"pronunciation,omitempty"`

	PronunciationUrl string `json:"pronunciationUrl,omitempty"`

	Source *BooksLayerDictDataDictWordsSensesSource `json:"source,omitempty"`

	Syllabification string `json:"syllabification,omitempty"`

	Synonyms []*BooksLayerDictDataDictWordsSensesSynonyms `json:"synonyms,omitempty"`
}

type BooksLayerDictDataDictWordsSensesConjugations struct {
	Type string `json:"type,omitempty"`

	Value string `json:"value,omitempty"`
}

type BooksLayerDictDataDictWordsSensesDefinitions struct {
	Definition string `json:"definition,omitempty"`

	Examples []*BooksLayerDictDataDictWordsSensesDefinitionsExamples `json:"examples,omitempty"`
}

type BooksLayerDictDataDictWordsSensesDefinitionsExamples struct {
	Source *BooksLayerDictDataDictWordsSensesDefinitionsExamplesSource `json:"source,omitempty"`

	Text string `json:"text,omitempty"`
}

type BooksLayerDictDataDictWordsSensesDefinitionsExamplesSource struct {
	Attribution string `json:"attribution,omitempty"`

	Url string `json:"url,omitempty"`
}

type BooksLayerDictDataDictWordsSensesSource struct {
	Attribution string `json:"attribution,omitempty"`

	Url string `json:"url,omitempty"`
}

type BooksLayerDictDataDictWordsSensesSynonyms struct {
	Source *BooksLayerDictDataDictWordsSensesSynonymsSource `json:"source,omitempty"`

	Text string `json:"text,omitempty"`
}

type BooksLayerDictDataDictWordsSensesSynonymsSource struct {
	Attribution string `json:"attribution,omitempty"`

	Url string `json:"url,omitempty"`
}

type BooksLayerDictDataDictWordsSource struct {
	Attribution string `json:"attribution,omitempty"`

	Url string `json:"url,omitempty"`
}

type BooksLayerGeoData struct {
	Common *BooksLayerGeoDataCommon `json:"common,omitempty"`

	Geo *BooksLayerGeoDataGeo `json:"geo,omitempty"`
}

type BooksLayerGeoDataCommon struct {
	// Lang: The language of the information url and description.
	Lang string `json:"lang,omitempty"`

	// PreviewImageUrl: The URL for the preview image information.
	PreviewImageUrl string `json:"previewImageUrl,omitempty"`

	// Snippet: The description for this location.
	Snippet string `json:"snippet,omitempty"`

	// SnippetUrl: The URL for information for this location. Ex: wikipedia
	// link.
	SnippetUrl string `json:"snippetUrl,omitempty"`

	// Title: The display title and localized canonical name to use when
	// searching for this entity on Google search.
	Title string `json:"title,omitempty"`
}

type BooksLayerGeoDataGeo struct {
	// Boundary: The boundary of the location as a set of loops containing
	// pairs of latitude, longitude coordinates.
	Boundary [][]*BooksLayerGeoDataGeoBoundaryItem `json:"boundary,omitempty"`

	// CachePolicy: The cache policy active for this data. EX: UNRESTRICTED,
	// RESTRICTED, NEVER
	CachePolicy string `json:"cachePolicy,omitempty"`

	// CountryCode: The country code of the location.
	CountryCode string `json:"countryCode,omitempty"`

	// Latitude: The latitude of the location.
	Latitude float64 `json:"latitude,omitempty"`

	// Longitude: The longitude of the location.
	Longitude float64 `json:"longitude,omitempty"`

	// MapType: The type of map that should be used for this location. EX:
	// HYBRID, ROADMAP, SATELLITE, TERRAIN
	MapType string `json:"mapType,omitempty"`

	// Viewport: The viewport for showing this location. This is a latitude,
	// longitude rectangle.
	Viewport *BooksLayerGeoDataGeoViewport `json:"viewport,omitempty"`

	// Zoom: The Zoom level to use for the map. Zoom levels between 0 (the
	// lowest zoom level, in which the entire world can be seen on one map)
	// to 21+ (down to individual buildings). See:
	// https://developers.google.com/maps/documentation/staticmaps/#Zoomlevel
	// s
	Zoom int64 `json:"zoom,omitempty"`
}

type BooksLayerGeoDataGeoBoundaryItem struct {
	Latitude int64 `json:"latitude,omitempty"`

	Longitude int64 `json:"longitude,omitempty"`
}

type BooksLayerGeoDataGeoViewport struct {
	Hi *BooksLayerGeoDataGeoViewportHi `json:"hi,omitempty"`

	Lo *BooksLayerGeoDataGeoViewportLo `json:"lo,omitempty"`
}

type BooksLayerGeoDataGeoViewportHi struct {
	Latitude float64 `json:"latitude,omitempty"`

	Longitude float64 `json:"longitude,omitempty"`
}

type BooksLayerGeoDataGeoViewportLo struct {
	Latitude float64 `json:"latitude,omitempty"`

	Longitude float64 `json:"longitude,omitempty"`
}

type Bookshelf struct {
	// Access: Whether this bookshelf is PUBLIC or PRIVATE.
	Access string `json:"access,omitempty"`

	// Created: Created time for this bookshelf (formatted UTC timestamp
	// with millisecond resolution).
	Created string `json:"created,omitempty"`

	// Description: Description of this bookshelf.
	Description string `json:"description,omitempty"`

	// Id: Id of this bookshelf, only unique by user.
	Id int64 `json:"id,omitempty"`

	// Kind: Resource type for bookshelf metadata.
	Kind string `json:"kind,omitempty"`

	// SelfLink: URL to this resource.
	SelfLink string `json:"selfLink,omitempty"`

	// Title: Title of this bookshelf.
	Title string `json:"title,omitempty"`

	// Updated: Last modified time of this bookshelf (formatted UTC
	// timestamp with millisecond resolution).
	Updated string `json:"updated,omitempty"`

	// VolumeCount: Number of volumes in this bookshelf.
	VolumeCount int64 `json:"volumeCount,omitempty"`

	// VolumesLastUpdated: Last time a volume was added or removed from this
	// bookshelf (formatted UTC timestamp with millisecond resolution).
	VolumesLastUpdated string `json:"volumesLastUpdated,omitempty"`
}

type Bookshelves struct {
	// Items: A list of bookshelves.
	Items []*Bookshelf `json:"items,omitempty"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`
}

type ConcurrentAccessRestriction struct {
	// DeviceAllowed: Whether access is granted for this (user, device,
	// volume).
	DeviceAllowed bool `json:"deviceAllowed,omitempty"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`

	// MaxConcurrentDevices: The maximum number of concurrent access
	// licenses for this volume.
	MaxConcurrentDevices int64 `json:"maxConcurrentDevices,omitempty"`

	// Message: Error/warning message.
	Message string `json:"message,omitempty"`

	// Nonce: Client nonce for verification. Download access and
	// client-validation only.
	Nonce string `json:"nonce,omitempty"`

	// ReasonCode: Error/warning reason code.
	ReasonCode string `json:"reasonCode,omitempty"`

	// Restricted: Whether this volume has any concurrent access
	// restrictions.
	Restricted bool `json:"restricted,omitempty"`

	// Signature: Response signature.
	Signature string `json:"signature,omitempty"`

	// Source: Client app identifier for verification. Download access and
	// client-validation only.
	Source string `json:"source,omitempty"`

	// TimeWindowSeconds: Time in seconds for license auto-expiration.
	TimeWindowSeconds int64 `json:"timeWindowSeconds,omitempty"`

	// VolumeId: Identifies the volume for which this entry applies.
	VolumeId string `json:"volumeId,omitempty"`
}

type DownloadAccessRestriction struct {
	// DeviceAllowed: If restricted, whether access is granted for this
	// (user, device, volume).
	DeviceAllowed bool `json:"deviceAllowed,omitempty"`

	// DownloadsAcquired: If restricted, the number of content download
	// licenses already acquired (including the requesting client, if
	// licensed).
	DownloadsAcquired int64 `json:"downloadsAcquired,omitempty"`

	// JustAcquired: If deviceAllowed, whether access was just acquired with
	// this request.
	JustAcquired bool `json:"justAcquired,omitempty"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`

	// MaxDownloadDevices: If restricted, the maximum number of content
	// download licenses for this volume.
	MaxDownloadDevices int64 `json:"maxDownloadDevices,omitempty"`

	// Message: Error/warning message.
	Message string `json:"message,omitempty"`

	// Nonce: Client nonce for verification. Download access and
	// client-validation only.
	Nonce string `json:"nonce,omitempty"`

	// ReasonCode: Error/warning reason code. Additional codes may be added
	// in the future. 0 OK 100 ACCESS_DENIED_PUBLISHER_LIMIT 101
	// ACCESS_DENIED_LIMIT 200 WARNING_USED_LAST_ACCESS
	ReasonCode string `json:"reasonCode,omitempty"`

	// Restricted: Whether this volume has any download access restrictions.
	Restricted bool `json:"restricted,omitempty"`

	// Signature: Response signature.
	Signature string `json:"signature,omitempty"`

	// Source: Client app identifier for verification. Download access and
	// client-validation only.
	Source string `json:"source,omitempty"`

	// VolumeId: Identifies the volume for which this entry applies.
	VolumeId string `json:"volumeId,omitempty"`
}

type DownloadAccesses struct {
	// DownloadAccessList: A list of download access responses.
	DownloadAccessList []*DownloadAccessRestriction `json:"downloadAccessList,omitempty"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`
}

type Layersummaries struct {
	// Items: A list of layer summary items.
	Items []*Layersummary `json:"items,omitempty"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`

	// TotalItems: The total number of layer summaries found.
	TotalItems int64 `json:"totalItems,omitempty"`
}

type Layersummary struct {
	// AnnotationCount: The number of annotations for this layer.
	AnnotationCount int64 `json:"annotationCount,omitempty"`

	// AnnotationTypes: The list of annotation types contained for this
	// layer.
	AnnotationTypes []string `json:"annotationTypes,omitempty"`

	// AnnotationsDataLink: Link to get data for this annotation.
	AnnotationsDataLink string `json:"annotationsDataLink,omitempty"`

	// AnnotationsLink: The link to get the annotations for this layer.
	AnnotationsLink string `json:"annotationsLink,omitempty"`

	// ContentVersion: The content version this resource is for.
	ContentVersion string `json:"contentVersion,omitempty"`

	// DataCount: The number of data items for this layer.
	DataCount int64 `json:"dataCount,omitempty"`

	// Id: Unique id of this layer summary.
	Id string `json:"id,omitempty"`

	// Kind: Resource Type
	Kind string `json:"kind,omitempty"`

	// LayerId: The layer id for this summary.
	LayerId string `json:"layerId,omitempty"`

	// SelfLink: URL to this resource.
	SelfLink string `json:"selfLink,omitempty"`

	// Updated: Timestamp for the last time an item in this layer was
	// updated. (RFC 3339 UTC date-time format).
	Updated string `json:"updated,omitempty"`

	// VolumeAnnotationsVersion: The current version of this layer's volume
	// annotations. Note that this version applies only to the data in the
	// books.layers.volumeAnnotations.* responses. The actual annotation
	// data is versioned separately.
	VolumeAnnotationsVersion string `json:"volumeAnnotationsVersion,omitempty"`

	// VolumeId: The volume id this resource is for.
	VolumeId string `json:"volumeId,omitempty"`
}

type ReadingPosition struct {
	// EpubCfiPosition: Position in an EPUB as a CFI.
	EpubCfiPosition string `json:"epubCfiPosition,omitempty"`

	// GbImagePosition: Position in a volume for image-based content.
	GbImagePosition string `json:"gbImagePosition,omitempty"`

	// GbTextPosition: Position in a volume for text-based content.
	GbTextPosition string `json:"gbTextPosition,omitempty"`

	// Kind: Resource type for a reading position.
	Kind string `json:"kind,omitempty"`

	// PdfPosition: Position in a PDF file.
	PdfPosition string `json:"pdfPosition,omitempty"`

	// Updated: Timestamp when this reading position was last updated
	// (formatted UTC timestamp with millisecond resolution).
	Updated string `json:"updated,omitempty"`

	// VolumeId: Volume id associated with this reading position.
	VolumeId string `json:"volumeId,omitempty"`
}

type RequestAccess struct {
	// ConcurrentAccess: A concurrent access response.
	ConcurrentAccess *ConcurrentAccessRestriction `json:"concurrentAccess,omitempty"`

	// DownloadAccess: A download access response.
	DownloadAccess *DownloadAccessRestriction `json:"downloadAccess,omitempty"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`
}

type Review struct {
	// Author: Author of this review.
	Author *ReviewAuthor `json:"author,omitempty"`

	// Content: Review text.
	Content string `json:"content,omitempty"`

	// Date: Date of this review.
	Date string `json:"date,omitempty"`

	// FullTextUrl: URL for the full review text, for reviews gathered from
	// the web.
	FullTextUrl string `json:"fullTextUrl,omitempty"`

	// Kind: Resource type for a review.
	Kind string `json:"kind,omitempty"`

	// Rating: Star rating for this review. Possible values are ONE, TWO,
	// THREE, FOUR, FIVE or NOT_RATED.
	Rating string `json:"rating,omitempty"`

	// Source: Information regarding the source of this review, when the
	// review is not from a Google Books user.
	Source *ReviewSource `json:"source,omitempty"`

	// Title: Title for this review.
	Title string `json:"title,omitempty"`

	// Type: Source type for this review. Possible values are EDITORIAL,
	// WEB_USER or GOOGLE_USER.
	Type string `json:"type,omitempty"`

	// VolumeId: Volume that this review is for.
	VolumeId string `json:"volumeId,omitempty"`
}

type ReviewAuthor struct {
	// DisplayName: Name of this person.
	DisplayName string `json:"displayName,omitempty"`
}

type ReviewSource struct {
	// Description: Name of the source.
	Description string `json:"description,omitempty"`

	// ExtraDescription: Extra text about the source of the review.
	ExtraDescription string `json:"extraDescription,omitempty"`

	// Url: URL of the source of the review.
	Url string `json:"url,omitempty"`
}

type Volume struct {
	// AccessInfo: Any information about a volume related to reading or
	// obtaining that volume text. This information can depend on country
	// (books may be public domain in one country but not in another, e.g.).
	AccessInfo *VolumeAccessInfo `json:"accessInfo,omitempty"`

	// Etag: Opaque identifier for a specific version of a volume resource.
	// (In LITE projection)
	Etag string `json:"etag,omitempty"`

	// Id: Unique identifier for a volume. (In LITE projection.)
	Id string `json:"id,omitempty"`

	// Kind: Resource type for a volume. (In LITE projection.)
	Kind string `json:"kind,omitempty"`

	// LayerInfo: What layers exist in this volume and high level
	// information about them.
	LayerInfo *VolumeLayerInfo `json:"layerInfo,omitempty"`

	// RecommendedInfo: Recommendation related information for this volume.
	RecommendedInfo *VolumeRecommendedInfo `json:"recommendedInfo,omitempty"`

	// SaleInfo: Any information about a volume related to the eBookstore
	// and/or purchaseability. This information can depend on the country
	// where the request originates from (i.e. books may not be for sale in
	// certain countries).
	SaleInfo *VolumeSaleInfo `json:"saleInfo,omitempty"`

	// SearchInfo: Search result information related to this volume.
	SearchInfo *VolumeSearchInfo `json:"searchInfo,omitempty"`

	// SelfLink: URL to this resource. (In LITE projection.)
	SelfLink string `json:"selfLink,omitempty"`

	// UserInfo: User specific information related to this volume. (e.g.
	// page this user last read or whether they purchased this book)
	UserInfo *VolumeUserInfo `json:"userInfo,omitempty"`

	// VolumeInfo: General volume information.
	VolumeInfo *VolumeVolumeInfo `json:"volumeInfo,omitempty"`
}

type VolumeAccessInfo struct {
	// AccessViewStatus: Combines the access and viewability of this volume
	// into a single status field for this user. Values can be
	// FULL_PURCHASED, FULL_PUBLIC_DOMAIN, SAMPLE or NONE. (In LITE
	// projection.)
	AccessViewStatus string `json:"accessViewStatus,omitempty"`

	// Country: The two-letter ISO_3166-1 country code for which this access
	// information is valid. (In LITE projection.)
	Country string `json:"country,omitempty"`

	// DownloadAccess: Information about a volume's download license access
	// restrictions.
	DownloadAccess *DownloadAccessRestriction `json:"downloadAccess,omitempty"`

	// Embeddable: Whether this volume can be embedded in a viewport using
	// the Embedded Viewer API.
	Embeddable bool `json:"embeddable,omitempty"`

	// Epub: Information about epub content. (In LITE projection.)
	Epub *VolumeAccessInfoEpub `json:"epub,omitempty"`

	// Pdf: Information about pdf content. (In LITE projection.)
	Pdf *VolumeAccessInfoPdf `json:"pdf,omitempty"`

	// PublicDomain: Whether or not this book is public domain in the
	// country listed above.
	PublicDomain bool `json:"publicDomain,omitempty"`

	// TextToSpeechPermission: Whether text-to-speech is permitted for this
	// volume. Values can be ALLOWED, ALLOWED_FOR_ACCESSIBILITY, or
	// NOT_ALLOWED.
	TextToSpeechPermission string `json:"textToSpeechPermission,omitempty"`

	// ViewOrderUrl: For ordered but not yet processed orders, we give a URL
	// that can be used to go to the appropriate Google Wallet page.
	ViewOrderUrl string `json:"viewOrderUrl,omitempty"`

	// Viewability: The read access of a volume. Possible values are
	// PARTIAL, ALL_PAGES, NO_PAGES or UNKNOWN. This value depends on the
	// country listed above. A value of PARTIAL means that the publisher has
	// allowed some portion of the volume to be viewed publicly, without
	// purchase. This can apply to eBooks as well as non-eBooks. Public
	// domain books will always have a value of ALL_PAGES.
	Viewability string `json:"viewability,omitempty"`

	// WebReaderLink: URL to read this volume on the Google Books site. Link
	// will not allow users to read non-viewable volumes.
	WebReaderLink string `json:"webReaderLink,omitempty"`
}

type VolumeAccessInfoEpub struct {
	// AcsTokenLink: URL to retrieve ACS token for epub download. (In LITE
	// projection.)
	AcsTokenLink string `json:"acsTokenLink,omitempty"`

	// DownloadLink: URL to download epub. (In LITE projection.)
	DownloadLink string `json:"downloadLink,omitempty"`

	// IsAvailable: Is a flowing text epub available either as public domain
	// or for purchase. (In LITE projection.)
	IsAvailable bool `json:"isAvailable,omitempty"`
}

type VolumeAccessInfoPdf struct {
	// AcsTokenLink: URL to retrieve ACS token for pdf download. (In LITE
	// projection.)
	AcsTokenLink string `json:"acsTokenLink,omitempty"`

	// DownloadLink: URL to download pdf. (In LITE projection.)
	DownloadLink string `json:"downloadLink,omitempty"`

	// IsAvailable: Is a scanned image pdf available either as public domain
	// or for purchase. (In LITE projection.)
	IsAvailable bool `json:"isAvailable,omitempty"`
}

type VolumeLayerInfo struct {
	// Layers: A layer should appear here if and only if the layer exists
	// for this book.
	Layers []*VolumeLayerInfoLayers `json:"layers,omitempty"`
}

type VolumeLayerInfoLayers struct {
	// LayerId: The layer id of this layer (e.g. "geo").
	LayerId string `json:"layerId,omitempty"`

	// VolumeAnnotationsVersion: The current version of this layer's volume
	// annotations. Note that this version applies only to the data in the
	// books.layers.volumeAnnotations.* responses. The actual annotation
	// data is versioned separately.
	VolumeAnnotationsVersion string `json:"volumeAnnotationsVersion,omitempty"`
}

type VolumeRecommendedInfo struct {
	// Explanation: A text explaining why this volume is recommended.
	Explanation string `json:"explanation,omitempty"`
}

type VolumeSaleInfo struct {
	// BuyLink: URL to purchase this volume on the Google Books site. (In
	// LITE projection)
	BuyLink string `json:"buyLink,omitempty"`

	// Country: The two-letter ISO_3166-1 country code for which this sale
	// information is valid. (In LITE projection.)
	Country string `json:"country,omitempty"`

	// IsEbook: Whether or not this volume is an eBook (can be added to the
	// My eBooks shelf).
	IsEbook bool `json:"isEbook,omitempty"`

	// ListPrice: Suggested retail price. (In LITE projection.)
	ListPrice *VolumeSaleInfoListPrice `json:"listPrice,omitempty"`

	// OnSaleDate: The date on which this book is available for sale.
	OnSaleDate string `json:"onSaleDate,omitempty"`

	// RetailPrice: The actual selling price of the book. This is the same
	// as the suggested retail or list price unless there are offers or
	// discounts on this volume. (In LITE projection.)
	RetailPrice *VolumeSaleInfoRetailPrice `json:"retailPrice,omitempty"`

	// Saleability: Whether or not this book is available for sale or
	// offered for free in the Google eBookstore for the country listed
	// above. Possible values are FOR_SALE, FREE, NOT_FOR_SALE, or
	// FOR_PREORDER.
	Saleability string `json:"saleability,omitempty"`
}

type VolumeSaleInfoListPrice struct {
	// Amount: Amount in the currency listed below. (In LITE projection.)
	Amount float64 `json:"amount,omitempty"`

	// CurrencyCode: An ISO 4217, three-letter currency code. (In LITE
	// projection.)
	CurrencyCode string `json:"currencyCode,omitempty"`
}

type VolumeSaleInfoRetailPrice struct {
	// Amount: Amount in the currency listed below. (In LITE projection.)
	Amount float64 `json:"amount,omitempty"`

	// CurrencyCode: An ISO 4217, three-letter currency code. (In LITE
	// projection.)
	CurrencyCode string `json:"currencyCode,omitempty"`
}

type VolumeSearchInfo struct {
	// TextSnippet: A text snippet containing the search query.
	TextSnippet string `json:"textSnippet,omitempty"`
}

type VolumeUserInfo struct {
	// IsInMyBooks: Whether or not this volume is currently in "my books."
	IsInMyBooks bool `json:"isInMyBooks,omitempty"`

	// IsPreordered: Whether or not this volume was pre-ordered by the
	// authenticated user making the request. (In LITE projection.)
	IsPreordered bool `json:"isPreordered,omitempty"`

	// IsPurchased: Whether or not this volume was purchased by the
	// authenticated user making the request. (In LITE projection.)
	IsPurchased bool `json:"isPurchased,omitempty"`

	// ReadingPosition: The user's current reading position in the volume,
	// if one is available. (In LITE projection.)
	ReadingPosition *ReadingPosition `json:"readingPosition,omitempty"`

	// Review: This user's review of this volume, if one exists.
	Review *Review `json:"review,omitempty"`

	// Updated: Timestamp when this volume was last modified by a user
	// action, such as a reading position update, volume purchase or writing
	// a review. (RFC 3339 UTC date-time format).
	Updated string `json:"updated,omitempty"`
}

type VolumeVolumeInfo struct {
	// Authors: The names of the authors and/or editors for this volume. (In
	// LITE projection)
	Authors []string `json:"authors,omitempty"`

	// AverageRating: The mean review rating for this volume. (min = 1.0,
	// max = 5.0)
	AverageRating float64 `json:"averageRating,omitempty"`

	// CanonicalVolumeLink: Canonical URL for a volume. (In LITE
	// projection.)
	CanonicalVolumeLink string `json:"canonicalVolumeLink,omitempty"`

	// Categories: A list of subject categories, such as "Fiction",
	// "Suspense", etc.
	Categories []string `json:"categories,omitempty"`

	// ContentVersion: An identifier for the version of the volume content
	// (text & images). (In LITE projection)
	ContentVersion string `json:"contentVersion,omitempty"`

	// Description: A synopsis of the volume. The text of the description is
	// formatted in HTML and includes simple formatting elements, such as b,
	// i, and br tags. (In LITE projection.)
	Description string `json:"description,omitempty"`

	// Dimensions: Physical dimensions of this volume.
	Dimensions *VolumeVolumeInfoDimensions `json:"dimensions,omitempty"`

	// ImageLinks: A list of image links for all the sizes that are
	// available. (In LITE projection.)
	ImageLinks *VolumeVolumeInfoImageLinks `json:"imageLinks,omitempty"`

	// IndustryIdentifiers: Industry standard identifiers for this volume.
	IndustryIdentifiers []*VolumeVolumeInfoIndustryIdentifiers `json:"industryIdentifiers,omitempty"`

	// InfoLink: URL to view information about this volume on the Google
	// Books site. (In LITE projection)
	InfoLink string `json:"infoLink,omitempty"`

	// Language: Best language for this volume (based on content). It is the
	// two-letter ISO 639-1 code such as 'fr', 'en', etc.
	Language string `json:"language,omitempty"`

	// MainCategory: The main category to which this volume belongs. It will
	// be the category from the categories list returned below that has the
	// highest weight.
	MainCategory string `json:"mainCategory,omitempty"`

	// PageCount: Total number of pages.
	PageCount int64 `json:"pageCount,omitempty"`

	// PreviewLink: URL to preview this volume on the Google Books site.
	PreviewLink string `json:"previewLink,omitempty"`

	// PrintType: Type of publication of this volume. Possible values are
	// BOOK or MAGAZINE.
	PrintType string `json:"printType,omitempty"`

	// PublishedDate: Date of publication. (In LITE projection.)
	PublishedDate string `json:"publishedDate,omitempty"`

	// Publisher: Publisher of this volume. (In LITE projection.)
	Publisher string `json:"publisher,omitempty"`

	// RatingsCount: The number of review ratings for this volume.
	RatingsCount int64 `json:"ratingsCount,omitempty"`

	// Subtitle: Volume subtitle. (In LITE projection.)
	Subtitle string `json:"subtitle,omitempty"`

	// Title: Volume title. (In LITE projection.)
	Title string `json:"title,omitempty"`
}

type VolumeVolumeInfoDimensions struct {
	// Height: Height or length of this volume (in cm).
	Height string `json:"height,omitempty"`

	// Thickness: Thickness of this volume (in cm).
	Thickness string `json:"thickness,omitempty"`

	// Width: Width of this volume (in cm).
	Width string `json:"width,omitempty"`
}

type VolumeVolumeInfoImageLinks struct {
	// ExtraLarge: Image link for extra large size (width of ~1280 pixels).
	// (In LITE projection)
	ExtraLarge string `json:"extraLarge,omitempty"`

	// Large: Image link for large size (width of ~800 pixels). (In LITE
	// projection)
	Large string `json:"large,omitempty"`

	// Medium: Image link for medium size (width of ~575 pixels). (In LITE
	// projection)
	Medium string `json:"medium,omitempty"`

	// Small: Image link for small size (width of ~300 pixels). (In LITE
	// projection)
	Small string `json:"small,omitempty"`

	// SmallThumbnail: Image link for small thumbnail size (width of ~80
	// pixels). (In LITE projection)
	SmallThumbnail string `json:"smallThumbnail,omitempty"`

	// Thumbnail: Image link for thumbnail size (width of ~128 pixels). (In
	// LITE projection)
	Thumbnail string `json:"thumbnail,omitempty"`
}

type VolumeVolumeInfoIndustryIdentifiers struct {
	// Identifier: Industry specific volume identifier.
	Identifier string `json:"identifier,omitempty"`

	// Type: Identifier type. Possible values are ISBN_10, ISBN_13, ISSN and
	// OTHER.
	Type string `json:"type,omitempty"`
}

type Volumeannotation struct {
	// AnnotationDataId: The annotation data id for this volume annotation.
	AnnotationDataId string `json:"annotationDataId,omitempty"`

	// AnnotationDataLink: Link to get data for this annotation.
	AnnotationDataLink string `json:"annotationDataLink,omitempty"`

	// AnnotationType: The type of annotation this is.
	AnnotationType string `json:"annotationType,omitempty"`

	// ContentRanges: The content ranges to identify the selected text.
	ContentRanges *VolumeannotationContentRanges `json:"contentRanges,omitempty"`

	// Data: Data for this annotation.
	Data string `json:"data,omitempty"`

	// Deleted: Indicates that this annotation is deleted.
	Deleted bool `json:"deleted,omitempty"`

	// Id: Unique id of this volume annotation.
	Id string `json:"id,omitempty"`

	// Kind: Resource Type
	Kind string `json:"kind,omitempty"`

	// LayerId: The Layer this annotation is for.
	LayerId string `json:"layerId,omitempty"`

	// PageIds: Pages the annotation spans.
	PageIds []string `json:"pageIds,omitempty"`

	// SelectedText: Excerpt from the volume.
	SelectedText string `json:"selectedText,omitempty"`

	// SelfLink: URL to this resource.
	SelfLink string `json:"selfLink,omitempty"`

	// Updated: Timestamp for the last time this anntoation was updated.
	// (RFC 3339 UTC date-time format).
	Updated string `json:"updated,omitempty"`

	// VolumeId: The Volume this annotation is for.
	VolumeId string `json:"volumeId,omitempty"`
}

type VolumeannotationContentRanges struct {
	// CfiRange: Range in CFI format for this annotation for version above.
	CfiRange *BooksAnnotationsRange `json:"cfiRange,omitempty"`

	// ContentVersion: Content version applicable to ranges below.
	ContentVersion string `json:"contentVersion,omitempty"`

	// GbImageRange: Range in GB image format for this annotation for
	// version above.
	GbImageRange *BooksAnnotationsRange `json:"gbImageRange,omitempty"`

	// GbTextRange: Range in GB text format for this annotation for version
	// above.
	GbTextRange *BooksAnnotationsRange `json:"gbTextRange,omitempty"`
}

type Volumeannotations struct {
	// Items: A list of volume annotations.
	Items []*Volumeannotation `json:"items,omitempty"`

	// Kind: Resource type
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token to pass in for pagination for the next page.
	// This will not be present if this request does not have more results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalItems: The total number of volume annotations found.
	TotalItems int64 `json:"totalItems,omitempty"`

	// Version: The version string for all of the volume annotations in this
	// layer (not just the ones in this response). Note: the version string
	// doesn't apply to the annotation data, just the information in this
	// response (e.g. the location of annotations in the book).
	Version string `json:"version,omitempty"`
}

type Volumes struct {
	// Items: A list of volumes.
	Items []*Volume `json:"items,omitempty"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`

	// TotalItems: Total number of volumes found. This might be greater than
	// the number of volumes returned in this response if results have been
	// paginated.
	TotalItems int64 `json:"totalItems,omitempty"`
}

// method id "books.bookshelves.get":

type BookshelvesGetCall struct {
	s      *Service
	userId string
	shelf  string
	opt_   map[string]interface{}
}

// Get: Retrieves metadata for a specific bookshelf for the specified
// user.
func (r *BookshelvesService) Get(userId string, shelf string) *BookshelvesGetCall {
	c := &BookshelvesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.shelf = shelf
	return c
}

// Source sets the optional parameter "source": String to identify the
// originator of this request.
func (c *BookshelvesGetCall) Source(source string) *BookshelvesGetCall {
	c.opt_["source"] = source
	return c
}

func (c *BookshelvesGetCall) Do() (*Bookshelf, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["source"]; ok {
		params.Set("source", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/books/v1/", "users/{userId}/bookshelves/{shelf}")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls = strings.Replace(urls, "{shelf}", cleanPathString(c.shelf), 1)
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
	ret := new(Bookshelf)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves metadata for a specific bookshelf for the specified user.",
	//   "httpMethod": "GET",
	//   "id": "books.bookshelves.get",
	//   "parameterOrder": [
	//     "userId",
	//     "shelf"
	//   ],
	//   "parameters": {
	//     "shelf": {
	//       "description": "ID of bookshelf to retrieve.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "source": {
	//       "description": "String to identify the originator of this request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of user for whom to retrieve bookshelves.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userId}/bookshelves/{shelf}",
	//   "response": {
	//     "$ref": "Bookshelf"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/books"
	//   ]
	// }

}

// method id "books.bookshelves.list":

type BookshelvesListCall struct {
	s      *Service
	userId string
	opt_   map[string]interface{}
}

// List: Retrieves a list of public bookshelves for the specified user.
func (r *BookshelvesService) List(userId string) *BookshelvesListCall {
	c := &BookshelvesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	return c
}

// Source sets the optional parameter "source": String to identify the
// originator of this request.
func (c *BookshelvesListCall) Source(source string) *BookshelvesListCall {
	c.opt_["source"] = source
	return c
}

func (c *BookshelvesListCall) Do() (*Bookshelves, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["source"]; ok {
		params.Set("source", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/books/v1/", "users/{userId}/bookshelves")
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
	ret := new(Bookshelves)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of public bookshelves for the specified user.",
	//   "httpMethod": "GET",
	//   "id": "books.bookshelves.list",
	//   "parameterOrder": [
	//     "userId"
	//   ],
	//   "parameters": {
	//     "source": {
	//       "description": "String to identify the originator of this request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of user for whom to retrieve bookshelves.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userId}/bookshelves",
	//   "response": {
	//     "$ref": "Bookshelves"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/books"
	//   ]
	// }

}

// method id "books.layers.get":

type LayersGetCall struct {
	s         *Service
	volumeId  string
	summaryId string
	opt_      map[string]interface{}
}

// Get: Gets the layer summary for a volume.
func (r *LayersService) Get(volumeId string, summaryId string) *LayersGetCall {
	c := &LayersGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.volumeId = volumeId
	c.summaryId = summaryId
	return c
}

// ContentVersion sets the optional parameter "contentVersion": The
// content version for the requested volume.
func (c *LayersGetCall) ContentVersion(contentVersion string) *LayersGetCall {
	c.opt_["contentVersion"] = contentVersion
	return c
}

// Source sets the optional parameter "source": String to identify the
// originator of this request.
func (c *LayersGetCall) Source(source string) *LayersGetCall {
	c.opt_["source"] = source
	return c
}

func (c *LayersGetCall) Do() (*Layersummary, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["contentVersion"]; ok {
		params.Set("contentVersion", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["source"]; ok {
		params.Set("source", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/books/v1/", "volumes/{volumeId}/layersummary/{summaryId}")
	urls = strings.Replace(urls, "{volumeId}", cleanPathString(c.volumeId), 1)
	urls = strings.Replace(urls, "{summaryId}", cleanPathString(c.summaryId), 1)
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
	ret := new(Layersummary)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the layer summary for a volume.",
	//   "httpMethod": "GET",
	//   "id": "books.layers.get",
	//   "parameterOrder": [
	//     "volumeId",
	//     "summaryId"
	//   ],
	//   "parameters": {
	//     "contentVersion": {
	//       "description": "The content version for the requested volume.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "source": {
	//       "description": "String to identify the originator of this request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "summaryId": {
	//       "description": "The ID for the layer to get the summary for.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "volumeId": {
	//       "description": "The volume to retrieve layers for.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "volumes/{volumeId}/layersummary/{summaryId}",
	//   "response": {
	//     "$ref": "Layersummary"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/books"
	//   ]
	// }

}

// method id "books.layers.list":

type LayersListCall struct {
	s        *Service
	volumeId string
	opt_     map[string]interface{}
}

// List: List the layer summaries for a volume.
func (r *LayersService) List(volumeId string) *LayersListCall {
	c := &LayersListCall{s: r.s, opt_: make(map[string]interface{})}
	c.volumeId = volumeId
	return c
}

// ContentVersion sets the optional parameter "contentVersion": The
// content version for the requested volume.
func (c *LayersListCall) ContentVersion(contentVersion string) *LayersListCall {
	c.opt_["contentVersion"] = contentVersion
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return
func (c *LayersListCall) MaxResults(maxResults int64) *LayersListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": The value of the
// nextToken from the previous page.
func (c *LayersListCall) PageToken(pageToken string) *LayersListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Source sets the optional parameter "source": String to identify the
// originator of this request.
func (c *LayersListCall) Source(source string) *LayersListCall {
	c.opt_["source"] = source
	return c
}

func (c *LayersListCall) Do() (*Layersummaries, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["contentVersion"]; ok {
		params.Set("contentVersion", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["source"]; ok {
		params.Set("source", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/books/v1/", "volumes/{volumeId}/layersummary")
	urls = strings.Replace(urls, "{volumeId}", cleanPathString(c.volumeId), 1)
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
	ret := new(Layersummaries)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List the layer summaries for a volume.",
	//   "httpMethod": "GET",
	//   "id": "books.layers.list",
	//   "parameterOrder": [
	//     "volumeId"
	//   ],
	//   "parameters": {
	//     "contentVersion": {
	//       "description": "The content version for the requested volume.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "200",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The value of the nextToken from the previous page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "source": {
	//       "description": "String to identify the originator of this request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "volumeId": {
	//       "description": "The volume to retrieve layers for.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "volumes/{volumeId}/layersummary",
	//   "response": {
	//     "$ref": "Layersummaries"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/books"
	//   ]
	// }

}

// method id "books.myconfig.releaseDownloadAccess":

type MyconfigReleaseDownloadAccessCall struct {
	s         *Service
	volumeIds []string
	cpksver   string
	opt_      map[string]interface{}
}

// ReleaseDownloadAccess: Release downloaded content access restriction.
func (r *MyconfigService) ReleaseDownloadAccess(volumeIds []string, cpksver string) *MyconfigReleaseDownloadAccessCall {
	c := &MyconfigReleaseDownloadAccessCall{s: r.s, opt_: make(map[string]interface{})}
	c.volumeIds = volumeIds
	c.cpksver = cpksver
	return c
}

// Locale sets the optional parameter "locale": ISO-639-1, ISO-3166-1
// codes for message localization, i.e. en_US.
func (c *MyconfigReleaseDownloadAccessCall) Locale(locale string) *MyconfigReleaseDownloadAccessCall {
	c.opt_["locale"] = locale
	return c
}

// Source sets the optional parameter "source": String to identify the
// originator of this request.
func (c *MyconfigReleaseDownloadAccessCall) Source(source string) *MyconfigReleaseDownloadAccessCall {
	c.opt_["source"] = source
	return c
}

func (c *MyconfigReleaseDownloadAccessCall) Do() (*DownloadAccesses, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("cpksver", fmt.Sprintf("%v", c.cpksver))
	for _, v := range c.volumeIds {
		params.Add("volumeIds", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["locale"]; ok {
		params.Set("locale", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["source"]; ok {
		params.Set("source", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/books/v1/", "myconfig/releaseDownloadAccess")
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
	ret := new(DownloadAccesses)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Release downloaded content access restriction.",
	//   "httpMethod": "POST",
	//   "id": "books.myconfig.releaseDownloadAccess",
	//   "parameterOrder": [
	//     "volumeIds",
	//     "cpksver"
	//   ],
	//   "parameters": {
	//     "cpksver": {
	//       "description": "The device/version ID from which to release the restriction.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "locale": {
	//       "description": "ISO-639-1, ISO-3166-1 codes for message localization, i.e. en_US.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "source": {
	//       "description": "String to identify the originator of this request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "volumeIds": {
	//       "description": "The volume(s) to release restrictions for.",
	//       "location": "query",
	//       "repeated": true,
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "myconfig/releaseDownloadAccess",
	//   "response": {
	//     "$ref": "DownloadAccesses"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/books"
	//   ]
	// }

}

// method id "books.myconfig.requestAccess":

type MyconfigRequestAccessCall struct {
	s        *Service
	source   string
	volumeId string
	nonce    string
	cpksver  string
	opt_     map[string]interface{}
}

// RequestAccess: Request concurrent and download access restrictions.
func (r *MyconfigService) RequestAccess(source string, volumeId string, nonce string, cpksver string) *MyconfigRequestAccessCall {
	c := &MyconfigRequestAccessCall{s: r.s, opt_: make(map[string]interface{})}
	c.source = source
	c.volumeId = volumeId
	c.nonce = nonce
	c.cpksver = cpksver
	return c
}

// Locale sets the optional parameter "locale": ISO-639-1, ISO-3166-1
// codes for message localization, i.e. en_US.
func (c *MyconfigRequestAccessCall) Locale(locale string) *MyconfigRequestAccessCall {
	c.opt_["locale"] = locale
	return c
}

func (c *MyconfigRequestAccessCall) Do() (*RequestAccess, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("cpksver", fmt.Sprintf("%v", c.cpksver))
	params.Set("nonce", fmt.Sprintf("%v", c.nonce))
	params.Set("source", fmt.Sprintf("%v", c.source))
	params.Set("volumeId", fmt.Sprintf("%v", c.volumeId))
	if v, ok := c.opt_["locale"]; ok {
		params.Set("locale", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/books/v1/", "myconfig/requestAccess")
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
	ret := new(RequestAccess)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Request concurrent and download access restrictions.",
	//   "httpMethod": "POST",
	//   "id": "books.myconfig.requestAccess",
	//   "parameterOrder": [
	//     "source",
	//     "volumeId",
	//     "nonce",
	//     "cpksver"
	//   ],
	//   "parameters": {
	//     "cpksver": {
	//       "description": "The device/version ID from which to request the restrictions.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "locale": {
	//       "description": "ISO-639-1, ISO-3166-1 codes for message localization, i.e. en_US.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "nonce": {
	//       "description": "The client nonce value.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "source": {
	//       "description": "String to identify the originator of this request.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "volumeId": {
	//       "description": "The volume to request concurrent/download restrictions for.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "myconfig/requestAccess",
	//   "response": {
	//     "$ref": "RequestAccess"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/books"
	//   ]
	// }

}

// method id "books.myconfig.syncVolumeLicenses":

type MyconfigSyncVolumeLicensesCall struct {
	s       *Service
	source  string
	nonce   string
	cpksver string
	opt_    map[string]interface{}
}

// SyncVolumeLicenses: Request downloaded content access for specified
// volumes on the My eBooks shelf.
func (r *MyconfigService) SyncVolumeLicenses(source string, nonce string, cpksver string) *MyconfigSyncVolumeLicensesCall {
	c := &MyconfigSyncVolumeLicensesCall{s: r.s, opt_: make(map[string]interface{})}
	c.source = source
	c.nonce = nonce
	c.cpksver = cpksver
	return c
}

// Locale sets the optional parameter "locale": ISO-639-1, ISO-3166-1
// codes for message localization, i.e. en_US.
func (c *MyconfigSyncVolumeLicensesCall) Locale(locale string) *MyconfigSyncVolumeLicensesCall {
	c.opt_["locale"] = locale
	return c
}

// ShowPreorders sets the optional parameter "showPreorders": Set to
// true to show pre-ordered books. Defaults to false.
func (c *MyconfigSyncVolumeLicensesCall) ShowPreorders(showPreorders bool) *MyconfigSyncVolumeLicensesCall {
	c.opt_["showPreorders"] = showPreorders
	return c
}

// VolumeIds sets the optional parameter "volumeIds": The volume(s) to
// request download restrictions for.
func (c *MyconfigSyncVolumeLicensesCall) VolumeIds(volumeIds string) *MyconfigSyncVolumeLicensesCall {
	c.opt_["volumeIds"] = volumeIds
	return c
}

func (c *MyconfigSyncVolumeLicensesCall) Do() (*Volumes, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("cpksver", fmt.Sprintf("%v", c.cpksver))
	params.Set("nonce", fmt.Sprintf("%v", c.nonce))
	params.Set("source", fmt.Sprintf("%v", c.source))
	if v, ok := c.opt_["locale"]; ok {
		params.Set("locale", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["showPreorders"]; ok {
		params.Set("showPreorders", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["volumeIds"]; ok {
		params.Set("volumeIds", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/books/v1/", "myconfig/syncVolumeLicenses")
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
	ret := new(Volumes)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Request downloaded content access for specified volumes on the My eBooks shelf.",
	//   "httpMethod": "POST",
	//   "id": "books.myconfig.syncVolumeLicenses",
	//   "parameterOrder": [
	//     "source",
	//     "nonce",
	//     "cpksver"
	//   ],
	//   "parameters": {
	//     "cpksver": {
	//       "description": "The device/version ID from which to release the restriction.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "locale": {
	//       "description": "ISO-639-1, ISO-3166-1 codes for message localization, i.e. en_US.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "nonce": {
	//       "description": "The client nonce value.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "showPreorders": {
	//       "description": "Set to true to show pre-ordered books. Defaults to false.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "source": {
	//       "description": "String to identify the originator of this request.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "volumeIds": {
	//       "description": "The volume(s) to request download restrictions for.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "myconfig/syncVolumeLicenses",
	//   "response": {
	//     "$ref": "Volumes"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/books"
	//   ]
	// }

}

// method id "books.volumes.get":

type VolumesGetCall struct {
	s        *Service
	volumeId string
	opt_     map[string]interface{}
}

// Get: Gets volume information for a single volume.
func (r *VolumesService) Get(volumeId string) *VolumesGetCall {
	c := &VolumesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.volumeId = volumeId
	return c
}

// Country sets the optional parameter "country": ISO-3166-1 code to
// override the IP-based location.
func (c *VolumesGetCall) Country(country string) *VolumesGetCall {
	c.opt_["country"] = country
	return c
}

// Partner sets the optional parameter "partner": Brand results for
// partner ID.
func (c *VolumesGetCall) Partner(partner string) *VolumesGetCall {
	c.opt_["partner"] = partner
	return c
}

// Projection sets the optional parameter "projection": Restrict
// information returned to a set of selected fields.
func (c *VolumesGetCall) Projection(projection string) *VolumesGetCall {
	c.opt_["projection"] = projection
	return c
}

// Source sets the optional parameter "source": String to identify the
// originator of this request.
func (c *VolumesGetCall) Source(source string) *VolumesGetCall {
	c.opt_["source"] = source
	return c
}

func (c *VolumesGetCall) Do() (*Volume, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["country"]; ok {
		params.Set("country", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["partner"]; ok {
		params.Set("partner", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["projection"]; ok {
		params.Set("projection", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["source"]; ok {
		params.Set("source", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/books/v1/", "volumes/{volumeId}")
	urls = strings.Replace(urls, "{volumeId}", cleanPathString(c.volumeId), 1)
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
	ret := new(Volume)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets volume information for a single volume.",
	//   "httpMethod": "GET",
	//   "id": "books.volumes.get",
	//   "parameterOrder": [
	//     "volumeId"
	//   ],
	//   "parameters": {
	//     "country": {
	//       "description": "ISO-3166-1 code to override the IP-based location.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "partner": {
	//       "description": "Brand results for partner ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Restrict information returned to a set of selected fields.",
	//       "enum": [
	//         "full",
	//         "lite"
	//       ],
	//       "enumDescriptions": [
	//         "Includes all volume data.",
	//         "Includes a subset of fields in volumeInfo and accessInfo."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "source": {
	//       "description": "String to identify the originator of this request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "volumeId": {
	//       "description": "ID of volume to retrieve.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "volumes/{volumeId}",
	//   "response": {
	//     "$ref": "Volume"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/books"
	//   ]
	// }

}

// method id "books.volumes.list":

type VolumesListCall struct {
	s    *Service
	q    string
	opt_ map[string]interface{}
}

// List: Performs a book search.
func (r *VolumesService) List(q string) *VolumesListCall {
	c := &VolumesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.q = q
	return c
}

// Download sets the optional parameter "download": Restrict to volumes
// by download availability.
func (c *VolumesListCall) Download(download string) *VolumesListCall {
	c.opt_["download"] = download
	return c
}

// Filter sets the optional parameter "filter": Filter search results.
func (c *VolumesListCall) Filter(filter string) *VolumesListCall {
	c.opt_["filter"] = filter
	return c
}

// LangRestrict sets the optional parameter "langRestrict": Restrict
// results to books with this language code.
func (c *VolumesListCall) LangRestrict(langRestrict string) *VolumesListCall {
	c.opt_["langRestrict"] = langRestrict
	return c
}

// LibraryRestrict sets the optional parameter "libraryRestrict":
// Restrict search to this user's library.
func (c *VolumesListCall) LibraryRestrict(libraryRestrict string) *VolumesListCall {
	c.opt_["libraryRestrict"] = libraryRestrict
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return.
func (c *VolumesListCall) MaxResults(maxResults int64) *VolumesListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// OrderBy sets the optional parameter "orderBy": Sort search results.
func (c *VolumesListCall) OrderBy(orderBy string) *VolumesListCall {
	c.opt_["orderBy"] = orderBy
	return c
}

// Partner sets the optional parameter "partner": Restrict and brand
// results for partner ID.
func (c *VolumesListCall) Partner(partner string) *VolumesListCall {
	c.opt_["partner"] = partner
	return c
}

// PrintType sets the optional parameter "printType": Restrict to books
// or magazines.
func (c *VolumesListCall) PrintType(printType string) *VolumesListCall {
	c.opt_["printType"] = printType
	return c
}

// Projection sets the optional parameter "projection": Restrict
// information returned to a set of selected fields.
func (c *VolumesListCall) Projection(projection string) *VolumesListCall {
	c.opt_["projection"] = projection
	return c
}

// ShowPreorders sets the optional parameter "showPreorders": Set to
// true to show books available for preorder. Defaults to false.
func (c *VolumesListCall) ShowPreorders(showPreorders bool) *VolumesListCall {
	c.opt_["showPreorders"] = showPreorders
	return c
}

// Source sets the optional parameter "source": String to identify the
// originator of this request.
func (c *VolumesListCall) Source(source string) *VolumesListCall {
	c.opt_["source"] = source
	return c
}

// StartIndex sets the optional parameter "startIndex": Index of the
// first result to return (starts at 0)
func (c *VolumesListCall) StartIndex(startIndex int64) *VolumesListCall {
	c.opt_["startIndex"] = startIndex
	return c
}

func (c *VolumesListCall) Do() (*Volumes, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("q", fmt.Sprintf("%v", c.q))
	if v, ok := c.opt_["download"]; ok {
		params.Set("download", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["filter"]; ok {
		params.Set("filter", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["langRestrict"]; ok {
		params.Set("langRestrict", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["libraryRestrict"]; ok {
		params.Set("libraryRestrict", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["orderBy"]; ok {
		params.Set("orderBy", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["partner"]; ok {
		params.Set("partner", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["printType"]; ok {
		params.Set("printType", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["projection"]; ok {
		params.Set("projection", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["showPreorders"]; ok {
		params.Set("showPreorders", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["source"]; ok {
		params.Set("source", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["startIndex"]; ok {
		params.Set("startIndex", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/books/v1/", "volumes")
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
	ret := new(Volumes)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Performs a book search.",
	//   "httpMethod": "GET",
	//   "id": "books.volumes.list",
	//   "parameterOrder": [
	//     "q"
	//   ],
	//   "parameters": {
	//     "download": {
	//       "description": "Restrict to volumes by download availability.",
	//       "enum": [
	//         "epub"
	//       ],
	//       "enumDescriptions": [
	//         "All volumes with epub."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "filter": {
	//       "description": "Filter search results.",
	//       "enum": [
	//         "ebooks",
	//         "free-ebooks",
	//         "full",
	//         "paid-ebooks",
	//         "partial"
	//       ],
	//       "enumDescriptions": [
	//         "All Google eBooks.",
	//         "Google eBook with full volume text viewability.",
	//         "Public can view entire volume text.",
	//         "Google eBook with a price.",
	//         "Public able to see parts of text."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "langRestrict": {
	//       "description": "Restrict results to books with this language code.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "libraryRestrict": {
	//       "description": "Restrict search to this user's library.",
	//       "enum": [
	//         "my-library",
	//         "no-restrict"
	//       ],
	//       "enumDescriptions": [
	//         "Restrict to the user's library, any shelf.",
	//         "Do not restrict based on user's library."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "40",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "orderBy": {
	//       "description": "Sort search results.",
	//       "enum": [
	//         "newest",
	//         "relevance"
	//       ],
	//       "enumDescriptions": [
	//         "Most recently published.",
	//         "Relevance to search terms."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "partner": {
	//       "description": "Restrict and brand results for partner ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "printType": {
	//       "description": "Restrict to books or magazines.",
	//       "enum": [
	//         "all",
	//         "books",
	//         "magazines"
	//       ],
	//       "enumDescriptions": [
	//         "All volume content types.",
	//         "Just books.",
	//         "Just magazines."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Restrict information returned to a set of selected fields.",
	//       "enum": [
	//         "full",
	//         "lite"
	//       ],
	//       "enumDescriptions": [
	//         "Includes all volume data.",
	//         "Includes a subset of fields in volumeInfo and accessInfo."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "q": {
	//       "description": "Full-text search query string.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "showPreorders": {
	//       "description": "Set to true to show books available for preorder. Defaults to false.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "source": {
	//       "description": "String to identify the originator of this request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "startIndex": {
	//       "description": "Index of the first result to return (starts at 0)",
	//       "format": "uint32",
	//       "location": "query",
	//       "minimum": "0",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "volumes",
	//   "response": {
	//     "$ref": "Volumes"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/books"
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
