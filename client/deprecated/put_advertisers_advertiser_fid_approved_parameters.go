// Code generated by go-swagger; DO NOT EDIT.

package deprecated

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

// NewPutAdvertisersAdvertiserFidApprovedParams creates a new PutAdvertisersAdvertiserFidApprovedParams object
// with the default values initialized.
func NewPutAdvertisersAdvertiserFidApprovedParams() *PutAdvertisersAdvertiserFidApprovedParams {
	var ()
	return &PutAdvertisersAdvertiserFidApprovedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutAdvertisersAdvertiserFidApprovedParamsWithTimeout creates a new PutAdvertisersAdvertiserFidApprovedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutAdvertisersAdvertiserFidApprovedParamsWithTimeout(timeout time.Duration) *PutAdvertisersAdvertiserFidApprovedParams {
	var ()
	return &PutAdvertisersAdvertiserFidApprovedParams{

		timeout: timeout,
	}
}

// NewPutAdvertisersAdvertiserFidApprovedParamsWithContext creates a new PutAdvertisersAdvertiserFidApprovedParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutAdvertisersAdvertiserFidApprovedParamsWithContext(ctx context.Context) *PutAdvertisersAdvertiserFidApprovedParams {
	var ()
	return &PutAdvertisersAdvertiserFidApprovedParams{

		Context: ctx,
	}
}

// NewPutAdvertisersAdvertiserFidApprovedParamsWithHTTPClient creates a new PutAdvertisersAdvertiserFidApprovedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutAdvertisersAdvertiserFidApprovedParamsWithHTTPClient(client *http.Client) *PutAdvertisersAdvertiserFidApprovedParams {
	var ()
	return &PutAdvertisersAdvertiserFidApprovedParams{
		HTTPClient: client,
	}
}

/*PutAdvertisersAdvertiserFidApprovedParams contains all the parameters to send to the API endpoint
for the put advertisers advertiser fid approved operation typically these are written to a http.Request
*/
type PutAdvertisersAdvertiserFidApprovedParams struct {

	/*AdvertiserFid
	  Advertiser FID to use

	*/
	AdvertiserFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put advertisers advertiser fid approved params
func (o *PutAdvertisersAdvertiserFidApprovedParams) WithTimeout(timeout time.Duration) *PutAdvertisersAdvertiserFidApprovedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put advertisers advertiser fid approved params
func (o *PutAdvertisersAdvertiserFidApprovedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put advertisers advertiser fid approved params
func (o *PutAdvertisersAdvertiserFidApprovedParams) WithContext(ctx context.Context) *PutAdvertisersAdvertiserFidApprovedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put advertisers advertiser fid approved params
func (o *PutAdvertisersAdvertiserFidApprovedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put advertisers advertiser fid approved params
func (o *PutAdvertisersAdvertiserFidApprovedParams) WithHTTPClient(client *http.Client) *PutAdvertisersAdvertiserFidApprovedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put advertisers advertiser fid approved params
func (o *PutAdvertisersAdvertiserFidApprovedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdvertiserFid adds the advertiserFid to the put advertisers advertiser fid approved params
func (o *PutAdvertisersAdvertiserFidApprovedParams) WithAdvertiserFid(advertiserFid string) *PutAdvertisersAdvertiserFidApprovedParams {
	o.SetAdvertiserFid(advertiserFid)
	return o
}

// SetAdvertiserFid adds the advertiserFid to the put advertisers advertiser fid approved params
func (o *PutAdvertisersAdvertiserFidApprovedParams) SetAdvertiserFid(advertiserFid string) {
	o.AdvertiserFid = advertiserFid
}

// WriteToRequest writes these params to a swagger request
func (o *PutAdvertisersAdvertiserFidApprovedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param advertiserFid
	if err := r.SetPathParam("advertiserFid", o.AdvertiserFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}