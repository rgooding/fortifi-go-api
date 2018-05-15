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

// GetCustomersCustomerFidTicketsReader is a Reader for the GetCustomersCustomerFidTickets structure.
type GetCustomersCustomerFidTicketsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomersCustomerFidTicketsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCustomersCustomerFidTicketsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetCustomersCustomerFidTicketsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCustomersCustomerFidTicketsOK creates a GetCustomersCustomerFidTicketsOK with default headers values
func NewGetCustomersCustomerFidTicketsOK() *GetCustomersCustomerFidTicketsOK {
	return &GetCustomersCustomerFidTicketsOK{}
}

/*GetCustomersCustomerFidTicketsOK handles this case with default header values.

Ticket collection
*/
type GetCustomersCustomerFidTicketsOK struct {
	Payload *models.GetCustomersCustomerFidTicketsOKBody
}

func (o *GetCustomersCustomerFidTicketsOK) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/tickets][%d] getCustomersCustomerFidTicketsOK  %+v", 200, o.Payload)
}

func (o *GetCustomersCustomerFidTicketsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCustomersCustomerFidTicketsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomersCustomerFidTicketsDefault creates a GetCustomersCustomerFidTicketsDefault with default headers values
func NewGetCustomersCustomerFidTicketsDefault(code int) *GetCustomersCustomerFidTicketsDefault {
	return &GetCustomersCustomerFidTicketsDefault{
		_statusCode: code,
	}
}

/*GetCustomersCustomerFidTicketsDefault handles this case with default header values.

Error
*/
type GetCustomersCustomerFidTicketsDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get customers customer fid tickets default response
func (o *GetCustomersCustomerFidTicketsDefault) Code() int {
	return o._statusCode
}

func (o *GetCustomersCustomerFidTicketsDefault) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/tickets][%d] GetCustomersCustomerFidTickets default  %+v", o._statusCode, o.Payload)
}

func (o *GetCustomersCustomerFidTicketsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
