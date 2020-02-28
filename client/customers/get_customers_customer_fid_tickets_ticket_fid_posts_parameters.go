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

// NewGetCustomersCustomerFidTicketsTicketFidPostsParams creates a new GetCustomersCustomerFidTicketsTicketFidPostsParams object
// with the default values initialized.
func NewGetCustomersCustomerFidTicketsTicketFidPostsParams() *GetCustomersCustomerFidTicketsTicketFidPostsParams {
	var ()
	return &GetCustomersCustomerFidTicketsTicketFidPostsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomersCustomerFidTicketsTicketFidPostsParamsWithTimeout creates a new GetCustomersCustomerFidTicketsTicketFidPostsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCustomersCustomerFidTicketsTicketFidPostsParamsWithTimeout(timeout time.Duration) *GetCustomersCustomerFidTicketsTicketFidPostsParams {
	var ()
	return &GetCustomersCustomerFidTicketsTicketFidPostsParams{

		timeout: timeout,
	}
}

// NewGetCustomersCustomerFidTicketsTicketFidPostsParamsWithContext creates a new GetCustomersCustomerFidTicketsTicketFidPostsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCustomersCustomerFidTicketsTicketFidPostsParamsWithContext(ctx context.Context) *GetCustomersCustomerFidTicketsTicketFidPostsParams {
	var ()
	return &GetCustomersCustomerFidTicketsTicketFidPostsParams{

		Context: ctx,
	}
}

// NewGetCustomersCustomerFidTicketsTicketFidPostsParamsWithHTTPClient creates a new GetCustomersCustomerFidTicketsTicketFidPostsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCustomersCustomerFidTicketsTicketFidPostsParamsWithHTTPClient(client *http.Client) *GetCustomersCustomerFidTicketsTicketFidPostsParams {
	var ()
	return &GetCustomersCustomerFidTicketsTicketFidPostsParams{
		HTTPClient: client,
	}
}

/*GetCustomersCustomerFidTicketsTicketFidPostsParams contains all the parameters to send to the API endpoint
for the get customers customer fid tickets ticket fid posts operation typically these are written to a http.Request
*/
type GetCustomersCustomerFidTicketsTicketFidPostsParams struct {

	/*CustomerFid
	  Customer FID to use

	*/
	CustomerFid string
	/*TicketFid
	  Ticket FID to use

	*/
	TicketFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get customers customer fid tickets ticket fid posts params
func (o *GetCustomersCustomerFidTicketsTicketFidPostsParams) WithTimeout(timeout time.Duration) *GetCustomersCustomerFidTicketsTicketFidPostsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get customers customer fid tickets ticket fid posts params
func (o *GetCustomersCustomerFidTicketsTicketFidPostsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get customers customer fid tickets ticket fid posts params
func (o *GetCustomersCustomerFidTicketsTicketFidPostsParams) WithContext(ctx context.Context) *GetCustomersCustomerFidTicketsTicketFidPostsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get customers customer fid tickets ticket fid posts params
func (o *GetCustomersCustomerFidTicketsTicketFidPostsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get customers customer fid tickets ticket fid posts params
func (o *GetCustomersCustomerFidTicketsTicketFidPostsParams) WithHTTPClient(client *http.Client) *GetCustomersCustomerFidTicketsTicketFidPostsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get customers customer fid tickets ticket fid posts params
func (o *GetCustomersCustomerFidTicketsTicketFidPostsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the get customers customer fid tickets ticket fid posts params
func (o *GetCustomersCustomerFidTicketsTicketFidPostsParams) WithCustomerFid(customerFid string) *GetCustomersCustomerFidTicketsTicketFidPostsParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the get customers customer fid tickets ticket fid posts params
func (o *GetCustomersCustomerFidTicketsTicketFidPostsParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithTicketFid adds the ticketFid to the get customers customer fid tickets ticket fid posts params
func (o *GetCustomersCustomerFidTicketsTicketFidPostsParams) WithTicketFid(ticketFid string) *GetCustomersCustomerFidTicketsTicketFidPostsParams {
	o.SetTicketFid(ticketFid)
	return o
}

// SetTicketFid adds the ticketFid to the get customers customer fid tickets ticket fid posts params
func (o *GetCustomersCustomerFidTicketsTicketFidPostsParams) SetTicketFid(ticketFid string) {
	o.TicketFid = ticketFid
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomersCustomerFidTicketsTicketFidPostsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}

	// path param ticketFid
	if err := r.SetPathParam("ticketFid", o.TicketFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}