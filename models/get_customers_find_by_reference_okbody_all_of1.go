// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetCustomersFindByReferenceOKBodyAllOf1 get customers find by reference o k body all of1
// swagger:model getCustomersFindByReferenceOKBodyAllOf1
type GetCustomersFindByReferenceOKBodyAllOf1 struct {

	// data
	Data *Customer `json:"data,omitempty"`
}

// Validate validates this get customers find by reference o k body all of1
func (m *GetCustomersFindByReferenceOKBodyAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetCustomersFindByReferenceOKBodyAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCustomersFindByReferenceOKBodyAllOf1) UnmarshalBinary(b []byte) error {
	var res GetCustomersFindByReferenceOKBodyAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
