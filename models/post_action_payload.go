// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostActionPayload post action payload
// swagger:model PostActionPayload
type PostActionPayload struct {

	// FID of a valid Brand
	// Required: true
	BrandFid *string `json:"brandFid"`

	// Advertiser Campaign to track this action to (if not already locked)
	//
	CampaignHash string `json:"campaignHash,omitempty"`

	// IP Address of the visitor
	ClientIP string `json:"clientIp,omitempty"`

	// Coupon code used for the transaction
	CouponCode string `json:"couponCode,omitempty"`

	// Encoding from the visitors browser 'HTTP_ACCEPT_ENCODING'
	Encoding string `json:"encoding,omitempty"`

	// External (to Fortifi) Reference for this visitor e.g. a user ID
	ExternalReference string `json:"externalReference,omitempty"`

	// Language from visitors browser 'HTTP_ACCEPT_LANGUAGE'
	Language string `json:"language,omitempty"`

	// meta data
	MetaData MetaData `json:"metaData"`

	// Payment method used on this transaction
	PaymentMethod string `json:"paymentMethod,omitempty"`

	// Product Code linked to this action
	ProductCode string `json:"productCode,omitempty"`

	// Product Term
	ProductTerm string `json:"productTerm,omitempty"`

	// Setting to true will return advertiser pixels for this event
	//
	ReturnPixels *bool `json:"returnPixels,omitempty"`

	// Advertised sub tracking ID 1
	Sid1 string `json:"sid1,omitempty"`

	// Advertised sub tracking ID 2
	Sid2 string `json:"sid2,omitempty"`

	// Advertised sub tracking ID 3
	Sid3 string `json:"sid3,omitempty"`

	// Time in ISO 8601 standard with optional fractions of a second e.g 2015-12-05T13:11:59.888Z
	// Format: date-time
	Time strfmt.DateTime `json:"time,omitempty"`

	// Your unique transaction ID for this event e.g. Order ID
	TransactionID string `json:"transactionId,omitempty"`

	// Your unique transaction ID for this event
	TransactionValue float32 `json:"transactionValue,omitempty"`

	// If an existing device exists for the visitor, prefer that over the user agent sent in this payload
	UseExistingDeviceIfAvailable bool `json:"useExistingDeviceIfAvailable,omitempty"`

	// User Agent of the visitors browser 'HTTP_USER_AGENT'
	UserAgent string `json:"userAgent,omitempty"`

	// Username associated with this transaction (e.g. Join)
	Username string `json:"username,omitempty"`
}

// Validate validates this post action payload
func (m *PostActionPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBrandFid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetaData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostActionPayload) validateBrandFid(formats strfmt.Registry) error {

	if err := validate.Required("brandFid", "body", m.BrandFid); err != nil {
		return err
	}

	return nil
}

func (m *PostActionPayload) validateMetaData(formats strfmt.Registry) error {

	if swag.IsZero(m.MetaData) { // not required
		return nil
	}

	if err := m.MetaData.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("metaData")
		}
		return err
	}

	return nil
}

func (m *PostActionPayload) validateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("time", "body", "date-time", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostActionPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostActionPayload) UnmarshalBinary(b []byte) error {
	var res PostActionPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
