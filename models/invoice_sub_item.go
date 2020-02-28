// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InvoiceSubItem Generic Response
// swagger:model InvoiceSubItem
type InvoiceSubItem struct {
	Entity

	// Time in ISO 8601 standard with optional fractions of a second e.g 2015-12-05T13:11:59.888Z
	// Format: date-time
	FromDate strfmt.DateTime `json:"fromDate,omitempty"`

	// item type
	ItemType string `json:"itemType,omitempty"`

	// per unit amount
	PerUnitAmount float32 `json:"perUnitAmount,omitempty"`

	// quantity
	Quantity int32 `json:"quantity,omitempty"`

	// Time in ISO 8601 standard with optional fractions of a second e.g 2015-12-05T13:11:59.888Z
	// Format: date-time
	ToDate strfmt.DateTime `json:"toDate,omitempty"`

	// total amount
	TotalAmount float32 `json:"totalAmount,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *InvoiceSubItem) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Entity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Entity = aO0

	// AO1
	var dataAO1 struct {
		FromDate strfmt.DateTime `json:"fromDate,omitempty"`

		ItemType string `json:"itemType,omitempty"`

		PerUnitAmount float32 `json:"perUnitAmount,omitempty"`

		Quantity int32 `json:"quantity,omitempty"`

		ToDate strfmt.DateTime `json:"toDate,omitempty"`

		TotalAmount float32 `json:"totalAmount,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.FromDate = dataAO1.FromDate

	m.ItemType = dataAO1.ItemType

	m.PerUnitAmount = dataAO1.PerUnitAmount

	m.Quantity = dataAO1.Quantity

	m.ToDate = dataAO1.ToDate

	m.TotalAmount = dataAO1.TotalAmount

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m InvoiceSubItem) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Entity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		FromDate strfmt.DateTime `json:"fromDate,omitempty"`

		ItemType string `json:"itemType,omitempty"`

		PerUnitAmount float32 `json:"perUnitAmount,omitempty"`

		Quantity int32 `json:"quantity,omitempty"`

		ToDate strfmt.DateTime `json:"toDate,omitempty"`

		TotalAmount float32 `json:"totalAmount,omitempty"`
	}

	dataAO1.FromDate = m.FromDate

	dataAO1.ItemType = m.ItemType

	dataAO1.PerUnitAmount = m.PerUnitAmount

	dataAO1.Quantity = m.Quantity

	dataAO1.ToDate = m.ToDate

	dataAO1.TotalAmount = m.TotalAmount

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this invoice sub item
func (m *InvoiceSubItem) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Entity
	if err := m.Entity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFromDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InvoiceSubItem) validateFromDate(formats strfmt.Registry) error {

	if swag.IsZero(m.FromDate) { // not required
		return nil
	}

	if err := validate.FormatOf("fromDate", "body", "date-time", m.FromDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceSubItem) validateToDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ToDate) { // not required
		return nil
	}

	if err := validate.FormatOf("toDate", "body", "date-time", m.ToDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InvoiceSubItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InvoiceSubItem) UnmarshalBinary(b []byte) error {
	var res InvoiceSubItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
