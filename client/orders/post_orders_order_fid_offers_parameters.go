// Code generated by go-swagger; DO NOT EDIT.

package orders

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
	"github.com/go-openapi/swag"
)

// NewPostOrdersOrderFidOffersParams creates a new PostOrdersOrderFidOffersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostOrdersOrderFidOffersParams() *PostOrdersOrderFidOffersParams {
	return &PostOrdersOrderFidOffersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostOrdersOrderFidOffersParamsWithTimeout creates a new PostOrdersOrderFidOffersParams object
// with the ability to set a timeout on a request.
func NewPostOrdersOrderFidOffersParamsWithTimeout(timeout time.Duration) *PostOrdersOrderFidOffersParams {
	return &PostOrdersOrderFidOffersParams{
		timeout: timeout,
	}
}

// NewPostOrdersOrderFidOffersParamsWithContext creates a new PostOrdersOrderFidOffersParams object
// with the ability to set a context for a request.
func NewPostOrdersOrderFidOffersParamsWithContext(ctx context.Context) *PostOrdersOrderFidOffersParams {
	return &PostOrdersOrderFidOffersParams{
		Context: ctx,
	}
}

// NewPostOrdersOrderFidOffersParamsWithHTTPClient creates a new PostOrdersOrderFidOffersParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostOrdersOrderFidOffersParamsWithHTTPClient(client *http.Client) *PostOrdersOrderFidOffersParams {
	return &PostOrdersOrderFidOffersParams{
		HTTPClient: client,
	}
}

/* PostOrdersOrderFidOffersParams contains all the parameters to send to the API endpoint
   for the post orders order fid offers operation.

   Typically these are written to a http.Request.
*/
type PostOrdersOrderFidOffersParams struct {

	// OfferFid.
	OfferFid string

	// OrderFid.
	OrderFid string

	// OrderItemFid.
	OrderItemFid *string

	// ProductFid.
	ProductFid *string

	// Replace.
	Replace *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post orders order fid offers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostOrdersOrderFidOffersParams) WithDefaults() *PostOrdersOrderFidOffersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post orders order fid offers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostOrdersOrderFidOffersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post orders order fid offers params
func (o *PostOrdersOrderFidOffersParams) WithTimeout(timeout time.Duration) *PostOrdersOrderFidOffersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post orders order fid offers params
func (o *PostOrdersOrderFidOffersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post orders order fid offers params
func (o *PostOrdersOrderFidOffersParams) WithContext(ctx context.Context) *PostOrdersOrderFidOffersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post orders order fid offers params
func (o *PostOrdersOrderFidOffersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post orders order fid offers params
func (o *PostOrdersOrderFidOffersParams) WithHTTPClient(client *http.Client) *PostOrdersOrderFidOffersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post orders order fid offers params
func (o *PostOrdersOrderFidOffersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOfferFid adds the offerFid to the post orders order fid offers params
func (o *PostOrdersOrderFidOffersParams) WithOfferFid(offerFid string) *PostOrdersOrderFidOffersParams {
	o.SetOfferFid(offerFid)
	return o
}

// SetOfferFid adds the offerFid to the post orders order fid offers params
func (o *PostOrdersOrderFidOffersParams) SetOfferFid(offerFid string) {
	o.OfferFid = offerFid
}

// WithOrderFid adds the orderFid to the post orders order fid offers params
func (o *PostOrdersOrderFidOffersParams) WithOrderFid(orderFid string) *PostOrdersOrderFidOffersParams {
	o.SetOrderFid(orderFid)
	return o
}

// SetOrderFid adds the orderFid to the post orders order fid offers params
func (o *PostOrdersOrderFidOffersParams) SetOrderFid(orderFid string) {
	o.OrderFid = orderFid
}

// WithOrderItemFid adds the orderItemFid to the post orders order fid offers params
func (o *PostOrdersOrderFidOffersParams) WithOrderItemFid(orderItemFid *string) *PostOrdersOrderFidOffersParams {
	o.SetOrderItemFid(orderItemFid)
	return o
}

// SetOrderItemFid adds the orderItemFid to the post orders order fid offers params
func (o *PostOrdersOrderFidOffersParams) SetOrderItemFid(orderItemFid *string) {
	o.OrderItemFid = orderItemFid
}

// WithProductFid adds the productFid to the post orders order fid offers params
func (o *PostOrdersOrderFidOffersParams) WithProductFid(productFid *string) *PostOrdersOrderFidOffersParams {
	o.SetProductFid(productFid)
	return o
}

// SetProductFid adds the productFid to the post orders order fid offers params
func (o *PostOrdersOrderFidOffersParams) SetProductFid(productFid *string) {
	o.ProductFid = productFid
}

// WithReplace adds the replace to the post orders order fid offers params
func (o *PostOrdersOrderFidOffersParams) WithReplace(replace *bool) *PostOrdersOrderFidOffersParams {
	o.SetReplace(replace)
	return o
}

// SetReplace adds the replace to the post orders order fid offers params
func (o *PostOrdersOrderFidOffersParams) SetReplace(replace *bool) {
	o.Replace = replace
}

// WriteToRequest writes these params to a swagger request
func (o *PostOrdersOrderFidOffersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param offerFid
	frOfferFid := o.OfferFid
	fOfferFid := frOfferFid
	if fOfferFid != "" {
		if err := r.SetFormParam("offerFid", fOfferFid); err != nil {
			return err
		}
	}

	// path param orderFid
	if err := r.SetPathParam("orderFid", o.OrderFid); err != nil {
		return err
	}

	if o.OrderItemFid != nil {

		// form param orderItemFid
		var frOrderItemFid string
		if o.OrderItemFid != nil {
			frOrderItemFid = *o.OrderItemFid
		}
		fOrderItemFid := frOrderItemFid
		if fOrderItemFid != "" {
			if err := r.SetFormParam("orderItemFid", fOrderItemFid); err != nil {
				return err
			}
		}
	}

	if o.ProductFid != nil {

		// form param productFid
		var frProductFid string
		if o.ProductFid != nil {
			frProductFid = *o.ProductFid
		}
		fProductFid := frProductFid
		if fProductFid != "" {
			if err := r.SetFormParam("productFid", fProductFid); err != nil {
				return err
			}
		}
	}

	if o.Replace != nil {

		// form param replace
		var frReplace bool
		if o.Replace != nil {
			frReplace = *o.Replace
		}
		fReplace := swag.FormatBool(frReplace)
		if fReplace != "" {
			if err := r.SetFormParam("replace", fReplace); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
