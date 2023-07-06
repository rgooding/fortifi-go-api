// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RetentionFlowOption retention flow option
//
// swagger:model RetentionFlowOption
type RetentionFlowOption struct {

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// type
	// Enum: [flag value]
	Type string `json:"type,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this retention flow option
func (m *RetentionFlowOption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var retentionFlowOptionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["flag","value"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		retentionFlowOptionTypeTypePropEnum = append(retentionFlowOptionTypeTypePropEnum, v)
	}
}

const (

	// RetentionFlowOptionTypeFlag captures enum value "flag"
	RetentionFlowOptionTypeFlag string = "flag"

	// RetentionFlowOptionTypeValue captures enum value "value"
	RetentionFlowOptionTypeValue string = "value"
)

// prop value enum
func (m *RetentionFlowOption) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, retentionFlowOptionTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RetentionFlowOption) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this retention flow option based on context it is used
func (m *RetentionFlowOption) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RetentionFlowOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RetentionFlowOption) UnmarshalBinary(b []byte) error {
	var res RetentionFlowOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}