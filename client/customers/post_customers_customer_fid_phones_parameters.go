// Code generated by go-swagger; DO NOT EDIT.

package customers

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

// NewPostCustomersCustomerFidPhonesParams creates a new PostCustomersCustomerFidPhonesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostCustomersCustomerFidPhonesParams() *PostCustomersCustomerFidPhonesParams {
	return &PostCustomersCustomerFidPhonesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostCustomersCustomerFidPhonesParamsWithTimeout creates a new PostCustomersCustomerFidPhonesParams object
// with the ability to set a timeout on a request.
func NewPostCustomersCustomerFidPhonesParamsWithTimeout(timeout time.Duration) *PostCustomersCustomerFidPhonesParams {
	return &PostCustomersCustomerFidPhonesParams{
		timeout: timeout,
	}
}

// NewPostCustomersCustomerFidPhonesParamsWithContext creates a new PostCustomersCustomerFidPhonesParams object
// with the ability to set a context for a request.
func NewPostCustomersCustomerFidPhonesParamsWithContext(ctx context.Context) *PostCustomersCustomerFidPhonesParams {
	return &PostCustomersCustomerFidPhonesParams{
		Context: ctx,
	}
}

// NewPostCustomersCustomerFidPhonesParamsWithHTTPClient creates a new PostCustomersCustomerFidPhonesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostCustomersCustomerFidPhonesParamsWithHTTPClient(client *http.Client) *PostCustomersCustomerFidPhonesParams {
	return &PostCustomersCustomerFidPhonesParams{
		HTTPClient: client,
	}
}

/* PostCustomersCustomerFidPhonesParams contains all the parameters to send to the API endpoint
   for the post customers customer fid phones operation.

   Typically these are written to a http.Request.
*/
type PostCustomersCustomerFidPhonesParams struct {

	/* CustomerFid.

	   Customer FID to use
	*/
	CustomerFid string

	// DisplayName.
	DisplayName *string

	// PhoneNumber.
	PhoneNumber string

	// SetAsDefault.
	SetAsDefault *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post customers customer fid phones params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostCustomersCustomerFidPhonesParams) WithDefaults() *PostCustomersCustomerFidPhonesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post customers customer fid phones params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostCustomersCustomerFidPhonesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post customers customer fid phones params
func (o *PostCustomersCustomerFidPhonesParams) WithTimeout(timeout time.Duration) *PostCustomersCustomerFidPhonesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post customers customer fid phones params
func (o *PostCustomersCustomerFidPhonesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post customers customer fid phones params
func (o *PostCustomersCustomerFidPhonesParams) WithContext(ctx context.Context) *PostCustomersCustomerFidPhonesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post customers customer fid phones params
func (o *PostCustomersCustomerFidPhonesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post customers customer fid phones params
func (o *PostCustomersCustomerFidPhonesParams) WithHTTPClient(client *http.Client) *PostCustomersCustomerFidPhonesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post customers customer fid phones params
func (o *PostCustomersCustomerFidPhonesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the post customers customer fid phones params
func (o *PostCustomersCustomerFidPhonesParams) WithCustomerFid(customerFid string) *PostCustomersCustomerFidPhonesParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the post customers customer fid phones params
func (o *PostCustomersCustomerFidPhonesParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithDisplayName adds the displayName to the post customers customer fid phones params
func (o *PostCustomersCustomerFidPhonesParams) WithDisplayName(displayName *string) *PostCustomersCustomerFidPhonesParams {
	o.SetDisplayName(displayName)
	return o
}

// SetDisplayName adds the displayName to the post customers customer fid phones params
func (o *PostCustomersCustomerFidPhonesParams) SetDisplayName(displayName *string) {
	o.DisplayName = displayName
}

// WithPhoneNumber adds the phoneNumber to the post customers customer fid phones params
func (o *PostCustomersCustomerFidPhonesParams) WithPhoneNumber(phoneNumber string) *PostCustomersCustomerFidPhonesParams {
	o.SetPhoneNumber(phoneNumber)
	return o
}

// SetPhoneNumber adds the phoneNumber to the post customers customer fid phones params
func (o *PostCustomersCustomerFidPhonesParams) SetPhoneNumber(phoneNumber string) {
	o.PhoneNumber = phoneNumber
}

// WithSetAsDefault adds the setAsDefault to the post customers customer fid phones params
func (o *PostCustomersCustomerFidPhonesParams) WithSetAsDefault(setAsDefault *bool) *PostCustomersCustomerFidPhonesParams {
	o.SetSetAsDefault(setAsDefault)
	return o
}

// SetSetAsDefault adds the setAsDefault to the post customers customer fid phones params
func (o *PostCustomersCustomerFidPhonesParams) SetSetAsDefault(setAsDefault *bool) {
	o.SetAsDefault = setAsDefault
}

// WriteToRequest writes these params to a swagger request
func (o *PostCustomersCustomerFidPhonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}

	if o.DisplayName != nil {

		// form param displayName
		var frDisplayName string
		if o.DisplayName != nil {
			frDisplayName = *o.DisplayName
		}
		fDisplayName := frDisplayName
		if fDisplayName != "" {
			if err := r.SetFormParam("displayName", fDisplayName); err != nil {
				return err
			}
		}
	}

	// form param phoneNumber
	frPhoneNumber := o.PhoneNumber
	fPhoneNumber := frPhoneNumber
	if fPhoneNumber != "" {
		if err := r.SetFormParam("phoneNumber", fPhoneNumber); err != nil {
			return err
		}
	}

	if o.SetAsDefault != nil {

		// form param setAsDefault
		var frSetAsDefault bool
		if o.SetAsDefault != nil {
			frSetAsDefault = *o.SetAsDefault
		}
		fSetAsDefault := swag.FormatBool(frSetAsDefault)
		if fSetAsDefault != "" {
			if err := r.SetFormParam("setAsDefault", fSetAsDefault); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
