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

/*
GetOrdersOrderFidProductsOK describes a response with status code 200, with default header values.

List of order products
*/
type GetOrdersOrderFidProductsOK struct {
	Payload *GetOrdersOrderFidProductsOKBody
}

// IsSuccess returns true when this get orders order fid products o k response has a 2xx status code
func (o *GetOrdersOrderFidProductsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get orders order fid products o k response has a 3xx status code
func (o *GetOrdersOrderFidProductsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orders order fid products o k response has a 4xx status code
func (o *GetOrdersOrderFidProductsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get orders order fid products o k response has a 5xx status code
func (o *GetOrdersOrderFidProductsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get orders order fid products o k response a status code equal to that given
func (o *GetOrdersOrderFidProductsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOrdersOrderFidProductsOK) Error() string {
	return fmt.Sprintf("[GET /orders/{orderFid}/products][%d] getOrdersOrderFidProductsOK  %+v", 200, o.Payload)
}

func (o *GetOrdersOrderFidProductsOK) String() string {
	return fmt.Sprintf("[GET /orders/{orderFid}/products][%d] getOrdersOrderFidProductsOK  %+v", 200, o.Payload)
}

func (o *GetOrdersOrderFidProductsOK) GetPayload() *GetOrdersOrderFidProductsOKBody {
	return o.Payload
}

func (o *GetOrdersOrderFidProductsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOrdersOrderFidProductsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrdersOrderFidProductsDefault creates a GetOrdersOrderFidProductsDefault with default headers values
func NewGetOrdersOrderFidProductsDefault(code int) *GetOrdersOrderFidProductsDefault {
	return &GetOrdersOrderFidProductsDefault{
		_statusCode: code,
	}
}

/*
GetOrdersOrderFidProductsDefault describes a response with status code -1, with default header values.

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

// IsSuccess returns true when this get orders order fid products default response has a 2xx status code
func (o *GetOrdersOrderFidProductsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get orders order fid products default response has a 3xx status code
func (o *GetOrdersOrderFidProductsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get orders order fid products default response has a 4xx status code
func (o *GetOrdersOrderFidProductsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get orders order fid products default response has a 5xx status code
func (o *GetOrdersOrderFidProductsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get orders order fid products default response a status code equal to that given
func (o *GetOrdersOrderFidProductsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetOrdersOrderFidProductsDefault) Error() string {
	return fmt.Sprintf("[GET /orders/{orderFid}/products][%d] GetOrdersOrderFidProducts default  %+v", o._statusCode, o.Payload)
}

func (o *GetOrdersOrderFidProductsDefault) String() string {
	return fmt.Sprintf("[GET /orders/{orderFid}/products][%d] GetOrdersOrderFidProducts default  %+v", o._statusCode, o.Payload)
}

func (o *GetOrdersOrderFidProductsDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetOrdersOrderFidProductsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetOrdersOrderFidProductsOKBody get orders order fid products o k body
swagger:model GetOrdersOrderFidProductsOKBody
*/
type GetOrdersOrderFidProductsOKBody struct {
	models.Envelope

	// data
	Data *models.OrderProducts `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetOrdersOrderFidProductsOKBody) UnmarshalJSON(raw []byte) error {
	// GetOrdersOrderFidProductsOKBodyAO0
	var getOrdersOrderFidProductsOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getOrdersOrderFidProductsOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getOrdersOrderFidProductsOKBodyAO0

	// GetOrdersOrderFidProductsOKBodyAO1
	var dataGetOrdersOrderFidProductsOKBodyAO1 struct {
		Data *models.OrderProducts `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetOrdersOrderFidProductsOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetOrdersOrderFidProductsOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetOrdersOrderFidProductsOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getOrdersOrderFidProductsOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getOrdersOrderFidProductsOKBodyAO0)
	var dataGetOrdersOrderFidProductsOKBodyAO1 struct {
		Data *models.OrderProducts `json:"data,omitempty"`
	}

	dataGetOrdersOrderFidProductsOKBodyAO1.Data = o.Data

	jsonDataGetOrdersOrderFidProductsOKBodyAO1, errGetOrdersOrderFidProductsOKBodyAO1 := swag.WriteJSON(dataGetOrdersOrderFidProductsOKBodyAO1)
	if errGetOrdersOrderFidProductsOKBodyAO1 != nil {
		return nil, errGetOrdersOrderFidProductsOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetOrdersOrderFidProductsOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get orders order fid products o k body
func (o *GetOrdersOrderFidProductsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetOrdersOrderFidProductsOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOrdersOrderFidProductsOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getOrdersOrderFidProductsOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get orders order fid products o k body based on the context it is used
func (o *GetOrdersOrderFidProductsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetOrdersOrderFidProductsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOrdersOrderFidProductsOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getOrdersOrderFidProductsOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrdersOrderFidProductsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrdersOrderFidProductsOKBody) UnmarshalBinary(b []byte) error {
	var res GetOrdersOrderFidProductsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
