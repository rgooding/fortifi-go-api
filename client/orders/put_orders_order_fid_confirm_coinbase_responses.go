// Code generated by go-swagger; DO NOT EDIT.

package orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// PutOrdersOrderFidConfirmCoinbaseReader is a Reader for the PutOrdersOrderFidConfirmCoinbase structure.
type PutOrdersOrderFidConfirmCoinbaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOrdersOrderFidConfirmCoinbaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutOrdersOrderFidConfirmCoinbaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 503:
		result := NewPutOrdersOrderFidConfirmCoinbaseServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutOrdersOrderFidConfirmCoinbaseDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutOrdersOrderFidConfirmCoinbaseOK creates a PutOrdersOrderFidConfirmCoinbaseOK with default headers values
func NewPutOrdersOrderFidConfirmCoinbaseOK() *PutOrdersOrderFidConfirmCoinbaseOK {
	return &PutOrdersOrderFidConfirmCoinbaseOK{}
}

/*PutOrdersOrderFidConfirmCoinbaseOK handles this case with default header values.

Order confirmed; awaiting blockchain confirmation.
*/
type PutOrdersOrderFidConfirmCoinbaseOK struct {
	Payload *models.PutOrdersOrderFidConfirmCoinbaseOKBody
}

func (o *PutOrdersOrderFidConfirmCoinbaseOK) Error() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/confirmCoinbase][%d] putOrdersOrderFidConfirmCoinbaseOK  %+v", 200, o.Payload)
}

func (o *PutOrdersOrderFidConfirmCoinbaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PutOrdersOrderFidConfirmCoinbaseOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrdersOrderFidConfirmCoinbaseServiceUnavailable creates a PutOrdersOrderFidConfirmCoinbaseServiceUnavailable with default headers values
func NewPutOrdersOrderFidConfirmCoinbaseServiceUnavailable() *PutOrdersOrderFidConfirmCoinbaseServiceUnavailable {
	return &PutOrdersOrderFidConfirmCoinbaseServiceUnavailable{}
}

/*PutOrdersOrderFidConfirmCoinbaseServiceUnavailable handles this case with default header values.

There are no payment gateways available to handle your request
*/
type PutOrdersOrderFidConfirmCoinbaseServiceUnavailable struct {
}

func (o *PutOrdersOrderFidConfirmCoinbaseServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/confirmCoinbase][%d] putOrdersOrderFidConfirmCoinbaseServiceUnavailable ", 503)
}

func (o *PutOrdersOrderFidConfirmCoinbaseServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutOrdersOrderFidConfirmCoinbaseDefault creates a PutOrdersOrderFidConfirmCoinbaseDefault with default headers values
func NewPutOrdersOrderFidConfirmCoinbaseDefault(code int) *PutOrdersOrderFidConfirmCoinbaseDefault {
	return &PutOrdersOrderFidConfirmCoinbaseDefault{
		_statusCode: code,
	}
}

/*PutOrdersOrderFidConfirmCoinbaseDefault handles this case with default header values.

Error
*/
type PutOrdersOrderFidConfirmCoinbaseDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put orders order fid confirm coinbase default response
func (o *PutOrdersOrderFidConfirmCoinbaseDefault) Code() int {
	return o._statusCode
}

func (o *PutOrdersOrderFidConfirmCoinbaseDefault) Error() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/confirmCoinbase][%d] PutOrdersOrderFidConfirmCoinbase default  %+v", o._statusCode, o.Payload)
}

func (o *PutOrdersOrderFidConfirmCoinbaseDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
