// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TicketPost Post On a Ticket
// swagger:model TicketPost
type TicketPost struct {

	// attachment count
	AttachmentCount int64 `json:"attachmentCount,omitempty"`

	// author fid
	AuthorFid string `json:"authorFid,omitempty"`

	// has attachments
	HasAttachments bool `json:"hasAttachments,omitempty"`

	// language
	Language string `json:"language,omitempty"`

	// sender fid
	SenderFid string `json:"senderFid,omitempty"`

	// text body
	TextBody string `json:"textBody,omitempty"`

	// timestamp
	Timestamp int64 `json:"timestamp,omitempty"`
}

// Validate validates this ticket post
func (m *TicketPost) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TicketPost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TicketPost) UnmarshalBinary(b []byte) error {
	var res TicketPost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
