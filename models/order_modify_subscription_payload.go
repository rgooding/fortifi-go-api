// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OrderModifySubscriptionPayload order modify subscription payload
//
// swagger:model OrderModifySubscriptionPayload
type OrderModifySubscriptionPayload struct {
	ModifySubscriptionPayload

	// properties
	Properties *PropertyBulkSetPayload `json:"properties,omitempty"`

	// FID of the subscription to modify
	SourceSubscriptionFid string `json:"sourceSubscriptionFid,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *OrderModifySubscriptionPayload) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ModifySubscriptionPayload
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ModifySubscriptionPayload = aO0

	// AO1
	var dataAO1 struct {
		Properties *PropertyBulkSetPayload `json:"properties,omitempty"`

		SourceSubscriptionFid string `json:"sourceSubscriptionFid,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Properties = dataAO1.Properties

	m.SourceSubscriptionFid = dataAO1.SourceSubscriptionFid

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m OrderModifySubscriptionPayload) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ModifySubscriptionPayload)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Properties *PropertyBulkSetPayload `json:"properties,omitempty"`

		SourceSubscriptionFid string `json:"sourceSubscriptionFid,omitempty"`
	}

	dataAO1.Properties = m.Properties

	dataAO1.SourceSubscriptionFid = m.SourceSubscriptionFid

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this order modify subscription payload
func (m *OrderModifySubscriptionPayload) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ModifySubscriptionPayload
	if err := m.ModifySubscriptionPayload.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderModifySubscriptionPayload) validateProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	if m.Properties != nil {
		if err := m.Properties.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("properties")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this order modify subscription payload based on the context it is used
func (m *OrderModifySubscriptionPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ModifySubscriptionPayload
	if err := m.ModifySubscriptionPayload.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderModifySubscriptionPayload) contextValidateProperties(ctx context.Context, formats strfmt.Registry) error {

	if m.Properties != nil {
		if err := m.Properties.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("properties")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderModifySubscriptionPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderModifySubscriptionPayload) UnmarshalBinary(b []byte) error {
	var res OrderModifySubscriptionPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
