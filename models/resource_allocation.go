// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResourceAllocation Resource Allocation
// swagger:model ResourceAllocation
type ResourceAllocation struct {

	// allocation
	Allocation string `json:"allocation,omitempty"`

	// resource fid
	ResourceFid string `json:"resourceFid,omitempty"`

	// resource gateway path
	ResourceGatewayPath string `json:"resourceGatewayPath,omitempty"`
}

// Validate validates this resource allocation
func (m *ResourceAllocation) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResourceAllocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceAllocation) UnmarshalBinary(b []byte) error {
	var res ResourceAllocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
