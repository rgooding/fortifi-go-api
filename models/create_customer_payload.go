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

// CreateCustomerPayload Payload for creating a customer
//
// swagger:model CreateCustomerPayload
type CreateCustomerPayload struct {

	// FID of an Account Manager
	AccountManagerFid string `json:"accountManagerFid,omitempty"`

	// account status
	AccountStatus CustomerAccountStatus `json:"accountStatus,omitempty"`

	// account type
	AccountType CustomerAccountType `json:"accountType,omitempty"`

	// FID of a valid Brand
	// Required: true
	BrandFid *string `json:"brandFid"`

	// IP Address of the visitor
	ClientIP string `json:"clientIp,omitempty"`

	// Company registration number (If customer is a company)
	CompanyNumber string `json:"companyNumber,omitempty"`

	// Currency
	Currency string `json:"currency,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// Email Address
	Email string `json:"email,omitempty"`

	// External (to Fortifi) Reference e.g. your internal Unique ID
	ExternalReference string `json:"externalReference,omitempty"`

	// First Name
	FirstName string `json:"firstName,omitempty"`

	// Language
	Language string `json:"language,omitempty"`

	// Last Name
	LastName string `json:"lastName,omitempty"`

	// lifecycle
	Lifecycle CustomerLifecycle `json:"lifecycle,omitempty"`

	// Phone Number
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// subscription type
	SubscriptionType CustomerSubscriptionType `json:"subscriptionType,omitempty"`

	// VAT number
	TaxNumber string `json:"taxNumber,omitempty"`

	// Time in ISO 8601 standard with optional fractions of a second e.g 2015-12-05T13:11:59.888Z
	// Format: date-time
	Time strfmt.DateTime `json:"time,omitempty"`

	// Visitor ID of the visitor
	VisitorID string `json:"visitorId,omitempty"`
}

// Validate validates this create customer payload
func (m *CreateCustomerPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBrandFid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLifecycle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriptionType(formats); err != nil {
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

func (m *CreateCustomerPayload) validateAccountStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.AccountStatus) { // not required
		return nil
	}

	if err := m.AccountStatus.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("accountStatus")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("accountStatus")
		}
		return err
	}

	return nil
}

func (m *CreateCustomerPayload) validateAccountType(formats strfmt.Registry) error {
	if swag.IsZero(m.AccountType) { // not required
		return nil
	}

	if err := m.AccountType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("accountType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("accountType")
		}
		return err
	}

	return nil
}

func (m *CreateCustomerPayload) validateBrandFid(formats strfmt.Registry) error {

	if err := validate.Required("brandFid", "body", m.BrandFid); err != nil {
		return err
	}

	return nil
}

func (m *CreateCustomerPayload) validateLifecycle(formats strfmt.Registry) error {
	if swag.IsZero(m.Lifecycle) { // not required
		return nil
	}

	if err := m.Lifecycle.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lifecycle")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("lifecycle")
		}
		return err
	}

	return nil
}

func (m *CreateCustomerPayload) validateSubscriptionType(formats strfmt.Registry) error {
	if swag.IsZero(m.SubscriptionType) { // not required
		return nil
	}

	if err := m.SubscriptionType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("subscriptionType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("subscriptionType")
		}
		return err
	}

	return nil
}

func (m *CreateCustomerPayload) validateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("time", "body", "date-time", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create customer payload based on the context it is used
func (m *CreateCustomerPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccountStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAccountType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLifecycle(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubscriptionType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateCustomerPayload) contextValidateAccountStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AccountStatus.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("accountStatus")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("accountStatus")
		}
		return err
	}

	return nil
}

func (m *CreateCustomerPayload) contextValidateAccountType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AccountType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("accountType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("accountType")
		}
		return err
	}

	return nil
}

func (m *CreateCustomerPayload) contextValidateLifecycle(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Lifecycle.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lifecycle")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("lifecycle")
		}
		return err
	}

	return nil
}

func (m *CreateCustomerPayload) contextValidateSubscriptionType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.SubscriptionType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("subscriptionType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("subscriptionType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateCustomerPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateCustomerPayload) UnmarshalBinary(b []byte) error {
	var res CreateCustomerPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
