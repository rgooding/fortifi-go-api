// Code generated by go-swagger; DO NOT EDIT.

package custom_properties

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

// NewGetEntitiesEntityFidPropertiesValuesPropertyNameParams creates a new GetEntitiesEntityFidPropertiesValuesPropertyNameParams object
// with the default values initialized.
func NewGetEntitiesEntityFidPropertiesValuesPropertyNameParams() *GetEntitiesEntityFidPropertiesValuesPropertyNameParams {
	var ()
	return &GetEntitiesEntityFidPropertiesValuesPropertyNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEntitiesEntityFidPropertiesValuesPropertyNameParamsWithTimeout creates a new GetEntitiesEntityFidPropertiesValuesPropertyNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEntitiesEntityFidPropertiesValuesPropertyNameParamsWithTimeout(timeout time.Duration) *GetEntitiesEntityFidPropertiesValuesPropertyNameParams {
	var ()
	return &GetEntitiesEntityFidPropertiesValuesPropertyNameParams{

		timeout: timeout,
	}
}

// NewGetEntitiesEntityFidPropertiesValuesPropertyNameParamsWithContext creates a new GetEntitiesEntityFidPropertiesValuesPropertyNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEntitiesEntityFidPropertiesValuesPropertyNameParamsWithContext(ctx context.Context) *GetEntitiesEntityFidPropertiesValuesPropertyNameParams {
	var ()
	return &GetEntitiesEntityFidPropertiesValuesPropertyNameParams{

		Context: ctx,
	}
}

// NewGetEntitiesEntityFidPropertiesValuesPropertyNameParamsWithHTTPClient creates a new GetEntitiesEntityFidPropertiesValuesPropertyNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEntitiesEntityFidPropertiesValuesPropertyNameParamsWithHTTPClient(client *http.Client) *GetEntitiesEntityFidPropertiesValuesPropertyNameParams {
	var ()
	return &GetEntitiesEntityFidPropertiesValuesPropertyNameParams{
		HTTPClient: client,
	}
}

/*GetEntitiesEntityFidPropertiesValuesPropertyNameParams contains all the parameters to send to the API endpoint
for the get entities entity fid properties values property name operation typically these are written to a http.Request
*/
type GetEntitiesEntityFidPropertiesValuesPropertyNameParams struct {

	/*EntityFid
	  Entity FID to use

	*/
	EntityFid string
	/*PropertyName
	  Property Name

	*/
	PropertyName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get entities entity fid properties values property name params
func (o *GetEntitiesEntityFidPropertiesValuesPropertyNameParams) WithTimeout(timeout time.Duration) *GetEntitiesEntityFidPropertiesValuesPropertyNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get entities entity fid properties values property name params
func (o *GetEntitiesEntityFidPropertiesValuesPropertyNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get entities entity fid properties values property name params
func (o *GetEntitiesEntityFidPropertiesValuesPropertyNameParams) WithContext(ctx context.Context) *GetEntitiesEntityFidPropertiesValuesPropertyNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get entities entity fid properties values property name params
func (o *GetEntitiesEntityFidPropertiesValuesPropertyNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get entities entity fid properties values property name params
func (o *GetEntitiesEntityFidPropertiesValuesPropertyNameParams) WithHTTPClient(client *http.Client) *GetEntitiesEntityFidPropertiesValuesPropertyNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get entities entity fid properties values property name params
func (o *GetEntitiesEntityFidPropertiesValuesPropertyNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntityFid adds the entityFid to the get entities entity fid properties values property name params
func (o *GetEntitiesEntityFidPropertiesValuesPropertyNameParams) WithEntityFid(entityFid string) *GetEntitiesEntityFidPropertiesValuesPropertyNameParams {
	o.SetEntityFid(entityFid)
	return o
}

// SetEntityFid adds the entityFid to the get entities entity fid properties values property name params
func (o *GetEntitiesEntityFidPropertiesValuesPropertyNameParams) SetEntityFid(entityFid string) {
	o.EntityFid = entityFid
}

// WithPropertyName adds the propertyName to the get entities entity fid properties values property name params
func (o *GetEntitiesEntityFidPropertiesValuesPropertyNameParams) WithPropertyName(propertyName string) *GetEntitiesEntityFidPropertiesValuesPropertyNameParams {
	o.SetPropertyName(propertyName)
	return o
}

// SetPropertyName adds the propertyName to the get entities entity fid properties values property name params
func (o *GetEntitiesEntityFidPropertiesValuesPropertyNameParams) SetPropertyName(propertyName string) {
	o.PropertyName = propertyName
}

// WriteToRequest writes these params to a swagger request
func (o *GetEntitiesEntityFidPropertiesValuesPropertyNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param entityFid
	if err := r.SetPathParam("entityFid", o.EntityFid); err != nil {
		return err
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
