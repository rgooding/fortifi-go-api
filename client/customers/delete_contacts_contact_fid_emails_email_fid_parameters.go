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
)

// NewDeleteContactsContactFidEmailsEmailFidParams creates a new DeleteContactsContactFidEmailsEmailFidParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteContactsContactFidEmailsEmailFidParams() *DeleteContactsContactFidEmailsEmailFidParams {
	return &DeleteContactsContactFidEmailsEmailFidParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteContactsContactFidEmailsEmailFidParamsWithTimeout creates a new DeleteContactsContactFidEmailsEmailFidParams object
// with the ability to set a timeout on a request.
func NewDeleteContactsContactFidEmailsEmailFidParamsWithTimeout(timeout time.Duration) *DeleteContactsContactFidEmailsEmailFidParams {
	return &DeleteContactsContactFidEmailsEmailFidParams{
		timeout: timeout,
	}
}

// NewDeleteContactsContactFidEmailsEmailFidParamsWithContext creates a new DeleteContactsContactFidEmailsEmailFidParams object
// with the ability to set a context for a request.
func NewDeleteContactsContactFidEmailsEmailFidParamsWithContext(ctx context.Context) *DeleteContactsContactFidEmailsEmailFidParams {
	return &DeleteContactsContactFidEmailsEmailFidParams{
		Context: ctx,
	}
}

// NewDeleteContactsContactFidEmailsEmailFidParamsWithHTTPClient creates a new DeleteContactsContactFidEmailsEmailFidParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteContactsContactFidEmailsEmailFidParamsWithHTTPClient(client *http.Client) *DeleteContactsContactFidEmailsEmailFidParams {
	return &DeleteContactsContactFidEmailsEmailFidParams{
		HTTPClient: client,
	}
}

/*
DeleteContactsContactFidEmailsEmailFidParams contains all the parameters to send to the API endpoint

	for the delete contacts contact fid emails email fid operation.

	Typically these are written to a http.Request.
*/
type DeleteContactsContactFidEmailsEmailFidParams struct {

	/* ContactFid.

	   Contact FID to use
	*/
	ContactFid string

	/* EmailFid.

	   Email FID to use
	*/
	EmailFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete contacts contact fid emails email fid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteContactsContactFidEmailsEmailFidParams) WithDefaults() *DeleteContactsContactFidEmailsEmailFidParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete contacts contact fid emails email fid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteContactsContactFidEmailsEmailFidParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete contacts contact fid emails email fid params
func (o *DeleteContactsContactFidEmailsEmailFidParams) WithTimeout(timeout time.Duration) *DeleteContactsContactFidEmailsEmailFidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete contacts contact fid emails email fid params
func (o *DeleteContactsContactFidEmailsEmailFidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete contacts contact fid emails email fid params
func (o *DeleteContactsContactFidEmailsEmailFidParams) WithContext(ctx context.Context) *DeleteContactsContactFidEmailsEmailFidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete contacts contact fid emails email fid params
func (o *DeleteContactsContactFidEmailsEmailFidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete contacts contact fid emails email fid params
func (o *DeleteContactsContactFidEmailsEmailFidParams) WithHTTPClient(client *http.Client) *DeleteContactsContactFidEmailsEmailFidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete contacts contact fid emails email fid params
func (o *DeleteContactsContactFidEmailsEmailFidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactFid adds the contactFid to the delete contacts contact fid emails email fid params
func (o *DeleteContactsContactFidEmailsEmailFidParams) WithContactFid(contactFid string) *DeleteContactsContactFidEmailsEmailFidParams {
	o.SetContactFid(contactFid)
	return o
}

// SetContactFid adds the contactFid to the delete contacts contact fid emails email fid params
func (o *DeleteContactsContactFidEmailsEmailFidParams) SetContactFid(contactFid string) {
	o.ContactFid = contactFid
}

// WithEmailFid adds the emailFid to the delete contacts contact fid emails email fid params
func (o *DeleteContactsContactFidEmailsEmailFidParams) WithEmailFid(emailFid string) *DeleteContactsContactFidEmailsEmailFidParams {
	o.SetEmailFid(emailFid)
	return o
}

// SetEmailFid adds the emailFid to the delete contacts contact fid emails email fid params
func (o *DeleteContactsContactFidEmailsEmailFidParams) SetEmailFid(emailFid string) {
	o.EmailFid = emailFid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteContactsContactFidEmailsEmailFidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contactFid
	if err := r.SetPathParam("contactFid", o.ContactFid); err != nil {
		return err
	}

	// path param emailFid
	if err := r.SetPathParam("emailFid", o.EmailFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
