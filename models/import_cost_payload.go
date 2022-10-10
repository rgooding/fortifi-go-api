// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ImportCostPayload Import Publisher Cost csv
//
// swagger:model ImportCostPayload
type ImportCostPayload struct {

	// data
	Data string `json:"data,omitempty"`

	// filename
	Filename string `json:"filename,omitempty"`

	// mime type
	MimeType string `json:"mimeType,omitempty"`
}

// Validate validates this import cost payload
func (m *ImportCostPayload) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this import cost payload based on context it is used
func (m *ImportCostPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ImportCostPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImportCostPayload) UnmarshalBinary(b []byte) error {
	var res ImportCostPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
