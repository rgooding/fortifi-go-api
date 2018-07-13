// Code generated by go-swagger; DO NOT EDIT.

package orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/fortifi/go-api/models"
)

// GetOrdersOrderFidReader is a Reader for the GetOrdersOrderFid structure.
type GetOrdersOrderFidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrdersOrderFidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOrdersOrderFidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetOrdersOrderFidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOrdersOrderFidOK creates a GetOrdersOrderFidOK with default headers values
func NewGetOrdersOrderFidOK() *GetOrdersOrderFidOK {
	return &GetOrdersOrderFidOK{}
}

/*GetOrdersOrderFidOK handles this case with default header values.

Order retrieved
*/
type GetOrdersOrderFidOK struct {
	Payload *GetOrdersOrderFidOKBody
}

func (o *GetOrdersOrderFidOK) Error() string {
	return fmt.Sprintf("[GET /orders/{orderFid}][%d] getOrdersOrderFidOK  %+v", 200, o.Payload)
}

func (o *GetOrdersOrderFidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOrdersOrderFidOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrdersOrderFidDefault creates a GetOrdersOrderFidDefault with default headers values
func NewGetOrdersOrderFidDefault(code int) *GetOrdersOrderFidDefault {
	return &GetOrdersOrderFidDefault{
		_statusCode: code,
	}
}

/*GetOrdersOrderFidDefault handles this case with default header values.

Error
*/
type GetOrdersOrderFidDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get orders order fid default response
func (o *GetOrdersOrderFidDefault) Code() int {
	return o._statusCode
}

func (o *GetOrdersOrderFidDefault) Error() string {
	return fmt.Sprintf("[GET /orders/{orderFid}][%d] GetOrdersOrderFid default  %+v", o._statusCode, o.Payload)
}

func (o *GetOrdersOrderFidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetOrdersOrderFidOKBody get orders order fid o k body
swagger:model GetOrdersOrderFidOKBody
*/
type GetOrdersOrderFidOKBody struct {
	models.Envelope

	// data
	Data *models.Order `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetOrdersOrderFidOKBody) UnmarshalJSON(raw []byte) error {
	// GetOrdersOrderFidOKBodyAO0
	var getOrdersOrderFidOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getOrdersOrderFidOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getOrdersOrderFidOKBodyAO0

	// GetOrdersOrderFidOKBodyAO1
	var dataGetOrdersOrderFidOKBodyAO1 struct {
		Data *models.Order `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetOrdersOrderFidOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetOrdersOrderFidOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetOrdersOrderFidOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getOrdersOrderFidOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getOrdersOrderFidOKBodyAO0)

	var dataGetOrdersOrderFidOKBodyAO1 struct {
		Data *models.Order `json:"data,omitempty"`
	}

	dataGetOrdersOrderFidOKBodyAO1.Data = o.Data

	jsonDataGetOrdersOrderFidOKBodyAO1, errGetOrdersOrderFidOKBodyAO1 := swag.WriteJSON(dataGetOrdersOrderFidOKBodyAO1)
	if errGetOrdersOrderFidOKBodyAO1 != nil {
		return nil, errGetOrdersOrderFidOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetOrdersOrderFidOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get orders order fid o k body
func (o *GetOrdersOrderFidOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetOrdersOrderFidOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOrdersOrderFidOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrdersOrderFidOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrdersOrderFidOKBody) UnmarshalBinary(b []byte) error {
	var res GetOrdersOrderFidOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
