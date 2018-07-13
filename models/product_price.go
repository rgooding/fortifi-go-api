// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ProductPrice Generic Response
// swagger:model ProductPrice
type ProductPrice struct {
	Entity

	// currency
	Currency string `json:"currency,omitempty"`

	// Interval in ISO 8601 standard
	Cycle string `json:"cycle,omitempty"`

	// cycle exact
	CycleExact string `json:"cycleExact,omitempty"`

	// cycle term
	CycleTerm int32 `json:"cycleTerm,omitempty"`

	// cycle type
	CycleType CycleTermType `json:"cycleType,omitempty"`

	// price
	Price string `json:"price,omitempty"`

	// product fid
	ProductFid string `json:"productFid,omitempty"`

	// setup fee
	SetupFee string `json:"setupFee,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ProductPrice) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Entity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Entity = aO0

	// AO1
	var dataAO1 struct {
		Currency string `json:"currency,omitempty"`

		Cycle string `json:"cycle,omitempty"`

		CycleExact string `json:"cycleExact,omitempty"`

		CycleTerm int32 `json:"cycleTerm,omitempty"`

		CycleType CycleTermType `json:"cycleType,omitempty"`

		Price string `json:"price,omitempty"`

		ProductFid string `json:"productFid,omitempty"`

		SetupFee string `json:"setupFee,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Currency = dataAO1.Currency

	m.Cycle = dataAO1.Cycle

	m.CycleExact = dataAO1.CycleExact

	m.CycleTerm = dataAO1.CycleTerm

	m.CycleType = dataAO1.CycleType

	m.Price = dataAO1.Price

	m.ProductFid = dataAO1.ProductFid

	m.SetupFee = dataAO1.SetupFee

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ProductPrice) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Entity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Currency string `json:"currency,omitempty"`

		Cycle string `json:"cycle,omitempty"`

		CycleExact string `json:"cycleExact,omitempty"`

		CycleTerm int32 `json:"cycleTerm,omitempty"`

		CycleType CycleTermType `json:"cycleType,omitempty"`

		Price string `json:"price,omitempty"`

		ProductFid string `json:"productFid,omitempty"`

		SetupFee string `json:"setupFee,omitempty"`
	}

	dataAO1.Currency = m.Currency

	dataAO1.Cycle = m.Cycle

	dataAO1.CycleExact = m.CycleExact

	dataAO1.CycleTerm = m.CycleTerm

	dataAO1.CycleType = m.CycleType

	dataAO1.Price = m.Price

	dataAO1.ProductFid = m.ProductFid

	dataAO1.SetupFee = m.SetupFee

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this product price
func (m *ProductPrice) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Entity
	if err := m.Entity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCycleType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductPrice) validateCycleType(formats strfmt.Registry) error {

	if swag.IsZero(m.CycleType) { // not required
		return nil
	}

	if err := m.CycleType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("cycleType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProductPrice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductPrice) UnmarshalBinary(b []byte) error {
	var res ProductPrice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
