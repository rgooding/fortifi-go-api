// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetOrganisationOKBody get organisation o k body
// swagger:model getOrganisationOKBody
type GetOrganisationOKBody struct {
	Envelope

	GetOrganisationOKBodyAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GetOrganisationOKBody) UnmarshalJSON(raw []byte) error {

	var aO0 Envelope
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Envelope = aO0

	var aO1 GetOrganisationOKBodyAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.GetOrganisationOKBodyAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GetOrganisationOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.GetOrganisationOKBodyAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get organisation o k body
func (m *GetOrganisationOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.Envelope.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.GetOrganisationOKBodyAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetOrganisationOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetOrganisationOKBody) UnmarshalBinary(b []byte) error {
	var res GetOrganisationOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}