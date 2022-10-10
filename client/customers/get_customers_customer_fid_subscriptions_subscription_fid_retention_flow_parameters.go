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

// NewGetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams creates a new GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams() *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams {
	return &GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParamsWithTimeout creates a new GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams object
// with the ability to set a timeout on a request.
func NewGetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParamsWithTimeout(timeout time.Duration) *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams {
	return &GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams{
		timeout: timeout,
	}
}

// NewGetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParamsWithContext creates a new GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams object
// with the ability to set a context for a request.
func NewGetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParamsWithContext(ctx context.Context) *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams {
	return &GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams{
		Context: ctx,
	}
}

// NewGetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParamsWithHTTPClient creates a new GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParamsWithHTTPClient(client *http.Client) *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams {
	return &GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams{
		HTTPClient: client,
	}
}

/*
GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams contains all the parameters to send to the API endpoint

	for the get customers customer fid subscriptions subscription fid retention flow operation.

	Typically these are written to a http.Request.
*/
type GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams struct {

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

// WithDefaults hydrates default values in the get customers customer fid subscriptions subscription fid retention flow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams) WithDefaults() *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get customers customer fid subscriptions subscription fid retention flow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get customers customer fid subscriptions subscription fid retention flow params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams) WithTimeout(timeout time.Duration) *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get customers customer fid subscriptions subscription fid retention flow params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get customers customer fid subscriptions subscription fid retention flow params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams) WithContext(ctx context.Context) *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get customers customer fid subscriptions subscription fid retention flow params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get customers customer fid subscriptions subscription fid retention flow params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams) WithHTTPClient(client *http.Client) *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get customers customer fid subscriptions subscription fid retention flow params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the get customers customer fid subscriptions subscription fid retention flow params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams) WithCustomerFid(customerFid string) *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the get customers customer fid subscriptions subscription fid retention flow params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithSubscriptionFid adds the subscriptionFid to the get customers customer fid subscriptions subscription fid retention flow params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams) WithSubscriptionFid(subscriptionFid string) *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams {
	o.SetSubscriptionFid(subscriptionFid)
	return o
}

// SetSubscriptionFid adds the subscriptionFid to the get customers customer fid subscriptions subscription fid retention flow params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams) SetSubscriptionFid(subscriptionFid string) {
	o.SubscriptionFid = subscriptionFid
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
