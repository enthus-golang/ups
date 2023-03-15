package ups

type ShipmentRequest struct {
	// Container for Shipment Information.
	Shipment Shipment
	// Container used to define the properties required by the user to print
	// and/or display the UPS shipping label.
	// Required for shipment without return service or shipments with PRL
	// return service. Required for Electronic Return Label or Electronic Import
	// Control Label shipments with SubVersion greater than or equal to 1707.
	LabelSpecification *LabelSpecification

	// TODO: implement ReceiptSpecification
}

type Shipment struct {
	Description string `validate:"min=1,max=50"`
	Shipper     Shipper
	ShipTo      ShipTo

	// TODO: implement "AlternateDeliveryAddress"

	ShipFrom *ShipFrom `json:",omitempty"`
	// Payment information container for detailed shipment charges. The two
	// shipment charges that are available for specification are Transportation
	// charges and Duties and Taxes.
	// It is required for non-Ground Freight Pricing shipments only.
	PaymentInformation *PaymentInformation `json:",omitempty"`
	// Required for Ground Freight Pricing Shipments only.
	FRSPaymentInformation *FRSPaymentInformation `json:",omitempty"`

	// TODO: implement FreightShipmentInformation

	GoodsNotInFreeCirculationIndicator string `json:",omitempty"`

	// TODO: implement PromotionalDiscountInformation
	// TODO: implement DGSignatoryInfo
	// TODO: implement ShipmentRatingOptions

	MovementReferenceNumber string           `json:",omitempty"`
	ReferenceNumber         *ReferenceNumber `json:",omitempty"`
	// UPS service type.
	Service Service

	// TODO: implement InvoiceLineTotal

	// Total number of pieces in all pallets in a UPS Worldwide Express
	// Freight Shipment.
	// It is required for UPS Worldwide Express Freight and UPS Worldwide
	// Express Freight Midday Shipment. Valid values are 1 to 99999.
	NumOfPiecesInShipment string `json:",omitempty" validate:"max=5"`
	// USPS endorsement is a Special handling for UPS SurePost shipments
	// delivered by the USPS.
	// Valid values:
	// 1 = Return Service Requested
	// 2 = Forwarding Service Requested
	// 3 = Address Service Requested
	// 4 = Change Service Requested
	// If user does not select a value, UPS system will pass ""Carrier – Leave
	// if No Response”.
	// Note: For International Mail Innovations shipments use No Service
	// Selected. International Mail Innovations shipments are applicable for
	// Priority Mail Innovations and Mail Innovations Economy Mail
	// Innovations services only.
	// Required for Mail Innovations forward and return shipments.
	USPSEndorsement string `json:",omitempty" validate:"max=1"`
	// Indicates single label with both MI label and CN22 form.
	// International CN22 form is required.
	MILabelCN22Indicator string `json:",omitempty"`
	// A component encoded on the barcode of the Mail Innovations label.
	// Valid values: IR = Irregular MA = Machineable SubClass is only
	// required if the customer’s contract have them subclass the package not
	// UPS.
	SubClassification string `json:",omitempty"`
	// Customer assigned identifier for report and billing summarization
	// displays to the right of the Cost Center title.
	// Required for Mail Innovations Return shipments. It is shown on the
	// bottom of the shipping label as reference 2.
	// Cost Center length is alphanumeric with a max length of 30 for Mail
	// Innovations forward shipments.
	// Cost Center length is numeric with a max length of 4 for Mail
	// Innovations Return shipments.
	CostCenter string `json:",omitempty"`
	// Presence/Absence indicator. Presence of this indicator means that the
	// customer is requesting for the CostCenter field to be barcoded at the
	// bottom of the label.
	CostCenterBarcodeIndicator string `json:",omitempty"`
	// Customer-assigned unique piece identifier that returns visibility events.
	// Required only for Mail Innovations forward shipments. Alpha numeric
	// values only. It is shown on the bottom of the shipping label as reference
	// 1.
	PackageID string `json:",omitempty" validate:"max=30"`
	// Presence/Absence indicator. Presence of this indicator means that the
	// customer is requesting for the PackageID field to be barcoded at the
	// bottom of the label.
	PackageIDBarcodeIndicator string `json:",omitempty" validate:"max=30"`
	// Mail classification defined by the USPS.
	// Valid values: 1 = Balloon 2 = Oversize 3 = Not Applicable
	IrregularIndicator string `json:",omitempty" validate:"max=30"`

	// TODO: implement ShipmentIndicationType

	// MIDualReturnShipmentKey is unique key required to process Mail
	// Innovations Dual Return Shipment. The unique identifier (key) would be
	// returned in response of first phase of Mail Innovations Dual Return
	// Shipments. This unique identifier (key) would be part of request for
	// second phase of Mail Innovations Dual Return Shipments. Format: For
	// Package return shipments, the package tracking number is concatenated
	// with the system time (YYYY-MM-DDHH.MM.SS.NNN), followed by
	// service code. For MI Return shipments, the Mail Manifest ID (MMI) is
	// concatenated with the system time.
	// The unique identifier (key) is required to link the package and the Mail
	// Innovations portion of Dual Return shipment. If unique identifier (key) is
	// empty in the request for UPS Mail Innovations Return Service, the
	// request will be treated as the first phase of the Mail Innovations Dual
	// Returns Request. If the MIDualReturnShipmentIndicator is present with
	// empty or null MIDualReturnShipmentKey in UPS Package Return
	// Shipment, the request will be treated as the first phase of Dual MI Return
	// Label Shipment. This field would be ignored if
	// MIDualReturnShipmentIndicator is not present in UPS Package Return
	// Shipment request.
	MIDualReturnShipmentKey string `json:",omitempty" validate:"max=50"`
	// MIDualReturnShipmentIndicator is an indicator to identify a Package
	// Shipment is part of UPS Mail Innovations Dual Label Shipment. Its
	// presence means Package Shipment is part of UPS Mail Innovations
	// Dual Label shipment.
	// If the indicator is present in Package Shipment request, shipment would
	// be considered as part of a Dual Mail Innovations Returns. This indicator
	// is not valid with UPS Mail Innovations Returns Service code.
	MIDualReturnShipmentIndicator string `json:",omitempty"`
	// Presence/Absence Indicator. Any value inside is ignored.
	// RatingMethodRequestedIndicator is an indicator. If present, Billable
	// Weight Calculation method information and Rating Method information
	// would be returned in response.
	RatingMethodRequestedIndicator string `json:",omitempty"`
	// Presence/Absence Indicator. Any value inside is ignored.
	// TaxInformationIndicator is an indicator. If present, any taxes that may
	// be applicable to a shipment would be returned in response. If this
	// indicator is requested with NegotiatedRatesIndicator, Tax related
	// information, if applicable, would be returned only for Negotiated Rates
	// and not for Published Rates. The Tax related information includes any
	// type of Taxes, corresponding Monetary Values, Total Charges with
	// Taxes and disclaimers (if applicable) would be returned in response.
	TaxInformationIndicator string                  `json:",omitempty"`
	ShipmentServiceOptions  *ShipmentServiceOptions `json:",omitempty"`
	// Represents 5 character ISO Locale that allows the user to request
	// Reference Number Code on Label, Label instructions and Receipt
	// instructions (if applicable) in desired language.
	// Locale is specified by the combination of language code and country or
	// territory code - 2 character language code and 2 character country or
	// territory code seperated by an underscore ('_') character.
	// If Locale element is requested along with LabelLinksIndicator, the URL
	// to retrieve Label and Receipts (if applicable) will be returned in the
	// requested Locale. Please note only LabelURL and ReceiptURL (if
	// applicable) will be returned. LocalLanguageLabelURL and
	// LocalLanguageReceiptURL will not be returned if Locale element is
	// present in request.
	// Queen’s English (en_GB) is the default
	Locale string `json:",omitempty" validate:"max=5"`
	// Master Carton ID.
	// If Economy Service (17 or 72): Economy Shipment will be associated
	// with given Master Carton ID.
	// If Non-Economy Service: Master Carton Shipment will be created for
	// given Master Carton ID.
	MasterCartonID string `json:",omitempty" validate:"max=24"`
	// Master Carton Indicator. Presence of the indicator means Master
	// Carton ID will be created and returned to client.
	// To show false, omit the element.
	// To show True:
	// {“MasterCartonIndicator”: “any value”}
	//  Or
	// {“MasterCartonIndicator”: “ ”}
	// MasterCartonIndicator is only valid with Econmoy Shipment (Service
	// Code 17 or 72). Will be ignored if master carton id present
	MasterCartonIndicator string `json:",omitempty"`
	// Shipment Value Threshold Code.
	// 01 = Shipment value is below or equals to threshold value
	// 02 = Shipment value is above threshold value.
	// NA = Not Applicable
	ShipmentValueThresholdCode string `json:",omitempty" validate:"max=2"`
	// Package Information container.
	// For Return Shipments up to and including 20 packages are allowed.
	// US/PR origin return movements are limited to only one package. For
	// Mail Innovations shipments only one package is allowed.
	Packages []Package
	// For SubVersion 2205, user can send up to 7 days in the future with
	// current date as day zero. Format: YYYYMMDD
	ShipmentDate string `json:",omitempty"`
}

type Shipper struct {
	// For forward Shipment 35 characters are
	// accepted, but only 30 characters will be printed on the label.
	Name string `validate:"min=1,max=35"`
	// For forward Shipment 35 characters are
	// accepted, but only 30 characters will be printed on the label.
	// Required if destination is international. Required if Invoice and CO
	// International forms are requested and the ShipFrom address is not
	// present.
	AttentionName string `validate:"max=35"`
	// The CompanyDisplayableName will be displayed in tracking results and
	// notification messages in place of the name associated with the shipper account.
	// The original shipper account name will be displayed for all Return Services and Import
	// Control Shipments.
	// This is available for Shipper accounts enabled by UPS and applies to
	// Forward Shipments.
	CompanyDisplayableName string `json:",omitempty" validate:"max=35"`
	// Conditionally required if EEI form (International forms) is requested and
	// ship From is not mentioned.
	TaxIdentificationNumber string `json:",omitempty" validate:"max=15"`
	Phone                   *Phone `json:",omitempty"`
	// Shipper’s six digit alphanumeric account number. Must be associated
	// with the UserId specified in the AccessRequest. The account must be a
	// valid UPS account number that is active. For US, PR and CA accounts,
	// the account must be either a daily pickup account, an occasional
	// account, or a customer B.I.N account. Drop Shipper accounts are valid
	// for return service shipments only if the account is Trade Direct (TD)
	// enabled. All other accounts must be either a daily pickup account or an
	// occasional account.
	ShipperNumber string `validate:"len=6"`
	FaxNumber     string `json:",omitempty" validate:"max=14"`
	// Must be associated with the UserId specified in the AccessRequest.
	EMailAddress string `json:",omitempty" validate:"max=50"`
	// This address appears on the upper left hand corner of the label.
	// Note:
	// If the ShipFrom container is not present then this address will be used
	// as the ShipFrom address. If this address is used as the ShipFrom the
	// shipment will be rated from this origin address
	Address ShipperAddress
}

type Phone struct {
	// Valid values are 0 - 9. If Shipper country or territory is US, PR, CA, and
	// VI, the layout is: area code, 7 digit PhoneNumber or area code, 7 digit
	// PhoneNumber, 4 digit extension number. For other country or territory,
	// the layout is: CountryCode, area code, 7 digit number. A phone number
	// is required if destination is international.
	Number    string `validate:"min=1,max=15"`
	Extension string `json:",omitempty" validate:"max=4"`
}

type ShipperAddress struct {
	// Up to three occurrences are allowed; only the first is printed on the
	// label. 35 characters are accepted, but for the first occurrence, only 30
	// characters will be printed on the label for return shipments.
	AddressLines []string `json:"AddressLine" validate:"required,max=3,dive,min=1,max=35"`
	// For forward Shipment 30 characters are accepted, but only 15
	// characters will be printed on the label.
	City string `validate:"min=1,max=30"`
	// For forward Shipment 5 characters
	// are accepted, but only 2 characters will be printed on the label.
	// For US, PR and CA accounts, the account must be either a daily pickup
	// account, an occasional account, or a customer B.I.N account.
	StateProvinceCode string `validate:"max=5"`
	PostalCode        string `validate:"max=9"`
	CountryCode       string `validate:"len=2"`
}

type ShipTo struct {
	// All other accounts must be either a daily pickup account or an
	// occasional account.
	Name string `validate:"min=1,max=35"`
	// Required for: UPS Next Day Air® Early service, and when ShipTo
	// country or territory is different than ShipFrom country or territory.
	// Required if Invoice International form is requested.
	AttentionName           string `validate:"max=35"`
	CompanyDisplayableName  string `json:",omitempty" validate:"max=35"`
	TaxIdentificationNumber string `json:",omitempty" validate:"max=15"`
	Phone                   *Phone `json:",omitempty"`
	// Shipper’s six digit alphanumeric account number. Must be associated
	// with the UserId specified in the AccessRequest. The account must be a
	// valid UPS account number that is active. For US, PR and CA accounts,
	// the account must be either a daily pickup account, an occasional
	// account, or a customer B.I.N account. Drop Shipper accounts are valid
	// for return service shipments only if the account is Trade Direct (TD)
	// enabled. All other accounts must be either a daily pickup account or an
	// occasional account.
	ShipperNumber string `json:",omitempty" validate:"len=6"`
	// If ShipTo country or territory is US 10 digits allowed, otherwise 1-15
	// digits allowed.
	FaxNumber    string `json:",omitempty" validate:"max=15"`
	EMailAddress string `json:",omitempty" validate:"max=50"`
	// This address appears on the upper left hand corner of the label.
	// Note:
	// If the ShipFrom container is not present then this address will be used
	// as the ShipFrom address. If this address is used as the ShipFrom the
	// shipment will be rated from this origin address
	Address ShipToAddress
	// Location ID must be alphanumeric characters. All letters must be
	// capitalized.
	LocationID string `json:",omitempty" validate:"max=10"`
}

type ShipToAddress struct {
	// Max occurrence: 3 Only first two Address Lines will be printed on the
	// label.
	AddressLines []string `json:"AddressLine" validate:"required,max=3,dive,min=1,max=35"`
	// 30 characters are accepted, but only 15 characters
	// will be printed on the label.
	City string `validate:"min=1,max=30"`
	// If destination is US or CA, then the value must be a valid US State/
	// Canadian Province code. If the country or territory is Ireland, the
	// StateProvinceCode will contain the county.
	StateProvinceCode string `json:",omitempty" validate:"max=5"`
	// If the ShipTo country or territory is US or Puerto Rico, 5 or 9 digits are
	// required. If the ShipTo country or territory is CA, then the postal code is
	// required and must be 6 alphanumeric characters whose format is
	// A#A#A# where A is an uppercase letter and # is a digit. Otherwise
	// optional. For all other countries or territories the postal code is optional
	// and must be no more than 9 alphanumeric characters long.
	PostalCode string `validate:"max=9"`
	// Must be a valid UPS Billing country or territory code. For Return
	// Shipment the country or territory code must meet the following
	// conditions: 1) At least two of the following country or territory codes are
	// the same: ShipTo, ShipFrom, and Shipper. 2) None of the following
	// country or territory codes are the same and are a member of the EU:
	// ShipTo, ShipFrom, and Shipper. 3) If any of the two following country or
	// territory codes: ShipTo/ ShipFrom/ Shipper are members in EU
	// otherwise check if the shipper has Third Country or territory Contract.
	CountryCode string `validate:"len=2"`
	// This field is a flag to indicate if the receiver is a residential location.
	// True if ResidentialAddressIndicator tag exists.
	ResidentialAddressIndicator string `json:",omitempty"`
}

type ShipFrom struct {
	// 35 characters are accepted, but for return Shipment only 30 characters will be printed on
	// the label.
	Name string `validate:"min=1,max=35"`
	// 35 characters are accepted, but for
	// return Shipment only 30 characters will be printed on the label.
	// Required if ShipFrom tag is in the request and Invoice or CO
	// International forms is requested. If not present, will default to the
	// Shipper Attention Name.
	AttentionName          string `validate:"max=35"`
	CompanyDisplayableName string `validate:"max=35"`
	// Conditionally required if EEI form (International forms) is requested.
	// Applies to EEI Form only.
	TaxIdentificationNumber string `validate:"max=15"`
	// Applies to EEI form only
	TaxIDType *TaxIDType `json:",omitempty"`
	Phone     *Phone     `json:",omitempty"`
	// Shipper’s six digit alphanumeric account number. Must be associated
	// with the UserId specified in the AccessRequest. The account must be a
	// valid UPS account number that is active. For US, PR and CA accounts,
	// the account must be either a daily pickup account, an occasional
	// account, or a customer B.I.N account. Drop Shipper accounts are valid
	// for return service shipments only if the account is Trade Direct (TD)
	// enabled. All other accounts must be either a daily pickup account or an
	// occasional account.
	ShipperNumber string `validate:"len=6"`
	// If ShipTo country or territory is US 10 digits allowed, otherwise 1-15
	// digits allowed.
	FaxNumber string `validate:"max=15"`
	// This address appears on the upper left hand corner of the label.
	// Note:
	// If the ShipFrom container is not present then this address will be used
	// as the ShipFrom address. If this address is used as the ShipFrom the
	// shipment will be rated from this origin address
	Address ShipToAddress
	// Location ID must be alphanumeric characters. All letters must be
	// capitalized.
	LocationID string `validate:"max=10"`
}

type TaxIDType struct {
	// Valid values: EIN, DNS, and FGN. Applies to EEI form only.
	Code string
}

type PaymentInformation struct {
	ShipmentCharge ShipmentCharge
}

type ShipmentCharge struct {
	// Valid values:
	// 01 = Transportation
	// 02 = Duties and Taxes
	// 03 = Broker of Choice
	// A shipment charge type of 01 = Transportation is required. A shipment
	// charge type of 02 = Duties and Taxes is not required; however, this
	// charge type is invalid for Qualified Domestic Shipments. A Qualified
	// Domestic Shipment is any shipment in which one of the following
	// applies: 1) The origin and destination country or territory is the same. 2)
	// US to PR shipment. 3) PR to US shipment. 4) The origin and
	// destination country or territory are both European Union countries and
	// territories and the GoodsNotInFreeCirculation indicator is not present.
	// 5) The origin and destination IATA code is the same.
	Type        string       `validate:"len=2"`
	BillShipper *BillShipper `json:",omitempty"`
	// TODO: implement BillReceiver
	// TODO: implement BillThirdParty
}

type BillShipper struct {
	// UPS account number.
	// Must be the same UPS account number as the one provided in
	// Shipper/ShipperNumber. Either this element or one of the sibling
	// elements CreditCard or AlternatePaymentMethod must be provided, but
	// all of them may not be provided.
	AccountNumber string `validate:"len=6"`
	// Credit card information container.
	// Required if neither of the following is present: {“ShipmentRequest”:
	// {“Shipment”: {“PaymentInformation”: {“ShipmentCharge”:
	// {“BillShipper/AccountNumber or {“ShipmentRequest”: {“Shipment”:
	// {“PaymentInformation”: {“ShipmentCharge”:
	// {“BillShipper/AlternatePaymentMethod. Credit card payment is valid for
	// shipments without return service only.
	CreditCard *CreditCard `json:",omitempty"`
}

type CreditCard struct {
	// Valid values:
	// 01 = American Express
	// 03 = Discover
	// 04 = MasterCard
	// 05 = Optima
	// 06 = VISA
	// 07 = Bravo
	// 08 = Diners Club
	// 13=Dankort
	// 14=Hipercard
	// 15=JCB
	// 17=Postepay
	// 18=UnionPay/ExpressPay
	// 19=Visa Electron
	// 20=VPAY
	// 21=Carte Bleue
	Type string `validate:"len=2"`
	// Credit Card number.
	Number string `validate:"min=9,max=16"`
	// Format is MMYYYY where MM is the 2 digit month and YYYY is the 4
	// digit year.
	// Valid month values are 01-12 and valid year values are Present Year –
	// (Present Year + 10 years)
	ExpirationDate string `validate:"len=6"`
	// Three or four digits that can be found either on top of credit card
	// number or on the back of credit card. Number of digits varies for
	// different type of credit card.
	// Valid values are 3 or 4 digits. It is required to provide the security code
	// if credit card information is provided and when the ShipFrom countries
	// or territories are other than the below mentioned countries or territories.
	// Argentina, Bahamas, Costa Rica, Dominican Republic, Guatemala,
	// Panama, Puerto Rico and Russia.
	SecurityCode string `validate:"min=3,max=4"`
}

type FRSPaymentInformation struct {
	Type FRSPaymentInformationType
	// The UPS account number.
	// If the Ground Freight Pricing indicator and
	// FreightShipmentInformation/DensityEligibleIndicator is present in the
	// request, this account number must be validated to check if it is Ground
	// Freight Pricing Density Based Rating enabled.
	AccountNumber string                        `validate:"len=6"`
	Address       *FRSPaymentInformationAddress `json:",omitempty"`
}

type FRSPaymentInformationType struct {
	// Valid codes: 01=Prepaid 02=FreightCollect 03=ThirdParty
	Code string `validate:"len=2"`
	// Specifies the description for Ground Freight Pricing payment type.
	Description string `json:",omitempty" validate:"max=50"`
}

type FRSPaymentInformationAddress struct {
	PostalCode  string `json:",omitempty" validate:"max=9"`
	CountryCode string `validate:"len=2"`
}

type ReferenceNumber struct {
	// If the indicator is present then the reference number’s value will be bar
	// coded on the label.
	// This is an empty tag, any value inside is ignored. Only one shipmentlevel or package-level reference number can be bar coded per
	// shipment. In order to barcode a reference number, its value must be no
	// longer than 14 alphanumeric characters or 24 numeric characters and
	// cannot contain spaces.
	BarCodeIndicator string `json:",omitempty"`
	// Shipment Reference number type code. The code specifies the
	// Reference name. Refer to the Reference Number Code table.
	// Valid if the origin/destination pair is not US/US or PR/PR and character
	// should be alpha-numeric.
	Code string `json:",omitempty" validate:"max=2"`
	// Customer supplied reference number.
	// Valid if the origin/destination pair is not US/US or PR/PR
	Value string `json:",omitempty" validate:"max=35"`
}

type Service struct {
	// Valid values:
	// 01 = Next Day Air
	// 02 = 2nd Day Air
	// 03 = Ground
	// 07 = Express
	// 08 = Expedited
	// 11 = UPS Standard
	// 12 = 3 Day Select
	// 13 = Next Day Air Saver
	// 14 = UPS Next Day Air® Early
	// 17 = UPS Worldwide Economy DDU
	// 54 = Express Plus
	// 59 = 2nd Day Air A.M.
	// 65 = UPS Saver
	// M2 = First Class Mail
	// M3 = Priority Mail
	// M4 = Expedited MaiI Innovations
	// M5 = Priority Mail Innovations
	// M6 = Economy Mail Innovations
	// M7 = MaiI Innovations (MI) Returns 70 = UPS Access Point™ Economy
	// 71 = UPS Worldwide Express Freight Midday
	// 72 = UPS Worldwide Economy
	// 74 = UPS Express®12:00 82 = UPS Today Standard 83 = UPS Today
	// Dedicated Courier
	// 84 = UPS Today Intercity 85 = UPS Today Express 86 = UPS Today
	// Express Saver
	// 96 = UPS Worldwide Express Freight.
	// Note: Only service code 03 is used for Ground Freight Pricing
	// shipments
	// The following Services are not available to return shipment: 13, 59, 82,
	// 83, 84, 85, 86
	Code string `json:",omitempty" validate:"len=2"`
	// Description of the service code. Examples are Next Day Air, Worldwide
	// Express, and Ground.
	Description string `json:",omitempty" validate:"max=35"`
}

type ShipmentServiceOptions struct {
	// Saturday delivery indicator. The presence indicates Saturday delivery is
	// requested and the absence indicates Saturday delivery is not
	// requested.
	SaturdayDeliveryIndicator string `json:",omitempty"`
	// Saturday pickup indicator. The presence indicates Saturday pickup is
	// requested and the absence indicates Saturday pickup is not requested.
	SaturdayPickupIndicator string `json:",omitempty"`

	// TODO: implement COD
	// TODO: implement AccessPointCOD

	// Presence/Absence Indicator. Any value inside is ignored.
	// DeliverToAddresseeOnlyIndicator is shipper specified restriction that
	// requires the addressee to be the one who takes final delivery of the
	// "Hold For PickUp at UPS Access Point" package. Presence of indicator
	// means shipper restriction will apply to the shipment.
	// Only valid for Shipment Indication type "01 - Hold For PickUp at UPS
	// Access Point".
	DeliverToAddresseeOnlyIndicator string `json:",omitempty"`
	// Presence/Absence Indicator. Any value inside is ignored. Direct
	// Delivery Only (DDO) accessorial in a request would ensure that delivery
	// is made only to the ship to address on the shipping label.
	// This accessorial is not valid with Shipment Indication Type "01 - Hold
	// For Pickup At UPS Access Point" and "02 - UPS Access Point™
	// Delivery".
	DirectDeliveryOnlyIndicator string `json:",omitempty"`
	// Container for the Quantum View Notification (QVN) is valid for all
	// shipments including Return service, Import Control and Returns
	// Flexible Access. Valid return service types are: ERL, PRL, RS1, or
	// RS3.
	// The shipment level notification is valid for forward and return
	// international shipments as well as for domestic shipments (for US and
	// PR).
	Notifications []Notification `json:"Notification,omitempty" validate:"max=3,dive"`
}

type Notification struct {
	// The type of notification requested. Note: - QVN Exception notification
	// and return notification are not applicable to GFP. - QV In-transit and
	// Return Notifications are only valid for ImportControl and Return
	// shipment. - QV In-transit Notification is allowed for return shipments
	// only. - QV Ship Notification is allowed for forward moving shipments
	// only.
	// Valid values: 5 - QV In-transit Notification 6 - QV Ship Notification 7 -
	// QV Exception Notification 8 - QV Delivery Notification 2 - Return
	// Notification or Label Creation Notification 012 - Alternate Delivery
	// Location Notification 013 - UAP Shipper Notification.
	NotificationCode string `validate:"min=1,max=3"`
	EMail            EMail
}

type EMail struct {
	// Email address where the notification is sent.
	// Up to five email addresses are allowed for each type of Quantum View
	// TM shipment notification. Up to two email address for return notification
	EMailAddress string `validate:"min=1,max=50"`
	// The address where an undeliverable eMail message is sent if the eMail
	// with the notification is undeliverable.
	// There can be only one UndeliverableEMailAddress for each type of
	// Quantum View Shipment Notifications.
	UndeliverableEMailAddress string `json:",omitempty" validate:"max=50"`
	// The e-mail address specifies the Reply To E-mail address. The "From"
	// field of the message header contains pkginfo@ups.com.
	// Valid for Return Notification only.
	FromEMailAddress string `json:",omitempty" validate:"max=50"`
	// The name the email will appear to be from. Defaults to the Shipper Name.
	// The FromName must occur only once for each shipment with Quantum
	// View Shipment Notifications.
	FromName string `json:",omitempty" validate:"max=35"`
	// User defined text that will be included in the eMail.
	// The Memo must occur only once for each shipment with Quantum View
	// Shipment Notifications.
	Memo string `json:",omitempty" validate:"max=150"`

	// TODO: implement VoiceMessage
	// TODO: implement TextMessage
	// TODO: implement Locale
	// TODO: implement LabelDelivery
	// TODO: implement InternationalForms
	// TODO: implement DeliveryConfirmation

	// The flag indicates the ReturnOfDocument accessorial has been
	// requested.
	// Valid for Poland to Poland forward shipment only.
	ReturnOfDocumentIndicator string `json:",omitempty"`
	// Indicates that the Shipment is an ImportControl shipment.
	ImportControlIndicator string `json:",omitempty"`

	// TODO: implement LabelMethod

	// CommercialInvoiceRemovalIndicator allows a shipper to dictate UPS to
	// remove the Commercial Invoice from the user's shipment before the
	// shipment is delivered to the ultimate consignee.
	CommercialInvoiceRemovalIndicator string `json:",omitempty"`

	// TODO: implement PreAlertNotification

	// Exchange forward indicator presence at shipment level is required to
	// create exchange forward Shipments.
	// In the label routing Instruction text will be defaulted to "EXCHANGELIKE ITEM ONLY".
	ExchangeForwardIndicator string `json:",omitempty"`
	// Hold For Pickup indicator. The empty tag means indicator is present.
	// This accessorial is only valid for UPS Worldwide Express Freight and
	// UPS Worldwide Express Freight Midday Shipment.
	HoldForPickupIndicator string `json:",omitempty"`
	// Drop off At UPS Facility indicator. The empty tag means indicator is
	// present.
	// This accessorial is only valid for UPS Worldwide Express Freight and
	// UPS Worldwide Express Freight Midday Shipment.
	DropoffAtUPSFacilityIndicator string `json:",omitempty"`
	// Lift Gate For Pick Up indicator. The empty tag means indicator is
	// present.
	// Lift Gate for Pickup is not allowed with Drop Off At UPS Facility for a
	// UPS Worldwide Express Freight and UPS Worldwide Express Freight
	// Midday shipment. When both Hold for Pickup and Drop Off At Facility
	// are selected, neither of the Lift Gate accessorial (Pick Up or Delivery)
	// are allowed for a UPS Worldwide Express Freight and UPS Worldwide
	// Express Freight Midday shipment. This accessorial is only valid for UPS
	// Worldwide Express Freight and UPS Worldwide Express Freight
	// Midday Shipment.
	LiftGateForPickUpIndicator string `json:",omitempty"`
	// Lift Gate For Delivery indicator. The empty tag means indicator is
	// present.
	// Lift Gate for Delivery is not allowed with Hold For Pickup for a UPS
	// Worldwide Express Freight and UPS Worldwide Express Freight
	// Midday shipment. When both Hold for Pickup and Drop Off At UPS
	// Facility are selected, neither of the Lift Gate accessorial (Pick Up or
	// Delivery) are allowed for a UPS Worldwide Express Freight and UPS
	// Worldwide Express Freight Midday shipment. This accessorial is only
	// valid for UPS Worldwide Express Freight and UPS Worldwide Express
	// Freight Midday Shipment.
	LiftGateForDeliveryIndicator string `json:",omitempty"`
	// The presence of the tag SDLShipmentIndicator indicates Shipment is
	// SDL. SDLShipmentIndicator presence means EEI form/ EEI Filing
	// option required.
	SDLShipmentIndicator string `json:",omitempty"`
	// Package Release code allows the consignee or claimant to pick-up a
	// package at a UPS Access Point™. The shipper must provide the
	// Package Release Code to the consignee so that they can provide the
	// code to the UPS Access Point personnel as another item for
	// authentication before the package is released to them. Package
	// Release Code is only valid with ShipmentIndicationType 01 - Hold for
	// Pickup at UPS Access Point™.
	// The release code must be between length 4 and 6 and only contain
	// numbers.
	EPRAReleaseCode string `json:",omitempty" validate:"max=6"`

	// TODO: implement RestrictedArticles
}

type Package struct {
	// Merchandise description of package.
	// Required for shipment with return service.
	Description string `json:",omitempty" validate:"max=35"`
	// Description of articles & special marks. Applicable for Air Freight only
	PalletDescription string `json:",omitempty" validate:"max=150"`
	// Number of Pieces. Applicable for Air Freight only
	NumOfPieces string `json:",omitempty" validate:"max=5"`
	// Unit price of the commodity. Applicable for Air Freight only
	// Limit to 2 digit after the decimal. The maximum length of the field is 12
	// including ‘.’ and can hold up to 2 decimal place. (e.g. 999999999.99)
	UnitPrice string `json:",omitempty" validate:"max=12"`
	// Packaging container.
	// Container for Packaging Type.
	Packaging Packaging
	// Dimensions information container. Note: Currently dimensions are not
	// applicable to Ground Freight Pricing.
	// Length + 2*(Width + Height) must be less than or equal to 165 IN or
	// 330 CM.
	Dimensions Dimensions
	// Dimensional weight of shipment. Please visit ups.com for rules on
	// calculating. There is one implied decimal place (e.g. 115 = 11.5).
	// If dimensions are provided, dimensional weight is ignored. For
	// US/PR/CA shipments, dimensional weight is ignored
	DimWeight *DimWeight `json:",omitempty"`
	// Container to hold package weight information.
	// Package weight is a required for Ground Freight Pricing shipments.
	PackageWeight *PackageWeight `json:",omitempty"`
	// Presence of the indicator mentions that the package is Large Package.
	// This is an empty tag, any value inside is ignored.
	LargePackageIndicator string `json:",omitempty"`

	// TODO: implement ReferenceNumber

	// Additional Handling Required. The presence indicates additional
	// handling is required, the absence indicates no additional handling is
	// required. Additional Handling indicator indicates it’s a non-corrugated
	// package.
	AdditionalHandlingIndicator string `json:",omitempty"`

	// TODO: implement UPSPremier

	// Presence/Absence Indicator. Any value is ignored. If present, indicates
	// that the package is over size.
	// Applicable for UPS Worldwide Economy DDU service.
	OversizeIndicator string `json:",omitempty"`
	// Presence/Absence Indicator. Any value is ignored. If present, indicates
	// that the package is qualified for minimum billable weight.
	// Applicable for UPS Worldwide Economy DDU service.
	MinimumBillableWeightIndicator string `json:",omitempty"`

	// TODO: implement PackageServiceOptions
	// TODO: implement Commodity
	// TODO: implement HazMatPackageInformation
	// TODO: implement SimpleRate
}

type Packaging struct {
	// Package types. Values are: 01 = UPS Letter
	// 02 = Customer Supplied Package
	// 03 = Tube 04 = PAK
	// 21 = UPS Express Box
	// 24 = UPS 25KG Box
	// 25 = UPS 10KG Box
	// 30 = Pallet
	// 2a = Small Express Box
	// 2b = Medium Express Box 2c = Large Express Box
	// 56 = Flats
	// 57 = Parcels
	// 58 = BPM
	// 59 = First Class
	// 60 = Priority
	// 61 = Machineables
	// 62 = Irregulars
	// 63 = Parcel Post
	// 64 = BPM Parcel
	// 65 = Media Mail
	// 66 = BPM Flat
	// 67 = Standard Flat.
	// Note: Only packaging type code 02 is applicable to Ground Freight
	// Pricing.
	// Package type 24, or 25 is only allowed for shipment without return
	// service. Packaging type must be valid for all the following: ShipTo
	// country or territory, ShipFrom country or territory, a shipment going
	// from ShipTo country or territory to ShipFrom country or territory, all
	// Accessorials at both the shipment and package level, and the shipment
	// service type. UPS will not accept raw wood pallets and please refer the
	// UPS packaging guidelines for pallets on UPS.com.
	Code string `validate:"len=2"`
	// Description of packaging type. Examples are letter, customer supplied,
	// express box.
	Description string `json:",omitempty" validate:"max=35"`
}

type Dimensions struct {
	// UnitOfMeasurement container for dimensions.
	UnitOfMeasurement DimensionsUnitOfMeasurement
	// Length must be the longest dimension of the container.
	// Valid values are 0 to 108 IN and 0 to 274 CM.
	Length string `validate:"min=1,max=3"`
	Width  string `validate:"min=1,max=3"`
	Height string `validate:"min=1,max=3"`
}

type DimensionsUnitOfMeasurement struct {
	// Package dimensions measurement code. Valid codes: IN = Inches CM
	// = Centimeters 00 = Metric Units Of Measurement 01 = English Units of
	// Measurement
	// The unit of measurement must be valid for the Shipper country or
	// territory.
	Code string `validate:"len=2"`
	// Description of the package dimensions measurement units.
	Description string `json:",omitempty" validate:"max=35"`
}

type DimWeight struct {
	UnitOfMeasurement DimWeightUnitOfMeasurement
	Weight            string `validate:"len=6"`
}

type DimWeightUnitOfMeasurement struct {
	// Code representing the unit of measure associated with the package
	// weight.
	// Valid values: LBS = Pounds (default) KGS = Kilograms
	Code string `validate:"len=3"`
	// Text description of the code representing the unit of measure
	// associated with the package weight.
	// Length and value are not validated.
	Description string `json:",omitempty" validate:"max=35"`
}

type PackageWeight struct {
	// Container to hold UnitOfMeasurement information for package weight.
	UnitOfMeasurement PackageWeightUnitOfMeasurement
	// Packages weight. Weight accepted for letters/envelopes.
	// Only average package weight is required for Ground Freight Pricing
	// Shipment.
	Weight string `validate:"max=5"`
}

type PackageWeightUnitOfMeasurement struct {
	// Package weight unit of measurement code. Valid values:
	// LBS = Pounds
	// KGS = Kilograms
	// OZS = Ounces
	// Unit of Measurement "OZS" is the only valid UOM for some of the Mail
	// Innovations Forward Shipments and Worldwide Economy DDU
	// Shipments..
	// Please refer to Appendix for more details regarding the valid
	// combination of Mail Innovation Forward Shipment services, Package
	// Type and Unit of Measurement.
	Code string `validate:"len=2"`
	// Description of the unit of measurement for package weight.
	Description string `json:",omitempty" validate:"max=35"`
}

type LabelSpecification struct {
	// LabelImageFormat Container.
	LabelImageFormat LabelImageFormat
	// Browser HTTPUserAgent String. This is the preferred way of identifying
	// GIF image type to be generated.
	// Required if {“ShipmentRequest”:
	// {LabelSpecificationLabelSpecification/LabelImageFormat/Code = Gif.
	// Default to Mozilla/4.5 if this field is missing or has invalid value.
	HTTPUserAgent string `json:",omitempty" validate:"max=64"`
	// Container for the EPL2, ZPL or SPL label size.
	// Valid for EPL2, ZPL and SPL Labels.
	LabelStockSize LabelStockSize
	// Routing Instruction Container.
	Instruction *Instruction `json:",omitempty"`
	// Language character set expected on label. Valid values: dan = Danish
	// (Latin-1) nld = Dutch (Latin-1) fin = Finnish (Latin-1) fra = French (Latin1) deu = German (Latin-1) itl = Italian (Latin-1) nor = Norwegian (Latin1) pol = Polish (Latin-2) por = Portuguese (Latin-1) spa = Spanish
	// (Latin-1) swe = Swedish (Latin-1) ces = Czech (Latin-2) hun =
	// Hungarian (Latin-2) slk = Slovak (Latin-2) rus = Russian (Cyrillic) tur =
	// Turkish (Latin-5) ron = Romanian (Latin-2) bul = Bulgarian (Latin-2) est
	// = Estonian (Latin-2) ell = Greek (Latin-2) lav = Latvian (Latin-2) lit =
	// Lithuanian (Latin-2) eng = English (Latin-1)
	// Default is English (Latin-1).
	CharacterSet string `json:",omitempty"`
}

type LabelImageFormat struct {
	// Label print method code determines the format in which Labels are to
	// be generated. For EPL2 formatted Labels use EPL, for PNG formatted
	// Labels use PNG and for SPL formatted Labels use SPL, for ZPL
	// formatted Labels use ZPL and for image formats use GIF.
	// For shipments without return service the valid value is GIF, PNG, ZPL,
	// EPL and SPL. For shipments with PRL return service, the valid values
	// are EPL, ZPL, SPL, PNG and GIF.
	// For UPS Premier Silver shipment only ZPL is supported.
	Code string `validate:"min=1,max=4"`
	// Description of the label image format code.
	Description string `json:",omitempty" validate:"max=35"`
}

type LabelStockSize struct {
	// Height of the label image. For IN, use whole inches.
	// For EPL2, ZPL and SPL Labels. Only valid values are 6 or 8. Note:
	// Label Image will only scale up to 4 X 6, even when requesting 4 X 8.
	Height string `validate:"min=1,max=3"`
	// Width of the label image. For IN, use whole inches.
	// For EPL2, ZPL and SPL Labels. Valid value is 4. Note: Label Image will
	// only scale up to 4 X 6, even when requesting 4 X 8.
	Width string `validate:"min=1,max=3"`
}

type Instruction struct {
	// For Exchange Forward Shipment, by default Label will have Exchange
	// Routing instruction Text as EXCHANGE-LIKE ITEM ONLY. If code
	// value is:
	// 01- EXCHANGE-LIKE ITEM ONLY,
	// 02- 02- EXCHANGE-DRIVER INSTRUCTIONS INSIDE.
	Code string `validate:"len=2"`
	// Description of the label Instruction code.
	Description string `json:",omitempty" validate:"max=35"`
}
