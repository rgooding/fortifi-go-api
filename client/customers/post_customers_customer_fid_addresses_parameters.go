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

// NewPostCustomersCustomerFidAddressesParams creates a new PostCustomersCustomerFidAddressesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostCustomersCustomerFidAddressesParams() *PostCustomersCustomerFidAddressesParams {
	return &PostCustomersCustomerFidAddressesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostCustomersCustomerFidAddressesParamsWithTimeout creates a new PostCustomersCustomerFidAddressesParams object
// with the ability to set a timeout on a request.
func NewPostCustomersCustomerFidAddressesParamsWithTimeout(timeout time.Duration) *PostCustomersCustomerFidAddressesParams {
	return &PostCustomersCustomerFidAddressesParams{
		timeout: timeout,
	}
}

// NewPostCustomersCustomerFidAddressesParamsWithContext creates a new PostCustomersCustomerFidAddressesParams object
// with the ability to set a context for a request.
func NewPostCustomersCustomerFidAddressesParamsWithContext(ctx context.Context) *PostCustomersCustomerFidAddressesParams {
	return &PostCustomersCustomerFidAddressesParams{
		Context: ctx,
	}
}

// NewPostCustomersCustomerFidAddressesParamsWithHTTPClient creates a new PostCustomersCustomerFidAddressesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostCustomersCustomerFidAddressesParamsWithHTTPClient(client *http.Client) *PostCustomersCustomerFidAddressesParams {
	return &PostCustomersCustomerFidAddressesParams{
		HTTPClient: client,
	}
}

/* PostCustomersCustomerFidAddressesParams contains all the parameters to send to the API endpoint
   for the post customers customer fid addresses operation.

   Typically these are written to a http.Request.
*/
type PostCustomersCustomerFidAddressesParams struct {

	/* CustomerFid.

	   Customer FID to use
	*/
	CustomerFid string

	// Payload.
	Payload *models.AddressPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post customers customer fid addresses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostCustomersCustomerFidAddressesParams) WithDefaults() *PostCustomersCustomerFidAddressesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post customers customer fid addresses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostCustomersCustomerFidAddressesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post customers customer fid addresses params
func (o *PostCustomersCustomerFidAddressesParams) WithTimeout(timeout time.Duration) *PostCustomersCustomerFidAddressesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post customers customer fid addresses params
func (o *PostCustomersCustomerFidAddressesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post customers customer fid addresses params
func (o *PostCustomersCustomerFidAddressesParams) WithContext(ctx context.Context) *PostCustomersCustomerFidAddressesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post customers customer fid addresses params
func (o *PostCustomersCustomerFidAddressesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post customers customer fid addresses params
func (o *PostCustomersCustomerFidAddressesParams) WithHTTPClient(client *http.Client) *PostCustomersCustomerFidAddressesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post customers customer fid addresses params
func (o *PostCustomersCustomerFidAddressesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the post customers customer fid addresses params
func (o *PostCustomersCustomerFidAddressesParams) WithCustomerFid(customerFid string) *PostCustomersCustomerFidAddressesParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the post customers customer fid addresses params
func (o *PostCustomersCustomerFidAddressesParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithPayload adds the payload to the post customers customer fid addresses params
func (o *PostCustomersCustomerFidAddressesParams) WithPayload(payload *models.AddressPayload) *PostCustomersCustomerFidAddressesParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the post customers customer fid addresses params
func (o *PostCustomersCustomerFidAddressesParams) SetPayload(payload *models.AddressPayload) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *PostCustomersCustomerFidAddressesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
