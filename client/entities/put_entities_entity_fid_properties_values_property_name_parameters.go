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

	"github.com/fortifi/go-api/models"
)

// NewPutEntitiesEntityFidPropertiesValuesPropertyNameParams creates a new PutEntitiesEntityFidPropertiesValuesPropertyNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutEntitiesEntityFidPropertiesValuesPropertyNameParams() *PutEntitiesEntityFidPropertiesValuesPropertyNameParams {
	return &PutEntitiesEntityFidPropertiesValuesPropertyNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutEntitiesEntityFidPropertiesValuesPropertyNameParamsWithTimeout creates a new PutEntitiesEntityFidPropertiesValuesPropertyNameParams object
// with the ability to set a timeout on a request.
func NewPutEntitiesEntityFidPropertiesValuesPropertyNameParamsWithTimeout(timeout time.Duration) *PutEntitiesEntityFidPropertiesValuesPropertyNameParams {
	return &PutEntitiesEntityFidPropertiesValuesPropertyNameParams{
		timeout: timeout,
	}
}

// NewPutEntitiesEntityFidPropertiesValuesPropertyNameParamsWithContext creates a new PutEntitiesEntityFidPropertiesValuesPropertyNameParams object
// with the ability to set a context for a request.
func NewPutEntitiesEntityFidPropertiesValuesPropertyNameParamsWithContext(ctx context.Context) *PutEntitiesEntityFidPropertiesValuesPropertyNameParams {
	return &PutEntitiesEntityFidPropertiesValuesPropertyNameParams{
		Context: ctx,
	}
}

// NewPutEntitiesEntityFidPropertiesValuesPropertyNameParamsWithHTTPClient creates a new PutEntitiesEntityFidPropertiesValuesPropertyNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutEntitiesEntityFidPropertiesValuesPropertyNameParamsWithHTTPClient(client *http.Client) *PutEntitiesEntityFidPropertiesValuesPropertyNameParams {
	return &PutEntitiesEntityFidPropertiesValuesPropertyNameParams{
		HTTPClient: client,
	}
}

/*
PutEntitiesEntityFidPropertiesValuesPropertyNameParams contains all the parameters to send to the API endpoint

	for the put entities entity fid properties values property name operation.

	Typically these are written to a http.Request.
*/
type PutEntitiesEntityFidPropertiesValuesPropertyNameParams struct {

	/* EntityFid.

	   Entity FID to use
	*/
	EntityFid string

	// Payload.
	Payload *models.PropertyValuePayload

	/* PropertyName.

	   Property Name
	*/
	PropertyName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put entities entity fid properties values property name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutEntitiesEntityFidPropertiesValuesPropertyNameParams) WithDefaults() *PutEntitiesEntityFidPropertiesValuesPropertyNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put entities entity fid properties values property name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutEntitiesEntityFidPropertiesValuesPropertyNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put entities entity fid properties values property name params
func (o *PutEntitiesEntityFidPropertiesValuesPropertyNameParams) WithTimeout(timeout time.Duration) *PutEntitiesEntityFidPropertiesValuesPropertyNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put entities entity fid properties values property name params
func (o *PutEntitiesEntityFidPropertiesValuesPropertyNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put entities entity fid properties values property name params
func (o *PutEntitiesEntityFidPropertiesValuesPropertyNameParams) WithContext(ctx context.Context) *PutEntitiesEntityFidPropertiesValuesPropertyNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put entities entity fid properties values property name params
func (o *PutEntitiesEntityFidPropertiesValuesPropertyNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put entities entity fid properties values property name params
func (o *PutEntitiesEntityFidPropertiesValuesPropertyNameParams) WithHTTPClient(client *http.Client) *PutEntitiesEntityFidPropertiesValuesPropertyNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put entities entity fid properties values property name params
func (o *PutEntitiesEntityFidPropertiesValuesPropertyNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntityFid adds the entityFid to the put entities entity fid properties values property name params
func (o *PutEntitiesEntityFidPropertiesValuesPropertyNameParams) WithEntityFid(entityFid string) *PutEntitiesEntityFidPropertiesValuesPropertyNameParams {
	o.SetEntityFid(entityFid)
	return o
}

// SetEntityFid adds the entityFid to the put entities entity fid properties values property name params
func (o *PutEntitiesEntityFidPropertiesValuesPropertyNameParams) SetEntityFid(entityFid string) {
	o.EntityFid = entityFid
}

// WithPayload adds the payload to the put entities entity fid properties values property name params
func (o *PutEntitiesEntityFidPropertiesValuesPropertyNameParams) WithPayload(payload *models.PropertyValuePayload) *PutEntitiesEntityFidPropertiesValuesPropertyNameParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the put entities entity fid properties values property name params
func (o *PutEntitiesEntityFidPropertiesValuesPropertyNameParams) SetPayload(payload *models.PropertyValuePayload) {
	o.Payload = payload
}

// WithPropertyName adds the propertyName to the put entities entity fid properties values property name params
func (o *PutEntitiesEntityFidPropertiesValuesPropertyNameParams) WithPropertyName(propertyName string) *PutEntitiesEntityFidPropertiesValuesPropertyNameParams {
	o.SetPropertyName(propertyName)
	return o
}

// SetPropertyName adds the propertyName to the put entities entity fid properties values property name params
func (o *PutEntitiesEntityFidPropertiesValuesPropertyNameParams) SetPropertyName(propertyName string) {
	o.PropertyName = propertyName
}

// WriteToRequest writes these params to a swagger request
func (o *PutEntitiesEntityFidPropertiesValuesPropertyNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param entityFid
	if err := r.SetPathParam("entityFid", o.EntityFid); err != nil {
		return err
	}
	if o.Payload != nil {
		if err := r.SetBodyParam(o.Payload); err != nil {
			return err
		}
	}

	// path param propertyName
	if err := r.SetPathParam("propertyName", o.PropertyName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
