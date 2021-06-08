// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InvoiceSummary Generic Response
//
// swagger:model InvoiceSummary
type InvoiceSummary struct {
	Entity

	// amount paid
	AmountPaid float32 `json:"amountPaid,omitempty"`

	// base amount
	BaseAmount float32 `json:"baseAmount,omitempty"`

	// credited amount
	CreditedAmount float32 `json:"creditedAmount,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`

	// discount amount
	DiscountAmount float32 `json:"discountAmount,omitempty"`

	// Time in ISO 8601 standard with optional fractions of a second e.g 2015-12-05T13:11:59.888Z
	// Format: date-time
	DueDate strfmt.DateTime `json:"dueDate,omitempty"`

	// Time in ISO 8601 standard with optional fractions of a second e.g 2015-12-05T13:11:59.888Z
	// Format: date-time
	InvoiceDate strfmt.DateTime `json:"invoiceDate,omitempty"`

	// invoice name
	InvoiceName string `json:"invoiceName,omitempty"`

	// invoice number
	InvoiceNumber int32 `json:"invoiceNumber,omitempty"`

	// invoice status
	InvoiceStatus string `json:"invoiceStatus,omitempty"`

	// outstanding amount
	OutstandingAmount float32 `json:"outstandingAmount,omitempty"`

	// Time in ISO 8601 standard with optional fractions of a second e.g 2015-12-05T13:11:59.888Z
	// Format: date-time
	PaymentDate strfmt.DateTime `json:"paymentDate,omitempty"`

	// refund amount
	RefundAmount float32 `json:"refundAmount,omitempty"`

	// tax amount
	TaxAmount float32 `json:"taxAmount,omitempty"`

	// total amount
	TotalAmount float32 `json:"totalAmount,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *InvoiceSummary) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Entity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Entity = aO0

	// AO1
	var dataAO1 struct {
		AmountPaid float32 `json:"amountPaid,omitempty"`

		BaseAmount float32 `json:"baseAmount,omitempty"`

		CreditedAmount float32 `json:"creditedAmount,omitempty"`

		Currency string `json:"currency,omitempty"`

		DiscountAmount float32 `json:"discountAmount,omitempty"`

		DueDate strfmt.DateTime `json:"dueDate,omitempty"`

		InvoiceDate strfmt.DateTime `json:"invoiceDate,omitempty"`

		InvoiceName string `json:"invoiceName,omitempty"`

		InvoiceNumber int32 `json:"invoiceNumber,omitempty"`

		InvoiceStatus string `json:"invoiceStatus,omitempty"`

		OutstandingAmount float32 `json:"outstandingAmount,omitempty"`

		PaymentDate strfmt.DateTime `json:"paymentDate,omitempty"`

		RefundAmount float32 `json:"refundAmount,omitempty"`

		TaxAmount float32 `json:"taxAmount,omitempty"`

		TotalAmount float32 `json:"totalAmount,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AmountPaid = dataAO1.AmountPaid

	m.BaseAmount = dataAO1.BaseAmount

	m.CreditedAmount = dataAO1.CreditedAmount

	m.Currency = dataAO1.Currency

	m.DiscountAmount = dataAO1.DiscountAmount

	m.DueDate = dataAO1.DueDate

	m.InvoiceDate = dataAO1.InvoiceDate

	m.InvoiceName = dataAO1.InvoiceName

	m.InvoiceNumber = dataAO1.InvoiceNumber

	m.InvoiceStatus = dataAO1.InvoiceStatus

	m.OutstandingAmount = dataAO1.OutstandingAmount

	m.PaymentDate = dataAO1.PaymentDate

	m.RefundAmount = dataAO1.RefundAmount

	m.TaxAmount = dataAO1.TaxAmount

	m.TotalAmount = dataAO1.TotalAmount

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m InvoiceSummary) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Entity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		AmountPaid float32 `json:"amountPaid,omitempty"`

		BaseAmount float32 `json:"baseAmount,omitempty"`

		CreditedAmount float32 `json:"creditedAmount,omitempty"`

		Currency string `json:"currency,omitempty"`

		DiscountAmount float32 `json:"discountAmount,omitempty"`

		DueDate strfmt.DateTime `json:"dueDate,omitempty"`

		InvoiceDate strfmt.DateTime `json:"invoiceDate,omitempty"`

		InvoiceName string `json:"invoiceName,omitempty"`

		InvoiceNumber int32 `json:"invoiceNumber,omitempty"`

		InvoiceStatus string `json:"invoiceStatus,omitempty"`

		OutstandingAmount float32 `json:"outstandingAmount,omitempty"`

		PaymentDate strfmt.DateTime `json:"paymentDate,omitempty"`

		RefundAmount float32 `json:"refundAmount,omitempty"`

		TaxAmount float32 `json:"taxAmount,omitempty"`

		TotalAmount float32 `json:"totalAmount,omitempty"`
	}

	dataAO1.AmountPaid = m.AmountPaid

	dataAO1.BaseAmount = m.BaseAmount

	dataAO1.CreditedAmount = m.CreditedAmount

	dataAO1.Currency = m.Currency

	dataAO1.DiscountAmount = m.DiscountAmount

	dataAO1.DueDate = m.DueDate

	dataAO1.InvoiceDate = m.InvoiceDate

	dataAO1.InvoiceName = m.InvoiceName

	dataAO1.InvoiceNumber = m.InvoiceNumber

	dataAO1.InvoiceStatus = m.InvoiceStatus

	dataAO1.OutstandingAmount = m.OutstandingAmount

	dataAO1.PaymentDate = m.PaymentDate

	dataAO1.RefundAmount = m.RefundAmount

	dataAO1.TaxAmount = m.TaxAmount

	dataAO1.TotalAmount = m.TotalAmount

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this invoice summary
func (m *InvoiceSummary) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Entity
	if err := m.Entity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDueDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvoiceDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InvoiceSummary) validateDueDate(formats strfmt.Registry) error {

	if swag.IsZero(m.DueDate) { // not required
		return nil
	}

	if err := validate.FormatOf("dueDate", "body", "date-time", m.DueDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceSummary) validateInvoiceDate(formats strfmt.Registry) error {

	if swag.IsZero(m.InvoiceDate) { // not required
		return nil
	}

	if err := validate.FormatOf("invoiceDate", "body", "date-time", m.InvoiceDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceSummary) validatePaymentDate(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentDate) { // not required
		return nil
	}

	if err := validate.FormatOf("paymentDate", "body", "date-time", m.PaymentDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this invoice summary based on the context it is used
func (m *InvoiceSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Entity
	if err := m.Entity.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *InvoiceSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InvoiceSummary) UnmarshalBinary(b []byte) error {
	var res InvoiceSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
