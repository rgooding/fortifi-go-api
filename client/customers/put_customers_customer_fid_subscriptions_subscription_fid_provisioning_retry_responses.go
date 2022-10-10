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

// PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryReader is a Reader for the PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetry structure.
type PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryOK creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryOK with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryOK() *PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryOK {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryOK{}
}

/*
PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryOK describes a response with status code 200, with default header values.

Provisioning Queued
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryOK struct {
}

// IsSuccess returns true when this put customers customer fid subscriptions subscription fid provisioning retry o k response has a 2xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put customers customer fid subscriptions subscription fid provisioning retry o k response has a 3xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put customers customer fid subscriptions subscription fid provisioning retry o k response has a 4xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put customers customer fid subscriptions subscription fid provisioning retry o k response has a 5xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put customers customer fid subscriptions subscription fid provisioning retry o k response a status code equal to that given
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/provisioning/retry][%d] putCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryOK ", 200)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryOK) String() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/provisioning/retry][%d] putCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryOK ", 200)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryDefault creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryDefault with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryDefault(code int) *PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryDefault {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryDefault{
		_statusCode: code,
	}
}

/*
PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryDefault describes a response with status code -1, with default header values.

Error
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put customers customer fid subscriptions subscription fid provisioning retry default response
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this put customers customer fid subscriptions subscription fid provisioning retry default response has a 2xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this put customers customer fid subscriptions subscription fid provisioning retry default response has a 3xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this put customers customer fid subscriptions subscription fid provisioning retry default response has a 4xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this put customers customer fid subscriptions subscription fid provisioning retry default response has a 5xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this put customers customer fid subscriptions subscription fid provisioning retry default response a status code equal to that given
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryDefault) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/provisioning/retry][%d] PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetry default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryDefault) String() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/provisioning/retry][%d] PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetry default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidProvisioningRetryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}