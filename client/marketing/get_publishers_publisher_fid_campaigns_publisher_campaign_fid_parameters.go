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

// NewGetPublishersPublisherFidCampaignsPublisherCampaignFidParams creates a new GetPublishersPublisherFidCampaignsPublisherCampaignFidParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPublishersPublisherFidCampaignsPublisherCampaignFidParams() *GetPublishersPublisherFidCampaignsPublisherCampaignFidParams {
	return &GetPublishersPublisherFidCampaignsPublisherCampaignFidParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPublishersPublisherFidCampaignsPublisherCampaignFidParamsWithTimeout creates a new GetPublishersPublisherFidCampaignsPublisherCampaignFidParams object
// with the ability to set a timeout on a request.
func NewGetPublishersPublisherFidCampaignsPublisherCampaignFidParamsWithTimeout(timeout time.Duration) *GetPublishersPublisherFidCampaignsPublisherCampaignFidParams {
	return &GetPublishersPublisherFidCampaignsPublisherCampaignFidParams{
		timeout: timeout,
	}
}

// NewGetPublishersPublisherFidCampaignsPublisherCampaignFidParamsWithContext creates a new GetPublishersPublisherFidCampaignsPublisherCampaignFidParams object
// with the ability to set a context for a request.
func NewGetPublishersPublisherFidCampaignsPublisherCampaignFidParamsWithContext(ctx context.Context) *GetPublishersPublisherFidCampaignsPublisherCampaignFidParams {
	return &GetPublishersPublisherFidCampaignsPublisherCampaignFidParams{
		Context: ctx,
	}
}

// NewGetPublishersPublisherFidCampaignsPublisherCampaignFidParamsWithHTTPClient creates a new GetPublishersPublisherFidCampaignsPublisherCampaignFidParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPublishersPublisherFidCampaignsPublisherCampaignFidParamsWithHTTPClient(client *http.Client) *GetPublishersPublisherFidCampaignsPublisherCampaignFidParams {
	return &GetPublishersPublisherFidCampaignsPublisherCampaignFidParams{
		HTTPClient: client,
	}
}

/* GetPublishersPublisherFidCampaignsPublisherCampaignFidParams contains all the parameters to send to the API endpoint
   for the get publishers publisher fid campaigns publisher campaign fid operation.

   Typically these are written to a http.Request.
*/
type GetPublishersPublisherFidCampaignsPublisherCampaignFidParams struct {

	/* PublisherCampaignFid.

	   Publisher Campaign FID to use
	*/
	PublisherCampaignFid string

	/* PublisherFid.

	   Publisher FID to use
	*/
	PublisherFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get publishers publisher fid campaigns publisher campaign fid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPublishersPublisherFidCampaignsPublisherCampaignFidParams) WithDefaults() *GetPublishersPublisherFidCampaignsPublisherCampaignFidParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get publishers publisher fid campaigns publisher campaign fid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPublishersPublisherFidCampaignsPublisherCampaignFidParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get publishers publisher fid campaigns publisher campaign fid params
func (o *GetPublishersPublisherFidCampaignsPublisherCampaignFidParams) WithTimeout(timeout time.Duration) *GetPublishersPublisherFidCampaignsPublisherCampaignFidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get publishers publisher fid campaigns publisher campaign fid params
func (o *GetPublishersPublisherFidCampaignsPublisherCampaignFidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get publishers publisher fid campaigns publisher campaign fid params
func (o *GetPublishersPublisherFidCampaignsPublisherCampaignFidParams) WithContext(ctx context.Context) *GetPublishersPublisherFidCampaignsPublisherCampaignFidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get publishers publisher fid campaigns publisher campaign fid params
func (o *GetPublishersPublisherFidCampaignsPublisherCampaignFidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get publishers publisher fid campaigns publisher campaign fid params
func (o *GetPublishersPublisherFidCampaignsPublisherCampaignFidParams) WithHTTPClient(client *http.Client) *GetPublishersPublisherFidCampaignsPublisherCampaignFidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get publishers publisher fid campaigns publisher campaign fid params
func (o *GetPublishersPublisherFidCampaignsPublisherCampaignFidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPublisherCampaignFid adds the publisherCampaignFid to the get publishers publisher fid campaigns publisher campaign fid params
func (o *GetPublishersPublisherFidCampaignsPublisherCampaignFidParams) WithPublisherCampaignFid(publisherCampaignFid string) *GetPublishersPublisherFidCampaignsPublisherCampaignFidParams {
	o.SetPublisherCampaignFid(publisherCampaignFid)
	return o
}

// SetPublisherCampaignFid adds the publisherCampaignFid to the get publishers publisher fid campaigns publisher campaign fid params
func (o *GetPublishersPublisherFidCampaignsPublisherCampaignFidParams) SetPublisherCampaignFid(publisherCampaignFid string) {
	o.PublisherCampaignFid = publisherCampaignFid
}

// WithPublisherFid adds the publisherFid to the get publishers publisher fid campaigns publisher campaign fid params
func (o *GetPublishersPublisherFidCampaignsPublisherCampaignFidParams) WithPublisherFid(publisherFid string) *GetPublishersPublisherFidCampaignsPublisherCampaignFidParams {
	o.SetPublisherFid(publisherFid)
	return o
}

// SetPublisherFid adds the publisherFid to the get publishers publisher fid campaigns publisher campaign fid params
func (o *GetPublishersPublisherFidCampaignsPublisherCampaignFidParams) SetPublisherFid(publisherFid string) {
	o.PublisherFid = publisherFid
}

// WriteToRequest writes these params to a swagger request
func (o *GetPublishersPublisherFidCampaignsPublisherCampaignFidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param publisherCampaignFid
	if err := r.SetPathParam("publisherCampaignFid", o.PublisherCampaignFid); err != nil {
		return err
	}

	// path param publisherFid
	if err := r.SetPathParam("publisherFid", o.PublisherFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
