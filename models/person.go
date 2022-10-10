// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Person A person
//
// swagger:model Person
type Person struct {
	Entity

	// Only loaded with a flag
	Addresses *Addresses `json:"addresses,omitempty"`

	// birthday
	Birthday string `json:"birthday,omitempty"`

	// default address fid
	DefaultAddressFid string `json:"defaultAddressFid,omitempty"`

	// Default email
	DefaultEmail string `json:"defaultEmail,omitempty"`

	// default email fid
	DefaultEmailFid string `json:"defaultEmailFid,omitempty"`

	// Default phone
	DefaultPhone string `json:"defaultPhone,omitempty"`

	// default phone fid
	DefaultPhoneFid string `json:"defaultPhoneFid,omitempty"`

	// Only loaded with a flag
	Emails *Emails `json:"emails,omitempty"`

	// external reference
	ExternalReference string `json:"externalReference,omitempty"`

	// first name
	FirstName string `json:"firstName,omitempty"`

	// job title
	JobTitle string `json:"jobTitle,omitempty"`

	// last name
	LastName string `json:"lastName,omitempty"`

	// middle names
	MiddleNames string `json:"middleNames,omitempty"`

	// nickname
	Nickname string `json:"nickname,omitempty"`

	// owner fid
	OwnerFid string `json:"ownerFid,omitempty"`

	// Only loaded with a flag
	Phones *Phones `json:"phones,omitempty"`

	// prefix
	Prefix string `json:"prefix,omitempty"`

	// suffix
	Suffix string `json:"suffix,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Person) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Entity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Entity = aO0

	// AO1
	var dataAO1 struct {
		Addresses *Addresses `json:"addresses,omitempty"`

		Birthday string `json:"birthday,omitempty"`

		DefaultAddressFid string `json:"defaultAddressFid,omitempty"`

		DefaultEmail string `json:"defaultEmail,omitempty"`

		DefaultEmailFid string `json:"defaultEmailFid,omitempty"`

		DefaultPhone string `json:"defaultPhone,omitempty"`

		DefaultPhoneFid string `json:"defaultPhoneFid,omitempty"`

		Emails *Emails `json:"emails,omitempty"`

		ExternalReference string `json:"externalReference,omitempty"`

		FirstName string `json:"firstName,omitempty"`

		JobTitle string `json:"jobTitle,omitempty"`

		LastName string `json:"lastName,omitempty"`

		MiddleNames string `json:"middleNames,omitempty"`

		Nickname string `json:"nickname,omitempty"`

		OwnerFid string `json:"ownerFid,omitempty"`

		Phones *Phones `json:"phones,omitempty"`

		Prefix string `json:"prefix,omitempty"`

		Suffix string `json:"suffix,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Addresses = dataAO1.Addresses

	m.Birthday = dataAO1.Birthday

	m.DefaultAddressFid = dataAO1.DefaultAddressFid

	m.DefaultEmail = dataAO1.DefaultEmail

	m.DefaultEmailFid = dataAO1.DefaultEmailFid

	m.DefaultPhone = dataAO1.DefaultPhone

	m.DefaultPhoneFid = dataAO1.DefaultPhoneFid

	m.Emails = dataAO1.Emails

	m.ExternalReference = dataAO1.ExternalReference

	m.FirstName = dataAO1.FirstName

	m.JobTitle = dataAO1.JobTitle

	m.LastName = dataAO1.LastName

	m.MiddleNames = dataAO1.MiddleNames

	m.Nickname = dataAO1.Nickname

	m.OwnerFid = dataAO1.OwnerFid

	m.Phones = dataAO1.Phones

	m.Prefix = dataAO1.Prefix

	m.Suffix = dataAO1.Suffix

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Person) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Entity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Addresses *Addresses `json:"addresses,omitempty"`

		Birthday string `json:"birthday,omitempty"`

		DefaultAddressFid string `json:"defaultAddressFid,omitempty"`

		DefaultEmail string `json:"defaultEmail,omitempty"`

		DefaultEmailFid string `json:"defaultEmailFid,omitempty"`

		DefaultPhone string `json:"defaultPhone,omitempty"`

		DefaultPhoneFid string `json:"defaultPhoneFid,omitempty"`

		Emails *Emails `json:"emails,omitempty"`

		ExternalReference string `json:"externalReference,omitempty"`

		FirstName string `json:"firstName,omitempty"`

		JobTitle string `json:"jobTitle,omitempty"`

		LastName string `json:"lastName,omitempty"`

		MiddleNames string `json:"middleNames,omitempty"`

		Nickname string `json:"nickname,omitempty"`

		OwnerFid string `json:"ownerFid,omitempty"`

		Phones *Phones `json:"phones,omitempty"`

		Prefix string `json:"prefix,omitempty"`

		Suffix string `json:"suffix,omitempty"`
	}

	dataAO1.Addresses = m.Addresses

	dataAO1.Birthday = m.Birthday

	dataAO1.DefaultAddressFid = m.DefaultAddressFid

	dataAO1.DefaultEmail = m.DefaultEmail

	dataAO1.DefaultEmailFid = m.DefaultEmailFid

	dataAO1.DefaultPhone = m.DefaultPhone

	dataAO1.DefaultPhoneFid = m.DefaultPhoneFid

	dataAO1.Emails = m.Emails

	dataAO1.ExternalReference = m.ExternalReference

	dataAO1.FirstName = m.FirstName

	dataAO1.JobTitle = m.JobTitle

	dataAO1.LastName = m.LastName

	dataAO1.MiddleNames = m.MiddleNames

	dataAO1.Nickname = m.Nickname

	dataAO1.OwnerFid = m.OwnerFid

	dataAO1.Phones = m.Phones

	dataAO1.Prefix = m.Prefix

	dataAO1.Suffix = m.Suffix

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this person
func (m *Person) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Entity
	if err := m.Entity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhones(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Person) validateAddresses(formats strfmt.Registry) error {

	if swag.IsZero(m.Addresses) { // not required
		return nil
	}

	if m.Addresses != nil {
		if err := m.Addresses.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addresses")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addresses")
			}
			return err
		}
	}

	return nil
}

func (m *Person) validateEmails(formats strfmt.Registry) error {

	if swag.IsZero(m.Emails) { // not required
		return nil
	}

	if m.Emails != nil {
		if err := m.Emails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("emails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("emails")
			}
			return err
		}
	}

	return nil
}

func (m *Person) validatePhones(formats strfmt.Registry) error {

	if swag.IsZero(m.Phones) { // not required
		return nil
	}

	if m.Phones != nil {
		if err := m.Phones.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phones")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phones")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this person based on the context it is used
func (m *Person) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Entity
	if err := m.Entity.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAddresses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEmails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhones(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Person) contextValidateAddresses(ctx context.Context, formats strfmt.Registry) error {

	if m.Addresses != nil {
		if err := m.Addresses.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addresses")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addresses")
			}
			return err
		}
	}

	return nil
}

func (m *Person) contextValidateEmails(ctx context.Context, formats strfmt.Registry) error {

	if m.Emails != nil {
		if err := m.Emails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("emails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("emails")
			}
			return err
		}
	}

	return nil
}

func (m *Person) contextValidatePhones(ctx context.Context, formats strfmt.Registry) error {

	if m.Phones != nil {
		if err := m.Phones.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phones")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phones")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Person) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Person) UnmarshalBinary(b []byte) error {
	var res Person
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
