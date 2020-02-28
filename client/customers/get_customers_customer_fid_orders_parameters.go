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

// NewGetCustomersCustomerFidOrdersParams creates a new GetCustomersCustomerFidOrdersParams object
// with the default values initialized.
func NewGetCustomersCustomerFidOrdersParams() *GetCustomersCustomerFidOrdersParams {
	var ()
	return &GetCustomersCustomerFidOrdersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomersCustomerFidOrdersParamsWithTimeout creates a new GetCustomersCustomerFidOrdersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCustomersCustomerFidOrdersParamsWithTimeout(timeout time.Duration) *GetCustomersCustomerFidOrdersParams {
	var ()
	return &GetCustomersCustomerFidOrdersParams{

		timeout: timeout,
	}
}

// NewGetCustomersCustomerFidOrdersParamsWithContext creates a new GetCustomersCustomerFidOrdersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCustomersCustomerFidOrdersParamsWithContext(ctx context.Context) *GetCustomersCustomerFidOrdersParams {
	var ()
	return &GetCustomersCustomerFidOrdersParams{

		Context: ctx,
	}
}

// NewGetCustomersCustomerFidOrdersParamsWithHTTPClient creates a new GetCustomersCustomerFidOrdersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCustomersCustomerFidOrdersParamsWithHTTPClient(client *http.Client) *GetCustomersCustomerFidOrdersParams {
	var ()
	return &GetCustomersCustomerFidOrdersParams{
		HTTPClient: client,
	}
}

/*GetCustomersCustomerFidOrdersParams contains all the parameters to send to the API endpoint
for the get customers customer fid orders operation typically these are written to a http.Request
*/
type GetCustomersCustomerFidOrdersParams struct {

	/*CustomerFid
	  Customer FID to use

	*/
	CustomerFid string
	/*ExternalReference*/
	ExternalReference *string
	/*Filter*/
	Filter *string
	/*State*/
	State *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get customers customer fid orders params
func (o *GetCustomersCustomerFidOrdersParams) WithTimeout(timeout time.Duration) *GetCustomersCustomerFidOrdersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get customers customer fid orders params
func (o *GetCustomersCustomerFidOrdersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get customers customer fid orders params
func (o *GetCustomersCustomerFidOrdersParams) WithContext(ctx context.Context) *GetCustomersCustomerFidOrdersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get customers customer fid orders params
func (o *GetCustomersCustomerFidOrdersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get customers customer fid orders params
func (o *GetCustomersCustomerFidOrdersParams) WithHTTPClient(client *http.Client) *GetCustomersCustomerFidOrdersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get customers customer fid orders params
func (o *GetCustomersCustomerFidOrdersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the get customers customer fid orders params
func (o *GetCustomersCustomerFidOrdersParams) WithCustomerFid(customerFid string) *GetCustomersCustomerFidOrdersParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the get customers customer fid orders params
func (o *GetCustomersCustomerFidOrdersParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithExternalReference adds the externalReference to the get customers customer fid orders params
func (o *GetCustomersCustomerFidOrdersParams) WithExternalReference(externalReference *string) *GetCustomersCustomerFidOrdersParams {
	o.SetExternalReference(externalReference)
	return o
}

// SetExternalReference adds the externalReference to the get customers customer fid orders params
func (o *GetCustomersCustomerFidOrdersParams) SetExternalReference(externalReference *string) {
	o.ExternalReference = externalReference
}

// WithFilter adds the filter to the get customers customer fid orders params
func (o *GetCustomersCustomerFidOrdersParams) WithFilter(filter *string) *GetCustomersCustomerFidOrdersParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get customers customer fid orders params
func (o *GetCustomersCustomerFidOrdersParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithState adds the state to the get customers customer fid orders params
func (o *GetCustomersCustomerFidOrdersParams) WithState(state *string) *GetCustomersCustomerFidOrdersParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the get customers customer fid orders params
func (o *GetCustomersCustomerFidOrdersParams) SetState(state *string) {
	o.State = state
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomersCustomerFidOrdersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}

	if o.ExternalReference != nil {

		// query param externalReference
		var qrExternalReference string
		if o.ExternalReference != nil {
			qrExternalReference = *o.ExternalReference
		}
		qExternalReference := qrExternalReference
		if qExternalReference != "" {
			if err := r.SetQueryParam("externalReference", qExternalReference); err != nil {
				return err
			}
		}

	}

	if o.Filter != nil {

		// query param filter
		var qrFilter string
		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {
			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}

	}

	if o.State != nil {

		// query param state
		var qrState string
		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {
			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}