// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VisitorReversalAction Visitor Action Reversal
//
// swagger:model VisitorReversalAction
type VisitorReversalAction struct {

	// advertiser fid
	AdvertiserFid string `json:"advertiserFid,omitempty"`

	// campaign fid
	CampaignFid string `json:"campaignFid,omitempty"`

	// commission currency
	CommissionCurrency string `json:"commissionCurrency,omitempty"`

	// event Id
	EventID string `json:"eventId,omitempty"`

	// fee charged
	FeeCharged float32 `json:"feeCharged,omitempty"`

	// reversed commission
	ReversedCommission float32 `json:"reversedCommission,omitempty"`

	// visitor Id
	VisitorID string `json:"visitorId,omitempty"`
}

// Validate validates this visitor reversal action
func (m *VisitorReversalAction) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this visitor reversal action based on context it is used
func (m *VisitorReversalAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VisitorReversalAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VisitorReversalAction) UnmarshalBinary(b []byte) error {
	var res VisitorReversalAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
