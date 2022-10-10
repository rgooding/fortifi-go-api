// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewPostCustomersCustomerFidAnonymizeParams creates a new PostCustomersCustomerFidAnonymizeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostCustomersCustomerFidAnonymizeParams() *PostCustomersCustomerFidAnonymizeParams {
	return &PostCustomersCustomerFidAnonymizeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostCustomersCustomerFidAnonymizeParamsWithTimeout creates a new PostCustomersCustomerFidAnonymizeParams object
// with the ability to set a timeout on a request.
func NewPostCustomersCustomerFidAnonymizeParamsWithTimeout(timeout time.Duration) *PostCustomersCustomerFidAnonymizeParams {
	return &PostCustomersCustomerFidAnonymizeParams{
		timeout: timeout,
	}
}

// NewPostCustomersCustomerFidAnonymizeParamsWithContext creates a new PostCustomersCustomerFidAnonymizeParams object
// with the ability to set a context for a request.
func NewPostCustomersCustomerFidAnonymizeParamsWithContext(ctx context.Context) *PostCustomersCustomerFidAnonymizeParams {
	return &PostCustomersCustomerFidAnonymizeParams{
		Context: ctx,
	}
}

// NewPostCustomersCustomerFidAnonymizeParamsWithHTTPClient creates a new PostCustomersCustomerFidAnonymizeParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostCustomersCustomerFidAnonymizeParamsWithHTTPClient(client *http.Client) *PostCustomersCustomerFidAnonymizeParams {
	return &PostCustomersCustomerFidAnonymizeParams{
		HTTPClient: client,
	}
}

/*
PostCustomersCustomerFidAnonymizeParams contains all the parameters to send to the API endpoint

	for the post customers customer fid anonymize operation.

	Typically these are written to a http.Request.
*/
type PostCustomersCustomerFidAnonymizeParams struct {

	// Addresses.
	Addresses *bool

	// Chats.
	Chats *bool

	// Customer.
	Customer *bool

	/* CustomerFid.

	   Customer FID to use
	*/
	CustomerFid string

	// Emails.
	Emails *bool

	// Orders.
	Orders *bool

	// PaymentAccounts.
	PaymentAccounts *bool

	// Phones.
	Phones *bool

	// Tickets.
	Tickets *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post customers customer fid anonymize params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostCustomersCustomerFidAnonymizeParams) WithDefaults() *PostCustomersCustomerFidAnonymizeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post customers customer fid anonymize params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostCustomersCustomerFidAnonymizeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post customers customer fid anonymize params
func (o *PostCustomersCustomerFidAnonymizeParams) WithTimeout(timeout time.Duration) *PostCustomersCustomerFidAnonymizeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post customers customer fid anonymize params
func (o *PostCustomersCustomerFidAnonymizeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post customers customer fid anonymize params
func (o *PostCustomersCustomerFidAnonymizeParams) WithContext(ctx context.Context) *PostCustomersCustomerFidAnonymizeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post customers customer fid anonymize params
func (o *PostCustomersCustomerFidAnonymizeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post customers customer fid anonymize params
func (o *PostCustomersCustomerFidAnonymizeParams) WithHTTPClient(client *http.Client) *PostCustomersCustomerFidAnonymizeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post customers customer fid anonymize params
func (o *PostCustomersCustomerFidAnonymizeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddresses adds the addresses to the post customers customer fid anonymize params
func (o *PostCustomersCustomerFidAnonymizeParams) WithAddresses(addresses *bool) *PostCustomersCustomerFidAnonymizeParams {
	o.SetAddresses(addresses)
	return o
}

// SetAddresses adds the addresses to the post customers customer fid anonymize params
func (o *PostCustomersCustomerFidAnonymizeParams) SetAddresses(addresses *bool) {
	o.Addresses = addresses
}

// WithChats adds the chats to the post customers customer fid anonymize params
func (o *PostCustomersCustomerFidAnonymizeParams) WithChats(chats *bool) *PostCustomersCustomerFidAnonymizeParams {
	o.SetChats(chats)
	return o
}

// SetChats adds the chats to the post customers customer fid anonymize params
func (o *PostCustomersCustomerFidAnonymizeParams) SetChats(chats *bool) {
	o.Chats = chats
}

// WithCustomer adds the customer to the post customers customer fid anonymize params
func (o *PostCustomersCustomerFidAnonymizeParams) WithCustomer(customer *bool) *PostCustomersCustomerFidAnonymizeParams {
	o.SetCustomer(customer)
	return o
}

// SetCustomer adds the customer to the post customers customer fid anonymize params
func (o *PostCustomersCustomerFidAnonymizeParams) SetCustomer(customer *bool) {
	o.Customer = customer
}

// WithCustomerFid adds the customerFid to the post customers customer fid anonymize params
func (o *PostCustomersCustomerFidAnonymizeParams) WithCustomerFid(customerFid string) *PostCustomersCustomerFidAnonymizeParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the post customers customer fid anonymize params
func (o *PostCustomersCustomerFidAnonymizeParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithEmails adds the emails to the post customers customer fid anonymize params
func (o *PostCustomersCustomerFidAnonymizeParams) WithEmails(emails *bool) *PostCustomersCustomerFidAnonymizeParams {
	o.SetEmails(emails)
	return o
}

// SetEmails adds the emails to the post customers customer fid anonymize params
func (o *PostCustomersCustomerFidAnonymizeParams) SetEmails(emails *bool) {
	o.Emails = emails
}

// WithOrders adds the orders to the post customers customer fid anonymize params
func (o *PostCustomersCustomerFidAnonymizeParams) WithOrders(orders *bool) *PostCustomersCustomerFidAnonymizeParams {
	o.SetOrders(orders)
	return o
}

// SetOrders adds the orders to the post customers customer fid anonymize params
func (o *PostCustomersCustomerFidAnonymizeParams) SetOrders(orders *bool) {
	o.Orders = orders
}

// WithPaymentAccounts adds the paymentAccounts to the post customers customer fid anonymize params
func (o *PostCustomersCustomerFidAnonymizeParams) WithPaymentAccounts(paymentAccounts *bool) *PostCustomersCustomerFidAnonymizeParams {
	o.SetPaymentAccounts(paymentAccounts)
	return o
}

// SetPaymentAccounts adds the paymentAccounts to the post customers customer fid anonymize params
func (o *PostCustomersCustomerFidAnonymizeParams) SetPaymentAccounts(paymentAccounts *bool) {
	o.PaymentAccounts = paymentAccounts
}

// WithPhones adds the phones to the post customers customer fid anonymize params
func (o *PostCustomersCustomerFidAnonymizeParams) WithPhones(phones *bool) *PostCustomersCustomerFidAnonymizeParams {
	o.SetPhones(phones)
	return o
}

// SetPhones adds the phones to the post customers customer fid anonymize params
func (o *PostCustomersCustomerFidAnonymizeParams) SetPhones(phones *bool) {
	o.Phones = phones
}

// WithTickets adds the tickets to the post customers customer fid anonymize params
func (o *PostCustomersCustomerFidAnonymizeParams) WithTickets(tickets *bool) *PostCustomersCustomerFidAnonymizeParams {
	o.SetTickets(tickets)
	return o
}

// SetTickets adds the tickets to the post customers customer fid anonymize params
func (o *PostCustomersCustomerFidAnonymizeParams) SetTickets(tickets *bool) {
	o.Tickets = tickets
}

// WriteToRequest writes these params to a swagger request
func (o *PostCustomersCustomerFidAnonymizeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Addresses != nil {

		// form param addresses
		var frAddresses bool
		if o.Addresses != nil {
			frAddresses = *o.Addresses
		}
		fAddresses := swag.FormatBool(frAddresses)
		if fAddresses != "" {
			if err := r.SetFormParam("addresses", fAddresses); err != nil {
				return err
			}
		}
	}

	if o.Chats != nil {

		// form param chats
		var frChats bool
		if o.Chats != nil {
			frChats = *o.Chats
		}
		fChats := swag.FormatBool(frChats)
		if fChats != "" {
			if err := r.SetFormParam("chats", fChats); err != nil {
				return err
			}
		}
	}

	if o.Customer != nil {

		// form param customer
		var frCustomer bool
		if o.Customer != nil {
			frCustomer = *o.Customer
		}
		fCustomer := swag.FormatBool(frCustomer)
		if fCustomer != "" {
			if err := r.SetFormParam("customer", fCustomer); err != nil {
				return err
			}
		}
	}

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}

	if o.Emails != nil {

		// form param emails
		var frEmails bool
		if o.Emails != nil {
			frEmails = *o.Emails
		}
		fEmails := swag.FormatBool(frEmails)
		if fEmails != "" {
			if err := r.SetFormParam("emails", fEmails); err != nil {
				return err
			}
		}
	}

	if o.Orders != nil {

		// form param orders
		var frOrders bool
		if o.Orders != nil {
			frOrders = *o.Orders
		}
		fOrders := swag.FormatBool(frOrders)
		if fOrders != "" {
			if err := r.SetFormParam("orders", fOrders); err != nil {
				return err
			}
		}
	}

	if o.PaymentAccounts != nil {

		// form param paymentAccounts
		var frPaymentAccounts bool
		if o.PaymentAccounts != nil {
			frPaymentAccounts = *o.PaymentAccounts
		}
		fPaymentAccounts := swag.FormatBool(frPaymentAccounts)
		if fPaymentAccounts != "" {
			if err := r.SetFormParam("paymentAccounts", fPaymentAccounts); err != nil {
				return err
			}
		}
	}

	if o.Phones != nil {

		// form param phones
		var frPhones bool
		if o.Phones != nil {
			frPhones = *o.Phones
		}
		fPhones := swag.FormatBool(frPhones)
		if fPhones != "" {
			if err := r.SetFormParam("phones", fPhones); err != nil {
				return err
			}
		}
	}

	if o.Tickets != nil {

		// form param tickets
		var frTickets bool
		if o.Tickets != nil {
			frTickets = *o.Tickets
		}
		fTickets := swag.FormatBool(frTickets)
		if fTickets != "" {
			if err := r.SetFormParam("tickets", fTickets); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
