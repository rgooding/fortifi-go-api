// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// DiscountType Discount Type
//
// swagger:model discountType
type DiscountType string

func NewDiscountType(value DiscountType) *DiscountType {
	v := value
	return &v
}

const (

	// DiscountTypeFixed captures enum value "fixed"
	DiscountTypeFixed DiscountType = "fixed"

	// DiscountTypePercentage captures enum value "percentage"
	DiscountTypePercentage DiscountType = "percentage"
)

// for schema
var discountTypeEnum []interface{}

func init() {
	var res []DiscountType
	if err := json.Unmarshal([]byte(`["fixed","percentage"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		discountTypeEnum = append(discountTypeEnum, v)
	}
}

func (m DiscountType) validateDiscountTypeEnum(path, location string, value DiscountType) error {
	if err := validate.EnumCase(path, location, value, discountTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this discount type
func (m DiscountType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDiscountTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this discount type based on context it is used
func (m DiscountType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
