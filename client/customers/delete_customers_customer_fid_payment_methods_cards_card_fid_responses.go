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

// DeleteCustomersCustomerFidPaymentMethodsCardsCardFidReader is a Reader for the DeleteCustomersCustomerFidPaymentMethodsCardsCardFid structure.
type DeleteCustomersCustomerFidPaymentMethodsCardsCardFidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCustomersCustomerFidPaymentMethodsCardsCardFidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteCustomersCustomerFidPaymentMethodsCardsCardFidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCustomersCustomerFidPaymentMethodsCardsCardFidOK creates a DeleteCustomersCustomerFidPaymentMethodsCardsCardFidOK with default headers values
func NewDeleteCustomersCustomerFidPaymentMethodsCardsCardFidOK() *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidOK {
	return &DeleteCustomersCustomerFidPaymentMethodsCardsCardFidOK{}
}

/*
DeleteCustomersCustomerFidPaymentMethodsCardsCardFidOK describes a response with status code 200, with default header values.

Card Removed
*/
type DeleteCustomersCustomerFidPaymentMethodsCardsCardFidOK struct {
}

// IsSuccess returns true when this delete customers customer fid payment methods cards card fid o k response has a 2xx status code
func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete customers customer fid payment methods cards card fid o k response has a 3xx status code
func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete customers customer fid payment methods cards card fid o k response has a 4xx status code
func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete customers customer fid payment methods cards card fid o k response has a 5xx status code
func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete customers customer fid payment methods cards card fid o k response a status code equal to that given
func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete customers customer fid payment methods cards card fid o k response
func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidOK) Code() int {
	return 200
}

func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidOK) Error() string {
	return fmt.Sprintf("[DELETE /customers/{customerFid}/paymentMethods/cards/{cardFid}][%d] deleteCustomersCustomerFidPaymentMethodsCardsCardFidOK ", 200)
}

func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidOK) String() string {
	return fmt.Sprintf("[DELETE /customers/{customerFid}/paymentMethods/cards/{cardFid}][%d] deleteCustomersCustomerFidPaymentMethodsCardsCardFidOK ", 200)
}

func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCustomersCustomerFidPaymentMethodsCardsCardFidDefault creates a DeleteCustomersCustomerFidPaymentMethodsCardsCardFidDefault with default headers values
func NewDeleteCustomersCustomerFidPaymentMethodsCardsCardFidDefault(code int) *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidDefault {
	return &DeleteCustomersCustomerFidPaymentMethodsCardsCardFidDefault{
		_statusCode: code,
	}
}

/*
DeleteCustomersCustomerFidPaymentMethodsCardsCardFidDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteCustomersCustomerFidPaymentMethodsCardsCardFidDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this delete customers customer fid payment methods cards card fid default response has a 2xx status code
func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete customers customer fid payment methods cards card fid default response has a 3xx status code
func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete customers customer fid payment methods cards card fid default response has a 4xx status code
func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete customers customer fid payment methods cards card fid default response has a 5xx status code
func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete customers customer fid payment methods cards card fid default response a status code equal to that given
func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete customers customer fid payment methods cards card fid default response
func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidDefault) Error() string {
	return fmt.Sprintf("[DELETE /customers/{customerFid}/paymentMethods/cards/{cardFid}][%d] DeleteCustomersCustomerFidPaymentMethodsCardsCardFid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidDefault) String() string {
	return fmt.Sprintf("[DELETE /customers/{customerFid}/paymentMethods/cards/{cardFid}][%d] DeleteCustomersCustomerFidPaymentMethodsCardsCardFid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
