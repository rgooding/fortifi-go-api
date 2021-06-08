// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InvoiceCreditNote Generic Response
//
// swagger:model InvoiceCreditNote
type InvoiceCreditNote struct {
	Entity

	// amount
	Amount float32 `json:"amount,omitempty"`

	// charge request fid
	ChargeRequestFid string `json:"chargeRequestFid,omitempty"`

	// Time in ISO 8601 standard with optional fractions of a second e.g 2015-12-05T13:11:59.888Z
	// Format: date-time
	CreditDate strfmt.DateTime `json:"creditDate,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`

	// invoice fid
	InvoiceFid string `json:"invoiceFid,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *InvoiceCreditNote) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Entity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Entity = aO0

	// AO1
	var dataAO1 struct {
		Amount float32 `json:"amount,omitempty"`

		ChargeRequestFid string `json:"chargeRequestFid,omitempty"`

		CreditDate strfmt.DateTime `json:"creditDate,omitempty"`

		Currency string `json:"currency,omitempty"`

		InvoiceFid string `json:"invoiceFid,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Amount = dataAO1.Amount

	m.ChargeRequestFid = dataAO1.ChargeRequestFid

	m.CreditDate = dataAO1.CreditDate

	m.Currency = dataAO1.Currency

	m.InvoiceFid = dataAO1.InvoiceFid

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m InvoiceCreditNote) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Entity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Amount float32 `json:"amount,omitempty"`

		ChargeRequestFid string `json:"chargeRequestFid,omitempty"`

		CreditDate strfmt.DateTime `json:"creditDate,omitempty"`

		Currency string `json:"currency,omitempty"`

		InvoiceFid string `json:"invoiceFid,omitempty"`
	}

	dataAO1.Amount = m.Amount

	dataAO1.ChargeRequestFid = m.ChargeRequestFid

	dataAO1.CreditDate = m.CreditDate

	dataAO1.Currency = m.Currency

	dataAO1.InvoiceFid = m.InvoiceFid

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this invoice credit note
func (m *InvoiceCreditNote) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Entity
	if err := m.Entity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreditDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InvoiceCreditNote) validateCreditDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreditDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creditDate", "body", "date-time", m.CreditDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this invoice credit note based on the context it is used
func (m *InvoiceCreditNote) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Entity
	if err := m.Entity.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *InvoiceCreditNote) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InvoiceCreditNote) UnmarshalBinary(b []byte) error {
	var res InvoiceCreditNote
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
