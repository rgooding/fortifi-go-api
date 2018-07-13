// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ReasonGroup Reason Group
// swagger:model ReasonGroup
type ReasonGroup struct {
	Entity

	// built in
	BuiltIn bool `json:"builtIn,omitempty"`

	// reason count
	ReasonCount int64 `json:"reasonCount,omitempty"`

	// type
	Type ReasonGroupType `json:"type,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ReasonGroup) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Entity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Entity = aO0

	// AO1
	var dataAO1 struct {
		BuiltIn bool `json:"builtIn,omitempty"`

		ReasonCount int64 `json:"reasonCount,omitempty"`

		Type ReasonGroupType `json:"type,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.BuiltIn = dataAO1.BuiltIn

	m.ReasonCount = dataAO1.ReasonCount

	m.Type = dataAO1.Type

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ReasonGroup) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Entity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		BuiltIn bool `json:"builtIn,omitempty"`

		ReasonCount int64 `json:"reasonCount,omitempty"`

		Type ReasonGroupType `json:"type,omitempty"`
	}

	dataAO1.BuiltIn = m.BuiltIn

	dataAO1.ReasonCount = m.ReasonCount

	dataAO1.Type = m.Type

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this reason group
func (m *ReasonGroup) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Entity
	if err := m.Entity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReasonGroup) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReasonGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReasonGroup) UnmarshalBinary(b []byte) error {
	var res ReasonGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
