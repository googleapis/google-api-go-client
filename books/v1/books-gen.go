// Package books provides access to the Books API.
//
// See https://code.google.com/apis/books/docs/v1/getting_started.html
//
// Usage example:
//
//   import "google-api-go-client.googlecode.com/hg/books/v1"
//   ...
//   booksService, err := books.New(oauthHttpClient)
package books

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

const apiId = "books:v1"
const apiName = "books"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/books/v1/"

// OAuth2 scopes used by this API.
const (
	// Manage your books
	BooksScope = "https://www.googleapis.com/auth/books"
)

func New(client *http.Client) (*Service, os.Error) {
	if client == nil {
		return nil, os.NewError("client is nil")
	}
	s := &Service{client: client}
	s.Volumes = &VolumesService{s: s}
	s.Bookshelves = &BookshelvesService{s: s}
	s.Mylibrary = &MylibraryService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	Volumes *VolumesService

	Bookshelves *BookshelvesService

	Mylibrary *MylibraryService
}

type VolumesService struct {
	s *Service
}

type BookshelvesService struct {
	s *Service
}

type MylibraryService struct {
	s *Service
}

type VolumeSaleInfoListPrice struct {
	// CurrencyCode: An ISO 4217, three-letter currency code. (In LITE
	// projection.)
	CurrencyCode string `json:"currencyCode,omitempty"`

	// Amount: Amount in the currency listed below. (In LITE projection.)
	Amount float64 `json:"amount,omitempty"`
}

type ReadingPosition struct {
	// EpubCfiPosition: Position in an EPUB as a CFI.
	EpubCfiPosition string `json:"epubCfiPosition,omitempty"`

	// GbTextPosition: Position in a volume for text-based content.
	GbTextPosition string `json:"gbTextPosition,omitempty"`

	// VolumeId: Volume id associated with this reading position.
	VolumeId string `json:"volumeId,omitempty"`

	// Kind: Resource type for a reading position.
	Kind string `json:"kind,omitempty"`

	// Updated: Timestamp when this reading position was last updated
	// (formatted UTC timestamp with millisecond resolution).
	Updated string `json:"updated,omitempty"`

	// GbImagePosition: Position in a volume for image-based content.
	GbImagePosition string `json:"gbImagePosition,omitempty"`

	// PdfPosition: Position in a PDF file.
	PdfPosition string `json:"pdfPosition,omitempty"`
}

type ReviewAuthor struct {
	// DisplayName: Name of this person.
	DisplayName string `json:"displayName,omitempty"`
}

type Annotation struct {
	// CurrentVersionRanges: Selection ranges for the most recent content
	// version.
	CurrentVersionRanges *AnnotationCurrentVersionRanges `json:"currentVersionRanges,omitempty"`

	// SelfLink: URL to this resource.
	SelfLink string `json:"selfLink,omitempty"`

	// AfterSelectedText: Anchor text after excerpt.
	AfterSelectedText string `json:"afterSelectedText,omitempty"`

	// ClientVersionRanges: Selection ranges sent from the client.
	ClientVersionRanges *AnnotationClientVersionRanges `json:"clientVersionRanges,omitempty"`

	// VolumeId: The volume that this annotation belongs to.
	VolumeId string `json:"volumeId,omitempty"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`

	// Updated: Timestamp for the last time this annotation was modified.
	Updated string `json:"updated,omitempty"`

	// Data: User-created data for this annotation.
	Data string `json:"data,omitempty"`

	// Id: Id of this annotation, in the form of a GUID.
	Id string `json:"id,omitempty"`

	// Created: Timestamp for the created time of this annotation.
	Created string `json:"created,omitempty"`

	// HighlightStyle: The highlight style for this annotation.
	HighlightStyle string `json:"highlightStyle,omitempty"`

	// SelectedText: Excerpt from the volume.
	SelectedText string `json:"selectedText,omitempty"`

	// LayerId: The layer this annotation is for.
	LayerId string `json:"layerId,omitempty"`

	// PageIds: Pages that this annotation spans.
	PageIds []string `json:"pageIds,omitempty"`

	// BeforeSelectedText: Anchor text before excerpt.
	BeforeSelectedText string `json:"beforeSelectedText,omitempty"`
}

type VolumeSaleInfoRetailPrice struct {
	// CurrencyCode: An ISO 4217, three-letter currency code. (In LITE
	// projection.)
	CurrencyCode string `json:"currencyCode,omitempty"`

	// Amount: Amount in the currency listed below. (In LITE projection.)
	Amount float64 `json:"amount,omitempty"`
}

type Volume struct {
	// SelfLink: URL to this resource. (In LITE projection.)
	SelfLink string `json:"selfLink,omitempty"`

	// SaleInfo: Any information about a volume related to the eBookstore
	// and/or purchaseability. This information can depend on the country
	// where the request originates from (i.e. books may not be for sale in
	// certain countries).
	SaleInfo *VolumeSaleInfo `json:"saleInfo,omitempty"`

	// VolumeInfo: General volume information.
	VolumeInfo *VolumeVolumeInfo `json:"volumeInfo,omitempty"`

	// Etag: Opaque identifier for a specific version of a volume resource.
	// (In LITE projection)
	Etag string `json:"etag,omitempty"`

	// Kind: Resource type for a volume. (In LITE projection.)
	Kind string `json:"kind,omitempty"`

	// UserInfo: User specific information related to this volume. (e.g.
	// page this user last read or whether they purchased this book)
	UserInfo *VolumeUserInfo `json:"userInfo,omitempty"`

	// Id: Unique identifier for a volume. (In LITE projection.)
	Id string `json:"id,omitempty"`

	// AccessInfo: Any information about a volume related to reading or
	// obtaining that volume text. This information can depend on country
	// (books may be public domain in one country but not in another, e.g.).
	AccessInfo *VolumeAccessInfo `json:"accessInfo,omitempty"`
}

type VolumeUserInfo struct {
	// IsPreordered: Whether or not this volume was pre-ordered by the
	// authenticated user making the request. (In LITE projection.)
	IsPreordered bool `json:"isPreordered,omitempty"`

	// IsPurchased: Whether or not this volume was purchased by the
	// authenticated user making the request. (In LITE projection.)
	IsPurchased bool `json:"isPurchased,omitempty"`

	// Updated: Timestamp when this volume was last modified by a user
	// action, such as a reading position update, volume purchase or writing
	// a review. (RFC 3339 UTC date-time format).
	Updated string `json:"updated,omitempty"`

	// ReadingPosition: The user's current reading position in the volume,
	// if one is available. (In LITE projection.)
	ReadingPosition *ReadingPosition `json:"readingPosition,omitempty"`

	// Review: This user's review of this volume, if one exists.
	Review *Review `json:"review,omitempty"`
}

type AnnotationClientVersionRanges struct {
	// GbTextRange: Range in GB text format for this annotation sent by
	// client.
	GbTextRange *BooksAnnotationsRange `json:"gbTextRange,omitempty"`

	// ContentVersion: Content version the client sent in.
	ContentVersion string `json:"contentVersion,omitempty"`

	// CfiRange: Range in CFI format for this annotation sent by client.
	CfiRange *BooksAnnotationsRange `json:"cfiRange,omitempty"`

	// GbImageRange: Range in GB image format for this annotation sent by
	// client.
	GbImageRange *BooksAnnotationsRange `json:"gbImageRange,omitempty"`
}

type VolumeVolumeInfoImageLinks struct {
	// Large: Image link for large size (width of ~800 pixels). (In LITE
	// projection)
	Large string `json:"large,omitempty"`

	// Small: Image link for small size (width of ~300 pixels). (In LITE
	// projection)
	Small string `json:"small,omitempty"`

	// SmallThumbnail: Image link for small thumbnail size (width of ~80
	// pixels). (In LITE projection)
	SmallThumbnail string `json:"smallThumbnail,omitempty"`

	// Medium: Image link for medium size (width of ~575 pixels). (In LITE
	// projection)
	Medium string `json:"medium,omitempty"`

	// Thumbnail: Image link for thumbnail size (width of ~128 pixels). (In
	// LITE projection)
	Thumbnail string `json:"thumbnail,omitempty"`

	// ExtraLarge: Image link for extra large size (width of ~1280 pixels).
	// (In LITE projection)
	ExtraLarge string `json:"extraLarge,omitempty"`
}

type AnnotationCurrentVersionRanges struct {
	// GbTextRange: Range in GB text format for this annotation for version
	// above.
	GbTextRange *BooksAnnotationsRange `json:"gbTextRange,omitempty"`

	// ContentVersion: Content version applicable to ranges below.
	ContentVersion string `json:"contentVersion,omitempty"`

	// CfiRange: Range in CFI format for this annotation for version above.
	CfiRange *BooksAnnotationsRange `json:"cfiRange,omitempty"`

	// GbImageRange: Range in GB image format for this annotation for
	// version above.
	GbImageRange *BooksAnnotationsRange `json:"gbImageRange,omitempty"`
}

type VolumeAccessInfoEpub struct {
	// AcsTokenLink: URL to retrieve ACS token for epub download. (In LITE
	// projection.)
	AcsTokenLink string `json:"acsTokenLink,omitempty"`

	// DownloadLink: URL to download epub. (In LITE projection.)
	DownloadLink string `json:"downloadLink,omitempty"`
}

type VolumeVolumeInfoIndustryIdentifiers struct {
	// Identifier: Industry specific volume identifier.
	Identifier string `json:"identifier,omitempty"`

	// Type: Identifier type. Possible values are ISBN_10, ISBN_13, ISSN and
	// OTHER.
	Type string `json:"type,omitempty"`
}

type Volumes struct {
	// Items: A list of volumes.
	Items []*Volume `json:"items,omitempty"`

	// TotalItems: Total number of volumes found. This might be greater than
	// the number of volumes returned in this response if results have been
	// paginated.
	TotalItems int64 `json:"totalItems,omitempty"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`
}

type VolumeVolumeInfoDimensions struct {
	// Height: Height or length of this volume (in cm).
	Height string `json:"height,omitempty"`

	// Width: Width of this volume (in cm).
	Width string `json:"width,omitempty"`

	// Thickness: Thickness of this volume (in cm).
	Thickness string `json:"thickness,omitempty"`
}

type VolumeSaleInfo struct {
	// ListPrice: Suggested retail price. (In LITE projection.)
	ListPrice *VolumeSaleInfoListPrice `json:"listPrice,omitempty"`

	// Country: The two-letter ISO_3166-1 country code for which this sale
	// information is valid. (In LITE projection.)
	Country string `json:"country,omitempty"`

	// Saleability: Whether or not this book is available for sale or
	// offered for free in the Google eBookstore for the country listed
	// above. Possible values are FOR_SALE, FREE, NOT_FOR_SALE, or
	// FOR_PREORDER.
	Saleability string `json:"saleability,omitempty"`

	// RetailPrice: The actual selling price of the book. This is the same
	// as the suggested retail or list price unless there are offers or
	// discounts on this volume. (In LITE projection.)
	RetailPrice *VolumeSaleInfoRetailPrice `json:"retailPrice,omitempty"`

	// OnSaleDate: The date on which this book is available for sale.
	OnSaleDate string `json:"onSaleDate,omitempty"`

	// BuyLink: URL to purchase this volume on the Google Books site. (In
	// LITE projection)
	BuyLink string `json:"buyLink,omitempty"`

	// IsEbook: Whether or not this volume is an eBook (can be added to the
	// My eBooks shelf).
	IsEbook bool `json:"isEbook,omitempty"`
}

type Bookshelves struct {
	// Items: A list of bookshelves.
	Items []*Bookshelf `json:"items,omitempty"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`
}

type Bookshelf struct {
	// SelfLink: URL to this resource.
	SelfLink string `json:"selfLink,omitempty"`

	// Kind: Resource type for bookshelf metadata.
	Kind string `json:"kind,omitempty"`

	// Updated: Last modified time of this bookshelf (formatted UTC
	// timestamp with millisecond resolution).
	Updated string `json:"updated,omitempty"`

	// Access: Whether this bookshelf is PUBLIC or PRIVATE.
	Access string `json:"access,omitempty"`

	// Id: Id of this bookshelf, only unique by user.
	Id int64 `json:"id,omitempty"`

	// Created: Created time for this bookshelf (formatted UTC timestamp
	// with millisecond resolution).
	Created string `json:"created,omitempty"`

	// Title: Title of this bookshelf.
	Title string `json:"title,omitempty"`

	// VolumeCount: Number of volumes in this bookshelf.
	VolumeCount int64 `json:"volumeCount,omitempty"`

	// VolumesLastUpdated: Last time a volume was added or removed from this
	// bookshelf (formatted UTC timestamp with millisecond resolution).
	VolumesLastUpdated string `json:"volumesLastUpdated,omitempty"`

	// Description: Description of this bookshelf.
	Description string `json:"description,omitempty"`
}

type VolumeAccessInfoPdf struct {
	// AcsTokenLink: URL to retrieve ACS token for pdf download. (In LITE
	// projection.)
	AcsTokenLink string `json:"acsTokenLink,omitempty"`

	// DownloadLink: URL to download pdf. (In LITE projection.)
	DownloadLink string `json:"downloadLink,omitempty"`
}

type VolumeVolumeInfo struct {
	// Categories: A list of subject categories, such as "Fiction",
	// "Suspense", etc.
	Categories []string `json:"categories,omitempty"`

	// InfoLink: URL to view information about this volume on the Google
	// Books site. (In LITE projection)
	InfoLink string `json:"infoLink,omitempty"`

	// Subtitle: Volume subtitle. (In LITE projection.)
	Subtitle string `json:"subtitle,omitempty"`

	// Language: Best language for this volume (based on content). It is the
	// two-letter ISO 639-1 code such as 'fr', 'en', etc.
	Language string `json:"language,omitempty"`

	// AverageRating: The mean review rating for this volume. (min = 1.0,
	// max = 5.0)
	AverageRating float64 `json:"averageRating,omitempty"`

	// Authors: The names of the authors and/or editors for this volume. (In
	// LITE projection)
	Authors []string `json:"authors,omitempty"`

	// PreviewLink: URL to preview this volume on the Google Books site.
	PreviewLink string `json:"previewLink,omitempty"`

	// Publisher: Publisher of this volume. (In LITE projection.)
	Publisher string `json:"publisher,omitempty"`

	// PublishedDate: Date of publication. (In LITE projection.)
	PublishedDate string `json:"publishedDate,omitempty"`

	// ContentVersion: An identifier for the version of the volume content
	// (text & images). (In LITE projection)
	ContentVersion string `json:"contentVersion,omitempty"`

	// MainCategory: The main category to which this volume belongs. It will
	// be the category from the categories list returned below that has the
	// highest weight.
	MainCategory string `json:"mainCategory,omitempty"`

	// PrintType: Type of publication of this volume. Possible values are
	// BOOK or MAGAZINE.
	PrintType string `json:"printType,omitempty"`

	// PageCount: Total number of pages.
	PageCount int64 `json:"pageCount,omitempty"`

	// Title: Volume title. (In LITE projection.)
	Title string `json:"title,omitempty"`

	// Dimensions: Physical dimensions of this volume.
	Dimensions *VolumeVolumeInfoDimensions `json:"dimensions,omitempty"`

	// ImageLinks: A list of image links for all the sizes that are
	// available. (In LITE projection.)
	ImageLinks *VolumeVolumeInfoImageLinks `json:"imageLinks,omitempty"`

	// IndustryIdentifiers: Industry standard identifiers for this volume.
	IndustryIdentifiers []*VolumeVolumeInfoIndustryIdentifiers `json:"industryIdentifiers,omitempty"`

	// RatingsCount: The number of review ratings for this volume.
	RatingsCount int64 `json:"ratingsCount,omitempty"`

	// CanonicalVolumeLink: Canonical URL for a volume. Use this URL to plus
	// one a Google Book. (In LITE projection)
	CanonicalVolumeLink string `json:"canonicalVolumeLink,omitempty"`

	// Description: A synopsis of the volume. The text of the description is
	// formatted in HTML and includes simple formatting elements, such as b,
	// i, and br tags. (In LITE projection.)
	Description string `json:"description,omitempty"`
}

type DownloadAccessRestriction struct {
	// Restricted: Whether this volume has any download access restrictions.
	Restricted bool `json:"restricted,omitempty"`

	// DownloadsAcquired: If restricted, the number of content download
	// licenses already acquired (including the requesting client, if
	// licensed).
	DownloadsAcquired int64 `json:"downloadsAcquired,omitempty"`

	// JustAcquired: If deviceAllowed, whether access was just acquired with
	// this request.
	JustAcquired bool `json:"justAcquired,omitempty"`

	// VolumeId: Identifies the volume for which this entry applies.
	VolumeId string `json:"volumeId,omitempty"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`

	// Source: Client app identifier for verification. Download access and
	// client-validation only.
	Source string `json:"source,omitempty"`

	// DeviceAllowed: If restricted, whether access is granted for this
	// (user, device, volume).
	DeviceAllowed bool `json:"deviceAllowed,omitempty"`

	// Nonce: Client nonce for verification. Download access and
	// client-validation only.
	Nonce string `json:"nonce,omitempty"`

	// MaxDownloadDevices: If restricted, the maximum number of content
	// download licenses for this volume.
	MaxDownloadDevices int64 `json:"maxDownloadDevices,omitempty"`

	// Message: Error/warning message.
	Message string `json:"message,omitempty"`

	// ReasonCode: Error/warning reason code. Additional codes may be added
	// in the future. 0 OK 100 ACCESS_DENIED_PUBLISHER_LIMIT 101
	// ACCESS_DENIED_LIMIT 200 WARNING_USED_LAST_ACCESS
	ReasonCode string `json:"reasonCode,omitempty"`

	// Signature: Response signature.
	Signature string `json:"signature,omitempty"`
}

type ReviewSource struct {
	// ExtraDescription: Extra text about the source of the review.
	ExtraDescription string `json:"extraDescription,omitempty"`

	// Url: URL of the source of the review.
	Url string `json:"url,omitempty"`

	// Description: Name of the source.
	Description string `json:"description,omitempty"`
}

type Annotations struct {
	// Items: A list of annotations.
	Items []*Annotation `json:"items,omitempty"`

	// NextPageToken: Token to pass in for pagination for the next page.
	// This will not be present if this request does not have more results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalItems: Total number of annotations found. This may be greater
	// than the number of notes returned in this response if results have
	// been paginated.
	TotalItems int64 `json:"totalItems,omitempty"`

	// Kind: Resource type.
	Kind string `json:"kind,omitempty"`
}

type Review struct {
	// FullTextUrl: URL for the full review text, for reviews gathered from
	// the web.
	FullTextUrl string `json:"fullTextUrl,omitempty"`

	// Author: Author of this review.
	Author *ReviewAuthor `json:"author,omitempty"`

	// VolumeId: Volume that this review is for.
	VolumeId string `json:"volumeId,omitempty"`

	// Kind: Resource type for a review.
	Kind string `json:"kind,omitempty"`

	// Content: Review text.
	Content string `json:"content,omitempty"`

	// Source: Information regarding the source of this review, when the
	// review is not from a Google Books user.
	Source *ReviewSource `json:"source,omitempty"`

	// Date: Date of this review.
	Date string `json:"date,omitempty"`

	// Title: Title for this review.
	Title string `json:"title,omitempty"`

	// Rating: Star rating for this review. Possible values are ONE, TWO,
	// THREE, FOUR, FIVE or NOT_RATED.
	Rating string `json:"rating,omitempty"`

	// Type: Source type for this review. Possible values are EDITORIAL,
	// WEB_USER or GOOGLE_USER.
	Type string `json:"type,omitempty"`
}

type VolumeAccessInfo struct {
	// DownloadAccess: Information about a volume's download license access
	// restrictions.
	DownloadAccess *DownloadAccessRestriction `json:"downloadAccess,omitempty"`

	// AccessViewStatus: Combines the access and viewability of this volume
	// into a single status field for this user. Values can be
	// FULL_PURCHASED, FULL_PUBLIC_DOMAIN, SAMPLE or NONE. (In LITE
	// projection.)
	AccessViewStatus string `json:"accessViewStatus,omitempty"`

	// Embeddable: Whether this volume can be embedded in a viewport using
	// the Embedded Viewer API.
	Embeddable bool `json:"embeddable,omitempty"`

	// PublicDomain: Whether or not this book is public domain in the
	// country listed above.
	PublicDomain bool `json:"publicDomain,omitempty"`

	// Country: The two-letter ISO_3166-1 country code for which this access
	// information is valid. (In LITE projection.)
	Country string `json:"country,omitempty"`

	// Pdf: Information about pdf content. (In LITE projection.)
	Pdf *VolumeAccessInfoPdf `json:"pdf,omitempty"`

	// Viewability: The read access of a volume. Possible values are
	// PARTIAL, ALL_PAGES, NO_PAGES or UNKNOWN. This value depends on the
	// country listed above. A value of PARTIAL means that the publisher has
	// allowed some portion of the volume to be viewed publicly, without
	// purchase. This can apply to eBooks as well as non-eBooks. Public
	// domain books will always have a value of ALL_PAGES.
	Viewability string `json:"viewability,omitempty"`

	// Epub: Information about epub content. (In LITE projection.)
	Epub *VolumeAccessInfoEpub `json:"epub,omitempty"`

	// TextToSpeechPermission: Whether text-to-speech is permitted for this
	// volume. Values can be ALLOWED, ALLOWED_FOR_ACCESSIBILITY, or
	// NOT_ALLOWED.
	TextToSpeechPermission string `json:"textToSpeechPermission,omitempty"`
}

type BooksAnnotationsRange struct {
	// EndPosition: The ending position for the range.
	EndPosition string `json:"endPosition,omitempty"`

	// EndOffset: The offset from the ending position.
	EndOffset string `json:"endOffset,omitempty"`

	// StartPosition: The starting position for the range.
	StartPosition string `json:"startPosition,omitempty"`

	// StartOffset: The offset from the starting position.
	StartOffset string `json:"startOffset,omitempty"`
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

// OrderBy sets the optional parameter "orderBy": Sort search results.
func (c *VolumesListCall) OrderBy(orderBy string) *VolumesListCall {
	c.opt_["orderBy"] = orderBy
	return c
}

// Projection sets the optional parameter "projection": Restrict
// information returned to a set of selected fields.
func (c *VolumesListCall) Projection(projection string) *VolumesListCall {
	c.opt_["projection"] = projection
	return c
}

// StartIndex sets the optional parameter "startIndex": Index of the
// first result to return (starts at 0)
func (c *VolumesListCall) StartIndex(startIndex int64) *VolumesListCall {
	c.opt_["startIndex"] = startIndex
	return c
}

// Download sets the optional parameter "download": Restrict to volumes
// by download availability.
func (c *VolumesListCall) Download(download string) *VolumesListCall {
	c.opt_["download"] = download
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return.
func (c *VolumesListCall) MaxResults(maxResults int64) *VolumesListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// Filter sets the optional parameter "filter": Filter search results.
func (c *VolumesListCall) Filter(filter string) *VolumesListCall {
	c.opt_["filter"] = filter
	return c
}

// Country sets the optional parameter "country": ISO-3166-1 code to
// override the IP-based location.
func (c *VolumesListCall) Country(country string) *VolumesListCall {
	c.opt_["country"] = country
	return c
}

// LibraryRestrict sets the optional parameter "libraryRestrict":
// Restrict search to this user's library.
func (c *VolumesListCall) LibraryRestrict(libraryRestrict string) *VolumesListCall {
	c.opt_["libraryRestrict"] = libraryRestrict
	return c
}

// PrintType sets the optional parameter "printType": Restrict to books
// or magazines.
func (c *VolumesListCall) PrintType(printType string) *VolumesListCall {
	c.opt_["printType"] = printType
	return c
}

// Partner sets the optional parameter "partner": Identifier of partner
// for whom to restrict and brand results.
func (c *VolumesListCall) Partner(partner string) *VolumesListCall {
	c.opt_["partner"] = partner
	return c
}

// Source sets the optional parameter "source": String to identify the
// originator of this request.
func (c *VolumesListCall) Source(source string) *VolumesListCall {
	c.opt_["source"] = source
	return c
}

// ShowPreorders sets the optional parameter "showPreorders": Set to
// true to show books available for preorder. Defaults to false.
func (c *VolumesListCall) ShowPreorders(showPreorders bool) *VolumesListCall {
	c.opt_["showPreorders"] = showPreorders
	return c
}

// LangRestrict sets the optional parameter "langRestrict": Restrict
// results to books with this language code.
func (c *VolumesListCall) LangRestrict(langRestrict string) *VolumesListCall {
	c.opt_["langRestrict"] = langRestrict
	return c
}

func (c *VolumesListCall) Do() (*Volumes, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("q", fmt.Sprintf("%v", c.q))
	if v, ok := c.opt_["orderBy"]; ok {
		params.Set("orderBy", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["projection"]; ok {
		params.Set("projection", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["startIndex"]; ok {
		params.Set("startIndex", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["download"]; ok {
		params.Set("download", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["filter"]; ok {
		params.Set("filter", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["country"]; ok {
		params.Set("country", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["libraryRestrict"]; ok {
		params.Set("libraryRestrict", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["printType"]; ok {
		params.Set("printType", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["partner"]; ok {
		params.Set("partner", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["source"]; ok {
		params.Set("source", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["showPreorders"]; ok {
		params.Set("showPreorders", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["langRestrict"]; ok {
		params.Set("langRestrict", fmt.Sprintf("%v", v))
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
	//     "country": {
	//       "description": "ISO-3166-1 code to override the IP-based location.",
	//       "location": "query",
	//       "type": "string"
	//     },
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
	//       "description": "Identifier of partner for whom to restrict and brand results.",
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

// Projection sets the optional parameter "projection": Restrict
// information returned to a set of selected fields.
func (c *VolumesGetCall) Projection(projection string) *VolumesGetCall {
	c.opt_["projection"] = projection
	return c
}

// Country sets the optional parameter "country": ISO-3166-1 code to
// override the IP-based location.
func (c *VolumesGetCall) Country(country string) *VolumesGetCall {
	c.opt_["country"] = country
	return c
}

// Partner sets the optional parameter "partner": Identifier of partner
// for whom to brand results.
func (c *VolumesGetCall) Partner(partner string) *VolumesGetCall {
	c.opt_["partner"] = partner
	return c
}

// Source sets the optional parameter "source": String to identify the
// originator of this request.
func (c *VolumesGetCall) Source(source string) *VolumesGetCall {
	c.opt_["source"] = source
	return c
}

func (c *VolumesGetCall) Do() (*Volume, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["projection"]; ok {
		params.Set("projection", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["country"]; ok {
		params.Set("country", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["partner"]; ok {
		params.Set("partner", fmt.Sprintf("%v", v))
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
	//       "description": "Identifier of partner for whom to brand results.",
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
	//       "description": "Id of volume to retrieve.",
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

// Country sets the optional parameter "country": ISO-3166-1 code to
// override the IP-based location.
func (c *BookshelvesListCall) Country(country string) *BookshelvesListCall {
	c.opt_["country"] = country
	return c
}

// Source sets the optional parameter "source": String to identify the
// originator of this request.
func (c *BookshelvesListCall) Source(source string) *BookshelvesListCall {
	c.opt_["source"] = source
	return c
}

func (c *BookshelvesListCall) Do() (*Bookshelves, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["country"]; ok {
		params.Set("country", fmt.Sprintf("%v", v))
	}
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
	//     "country": {
	//       "description": "ISO-3166-1 code to override the IP-based location.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "source": {
	//       "description": "String to identify the originator of this request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "Id of user for whom to retrieve bookshelves.",
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

// method id "books.bookshelves.get":

type BookshelvesGetCall struct {
	s      *Service
	userId string
	shelf  string
	opt_   map[string]interface{}
}

// Get: Retrieves a specific bookshelf for the specified user.
func (r *BookshelvesService) Get(userId string, shelf string) *BookshelvesGetCall {
	c := &BookshelvesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.shelf = shelf
	return c
}

// Country sets the optional parameter "country": ISO-3166-1 code to
// override the IP-based location.
func (c *BookshelvesGetCall) Country(country string) *BookshelvesGetCall {
	c.opt_["country"] = country
	return c
}

// Source sets the optional parameter "source": String to identify the
// originator of this request.
func (c *BookshelvesGetCall) Source(source string) *BookshelvesGetCall {
	c.opt_["source"] = source
	return c
}

func (c *BookshelvesGetCall) Do() (*Bookshelf, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["country"]; ok {
		params.Set("country", fmt.Sprintf("%v", v))
	}
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
	//   "description": "Retrieves a specific bookshelf for the specified user.",
	//   "httpMethod": "GET",
	//   "id": "books.bookshelves.get",
	//   "parameterOrder": [
	//     "userId",
	//     "shelf"
	//   ],
	//   "parameters": {
	//     "country": {
	//       "description": "ISO-3166-1 code to override the IP-based location.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "shelf": {
	//       "description": "Id of bookshelf to retrieve.",
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
	//       "description": "Id of user for whom to retrieve bookshelves.",
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

func cleanPathString(s string) string {
	return strings.Map(func(r int) int {
		if r >= 0x30 && r <= 0x7a {
			return r
		}
		return -1
	}, s)
}
