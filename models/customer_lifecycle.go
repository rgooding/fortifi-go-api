// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// CustomerLifecycle Customer Lifecycle Stage
//
// swagger:model customerLifecycle
type CustomerLifecycle string

func NewCustomerLifecycle(value CustomerLifecycle) *CustomerLifecycle {
	return &value
}

// Pointer returns a pointer to a freshly-allocated CustomerLifecycle.
func (m CustomerLifecycle) Pointer() *CustomerLifecycle {
	return &m
}

const (

	// CustomerLifecycleClosed captures enum value "closed"
	CustomerLifecycleClosed CustomerLifecycle = "closed"

	// CustomerLifecycleCold captures enum value "cold"
	CustomerLifecycleCold CustomerLifecycle = "cold"

	// CustomerLifecycleCustomer captures enum value "customer"
	CustomerLifecycleCustomer CustomerLifecycle = "customer"

	// CustomerLifecycleEvangelist captures enum value "evangelist"
	CustomerLifecycleEvangelist CustomerLifecycle = "evangelist"

	// CustomerLifecycleLead captures enum value "lead"
	CustomerLifecycleLead CustomerLifecycle = "lead"

	// CustomerLifecycleMql captures enum value "mql"
	CustomerLifecycleMql CustomerLifecycle = "mql"

	// CustomerLifecycleOpportunity captures enum value "opportunity"
	CustomerLifecycleOpportunity CustomerLifecycle = "opportunity"

	// CustomerLifecycleOther captures enum value "other"
	CustomerLifecycleOther CustomerLifecycle = "other"

	// CustomerLifecycleProspect captures enum value "prospect"
	CustomerLifecycleProspect CustomerLifecycle = "prospect"

	// CustomerLifecycleRecommended captures enum value "recommended"
	CustomerLifecycleRecommended CustomerLifecycle = "recommended"

	// CustomerLifecycleSQL captures enum value "sql"
	CustomerLifecycleSQL CustomerLifecycle = "sql"

	// CustomerLifecycleSubscriber captures enum value "subscriber"
	CustomerLifecycleSubscriber CustomerLifecycle = "subscriber"
)

// for schema
var customerLifecycleEnum []interface{}

func init() {
	var res []CustomerLifecycle
	if err := json.Unmarshal([]byte(`["closed","cold","customer","evangelist","lead","mql","opportunity","other","prospect","recommended","sql","subscriber"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerLifecycleEnum = append(customerLifecycleEnum, v)
	}
}

func (m CustomerLifecycle) validateCustomerLifecycleEnum(path, location string, value CustomerLifecycle) error {
	if err := validate.EnumCase(path, location, value, customerLifecycleEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this customer lifecycle
func (m CustomerLifecycle) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCustomerLifecycleEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this customer lifecycle based on context it is used
func (m CustomerLifecycle) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
