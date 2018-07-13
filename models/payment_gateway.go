// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PaymentGateway Generic Response
// swagger:model PaymentGateway
type PaymentGateway struct {
	Entity

	// brands
	Brands []string `json:"brands"`

	// card types
	CardTypes []string `json:"cardTypes"`

	// currencies
	Currencies []string `json:"currencies"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// external Id
	ExternalID string `json:"externalId,omitempty"`

	// provider
	Provider string `json:"provider,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PaymentGateway) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Entity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Entity = aO0

	// AO1
	var dataAO1 struct {
		Brands []string `json:"brands,omitempty"`

		CardTypes []string `json:"cardTypes,omitempty"`

		Currencies []string `json:"currencies,omitempty"`

		Enabled bool `json:"enabled,omitempty"`

		ExternalID string `json:"externalId,omitempty"`

		Provider string `json:"provider,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Brands = dataAO1.Brands

	m.CardTypes = dataAO1.CardTypes

	m.Currencies = dataAO1.Currencies

	m.Enabled = dataAO1.Enabled

	m.ExternalID = dataAO1.ExternalID

	m.Provider = dataAO1.Provider

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PaymentGateway) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Entity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Brands []string `json:"brands,omitempty"`

		CardTypes []string `json:"cardTypes,omitempty"`

		Currencies []string `json:"currencies,omitempty"`

		Enabled bool `json:"enabled,omitempty"`

		ExternalID string `json:"externalId,omitempty"`

		Provider string `json:"provider,omitempty"`
	}

	dataAO1.Brands = m.Brands

	dataAO1.CardTypes = m.CardTypes

	dataAO1.Currencies = m.Currencies

	dataAO1.Enabled = m.Enabled

	dataAO1.ExternalID = m.ExternalID

	dataAO1.Provider = m.Provider

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this payment gateway
func (m *PaymentGateway) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Entity
	if err := m.Entity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PaymentGateway) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentGateway) UnmarshalBinary(b []byte) error {
	var res PaymentGateway
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
