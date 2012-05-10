// Package gan provides access to the Google Affiliate Network API.
//
// See https://code.google.com/apis/gan/
//
// Usage example:
//
//   import "code.google.com/p/google-api-go-client/gan/v1beta1"
//   ...
//   ganService, err := gan.New(oauthHttpClient)
package gan

import (
	"bytes"
	"fmt"
	"net/http"
	"io"
	"encoding/json"
	"errors"
	"strings"
	"strconv"
	"net/url"
	"code.google.com/p/google-api-go-client/googleapi"
)

var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New

const apiId = "gan:v1beta1"
const apiName = "gan"
const apiVersion = "v1beta1"
const basePath = "https://www.googleapis.com/gan/v1beta1/"

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Advertisers = &AdvertisersService{s: s}
	s.CcOffers = &CcOffersService{s: s}
	s.Events = &EventsService{s: s}
	s.Publishers = &PublishersService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	Advertisers *AdvertisersService

	CcOffers *CcOffersService

	Events *EventsService

	Publishers *PublishersService
}

type AdvertisersService struct {
	s *Service
}

type CcOffersService struct {
	s *Service
}

type EventsService struct {
	s *Service
}

type PublishersService struct {
	s *Service
}

type Advertiser struct {
	// ProductFeedsEnabled: Allows advertisers to submit product listings to
	// Google Product Search.
	ProductFeedsEnabled bool `json:"productFeedsEnabled,omitempty"`

	// ContactEmail: Email that this advertiser would like publishers to
	// contact them with.
	ContactEmail string `json:"contactEmail,omitempty"`

	// SiteUrl: URL of the website this advertiser advertises from.
	SiteUrl string `json:"siteUrl,omitempty"`

	// PayoutRank: A rank based on commissions paid to publishers over the
	// past 90 days. A number between 1 and 4 where 4 means the top quartile
	// (most money paid) and 1 means the bottom quartile (least money paid).
	PayoutRank string `json:"payoutRank,omitempty"`

	// Description: Description of the website the advertiser advertises
	// from.
	Description string `json:"description,omitempty"`

	// Item: The requested advertiser.
	Item *Advertiser `json:"item,omitempty"`

	// Category: Category that this advertiser belongs to. A valid list of
	// categories can be found here:
	// http://www.google.com/support/affiliatenetwork/advertiser/bin/answer.p
	// y?hl=en&answer=107581
	Category string `json:"category,omitempty"`

	// ContactPhone: Phone that this advertiser would like publishers to
	// contact them with.
	ContactPhone string `json:"contactPhone,omitempty"`

	// Name: The name of this advertiser.
	Name string `json:"name,omitempty"`

	// LogoUrl: URL to the logo this advertiser uses on the Google Affiliate
	// Network.
	LogoUrl string `json:"logoUrl,omitempty"`

	// CommissionDuration: The longest possible length of a commission (how
	// long the cookies on the customer's browser last before they expire).
	CommissionDuration int64 `json:"commissionDuration,omitempty"`

	// Kind: The kind for an advertiser.
	Kind string `json:"kind,omitempty"`

	// JoinDate: Date that this advertiser was approved as a Google
	// Affiliate Network advertiser.
	JoinDate string `json:"joinDate,omitempty"`

	// Id: The ID of this advertiser.
	Id int64 `json:"id,omitempty,string"`

	// Status: The status of the requesting publisher's relationship this
	// advertiser.
	Status string `json:"status,omitempty"`

	// EpcNinetyDayAverage: The sum of fees paid to publishers divided by
	// the total number of clicks over the past three months. Values are
	// multiplied by 100 for display purposes.
	EpcNinetyDayAverage *Money `json:"epcNinetyDayAverage,omitempty"`

	// EpcSevenDayAverage: The sum of fees paid to publishers divided by the
	// total number of clicks over the past seven days. Values are
	// multiplied by 100 for display purposes.
	EpcSevenDayAverage *Money `json:"epcSevenDayAverage,omitempty"`
}

type Advertisers struct {
	// Items: The advertiser list.
	Items []*Advertiser `json:"items,omitempty"`

	// NextPageToken: The 'pageToken' to pass to the next request to get the
	// next page, if there are more to retrieve.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Kind: The kind for a page of advertisers.
	Kind string `json:"kind,omitempty"`
}

type CcOffer struct {
	// RewardsExpire: Whether accumulated rewards ever expire.
	RewardsExpire bool `json:"rewardsExpire,omitempty"`

	// BonusRewards: For cards with rewards programs, extra circumstances
	// whereby additional rewards may be granted.
	BonusRewards []*CcOfferBonusRewards `json:"bonusRewards,omitempty"`

	// AnnualFeeDisplay: Text describing the annual fee, including any
	// difference for the first year. A summary field.
	AnnualFeeDisplay string `json:"annualFeeDisplay,omitempty"`

	// AgeMinimum: The youngest a recipient of this card may be.
	AgeMinimum float64 `json:"ageMinimum,omitempty"`

	// StatementCopyFee: Fee for requesting a copy of your statement.
	StatementCopyFee string `json:"statementCopyFee,omitempty"`

	// CardType: What kind of credit card this is, for example secured or
	// unsecured.
	CardType string `json:"cardType,omitempty"`

	// CashAdvanceTerms: Text describing the terms for cash advances. A
	// summary field.
	CashAdvanceTerms string `json:"cashAdvanceTerms,omitempty"`

	// CreditRatingDisplay: Text describing the credit ratings required for
	// recipients of this card, for example "Excellent/Good." A summary
	// field.
	CreditRatingDisplay string `json:"creditRatingDisplay,omitempty"`

	// CardBenefits: A list of what the issuer thinks are the most important
	// benefits of the card. Usually summarizes the rewards program, if
	// there is one. A summary field.
	CardBenefits []string `json:"cardBenefits,omitempty"`

	// FlightAccidentInsurance: If you get coverage when you use the card
	// for the given activity, this field describes it.
	FlightAccidentInsurance string `json:"flightAccidentInsurance,omitempty"`

	// IssuerWebsite: The generic link to the issuer's site.
	IssuerWebsite string `json:"issuerWebsite,omitempty"`

	// VariableRatesLastUpdated: When variable rates were last updated.
	VariableRatesLastUpdated string `json:"variableRatesLastUpdated,omitempty"`

	// LuggageInsurance: If you get coverage when you use the card for the
	// given activity, this field describes it.
	LuggageInsurance string `json:"luggageInsurance,omitempty"`

	// AnnualFee: The ongoing annual fee, in dollars.
	AnnualFee float64 `json:"annualFee,omitempty"`

	// IssuerId: The Google Affiliate Network ID of the advertiser making
	// this offer.
	IssuerId string `json:"issuerId,omitempty"`

	// AnnualRewardMaximum: The largest number of units you may accumulate
	// in a year.
	AnnualRewardMaximum float64 `json:"annualRewardMaximum,omitempty"`

	// TrackingUrl: The link to ping to register a click on this offer. A
	// summary field.
	TrackingUrl string `json:"trackingUrl,omitempty"`

	// ReturnedPaymentFee: Text describing the fee for a payment that
	// doesn't clear. A summary field.
	ReturnedPaymentFee string `json:"returnedPaymentFee,omitempty"`

	// PurchaseRateAdditionalDetails: Text describing any additional details
	// for the purchase rate. A summary field.
	PurchaseRateAdditionalDetails string `json:"purchaseRateAdditionalDetails,omitempty"`

	// ProhibitedCategories: Categories in which the issuer does not wish
	// the card to be displayed. A summary field.
	ProhibitedCategories []string `json:"prohibitedCategories,omitempty"`

	// OverLimitFee: Fee for exceeding the card's charge limit.
	OverLimitFee string `json:"overLimitFee,omitempty"`

	// ImageUrl: The link to the image of the card that is shown on Connect
	// Commerce. A summary field.
	ImageUrl string `json:"imageUrl,omitempty"`

	// BalanceTransferTerms: Text describing the terms for balance
	// transfers. A summary field.
	BalanceTransferTerms string `json:"balanceTransferTerms,omitempty"`

	// Network: Which network (eg Visa) the card belongs to. A summary
	// field.
	Network string `json:"network,omitempty"`

	// Disclaimer: A notice that, if present, is referenced via an asterisk
	// by many of the other summary fields. If this field is present, it
	// will always start with an asterisk ("*"), and must be prominently
	// displayed with the offer. A summary field.
	Disclaimer string `json:"disclaimer,omitempty"`

	// TravelInsurance: If you get coverage when you use the card for the
	// given activity, this field describes it.
	TravelInsurance string `json:"travelInsurance,omitempty"`

	// AprDisplay: Text describing the purchase APR. A summary field.
	AprDisplay string `json:"aprDisplay,omitempty"`

	// VariableRatesUpdateFrequency: How often variable rates are updated.
	VariableRatesUpdateFrequency string `json:"variableRatesUpdateFrequency,omitempty"`

	// CardName: The issuer's name for the card, including any trademark or
	// service mark designators. A summary field.
	CardName string `json:"cardName,omitempty"`

	// ForeignCurrencyTransactionFee: Fee for each transaction involving a
	// foreign currency.
	ForeignCurrencyTransactionFee string `json:"foreignCurrencyTransactionFee,omitempty"`

	// ExistingCustomerOnly: Whether this card is only available to existing
	// customers of the issuer.
	ExistingCustomerOnly bool `json:"existingCustomerOnly,omitempty"`

	// EmergencyInsurance: If you get coverage when you use the card for the
	// given activity, this field describes it.
	EmergencyInsurance string `json:"emergencyInsurance,omitempty"`

	// CreditLimitMax: The high end for credit limits the issuer imposes on
	// recipients of this card.
	CreditLimitMax float64 `json:"creditLimitMax,omitempty"`

	// AgeMinimumDetails: Text describing the details of the age minimum
	// restriction.
	AgeMinimumDetails string `json:"ageMinimumDetails,omitempty"`

	// GracePeriodDisplay: Text describing the grace period before finance
	// charges apply. A summary field.
	GracePeriodDisplay string `json:"gracePeriodDisplay,omitempty"`

	// MinPurchaseRate: The lowest interest rate the issuer charges on this
	// card. Expressed as an absolute number, not as a percentage.
	MinPurchaseRate float64 `json:"minPurchaseRate,omitempty"`

	// RewardPartner: The company that redeems the rewards, if different
	// from the issuer.
	RewardPartner string `json:"rewardPartner,omitempty"`

	// MaxPurchaseRate: The highest interest rate the issuer charges on this
	// card. Expressed as an absolute number, not as a percentage.
	MaxPurchaseRate float64 `json:"maxPurchaseRate,omitempty"`

	// ExtendedWarranty: If you get coverage when you use the card for the
	// given activity, this field describes it.
	ExtendedWarranty string `json:"extendedWarranty,omitempty"`

	// OfferId: This offer's ID. A summary field.
	OfferId string `json:"offerId,omitempty"`

	// InitialSetupAndProcessingFee: Fee for setting up the card.
	InitialSetupAndProcessingFee string `json:"initialSetupAndProcessingFee,omitempty"`

	// IntroBalanceTransferTerms: Text describing the terms for introductory
	// period balance transfers. A summary field.
	IntroBalanceTransferTerms string `json:"introBalanceTransferTerms,omitempty"`

	// PurchaseRateType: Fixed or variable.
	PurchaseRateType string `json:"purchaseRateType,omitempty"`

	// Issuer: Name of card issuer. A summary field.
	Issuer string `json:"issuer,omitempty"`

	// CarRentalInsurance: If you get coverage when you use the card for the
	// given activity, this field describes it.
	CarRentalInsurance string `json:"carRentalInsurance,omitempty"`

	// AdditionalCardBenefits: More marketing copy about the card's
	// benefits. A summary field.
	AdditionalCardBenefits []string `json:"additionalCardBenefits,omitempty"`

	// AdditionalCardHolderFee: Any extra fees levied on card holders.
	AdditionalCardHolderFee string `json:"additionalCardHolderFee,omitempty"`

	// RewardUnit: For cards with rewards programs, the unit of reward. For
	// example, miles, cash back, points.
	RewardUnit string `json:"rewardUnit,omitempty"`

	// LandingPageUrl: The link to the issuer's page for this card. A
	// summary field.
	LandingPageUrl string `json:"landingPageUrl,omitempty"`

	// IntroPurchaseTerms: Text describing the terms for introductory period
	// purchases. A summary field.
	IntroPurchaseTerms string `json:"introPurchaseTerms,omitempty"`

	// RewardsHaveBlackoutDates: For airline miles rewards, tells whether
	// blackout dates apply to the miles.
	RewardsHaveBlackoutDates bool `json:"rewardsHaveBlackoutDates,omitempty"`

	// CreditLimitMin: The low end for credit limits the issuer imposes on
	// recipients of this card.
	CreditLimitMin float64 `json:"creditLimitMin,omitempty"`

	// FraudLiability: If you get coverage when you use the card for the
	// given activity, this field describes it.
	FraudLiability string `json:"fraudLiability,omitempty"`

	// MinimumFinanceCharge: Text describing how much missing the grace
	// period will cost.
	MinimumFinanceCharge string `json:"minimumFinanceCharge,omitempty"`

	// IntroCashAdvanceTerms: Text describing the terms for introductory
	// period cash advances. A summary field.
	IntroCashAdvanceTerms string `json:"introCashAdvanceTerms,omitempty"`

	// DefaultFees: Fees for defaulting on your payments.
	DefaultFees []*CcOfferDefaultFees `json:"defaultFees,omitempty"`

	// Rewards: For cards with rewards programs, detailed rules about how
	// the program works.
	Rewards []*CcOfferRewards `json:"rewards,omitempty"`

	// ApprovedCategories: Possible categories for this card, eg "Low
	// Interest" or "Good." A summary field.
	ApprovedCategories []string `json:"approvedCategories,omitempty"`

	// Kind: The kind for one credit card offer. A summary field.
	Kind string `json:"kind,omitempty"`

	// LatePaymentFee: Text describing how much a late payment will cost, eg
	// "up to $35." A summary field.
	LatePaymentFee string `json:"latePaymentFee,omitempty"`

	// FirstYearAnnualFee: The annual fee for the first year, if different
	// from the ongoing fee. Optional.
	FirstYearAnnualFee float64 `json:"firstYearAnnualFee,omitempty"`

	// BalanceComputationMethod: Text describing how the balance is
	// computed. A summary field.
	BalanceComputationMethod string `json:"balanceComputationMethod,omitempty"`

	// OffersImmediateCashReward: Whether a cash reward program lets you get
	// cash back sooner than end of year or other longish period.
	OffersImmediateCashReward bool `json:"offersImmediateCashReward,omitempty"`
}

type CcOfferBonusRewards struct {
	// Details: The circumstances under which this rule applies, for
	// example, booking a flight via Orbitz.
	Details string `json:"details,omitempty"`

	// Amount: How many units of reward will be granted.
	Amount float64 `json:"amount,omitempty"`
}

type CcOfferDefaultFees struct {
	// Category: The type of charge, for example Purchases.
	Category string `json:"category,omitempty"`

	// MinRate: The lowest rate the issuer may charge for defaulting on debt
	// in this category. Expressed as an absolute number, not as a
	// percentage.
	MinRate float64 `json:"minRate,omitempty"`

	// MaxRate: The highest rate the issuer may charge for defaulting on
	// debt in this category. Expressed as an absolute number, not as a
	// percentage.
	MaxRate float64 `json:"maxRate,omitempty"`

	// RateType: Fixed or variable.
	RateType string `json:"rateType,omitempty"`
}

type CcOfferRewards struct {
	// MaxRewardTier: The maximum purchase amount in the given category for
	// this rule to apply.
	MaxRewardTier float64 `json:"maxRewardTier,omitempty"`

	// ExpirationMonths: How long rewards granted by this rule last.
	ExpirationMonths float64 `json:"expirationMonths,omitempty"`

	// AdditionalDetails: Other limits, for example, if this rule only
	// applies during an introductory period.
	AdditionalDetails string `json:"additionalDetails,omitempty"`

	// Amount: The number of units rewarded per purchase dollar.
	Amount float64 `json:"amount,omitempty"`

	// Category: The kind of purchases covered by this rule.
	Category string `json:"category,omitempty"`

	// MinRewardTier: The minimum purchase amount in the given category
	// before this rule applies.
	MinRewardTier float64 `json:"minRewardTier,omitempty"`
}

type CcOffers struct {
	// Items: The credit card offers.
	Items []*CcOffer `json:"items,omitempty"`

	// Kind: The kind for a page of credit card offers.
	Kind string `json:"kind,omitempty"`
}

type Event struct {
	// NetworkFee: Fee that the advertiser paid to the Google Affiliate
	// Network.
	NetworkFee *Money `json:"networkFee,omitempty"`

	// ChargeType: Charge type of the event
	// (other|slotting_fee|monthly_minimum|tier_bonus|debit|credit). Only
	// returned for charge events.
	ChargeType string `json:"chargeType,omitempty"`

	// Kind: The kind for one event.
	Kind string `json:"kind,omitempty"`

	// ChargeId: The charge ID for this event. Only returned for charge
	// events.
	ChargeId string `json:"chargeId,omitempty"`

	// PublisherFee: Fee that the advertiser paid to the publisher.
	PublisherFee *Money `json:"publisherFee,omitempty"`

	// CommissionableSales: Amount of money exchanged during the
	// transaction. Only returned for charge and conversion events.
	CommissionableSales *Money `json:"commissionableSales,omitempty"`

	// Products: Products associated with the event.
	Products []*EventProducts `json:"products,omitempty"`

	// Status: Status of the event (active|canceled). Only returned for
	// charge and conversion events.
	Status string `json:"status,omitempty"`

	// EventDate: The date-time this event was initiated as a RFC 3339
	// date-time value.
	EventDate string `json:"eventDate,omitempty"`

	// ModifyDate: The date-time this event was last modified as a RFC 3339
	// date-time value.
	ModifyDate string `json:"modifyDate,omitempty"`

	// Type: Type of the event (action|transaction|charge).
	Type string `json:"type,omitempty"`

	// PublisherId: The ID of the publisher for this event.
	PublisherId int64 `json:"publisherId,omitempty,string"`

	// MemberId: The ID of the member attached to this event. Only returned
	// for conversion events.
	MemberId string `json:"memberId,omitempty"`

	// AdvertiserId: The ID of advertiser for this event.
	AdvertiserId int64 `json:"advertiserId,omitempty,string"`

	// PublisherName: The name of the publisher for this event.
	PublisherName string `json:"publisherName,omitempty"`

	// Earnings: Earnings by the publisher.
	Earnings *Money `json:"earnings,omitempty"`

	// OrderId: The order ID for this event. Only returned for conversion
	// events.
	OrderId string `json:"orderId,omitempty"`

	// AdvertiserName: The name of the advertiser for this event.
	AdvertiserName string `json:"advertiserName,omitempty"`
}

type EventProducts struct {
	// PublisherFee: Fee that the advertiser paid to the publisehr for this
	// product.
	PublisherFee *Money `json:"publisherFee,omitempty"`

	// CategoryId: Id of the category this product belongs to.
	CategoryId string `json:"categoryId,omitempty"`

	// CategoryName: Name of the category this product belongs to.
	CategoryName string `json:"categoryName,omitempty"`

	// SkuName: Sku name of this product.
	SkuName string `json:"skuName,omitempty"`

	// UnitPrice: Price per unit of this product.
	UnitPrice *Money `json:"unitPrice,omitempty"`

	// Sku: Sku of this product.
	Sku string `json:"sku,omitempty"`

	// Quantity: Quantity of this product bought/exchanged.
	Quantity int64 `json:"quantity,omitempty,string"`

	// Earnings: Amount earned by the publisher on this product.
	Earnings *Money `json:"earnings,omitempty"`

	// NetworkFee: Fee that the advertiser paid to the Google Affiliate
	// Network for this product.
	NetworkFee *Money `json:"networkFee,omitempty"`
}

type Events struct {
	// Items: The event list.
	Items []*Event `json:"items,omitempty"`

	// NextPageToken: The 'pageToken' to pass to the next request to get the
	// next page, if there are more to retrieve.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Kind: The kind for a page of events.
	Kind string `json:"kind,omitempty"`
}

type Money struct {
	// Amount: The amount of money.
	Amount float64 `json:"amount,omitempty"`

	// CurrencyCode: The 3-letter code of the currency in question.
	CurrencyCode string `json:"currencyCode,omitempty"`
}

type Publisher struct {
	// PayoutRank: A rank based on commissions paid to this publisher over
	// the past 90 days. A number between 1 and 4 where 4 means the top
	// quartile (most money paid) and 1 means the bottom quartile (least
	// money paid).
	PayoutRank string `json:"payoutRank,omitempty"`

	// Sites: Websites that this publisher uses to advertise.
	Sites []string `json:"sites,omitempty"`

	// Item: The requested publisher.
	Item *Publisher `json:"item,omitempty"`

	// Name: The name of this publisher.
	Name string `json:"name,omitempty"`

	// Kind: The kind for a publisher.
	Kind string `json:"kind,omitempty"`

	// JoinDate: Date that this publisher was approved as a Google Affiliate
	// Network publisher.
	JoinDate string `json:"joinDate,omitempty"`

	// Classification: Classification that this publisher belongs to. See
	// this link for all publisher classifications:
	// http://www.google.com/support/affiliatenetwork/advertiser/bin/answer.p
	// y?hl=en&answer=107625&ctx=cb&src=cb&cbid=-k5fihzthfaik&cbrank=4
	Classification string `json:"classification,omitempty"`

	// Id: The ID of this publisher.
	Id int64 `json:"id,omitempty,string"`

	// Status: The status of the requesting advertiser's relationship with
	// this publisher.
	Status string `json:"status,omitempty"`

	// EpcNinetyDayAverage: The sum of fees paid to this publisher divided
	// by the total number of clicks over the past three months. Values are
	// multiplied by 100 for display purposes.
	EpcNinetyDayAverage *Money `json:"epcNinetyDayAverage,omitempty"`

	// EpcSevenDayAverage: The sum of fees paid to this publisher divided by
	// the total number of clicks over the past seven days. Values are
	// multiplied by 100 for display purposes.
	EpcSevenDayAverage *Money `json:"epcSevenDayAverage,omitempty"`
}

type Publishers struct {
	// Items: The entity list.
	Items []*Publisher `json:"items,omitempty"`

	// NextPageToken: The 'pageToken' to pass to the next request to get the
	// next page, if there are more to retrieve.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Kind: The kind for a page of entities.
	Kind string `json:"kind,omitempty"`
}

// method id "gan.advertisers.get":

type AdvertisersGetCall struct {
	s      *Service
	role   string
	roleId string
	opt_   map[string]interface{}
}

// Get: Retrieves data about a single advertiser if that the requesting
// advertiser/publisher has access to it. Only publishers can lookup
// advertisers. Advertisers can request information about themselves by
// omitting the advertiserId query parameter.
func (r *AdvertisersService) Get(role string, roleId string) *AdvertisersGetCall {
	c := &AdvertisersGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.role = role
	c.roleId = roleId
	return c
}

// AdvertiserId sets the optional parameter "advertiserId": The ID of
// the advertiser to look up.
func (c *AdvertisersGetCall) AdvertiserId(advertiserId string) *AdvertisersGetCall {
	c.opt_["advertiserId"] = advertiserId
	return c
}

func (c *AdvertisersGetCall) Do() (*Advertiser, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["advertiserId"]; ok {
		params.Set("advertiserId", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/gan/v1beta1/", "{role}/{roleId}/advertiser")
	urls = strings.Replace(urls, "{role}", cleanPathString(c.role), 1)
	urls = strings.Replace(urls, "{roleId}", cleanPathString(c.roleId), 1)
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
	ret := new(Advertiser)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves data about a single advertiser if that the requesting advertiser/publisher has access to it. Only publishers can lookup advertisers. Advertisers can request information about themselves by omitting the advertiserId query parameter.",
	//   "httpMethod": "GET",
	//   "id": "gan.advertisers.get",
	//   "parameterOrder": [
	//     "role",
	//     "roleId"
	//   ],
	//   "parameters": {
	//     "advertiserId": {
	//       "description": "The ID of the advertiser to look up. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "role": {
	//       "description": "The role of the requester. Valid values: 'advertisers' or 'publishers'.",
	//       "enum": [
	//         "advertisers",
	//         "publishers"
	//       ],
	//       "enumDescriptions": [
	//         "The requester is requesting as an advertiser.",
	//         "The requester is requesting as a publisher."
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "roleId": {
	//       "description": "The ID of the requesting advertiser or publisher.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{role}/{roleId}/advertiser",
	//   "response": {
	//     "$ref": "Advertiser"
	//   }
	// }

}

// method id "gan.advertisers.list":

type AdvertisersListCall struct {
	s      *Service
	role   string
	roleId string
	opt_   map[string]interface{}
}

// List: Retrieves data about all advertisers that the requesting
// advertiser/publisher has access to.
func (r *AdvertisersService) List(role string, roleId string) *AdvertisersListCall {
	c := &AdvertisersListCall{s: r.s, opt_: make(map[string]interface{})}
	c.role = role
	c.roleId = roleId
	return c
}

// AdvertiserCategory sets the optional parameter "advertiserCategory":
// Caret(^) delimted list of advertiser categories. Valid categories are
// defined here:
// http://www.google.com/support/affiliatenetwork/advertiser/bin/answer.p
// y?hl=en&answer=107581. Filters out all advertisers not in one of the
// given advertiser categories.
func (c *AdvertisersListCall) AdvertiserCategory(advertiserCategory string) *AdvertisersListCall {
	c.opt_["advertiserCategory"] = advertiserCategory
	return c
}

// MaxResults sets the optional parameter "maxResults": Max number of
// items to return in this page.  Defaults to 20.
func (c *AdvertisersListCall) MaxResults(maxResults int64) *AdvertisersListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// MinNinetyDayEpc sets the optional parameter "minNinetyDayEpc":
// Filters out all advertisers that have a ninety day EPC average lower
// than the given value (inclusive). Min value: 0.0.
func (c *AdvertisersListCall) MinNinetyDayEpc(minNinetyDayEpc float64) *AdvertisersListCall {
	c.opt_["minNinetyDayEpc"] = minNinetyDayEpc
	return c
}

// MinPayoutRank sets the optional parameter "minPayoutRank": A value
// between 1 and 4, where 1 represents the quartile of advertisers with
// the lowest ranks and 4 represents the quartile of advertisers with
// the highest ranks. Filters out all advertisers with a lower rank than
// the given quartile. For example if a 2 was given only advertisers
// with a payout rank of 25 or higher would be included.
func (c *AdvertisersListCall) MinPayoutRank(minPayoutRank int64) *AdvertisersListCall {
	c.opt_["minPayoutRank"] = minPayoutRank
	return c
}

// MinSevenDayEpc sets the optional parameter "minSevenDayEpc": Filters
// out all advertisers that have a seven day EPC average lower than the
// given value (inclusive). Min value: 0.0.
func (c *AdvertisersListCall) MinSevenDayEpc(minSevenDayEpc float64) *AdvertisersListCall {
	c.opt_["minSevenDayEpc"] = minSevenDayEpc
	return c
}

// PageToken sets the optional parameter "pageToken": The value of
// 'nextPageToken' from the previous page.
func (c *AdvertisersListCall) PageToken(pageToken string) *AdvertisersListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// RelationshipStatus sets the optional parameter "relationshipStatus":
// Filters out all advertisers for which do not have the given
// relationship status with the requesting publisher.
func (c *AdvertisersListCall) RelationshipStatus(relationshipStatus string) *AdvertisersListCall {
	c.opt_["relationshipStatus"] = relationshipStatus
	return c
}

func (c *AdvertisersListCall) Do() (*Advertisers, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["advertiserCategory"]; ok {
		params.Set("advertiserCategory", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["minNinetyDayEpc"]; ok {
		params.Set("minNinetyDayEpc", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["minPayoutRank"]; ok {
		params.Set("minPayoutRank", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["minSevenDayEpc"]; ok {
		params.Set("minSevenDayEpc", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["relationshipStatus"]; ok {
		params.Set("relationshipStatus", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/gan/v1beta1/", "{role}/{roleId}/advertisers")
	urls = strings.Replace(urls, "{role}", cleanPathString(c.role), 1)
	urls = strings.Replace(urls, "{roleId}", cleanPathString(c.roleId), 1)
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
	ret := new(Advertisers)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves data about all advertisers that the requesting advertiser/publisher has access to.",
	//   "httpMethod": "GET",
	//   "id": "gan.advertisers.list",
	//   "parameterOrder": [
	//     "role",
	//     "roleId"
	//   ],
	//   "parameters": {
	//     "advertiserCategory": {
	//       "description": "Caret(^) delimted list of advertiser categories. Valid categories are defined here: http://www.google.com/support/affiliatenetwork/advertiser/bin/answer.py?hl=en&answer=107581. Filters out all advertisers not in one of the given advertiser categories. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Max number of items to return in this page. Optional. Defaults to 20.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "100",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "minNinetyDayEpc": {
	//       "description": "Filters out all advertisers that have a ninety day EPC average lower than the given value (inclusive). Min value: 0.0. Optional.",
	//       "format": "double",
	//       "location": "query",
	//       "type": "number"
	//     },
	//     "minPayoutRank": {
	//       "description": "A value between 1 and 4, where 1 represents the quartile of advertisers with the lowest ranks and 4 represents the quartile of advertisers with the highest ranks. Filters out all advertisers with a lower rank than the given quartile. For example if a 2 was given only advertisers with a payout rank of 25 or higher would be included. Optional.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "4",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "minSevenDayEpc": {
	//       "description": "Filters out all advertisers that have a seven day EPC average lower than the given value (inclusive). Min value: 0.0. Optional.",
	//       "format": "double",
	//       "location": "query",
	//       "type": "number"
	//     },
	//     "pageToken": {
	//       "description": "The value of 'nextPageToken' from the previous page. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "relationshipStatus": {
	//       "description": "Filters out all advertisers for which do not have the given relationship status with the requesting publisher.",
	//       "enum": [
	//         "approved",
	//         "available",
	//         "deactivated",
	//         "declined",
	//         "pending"
	//       ],
	//       "enumDescriptions": [
	//         "An advertiser that has approved your application.",
	//         "An advertiser program that's accepting new publishers.",
	//         "Deactivated means either the advertiser has removed you from their program, or it could also mean that you chose to remove yourself from the advertiser's program.",
	//         "An advertiser that did not approve your application.",
	//         "An advertiser program that you've already applied to, but they haven't yet decided to approve or decline your application."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "role": {
	//       "description": "The role of the requester. Valid values: 'advertisers' or 'publishers'.",
	//       "enum": [
	//         "advertisers",
	//         "publishers"
	//       ],
	//       "enumDescriptions": [
	//         "The requester is requesting as an advertiser.",
	//         "The requester is requesting as a publisher."
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "roleId": {
	//       "description": "The ID of the requesting advertiser or publisher.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{role}/{roleId}/advertisers",
	//   "response": {
	//     "$ref": "Advertisers"
	//   }
	// }

}

// method id "gan.ccOffers.list":

type CcOffersListCall struct {
	s         *Service
	publisher string
	opt_      map[string]interface{}
}

// List: Retrieves credit card offers for the given publisher.
func (r *CcOffersService) List(publisher string) *CcOffersListCall {
	c := &CcOffersListCall{s: r.s, opt_: make(map[string]interface{})}
	c.publisher = publisher
	return c
}

// Advertiser sets the optional parameter "advertiser": The advertiser
// ID of a card issuer whose offers to include. Optional, may be
// repeated.
func (c *CcOffersListCall) Advertiser(advertiser string) *CcOffersListCall {
	c.opt_["advertiser"] = advertiser
	return c
}

// Projection sets the optional parameter "projection": The set of
// fields to return.
func (c *CcOffersListCall) Projection(projection string) *CcOffersListCall {
	c.opt_["projection"] = projection
	return c
}

func (c *CcOffersListCall) Do() (*CcOffers, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["advertiser"]; ok {
		params.Set("advertiser", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["projection"]; ok {
		params.Set("projection", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/gan/v1beta1/", "publishers/{publisher}/ccOffers")
	urls = strings.Replace(urls, "{publisher}", cleanPathString(c.publisher), 1)
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
	ret := new(CcOffers)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves credit card offers for the given publisher.",
	//   "httpMethod": "GET",
	//   "id": "gan.ccOffers.list",
	//   "parameterOrder": [
	//     "publisher"
	//   ],
	//   "parameters": {
	//     "advertiser": {
	//       "description": "The advertiser ID of a card issuer whose offers to include. Optional, may be repeated.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "The set of fields to return.",
	//       "enum": [
	//         "full",
	//         "summary"
	//       ],
	//       "enumDescriptions": [
	//         "Include all offer fields",
	//         "Include only the basic fields needed to display an offer. This is the default."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "publisher": {
	//       "description": "The ID of the publisher in question.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "publishers/{publisher}/ccOffers",
	//   "response": {
	//     "$ref": "CcOffers"
	//   }
	// }

}

// method id "gan.events.list":

type EventsListCall struct {
	s      *Service
	role   string
	roleId string
	opt_   map[string]interface{}
}

// List: Retrieves event data for a given advertiser/publisher.
func (r *EventsService) List(role string, roleId string) *EventsListCall {
	c := &EventsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.role = role
	c.roleId = roleId
	return c
}

// AdvertiserId sets the optional parameter "advertiserId": Caret(^)
// delimited list of advertiser IDs. Filters out all events that do not
// reference one of the given advertiser IDs. Only used when under
// publishers role.
func (c *EventsListCall) AdvertiserId(advertiserId string) *EventsListCall {
	c.opt_["advertiserId"] = advertiserId
	return c
}

// ChargeType sets the optional parameter "chargeType": Filters out all
// charge events that are not of the given charge type. Valid values:
// 'other', 'slotting_fee', 'monthly_minimum', 'tier_bonus', 'credit',
// 'debit'.
func (c *EventsListCall) ChargeType(chargeType string) *EventsListCall {
	c.opt_["chargeType"] = chargeType
	return c
}

// EventDateMax sets the optional parameter "eventDateMax": Filters out
// all events later than given date.  Defaults to 24 hours after
// eventMin.
func (c *EventsListCall) EventDateMax(eventDateMax string) *EventsListCall {
	c.opt_["eventDateMax"] = eventDateMax
	return c
}

// EventDateMin sets the optional parameter "eventDateMin": Filters out
// all events earlier than given date.  Defaults to 24 hours from
// current date/time.
func (c *EventsListCall) EventDateMin(eventDateMin string) *EventsListCall {
	c.opt_["eventDateMin"] = eventDateMin
	return c
}

// LinkId sets the optional parameter "linkId": Caret(^) delimited list
// of link IDs. Filters out all events that do not reference one of the
// given link IDs.
func (c *EventsListCall) LinkId(linkId string) *EventsListCall {
	c.opt_["linkId"] = linkId
	return c
}

// MaxResults sets the optional parameter "maxResults": Max number of
// offers to return in this page.  Defaults to 20.
func (c *EventsListCall) MaxResults(maxResults int64) *EventsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// MemberId sets the optional parameter "memberId": Caret(^) delimited
// list of member IDs. Filters out all events that do not reference one
// of the given member IDs.
func (c *EventsListCall) MemberId(memberId string) *EventsListCall {
	c.opt_["memberId"] = memberId
	return c
}

// ModifyDateMax sets the optional parameter "modifyDateMax": Filters
// out all events modified later than given date.  Defaults to 24 hours
// after modifyDateMin, if modifyDateMin is explicitly set.
func (c *EventsListCall) ModifyDateMax(modifyDateMax string) *EventsListCall {
	c.opt_["modifyDateMax"] = modifyDateMax
	return c
}

// ModifyDateMin sets the optional parameter "modifyDateMin": Filters
// out all events modified earlier than given date.  Defaults to 24
// hours before the current modifyDateMax, if modifyDateMax is
// explicitly set.
func (c *EventsListCall) ModifyDateMin(modifyDateMin string) *EventsListCall {
	c.opt_["modifyDateMin"] = modifyDateMin
	return c
}

// OrderId sets the optional parameter "orderId": Caret(^) delimited
// list of order IDs. Filters out all events that do not reference one
// of the given order IDs.
func (c *EventsListCall) OrderId(orderId string) *EventsListCall {
	c.opt_["orderId"] = orderId
	return c
}

// PageToken sets the optional parameter "pageToken": The value of
// 'nextPageToken' from the previous page.
func (c *EventsListCall) PageToken(pageToken string) *EventsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// ProductCategory sets the optional parameter "productCategory":
// Caret(^) delimited list of product categories. Filters out all events
// that do not reference a product in one of the given product
// categories.
func (c *EventsListCall) ProductCategory(productCategory string) *EventsListCall {
	c.opt_["productCategory"] = productCategory
	return c
}

// PublisherId sets the optional parameter "publisherId": Caret(^)
// delimited list of publisher IDs. Filters out all events that do not
// reference one of the given publishers IDs. Only used when under
// advertiser role.
func (c *EventsListCall) PublisherId(publisherId string) *EventsListCall {
	c.opt_["publisherId"] = publisherId
	return c
}

// Sku sets the optional parameter "sku": Caret(^) delimited list of
// SKUs. Filters out all events that do not reference one of the given
// SKU.
func (c *EventsListCall) Sku(sku string) *EventsListCall {
	c.opt_["sku"] = sku
	return c
}

// Status sets the optional parameter "status": Filters out all events
// that do not have the given status. Valid values: 'active',
// 'canceled'.
func (c *EventsListCall) Status(status string) *EventsListCall {
	c.opt_["status"] = status
	return c
}

// Type sets the optional parameter "type": Filters out all events that
// are not of the given type. Valid values: 'action', 'transaction',
// 'charge'.
func (c *EventsListCall) Type(type_ string) *EventsListCall {
	c.opt_["type"] = type_
	return c
}

func (c *EventsListCall) Do() (*Events, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["advertiserId"]; ok {
		params.Set("advertiserId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["chargeType"]; ok {
		params.Set("chargeType", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["eventDateMax"]; ok {
		params.Set("eventDateMax", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["eventDateMin"]; ok {
		params.Set("eventDateMin", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["linkId"]; ok {
		params.Set("linkId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["memberId"]; ok {
		params.Set("memberId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["modifyDateMax"]; ok {
		params.Set("modifyDateMax", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["modifyDateMin"]; ok {
		params.Set("modifyDateMin", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["orderId"]; ok {
		params.Set("orderId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["productCategory"]; ok {
		params.Set("productCategory", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["publisherId"]; ok {
		params.Set("publisherId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["sku"]; ok {
		params.Set("sku", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["status"]; ok {
		params.Set("status", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["type"]; ok {
		params.Set("type", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/gan/v1beta1/", "{role}/{roleId}/events")
	urls = strings.Replace(urls, "{role}", cleanPathString(c.role), 1)
	urls = strings.Replace(urls, "{roleId}", cleanPathString(c.roleId), 1)
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
	ret := new(Events)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves event data for a given advertiser/publisher.",
	//   "httpMethod": "GET",
	//   "id": "gan.events.list",
	//   "parameterOrder": [
	//     "role",
	//     "roleId"
	//   ],
	//   "parameters": {
	//     "advertiserId": {
	//       "description": "Caret(^) delimited list of advertiser IDs. Filters out all events that do not reference one of the given advertiser IDs. Only used when under publishers role. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "chargeType": {
	//       "description": "Filters out all charge events that are not of the given charge type. Valid values: 'other', 'slotting_fee', 'monthly_minimum', 'tier_bonus', 'credit', 'debit'. Optional.",
	//       "enum": [
	//         "credit",
	//         "debit",
	//         "monthly_minimum",
	//         "other",
	//         "slotting_fee",
	//         "tier_bonus"
	//       ],
	//       "enumDescriptions": [
	//         "A credit increases the publisher's payout amount and decreases the advertiser's invoice amount.",
	//         "A debit reduces the publisher's payout and increases the advertiser's invoice amount.",
	//         "A payment made to Google by an advertiser as a minimum monthly network fee.",
	//         "Catch all. Default if unset",
	//         "A one time payment made from an advertiser to a publisher.",
	//         "A payment from an advertiser to a publisher for the publisher maintaining a high tier level"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "eventDateMax": {
	//       "description": "Filters out all events later than given date. Optional. Defaults to 24 hours after eventMin.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "eventDateMin": {
	//       "description": "Filters out all events earlier than given date. Optional. Defaults to 24 hours from current date/time.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "linkId": {
	//       "description": "Caret(^) delimited list of link IDs. Filters out all events that do not reference one of the given link IDs. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Max number of offers to return in this page. Optional. Defaults to 20.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "100",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "memberId": {
	//       "description": "Caret(^) delimited list of member IDs. Filters out all events that do not reference one of the given member IDs. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "modifyDateMax": {
	//       "description": "Filters out all events modified later than given date. Optional. Defaults to 24 hours after modifyDateMin, if modifyDateMin is explicitly set.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "modifyDateMin": {
	//       "description": "Filters out all events modified earlier than given date. Optional. Defaults to 24 hours before the current modifyDateMax, if modifyDateMax is explicitly set.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "orderId": {
	//       "description": "Caret(^) delimited list of order IDs. Filters out all events that do not reference one of the given order IDs. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The value of 'nextPageToken' from the previous page. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "productCategory": {
	//       "description": "Caret(^) delimited list of product categories. Filters out all events that do not reference a product in one of the given product categories. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "publisherId": {
	//       "description": "Caret(^) delimited list of publisher IDs. Filters out all events that do not reference one of the given publishers IDs. Only used when under advertiser role. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "role": {
	//       "description": "The role of the requester. Valid values: 'advertisers' or 'publishers'.",
	//       "enum": [
	//         "advertisers",
	//         "publishers"
	//       ],
	//       "enumDescriptions": [
	//         "The requester is requesting as an advertiser.",
	//         "The requester is requesting as a publisher."
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "roleId": {
	//       "description": "The ID of the requesting advertiser or publisher.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sku": {
	//       "description": "Caret(^) delimited list of SKUs. Filters out all events that do not reference one of the given SKU. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "status": {
	//       "description": "Filters out all events that do not have the given status. Valid values: 'active', 'canceled'. Optional.",
	//       "enum": [
	//         "active",
	//         "canceled"
	//       ],
	//       "enumDescriptions": [
	//         "Event is currently active.",
	//         "Event is currently canceled."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "type": {
	//       "description": "Filters out all events that are not of the given type. Valid values: 'action', 'transaction', 'charge'. Optional.",
	//       "enum": [
	//         "action",
	//         "charge",
	//         "transaction"
	//       ],
	//       "enumDescriptions": [
	//         "The completion of an application, sign-up, or other process. For example, an action occurs if a user clicks an ad for a credit card and completes an application for that card.",
	//         "A charge event is typically a payment between an advertiser, publisher or Google.",
	//         "A conversion event, typically an e-commerce transaction. Some advertisers use a transaction to record other types of events, such as magazine subscriptions."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "{role}/{roleId}/events",
	//   "response": {
	//     "$ref": "Events"
	//   }
	// }

}

// method id "gan.publishers.get":

type PublishersGetCall struct {
	s      *Service
	role   string
	roleId string
	opt_   map[string]interface{}
}

// Get: Retrieves data about a single advertiser if that the requesting
// advertiser/publisher has access to it. Only advertisers can look up
// publishers. Publishers can request information about themselves by
// omitting the publisherId query parameter.
func (r *PublishersService) Get(role string, roleId string) *PublishersGetCall {
	c := &PublishersGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.role = role
	c.roleId = roleId
	return c
}

// PublisherId sets the optional parameter "publisherId": The ID of the
// publisher to look up.
func (c *PublishersGetCall) PublisherId(publisherId string) *PublishersGetCall {
	c.opt_["publisherId"] = publisherId
	return c
}

func (c *PublishersGetCall) Do() (*Publisher, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["publisherId"]; ok {
		params.Set("publisherId", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/gan/v1beta1/", "{role}/{roleId}/publisher")
	urls = strings.Replace(urls, "{role}", cleanPathString(c.role), 1)
	urls = strings.Replace(urls, "{roleId}", cleanPathString(c.roleId), 1)
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
	ret := new(Publisher)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves data about a single advertiser if that the requesting advertiser/publisher has access to it. Only advertisers can look up publishers. Publishers can request information about themselves by omitting the publisherId query parameter.",
	//   "httpMethod": "GET",
	//   "id": "gan.publishers.get",
	//   "parameterOrder": [
	//     "role",
	//     "roleId"
	//   ],
	//   "parameters": {
	//     "publisherId": {
	//       "description": "The ID of the publisher to look up. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "role": {
	//       "description": "The role of the requester. Valid values: 'advertisers' or 'publishers'.",
	//       "enum": [
	//         "advertisers",
	//         "publishers"
	//       ],
	//       "enumDescriptions": [
	//         "The requester is requesting as an advertiser.",
	//         "The requester is requesting as a publisher."
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "roleId": {
	//       "description": "The ID of the requesting advertiser or publisher.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{role}/{roleId}/publisher",
	//   "response": {
	//     "$ref": "Publisher"
	//   }
	// }

}

// method id "gan.publishers.list":

type PublishersListCall struct {
	s      *Service
	role   string
	roleId string
	opt_   map[string]interface{}
}

// List: Retrieves data about all publishers that the requesting
// advertiser/publisher has access to.
func (r *PublishersService) List(role string, roleId string) *PublishersListCall {
	c := &PublishersListCall{s: r.s, opt_: make(map[string]interface{})}
	c.role = role
	c.roleId = roleId
	return c
}

// MaxResults sets the optional parameter "maxResults": Max number of
// items to return in this page.  Defaults to 20.
func (c *PublishersListCall) MaxResults(maxResults int64) *PublishersListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// MinNinetyDayEpc sets the optional parameter "minNinetyDayEpc":
// Filters out all publishers that have a ninety day EPC average lower
// than the given value (inclusive). Min value: 0.0.
func (c *PublishersListCall) MinNinetyDayEpc(minNinetyDayEpc float64) *PublishersListCall {
	c.opt_["minNinetyDayEpc"] = minNinetyDayEpc
	return c
}

// MinPayoutRank sets the optional parameter "minPayoutRank": A value
// between 1 and 4, where 1 represents the quartile of publishers with
// the lowest ranks and 4 represents the quartile of publishers with the
// highest ranks. Filters out all publishers with a lower rank than the
// given quartile. For example if a 2 was given only publishers with a
// payout rank of 25 or higher would be included.
func (c *PublishersListCall) MinPayoutRank(minPayoutRank int64) *PublishersListCall {
	c.opt_["minPayoutRank"] = minPayoutRank
	return c
}

// MinSevenDayEpc sets the optional parameter "minSevenDayEpc": Filters
// out all publishers that have a seven day EPC average lower than the
// given value (inclusive). Min value 0.0.
func (c *PublishersListCall) MinSevenDayEpc(minSevenDayEpc float64) *PublishersListCall {
	c.opt_["minSevenDayEpc"] = minSevenDayEpc
	return c
}

// PageToken sets the optional parameter "pageToken": The value of
// 'nextPageToken' from the previous page.
func (c *PublishersListCall) PageToken(pageToken string) *PublishersListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// PublisherCategory sets the optional parameter "publisherCategory":
// Caret(^) delimted list of publisher categories. Valid categories:
// (unclassified|community_and_content|shopping_and_promotion|loyalty_and
// _rewards|network|search_specialist|comparison_shopping|email).
// Filters out all publishers not in one of the given advertiser
// categories.
func (c *PublishersListCall) PublisherCategory(publisherCategory string) *PublishersListCall {
	c.opt_["publisherCategory"] = publisherCategory
	return c
}

// RelationshipStatus sets the optional parameter "relationshipStatus":
// Filters out all publishers for which do not have the given
// relationship status with the requesting publisher.
func (c *PublishersListCall) RelationshipStatus(relationshipStatus string) *PublishersListCall {
	c.opt_["relationshipStatus"] = relationshipStatus
	return c
}

func (c *PublishersListCall) Do() (*Publishers, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["minNinetyDayEpc"]; ok {
		params.Set("minNinetyDayEpc", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["minPayoutRank"]; ok {
		params.Set("minPayoutRank", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["minSevenDayEpc"]; ok {
		params.Set("minSevenDayEpc", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["publisherCategory"]; ok {
		params.Set("publisherCategory", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["relationshipStatus"]; ok {
		params.Set("relationshipStatus", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/gan/v1beta1/", "{role}/{roleId}/publishers")
	urls = strings.Replace(urls, "{role}", cleanPathString(c.role), 1)
	urls = strings.Replace(urls, "{roleId}", cleanPathString(c.roleId), 1)
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
	ret := new(Publishers)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves data about all publishers that the requesting advertiser/publisher has access to.",
	//   "httpMethod": "GET",
	//   "id": "gan.publishers.list",
	//   "parameterOrder": [
	//     "role",
	//     "roleId"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "description": "Max number of items to return in this page. Optional. Defaults to 20.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "100",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "minNinetyDayEpc": {
	//       "description": "Filters out all publishers that have a ninety day EPC average lower than the given value (inclusive). Min value: 0.0. Optional.",
	//       "format": "double",
	//       "location": "query",
	//       "type": "number"
	//     },
	//     "minPayoutRank": {
	//       "description": "A value between 1 and 4, where 1 represents the quartile of publishers with the lowest ranks and 4 represents the quartile of publishers with the highest ranks. Filters out all publishers with a lower rank than the given quartile. For example if a 2 was given only publishers with a payout rank of 25 or higher would be included. Optional.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "4",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "minSevenDayEpc": {
	//       "description": "Filters out all publishers that have a seven day EPC average lower than the given value (inclusive). Min value 0.0. Optional.",
	//       "format": "double",
	//       "location": "query",
	//       "type": "number"
	//     },
	//     "pageToken": {
	//       "description": "The value of 'nextPageToken' from the previous page. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "publisherCategory": {
	//       "description": "Caret(^) delimted list of publisher categories. Valid categories: (unclassified|community_and_content|shopping_and_promotion|loyalty_and_rewards|network|search_specialist|comparison_shopping|email). Filters out all publishers not in one of the given advertiser categories. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "relationshipStatus": {
	//       "description": "Filters out all publishers for which do not have the given relationship status with the requesting publisher.",
	//       "enum": [
	//         "approved",
	//         "available",
	//         "deactivated",
	//         "declined",
	//         "pending"
	//       ],
	//       "enumDescriptions": [
	//         "Publishers you've approved to your program.",
	//         "Publishers available for you to recruit.",
	//         "A publisher that you terminated from your program. Publishers also have the ability to remove themselves from your program.",
	//         "A publisher that you did not approve to your program.",
	//         "Publishers that have applied to your program. We recommend reviewing and deciding on pending publishers on a weekly basis."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "role": {
	//       "description": "The role of the requester. Valid values: 'advertisers' or 'publishers'.",
	//       "enum": [
	//         "advertisers",
	//         "publishers"
	//       ],
	//       "enumDescriptions": [
	//         "The requester is requesting as an advertiser.",
	//         "The requester is requesting as a publisher."
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "roleId": {
	//       "description": "The ID of the requesting advertiser or publisher.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{role}/{roleId}/publishers",
	//   "response": {
	//     "$ref": "Publishers"
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
