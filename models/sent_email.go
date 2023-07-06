// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SentEmail sent email
//
// swagger:model SentEmail
type SentEmail struct {

	// bounced
	Bounced bool `json:"bounced,omitempty"`

	// campaign fid
	CampaignFid string `json:"campaignFid,omitempty"`

	// clicked
	Clicked bool `json:"clicked,omitempty"`

	// company fid
	CompanyFid string `json:"companyFid,omitempty"`

	// complained
	Complained bool `json:"complained,omitempty"`

	// message fid
	MessageFid string `json:"messageFid,omitempty"`

	// opened
	Opened bool `json:"opened,omitempty"`

	// recipient emails
	RecipientEmails []string `json:"recipientEmails"`

	// sent timestamp
	SentTimestamp string `json:"sentTimestamp,omitempty"`

	// subject
	Subject string `json:"subject,omitempty"`
}

// Validate validates this sent email
func (m *SentEmail) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sent email based on context it is used
func (m *SentEmail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SentEmail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SentEmail) UnmarshalBinary(b []byte) error {
	var res SentEmail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}