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

// PutCustomersCustomerFidAccountStatusReader is a Reader for the PutCustomersCustomerFidAccountStatus structure.
type PutCustomersCustomerFidAccountStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidAccountStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCustomersCustomerFidAccountStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutCustomersCustomerFidAccountStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCustomersCustomerFidAccountStatusOK creates a PutCustomersCustomerFidAccountStatusOK with default headers values
func NewPutCustomersCustomerFidAccountStatusOK() *PutCustomersCustomerFidAccountStatusOK {
	return &PutCustomersCustomerFidAccountStatusOK{}
}

/* PutCustomersCustomerFidAccountStatusOK describes a response with status code 200, with default header values.

Customer Status Updated
*/
type PutCustomersCustomerFidAccountStatusOK struct {
}

func (o *PutCustomersCustomerFidAccountStatusOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/accountStatus][%d] putCustomersCustomerFidAccountStatusOK ", 200)
}

func (o *PutCustomersCustomerFidAccountStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomersCustomerFidAccountStatusDefault creates a PutCustomersCustomerFidAccountStatusDefault with default headers values
func NewPutCustomersCustomerFidAccountStatusDefault(code int) *PutCustomersCustomerFidAccountStatusDefault {
	return &PutCustomersCustomerFidAccountStatusDefault{
		_statusCode: code,
	}
}

/* PutCustomersCustomerFidAccountStatusDefault describes a response with status code -1, with default header values.

Error
*/
type PutCustomersCustomerFidAccountStatusDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put customers customer fid account status default response
func (o *PutCustomersCustomerFidAccountStatusDefault) Code() int {
	return o._statusCode
}

func (o *PutCustomersCustomerFidAccountStatusDefault) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/accountStatus][%d] PutCustomersCustomerFidAccountStatus default  %+v", o._statusCode, o.Payload)
}
func (o *PutCustomersCustomerFidAccountStatusDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutCustomersCustomerFidAccountStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
