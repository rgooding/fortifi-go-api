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
	"github.com/go-openapi/strfmt"
)

// NewGetProductsParams creates a new GetProductsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProductsParams() *GetProductsParams {
	return &GetProductsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProductsParamsWithTimeout creates a new GetProductsParams object
// with the ability to set a timeout on a request.
func NewGetProductsParamsWithTimeout(timeout time.Duration) *GetProductsParams {
	return &GetProductsParams{
		timeout: timeout,
	}
}

// NewGetProductsParamsWithContext creates a new GetProductsParams object
// with the ability to set a context for a request.
func NewGetProductsParamsWithContext(ctx context.Context) *GetProductsParams {
	return &GetProductsParams{
		Context: ctx,
	}
}

// NewGetProductsParamsWithHTTPClient creates a new GetProductsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProductsParamsWithHTTPClient(client *http.Client) *GetProductsParams {
	return &GetProductsParams{
		HTTPClient: client,
	}
}

/*
GetProductsParams contains all the parameters to send to the API endpoint

	for the get products operation.

	Typically these are written to a http.Request.
*/
type GetProductsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get products params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProductsParams) WithDefaults() *GetProductsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get products params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProductsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get products params
func (o *GetProductsParams) WithTimeout(timeout time.Duration) *GetProductsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get products params
func (o *GetProductsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get products params
func (o *GetProductsParams) WithContext(ctx context.Context) *GetProductsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get products params
func (o *GetProductsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get products params
func (o *GetProductsParams) WithHTTPClient(client *http.Client) *GetProductsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get products params
func (o *GetProductsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetProductsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
