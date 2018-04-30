// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetOrdersOrderFidProductsOKBody get orders order fid products o k body
// swagger:model getOrdersOrderFidProductsOKBody
type GetOrdersOrderFidProductsOKBody struct {
	Envelope

	GetOrdersOrderFidProductsOKBodyAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GetOrdersOrderFidProductsOKBody) UnmarshalJSON(raw []byte) error {

	var aO0 Envelope
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Envelope = aO0

	var aO1 GetOrdersOrderFidProductsOKBodyAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.GetOrdersOrderFidProductsOKBodyAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GetOrdersOrderFidProductsOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.GetOrdersOrderFidProductsOKBodyAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get orders order fid products o k body
func (m *GetOrdersOrderFidProductsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.Envelope.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.GetOrdersOrderFidProductsOKBodyAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetOrdersOrderFidProductsOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetOrdersOrderFidProductsOKBody) UnmarshalBinary(b []byte) error {
	var res GetOrdersOrderFidProductsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
