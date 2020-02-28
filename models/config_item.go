// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConfigItem Generic Response
// swagger:model ConfigItem
type ConfigItem struct {
	Entity

	// entity fid
	EntityFid string `json:"entityFid,omitempty"`

	// item name
	ItemName string `json:"itemName,omitempty"`

	// section name
	SectionName string `json:"sectionName,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ConfigItem) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Entity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Entity = aO0

	// AO1
	var dataAO1 struct {
		EntityFid string `json:"entityFid,omitempty"`

		ItemName string `json:"itemName,omitempty"`

		SectionName string `json:"sectionName,omitempty"`

		Value string `json:"value,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.EntityFid = dataAO1.EntityFid

	m.ItemName = dataAO1.ItemName

	m.SectionName = dataAO1.SectionName

	m.Value = dataAO1.Value

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ConfigItem) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Entity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		EntityFid string `json:"entityFid,omitempty"`

		ItemName string `json:"itemName,omitempty"`

		SectionName string `json:"sectionName,omitempty"`

		Value string `json:"value,omitempty"`
	}

	dataAO1.EntityFid = m.EntityFid

	dataAO1.ItemName = m.ItemName

	dataAO1.SectionName = m.SectionName

	dataAO1.Value = m.Value

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this config item
func (m *ConfigItem) Validate(formats strfmt.Registry) error {
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
func (m *ConfigItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigItem) UnmarshalBinary(b []byte) error {
	var res ConfigItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
