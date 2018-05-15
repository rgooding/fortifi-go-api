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

/*DeleteCustomersCustomerFidOK handles this case with default header values.

Customer Archived
*/
type DeleteCustomersCustomerFidOK struct {
}

func (o *DeleteCustomersCustomerFidOK) Error() string {
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

/*DeleteCustomersCustomerFidDefault handles this case with default header values.

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

func (o *DeleteCustomersCustomerFidDefault) Error() string {
	return fmt.Sprintf("[DELETE /customers/{customerFid}][%d] DeleteCustomersCustomerFid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteCustomersCustomerFidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
