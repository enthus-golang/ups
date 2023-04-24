package ups

import (
	"encoding/json"
)

type ShipmentResponse struct {
	// Response container for Shipment response.
	Response Response
	// Shipment Results container.
	ShipmentResults ShipmentResults
}

type Response struct {
	// Response status container.
	ResponseStatus ResponseStatus
	// Alert Container. There can be zero to many alert containers with
	// code and description.
	Alerts []Alert `json:"Alert"`
	// Transaction Reference Container.
	// TransactionReference *TransactionReference // buggy
}

func (s *Response) UnmarshalJSON(data []byte) error {
	var v map[string]json.RawMessage
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	if status, ok := v["ResponseStatus"]; ok {
		err := json.Unmarshal(status, &s.ResponseStatus)
		if err != nil {
			return err
		}
	}

	if alert, ok := v["Alert"]; ok {
		if alert[0] == '{' {
			s.Alerts = make([]Alert, 1)

			err := json.Unmarshal(alert, &s.Alerts[0])
			if err != nil {
				return err
			}
		} else {
			err := json.Unmarshal(alert, &s.Alerts)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

type ResponseStatus struct {
	// Identifies the success or failure of the transaction. 1 = Successful
	Code string
	// Describes Response Status Code. Returns text of Success.
	Description string
}

type Alert struct {
	// Warning code returned by the system.
	Code string
	// Warning messages returned by the system.
	Description string
}

type TransactionReference struct {
	// The CustomerContext Information which will be echoed during
	// response.
	CustomerContext string
}

type ShipmentResults struct {
	// Returned UPS shipment ID number.1Z Number of the first package
	// in the shipment.
	ShipmentIdentificationNumber string
	// Returned Package Information.
	// Applicable only for ShipmentResponse and ShipAcceptResponse.
	PackageResults []PackageResults
}

func (s *ShipmentResults) UnmarshalJSON(data []byte) error {
	var v map[string]json.RawMessage
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	if number, ok := v["ShipmentIdentificationNumber"]; ok {
		err := json.Unmarshal(number, &s.ShipmentIdentificationNumber)
		if err != nil {
			return err
		}
	}

	if packageResults, ok := v["PackageResults"]; ok {
		if packageResults[0] == '{' {
			s.PackageResults = make([]PackageResults, 1)

			err := json.Unmarshal(packageResults, &s.PackageResults[0])
			if err != nil {
				return err
			}
		} else {
			err := json.Unmarshal(packageResults, &s.PackageResults)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

type PackageResults struct {
	// Package 1Z number. For Mail Innovations shipments, please use
	// the USPSPICNumber when tracking packages (a non-1Z number
	// Mail Manifest Id is returned).
	// Applicable only for ShipmentResponse and ShipAcceptResponse.
	TrackingNumber string
	// The container for UPS shipping label. Returned for following
	// shipments - Forward shipments, Shipments with PRL returns
	// service, Electronic Return Label or Electronic Import Control Label
	// shipments with SubVersion greater than or equal to 1707.
	// Applicable only for ShipmentResponse and ShipAcceptResponse.
	ShippingLabel *ShippingLabel
}

type ShippingLabel struct {
	// The container image format.
	// Applicable only for ShipmentResponse and ShipAcceptResponse.
	ImageFormat ImageFormat
	// Base 64 encoded graphic image.
	// Applicable only for ShipmentResponse and ShipAcceptResponse.
	GraphicImage string
	// Base 64 encoded graphic image.
	// Applicable only for ShipmentResponse and ShipAcceptResponse
	// for Mail Innovations CN22 Combination Forward Label with more
	// than 3 commodities.
	GraphicImagePart string
}

type ImageFormat struct {
	// Label image code that the labels are generated. Valid values: EPL
	// = EPL2 SPL = SPL ZPL = ZPL GIF = gif images PNG = PNG
	// images. Only EPL, SPL, ZPL and GIF are currently supported.
	// For multi piece COD shipments, the label image format for the first
	// package will always be a GIF for any form of label requested.
	// Applicable only for ShipmentResponse and ShipAcceptResponse.
	Code string
	// Description of the label image format code.
	// Applicable only for ShipmentResponse and ShipAcceptResponse.
	Description string
}

type VoidShipmentResponse struct {
	// Response Container.
	Response Response
	// Container for the Summary Result.
	SummaryResult SummaryResult
}

type SummaryResult struct {
	// Container for the status of the Summary Result
	Status Status
}

type Status struct {
	// Code for the status of the Summary Result
	Code string
	// Description of the status of the Summary Result
	Description string
}

type PackageLevelResult struct {
	// Container for the status of the Summary Result
	Status Status
	// The package's identification number
	TrackingNumber string
}
