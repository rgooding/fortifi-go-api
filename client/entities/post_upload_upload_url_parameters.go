// Code generated by go-swagger; DO NOT EDIT.

package entities

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

// NewPostUploadUploadURLParams creates a new PostUploadUploadURLParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostUploadUploadURLParams() *PostUploadUploadURLParams {
	return &PostUploadUploadURLParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostUploadUploadURLParamsWithTimeout creates a new PostUploadUploadURLParams object
// with the ability to set a timeout on a request.
func NewPostUploadUploadURLParamsWithTimeout(timeout time.Duration) *PostUploadUploadURLParams {
	return &PostUploadUploadURLParams{
		timeout: timeout,
	}
}

// NewPostUploadUploadURLParamsWithContext creates a new PostUploadUploadURLParams object
// with the ability to set a context for a request.
func NewPostUploadUploadURLParamsWithContext(ctx context.Context) *PostUploadUploadURLParams {
	return &PostUploadUploadURLParams{
		Context: ctx,
	}
}

// NewPostUploadUploadURLParamsWithHTTPClient creates a new PostUploadUploadURLParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostUploadUploadURLParamsWithHTTPClient(client *http.Client) *PostUploadUploadURLParams {
	return &PostUploadUploadURLParams{
		HTTPClient: client,
	}
}

/* PostUploadUploadURLParams contains all the parameters to send to the API endpoint
   for the post upload upload URL operation.

   Typically these are written to a http.Request.
*/
type PostUploadUploadURLParams struct {

	// Payload.
	Payload *models.RequestUploadURLPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post upload upload URL params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUploadUploadURLParams) WithDefaults() *PostUploadUploadURLParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post upload upload URL params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUploadUploadURLParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post upload upload URL params
func (o *PostUploadUploadURLParams) WithTimeout(timeout time.Duration) *PostUploadUploadURLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post upload upload URL params
func (o *PostUploadUploadURLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post upload upload URL params
func (o *PostUploadUploadURLParams) WithContext(ctx context.Context) *PostUploadUploadURLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post upload upload URL params
func (o *PostUploadUploadURLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post upload upload URL params
func (o *PostUploadUploadURLParams) WithHTTPClient(client *http.Client) *PostUploadUploadURLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post upload upload URL params
func (o *PostUploadUploadURLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPayload adds the payload to the post upload upload URL params
func (o *PostUploadUploadURLParams) WithPayload(payload *models.RequestUploadURLPayload) *PostUploadUploadURLParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the post upload upload URL params
func (o *PostUploadUploadURLParams) SetPayload(payload *models.RequestUploadURLPayload) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *PostUploadUploadURLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
