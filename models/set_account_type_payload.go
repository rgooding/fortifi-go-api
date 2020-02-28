// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SetAccountTypePayload set account type payload
// swagger:model SetAccountTypePayload
type SetAccountTypePayload struct {

	// account type
	AccountType CustomerAccountType `json:"accountType,omitempty"`
}

// Validate validates this set account type payload
func (m *SetAccountTypePayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SetAccountTypePayload) validateAccountType(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountType) { // not required
		return nil
	}

	if err := m.AccountType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("accountType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SetAccountTypePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetAccountTypePayload) UnmarshalBinary(b []byte) error {
	var res SetAccountTypePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
