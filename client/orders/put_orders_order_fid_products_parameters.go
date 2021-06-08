// Code generated by go-swagger; DO NOT EDIT.

package orders

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

// NewPutOrdersOrderFidProductsParams creates a new PutOrdersOrderFidProductsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutOrdersOrderFidProductsParams() *PutOrdersOrderFidProductsParams {
	return &PutOrdersOrderFidProductsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutOrdersOrderFidProductsParamsWithTimeout creates a new PutOrdersOrderFidProductsParams object
// with the ability to set a timeout on a request.
func NewPutOrdersOrderFidProductsParamsWithTimeout(timeout time.Duration) *PutOrdersOrderFidProductsParams {
	return &PutOrdersOrderFidProductsParams{
		timeout: timeout,
	}
}

// NewPutOrdersOrderFidProductsParamsWithContext creates a new PutOrdersOrderFidProductsParams object
// with the ability to set a context for a request.
func NewPutOrdersOrderFidProductsParamsWithContext(ctx context.Context) *PutOrdersOrderFidProductsParams {
	return &PutOrdersOrderFidProductsParams{
		Context: ctx,
	}
}

// NewPutOrdersOrderFidProductsParamsWithHTTPClient creates a new PutOrdersOrderFidProductsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutOrdersOrderFidProductsParamsWithHTTPClient(client *http.Client) *PutOrdersOrderFidProductsParams {
	return &PutOrdersOrderFidProductsParams{
		HTTPClient: client,
	}
}

/* PutOrdersOrderFidProductsParams contains all the parameters to send to the API endpoint
   for the put orders order fid products operation.

   Typically these are written to a http.Request.
*/
type PutOrdersOrderFidProductsParams struct {

	// OrderFid.
	OrderFid string

	// Payload.
	Payload *models.OrderProductsPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put orders order fid products params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutOrdersOrderFidProductsParams) WithDefaults() *PutOrdersOrderFidProductsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put orders order fid products params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutOrdersOrderFidProductsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put orders order fid products params
func (o *PutOrdersOrderFidProductsParams) WithTimeout(timeout time.Duration) *PutOrdersOrderFidProductsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put orders order fid products params
func (o *PutOrdersOrderFidProductsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put orders order fid products params
func (o *PutOrdersOrderFidProductsParams) WithContext(ctx context.Context) *PutOrdersOrderFidProductsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put orders order fid products params
func (o *PutOrdersOrderFidProductsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put orders order fid products params
func (o *PutOrdersOrderFidProductsParams) WithHTTPClient(client *http.Client) *PutOrdersOrderFidProductsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put orders order fid products params
func (o *PutOrdersOrderFidProductsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderFid adds the orderFid to the put orders order fid products params
func (o *PutOrdersOrderFidProductsParams) WithOrderFid(orderFid string) *PutOrdersOrderFidProductsParams {
	o.SetOrderFid(orderFid)
	return o
}

// SetOrderFid adds the orderFid to the put orders order fid products params
func (o *PutOrdersOrderFidProductsParams) SetOrderFid(orderFid string) {
	o.OrderFid = orderFid
}

// WithPayload adds the payload to the put orders order fid products params
func (o *PutOrdersOrderFidProductsParams) WithPayload(payload *models.OrderProductsPayload) *PutOrdersOrderFidProductsParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the put orders order fid products params
func (o *PutOrdersOrderFidProductsParams) SetPayload(payload *models.OrderProductsPayload) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *PutOrdersOrderFidProductsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param orderFid
	if err := r.SetPathParam("orderFid", o.OrderFid); err != nil {
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
