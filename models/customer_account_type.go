// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// CustomerAccountType Account Type
// swagger:model customerAccountType
type CustomerAccountType string

const (

	// CustomerAccountTypeUnknown captures enum value "unknown"
	CustomerAccountTypeUnknown CustomerAccountType = "unknown"

	// CustomerAccountTypeStudent captures enum value "student"
	CustomerAccountTypeStudent CustomerAccountType = "student"

	// CustomerAccountTypeCharity captures enum value "charity"
	CustomerAccountTypeCharity CustomerAccountType = "charity"

	// CustomerAccountTypeBusiness captures enum value "business"
	CustomerAccountTypeBusiness CustomerAccountType = "business"

	// CustomerAccountTypeResidential captures enum value "residential"
	CustomerAccountTypeResidential CustomerAccountType = "residential"

	// CustomerAccountTypeEnterprise captures enum value "enterprise"
	CustomerAccountTypeEnterprise CustomerAccountType = "enterprise"

	// CustomerAccountTypeGroup captures enum value "group"
	CustomerAccountTypeGroup CustomerAccountType = "group"
)

// for schema
var customerAccountTypeEnum []interface{}

func init() {
	var res []CustomerAccountType
	if err := json.Unmarshal([]byte(`["unknown","student","charity","business","residential","enterprise","group"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerAccountTypeEnum = append(customerAccountTypeEnum, v)
	}
}

func (m CustomerAccountType) validateCustomerAccountTypeEnum(path, location string, value CustomerAccountType) error {
	if err := validate.Enum(path, location, value, customerAccountTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this customer account type
func (m CustomerAccountType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCustomerAccountTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
