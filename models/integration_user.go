// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IntegrationUser Information about the user attempting to integrate
//
// swagger:model IntegrationUser
type IntegrationUser struct {

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// employee fid
	EmployeeFid string `json:"employeeFid,omitempty"`

	// Role Aliases
	Roles []string `json:"roles"`

	// user fid
	UserFid string `json:"userFid,omitempty"`
}

// Validate validates this integration user
func (m *IntegrationUser) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this integration user based on context it is used
func (m *IntegrationUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IntegrationUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntegrationUser) UnmarshalBinary(b []byte) error {
	var res IntegrationUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
