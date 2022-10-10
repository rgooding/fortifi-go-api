// Code generated by go-swagger; DO NOT EDIT.

package contacts

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
	"github.com/go-openapi/swag"
)

// NewPutContactsDeviceHardwareIDSubscribeParams creates a new PutContactsDeviceHardwareIDSubscribeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutContactsDeviceHardwareIDSubscribeParams() *PutContactsDeviceHardwareIDSubscribeParams {
	return &PutContactsDeviceHardwareIDSubscribeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutContactsDeviceHardwareIDSubscribeParamsWithTimeout creates a new PutContactsDeviceHardwareIDSubscribeParams object
// with the ability to set a timeout on a request.
func NewPutContactsDeviceHardwareIDSubscribeParamsWithTimeout(timeout time.Duration) *PutContactsDeviceHardwareIDSubscribeParams {
	return &PutContactsDeviceHardwareIDSubscribeParams{
		timeout: timeout,
	}
}

// NewPutContactsDeviceHardwareIDSubscribeParamsWithContext creates a new PutContactsDeviceHardwareIDSubscribeParams object
// with the ability to set a context for a request.
func NewPutContactsDeviceHardwareIDSubscribeParamsWithContext(ctx context.Context) *PutContactsDeviceHardwareIDSubscribeParams {
	return &PutContactsDeviceHardwareIDSubscribeParams{
		Context: ctx,
	}
}

// NewPutContactsDeviceHardwareIDSubscribeParamsWithHTTPClient creates a new PutContactsDeviceHardwareIDSubscribeParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutContactsDeviceHardwareIDSubscribeParamsWithHTTPClient(client *http.Client) *PutContactsDeviceHardwareIDSubscribeParams {
	return &PutContactsDeviceHardwareIDSubscribeParams{
		HTTPClient: client,
	}
}

/*
PutContactsDeviceHardwareIDSubscribeParams contains all the parameters to send to the API endpoint

	for the put contacts device hardware ID subscribe operation.

	Typically these are written to a http.Request.
*/
type PutContactsDeviceHardwareIDSubscribeParams struct {

	/* BrandFid.

	   Brand to subscribe the device to
	*/
	BrandFid *string

	/* ClientIP.

	   IP Address of the visitor
	*/
	ClientIP *string

	/* Encoding.

	   Encoding from the visitors browser 'HTTP_ACCEPT_ENCODING'
	*/
	Encoding *string

	/* EntityFid.

	   Customer or Contact Fid linked to device (Fid)
	*/
	EntityFid *string

	/* GroupFid.

	   Group to subscribe the device to
	*/
	GroupFid *string

	/* HardwareID.

	   Hardware ID
	*/
	HardwareID string

	/* Language.

	   Language from visitors browser 'HTTP_ACCEPT_LANGUAGE'
	*/
	Language *string

	/* OptInData.

	   Additional data to store against opt-in
	*/
	OptInData []string

	/* OptInStatus.

	   Status of customer email opt-in
	*/
	OptInStatus *string

	/* UserAgent.

	   User Agent of the visitors browser 'HTTP_USER_AGENT'
	*/
	UserAgent *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put contacts device hardware ID subscribe params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutContactsDeviceHardwareIDSubscribeParams) WithDefaults() *PutContactsDeviceHardwareIDSubscribeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put contacts device hardware ID subscribe params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutContactsDeviceHardwareIDSubscribeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put contacts device hardware ID subscribe params
func (o *PutContactsDeviceHardwareIDSubscribeParams) WithTimeout(timeout time.Duration) *PutContactsDeviceHardwareIDSubscribeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put contacts device hardware ID subscribe params
func (o *PutContactsDeviceHardwareIDSubscribeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put contacts device hardware ID subscribe params
func (o *PutContactsDeviceHardwareIDSubscribeParams) WithContext(ctx context.Context) *PutContactsDeviceHardwareIDSubscribeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put contacts device hardware ID subscribe params
func (o *PutContactsDeviceHardwareIDSubscribeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put contacts device hardware ID subscribe params
func (o *PutContactsDeviceHardwareIDSubscribeParams) WithHTTPClient(client *http.Client) *PutContactsDeviceHardwareIDSubscribeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put contacts device hardware ID subscribe params
func (o *PutContactsDeviceHardwareIDSubscribeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBrandFid adds the brandFid to the put contacts device hardware ID subscribe params
func (o *PutContactsDeviceHardwareIDSubscribeParams) WithBrandFid(brandFid *string) *PutContactsDeviceHardwareIDSubscribeParams {
	o.SetBrandFid(brandFid)
	return o
}

// SetBrandFid adds the brandFid to the put contacts device hardware ID subscribe params
func (o *PutContactsDeviceHardwareIDSubscribeParams) SetBrandFid(brandFid *string) {
	o.BrandFid = brandFid
}

// WithClientIP adds the clientIP to the put contacts device hardware ID subscribe params
func (o *PutContactsDeviceHardwareIDSubscribeParams) WithClientIP(clientIP *string) *PutContactsDeviceHardwareIDSubscribeParams {
	o.SetClientIP(clientIP)
	return o
}

// SetClientIP adds the clientIp to the put contacts device hardware ID subscribe params
func (o *PutContactsDeviceHardwareIDSubscribeParams) SetClientIP(clientIP *string) {
	o.ClientIP = clientIP
}

// WithEncoding adds the encoding to the put contacts device hardware ID subscribe params
func (o *PutContactsDeviceHardwareIDSubscribeParams) WithEncoding(encoding *string) *PutContactsDeviceHardwareIDSubscribeParams {
	o.SetEncoding(encoding)
	return o
}

// SetEncoding adds the encoding to the put contacts device hardware ID subscribe params
func (o *PutContactsDeviceHardwareIDSubscribeParams) SetEncoding(encoding *string) {
	o.Encoding = encoding
}

// WithEntityFid adds the entityFid to the put contacts device hardware ID subscribe params
func (o *PutContactsDeviceHardwareIDSubscribeParams) WithEntityFid(entityFid *string) *PutContactsDeviceHardwareIDSubscribeParams {
	o.SetEntityFid(entityFid)
	return o
}

// SetEntityFid adds the entityFid to the put contacts device hardware ID subscribe params
func (o *PutContactsDeviceHardwareIDSubscribeParams) SetEntityFid(entityFid *string) {
	o.EntityFid = entityFid
}

// WithGroupFid adds the groupFid to the put contacts device hardware ID subscribe params
func (o *PutContactsDeviceHardwareIDSubscribeParams) WithGroupFid(groupFid *string) *PutContactsDeviceHardwareIDSubscribeParams {
	o.SetGroupFid(groupFid)
	return o
}

// SetGroupFid adds the groupFid to the put contacts device hardware ID subscribe params
func (o *PutContactsDeviceHardwareIDSubscribeParams) SetGroupFid(groupFid *string) {
	o.GroupFid = groupFid
}

// WithHardwareID adds the hardwareID to the put contacts device hardware ID subscribe params
func (o *PutContactsDeviceHardwareIDSubscribeParams) WithHardwareID(hardwareID string) *PutContactsDeviceHardwareIDSubscribeParams {
	o.SetHardwareID(hardwareID)
	return o
}

// SetHardwareID adds the hardwareId to the put contacts device hardware ID subscribe params
func (o *PutContactsDeviceHardwareIDSubscribeParams) SetHardwareID(hardwareID string) {
	o.HardwareID = hardwareID
}

// WithLanguage adds the language to the put contacts device hardware ID subscribe params
func (o *PutContactsDeviceHardwareIDSubscribeParams) WithLanguage(language *string) *PutContactsDeviceHardwareIDSubscribeParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the put contacts device hardware ID subscribe params
func (o *PutContactsDeviceHardwareIDSubscribeParams) SetLanguage(language *string) {
	o.Language = language
}

// WithOptInData adds the optInData to the put contacts device hardware ID subscribe params
func (o *PutContactsDeviceHardwareIDSubscribeParams) WithOptInData(optInData []string) *PutContactsDeviceHardwareIDSubscribeParams {
	o.SetOptInData(optInData)
	return o
}

// SetOptInData adds the optInData to the put contacts device hardware ID subscribe params
func (o *PutContactsDeviceHardwareIDSubscribeParams) SetOptInData(optInData []string) {
	o.OptInData = optInData
}

// WithOptInStatus adds the optInStatus to the put contacts device hardware ID subscribe params
func (o *PutContactsDeviceHardwareIDSubscribeParams) WithOptInStatus(optInStatus *string) *PutContactsDeviceHardwareIDSubscribeParams {
	o.SetOptInStatus(optInStatus)
	return o
}

// SetOptInStatus adds the optInStatus to the put contacts device hardware ID subscribe params
func (o *PutContactsDeviceHardwareIDSubscribeParams) SetOptInStatus(optInStatus *string) {
	o.OptInStatus = optInStatus
}

// WithUserAgent adds the userAgent to the put contacts device hardware ID subscribe params
func (o *PutContactsDeviceHardwareIDSubscribeParams) WithUserAgent(userAgent *string) *PutContactsDeviceHardwareIDSubscribeParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the put contacts device hardware ID subscribe params
func (o *PutContactsDeviceHardwareIDSubscribeParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *PutContactsDeviceHardwareIDSubscribeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BrandFid != nil {

		// form param brandFid
		var frBrandFid string
		if o.BrandFid != nil {
			frBrandFid = *o.BrandFid
		}
		fBrandFid := frBrandFid
		if fBrandFid != "" {
			if err := r.SetFormParam("brandFid", fBrandFid); err != nil {
				return err
			}
		}
	}

	if o.ClientIP != nil {

		// form param clientIp
		var frClientIP string
		if o.ClientIP != nil {
			frClientIP = *o.ClientIP
		}
		fClientIP := frClientIP
		if fClientIP != "" {
			if err := r.SetFormParam("clientIp", fClientIP); err != nil {
				return err
			}
		}
	}

	if o.Encoding != nil {

		// form param encoding
		var frEncoding string
		if o.Encoding != nil {
			frEncoding = *o.Encoding
		}
		fEncoding := frEncoding
		if fEncoding != "" {
			if err := r.SetFormParam("encoding", fEncoding); err != nil {
				return err
			}
		}
	}

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

	if o.GroupFid != nil {

		// form param groupFid
		var frGroupFid string
		if o.GroupFid != nil {
			frGroupFid = *o.GroupFid
		}
		fGroupFid := frGroupFid
		if fGroupFid != "" {
			if err := r.SetFormParam("groupFid", fGroupFid); err != nil {
				return err
			}
		}
	}

	// path param hardwareId
	if err := r.SetPathParam("hardwareId", o.HardwareID); err != nil {
		return err
	}

	if o.Language != nil {

		// form param language
		var frLanguage string
		if o.Language != nil {
			frLanguage = *o.Language
		}
		fLanguage := frLanguage
		if fLanguage != "" {
			if err := r.SetFormParam("language", fLanguage); err != nil {
				return err
			}
		}
	}

	if o.OptInData != nil {

		// binding items for optInData
		joinedOptInData := o.bindParamOptInData(reg)

		// form array param optInData
		if err := r.SetFormParam("optInData", joinedOptInData...); err != nil {
			return err
		}
	}

	if o.OptInStatus != nil {

		// form param optInStatus
		var frOptInStatus string
		if o.OptInStatus != nil {
			frOptInStatus = *o.OptInStatus
		}
		fOptInStatus := frOptInStatus
		if fOptInStatus != "" {
			if err := r.SetFormParam("optInStatus", fOptInStatus); err != nil {
				return err
			}
		}
	}

	if o.UserAgent != nil {

		// form param userAgent
		var frUserAgent string
		if o.UserAgent != nil {
			frUserAgent = *o.UserAgent
		}
		fUserAgent := frUserAgent
		if fUserAgent != "" {
			if err := r.SetFormParam("userAgent", fUserAgent); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamPutContactsDeviceHardwareIDSubscribe binds the parameter optInData
func (o *PutContactsDeviceHardwareIDSubscribeParams) bindParamOptInData(formats strfmt.Registry) []string {
	optInDataIR := o.OptInData

	var optInDataIC []string
	for _, optInDataIIR := range optInDataIR { // explode []string

		optInDataIIV := optInDataIIR // string as string
		optInDataIC = append(optInDataIC, optInDataIIV)
	}

	// items.CollectionFormat: ""
	optInDataIS := swag.JoinByFormat(optInDataIC, "")

	return optInDataIS
}
