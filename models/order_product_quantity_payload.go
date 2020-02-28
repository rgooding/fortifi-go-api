// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OrderProductQuantityPayload order product quantity payload
// swagger:model OrderProductQuantityPayload
type OrderProductQuantityPayload struct {

	// Price FID to modify subscription with
	// Required: true
	PriceFid *string `json:"priceFid"`

	// quantity
	Quantity *int64 `json:"quantity,omitempty"`
}

// Validate validates this order product quantity payload
func (m *OrderProductQuantityPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePriceFid(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderProductQuantityPayload) validatePriceFid(formats strfmt.Registry) error {

	if err := validate.Required("priceFid", "body", m.PriceFid); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderProductQuantityPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderProductQuantityPayload) UnmarshalBinary(b []byte) error {
	var res OrderProductQuantityPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
