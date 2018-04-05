// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// PaymentMethod Payment Method
// swagger:model PaymentMethod
type PaymentMethod string

const (
	// PaymentMethodCreditcard captures enum value "creditcard"
	PaymentMethodCreditcard PaymentMethod = "creditcard"
	// PaymentMethodDebitcard captures enum value "debitcard"
	PaymentMethodDebitcard PaymentMethod = "debitcard"
	// PaymentMethodCheque captures enum value "cheque"
	PaymentMethodCheque PaymentMethod = "cheque"
	// PaymentMethodCash captures enum value "cash"
	PaymentMethodCash PaymentMethod = "cash"
	// PaymentMethodPrepaidcard captures enum value "prepaidcard"
	PaymentMethodPrepaidcard PaymentMethod = "prepaidcard"
	// PaymentMethodDirectdebit captures enum value "directdebit"
	PaymentMethodDirectdebit PaymentMethod = "directdebit"
	// PaymentMethodBacs captures enum value "bacs"
	PaymentMethodBacs PaymentMethod = "bacs"
	// PaymentMethodStandingorder captures enum value "standingorder"
	PaymentMethodStandingorder PaymentMethod = "standingorder"
	// PaymentMethodChaps captures enum value "chaps"
	PaymentMethodChaps PaymentMethod = "chaps"
	// PaymentMethodOnlineservice captures enum value "onlineservice"
	PaymentMethodOnlineservice PaymentMethod = "onlineservice"
	// PaymentMethodTelephone captures enum value "telephone"
	PaymentMethodTelephone PaymentMethod = "telephone"
	// PaymentMethodCreditnote captures enum value "creditnote"
	PaymentMethodCreditnote PaymentMethod = "creditnote"
	// PaymentMethodVirtualcard captures enum value "virtualcard"
	PaymentMethodVirtualcard PaymentMethod = "virtualcard"
	// PaymentMethodGiftcard captures enum value "giftcard"
	PaymentMethodGiftcard PaymentMethod = "giftcard"
	// PaymentMethodUnknown captures enum value "unknown"
	PaymentMethodUnknown PaymentMethod = "unknown"
	// PaymentMethodMultiple captures enum value "multiple"
	PaymentMethodMultiple PaymentMethod = "multiple"
	// PaymentMethodPaypal captures enum value "paypal"
	PaymentMethodPaypal PaymentMethod = "paypal"
	// PaymentMethodBitcoin captures enum value "bitcoin"
	PaymentMethodBitcoin PaymentMethod = "bitcoin"
	// PaymentMethodAccountBalance captures enum value "account_balance"
	PaymentMethodAccountBalance PaymentMethod = "account_balance"
)

// for schema
var paymentMethodEnum []interface{}

func init() {
	var res []PaymentMethod
	if err := json.Unmarshal([]byte(`["creditcard","debitcard","cheque","cash","prepaidcard","directdebit","bacs","standingorder","chaps","onlineservice","telephone","creditnote","virtualcard","giftcard","unknown","multiple","paypal","bitcoin","account_balance"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentMethodEnum = append(paymentMethodEnum, v)
	}
}

func (m PaymentMethod) validatePaymentMethodEnum(path, location string, value PaymentMethod) error {
	if err := validate.Enum(path, location, value, paymentMethodEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this payment method
func (m PaymentMethod) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePaymentMethodEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
