// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GetVersionOKBodyAllOf1 get version o k body all of1
// swagger:model getVersionOKBodyAllOf1
type GetVersionOKBodyAllOf1 struct {

	// data
	Data string `json:"data,omitempty"`
}

// Validate validates this get version o k body all of1
func (m *GetVersionOKBodyAllOf1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetVersionOKBodyAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetVersionOKBodyAllOf1) UnmarshalBinary(b []byte) error {
	var res GetVersionOKBodyAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}