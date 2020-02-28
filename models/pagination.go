// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Pagination Generic Response
// swagger:model Pagination
type Pagination struct {

	// limit
	Limit int32 `json:"limit,omitempty"`

	// page
	Page int32 `json:"page,omitempty"`

	// total items
	TotalItems int32 `json:"totalItems,omitempty"`
}

// Validate validates this pagination
func (m *Pagination) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Pagination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Pagination) UnmarshalBinary(b []byte) error {
	var res Pagination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
