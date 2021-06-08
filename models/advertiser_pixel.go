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

// AdvertiserPixel Pixel
//
// swagger:model AdvertiserPixel
type AdvertiserPixel struct {

	// content
	Content string `json:"content,omitempty"`

	// method
	// Enum: [iframe img js curl html]
	Method string `json:"method,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this advertiser pixel
func (m *AdvertiserPixel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMethod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var advertiserPixelTypeMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["iframe","img","js","curl","html"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		advertiserPixelTypeMethodPropEnum = append(advertiserPixelTypeMethodPropEnum, v)
	}
}

const (

	// AdvertiserPixelMethodIframe captures enum value "iframe"
	AdvertiserPixelMethodIframe string = "iframe"

	// AdvertiserPixelMethodImg captures enum value "img"
	AdvertiserPixelMethodImg string = "img"

	// AdvertiserPixelMethodJs captures enum value "js"
	AdvertiserPixelMethodJs string = "js"

	// AdvertiserPixelMethodCurl captures enum value "curl"
	AdvertiserPixelMethodCurl string = "curl"

	// AdvertiserPixelMethodHTML captures enum value "html"
	AdvertiserPixelMethodHTML string = "html"
)

// prop value enum
func (m *AdvertiserPixel) validateMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, advertiserPixelTypeMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AdvertiserPixel) validateMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.Method) { // not required
		return nil
	}

	// value enum
	if err := m.validateMethodEnum("method", "body", m.Method); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this advertiser pixel based on context it is used
func (m *AdvertiserPixel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AdvertiserPixel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdvertiserPixel) UnmarshalBinary(b []byte) error {
	var res AdvertiserPixel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
