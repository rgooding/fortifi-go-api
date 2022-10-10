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

// PutCustomersCustomerFidPaymentMethodsCardsCardFidReader is a Reader for the PutCustomersCustomerFidPaymentMethodsCardsCardFid structure.
type PutCustomersCustomerFidPaymentMethodsCardsCardFidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCustomersCustomerFidPaymentMethodsCardsCardFidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutCustomersCustomerFidPaymentMethodsCardsCardFidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCustomersCustomerFidPaymentMethodsCardsCardFidOK creates a PutCustomersCustomerFidPaymentMethodsCardsCardFidOK with default headers values
func NewPutCustomersCustomerFidPaymentMethodsCardsCardFidOK() *PutCustomersCustomerFidPaymentMethodsCardsCardFidOK {
	return &PutCustomersCustomerFidPaymentMethodsCardsCardFidOK{}
}

/*
PutCustomersCustomerFidPaymentMethodsCardsCardFidOK describes a response with status code 200, with default header values.

Card Updated
*/
type PutCustomersCustomerFidPaymentMethodsCardsCardFidOK struct {
}

// IsSuccess returns true when this put customers customer fid payment methods cards card fid o k response has a 2xx status code
func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put customers customer fid payment methods cards card fid o k response has a 3xx status code
func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put customers customer fid payment methods cards card fid o k response has a 4xx status code
func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put customers customer fid payment methods cards card fid o k response has a 5xx status code
func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put customers customer fid payment methods cards card fid o k response a status code equal to that given
func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/paymentMethods/cards/{cardFid}][%d] putCustomersCustomerFidPaymentMethodsCardsCardFidOK ", 200)
}

func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidOK) String() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/paymentMethods/cards/{cardFid}][%d] putCustomersCustomerFidPaymentMethodsCardsCardFidOK ", 200)
}

func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomersCustomerFidPaymentMethodsCardsCardFidDefault creates a PutCustomersCustomerFidPaymentMethodsCardsCardFidDefault with default headers values
func NewPutCustomersCustomerFidPaymentMethodsCardsCardFidDefault(code int) *PutCustomersCustomerFidPaymentMethodsCardsCardFidDefault {
	return &PutCustomersCustomerFidPaymentMethodsCardsCardFidDefault{
		_statusCode: code,
	}
}

/*
PutCustomersCustomerFidPaymentMethodsCardsCardFidDefault describes a response with status code -1, with default header values.

Error
*/
type PutCustomersCustomerFidPaymentMethodsCardsCardFidDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put customers customer fid payment methods cards card fid default response
func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this put customers customer fid payment methods cards card fid default response has a 2xx status code
func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this put customers customer fid payment methods cards card fid default response has a 3xx status code
func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this put customers customer fid payment methods cards card fid default response has a 4xx status code
func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this put customers customer fid payment methods cards card fid default response has a 5xx status code
func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this put customers customer fid payment methods cards card fid default response a status code equal to that given
func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidDefault) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/paymentMethods/cards/{cardFid}][%d] PutCustomersCustomerFidPaymentMethodsCardsCardFid default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidDefault) String() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/paymentMethods/cards/{cardFid}][%d] PutCustomersCustomerFidPaymentMethodsCardsCardFid default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
