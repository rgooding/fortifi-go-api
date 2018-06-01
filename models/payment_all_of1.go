// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PaymentAllOf1 payment all of1
// swagger:model paymentAllOf1
type PaymentAllOf1 struct {

	// amount
	Amount float32 `json:"amount,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`

	// direction
	Direction string `json:"direction,omitempty"`

	// dispute fid
	DisputeFid string `json:"disputeFid,omitempty"`

	// fee
	Fee float32 `json:"fee,omitempty"`

	// fee currency
	FeeCurrency string `json:"feeCurrency,omitempty"`

	// fraud fid
	FraudFid string `json:"fraudFid,omitempty"`

	// invoice fid
	InvoiceFid string `json:"invoiceFid,omitempty"`

	// order fid
	OrderFid string `json:"orderFid,omitempty"`

	// processed
	Processed bool `json:"processed,omitempty"`

	// processed date
	ProcessedDate int64 `json:"processedDate,omitempty"`

	// statement descriptor
	StatementDescriptor string `json:"statementDescriptor,omitempty"`

	// status
	// Enum: [pending paid refund-pending refunded partially-refunded chargeback chargeback-pending]
	Status string `json:"status,omitempty"`

	// status code
	StatusCode string `json:"statusCode,omitempty"`

	// status message
	StatusMessage string `json:"statusMessage,omitempty"`

	// sub gateway fid
	SubGatewayFid string `json:"subGatewayFid,omitempty"`

	// sub gateway transaction Id
	SubGatewayTransactionID string `json:"subGatewayTransactionId,omitempty"`

	// transaction Id
	TransactionID string `json:"transactionId,omitempty"`

	// type
	// Enum: [affiliate invoice order unknown preauth]
	Type string `json:"type,omitempty"`
}

// Validate validates this payment all of1
func (m *PaymentAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var paymentAllOf1TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pending","paid","refund-pending","refunded","partially-refunded","chargeback","chargeback-pending"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentAllOf1TypeStatusPropEnum = append(paymentAllOf1TypeStatusPropEnum, v)
	}
}

const (

	// PaymentAllOf1StatusPending captures enum value "pending"
	PaymentAllOf1StatusPending string = "pending"

	// PaymentAllOf1StatusPaid captures enum value "paid"
	PaymentAllOf1StatusPaid string = "paid"

	// PaymentAllOf1StatusRefundPending captures enum value "refund-pending"
	PaymentAllOf1StatusRefundPending string = "refund-pending"

	// PaymentAllOf1StatusRefunded captures enum value "refunded"
	PaymentAllOf1StatusRefunded string = "refunded"

	// PaymentAllOf1StatusPartiallyRefunded captures enum value "partially-refunded"
	PaymentAllOf1StatusPartiallyRefunded string = "partially-refunded"

	// PaymentAllOf1StatusChargeback captures enum value "chargeback"
	PaymentAllOf1StatusChargeback string = "chargeback"

	// PaymentAllOf1StatusChargebackPending captures enum value "chargeback-pending"
	PaymentAllOf1StatusChargebackPending string = "chargeback-pending"
)

// prop value enum
func (m *PaymentAllOf1) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, paymentAllOf1TypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PaymentAllOf1) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var paymentAllOf1TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["affiliate","invoice","order","unknown","preauth"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentAllOf1TypeTypePropEnum = append(paymentAllOf1TypeTypePropEnum, v)
	}
}

const (

	// PaymentAllOf1TypeAffiliate captures enum value "affiliate"
	PaymentAllOf1TypeAffiliate string = "affiliate"

	// PaymentAllOf1TypeInvoice captures enum value "invoice"
	PaymentAllOf1TypeInvoice string = "invoice"

	// PaymentAllOf1TypeOrder captures enum value "order"
	PaymentAllOf1TypeOrder string = "order"

	// PaymentAllOf1TypeUnknown captures enum value "unknown"
	PaymentAllOf1TypeUnknown string = "unknown"

	// PaymentAllOf1TypePreauth captures enum value "preauth"
	PaymentAllOf1TypePreauth string = "preauth"
)

// prop value enum
func (m *PaymentAllOf1) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, paymentAllOf1TypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PaymentAllOf1) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAllOf1) UnmarshalBinary(b []byte) error {
	var res PaymentAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}