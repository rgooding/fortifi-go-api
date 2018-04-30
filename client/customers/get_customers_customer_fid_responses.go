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

// GetCustomersCustomerFidReader is a Reader for the GetCustomersCustomerFid structure.
type GetCustomersCustomerFidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomersCustomerFidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCustomersCustomerFidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetCustomersCustomerFidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetCustomersCustomerFidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCustomersCustomerFidOK creates a GetCustomersCustomerFidOK with default headers values
func NewGetCustomersCustomerFidOK() *GetCustomersCustomerFidOK {
	return &GetCustomersCustomerFidOK{}
}

/*GetCustomersCustomerFidOK handles this case with default header values.

Loaded Customer
*/
type GetCustomersCustomerFidOK struct {
	Payload *models.GetCustomersCustomerFidOKBody
}

func (o *GetCustomersCustomerFidOK) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}][%d] getCustomersCustomerFidOK  %+v", 200, o.Payload)
}

func (o *GetCustomersCustomerFidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCustomersCustomerFidOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomersCustomerFidNotFound creates a GetCustomersCustomerFidNotFound with default headers values
func NewGetCustomersCustomerFidNotFound() *GetCustomersCustomerFidNotFound {
	return &GetCustomersCustomerFidNotFound{}
}

/*GetCustomersCustomerFidNotFound handles this case with default header values.

Customer not found
*/
type GetCustomersCustomerFidNotFound struct {
}

func (o *GetCustomersCustomerFidNotFound) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}][%d] getCustomersCustomerFidNotFound ", 404)
}

func (o *GetCustomersCustomerFidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCustomersCustomerFidDefault creates a GetCustomersCustomerFidDefault with default headers values
func NewGetCustomersCustomerFidDefault(code int) *GetCustomersCustomerFidDefault {
	return &GetCustomersCustomerFidDefault{
		_statusCode: code,
	}
}

/*GetCustomersCustomerFidDefault handles this case with default header values.

Error
*/
type GetCustomersCustomerFidDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get customers customer fid default response
func (o *GetCustomersCustomerFidDefault) Code() int {
	return o._statusCode
}

func (o *GetCustomersCustomerFidDefault) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}][%d] GetCustomersCustomerFid default  %+v", o._statusCode, o.Payload)
}

func (o *GetCustomersCustomerFidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
