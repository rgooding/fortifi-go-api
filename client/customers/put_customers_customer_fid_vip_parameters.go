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

// NewPutCustomersCustomerFidVipParams creates a new PutCustomersCustomerFidVipParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutCustomersCustomerFidVipParams() *PutCustomersCustomerFidVipParams {
	return &PutCustomersCustomerFidVipParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomersCustomerFidVipParamsWithTimeout creates a new PutCustomersCustomerFidVipParams object
// with the ability to set a timeout on a request.
func NewPutCustomersCustomerFidVipParamsWithTimeout(timeout time.Duration) *PutCustomersCustomerFidVipParams {
	return &PutCustomersCustomerFidVipParams{
		timeout: timeout,
	}
}

// NewPutCustomersCustomerFidVipParamsWithContext creates a new PutCustomersCustomerFidVipParams object
// with the ability to set a context for a request.
func NewPutCustomersCustomerFidVipParamsWithContext(ctx context.Context) *PutCustomersCustomerFidVipParams {
	return &PutCustomersCustomerFidVipParams{
		Context: ctx,
	}
}

// NewPutCustomersCustomerFidVipParamsWithHTTPClient creates a new PutCustomersCustomerFidVipParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutCustomersCustomerFidVipParamsWithHTTPClient(client *http.Client) *PutCustomersCustomerFidVipParams {
	return &PutCustomersCustomerFidVipParams{
		HTTPClient: client,
	}
}

/*
PutCustomersCustomerFidVipParams contains all the parameters to send to the API endpoint

	for the put customers customer fid vip operation.

	Typically these are written to a http.Request.
*/
type PutCustomersCustomerFidVipParams struct {

	/* CustomerFid.

	   Customer FID to use
	*/
	CustomerFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put customers customer fid vip params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutCustomersCustomerFidVipParams) WithDefaults() *PutCustomersCustomerFidVipParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put customers customer fid vip params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutCustomersCustomerFidVipParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put customers customer fid vip params
func (o *PutCustomersCustomerFidVipParams) WithTimeout(timeout time.Duration) *PutCustomersCustomerFidVipParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put customers customer fid vip params
func (o *PutCustomersCustomerFidVipParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put customers customer fid vip params
func (o *PutCustomersCustomerFidVipParams) WithContext(ctx context.Context) *PutCustomersCustomerFidVipParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put customers customer fid vip params
func (o *PutCustomersCustomerFidVipParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put customers customer fid vip params
func (o *PutCustomersCustomerFidVipParams) WithHTTPClient(client *http.Client) *PutCustomersCustomerFidVipParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put customers customer fid vip params
func (o *PutCustomersCustomerFidVipParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the put customers customer fid vip params
func (o *PutCustomersCustomerFidVipParams) WithCustomerFid(customerFid string) *PutCustomersCustomerFidVipParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the put customers customer fid vip params
func (o *PutCustomersCustomerFidVipParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomersCustomerFidVipParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
