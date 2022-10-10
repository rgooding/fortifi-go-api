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

// NewPutPublishersPublisherFidApprovedParams creates a new PutPublishersPublisherFidApprovedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutPublishersPublisherFidApprovedParams() *PutPublishersPublisherFidApprovedParams {
	return &PutPublishersPublisherFidApprovedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutPublishersPublisherFidApprovedParamsWithTimeout creates a new PutPublishersPublisherFidApprovedParams object
// with the ability to set a timeout on a request.
func NewPutPublishersPublisherFidApprovedParamsWithTimeout(timeout time.Duration) *PutPublishersPublisherFidApprovedParams {
	return &PutPublishersPublisherFidApprovedParams{
		timeout: timeout,
	}
}

// NewPutPublishersPublisherFidApprovedParamsWithContext creates a new PutPublishersPublisherFidApprovedParams object
// with the ability to set a context for a request.
func NewPutPublishersPublisherFidApprovedParamsWithContext(ctx context.Context) *PutPublishersPublisherFidApprovedParams {
	return &PutPublishersPublisherFidApprovedParams{
		Context: ctx,
	}
}

// NewPutPublishersPublisherFidApprovedParamsWithHTTPClient creates a new PutPublishersPublisherFidApprovedParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutPublishersPublisherFidApprovedParamsWithHTTPClient(client *http.Client) *PutPublishersPublisherFidApprovedParams {
	return &PutPublishersPublisherFidApprovedParams{
		HTTPClient: client,
	}
}

/*
PutPublishersPublisherFidApprovedParams contains all the parameters to send to the API endpoint

	for the put publishers publisher fid approved operation.

	Typically these are written to a http.Request.
*/
type PutPublishersPublisherFidApprovedParams struct {

	/* PublisherFid.

	   Publisher FID to use
	*/
	PublisherFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put publishers publisher fid approved params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutPublishersPublisherFidApprovedParams) WithDefaults() *PutPublishersPublisherFidApprovedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put publishers publisher fid approved params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutPublishersPublisherFidApprovedParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put publishers publisher fid approved params
func (o *PutPublishersPublisherFidApprovedParams) WithTimeout(timeout time.Duration) *PutPublishersPublisherFidApprovedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put publishers publisher fid approved params
func (o *PutPublishersPublisherFidApprovedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put publishers publisher fid approved params
func (o *PutPublishersPublisherFidApprovedParams) WithContext(ctx context.Context) *PutPublishersPublisherFidApprovedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put publishers publisher fid approved params
func (o *PutPublishersPublisherFidApprovedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put publishers publisher fid approved params
func (o *PutPublishersPublisherFidApprovedParams) WithHTTPClient(client *http.Client) *PutPublishersPublisherFidApprovedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put publishers publisher fid approved params
func (o *PutPublishersPublisherFidApprovedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPublisherFid adds the publisherFid to the put publishers publisher fid approved params
func (o *PutPublishersPublisherFidApprovedParams) WithPublisherFid(publisherFid string) *PutPublishersPublisherFidApprovedParams {
	o.SetPublisherFid(publisherFid)
	return o
}

// SetPublisherFid adds the publisherFid to the put publishers publisher fid approved params
func (o *PutPublishersPublisherFidApprovedParams) SetPublisherFid(publisherFid string) {
	o.PublisherFid = publisherFid
}

// WriteToRequest writes these params to a swagger request
func (o *PutPublishersPublisherFidApprovedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param publisherFid
	if err := r.SetPathParam("publisherFid", o.PublisherFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
