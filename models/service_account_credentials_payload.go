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

// ServiceAccountCredentialsPayload Service credentials used to verify ownership of a user account
// swagger:model ServiceAccountCredentialsPayload
type ServiceAccountCredentialsPayload struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// key
	// Required: true
	Key *string `json:"key"`
}

// Validate validates this service account credentials payload
func (m *ServiceAccountCredentialsPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceAccountCredentialsPayload) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceAccountCredentialsPayload) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceAccountCredentialsPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceAccountCredentialsPayload) UnmarshalBinary(b []byte) error {
	var res ServiceAccountCredentialsPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
