// Code generated by go-swagger; DO NOT EDIT.

package devices

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

// NewDeleteDeviceHardwareIDParams creates a new DeleteDeviceHardwareIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDeviceHardwareIDParams() *DeleteDeviceHardwareIDParams {
	return &DeleteDeviceHardwareIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDeviceHardwareIDParamsWithTimeout creates a new DeleteDeviceHardwareIDParams object
// with the ability to set a timeout on a request.
func NewDeleteDeviceHardwareIDParamsWithTimeout(timeout time.Duration) *DeleteDeviceHardwareIDParams {
	return &DeleteDeviceHardwareIDParams{
		timeout: timeout,
	}
}

// NewDeleteDeviceHardwareIDParamsWithContext creates a new DeleteDeviceHardwareIDParams object
// with the ability to set a context for a request.
func NewDeleteDeviceHardwareIDParamsWithContext(ctx context.Context) *DeleteDeviceHardwareIDParams {
	return &DeleteDeviceHardwareIDParams{
		Context: ctx,
	}
}

// NewDeleteDeviceHardwareIDParamsWithHTTPClient creates a new DeleteDeviceHardwareIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDeviceHardwareIDParamsWithHTTPClient(client *http.Client) *DeleteDeviceHardwareIDParams {
	return &DeleteDeviceHardwareIDParams{
		HTTPClient: client,
	}
}

/*
DeleteDeviceHardwareIDParams contains all the parameters to send to the API endpoint

	for the delete device hardware ID operation.

	Typically these are written to a http.Request.
*/
type DeleteDeviceHardwareIDParams struct {

	/* EntityFid.

	   Entity FID assigned to the device (CST or CONT)
	*/
	EntityFid *string

	/* HardwareID.

	   Hardware ID
	*/
	HardwareID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete device hardware ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDeviceHardwareIDParams) WithDefaults() *DeleteDeviceHardwareIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete device hardware ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDeviceHardwareIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete device hardware ID params
func (o *DeleteDeviceHardwareIDParams) WithTimeout(timeout time.Duration) *DeleteDeviceHardwareIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete device hardware ID params
func (o *DeleteDeviceHardwareIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete device hardware ID params
func (o *DeleteDeviceHardwareIDParams) WithContext(ctx context.Context) *DeleteDeviceHardwareIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete device hardware ID params
func (o *DeleteDeviceHardwareIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete device hardware ID params
func (o *DeleteDeviceHardwareIDParams) WithHTTPClient(client *http.Client) *DeleteDeviceHardwareIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete device hardware ID params
func (o *DeleteDeviceHardwareIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntityFid adds the entityFid to the delete device hardware ID params
func (o *DeleteDeviceHardwareIDParams) WithEntityFid(entityFid *string) *DeleteDeviceHardwareIDParams {
	o.SetEntityFid(entityFid)
	return o
}

// SetEntityFid adds the entityFid to the delete device hardware ID params
func (o *DeleteDeviceHardwareIDParams) SetEntityFid(entityFid *string) {
	o.EntityFid = entityFid
}

// WithHardwareID adds the hardwareID to the delete device hardware ID params
func (o *DeleteDeviceHardwareIDParams) WithHardwareID(hardwareID string) *DeleteDeviceHardwareIDParams {
	o.SetHardwareID(hardwareID)
	return o
}

// SetHardwareID adds the hardwareId to the delete device hardware ID params
func (o *DeleteDeviceHardwareIDParams) SetHardwareID(hardwareID string) {
	o.HardwareID = hardwareID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDeviceHardwareIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EntityFid != nil {

		// form param entityFid
		var frEntityFid string
		if o.EntityFid != nil {
			frEntityFid = *o.EntityFid
		}
		fEntityFid := frEntityFid
		if fEntityFid != "" {
			if err := r.SetFormParam("entityFid", fEntityFid); err != nil {
				return err
			}
		}
	}

	// path param hardwareId
	if err := r.SetPathParam("hardwareId", o.HardwareID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
