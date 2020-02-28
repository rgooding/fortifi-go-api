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

// NewPostCustomersCustomerFidPaymentsPaymentFidChargebackParams creates a new PostCustomersCustomerFidPaymentsPaymentFidChargebackParams object
// with the default values initialized.
func NewPostCustomersCustomerFidPaymentsPaymentFidChargebackParams() *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams {
	var ()
	return &PostCustomersCustomerFidPaymentsPaymentFidChargebackParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCustomersCustomerFidPaymentsPaymentFidChargebackParamsWithTimeout creates a new PostCustomersCustomerFidPaymentsPaymentFidChargebackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCustomersCustomerFidPaymentsPaymentFidChargebackParamsWithTimeout(timeout time.Duration) *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams {
	var ()
	return &PostCustomersCustomerFidPaymentsPaymentFidChargebackParams{

		timeout: timeout,
	}
}

// NewPostCustomersCustomerFidPaymentsPaymentFidChargebackParamsWithContext creates a new PostCustomersCustomerFidPaymentsPaymentFidChargebackParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCustomersCustomerFidPaymentsPaymentFidChargebackParamsWithContext(ctx context.Context) *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams {
	var ()
	return &PostCustomersCustomerFidPaymentsPaymentFidChargebackParams{

		Context: ctx,
	}
}

// NewPostCustomersCustomerFidPaymentsPaymentFidChargebackParamsWithHTTPClient creates a new PostCustomersCustomerFidPaymentsPaymentFidChargebackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCustomersCustomerFidPaymentsPaymentFidChargebackParamsWithHTTPClient(client *http.Client) *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams {
	var ()
	return &PostCustomersCustomerFidPaymentsPaymentFidChargebackParams{
		HTTPClient: client,
	}
}

/*PostCustomersCustomerFidPaymentsPaymentFidChargebackParams contains all the parameters to send to the API endpoint
for the post customers customer fid payments payment fid chargeback operation typically these are written to a http.Request
*/
type PostCustomersCustomerFidPaymentsPaymentFidChargebackParams struct {

	/*Amount
	  Disputed Amount

	*/
	Amount float32
	/*CaseNumber
	  Case Number

	*/
	CaseNumber string
	/*CustomerFid
	  Customer FID to use

	*/
	CustomerFid string
	/*DateSubmitted
	  Date the chargeback was received (Time in ISO 8601 standard with optional fractions of a second e.g 2015-12-05T13:11:59.888Z)

	*/
	DateSubmitted strfmt.DateTime
	/*Description
	  Notes

	*/
	Description *string
	/*FeeAmount
	  Fee Amount

	*/
	FeeAmount *float32
	/*FeeCurrency
	  3 Character Currency code for the fee

	*/
	FeeCurrency *string
	/*PaymentFid
	  Payment FID to use

	*/
	PaymentFid string
	/*ReasonCode
	  Chargeback Reason Code (see https://github.com/packaged/rwd/blob/master/src/Finance/Chargeback/ChargebackReasonHelper.php)

	*/
	ReasonCode string
	/*Refunded
	  If this payment has been refunded by the alert

	*/
	Refunded bool
	/*Source
	  Source of the alert, or gateway for everything else

	*/
	Source string
	/*State
	  Current State

	*/
	State string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post customers customer fid payments payment fid chargeback params
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) WithTimeout(timeout time.Duration) *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post customers customer fid payments payment fid chargeback params
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post customers customer fid payments payment fid chargeback params
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) WithContext(ctx context.Context) *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post customers customer fid payments payment fid chargeback params
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post customers customer fid payments payment fid chargeback params
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) WithHTTPClient(client *http.Client) *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post customers customer fid payments payment fid chargeback params
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAmount adds the amount to the post customers customer fid payments payment fid chargeback params
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) WithAmount(amount float32) *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams {
	o.SetAmount(amount)
	return o
}

// SetAmount adds the amount to the post customers customer fid payments payment fid chargeback params
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) SetAmount(amount float32) {
	o.Amount = amount
}

// WithCaseNumber adds the caseNumber to the post customers customer fid payments payment fid chargeback params
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) WithCaseNumber(caseNumber string) *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams {
	o.SetCaseNumber(caseNumber)
	return o
}

// SetCaseNumber adds the caseNumber to the post customers customer fid payments payment fid chargeback params
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) SetCaseNumber(caseNumber string) {
	o.CaseNumber = caseNumber
}

// WithCustomerFid adds the customerFid to the post customers customer fid payments payment fid chargeback params
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) WithCustomerFid(customerFid string) *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the post customers customer fid payments payment fid chargeback params
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithDateSubmitted adds the dateSubmitted to the post customers customer fid payments payment fid chargeback params
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) WithDateSubmitted(dateSubmitted strfmt.DateTime) *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams {
	o.SetDateSubmitted(dateSubmitted)
	return o
}

// SetDateSubmitted adds the dateSubmitted to the post customers customer fid payments payment fid chargeback params
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) SetDateSubmitted(dateSubmitted strfmt.DateTime) {
	o.DateSubmitted = dateSubmitted
}

// WithDescription adds the description to the post customers customer fid payments payment fid chargeback params
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) WithDescription(description *string) *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the post customers customer fid payments payment fid chargeback params
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) SetDescription(description *string) {
	o.Description = description
}

// WithFeeAmount adds the feeAmount to the post customers customer fid payments payment fid chargeback params
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) WithFeeAmount(feeAmount *float32) *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams {
	o.SetFeeAmount(feeAmount)
	return o
}

// SetFeeAmount adds the feeAmount to the post customers customer fid payments payment fid chargeback params
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) SetFeeAmount(feeAmount *float32) {
	o.FeeAmount = feeAmount
}

// WithFeeCurrency adds the feeCurrency to the post customers customer fid payments payment fid chargeback params
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) WithFeeCurrency(feeCurrency *string) *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams {
	o.SetFeeCurrency(feeCurrency)
	return o
}

// SetFeeCurrency adds the feeCurrency to the post customers customer fid payments payment fid chargeback params
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) SetFeeCurrency(feeCurrency *string) {
	o.FeeCurrency = feeCurrency
}

// WithPaymentFid adds the paymentFid to the post customers customer fid payments payment fid chargeback params
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) WithPaymentFid(paymentFid string) *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams {
	o.SetPaymentFid(paymentFid)
	return o
}

// SetPaymentFid adds the paymentFid to the post customers customer fid payments payment fid chargeback params
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) SetPaymentFid(paymentFid string) {
	o.PaymentFid = paymentFid
}

// WithReasonCode adds the reasonCode to the post customers customer fid payments payment fid chargeback params
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) WithReasonCode(reasonCode string) *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams {
	o.SetReasonCode(reasonCode)
	return o
}

// SetReasonCode adds the reasonCode to the post customers customer fid payments payment fid chargeback params
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) SetReasonCode(reasonCode string) {
	o.ReasonCode = reasonCode
}

// WithRefunded adds the refunded to the post customers customer fid payments payment fid chargeback params
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) WithRefunded(refunded bool) *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams {
	o.SetRefunded(refunded)
	return o
}

// SetRefunded adds the refunded to the post customers customer fid payments payment fid chargeback params
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) SetRefunded(refunded bool) {
	o.Refunded = refunded
}

// WithSource adds the source to the post customers customer fid payments payment fid chargeback params
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) WithSource(source string) *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams {
	o.SetSource(source)
	return o
}

// SetSource adds the source to the post customers customer fid payments payment fid chargeback params
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) SetSource(source string) {
	o.Source = source
}

// WithState adds the state to the post customers customer fid payments payment fid chargeback params
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) WithState(state string) *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the post customers customer fid payments payment fid chargeback params
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) SetState(state string) {
	o.State = state
}

// WriteToRequest writes these params to a swagger request
func (o *PostCustomersCustomerFidPaymentsPaymentFidChargebackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param amount
	frAmount := o.Amount
	fAmount := swag.FormatFloat32(frAmount)
	if fAmount != "" {
		if err := r.SetFormParam("amount", fAmount); err != nil {
			return err
		}
	}

	// form param caseNumber
	frCaseNumber := o.CaseNumber
	fCaseNumber := frCaseNumber
	if fCaseNumber != "" {
		if err := r.SetFormParam("caseNumber", fCaseNumber); err != nil {
			return err
		}
	}

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}

	// form param dateSubmitted
	frDateSubmitted := o.DateSubmitted
	fDateSubmitted := frDateSubmitted.String()
	if fDateSubmitted != "" {
		if err := r.SetFormParam("dateSubmitted", fDateSubmitted); err != nil {
			return err
		}
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

	if o.FeeAmount != nil {

		// form param feeAmount
		var frFeeAmount float32
		if o.FeeAmount != nil {
			frFeeAmount = *o.FeeAmount
		}
		fFeeAmount := swag.FormatFloat32(frFeeAmount)
		if fFeeAmount != "" {
			if err := r.SetFormParam("feeAmount", fFeeAmount); err != nil {
				return err
			}
		}

	}

	if o.FeeCurrency != nil {

		// form param feeCurrency
		var frFeeCurrency string
		if o.FeeCurrency != nil {
			frFeeCurrency = *o.FeeCurrency
		}
		fFeeCurrency := frFeeCurrency
		if fFeeCurrency != "" {
			if err := r.SetFormParam("feeCurrency", fFeeCurrency); err != nil {
				return err
			}
		}

	}

	// path param paymentFid
	if err := r.SetPathParam("paymentFid", o.PaymentFid); err != nil {
		return err
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

	// form param source
	frSource := o.Source
	fSource := frSource
	if fSource != "" {
		if err := r.SetFormParam("source", fSource); err != nil {
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
