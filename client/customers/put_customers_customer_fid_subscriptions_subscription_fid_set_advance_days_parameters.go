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

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams() *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParamsWithTimeout creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams object
// with the ability to set a timeout on a request.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParamsWithTimeout(timeout time.Duration) *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams{
		timeout: timeout,
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParamsWithContext creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams object
// with the ability to set a context for a request.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParamsWithContext(ctx context.Context) *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams{
		Context: ctx,
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParamsWithHTTPClient creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParamsWithHTTPClient(client *http.Client) *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams{
		HTTPClient: client,
	}
}

/*
PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams contains all the parameters to send to the API endpoint

	for the put customers customer fid subscriptions subscription fid set advance days operation.

	Typically these are written to a http.Request.
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams struct {

	/* CustomerFid.

	   Customer FID to use
	*/
	CustomerFid string

	// Days.
	//
	// Format: int32
	Days int32

	/* SubscriptionFid.

	   Subscription FID to use
	*/
	SubscriptionFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put customers customer fid subscriptions subscription fid set advance days params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams) WithDefaults() *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put customers customer fid subscriptions subscription fid set advance days params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put customers customer fid subscriptions subscription fid set advance days params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams) WithTimeout(timeout time.Duration) *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put customers customer fid subscriptions subscription fid set advance days params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put customers customer fid subscriptions subscription fid set advance days params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams) WithContext(ctx context.Context) *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put customers customer fid subscriptions subscription fid set advance days params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put customers customer fid subscriptions subscription fid set advance days params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams) WithHTTPClient(client *http.Client) *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put customers customer fid subscriptions subscription fid set advance days params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the put customers customer fid subscriptions subscription fid set advance days params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams) WithCustomerFid(customerFid string) *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the put customers customer fid subscriptions subscription fid set advance days params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithDays adds the days to the put customers customer fid subscriptions subscription fid set advance days params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams) WithDays(days int32) *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams {
	o.SetDays(days)
	return o
}

// SetDays adds the days to the put customers customer fid subscriptions subscription fid set advance days params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams) SetDays(days int32) {
	o.Days = days
}

// WithSubscriptionFid adds the subscriptionFid to the put customers customer fid subscriptions subscription fid set advance days params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams) WithSubscriptionFid(subscriptionFid string) *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams {
	o.SetSubscriptionFid(subscriptionFid)
	return o
}

// SetSubscriptionFid adds the subscriptionFid to the put customers customer fid subscriptions subscription fid set advance days params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams) SetSubscriptionFid(subscriptionFid string) {
	o.SubscriptionFid = subscriptionFid
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidSetAdvanceDaysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}

	// form param days
	frDays := o.Days
	fDays := swag.FormatInt32(frDays)
	if fDays != "" {
		if err := r.SetFormParam("days", fDays); err != nil {
			return err
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
