// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ReasonAllOf1 reason all of1
// swagger:model reasonAllOf1
type ReasonAllOf1 struct {

	// FID for the reason group
	GroupFid string `json:"groupFid,omitempty"`

	// used count
	UsedCount int64 `json:"usedCount,omitempty"`
}

// Validate validates this reason all of1
func (m *ReasonAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ReasonAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReasonAllOf1) UnmarshalBinary(b []byte) error {
	var res ReasonAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
