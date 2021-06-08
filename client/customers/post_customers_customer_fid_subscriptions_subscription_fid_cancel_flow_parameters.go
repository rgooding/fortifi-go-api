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

	"github.com/fortifi/go-api/models"
)

// NewPostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams creates a new PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams() *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams {
	return &PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParamsWithTimeout creates a new PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams object
// with the ability to set a timeout on a request.
func NewPostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParamsWithTimeout(timeout time.Duration) *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams {
	return &PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams{
		timeout: timeout,
	}
}

// NewPostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParamsWithContext creates a new PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams object
// with the ability to set a context for a request.
func NewPostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParamsWithContext(ctx context.Context) *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams {
	return &PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams{
		Context: ctx,
	}
}

// NewPostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParamsWithHTTPClient creates a new PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParamsWithHTTPClient(client *http.Client) *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams {
	return &PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams{
		HTTPClient: client,
	}
}

/* PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams contains all the parameters to send to the API endpoint
   for the post customers customer fid subscriptions subscription fid cancel flow operation.

   Typically these are written to a http.Request.
*/
type PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams struct {

	/* CustomerFid.

	   Customer FID to use
	*/
	CustomerFid string

	// Payload.
	Payload *models.ActionCancelFlowPayload

	/* SubscriptionFid.

	   Subscription FID to use
	*/
	SubscriptionFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post customers customer fid subscriptions subscription fid cancel flow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams) WithDefaults() *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post customers customer fid subscriptions subscription fid cancel flow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post customers customer fid subscriptions subscription fid cancel flow params
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams) WithTimeout(timeout time.Duration) *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post customers customer fid subscriptions subscription fid cancel flow params
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post customers customer fid subscriptions subscription fid cancel flow params
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams) WithContext(ctx context.Context) *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post customers customer fid subscriptions subscription fid cancel flow params
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post customers customer fid subscriptions subscription fid cancel flow params
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams) WithHTTPClient(client *http.Client) *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post customers customer fid subscriptions subscription fid cancel flow params
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the post customers customer fid subscriptions subscription fid cancel flow params
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams) WithCustomerFid(customerFid string) *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the post customers customer fid subscriptions subscription fid cancel flow params
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithPayload adds the payload to the post customers customer fid subscriptions subscription fid cancel flow params
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams) WithPayload(payload *models.ActionCancelFlowPayload) *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the post customers customer fid subscriptions subscription fid cancel flow params
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams) SetPayload(payload *models.ActionCancelFlowPayload) {
	o.Payload = payload
}

// WithSubscriptionFid adds the subscriptionFid to the post customers customer fid subscriptions subscription fid cancel flow params
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams) WithSubscriptionFid(subscriptionFid string) *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams {
	o.SetSubscriptionFid(subscriptionFid)
	return o
}

// SetSubscriptionFid adds the subscriptionFid to the post customers customer fid subscriptions subscription fid cancel flow params
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams) SetSubscriptionFid(subscriptionFid string) {
	o.SubscriptionFid = subscriptionFid
}

// WriteToRequest writes these params to a swagger request
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}
	if o.Payload != nil {
		if err := r.SetBodyParam(o.Payload); err != nil {
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
