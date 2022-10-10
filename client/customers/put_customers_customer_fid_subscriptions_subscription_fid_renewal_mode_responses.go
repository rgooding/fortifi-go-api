// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeReader is a Reader for the PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalMode structure.
type PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeOK creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeOK with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeOK() *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeOK {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeOK{}
}

/*
PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeOK describes a response with status code 200, with default header values.

Renewal mode set
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeOK struct {
}

// IsSuccess returns true when this put customers customer fid subscriptions subscription fid renewal mode o k response has a 2xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put customers customer fid subscriptions subscription fid renewal mode o k response has a 3xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put customers customer fid subscriptions subscription fid renewal mode o k response has a 4xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put customers customer fid subscriptions subscription fid renewal mode o k response has a 5xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put customers customer fid subscriptions subscription fid renewal mode o k response a status code equal to that given
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/renewalMode][%d] putCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeOK ", 200)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeOK) String() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/renewalMode][%d] putCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeOK ", 200)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeDefault creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeDefault with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeDefault(code int) *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeDefault {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeDefault{
		_statusCode: code,
	}
}

/*
PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeDefault describes a response with status code -1, with default header values.

Error
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put customers customer fid subscriptions subscription fid renewal mode default response
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this put customers customer fid subscriptions subscription fid renewal mode default response has a 2xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this put customers customer fid subscriptions subscription fid renewal mode default response has a 3xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this put customers customer fid subscriptions subscription fid renewal mode default response has a 4xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this put customers customer fid subscriptions subscription fid renewal mode default response has a 5xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this put customers customer fid subscriptions subscription fid renewal mode default response a status code equal to that given
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeDefault) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/renewalMode][%d] PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalMode default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeDefault) String() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/renewalMode][%d] PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalMode default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
