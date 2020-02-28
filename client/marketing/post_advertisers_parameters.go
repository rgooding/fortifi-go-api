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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// NewPostAdvertisersParams creates a new PostAdvertisersParams object
// with the default values initialized.
func NewPostAdvertisersParams() *PostAdvertisersParams {
	var ()
	return &PostAdvertisersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAdvertisersParamsWithTimeout creates a new PostAdvertisersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAdvertisersParamsWithTimeout(timeout time.Duration) *PostAdvertisersParams {
	var ()
	return &PostAdvertisersParams{

		timeout: timeout,
	}
}

// NewPostAdvertisersParamsWithContext creates a new PostAdvertisersParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAdvertisersParamsWithContext(ctx context.Context) *PostAdvertisersParams {
	var ()
	return &PostAdvertisersParams{

		Context: ctx,
	}
}

// NewPostAdvertisersParamsWithHTTPClient creates a new PostAdvertisersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAdvertisersParamsWithHTTPClient(client *http.Client) *PostAdvertisersParams {
	var ()
	return &PostAdvertisersParams{
		HTTPClient: client,
	}
}

/*PostAdvertisersParams contains all the parameters to send to the API endpoint
for the post advertisers operation typically these are written to a http.Request
*/
type PostAdvertisersParams struct {

	/*Payload*/
	Payload *models.CreateAdvertiserPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post advertisers params
func (o *PostAdvertisersParams) WithTimeout(timeout time.Duration) *PostAdvertisersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post advertisers params
func (o *PostAdvertisersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post advertisers params
func (o *PostAdvertisersParams) WithContext(ctx context.Context) *PostAdvertisersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post advertisers params
func (o *PostAdvertisersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post advertisers params
func (o *PostAdvertisersParams) WithHTTPClient(client *http.Client) *PostAdvertisersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post advertisers params
func (o *PostAdvertisersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPayload adds the payload to the post advertisers params
func (o *PostAdvertisersParams) WithPayload(payload *models.CreateAdvertiserPayload) *PostAdvertisersParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the post advertisers params
func (o *PostAdvertisersParams) SetPayload(payload *models.CreateAdvertiserPayload) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *PostAdvertisersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
