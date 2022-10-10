// Code generated by go-swagger; DO NOT EDIT.

package marketing

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

// NewGetVisitorsVisitorIDParams creates a new GetVisitorsVisitorIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVisitorsVisitorIDParams() *GetVisitorsVisitorIDParams {
	return &GetVisitorsVisitorIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVisitorsVisitorIDParamsWithTimeout creates a new GetVisitorsVisitorIDParams object
// with the ability to set a timeout on a request.
func NewGetVisitorsVisitorIDParamsWithTimeout(timeout time.Duration) *GetVisitorsVisitorIDParams {
	return &GetVisitorsVisitorIDParams{
		timeout: timeout,
	}
}

// NewGetVisitorsVisitorIDParamsWithContext creates a new GetVisitorsVisitorIDParams object
// with the ability to set a context for a request.
func NewGetVisitorsVisitorIDParamsWithContext(ctx context.Context) *GetVisitorsVisitorIDParams {
	return &GetVisitorsVisitorIDParams{
		Context: ctx,
	}
}

// NewGetVisitorsVisitorIDParamsWithHTTPClient creates a new GetVisitorsVisitorIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVisitorsVisitorIDParamsWithHTTPClient(client *http.Client) *GetVisitorsVisitorIDParams {
	return &GetVisitorsVisitorIDParams{
		HTTPClient: client,
	}
}

/*
GetVisitorsVisitorIDParams contains all the parameters to send to the API endpoint

	for the get visitors visitor ID operation.

	Typically these are written to a http.Request.
*/
type GetVisitorsVisitorIDParams struct {

	/* VisitorID.

	     'Visitor ID from the cookie.
	If providing a pre-linked external reference, should be set to 'byref'.
	If no visitor ID is known, client IP should be provided and visitorId should be set to 'unknown''

	*/
	VisitorID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get visitors visitor ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVisitorsVisitorIDParams) WithDefaults() *GetVisitorsVisitorIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get visitors visitor ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVisitorsVisitorIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get visitors visitor ID params
func (o *GetVisitorsVisitorIDParams) WithTimeout(timeout time.Duration) *GetVisitorsVisitorIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get visitors visitor ID params
func (o *GetVisitorsVisitorIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get visitors visitor ID params
func (o *GetVisitorsVisitorIDParams) WithContext(ctx context.Context) *GetVisitorsVisitorIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get visitors visitor ID params
func (o *GetVisitorsVisitorIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get visitors visitor ID params
func (o *GetVisitorsVisitorIDParams) WithHTTPClient(client *http.Client) *GetVisitorsVisitorIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get visitors visitor ID params
func (o *GetVisitorsVisitorIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVisitorID adds the visitorID to the get visitors visitor ID params
func (o *GetVisitorsVisitorIDParams) WithVisitorID(visitorID string) *GetVisitorsVisitorIDParams {
	o.SetVisitorID(visitorID)
	return o
}

// SetVisitorID adds the visitorId to the get visitors visitor ID params
func (o *GetVisitorsVisitorIDParams) SetVisitorID(visitorID string) {
	o.VisitorID = visitorID
}

// WriteToRequest writes these params to a swagger request
func (o *GetVisitorsVisitorIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param visitorId
	if err := r.SetPathParam("visitorId", o.VisitorID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
