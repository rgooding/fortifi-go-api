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

// InteractionType interaction type
//
// swagger:model InteractionType
type InteractionType struct {

	// name
	// Enum: [unknown voice email chat ticket meeting remote letter social]
	Name string `json:"name,omitempty"`
}

// Validate validates this interaction type
func (m *InteractionType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var interactionTypeTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unknown","voice","email","chat","ticket","meeting","remote","letter","social"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		interactionTypeTypeNamePropEnum = append(interactionTypeTypeNamePropEnum, v)
	}
}

const (

	// InteractionTypeNameUnknown captures enum value "unknown"
	InteractionTypeNameUnknown string = "unknown"

	// InteractionTypeNameVoice captures enum value "voice"
	InteractionTypeNameVoice string = "voice"

	// InteractionTypeNameEmail captures enum value "email"
	InteractionTypeNameEmail string = "email"

	// InteractionTypeNameChat captures enum value "chat"
	InteractionTypeNameChat string = "chat"

	// InteractionTypeNameTicket captures enum value "ticket"
	InteractionTypeNameTicket string = "ticket"

	// InteractionTypeNameMeeting captures enum value "meeting"
	InteractionTypeNameMeeting string = "meeting"

	// InteractionTypeNameRemote captures enum value "remote"
	InteractionTypeNameRemote string = "remote"

	// InteractionTypeNameLetter captures enum value "letter"
	InteractionTypeNameLetter string = "letter"

	// InteractionTypeNameSocial captures enum value "social"
	InteractionTypeNameSocial string = "social"
)

// prop value enum
func (m *InteractionType) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, interactionTypeTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InteractionType) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this interaction type based on context it is used
func (m *InteractionType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InteractionType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InteractionType) UnmarshalBinary(b []byte) error {
	var res InteractionType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
