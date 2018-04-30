// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ReverseActionPayload reverse action payload
// swagger:model ReverseActionPayload
type ReverseActionPayload struct {

	// IP Address of the visitor
	ClientIP string `json:"clientIp,omitempty"`

	// Encoding from the visitors browser 'HTTP_ACCEPT_ENCODING'
	Encoding string `json:"encoding,omitempty"`

	// If known, the Fortifi event ID for this visitors action
	EventID string `json:"eventId,omitempty"`

	// External (to Fortifi) Reference for this visitor e.g. a user ID
	ExternalReference string `json:"externalReference,omitempty"`

	// Language from visitors browser 'HTTP_ACCEPT_LANGUAGE'
	Language string `json:"language,omitempty"`

	// meta data
	MetaData MetaData `json:"metaData"`

	// reason
	Reason ReversalReason `json:"reason,omitempty"`

	// Amount of revene refunded to the client from the original transaction / chargeback amount
	ReversalAmount float32 `json:"reversalAmount,omitempty"`

	// Your unique transaction ID for this event e.g. Refund ID
	ReversalID string `json:"reversalId,omitempty"`

	// Your unique transaction ID for the event to be reversed
	SourceTransactionID string `json:"sourceTransactionId,omitempty"`

	// Time in ISO 8601 standard with optional fractions of a second e.g 2015-12-05T13:11:59.888Z
	Time strfmt.DateTime `json:"time,omitempty"`

	// User Agent of the visitors browser 'HTTP_USER_AGENT'
	UserAgent string `json:"userAgent,omitempty"`
}

// Validate validates this reverse action payload
func (m *ReverseActionPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReason(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReverseActionPayload) validateReason(formats strfmt.Registry) error {

	if swag.IsZero(m.Reason) { // not required
		return nil
	}

	if err := m.Reason.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("reason")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReverseActionPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReverseActionPayload) UnmarshalBinary(b []byte) error {
	var res ReverseActionPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
