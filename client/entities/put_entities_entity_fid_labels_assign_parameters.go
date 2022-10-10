// Code generated by go-swagger; DO NOT EDIT.

package entities

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

// NewPutEntitiesEntityFidLabelsAssignParams creates a new PutEntitiesEntityFidLabelsAssignParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutEntitiesEntityFidLabelsAssignParams() *PutEntitiesEntityFidLabelsAssignParams {
	return &PutEntitiesEntityFidLabelsAssignParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutEntitiesEntityFidLabelsAssignParamsWithTimeout creates a new PutEntitiesEntityFidLabelsAssignParams object
// with the ability to set a timeout on a request.
func NewPutEntitiesEntityFidLabelsAssignParamsWithTimeout(timeout time.Duration) *PutEntitiesEntityFidLabelsAssignParams {
	return &PutEntitiesEntityFidLabelsAssignParams{
		timeout: timeout,
	}
}

// NewPutEntitiesEntityFidLabelsAssignParamsWithContext creates a new PutEntitiesEntityFidLabelsAssignParams object
// with the ability to set a context for a request.
func NewPutEntitiesEntityFidLabelsAssignParamsWithContext(ctx context.Context) *PutEntitiesEntityFidLabelsAssignParams {
	return &PutEntitiesEntityFidLabelsAssignParams{
		Context: ctx,
	}
}

// NewPutEntitiesEntityFidLabelsAssignParamsWithHTTPClient creates a new PutEntitiesEntityFidLabelsAssignParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutEntitiesEntityFidLabelsAssignParamsWithHTTPClient(client *http.Client) *PutEntitiesEntityFidLabelsAssignParams {
	return &PutEntitiesEntityFidLabelsAssignParams{
		HTTPClient: client,
	}
}

/*
PutEntitiesEntityFidLabelsAssignParams contains all the parameters to send to the API endpoint

	for the put entities entity fid labels assign operation.

	Typically these are written to a http.Request.
*/
type PutEntitiesEntityFidLabelsAssignParams struct {

	/* EntityFid.

	   Entity FID to use
	*/
	EntityFid string

	// Label.
	Label *string

	// Value.
	Value *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put entities entity fid labels assign params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutEntitiesEntityFidLabelsAssignParams) WithDefaults() *PutEntitiesEntityFidLabelsAssignParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put entities entity fid labels assign params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutEntitiesEntityFidLabelsAssignParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put entities entity fid labels assign params
func (o *PutEntitiesEntityFidLabelsAssignParams) WithTimeout(timeout time.Duration) *PutEntitiesEntityFidLabelsAssignParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put entities entity fid labels assign params
func (o *PutEntitiesEntityFidLabelsAssignParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put entities entity fid labels assign params
func (o *PutEntitiesEntityFidLabelsAssignParams) WithContext(ctx context.Context) *PutEntitiesEntityFidLabelsAssignParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put entities entity fid labels assign params
func (o *PutEntitiesEntityFidLabelsAssignParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put entities entity fid labels assign params
func (o *PutEntitiesEntityFidLabelsAssignParams) WithHTTPClient(client *http.Client) *PutEntitiesEntityFidLabelsAssignParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put entities entity fid labels assign params
func (o *PutEntitiesEntityFidLabelsAssignParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntityFid adds the entityFid to the put entities entity fid labels assign params
func (o *PutEntitiesEntityFidLabelsAssignParams) WithEntityFid(entityFid string) *PutEntitiesEntityFidLabelsAssignParams {
	o.SetEntityFid(entityFid)
	return o
}

// SetEntityFid adds the entityFid to the put entities entity fid labels assign params
func (o *PutEntitiesEntityFidLabelsAssignParams) SetEntityFid(entityFid string) {
	o.EntityFid = entityFid
}

// WithLabel adds the label to the put entities entity fid labels assign params
func (o *PutEntitiesEntityFidLabelsAssignParams) WithLabel(label *string) *PutEntitiesEntityFidLabelsAssignParams {
	o.SetLabel(label)
	return o
}

// SetLabel adds the label to the put entities entity fid labels assign params
func (o *PutEntitiesEntityFidLabelsAssignParams) SetLabel(label *string) {
	o.Label = label
}

// WithValue adds the value to the put entities entity fid labels assign params
func (o *PutEntitiesEntityFidLabelsAssignParams) WithValue(value *string) *PutEntitiesEntityFidLabelsAssignParams {
	o.SetValue(value)
	return o
}

// SetValue adds the value to the put entities entity fid labels assign params
func (o *PutEntitiesEntityFidLabelsAssignParams) SetValue(value *string) {
	o.Value = value
}

// WriteToRequest writes these params to a swagger request
func (o *PutEntitiesEntityFidLabelsAssignParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param entityFid
	if err := r.SetPathParam("entityFid", o.EntityFid); err != nil {
		return err
	}

	if o.Label != nil {

		// form param label
		var frLabel string
		if o.Label != nil {
			frLabel = *o.Label
		}
		fLabel := frLabel
		if fLabel != "" {
			if err := r.SetFormParam("label", fLabel); err != nil {
				return err
			}
		}
	}

	if o.Value != nil {

		// form param value
		var frValue string
		if o.Value != nil {
			frValue = *o.Value
		}
		fValue := frValue
		if fValue != "" {
			if err := r.SetFormParam("value", fValue); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
