// Package shopping provides access to the Search API for Shopping.
//
// See http://code.google.com/apis/shopping/search/v1/getting_started.html
//
// Usage example:
//
//   import "google-api-go-client.googlecode.com/hg/shopping/v1"
//   ...
//   shoppingService, err := shopping.New(oauthHttpClient)
package shopping

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

const apiId = "shopping:v1"
const apiName = "shopping"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/shopping/search/v1/"

func New(client *http.Client) (*Service, os.Error) {
	if client == nil {
		return nil, os.NewError("client is nil")
	}
	s := &Service{client: client}
	s.Products = &ProductsService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	Products *ProductsService
}

type ProductsService struct {
	s *Service
}

type ShoppingModelProductJsonV1ImagesThumbnails struct {
	// Height: Height of thumbnail (omitted if not specified in the
	// request).
	Height int64 `json:"height,omitempty"`

	// Width: Width of thumbnail (omitted if not specified in the request).
	Width int64 `json:"width,omitempty"`

	// Link: Link to thumbnail.
	Link string `json:"link,omitempty"`

	// Content: Content of thumbnail (only available for the first thumbnail
	// of the top results if SAYT is enabled).
	Content string `json:"content,omitempty"`
}

type ProductsFacets struct {
	// Count: Number of products matching the query that have a value for
	// the facet's property or attribute.
	Count int64 `json:"count,omitempty"`

	// Name: Name of the facet's attribute (omitted for property facets).
	Name string `json:"name,omitempty"`

	// DisplayName: Display name of facet.
	DisplayName string `json:"displayName,omitempty"`

	// Property: Property of facet (omitted for attribute facets).
	Property string `json:"property,omitempty"`

	// Unit: Unit of the facet's property or attribute (omitted if the
	// facet's property or attribute has no unit).
	Unit string `json:"unit,omitempty"`

	// Buckets: List of Buckets within facet.
	Buckets []*ProductsFacetsBuckets `json:"buckets,omitempty"`

	// Type: Type of facet's attribute (omitted for property facets, one of:
	// text, bool, int, float).
	Type string `json:"type,omitempty"`
}

type ShoppingModelDebugJsonV1BackendTimes struct {
	// Name: Google internal
	Name string `json:"name,omitempty"`

	// ServerMillis: Google internal
	ServerMillis int64 `json:"serverMillis,omitempty,string"`

	// ElapsedMillis: Google internal
	ElapsedMillis int64 `json:"elapsedMillis,omitempty,string"`

	// HostName: Google internal
	HostName string `json:"hostName,omitempty"`
}

type Products struct {
	// Categories: List of categories.
	Categories []*ShoppingModelCategoryJsonV1 `json:"categories,omitempty"`

	// StartIndex: 1-based index of the first item in the search results.
	StartIndex int64 `json:"startIndex,omitempty"`

	// SelfLink: Self link of feed.
	SelfLink string `json:"selfLink,omitempty"`

	// NextLink: Next link of feed.
	NextLink string `json:"nextLink,omitempty"`

	// Items: List of returned products.
	Items []*Product `json:"items,omitempty"`

	// TotalItems: Total number of search results.
	TotalItems int64 `json:"totalItems,omitempty"`

	// Etag: Etag of feed.
	Etag string `json:"etag,omitempty"`

	// Kind: The fixed string "shopping#products". The kind of feed
	// returned.
	Kind string `json:"kind,omitempty"`

	// Stores: List of returned stores.
	Stores []*ProductsStores `json:"stores,omitempty"`

	// Id: Id of feed.
	Id string `json:"id,omitempty"`

	// CurrentItemCount: Current item count.
	CurrentItemCount int64 `json:"currentItemCount,omitempty"`

	// Facets: List of facets.
	Facets []*ProductsFacets `json:"facets,omitempty"`

	ShelfSpaceAds []*ProductsShelfSpaceAds `json:"shelfSpaceAds,omitempty"`

	// RelatedQueries: Related queries.
	RelatedQueries []string `json:"relatedQueries,omitempty"`

	// Debug: Google internal.
	Debug *ShoppingModelDebugJsonV1 `json:"debug,omitempty"`

	// Promotions: List of promotions.
	Promotions []*ProductsPromotions `json:"promotions,omitempty"`

	// Spelling: Spelling.
	Spelling *ProductsSpelling `json:"spelling,omitempty"`

	// Redirects: Redirects.
	Redirects []string `json:"redirects,omitempty"`

	// ItemsPerPage: Number of items per page of results.
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`

	// PreviousLink: Previous link of feed.
	PreviousLink string `json:"previousLink,omitempty"`
}

type ShoppingModelProductJsonV1Images struct {
	// Thumbnails: Thumbnails of product image.
	Thumbnails []*ShoppingModelProductJsonV1ImagesThumbnails `json:"thumbnails,omitempty"`

	// Link: Link to product image.
	Link string `json:"link,omitempty"`
}

type ShoppingModelProductJsonV1Attributes struct {
	// Value: Value of product attribute.
	Value interface{} `json:"value,omitempty"`

	// Name: Name of product attribute.
	Name string `json:"name,omitempty"`

	// DisplayName: Display Name of prodct attribute.
	DisplayName string `json:"displayName,omitempty"`

	// Unit: Unit of product attribute.
	Unit string `json:"unit,omitempty"`

	// Type: Type of product attribute (one of: text, bool, int, float,
	// dateRange, url).
	Type string `json:"type,omitempty"`
}

type ProductsSpelling struct {
	// Suggestion: Suggestion for spelling.
	Suggestion string `json:"suggestion,omitempty"`
}

type ShoppingModelProductJsonV1 struct {
	// Categories: Categories of product according to the selected taxonomy,
	// omitted if no taxonomy is selected.
	Categories []string `json:"categories,omitempty"`

	// Author: Author of product.
	Author *ShoppingModelProductJsonV1Author `json:"author,omitempty"`

	// Gtin: The first GTIN of the product. Deprecated in favor of "gtins".
	Gtin string `json:"gtin,omitempty"`

	// Language: BCP 47 language tag of language of product.
	Language string `json:"language,omitempty"`

	// Gtins: List of all the product's GTINs (in GTIN-14 format).
	Gtins []string `json:"gtins,omitempty"`

	// ModificationTime: RFC 3339 formatted modification time and date of
	// product.
	ModificationTime string `json:"modificationTime,omitempty"`

	// CreationTime: RFC 3339 formatted creation time and date of product.
	CreationTime string `json:"creationTime,omitempty"`

	// Condition: Condition of product (one of: new, refurbished, used).
	Condition string `json:"condition,omitempty"`

	// GoogleId: Google id of product.
	GoogleId string `json:"googleId,omitempty"`

	// Country: ISO 3166 code of target country of product.
	Country string `json:"country,omitempty"`

	// Attributes: Attributes of product (available only with a cx source).
	Attributes []*ShoppingModelProductJsonV1Attributes `json:"attributes,omitempty"`

	// Link: Link to product.
	Link string `json:"link,omitempty"`

	// Inventories: Inventories of product.
	Inventories []*ShoppingModelProductJsonV1Inventories `json:"inventories,omitempty"`

	// ProvidedId: Merchant-provided id of product (available only with a cx
	// source).
	ProvidedId string `json:"providedId,omitempty"`

	// Title: Title of product.
	Title string `json:"title,omitempty"`

	// PlusOne: Code to add to the page to render the +1 content.
	PlusOne string `json:"plusOne,omitempty"`

	// Images: Images of product.
	Images []*ShoppingModelProductJsonV1Images `json:"images,omitempty"`

	// Brand: Brand of product.
	Brand string `json:"brand,omitempty"`

	// Description: Description of product.
	Description string `json:"description,omitempty"`
}

type ProductRecommendations struct {
	// RecommendationList: List of recommendations.
	RecommendationList []*ProductRecommendationsRecommendationList `json:"recommendationList,omitempty"`

	// Type: Type of recommendation list (one of: all, purchaseToPurchase,
	// visitToVisit, visitToPurchase).
	Type string `json:"type,omitempty"`
}

type ProductsFacetsBuckets struct {
	// Count: Number of products matching the query that have a value for
	// the facet's property or attribute that matches the bucket.
	Count int64 `json:"count,omitempty"`

	// Value: Value of the bucket (omitted for range buckets).
	Value interface{} `json:"value,omitempty"`

	// Min: Lower bound of the bucket (omitted for value buckets or if the
	// range has no lower bound).
	Min interface{} `json:"min,omitempty"`

	// MinExclusive: Whether the lower bound of the bucket is exclusive
	// (omitted for value buckets or if the range has no lower bound).
	MinExclusive bool `json:"minExclusive,omitempty"`

	// Max: Upper bound of the bucket (omitted for value buckets or if the
	// range has no upper bound).
	Max interface{} `json:"max,omitempty"`

	// MaxExclusive: Whether the upper bound of the bucket is exclusive
	// (omitted for value buckets or if the range has no upper bound).
	MaxExclusive bool `json:"maxExclusive,omitempty"`
}

type ShoppingModelCategoryJsonV1 struct {
	// Url: URL of category.
	Url string `json:"url,omitempty"`

	// Id: Id of category.
	Id string `json:"id,omitempty"`

	// Parents: Ids of the parents of the category.
	Parents []string `json:"parents,omitempty"`

	// ShortName: Short name of category.
	ShortName string `json:"shortName,omitempty"`
}

type ProductsStores struct {
	// Location: Location of store.
	Location string `json:"location,omitempty"`

	// StoreCode: Merchant-supplied store code.
	StoreCode string `json:"storeCode,omitempty"`

	// Name: Name of store.
	Name string `json:"name,omitempty"`

	// StoreId: Id of store.
	StoreId string `json:"storeId,omitempty"`

	// Address: Address of store.
	Address string `json:"address,omitempty"`

	// Telephone: Telephone number of store.
	Telephone string `json:"telephone,omitempty"`
}

type ProductsPromotions struct {
	// DestLink: Link to promotion (omitted if type is not standard).
	DestLink string `json:"destLink,omitempty"`

	// Name: Name of promotion (omitted if type is not standard).
	Name string `json:"name,omitempty"`

	// Product: Product of promotion (omitted if type is not product).
	Product *ShoppingModelProductJsonV1 `json:"product,omitempty"`

	// Link: Link to promotion without scheme. DEPRECATED. WILL BE REMOVED
	// SOON. USE destLink.
	Link string `json:"link,omitempty"`

	// CustomFields: List of custom fields of promotion.
	CustomFields []*ProductsPromotionsCustomFields `json:"customFields,omitempty"`

	// ImageLink: Link to promotion image (omitted if type is not standard).
	ImageLink string `json:"imageLink,omitempty"`

	// CustomHtml: Custom HTML of promotion (omitted if type is not custom).
	CustomHtml string `json:"customHtml,omitempty"`

	// Type: Type of promotion (one of: standard, product, custom).
	Type string `json:"type,omitempty"`

	// Description: Description of promotion (omitted if type is not
	// standard).
	Description string `json:"description,omitempty"`
}

type ShoppingModelProductJsonV1Inventories struct {
	// Currency: Currency of product inventory (an ISO 4217 alphabetic
	// code).
	Currency string `json:"currency,omitempty"`

	// Distance: Distance of product inventory.
	Distance float64 `json:"distance,omitempty"`

	// Price: Price of product inventory.
	Price float64 `json:"price,omitempty"`

	// Shipping: Shipping cost of product inventory.
	Shipping float64 `json:"shipping,omitempty"`

	// DistanceUnit: Distance unit of product inventory.
	DistanceUnit string `json:"distanceUnit,omitempty"`

	// StoreId: Store ID of product inventory.
	StoreId string `json:"storeId,omitempty"`

	// Availability: Availability of product inventory.
	Availability string `json:"availability,omitempty"`

	// Channel: Channel of product inventory (one of: online, local).
	Channel string `json:"channel,omitempty"`

	// Tax: Tax of product inventory.
	Tax float64 `json:"tax,omitempty"`
}

type ProductsShelfSpaceAds struct {
	Product *ShoppingModelProductJsonV1 `json:"product,omitempty"`
}

type Product struct {
	// Categories: List of categories for product.
	Categories []*ShoppingModelCategoryJsonV1 `json:"categories,omitempty"`

	// SelfLink: Self link of product.
	SelfLink string `json:"selfLink,omitempty"`

	// Product: Product.
	Product *ShoppingModelProductJsonV1 `json:"product,omitempty"`

	// Kind: The kind of item, always shopping#product.
	Kind string `json:"kind,omitempty"`

	// Id: Id of product.
	Id string `json:"id,omitempty"`

	// Debug: Google internal.
	Debug *ShoppingModelDebugJsonV1 `json:"debug,omitempty"`

	// Recommendations: Recommendations for product.
	Recommendations []*ProductRecommendations `json:"recommendations,omitempty"`
}

type ShoppingModelDebugJsonV1 struct {
	// SearchResponse: Google internal.
	SearchResponse string `json:"searchResponse,omitempty"`

	// BackendTimes: Google internal
	BackendTimes []*ShoppingModelDebugJsonV1BackendTimes `json:"backendTimes,omitempty"`

	// RdcResponse: Google internal.
	RdcResponse string `json:"rdcResponse,omitempty"`

	// ElapsedMillis: Google internal.
	ElapsedMillis int64 `json:"elapsedMillis,omitempty,string"`

	// SearchRequest: Google internal.
	SearchRequest string `json:"searchRequest,omitempty"`
}

type ProductsPromotionsCustomFields struct {
	// Value: Value of field.
	Value string `json:"value,omitempty"`

	// Name: Name of field.
	Name string `json:"name,omitempty"`
}

type ShoppingModelProductJsonV1Author struct {
	// Email: Email address of product author.
	Email string `json:"email,omitempty"`

	// Name: Name of product author.
	Name string `json:"name,omitempty"`

	// AccountId: Account id of product author.
	AccountId string `json:"accountId,omitempty"`

	// AggregatorId: Aggregator id of product author.
	AggregatorId string `json:"aggregatorId,omitempty"`

	// Uri: URI of product author.
	Uri string `json:"uri,omitempty"`
}

type ProductRecommendationsRecommendationList struct {
	// Product: Recommended product.
	Product *ShoppingModelProductJsonV1 `json:"product,omitempty"`
}

// method id "shopping.products.list":

type ProductsListCall struct {
	s      *Service
	source string
	opt_   map[string]interface{}
}

// List: Returns a list of products and content modules
func (r *ProductsService) List(source string) *ProductsListCall {
	c := &ProductsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.source = source
	return c
}

// Thumbnails sets the optional parameter "thumbnails": Image thumbnails
// specification
func (c *ProductsListCall) Thumbnails(thumbnails string) *ProductsListCall {
	c.opt_["thumbnails"] = thumbnails
	return c
}

// Channels sets the optional parameter "channels": Channels
// specification
func (c *ProductsListCall) Channels(channels string) *ProductsListCall {
	c.opt_["channels"] = channels
	return c
}

// Taxonomy sets the optional parameter "taxonomy": Taxonomy name
func (c *ProductsListCall) Taxonomy(taxonomy string) *ProductsListCall {
	c.opt_["taxonomy"] = taxonomy
	return c
}

// StartIndex sets the optional parameter "startIndex": Index (1-based)
// of first product to return
func (c *ProductsListCall) StartIndex(startIndex int64) *ProductsListCall {
	c.opt_["startIndex"] = startIndex
	return c
}

// Q sets the optional parameter "q": Search query
func (c *ProductsListCall) Q(q string) *ProductsListCall {
	c.opt_["q"] = q
	return c
}

// CategoriesUseGcsConfig sets the optional parameter
// "categories.useGcsConfig": This parameter is currently ignored
func (c *ProductsListCall) CategoriesUseGcsConfig(categoriesUseGcsConfig bool) *ProductsListCall {
	c.opt_["categories.useGcsConfig"] = categoriesUseGcsConfig
	return c
}

// Currency sets the optional parameter "currency": Currency restriction
// (ISO 4217)
func (c *ProductsListCall) Currency(currency string) *ProductsListCall {
	c.opt_["currency"] = currency
	return c
}

// ProductFields sets the optional parameter "productFields": Google
// Internal
func (c *ProductsListCall) ProductFields(productFields string) *ProductsListCall {
	c.opt_["productFields"] = productFields
	return c
}

// Language sets the optional parameter "language": Language restriction
// (BCP 47)
func (c *ProductsListCall) Language(language string) *ProductsListCall {
	c.opt_["language"] = language
	return c
}

// DebugEnableLogging sets the optional parameter "debug.enableLogging":
// Google Internal
func (c *ProductsListCall) DebugEnableLogging(debugEnableLogging bool) *ProductsListCall {
	c.opt_["debug.enableLogging"] = debugEnableLogging
	return c
}

// DebugGeocodeRequest sets the optional parameter
// "debug.geocodeRequest": Google Internal
func (c *ProductsListCall) DebugGeocodeRequest(debugGeocodeRequest bool) *ProductsListCall {
	c.opt_["debug.geocodeRequest"] = debugGeocodeRequest
	return c
}

// DebugSearchRequest sets the optional parameter "debug.searchRequest":
// Google Internal
func (c *ProductsListCall) DebugSearchRequest(debugSearchRequest bool) *ProductsListCall {
	c.opt_["debug.searchRequest"] = debugSearchRequest
	return c
}

// FacetsUseGcsConfig sets the optional parameter "facets.useGcsConfig":
// Whether to return facet information as configured in the GCS account
func (c *ProductsListCall) FacetsUseGcsConfig(facetsUseGcsConfig bool) *ProductsListCall {
	c.opt_["facets.useGcsConfig"] = facetsUseGcsConfig
	return c
}

// Location sets the optional parameter "location": Location used to
// determine tax and shipping
func (c *ProductsListCall) Location(location string) *ProductsListCall {
	c.opt_["location"] = location
	return c
}

// RelatedQueriesUseGcsConfig sets the optional parameter
// "relatedQueries.useGcsConfig": This parameter is currently ignored
func (c *ProductsListCall) RelatedQueriesUseGcsConfig(relatedQueriesUseGcsConfig bool) *ProductsListCall {
	c.opt_["relatedQueries.useGcsConfig"] = relatedQueriesUseGcsConfig
	return c
}

// AttributeFilter sets the optional parameter "attributeFilter": Comma
// separated list of attributes to return
func (c *ProductsListCall) AttributeFilter(attributeFilter string) *ProductsListCall {
	c.opt_["attributeFilter"] = attributeFilter
	return c
}

// PromotionsEnabled sets the optional parameter "promotions.enabled":
// Whether to return promotion information
func (c *ProductsListCall) PromotionsEnabled(promotionsEnabled bool) *ProductsListCall {
	c.opt_["promotions.enabled"] = promotionsEnabled
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return
func (c *ProductsListCall) MaxResults(maxResults int64) *ProductsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// ShelfSpaceAdsUseGcsConfig sets the optional parameter
// "shelfSpaceAds.useGcsConfig": This parameter is currently ignored
func (c *ProductsListCall) ShelfSpaceAdsUseGcsConfig(shelfSpaceAdsUseGcsConfig bool) *ProductsListCall {
	c.opt_["shelfSpaceAds.useGcsConfig"] = shelfSpaceAdsUseGcsConfig
	return c
}

// UseCase sets the optional parameter "useCase": One of
// CommerceSearchUseCase, ShoppingApiUseCase
func (c *ProductsListCall) UseCase(useCase string) *ProductsListCall {
	c.opt_["useCase"] = useCase
	return c
}

// FacetsInclude sets the optional parameter "facets.include": Facets to
// include (applies when useGcsConfig == false)
func (c *ProductsListCall) FacetsInclude(facetsInclude string) *ProductsListCall {
	c.opt_["facets.include"] = facetsInclude
	return c
}

// SaytUseGcsConfig sets the optional parameter "sayt.useGcsConfig":
// Google Internal
func (c *ProductsListCall) SaytUseGcsConfig(saytUseGcsConfig bool) *ProductsListCall {
	c.opt_["sayt.useGcsConfig"] = saytUseGcsConfig
	return c
}

// ShelfSpaceAdsMaxResults sets the optional parameter
// "shelfSpaceAds.maxResults": The maximum number of shelf space ads to
// return
func (c *ProductsListCall) ShelfSpaceAdsMaxResults(shelfSpaceAdsMaxResults int64) *ProductsListCall {
	c.opt_["shelfSpaceAds.maxResults"] = shelfSpaceAdsMaxResults
	return c
}

// CrowdBy sets the optional parameter "crowdBy": Crowding specification
func (c *ProductsListCall) CrowdBy(crowdBy string) *ProductsListCall {
	c.opt_["crowdBy"] = crowdBy
	return c
}

// Safe sets the optional parameter "safe": Whether safe search is
// enabled. Default: true
func (c *ProductsListCall) Safe(safe bool) *ProductsListCall {
	c.opt_["safe"] = safe
	return c
}

// CategoriesEnabled sets the optional parameter "categories.enabled":
// Whether to return category information
func (c *ProductsListCall) CategoriesEnabled(categoriesEnabled bool) *ProductsListCall {
	c.opt_["categories.enabled"] = categoriesEnabled
	return c
}

// Country sets the optional parameter "country": Country restriction
// (ISO 3166)
func (c *ProductsListCall) Country(country string) *ProductsListCall {
	c.opt_["country"] = country
	return c
}

// FacetsEnabled sets the optional parameter "facets.enabled": Whether
// to return facet information
func (c *ProductsListCall) FacetsEnabled(facetsEnabled bool) *ProductsListCall {
	c.opt_["facets.enabled"] = facetsEnabled
	return c
}

// DebugEnabled sets the optional parameter "debug.enabled": Google
// Internal
func (c *ProductsListCall) DebugEnabled(debugEnabled bool) *ProductsListCall {
	c.opt_["debug.enabled"] = debugEnabled
	return c
}

// BoostBy sets the optional parameter "boostBy": Boosting specification
func (c *ProductsListCall) BoostBy(boostBy string) *ProductsListCall {
	c.opt_["boostBy"] = boostBy
	return c
}

// SpellingEnabled sets the optional parameter "spelling.enabled":
// Whether to return spelling suggestions
func (c *ProductsListCall) SpellingEnabled(spellingEnabled bool) *ProductsListCall {
	c.opt_["spelling.enabled"] = spellingEnabled
	return c
}

// RankBy sets the optional parameter "rankBy": Ranking specification
func (c *ProductsListCall) RankBy(rankBy string) *ProductsListCall {
	c.opt_["rankBy"] = rankBy
	return c
}

// CategoriesInclude sets the optional parameter "categories.include":
// Category specification
func (c *ProductsListCall) CategoriesInclude(categoriesInclude string) *ProductsListCall {
	c.opt_["categories.include"] = categoriesInclude
	return c
}

// DebugGeocodeResponse sets the optional parameter
// "debug.geocodeResponse": Google Internal
func (c *ProductsListCall) DebugGeocodeResponse(debugGeocodeResponse bool) *ProductsListCall {
	c.opt_["debug.geocodeResponse"] = debugGeocodeResponse
	return c
}

// SpellingUseGcsConfig sets the optional parameter
// "spelling.useGcsConfig": This parameter is currently ignored
func (c *ProductsListCall) SpellingUseGcsConfig(spellingUseGcsConfig bool) *ProductsListCall {
	c.opt_["spelling.useGcsConfig"] = spellingUseGcsConfig
	return c
}

// PlusOne sets the optional parameter "plusOne": +1 rendering
// specification.
func (c *ProductsListCall) PlusOne(plusOne string) *ProductsListCall {
	c.opt_["plusOne"] = plusOne
	return c
}

// RelatedQueriesEnabled sets the optional parameter
// "relatedQueries.enabled": Whether to return related queries
func (c *ProductsListCall) RelatedQueriesEnabled(relatedQueriesEnabled bool) *ProductsListCall {
	c.opt_["relatedQueries.enabled"] = relatedQueriesEnabled
	return c
}

// DebugSearchResponse sets the optional parameter
// "debug.searchResponse": Google Internal
func (c *ProductsListCall) DebugSearchResponse(debugSearchResponse bool) *ProductsListCall {
	c.opt_["debug.searchResponse"] = debugSearchResponse
	return c
}

// MinAvailability sets the optional parameter "minAvailability": 
func (c *ProductsListCall) MinAvailability(minAvailability string) *ProductsListCall {
	c.opt_["minAvailability"] = minAvailability
	return c
}

// FacetsDiscover sets the optional parameter "facets.discover": Facets
// to discover
func (c *ProductsListCall) FacetsDiscover(facetsDiscover string) *ProductsListCall {
	c.opt_["facets.discover"] = facetsDiscover
	return c
}

// RestrictBy sets the optional parameter "restrictBy": Restriction
// specification
func (c *ProductsListCall) RestrictBy(restrictBy string) *ProductsListCall {
	c.opt_["restrictBy"] = restrictBy
	return c
}

// DebugRdcRequest sets the optional parameter "debug.rdcRequest":
// Google Internal
func (c *ProductsListCall) DebugRdcRequest(debugRdcRequest bool) *ProductsListCall {
	c.opt_["debug.rdcRequest"] = debugRdcRequest
	return c
}

// PromotionsUseGcsConfig sets the optional parameter
// "promotions.useGcsConfig": Whether to return promotion information as
// configured in the GCS account
func (c *ProductsListCall) PromotionsUseGcsConfig(promotionsUseGcsConfig bool) *ProductsListCall {
	c.opt_["promotions.useGcsConfig"] = promotionsUseGcsConfig
	return c
}

// RedirectsUseGcsConfig sets the optional parameter
// "redirects.useGcsConfig": Whether to return redirect information as
// configured in the GCS account
func (c *ProductsListCall) RedirectsUseGcsConfig(redirectsUseGcsConfig bool) *ProductsListCall {
	c.opt_["redirects.useGcsConfig"] = redirectsUseGcsConfig
	return c
}

// DebugRdcResponse sets the optional parameter "debug.rdcResponse":
// Google Internal
func (c *ProductsListCall) DebugRdcResponse(debugRdcResponse bool) *ProductsListCall {
	c.opt_["debug.rdcResponse"] = debugRdcResponse
	return c
}

// SaytEnabled sets the optional parameter "sayt.enabled": Google
// Internal
func (c *ProductsListCall) SaytEnabled(saytEnabled bool) *ProductsListCall {
	c.opt_["sayt.enabled"] = saytEnabled
	return c
}

// RedirectsEnabled sets the optional parameter "redirects.enabled":
// Whether to return redirect information
func (c *ProductsListCall) RedirectsEnabled(redirectsEnabled bool) *ProductsListCall {
	c.opt_["redirects.enabled"] = redirectsEnabled
	return c
}

// ShelfSpaceAdsEnabled sets the optional parameter
// "shelfSpaceAds.enabled": Whether to return shelf space ads
func (c *ProductsListCall) ShelfSpaceAdsEnabled(shelfSpaceAdsEnabled bool) *ProductsListCall {
	c.opt_["shelfSpaceAds.enabled"] = shelfSpaceAdsEnabled
	return c
}

func (c *ProductsListCall) Do() (*Products, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["thumbnails"]; ok {
		params.Set("thumbnails", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["channels"]; ok {
		params.Set("channels", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["taxonomy"]; ok {
		params.Set("taxonomy", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["startIndex"]; ok {
		params.Set("startIndex", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["q"]; ok {
		params.Set("q", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["categories.useGcsConfig"]; ok {
		params.Set("categories.useGcsConfig", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["currency"]; ok {
		params.Set("currency", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["productFields"]; ok {
		params.Set("productFields", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["language"]; ok {
		params.Set("language", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["debug.enableLogging"]; ok {
		params.Set("debug.enableLogging", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["debug.geocodeRequest"]; ok {
		params.Set("debug.geocodeRequest", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["debug.searchRequest"]; ok {
		params.Set("debug.searchRequest", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["facets.useGcsConfig"]; ok {
		params.Set("facets.useGcsConfig", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["location"]; ok {
		params.Set("location", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["relatedQueries.useGcsConfig"]; ok {
		params.Set("relatedQueries.useGcsConfig", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["attributeFilter"]; ok {
		params.Set("attributeFilter", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["promotions.enabled"]; ok {
		params.Set("promotions.enabled", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["shelfSpaceAds.useGcsConfig"]; ok {
		params.Set("shelfSpaceAds.useGcsConfig", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["useCase"]; ok {
		params.Set("useCase", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["facets.include"]; ok {
		params.Set("facets.include", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["sayt.useGcsConfig"]; ok {
		params.Set("sayt.useGcsConfig", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["shelfSpaceAds.maxResults"]; ok {
		params.Set("shelfSpaceAds.maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["crowdBy"]; ok {
		params.Set("crowdBy", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["safe"]; ok {
		params.Set("safe", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["categories.enabled"]; ok {
		params.Set("categories.enabled", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["country"]; ok {
		params.Set("country", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["facets.enabled"]; ok {
		params.Set("facets.enabled", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["debug.enabled"]; ok {
		params.Set("debug.enabled", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["boostBy"]; ok {
		params.Set("boostBy", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["spelling.enabled"]; ok {
		params.Set("spelling.enabled", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["rankBy"]; ok {
		params.Set("rankBy", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["categories.include"]; ok {
		params.Set("categories.include", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["debug.geocodeResponse"]; ok {
		params.Set("debug.geocodeResponse", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["spelling.useGcsConfig"]; ok {
		params.Set("spelling.useGcsConfig", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["plusOne"]; ok {
		params.Set("plusOne", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["relatedQueries.enabled"]; ok {
		params.Set("relatedQueries.enabled", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["debug.searchResponse"]; ok {
		params.Set("debug.searchResponse", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["minAvailability"]; ok {
		params.Set("minAvailability", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["facets.discover"]; ok {
		params.Set("facets.discover", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["restrictBy"]; ok {
		params.Set("restrictBy", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["debug.rdcRequest"]; ok {
		params.Set("debug.rdcRequest", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["promotions.useGcsConfig"]; ok {
		params.Set("promotions.useGcsConfig", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["redirects.useGcsConfig"]; ok {
		params.Set("redirects.useGcsConfig", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["debug.rdcResponse"]; ok {
		params.Set("debug.rdcResponse", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["sayt.enabled"]; ok {
		params.Set("sayt.enabled", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["redirects.enabled"]; ok {
		params.Set("redirects.enabled", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["shelfSpaceAds.enabled"]; ok {
		params.Set("shelfSpaceAds.enabled", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/shopping/search/v1/", "{source}/products")
	urls = strings.Replace(urls, "{source}", cleanPathString(c.source), 1)
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
	ret := new(Products)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns a list of products and content modules",
	//   "httpMethod": "GET",
	//   "id": "shopping.products.list",
	//   "parameterOrder": [
	//     "source"
	//   ],
	//   "parameters": {
	//     "attributeFilter": {
	//       "description": "Comma separated list of attributes to return",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "boostBy": {
	//       "description": "Boosting specification",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "categories.enabled": {
	//       "description": "Whether to return category information",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "categories.include": {
	//       "description": "Category specification",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "categories.useGcsConfig": {
	//       "description": "This parameter is currently ignored",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "channels": {
	//       "description": "Channels specification",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "country": {
	//       "description": "Country restriction (ISO 3166)",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "crowdBy": {
	//       "description": "Crowding specification",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "currency": {
	//       "description": "Currency restriction (ISO 4217)",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "debug.enableLogging": {
	//       "description": "Google Internal",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "debug.enabled": {
	//       "description": "Google Internal",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "debug.geocodeRequest": {
	//       "description": "Google Internal",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "debug.geocodeResponse": {
	//       "description": "Google Internal",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "debug.rdcRequest": {
	//       "description": "Google Internal",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "debug.rdcResponse": {
	//       "description": "Google Internal",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "debug.searchRequest": {
	//       "description": "Google Internal",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "debug.searchResponse": {
	//       "description": "Google Internal",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "facets.discover": {
	//       "description": "Facets to discover",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "facets.enabled": {
	//       "description": "Whether to return facet information",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "facets.include": {
	//       "description": "Facets to include (applies when useGcsConfig == false)",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "facets.useGcsConfig": {
	//       "description": "Whether to return facet information as configured in the GCS account",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "language": {
	//       "description": "Language restriction (BCP 47)",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "location": {
	//       "description": "Location used to determine tax and shipping",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "minAvailability": {
	//       "enum": [
	//         "inStock",
	//         "limited",
	//         "outOfStock",
	//         "unknown"
	//       ],
	//       "enumDescriptions": [
	//         "Only in stcok prodicts will be returned",
	//         "Limited availability and in stock products will be returned",
	//         "Out of stock, limited availability and in stock products will be returned",
	//         "All products will be returned"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "plusOne": {
	//       "description": "+1 rendering specification.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "productFields": {
	//       "description": "Google Internal",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "promotions.enabled": {
	//       "description": "Whether to return promotion information",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "promotions.useGcsConfig": {
	//       "description": "Whether to return promotion information as configured in the GCS account",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "q": {
	//       "description": "Search query",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "rankBy": {
	//       "description": "Ranking specification",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "redirects.enabled": {
	//       "description": "Whether to return redirect information",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "redirects.useGcsConfig": {
	//       "description": "Whether to return redirect information as configured in the GCS account",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "relatedQueries.enabled": {
	//       "description": "Whether to return related queries",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "relatedQueries.useGcsConfig": {
	//       "description": "This parameter is currently ignored",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "restrictBy": {
	//       "description": "Restriction specification",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "safe": {
	//       "description": "Whether safe search is enabled. Default: true",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "sayt.enabled": {
	//       "description": "Google Internal",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "sayt.useGcsConfig": {
	//       "description": "Google Internal",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "shelfSpaceAds.enabled": {
	//       "description": "Whether to return shelf space ads",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "shelfSpaceAds.maxResults": {
	//       "description": "The maximum number of shelf space ads to return",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "shelfSpaceAds.useGcsConfig": {
	//       "description": "This parameter is currently ignored",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "source": {
	//       "description": "Query source",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "spelling.enabled": {
	//       "description": "Whether to return spelling suggestions",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "spelling.useGcsConfig": {
	//       "description": "This parameter is currently ignored",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "startIndex": {
	//       "description": "Index (1-based) of first product to return",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "taxonomy": {
	//       "description": "Taxonomy name",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "thumbnails": {
	//       "description": "Image thumbnails specification",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "useCase": {
	//       "description": "One of CommerceSearchUseCase, ShoppingApiUseCase",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "{source}/products",
	//   "response": {
	//     "$ref": "Products"
	//   }
	// }

}

// method id "shopping.products.get":

type ProductsGetCall struct {
	s             *Service
	source        string
	accountId     int64
	productIdType string
	productId     string
	opt_          map[string]interface{}
}

// Get: Returns a single product
func (r *ProductsService) Get(source string, accountId int64, productIdType string, productId string) *ProductsGetCall {
	c := &ProductsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.source = source
	c.accountId = accountId
	c.productIdType = productIdType
	c.productId = productId
	return c
}

// Thumbnails sets the optional parameter "thumbnails": Thumbnail
// specification
func (c *ProductsGetCall) Thumbnails(thumbnails string) *ProductsGetCall {
	c.opt_["thumbnails"] = thumbnails
	return c
}

// Taxonomy sets the optional parameter "taxonomy": Merchant taxonomy
func (c *ProductsGetCall) Taxonomy(taxonomy string) *ProductsGetCall {
	c.opt_["taxonomy"] = taxonomy
	return c
}

// CategoriesUseGcsConfig sets the optional parameter
// "categories.useGcsConfig": This parameter is currently ignored
func (c *ProductsGetCall) CategoriesUseGcsConfig(categoriesUseGcsConfig bool) *ProductsGetCall {
	c.opt_["categories.useGcsConfig"] = categoriesUseGcsConfig
	return c
}

// ProductFields sets the optional parameter "productFields": Google
// Internal
func (c *ProductsGetCall) ProductFields(productFields string) *ProductsGetCall {
	c.opt_["productFields"] = productFields
	return c
}

// DebugEnableLogging sets the optional parameter "debug.enableLogging":
// Google Internal
func (c *ProductsGetCall) DebugEnableLogging(debugEnableLogging bool) *ProductsGetCall {
	c.opt_["debug.enableLogging"] = debugEnableLogging
	return c
}

// DebugSearchRequest sets the optional parameter "debug.searchRequest":
// Google Internal
func (c *ProductsGetCall) DebugSearchRequest(debugSearchRequest bool) *ProductsGetCall {
	c.opt_["debug.searchRequest"] = debugSearchRequest
	return c
}

// Location sets the optional parameter "location": Location used to
// determine tax and shipping
func (c *ProductsGetCall) Location(location string) *ProductsGetCall {
	c.opt_["location"] = location
	return c
}

// RecommendationsEnabled sets the optional parameter
// "recommendations.enabled": Whether to return recommendation
// information
func (c *ProductsGetCall) RecommendationsEnabled(recommendationsEnabled bool) *ProductsGetCall {
	c.opt_["recommendations.enabled"] = recommendationsEnabled
	return c
}

// AttributeFilter sets the optional parameter "attributeFilter": Comma
// separated list of attributes to return
func (c *ProductsGetCall) AttributeFilter(attributeFilter string) *ProductsGetCall {
	c.opt_["attributeFilter"] = attributeFilter
	return c
}

// CategoriesEnabled sets the optional parameter "categories.enabled":
// Whether to return category information
func (c *ProductsGetCall) CategoriesEnabled(categoriesEnabled bool) *ProductsGetCall {
	c.opt_["categories.enabled"] = categoriesEnabled
	return c
}

// DebugEnabled sets the optional parameter "debug.enabled": Google
// Internal
func (c *ProductsGetCall) DebugEnabled(debugEnabled bool) *ProductsGetCall {
	c.opt_["debug.enabled"] = debugEnabled
	return c
}

// RecommendationsInclude sets the optional parameter
// "recommendations.include": Recommendation specification
func (c *ProductsGetCall) RecommendationsInclude(recommendationsInclude string) *ProductsGetCall {
	c.opt_["recommendations.include"] = recommendationsInclude
	return c
}

// CategoriesInclude sets the optional parameter "categories.include":
// Category specification
func (c *ProductsGetCall) CategoriesInclude(categoriesInclude string) *ProductsGetCall {
	c.opt_["categories.include"] = categoriesInclude
	return c
}

// PlusOne sets the optional parameter "plusOne": +1 rendering
// specification.
func (c *ProductsGetCall) PlusOne(plusOne string) *ProductsGetCall {
	c.opt_["plusOne"] = plusOne
	return c
}

// DebugSearchResponse sets the optional parameter
// "debug.searchResponse": Google Internal
func (c *ProductsGetCall) DebugSearchResponse(debugSearchResponse bool) *ProductsGetCall {
	c.opt_["debug.searchResponse"] = debugSearchResponse
	return c
}

// RecommendationsUseGcsConfig sets the optional parameter
// "recommendations.useGcsConfig": This parameter is currently ignored
func (c *ProductsGetCall) RecommendationsUseGcsConfig(recommendationsUseGcsConfig bool) *ProductsGetCall {
	c.opt_["recommendations.useGcsConfig"] = recommendationsUseGcsConfig
	return c
}

func (c *ProductsGetCall) Do() (*Product, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["thumbnails"]; ok {
		params.Set("thumbnails", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["taxonomy"]; ok {
		params.Set("taxonomy", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["categories.useGcsConfig"]; ok {
		params.Set("categories.useGcsConfig", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["productFields"]; ok {
		params.Set("productFields", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["debug.enableLogging"]; ok {
		params.Set("debug.enableLogging", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["debug.searchRequest"]; ok {
		params.Set("debug.searchRequest", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["location"]; ok {
		params.Set("location", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["recommendations.enabled"]; ok {
		params.Set("recommendations.enabled", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["attributeFilter"]; ok {
		params.Set("attributeFilter", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["categories.enabled"]; ok {
		params.Set("categories.enabled", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["debug.enabled"]; ok {
		params.Set("debug.enabled", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["recommendations.include"]; ok {
		params.Set("recommendations.include", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["categories.include"]; ok {
		params.Set("categories.include", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["plusOne"]; ok {
		params.Set("plusOne", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["debug.searchResponse"]; ok {
		params.Set("debug.searchResponse", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["recommendations.useGcsConfig"]; ok {
		params.Set("recommendations.useGcsConfig", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/shopping/search/v1/", "{source}/products/{accountId}/{productIdType}/{productId}")
	urls = strings.Replace(urls, "{source}", cleanPathString(c.source), 1)
	urls = strings.Replace(urls, "{accountId}", strconv.Itoa64(c.accountId), 1)
	urls = strings.Replace(urls, "{productIdType}", cleanPathString(c.productIdType), 1)
	urls = strings.Replace(urls, "{productId}", cleanPathString(c.productId), 1)
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
	ret := new(Product)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns a single product",
	//   "httpMethod": "GET",
	//   "id": "shopping.products.get",
	//   "parameterOrder": [
	//     "source",
	//     "accountId",
	//     "productIdType",
	//     "productId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Merchant center account id",
	//       "format": "uint32",
	//       "location": "path",
	//       "required": true,
	//       "type": "integer"
	//     },
	//     "attributeFilter": {
	//       "description": "Comma separated list of attributes to return",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "categories.enabled": {
	//       "description": "Whether to return category information",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "categories.include": {
	//       "description": "Category specification",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "categories.useGcsConfig": {
	//       "description": "This parameter is currently ignored",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "debug.enableLogging": {
	//       "description": "Google Internal",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "debug.enabled": {
	//       "description": "Google Internal",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "debug.searchRequest": {
	//       "description": "Google Internal",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "debug.searchResponse": {
	//       "description": "Google Internal",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "location": {
	//       "description": "Location used to determine tax and shipping",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "plusOne": {
	//       "description": "+1 rendering specification.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "productFields": {
	//       "description": "Google Internal",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "productId": {
	//       "description": "Id of product",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "productIdType": {
	//       "description": "Type of productId",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "recommendations.enabled": {
	//       "description": "Whether to return recommendation information",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "recommendations.include": {
	//       "description": "Recommendation specification",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "recommendations.useGcsConfig": {
	//       "description": "This parameter is currently ignored",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "source": {
	//       "description": "Query source",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "taxonomy": {
	//       "description": "Merchant taxonomy",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "thumbnails": {
	//       "description": "Thumbnail specification",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "{source}/products/{accountId}/{productIdType}/{productId}",
	//   "response": {
	//     "$ref": "Product"
	//   }
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
