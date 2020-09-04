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

// GetProductsProductFidSkusReader is a Reader for the GetProductsProductFidSkus structure.
type GetProductsProductFidSkusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProductsProductFidSkusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProductsProductFidSkusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetProductsProductFidSkusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProductsProductFidSkusOK creates a GetProductsProductFidSkusOK with default headers values
func NewGetProductsProductFidSkusOK() *GetProductsProductFidSkusOK {
	return &GetProductsProductFidSkusOK{}
}

/*GetProductsProductFidSkusOK handles this case with default header values.

List of product skus
*/
type GetProductsProductFidSkusOK struct {
	Payload *GetProductsProductFidSkusOKBody
}

func (o *GetProductsProductFidSkusOK) Error() string {
	return fmt.Sprintf("[GET /products/{productFid}/skus][%d] getProductsProductFidSkusOK  %+v", 200, o.Payload)
}

func (o *GetProductsProductFidSkusOK) GetPayload() *GetProductsProductFidSkusOKBody {
	return o.Payload
}

func (o *GetProductsProductFidSkusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetProductsProductFidSkusOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProductsProductFidSkusDefault creates a GetProductsProductFidSkusDefault with default headers values
func NewGetProductsProductFidSkusDefault(code int) *GetProductsProductFidSkusDefault {
	return &GetProductsProductFidSkusDefault{
		_statusCode: code,
	}
}

/*GetProductsProductFidSkusDefault handles this case with default header values.

Error
*/
type GetProductsProductFidSkusDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get products product fid skus default response
func (o *GetProductsProductFidSkusDefault) Code() int {
	return o._statusCode
}

func (o *GetProductsProductFidSkusDefault) Error() string {
	return fmt.Sprintf("[GET /products/{productFid}/skus][%d] GetProductsProductFidSkus default  %+v", o._statusCode, o.Payload)
}

func (o *GetProductsProductFidSkusDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetProductsProductFidSkusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetProductsProductFidSkusOKBody get products product fid skus o k body
swagger:model GetProductsProductFidSkusOKBody
*/
type GetProductsProductFidSkusOKBody struct {
	models.Envelope

	// data
	Data *models.ProductSkus `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetProductsProductFidSkusOKBody) UnmarshalJSON(raw []byte) error {
	// GetProductsProductFidSkusOKBodyAO0
	var getProductsProductFidSkusOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getProductsProductFidSkusOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getProductsProductFidSkusOKBodyAO0

	// GetProductsProductFidSkusOKBodyAO1
	var dataGetProductsProductFidSkusOKBodyAO1 struct {
		Data *models.ProductSkus `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetProductsProductFidSkusOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetProductsProductFidSkusOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetProductsProductFidSkusOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getProductsProductFidSkusOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getProductsProductFidSkusOKBodyAO0)

	var dataGetProductsProductFidSkusOKBodyAO1 struct {
		Data *models.ProductSkus `json:"data,omitempty"`
	}

	dataGetProductsProductFidSkusOKBodyAO1.Data = o.Data

	jsonDataGetProductsProductFidSkusOKBodyAO1, errGetProductsProductFidSkusOKBodyAO1 := swag.WriteJSON(dataGetProductsProductFidSkusOKBodyAO1)
	if errGetProductsProductFidSkusOKBodyAO1 != nil {
		return nil, errGetProductsProductFidSkusOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetProductsProductFidSkusOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get products product fid skus o k body
func (o *GetProductsProductFidSkusOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetProductsProductFidSkusOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getProductsProductFidSkusOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetProductsProductFidSkusOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetProductsProductFidSkusOKBody) UnmarshalBinary(b []byte) error {
	var res GetProductsProductFidSkusOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}