// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PaymentAccount Generic Response
// swagger:model PaymentAccount
type PaymentAccount struct {
	Entity

	// account holder
	AccountHolder string `json:"accountHolder,omitempty"`

	// account type
	AccountType PaymentAccountType `json:"accountType,omitempty"`

	// payment method
	PaymentMethod PaymentMethod `json:"paymentMethod,omitempty"`

	// payment mode
	PaymentMode PaymentMode `json:"paymentMode,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PaymentAccount) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Entity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Entity = aO0

	// AO1
	var dataAO1 struct {
		AccountHolder string `json:"accountHolder,omitempty"`

		AccountType PaymentAccountType `json:"accountType,omitempty"`

		PaymentMethod PaymentMethod `json:"paymentMethod,omitempty"`

		PaymentMode PaymentMode `json:"paymentMode,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AccountHolder = dataAO1.AccountHolder

	m.AccountType = dataAO1.AccountType

	m.PaymentMethod = dataAO1.PaymentMethod

	m.PaymentMode = dataAO1.PaymentMode

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PaymentAccount) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Entity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		AccountHolder string `json:"accountHolder,omitempty"`

		AccountType PaymentAccountType `json:"accountType,omitempty"`

		PaymentMethod PaymentMethod `json:"paymentMethod,omitempty"`

		PaymentMode PaymentMode `json:"paymentMode,omitempty"`
	}

	dataAO1.AccountHolder = m.AccountHolder

	dataAO1.AccountType = m.AccountType

	dataAO1.PaymentMethod = m.PaymentMethod

	dataAO1.PaymentMode = m.PaymentMode

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this payment account
func (m *PaymentAccount) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Entity
	if err := m.Entity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentAccount) validateAccountType(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountType) { // not required
		return nil
	}

	if err := m.AccountType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("accountType")
		}
		return err
	}

	return nil
}

func (m *PaymentAccount) validatePaymentMethod(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentMethod) { // not required
		return nil
	}

	if err := m.PaymentMethod.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("paymentMethod")
		}
		return err
	}

	return nil
}

func (m *PaymentAccount) validatePaymentMode(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentMode) { // not required
		return nil
	}

	if err := m.PaymentMode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("paymentMode")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAccount) UnmarshalBinary(b []byte) error {
	var res PaymentAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
