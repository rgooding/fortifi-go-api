// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModifySubscriptionPayload modify subscription payload
// swagger:model ModifySubscriptionPayload
type ModifySubscriptionPayload struct {

	// mode
	Mode ModifySubscriptionMode `json:"mode,omitempty"`

	// Price FID to modify subscription with
	OfferFid string `json:"offerFid,omitempty"`

	// Price FID to modify subscription with
	// Required: true
	PriceFid *string `json:"priceFid"`
}

// Validate validates this modify subscription payload
func (m *ModifySubscriptionPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePriceFid(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModifySubscriptionPayload) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	if err := m.Mode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("mode")
		}
		return err
	}

	return nil
}

func (m *ModifySubscriptionPayload) validatePriceFid(formats strfmt.Registry) error {

	if err := validate.Required("priceFid", "body", m.PriceFid); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModifySubscriptionPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModifySubscriptionPayload) UnmarshalBinary(b []byte) error {
	var res ModifySubscriptionPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
