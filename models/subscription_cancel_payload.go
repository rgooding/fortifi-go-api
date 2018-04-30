// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SubscriptionCancelPayload subscription cancel payload
// swagger:model SubscriptionCancelPayload
type SubscriptionCancelPayload struct {

	// Reason FID
	ReasonFid string `json:"reasonFid,omitempty"`

	// subscription refund type
	// Required: true
	SubscriptionRefundType SubscriptionRefundType `json:"subscriptionRefundType"`
}

// Validate validates this subscription cancel payload
func (m *SubscriptionCancelPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubscriptionRefundType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubscriptionCancelPayload) validateSubscriptionRefundType(formats strfmt.Registry) error {

	if err := m.SubscriptionRefundType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("subscriptionRefundType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubscriptionCancelPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriptionCancelPayload) UnmarshalBinary(b []byte) error {
	var res SubscriptionCancelPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
