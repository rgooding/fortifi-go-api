// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CardUpdatePayload Update card details
// swagger:model CardUpdatePayload
type CardUpdatePayload struct {

	// City for the card
	AddressCity string `json:"addressCity,omitempty"`

	// Country of the card
	AddressCountry string `json:"addressCountry,omitempty"`

	// FID of an existing address to use with the card
	AddressFid string `json:"addressFid,omitempty"`

	// Address Line 1 of the card
	AddressLine1 string `json:"addressLine1,omitempty"`

	// Address Line 2 of the card
	AddressLine2 string `json:"addressLine2,omitempty"`

	// Address Line 3 of the card
	AddressLine3 string `json:"addressLine3,omitempty"`

	// Postal/Zip Code of the card
	AddressPostal string `json:"addressPostal,omitempty"`

	// State/County of the card
	AddressState string `json:"addressState,omitempty"`

	// as default
	AsDefault bool `json:"asDefault,omitempty"`

	// Name as appears on card
	CardHolder string `json:"cardHolder,omitempty"`

	// Expiration Month of the card
	ExpiryMonth int32 `json:"expiryMonth,omitempty"`

	// Expiration Year of the card
	ExpiryYear int32 `json:"expiryYear,omitempty"`

	// Issue number of the card
	Issue int32 `json:"issue,omitempty"`

	// Start Month of the card
	StartMonth int32 `json:"startMonth,omitempty"`

	// Start Year of the card
	StartYear int32 `json:"startYear,omitempty"`
}

// Validate validates this card update payload
func (m *CardUpdatePayload) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CardUpdatePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CardUpdatePayload) UnmarshalBinary(b []byte) error {
	var res CardUpdatePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
