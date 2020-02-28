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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// NewPutCustomersCustomerFidLocationParams creates a new PutCustomersCustomerFidLocationParams object
// with the default values initialized.
func NewPutCustomersCustomerFidLocationParams() *PutCustomersCustomerFidLocationParams {
	var ()
	return &PutCustomersCustomerFidLocationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomersCustomerFidLocationParamsWithTimeout creates a new PutCustomersCustomerFidLocationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCustomersCustomerFidLocationParamsWithTimeout(timeout time.Duration) *PutCustomersCustomerFidLocationParams {
	var ()
	return &PutCustomersCustomerFidLocationParams{

		timeout: timeout,
	}
}

// NewPutCustomersCustomerFidLocationParamsWithContext creates a new PutCustomersCustomerFidLocationParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCustomersCustomerFidLocationParamsWithContext(ctx context.Context) *PutCustomersCustomerFidLocationParams {
	var ()
	return &PutCustomersCustomerFidLocationParams{

		Context: ctx,
	}
}

// NewPutCustomersCustomerFidLocationParamsWithHTTPClient creates a new PutCustomersCustomerFidLocationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCustomersCustomerFidLocationParamsWithHTTPClient(client *http.Client) *PutCustomersCustomerFidLocationParams {
	var ()
	return &PutCustomersCustomerFidLocationParams{
		HTTPClient: client,
	}
}

/*PutCustomersCustomerFidLocationParams contains all the parameters to send to the API endpoint
for the put customers customer fid location operation typically these are written to a http.Request
*/
type PutCustomersCustomerFidLocationParams struct {

	/*CustomerFid
	  Customer FID to use

	*/
	CustomerFid string
	/*Payload*/
	Payload *models.SetCustomerLocationPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put customers customer fid location params
func (o *PutCustomersCustomerFidLocationParams) WithTimeout(timeout time.Duration) *PutCustomersCustomerFidLocationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put customers customer fid location params
func (o *PutCustomersCustomerFidLocationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put customers customer fid location params
func (o *PutCustomersCustomerFidLocationParams) WithContext(ctx context.Context) *PutCustomersCustomerFidLocationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put customers customer fid location params
func (o *PutCustomersCustomerFidLocationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put customers customer fid location params
func (o *PutCustomersCustomerFidLocationParams) WithHTTPClient(client *http.Client) *PutCustomersCustomerFidLocationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put customers customer fid location params
func (o *PutCustomersCustomerFidLocationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the put customers customer fid location params
func (o *PutCustomersCustomerFidLocationParams) WithCustomerFid(customerFid string) *PutCustomersCustomerFidLocationParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the put customers customer fid location params
func (o *PutCustomersCustomerFidLocationParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithPayload adds the payload to the put customers customer fid location params
func (o *PutCustomersCustomerFidLocationParams) WithPayload(payload *models.SetCustomerLocationPayload) *PutCustomersCustomerFidLocationParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the put customers customer fid location params
func (o *PutCustomersCustomerFidLocationParams) SetPayload(payload *models.SetCustomerLocationPayload) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomersCustomerFidLocationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
