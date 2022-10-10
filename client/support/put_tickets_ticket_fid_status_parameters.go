// Code generated by go-swagger; DO NOT EDIT.

package support

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

// NewPutTicketsTicketFidStatusParams creates a new PutTicketsTicketFidStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutTicketsTicketFidStatusParams() *PutTicketsTicketFidStatusParams {
	return &PutTicketsTicketFidStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutTicketsTicketFidStatusParamsWithTimeout creates a new PutTicketsTicketFidStatusParams object
// with the ability to set a timeout on a request.
func NewPutTicketsTicketFidStatusParamsWithTimeout(timeout time.Duration) *PutTicketsTicketFidStatusParams {
	return &PutTicketsTicketFidStatusParams{
		timeout: timeout,
	}
}

// NewPutTicketsTicketFidStatusParamsWithContext creates a new PutTicketsTicketFidStatusParams object
// with the ability to set a context for a request.
func NewPutTicketsTicketFidStatusParamsWithContext(ctx context.Context) *PutTicketsTicketFidStatusParams {
	return &PutTicketsTicketFidStatusParams{
		Context: ctx,
	}
}

// NewPutTicketsTicketFidStatusParamsWithHTTPClient creates a new PutTicketsTicketFidStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutTicketsTicketFidStatusParamsWithHTTPClient(client *http.Client) *PutTicketsTicketFidStatusParams {
	return &PutTicketsTicketFidStatusParams{
		HTTPClient: client,
	}
}

/*
PutTicketsTicketFidStatusParams contains all the parameters to send to the API endpoint

	for the put tickets ticket fid status operation.

	Typically these are written to a http.Request.
*/
type PutTicketsTicketFidStatusParams struct {

	// Payload.
	Payload *models.TicketStatusPayload

	/* TicketFid.

	   Ticket FID to use
	*/
	TicketFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put tickets ticket fid status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutTicketsTicketFidStatusParams) WithDefaults() *PutTicketsTicketFidStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put tickets ticket fid status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutTicketsTicketFidStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put tickets ticket fid status params
func (o *PutTicketsTicketFidStatusParams) WithTimeout(timeout time.Duration) *PutTicketsTicketFidStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put tickets ticket fid status params
func (o *PutTicketsTicketFidStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put tickets ticket fid status params
func (o *PutTicketsTicketFidStatusParams) WithContext(ctx context.Context) *PutTicketsTicketFidStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put tickets ticket fid status params
func (o *PutTicketsTicketFidStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put tickets ticket fid status params
func (o *PutTicketsTicketFidStatusParams) WithHTTPClient(client *http.Client) *PutTicketsTicketFidStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put tickets ticket fid status params
func (o *PutTicketsTicketFidStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPayload adds the payload to the put tickets ticket fid status params
func (o *PutTicketsTicketFidStatusParams) WithPayload(payload *models.TicketStatusPayload) *PutTicketsTicketFidStatusParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the put tickets ticket fid status params
func (o *PutTicketsTicketFidStatusParams) SetPayload(payload *models.TicketStatusPayload) {
	o.Payload = payload
}

// WithTicketFid adds the ticketFid to the put tickets ticket fid status params
func (o *PutTicketsTicketFidStatusParams) WithTicketFid(ticketFid string) *PutTicketsTicketFidStatusParams {
	o.SetTicketFid(ticketFid)
	return o
}

// SetTicketFid adds the ticketFid to the put tickets ticket fid status params
func (o *PutTicketsTicketFidStatusParams) SetTicketFid(ticketFid string) {
	o.TicketFid = ticketFid
}

// WriteToRequest writes these params to a swagger request
func (o *PutTicketsTicketFidStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Payload != nil {
		if err := r.SetBodyParam(o.Payload); err != nil {
			return err
		}
	}

	// path param ticketFid
	if err := r.SetPathParam("ticketFid", o.TicketFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
