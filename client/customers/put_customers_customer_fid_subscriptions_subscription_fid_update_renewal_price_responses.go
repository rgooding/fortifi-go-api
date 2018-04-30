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

// PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceReader is a Reader for the PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPrice structure.
type PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOK creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOK with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOK() *PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOK {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOK{}
}

/*PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOK handles this case with default header values.

Subscription price updated
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOK struct {
	Payload *models.PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBody
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/updateRenewalPrice][%d] putCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOK  %+v", 200, o.Payload)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceBadRequest creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceBadRequest with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceBadRequest() *PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceBadRequest {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceBadRequest{}
}

/*PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceBadRequest handles this case with default header values.

Invalid payload data
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceBadRequest struct {
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceBadRequest) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/updateRenewalPrice][%d] putCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceBadRequest ", 400)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceNotFound creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceNotFound with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceNotFound() *PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceNotFound {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceNotFound{}
}

/*PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceNotFound handles this case with default header values.

Subscription not found
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceNotFound struct {
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceNotFound) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/updateRenewalPrice][%d] putCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceNotFound ", 404)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceDefault creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceDefault with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceDefault(code int) *PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceDefault {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceDefault{
		_statusCode: code,
	}
}

/*PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceDefault handles this case with default header values.

Error
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put customers customer fid subscriptions subscription fid update renewal price default response
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceDefault) Code() int {
	return o._statusCode
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceDefault) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/updateRenewalPrice][%d] PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPrice default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidUpdateRenewalPriceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
