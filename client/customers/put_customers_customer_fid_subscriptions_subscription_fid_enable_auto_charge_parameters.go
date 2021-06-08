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

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams() *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParamsWithTimeout creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams object
// with the ability to set a timeout on a request.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParamsWithTimeout(timeout time.Duration) *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams{
		timeout: timeout,
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParamsWithContext creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams object
// with the ability to set a context for a request.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParamsWithContext(ctx context.Context) *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams{
		Context: ctx,
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParamsWithHTTPClient creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParamsWithHTTPClient(client *http.Client) *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams{
		HTTPClient: client,
	}
}

/* PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams contains all the parameters to send to the API endpoint
   for the put customers customer fid subscriptions subscription fid enable auto charge operation.

   Typically these are written to a http.Request.
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams struct {

	/* CustomerFid.

	   Customer FID to use
	*/
	CustomerFid string

	/* SubscriptionFid.

	   Subscription FID to use
	*/
	SubscriptionFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put customers customer fid subscriptions subscription fid enable auto charge params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams) WithDefaults() *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put customers customer fid subscriptions subscription fid enable auto charge params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put customers customer fid subscriptions subscription fid enable auto charge params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams) WithTimeout(timeout time.Duration) *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put customers customer fid subscriptions subscription fid enable auto charge params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put customers customer fid subscriptions subscription fid enable auto charge params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams) WithContext(ctx context.Context) *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put customers customer fid subscriptions subscription fid enable auto charge params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put customers customer fid subscriptions subscription fid enable auto charge params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams) WithHTTPClient(client *http.Client) *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put customers customer fid subscriptions subscription fid enable auto charge params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the put customers customer fid subscriptions subscription fid enable auto charge params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams) WithCustomerFid(customerFid string) *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the put customers customer fid subscriptions subscription fid enable auto charge params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithSubscriptionFid adds the subscriptionFid to the put customers customer fid subscriptions subscription fid enable auto charge params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams) WithSubscriptionFid(subscriptionFid string) *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams {
	o.SetSubscriptionFid(subscriptionFid)
	return o
}

// SetSubscriptionFid adds the subscriptionFid to the put customers customer fid subscriptions subscription fid enable auto charge params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams) SetSubscriptionFid(subscriptionFid string) {
	o.SubscriptionFid = subscriptionFid
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
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
