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
)

// NewDeleteCustomersCustomerFidContactsContactFidParams creates a new DeleteCustomersCustomerFidContactsContactFidParams object
// with the default values initialized.
func NewDeleteCustomersCustomerFidContactsContactFidParams() *DeleteCustomersCustomerFidContactsContactFidParams {
	var ()
	return &DeleteCustomersCustomerFidContactsContactFidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCustomersCustomerFidContactsContactFidParamsWithTimeout creates a new DeleteCustomersCustomerFidContactsContactFidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCustomersCustomerFidContactsContactFidParamsWithTimeout(timeout time.Duration) *DeleteCustomersCustomerFidContactsContactFidParams {
	var ()
	return &DeleteCustomersCustomerFidContactsContactFidParams{

		timeout: timeout,
	}
}

// NewDeleteCustomersCustomerFidContactsContactFidParamsWithContext creates a new DeleteCustomersCustomerFidContactsContactFidParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCustomersCustomerFidContactsContactFidParamsWithContext(ctx context.Context) *DeleteCustomersCustomerFidContactsContactFidParams {
	var ()
	return &DeleteCustomersCustomerFidContactsContactFidParams{

		Context: ctx,
	}
}

// NewDeleteCustomersCustomerFidContactsContactFidParamsWithHTTPClient creates a new DeleteCustomersCustomerFidContactsContactFidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCustomersCustomerFidContactsContactFidParamsWithHTTPClient(client *http.Client) *DeleteCustomersCustomerFidContactsContactFidParams {
	var ()
	return &DeleteCustomersCustomerFidContactsContactFidParams{
		HTTPClient: client,
	}
}

/*DeleteCustomersCustomerFidContactsContactFidParams contains all the parameters to send to the API endpoint
for the delete customers customer fid contacts contact fid operation typically these are written to a http.Request
*/
type DeleteCustomersCustomerFidContactsContactFidParams struct {

	/*ContactFid
	  Contact FID to use

	*/
	ContactFid string
	/*CustomerFid
	  Customer FID to use

	*/
	CustomerFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete customers customer fid contacts contact fid params
func (o *DeleteCustomersCustomerFidContactsContactFidParams) WithTimeout(timeout time.Duration) *DeleteCustomersCustomerFidContactsContactFidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete customers customer fid contacts contact fid params
func (o *DeleteCustomersCustomerFidContactsContactFidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete customers customer fid contacts contact fid params
func (o *DeleteCustomersCustomerFidContactsContactFidParams) WithContext(ctx context.Context) *DeleteCustomersCustomerFidContactsContactFidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete customers customer fid contacts contact fid params
func (o *DeleteCustomersCustomerFidContactsContactFidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete customers customer fid contacts contact fid params
func (o *DeleteCustomersCustomerFidContactsContactFidParams) WithHTTPClient(client *http.Client) *DeleteCustomersCustomerFidContactsContactFidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete customers customer fid contacts contact fid params
func (o *DeleteCustomersCustomerFidContactsContactFidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactFid adds the contactFid to the delete customers customer fid contacts contact fid params
func (o *DeleteCustomersCustomerFidContactsContactFidParams) WithContactFid(contactFid string) *DeleteCustomersCustomerFidContactsContactFidParams {
	o.SetContactFid(contactFid)
	return o
}

// SetContactFid adds the contactFid to the delete customers customer fid contacts contact fid params
func (o *DeleteCustomersCustomerFidContactsContactFidParams) SetContactFid(contactFid string) {
	o.ContactFid = contactFid
}

// WithCustomerFid adds the customerFid to the delete customers customer fid contacts contact fid params
func (o *DeleteCustomersCustomerFidContactsContactFidParams) WithCustomerFid(customerFid string) *DeleteCustomersCustomerFidContactsContactFidParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the delete customers customer fid contacts contact fid params
func (o *DeleteCustomersCustomerFidContactsContactFidParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCustomersCustomerFidContactsContactFidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contactFid
	if err := r.SetPathParam("contactFid", o.ContactFid); err != nil {
		return err
	}

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
