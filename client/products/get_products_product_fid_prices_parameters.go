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
)

// NewGetProductsProductFidPricesParams creates a new GetProductsProductFidPricesParams object
// with the default values initialized.
func NewGetProductsProductFidPricesParams() *GetProductsProductFidPricesParams {
	var ()
	return &GetProductsProductFidPricesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProductsProductFidPricesParamsWithTimeout creates a new GetProductsProductFidPricesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProductsProductFidPricesParamsWithTimeout(timeout time.Duration) *GetProductsProductFidPricesParams {
	var ()
	return &GetProductsProductFidPricesParams{

		timeout: timeout,
	}
}

// NewGetProductsProductFidPricesParamsWithContext creates a new GetProductsProductFidPricesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProductsProductFidPricesParamsWithContext(ctx context.Context) *GetProductsProductFidPricesParams {
	var ()
	return &GetProductsProductFidPricesParams{

		Context: ctx,
	}
}

// NewGetProductsProductFidPricesParamsWithHTTPClient creates a new GetProductsProductFidPricesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProductsProductFidPricesParamsWithHTTPClient(client *http.Client) *GetProductsProductFidPricesParams {
	var ()
	return &GetProductsProductFidPricesParams{
		HTTPClient: client,
	}
}

/*GetProductsProductFidPricesParams contains all the parameters to send to the API endpoint
for the get products product fid prices operation typically these are written to a http.Request
*/
type GetProductsProductFidPricesParams struct {

	/*Currency*/
	Currency *string
	/*ProductFid*/
	ProductFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get products product fid prices params
func (o *GetProductsProductFidPricesParams) WithTimeout(timeout time.Duration) *GetProductsProductFidPricesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get products product fid prices params
func (o *GetProductsProductFidPricesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get products product fid prices params
func (o *GetProductsProductFidPricesParams) WithContext(ctx context.Context) *GetProductsProductFidPricesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get products product fid prices params
func (o *GetProductsProductFidPricesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get products product fid prices params
func (o *GetProductsProductFidPricesParams) WithHTTPClient(client *http.Client) *GetProductsProductFidPricesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get products product fid prices params
func (o *GetProductsProductFidPricesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCurrency adds the currency to the get products product fid prices params
func (o *GetProductsProductFidPricesParams) WithCurrency(currency *string) *GetProductsProductFidPricesParams {
	o.SetCurrency(currency)
	return o
}

// SetCurrency adds the currency to the get products product fid prices params
func (o *GetProductsProductFidPricesParams) SetCurrency(currency *string) {
	o.Currency = currency
}

// WithProductFid adds the productFid to the get products product fid prices params
func (o *GetProductsProductFidPricesParams) WithProductFid(productFid string) *GetProductsProductFidPricesParams {
	o.SetProductFid(productFid)
	return o
}

// SetProductFid adds the productFid to the get products product fid prices params
func (o *GetProductsProductFidPricesParams) SetProductFid(productFid string) {
	o.ProductFid = productFid
}

// WriteToRequest writes these params to a swagger request
func (o *GetProductsProductFidPricesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Currency != nil {

		// form param currency
		var frCurrency string
		if o.Currency != nil {
			frCurrency = *o.Currency
		}
		fCurrency := frCurrency
		if fCurrency != "" {
			if err := r.SetFormParam("currency", fCurrency); err != nil {
				return err
			}
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
