// Code generated by go-swagger; DO NOT EDIT.

package interactions

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

// NewGetInteractionsInteractionFidParams creates a new GetInteractionsInteractionFidParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInteractionsInteractionFidParams() *GetInteractionsInteractionFidParams {
	return &GetInteractionsInteractionFidParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInteractionsInteractionFidParamsWithTimeout creates a new GetInteractionsInteractionFidParams object
// with the ability to set a timeout on a request.
func NewGetInteractionsInteractionFidParamsWithTimeout(timeout time.Duration) *GetInteractionsInteractionFidParams {
	return &GetInteractionsInteractionFidParams{
		timeout: timeout,
	}
}

// NewGetInteractionsInteractionFidParamsWithContext creates a new GetInteractionsInteractionFidParams object
// with the ability to set a context for a request.
func NewGetInteractionsInteractionFidParamsWithContext(ctx context.Context) *GetInteractionsInteractionFidParams {
	return &GetInteractionsInteractionFidParams{
		Context: ctx,
	}
}

// NewGetInteractionsInteractionFidParamsWithHTTPClient creates a new GetInteractionsInteractionFidParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInteractionsInteractionFidParamsWithHTTPClient(client *http.Client) *GetInteractionsInteractionFidParams {
	return &GetInteractionsInteractionFidParams{
		HTTPClient: client,
	}
}

/*
GetInteractionsInteractionFidParams contains all the parameters to send to the API endpoint

	for the get interactions interaction fid operation.

	Typically these are written to a http.Request.
*/
type GetInteractionsInteractionFidParams struct {

	// InteractionFid.
	InteractionFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get interactions interaction fid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInteractionsInteractionFidParams) WithDefaults() *GetInteractionsInteractionFidParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get interactions interaction fid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInteractionsInteractionFidParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get interactions interaction fid params
func (o *GetInteractionsInteractionFidParams) WithTimeout(timeout time.Duration) *GetInteractionsInteractionFidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get interactions interaction fid params
func (o *GetInteractionsInteractionFidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get interactions interaction fid params
func (o *GetInteractionsInteractionFidParams) WithContext(ctx context.Context) *GetInteractionsInteractionFidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get interactions interaction fid params
func (o *GetInteractionsInteractionFidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get interactions interaction fid params
func (o *GetInteractionsInteractionFidParams) WithHTTPClient(client *http.Client) *GetInteractionsInteractionFidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get interactions interaction fid params
func (o *GetInteractionsInteractionFidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInteractionFid adds the interactionFid to the get interactions interaction fid params
func (o *GetInteractionsInteractionFidParams) WithInteractionFid(interactionFid string) *GetInteractionsInteractionFidParams {
	o.SetInteractionFid(interactionFid)
	return o
}

// SetInteractionFid adds the interactionFid to the get interactions interaction fid params
func (o *GetInteractionsInteractionFidParams) SetInteractionFid(interactionFid string) {
	o.InteractionFid = interactionFid
}

// WriteToRequest writes these params to a swagger request
func (o *GetInteractionsInteractionFidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param interactionFid
	if err := r.SetPathParam("interactionFid", o.InteractionFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}