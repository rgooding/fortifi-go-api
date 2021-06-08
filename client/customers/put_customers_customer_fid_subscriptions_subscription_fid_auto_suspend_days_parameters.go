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

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams() *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParamsWithTimeout creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams object
// with the ability to set a timeout on a request.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParamsWithTimeout(timeout time.Duration) *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams{
		timeout: timeout,
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParamsWithContext creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams object
// with the ability to set a context for a request.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParamsWithContext(ctx context.Context) *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams{
		Context: ctx,
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParamsWithHTTPClient creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParamsWithHTTPClient(client *http.Client) *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams{
		HTTPClient: client,
	}
}

/* PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams contains all the parameters to send to the API endpoint
   for the put customers customer fid subscriptions subscription fid auto suspend days operation.

   Typically these are written to a http.Request.
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams struct {

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

// WithDefaults hydrates default values in the put customers customer fid subscriptions subscription fid auto suspend days params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams) WithDefaults() *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put customers customer fid subscriptions subscription fid auto suspend days params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put customers customer fid subscriptions subscription fid auto suspend days params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams) WithTimeout(timeout time.Duration) *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put customers customer fid subscriptions subscription fid auto suspend days params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put customers customer fid subscriptions subscription fid auto suspend days params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams) WithContext(ctx context.Context) *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put customers customer fid subscriptions subscription fid auto suspend days params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put customers customer fid subscriptions subscription fid auto suspend days params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams) WithHTTPClient(client *http.Client) *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put customers customer fid subscriptions subscription fid auto suspend days params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the put customers customer fid subscriptions subscription fid auto suspend days params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams) WithCustomerFid(customerFid string) *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the put customers customer fid subscriptions subscription fid auto suspend days params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithDays adds the days to the put customers customer fid subscriptions subscription fid auto suspend days params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams) WithDays(days int32) *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams {
	o.SetDays(days)
	return o
}

// SetDays adds the days to the put customers customer fid subscriptions subscription fid auto suspend days params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams) SetDays(days int32) {
	o.Days = days
}

// WithSubscriptionFid adds the subscriptionFid to the put customers customer fid subscriptions subscription fid auto suspend days params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams) WithSubscriptionFid(subscriptionFid string) *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams {
	o.SetSubscriptionFid(subscriptionFid)
	return o
}

// SetSubscriptionFid adds the subscriptionFid to the put customers customer fid subscriptions subscription fid auto suspend days params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams) SetSubscriptionFid(subscriptionFid string) {
	o.SubscriptionFid = subscriptionFid
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoSuspendDaysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
