// Code generated by go-swagger; DO NOT EDIT.

package products

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

// NewPostProductsProductFidAvailabilityCheckParams creates a new PostProductsProductFidAvailabilityCheckParams object
// with the default values initialized.
func NewPostProductsProductFidAvailabilityCheckParams() *PostProductsProductFidAvailabilityCheckParams {
	var ()
	return &PostProductsProductFidAvailabilityCheckParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostProductsProductFidAvailabilityCheckParamsWithTimeout creates a new PostProductsProductFidAvailabilityCheckParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostProductsProductFidAvailabilityCheckParamsWithTimeout(timeout time.Duration) *PostProductsProductFidAvailabilityCheckParams {
	var ()
	return &PostProductsProductFidAvailabilityCheckParams{

		timeout: timeout,
	}
}

// NewPostProductsProductFidAvailabilityCheckParamsWithContext creates a new PostProductsProductFidAvailabilityCheckParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostProductsProductFidAvailabilityCheckParamsWithContext(ctx context.Context) *PostProductsProductFidAvailabilityCheckParams {
	var ()
	return &PostProductsProductFidAvailabilityCheckParams{

		Context: ctx,
	}
}

// NewPostProductsProductFidAvailabilityCheckParamsWithHTTPClient creates a new PostProductsProductFidAvailabilityCheckParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostProductsProductFidAvailabilityCheckParamsWithHTTPClient(client *http.Client) *PostProductsProductFidAvailabilityCheckParams {
	var ()
	return &PostProductsProductFidAvailabilityCheckParams{
		HTTPClient: client,
	}
}

/*PostProductsProductFidAvailabilityCheckParams contains all the parameters to send to the API endpoint
for the post products product fid availability check operation typically these are written to a http.Request
*/
type PostProductsProductFidAvailabilityCheckParams struct {

	/*Payload*/
	Payload *models.AvailabilityCheckPayload
	/*ProductFid*/
	ProductFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post products product fid availability check params
func (o *PostProductsProductFidAvailabilityCheckParams) WithTimeout(timeout time.Duration) *PostProductsProductFidAvailabilityCheckParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post products product fid availability check params
func (o *PostProductsProductFidAvailabilityCheckParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post products product fid availability check params
func (o *PostProductsProductFidAvailabilityCheckParams) WithContext(ctx context.Context) *PostProductsProductFidAvailabilityCheckParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post products product fid availability check params
func (o *PostProductsProductFidAvailabilityCheckParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post products product fid availability check params
func (o *PostProductsProductFidAvailabilityCheckParams) WithHTTPClient(client *http.Client) *PostProductsProductFidAvailabilityCheckParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post products product fid availability check params
func (o *PostProductsProductFidAvailabilityCheckParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPayload adds the payload to the post products product fid availability check params
func (o *PostProductsProductFidAvailabilityCheckParams) WithPayload(payload *models.AvailabilityCheckPayload) *PostProductsProductFidAvailabilityCheckParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the post products product fid availability check params
func (o *PostProductsProductFidAvailabilityCheckParams) SetPayload(payload *models.AvailabilityCheckPayload) {
	o.Payload = payload
}

// WithProductFid adds the productFid to the post products product fid availability check params
func (o *PostProductsProductFidAvailabilityCheckParams) WithProductFid(productFid string) *PostProductsProductFidAvailabilityCheckParams {
	o.SetProductFid(productFid)
	return o
}

// SetProductFid adds the productFid to the post products product fid availability check params
func (o *PostProductsProductFidAvailabilityCheckParams) SetProductFid(productFid string) {
	o.ProductFid = productFid
}

// WriteToRequest writes these params to a swagger request
func (o *PostProductsProductFidAvailabilityCheckParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Payload != nil {
		if err := r.SetBodyParam(o.Payload); err != nil {
			return err
		}
	}

	// path param productFid
	if err := r.SetPathParam("productFid", o.ProductFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
