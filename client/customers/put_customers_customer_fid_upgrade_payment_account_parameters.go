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

// NewPutCustomersCustomerFidUpgradePaymentAccountParams creates a new PutCustomersCustomerFidUpgradePaymentAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutCustomersCustomerFidUpgradePaymentAccountParams() *PutCustomersCustomerFidUpgradePaymentAccountParams {
	return &PutCustomersCustomerFidUpgradePaymentAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomersCustomerFidUpgradePaymentAccountParamsWithTimeout creates a new PutCustomersCustomerFidUpgradePaymentAccountParams object
// with the ability to set a timeout on a request.
func NewPutCustomersCustomerFidUpgradePaymentAccountParamsWithTimeout(timeout time.Duration) *PutCustomersCustomerFidUpgradePaymentAccountParams {
	return &PutCustomersCustomerFidUpgradePaymentAccountParams{
		timeout: timeout,
	}
}

// NewPutCustomersCustomerFidUpgradePaymentAccountParamsWithContext creates a new PutCustomersCustomerFidUpgradePaymentAccountParams object
// with the ability to set a context for a request.
func NewPutCustomersCustomerFidUpgradePaymentAccountParamsWithContext(ctx context.Context) *PutCustomersCustomerFidUpgradePaymentAccountParams {
	return &PutCustomersCustomerFidUpgradePaymentAccountParams{
		Context: ctx,
	}
}

// NewPutCustomersCustomerFidUpgradePaymentAccountParamsWithHTTPClient creates a new PutCustomersCustomerFidUpgradePaymentAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutCustomersCustomerFidUpgradePaymentAccountParamsWithHTTPClient(client *http.Client) *PutCustomersCustomerFidUpgradePaymentAccountParams {
	return &PutCustomersCustomerFidUpgradePaymentAccountParams{
		HTTPClient: client,
	}
}

/* PutCustomersCustomerFidUpgradePaymentAccountParams contains all the parameters to send to the API endpoint
   for the put customers customer fid upgrade payment account operation.

   Typically these are written to a http.Request.
*/
type PutCustomersCustomerFidUpgradePaymentAccountParams struct {

	/* CustomerFid.

	   Customer FID to use
	*/
	CustomerFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put customers customer fid upgrade payment account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutCustomersCustomerFidUpgradePaymentAccountParams) WithDefaults() *PutCustomersCustomerFidUpgradePaymentAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put customers customer fid upgrade payment account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutCustomersCustomerFidUpgradePaymentAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put customers customer fid upgrade payment account params
func (o *PutCustomersCustomerFidUpgradePaymentAccountParams) WithTimeout(timeout time.Duration) *PutCustomersCustomerFidUpgradePaymentAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put customers customer fid upgrade payment account params
func (o *PutCustomersCustomerFidUpgradePaymentAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put customers customer fid upgrade payment account params
func (o *PutCustomersCustomerFidUpgradePaymentAccountParams) WithContext(ctx context.Context) *PutCustomersCustomerFidUpgradePaymentAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put customers customer fid upgrade payment account params
func (o *PutCustomersCustomerFidUpgradePaymentAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put customers customer fid upgrade payment account params
func (o *PutCustomersCustomerFidUpgradePaymentAccountParams) WithHTTPClient(client *http.Client) *PutCustomersCustomerFidUpgradePaymentAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put customers customer fid upgrade payment account params
func (o *PutCustomersCustomerFidUpgradePaymentAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the put customers customer fid upgrade payment account params
func (o *PutCustomersCustomerFidUpgradePaymentAccountParams) WithCustomerFid(customerFid string) *PutCustomersCustomerFidUpgradePaymentAccountParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the put customers customer fid upgrade payment account params
func (o *PutCustomersCustomerFidUpgradePaymentAccountParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomersCustomerFidUpgradePaymentAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
