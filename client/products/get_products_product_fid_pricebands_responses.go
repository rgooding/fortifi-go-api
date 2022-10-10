// Code generated by go-swagger; DO NOT EDIT.

package products

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

// GetProductsProductFidPricebandsReader is a Reader for the GetProductsProductFidPricebands structure.
type GetProductsProductFidPricebandsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProductsProductFidPricebandsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProductsProductFidPricebandsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetProductsProductFidPricebandsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProductsProductFidPricebandsOK creates a GetProductsProductFidPricebandsOK with default headers values
func NewGetProductsProductFidPricebandsOK() *GetProductsProductFidPricebandsOK {
	return &GetProductsProductFidPricebandsOK{}
}

/*
GetProductsProductFidPricebandsOK describes a response with status code 200, with default header values.

List of product price bands
*/
type GetProductsProductFidPricebandsOK struct {
	Payload *GetProductsProductFidPricebandsOKBody
}

// IsSuccess returns true when this get products product fid pricebands o k response has a 2xx status code
func (o *GetProductsProductFidPricebandsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get products product fid pricebands o k response has a 3xx status code
func (o *GetProductsProductFidPricebandsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get products product fid pricebands o k response has a 4xx status code
func (o *GetProductsProductFidPricebandsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get products product fid pricebands o k response has a 5xx status code
func (o *GetProductsProductFidPricebandsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get products product fid pricebands o k response a status code equal to that given
func (o *GetProductsProductFidPricebandsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetProductsProductFidPricebandsOK) Error() string {
	return fmt.Sprintf("[GET /products/{productFid}/pricebands][%d] getProductsProductFidPricebandsOK  %+v", 200, o.Payload)
}

func (o *GetProductsProductFidPricebandsOK) String() string {
	return fmt.Sprintf("[GET /products/{productFid}/pricebands][%d] getProductsProductFidPricebandsOK  %+v", 200, o.Payload)
}

func (o *GetProductsProductFidPricebandsOK) GetPayload() *GetProductsProductFidPricebandsOKBody {
	return o.Payload
}

func (o *GetProductsProductFidPricebandsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetProductsProductFidPricebandsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProductsProductFidPricebandsDefault creates a GetProductsProductFidPricebandsDefault with default headers values
func NewGetProductsProductFidPricebandsDefault(code int) *GetProductsProductFidPricebandsDefault {
	return &GetProductsProductFidPricebandsDefault{
		_statusCode: code,
	}
}

/*
GetProductsProductFidPricebandsDefault describes a response with status code -1, with default header values.

Error
*/
type GetProductsProductFidPricebandsDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get products product fid pricebands default response
func (o *GetProductsProductFidPricebandsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get products product fid pricebands default response has a 2xx status code
func (o *GetProductsProductFidPricebandsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get products product fid pricebands default response has a 3xx status code
func (o *GetProductsProductFidPricebandsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get products product fid pricebands default response has a 4xx status code
func (o *GetProductsProductFidPricebandsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get products product fid pricebands default response has a 5xx status code
func (o *GetProductsProductFidPricebandsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get products product fid pricebands default response a status code equal to that given
func (o *GetProductsProductFidPricebandsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetProductsProductFidPricebandsDefault) Error() string {
	return fmt.Sprintf("[GET /products/{productFid}/pricebands][%d] GetProductsProductFidPricebands default  %+v", o._statusCode, o.Payload)
}

func (o *GetProductsProductFidPricebandsDefault) String() string {
	return fmt.Sprintf("[GET /products/{productFid}/pricebands][%d] GetProductsProductFidPricebands default  %+v", o._statusCode, o.Payload)
}

func (o *GetProductsProductFidPricebandsDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetProductsProductFidPricebandsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetProductsProductFidPricebandsOKBody get products product fid pricebands o k body
swagger:model GetProductsProductFidPricebandsOKBody
*/
type GetProductsProductFidPricebandsOKBody struct {
	models.Envelope

	// data
	Data *models.ProductPriceBands `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetProductsProductFidPricebandsOKBody) UnmarshalJSON(raw []byte) error {
	// GetProductsProductFidPricebandsOKBodyAO0
	var getProductsProductFidPricebandsOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getProductsProductFidPricebandsOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getProductsProductFidPricebandsOKBodyAO0

	// GetProductsProductFidPricebandsOKBodyAO1
	var dataGetProductsProductFidPricebandsOKBodyAO1 struct {
		Data *models.ProductPriceBands `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetProductsProductFidPricebandsOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetProductsProductFidPricebandsOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetProductsProductFidPricebandsOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getProductsProductFidPricebandsOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getProductsProductFidPricebandsOKBodyAO0)
	var dataGetProductsProductFidPricebandsOKBodyAO1 struct {
		Data *models.ProductPriceBands `json:"data,omitempty"`
	}

	dataGetProductsProductFidPricebandsOKBodyAO1.Data = o.Data

	jsonDataGetProductsProductFidPricebandsOKBodyAO1, errGetProductsProductFidPricebandsOKBodyAO1 := swag.WriteJSON(dataGetProductsProductFidPricebandsOKBodyAO1)
	if errGetProductsProductFidPricebandsOKBodyAO1 != nil {
		return nil, errGetProductsProductFidPricebandsOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetProductsProductFidPricebandsOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get products product fid pricebands o k body
func (o *GetProductsProductFidPricebandsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetProductsProductFidPricebandsOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getProductsProductFidPricebandsOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getProductsProductFidPricebandsOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get products product fid pricebands o k body based on the context it is used
func (o *GetProductsProductFidPricebandsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetProductsProductFidPricebandsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getProductsProductFidPricebandsOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getProductsProductFidPricebandsOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetProductsProductFidPricebandsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetProductsProductFidPricebandsOKBody) UnmarshalBinary(b []byte) error {
	var res GetProductsProductFidPricebandsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}