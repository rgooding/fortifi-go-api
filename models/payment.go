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

// Payment Payment
//
// swagger:model Payment
type Payment struct {
	Fid

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

	// source account fid
	SourceAccountFid string `json:"sourceAccountFid,omitempty"`

	// source account type
	SourceAccountType string `json:"sourceAccountType,omitempty"`

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

	// sub gateway name
	SubGatewayName string `json:"subGatewayName,omitempty"`

	// sub gateway transaction Id
	SubGatewayTransactionID string `json:"subGatewayTransactionId,omitempty"`

	// transaction Id
	TransactionID string `json:"transactionId,omitempty"`

	// type
	// Enum: [affiliate invoice order unknown preauth]
	Type string `json:"type,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Payment) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Fid
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Fid = aO0

	// AO1
	var dataAO1 struct {
		Amount float32 `json:"amount,omitempty"`

		Currency string `json:"currency,omitempty"`

		Direction string `json:"direction,omitempty"`

		DisputeFid string `json:"disputeFid,omitempty"`

		Fee float32 `json:"fee,omitempty"`

		FeeCurrency string `json:"feeCurrency,omitempty"`

		FraudFid string `json:"fraudFid,omitempty"`

		InvoiceFid string `json:"invoiceFid,omitempty"`

		OrderFid string `json:"orderFid,omitempty"`

		Processed bool `json:"processed,omitempty"`

		ProcessedDate int64 `json:"processedDate,omitempty"`

		SourceAccountFid string `json:"sourceAccountFid,omitempty"`

		SourceAccountType string `json:"sourceAccountType,omitempty"`

		StatementDescriptor string `json:"statementDescriptor,omitempty"`

		Status string `json:"status,omitempty"`

		StatusCode string `json:"statusCode,omitempty"`

		StatusMessage string `json:"statusMessage,omitempty"`

		SubGatewayFid string `json:"subGatewayFid,omitempty"`

		SubGatewayName string `json:"subGatewayName,omitempty"`

		SubGatewayTransactionID string `json:"subGatewayTransactionId,omitempty"`

		TransactionID string `json:"transactionId,omitempty"`

		Type string `json:"type,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Amount = dataAO1.Amount

	m.Currency = dataAO1.Currency

	m.Direction = dataAO1.Direction

	m.DisputeFid = dataAO1.DisputeFid

	m.Fee = dataAO1.Fee

	m.FeeCurrency = dataAO1.FeeCurrency

	m.FraudFid = dataAO1.FraudFid

	m.InvoiceFid = dataAO1.InvoiceFid

	m.OrderFid = dataAO1.OrderFid

	m.Processed = dataAO1.Processed

	m.ProcessedDate = dataAO1.ProcessedDate

	m.SourceAccountFid = dataAO1.SourceAccountFid

	m.SourceAccountType = dataAO1.SourceAccountType

	m.StatementDescriptor = dataAO1.StatementDescriptor

	m.Status = dataAO1.Status

	m.StatusCode = dataAO1.StatusCode

	m.StatusMessage = dataAO1.StatusMessage

	m.SubGatewayFid = dataAO1.SubGatewayFid

	m.SubGatewayName = dataAO1.SubGatewayName

	m.SubGatewayTransactionID = dataAO1.SubGatewayTransactionID

	m.TransactionID = dataAO1.TransactionID

	m.Type = dataAO1.Type

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Payment) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Fid)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Amount float32 `json:"amount,omitempty"`

		Currency string `json:"currency,omitempty"`

		Direction string `json:"direction,omitempty"`

		DisputeFid string `json:"disputeFid,omitempty"`

		Fee float32 `json:"fee,omitempty"`

		FeeCurrency string `json:"feeCurrency,omitempty"`

		FraudFid string `json:"fraudFid,omitempty"`

		InvoiceFid string `json:"invoiceFid,omitempty"`

		OrderFid string `json:"orderFid,omitempty"`

		Processed bool `json:"processed,omitempty"`

		ProcessedDate int64 `json:"processedDate,omitempty"`

		SourceAccountFid string `json:"sourceAccountFid,omitempty"`

		SourceAccountType string `json:"sourceAccountType,omitempty"`

		StatementDescriptor string `json:"statementDescriptor,omitempty"`

		Status string `json:"status,omitempty"`

		StatusCode string `json:"statusCode,omitempty"`

		StatusMessage string `json:"statusMessage,omitempty"`

		SubGatewayFid string `json:"subGatewayFid,omitempty"`

		SubGatewayName string `json:"subGatewayName,omitempty"`

		SubGatewayTransactionID string `json:"subGatewayTransactionId,omitempty"`

		TransactionID string `json:"transactionId,omitempty"`

		Type string `json:"type,omitempty"`
	}

	dataAO1.Amount = m.Amount

	dataAO1.Currency = m.Currency

	dataAO1.Direction = m.Direction

	dataAO1.DisputeFid = m.DisputeFid

	dataAO1.Fee = m.Fee

	dataAO1.FeeCurrency = m.FeeCurrency

	dataAO1.FraudFid = m.FraudFid

	dataAO1.InvoiceFid = m.InvoiceFid

	dataAO1.OrderFid = m.OrderFid

	dataAO1.Processed = m.Processed

	dataAO1.ProcessedDate = m.ProcessedDate

	dataAO1.SourceAccountFid = m.SourceAccountFid

	dataAO1.SourceAccountType = m.SourceAccountType

	dataAO1.StatementDescriptor = m.StatementDescriptor

	dataAO1.Status = m.Status

	dataAO1.StatusCode = m.StatusCode

	dataAO1.StatusMessage = m.StatusMessage

	dataAO1.SubGatewayFid = m.SubGatewayFid

	dataAO1.SubGatewayName = m.SubGatewayName

	dataAO1.SubGatewayTransactionID = m.SubGatewayTransactionID

	dataAO1.TransactionID = m.TransactionID

	dataAO1.Type = m.Type

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this payment
func (m *Payment) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Fid
	if err := m.Fid.Validate(formats); err != nil {
		res = append(res, err)
	}

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

var paymentTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pending","paid","refund-pending","refunded","partially-refunded","chargeback","chargeback-pending"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentTypeStatusPropEnum = append(paymentTypeStatusPropEnum, v)
	}
}

// property enum
func (m *Payment) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, paymentTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Payment) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var paymentTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["affiliate","invoice","order","unknown","preauth"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentTypeTypePropEnum = append(paymentTypeTypePropEnum, v)
	}
}

// property enum
func (m *Payment) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, paymentTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Payment) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this payment based on the context it is used
func (m *Payment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Fid
	if err := m.Fid.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Payment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Payment) UnmarshalBinary(b []byte) error {
	var res Payment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
