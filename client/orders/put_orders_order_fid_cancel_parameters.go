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

// NewPutOrdersOrderFidCancelParams creates a new PutOrdersOrderFidCancelParams object
// with the default values initialized.
func NewPutOrdersOrderFidCancelParams() *PutOrdersOrderFidCancelParams {
	var ()
	return &PutOrdersOrderFidCancelParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutOrdersOrderFidCancelParamsWithTimeout creates a new PutOrdersOrderFidCancelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutOrdersOrderFidCancelParamsWithTimeout(timeout time.Duration) *PutOrdersOrderFidCancelParams {
	var ()
	return &PutOrdersOrderFidCancelParams{

		timeout: timeout,
	}
}

// NewPutOrdersOrderFidCancelParamsWithContext creates a new PutOrdersOrderFidCancelParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutOrdersOrderFidCancelParamsWithContext(ctx context.Context) *PutOrdersOrderFidCancelParams {
	var ()
	return &PutOrdersOrderFidCancelParams{

		Context: ctx,
	}
}

// NewPutOrdersOrderFidCancelParamsWithHTTPClient creates a new PutOrdersOrderFidCancelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutOrdersOrderFidCancelParamsWithHTTPClient(client *http.Client) *PutOrdersOrderFidCancelParams {
	var ()
	return &PutOrdersOrderFidCancelParams{
		HTTPClient: client,
	}
}

/*PutOrdersOrderFidCancelParams contains all the parameters to send to the API endpoint
for the put orders order fid cancel operation typically these are written to a http.Request
*/
type PutOrdersOrderFidCancelParams struct {

	/*OrderFid*/
	OrderFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put orders order fid cancel params
func (o *PutOrdersOrderFidCancelParams) WithTimeout(timeout time.Duration) *PutOrdersOrderFidCancelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put orders order fid cancel params
func (o *PutOrdersOrderFidCancelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put orders order fid cancel params
func (o *PutOrdersOrderFidCancelParams) WithContext(ctx context.Context) *PutOrdersOrderFidCancelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put orders order fid cancel params
func (o *PutOrdersOrderFidCancelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put orders order fid cancel params
func (o *PutOrdersOrderFidCancelParams) WithHTTPClient(client *http.Client) *PutOrdersOrderFidCancelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put orders order fid cancel params
func (o *PutOrdersOrderFidCancelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderFid adds the orderFid to the put orders order fid cancel params
func (o *PutOrdersOrderFidCancelParams) WithOrderFid(orderFid string) *PutOrdersOrderFidCancelParams {
	o.SetOrderFid(orderFid)
	return o
}

// SetOrderFid adds the orderFid to the put orders order fid cancel params
func (o *PutOrdersOrderFidCancelParams) SetOrderFid(orderFid string) {
	o.OrderFid = orderFid
}

// WriteToRequest writes these params to a swagger request
func (o *PutOrdersOrderFidCancelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param orderFid
	if err := r.SetPathParam("orderFid", o.OrderFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
