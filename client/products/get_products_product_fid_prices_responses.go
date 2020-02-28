// Code generated by go-swagger; DO NOT EDIT.

package products

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

// GetProductsProductFidPricesReader is a Reader for the GetProductsProductFidPrices structure.
type GetProductsProductFidPricesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProductsProductFidPricesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProductsProductFidPricesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetProductsProductFidPricesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProductsProductFidPricesOK creates a GetProductsProductFidPricesOK with default headers values
func NewGetProductsProductFidPricesOK() *GetProductsProductFidPricesOK {
	return &GetProductsProductFidPricesOK{}
}

/*GetProductsProductFidPricesOK handles this case with default header values.

List of product prices
*/
type GetProductsProductFidPricesOK struct {
	Payload *GetProductsProductFidPricesOKBody
}

func (o *GetProductsProductFidPricesOK) Error() string {
	return fmt.Sprintf("[GET /products/{productFid}/prices][%d] getProductsProductFidPricesOK  %+v", 200, o.Payload)
}

func (o *GetProductsProductFidPricesOK) GetPayload() *GetProductsProductFidPricesOKBody {
	return o.Payload
}

func (o *GetProductsProductFidPricesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetProductsProductFidPricesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProductsProductFidPricesDefault creates a GetProductsProductFidPricesDefault with default headers values
func NewGetProductsProductFidPricesDefault(code int) *GetProductsProductFidPricesDefault {
	return &GetProductsProductFidPricesDefault{
		_statusCode: code,
	}
}

/*GetProductsProductFidPricesDefault handles this case with default header values.

Error
*/
type GetProductsProductFidPricesDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get products product fid prices default response
func (o *GetProductsProductFidPricesDefault) Code() int {
	return o._statusCode
}

func (o *GetProductsProductFidPricesDefault) Error() string {
	return fmt.Sprintf("[GET /products/{productFid}/prices][%d] GetProductsProductFidPrices default  %+v", o._statusCode, o.Payload)
}

func (o *GetProductsProductFidPricesDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetProductsProductFidPricesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetProductsProductFidPricesOKBody get products product fid prices o k body
swagger:model GetProductsProductFidPricesOKBody
*/
type GetProductsProductFidPricesOKBody struct {
	models.Envelope

	// data
	Data *models.ProductPrices `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetProductsProductFidPricesOKBody) UnmarshalJSON(raw []byte) error {
	// GetProductsProductFidPricesOKBodyAO0
	var getProductsProductFidPricesOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getProductsProductFidPricesOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getProductsProductFidPricesOKBodyAO0

	// GetProductsProductFidPricesOKBodyAO1
	var dataGetProductsProductFidPricesOKBodyAO1 struct {
		Data *models.ProductPrices `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetProductsProductFidPricesOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetProductsProductFidPricesOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetProductsProductFidPricesOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getProductsProductFidPricesOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getProductsProductFidPricesOKBodyAO0)

	var dataGetProductsProductFidPricesOKBodyAO1 struct {
		Data *models.ProductPrices `json:"data,omitempty"`
	}

	dataGetProductsProductFidPricesOKBodyAO1.Data = o.Data

	jsonDataGetProductsProductFidPricesOKBodyAO1, errGetProductsProductFidPricesOKBodyAO1 := swag.WriteJSON(dataGetProductsProductFidPricesOKBodyAO1)
	if errGetProductsProductFidPricesOKBodyAO1 != nil {
		return nil, errGetProductsProductFidPricesOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetProductsProductFidPricesOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get products product fid prices o k body
func (o *GetProductsProductFidPricesOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetProductsProductFidPricesOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getProductsProductFidPricesOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetProductsProductFidPricesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetProductsProductFidPricesOKBody) UnmarshalBinary(b []byte) error {
	var res GetProductsProductFidPricesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
