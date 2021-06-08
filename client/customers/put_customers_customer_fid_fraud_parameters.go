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

// NewPutCustomersCustomerFidFraudParams creates a new PutCustomersCustomerFidFraudParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutCustomersCustomerFidFraudParams() *PutCustomersCustomerFidFraudParams {
	return &PutCustomersCustomerFidFraudParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomersCustomerFidFraudParamsWithTimeout creates a new PutCustomersCustomerFidFraudParams object
// with the ability to set a timeout on a request.
func NewPutCustomersCustomerFidFraudParamsWithTimeout(timeout time.Duration) *PutCustomersCustomerFidFraudParams {
	return &PutCustomersCustomerFidFraudParams{
		timeout: timeout,
	}
}

// NewPutCustomersCustomerFidFraudParamsWithContext creates a new PutCustomersCustomerFidFraudParams object
// with the ability to set a context for a request.
func NewPutCustomersCustomerFidFraudParamsWithContext(ctx context.Context) *PutCustomersCustomerFidFraudParams {
	return &PutCustomersCustomerFidFraudParams{
		Context: ctx,
	}
}

// NewPutCustomersCustomerFidFraudParamsWithHTTPClient creates a new PutCustomersCustomerFidFraudParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutCustomersCustomerFidFraudParamsWithHTTPClient(client *http.Client) *PutCustomersCustomerFidFraudParams {
	return &PutCustomersCustomerFidFraudParams{
		HTTPClient: client,
	}
}

/* PutCustomersCustomerFidFraudParams contains all the parameters to send to the API endpoint
   for the put customers customer fid fraud operation.

   Typically these are written to a http.Request.
*/
type PutCustomersCustomerFidFraudParams struct {

	/* CustomerFid.

	   Customer FID to use
	*/
	CustomerFid string

	/* IsoTime.

	   Time in ISO 8601 standard with optional fractions of a second e.g 2015-12-05T13:11:59.888Z

	   Format: date-time
	*/
	IsoTime *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put customers customer fid fraud params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutCustomersCustomerFidFraudParams) WithDefaults() *PutCustomersCustomerFidFraudParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put customers customer fid fraud params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutCustomersCustomerFidFraudParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put customers customer fid fraud params
func (o *PutCustomersCustomerFidFraudParams) WithTimeout(timeout time.Duration) *PutCustomersCustomerFidFraudParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put customers customer fid fraud params
func (o *PutCustomersCustomerFidFraudParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put customers customer fid fraud params
func (o *PutCustomersCustomerFidFraudParams) WithContext(ctx context.Context) *PutCustomersCustomerFidFraudParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put customers customer fid fraud params
func (o *PutCustomersCustomerFidFraudParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put customers customer fid fraud params
func (o *PutCustomersCustomerFidFraudParams) WithHTTPClient(client *http.Client) *PutCustomersCustomerFidFraudParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put customers customer fid fraud params
func (o *PutCustomersCustomerFidFraudParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the put customers customer fid fraud params
func (o *PutCustomersCustomerFidFraudParams) WithCustomerFid(customerFid string) *PutCustomersCustomerFidFraudParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the put customers customer fid fraud params
func (o *PutCustomersCustomerFidFraudParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithIsoTime adds the isoTime to the put customers customer fid fraud params
func (o *PutCustomersCustomerFidFraudParams) WithIsoTime(isoTime *strfmt.DateTime) *PutCustomersCustomerFidFraudParams {
	o.SetIsoTime(isoTime)
	return o
}

// SetIsoTime adds the isoTime to the put customers customer fid fraud params
func (o *PutCustomersCustomerFidFraudParams) SetIsoTime(isoTime *strfmt.DateTime) {
	o.IsoTime = isoTime
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomersCustomerFidFraudParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}

	if o.IsoTime != nil {

		// form param isoTime
		var frIsoTime strfmt.DateTime
		if o.IsoTime != nil {
			frIsoTime = *o.IsoTime
		}
		fIsoTime := frIsoTime.String()
		if fIsoTime != "" {
			if err := r.SetFormParam("isoTime", fIsoTime); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
