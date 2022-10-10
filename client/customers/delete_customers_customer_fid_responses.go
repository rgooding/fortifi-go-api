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

// DeleteCustomersCustomerFidReader is a Reader for the DeleteCustomersCustomerFid structure.
type DeleteCustomersCustomerFidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCustomersCustomerFidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCustomersCustomerFidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteCustomersCustomerFidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCustomersCustomerFidOK creates a DeleteCustomersCustomerFidOK with default headers values
func NewDeleteCustomersCustomerFidOK() *DeleteCustomersCustomerFidOK {
	return &DeleteCustomersCustomerFidOK{}
}

/*
DeleteCustomersCustomerFidOK describes a response with status code 200, with default header values.

Customer Archived
*/
type DeleteCustomersCustomerFidOK struct {
}

// IsSuccess returns true when this delete customers customer fid o k response has a 2xx status code
func (o *DeleteCustomersCustomerFidOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete customers customer fid o k response has a 3xx status code
func (o *DeleteCustomersCustomerFidOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete customers customer fid o k response has a 4xx status code
func (o *DeleteCustomersCustomerFidOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete customers customer fid o k response has a 5xx status code
func (o *DeleteCustomersCustomerFidOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete customers customer fid o k response a status code equal to that given
func (o *DeleteCustomersCustomerFidOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteCustomersCustomerFidOK) Error() string {
	return fmt.Sprintf("[DELETE /customers/{customerFid}][%d] deleteCustomersCustomerFidOK ", 200)
}

func (o *DeleteCustomersCustomerFidOK) String() string {
	return fmt.Sprintf("[DELETE /customers/{customerFid}][%d] deleteCustomersCustomerFidOK ", 200)
}

func (o *DeleteCustomersCustomerFidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCustomersCustomerFidDefault creates a DeleteCustomersCustomerFidDefault with default headers values
func NewDeleteCustomersCustomerFidDefault(code int) *DeleteCustomersCustomerFidDefault {
	return &DeleteCustomersCustomerFidDefault{
		_statusCode: code,
	}
}

/*
DeleteCustomersCustomerFidDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteCustomersCustomerFidDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the delete customers customer fid default response
func (o *DeleteCustomersCustomerFidDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete customers customer fid default response has a 2xx status code
func (o *DeleteCustomersCustomerFidDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete customers customer fid default response has a 3xx status code
func (o *DeleteCustomersCustomerFidDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete customers customer fid default response has a 4xx status code
func (o *DeleteCustomersCustomerFidDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete customers customer fid default response has a 5xx status code
func (o *DeleteCustomersCustomerFidDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete customers customer fid default response a status code equal to that given
func (o *DeleteCustomersCustomerFidDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteCustomersCustomerFidDefault) Error() string {
	return fmt.Sprintf("[DELETE /customers/{customerFid}][%d] DeleteCustomersCustomerFid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteCustomersCustomerFidDefault) String() string {
	return fmt.Sprintf("[DELETE /customers/{customerFid}][%d] DeleteCustomersCustomerFid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteCustomersCustomerFidDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *DeleteCustomersCustomerFidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
