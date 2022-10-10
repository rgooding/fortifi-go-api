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

// PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundReader is a Reader for the PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefund structure.
type PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundOK creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundOK with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundOK() *PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundOK {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundOK{}
}

/*
PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundOK describes a response with status code 200, with default header values.

Refund Actioned
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundOK struct {
}

// IsSuccess returns true when this put customers customer fid subscriptions subscription fid periods period fid refund o k response has a 2xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put customers customer fid subscriptions subscription fid periods period fid refund o k response has a 3xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put customers customer fid subscriptions subscription fid periods period fid refund o k response has a 4xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put customers customer fid subscriptions subscription fid periods period fid refund o k response has a 5xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put customers customer fid subscriptions subscription fid periods period fid refund o k response a status code equal to that given
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/periods/{periodFid}/refund][%d] putCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundOK ", 200)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundOK) String() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/periods/{periodFid}/refund][%d] putCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundOK ", 200)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundDefault creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundDefault with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundDefault(code int) *PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundDefault {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundDefault{
		_statusCode: code,
	}
}

/*
PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundDefault describes a response with status code -1, with default header values.

Error
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put customers customer fid subscriptions subscription fid periods period fid refund default response
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this put customers customer fid subscriptions subscription fid periods period fid refund default response has a 2xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this put customers customer fid subscriptions subscription fid periods period fid refund default response has a 3xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this put customers customer fid subscriptions subscription fid periods period fid refund default response has a 4xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this put customers customer fid subscriptions subscription fid periods period fid refund default response has a 5xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this put customers customer fid subscriptions subscription fid periods period fid refund default response a status code equal to that given
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundDefault) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/periods/{periodFid}/refund][%d] PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefund default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundDefault) String() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/periods/{periodFid}/refund][%d] PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefund default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsPeriodFidRefundDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
