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
)

// NewPutOrdersOrderFidSetChargeIDParams creates a new PutOrdersOrderFidSetChargeIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutOrdersOrderFidSetChargeIDParams() *PutOrdersOrderFidSetChargeIDParams {
	return &PutOrdersOrderFidSetChargeIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutOrdersOrderFidSetChargeIDParamsWithTimeout creates a new PutOrdersOrderFidSetChargeIDParams object
// with the ability to set a timeout on a request.
func NewPutOrdersOrderFidSetChargeIDParamsWithTimeout(timeout time.Duration) *PutOrdersOrderFidSetChargeIDParams {
	return &PutOrdersOrderFidSetChargeIDParams{
		timeout: timeout,
	}
}

// NewPutOrdersOrderFidSetChargeIDParamsWithContext creates a new PutOrdersOrderFidSetChargeIDParams object
// with the ability to set a context for a request.
func NewPutOrdersOrderFidSetChargeIDParamsWithContext(ctx context.Context) *PutOrdersOrderFidSetChargeIDParams {
	return &PutOrdersOrderFidSetChargeIDParams{
		Context: ctx,
	}
}

// NewPutOrdersOrderFidSetChargeIDParamsWithHTTPClient creates a new PutOrdersOrderFidSetChargeIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutOrdersOrderFidSetChargeIDParamsWithHTTPClient(client *http.Client) *PutOrdersOrderFidSetChargeIDParams {
	return &PutOrdersOrderFidSetChargeIDParams{
		HTTPClient: client,
	}
}

/* PutOrdersOrderFidSetChargeIDParams contains all the parameters to send to the API endpoint
   for the put orders order fid set charge ID operation.

   Typically these are written to a http.Request.
*/
type PutOrdersOrderFidSetChargeIDParams struct {

	/* ChargeID.

	   Charge ID provided by ChargeHive.com
	*/
	ChargeID string

	// OrderFid.
	OrderFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put orders order fid set charge ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutOrdersOrderFidSetChargeIDParams) WithDefaults() *PutOrdersOrderFidSetChargeIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put orders order fid set charge ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutOrdersOrderFidSetChargeIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put orders order fid set charge ID params
func (o *PutOrdersOrderFidSetChargeIDParams) WithTimeout(timeout time.Duration) *PutOrdersOrderFidSetChargeIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put orders order fid set charge ID params
func (o *PutOrdersOrderFidSetChargeIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put orders order fid set charge ID params
func (o *PutOrdersOrderFidSetChargeIDParams) WithContext(ctx context.Context) *PutOrdersOrderFidSetChargeIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put orders order fid set charge ID params
func (o *PutOrdersOrderFidSetChargeIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put orders order fid set charge ID params
func (o *PutOrdersOrderFidSetChargeIDParams) WithHTTPClient(client *http.Client) *PutOrdersOrderFidSetChargeIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put orders order fid set charge ID params
func (o *PutOrdersOrderFidSetChargeIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChargeID adds the chargeID to the put orders order fid set charge ID params
func (o *PutOrdersOrderFidSetChargeIDParams) WithChargeID(chargeID string) *PutOrdersOrderFidSetChargeIDParams {
	o.SetChargeID(chargeID)
	return o
}

// SetChargeID adds the chargeId to the put orders order fid set charge ID params
func (o *PutOrdersOrderFidSetChargeIDParams) SetChargeID(chargeID string) {
	o.ChargeID = chargeID
}

// WithOrderFid adds the orderFid to the put orders order fid set charge ID params
func (o *PutOrdersOrderFidSetChargeIDParams) WithOrderFid(orderFid string) *PutOrdersOrderFidSetChargeIDParams {
	o.SetOrderFid(orderFid)
	return o
}

// SetOrderFid adds the orderFid to the put orders order fid set charge ID params
func (o *PutOrdersOrderFidSetChargeIDParams) SetOrderFid(orderFid string) {
	o.OrderFid = orderFid
}

// WriteToRequest writes these params to a swagger request
func (o *PutOrdersOrderFidSetChargeIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param chargeId
	frChargeID := o.ChargeID
	fChargeID := frChargeID
	if fChargeID != "" {
		if err := r.SetFormParam("chargeId", fChargeID); err != nil {
			return err
		}
	}

	// path param orderFid
	if err := r.SetPathParam("orderFid", o.OrderFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
