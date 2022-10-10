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
)

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams() *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParamsWithTimeout creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams object
// with the ability to set a timeout on a request.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParamsWithTimeout(timeout time.Duration) *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams{
		timeout: timeout,
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParamsWithContext creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams object
// with the ability to set a context for a request.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParamsWithContext(ctx context.Context) *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams{
		Context: ctx,
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParamsWithHTTPClient creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParamsWithHTTPClient(client *http.Client) *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams{
		HTTPClient: client,
	}
}

/*
PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams contains all the parameters to send to the API endpoint

	for the put customers customer fid subscriptions subscription fid set payment account operation.

	Typically these are written to a http.Request.
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams struct {

	/* CustomerFid.

	   Customer FID to use
	*/
	CustomerFid string

	// PaymentAccountFid.
	PaymentAccountFid *string

	/* SubscriptionFid.

	   Subscription FID to use
	*/
	SubscriptionFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put customers customer fid subscriptions subscription fid set payment account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams) WithDefaults() *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put customers customer fid subscriptions subscription fid set payment account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put customers customer fid subscriptions subscription fid set payment account params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams) WithTimeout(timeout time.Duration) *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put customers customer fid subscriptions subscription fid set payment account params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put customers customer fid subscriptions subscription fid set payment account params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams) WithContext(ctx context.Context) *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put customers customer fid subscriptions subscription fid set payment account params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put customers customer fid subscriptions subscription fid set payment account params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams) WithHTTPClient(client *http.Client) *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put customers customer fid subscriptions subscription fid set payment account params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the put customers customer fid subscriptions subscription fid set payment account params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams) WithCustomerFid(customerFid string) *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the put customers customer fid subscriptions subscription fid set payment account params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithPaymentAccountFid adds the paymentAccountFid to the put customers customer fid subscriptions subscription fid set payment account params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams) WithPaymentAccountFid(paymentAccountFid *string) *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams {
	o.SetPaymentAccountFid(paymentAccountFid)
	return o
}

// SetPaymentAccountFid adds the paymentAccountFid to the put customers customer fid subscriptions subscription fid set payment account params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams) SetPaymentAccountFid(paymentAccountFid *string) {
	o.PaymentAccountFid = paymentAccountFid
}

// WithSubscriptionFid adds the subscriptionFid to the put customers customer fid subscriptions subscription fid set payment account params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams) WithSubscriptionFid(subscriptionFid string) *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams {
	o.SetSubscriptionFid(subscriptionFid)
	return o
}

// SetSubscriptionFid adds the subscriptionFid to the put customers customer fid subscriptions subscription fid set payment account params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams) SetSubscriptionFid(subscriptionFid string) {
	o.SubscriptionFid = subscriptionFid
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetPaymentAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}

	if o.PaymentAccountFid != nil {

		// form param paymentAccountFid
		var frPaymentAccountFid string
		if o.PaymentAccountFid != nil {
			frPaymentAccountFid = *o.PaymentAccountFid
		}
		fPaymentAccountFid := frPaymentAccountFid
		if fPaymentAccountFid != "" {
			if err := r.SetFormParam("paymentAccountFid", fPaymentAccountFid); err != nil {
				return err
			}
		}
	}

	// path param subscriptionFid
	if err := r.SetPathParam("subscriptionFid", o.SubscriptionFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
