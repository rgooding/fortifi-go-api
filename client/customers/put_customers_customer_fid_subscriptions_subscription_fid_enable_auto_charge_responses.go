// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeReader is a Reader for the PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoCharge structure.
type PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeOK creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeOK with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeOK() *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeOK {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeOK{}
}

/*PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeOK handles this case with default header values.

Auto charge enabled
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeOK struct {
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/enableAutoCharge][%d] putCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeOK ", 200)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeDefault creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeDefault with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeDefault(code int) *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeDefault {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeDefault{
		_statusCode: code,
	}
}

/*PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeDefault handles this case with default header values.

Error
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put customers customer fid subscriptions subscription fid enable auto charge default response
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeDefault) Code() int {
	return o._statusCode
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeDefault) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/enableAutoCharge][%d] PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoCharge default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
