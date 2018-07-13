// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// GetBrandsReader is a Reader for the GetBrands structure.
type GetBrandsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBrandsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBrandsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetBrandsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBrandsOK creates a GetBrandsOK with default headers values
func NewGetBrandsOK() *GetBrandsOK {
	return &GetBrandsOK{}
}

/*GetBrandsOK handles this case with default header values.

Brand Information
*/
type GetBrandsOK struct {
	Payload *GetBrandsOKBody
}

func (o *GetBrandsOK) Error() string {
	return fmt.Sprintf("[GET /brands][%d] getBrandsOK  %+v", 200, o.Payload)
}

func (o *GetBrandsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetBrandsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBrandsDefault creates a GetBrandsDefault with default headers values
func NewGetBrandsDefault(code int) *GetBrandsDefault {
	return &GetBrandsDefault{
		_statusCode: code,
	}
}

/*GetBrandsDefault handles this case with default header values.

Error
*/
type GetBrandsDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get brands default response
func (o *GetBrandsDefault) Code() int {
	return o._statusCode
}

func (o *GetBrandsDefault) Error() string {
	return fmt.Sprintf("[GET /brands][%d] getBrands default  %+v", o._statusCode, o.Payload)
}

func (o *GetBrandsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetBrandsOKBody get brands o k body
swagger:model GetBrandsOKBody
*/
type GetBrandsOKBody struct {
	models.Envelope

	// data
	Data *models.Brands `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetBrandsOKBody) UnmarshalJSON(raw []byte) error {
	// GetBrandsOKBodyAO0
	var getBrandsOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getBrandsOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getBrandsOKBodyAO0

	// GetBrandsOKBodyAO1
	var dataGetBrandsOKBodyAO1 struct {
		Data *models.Brands `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetBrandsOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetBrandsOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetBrandsOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getBrandsOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getBrandsOKBodyAO0)

	var dataGetBrandsOKBodyAO1 struct {
		Data *models.Brands `json:"data,omitempty"`
	}

	dataGetBrandsOKBodyAO1.Data = o.Data

	jsonDataGetBrandsOKBodyAO1, errGetBrandsOKBodyAO1 := swag.WriteJSON(dataGetBrandsOKBodyAO1)
	if errGetBrandsOKBodyAO1 != nil {
		return nil, errGetBrandsOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetBrandsOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get brands o k body
func (o *GetBrandsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetBrandsOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getBrandsOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetBrandsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetBrandsOKBody) UnmarshalBinary(b []byte) error {
	var res GetBrandsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
