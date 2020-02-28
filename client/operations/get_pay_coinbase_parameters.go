// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewGetPayCoinbaseParams creates a new GetPayCoinbaseParams object
// with the default values initialized.
func NewGetPayCoinbaseParams() *GetPayCoinbaseParams {
	var ()
	return &GetPayCoinbaseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPayCoinbaseParamsWithTimeout creates a new GetPayCoinbaseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPayCoinbaseParamsWithTimeout(timeout time.Duration) *GetPayCoinbaseParams {
	var ()
	return &GetPayCoinbaseParams{

		timeout: timeout,
	}
}

// NewGetPayCoinbaseParamsWithContext creates a new GetPayCoinbaseParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPayCoinbaseParamsWithContext(ctx context.Context) *GetPayCoinbaseParams {
	var ()
	return &GetPayCoinbaseParams{

		Context: ctx,
	}
}

// NewGetPayCoinbaseParamsWithHTTPClient creates a new GetPayCoinbaseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPayCoinbaseParamsWithHTTPClient(client *http.Client) *GetPayCoinbaseParams {
	var ()
	return &GetPayCoinbaseParams{
		HTTPClient: client,
	}
}

/*GetPayCoinbaseParams contains all the parameters to send to the API endpoint
for the get pay coinbase operation typically these are written to a http.Request
*/
type GetPayCoinbaseParams struct {

	/*OrderFID
	  FID of the order for which to retrieve a checkout ID

	*/
	OrderFID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get pay coinbase params
func (o *GetPayCoinbaseParams) WithTimeout(timeout time.Duration) *GetPayCoinbaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pay coinbase params
func (o *GetPayCoinbaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pay coinbase params
func (o *GetPayCoinbaseParams) WithContext(ctx context.Context) *GetPayCoinbaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pay coinbase params
func (o *GetPayCoinbaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pay coinbase params
func (o *GetPayCoinbaseParams) WithHTTPClient(client *http.Client) *GetPayCoinbaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pay coinbase params
func (o *GetPayCoinbaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderFID adds the orderFID to the get pay coinbase params
func (o *GetPayCoinbaseParams) WithOrderFID(orderFID *string) *GetPayCoinbaseParams {
	o.SetOrderFID(orderFID)
	return o
}

// SetOrderFID adds the orderFId to the get pay coinbase params
func (o *GetPayCoinbaseParams) SetOrderFID(orderFID *string) {
	o.OrderFID = orderFID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPayCoinbaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.OrderFID != nil {

		// query param orderFID
		var qrOrderFID string
		if o.OrderFID != nil {
			qrOrderFID = *o.OrderFID
		}
		qOrderFID := qrOrderFID
		if qOrderFID != "" {
			if err := r.SetQueryParam("orderFID", qOrderFID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
