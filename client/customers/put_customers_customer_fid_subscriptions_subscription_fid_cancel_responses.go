// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelReader is a Reader for the PutCustomersCustomerFidSubscriptionsSubscriptionFidCancel structure.
type PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOK creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOK with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOK() *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOK {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOK{}
}

/*PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOK handles this case with default header values.

Subscription cancelled
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOK struct {
	Payload *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBody
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/cancel][%d] putCustomersCustomerFidSubscriptionsSubscriptionFidCancelOK  %+v", 200, o.Payload)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOK) GetPayload() *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBody {
	return o.Payload
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelDefault creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelDefault with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelDefault(code int) *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelDefault {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelDefault{
		_statusCode: code,
	}
}

/*PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelDefault handles this case with default header values.

Error
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put customers customer fid subscriptions subscription fid cancel default response
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelDefault) Code() int {
	return o._statusCode
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelDefault) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/cancel][%d] PutCustomersCustomerFidSubscriptionsSubscriptionFidCancel default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBody put customers customer fid subscriptions subscription fid cancel o k body
swagger:model PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBody
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBody struct {
	models.Envelope

	// data
	Data *models.Subscription `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBody) UnmarshalJSON(raw []byte) error {
	// PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBodyAO0
	var putCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &putCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = putCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBodyAO0

	// PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBodyAO1
	var dataPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBodyAO1 struct {
		Data *models.Subscription `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	putCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, putCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBodyAO0)

	var dataPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBodyAO1 struct {
		Data *models.Subscription `json:"data,omitempty"`
	}

	dataPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBodyAO1.Data = o.Data

	jsonDataPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBodyAO1, errPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBodyAO1 := swag.WriteJSON(dataPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBodyAO1)
	if errPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBodyAO1 != nil {
		return nil, errPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBodyAO1
	}
	_parts = append(_parts, jsonDataPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this put customers customer fid subscriptions subscription fid cancel o k body
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Envelope
	if err := o.Envelope.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putCustomersCustomerFidSubscriptionsSubscriptionFidCancelOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBody) UnmarshalBinary(b []byte) error {
	var res PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
