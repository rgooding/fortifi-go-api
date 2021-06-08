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

// NewPostEntitiesEntityFidConfigSectionNameParams creates a new PostEntitiesEntityFidConfigSectionNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostEntitiesEntityFidConfigSectionNameParams() *PostEntitiesEntityFidConfigSectionNameParams {
	return &PostEntitiesEntityFidConfigSectionNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostEntitiesEntityFidConfigSectionNameParamsWithTimeout creates a new PostEntitiesEntityFidConfigSectionNameParams object
// with the ability to set a timeout on a request.
func NewPostEntitiesEntityFidConfigSectionNameParamsWithTimeout(timeout time.Duration) *PostEntitiesEntityFidConfigSectionNameParams {
	return &PostEntitiesEntityFidConfigSectionNameParams{
		timeout: timeout,
	}
}

// NewPostEntitiesEntityFidConfigSectionNameParamsWithContext creates a new PostEntitiesEntityFidConfigSectionNameParams object
// with the ability to set a context for a request.
func NewPostEntitiesEntityFidConfigSectionNameParamsWithContext(ctx context.Context) *PostEntitiesEntityFidConfigSectionNameParams {
	return &PostEntitiesEntityFidConfigSectionNameParams{
		Context: ctx,
	}
}

// NewPostEntitiesEntityFidConfigSectionNameParamsWithHTTPClient creates a new PostEntitiesEntityFidConfigSectionNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostEntitiesEntityFidConfigSectionNameParamsWithHTTPClient(client *http.Client) *PostEntitiesEntityFidConfigSectionNameParams {
	return &PostEntitiesEntityFidConfigSectionNameParams{
		HTTPClient: client,
	}
}

/* PostEntitiesEntityFidConfigSectionNameParams contains all the parameters to send to the API endpoint
   for the post entities entity fid config section name operation.

   Typically these are written to a http.Request.
*/
type PostEntitiesEntityFidConfigSectionNameParams struct {

	/* EntityFid.

	   Entity FID to use
	*/
	EntityFid string

	// ItemName.
	ItemName string

	/* SectionName.

	   Section Name
	*/
	SectionName string

	// Value.
	Value string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post entities entity fid config section name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostEntitiesEntityFidConfigSectionNameParams) WithDefaults() *PostEntitiesEntityFidConfigSectionNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post entities entity fid config section name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostEntitiesEntityFidConfigSectionNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post entities entity fid config section name params
func (o *PostEntitiesEntityFidConfigSectionNameParams) WithTimeout(timeout time.Duration) *PostEntitiesEntityFidConfigSectionNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post entities entity fid config section name params
func (o *PostEntitiesEntityFidConfigSectionNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post entities entity fid config section name params
func (o *PostEntitiesEntityFidConfigSectionNameParams) WithContext(ctx context.Context) *PostEntitiesEntityFidConfigSectionNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post entities entity fid config section name params
func (o *PostEntitiesEntityFidConfigSectionNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post entities entity fid config section name params
func (o *PostEntitiesEntityFidConfigSectionNameParams) WithHTTPClient(client *http.Client) *PostEntitiesEntityFidConfigSectionNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post entities entity fid config section name params
func (o *PostEntitiesEntityFidConfigSectionNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntityFid adds the entityFid to the post entities entity fid config section name params
func (o *PostEntitiesEntityFidConfigSectionNameParams) WithEntityFid(entityFid string) *PostEntitiesEntityFidConfigSectionNameParams {
	o.SetEntityFid(entityFid)
	return o
}

// SetEntityFid adds the entityFid to the post entities entity fid config section name params
func (o *PostEntitiesEntityFidConfigSectionNameParams) SetEntityFid(entityFid string) {
	o.EntityFid = entityFid
}

// WithItemName adds the itemName to the post entities entity fid config section name params
func (o *PostEntitiesEntityFidConfigSectionNameParams) WithItemName(itemName string) *PostEntitiesEntityFidConfigSectionNameParams {
	o.SetItemName(itemName)
	return o
}

// SetItemName adds the itemName to the post entities entity fid config section name params
func (o *PostEntitiesEntityFidConfigSectionNameParams) SetItemName(itemName string) {
	o.ItemName = itemName
}

// WithSectionName adds the sectionName to the post entities entity fid config section name params
func (o *PostEntitiesEntityFidConfigSectionNameParams) WithSectionName(sectionName string) *PostEntitiesEntityFidConfigSectionNameParams {
	o.SetSectionName(sectionName)
	return o
}

// SetSectionName adds the sectionName to the post entities entity fid config section name params
func (o *PostEntitiesEntityFidConfigSectionNameParams) SetSectionName(sectionName string) {
	o.SectionName = sectionName
}

// WithValue adds the value to the post entities entity fid config section name params
func (o *PostEntitiesEntityFidConfigSectionNameParams) WithValue(value string) *PostEntitiesEntityFidConfigSectionNameParams {
	o.SetValue(value)
	return o
}

// SetValue adds the value to the post entities entity fid config section name params
func (o *PostEntitiesEntityFidConfigSectionNameParams) SetValue(value string) {
	o.Value = value
}

// WriteToRequest writes these params to a swagger request
func (o *PostEntitiesEntityFidConfigSectionNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param entityFid
	if err := r.SetPathParam("entityFid", o.EntityFid); err != nil {
		return err
	}

	// form param itemName
	frItemName := o.ItemName
	fItemName := frItemName
	if fItemName != "" {
		if err := r.SetFormParam("itemName", fItemName); err != nil {
			return err
		}
	}

	// path param sectionName
	if err := r.SetPathParam("sectionName", o.SectionName); err != nil {
		return err
	}

	// form param value
	frValue := o.Value
	fValue := frValue
	if fValue != "" {
		if err := r.SetFormParam("value", fValue); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
