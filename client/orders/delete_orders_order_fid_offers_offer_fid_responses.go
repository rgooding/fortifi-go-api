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

// DeleteOrdersOrderFidOffersOfferFidReader is a Reader for the DeleteOrdersOrderFidOffersOfferFid structure.
type DeleteOrdersOrderFidOffersOfferFidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrdersOrderFidOffersOfferFidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteOrdersOrderFidOffersOfferFidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteOrdersOrderFidOffersOfferFidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteOrdersOrderFidOffersOfferFidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteOrdersOrderFidOffersOfferFidOK creates a DeleteOrdersOrderFidOffersOfferFidOK with default headers values
func NewDeleteOrdersOrderFidOffersOfferFidOK() *DeleteOrdersOrderFidOffersOfferFidOK {
	return &DeleteOrdersOrderFidOffersOfferFidOK{}
}

/*DeleteOrdersOrderFidOffersOfferFidOK handles this case with default header values.

Offer removed from the order successfully
*/
type DeleteOrdersOrderFidOffersOfferFidOK struct {
}

func (o *DeleteOrdersOrderFidOffersOfferFidOK) Error() string {
	return fmt.Sprintf("[DELETE /orders/{orderFid}/offers/{offerFid}][%d] deleteOrdersOrderFidOffersOfferFidOK ", 200)
}

func (o *DeleteOrdersOrderFidOffersOfferFidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrdersOrderFidOffersOfferFidNotFound creates a DeleteOrdersOrderFidOffersOfferFidNotFound with default headers values
func NewDeleteOrdersOrderFidOffersOfferFidNotFound() *DeleteOrdersOrderFidOffersOfferFidNotFound {
	return &DeleteOrdersOrderFidOffersOfferFidNotFound{}
}

/*DeleteOrdersOrderFidOffersOfferFidNotFound handles this case with default header values.

Order not found
*/
type DeleteOrdersOrderFidOffersOfferFidNotFound struct {
}

func (o *DeleteOrdersOrderFidOffersOfferFidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /orders/{orderFid}/offers/{offerFid}][%d] deleteOrdersOrderFidOffersOfferFidNotFound ", 404)
}

func (o *DeleteOrdersOrderFidOffersOfferFidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrdersOrderFidOffersOfferFidDefault creates a DeleteOrdersOrderFidOffersOfferFidDefault with default headers values
func NewDeleteOrdersOrderFidOffersOfferFidDefault(code int) *DeleteOrdersOrderFidOffersOfferFidDefault {
	return &DeleteOrdersOrderFidOffersOfferFidDefault{
		_statusCode: code,
	}
}

/*DeleteOrdersOrderFidOffersOfferFidDefault handles this case with default header values.

Error
*/
type DeleteOrdersOrderFidOffersOfferFidDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the delete orders order fid offers offer fid default response
func (o *DeleteOrdersOrderFidOffersOfferFidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteOrdersOrderFidOffersOfferFidDefault) Error() string {
	return fmt.Sprintf("[DELETE /orders/{orderFid}/offers/{offerFid}][%d] DeleteOrdersOrderFidOffersOfferFid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteOrdersOrderFidOffersOfferFidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
