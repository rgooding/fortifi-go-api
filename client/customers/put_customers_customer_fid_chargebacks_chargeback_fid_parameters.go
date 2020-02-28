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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPutCustomersCustomerFidChargebacksChargebackFidParams creates a new PutCustomersCustomerFidChargebacksChargebackFidParams object
// with the default values initialized.
func NewPutCustomersCustomerFidChargebacksChargebackFidParams() *PutCustomersCustomerFidChargebacksChargebackFidParams {
	var ()
	return &PutCustomersCustomerFidChargebacksChargebackFidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomersCustomerFidChargebacksChargebackFidParamsWithTimeout creates a new PutCustomersCustomerFidChargebacksChargebackFidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCustomersCustomerFidChargebacksChargebackFidParamsWithTimeout(timeout time.Duration) *PutCustomersCustomerFidChargebacksChargebackFidParams {
	var ()
	return &PutCustomersCustomerFidChargebacksChargebackFidParams{

		timeout: timeout,
	}
}

// NewPutCustomersCustomerFidChargebacksChargebackFidParamsWithContext creates a new PutCustomersCustomerFidChargebacksChargebackFidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCustomersCustomerFidChargebacksChargebackFidParamsWithContext(ctx context.Context) *PutCustomersCustomerFidChargebacksChargebackFidParams {
	var ()
	return &PutCustomersCustomerFidChargebacksChargebackFidParams{

		Context: ctx,
	}
}

// NewPutCustomersCustomerFidChargebacksChargebackFidParamsWithHTTPClient creates a new PutCustomersCustomerFidChargebacksChargebackFidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCustomersCustomerFidChargebacksChargebackFidParamsWithHTTPClient(client *http.Client) *PutCustomersCustomerFidChargebacksChargebackFidParams {
	var ()
	return &PutCustomersCustomerFidChargebacksChargebackFidParams{
		HTTPClient: client,
	}
}

/*PutCustomersCustomerFidChargebacksChargebackFidParams contains all the parameters to send to the API endpoint
for the put customers customer fid chargebacks chargeback fid operation typically these are written to a http.Request
*/
type PutCustomersCustomerFidChargebacksChargebackFidParams struct {

	/*CaseNumber
	  Case Number

	*/
	CaseNumber string
	/*ChargebackFid
	  Chargeback FID to use

	*/
	ChargebackFid string
	/*CustomerFid
	  Customer FID to use

	*/
	CustomerFid string
	/*Description
	  Notes

	*/
	Description *string
	/*ReasonCode
	  Chargeback Reason Code (see https://github.com/packaged/rwd/blob/master/src/Finance/Chargeback/ChargebackReasonHelper.php)

	*/
	ReasonCode string
	/*Refunded
	  If this payment has been refunded by the alert

	*/
	Refunded bool
	/*State
	  Current State

	*/
	State string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put customers customer fid chargebacks chargeback fid params
func (o *PutCustomersCustomerFidChargebacksChargebackFidParams) WithTimeout(timeout time.Duration) *PutCustomersCustomerFidChargebacksChargebackFidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put customers customer fid chargebacks chargeback fid params
func (o *PutCustomersCustomerFidChargebacksChargebackFidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put customers customer fid chargebacks chargeback fid params
func (o *PutCustomersCustomerFidChargebacksChargebackFidParams) WithContext(ctx context.Context) *PutCustomersCustomerFidChargebacksChargebackFidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put customers customer fid chargebacks chargeback fid params
func (o *PutCustomersCustomerFidChargebacksChargebackFidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put customers customer fid chargebacks chargeback fid params
func (o *PutCustomersCustomerFidChargebacksChargebackFidParams) WithHTTPClient(client *http.Client) *PutCustomersCustomerFidChargebacksChargebackFidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put customers customer fid chargebacks chargeback fid params
func (o *PutCustomersCustomerFidChargebacksChargebackFidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCaseNumber adds the caseNumber to the put customers customer fid chargebacks chargeback fid params
func (o *PutCustomersCustomerFidChargebacksChargebackFidParams) WithCaseNumber(caseNumber string) *PutCustomersCustomerFidChargebacksChargebackFidParams {
	o.SetCaseNumber(caseNumber)
	return o
}

// SetCaseNumber adds the caseNumber to the put customers customer fid chargebacks chargeback fid params
func (o *PutCustomersCustomerFidChargebacksChargebackFidParams) SetCaseNumber(caseNumber string) {
	o.CaseNumber = caseNumber
}

// WithChargebackFid adds the chargebackFid to the put customers customer fid chargebacks chargeback fid params
func (o *PutCustomersCustomerFidChargebacksChargebackFidParams) WithChargebackFid(chargebackFid string) *PutCustomersCustomerFidChargebacksChargebackFidParams {
	o.SetChargebackFid(chargebackFid)
	return o
}

// SetChargebackFid adds the chargebackFid to the put customers customer fid chargebacks chargeback fid params
func (o *PutCustomersCustomerFidChargebacksChargebackFidParams) SetChargebackFid(chargebackFid string) {
	o.ChargebackFid = chargebackFid
}

// WithCustomerFid adds the customerFid to the put customers customer fid chargebacks chargeback fid params
func (o *PutCustomersCustomerFidChargebacksChargebackFidParams) WithCustomerFid(customerFid string) *PutCustomersCustomerFidChargebacksChargebackFidParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the put customers customer fid chargebacks chargeback fid params
func (o *PutCustomersCustomerFidChargebacksChargebackFidParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithDescription adds the description to the put customers customer fid chargebacks chargeback fid params
func (o *PutCustomersCustomerFidChargebacksChargebackFidParams) WithDescription(description *string) *PutCustomersCustomerFidChargebacksChargebackFidParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the put customers customer fid chargebacks chargeback fid params
func (o *PutCustomersCustomerFidChargebacksChargebackFidParams) SetDescription(description *string) {
	o.Description = description
}

// WithReasonCode adds the reasonCode to the put customers customer fid chargebacks chargeback fid params
func (o *PutCustomersCustomerFidChargebacksChargebackFidParams) WithReasonCode(reasonCode string) *PutCustomersCustomerFidChargebacksChargebackFidParams {
	o.SetReasonCode(reasonCode)
	return o
}

// SetReasonCode adds the reasonCode to the put customers customer fid chargebacks chargeback fid params
func (o *PutCustomersCustomerFidChargebacksChargebackFidParams) SetReasonCode(reasonCode string) {
	o.ReasonCode = reasonCode
}

// WithRefunded adds the refunded to the put customers customer fid chargebacks chargeback fid params
func (o *PutCustomersCustomerFidChargebacksChargebackFidParams) WithRefunded(refunded bool) *PutCustomersCustomerFidChargebacksChargebackFidParams {
	o.SetRefunded(refunded)
	return o
}

// SetRefunded adds the refunded to the put customers customer fid chargebacks chargeback fid params
func (o *PutCustomersCustomerFidChargebacksChargebackFidParams) SetRefunded(refunded bool) {
	o.Refunded = refunded
}

// WithState adds the state to the put customers customer fid chargebacks chargeback fid params
func (o *PutCustomersCustomerFidChargebacksChargebackFidParams) WithState(state string) *PutCustomersCustomerFidChargebacksChargebackFidParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the put customers customer fid chargebacks chargeback fid params
func (o *PutCustomersCustomerFidChargebacksChargebackFidParams) SetState(state string) {
	o.State = state
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomersCustomerFidChargebacksChargebackFidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param caseNumber
	frCaseNumber := o.CaseNumber
	fCaseNumber := frCaseNumber
	if fCaseNumber != "" {
		if err := r.SetFormParam("caseNumber", fCaseNumber); err != nil {
			return err
		}
	}

	// path param chargebackFid
	if err := r.SetPathParam("chargebackFid", o.ChargebackFid); err != nil {
		return err
	}

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}

	if o.Description != nil {

		// form param description
		var frDescription string
		if o.Description != nil {
			frDescription = *o.Description
		}
		fDescription := frDescription
		if fDescription != "" {
			if err := r.SetFormParam("description", fDescription); err != nil {
				return err
			}
		}

	}

	// form param reasonCode
	frReasonCode := o.ReasonCode
	fReasonCode := frReasonCode
	if fReasonCode != "" {
		if err := r.SetFormParam("reasonCode", fReasonCode); err != nil {
			return err
		}
	}

	// form param refunded
	frRefunded := o.Refunded
	fRefunded := swag.FormatBool(frRefunded)
	if fRefunded != "" {
		if err := r.SetFormParam("refunded", fRefunded); err != nil {
			return err
		}
	}

	// form param state
	frState := o.State
	fState := frState
	if fState != "" {
		if err := r.SetFormParam("state", fState); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
