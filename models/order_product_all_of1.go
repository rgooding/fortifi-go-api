// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OrderProductAllOf1 order product all of1
// swagger:model orderProductAllOf1
type OrderProductAllOf1 struct {

	// currency
	Currency string `json:"currency,omitempty"`

	// cycle exact
	CycleExact int32 `json:"cycleExact,omitempty"`

	// cycle term
	CycleTerm int32 `json:"cycleTerm,omitempty"`

	// cycle type
	CycleType int32 `json:"cycleType,omitempty"`

	// discount amount
	DiscountAmount float32 `json:"discountAmount,omitempty"`

	// offer fid
	OfferFid string `json:"offerFid,omitempty"`

	// parent fid
	ParentFid string `json:"parentFid,omitempty"`

	// price
	Price float32 `json:"price,omitempty"`

	// price fid
	PriceFid string `json:"priceFid,omitempty"`

	// product fid
	ProductFid string `json:"productFid,omitempty"`

	// purchase fid
	PurchaseFid string `json:"purchaseFid,omitempty"`

	// quantity
	Quantity int64 `json:"quantity,omitempty"`

	// renewal date
	RenewalDate int32 `json:"renewalDate,omitempty"`

	// setup discount amount
	SetupDiscountAmount float32 `json:"setupDiscountAmount,omitempty"`

	// setup fee
	SetupFee float32 `json:"setupFee,omitempty"`

	// tax amount
	TaxAmount float32 `json:"taxAmount,omitempty"`

	// total amount
	TotalAmount float32 `json:"totalAmount,omitempty"`
}

// Validate validates this order product all of1
func (m *OrderProductAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OrderProductAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderProductAllOf1) UnmarshalBinary(b []byte) error {
	var res OrderProductAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
