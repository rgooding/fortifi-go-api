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

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams() *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParamsWithTimeout creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams object
// with the ability to set a timeout on a request.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParamsWithTimeout(timeout time.Duration) *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams{
		timeout: timeout,
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParamsWithContext creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams object
// with the ability to set a context for a request.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParamsWithContext(ctx context.Context) *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams{
		Context: ctx,
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParamsWithHTTPClient creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParamsWithHTTPClient(client *http.Client) *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams{
		HTTPClient: client,
	}
}

/*
PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams contains all the parameters to send to the API endpoint

	for the put customers customer fid subscriptions subscription fid cancel flow flow search abandon operation.

	Typically these are written to a http.Request.
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams struct {

	/* CustomerFid.

	   Customer FID to use
	*/
	CustomerFid string

	/* FlowSearch.

	   The unique code or fid for the flow you wish to retrieve
	*/
	FlowSearch string

	/* SubscriptionFid.

	   Subscription FID to use
	*/
	SubscriptionFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put customers customer fid subscriptions subscription fid cancel flow flow search abandon params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams) WithDefaults() *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put customers customer fid subscriptions subscription fid cancel flow flow search abandon params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put customers customer fid subscriptions subscription fid cancel flow flow search abandon params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams) WithTimeout(timeout time.Duration) *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put customers customer fid subscriptions subscription fid cancel flow flow search abandon params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put customers customer fid subscriptions subscription fid cancel flow flow search abandon params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams) WithContext(ctx context.Context) *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put customers customer fid subscriptions subscription fid cancel flow flow search abandon params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put customers customer fid subscriptions subscription fid cancel flow flow search abandon params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams) WithHTTPClient(client *http.Client) *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put customers customer fid subscriptions subscription fid cancel flow flow search abandon params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the put customers customer fid subscriptions subscription fid cancel flow flow search abandon params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams) WithCustomerFid(customerFid string) *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the put customers customer fid subscriptions subscription fid cancel flow flow search abandon params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithFlowSearch adds the flowSearch to the put customers customer fid subscriptions subscription fid cancel flow flow search abandon params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams) WithFlowSearch(flowSearch string) *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams {
	o.SetFlowSearch(flowSearch)
	return o
}

// SetFlowSearch adds the flowSearch to the put customers customer fid subscriptions subscription fid cancel flow flow search abandon params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams) SetFlowSearch(flowSearch string) {
	o.FlowSearch = flowSearch
}

// WithSubscriptionFid adds the subscriptionFid to the put customers customer fid subscriptions subscription fid cancel flow flow search abandon params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams) WithSubscriptionFid(subscriptionFid string) *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams {
	o.SetSubscriptionFid(subscriptionFid)
	return o
}

// SetSubscriptionFid adds the subscriptionFid to the put customers customer fid subscriptions subscription fid cancel flow flow search abandon params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams) SetSubscriptionFid(subscriptionFid string) {
	o.SubscriptionFid = subscriptionFid
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchAbandonParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}

	// path param flowSearch
	if err := r.SetPathParam("flowSearch", o.FlowSearch); err != nil {
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