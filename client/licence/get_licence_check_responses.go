// Code generated by go-swagger; DO NOT EDIT.

package licence

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

// GetLicenceCheckReader is a Reader for the GetLicenceCheck structure.
type GetLicenceCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLicenceCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLicenceCheckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetLicenceCheckDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLicenceCheckOK creates a GetLicenceCheckOK with default headers values
func NewGetLicenceCheckOK() *GetLicenceCheckOK {
	return &GetLicenceCheckOK{}
}

/*GetLicenceCheckOK handles this case with default header values.

Licence Information
*/
type GetLicenceCheckOK struct {
	Payload *GetLicenceCheckOKBody
}

func (o *GetLicenceCheckOK) Error() string {
	return fmt.Sprintf("[GET /licence/check][%d] getLicenceCheckOK  %+v", 200, o.Payload)
}

func (o *GetLicenceCheckOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetLicenceCheckOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenceCheckDefault creates a GetLicenceCheckDefault with default headers values
func NewGetLicenceCheckDefault(code int) *GetLicenceCheckDefault {
	return &GetLicenceCheckDefault{
		_statusCode: code,
	}
}

/*GetLicenceCheckDefault handles this case with default header values.

Error
*/
type GetLicenceCheckDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get licence check default response
func (o *GetLicenceCheckDefault) Code() int {
	return o._statusCode
}

func (o *GetLicenceCheckDefault) Error() string {
	return fmt.Sprintf("[GET /licence/check][%d] GetLicenceCheck default  %+v", o._statusCode, o.Payload)
}

func (o *GetLicenceCheckDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetLicenceCheckOKBody get licence check o k body
swagger:model GetLicenceCheckOKBody
*/
type GetLicenceCheckOKBody struct {
	models.Envelope

	// data
	Data *models.Licence `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetLicenceCheckOKBody) UnmarshalJSON(raw []byte) error {
	// GetLicenceCheckOKBodyAO0
	var getLicenceCheckOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getLicenceCheckOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getLicenceCheckOKBodyAO0

	// GetLicenceCheckOKBodyAO1
	var dataGetLicenceCheckOKBodyAO1 struct {
		Data *models.Licence `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetLicenceCheckOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetLicenceCheckOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetLicenceCheckOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getLicenceCheckOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getLicenceCheckOKBodyAO0)

	var dataGetLicenceCheckOKBodyAO1 struct {
		Data *models.Licence `json:"data,omitempty"`
	}

	dataGetLicenceCheckOKBodyAO1.Data = o.Data

	jsonDataGetLicenceCheckOKBodyAO1, errGetLicenceCheckOKBodyAO1 := swag.WriteJSON(dataGetLicenceCheckOKBodyAO1)
	if errGetLicenceCheckOKBodyAO1 != nil {
		return nil, errGetLicenceCheckOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetLicenceCheckOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get licence check o k body
func (o *GetLicenceCheckOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetLicenceCheckOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLicenceCheckOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLicenceCheckOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLicenceCheckOKBody) UnmarshalBinary(b []byte) error {
	var res GetLicenceCheckOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
