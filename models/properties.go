// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Properties Properties set on an entity
//
// swagger:model Properties
type Properties struct {

	// Counter properties
	Counters []*PropertyCounter `json:"counters"`

	// Flag properties
	Flags []*PropertyFlag `json:"flags"`

	// Value properties
	Values []*PropertyValue `json:"values"`
}

// Validate validates this properties
func (m *Properties) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCounters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Properties) validateCounters(formats strfmt.Registry) error {
	if swag.IsZero(m.Counters) { // not required
		return nil
	}

	for i := 0; i < len(m.Counters); i++ {
		if swag.IsZero(m.Counters[i]) { // not required
			continue
		}

		if m.Counters[i] != nil {
			if err := m.Counters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("counters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("counters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Properties) validateFlags(formats strfmt.Registry) error {
	if swag.IsZero(m.Flags) { // not required
		return nil
	}

	for i := 0; i < len(m.Flags); i++ {
		if swag.IsZero(m.Flags[i]) { // not required
			continue
		}

		if m.Flags[i] != nil {
			if err := m.Flags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("flags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("flags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Properties) validateValues(formats strfmt.Registry) error {
	if swag.IsZero(m.Values) { // not required
		return nil
	}

	for i := 0; i < len(m.Values); i++ {
		if swag.IsZero(m.Values[i]) { // not required
			continue
		}

		if m.Values[i] != nil {
			if err := m.Values[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("values" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this properties based on the context it is used
func (m *Properties) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCounters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFlags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Properties) contextValidateCounters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Counters); i++ {

		if m.Counters[i] != nil {
			if err := m.Counters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("counters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("counters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Properties) contextValidateFlags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Flags); i++ {

		if m.Flags[i] != nil {
			if err := m.Flags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("flags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("flags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Properties) contextValidateValues(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Values); i++ {

		if m.Values[i] != nil {
			if err := m.Values[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("values" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Properties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Properties) UnmarshalBinary(b []byte) error {
	var res Properties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
