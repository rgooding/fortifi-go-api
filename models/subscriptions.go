// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Subscriptions Generic Response
// swagger:model Subscriptions
type Subscriptions struct {
	Pagination

	SubscriptionsAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Subscriptions) UnmarshalJSON(raw []byte) error {

	var aO0 Pagination
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Pagination = aO0

	var aO1 SubscriptionsAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.SubscriptionsAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Subscriptions) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.Pagination)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.SubscriptionsAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this subscriptions
func (m *Subscriptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.Pagination.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.SubscriptionsAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Subscriptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Subscriptions) UnmarshalBinary(b []byte) error {
	var res Subscriptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
