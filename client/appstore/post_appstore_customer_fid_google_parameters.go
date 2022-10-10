// Code generated by go-swagger; DO NOT EDIT.

package appstore

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

// NewPostAppstoreCustomerFidGoogleParams creates a new PostAppstoreCustomerFidGoogleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAppstoreCustomerFidGoogleParams() *PostAppstoreCustomerFidGoogleParams {
	return &PostAppstoreCustomerFidGoogleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAppstoreCustomerFidGoogleParamsWithTimeout creates a new PostAppstoreCustomerFidGoogleParams object
// with the ability to set a timeout on a request.
func NewPostAppstoreCustomerFidGoogleParamsWithTimeout(timeout time.Duration) *PostAppstoreCustomerFidGoogleParams {
	return &PostAppstoreCustomerFidGoogleParams{
		timeout: timeout,
	}
}

// NewPostAppstoreCustomerFidGoogleParamsWithContext creates a new PostAppstoreCustomerFidGoogleParams object
// with the ability to set a context for a request.
func NewPostAppstoreCustomerFidGoogleParamsWithContext(ctx context.Context) *PostAppstoreCustomerFidGoogleParams {
	return &PostAppstoreCustomerFidGoogleParams{
		Context: ctx,
	}
}

// NewPostAppstoreCustomerFidGoogleParamsWithHTTPClient creates a new PostAppstoreCustomerFidGoogleParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAppstoreCustomerFidGoogleParamsWithHTTPClient(client *http.Client) *PostAppstoreCustomerFidGoogleParams {
	return &PostAppstoreCustomerFidGoogleParams{
		HTTPClient: client,
	}
}

/*
PostAppstoreCustomerFidGoogleParams contains all the parameters to send to the API endpoint

	for the post appstore customer fid google operation.

	Typically these are written to a http.Request.
*/
type PostAppstoreCustomerFidGoogleParams struct {

	/* CustomerFid.

	   Customer FID to use
	*/
	CustomerFid string

	// Payload.
	Payload *models.CreateGoogleNotificationPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post appstore customer fid google params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAppstoreCustomerFidGoogleParams) WithDefaults() *PostAppstoreCustomerFidGoogleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post appstore customer fid google params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAppstoreCustomerFidGoogleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post appstore customer fid google params
func (o *PostAppstoreCustomerFidGoogleParams) WithTimeout(timeout time.Duration) *PostAppstoreCustomerFidGoogleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post appstore customer fid google params
func (o *PostAppstoreCustomerFidGoogleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post appstore customer fid google params
func (o *PostAppstoreCustomerFidGoogleParams) WithContext(ctx context.Context) *PostAppstoreCustomerFidGoogleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post appstore customer fid google params
func (o *PostAppstoreCustomerFidGoogleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post appstore customer fid google params
func (o *PostAppstoreCustomerFidGoogleParams) WithHTTPClient(client *http.Client) *PostAppstoreCustomerFidGoogleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post appstore customer fid google params
func (o *PostAppstoreCustomerFidGoogleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the post appstore customer fid google params
func (o *PostAppstoreCustomerFidGoogleParams) WithCustomerFid(customerFid string) *PostAppstoreCustomerFidGoogleParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the post appstore customer fid google params
func (o *PostAppstoreCustomerFidGoogleParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithPayload adds the payload to the post appstore customer fid google params
func (o *PostAppstoreCustomerFidGoogleParams) WithPayload(payload *models.CreateGoogleNotificationPayload) *PostAppstoreCustomerFidGoogleParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the post appstore customer fid google params
func (o *PostAppstoreCustomerFidGoogleParams) SetPayload(payload *models.CreateGoogleNotificationPayload) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *PostAppstoreCustomerFidGoogleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}
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
