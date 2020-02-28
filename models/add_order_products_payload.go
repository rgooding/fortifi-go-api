// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AddOrderProductsPayload add order products payload
// swagger:model AddOrderProductsPayload
type AddOrderProductsPayload struct {
	OrderProductsPayload
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *AddOrderProductsPayload) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 OrderProductsPayload
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.OrderProductsPayload = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m AddOrderProductsPayload) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.OrderProductsPayload)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this add order products payload
func (m *AddOrderProductsPayload) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with OrderProductsPayload
	if err := m.OrderProductsPayload.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AddOrderProductsPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddOrderProductsPayload) UnmarshalBinary(b []byte) error {
	var res AddOrderProductsPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
