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

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostCustomersCustomerFidTicketsParams creates a new PostCustomersCustomerFidTicketsParams object
// with the default values initialized.
func NewPostCustomersCustomerFidTicketsParams() *PostCustomersCustomerFidTicketsParams {
	var ()
	return &PostCustomersCustomerFidTicketsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCustomersCustomerFidTicketsParamsWithTimeout creates a new PostCustomersCustomerFidTicketsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCustomersCustomerFidTicketsParamsWithTimeout(timeout time.Duration) *PostCustomersCustomerFidTicketsParams {
	var ()
	return &PostCustomersCustomerFidTicketsParams{

		timeout: timeout,
	}
}

// NewPostCustomersCustomerFidTicketsParamsWithContext creates a new PostCustomersCustomerFidTicketsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCustomersCustomerFidTicketsParamsWithContext(ctx context.Context) *PostCustomersCustomerFidTicketsParams {
	var ()
	return &PostCustomersCustomerFidTicketsParams{

		Context: ctx,
	}
}

// NewPostCustomersCustomerFidTicketsParamsWithHTTPClient creates a new PostCustomersCustomerFidTicketsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCustomersCustomerFidTicketsParamsWithHTTPClient(client *http.Client) *PostCustomersCustomerFidTicketsParams {
	var ()
	return &PostCustomersCustomerFidTicketsParams{
		HTTPClient: client,
	}
}

/*PostCustomersCustomerFidTicketsParams contains all the parameters to send to the API endpoint
for the post customers customer fid tickets operation typically these are written to a http.Request
*/
type PostCustomersCustomerFidTicketsParams struct {

	/*Content
	  Content of the support ticket

	*/
	Content string
	/*CustomerFid
	  Customer FID to use

	*/
	CustomerFid string
	/*DepartmentFid
	  Department FID

	*/
	DepartmentFid *string
	/*Recipient
	  Receiver email address e.g. support@yourdomain.com

	*/
	Recipient string
	/*Sender
	  Sender email address e.g. user@customer.com

	*/
	Sender string
	/*Subject
	  Subject of the support ticket

	*/
	Subject string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post customers customer fid tickets params
func (o *PostCustomersCustomerFidTicketsParams) WithTimeout(timeout time.Duration) *PostCustomersCustomerFidTicketsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post customers customer fid tickets params
func (o *PostCustomersCustomerFidTicketsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post customers customer fid tickets params
func (o *PostCustomersCustomerFidTicketsParams) WithContext(ctx context.Context) *PostCustomersCustomerFidTicketsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post customers customer fid tickets params
func (o *PostCustomersCustomerFidTicketsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post customers customer fid tickets params
func (o *PostCustomersCustomerFidTicketsParams) WithHTTPClient(client *http.Client) *PostCustomersCustomerFidTicketsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post customers customer fid tickets params
func (o *PostCustomersCustomerFidTicketsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContent adds the content to the post customers customer fid tickets params
func (o *PostCustomersCustomerFidTicketsParams) WithContent(content string) *PostCustomersCustomerFidTicketsParams {
	o.SetContent(content)
	return o
}

// SetContent adds the content to the post customers customer fid tickets params
func (o *PostCustomersCustomerFidTicketsParams) SetContent(content string) {
	o.Content = content
}

// WithCustomerFid adds the customerFid to the post customers customer fid tickets params
func (o *PostCustomersCustomerFidTicketsParams) WithCustomerFid(customerFid string) *PostCustomersCustomerFidTicketsParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the post customers customer fid tickets params
func (o *PostCustomersCustomerFidTicketsParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithDepartmentFid adds the departmentFid to the post customers customer fid tickets params
func (o *PostCustomersCustomerFidTicketsParams) WithDepartmentFid(departmentFid *string) *PostCustomersCustomerFidTicketsParams {
	o.SetDepartmentFid(departmentFid)
	return o
}

// SetDepartmentFid adds the departmentFid to the post customers customer fid tickets params
func (o *PostCustomersCustomerFidTicketsParams) SetDepartmentFid(departmentFid *string) {
	o.DepartmentFid = departmentFid
}

// WithRecipient adds the recipient to the post customers customer fid tickets params
func (o *PostCustomersCustomerFidTicketsParams) WithRecipient(recipient string) *PostCustomersCustomerFidTicketsParams {
	o.SetRecipient(recipient)
	return o
}

// SetRecipient adds the recipient to the post customers customer fid tickets params
func (o *PostCustomersCustomerFidTicketsParams) SetRecipient(recipient string) {
	o.Recipient = recipient
}

// WithSender adds the sender to the post customers customer fid tickets params
func (o *PostCustomersCustomerFidTicketsParams) WithSender(sender string) *PostCustomersCustomerFidTicketsParams {
	o.SetSender(sender)
	return o
}

// SetSender adds the sender to the post customers customer fid tickets params
func (o *PostCustomersCustomerFidTicketsParams) SetSender(sender string) {
	o.Sender = sender
}

// WithSubject adds the subject to the post customers customer fid tickets params
func (o *PostCustomersCustomerFidTicketsParams) WithSubject(subject string) *PostCustomersCustomerFidTicketsParams {
	o.SetSubject(subject)
	return o
}

// SetSubject adds the subject to the post customers customer fid tickets params
func (o *PostCustomersCustomerFidTicketsParams) SetSubject(subject string) {
	o.Subject = subject
}

// WriteToRequest writes these params to a swagger request
func (o *PostCustomersCustomerFidTicketsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param content
	frContent := o.Content
	fContent := frContent
	if fContent != "" {
		if err := r.SetFormParam("content", fContent); err != nil {
			return err
		}
	}

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}

	if o.DepartmentFid != nil {

		// form param departmentFid
		var frDepartmentFid string
		if o.DepartmentFid != nil {
			frDepartmentFid = *o.DepartmentFid
		}
		fDepartmentFid := frDepartmentFid
		if fDepartmentFid != "" {
			if err := r.SetFormParam("departmentFid", fDepartmentFid); err != nil {
				return err
			}
		}

	}

	// form param recipient
	frRecipient := o.Recipient
	fRecipient := frRecipient
	if fRecipient != "" {
		if err := r.SetFormParam("recipient", fRecipient); err != nil {
			return err
		}
	}

	// form param sender
	frSender := o.Sender
	fSender := frSender
	if fSender != "" {
		if err := r.SetFormParam("sender", fSender); err != nil {
			return err
		}
	}

	// form param subject
	frSubject := o.Subject
	fSubject := frSubject
	if fSubject != "" {
		if err := r.SetFormParam("subject", fSubject); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}