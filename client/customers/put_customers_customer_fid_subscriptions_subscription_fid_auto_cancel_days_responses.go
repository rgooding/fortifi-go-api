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

// PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoCancelDaysReader is a Reader for the PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoCancelDays structure.
type PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoCancelDaysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoCancelDaysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidAutoCancelDaysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidAutoCancelDaysDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidAutoCancelDaysOK creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoCancelDaysOK with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidAutoCancelDaysOK() *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoCancelDaysOK {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoCancelDaysOK{}
}

/*PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoCancelDaysOK handles this case with default header values.

Auto cancel days set
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoCancelDaysOK struct {
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoCancelDaysOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/autoCancelDays][%d] putCustomersCustomerFidSubscriptionsSubscriptionFidAutoCancelDaysOK ", 200)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoCancelDaysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidAutoCancelDaysDefault creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoCancelDaysDefault with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidAutoCancelDaysDefault(code int) *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoCancelDaysDefault {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoCancelDaysDefault{
		_statusCode: code,
	}
}

/*PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoCancelDaysDefault handles this case with default header values.

Error
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoCancelDaysDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put customers customer fid subscriptions subscription fid auto cancel days default response
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoCancelDaysDefault) Code() int {
	return o._statusCode
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoCancelDaysDefault) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/autoCancelDays][%d] PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoCancelDays default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoCancelDaysDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidAutoCancelDaysDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}