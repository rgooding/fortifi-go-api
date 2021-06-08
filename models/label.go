// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Label Label
//
// swagger:model Label
type Label struct {
	Fid

	// Label
	Label string `json:"label,omitempty"`

	// Value
	Value string `json:"value,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Label) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Fid
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Fid = aO0

	// AO1
	var dataAO1 struct {
		Label string `json:"label,omitempty"`

		Value string `json:"value,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Label = dataAO1.Label

	m.Value = dataAO1.Value

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Label) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Fid)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Label string `json:"label,omitempty"`

		Value string `json:"value,omitempty"`
	}

	dataAO1.Label = m.Label

	dataAO1.Value = m.Value

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this label
func (m *Label) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Fid
	if err := m.Fid.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this label based on the context it is used
func (m *Label) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Fid
	if err := m.Fid.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Label) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Label) UnmarshalBinary(b []byte) error {
	var res Label
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
