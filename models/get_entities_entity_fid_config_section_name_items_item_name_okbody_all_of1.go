// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetEntitiesEntityFidConfigSectionNameItemsItemNameOKBodyAllOf1 get entities entity fid config section name items item name o k body all of1
// swagger:model getEntitiesEntityFidConfigSectionNameItemsItemNameOKBodyAllOf1
type GetEntitiesEntityFidConfigSectionNameItemsItemNameOKBodyAllOf1 struct {

	// data
	Data *ConfigItem `json:"data,omitempty"`
}

// Validate validates this get entities entity fid config section name items item name o k body all of1
func (m *GetEntitiesEntityFidConfigSectionNameItemsItemNameOKBodyAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetEntitiesEntityFidConfigSectionNameItemsItemNameOKBodyAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetEntitiesEntityFidConfigSectionNameItemsItemNameOKBodyAllOf1) UnmarshalBinary(b []byte) error {
	var res GetEntitiesEntityFidConfigSectionNameItemsItemNameOKBodyAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
