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

// PutCustomersCustomerFidLoyalReader is a Reader for the PutCustomersCustomerFidLoyal structure.
type PutCustomersCustomerFidLoyalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidLoyalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCustomersCustomerFidLoyalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutCustomersCustomerFidLoyalDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCustomersCustomerFidLoyalOK creates a PutCustomersCustomerFidLoyalOK with default headers values
func NewPutCustomersCustomerFidLoyalOK() *PutCustomersCustomerFidLoyalOK {
	return &PutCustomersCustomerFidLoyalOK{}
}

/*
PutCustomersCustomerFidLoyalOK describes a response with status code 200, with default header values.

Customer Loyalty Acknowledged
*/
type PutCustomersCustomerFidLoyalOK struct {
}

// IsSuccess returns true when this put customers customer fid loyal o k response has a 2xx status code
func (o *PutCustomersCustomerFidLoyalOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put customers customer fid loyal o k response has a 3xx status code
func (o *PutCustomersCustomerFidLoyalOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put customers customer fid loyal o k response has a 4xx status code
func (o *PutCustomersCustomerFidLoyalOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put customers customer fid loyal o k response has a 5xx status code
func (o *PutCustomersCustomerFidLoyalOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put customers customer fid loyal o k response a status code equal to that given
func (o *PutCustomersCustomerFidLoyalOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutCustomersCustomerFidLoyalOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/loyal][%d] putCustomersCustomerFidLoyalOK ", 200)
}

func (o *PutCustomersCustomerFidLoyalOK) String() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/loyal][%d] putCustomersCustomerFidLoyalOK ", 200)
}

func (o *PutCustomersCustomerFidLoyalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomersCustomerFidLoyalDefault creates a PutCustomersCustomerFidLoyalDefault with default headers values
func NewPutCustomersCustomerFidLoyalDefault(code int) *PutCustomersCustomerFidLoyalDefault {
	return &PutCustomersCustomerFidLoyalDefault{
		_statusCode: code,
	}
}

/*
PutCustomersCustomerFidLoyalDefault describes a response with status code -1, with default header values.

Error
*/
type PutCustomersCustomerFidLoyalDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put customers customer fid loyal default response
func (o *PutCustomersCustomerFidLoyalDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this put customers customer fid loyal default response has a 2xx status code
func (o *PutCustomersCustomerFidLoyalDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this put customers customer fid loyal default response has a 3xx status code
func (o *PutCustomersCustomerFidLoyalDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this put customers customer fid loyal default response has a 4xx status code
func (o *PutCustomersCustomerFidLoyalDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this put customers customer fid loyal default response has a 5xx status code
func (o *PutCustomersCustomerFidLoyalDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this put customers customer fid loyal default response a status code equal to that given
func (o *PutCustomersCustomerFidLoyalDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PutCustomersCustomerFidLoyalDefault) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/loyal][%d] PutCustomersCustomerFidLoyal default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidLoyalDefault) String() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/loyal][%d] PutCustomersCustomerFidLoyal default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidLoyalDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutCustomersCustomerFidLoyalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
