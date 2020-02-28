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

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteOrdersOrderFidProductsOrderProductFidParams creates a new DeleteOrdersOrderFidProductsOrderProductFidParams object
// with the default values initialized.
func NewDeleteOrdersOrderFidProductsOrderProductFidParams() *DeleteOrdersOrderFidProductsOrderProductFidParams {
	var ()
	return &DeleteOrdersOrderFidProductsOrderProductFidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOrdersOrderFidProductsOrderProductFidParamsWithTimeout creates a new DeleteOrdersOrderFidProductsOrderProductFidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteOrdersOrderFidProductsOrderProductFidParamsWithTimeout(timeout time.Duration) *DeleteOrdersOrderFidProductsOrderProductFidParams {
	var ()
	return &DeleteOrdersOrderFidProductsOrderProductFidParams{

		timeout: timeout,
	}
}

// NewDeleteOrdersOrderFidProductsOrderProductFidParamsWithContext creates a new DeleteOrdersOrderFidProductsOrderProductFidParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteOrdersOrderFidProductsOrderProductFidParamsWithContext(ctx context.Context) *DeleteOrdersOrderFidProductsOrderProductFidParams {
	var ()
	return &DeleteOrdersOrderFidProductsOrderProductFidParams{

		Context: ctx,
	}
}

// NewDeleteOrdersOrderFidProductsOrderProductFidParamsWithHTTPClient creates a new DeleteOrdersOrderFidProductsOrderProductFidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteOrdersOrderFidProductsOrderProductFidParamsWithHTTPClient(client *http.Client) *DeleteOrdersOrderFidProductsOrderProductFidParams {
	var ()
	return &DeleteOrdersOrderFidProductsOrderProductFidParams{
		HTTPClient: client,
	}
}

/*DeleteOrdersOrderFidProductsOrderProductFidParams contains all the parameters to send to the API endpoint
for the delete orders order fid products order product fid operation typically these are written to a http.Request
*/
type DeleteOrdersOrderFidProductsOrderProductFidParams struct {

	/*OrderFid*/
	OrderFid string
	/*OrderProductFid*/
	OrderProductFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete orders order fid products order product fid params
func (o *DeleteOrdersOrderFidProductsOrderProductFidParams) WithTimeout(timeout time.Duration) *DeleteOrdersOrderFidProductsOrderProductFidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete orders order fid products order product fid params
func (o *DeleteOrdersOrderFidProductsOrderProductFidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete orders order fid products order product fid params
func (o *DeleteOrdersOrderFidProductsOrderProductFidParams) WithContext(ctx context.Context) *DeleteOrdersOrderFidProductsOrderProductFidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete orders order fid products order product fid params
func (o *DeleteOrdersOrderFidProductsOrderProductFidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete orders order fid products order product fid params
func (o *DeleteOrdersOrderFidProductsOrderProductFidParams) WithHTTPClient(client *http.Client) *DeleteOrdersOrderFidProductsOrderProductFidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete orders order fid products order product fid params
func (o *DeleteOrdersOrderFidProductsOrderProductFidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderFid adds the orderFid to the delete orders order fid products order product fid params
func (o *DeleteOrdersOrderFidProductsOrderProductFidParams) WithOrderFid(orderFid string) *DeleteOrdersOrderFidProductsOrderProductFidParams {
	o.SetOrderFid(orderFid)
	return o
}

// SetOrderFid adds the orderFid to the delete orders order fid products order product fid params
func (o *DeleteOrdersOrderFidProductsOrderProductFidParams) SetOrderFid(orderFid string) {
	o.OrderFid = orderFid
}

// WithOrderProductFid adds the orderProductFid to the delete orders order fid products order product fid params
func (o *DeleteOrdersOrderFidProductsOrderProductFidParams) WithOrderProductFid(orderProductFid string) *DeleteOrdersOrderFidProductsOrderProductFidParams {
	o.SetOrderProductFid(orderProductFid)
	return o
}

// SetOrderProductFid adds the orderProductFid to the delete orders order fid products order product fid params
func (o *DeleteOrdersOrderFidProductsOrderProductFidParams) SetOrderProductFid(orderProductFid string) {
	o.OrderProductFid = orderProductFid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOrdersOrderFidProductsOrderProductFidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param orderFid
	if err := r.SetPathParam("orderFid", o.OrderFid); err != nil {
		return err
	}

	// path param orderProductFid
	if err := r.SetPathParam("orderProductFid", o.OrderProductFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
