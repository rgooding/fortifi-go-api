// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// PostCustomersCustomerFidPaymentMethodsPaypalInitialiseReader is a Reader for the PostCustomersCustomerFidPaymentMethodsPaypalInitialise structure.
type PostCustomersCustomerFidPaymentMethodsPaypalInitialiseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostCustomersCustomerFidPaymentMethodsPaypalInitialiseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostCustomersCustomerFidPaymentMethodsPaypalInitialiseDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCustomersCustomerFidPaymentMethodsPaypalInitialiseOK creates a PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOK with default headers values
func NewPostCustomersCustomerFidPaymentMethodsPaypalInitialiseOK() *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOK {
	return &PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOK{}
}

/*PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOK handles this case with default header values.

Redirect instructions
*/
type PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOK struct {
	Payload *models.PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBody
}

func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOK) Error() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/paymentMethods/paypal/initialise][%d] postCustomersCustomerFidPaymentMethodsPaypalInitialiseOK  %+v", 200, o.Payload)
}

func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostCustomersCustomerFidPaymentMethodsPaypalInitialiseOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCustomersCustomerFidPaymentMethodsPaypalInitialiseDefault creates a PostCustomersCustomerFidPaymentMethodsPaypalInitialiseDefault with default headers values
func NewPostCustomersCustomerFidPaymentMethodsPaypalInitialiseDefault(code int) *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseDefault {
	return &PostCustomersCustomerFidPaymentMethodsPaypalInitialiseDefault{
		_statusCode: code,
	}
}

/*PostCustomersCustomerFidPaymentMethodsPaypalInitialiseDefault handles this case with default header values.

Error
*/
type PostCustomersCustomerFidPaymentMethodsPaypalInitialiseDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the post customers customer fid payment methods paypal initialise default response
func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseDefault) Code() int {
	return o._statusCode
}

func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseDefault) Error() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/paymentMethods/paypal/initialise][%d] PostCustomersCustomerFidPaymentMethodsPaypalInitialise default  %+v", o._statusCode, o.Payload)
}

func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
