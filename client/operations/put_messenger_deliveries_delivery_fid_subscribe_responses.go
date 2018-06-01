// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/fortifi/go-api/models"
)

// PutMessengerDeliveriesDeliveryFidSubscribeReader is a Reader for the PutMessengerDeliveriesDeliveryFidSubscribe structure.
type PutMessengerDeliveriesDeliveryFidSubscribeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutMessengerDeliveriesDeliveryFidSubscribeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutMessengerDeliveriesDeliveryFidSubscribeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPutMessengerDeliveriesDeliveryFidSubscribeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutMessengerDeliveriesDeliveryFidSubscribeOK creates a PutMessengerDeliveriesDeliveryFidSubscribeOK with default headers values
func NewPutMessengerDeliveriesDeliveryFidSubscribeOK() *PutMessengerDeliveriesDeliveryFidSubscribeOK {
	return &PutMessengerDeliveriesDeliveryFidSubscribeOK{}
}

/*PutMessengerDeliveriesDeliveryFidSubscribeOK handles this case with default header values.

Email Address Subscribed
*/
type PutMessengerDeliveriesDeliveryFidSubscribeOK struct {
}

func (o *PutMessengerDeliveriesDeliveryFidSubscribeOK) Error() string {
	return fmt.Sprintf("[PUT /messenger/deliveries/{deliveryFid}/subscribe][%d] putMessengerDeliveriesDeliveryFidSubscribeOK ", 200)
}

func (o *PutMessengerDeliveriesDeliveryFidSubscribeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutMessengerDeliveriesDeliveryFidSubscribeDefault creates a PutMessengerDeliveriesDeliveryFidSubscribeDefault with default headers values
func NewPutMessengerDeliveriesDeliveryFidSubscribeDefault(code int) *PutMessengerDeliveriesDeliveryFidSubscribeDefault {
	return &PutMessengerDeliveriesDeliveryFidSubscribeDefault{
		_statusCode: code,
	}
}

/*PutMessengerDeliveriesDeliveryFidSubscribeDefault handles this case with default header values.

Error
*/
type PutMessengerDeliveriesDeliveryFidSubscribeDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put messenger deliveries delivery fid subscribe default response
func (o *PutMessengerDeliveriesDeliveryFidSubscribeDefault) Code() int {
	return o._statusCode
}

func (o *PutMessengerDeliveriesDeliveryFidSubscribeDefault) Error() string {
	return fmt.Sprintf("[PUT /messenger/deliveries/{deliveryFid}/subscribe][%d] PutMessengerDeliveriesDeliveryFidSubscribe default  %+v", o._statusCode, o.Payload)
}

func (o *PutMessengerDeliveriesDeliveryFidSubscribeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}