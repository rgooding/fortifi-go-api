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

// PostCustomersCustomerFidTicketsReader is a Reader for the PostCustomersCustomerFidTickets structure.
type PostCustomersCustomerFidTicketsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCustomersCustomerFidTicketsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostCustomersCustomerFidTicketsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostCustomersCustomerFidTicketsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCustomersCustomerFidTicketsOK creates a PostCustomersCustomerFidTicketsOK with default headers values
func NewPostCustomersCustomerFidTicketsOK() *PostCustomersCustomerFidTicketsOK {
	return &PostCustomersCustomerFidTicketsOK{}
}

/* PostCustomersCustomerFidTicketsOK describes a response with status code 200, with default header values.

Ticket Created
*/
type PostCustomersCustomerFidTicketsOK struct {
}

func (o *PostCustomersCustomerFidTicketsOK) Error() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/tickets][%d] postCustomersCustomerFidTicketsOK ", 200)
}

func (o *PostCustomersCustomerFidTicketsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomersCustomerFidTicketsDefault creates a PostCustomersCustomerFidTicketsDefault with default headers values
func NewPostCustomersCustomerFidTicketsDefault(code int) *PostCustomersCustomerFidTicketsDefault {
	return &PostCustomersCustomerFidTicketsDefault{
		_statusCode: code,
	}
}

/* PostCustomersCustomerFidTicketsDefault describes a response with status code -1, with default header values.

Error
*/
type PostCustomersCustomerFidTicketsDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the post customers customer fid tickets default response
func (o *PostCustomersCustomerFidTicketsDefault) Code() int {
	return o._statusCode
}

func (o *PostCustomersCustomerFidTicketsDefault) Error() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/tickets][%d] PostCustomersCustomerFidTickets default  %+v", o._statusCode, o.Payload)
}
func (o *PostCustomersCustomerFidTicketsDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostCustomersCustomerFidTicketsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
