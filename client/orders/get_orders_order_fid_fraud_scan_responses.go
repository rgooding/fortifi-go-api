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

// GetOrdersOrderFidFraudScanReader is a Reader for the GetOrdersOrderFidFraudScan structure.
type GetOrdersOrderFidFraudScanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrdersOrderFidFraudScanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrdersOrderFidFraudScanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetOrdersOrderFidFraudScanDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOrdersOrderFidFraudScanOK creates a GetOrdersOrderFidFraudScanOK with default headers values
func NewGetOrdersOrderFidFraudScanOK() *GetOrdersOrderFidFraudScanOK {
	return &GetOrdersOrderFidFraudScanOK{}
}

/* GetOrdersOrderFidFraudScanOK describes a response with status code 200, with default header values.

Fraud scan result
*/
type GetOrdersOrderFidFraudScanOK struct {
	Payload *GetOrdersOrderFidFraudScanOKBody
}

func (o *GetOrdersOrderFidFraudScanOK) Error() string {
	return fmt.Sprintf("[GET /orders/{orderFid}/fraudScan][%d] getOrdersOrderFidFraudScanOK  %+v", 200, o.Payload)
}
func (o *GetOrdersOrderFidFraudScanOK) GetPayload() *GetOrdersOrderFidFraudScanOKBody {
	return o.Payload
}

func (o *GetOrdersOrderFidFraudScanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOrdersOrderFidFraudScanOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrdersOrderFidFraudScanDefault creates a GetOrdersOrderFidFraudScanDefault with default headers values
func NewGetOrdersOrderFidFraudScanDefault(code int) *GetOrdersOrderFidFraudScanDefault {
	return &GetOrdersOrderFidFraudScanDefault{
		_statusCode: code,
	}
}

/* GetOrdersOrderFidFraudScanDefault describes a response with status code -1, with default header values.

Error
*/
type GetOrdersOrderFidFraudScanDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get orders order fid fraud scan default response
func (o *GetOrdersOrderFidFraudScanDefault) Code() int {
	return o._statusCode
}

func (o *GetOrdersOrderFidFraudScanDefault) Error() string {
	return fmt.Sprintf("[GET /orders/{orderFid}/fraudScan][%d] GetOrdersOrderFidFraudScan default  %+v", o._statusCode, o.Payload)
}
func (o *GetOrdersOrderFidFraudScanDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetOrdersOrderFidFraudScanDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetOrdersOrderFidFraudScanOKBody get orders order fid fraud scan o k body
swagger:model GetOrdersOrderFidFraudScanOKBody
*/
type GetOrdersOrderFidFraudScanOKBody struct {
	models.Envelope

	// data
	Data *models.FraudScan `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetOrdersOrderFidFraudScanOKBody) UnmarshalJSON(raw []byte) error {
	// GetOrdersOrderFidFraudScanOKBodyAO0
	var getOrdersOrderFidFraudScanOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getOrdersOrderFidFraudScanOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getOrdersOrderFidFraudScanOKBodyAO0

	// GetOrdersOrderFidFraudScanOKBodyAO1
	var dataGetOrdersOrderFidFraudScanOKBodyAO1 struct {
		Data *models.FraudScan `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetOrdersOrderFidFraudScanOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetOrdersOrderFidFraudScanOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetOrdersOrderFidFraudScanOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getOrdersOrderFidFraudScanOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getOrdersOrderFidFraudScanOKBodyAO0)
	var dataGetOrdersOrderFidFraudScanOKBodyAO1 struct {
		Data *models.FraudScan `json:"data,omitempty"`
	}

	dataGetOrdersOrderFidFraudScanOKBodyAO1.Data = o.Data

	jsonDataGetOrdersOrderFidFraudScanOKBodyAO1, errGetOrdersOrderFidFraudScanOKBodyAO1 := swag.WriteJSON(dataGetOrdersOrderFidFraudScanOKBodyAO1)
	if errGetOrdersOrderFidFraudScanOKBodyAO1 != nil {
		return nil, errGetOrdersOrderFidFraudScanOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetOrdersOrderFidFraudScanOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get orders order fid fraud scan o k body
func (o *GetOrdersOrderFidFraudScanOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetOrdersOrderFidFraudScanOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOrdersOrderFidFraudScanOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get orders order fid fraud scan o k body based on the context it is used
func (o *GetOrdersOrderFidFraudScanOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetOrdersOrderFidFraudScanOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOrdersOrderFidFraudScanOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrdersOrderFidFraudScanOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrdersOrderFidFraudScanOKBody) UnmarshalBinary(b []byte) error {
	var res GetOrdersOrderFidFraudScanOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
