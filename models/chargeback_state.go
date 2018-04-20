// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ChargebackState Current State
// swagger:model ChargebackState
type ChargebackState string

const (
	// ChargebackStateAlert captures enum value "alert"
	ChargebackStateAlert ChargebackState = "alert"
	// ChargebackStateInitiated captures enum value "initiated"
	ChargebackStateInitiated ChargebackState = "initiated"
	// ChargebackStateDisputed captures enum value "disputed"
	ChargebackStateDisputed ChargebackState = "disputed"
	// ChargebackStateWon captures enum value "won"
	ChargebackStateWon ChargebackState = "won"
	// ChargebackStateLost captures enum value "lost"
	ChargebackStateLost ChargebackState = "lost"
	// ChargebackStateUndisputedLoss captures enum value "undisputed_loss"
	ChargebackStateUndisputedLoss ChargebackState = "undisputed_loss"
)

// for schema
var chargebackStateEnum []interface{}

func init() {
	var res []ChargebackState
	if err := json.Unmarshal([]byte(`["alert","initiated","disputed","won","lost","undisputed_loss"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		chargebackStateEnum = append(chargebackStateEnum, v)
	}
}

func (m ChargebackState) validateChargebackStateEnum(path, location string, value ChargebackState) error {
	if err := validate.Enum(path, location, value, chargebackStateEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this chargeback state
func (m ChargebackState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateChargebackStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}