// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FinalizeOrderPayload finalize order payload
// swagger:model FinalizeOrderPayload
type FinalizeOrderPayload struct {

	// ChargeHive Charge ID
	ChargeID string `json:"chargeId,omitempty"`

	// Payment Method ID
	MethodID string `json:"methodId,omitempty"`

	// Transaction ID which authorized this order
	TransactionID string `json:"transactionId,omitempty"`
}

// Validate validates this finalize order payload
func (m *FinalizeOrderPayload) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FinalizeOrderPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FinalizeOrderPayload) UnmarshalBinary(b []byte) error {
	var res FinalizeOrderPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
