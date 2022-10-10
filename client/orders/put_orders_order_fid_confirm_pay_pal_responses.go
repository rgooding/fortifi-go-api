// Code generated by go-swagger; DO NOT EDIT.

package orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/fortifi/go-api/models"
)

// PutOrdersOrderFidConfirmPayPalReader is a Reader for the PutOrdersOrderFidConfirmPayPal structure.
type PutOrdersOrderFidConfirmPayPalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOrdersOrderFidConfirmPayPalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutOrdersOrderFidConfirmPayPalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutOrdersOrderFidConfirmPayPalDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutOrdersOrderFidConfirmPayPalOK creates a PutOrdersOrderFidConfirmPayPalOK with default headers values
func NewPutOrdersOrderFidConfirmPayPalOK() *PutOrdersOrderFidConfirmPayPalOK {
	return &PutOrdersOrderFidConfirmPayPalOK{}
}

/*
PutOrdersOrderFidConfirmPayPalOK describes a response with status code 200, with default header values.

Order confirmed and payment authroized
*/
type PutOrdersOrderFidConfirmPayPalOK struct {
	Payload *PutOrdersOrderFidConfirmPayPalOKBody
}

// IsSuccess returns true when this put orders order fid confirm pay pal o k response has a 2xx status code
func (o *PutOrdersOrderFidConfirmPayPalOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put orders order fid confirm pay pal o k response has a 3xx status code
func (o *PutOrdersOrderFidConfirmPayPalOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put orders order fid confirm pay pal o k response has a 4xx status code
func (o *PutOrdersOrderFidConfirmPayPalOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put orders order fid confirm pay pal o k response has a 5xx status code
func (o *PutOrdersOrderFidConfirmPayPalOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put orders order fid confirm pay pal o k response a status code equal to that given
func (o *PutOrdersOrderFidConfirmPayPalOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutOrdersOrderFidConfirmPayPalOK) Error() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/confirmPayPal][%d] putOrdersOrderFidConfirmPayPalOK  %+v", 200, o.Payload)
}

func (o *PutOrdersOrderFidConfirmPayPalOK) String() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/confirmPayPal][%d] putOrdersOrderFidConfirmPayPalOK  %+v", 200, o.Payload)
}

func (o *PutOrdersOrderFidConfirmPayPalOK) GetPayload() *PutOrdersOrderFidConfirmPayPalOKBody {
	return o.Payload
}

func (o *PutOrdersOrderFidConfirmPayPalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutOrdersOrderFidConfirmPayPalOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrdersOrderFidConfirmPayPalDefault creates a PutOrdersOrderFidConfirmPayPalDefault with default headers values
func NewPutOrdersOrderFidConfirmPayPalDefault(code int) *PutOrdersOrderFidConfirmPayPalDefault {
	return &PutOrdersOrderFidConfirmPayPalDefault{
		_statusCode: code,
	}
}

/*
PutOrdersOrderFidConfirmPayPalDefault describes a response with status code -1, with default header values.

Error
*/
type PutOrdersOrderFidConfirmPayPalDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put orders order fid confirm pay pal default response
func (o *PutOrdersOrderFidConfirmPayPalDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this put orders order fid confirm pay pal default response has a 2xx status code
func (o *PutOrdersOrderFidConfirmPayPalDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this put orders order fid confirm pay pal default response has a 3xx status code
func (o *PutOrdersOrderFidConfirmPayPalDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this put orders order fid confirm pay pal default response has a 4xx status code
func (o *PutOrdersOrderFidConfirmPayPalDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this put orders order fid confirm pay pal default response has a 5xx status code
func (o *PutOrdersOrderFidConfirmPayPalDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this put orders order fid confirm pay pal default response a status code equal to that given
func (o *PutOrdersOrderFidConfirmPayPalDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PutOrdersOrderFidConfirmPayPalDefault) Error() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/confirmPayPal][%d] PutOrdersOrderFidConfirmPayPal default  %+v", o._statusCode, o.Payload)
}

func (o *PutOrdersOrderFidConfirmPayPalDefault) String() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/confirmPayPal][%d] PutOrdersOrderFidConfirmPayPal default  %+v", o._statusCode, o.Payload)
}

func (o *PutOrdersOrderFidConfirmPayPalDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutOrdersOrderFidConfirmPayPalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PutOrdersOrderFidConfirmPayPalOKBody put orders order fid confirm pay pal o k body
swagger:model PutOrdersOrderFidConfirmPayPalOKBody
*/
type PutOrdersOrderFidConfirmPayPalOKBody struct {
	models.Envelope

	// data
	Data *models.OrderConfirmation `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PutOrdersOrderFidConfirmPayPalOKBody) UnmarshalJSON(raw []byte) error {
	// PutOrdersOrderFidConfirmPayPalOKBodyAO0
	var putOrdersOrderFidConfirmPayPalOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &putOrdersOrderFidConfirmPayPalOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = putOrdersOrderFidConfirmPayPalOKBodyAO0

	// PutOrdersOrderFidConfirmPayPalOKBodyAO1
	var dataPutOrdersOrderFidConfirmPayPalOKBodyAO1 struct {
		Data *models.OrderConfirmation `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPutOrdersOrderFidConfirmPayPalOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPutOrdersOrderFidConfirmPayPalOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PutOrdersOrderFidConfirmPayPalOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	putOrdersOrderFidConfirmPayPalOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, putOrdersOrderFidConfirmPayPalOKBodyAO0)
	var dataPutOrdersOrderFidConfirmPayPalOKBodyAO1 struct {
		Data *models.OrderConfirmation `json:"data,omitempty"`
	}

	dataPutOrdersOrderFidConfirmPayPalOKBodyAO1.Data = o.Data

	jsonDataPutOrdersOrderFidConfirmPayPalOKBodyAO1, errPutOrdersOrderFidConfirmPayPalOKBodyAO1 := swag.WriteJSON(dataPutOrdersOrderFidConfirmPayPalOKBodyAO1)
	if errPutOrdersOrderFidConfirmPayPalOKBodyAO1 != nil {
		return nil, errPutOrdersOrderFidConfirmPayPalOKBodyAO1
	}
	_parts = append(_parts, jsonDataPutOrdersOrderFidConfirmPayPalOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this put orders order fid confirm pay pal o k body
func (o *PutOrdersOrderFidConfirmPayPalOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PutOrdersOrderFidConfirmPayPalOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putOrdersOrderFidConfirmPayPalOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("putOrdersOrderFidConfirmPayPalOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this put orders order fid confirm pay pal o k body based on the context it is used
func (o *PutOrdersOrderFidConfirmPayPalOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Envelope
	if err := o.Envelope.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutOrdersOrderFidConfirmPayPalOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putOrdersOrderFidConfirmPayPalOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("putOrdersOrderFidConfirmPayPalOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutOrdersOrderFidConfirmPayPalOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutOrdersOrderFidConfirmPayPalOKBody) UnmarshalBinary(b []byte) error {
	var res PutOrdersOrderFidConfirmPayPalOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
