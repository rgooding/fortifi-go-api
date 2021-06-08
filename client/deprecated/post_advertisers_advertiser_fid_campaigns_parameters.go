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
	"github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// NewPostAdvertisersAdvertiserFidCampaignsParams creates a new PostAdvertisersAdvertiserFidCampaignsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAdvertisersAdvertiserFidCampaignsParams() *PostAdvertisersAdvertiserFidCampaignsParams {
	return &PostAdvertisersAdvertiserFidCampaignsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAdvertisersAdvertiserFidCampaignsParamsWithTimeout creates a new PostAdvertisersAdvertiserFidCampaignsParams object
// with the ability to set a timeout on a request.
func NewPostAdvertisersAdvertiserFidCampaignsParamsWithTimeout(timeout time.Duration) *PostAdvertisersAdvertiserFidCampaignsParams {
	return &PostAdvertisersAdvertiserFidCampaignsParams{
		timeout: timeout,
	}
}

// NewPostAdvertisersAdvertiserFidCampaignsParamsWithContext creates a new PostAdvertisersAdvertiserFidCampaignsParams object
// with the ability to set a context for a request.
func NewPostAdvertisersAdvertiserFidCampaignsParamsWithContext(ctx context.Context) *PostAdvertisersAdvertiserFidCampaignsParams {
	return &PostAdvertisersAdvertiserFidCampaignsParams{
		Context: ctx,
	}
}

// NewPostAdvertisersAdvertiserFidCampaignsParamsWithHTTPClient creates a new PostAdvertisersAdvertiserFidCampaignsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAdvertisersAdvertiserFidCampaignsParamsWithHTTPClient(client *http.Client) *PostAdvertisersAdvertiserFidCampaignsParams {
	return &PostAdvertisersAdvertiserFidCampaignsParams{
		HTTPClient: client,
	}
}

/* PostAdvertisersAdvertiserFidCampaignsParams contains all the parameters to send to the API endpoint
   for the post advertisers advertiser fid campaigns operation.

   Typically these are written to a http.Request.
*/
type PostAdvertisersAdvertiserFidCampaignsParams struct {

	/* AdvertiserFid.

	   Advertiser FID to use
	*/
	AdvertiserFid string

	// Payload.
	Payload *models.CreateAdvertiserCampaignPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post advertisers advertiser fid campaigns params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAdvertisersAdvertiserFidCampaignsParams) WithDefaults() *PostAdvertisersAdvertiserFidCampaignsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post advertisers advertiser fid campaigns params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAdvertisersAdvertiserFidCampaignsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post advertisers advertiser fid campaigns params
func (o *PostAdvertisersAdvertiserFidCampaignsParams) WithTimeout(timeout time.Duration) *PostAdvertisersAdvertiserFidCampaignsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post advertisers advertiser fid campaigns params
func (o *PostAdvertisersAdvertiserFidCampaignsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post advertisers advertiser fid campaigns params
func (o *PostAdvertisersAdvertiserFidCampaignsParams) WithContext(ctx context.Context) *PostAdvertisersAdvertiserFidCampaignsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post advertisers advertiser fid campaigns params
func (o *PostAdvertisersAdvertiserFidCampaignsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post advertisers advertiser fid campaigns params
func (o *PostAdvertisersAdvertiserFidCampaignsParams) WithHTTPClient(client *http.Client) *PostAdvertisersAdvertiserFidCampaignsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post advertisers advertiser fid campaigns params
func (o *PostAdvertisersAdvertiserFidCampaignsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdvertiserFid adds the advertiserFid to the post advertisers advertiser fid campaigns params
func (o *PostAdvertisersAdvertiserFidCampaignsParams) WithAdvertiserFid(advertiserFid string) *PostAdvertisersAdvertiserFidCampaignsParams {
	o.SetAdvertiserFid(advertiserFid)
	return o
}

// SetAdvertiserFid adds the advertiserFid to the post advertisers advertiser fid campaigns params
func (o *PostAdvertisersAdvertiserFidCampaignsParams) SetAdvertiserFid(advertiserFid string) {
	o.AdvertiserFid = advertiserFid
}

// WithPayload adds the payload to the post advertisers advertiser fid campaigns params
func (o *PostAdvertisersAdvertiserFidCampaignsParams) WithPayload(payload *models.CreateAdvertiserCampaignPayload) *PostAdvertisersAdvertiserFidCampaignsParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the post advertisers advertiser fid campaigns params
func (o *PostAdvertisersAdvertiserFidCampaignsParams) SetPayload(payload *models.CreateAdvertiserCampaignPayload) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *PostAdvertisersAdvertiserFidCampaignsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param advertiserFid
	if err := r.SetPathParam("advertiserFid", o.AdvertiserFid); err != nil {
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
