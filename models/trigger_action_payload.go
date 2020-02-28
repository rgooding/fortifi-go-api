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

// TriggerActionPayload trigger action payload
// swagger:model TriggerActionPayload
type TriggerActionPayload struct {

	// Your alias for the event to be triggered
	Alias string `json:"alias,omitempty"`

	// meta data
	MetaData MetaData `json:"metaData,omitempty"`

	// Time in ISO 8601 standard with optional fractions of a second e.g 2015-12-05T13:11:59.888Z
	// Format: date-time
	Time strfmt.DateTime `json:"time,omitempty"`

	// If set to true, transactional messenger chains will be triggered
	TriggerMessenger bool `json:"triggerMessenger,omitempty"`
}

// Validate validates this trigger action payload
func (m *TriggerActionPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetaData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TriggerActionPayload) validateMetaData(formats strfmt.Registry) error {

	if swag.IsZero(m.MetaData) { // not required
		return nil
	}

	if err := m.MetaData.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("metaData")
		}
		return err
	}

	return nil
}

func (m *TriggerActionPayload) validateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("time", "body", "date-time", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TriggerActionPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TriggerActionPayload) UnmarshalBinary(b []byte) error {
	var res TriggerActionPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
