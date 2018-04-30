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

// GetOrdersOrderFidProductsReader is a Reader for the GetOrdersOrderFidProducts structure.
type GetOrdersOrderFidProductsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrdersOrderFidProductsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOrdersOrderFidProductsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetOrdersOrderFidProductsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetOrdersOrderFidProductsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOrdersOrderFidProductsOK creates a GetOrdersOrderFidProductsOK with default headers values
func NewGetOrdersOrderFidProductsOK() *GetOrdersOrderFidProductsOK {
	return &GetOrdersOrderFidProductsOK{}
}

/*GetOrdersOrderFidProductsOK handles this case with default header values.

List of order products
*/
type GetOrdersOrderFidProductsOK struct {
	Payload *models.GetOrdersOrderFidProductsOKBody
}

func (o *GetOrdersOrderFidProductsOK) Error() string {
	return fmt.Sprintf("[GET /orders/{orderFid}/products][%d] getOrdersOrderFidProductsOK  %+v", 200, o.Payload)
}

func (o *GetOrdersOrderFidProductsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetOrdersOrderFidProductsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrdersOrderFidProductsNotFound creates a GetOrdersOrderFidProductsNotFound with default headers values
func NewGetOrdersOrderFidProductsNotFound() *GetOrdersOrderFidProductsNotFound {
	return &GetOrdersOrderFidProductsNotFound{}
}

/*GetOrdersOrderFidProductsNotFound handles this case with default header values.

Order not found
*/
type GetOrdersOrderFidProductsNotFound struct {
}

func (o *GetOrdersOrderFidProductsNotFound) Error() string {
	return fmt.Sprintf("[GET /orders/{orderFid}/products][%d] getOrdersOrderFidProductsNotFound ", 404)
}

func (o *GetOrdersOrderFidProductsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrdersOrderFidProductsDefault creates a GetOrdersOrderFidProductsDefault with default headers values
func NewGetOrdersOrderFidProductsDefault(code int) *GetOrdersOrderFidProductsDefault {
	return &GetOrdersOrderFidProductsDefault{
		_statusCode: code,
	}
}

/*GetOrdersOrderFidProductsDefault handles this case with default header values.

Error
*/
type GetOrdersOrderFidProductsDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get orders order fid products default response
func (o *GetOrdersOrderFidProductsDefault) Code() int {
	return o._statusCode
}

func (o *GetOrdersOrderFidProductsDefault) Error() string {
	return fmt.Sprintf("[GET /orders/{orderFid}/products][%d] GetOrdersOrderFidProducts default  %+v", o._statusCode, o.Payload)
}

func (o *GetOrdersOrderFidProductsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
