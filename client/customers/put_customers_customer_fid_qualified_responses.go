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

// PutCustomersCustomerFidQualifiedReader is a Reader for the PutCustomersCustomerFidQualified structure.
type PutCustomersCustomerFidQualifiedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidQualifiedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCustomersCustomerFidQualifiedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutCustomersCustomerFidQualifiedDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCustomersCustomerFidQualifiedOK creates a PutCustomersCustomerFidQualifiedOK with default headers values
func NewPutCustomersCustomerFidQualifiedOK() *PutCustomersCustomerFidQualifiedOK {
	return &PutCustomersCustomerFidQualifiedOK{}
}

/*PutCustomersCustomerFidQualifiedOK handles this case with default header values.

Customer Marked
*/
type PutCustomersCustomerFidQualifiedOK struct {
}

func (o *PutCustomersCustomerFidQualifiedOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/qualified][%d] putCustomersCustomerFidQualifiedOK ", 200)
}

func (o *PutCustomersCustomerFidQualifiedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomersCustomerFidQualifiedDefault creates a PutCustomersCustomerFidQualifiedDefault with default headers values
func NewPutCustomersCustomerFidQualifiedDefault(code int) *PutCustomersCustomerFidQualifiedDefault {
	return &PutCustomersCustomerFidQualifiedDefault{
		_statusCode: code,
	}
}

/*PutCustomersCustomerFidQualifiedDefault handles this case with default header values.

Error
*/
type PutCustomersCustomerFidQualifiedDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put customers customer fid qualified default response
func (o *PutCustomersCustomerFidQualifiedDefault) Code() int {
	return o._statusCode
}

func (o *PutCustomersCustomerFidQualifiedDefault) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/qualified][%d] PutCustomersCustomerFidQualified default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidQualifiedDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutCustomersCustomerFidQualifiedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
