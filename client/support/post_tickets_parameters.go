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

// NewPostTicketsParams creates a new PostTicketsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostTicketsParams() *PostTicketsParams {
	return &PostTicketsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostTicketsParamsWithTimeout creates a new PostTicketsParams object
// with the ability to set a timeout on a request.
func NewPostTicketsParamsWithTimeout(timeout time.Duration) *PostTicketsParams {
	return &PostTicketsParams{
		timeout: timeout,
	}
}

// NewPostTicketsParamsWithContext creates a new PostTicketsParams object
// with the ability to set a context for a request.
func NewPostTicketsParamsWithContext(ctx context.Context) *PostTicketsParams {
	return &PostTicketsParams{
		Context: ctx,
	}
}

// NewPostTicketsParamsWithHTTPClient creates a new PostTicketsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostTicketsParamsWithHTTPClient(client *http.Client) *PostTicketsParams {
	return &PostTicketsParams{
		HTTPClient: client,
	}
}

/*
PostTicketsParams contains all the parameters to send to the API endpoint

	for the post tickets operation.

	Typically these are written to a http.Request.
*/
type PostTicketsParams struct {

	// Payload.
	Payload *models.CreateTicketPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post tickets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostTicketsParams) WithDefaults() *PostTicketsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post tickets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostTicketsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post tickets params
func (o *PostTicketsParams) WithTimeout(timeout time.Duration) *PostTicketsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post tickets params
func (o *PostTicketsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post tickets params
func (o *PostTicketsParams) WithContext(ctx context.Context) *PostTicketsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post tickets params
func (o *PostTicketsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post tickets params
func (o *PostTicketsParams) WithHTTPClient(client *http.Client) *PostTicketsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post tickets params
func (o *PostTicketsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPayload adds the payload to the post tickets params
func (o *PostTicketsParams) WithPayload(payload *models.CreateTicketPayload) *PostTicketsParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the post tickets params
func (o *PostTicketsParams) SetPayload(payload *models.CreateTicketPayload) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *PostTicketsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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
